package node

import (
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"
	"github.com/sirupsen/logrus"
)

type TestStrategy interface {
	StartTest(*PeerNode)
	CheckResult(string, *internalMsg.Cmd, *PeerNode)
}

type node struct {
	p         p2p.P2PInterface
	peerEvent chan p2p.P2PMessage
	members   [][]byte
	allIP     []string
	peerSize  int
	done      chan bool
	log       *logrus.Entry
}
