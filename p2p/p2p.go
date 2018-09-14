package p2p

import (
	"bufio"
	"fmt"
	"github.com/DOSNetwork/core/p2p/nat"
	"log"
	"net"
	"strconv"

	"sync"

	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/golang/protobuf/proto"
)

type P2P struct {
	//Map of connection addresses (string) <-> *p2p.PeerClient
	peers *sync.Map
	// Channels are thread safe
	messageChan chan P2PMessage
	suite       suites.Suite
	id          string
	ip          string
	secKey      kyber.Scalar
	pubKey      kyber.Point
}

func (n *P2P) Listen() error {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}

	isPrivateIp, err := nat.IsPrivateIp()
	if err != nil {
		return err
	}

	if isPrivateIp {
		externalPort := nat.RandomPort()
		nat, err := nat.SetMapping("tcp", externalPort, listener.Addr().(*net.TCPAddr).Port, "DosNode")
		if err != nil {
			return err
		}

		externalIp, err := nat.GetExternalAddress()
		if err != nil {
			return err
		}

		n.ip = externalIp.String() + ":" + strconv.Itoa(externalPort)
	} else {
		n.ip = n.ip + ":" + strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	}
	n.id = n.ip + n.pubKey.String()[:15]


	fmt.Println("listen ", n.ip)
	// Handle new clients.

	go func() {
		for {
			if conn, err := listener.Accept(); err == nil {
				//Create a peer client
				err := n.CreatePeer("", &conn)
				if err != nil {
					log.Fatal(err)
				}

			} else {
				fmt.Println("Failed accepting a connection request:", err)
				log.Fatal(err)
			}
		}
	}()

	return nil
}

func (n *P2P) Broadcast(m *proto.Message) {
	n.peers.Range(func(key, value interface{}) bool {
		ip := key.(string)
		client := value.(*PeerClient)
		fmt.Printf("key[%s]\n", ip)
		client.SendPackage(m)
		return true
	})
}

func (n *P2P) SendMessageById(id string, m *proto.Message) {
	value, loaded := n.peers.Load(id)
	if loaded {
		client := value.(*PeerClient)
		client.SendPackage(m)
	}
}

func (n *P2P) GetTunnel() chan P2PMessage {
	return n.messageChan
}

func (n *P2P) CreatePeer(addr string, c *net.Conn) error {
	peer := &PeerClient{
		conn:   c,
		p2pnet: n,
	}
	if addr != "" {
		peer.Dial(addr)
	}
	peer.rw = bufio.NewReadWriter(bufio.NewReader(*peer.conn), bufio.NewWriter(*peer.conn))
	//n.peers.LoadOrStore(peer.id, peer)
	peer.messageChan = n.messageChan

	//fmt.Println("InitClient id ", peer.id)
	go peer.HandlePackages()
	peer.SayHi()
	return nil
}
