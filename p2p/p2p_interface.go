package p2p

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p/discover"
	"github.com/DOSNetwork/core/p2p/nat"
	"github.com/DOSNetwork/core/suites"

	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

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
	GetPort() string
	Listen() error
	Join(bootstrapIp []string) (num int, err error)
	DisConnectTo(id []byte) error
	Leave()
	Request(ctx context.Context, id []byte, m proto.Message) (msg P2PMessage, err error)
	Reply(ctx context.Context, id []byte, nonce uint64, m proto.Message) (err error)
	SubscribeEvent() (subID int, outch chan discover.P2PEvent, err error)
	UnSubscribeEvent(int)
	SubscribeMsg(chanBuffer int, messages ...interface{}) (outch chan P2PMessage, err error)
	UnSubscribeMsg(messages ...interface{})
	NumOfMembers() int
	MembersID() [][]byte
	RandomPeerIP() []string
	//ConnectToAll(ctx context.Context, groupIds [][]byte, sessionID string) (out chan bool, errc chan error)
	numOfClient() (int, int)
}

// CreateP2PNetwork creates a P2PInterface implementation , gets a public IP and generates a secret key
func CreateP2PNetwork(id []byte, ip, port string, netType int) (P2PInterface, error) {
	suite := suites.MustFind("bn256")

	p := &server{
		suite:           suite,
		addIncomingC:    make(chan *client),
		removeIncomingC: make(chan []byte),
		replying:        make(chan p2pRequest),
		calling:         make(chan p2pRequest),
		removeCallingC:  make(chan []byte),
		peersFeed:       make(chan P2PMessage, 5),
		peersEvent:      make(chan discover.P2PEvent, 5),
		subscribeMsg:    make(chan *subscription),
		unscribeMsg:     make(chan string),
		subscribeEvent:  make(chan *subscription),
		unscribeEvent:   make(chan int),
		port:            port,
		logger:          log.New("module", "p2p"),
	}
	p.ctx, p.cancel = context.WithCancel(context.Background())
	p.secKey = suite.Scalar().Pick(suite.RandomStream())
	p.pubKey = suite.Point().Mul(p.secKey, nil)
	p.id = id

	// If user specify a public ip from the env variable,use it as external IP.
	if addr := net.ParseIP(ip); addr != nil {
		p.addr = addr
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
		network, err := discover.NewSerfNet(p.addr, string(p.id), p.port)
		if err != nil {
			p.cancel()
			return nil, err
		}
		p.members = network
	default:
		network, _ := discover.NewSimulator()
		p.members = network
	}
	return p, nil
}
