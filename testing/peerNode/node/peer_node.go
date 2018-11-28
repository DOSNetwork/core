package node

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

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

func (d *PeerNode) Init(bootStrapIp string, port, peerSize int, numMessages int, tStrategy string) {
	d.peerSize = peerSize
	d.checkCount = 1
	d.bootStrapIp = bootStrapIp
	d.checkroll = make(map[string]int)
	d.numMessages = numMessages
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

	//2)Build a p2p network
	d.p, d.peerEvent, _ = p2p.CreateP2PNetwork(d.nodeID[:], port)
	go d.p.Listen()
}

func (d *PeerNode) EventLoop() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			//PrintMemUsage()
		//event from peer
		case msg := <-d.peerEvent:
			switch content := msg.Msg.Message.(type) {
			case *internalMsg.Cmd:
				fmt.Println("!! content.Ctype ", content.Ctype)
				if content.Ctype == internalMsg.Cmd_ALLIP {
					nodes := strings.Split(content.Args, ",")
					allIP := []string{}
					ip, _ := p2p.GetLocalIp()
					ip = ip + ":44460"
					for _, node := range nodes {
						if ip != node {
							allIP = append(allIP, node)
						}
					}
					d.nodeIPs = allIP
					d.tStrategy.StartTest(content, d)
				} else if content.Ctype == internalMsg.Cmd_ALLID {
					nodes := strings.Split(content.Args, ",")
					fmt.Println("!! Cmd_ALLID ", nodes)
					allID := [][]byte{}
					for _, node := range nodes {
						id := []byte(node)
						if !bytes.Equal(d.nodeID, id) {
							allID = append(allID, id)
						}
					}
					d.nodeIDs = allID
					d.tStrategy.StartTest(content, d)
				} else {
					sender := string(msg.Sender)
					d.tStrategy.CheckResult(sender, content, d)
				}
			default:
				fmt.Println(content)
			}
		}
	}
}
