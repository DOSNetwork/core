package dkg

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
)

type PDKGInterface interface {
	GetGroupPublicPoly(groupId string) *share.PubPoly
	GetShareSecurity(groupId string) *share.PriShare
	GetGroupIDs(groupId string) [][]byte
	GetGroupNumber() int
	Grouping(ctx context.Context, groupId string, Participants [][]byte) (chan [5]*big.Int, chan error, error)
	GroupDissolve(groupId string)
}

type pdkg struct {
	p         p2p.P2PInterface
	suite     suites.Suite
	bufToNode chan interface{}
	register  chan *group
	groups    sync.Map
	logger    log.Logger
}

type group struct {
	participants [][]byte
	secShare     *DistKeyShare
	pubPoly      *share.PubPoly
}

type request struct {
	ctx        context.Context
	reqType    int
	sessionID  string
	numOfResps int
	reply      chan []interface{}
}

func NewPDKG(p p2p.P2PInterface, suite suites.Suite) PDKGInterface {
	d := &pdkg{
		p:         p,
		bufToNode: make(chan interface{}, 1),
		register:  make(chan *group),
		suite:     suite,
		logger:    log.New("module", "dkg"),
	}
	d.listen()
	return d
}

func handlePeerMsg(sessionMap map[string][]interface{}, sessionReq map[string]request, p p2p.P2PInterface, sessionID string, content interface{}) {
	sessionMap[sessionID] = append(sessionMap[sessionID], content)
	if sessionMap[sessionID] != nil {
		if len(sessionMap[sessionID]) == sessionReq[sessionID].numOfResps {
			go func(req request, m []interface{}) {
				select {
				case <-req.ctx.Done():
				case req.reply <- m:
				}
				close(req.reply)
			}(sessionReq[sessionID], sessionMap[sessionID])

			delete(sessionMap, sessionID)
			delete(sessionReq, sessionID)
		}
	}
}

func handleRequest(sessionMap map[string][]interface{}, sessionReq map[string]request, req request) {
	sessionReq[req.sessionID] = req
	if len(sessionMap[req.sessionID]) == req.numOfResps {
		go func(req request, m []interface{}) {
			select {
			case <-req.ctx.Done():
			case req.reply <- m:
			}
			close(req.reply)
		}(sessionReq[req.sessionID], sessionMap[req.sessionID])

		delete(sessionMap, req.sessionID)
		delete(sessionReq, req.sessionID)
	}
}
func (d *pdkg) listen() {
	peersToBuf, _ := d.p.SubscribeEvent(50, PublicKey{}, Deal{}, Responses{})
	go func() {
		sessionPubKeys := make(map[string][]interface{})
		sessionDeals := make(map[string][]interface{})
		sessionResps := make(map[string][]interface{})
		sessionReqPubs := map[string]request{}
		sessionReqDeals := map[string]request{}
		sessionReResps := map[string]request{}
		for {
			select {
			case msg, ok := <-peersToBuf:
				if !ok {
					return
				}
				switch content := msg.Msg.Message.(type) {
				case *PublicKey:
					d.p.Reply(msg.Sender, msg.RequestNonce, &PublicKey{})
					handlePeerMsg(sessionPubKeys, sessionReqPubs, d.p, content.SessionId, content)
				case *Deal:
					d.p.Reply(msg.Sender, msg.RequestNonce, &Deal{})
					handlePeerMsg(sessionDeals, sessionReqDeals, d.p, content.SessionId, content)
				case *Responses:
					d.p.Reply(msg.Sender, msg.RequestNonce, &Responses{})
					resps := content.Response
					for _, resp := range resps {
						handlePeerMsg(sessionResps, sessionReResps, d.p, content.SessionId, resp)
					}
				}

			case req, ok := <-d.bufToNode:
				if !ok {
					return
				}
				if r, ok := req.(request); ok {
					switch r.reqType {
					case 0:
						handleRequest(sessionPubKeys, sessionReqPubs, r)
					case 1:
						handleRequest(sessionDeals, sessionReqDeals, r)
					case 2:
						handleRequest(sessionResps, sessionReResps, r)
					}
				} else {
					fmt.Println("handleRequest cast error")

				}
			}
		}
	}()
}

func (d *pdkg) Grouping(ctx context.Context, sessionID string, groupIds [][]byte) (chan [5]*big.Int, chan error, error) {
	group := &group{participants: groupIds}
	var errcList []chan error
	if _, loaded := d.groups.LoadOrStore(sessionID, group); loaded {
		return nil, nil, errors.New("dkg: duplicate share public key")
	}

	//Check if all members are reachable
	connc, errc := d.p.ConnectToAll(ctx, groupIds)
	errcList = append(errcList, errc)
	//exchange pub key
	selfPubc, secrc, errc := genPub(ctx, d.logger, connc, d.suite, d.p.GetID(), groupIds, sessionID)
	errcList = append(errcList, errc)
	selfPubcs := fanOut(ctx, selfPubc, 2)
	errcList = append(errcList, sendToMembers(ctx, d.logger, selfPubcs[0], d.p, groupIds))
	peerPubc := askMembers(ctx, d.logger, d.bufToNode, len(groupIds)-1, 0, sessionID)
	errcList = append(errcList, errc)
	partPubsc, errc := exchangePub(ctx, d.logger, selfPubcs[1], peerPubc, d.p, groupIds)
	errcList = append(errcList, errc)

	//generate a dkg
	dkgcStep1, errc := genDistKeyGenerator(ctx, d.logger, secrc, partPubsc, len(groupIds), d.suite)
	errcList = append(errcList, errc)

	//generate deals for other member and process deals from other member
	dkgcStep2, errc := genDealsAndSend(ctx, d.logger, dkgcStep1, d.p, groupIds, sessionID)
	errcList = append(errcList, errc)
	dkgcStep3, respsc, errc := getAndProcessDeals(ctx, d.logger, dkgcStep2, askMembers(ctx, d.logger, d.bufToNode, len(groupIds)-1, 1, sessionID), sessionID)
	errcList = append(errcList, errc)
	errcList = append(errcList, sendToMembers(ctx, d.logger, respsc, d.p, groupIds))

	//process responces to certify dkg and generate a group sec and pub key
	cetifiedDkgc, errc := getAndProcessResponses(ctx, d.logger, dkgcStep3, askMembers(ctx, d.logger, d.bufToNode, (len(groupIds)-1)*(len(groupIds)-1), 2, sessionID))
	outc, errc := genGroup(ctx, d.logger, group, d.suite, cetifiedDkgc, sessionID)
	errcList = append(errcList, errc)
	errc = mergeErrors(ctx, d.logger, errcList...)
	return outc, errc, nil
}

func genPub(ctx context.Context, logger log.Logger, conn chan bool, suite suites.Suite, id []byte, groupIds [][]byte, sessionID string) (out chan interface{}, secrc chan kyber.Scalar, errc chan error) {
	out = make(chan interface{})
	secrc = make(chan kyber.Scalar)
	errc = make(chan error)
	go func() {
		//defer fmt.Println("1 ) genPub")
		defer close(out)
		defer close(errc)
		select {
		case c, ok := <-conn:
			if ok && c == true {
				defer logger.TimeTrack(time.Now(), "genPub", nil)
				//Index pub key
				index := -1
				for i, groupId := range groupIds {
					if r := bytes.Compare(id, groupId); r == 0 {
						index = i
						break
					}
				}
				if index == -1 {
					reportErr(ctx, errc, errors.New("Can't find id in group IDs"))
				}
				//Generate secret and public key
				sec := suite.Scalar().Pick(suite.RandomStream())
				select {
				case secrc <- sec:
				case <-ctx.Done():
				}
				pub := suite.Point().Mul(sec, nil)
				if bin, err := pub.MarshalBinary(); err != nil {
					reportErr(ctx, errc, err)
				} else {
					pubkey := &PublicKey{SessionId: sessionID, Index: uint32(index), Publickey: &vss.PublicKey{Binary: bin}}
					select {
					case out <- pubkey:
					case <-ctx.Done():
					}
					return
				}
			}
		case <-ctx.Done():
			return
		}
	}()
	return
}

func exchangePub(ctx context.Context, logger log.Logger, selfPubc chan interface{}, peerPubc chan []interface{}, p p2p.P2PInterface, groupIds [][]byte) (out chan []*PublicKey, errc chan error) {
	out = make(chan []*PublicKey)
	errc = make(chan error)
	go func() {
		defer logger.TimeTrack(time.Now(), "exchangePub", nil)
		defer close(out)
		defer close(errc)
		var partPubs []*PublicKey
		for {
			select {
			case <-ctx.Done():
			case resp, ok := <-selfPubc:
				if ok {
					if pubkey, ok := resp.(*PublicKey); ok {
						partPubs = append(partPubs, pubkey)
					}
				}
			case resps, ok := <-peerPubc:
				if ok {
					for _, resp := range resps {
						if pubkey, ok := resp.(*PublicKey); ok {
							partPubs = append(partPubs, pubkey)
						}
					}
				}
			}
			if len(partPubs) == len(groupIds) {
				select {
				case <-ctx.Done():
				case out <- partPubs:
				}
				return
			}
		}
	}()
	return
}

func sendToMembers(ctx context.Context, logger log.Logger, msgc chan interface{}, p p2p.P2PInterface, groupIds [][]byte) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case <-ctx.Done():
		case msg, ok := <-msgc:
			if ok {
				defer logger.TimeTrack(time.Now(), "sendToMembers", nil)

				if m, ok := msg.(proto.Message); ok {
					var wg sync.WaitGroup
					wg.Add(len(groupIds) - 1)
					for i, id := range groupIds {
						if r := bytes.Compare(p.GetID(), id); r != 0 {
							go func(i int, id []byte) {
								defer wg.Done()
								for {
									//retry until success or ctx.Done
									if _, err := p.Request(id, m); err != nil {
										reportErr(ctx, errc, err)
									} else {
										return
									}
								}
							}(i, id)
						}
					}
					wg.Wait()
				}
			}
		}
	}()
	return
}

func askMembers(ctx context.Context, logger log.Logger, bufToNode chan interface{}, numOfResp, reqTpe int, sessionID string) (out chan []interface{}) {
	out = make(chan []interface{})
	go func() {
		defer logger.TimeTrack(time.Now(), "askMembers", nil)

		req := request{ctx: ctx, reqType: reqTpe, sessionID: sessionID, numOfResps: numOfResp, reply: out}
		select {
		case <-ctx.Done():
		case bufToNode <- req:
		}
	}()
	return
}

func genDistKeyGenerator(ctx context.Context, logger log.Logger, secrc chan kyber.Scalar, partPubs chan []*PublicKey, numOfPubkeys int, suite suites.Suite) (out chan *DistKeyGenerator, errc chan error) {
	out = make(chan *DistKeyGenerator)
	errc = make(chan error)
	go func() {
		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case sec, ok := <-secrc:
			if ok {
				defer logger.TimeTrack(time.Now(), "genDistKeyGenerator", nil)

				select {
				case <-ctx.Done():
				case pubs, ok := <-partPubs:
					if ok {
						pubPoints := make([]kyber.Point, numOfPubkeys)
						for _, pubkey := range pubs {
							pubPoints[pubkey.Index] = suite.Point()
							if err := pubPoints[pubkey.Index].UnmarshalBinary(pubkey.Publickey.Binary); err != nil {
								reportErr(ctx, errc, err)
								return
							}
						}
						dkg, err := NewDistKeyGenerator(suite, sec, pubPoints, numOfPubkeys/2+1)
						if err != nil {
							reportErr(ctx, errc, err)
						} else {
							select {
							case <-ctx.Done():
							case out <- dkg:
							}
							return
						}
					}
				}
			}

		}
	}()
	return
}

func genDealsAndSend(ctx context.Context, logger log.Logger, dkgc chan *DistKeyGenerator, p p2p.P2PInterface, groupIds [][]byte, sessionID string) (out chan *DistKeyGenerator, errc chan error) {
	out = make(chan *DistKeyGenerator)
	errc = make(chan error)
	go func() {
		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case dkg, ok := <-dkgc:
			if ok {
				defer logger.TimeTrack(time.Now(), "genDealsAndSend", nil)

				if deals, err := dkg.Deals(); err == nil {
					var wg sync.WaitGroup
					wg.Add(len(groupIds) - 1)
					for i, d := range deals {
						d.SessionId = sessionID
						func(id []byte, d *Deal) {
							defer wg.Done()
							if _, err := p.Request(id, d); err != nil {
								reportErr(ctx, errc, err)
								return
							}
						}(groupIds[i], d)
					}
					wg.Wait()
					select {
					case <-ctx.Done():
					case out <- dkg:
					}
				} else {
					reportErr(ctx, errc, err)
				}
			}
		}
	}()
	return
}
func getAndProcessDeals(ctx context.Context, logger log.Logger, dkgc chan *DistKeyGenerator, dealsc chan []interface{}, sessionID string) (dkgOut chan *DistKeyGenerator, out chan interface{}, errc chan error) {
	dkgOut = make(chan *DistKeyGenerator)
	out = make(chan interface{})
	errc = make(chan error)
	go func() {
		defer close(dkgOut)
		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case dkg, ok := <-dkgc:
			if ok {
				defer logger.TimeTrack(time.Now(), "getAndProcessDeals", nil)
				select {
				case <-ctx.Done():
				case deals, ok := <-dealsc:
					if ok {
						var resps []*Response
						for _, d := range deals {
							if deal, ok := d.(*Deal); ok {
								if resp, err := dkg.ProcessDeal(deal); err == nil {
									resp.SessionId = sessionID
									if vss.StatusApproval == resp.Response.Status {
										resps = append(resps, resp)
									} else {
										reportErr(ctx, errc, errors.New("resp StatusNotApproval"))
									}
								} else {
									reportErr(ctx, errc, err)
								}
							}
						}

						select {
						case out <- &Responses{SessionId: sessionID, Response: resps}:
						case <-ctx.Done():
						}
						select {
						case dkgOut <- dkg:
						case <-ctx.Done():
						}
					}
				}
			}
		}
	}()
	return
}

func getAndProcessResponses(ctx context.Context, logger log.Logger, dkgc chan *DistKeyGenerator, respsc chan []interface{}) (out chan *DistKeyGenerator, errc chan error) {
	out = make(chan *DistKeyGenerator)
	errc = make(chan error)
	go func() {
		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case dkg, ok := <-dkgc:
			if ok {
				defer logger.TimeTrack(time.Now(), "getAndProcessResponses", nil)
				select {
				case <-ctx.Done():
				case resps, ok := <-respsc:
					if ok {
						for _, r := range resps {
							if resp, ok := r.(*Response); ok {
								if _, err := dkg.ProcessResponse(resp); err != nil {
									reportErr(ctx, errc, err)
								}
							} else {
								reportErr(ctx, errc, errors.New("Response cast error"))
							}
						}

						select {
						case <-ctx.Done():
						case out <- dkg:
						}
					}

				}
			}
		}
	}()
	return
}
func genGroup(ctx context.Context, logger log.Logger, group *group, suite suites.Suite, dkgc <-chan *DistKeyGenerator, sessionID string) (out chan [5]*big.Int, errc chan error) {
	out = make(chan [5]*big.Int)
	errc = make(chan error)
	go func() {
		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case dkg, ok := <-dkgc:
			if ok {
				defer logger.TimeTrack(time.Now(), "genGroup", nil)
				if !dkg.Certified() {
					reportErr(ctx, errc, errors.New("dkg is not certified"))
				}
				if secShare, err := dkg.DistKeyShare(); err == nil {
					group.secShare = secShare
					group.pubPoly = share.NewPubPoly(suite, suite.Point().Base(), group.secShare.Commitments())
					pubKey := group.pubPoly.Commit()
					if pubKeyCoor, err := decodePubKey(pubKey); err == nil {
						if groupId, ok := new(big.Int).SetString(sessionID, 16); ok {
							dataReturn := [5]*big.Int{groupId}
							copy(dataReturn[1:], pubKeyCoor[:])
							select {
							case <-ctx.Done():
							case out <- dataReturn:
							}
						} else {
							reportErr(ctx, errc, errors.New("sessionID cast error "))
						}
					} else {
						reportErr(ctx, errc, err)
					}
				} else {
					reportErr(ctx, errc, err)
				}
			}
		}
	}()
	return
}

func (d *pdkg) GetGroupPublicPoly(groupId string) (pubPoly *share.PubPoly) {
	if g, loaded := d.groups.Load(groupId); loaded {
		pubPoly = g.(*group).pubPoly
	}
	return
}

func (d *pdkg) GetShareSecurity(groupId string) (secShare *share.PriShare) {
	if g, loaded := d.groups.Load(groupId); loaded {
		secShare = g.(*group).secShare.Share
	}
	return
}

func (d *pdkg) GetGroupIDs(groupId string) (participants [][]byte) {
	if g, loaded := d.groups.Load(groupId); loaded {
		participants = g.(*group).participants
		d.logger.Event("GetGroupIDsSucc", map[string]interface{}{"GroupID": groupId, "GroupNumber": d.GetGroupNumber()})
	} else {
		d.logger.Event("GetGroupIDsFail", map[string]interface{}{"GroupID": groupId, "GroupNumber": d.GetGroupNumber()})

	}

	return
}

func (d *pdkg) GetGroupNumber() int {
	length := 0
	d.groups.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	return length
}

func (d *pdkg) GroupDissolve(groupId string) {
	d.groups.Delete(groupId)
	d.logger.Event("GroupDissolve", map[string]interface{}{"GroupID": groupId, "GroupNumber": d.GetGroupNumber()})
}

func decodePubKey(pubKey kyber.Point) (pubKeyCoor [4]*big.Int, err error) {
	pubKeyMar, err := pubKey.MarshalBinary()
	if err != nil {
		return
	}

	for i := 0; i < 4; i++ {
		pubKeyCoor[i] = new(big.Int).SetBytes(pubKeyMar[32*i+1 : 32*i+33])
	}

	return
}

func reportErr(ctx context.Context, errc chan error, err error) {
	fmt.Println("reportErr err ", err)
	select {
	case errc <- err:
	case <-ctx.Done():
	}
	return
}

func fanOut(ctx context.Context, ch chan interface{}, size int) (cs []chan interface{}) {
	cs = make([]chan interface{}, size)
	for i, _ := range cs {
		cs[i] = make(chan interface{})
	}
	go func() {
		for i := range ch {
			for _, c := range cs {
				select {
				case c <- i:
				case <-ctx.Done():
				}
			}
		}
		for _, c := range cs {
			// close all our fanOut channels when the input channel is exhausted.
			close(c)
		}
	}()
	return
}

func mergeErrors(ctx context.Context, logger log.Logger, cs ...chan error) chan error {
	var wg sync.WaitGroup

	out := make(chan error, len(cs))
	output := func(c <-chan error) {
		for n := range c {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		defer logger.TimeTrack(time.Now(), "Grouping", nil)
		wg.Wait()
		close(out)
	}()
	return out
}
