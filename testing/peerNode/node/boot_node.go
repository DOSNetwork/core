package node

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"

	"time"

	"github.com/bshuster-repo/logrus-logstash-hook"
	//	"github.com/ethereum/go-ethereum/common"

	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"github.com/DOSNetwork/core/p2p"
	//	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/golang/protobuf/proto"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const ALLNODEREADY = 1
const ALLNODENOTREADY = 0

type BootNode struct {
	node
	count     int
	ipIdMap   map[string][]byte
	lock      sync.Mutex
	checkroll map[string]bool
	readynode map[string]bool
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

func (b *BootNode) Init(port, peerSize int, logger *logrus.Logger) {
	//1)Generate member ID
	b.peerSize = peerSize
	b.members = [][]byte{}
	b.allIP = []string{}
	bootID := []byte{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1}
	//bootID := []byte{152, 37, 111, 123, 113, 163, 23, 27, 114, 171, 254, 13, 147, 217, 232, 54, 148, 166, 46, 131}
	b.members = append(b.members, bootID)
	b.checkroll = make(map[string]bool)
	b.ipIdMap = make(map[string][]byte)
	b.done = make(chan bool)
	b.readynode = make(map[string]bool)
	b.log = logger

	for i := 1; i <= b.peerSize; i++ {
		id, _ := GenerateRandomBytes(len(bootID))
		for b.checkroll[string(id)] {
			id, _ = GenerateRandomBytes(len(bootID))
		}
		b.checkroll[string(id)] = true
		b.lock.Lock()
		b.members = append(b.members, id)
		b.lock.Unlock()
	}

	//2)Declare a new router to handle REST API call
	r := mux.NewRouter()
	r.HandleFunc("/getID", b.getID).Methods("POST")
	r.HandleFunc("/getMembers", b.getMembers).Methods("GET")
	r.HandleFunc("/getAllIDs", b.getAllIDs).Methods("POST")
	r.HandleFunc("/getAllIPs", b.getAllIPs).Methods("POST")
	r.HandleFunc("/isTestReady", b.isTestReady).Methods("POST")
	r.HandleFunc("/post", b.postHandler)
	go http.ListenAndServe(":8080", r)

	//3)Build a p2p network
	b.p, b.peerEvent, _ = p2p.CreateP2PNetwork(bootID, port, b.log)
	tcpAddr, err := net.ResolveTCPAddr("tcp", "163.172.36.173:9500")
	if err != nil {
		b.log.Error(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		b.log.Error(err)
	}

	if err = conn.SetKeepAlivePeriod(time.Minute); err != nil {
		b.log.Warn(err)
	}

	if err = conn.SetKeepAlive(true); err != nil {
		b.log.Warn(err)
	}

	hook, err := logrustash.NewHookWithFieldsAndConn(conn, "peer_node", logrus.Fields{
		"queryType":         "peer_node",
		"startingTimestamp": time.Now(),
	})
	if err != nil {
		b.log.Error(err)
	}

	b.log.Hooks.Add(hook)
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
	b.log.WithFields(logrus.Fields{
		"bootGetID": true,
	}).Info()
	if len(b.node.members) > b.count {
		w.Write(b.ipIdMap[ip])
	} else {
		fmt.Println(b.node.members)

	}
	//TODO:Send addresses and IDs after all peer node already build p2p connection to boot node
	//if b.count == b.peerSize {
	//	go b.sendPeerAddresses()
	//	go b.sendPeerIDs()
	//}
}

func (b *BootNode) getAllIDs(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("getAllIDs ,", b.peerSize)
	for i := 1; i <= b.peerSize; i++ {
		w.Write(b.members[i])
	}
}

func (b *BootNode) getAllIPs(w http.ResponseWriter, r *http.Request) {
	if b.count == b.peerSize {
		for i := 0; i < len(b.allIP); i++ {
			w.Write([]byte(b.allIP[i] + ","))
		}
	}
}

func (b *BootNode) isAllnodeready() bool {
	//fmt.Println("isAllnodeready ", len(b.readynode) == b.peerSize, "", len(b.readynode))
	return len(b.readynode) == b.peerSize
}

func (b *BootNode) isTestReady(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	ip := string(body)
	b.lock.Lock()
	if b.readynode[ip] == false {
		b.readynode[ip] = true
	}
	b.lock.Unlock()
	if b.isAllnodeready() {
		w.Write([]byte{ALLNODEREADY})
		b.log.WithFields(logrus.Fields{
			"bootTestReady": true,
		}).Info()
	} else {
		w.Write([]byte{ALLNODENOTREADY})
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

	b.log.WithFields(logrus.Fields{
		"bootTestDone": true,
	}).Info()

	if len(b.checkroll) == 0 {
		b.finishTest()
		b.done <- true
	}
}

func (b *BootNode) startTest() {
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_STARTTEST,
		Args:  []byte{},
	}
	pb := proto.Message(cmd)
	for i := 1; i <= b.peerSize; i++ {
		b.p.SendMessage(b.members[i], pb)
	}
}

func (b *BootNode) finishTest() {
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_TESTDONE,
		Args:  []byte{},
	}
	pb := proto.Message(cmd)
	for i := 1; i <= b.peerSize; i++ {
		b.p.SendMessage(b.members[i], pb)
	}
}

func (b *BootNode) EventLoop() {
L:
	for {
		select {
		//event from peer
		case _ = <-b.peerEvent:
		case <-b.done:
			fmt.Println("EventLoop done")
			break L
		default:
		}
	}
	os.Exit(0)
}
