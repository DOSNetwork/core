package p2p

import (
	"testing"

	"github.com/sirupsen/logrus"
)

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

	bContactsBefore := boot.GetPeers()
	pContactsBefore := peer.GetPeers()
	if len(pContactsBefore) != 0 || len(pContactsBefore) != 0 {
		t.Errorf("Test Failed: Routing Table should be empty boot(%d) peer(%d)", len(bContactsBefore), len(pContactsBefore))
	}

	peer.BootStrap(boot.GetIPAddress())

	bContactsAfter := boot.GetPeers()
	pContactsAfter := peer.GetPeers()
	if len(bContactsAfter) != 1 || bContactsAfter[string(peerid)] != peer.GetIPAddress() {
		t.Errorf("boot contacts(%d) : peer contact %s peer ip %s", len(bContactsAfter), bContactsAfter[string(peerid)], peer.GetIPAddress())
	}
	if len(pContactsAfter) != 2 ||
		pContactsAfter[string(peerid)] != peer.GetIPAddress() ||
		pContactsAfter[string(bootid)] != boot.GetIPAddress() {
		t.Errorf("peer contacts(%d) : peer contact %s  peer ip %s , boot contact %s  boot ip %s ",
			len(pContactsAfter), pContactsAfter[string(peerid)], peer.GetIPAddress(),
			pContactsAfter[string(bootid)], boot.GetIPAddress())
	}
}
