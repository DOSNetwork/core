package p2p

import (
	"testing"
	"fmt"
	"strconv"
	"time"
)

func TestPeerManagerAddPeer(test *testing.T) {
    pm := new(PeerManager)
    for i:=0; i<50; i++ {
    	pconn := &PeerConn{
			p2pnet:    nil,
			conn:      nil,
			rxMessage: make(chan P2PMessage, 1),
			waitForHi: make(chan bool, 2),
			rw:        nil,
			lastusedtime: time.Now(),
		}
    	pconn.identity.Address = strconv.Itoa(i)
    	pconn.lastusedtime = time.Now()
    	fmt.Println("add peer"+strconv.Itoa(i))
    	pm.LoadOrStore(strconv.Itoa(i), pconn)
    	toppconn := pm.parray[0]
    	if i < MAXPEERCOUNT {
    		if toppconn.identity.Address != strconv.Itoa(0) {
    			test.Fail()
				fmt.Println("top peer id not equel to expected id")
			}
		}else {
			if pm.parray.Len() > MAXPEERCOUNT {
				test.Fail()
				fmt.Println(pm.parray.Len())
				fmt.Println("peer size bigger then maxpeercount")
			}
			if toppconn.identity.Address != strconv.Itoa(i-20) {
				test.Fail()
				fmt.Println(i)
				fmt.Println("top peer id not equel to expected id")
			}
		}
		time.Sleep(time.Second)
	}

}
