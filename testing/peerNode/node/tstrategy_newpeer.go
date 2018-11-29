package node

import (
	"bytes"
	"fmt"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"github.com/golang/protobuf/proto"
)

type test1 struct{}

func (r test1) StartTest(content *internalMsg.Cmd, d *PeerNode) {
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_SIGNIN,
		Args:  []byte{},
	}
	pb := proto.Message(cmd)
	for i := 0; i < len(d.nodeIPs); i++ {
		ip := d.nodeIPs[i]
		id, _ := d.p.NewPeer(ip)
		d.checkroll[string(id)] = 0
	}

	for i := 0; i < d.numMessages; i++ {
		for id, _ := range d.checkroll {
			go d.p.SendMessageById([]byte(id), pb)
		}
	}
}

func (r test1) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {
	if content.Ctype == internalMsg.Cmd_SIGNIN {
		d.checkroll[sender] = d.checkroll[sender] + 1
		fmt.Println("sender", []byte(sender), "  ", d.checkroll[sender], content.Ctype, " ", len(d.checkroll))
		if d.checkroll[sender] == d.numMessages {
			delete(d.checkroll, sender)
		}

		if len(d.checkroll) == 0 {
			d.MakeRequest(d.bootStrapIp, "post", d.nodeID)
		}
	}
}

type test2 struct{}

func (r test2) StartTest(content *internalMsg.Cmd, d *PeerNode) {
	hostid := d.p.GetId()
	fmt.Println("!!!!StartTest ", hostid.GetId())

	for i := 0; i < len(d.nodeIDs); i++ {
		if !bytes.Equal(hostid.GetId(), d.nodeIDs[i]) {
			d.checkroll[string(d.nodeIDs[i])] = 0
		}
	}

	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_SIGNIN,
		Args:  []byte{},
	}
	pb := proto.Message(cmd)

	for i := 0; i < d.numMessages; i++ {
		for id, _ := range d.checkroll {
			go d.p.SendMessageById([]byte(id), pb)
		}
	}
}

func (r test2) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {
	if content.Ctype == internalMsg.Cmd_SIGNIN {
		d.checkroll[sender] = d.checkroll[sender] + 1
		fmt.Println("sender", []byte(sender), "  ", d.checkroll[sender], content.Ctype, " ", len(d.checkroll))
		if d.checkroll[sender] == d.numMessages {
			delete(d.checkroll, sender)
		}

		if len(d.checkroll) == 0 {
			d.MakeRequest(d.bootStrapIp, "post", d.nodeID)
		}
	}
}
