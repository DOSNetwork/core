package node

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bshuster-repo/logrus-logstash-hook"
	//	"github.com/ethereum/go-ethereum/common"

	//"github.com/DOSNetwork/core/p2p/dht"
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

func (d *PeerNode) Init(bootStrapIp string, port, peerSize int, numMessages int, tStrategy string, logger *logrus.Logger) {
	d.peerSize = peerSize
	d.checkCount = 1
	d.bootStrapIp = bootStrapIp
	d.checkroll = make(map[string]int)
	d.done = make(chan bool)
	d.numMessages = numMessages
	d.log = logger
	if tStrategy == "SENDMESSAGE" {
		d.tStrategy = &test1{}
	} else {
		d.tStrategy = &test2{}
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

	d.requestAllIDs() //get all ids
	d.requestAllIPs() //get all ips

	//teststart todo
	//send test result todo

	//2)Build a p2p network
	d.p, d.peerEvent, _ = p2p.CreateP2PNetwork(d.nodeID[:], port, d.log)

	tcpAddr, err := net.ResolveTCPAddr("tcp", "163.172.36.173:9500")
	if err != nil {
		d.log.Error(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		d.log.Error(err)
	}

	if err = conn.SetKeepAlivePeriod(time.Minute); err != nil {
		d.log.Warn(err)
	}

	if err = conn.SetKeepAlive(true); err != nil {
		d.log.Warn(err)
	}

	hook, err := logrustash.NewHookWithFieldsAndConn(conn, "peer_node", logrus.Fields{
		"queryType":         "peer_node",
		"startingTimestamp": time.Now(),
	})
	if err != nil {
		d.log.Error(err)
	}

	d.log.Hooks.Add(hook)

	d.p.Listen()
	d.p.Join(bootStrapIp + ":44460")
	fmt.Println("nodeIP = ", d.p.GetIP())
}

func (d *PeerNode) EventLoop() {
	ticker := time.NewTicker(5 * time.Second)
L:
	for {
		select {
		case <-ticker.C:
			if d.requestIsReady() {
				d.tStrategy.StartTest(d)
				ticker.Stop()
			}
		case <-d.done:
			fmt.Println("EventLoop done")
			break L
		//event from peer
		case msg := <-d.peerEvent:
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
				} else if content.Ctype == internalMsg.Cmd_STARTTEST {
					d.tStrategy.StartTest(d)
				} else if content.Ctype == internalMsg.Cmd_TESTDONE {
					go func() {
						d.done <- true
					}()
				} else {
					sender := string(msg.Sender)
					d.tStrategy.CheckResult(sender, content, d)
				}
			default:
			}
		default:
		}
	}
	os.Exit(0)
}
