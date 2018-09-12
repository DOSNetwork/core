package p2p

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

type P2P struct {
	Port int
	//Map of connection addresses (string) <-> *p2p.PeerClient
	Peers *sync.Map
	// Channels are thread safe
	MessageChan chan []byte
	Kill        chan struct{}
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
	Port := fmt.Sprintf(":%v", n.Port)
	listener, err = net.Listen("tcp", Port)
	if err != nil {
		return err
	}
	fmt.Println("Listen on", n.getLocalIp(), "", n.Port)

	// Handle new clients.
	for {
		if conn, err := listener.Accept(); err == nil {
			//Create a peer client
			peer, err := n.CreatePeer("", &conn)
			if err != nil {
				return err
			}
			go peer.HandleMessages()

		} else {
			// if the Shutdown flag is set, no need to continue with the for loop
			select {
			case <-n.Kill:
				fmt.Println("Shutting down server")
				return nil
			default:
				fmt.Println("Failed accepting a connection request:", err)
				return err
			}
		}
	}
}

func (n *P2P) CreatePeer(addr string, c *net.Conn) (*PeerClient, error) {

	client := &PeerClient{
		conn: c,
	}
	if addr != "" {
		client.Dial(addr)
		client.id = addr
	} else {
		conn := *c
		client.id = conn.RemoteAddr().String()
	}
	client.rw = bufio.NewReadWriter(bufio.NewReader(*client.conn), bufio.NewWriter(*client.conn))
	n.Peers.LoadOrStore(client.id, client)
	client.MessageChan = n.MessageChan
	fmt.Println("InitClient id ", client.id)
	return client, nil
}
