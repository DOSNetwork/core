package p2p

import (
	"fmt"
	"net"
	"sync"
	"testing"

	"github.com/DOSNetwork/core/suites"
	"github.com/golang/protobuf/proto"
)

func listen(listener net.Listener, sID string) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn, sID string) {
			suite := suites.MustFind("bn256")
			secret := suite.Scalar().Pick(suite.RandomStream())
			public := suite.Point().Mul(secret, nil)
			client, err := NewClient(suite, secret, public, []byte(sID), conn, true)
			if err != nil {
				fmt.Println("err", err)

				return
			}
			for msg := range client.receiver {
				p, ok := msg.Msg.Message.(*Ping)
				if !ok {
					fmt.Println("not ok")
				}
				client.Reply(msg.RequestNonce, proto.Message(&Pong{Count: p.Count + 10}))
			}
		}(conn, sID)
	}
}

func TestExchangeID(t *testing.T) {
	var listenerA net.Listener
	var err error
	if listenerA, err = net.Listen("tcp", ":9901"); err != nil {
		return
	}
	go listen(listenerA, "server")

	add := "localhost:9901"
	var conn net.Conn
	var client *Client
	if conn, err = net.Dial("tcp", add); err != nil {
		t.Errorf("Can't Dial to %s ,Error %s", add, err.Error())
	}
	suite := suites.MustFind("bn256")
	secret := suite.Scalar().Pick(suite.RandomStream())
	public := suite.Point().Mul(secret, nil)
	client, err = NewClient(suite, secret, public, []byte("local"), conn, false)
	if err != nil {
		t.Errorf("NewClient,Error %s", err.Error())
	}
	if string(client.localID) != "local" || string(client.remoteID) != "server" {
		t.Errorf("ExchangeID ,Error %s %s", string(client.localID), string(client.remoteID))
	}
}

func TestRequest(t *testing.T) {
	var listenerA net.Listener
	var err error
	if listenerA, err = net.Listen("tcp", ":9902"); err != nil {
		return
	}
	go listen(listenerA, "server")

	add := "localhost:9902"
	var conn net.Conn
	var client *Client
	if conn, err = net.Dial("tcp", add); err != nil {
		t.Errorf("Can't Dial to %s ,Error %s", add, err.Error())
	}
	suite := suites.MustFind("bn256")
	secret := suite.Scalar().Pick(suite.RandomStream())
	public := suite.Point().Mul(secret, nil)

	client, err = NewClient(suite, secret, public, []byte("local"), conn, false)
	if err != nil {
		t.Errorf("NewClient,Error %s", err.Error())
	}
	var count uint64
	checkRoll := make(map[uint64]uint64)
	var wg sync.WaitGroup
	wg.Add(5)
	for count = 0; count < 5; count++ {
		go func(count uint64) {
			defer wg.Done()
			cmd := &Ping{Count: count}
			pb := proto.Message(cmd)
			reply, _ := client.Request(pb)
			p, ok := reply.Msg.Message.(*Pong)
			if !ok {
				return
			}
			fmt.Println("reply count ", p.Count)
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
