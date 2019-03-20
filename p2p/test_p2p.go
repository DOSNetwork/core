package p2p

import (
	"errors"
	"fmt"
	"net"
	"time"
)

type TestP2P struct {
	P2P
	testStrategy string
}

/*
This is a block call
*/

func (n *TestP2P) NewPeer(addr string) (id []byte, err error) {
	retryCount := 0
retry:
	retryCount++
	if retryCount >= 10 {
		err = errors.New("Peer : retried over 10 times")
		return
	}
	var conn net.Conn
	//1)Check if this address has been in peers map
	existing := false
	n.requestPeers.Range(func(key, value interface{}) bool {
		client := value.(*PeerConn)
		if client.identity.Address == addr {

			existing = true
			id = make([]byte, len(client.identity.Id))
			copy(id, client.identity.Id)
		}
		return true
	})
	if existing {
		return
	}

	//2)Dial to peer to get a connection
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Dial err ", err)
		time.Sleep(1 * time.Second)
		goto retry
	}

	_, err = NewTestPeerConn(n, &conn, n.messages)
	if err != nil {
		fmt.Println("NewPeerConn err ", err)
		goto retry
	}

	n.requestPeers.Range(func(key, value interface{}) bool {
		client := value.(*PeerConn)
		if client.identity.Address == addr {
			existing = true
			id = make([]byte, len(client.identity.Id))
			copy(id, client.identity.Id)
		}
		return true
	})
	return
}
