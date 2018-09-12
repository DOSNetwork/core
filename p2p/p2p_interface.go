package p2p

import (
	"net"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func CreateP2PNetwork(tunnel chan P2PMessage) (P2PInterface, error) {
	p := &P2P{
		peers:       new(sync.Map),
		messageChan: tunnel,
	}
	return p, nil
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
