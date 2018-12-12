package dkg

import (
	//"fmt"
	"testing"

	"github.com/DOSNetwork/core/p2p"

	"github.com/sirupsen/logrus"
)

func TestPeerConnEnd(t *testing.T) {
	numOfNodes := 10
	log := logrus.New()
	var ok bool
	peerPort := 55550
	peers := make([]*p2p.P2P, numOfNodes)
	peerIDs := make([][]byte, numOfNodes)
	dkgs := make([]P2PDkgInterface, numOfNodes)

	peerIDs[0] = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	peer, peerEvent, err := p2p.CreateP2PNetwork(peerIDs[0][:], peerPort, log)
	if err := peer.Listen(); err != nil {
		t.Error("Test Failed: Listen failed ", err)
	}
	peers[0], ok = peer.(*p2p.P2P)
	if err != nil {
		t.Error("Test Failed: CreateP2PNetwork failed ", err)
	}
	if !ok {
		t.Error("Test Failed: ")
	}
	dkgs[0], _ = CreateP2PDkg(peers[0], suite, peerEvent, numOfNodes, log)
	go dkgs[0].EventLoop()
	for i := 1; i < numOfNodes; i++ {
		peerIDs[i] = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12 + byte(i), 13, 14, 15, 16, 17, 18, 19, 21 + byte(i)}
		peerPort := 55550 + i
		peer, peerEvent, err := p2p.CreateP2PNetwork(peerIDs[i][:], peerPort, log)
		if err := peer.Listen(); err != nil {
			t.Error("Test Failed: Listen failed ", err)
		}
		peers[i], ok = peer.(*p2p.P2P)
		if err != nil {
			t.Error("Test Failed: CreateP2PNetwork failed ", err)
		}
		if !ok {
			t.Error("Test Failed: ")
		}
		dkgs[i], _ = CreateP2PDkg(peers[i], suite, peerEvent, numOfNodes, log)
		go dkgs[i].EventLoop()
		peer.Join(peers[0].GetIP())

	}
	for i := 0; i < numOfNodes; i++ {
		dkgs[i].SetNbParticipants(numOfNodes)
		dkgs[i].SetGroupMembers(peerIDs)
		dkgs[i].RunDKG()
	}
	for !dkgs[0].IsCertified() {
		//fmt.Println("dosDkg.IsCetified() ", dkgs[0].IsCertified())
	}
}
