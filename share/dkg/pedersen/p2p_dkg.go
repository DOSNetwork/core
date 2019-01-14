package dkg

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"

	"github.com/dedis/kyber"

	"github.com/golang/protobuf/proto"

	"github.com/sirupsen/logrus"
)

const (
	INIT = iota
	VERIFIED
)

var log = logrus.New()

type P2PDkgInterface interface {
	SetGroupId([]byte)
	GetGroupId() []byte
	GetDKGIndex() int
	SetGroupMembers([][]byte)
	GetGroupCmd() chan [][]byte
	GetDkgEvent() chan int
	SetNbParticipants(int)
	GetGroupPublicPoly() *share.PubPoly
	GetShareSecurity() *share.PriShare
	IsCertified() bool
	Reset()
	SubscribeEvent(chan int)
}

func CreateP2PDkg(p p2p.P2PInterface, suite suites.Suite, peerEvent chan p2p.P2PMessage) P2PDkgInterface {
	sec := suite.Scalar().Pick(suite.RandomStream())
	d := &P2PDkg{
		suite:          suite,
		publicKeys:     Pubkeys{},
		pubkeyIdMap:    make(map[string]string),
		partSec:        sec,
		partPub:        suite.Point().Mul(sec, nil),
		groupCmd:       make(chan [][]byte),
		network:        &p,
		chPeerEvent:    peerEvent,
		currentState:   INIT,
		subscribeEvent: make(chan int),
	}
	d.ctx, d.cancel = context.WithCancel(context.Background())

	go func() {
		for groupIds := range d.groupCmd {
			d.Reset()
			go d.eventLoop(groupIds)
		}
	}()

	return d
}

func (d *P2PDkg) GetGroupCmd() chan [][]byte {
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
	responses  []Responses
	//Group member ID
	groupIds [][]byte
	//node key pair
	partPub kyber.Point
	partSec kyber.Scalar
	//
	partDkg      *DistKeyGenerator
	partDks      *DistKeyShare
	groupPubPoly *share.PubPoly
	//
	pubkeyIdMap map[string]string
	//
	nbParticipants int
	subscribeEvent chan int
	network        *p2p.P2PInterface
	chPeerEvent    chan p2p.P2PMessage

	groupingStart time.Time
	dkgIndex      int

	currentState int
	groupCmd     chan [][]byte
	ctx          context.Context
	cancel       context.CancelFunc
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

func (d *P2PDkg) Reset() {
	d.cancel()
	d.ctx, d.cancel = context.WithCancel(context.Background())
	d.publicKeys = d.publicKeys[:0]
	d.deals = d.deals[:0]
	d.responses = d.responses[:0]
	d.pubkeyIdMap = make(map[string]string)
	d.partSec = d.suite.Scalar().Pick(d.suite.RandomStream())
	d.partPub = d.suite.Point().Mul(d.partSec, nil)
	d.currentState = INIT
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

func (d *P2PDkg) eventLoop(groupIds [][]byte) {
	pubKeyCount := 0
	dealCount := 0
	responseCount := 0
	unknown := 0

	pubKeyCh := make(chan vss.PublicKey)
	defer close(pubKeyCh)
	dealCh := make(chan Deal)
	defer close(dealCh)
	respsCh := make(chan Responses)
	defer close(respsCh)

	d.pipeExchangePubKey(groupIds)
	d.pipeProcessResponses(d.pipeProcessDeal(d.pipeNewDistKeyGenerator(pubKeyCh), dealCh), respsCh)

	for {
		select {
		//event from peer
		case msg := <-d.chPeerEvent:
			switch content := msg.Msg.Message.(type) {
			case *vss.PublicKey:
				pubKeyCount++
				log.WithFields(logrus.Fields{
					"function":    "eventLoop",
					"pubKeyCount": pubKeyCount,
				}).Info()
				pubKeyCh <- *content
			case *Deal:
				dealCount++
				log.WithFields(logrus.Fields{
					"function":  "eventLoop",
					"dealCount": dealCount,
				}).Info()
				dealCh <- *content
			case *Responses:
				responseCount++
				log.WithFields(logrus.Fields{
					"function":      "eventLoop",
					"responseCount": responseCount,
				}).Info()
				respsCh <- *content
			default:
				unknown++
				fmt.Println("unknown", content)
			}
		case <-d.ctx.Done():
			fmt.Println("eventLoop closed")
			return
		}
	}
}

func (d *P2PDkg) pipeExchangePubKey(groupIds [][]byte) {
	go func() {
		if d.currentState == INIT {
			log.WithFields(logrus.Fields{
				"function":    "runDKG",
				"eventRunDKG": true,
			}).Info()
			d.groupIds = groupIds
			d.nbParticipants = len(groupIds)
			d.groupingStart = time.Now()

			//send public key to groupIds
			id := (*d.network).GetID()
			public := vss.PublicKey{SenderId: id}
			err := public.SetPoint(d.suite, d.partPub)
			if err != nil {
				log.WithField("function", "setPoint").Warn(err)
			}
			d.broadcast(&public)

			log.WithFields(logrus.Fields{
				"function":           "enterExchangePubKey",
				"eventSendAllPubKey": true,
			}).Info()
		}
		fmt.Println("pipeExchangePubKey closed")
	}()
}

func (d *P2PDkg) pipeNewDistKeyGenerator(pubKeych <-chan vss.PublicKey) (genDoneCh chan bool) {
	genDoneCh = make(chan bool)
	d.pubkeyIdMap[d.partPub.String()] = string((*d.network).GetID())
	d.publicKeys = append(d.publicKeys, d.partPub)

	go func() {
		for pubKey := range pubKeych {
			p, err := pubKey.GetPoint(d.suite)
			if err != nil {
				log.WithField("function", "getPoint").Warn(err)
			}
			d.pubkeyIdMap[p.String()] = string(pubKey.SenderId)
			d.publicKeys = append(d.publicKeys, p)

			if len(d.publicKeys) == d.nbParticipants {
				sort.Sort(d.publicKeys)
				d.partDkg, err = NewDistKeyGenerator(d.suite, d.partSec, d.publicKeys, d.nbParticipants/2+1)
				if err != nil {
					log.WithField("function", "newDistKeyGenerator").Warn(err)
				}
				deals, err := d.partDkg.Deals()
				if err != nil {
					log.WithField("function", "deals").Warn(err)
				}
				for i, pub := range d.publicKeys {
					if !pub.Equal(d.partPub) {
						err = (*d.network).SendMessage([]byte(d.pubkeyIdMap[pub.String()]), deals[i])
						if err != nil {
							log.WithField("function", "sendMessage").Warn(err)
						}
					} else {
						d.dkgIndex = i
					}
				}
				genDoneCh <- true
				log.WithFields(logrus.Fields{
					"function":         "enterNewDistKeyGenerator",
					"eventSendAllDeal": true,
				}).Info()
			}
		}
		close(genDoneCh)
		fmt.Println("pipeNewDistKeyGenerator closed")
	}()
	return
}

func (d *P2PDkg) pipeProcessDeal(genDoneCh <-chan bool, dealCh <-chan Deal) (dealDoneCh chan bool) {
	dealDoneCh = make(chan bool)
	go func() {
		for deal := range dealCh {
			d.deals = append(d.deals, deal)
			if len(d.deals) == d.nbParticipants-1 {
				if <-genDoneCh == true {
					var resps []*Response
					for _, deal := range d.deals {
						resp, err := (*d.partDkg).ProcessDeal(&deal)
						if err != nil {
							log.WithField("function", "processDeal").Warn(err)
						} else {
							resps = append(resps, resp)
						}
					}
					d.broadcast(&Responses{Response: resps})
					dealDoneCh <- true
				}
			}
		}
		close(dealDoneCh)
		fmt.Println("pipeProcessDeal closed")
	}()
	return
}

func (d *P2PDkg) pipeProcessResponses(dealDoneCh chan bool, respsCh chan Responses) {
	go func() {
		for resps := range respsCh {
			d.responses = append(d.responses, resps)
			if len(d.responses) == d.nbParticipants-1 {
				if <-dealDoneCh == true {
					for _, rs := range d.responses {
						for _, r := range rs.GetResponse() {
							_, err := (*d.partDkg).ProcessResponse(r)
							if err != nil {
								log.WithField("function", "ProcessResponse").Warn(err)
							}
						}
					}
					if (*d.partDkg).Certified() {
						var err error
						d.partDks, err = d.partDkg.DistKeyShare()
						if err != nil {
							log.WithField("function", "distKeyShare").Warn(err)
						}
						d.groupPubPoly = share.NewPubPoly(d.suite, d.suite.Point().Base(), d.partDks.Commitments())
						d.currentState = VERIFIED
						timeCost := time.Since(d.groupingStart).Seconds()
						fmt.Println("DistKeyShare SUCCESS ", timeCost)
						log.WithFields(logrus.Fields{
							"function":        "enterVerified",
							"eventDKGSucceed": true,
							"DKGTimeCost":     timeCost,
						}).Info()
						d.cancel()
						fmt.Println("pipeProcessResponses closed")
						if d.subscribeEvent != nil {
							d.subscribeEvent <- VERIFIED
						}
						return
					}
				}
			}
		}
		fmt.Println("pipeProcessResponses closed")
	}()
}

func (d *P2PDkg) broadcast(m proto.Message) {
	for _, member := range d.groupIds {
		if string(member) != string((*d.network).GetID()) {
			if err := (*d.network).SendMessage(member, m); err != nil {
			}
		}
	}
}
