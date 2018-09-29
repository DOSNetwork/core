package p2p

import (
	"errors"
	"net"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/suites"

	"github.com/dedis/kyber"
)

var suite = suites.MustFind("bn256")

func genPair() (kyber.Scalar, kyber.Point) {
	secret := suite.Scalar().Pick(suite.RandomStream())
	public := suite.Point().Mul(secret, nil)
	return secret, public
}

func CreateP2PNetwork(tunnel chan P2PMessage, port int) (P2PInterface, error) {
	p := &P2P{
		peers:			new(sync.Map),
		messageChan:	tunnel,
		suite:       	suite,
		port:        	port,
	}
	return p, nil
}

func getLocalIp() (string, error) {
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
	// Listen starts listening for peers on a port.
	GetId() dht.ID
	Listen() error
	Broadcast(proto.Message)
	SendMessageById([]byte, proto.Message)
	CreatePeer(string, *net.Conn)
	GetTunnel() chan P2PMessage
	FindNodeById(id []byte) []dht.ID
	FindNode(targetID dht.ID, alpha int, disjointPaths int) (results []dht.ID)
	GetRoutingTable() *dht.RoutingTable
}
