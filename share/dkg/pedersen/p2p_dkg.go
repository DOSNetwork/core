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
)

type P2PDkgInterface interface {
	SetGroupId([]byte)
	GetGroupId() []byte
	GetDKGIndex() int
	SetGroupMembers([][]byte)
	IsCetified() bool
	RunDKG()
	Reset()
	SetNbParticipants(int)
	GetGroupPublicPoly() *share.PubPoly
	GetShareSecuirty() *share.PriShare
	EventLoop()
	Event(string)
	SubscribeEvent(chan string)
}

func CreateP2PDkg(p p2p.P2PInterface, suite suites.Suite, peerEvent chan p2p.P2PMessage, nbParticipants int) (P2PDkgInterface, error) {
	sec := suite.Scalar().Pick(suite.RandomStream())
	d := &P2PDkg{
		suite:       suite,
		publicKeys:  Pubkeys{},
		chResponse:  make(chan Response, 1),
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
			{Name: "receivePubkey", Src: []string{"Init"}, Dst: "Init"},
			{Name: "grouping", Src: []string{"Init"}, Dst: "ExchangePubKey"},
			{Name: "receivePubkey", Src: []string{"ExchangePubKey"}, Dst: "ExchangePubKey"},
			{Name: "receiveAllPubkey", Src: []string{"ExchangePubKey"}, Dst: "NewDistKeyGenerator"},
			{Name: "receiveDeal", Src: []string{"NewDistKeyGenerator"}, Dst: "NewDistKeyGenerator"},
			{Name: "receiveAllDeal", Src: []string{"NewDistKeyGenerator"}, Dst: "ProcessDeal"},
			{Name: "ReadyForResponse", Src: []string{"ProcessDeal"}, Dst: "ProcessResponse"},
			{Name: "cetified", Src: []string{"ProcessResponse"}, Dst: "Verified"},
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
	responses  []Response
	chResponse chan Response
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

func (d *P2PDkg) IsCetified() bool {
	if d.FSM.Current() == "Verified" {
		return true
	} else {
		return false
	}
}

func (d *P2PDkg) RunDKG() {
	if d.FSM.Current() == "Init" {
		d.chFsmEvent <- "grouping"
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

func (d *P2PDkg) GetShareSecuirty() *share.PriShare {
	if d.FSM.Current() == "Verified" {
		return d.partDks.Share
	}
	return nil
}
func (d *P2PDkg) SubscribeEvent(dkgEvent chan string) {
	d.subscribeEvent = dkgEvent
}
func (d *P2PDkg) EventLoop() {
	for {
		select {
		//event from FSM
		case event := <-d.chFsmEvent:
			d.Event(event)
		//event from peer
		case msg := <-d.chPeerEvent:
			switch content := msg.Msg.Message.(type) {
			case *vss.PublicKey:
				pubkey := *content
				p, err := pubkey.GetPoint(d.suite)
				if err != nil {
					fmt.Println(err)
				}
				d.pubkeyIdMap[p.String()] = string(pubkey.SenderId)
				d.publicKeys = append(d.publicKeys, p)
				d.Event("receivePubkey")
			case *Deal:
				d.deals = append(d.deals, *content)
				d.Event("receiveDeal")
			case *Response:
				d.chResponse <- *content
			default:
				fmt.Println("unknown", content)
			}
		}
	}
}

func (d *P2PDkg) enterInit(e *fsm.Event) {
	//Data Buffer
	d.publicKeys = d.publicKeys[:0]
	d.deals = d.deals[:0]
	d.chResponse = make(chan Response, 1)

	d.pubkeyIdMap = make(map[string]string)
	d.partSec = d.suite.Scalar().Pick(d.suite.RandomStream())
	d.partPub = d.suite.Point().Mul(d.partSec, nil)
}

func (d *P2PDkg) enterExchangePubKey(e *fsm.Event) {
	d.groupingStart = time.Now()
	id := (*d.network).GetId().Id
	d.pubkeyIdMap[d.partPub.String()] = string(id)
	d.publicKeys = append(d.publicKeys, d.partPub)

	//send publick key to groupIds
	public := vss.PublicKey{SenderId: id}
	err := public.SetPoint(d.suite, d.partPub)
	if err != nil {
		fmt.Println(err)
	}
	d.broadcast(&public)

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

	go func() {
		for r := range d.chResponse {
			if d.FSM.Current() != "ProcessResponse" {
				d.responses = append(d.responses, r)
			} else {
				for _, savedResponses := range d.responses {
					_, err := (*d.partDkg).ProcessResponse(&savedResponses)
					if err != nil {
						fmt.Println(err)
					}
				}
				d.responses = d.responses[:0]
				_, err := (*d.partDkg).ProcessResponse(&r)
				if err != nil {
					fmt.Println(err)
				}
				if (*d.partDkg).Certified() && d.FSM.Current() == "ProcessResponse" {
					fmt.Println("resp Certified ")
					d.chFsmEvent <- "cetified"
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
			err = (*d.network).SendMessageById([]byte(d.pubkeyIdMap[pub.String()]), deals[i])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			d.dkgIndex = i
		}
	}
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
	for _, deal := range d.deals {
		resp, err := (*d.partDkg).ProcessDeal(&deal)
		if err != nil {
			fmt.Println(err)
		} else {
			d.broadcast(resp)
		}
	}
	d.chFsmEvent <- "ReadyForResponse"
}

func (d *P2PDkg) enterVerified(e *fsm.Event) {
	close(d.chResponse)
	var err error
	d.partDks, err = d.partDkg.DistKeyShare()
	if err != nil {
		fmt.Println(err)
	}
	d.groupPubPoly = share.NewPubPoly(d.suite, d.suite.Point().Base(), d.partDks.Commitments())
	fmt.Println("DistKeyShare SUCCESS ")
	fmt.Println(time.Since(d.groupingStart))
	if d.subscribeEvent != nil {
		d.subscribeEvent <- "cetified"
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
		if string(member) != string((*d.network).GetId().Id) {
			go (*d.network).SendMessageById(member, m)
		}
	}
}
