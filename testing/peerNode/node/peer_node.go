package node

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"

	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"github.com/DOSNetwork/core/p2p"
	"github.com/sirupsen/logrus"
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
	dkgChan     chan p2p.P2PMessage
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
	fmt.Println("requestAllIDs")
	for {
		r, err := d.MakeRequest(d.bootStrapIp, "getAllIDs", []byte{})
		for err != nil {
			time.Sleep(10 * time.Second)
			r, err = d.MakeRequest(d.bootStrapIp, "getAllIDs", []byte{})
		}

		if err != nil {
			fmt.Println(err)
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
	//d.log.WithField("event", "requestAllIds done").Info()
}

func (d *PeerNode) requestAllIPs() {
	fmt.Println("requestAllIPs")
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
	d.log.WithField("event", "requestAllIPs done").Info()
}

func (d *PeerNode) requestIsReady() bool {
	ip, _ := p2p.GetLocalIp()

	r, err := d.MakeRequest(d.bootStrapIp, "isTestReady", []byte(ip))
	for err != nil {
		time.Sleep(10 * time.Second)
		r, err = d.MakeRequest(d.bootStrapIp, "isTestReady", []byte(ip))
	}

	if err != nil {
		fmt.Println(err)
	} else {
		if len(r) != 0 {
			if r[0] == byte(ALLNODEREADY) {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return false
}

func (d *PeerNode) Init(bootStrapIp string, port, peerSize int, numMessages int, tStrategy string, logger *logrus.Entry) {
	d.peerSize = peerSize
	d.checkCount = 1
	d.bootStrapIp = bootStrapIp
	d.checkroll = make(map[string]int)
	d.done = make(chan bool)
	d.dkgChan = make(chan p2p.P2PMessage)
	d.numMessages = numMessages
	d.log = logger

	switch tStrategy {
	case "SENDMESSAGE":
		d.tStrategy = &test1{}
	case "FINDNODE":
		d.tStrategy = &test2{}
	case "DKG":
		d.tStrategy = &test3{}
	}

	//1)Wait until bootstrap node assign an ID
	for {
		ip, _ := p2p.GetLocalIp()
		ip = ip + ":44460"
		fmt.Println("IP : ", ip)
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
	fmt.Println("nodeID = ", d.nodeID[:])
	//2)Build a p2p network
	d.p, d.peerEvent, _ = p2p.CreateP2PNetwork(d.nodeID[:], port, d.log)
	d.p.Listen()
	go func() {
		for msg := range d.peerEvent {
			switch content := msg.Msg.Message.(type) {
			case *internalMsg.Cmd:
				if content.Ctype == internalMsg.Cmd_ALLIP {
					nodes := strings.Split(string(content.Args), ",")
					allIP := []string{}
					ip, _ := p2p.GetLocalIp()
					ip = ip + ":44460"
					for _, node := range nodes {
						if ip != node {
							allIP = append(allIP, node)
						}
					}
					d.nodeIPs = allIP
				} else if content.Ctype == internalMsg.Cmd_ALLID {
					buf := []byte(content.Args)
					allID := [][]byte{}
					var chunk []byte
					lim := 20
					for len(buf) >= lim {
						chunk, buf = buf[:lim], buf[lim:]
						allID = append(allID, chunk)
					}
					d.nodeIDs = allID
				} else if content.Ctype == internalMsg.Cmd_SIGNIN {
					sender := string(msg.Sender)
					go d.tStrategy.CheckResult(sender, content, d)
				}
			case *vss.PublicKey:
				d.dkgChan <- msg
				fmt.Println("PublicKey")
			case *dkg.Deal:
				d.dkgChan <- msg
				fmt.Println("Deal")
			case *dkg.Response:
				d.dkgChan <- msg
				fmt.Println("Response")
			default:
			}
		}
	}()

	//d.p.Join(bootStrapIp + ":44460")
	fmt.Println("nodeIP = ", d.p.GetIP())

	d.log.Data["testType"] = tStrategy
	d.log.Data["role"] = "peernode"
	d.log.Data["nodeId"] = new(big.Int).SetBytes(d.nodeID).String()

	d.requestAllIDs() //get all ids
	d.requestAllIPs() //get all ips

	if tStrategy != "SENDMESSAGE" {
		for i := 0; i < int(math.Min(4, float64(len(d.nodeIPs)))); i++ {
			if d.p.GetIP() != d.nodeIPs[i] {
				d.p.Join(d.nodeIPs[i])
			}
		}
	}
	//teststart todo
	//send test result todo
}

func (d *PeerNode) EventLoop() {
	ticker := time.NewTicker(5 * time.Second)
	count := 0
L:
	for {
		select {
		case <-ticker.C:
			if d.requestIsReady() {
				ticker.Stop()
				go d.tStrategy.StartTest(d)
			}
		case <-d.done:
			count++
			fmt.Println("done  count ", count)
			if count >= 2 {
				fmt.Println("EventLoop done")
				d.log.WithField("event", "EventLoop done").Info()
				break L
			}
		default:
		}
	}
	os.Exit(0)
}
