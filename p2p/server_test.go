package p2p

import (
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/golang/protobuf/proto"
)

func TestServer(t *testing.T) {
	listener := []byte("9")
	os.Setenv("PUBLICIP", "0.0.0.0")
	log.Init(listener[:])

	pListener, _ := CreateP2PNetwork(listener, "9905", nodiscover)
	pListener.Listen()
	events, _ := pListener.SubscribeEvent(1, Ping{})
	go func(pListener P2PInterface) {
		for msg := range events {
			r, ok := msg.Msg.Message.(*Ping)
			if !ok {
				t.Errorf("Not ok")
			}
			pListener.Reply(msg.Sender, msg.RequestNonce, proto.Message(&Pong{Count: r.Count + 10}))
		}
	}(pListener)

	var wgForPeer sync.WaitGroup
	wgForPeer.Add(3)
	for c := 9902; c < 9905; c++ {
		go func(c int) {
			defer wgForPeer.Done()
			id := []byte(strconv.Itoa(c))
			p, _ := CreateP2PNetwork(id, strconv.Itoa(c), nodiscover)
			p.Listen()
			p.SetPort("9905")
			connected, err := p.ConnectTo("0.0.0.0", nil)
			if err != nil {
				t.Errorf("TestRequest ,Error %s", err)
			}

			var count uint64
			var wgForMsg sync.WaitGroup
			wgForMsg.Add(3)
			for count = 0; count < 3; count++ {
				go func(count uint64) {
					defer wgForMsg.Done()
					cmd := &Ping{Count: count}
					pb := proto.Message(cmd)
					reply, _ := p.Request(connected, pb)
					p, ok := reply.Msg.Message.(*Pong)
					if !ok {
						return
					}
					if p.Count-count != 10 {
						t.Errorf("TestRequest ,Expected %d Actual %d", count+10, p.Count)
					}
				}(count)
			}
			wgForMsg.Wait()
			p.DisConnectTo(connected)
			retryLimit := 5
			for {
				rNum, cNum := p.numOfClient()
				if rNum == 0 && cNum == 0 {
					break
				}
				retryLimit--
				if retryLimit == 0 {
					t.Errorf("TestServer ,Expected %d Actual %d", 1, cNum)
				}
				time.Sleep(1 * time.Second)
			}
		}(c)
	}
	wgForPeer.Wait()
	retryLimit := 5
	for {
		prNum, pcNum := pListener.numOfClient()
		if prNum == 0 && pcNum == 0 {
			break
		}
		retryLimit--
		if retryLimit == 0 {
			t.Errorf("TestServer ,Expected %d Actual %d", 0, prNum)
		}
		time.Sleep(1 * time.Second)
	}
}
