package dkg

import (
	"bytes"
	"context"
	"encoding/hex"
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
	"golang.org/x/crypto/sha3"
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
	GetGroupNumber() int
	Grouping(ctx context.Context, groupId string, Participants [][]byte) (chan [5]*big.Int, <-chan error, error)
	GroupDissolve(groupId string)
}

func NewPDKG(p p2p.P2PInterface, suite suites.Suite) PDKGInterface {
	logger = log.New("module", "dkg")
	pdkg := &PDKG{
		p:         p,
		bufToNode: make(chan interface{}, 50),
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
		logger.Event("GetGroupIDsSucc", map[string]interface{}{"GroupID": groupId, "GroupNumber": d.GetGroupNumber()})

	} else {
		logger.Event("GetGroupIDsFail", map[string]interface{}{"GroupID": groupId, "GroupNumber": d.GetGroupNumber()})

	}

	return
}
func mergeErrors(ctx context.Context, cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	// We must ensure that the output channel has the capacity to
	// hold as many errors
	// as there are error channels.
	// This will ensure that it never blocks, even
	// if WaitForPipeline returns early.
	out := make(chan error, len(cs))
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls
	// wg.Done.
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
	// Start a goroutine to close out once all the output goroutines
	// are done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func (d *PDKG) Grouping(ctx context.Context, groupId string, participants [][]byte) (chan [5]*big.Int, <-chan error, error) {
	group := &Group{Participants: participants}
	var errcList []<-chan error
	if _, loaded := d.groups.LoadOrStore(groupId, group); loaded {
		return nil, nil, errors.New("dkg: duplicate share public key")
	} else {
		dkgc, errc := exchangePub(ctx, d.suite, d.bufToNode, participants, d.p, groupId)
		errcList = append(errcList, errc)
		dkgCetifiedc, errc := processDeal(ctx, dkgc, d.bufToNode, participants, d.p, groupId)
		errcList = append(errcList, errc)
		outc, errc := genPubKey(ctx, group, d.suite, dkgCetifiedc, groupId)
		errcList = append(errcList, errc)
		errc = mergeErrors(ctx, errcList...)
		return outc, errc, nil
	}
}

func (d *PDKG) GetGroupNumber() int {
	length := 0
	d.groups.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	return length
}

func (d *PDKG) GroupDissolve(groupId string) {
	d.groups.Delete(groupId)
	logger.Event("GroupDissolve", map[string]interface{}{"GroupID": groupId, "GroupNumber": d.GetGroupNumber()})
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
					logger.Event("PublicKeyFromPeer", map[string]interface{}{"GroupID": content.SessionId})

					go func() { d.p.Reply(msg.Sender, msg.RequestNonce, &PublicKey{}) }()
				case *Deal:
					sessionDeals[content.SessionId] = append(sessionDeals[content.SessionId], content)
					logger.Event("DealsFromPeer", map[string]interface{}{"GroupID": content.SessionId})

					go func() { d.p.Reply(msg.Sender, msg.RequestNonce, &Deal{}) }()
				case *Responses:
					sessionResps[content.SessionId] = append(sessionResps[content.SessionId], content.Response...)
					logger.Event("ResponsesFromPeer", map[string]interface{}{"GroupID": content.SessionId})
					go func() { d.p.Reply(msg.Sender, msg.RequestNonce, &Responses{}) }()
				}
			case msg, ok := <-d.bufToNode:
				if !ok {
					return
				}
				switch content := msg.(type) {
				case ReqPubs:
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
		defer logger.TimeTrack(time.Now(), "genPubKeyDone", map[string]interface{}{"GroupID": sessionID})

		defer close(out)
		defer close(errc)
		logger.Event("genPubKey", map[string]interface{}{"GroupID": sessionID})
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
			err := errors.New("!dkg.Certified")
			logger.Error(err)
			errc <- err
			return
		}
		var err error
		group.SecShare, err = dkg.DistKeyShare()
		if err != nil {
			logger.Error(err)
			errc <- err
			return
		}
		group.PubPoly = share.NewPubPoly(suite, suite.Point().Base(), group.SecShare.Commitments())
		pubKey := group.PubPoly.Commit()
		pubKeyCoor, err := decodePubKey(pubKey)
		if err != nil {
			select {
			case <-ctx.Done():
				return
			case errc <- err:
				logger.Error(err)
			}
			return
		}
		groupId, ok := new(big.Int).SetString(sessionID, 16)
		if !ok {
			select {
			case <-ctx.Done():
				return
			case errc <- err:
				logger.Error(err)
			}
			return
		}
		dataReturn := [5]*big.Int{groupId}
		copy(dataReturn[1:], pubKeyCoor[:])

		select {
		case <-ctx.Done():
			return
		case out <- dataReturn:
		}
	}()
	return out, errc
}
func ByteTohex(a []byte) string {
	unchecksummed := hex.EncodeToString(a[:])
	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(unchecksummed))
	hash := sha.Sum(nil)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return "0x" + string(result)
}
func exchangePub(ctx context.Context, suite suites.Suite, bufToNode chan interface{}, groupIds [][]byte, p p2p.P2PInterface, sessionID string) (<-chan *DistKeyGenerator, <-chan error) {
	out := make(chan *DistKeyGenerator)
	errc := make(chan error)
	go func() {
		defer logger.TimeTrack(time.Now(), "exchangePubDone", map[string]interface{}{"GroupID": sessionID})
		defer close(out)
		defer close(errc)
		logger.Event("exchangePub", map[string]interface{}{"GroupID": sessionID})

		for i := 0; i < len(groupIds); i++ {
			start := time.Now()
			if !bytes.Equal(p.GetID(), groupIds[i]) {
				retry := 0
				for {
					if retry >= 10 {
						break
					}
					if _, err := p.ConnectTo("", groupIds[i]); err != nil {
						fmt.Println("ConnectTo done retry=", retry, err)
						retry++
						time.Sleep(1 * time.Second)
					} else {
						break
					}
				}

				f := map[string]interface{}{
					"GroupID":  sessionID,
					"retry":    retry,
					"costTime": time.Since(start).Nanoseconds() / 1000,
					"From":     ByteTohex(p.GetID()),
					"To":       ByteTohex(groupIds[i])}
				if retry >= 10 {
					logger.Event("DKGConnectToFaile", f)
				} else {
					logger.Event("DKGConnectToSuccess", f)
				}
			}
		}

		//Generate secret and public key
		sec := suite.Scalar().Pick(suite.RandomStream())
		pub := suite.Point().Mul(sec, nil)
		bin, err := pub.MarshalBinary()
		if err != nil {
			fmt.Println(err)
			logger.Error(err)
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

		pubkey := &PublicKey{SessionId: sessionID, Index: uint32(index), Publickey: &vss.PublicKey{Binary: bin}}
		var partPubs []*PublicKey
		partPubs = append(partPubs, pubkey)

		var wg sync.WaitGroup
		wg.Add(len(groupIds) - 1)
		for i, id := range groupIds {
			if r := bytes.Compare(p.GetID(), id); r != 0 {
				go func(i int, id []byte) {
					defer wg.Done()

					_, err := p.Request(id, pubkey)
					if err != nil {
						select {
						case <-ctx.Done():
							return
						case errc <- err:
							logger.Error(err)
						}
						return
					}

				}(i, id)
			}
		}
		wg.Wait()

		select {
		case err := <-errc:
			logger.Error(err)
			return
		default:
		}

		reqChan := make(chan []*PublicKey)
		req := ReqPubs{ctx: ctx, SessionId: sessionID, numOfPubs: len(groupIds) - 1, Pubkeys: reqChan}
		select {
		case <-ctx.Done():
			return
		case bufToNode <- req:
		}

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
						logger.Error(err)
					}
					return
				}
			}
			//var partPubs []kyber.Point

			dkg, err := NewDistKeyGenerator(suite, sec, pubPoints, len(groupIds)/2+1)
			if err != nil {
				fmt.Println("NewDistKeyGenerator err ", err)
				logger.Error(err)
				errc <- err
				return
			}
			pubPoints = nil
			out <- dkg
		} else {
			logger.Error(err)
			errc <- err
		}
	}()
	return out, errc
}

func processDeal(ctx context.Context, dkgc <-chan *DistKeyGenerator, bufToNode chan interface{}, groupIds [][]byte, p p2p.P2PInterface, sessionID string) (<-chan *DistKeyGenerator, <-chan error) {
	out := make(chan *DistKeyGenerator)
	errc := make(chan error)
	go func() {
		defer close(out)
		defer close(errc)
		defer logger.TimeTrack(time.Now(), "processDealDone", map[string]interface{}{"GroupID": sessionID})
		logger.Event("processDeal", map[string]interface{}{"GroupID": sessionID})

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
			logger.Error(err)
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
			go func(id []byte, d *Deal) {
				defer wg.Done()
				_, err := p.Request(id, d)
				if err != nil {
					logger.Error(err)
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
					logger.Error(err)
					return
				}
				if vss.StatusApproval != resp.Response.Status {
					var err = errors.New("resp StatusNotApproval")
					logger.Error(err)
					select {
					case <-ctx.Done():
						return
					case errc <- err:
						return
					}
				}
				resps = append(resps, resp)
			}
		}

		wg.Add(len(groupIds) - 1)
		for i, id := range groupIds {
			if r := bytes.Compare(p.GetID(), id); r != 0 {
				go func(i int, id []byte) {
					defer wg.Done()
					_, err := p.Request(id, &Responses{SessionId: sessionID, Response: resps})
					if err != nil {
						logger.Error(err)
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
					logger.Error(err)
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
			err = errors.New("dkg not cetified")
			logger.Error(err)
			select {
			case <-ctx.Done():
				return
			case errc <- err:
			}
		}
	}()
	return out, errc
}
