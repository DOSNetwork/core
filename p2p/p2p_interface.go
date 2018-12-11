package p2p

import (
	"errors"
	"net"
	"os"
	"sync"

	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
)

var suite = suites.MustFind("bn256")

func genPair() (kyber.Scalar, kyber.Point) {
	secret := suite.Scalar().Pick(suite.RandomStream())
	public := suite.Point().Mul(secret, nil)
	return secret, public
}

func CreateP2PNetwork(id []byte, port int, logger *logrus.Logger) (P2PInterface, chan P2PMessage, error) {
	testStrategy := os.Getenv("TESTSTRATEGY")
	if testStrategy == "DELAY_BEFORE_RECEIVELOOP" {
		p := &TestP2P{
			P2P{
				peers:    new(sync.Map),
				suite:    suite,
				messages: make(chan P2PMessage, 100),
				port:     port,
				log:      logger,
			},
			testStrategy,
		}
		return p, p.messages, nil
	} else {
		p := &P2P{
			peers:    new(sync.Map),
			suite:    suite,
			messages: make(chan P2PMessage, 100),
			port:     port,
			log:      logger,
		}
		p.identity.Id = id
		return p, p.messages, nil
	}

}

func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("ip not found")
}

type P2PMessage struct {
	Msg    ptypes.DynamicAny
	Sender []byte
}

type P2PInterface interface {
	GetIP() string
	GetID() []byte
	Listen() error
	Broadcast(proto.Message)
	Join(bootstrapIp string) error
	Leave()
	SendMessage(id []byte, msg proto.Message) error
}
