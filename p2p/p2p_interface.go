package p2p

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

var (
	suite = suites.MustFind("bn256")
)

func genPair() (kyber.Scalar, kyber.Point) {
	secret := suite.Scalar().Pick(suite.RandomStream())
	public := suite.Point().Mul(secret, nil)
	return secret, public
}

func CreateP2PNetwork(id []byte, port string) (P2PInterface, error) {
	p := &P2P{
		replyPeers:   new(sync.Map),
		requestPeers: new(sync.Map),
		suite:        suite,
		messages:     make(chan P2PMessage, 100),
		port:         port,
		logger:       log.New("module", "p2p"),
	}
	p.identity.Id = id
	ip, err := GetLocalIP()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	p.cluster, err = SetupCluster(ip, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Done SetupCluster")
	return p, nil
}

func GetLocalIP() (ip string, err error) {
	var addrs []net.Addr

	if addrs, err = net.InterfaceAddrs(); err != nil {
		return "", err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("IP not found")
}

type P2PMessage struct {
	Msg          ptypes.DynamicAny
	Sender       []byte
	RequestNonce uint64
	PeerConn     *PeerConn
}

type P2PInterface interface {
	GetIP() string
	GetID() []byte
	Listen() error
	Join(bootstrapIp string) error
	ConnectTo(IpAddr string) (id []byte, err error)
	Leave()
	SendMessage(id []byte, msg proto.Message) error
	Members() int
	Request(id []byte, m proto.Message) (msg proto.Message, err error)
	Reply(id []byte, nonce uint64, m proto.Message) (err error)
	SubscribeEvent(chanBuffer int, messages ...interface{}) (outch chan P2PMessage, err error)
	UnSubscribeEvent(messages ...interface{})
	CloseMessagesChannel()
}
