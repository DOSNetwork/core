package p2p

import (
	"net"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

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
	secKey, pubKey := genPair()
	id := getLocalIp() + pubKey.String()[:15]
	p := &P2P{
		peers:       new(sync.Map),
		messageChan: tunnel,
		suite:       suite,
		secKey:      secKey,
		pubKey:      pubKey,
		id:          id,
		ip:          getLocalIp(),
		port:        port,
	}
	return p, nil
}
func getLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

type P2PMessage struct {
	Msg    ptypes.DynamicAny
	Sender string
}

type P2PInterface interface {
	// Listen starts listening for peers on a port.
	Listen() error
	Broadcast(*proto.Message)
	SendMessageById(string, *proto.Message)
	CreatePeer(string, *net.Conn) error
	GetTunnel() chan P2PMessage
}
