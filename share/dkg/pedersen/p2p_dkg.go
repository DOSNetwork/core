package dkg

import (
	"fmt"
	"sort"
	"time"

	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"

	"github.com/dedis/kyber"

	"github.com/golang/protobuf/proto"

	"github.com/looplab/fsm"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

type P2PDkgInterface interface {
	SetGroupId([]byte)
	GetGroupId() []byte
	GetDKGIndex() int
	SetGroupMembers([][]byte)
	IsCertified() bool
	RunDKG()
	Reset()
	SetNbParticipants(int)
	GetGroupPublicPoly() *share.PubPoly
	GetShareSecurity() *share.PriShare
	EventLoop()
	Event(string)
	SubscribeEvent(chan string)
}

func CreateP2PDkg(p p2p.P2PInterface, suite suites.Suite, peerEvent chan p2p.P2PMessage, nbParticipants int, logger *logrus.Entry) (P2PDkgInterface, error) {
	log = logger
	sec := suite.Scalar().Pick(suite.RandomStream())
	d := &P2PDkg{
		suite:       suite,
		publicKeys:  Pubkeys{},
		chResponse:  make(chan Responses, 1),
		pubkeyIdMap: make(map[string]string),
		partSec:     sec,
		partPub:     suite.Point().Mul(sec, nil),

		nbParticipants: nbParticipants,
		network:        &p,
		chFsmEvent:     make(chan string, 1),
		chPeerEvent:    peerEvent,
	}
	d.FSM = fsm.NewFSM(
		"Init",
		fsm.Events{
			{Name: "reset", Src: []string{"Init"}, Dst: "Init"},
			{Name: "grouping", Src: []string{"Init"}, Dst: "ExchangePubKey"},
			{Name: "receivePubkey", Src: []string{"Init"}, Dst: "Init"},
			{Name: "receivePubkey", Src: []string{"ExchangePubKey"}, Dst: "ExchangePubKey"},
			{Name: "receiveAllPubkey", Src: []string{"ExchangePubKey"}, Dst: "NewDistKeyGenerator"},
			{Name: "receiveDeal", Src: []string{"ExchangePubKey"}, Dst: "ExchangePubKey"},
			{Name: "receiveDeal", Src: []string{"NewDistKeyGenerator"}, Dst: "NewDistKeyGenerator"},
			{Name: "receiveAllDeal", Src: []string{"NewDistKeyGenerator"}, Dst: "ProcessDeal"},
			{Name: "certified", Src: []string{"ProcessDeal"}, Dst: "Verified"},
			{Name: "checkURL", Src: []string{"Verified"}, Dst: "Verified"},
			{Name: "receiveSignature", Src: []string{"Verified"}, Dst: "Verified"},
			{Name: "verify", Src: []string{"Verified"}, Dst: "Verified"},
			{Name: "reset", Src: []string{"Verified"}, Dst: "Init"},
		},
		fsm.Callbacks{
			"enter_Init":                func(e *fsm.Event) { d.enterInit(e) },
			"enter_ExchangePubKey":      func(e *fsm.Event) { d.enterExchangePubKey(e) },
			"after_receivePubkey":       func(e *fsm.Event) { d.afterReceivePubkey(e) },
			"enter_NewDistKeyGenerator": func(e *fsm.Event) { d.enterNewDistKeyGenerator(e) },
			"after_receiveDeal":         func(e *fsm.Event) { d.afterReceiveDeal(e) },
			"enter_ProcessDeal":         func(e *fsm.Event) { d.enterProcessDeal(e) },
			"enter_Verified":            func(e *fsm.Event) { d.enterVerified(e) },
		},
	)
	return d, nil
}

type P2PDkg struct {
	groupId []byte
	suite   suites.Suite
	//Data Buffer
	publicKeys Pubkeys
	deals      []Deal
	chResponse chan Responses
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
	subscribeEvent chan string
	FSM            *fsm.FSM
	chFsmEvent     chan string
	network        *p2p.P2PInterface
	chPeerEvent    chan p2p.P2PMessage

	groupingStart time.Time
	dkgIndex      int
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
	if d.FSM.Current() == "Init" {
		d.nbParticipants = n
	}
}
func (d *P2PDkg) SetGroupMembers(members [][]byte) {
	if d.FSM.Current() == "Init" {
		d.groupIds = members
	}
}

func (d *P2PDkg) IsCertified() bool {
	if d.FSM.Current() == "Verified" {
		return true
	} else {
		return false
	}
}

func (d *P2PDkg) RunDKG() {
	if d.FSM.Current() == "Init" {
		d.chFsmEvent <- "grouping"
		log.WithFields(logrus.Fields{
			"eventRunDKG": true,
		}).Info()
	}
}

func (d *P2PDkg) Reset() {
	d.chFsmEvent <- "reset"
}

func (d *P2PDkg) GetGroupPublicPoly() *share.PubPoly {
	if d.FSM.Current() == "Verified" {
		return d.groupPubPoly
	}
	return nil
}

func (d *P2PDkg) GetShareSecurity() *share.PriShare {
	if d.FSM.Current() == "Verified" {
		return d.partDks.Share
	}
	return nil
}
func (d *P2PDkg) SubscribeEvent(dkgEvent chan string) {
	d.subscribeEvent = dkgEvent
}
func (d *P2PDkg) EventLoop() {
	pubKeyCount := 0
	dealCount := 0
	responseCount := 0
	unknown := 0
	for {
		select {
		//event from FSM
		case event := <-d.chFsmEvent:
			d.Event(event)
		//event from peer
		case msg := <-d.chPeerEvent:
			switch content := msg.Msg.Message.(type) {
			case *vss.PublicKey:
				pubKeyCount++
				log.WithFields(logrus.Fields{
					"pubKeyCount": pubKeyCount,
				}).Info()
				pubkey := *content
				p, err := pubkey.GetPoint(d.suite)
				if err != nil {
					fmt.Println(err)
				}
				d.pubkeyIdMap[p.String()] = string(pubkey.SenderId)
				d.publicKeys = append(d.publicKeys, p)
				d.Event("receivePubkey")
			case *Deal:
				dealCount++
				log.WithFields(logrus.Fields{
					"dealCount": dealCount,
				}).Info()
				d.deals = append(d.deals, *content)
				d.Event("receiveDeal")
			case *Responses:
				responseCount++
				log.WithFields(logrus.Fields{
					"responseCount": responseCount,
				}).Info()
				d.chResponse <- *content
			default:
				unknown++
				log.WithFields(logrus.Fields{
					"unknown": unknown,
				}).Info()
				fmt.Println("unknown", content)
			}
		}
	}
}

func (d *P2PDkg) enterInit(e *fsm.Event) {
	//Data Buffer
	d.publicKeys = d.publicKeys[:0]
	d.deals = d.deals[:0]
	d.chResponse = make(chan Responses, 1)

	d.pubkeyIdMap = make(map[string]string)
	d.partSec = d.suite.Scalar().Pick(d.suite.RandomStream())
	d.partPub = d.suite.Point().Mul(d.partSec, nil)
}

func (d *P2PDkg) enterExchangePubKey(e *fsm.Event) {
	d.groupingStart = time.Now()
	id := (*d.network).GetID()
	d.pubkeyIdMap[d.partPub.String()] = string(id)
	d.publicKeys = append(d.publicKeys, d.partPub)

	//send public key to groupIds
	public := vss.PublicKey{SenderId: id}
	err := public.SetPoint(d.suite, d.partPub)
	if err != nil {
		fmt.Println(err)
	}
	d.broadcast(&public)

	log.WithFields(logrus.Fields{
		"eventSendAllPubKey": true,
	}).Info()

	//TODO: Should check if publicKeys are all belong to the members
	if len(d.publicKeys) == d.nbParticipants {
		d.chFsmEvent <- "receiveAllPubkey"
	}
}

func (d *P2PDkg) afterReceivePubkey(e *fsm.Event) {
	//TODO: Should check if publicKeys are all belong to the members
	if len(d.publicKeys) == d.nbParticipants {
		d.chFsmEvent <- "receiveAllPubkey"
	}
}

//Can't call NewDistKeyGenerator .need to wait unitl all deal has been received
func (d *P2PDkg) enterNewDistKeyGenerator(e *fsm.Event) {
	log.WithFields(logrus.Fields{
		"eventReceiveAllPubKey": true,
	}).Info()

	go func() {
		for rs := range d.chResponse {
			for _, r := range rs.GetResponse() {
				_, err := (*d.partDkg).ProcessResponse(r)
				if err != nil {
					fmt.Println(err)
				}
				if (*d.partDkg).Certified() && d.FSM.Current() == "ProcessDeal" {
					fmt.Println("resp Certified ")
					d.chFsmEvent <- "certified"
				}
			}
		}
	}()

	var err error
	sort.Sort(d.publicKeys)
	d.partDkg, err = NewDistKeyGenerator(d.suite, d.partSec, d.publicKeys, d.nbParticipants/2+1)
	if err != nil {
		fmt.Println(err)
	}
	deals, err := d.partDkg.Deals()
	if err != nil {
		fmt.Println(err)
	}
	for i, pub := range d.publicKeys {
		if !pub.Equal(d.partPub) {
			err = (*d.network).SendMessage([]byte(d.pubkeyIdMap[pub.String()]), deals[i])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			d.dkgIndex = i
		}
	}

	log.WithFields(logrus.Fields{
		"eventSendAllDeal": true,
	}).Info()

	condition := d.nbParticipants - 1
	if len(d.deals) == condition {
		d.chFsmEvent <- "receiveAllDeal"
	}
}
func (d *P2PDkg) afterReceiveDeal(e *fsm.Event) {
	condition := d.nbParticipants - 1
	if len(d.deals) == condition {
		d.chFsmEvent <- "receiveAllDeal"
	}
}

//This is only happened after all peers has all public keys
func (d *P2PDkg) enterProcessDeal(e *fsm.Event) {
	log.WithFields(logrus.Fields{
		"eventReceiveAllDeal": true,
	}).Info()

	var resps []*Response

	for _, deal := range d.deals {
		resp, err := (*d.partDkg).ProcessDeal(&deal)
		if err != nil {
			fmt.Println(err)
		} else {
			resps = append(resps, resp)
		}
	}

	d.broadcast(&Responses{Response: resps})

	if (*d.partDkg).Certified() {
		fmt.Println("resp Certified ")
		d.chFsmEvent <- "certified"
	}

	log.WithFields(logrus.Fields{
		"eventProcessAllDeal": true,
	}).Info()
}

func (d *P2PDkg) enterVerified(e *fsm.Event) {
	log.WithFields(logrus.Fields{
		"eventProcessAllResponse": true,
	}).Info()

	close(d.chResponse)
	var err error
	d.partDks, err = d.partDkg.DistKeyShare()
	if err != nil {
		fmt.Println(err)
	}
	d.groupPubPoly = share.NewPubPoly(d.suite, d.suite.Point().Base(), d.partDks.Commitments())
	timeCost := time.Since(d.groupingStart).Seconds()
	fmt.Println("DistKeyShare SUCCESS ", timeCost)
	log.WithFields(logrus.Fields{
		"eventDKGSucceed": true,
		"DKGTimeCost":     timeCost,
	}).Info()
	if d.subscribeEvent != nil {
		d.subscribeEvent <- "certified"
	}
}

func (d *P2PDkg) Event(event string) {
	err := d.FSM.Event(event)
	if err != nil {
		fmt.Println(err)
	}
}

func (d *P2PDkg) broadcast(m proto.Message) {
	for _, member := range d.groupIds {
		if string(member) != string((*d.network).GetID()) {
			if err := (*d.network).SendMessage(member, m); err != nil {
				fmt.Println("DKG SendMessage err ", err)
			}
		}
	}
}
