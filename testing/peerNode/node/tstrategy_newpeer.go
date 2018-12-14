package node

import (
	"bytes"
	"fmt"
	//	"math/big"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
)

type test1 struct{}

func (r test1) StartTest(d *PeerNode) {
	d.log.WithFields(logrus.Fields{
		"eventStartTest": true,
	}).Info()
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
					d.log.WithFields(logrus.Fields{
						"eventConnectToErr": err,
					}).Info()
				} else {
					d.log.WithFields(logrus.Fields{
						"eventConnectTo": true,
					}).Info()

				}
				fmt.Println("ConnectTo To", id, " ", ip)
				d.checkroll[string(id)] = 0
			}
		}
		for i := 0; i < d.numMessages; i++ {
			for id, _ := range d.checkroll {
				if err := d.p.SendMessage([]byte(id), pb); err != nil {
					retry := 1
					for err != nil {
						d.log.WithFields(logrus.Fields{
							"SendMessageErr": err,
						}).Info()
						retry++
						err = d.p.SendMessage([]byte(id), pb)
						if retry > 20 {
							break
						}
					}
				} else {
					d.log.WithFields(logrus.Fields{
						"SendMessageSuccess": true,
					}).Info()
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
			d.log.WithFields(logrus.Fields{
				"eventWaitFor": len(d.checkroll),
			}).Info()
			if len(d.checkroll) == 0 {
				d.log.WithFields(logrus.Fields{
					"eventCheckDone": true,
				}).Info()
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
			retry := 0
			for err != nil {
				retry++
				err = d.p.SendMessage([]byte(sender), pb)
				if retry >= 10 {
					return
				}
			}
		}
		d.log.WithFields(logrus.Fields{
			"eventCheckDone": true,
		}).Info()
		d.MakeRequest(d.bootStrapIp, "post", d.nodeID)
	}
}

type test2 struct{}

func (r test2) StartTest(d *PeerNode) {
	d.log.WithFields(logrus.Fields{
		"eventStartTest": true,
	}).Info()
	if d.p.GetIP() == d.nodeIPs[5] {
		cmd := &internalMsg.Cmd{
			Ctype: internalMsg.Cmd_SIGNIN,
			Args:  []byte{},
		}
		pb := proto.Message(cmd)
		for i := 0; i < len(d.nodeIDs); i++ {
			if !bytes.Equal(d.p.GetID(), d.nodeIDs[i]) {
				if err := d.p.SendMessage(d.nodeIDs[i], pb); err != nil {
					retry := 0
					for err != nil {
						err = d.p.SendMessage(d.nodeIDs[i], pb)
						if retry > 20 {
							break
						}
						retry++
					}
				}
			}
		}
		d.log.WithFields(logrus.Fields{
			"eventCheckDone": true,
		}).Info()
	}
}

func (r test2) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {
	if d.p.GetIP() != d.nodeIPs[5] {
		if content.Ctype == internalMsg.Cmd_SIGNIN {
			d.log.WithFields(logrus.Fields{
				"eventCheckDone": true,
			}).Info()
		}
	}
}
