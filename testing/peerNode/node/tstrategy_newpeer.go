package node

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
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

	//	log "github.com/DOSNetwork/core/log"
	"github.com/golang/protobuf/proto"
)

type test1 struct{}

func (r test1) StartTest(d *PeerNode) {
	fmt.Println("StartTest")
	if d.p.GetIP() == d.nodeIPs[0] {
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
				d.checkroll[string(id)] = 0
			}
		}
		for i := 0; i < d.numMessages; i++ {
			for id := range d.checkroll {
				var err error
				fmt.Println("SendMessage ", []byte(id))
				if err = d.p.SendMessage([]byte(id), pb); err != nil {
					retry := 1
					for err != nil {
						fmt.Println("SendMessage err", err)

						retry++
						if retry > 20 {
							break
						}
						err = d.p.SendMessage([]byte(id), pb)
					}
				}
			}
		}
	}
}

func (r test1) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {
	fmt.Println("CheckResult ")

	if d.p.GetIP() == d.nodeIPs[0] {
		fmt.Println("CheckResult ")

		if content.Ctype == internalMsg.Cmd_SIGNIN {
			d.checkroll[sender] = d.checkroll[sender] + 1
			if d.checkroll[sender] == d.numMessages {
				delete(d.checkroll, sender)

				if len(d.checkroll) == 0 {
					d.FinishTest()
				} else {
					fmt.Println("wait for  = ", len(d.checkroll))
					for id := range d.checkroll {
						fmt.Println("wait for ", []byte(id))
					}
					fmt.Println("==================== ")
				}
			}
		}
	} else {
		cmd := &internalMsg.Cmd{
			Ctype: internalMsg.Cmd_SIGNIN,
			Args:  []byte{},
		}
		fmt.Println("SendMessage 11")

		pb := proto.Message(cmd)
		if err := d.p.SendMessage([]byte(sender), pb); err != nil {
			retry := 0
			for err != nil {
				retry++
				fmt.Println("SendMessage ")

				err = d.p.SendMessage([]byte(sender), pb)
				if retry >= 10 {
					return
				}
			}
		}
		d.FinishTest()
	}
}

type test2 struct{}

func (r test2) StartTest(d *PeerNode) {
	id := len(d.nodeIPs) - 1

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
		d.FinishTest()
	}
}

func (r test2) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {
	id := len(d.nodeIPs) - 1
	if d.p.GetIP() != d.nodeIPs[id] {
		d.FinishTest()
	}
}

type test3 struct{}

func (r test3) StartTest(d *PeerNode) {
	groupSizeStr := os.Getenv("GROUPSIZE")
	groupSize, err := strconv.Atoi(groupSizeStr)
	if err != nil {
		//d.log.Fatal(err)
	}

	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(d.p, suite)
	if err != nil {
		fmt.Println(err)
	}

	for idx, id := range d.nodeIDs {
		if bytes.Compare(d.p.GetID(), id) == 0 {
			start := idx / groupSize * groupSize
			group := d.nodeIDs[start : start+groupSize]
			dkgEvent, errChan := p2pDkg.Start(context.Background(), group, dkg.GIdsToSessionID(group))
			go func() {
				for err := range errChan {
					fmt.Println("errorChan", err)
				}
			}()
			if pubKey, succ := <-dkgEvent; succ {
				fmt.Println("eventCheckDone", pubKey)
				d.FinishTest()
			} else {
				fmt.Println(errors.New("err: dkg fail"))
				return
			}
			break
		}
	}
}

func (r test3) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {}

type test4 struct{}

func (r test4) StartTest(d *PeerNode) {
	groupSizeStr := os.Getenv("GROUPSIZE")
	groupSize, err := strconv.Atoi(groupSizeStr)

	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(d.p, suite)
	if err != nil {
		fmt.Println(err)
	}

	var pubKey [4]*big.Int
	var succ bool

	for idx, id := range d.nodeIDs {
		if bytes.Compare(d.p.GetID(), id) == 0 {
			start := idx / groupSize * groupSize
			group := d.nodeIDs[start : start+groupSize]
			dkgEvent, errChan := p2pDkg.Start(context.Background(), group, dkg.GIdsToSessionID(group))
			go func() {
				for err := range errChan {
					fmt.Println("errorChan", err)
				}
			}()
			if pubKey, succ = <-dkgEvent; !succ {
				fmt.Println(errors.New("err: dkg fail"))
				return
			}
			break
		}
	}

	rawMsg, err := dataFetch(os.Getenv("URL"))
	if err != nil {
		fmt.Println(err)
	}

	var signatures [][]byte
	sig, err := tbls.Sign(suite, p2pDkg.GetShareSecurity(pubKey), rawMsg)
	if err != nil {
		fmt.Println(err)
	}
	sign := &vss.Signature{
		Content:   rawMsg,
		Signature: sig,
	}
	signatures = append(signatures, sig)
	for _, id := range d.nodeIDs {
		if bytes.Compare(d.p.GetID(), id) != 0 {
			if err = d.p.SendMessage(id, sign); err != nil {
				fmt.Println(err)
			}
		}
	}
	for sig := range d.tblsChan {
		signatures = append(signatures, sig.Signature)
		if len(signatures) > groupSize/2 {
			finalSig, err := tbls.Recover(suite, p2pDkg.GetGroupPublicPoly(pubKey), sig.Content, signatures, groupSize/2+1, groupSize)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if err = bls.Verify(suite, p2pDkg.GetGroupPublicPoly(pubKey).Commit(), sig.Content, finalSig); err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Println("TBLS SUCCESS!!!")
				d.FinishTest()
				break
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
	groupSizeStr := os.Getenv("GROUPSIZE")
	groupSize, err := strconv.Atoi(groupSizeStr)
	if err != nil {
		//d.log.Fatal(err)
	}

	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(d.p, suite)
	if err != nil {
		fmt.Println(err)
	}

	roundCount := uint16(1)
	for {
		var (
			group    [][]byte
			dkgEvent chan [4]*big.Int
			errChan  <-chan error
		)
		for idx, id := range d.nodeIDs {
			if bytes.Compare(d.p.GetID(), id) == 0 {
				start := idx / groupSize * groupSize
				group = d.nodeIDs[start : start+groupSize]
				dkgEvent, errChan = p2pDkg.Start(context.Background(), group, dkg.GIdsToSessionID(group))
				go func() {
					for err := range errChan {
						fmt.Println("errorChan", err)
					}
				}()
				break
			}
		}
		if _, succ := <-dkgEvent; !succ {
			fmt.Println(errors.New("err: dkg fail"))
			return
		}
		fmt.Println("\n certified!!!!!!")
		fmt.Println("eventCheckRoundDone", roundCount)
		next := d.requestIsNextRoundReady(roundCount)
		if next == byte(DKGROUNDFINISH) {
			break
		} else {
			roundCount++
			rand.Shuffle(len(d.nodeIDs), func(i, j int) {
				d.nodeIDs[i], d.nodeIDs[j] = d.nodeIDs[j], d.nodeIDs[i]
			})
		}
	}
	d.FinishTest()
}

func (r test5) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {}
