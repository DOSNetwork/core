package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/p2p"
	//	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

/*
The purpose of meeting is to test findNode and sendMessageById

*/
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

type node struct {
	p         p2p.P2PInterface
	peerEvent chan p2p.P2PMessage
	members   [][]byte
	allIP     []string
	peerSize  int
	done      chan bool
}

type bootNode struct {
	node
	count     int
	ipIdMap   map[string][]byte
	lock      sync.Mutex
	checkroll map[string]bool
}

type peerNode struct {
	node
	bootStrapIp string
	nodeID      []byte
	checkCount  int
	findNodeDur time.Duration
	checkroll   map[string]int
	numMessages int
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (n *node) EventLoop() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			//PrintMemUsage()
		//event from peer
		case _ = <-n.peerEvent:
		case <-n.done:
			//os.Exit(0)
			break
		}
	}
}

func (b *bootNode) Init(port, peerSize int) {
	//1)Generate member ID
	b.peerSize = peerSize
	b.members = [][]byte{}
	b.allIP = []string{}
	bootID := []byte{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1}
	b.members = append(b.members, bootID)
	b.checkroll = make(map[string]bool)
	b.ipIdMap = make(map[string][]byte)
	for i := 1; i <= b.peerSize; i++ {
		id, _ := GenerateRandomBytes(len(bootID))
		for b.checkroll[string(id)] {
			id, _ = GenerateRandomBytes(len(bootID))
		}
		b.checkroll[string(id)] = true
		b.members = append(b.members, id)
	}

	//2)Declare a new router to handle REST API call
	r := mux.NewRouter()
	r.HandleFunc("/getID", b.getID).Methods("POST")
	r.HandleFunc("/getMembers", b.getMembers).Methods("GET")
	r.HandleFunc("/post", b.postHandler)
	go http.ListenAndServe(":8080", r)

	//3)Build a p2p network
	b.p, b.peerEvent, _ = p2p.CreateP2PNetwork(bootID, port)
	go b.p.Listen()
}

func (b *bootNode) getMembers(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= b.peerSize; i++ {
		w.Write(b.members[i])
	}
}

func (b *bootNode) getID(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	ip := string(body)
	fmt.Println("getID count ", b.count, " ip", ip, " members", b.ipIdMap[ip])
	if b.ipIdMap[ip] == nil {
		b.allIP = append(b.allIP, ip)
		b.count++
		if b.count > b.peerSize {
			fmt.Println("getID over size error!!!")
			return
		}
		fmt.Println("getID count ", b.count, " peerSize", b.peerSize, " members", b.node.members[b.count])
		b.ipIdMap[ip] = b.node.members[b.count]
	}

	if len(b.node.members) > b.count {
		w.Write(b.ipIdMap[ip])
	} else {
		fmt.Println(b.node.members)
	}
	if b.count == b.peerSize {
		go b.sendPeerAddresses()
	}
}

type Result struct {
	id string
}

func (b *bootNode) postHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	b.lock.Lock()
	delete(b.checkroll, string(body))
	b.lock.Unlock()
	fmt.Println("check done ", len(b.checkroll))
	if len(b.checkroll) == 0 {
		b.done <- true
	}
}

func (b *bootNode) sendPeerAddresses() {
	fmt.Println("sendPeerAddresses ", b.members)
	s := []string{}
	for i := 0; i < len(b.allIP); i++ {
		id, err := b.p.NewPeer(b.allIP[i])
		if err != nil {
			fmt.Println("bootNode err ", err)
		} else {
			fmt.Println("NewPeer id ", id, " address ", b.allIP[i])
			s = append(s, b.allIP[i])
		}
	}
	allIP := strings.Join(s, ",")
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_ALLIP,
		Args:  allIP,
	}
	pb := proto.Message(cmd)
	b.p.Broadcast(pb)
}

func (d *peerNode) MakeRequest(bootStrapIp, f string, args []byte) ([]byte, error) {

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

func (d *peerNode) Init(bootStrapIp string, port, peerSize int) {
	var err error
	d.peerSize = peerSize
	d.checkCount = 1
	d.bootStrapIp = bootStrapIp
	d.checkroll = make(map[string]int)
	d.numMessages, err = strconv.Atoi(os.Getenv("NUMOFMESSAGS"))
	fmt.Println("!!!!!! d.numMessages ", d.numMessages)
	if err != nil {
		fmt.Println(err)
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

func (d *peerNode) EventLoop() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			//PrintMemUsage()
		//event from peer
		case msg := <-d.peerEvent:
			switch content := msg.Msg.Message.(type) {
			case *internalMsg.Cmd:
				if content.Ctype == internalMsg.Cmd_ALLIP {
					//msg.Sender
					cmd := &internalMsg.Cmd{
						Ctype: internalMsg.Cmd_SIGNIN,
						Args:  "check",
					}
					pb := proto.Message(cmd)
					nodes := strings.Split(content.Args, ",")
					allIP := []string{}
					ip, _ := p2p.GetLocalIp()
					ip = ip + ":44460"
					for _, node := range nodes {
						if ip != node {
							allIP = append(allIP, node)
						}
					}
					for i := 0; i < len(allIP); i++ {
						ip := allIP[i]
						id, _ := d.p.NewPeer(ip)
						d.checkroll[string(id)] = 0
					}
					for i := 0; i < d.numMessages; i++ {
						for id, _ := range d.checkroll {
							go d.p.SendMessageById([]byte(id), pb)
						}
					}
				} else if content.Ctype == internalMsg.Cmd_SIGNIN {
					sender := string(msg.Sender)
					d.checkroll[sender] = d.checkroll[sender] + 1
					fmt.Println("sender", []byte(sender), "  ", d.checkroll[sender], content.Ctype, " ", len(d.checkroll))
					if d.checkroll[sender] == d.numMessages {
						delete(d.checkroll, sender)
					}

					if len(d.checkroll) == 0 {
						d.MakeRequest(d.bootStrapIp, "post", d.nodeID)
					}
				}
			default:
				fmt.Println(content)
			}
		}
	}
}

// main
func main() {
	var ip string
	var err error
	var peerSize int
	debug.FreeOSMemory()
	//1)Load config
	offChainConfig := configuration.OffChainConfig{}
	offChainConfig.LoadConfig()
	port := offChainConfig.Port
	//It also need to connect to bootstrape node to get crential
	bootStrapIP := os.Getenv("BOOTSTRAPIP")
	noderole := os.Getenv("NODEROLE")
	peerSize, err = strconv.Atoi(os.Getenv("PEERSIZE"))
	if err != nil {
		fmt.Println(err)
	}

	//boot node
	if noderole == "boot" {
		b := new(bootNode)
		b.Init(port, peerSize)
		b.EventLoop()
	} else {
		s := strings.Split(bootStrapIP, ":")
		ip, _ = s[0], s[1]
		d := new(peerNode)
		d.Init(ip, port, peerSize)
		d.EventLoop()
	}
}
