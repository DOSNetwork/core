package p2p

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/suites"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

var logger log.Logger

func CreateP2PNetwork(id []byte, port string) (P2PInterface, error) {
	suite := suites.MustFind("bn256")
	logger = log.New("module", "p2p")
	p := &Server{
		suite:     suite,
		messages:  make(chan P2PMessage, 100),
		subscribe: make(chan Subscription),
		unscribe:  make(chan Subscription),
		port:      port,
	}
	p.ctx, p.cancel = context.WithCancel(context.Background())
	p.secKey = suite.Scalar().Pick(suite.RandomStream())
	p.pubKey = suite.Point().Mul(p.secKey, nil)
	p.id = id
	ip, err := GetLocalIP()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	p.address = ip

	p.cluster, err = SetupCluster(ip, id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
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
}

type P2PInterface interface {
	GetIP() string
	GetID() []byte
	Listen() error
	Join(bootstrapIp string) error
	ConnectTo(IpAddr string) (id []byte, err error)
	Leave()
	Members() int
	Request(id []byte, m proto.Message) (msg P2PMessage, err error)
	Reply(id []byte, nonce uint64, m proto.Message) (err error)
	SubscribeEvent(chanBuffer int, messages ...interface{}) (outch chan P2PMessage, err error)
	UnSubscribeEvent(messages ...interface{})
}
