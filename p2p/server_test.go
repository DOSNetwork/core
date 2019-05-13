package p2p

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/DOSNetwork/core/log"
	"github.com/golang/protobuf/proto"
)

func TestServer(t *testing.T) {
	listener := []byte("9")
	os.Setenv("PUBLICIP", "0.0.0.0")
	log.Init(listener[:])

	pListener, _ := CreateP2PNetwork(listener, "9905", NONE)
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
	PrintMemUsage()

	var wgForPeer sync.WaitGroup
	wgForPeer.Add(3)
	for c := 9902; c < 9905; c++ {
		go func(c int) {
			defer wgForPeer.Done()
			id := []byte(strconv.Itoa(c))
			p, _ := CreateP2PNetwork(id, strconv.Itoa(c), NONE)
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
		}(c)
	}
	wgForPeer.Wait()
	PrintMemUsage()

}

func TestDisconn(t *testing.T) {
	listener := []byte("9")
	os.Setenv("PUBLICIP", "0.0.0.0")
	log.Init(listener[:])

	pListener, _ := CreateP2PNetwork(listener, "9905", NONE)
	pListener.Listen()
	events, _ := pListener.SubscribeEvent(1, Ping{})
	go func(pListener P2PInterface) {
		for msg := range events {
			r, ok := msg.Msg.Message.(*Ping)
			if !ok {
				t.Errorf("Not ok")
			}
			pListener.Reply(msg.Sender, msg.RequestNonce, proto.Message(&Pong{Count: r.Count + 10}))
			pListener.DisConnectTo(msg.Sender)
		}
	}(pListener)

	p, _ := CreateP2PNetwork([]byte("0"), "9904", NONE)
	p.Listen()
	p.SetPort("9905")
	connected, err := p.ConnectTo("0.0.0.0", nil)
	if err != nil {
		t.Errorf("TestRequest ,Error %s", err)
	}
	var count uint64
	for count = 0; count < 10; count++ {
		cmd := &Ping{Count: count}
		pb := proto.Message(cmd)
		reply, err := p.Request(connected, pb)
		if err != nil {
			fmt.Println("Request,Err ", err)
			continue
		}
		p, ok := reply.Msg.Message.(*Pong)
		if !ok {
			fmt.Println("test !ok")
			t.Errorf("TestRequest ,!ok")
		}
		if p.Count-count != 10 {
			t.Errorf("TestRequest ,Expected %d Actual %d", count+10, p.Count)
		}
	}
	fmt.Println("test")
	//p.DisConnectTo(connected)

}
