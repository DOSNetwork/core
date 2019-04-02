package p2p

import (
	"fmt"
	"sync"
	"testing"

	"github.com/DOSNetwork/core/log"
	"github.com/golang/protobuf/proto"
)

func TestServer(t *testing.T) {
	id1 := []byte("1")
	id2 := []byte("9")

	log.Init(id1[:])

	p1, _ := CreateP2PNetwork(id1, "9906", NONE)
	p1.Listen()
	p2, _ := CreateP2PNetwork(id2, "9905", NONE)
	p2.Listen()
	events, _ := p2.SubscribeEvent(1, Ping{})
	go func(p2 P2PInterface) {
		for msg := range events {
			r, ok := msg.Msg.Message.(*Ping)
			if !ok {
				t.Errorf("Not ok")
			}
			p2.Reply(msg.Sender, msg.RequestNonce, proto.Message(&Pong{Count: r.Count + 10}))
		}
	}(p2)
	p1.SetPort("9905")
	connected, _ := p1.ConnectTo("0.0.0.0", nil)

	var count uint64
	checkRoll := make(map[uint64]uint64)
	var wg sync.WaitGroup
	wg.Add(5)
	for count = 0; count < 5; count++ {
		go func(count uint64) {
			defer wg.Done()
			cmd := &Ping{Count: count}
			pb := proto.Message(cmd)
			reply, _ := p1.Request(connected, pb)
			p, ok := reply.Msg.Message.(*Pong)
			if !ok {
				return
			}
			checkRoll[count] = p.Count
		}(count)
	}
	wg.Wait()
	for count = 0; count < 5; count++ {
		if checkRoll[count]-count != 10 {
			t.Errorf("TestRequest ,Expected %d Actual %d", count+10, checkRoll[count])
		}
	}

}
