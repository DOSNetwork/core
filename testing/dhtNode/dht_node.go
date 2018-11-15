package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/testing/dhtNode/internalMsg"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

/*
The purpose of meeting is to test findNode and sendMessageById
1)What is the average time for findNode
2)What is the average time for SendMessageById.
3)How many times of cant' find node when it call SendMessageById
4)The pass criteria for this test is that all nodes should receive check message from all members.
*/

type node struct {
	p         p2p.P2PInterface
	peerEvent chan p2p.P2PMessage
	members   [][]byte
	dhtSize   int
	checkroll map[string]bool
	done      chan bool
}

type bootNode struct {
	node
	count int
	lock  sync.Mutex
}

type dhtNode struct {
	node
	bootStrapIp string
	nodeID      []byte
	checkCount  int
	findNodeDur time.Duration
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

func (n *node) Init(port, dhtSize int) {
	//fmt.Println("node Init", n.members)
	n.done = make(chan bool)
	n.dhtSize = dhtSize
	//Build a p2p network
	n.peerEvent = make(chan p2p.P2PMessage, 100)
	n.p, _ = p2p.CreateP2PNetwork(n.peerEvent, port)
}

func (n *node) EventLoop() {
	for {
		select {
		//event from peer
		case _ = <-n.peerEvent:
		case <-n.done:
			//os.Exit(0)
			break
		}
	}
}

func (b *bootNode) Init(port, dhtSize int) {
	//Generate member ID
	b.dhtSize = dhtSize
	b.members = [][]byte{}
	bootID := []byte{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1}
	b.members = append(b.members, bootID)
	b.checkroll = make(map[string]bool)
	for i := 1; i <= b.dhtSize; i++ {
		id, _ := GenerateRandomBytes(len(bootID))
		for b.checkroll[string(id)] {
			id, _ = GenerateRandomBytes(len(bootID))
		}
		b.checkroll[string(id)] = true
		b.members = append(b.members, id)
		fmt.Println("i = ", i, " ", id)
	}

	fmt.Println("bootNode members ", b.members)
	b.node.Init(port, dhtSize)
	b.p.SetId(b.node.members[0])
	go b.p.Listen()

	//Declare a new router to handle REST API call
	r := mux.NewRouter()
	r.HandleFunc("/getID", b.getID).Methods("GET")
	r.HandleFunc("/getMembers", b.getMembers).Methods("GET")
	r.HandleFunc("/post", b.postHandler)
	go http.ListenAndServe(":8080", r)
}

func (b *bootNode) getMembers(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= b.dhtSize; i++ {
		w.Write(b.members[i])
	}
}

func (b *bootNode) getID(w http.ResponseWriter, r *http.Request) {
	b.count++
	fmt.Println("getID count ", b.count, " dhtSize", b.dhtSize, " members", b.node.members[b.count])

	if len(b.node.members) > b.count {
		w.Write(b.node.members[b.count])
	} else {
		fmt.Println(b.node.members)
	}
	if b.count == b.dhtSize {
		go b.requestGrouping()
	}
}

type Result struct {
	id string
}

func (b *bootNode) postHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	delete(b.checkroll, string(body))
	fmt.Println("check done ", len(b.checkroll))
	if len(b.checkroll) == 0 {
		b.done <- true
	}
}

func (b *bootNode) requestGrouping() {
	fmt.Println("requestGrouping ", b.members)
	time.Sleep(5 * time.Second)
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_FINDNODE,
		Args:  "",
	}
	pb := proto.Message(cmd)
	for i, member := range b.members {
		if i != 0 {
			b.p.SendMessageById(member, pb)
		}
	}
}

func (d *dhtNode) Init(bootStrapIp string, port, dhtSize int) {
	d.node.Init(port, dhtSize)
	d.dhtSize = dhtSize
	d.checkCount = 1
	d.bootStrapIp = bootStrapIp
	//Wait until bootstrap node assign an ID
	for {
		tServer := "http://" + bootStrapIp + ":8080/getID"
		resp, err := http.Get(tServer)
		if err != nil {
			fmt.Println(err)
		}

		for err != nil {
			time.Sleep(10 * time.Second)
			resp, err = http.Get(tServer)
		}

		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(" dhtNode getID, r ", len(r))
			d.nodeID = r
			d.p.SetId(d.nodeID[:])
			fmt.Println(" dhtNode getID, ", len(d.nodeID))
			break
		}
	}

	fmt.Println("nodeID = ", d.p.GetId().Id)
	d.checkroll = make(map[string]bool)
	tServer := "http://" + bootStrapIp + ":8080/getMembers"
	resp, err := http.Get(tServer)
	if err != nil {
		fmt.Println(err)
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("!!!! members len ", len(r))
		idSize := len(d.p.GetId().Id)
		for i := 0; i < d.dhtSize; i++ {
			d.members = append(d.members, r[i*idSize:(i+1)*idSize])
		}
	}
	fmt.Println("=============================")
	for _, member := range d.members {
		fmt.Println(member)
	}
	fmt.Println("=============================")

	//Start to listen incoming connection
	go d.p.Listen()
}

func (d *dhtNode) MakeRequest(bootStrapIp string) {

	tServer := "http://" + bootStrapIp + ":8080/post"

	req, err := http.NewRequest("POST", tServer, bytes.NewBuffer(d.nodeID))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
func (d *dhtNode) EventLoop() {
	for {
		select {
		//event from peer
		case msg := <-d.peerEvent:
			switch content := msg.Msg.Message.(type) {
			case *internalMsg.Cmd:
				if content.Ctype == internalMsg.Cmd_FINDNODE {
					//msg.Sender
					cmd := &internalMsg.Cmd{
						Ctype: 2,
						Args:  "check",
					}
					pb := proto.Message(cmd)
					for i := 0; i < len(d.members); i++ {
						member := d.members[i]
						if string(member[:]) != string(d.p.GetId().Id) {
							d.p.SendMessageById(member, pb)
						}
					}
				} else {
					sender := string(msg.Sender)
					if !d.checkroll[sender] {
						d.checkroll[sender] = true
						d.checkCount++
						if d.checkCount == d.dhtSize {
							d.MakeRequest(d.bootStrapIp)
							fmt.Println("done check")
							//os.Exit(0)
						}
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
	var dhtSize int
	//1)Load config
	offChainConfig := configuration.OffChainConfig{}
	offChainConfig.LoadConfig()
	port := offChainConfig.Port
	//It also need to connect to bootstrape node to get crential
	bootStrapIP := os.Getenv("BOOTSTRAPIP")
	noderole := os.Getenv("NODEROLE")
	dhtSize, err = strconv.Atoi(os.Getenv("DHTSIZE"))
	if err != nil {
		fmt.Println(err)
	}

	//boot node
	if noderole == "boot" {
		b := new(bootNode)
		b.Init(port, dhtSize)
		b.EventLoop()
	} else {
		s := strings.Split(bootStrapIP, ":")
		ip, _ = s[0], s[1]
		d := new(dhtNode)
		d.Init(ip, port, dhtSize)
		//TODO : Fix race conditation between Listen and CreatePeer
		time.Sleep(2 * time.Second)
		d.p.CreatePeer(bootStrapIP, nil)
		results := d.p.FindNode(d.p.GetId(), dht.BucketSize, 20)
		for _, result := range results {
			d.p.GetRoutingTable().Update(result)
			//fmt.Println(d.p.GetId().Address, "Update peer: ", result.Address)
		}
		d.EventLoop()
	}
}
