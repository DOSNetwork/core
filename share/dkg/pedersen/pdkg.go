package dkg

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/go-stack/stack"
	errors "golang.org/x/xerrors"
)

// PDKGInterface is a interface for DKG
type PDKGInterface interface {
	Loop()
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

// NewPDKG creates a pdkg struct
func NewPDKG(p p2p.P2PInterface, suite suites.Suite) PDKGInterface {
	d := &pdkg{
		p:         p,
		bufToNode: make(chan interface{}, 50),
		register:  make(chan *group),
		suite:     suite,
		logger:    log.New("module", "dkg"),
	}
	return d
}

func handlePeerMsg(sessionMap map[string][]interface{}, sessionReq map[string]request, p p2p.P2PInterface, sessionID string, content interface{}) {
	switch pubkeyFromPeer := content.(type) {
	case *PublicKey:
		pubkeys := sessionMap[sessionID]
		for _, p := range pubkeys {
			pubkey, ok := p.(*PublicKey)
			if ok {
				if pubkey.Index == pubkeyFromPeer.Index {
					return
				}
			}
		}
	default:
	}
	sessionMap[sessionID] = append(sessionMap[sessionID], content)

	//if sessionMap[sessionID] != nil {
	fmt.Println("sessionMap len ", len(sessionMap[sessionID]))
	if len(sessionMap[sessionID]) == sessionReq[sessionID].numOfResps {
		//go func(req request, m []interface{}) {
		select {
		case <-sessionReq[sessionID].ctx.Done():
		case sessionReq[sessionID].reply <- sessionMap[sessionID]:
		}
		close(sessionReq[sessionID].reply)
		//}(sessionReq[sessionID], sessionMap[sessionID])

		delete(sessionMap, sessionID)
		delete(sessionReq, sessionID)
	}
	//}
}

func handleRequest(sessionMap map[string][]interface{}, sessionReq map[string]request, req request) {
	sessionReq[req.sessionID] = req
	if len(sessionMap[req.sessionID]) == req.numOfResps {
		//go func(req request, m []interface{}) {
		select {
		case <-sessionReq[req.sessionID].ctx.Done():
		case sessionReq[req.sessionID].reply <- sessionMap[req.sessionID]:
		}
		close(req.reply)
		//}(sessionReq[req.sessionID], sessionMap[req.sessionID])

		delete(sessionMap, req.sessionID)
		delete(sessionReq, req.sessionID)
	}
}
func (d *pdkg) Loop() {

	peersToBuf, _ := d.p.SubscribeMsg(50, PublicKey{}, Deal{}, Responses{})
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
				d.p.Reply(context.Background(), msg.Sender, msg.RequestNonce, content)
				handlePeerMsg(sessionPubKeys, sessionReqPubs, d.p, content.SessionId, content)
				fmt.Println("pkdg :PublicKey nonce", msg.RequestNonce)
			case *Deal:
				d.p.Reply(context.Background(), msg.Sender, msg.RequestNonce, content)
				handlePeerMsg(sessionDeals, sessionReqDeals, d.p, content.SessionId, content)
				fmt.Println("pkdg :Deal nonce", msg.RequestNonce)
			case *Responses:
				d.p.Reply(context.Background(), msg.Sender, msg.RequestNonce, content)
				resps := content.Response
				for _, resp := range resps {
					handlePeerMsg(sessionResps, sessionReResps, d.p, content.SessionId, resp)
				}
				fmt.Println("pkdg :Responses")
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
}

func (d *pdkg) Grouping(ctx context.Context, sessionID string, groupIds [][]byte) (chan [5]*big.Int, chan error, error) {
	group := &group{participants: groupIds}
	var errcList []chan error
	if _, loaded := d.groups.LoadOrStore(sessionID, group); loaded {
		return nil, nil, errors.New("dkg: duplicate share public key")
	}

	//Check if all members are reachable
	//connc, errc := d.p.ConnectToAll(ctx, groupIds, sessionID)
	//errcList = append(errcList, errc)
	//exchange pub key
	selfPubc, secrc, errc := genPub(ctx, d.logger, d.suite, d.p.GetID(), groupIds, sessionID)
	errcList = append(errcList, errc)
	selfPubcs := fanOut(ctx, selfPubc, 2)
	errcList = append(errcList, sendToMembers(ctx, d.logger, selfPubcs[0], d.p, groupIds, sessionID))
	peerPubc := askMembers(ctx, d.logger, d.bufToNode, len(groupIds)-1, 0, sessionID)
	errcList = append(errcList, errc)
	partPubsc, errc := exchangePub(ctx, d.logger, selfPubcs[1], peerPubc, d.p, groupIds, sessionID)
	errcList = append(errcList, errc)

	//generate a dkg
	dkgcStep1, errc := genDistKeyGenerator(ctx, d.logger, secrc, partPubsc, len(groupIds), d.suite, sessionID)
	errcList = append(errcList, errc)

	//generate deals for other member and process deals from other member
	dkgcStep2, errc := genDealsAndSend(ctx, d.logger, dkgcStep1, d.p, groupIds, sessionID)
	errcList = append(errcList, errc)
	dkgcStep3, respsc, errc := getAndProcessDeals(ctx, d.logger, dkgcStep2, askMembers(ctx, d.logger, d.bufToNode, len(groupIds)-1, 1, sessionID), sessionID)
	errcList = append(errcList, errc)
	errcList = append(errcList, sendToMembers(ctx, d.logger, respsc, d.p, groupIds, sessionID))

	//process response to certify dkg and generate a group sec and pub key
	cetifiedDkgc, errc := getAndProcessResponses(ctx, d.logger, dkgcStep3, askMembers(ctx, d.logger, d.bufToNode, (len(groupIds)-1)*(len(groupIds)-1), 2, sessionID), sessionID)
	outc, errc := genGroup(ctx, d.logger, group, d.suite, cetifiedDkgc, sessionID)
	errcList = append(errcList, errc)
	errc = mergeErrors(d.logger, sessionID, errcList...)
	return outc, errc, nil
}

func (d *pdkg) GetGroupPublicPoly(groupId string) (pubPoly *share.PubPoly) {
	if g, loaded := d.groups.Load(groupId); loaded {
		pubPoly = g.(*group).pubPoly
	}
	return
}

func (d *pdkg) GetShareSecurity(groupId string) (secShare *share.PriShare) {
	if g, loaded := d.groups.Load(groupId); loaded {
		if g != nil {
			if dks := g.(*group).secShare; dks != nil {
				secShare = dks.Share
			} else {
				err := &DKGError{err: errors.Errorf("GetShareSecurity from %s failed : %w", groupId, ErrCanNotLoadSec)}
				d.logger.Error(err)
			}
		} else {
			err := &DKGError{err: errors.Errorf("GetShareSecurity from %s failed : %w", groupId, ErrCanNotLoadGroup)}
			d.logger.Error(err)
		}
	}
	return
}

func (d *pdkg) GetGroupIDs(groupId string) (participants [][]byte) {
	if g, loaded := d.groups.Load(groupId); loaded {
		participants = g.(*group).participants
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
	s := stack.Trace().TrimRuntime()
	//d.logger.Error(err)
	fmt.Println("reportErr err ", err, s)
	select {
	case errc <- err:
	case <-ctx.Done():
	}
	return
}

func fanOut(ctx context.Context, ch chan interface{}, size int) (cs []chan interface{}) {
	cs = make([]chan interface{}, size)
	for i := range cs {
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

func mergeErrors(logger log.Logger, sessionID string, cs ...chan error) chan error {
	var wg sync.WaitGroup

	out := make(chan error, len(cs))
	output := func(c <-chan error) {
		for n := range c {
			select {
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
		defer logger.TimeTrack(time.Now(), "Grouping", map[string]interface{}{"GroupID": sessionID})
		wg.Wait()
		close(out)
	}()
	return out
}
