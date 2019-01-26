package node

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/DOSNetwork/core/share/vss/pedersen"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"math/rand"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
)

type PeerNode struct {
	node
	bootStrapIp string
	nodeID      []byte
	nodeIDs     [][]byte
	nodeIPs     []string
	checkCount  int
	findNodeDur time.Duration
	checkroll   map[string]int
	numMessages int
	tStrategy   TestStrategy
	tblsChan    chan vss.Signature
}

func (d *PeerNode) MakeRequest(bootStrapIp, f string, args []byte) ([]byte, error) {

	tServer := "http://" + bootStrapIp + ":8080/" + f

	req, err := http.NewRequest("POST", tServer, bytes.NewBuffer(args))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	return r, err
}

func (d *PeerNode) requestAllIDs() {
	for {
		r, err := d.MakeRequest(d.bootStrapIp, "getAllIDs", []byte{})
		for err != nil {
			time.Sleep(10 * time.Second)
			r, err = d.MakeRequest(d.bootStrapIp, "getAllIDs", []byte{})
		}

		if err != nil {
		} else {
			if len(r) != 0 {
				num := len(r) / len(d.nodeID)
				d.nodeIDs = make([][]byte, num)
				for i := 0; i < num; i++ {
					d.nodeIDs[i] = make([]byte, len(d.nodeID))
					copy(d.nodeIDs[i], r[i*len(d.nodeID):i*len(d.nodeID)+len(d.nodeID)])
				}
				break
			} else {
				os.Exit(0)
			}
		}
	}
}

func (d *PeerNode) requestAllIPs() {
	for {
		r, err := d.MakeRequest(d.bootStrapIp, "getAllIPs", []byte{})
		for err != nil {
			time.Sleep(10 * time.Second)
			r, err = d.MakeRequest(d.bootStrapIp, "getAllIPs", []byte{})
		}

		if err != nil {
			fmt.Println(err)
		} else {
			if len(r) != 0 {
				str := string(r)
				strlist := strings.Split(str, ",")
				d.nodeIPs = make([]string, len(strlist)-1)
				for i := 0; i < len(strlist)-1; i++ {
					d.nodeIPs[i] = strlist[i]
				}
				break
			}
		}
	}
}

func (d *PeerNode) requestIsReady() bool {
	ip, _ := p2p.GetLocalIP()
	ip += ":44460"

	r, err := d.MakeRequest(d.bootStrapIp, "isTestReady", []byte(ip))
	for err != nil {
		time.Sleep(10 * time.Second)
		r, err = d.MakeRequest(d.bootStrapIp, "isTestReady", []byte(ip))
	}

	if err != nil {
		fmt.Println(err)
		return false
	} else if len(r) == 0 {
		return false
	} else if r[0] == byte(ALLNODEREADY) {
		return true
	} else {
		return false
	}
}

func (d *PeerNode) requestIsNextRoundReady(roundCount uint16) byte {
	ip, _ := p2p.GetLocalIP()
	ip += ":44460"
	roundCountBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(roundCountBytes, roundCount)
	request := append([]byte(ip), roundCountBytes...)

	r, err := d.MakeRequest(d.bootStrapIp, "isNextRoundReady", request)
	for r[0] == byte(ALLNODENOTREADY) || err != nil {
		time.Sleep(1 * time.Second)
		r, err = d.MakeRequest(d.bootStrapIp, "isNextRoundReady", request)
	}

	return r[0]
}

func (d *PeerNode) requestIsFinish() bool {
	ip, _ := p2p.GetLocalIP()
	ip += ":44460"

	r, err := d.MakeRequest(d.bootStrapIp, "isTestFinish", []byte(ip))
	for err != nil {
		time.Sleep(10 * time.Second)
		r, err = d.MakeRequest(d.bootStrapIp, "isTestFinish", []byte(ip))
	}

	if err != nil {
		fmt.Println(err)
		return false
	} else if len(r) == 0 {
		return false
	} else if r[0] == byte(ALLNODEFINISH) {
		return true
	} else {
		return false
	}

}

func (d *PeerNode) Init(bootStrapIp string, port, peerSize int, numMessages int, tStrategy string) {
	d.peerSize = peerSize
	d.checkCount = 1
	d.bootStrapIp = bootStrapIp
	d.checkroll = make(map[string]int)
	d.done = make(chan bool)
	d.tblsChan = make(chan vss.Signature)
	d.numMessages = numMessages

	switch tStrategy {
	case "SENDMESSAGE":
		d.tStrategy = &test1{}
	case "FINDNODE":
		d.tStrategy = &test2{}
	case "DKG":
		d.tStrategy = &test3{}
	case "TBLS":
		d.tStrategy = &test4{}
	case "DKGMULTIGROUPING":
		d.tStrategy = &test5{}
	}

	//1)Wait until bootstrap node assign an ID
	for {
		ip, _ := p2p.GetLocalIP()
		ip = ip + ":44460"
		r, err := d.MakeRequest(bootStrapIp, "getID", []byte(ip))
		for err != nil {
			time.Sleep(10 * time.Second)
			r, err = d.MakeRequest(bootStrapIp, "getID", []byte(ip))
		}

		if err != nil {
			fmt.Println(err)
		} else {
			if len(r) != 0 {
				d.nodeID = r
				break
			} else {
				os.Exit(0)
			}
		}
	}
	log.Init(d.nodeID[:])

	//2)Build a p2p network
	d.p, _ = p2p.CreateP2PNetwork(d.nodeID[:], port)
	peerEvent, _ := d.p.SubscribeEvent(100, internalMsg.Cmd{}, vss.Signature{})
	d.p.Listen()
	go func() {
		for msg := range peerEvent {
			switch content := msg.Msg.Message.(type) {
			case *internalMsg.Cmd:

				if content.Ctype == internalMsg.Cmd_SIGNIN {
					sender := string(msg.Sender)
					go d.tStrategy.CheckResult(sender, content, d)
					response := &internalMsg.Cmd{}
					replyNonce := msg.RequestNonce
					d.p.Reply([]byte(sender), replyNonce, response)
				}
			case *vss.Signature:
				go func() { d.tblsChan <- *content }()
				if err := d.p.Reply(msg.Sender, msg.RequestNonce, &vss.Signature{}); err != nil {
					fmt.Println("signature reply", err)
				}
			default:
			}
		}
	}()

	//fmt.Println("nodeIP = ", d.p.GetIP())

	d.requestAllIDs() //get all ids
	d.requestAllIPs() //get all ips

	if tStrategy != "SENDMESSAGE" {
		for i := 0; i < int(math.Min(4, float64(len(d.nodeIPs)))); i++ {
			if d.p.GetIP() != d.nodeIPs[i] {
				d.p.Join(d.nodeIPs[i])
			}
		}
	}
}

func (d *PeerNode) FinishTest() {
	ticker := time.NewTicker(5 * time.Second)
L:
	for {
		select {
		case <-ticker.C:
			if d.requestIsFinish() {
				ticker.Stop()
				d.done <- true
				break L
			}
		default:
		}
	}
}

func (d *PeerNode) EventLoop() {
	ticker := time.NewTicker(5 * time.Second)
L:
	for {
		select {
		case <-ticker.C:
			if d.requestIsReady() {
				ticker.Stop()
				log.Progress("StartTest")
				go d.tStrategy.StartTest(d)
			}
		case <-d.done:
			log.Progress("EndTest")
			break L
		default:
		}
	}
	os.Exit(0)
}

func (d *PeerNode) CloseConnectionRandom(interval int) {
	fmt.Println("CloseConnectionLoop begin")
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			n := d.p.GetPeerConnManager().PeerConnNum()
			if n > 0 {
				rn := rand.Uint32() % n
				count := 0
				var peerid string
				peerid = ""
				d.p.GetPeerConnManager().Range(func(key, value interface{}) bool {
					if uint32(count) == rn {
						peerid = key.(string)
						return false
					}
					count++
					return true
				})
				if peerid != "" {
					d.p.GetPeerConnManager().DeletePeer(peerid)
				}
			}

		default:
		}
	}
}
