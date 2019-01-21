package node

import (
	"crypto/rand"
	"encoding/binary"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"

	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	"github.com/gorilla/mux"
)

const (
	ALLNODEREADY     = 1
	ALLNODENOTREADY  = 0
	DKGROUNDFINISH   = 2
	ALLNODEFINISH    = 1
	ALLNODENOTFINISH = 0
)

var dkgRound = 0

type BootNode struct {
	node
	count           int
	testFinishCount int
	ipIdMap         map[string][]byte
	lock            sync.Mutex
	checkroll       map[string]bool
	readynode       map[string]bool
	finishnode      map[string]bool
	dkgProgress     map[string]uint16
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
	//0)prepare round count for dkg test
	dkgRoundStr := os.Getenv("DKGROUND")
	if len(dkgRoundStr) > 0 {
		var err error
		dkgRound, err = strconv.Atoi(dkgRoundStr)
		if err != nil {
			//b.log.Fatal(err)
		}
	}
	//1)Generate member ID
	b.testFinishCount = 0
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
	b.dkgProgress = make(map[string]uint16)

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
	r.HandleFunc("/isNextRoundReady", b.isNextRoundReady).Methods("POST")
	r.HandleFunc("/isTestFinish", b.isTestFinish).Methods("POST")
	r.HandleFunc("/post", b.postHandler)
	go http.ListenAndServe(":8080", r)

	//3)Build a p2p network
	b.p, _ = p2p.CreateP2PNetwork(bootID, port)

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
			return
		}
		b.allIP = append(b.allIP, ip)
		//fmt.Println("getID count ", b.count, " peerSize", b.peerSize, " len(b.allIP)", len(b.allIP))
		b.ipIdMap[ip] = b.node.members[b.count]
		b.dkgProgress[ip] = 0
		b.count++
		log.Progress("GetID")
	}

	w.Write(b.ipIdMap[ip])
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
	return len(b.readynode) >= b.peerSize
}

func (b *BootNode) isAllnodefinish() bool {
	return len(b.finishnode) >= b.peerSize
}

func (b *BootNode) isAllNextRoundReady(roundCount uint16) bool {
	b.lock.Lock()
	defer b.lock.Unlock()

	for _, value := range b.dkgProgress {
		if value != roundCount {
			return false
		}
	}
	return true
}

func (b *BootNode) isTestReady(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	ip := string(body)
	b.lock.Lock()
	if b.readynode[ip] == false {
		b.readynode[ip] = true
		log.Progress("TestReady")
	}
	b.lock.Unlock()

	if b.isAllnodeready() {
		w.Write([]byte{ALLNODEREADY})
	} else {
		w.Write([]byte{ALLNODENOTREADY})
	}
}

func (b *BootNode) isNextRoundReady(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	ipBytes := body[:len(body)-2]
	roundCountBytes := body[len(body)-2:]
	ip := string(ipBytes)
	roundCount := binary.LittleEndian.Uint16(roundCountBytes)
	b.lock.Lock()
	b.dkgProgress[ip] = roundCount
	b.lock.Unlock()

	if roundCount == uint16(dkgRound) {
		w.Write([]byte{DKGROUNDFINISH})
	} else if b.isAllNextRoundReady(roundCount) {
		w.Write([]byte{ALLNODEREADY})
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
		w.Write([]byte{ALLNODEFINISH})
		b.testFinishCount++
		if b.testFinishCount >= b.peerSize {
			b.done <- true
		}
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
}

func (b *BootNode) EventLoop() {
L:
	for {
		select {
		//event from peer
		case _ = <-b.peerEvent:
		case <-b.done:
			//fmt.Println("EventLoop done")
			break L
		default:
		}
	}
	time.Sleep(time.Second)
}
