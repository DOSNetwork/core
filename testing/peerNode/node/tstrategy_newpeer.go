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
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	//log "github.com/DOSNetwork/core/log"
	"github.com/golang/protobuf/proto"
)

type test1 struct{}

func (r test1) StartTest(d *PeerNode) {
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_SIGNIN,
		Args:  []byte{},
	}
	pb := proto.Message(cmd)
	var wg sync.WaitGroup
	wg.Add(len(d.nodeIDs) - 1)
	for i := 0; i < len(d.nodeIPs); i++ {
		if !strings.Contains(d.p.GetIP().String(), d.nodeIPs[i]) {
			go func(j int) {
				defer wg.Done()
				ip := d.nodeIPs[j]
				id, err := d.p.ConnectTo(ip, nil)
				if err != nil {
					os.Exit(1)
				}
				for k := 0; k < d.numMessages; k++ {
					start := time.Now()
					if _, err := d.p.Request(id, pb); err != nil {
						os.Exit(1)
					} else {
						fmt.Println("Request done", time.Since(start).Nanoseconds()/1000)
					}

				}
			}(i)
		}
	}
	wg.Wait()

	d.FinishTest()
}

func (r test1) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {

}

type test2 struct{}

func (r test2) StartTest(d *PeerNode) {
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_SIGNIN,
		Args:  []byte{},
	}
	pb := proto.Message(cmd)

	for i := 0; i < len(d.nodeIDs); i++ {
		if !bytes.Equal(d.p.GetID(), d.nodeIDs[i]) {
			start := time.Now()
			retry := 0
			for {
				if _, err := d.p.Request(d.nodeIDs[i], pb); err != nil {
					fmt.Println(start, " ->", time.Now(), "retry ", retry, d.p.GetID(), " ->", d.nodeIDs[i], err)
					retry++
					time.Sleep(1 * time.Second)
				} else {
					break
				}
			}
			fmt.Println("Request done", time.Since(start).Nanoseconds()/1000)
		}
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Start Test FindNode")

	var wg sync.WaitGroup
	/*
		wg.Add((len(d.nodeIDs) - 1) * d.numMessages)
		fmt.Println("Test start test2", (len(d.nodeIDs)-1)*d.numMessages)

		for i := 0; i < len(d.nodeIDs); i++ {
			if !bytes.Equal(d.p.GetID(), d.nodeIDs[i]) {
				for j := 0; j < d.numMessages; j++ {
					fmt.Println("Test start count", i, j)
					go func(p p2p.P2PInterface, id []byte) {
						fmt.Println("Test start goroutine")
						start := time.Now()
						defer wg.Done()
						if _, err := p.Request(id, pb); err != nil {
							fmt.Println(start, " ->", time.Now(), "testfailed ", p.GetID(), " ->", id, err)
							os.Exit(1)
						} else {
							fmt.Println("Test done", time.Since(start).Nanoseconds()/1000)
						}
					}(d.p, d.nodeIDs[i])
				}
			}
		}
	*/
	start := time.Now()

	wg.Add((len(d.nodeIDs) - 1))
	for i := 0; i < len(d.nodeIDs); i++ {
		if !bytes.Equal(d.p.GetID(), d.nodeIDs[i]) {
			go func(i int) {
				defer wg.Done()
				start := time.Now()
				for j := 0; j < d.numMessages; j++ {
					if _, err := d.p.Request(d.nodeIDs[i], pb); err != nil {
						os.Exit(1)
					} else {
						fmt.Println("Test done", time.Since(start).Nanoseconds()/1000)
					}
				}
			}(i)
		}
	}

	wg.Wait()
	fmt.Println("Test done", time.Since(start).Nanoseconds()/1000)
	time.Sleep(10 * time.Second)

	d.FinishTest()
}

func (r test2) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {

}

type test3 struct{}

func (r test3) StartTest(d *PeerNode) {
	groupSizeStr := os.Getenv("GROUPSIZE")
	groupSize, err := strconv.Atoi(groupSizeStr)
	if err != nil {
		//d.log.Fatal(err)
	}

	for i := 0; i < len(d.nodeIDs); i++ {
		if !bytes.Equal(d.p.GetID(), d.nodeIDs[i]) {
			start := time.Now()
			retry := 0
			for {
				fmt.Println(start, " ->", time.Now(), "ConnectTo ", d.p.GetID(), " ->", d.nodeIDs[i])

				if _, err := d.p.ConnectTo("", d.nodeIDs[i]); err != nil {
					fmt.Println(start, " ->", time.Now(), "testfailed ", d.p.GetID(), " ->", d.nodeIDs[i], err)
					retry++
				} else {
					break
				}
			}
			fmt.Println("Request ConnectTo retry=", retry, d.p.GetID(), " ->", d.nodeIDs[i], time.Since(start).Nanoseconds()/1000)
		}
	}

	time.Sleep(10 * time.Second)
	startTime := time.Now()
	for idx, id := range d.nodeIDs {
		if bytes.Compare(d.p.GetID(), id) == 0 {
			start := idx / groupSize * groupSize
			group := d.nodeIDs[start : start+groupSize]
			dkgEvent, errChan := d.p2pDkg.Start(context.Background(), group, fmt.Sprintf("%x", group))
			go func() {
				for err := range errChan {
					fmt.Println("errorChan", err)
				}
			}()
			if _, succ := <-dkgEvent; succ {
				fmt.Println("dkgTest done", time.Since(startTime).Nanoseconds()/1000)
				d.FinishTest()
			} else {
				fmt.Println("dkgTest Failed", time.Since(startTime).Nanoseconds()/1000)
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

	var group [][]byte
	for idx, id := range d.nodeIDs {
		if bytes.Compare(d.p.GetID(), id) == 0 {
			start := idx / groupSize * groupSize
			group = d.nodeIDs[start : start+groupSize]
			dkgEvent, errChan := p2pDkg.Start(context.Background(), group, fmt.Sprintf("%x", group))
			go func() {
				for err := range errChan {
					fmt.Println("errorChan", err)
				}
			}()
			var IdWithPubKey [5]*big.Int
			if IdWithPubKey, succ = <-dkgEvent; !succ {
				fmt.Println(errors.New("err: dkg fail"))
				return
			}
			copy(pubKey[:], IdWithPubKey[1:])
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

	peerEvent, err := d.p.SubscribeEvent(100, vss.Signature{})
	if err != nil {
		fmt.Println(err)
	}

	groupPeers := make(map[string][]byte)
	for _, id := range group {
		if bytes.Compare(d.p.GetID(), id) != 0 {
			groupPeers[string(id)] = id
		}
	}

	go func() {
		for len(groupPeers) > 0 {
			for key, id := range groupPeers {
				if msg, err := d.p.Request(id, &vss.Signature{}); err == nil {
					delete(groupPeers, key)
					sigMsg := msg.Msg.Message.(*vss.Signature)
					signatures = append(signatures, sigMsg.Signature)
					if len(signatures) > groupSize/2 {
						finalSig, err := tbls.Recover(suite, p2pDkg.GetGroupPublicPoly(pubKey), sigMsg.Content, signatures, groupSize/2+1, groupSize)
						if err != nil {
							fmt.Println(err)
							continue
						}
						if err = bls.Verify(suite, p2pDkg.GetGroupPublicPoly(pubKey).Commit(), sigMsg.Content, finalSig); err != nil {
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
		}
	}()

	for msg := range peerEvent {
		if err = d.p.Reply(msg.Sender, msg.RequestNonce, sign); err != nil {
			fmt.Println(err)
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
			dkgEvent chan [5]*big.Int
			errChan  <-chan error
		)
		for idx, id := range d.nodeIDs {
			if bytes.Compare(d.p.GetID(), id) == 0 {
				start := idx / groupSize * groupSize
				group = d.nodeIDs[start : start+groupSize]
				dkgEvent, errChan = p2pDkg.Start(context.Background(), group, fmt.Sprintf("%x", group))
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
			rdm := rand.New(rand.NewSource(0))
			rdm.Shuffle(len(d.nodeIDs), func(i, j int) {
				d.nodeIDs[i], d.nodeIDs[j] = d.nodeIDs[j], d.nodeIDs[i]
			})
		}
	}
	d.FinishTest()
}

func (r test5) CheckResult(sender string, content *internalMsg.Cmd, d *PeerNode) {}
