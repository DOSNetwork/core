package dkg

import (
	"bytes"
	"fmt"
	"sort"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"

	"github.com/dedis/kyber"
)

const (
	INIT = iota
	PUBLICKEYDONE
	DEALDONE
	RESPONSEDONE
	VERIFIED
)

type P2PDkgInterface interface {
	SetGroupId([]byte)
	GetGroupId() []byte
	GetDKGIndex() int
	SetGroupMembers([][]byte)
	GetGroupCmd() chan DkgSession
	GetDkgEvent() chan int
	SetNbParticipants(int)
	GetGroupPublicPoly() *share.PubPoly
	GetShareSecurity() *share.PriShare
	IsCertified() bool
	Reset(sessionId string)
	SubscribeEvent(chan int)
}

type DkgSession struct {
	SessionId string
	GroupIds  [][]byte
}

func CreateP2PDkg(p p2p.P2PInterface, suite suites.Suite, peerEvent chan p2p.P2PMessage) P2PDkgInterface {
	sec := suite.Scalar().Pick(suite.RandomStream())
	d := &P2PDkg{
		suite:                  suite,
		publicKeys:             Pubkeys{},
		responses:              Responses{},
		idPubkeyMap:            make(map[string]string),
		dealIdMap:              make(map[int]*Deal),
		partSec:                sec,
		partPub:                suite.Point().Mul(sec, nil),
		groupCmd:               make(chan DkgSession),
		network:                &p,
		chPeerEvent:            peerEvent,
		currentState:           INIT,
		subscribeEvent:         make(chan int),
		finExchangePubKey:      make(chan struct{}),
		finNewDistKeyGenerator: make(chan struct{}),
		finProcessDeal:         make(chan struct{}),
		finProcessResponses:    make(chan struct{}),
		logger:                 log.New("module", "dkg"),
	}
	go d.eventLoop()
	return d
}

func (d *P2PDkg) GetGroupCmd() chan DkgSession {
	return d.groupCmd
}

func (d *P2PDkg) GetDkgEvent() chan int {
	return d.subscribeEvent
}

type P2PDkg struct {
	groupId []byte
	suite   suites.Suite
	//Data Buffer
	publicKeys Pubkeys
	deals      []Deal
	responses  Responses
	//Group member ID
	groupIds [][]byte
	//node key pair
	partPub   kyber.Point
	partSec   kyber.Scalar
	pubToSend vss.PublicKey
	//
	partDkg      *DistKeyGenerator
	partDks      *DistKeyShare
	groupPubPoly *share.PubPoly
	//
	idPubkeyMap map[string]string
	dealIdMap   map[int]*Deal
	//
	nbParticipants int
	subscribeEvent chan int
	network        *p2p.P2PInterface
	chPeerEvent    chan p2p.P2PMessage

	groupingStart time.Time
	dkgIndex      int

	currentState           int
	currentSession         string
	groupCmd               chan DkgSession
	finExchangePubKey      chan struct{}
	finNewDistKeyGenerator chan struct{}
	finProcessDeal         chan struct{}
	finProcessResponses    chan struct{}
	logger                 log.Logger
}

func (d *P2PDkg) SetGroupId(id []byte) {
	d.groupId = id
}

func (d *P2PDkg) GetGroupId() []byte {
	return d.groupId
}

func (d *P2PDkg) GetDKGIndex() int {
	return d.dkgIndex
}

func (d *P2PDkg) SetNbParticipants(n int) {
	if d.currentState == INIT {
		d.nbParticipants = n
	}
}

func (d *P2PDkg) SetGroupMembers(members [][]byte) {
	if d.currentState == INIT {
		d.groupIds = members
	}
}

func (d *P2PDkg) IsCertified() bool {
	if d.currentState == VERIFIED {
		return true
	} else {
		return false
	}
}

func (d *P2PDkg) Reset(sessionId string) {
	d.currentState = INIT
	if d.currentSession != "" {
		d.allFinished()
	}
	d.publicKeys = d.publicKeys[:0]
	d.deals = d.deals[:0]
	d.responses = Responses{}
	d.idPubkeyMap = make(map[string]string)
	d.dealIdMap = make(map[int]*Deal)
	d.partSec = d.suite.Scalar().Pick(d.suite.RandomStream())
	d.partPub = d.suite.Point().Mul(d.partSec, nil)
	d.pubToSend = vss.PublicKey{}
	d.currentSession = sessionId
}

func (d *P2PDkg) GetGroupPublicPoly() *share.PubPoly {
	if d.currentState == VERIFIED {
		return d.groupPubPoly
	}
	return nil
}

func (d *P2PDkg) GetShareSecurity() *share.PriShare {
	if d.currentState == VERIFIED {
		return d.partDks.Share
	}
	return nil
}

func (d *P2PDkg) SubscribeEvent(dkgEvent chan int) {
	d.subscribeEvent = dkgEvent
}

func (d *P2PDkg) allFinished() {
	<-d.finExchangePubKey
	<-d.finNewDistKeyGenerator
	<-d.finProcessDeal
	<-d.finProcessResponses
	d.finExchangePubKey = make(chan struct{})
	d.finNewDistKeyGenerator = make(chan struct{})
	d.finProcessDeal = make(chan struct{})
	d.finProcessResponses = make(chan struct{})
}

func (d *P2PDkg) eventLoop() {
	for {
		select {
		//event from eth
		case groupCmd := <-d.GetGroupCmd():
			d.Reset(groupCmd.SessionId)
			d.pipeExchangePubKey(groupCmd.GroupIds, groupCmd.SessionId)
			d.pipeNewDistKeyGenerator(groupCmd.SessionId)
			d.pipeProcessDeal(groupCmd.SessionId)
			d.pipeProcessResponses(groupCmd.SessionId)
		//event from peer
		case msg := <-d.chPeerEvent:
			switch content := msg.Msg.Message.(type) {
			case *ReqPublicKey:
				if d.currentState >= PUBLICKEYDONE && d.currentSession == content.SessionId {
					msg.PeerConn.Reply(msg.RequestNonce, &d.pubToSend)
				} else {
					msg.PeerConn.Reply(msg.RequestNonce, &vss.PublicKey{})
				}
			case *ReqDeal:
				if d.currentState >= DEALDONE && d.currentSession == content.SessionId {
					targetPubKey := d.idPubkeyMap[string(msg.Sender)]
					for index, pubKey := range d.publicKeys {
						if pubKey.String() == targetPubKey {
							msg.PeerConn.Reply(msg.RequestNonce, d.dealIdMap[index])
							break
						}
					}
				} else {
					msg.PeerConn.Reply(msg.RequestNonce, &Deal{})
				}
			case *ReqResponses:
				if d.currentState >= RESPONSEDONE && d.currentSession == content.SessionId {
					msg.PeerConn.Reply(msg.RequestNonce, &d.responses)
				} else {
					msg.PeerConn.Reply(msg.RequestNonce, &Responses{})
				}
			default:
			}
		}
	}
}

func (d *P2PDkg) pipeExchangePubKey(groupIds [][]byte, sessionId string) {
	go func() {
		if d.currentState == INIT && d.currentSession == sessionId {
			defer close(d.finExchangePubKey)
			d.groupIds = groupIds
			d.nbParticipants = len(groupIds)
			d.groupingStart = time.Now()

			id := (*d.network).GetID()
			d.idPubkeyMap[string(id)] = d.partPub.String()
			d.publicKeys = append(d.publicKeys, d.partPub)
			public := vss.PublicKey{SenderId: id}
			if err := public.SetPoint(d.suite, d.partPub); err != nil {
				d.logger.Error(err)
			}
			d.pubToSend = public
			d.currentState = PUBLICKEYDONE

			groupPeers := make(map[string][]byte)
			for _, id := range d.groupIds {
				if bytes.Compare(id, (*d.network).GetID()) != 0 {
					groupPeers[string(id)] = id
				}
			}
			for len(groupPeers) > 0 {
				for key, id := range groupPeers {
					if d.currentSession == sessionId {
						pubkey, err := (*d.network).Request(id, &ReqPublicKey{SessionId: sessionId})
						if err != nil {
							continue
						}
						switch content := pubkey.(type) {
						case *vss.PublicKey:
							if content.GetBinary() != nil {
								p, err := content.GetPoint(d.suite)
								if err != nil {
									d.logger.Error(err)
								}
								d.idPubkeyMap[string(content.SenderId)] = p.String()
								d.publicKeys = append(d.publicKeys, p)
								delete(groupPeers, key)
							}
						}
					} else {
						return
					}
				}
			}
		}
		fmt.Println("pipeExchangePubKey closed")
	}()
}

func (d *P2PDkg) pipeNewDistKeyGenerator(sessionId string) {
	go func() {
		defer close(d.finNewDistKeyGenerator)
		<-d.finExchangePubKey
		sort.Sort(d.publicKeys)
		var err error
		d.partDkg, err = NewDistKeyGenerator(d.suite, d.partSec, d.publicKeys, d.nbParticipants/2+1)
		if err != nil {
			d.logger.Error(err)
		}
		d.dealIdMap, err = d.partDkg.Deals()
		if err != nil {
			d.logger.Error(err)
		}
		for i, pub := range d.publicKeys {
			if pub.Equal(d.partPub) {
				d.dkgIndex = i
				break
			}
		}
		d.currentState = DEALDONE
		fmt.Println("pipeNewDistKeyGenerator closed")
	}()
}

func (d *P2PDkg) pipeProcessDeal(sessionId string) {
	go func() {
		defer close(d.finProcessDeal)
		<-d.finNewDistKeyGenerator
		groupPeers := make(map[string][]byte)
		for _, id := range d.groupIds {
			if bytes.Compare(id, (*d.network).GetID()) != 0 {
				groupPeers[string(id)] = id
			}
		}
		for len(groupPeers) > 0 {
			for key, id := range groupPeers {
				if d.currentSession == sessionId {
					deal, err := (*d.network).Request(id, &ReqDeal{SessionId: sessionId})
					if err != nil {
						continue
					}
					switch content := deal.(type) {
					case *Deal:
						if content.GetDeal() != nil {
							d.deals = append(d.deals, *content)
							delete(groupPeers, key)
						}
					}
				} else {
					return
				}
			}
		}
		var resps []*Response
		for _, deal := range d.deals {
			resp, err := (*d.partDkg).ProcessDeal(&deal)
			if err != nil {
				d.logger.Error(err)
			} else {
				resps = append(resps, resp)
			}
		}
		d.responses = Responses{Response: resps}
		d.currentState = RESPONSEDONE
		fmt.Println("pipeProcessDeal closed")
	}()
}

func (d *P2PDkg) pipeProcessResponses(sessionId string) {
	go func() {
		defer close(d.finProcessResponses)
		<-d.finProcessDeal
		groupPeers := make(map[string][]byte)
		for _, id := range d.groupIds {
			if bytes.Compare(id, (*d.network).GetID()) != 0 {
				groupPeers[string(id)] = id
			}
		}
		for len(groupPeers) > 0 {
			for key, id := range groupPeers {
				if d.currentSession == sessionId {
					responses, err := (*d.network).Request(id, &ReqResponses{SessionId: sessionId})
					if err != nil {
						continue
					}
					switch content := responses.(type) {
					case *Responses:
						if resps := content.GetResponse(); resps != nil {
							for _, r := range resps {
								if _, err := (*d.partDkg).ProcessResponse(r); err != nil {
									d.logger.Error(err)
								}
							}
							delete(groupPeers, key)
						}
					}
				} else {
					return
				}
			}
		}
		if (*d.partDkg).Certified() {
			var err error
			d.partDks, err = d.partDkg.DistKeyShare()
			if err != nil {
				d.logger.Error(err)
			}
			d.groupPubPoly = share.NewPubPoly(d.suite, d.suite.Point().Base(), d.partDks.Commitments())
			d.currentState = VERIFIED
			timeCost := time.Since(d.groupingStart).Seconds()
			fmt.Println("DistKeyShare SUCCESS ", timeCost)
			if d.subscribeEvent != nil && d.currentSession == sessionId {
				d.subscribeEvent <- VERIFIED
			}
		}
		fmt.Println("pipeProcessResponses closed")
	}()
}
