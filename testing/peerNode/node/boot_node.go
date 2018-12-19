package node

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	//	"github.com/ethereum/go-ethereum/common"

	//	"github.com/DOSNetwork/core/testing/peerNode/internalMsg"

	"github.com/DOSNetwork/core/p2p"
	//	"github.com/DOSNetwork/core/p2p/dht"
	//	"github.com/golang/protobuf/proto"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const ALLNODEREADY = 1
const ALLNODENOTREADY = 0
const ALLNODEFINISH = 1
const ALLNODENOTFINISH = 0

type BootNode struct {
	node
	count          int
	testReadyCount int
	ipIdMap        map[string][]byte
	lock           sync.Mutex
	checkroll      map[string]bool
	readynode      map[string]bool
	finishnode     map[string]bool
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

func (b *BootNode) Init(port, peerSize int, logger *logrus.Entry) {
	//1)Generate member ID
	b.peerSize = peerSize
	b.members = [][]byte{}
	b.allIP = []string{}
	bootID := []byte{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1}
	//bootID := []byte{152, 37, 111, 123, 113, 163, 23, 27, 114, 171, 254, 13, 147, 217, 232, 54, 148, 166, 46, 131}
	//	b.members = append(b.members, bootID)

	b.checkroll = make(map[string]bool)
	b.ipIdMap = make(map[string][]byte)
	b.done = make(chan bool)
	b.readynode = make(map[string]bool)
	b.finishnode = make(map[string]bool)
	b.log = logger
	b.log.Data["role"] = "bootstrap"

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
	r.HandleFunc("/getAllIDs", b.getAllIDs).Methods("POST")
	r.HandleFunc("/getAllIPs", b.getAllIPs).Methods("POST")
	r.HandleFunc("/isTestReady", b.isTestReady).Methods("POST")
	r.HandleFunc("/isTestFinish", b.isTestFinish).Methods("POST")
	r.HandleFunc("/post", b.postHandler)
	go http.ListenAndServe(":8080", r)

	//3)Build a p2p network
	b.p, b.peerEvent, _ = p2p.CreateP2PNetwork(bootID, port, b.log)

	go b.p.Listen()
}

func (b *BootNode) getMembers(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < b.peerSize; i++ {
		w.Write(b.members[i])
	}
}

func (b *BootNode) getID(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	ip := string(body)
	if b.ipIdMap[ip] == nil {

		if b.count >= b.peerSize {
			fmt.Println("getID over size error!!!")
			return
		}
		b.allIP = append(b.allIP, ip)
		fmt.Println("getID count ", b.count, " peerSize", b.peerSize, " len(b.allIP)", len(b.allIP))
		b.ipIdMap[ip] = b.node.members[b.count]
		b.count++
		b.log.WithFields(logrus.Fields{
			"eventGetID": b.count,
		}).Info()
	}

	w.Write(b.ipIdMap[ip])

	//TODO:Send addresses and IDs after all peer node already build p2p connection to boot node
	//if b.count == b.peerSize {
	//	go b.sendPeerAddresses()
	//	go b.sendPeerIDs()
	//}
}

func (b *BootNode) getAllIDs(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("getAllIDs ,", b.peerSize)
	//d.log.WithField("requestAllIds", len(b.allIP)).Info()
	for i := 0; i < b.peerSize; i++ {
		w.Write(b.members[i])
	}
}

func (b *BootNode) getAllIPs(w http.ResponseWriter, r *http.Request) {
	if len(b.allIP) >= b.peerSize {
		for i := 0; i < b.peerSize; i++ {
			w.Write([]byte(b.allIP[i] + ","))
		}
	}
}

func (b *BootNode) isAllnodeready() bool {
	fmt.Println("isAllnodeready ", len(b.readynode) >= b.peerSize, "", len(b.readynode))
	return len(b.readynode) >= b.peerSize
}

func (b *BootNode) isAllnodefinish() bool {
	fmt.Println("isAllnodefinish ", len(b.finishnode) >= b.peerSize, "", len(b.finishnode))
	return len(b.finishnode) >= b.peerSize
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
		b.log.WithFields(logrus.Fields{
			"eventIsAllnodeready": len(b.readynode),
		}).Info()
		w.Write([]byte{ALLNODEREADY})
		b.testReadyCount++
		if b.testReadyCount >= b.peerSize {
			b.done <- true
		}
	} else {
		w.Write([]byte{ALLNODENOTREADY})
	}
}

func (b *BootNode) isTestFinish(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	ip := string(body)
	b.lock.Lock()
	if b.finishnode[ip] == false {
		b.finishnode[ip] = true
	}
	b.lock.Unlock()

	if b.isAllnodefinish() {
		b.log.WithFields(logrus.Fields{
			"eventIsAllnodefinish": len(b.finishnode),
		}).Info()
		w.Write([]byte{ALLNODEFINISH})
	} else {
		w.Write([]byte{ALLNODENOTFINISH})
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
		"eventTestDone": len(b.checkroll),
	}).Info()

	if len(b.checkroll) == 0 {
		//		b.finishTest()
		b.done <- true
	}
}

func (b *BootNode) EventLoop() {
	for {
		select {
		//event from peer
		case _ = <-b.peerEvent:
		case <-b.done:
			fmt.Println("EventLoop done")
			b.log.WithField("event", "EventLoop done").Info()
		default:
		}
	}
}