package p2p

import (
	"testing"
	"time"
	"strconv"
	"fmt"
)

func TestPeerManagerAddPeer(test *testing.T) {
    pm := new(PeerManager)
    for i:=0; i<50; i++ {
    	pconn := new(PeerConn)
    	pconn.identity.Address = strconv.Itoa(i)
    	pconn.lastusedtime = time.Now()
    	pm.LoadOrStore(strconv.Itoa(i), pconn)
    	toppconn := pm.parray[0]
    	if i < MAXPEERCOUNT {
    		if toppconn.identity.Address != strconv.Itoa(0) {
    			test.Fail()
    			fmt.Println("Add Peer Test Failed")
			}
		}else {
			if pm.parray.Len() > MAXPEERCOUNT {
				test.Fail()
				fmt.Println("Add Peer Test Failed")
			}
			if toppconn.identity.Address != strconv.Itoa(i-20) {
				test.Fail()
				fmt.Println("Add Peer Test Failed")
			}
		}
		time.Sleep(1000)
	}

}
