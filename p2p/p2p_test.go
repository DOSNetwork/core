package p2p

import (
	"testing"

	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/sirupsen/logrus"
	"math/rand"
	"fmt"
)

const peerSize =105
const connSize =100

//test scenario:
func TestBootStrap(t *testing.T) {
	bootid := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	peerid := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 19}
	bootPort := 44460
	peerPort := 44459
	log := logrus.New()

	boot, _, err := CreateP2PNetwork(bootid[:], bootPort, log)
	if err != nil {
		t.Error("Test Failed: CreateP2PNetwork failed ", err)
	}

	if err := boot.Listen(); err != nil {
		t.Error("Test Failed: Listen failed ", err)

	}

	peer, _, err := CreateP2PNetwork(peerid[:], peerPort, log)
	if err != nil {
		t.Error("Test Failed: CreateP2PNetwork failed ", err)
	}
	if err := peer.Listen(); err != nil {
		t.Error("Test Failed: Listen failed ", err)
	}

	boot_p2p, ok := boot.(*P2P)
	if !ok {
		t.Error("Test Failed: ")
	}
	peer_p2p, ok := peer.(*P2P)
	if !ok {
		t.Error("Test Failed: ")
	}
	bContactsBefore := boot_p2p.routingTable.GetPeers()
	pContactsBefore := peer_p2p.routingTable.GetPeers()

	if len(pContactsBefore) != 0 || len(pContactsBefore) != 0 {
		t.Errorf("Test Failed: Routing Table should be empty boot(%d) peer(%d)", len(bContactsBefore), len(pContactsBefore))
	}

	peer.Join(boot.GetIP())

	bContactsAfter := boot_p2p.routingTable.GetPeers()
	pContactsAfter := peer_p2p.routingTable.GetPeers()

	if len(bContactsAfter) != 1 || bContactsAfter[string(peerid)] != peer.GetIP() {
		t.Errorf("boot contacts(%d) : peer contact %s peer ip %s", len(bContactsAfter), bContactsAfter[string(peerid)], peer.GetIP())
	}
	if len(pContactsAfter) != 1 || pContactsAfter[string(bootid)] != boot.GetIP() {
		t.Errorf("peer contacts(%d) : peer contact %s peer ip %s", len(pContactsAfter), pContactsAfter[string(bootid)], boot.GetIP())
	}

	//TODO:Implement Leave function to turn off connection

}

//test scenario:numOfNodes is dht.BucketSize ,so any of nodeID should existing in every routing table.
func TestBootStrapWithMultipleNode(t *testing.T) {
	numOfNodes := dht.BucketSize
	log := logrus.New()
	var ok bool
	peerPort := 55550
	peers := make([]*P2P, numOfNodes)
	peerIDs := make([][]byte, numOfNodes)

	peerIDs[0] = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	peer, _, err := CreateP2PNetwork(peerIDs[0][:], peerPort, log)
	if err := peer.Listen(); err != nil {
		t.Error("Test Failed: Listen failed ", err)
	}
	peers[0], ok = peer.(*P2P)
	if err != nil {
		t.Error("Test Failed: CreateP2PNetwork failed ", err)
	}
	if !ok {
		t.Error("Test Failed: ")
	}
	for i := 1; i < numOfNodes; i++ {
		peerIDs[i] = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12 + byte(i), 13, 14, 15, 16, 17, 18, 19, 21 + byte(i)}
		peerPort := 55550 + i
		peer, _, err := CreateP2PNetwork(peerIDs[i][:], peerPort, log)
		if err := peer.Listen(); err != nil {
			t.Error("Test Failed: Listen failed ", err)
		}
		peers[i], ok = peer.(*P2P)
		if err != nil {
			t.Error("Test Failed: CreateP2PNetwork failed ", err)
		}
		if !ok {
			t.Error("Test Failed: ")
		}
		peer.Join(peers[0].GetIP())
		count := 0
		for j := 0; j < i; j++ {
			pContactsAfter := peers[j].routingTable.GetPeers()
			if pContactsAfter[string(peerIDs[i])] == peers[i].GetIP() {
				count++
			}
		}
		if count == 0 {
			t.Errorf("Can't find %s", peers[i].GetIP())
		}
	}
	for i := 0; i < numOfNodes; i++ {
		count := 0
		for j := 0; j < numOfNodes; j++ {
			if i != j {
				pContactsAfter := peers[j].routingTable.GetPeers()
				if pContactsAfter[string(peerIDs[i])] == peers[i].GetIP() {
					count++
				}
			}
		}
		if count != (numOfNodes - 1) {
			t.Errorf("Can't find %s in all routing table, %d", peers[i].GetIP(), count)
		}
	}
}



//test scenario: test max connection for single node
func TestMaxConnectionBySingleNode (t *testing.T){
	bootId := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	bootPort := 33333
	log := logrus.New()
	//peers := make([]*P2P, peerSize)
	peerIds := make([][]byte, peerSize)

	boot, _, err := CreateP2PNetwork(bootId[:], bootPort, log)
	if err != nil {
		t.Error("Test Failed: CreateP2PNetwork failed ", err)
	}
	boot_p2p, _ := boot.(*P2P)

	if err := boot.Listen(); err != nil {
		t.Error("Test Failed: Listen failed ", err)

	}

	peerIds[0] = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i:=1;i<peerSize;i++{
		peerIds[i] = []byte{1+byte(i), 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12 , 13, 14, 15, 16, 17, 18, 19, 21 + byte(rand.Intn(10))}
		peerport := 22222+i

		peer,_,err := CreateP2PNetwork(peerIds[i],peerport,log)
		if err != nil{
			t.Error("Test Failed: CreateP2PNetwork failed",err)
		}

		if err:=peer.Listen();err !=nil{
			t.Error("Test Failed: Listen failed",err)
		}

		peer.Join(boot.GetIP())

		fmt.Println("aaaaaaaa------>")

		if i>connSize  && boot_p2p.peers.PeerNum() > connSize{
			boot_p2p.peers.DeletePeer(boot_p2p.peers.parray[0].identity.Address)
			break
		}
	}

	if boot_p2p.peers.PeerNum() >= connSize{
		t.Fatal("test failed ,connection out of indicated range")
	}
}