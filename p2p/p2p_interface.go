package p2p

import (
	"net"
)

type P2PInterface interface {
	// Listen starts listening for peers on a port.
	Listen()
	CreatePeer(string, *net.Conn)
	getLocalIp()
}
