package p2p

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
)

type P2P struct {
	//Map of connection addresses (string) <-> *p2p.PeerClient
	peers *sync.Map
	// Channels are thread safe
	messageChan chan P2PMessage
}

func (n *P2P) getLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				fmt.Println("Your IP is:", ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func (n *P2P) Listen() error {
	var err error
	var listener net.Listener
	listener, err = net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	fmt.Println("Listen on ", n.getLocalIp(), ":", listener.Addr())

	// Handle new clients.
	for {
		if conn, err := listener.Accept(); err == nil {
			//Create a peer client
			err := n.CreatePeer("", &conn)
			if err != nil {
				return err
			}

		} else {
			fmt.Println("Failed accepting a connection request:", err)
			return err
		}
	}
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
		conn: c,
	}
	if addr != "" {
		peer.Dial(addr)
		peer.id = addr
	} else {
		conn := *c
		peer.id = conn.RemoteAddr().String()
	}
	peer.rw = bufio.NewReadWriter(bufio.NewReader(*peer.conn), bufio.NewWriter(*peer.conn))
	n.peers.LoadOrStore(peer.id, peer)
	peer.messageChan = n.messageChan
	fmt.Println("InitClient id ", peer.id)
	go peer.HandlePackages()
	return nil
}
