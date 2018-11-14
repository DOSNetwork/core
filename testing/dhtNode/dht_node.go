package main

import (
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

type node struct {
	p         p2p.P2PInterface
	peerEvent chan p2p.P2PMessage
	members   [][]byte
	dhtSize   int
}

type bootNode struct {
	node
	count int
	lock  sync.Mutex
}

type dhtNode struct {
	node
	nodeID      []byte
	checkCount  int
	checkroll   map[string]bool
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
		}
	}
}

func (b *bootNode) Init(port, dhtSize int) {
	//Generate member ID
	b.dhtSize = dhtSize
	b.members = [][]byte{}
	bootID := []byte{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1}
	existing := make(map[string]bool)
	b.members = append(b.members, bootID)
	for i := 1; i <= b.dhtSize; i++ {
		id, _ := GenerateRandomBytes(len(bootID))
		for existing[string(id)] {
			id, _ = GenerateRandomBytes(len(bootID))
		}
		existing[string(id)] = true
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
					//					start := time.Now()
					for i := 0; i < len(d.members); i++ {
						member := d.members[i]
						if string(member[:]) != string(d.p.GetId().Id) {
							d.p.SendMessageById(member, pb)
						}
					}
					//d.findNodeDur = time.Since(start)
					//d := float64(d.findNodeDur/time.Millisecond) / float64(len(d.members))
					//fmt.Println("===================================================")
					//fmt.Println("From receiving FINDNODE to boradcast", d)
					//fmt.Println("===================================================")

				} else {
					sender := string(msg.Sender)
					if !d.checkroll[sender] {
						d.checkroll[sender] = true
						d.checkCount++
						if d.checkCount == d.dhtSize {
							fmt.Println("===================================================")
							fmt.Println("check all done")
							fmt.Println("===================================================")
						}
						//fmt.Println("receive check from ", msg.Sender, " ", d.checkCount, " ", d.dhtSize, " ", d.checkroll[sender])

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
			fmt.Println(d.p.GetId().Address, "Update peer: ", result.Address)
		}
		d.EventLoop()
	}
}
