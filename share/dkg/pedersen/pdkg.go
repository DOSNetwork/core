package dkg

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
)

var logger log.Logger

type ReqPubs struct {
	ctx       context.Context
	SessionId string
	numOfPubs int
	Pubkeys   chan []*PublicKey
}

type ReqDeals struct {
	ctx        context.Context
	SessionId  string
	numOfDeals int
	Reply      chan []*Deal
}

type ReResps struct {
	ctx        context.Context
	SessionId  string
	numOfResps int
	Reply      chan []*Response
}

type PDKGInterface interface {
	GetGroupPublicPoly(groupId string) *share.PubPoly
	GetShareSecurity(groupId string) *share.PriShare
	GetGroupIDs(groupId string) [][]byte
	Grouping(ctx context.Context, groupId string, Participants [][]byte) (chan [5]*big.Int, <-chan error, error)
	GroupDissolve(groupId string)
}

func NewPDKG(p p2p.P2PInterface, suite suites.Suite) PDKGInterface {
	logger = log.New("module", "dkg")
	pdkg := &PDKG{
		p:         p,
		bufToNode: make(chan interface{}, 10),
		register:  make(chan *Group),
		suite:     suite,
	}
	pdkg.Listen()
	return pdkg
}

type PDKG struct {
	p         p2p.P2PInterface
	suite     suites.Suite
	bufToNode chan interface{}
	register  chan *Group
	groups    sync.Map
}

type Group struct {
	Participants [][]byte
	SecShare     *DistKeyShare
	PubPoly      *share.PubPoly
}

func (d *PDKG) GetGroupPublicPoly(groupId string) (pubPoly *share.PubPoly) {
	if group, loaded := d.groups.Load(groupId); loaded {
		pubPoly = group.(*Group).PubPoly
	}
	return
}

func (d *PDKG) GetShareSecurity(groupId string) (secShare *share.PriShare) {
	if group, loaded := d.groups.Load(groupId); loaded {
		secShare = group.(*Group).SecShare.Share
	}
	return
}

func (d *PDKG) GetGroupIDs(groupId string) (participants [][]byte) {
	if group, loaded := d.groups.Load(groupId); loaded {
		participants = group.(*Group).Participants
	}
	return
}

func (d *PDKG) Grouping(ctx context.Context, groupId string, participants [][]byte) (chan [5]*big.Int, <-chan error, error) {
	group := &Group{Participants: participants}
	if _, loaded := d.groups.LoadOrStore(groupId, group); loaded {
		return nil, nil, errors.New("dkg: duplicate share public key")
	} else {
		dkgc, _ := exchangePub(ctx, d.suite, d.bufToNode, participants, d.p, groupId)
		dkgCetifiedc, _ := processDeal(ctx, dkgc, d.bufToNode, participants, d.p, groupId)
		outc, errc := genPubKey(ctx, group, d.suite, dkgCetifiedc, groupId)
		return outc, errc, nil
	}
}

func (d *PDKG) GroupDissolve(groupId string) {
	//	for _, id := range d.dkgs[groupId].Participants {
	//		d.p.DisConnectTo(id)
	//	}
	//	d.dkgs[groupId] = nil
}

func (d *PDKG) Listen() {
	peersToBuf, _ := d.p.SubscribeEvent(10, PublicKey{}, Deal{}, Responses{})
	go func() {
		sessionPubKeys := map[string][]*PublicKey{}
		sessionDeals := map[string][]*Deal{}
		sessionResps := map[string][]*Response{}
		for {
			select {
			case msg, ok := <-peersToBuf:
				if !ok {
					return
				}
				switch content := msg.Msg.Message.(type) {
				case *PublicKey:
					sessionPubKeys[content.SessionId] = append(sessionPubKeys[content.SessionId], content)
					go func() { d.p.Reply(msg.Sender, msg.RequestNonce, &PublicKey{}) }()
					fmt.Println(content.SessionId, " - ", string(d.p.GetID()), " Got sessionPubKeys ", len(sessionPubKeys[content.SessionId]))
				case *Deal:
					sessionDeals[content.SessionId] = append(sessionDeals[content.SessionId], content)
					go func() { d.p.Reply(msg.Sender, msg.RequestNonce, &Deal{}) }()
					fmt.Println(content.SessionId, " - ", string(d.p.GetID()), "Got Deal ", len(sessionDeals[content.SessionId]), content.SessionId)
				case *Responses:
					sessionResps[content.SessionId] = append(sessionResps[content.SessionId], content.Response...)
					go func() { d.p.Reply(msg.Sender, msg.RequestNonce, &Responses{}) }()
				}
			case msg, ok := <-d.bufToNode:
				if !ok {
					return
				}
				switch content := msg.(type) {
				case ReqPubs:
					fmt.Println("GROUPid:", content.SessionId, "    Request sessionPubKeys ", len(sessionPubKeys[content.SessionId]), content.numOfPubs)
					if len(sessionPubKeys[content.SessionId]) != content.numOfPubs {
						go func() {
							time.Sleep(1 * time.Second)
							select {
							case d.bufToNode <- content:
							case <-content.ctx.Done():
							}
							return
						}()
					} else {
						pubkeys := sessionPubKeys[content.SessionId]
						sessionPubKeys[content.SessionId] = nil
						content.Pubkeys <- pubkeys
					}
				case ReqDeals:
					if len(sessionDeals[content.SessionId]) != content.numOfDeals {
						go func() {
							time.Sleep(1 * time.Second)
							select {
							case d.bufToNode <- content:
							case <-content.ctx.Done():
							}
							return
						}()
					} else {
						deals := sessionDeals[content.SessionId]
						sessionDeals[content.SessionId] = nil
						content.Reply <- deals
					}
				case ReResps:
					if len(sessionResps[content.SessionId]) != content.numOfResps {
						go func() {
							time.Sleep(1 * time.Second)
							select {
							case d.bufToNode <- content:
							case <-content.ctx.Done():
							}
							return
						}()
					} else {
						resps := sessionResps[content.SessionId]
						sessionResps[content.SessionId] = nil
						content.Reply <- resps
					}
				}
			}
		}
	}()
	return
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

func genPubKey(ctx context.Context, group *Group, suite suites.Suite, dkgc <-chan *DistKeyGenerator, sessionID string) (chan [5]*big.Int, <-chan error) {
	out := make(chan [5]*big.Int)
	errc := make(chan error)
	go func() {
		defer close(out)
		defer close(errc)
		fmt.Println("genPubKey")

		var dkg *DistKeyGenerator
		var ok bool
		select {
		case dkg, ok = <-dkgc:
			if !ok {
				return
			}
		case <-ctx.Done():
			return
		}
		if !dkg.Certified() {
			errc <- errors.New("!dkg.Certified")
			return
		}
		var err error
		group.SecShare, err = dkg.DistKeyShare()
		if err != nil {
			errc <- err
			return
		}
		group.PubPoly = share.NewPubPoly(suite, suite.Point().Base(), group.SecShare.Commitments())
		pubKey := group.PubPoly.Commit()
		pubKeyCoor, err := decodePubKey(pubKey)
		groupId, _ := new(big.Int).SetString(sessionID, 16)
		dataReturn := [5]*big.Int{groupId}
		copy(dataReturn[1:], pubKeyCoor[:])
		fmt.Println("genPubKey ", dataReturn)

		select {
		case <-ctx.Done():
			return
		case out <- dataReturn:
		}
	}()
	return out, errc
}
func exchangePub(ctx context.Context, suite suites.Suite, bufToNode chan interface{}, groupIds [][]byte, p p2p.P2PInterface, sessionID string) (<-chan *DistKeyGenerator, <-chan error) {
	out := make(chan *DistKeyGenerator)
	errc := make(chan error)
	go func() {
		defer fmt.Println(sessionID, "exchangePub close !!!!!!!!!!!!!!!!!!")
		defer close(out)
		defer close(errc)
		//Generate secret and public key
		fmt.Println(sessionID, "exchangePub")
		sec := suite.Scalar().Pick(suite.RandomStream())
		pub := suite.Point().Mul(sec, nil)
		bin, err := pub.MarshalBinary()
		if err != nil {
			fmt.Println(err)
			return
		}
		//Index pub key
		index := 0
		for i, id := range groupIds {
			if r := bytes.Compare(p.GetID(), id); r == 0 {
				index = i
				break
			}
		}
		//fmt.Println(string(p.GetID()), "index ", index, pub)

		pubkey := &PublicKey{SessionId: sessionID, Index: uint32(index), Publickey: &vss.PublicKey{Binary: bin}}
		var partPubs []*PublicKey
		partPubs = append(partPubs, pubkey)

		var wg sync.WaitGroup
		wg.Add(len(groupIds) - 1)
		for i, id := range groupIds {
			if r := bytes.Compare(p.GetID(), id); r != 0 {
				go func(i int, id []byte) {
					defer wg.Done()
					fmt.Println(sessionID, "Send pub key to ", string(id))

					_, err := p.Request(id, pubkey)
					if err != nil {
						fmt.Println(sessionID, "Send pub key to ", string(id), " err ", err)
						select {
						case <-ctx.Done():
							return
						case errc <- err:
						}
						return
					}
				}(i, id)
			}
		}
		wg.Wait()

		select {
		case err := <-errc:
			fmt.Println("Err ", err)
			return
		default:
		}
		fmt.Println(sessionID, "wg.Wait()111111111")

		reqChan := make(chan []*PublicKey)
		req := ReqPubs{ctx: ctx, SessionId: sessionID, numOfPubs: len(groupIds) - 1, Pubkeys: reqChan}
		select {
		case <-ctx.Done():
			return
		case bufToNode <- req:
		}
		fmt.Println(sessionID, "wg.Wait()22222222")

		select {
		case <-ctx.Done():
			return
		case pubs := <-reqChan:
			partPubs = append(partPubs, pubs...)
		}
		//}

		if err == nil {
			pubPoints := make([]kyber.Point, len(groupIds))
			for _, pub := range partPubs {
				pubPoints[pub.Index] = suite.Point()
				if err := pubPoints[pub.Index].UnmarshalBinary(pub.Publickey.Binary); err != nil {
					select {
					case <-ctx.Done():
						return
					case errc <- err:
					}
					return
				}
			}
			//var partPubs []kyber.Point

			dkg, err := NewDistKeyGenerator(suite, sec, pubPoints, len(groupIds)/2+1)
			if err != nil {
				fmt.Println("NewDistKeyGenerator err ", err)
				errc <- err
				return
			}
			fmt.Println("Result:", string(p.GetID()))
			pubPoints = nil
			out <- dkg
		} else {
			errc <- err
		}
	}()
	return out, errc
}

func processDeal(ctx context.Context, dkgc <-chan *DistKeyGenerator, bufToNode chan interface{}, groupIds [][]byte, p p2p.P2PInterface, sessionID string) (<-chan *DistKeyGenerator, <-chan error) {
	out := make(chan *DistKeyGenerator)
	errc := make(chan error)
	go func() {
		defer fmt.Println("processDeal close !!!!!!!!!!!!!!!!!!")
		defer close(out)
		defer close(errc)
		fmt.Println("processDeal")

		var dkg *DistKeyGenerator
		var ok bool
		select {
		case dkg, ok = <-dkgc:
			if !ok {
				return
			}
		case <-ctx.Done():
			return
		}
		deals, err := dkg.Deals()
		if err != nil {
			select {
			case <-ctx.Done():
				return
			case errc <- err:
			}
		}
		var wg sync.WaitGroup
		wg.Add(len(groupIds) - 1)
		for i, d := range deals {
			d.SessionId = sessionID
			fmt.Println(string(p.GetID()), " SendDeal_", i, " to ", groupIds[i])
			go func(id []byte, d *Deal) {
				defer wg.Done()
				_, err := p.Request(id, d)
				if err != nil {
					select {
					case <-ctx.Done():
						return
					case errc <- err:
					}
					return
				}
			}(groupIds[i], d)
		}
		wg.Wait()
		fmt.Println("processDeal	")
		reply := make(chan []*Deal)
		req := ReqDeals{ctx: ctx, SessionId: sessionID, numOfDeals: len(groupIds) - 1, Reply: reply}
		select {
		case <-ctx.Done():
			return
		case bufToNode <- req:
		}

		var resps []*Response
		select {
		case <-ctx.Done():
			return
		case deals := <-reply:
			for _, d := range deals {
				resp, err := dkg.ProcessDeal(d)
				resp.SessionId = sessionID
				if err != nil {
					fmt.Println(" ProcessDeal err ", err)
					return
				}
				fmt.Println(" resp Status ", resp.Response.Status, " index ", resp.Index)
				if vss.StatusApproval != resp.Response.Status {
					select {
					case <-ctx.Done():
						return
					case errc <- errors.New("resp StatusNotApproval"):
						return
					}
				}
				resps = append(resps, resp)
			}
		}
		fmt.Println("len of resps ", len(resps))

		wg.Add(len(groupIds) - 1)
		for i, id := range groupIds {
			if r := bytes.Compare(p.GetID(), id); r != 0 {
				go func(i int, id []byte) {
					defer wg.Done()
					_, err := p.Request(id, &Responses{SessionId: sessionID, Response: resps})
					if err != nil {
						select {
						case <-ctx.Done():
							return
						case errc <- err:
						}
						return
					}
				}(i, id)
			}
		}
		wg.Wait()

		replyResp := make(chan []*Response)
		reqResp := ReResps{ctx: ctx, SessionId: sessionID, numOfResps: (len(groupIds) - 1) * (len(groupIds) - 1), Reply: replyResp}
		select {
		case <-ctx.Done():
			return
		case bufToNode <- reqResp:
		}
		select {
		case <-ctx.Done():
			return
		case resps := <-replyResp:
			for _, r := range resps {
				_, err := dkg.ProcessResponse(r)
				if err != nil {
					fmt.Println("ProcessResponse err ", err)
				}
			}
		}
		if dkg.Certified() {
			select {
			case <-ctx.Done():
				return
			case out <- dkg:
			}
		} else {
			select {
			case <-ctx.Done():
				return
			case errc <- errors.New("dkg not cetified"):
			}
		}
	}()
	return out, errc
}
