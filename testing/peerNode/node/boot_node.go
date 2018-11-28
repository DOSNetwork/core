package node

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"github.com/DOSNetwork/core/p2p"
	"github.com/golang/protobuf/proto"

	"github.com/gorilla/mux"
)

type BootNode struct {
	node
	count     int
	ipIdMap   map[string][]byte
	lock      sync.Mutex
	checkroll map[string]bool
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

func (b *BootNode) Init(port, peerSize int) {
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

func (b *BootNode) getMembers(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= b.peerSize; i++ {
		w.Write(b.members[i])
	}
}

func (b *BootNode) getID(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	ip := string(body)
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
	//TODO:Send addresses and IDs after all peer node already build p2p connection to boot node
	if b.count == b.peerSize {
		go b.sendPeerAddresses()
	}
}

type Result struct {
	id string
}

func (b *BootNode) postHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	b.lock.Lock()
	delete(b.checkroll, string(body))
	b.lock.Unlock()
	fmt.Println("check done ", len(b.checkroll))
	if len(b.checkroll) == 0 {
		b.done <- true
	}
}

func (b *BootNode) sendPeerAddresses() {
	s := []string{}
	for i := 0; i < len(b.allIP); i++ {
		id, err := b.p.NewPeer(b.allIP[i])
		if err != nil {
			fmt.Println("BootNode err ", err)
		} else {
			fmt.Println("NewPeer id ", id, " address ", b.allIP[i])
			s = append(s, b.allIP[i])
		}
	}
	fmt.Println("sendPeerAddresses ", s)

	allIP := strings.Join(s, ",")
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_ALLIP,
		Args:  allIP,
	}
	pb := proto.Message(cmd)
	b.p.Broadcast(pb)
}
