package p2p

import (
	"testing"
	"fmt"
	"strconv"
	"time"
)

func TestPeerManagerAddPeer(test *testing.T) {
	pm := new(PeerConnManager)
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
		pconn.identity.Id = []byte{byte(i)}
		pm.LoadOrStore(string(pconn.identity.Id), pconn)
		toppconn := pm.FindLessUsedPeerConn()
		if i < MAXPEERCOUNT {
			if toppconn.identity.Address != strconv.Itoa(0) {
				test.Fail()
				fmt.Println("top peer id not equel to expected id")
			}
		}else {
			if pm.PeerConnNum() > MAXPEERCOUNT {
				test.Fail()
				fmt.Println(pm.PeerConnNum())
				fmt.Println("peer size bigger then maxpeercount")
			}
			if toppconn.identity.Address != strconv.Itoa(i-20) {
				test.Fail()
				fmt.Println("top peer id not equel to expected id",i)
			}
		}
		time.Sleep(time.Second)
	}

}

func TestPeerManagerHeap(test *testing.T) {
	t := []int64{100,34,5,87,24,65,8,37,11,14,17,14,20,39,33,87,89,25,74,66,9,1,4,24,54,77,21,29,30,13}
	ans := []int64{100,34,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,5,8,9,11,14,14,17,17}
	pm := new(PeerConnManager)
	curt := time.Now()
	for i:=0; i<len(t); i++ {
		pconn := &PeerConn{
			p2pnet:    nil,
			conn:      nil,
			rxMessage: make(chan P2PMessage, 1),
			waitForHi: make(chan bool, 2),
			rw:        nil,
			lastusedtime: curt.Add(time.Duration(t[i])*time.Second),
		}
		pconn.identity.Id = []byte{byte(i)}
		pconn.identity.Address = strconv.Itoa(i)
		pm.LoadOrStore(string(pconn.identity.Id), pconn)
		toppconn := pm.FindLessUsedPeerConn()
		temp := int64(toppconn.lastusedtime.Sub(curt)/time.Second)
		if int64(temp) != ans[i] {
			test.Fail()
			fmt.Println("failed"+strconv.Itoa(i))
			fmt.Println(toppconn.lastusedtime.Unix())
			fmt.Println(temp)
			fmt.Println(ans[i])
		}

	}

}
