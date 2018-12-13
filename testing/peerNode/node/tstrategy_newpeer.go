package node

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"github.com/golang/protobuf/proto"
)

type test1 struct{}

func (r test1) StartTest(d *PeerNode) {
	if d.p.GetIP() == d.nodeIPs[0] {
		fmt.Println("\nStartTest ", d.p.GetID(), " ", d.p.GetIP())
		cmd := &internalMsg.Cmd{
			Ctype: internalMsg.Cmd_SIGNIN,
			Args:  []byte{},
		}
		pb := proto.Message(cmd)
		for i := 1; i < len(d.nodeIPs); i++ {
			if d.p.GetIP() != d.nodeIPs[i] {
				ip := d.nodeIPs[i]
				id, err := d.p.ConnectTo(ip)
				if err != nil {
					fmt.Println("NewPeer err", err)
				}
				fmt.Println("ConnectTo To", id, " ", ip)
				d.checkroll[string(id)] = 0
			}
		}
		for i := 0; i < d.numMessages; i++ {
			for id, _ := range d.checkroll {
				if err := d.p.SendMessage([]byte(id), pb); err != nil {
					for err != nil {
						err = d.p.SendMessage([]byte(id), pb)
					}
				}
			}
		}
	}
}

func (r test1) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {
	if d.p.GetIP() == d.nodeIPs[0] {
		if content.Ctype == internalMsg.Cmd_SIGNIN {
			d.checkroll[sender] = d.checkroll[sender] + 1
			if d.checkroll[sender] == d.numMessages {
				delete(d.checkroll, sender)
			}
			d.log.WithField("event", "peer remain").Info(len(d.checkroll))
			if len(d.checkroll) == 0 {
				fmt.Println("test done")
				d.MakeRequest(d.bootStrapIp, "post", d.nodeID)
			} else {
				fmt.Println("wait for  = ", len(d.checkroll))
				for id, _ := range d.checkroll {
					fmt.Println("wait for ", []byte(id))
				}
				fmt.Println("==================== ")
			}
		}
	} else {
		cmd := &internalMsg.Cmd{
			Ctype: internalMsg.Cmd_SIGNIN,
			Args:  []byte{},
		}
		pb := proto.Message(cmd)
		if err := d.p.SendMessage([]byte(sender), pb); err != nil {
			for err != nil {
				err = d.p.SendMessage([]byte(sender), pb)
			}
		}
		d.MakeRequest(d.bootStrapIp, "post", d.nodeID)
	}
}

type test2 struct{}

func (r test2) StartTest(d *PeerNode) {

	for i := 0; i < len(d.nodeIDs); i++ {
		if !bytes.Equal(d.p.GetID(), d.nodeIDs[i]) {
			d.checkroll[string(d.nodeIDs[i])] = 0
		}
	}
	/*
		for i := 0; i < len(d.nodeIPs[i]); i++ {
			if d.p.GetIP() != d.nodeIPs[i] {
				fmt.Println(d.p.GetIP(), " join ", d.nodeIPs[i])
				d.p.Join(d.nodeIPs[i])
				break
			}
		}
	*/
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_SIGNIN,
		Args:  []byte{},
	}
	pb := proto.Message(cmd)

	for i := 0; i < d.numMessages; i++ {
		for id, _ := range d.checkroll {
			err := d.p.SendMessage([]byte(id), pb)
			if err != nil {
				fmt.Println("SendMessage err ", err)
				d.log.WithField("event", "SendMessage err").Warn(new(big.Int).SetBytes([]byte(id)).String())
			}
		}
	}
}

func (r test2) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {
	if content.Ctype == internalMsg.Cmd_SIGNIN {
		d.checkroll[sender] = d.checkroll[sender] + 1
		//fmt.Println("sender", []byte(sender), "  ", d.checkroll[sender], content.Ctype, " ", len(d.checkroll))
		if d.checkroll[sender] == d.numMessages {
			delete(d.checkroll, sender)
		}

		d.log.WithField("event", "peer remain").Info(len(d.checkroll))

		if len(d.checkroll) == 0 {
			d.log.WithField("event", "test done").Info()
			d.MakeRequest(d.bootStrapIp, "post", d.nodeID)
		}
	}
}
