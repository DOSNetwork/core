package node

import (
	"time"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"github.com/DOSNetwork/core/p2p"
	"github.com/sirupsen/logrus"
)

type TestStrategy interface {
	StartTest(*internalMsg.Cmd, *PeerNode)
	CheckResult(string, *internalMsg.Cmd, *PeerNode)
}

type node struct {
	p         p2p.P2PInterface
	peerEvent chan p2p.P2PMessage
	members   [][]byte
	allIP     []string
	peerSize  int
	done      chan bool
	log       *logrus.Logger
}

func (n *node) EventLoop() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			//PrintMemUsage()
		//event from peer
		case _ = <-n.peerEvent:
		case <-n.done:
			//os.Exit(0)
			break
		}
	}
}
