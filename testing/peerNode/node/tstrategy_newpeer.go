package node

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
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
			for id := range d.checkroll {
				var err error
				if err = d.p.SendMessage([]byte(id), pb); err != nil {
					retry := 1
					for err != nil {
						d.log.WithFields(logrus.Fields{
							"SendMessageErr": err,
						}).Info()
						retry++
						if retry > 20 {
							break
						}
						err = d.p.SendMessage([]byte(id), pb)
					}
				}
				if err == nil {
					d.log.WithFields(logrus.Fields{
						"SendMessageSuccess": true,
					}).Info()
				}
			}
		}
	}
	//d.done <- true
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
				d.FinishTest()
				//d.done <- true
			} else {
				fmt.Println("wait for  = ", len(d.checkroll))
				for id := range d.checkroll {
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
		d.FinishTest()
		//d.done <- true
	}
}

type test2 struct{}

func (r test2) StartTest(d *PeerNode) {
	fmt.Println("StartTest")
	d.log.WithFields(logrus.Fields{
		"eventStartTest": true,
	}).Info()
	id := len(d.nodeIPs) - 1
	//d.done <- true

	if d.p.GetIP() == d.nodeIPs[id] {
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
		d.FinishTest()
		//d.done <- true
	}
}

func (r test2) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {
	id := len(d.nodeIPs) - 1
	if d.p.GetIP() != d.nodeIPs[id] {
		if content.Ctype == internalMsg.Cmd_SIGNIN {
			d.log.WithFields(logrus.Fields{
				"eventCheckDone": true,
			}).Info()
		}
		d.FinishTest()
	}
}

type test3 struct{}

func (r test3) StartTest(d *PeerNode) {
	d.log.WithFields(logrus.Fields{
		"eventStartTest": true,
	}).Info()

	groupSizeStr := os.Getenv("GROUPSIZE")
	groupSize, err := strconv.Atoi(groupSizeStr)
	if err != nil {
		d.log.Fatal(err)
	}

	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(d.p, suite, d.dkgChan, groupSize, d.log)
	if err != nil {
		d.log.Fatal(err)
	}
	go p2pDkg.EventLoop()
	dkgEvent := make(chan string, 1)
	p2pDkg.SubscribeEvent(dkgEvent)
	defer close(dkgEvent)

	var group [][]byte
	for idx, id := range d.nodeIDs {
		if bytes.Compare(d.p.GetID(), id) == 0 {
			start := idx / groupSize * groupSize
			group = d.nodeIDs[start : start+groupSize]
			break
		}
	}

	p2pDkg.SetGroupMembers(group)
	p2pDkg.RunDKG()

	result := <-dkgEvent
	if result == "certified" {
		d.log.WithFields(logrus.Fields{
			"eventCheckDone": true,
		}).Info()
		d.FinishTest()
	}

}

func (r test3) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {}

type test4 struct{}

func (r test4) StartTest(d *PeerNode) {
	d.log.WithFields(logrus.Fields{
		"eventStartTest": true,
	}).Info()

	groupSizeStr := os.Getenv("GROUPSIZE")
	groupSize, err := strconv.Atoi(groupSizeStr)
	if err != nil {
		d.log.Fatal(err)
	}

	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(d.p, suite, d.dkgChan, groupSize, d.log)
	if err != nil {
		d.log.Fatal(err)
	}
	go p2pDkg.EventLoop()
	dkgEvent := make(chan string, 1)
	p2pDkg.SubscribeEvent(dkgEvent)
	defer close(dkgEvent)

	var group [][]byte
	for idx, id := range d.nodeIDs {
		if bytes.Compare(d.p.GetID(), id) == 0 {
			start := idx / groupSize * groupSize
			group = d.nodeIDs[start : start+groupSize]
			break
		}
	}

	p2pDkg.SetGroupMembers(group)
	p2pDkg.RunDKG()

	var signatures [][]byte
	go func() {
		for sig := range d.tblsChan {
			signatures = append(signatures, sig.Signature)
			if len(signatures) > groupSize/2 {
				finalSig, err := tbls.Recover(suite, p2pDkg.GetGroupPublicPoly(), sig.Content, signatures, groupSize/2+1, groupSize)
				if err != nil {
					d.log.WithFields(logrus.Fields{
						"eventTblsRecoverErr": true,
					}).Info(err)
					continue
				}
				if err = bls.Verify(suite, p2pDkg.GetGroupPublicPoly().Commit(), sig.Content, finalSig); err != nil {
					d.log.WithFields(logrus.Fields{
						"eventTblsVerifyErr": true,
					}).Info(err)
					continue
				} else {
					d.FinishTest()
					break
				}
			}
		}
	}()

	result := <-dkgEvent
	if result == "certified" {
		d.log.WithFields(logrus.Fields{
			"eventCheckDone": true,
		}).Info()
	}

	rawMsg, err := dataFetch(os.Getenv("URL"))
	if err != nil {
		d.log.WithFields(logrus.Fields{
			"eventFetchURLFail": true,
		}).Info(err)
	} else {
		d.log.WithFields(logrus.Fields{
			"eventFetchURL": true,
		}).Info(string(rawMsg))
	}

	sig, err := tbls.Sign(suite, p2pDkg.GetShareSecurity(), rawMsg)
	sign := &vss.Signature{
		Content:   rawMsg,
		Signature: sig,
	}
	signatures = append(signatures, sig)
	for _, id := range d.nodeIDs {
		if bytes.Compare(d.p.GetID(), id) != 0 {
			if err = d.p.SendMessage(id, sign); err != nil {
				d.log.WithFields(logrus.Fields{
					"eventSendSignatureErr": true,
				}).Info(err)
			}
		}
	}
}

func (r test4) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {}

func dataFetch(url string) (body []byte, err error) {
	client := &http.Client{Timeout: 60 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = r.Body.Close()
	return
}

type test5 struct{}

func (r test5) StartTest(d *PeerNode) {
	d.log.WithFields(logrus.Fields{
		"eventStartTest": true,
	}).Info()

	groupSizeStr := os.Getenv("GROUPSIZE")
	groupSize, err := strconv.Atoi(groupSizeStr)
	if err != nil {
		d.log.Fatal(err)
	}

	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(d.p, suite, d.dkgChan, groupSize, d.log)
	if err != nil {
		d.log.Fatal(err)
	}
	go p2pDkg.EventLoop()
	dkgEvent := make(chan string, 1)
	p2pDkg.SubscribeEvent(dkgEvent)
	defer close(dkgEvent)

	roundCount := uint16(1)
	for {
		var group [][]byte
		for idx, id := range d.nodeIDs {
			if bytes.Compare(d.p.GetID(), id) == 0 {
				start := idx / groupSize * groupSize
				group = d.nodeIDs[start : start+groupSize]
				break
			}
		}

		p2pDkg.SetGroupMembers(group)
		p2pDkg.RunDKG()

		result := <-dkgEvent
		if result == "certified" {
			d.log.WithFields(logrus.Fields{
				"eventCheckRoundDone": roundCount,
			}).Info()
			p2pDkg.Reset()
			next := d.requestIsNextRoundReady(roundCount)
			if next == byte(DKGROUNDFINISH) {
				break
			} else {
				roundCount++
				//rand.Shuffle(len(d.nodeIDs), func(i, j int) {
				//	d.nodeIDs[i], d.nodeIDs[j] = d.nodeIDs[j], d.nodeIDs[i]
				//})
			}
		}
	}
	//d.FinishTest()
}

func (r test5) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {}
