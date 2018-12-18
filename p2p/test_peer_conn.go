package p2p

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/DOSNetwork/core/p2p/internal"
)

type TestPeerConn struct {
	PeerConn
	testStrategy string
}

func NewTestPeerConn(p2pnet *TestP2P, conn *net.Conn, rxMessage chan P2PMessage) (peer *TestPeerConn, err error) {
	//fmt.Println("TestPeerConn receiveLoop !!! ", p.testStrategy)
	peer = &TestPeerConn{
		PeerConn{
			p2pnet:    &p2pnet.P2P,
			conn:      conn,
			rxMessage: rxMessage,
			waitForHi: make(chan bool, 2),
			rw:        bufio.NewReadWriter(bufio.NewReader(*conn), bufio.NewWriter(*conn)),
		},
		p2pnet.testStrategy,
	}
	fmt.Println("!!!! NewTestPeerConn ", 10)
	go peer.receiveLoop()
	if err = peer.sayHi(); err != nil {
		return
	}

	if err = peer.heardHi(); err != nil {
		return
	}

	go peer.heartBeat()
	return
}

func (p *TestPeerConn) receiveLoop() {
	if p.testStrategy == "DELAY_BEFORE_RECEIVELOOP" {
		time.Sleep(10 * time.Second)
	}
	//fmt.Println("TestPeerConn receiveLoop !!! ", p.testStrategy)
	p.PeerConn.receiveLoop()
}

func (p *TestPeerConn) sayHi() (err error) {
	fmt.Println("!!!! TestPeerConn sayHi")
	pa := &internal.Hi{
		PublicKey: p.p2pnet.identity.PublicKey,
		Address:   p.p2pnet.identity.Address,
		Id:        p.p2pnet.identity.Id,
	}

	err = p.SendMessage(pa)
	return
}

//Add a timer to avoid wait for Hi forever
func (p *TestPeerConn) heardHi() (err error) {
	timer := time.NewTimer(10 * time.Second)

	select {
	case <-timer.C:
		fmt.Println("Time expire")
		err = errors.New("PeerConn: Waiting for hi time expire")
	case <-p.waitForHi:
		timer.Stop()
	}

	close(p.waitForHi)
	return
}
