package p2p

import (
	"fmt"
	"sync"
	"testing"

	"github.com/DOSNetwork/core/log"
	"github.com/golang/protobuf/proto"
)

func TestServer(t *testing.T) {

	id1 := []byte("AES256Key-32Characters1234567890")
	fmt.Println("id1  ", id1)
	id2 := []byte("AES256Key-32Characters0987654321")
	id3 := []byte("AES256Key-32Characters0987654322")

	fmt.Println("id2  ", id2)
	log.Init(id1[:])

	p1, _ := CreateP2PNetwork(id1, "9906")
	p1.Listen()
	p2, _ := CreateP2PNetwork(id2, "9905")
	p2.Listen()
	p3, _ := CreateP2PNetwork(id3, "9904")
	p3.Listen()
	events3, _ := p3.SubscribeEvent(1, Ping{})
	go func(p3 P2PInterface) {
		for msg := range events3 {
			r, ok := msg.Msg.Message.(*Ping)
			if !ok {
				t.Errorf("Not ok")
			}
			fmt.Println("Reply to ", msg.Sender)
			p3.Reply(msg.Sender, msg.RequestNonce, proto.Message(&Pong{Count: r.Count + 10}))
		}
	}(p3)
	events, _ := p2.SubscribeEvent(1, Ping{})
	go func(p2 P2PInterface) {
		for msg := range events {
			r, ok := msg.Msg.Message.(*Ping)
			if !ok {
				t.Errorf("Not ok")
			}
			fmt.Println("Reply to ", msg.Sender)
			p2.Reply(msg.Sender, msg.RequestNonce, proto.Message(&Pong{Count: r.Count + 10}))
		}
	}(p2)

	connected, _ := p1.ConnectTo("localhost:9905")
	_ = connected
	connected3, _ := p1.ConnectTo("localhost:9904")
	_ = connected3
	fmt.Println("Start Test!!!!!!!!!!!")

	var count uint64
	checkRoll := make(map[uint64]uint64)
	var wg sync.WaitGroup
	wg.Add(10)
	for count = 0; count < 5; count++ {
		go func(count uint64) {
			defer wg.Done()
			cmd := &Ping{Count: count}
			pb := proto.Message(cmd)
			reply, _ := p1.Request(connected, pb)
			p, ok := reply.Msg.Message.(*Pong)
			if !ok {
				fmt.Println("!!!!!!!!!!!!!not ok")
				return
			}
			fmt.Println("!!!!!!!!!!!!!reply count ", p.Count)
			checkRoll[count] = p.Count
		}(count)
		go func(count uint64) {
			defer wg.Done()
			cmd := &Ping{Count: count}
			pb := proto.Message(cmd)
			reply, _ := p1.Request(connected3, pb)
			p, ok := reply.Msg.Message.(*Pong)
			if !ok {
				fmt.Println("!!!!!!!!!!!!!not ok")
				return
			}
			fmt.Println("!!!!!!!!!!!!!reply count ", p.Count)
			checkRoll[count] = p.Count
		}(count + 5)
	}
	wg.Wait()
	for count = 0; count < 10; count++ {
		if checkRoll[count]-count != 10 {
			t.Errorf("TestRequest ,Expected %d Actual %d", count+10, checkRoll[count])
		}
	}

}
