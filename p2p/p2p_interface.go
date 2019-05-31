package p2p

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p/discover"
	"github.com/DOSNetwork/core/p2p/nat"
	"github.com/DOSNetwork/core/suites"

	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

var logger log.Logger

const (
	//NoDiscover means that p2p don't use any discover protocol
	NoDiscover = iota // 0
	//GossipDiscover means that p2p use gossip as a discover protocol
	GossipDiscover
	swimPort = 7946
)

// P2PMessage is a struct that includes message,sender and nonce
type P2PMessage struct {
	Msg          ptypes.DynamicAny
	Sender       []byte
	RequestNonce uint64
}

// P2PInterface represents a p2p network
type P2PInterface interface {
	GetIP() net.IP
	GetID() []byte
	SetPort(port string)
	Listen() error
	Join(bootstrapIp []string) (num int, err error)
	ConnectTo(ip string, id []byte) ([]byte, error)
	DisConnectTo(id []byte) error
	Leave()
	Request(id []byte, m proto.Message) (msg P2PMessage, err error)
	Reply(id []byte, nonce uint64, m proto.Message) (err error)
	SubscribeEvent(chanBuffer int, messages ...interface{}) (outch chan P2PMessage, err error)
	UnSubscribeEvent(messages ...interface{})
	Members() int
	ConnectToAll() (memNum, connNum int)
	numOfClient() (int, int)
}

// CreateP2PNetwork creates a P2PInterface implementation , gets a public IP and generates a secret key
func CreateP2PNetwork(id []byte, port string, netType int) (P2PInterface, error) {
	suite := suites.MustFind("bn256")
	logger = log.New("module", "p2p")
	p := &server{
		suite:     suite,
		messages:  make(chan P2PMessage, 100),
		subscribe: make(chan subscription),
		unscribe:  make(chan subscription),
		port:      port,
	}
	p.ctx, p.cancel = context.WithCancel(context.Background())
	p.secKey = suite.Scalar().Pick(suite.RandomStream())
	p.pubKey = suite.Point().Mul(p.secKey, nil)
	p.id = id

	// If user specify a public ip from the env variable,use it as external IP.
	if ip := net.ParseIP(os.Getenv("PUBLICIP")); ip != nil {
		p.addr = ip
	} else {
		ip, err := getIP()
		if err != nil {
			return nil, err
		}
		if netutil.IsLAN(ip) {
			natdev, err := nat.DiscoverGateway()
			if err != nil {
				return nil, err
			}

			externalIp, err := natdev.GetExternalAddress()
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			if netutil.IsLAN(externalIp) {
				return nil, errors.New("NAT IP is a local IP Address")
			}

			portInt, err := strconv.Atoi(port)
			if err != nil {
				return nil, err
			}

			if err := nat.SetMapping(p.ctx, natdev, "tcp", portInt, portInt, "DosClient"); err != nil {
				fmt.Println(err)
				return nil, err
			}

			if netType == GossipDiscover {
				if err := nat.SetMapping(p.ctx, natdev, "tcp", swimPort, swimPort, "DosGossip"); err != nil {
					fmt.Println(err)
					return nil, err
				}
			}
			// If NAT port mapping success, then return NAT external IP.
			p.addr = externalIp
		} else {
			// If the local IP is a public IP, then return it as a external IP.
			p.addr = ip
		}
	}

	switch netType {
	case GossipDiscover:
		network, err := discover.NewSerfNet(p.addr, p.id)
		if err != nil {
			p.cancel()
			return nil, err
		}
		p.members = network
	default:
	}
	return p, nil
}
