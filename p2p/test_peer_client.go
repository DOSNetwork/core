package p2p

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/DOSNetwork/core/p2p/internal"
	"github.com/golang/protobuf/proto"
)

type TestPeerClient struct {
	PeerClient
	testStrategy string
}

const (
	DELAY_BEFORE_RECEIVELOOP = "DELAY_BEFORE_RECEIVELOOP"
)

const TESTTIMEOUTFORHI = 10

func NewTestPeerClient(p2pnet *TestP2P, conn *net.Conn, rxMessage chan P2PMessage) (peer *TestPeerClient, err error) {
	//fmt.Println("TestPeerClient receiveLoop !!! ", p.testStrategy)
	peer = &TestPeerClient{
		PeerClient{
			p2pnet:    &p2pnet.P2P,
			conn:      conn,
			rxMessage: rxMessage,
			txMessage: make(chan proto.Message, 100),
			waitForHi: make(chan bool, 2),
			done:      make(chan bool, 2),
			rw:        bufio.NewReadWriter(bufio.NewReader(*conn), bufio.NewWriter(*conn)),
		},
		p2pnet.testStrategy,
	}
	fmt.Println("!!!! NewTestPeerClient ", TESTTIMEOUTFORHI)
	go peer.receiveLoop()
	err = peer.sayHi()
	if err != nil {
		close(peer.txMessage)
		peer = nil
	}
	return
}

func (p *TestPeerClient) receiveLoop() {
	if p.testStrategy == DELAY_BEFORE_RECEIVELOOP {
		time.Sleep(TESTTIMEOUTFORHI * time.Second)
	}
	//fmt.Println("TestPeerClient receiveLoop !!! ", p.testStrategy)
	p.PeerClient.receiveLoop()
}

func (p *TestPeerClient) sayHi() (err error) {
	fmt.Println("!!!! TestPeerClient sayHi")
	pa := &internal.Hi{
		PublicKey: p.p2pnet.identity.PublicKey,
		Address:   p.p2pnet.identity.Address,
		Id:        p.p2pnet.identity.Id,
	}

	err = p.SendMessage(pa)

	//Add a timer to avoid wait for Hi forever
	timer := time.NewTimer(TESTTIMEOUTFORHI * time.Second)
L:
	for {
		select {
		case <-timer.C:
			p.done <- true
			fmt.Println("Time expire")
			err = errors.New("PeerClient: Time expire")
			break L
		case <-p.waitForHi:
			_ = timer.Stop()
			break L
		}
	}

	return
}
