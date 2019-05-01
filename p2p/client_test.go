package p2p

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/suites"
	"github.com/golang/protobuf/proto"
)

func Init() {
	logger = log.New("module", "p2p")
}
func listen(t *testing.T, listener net.Listener, sID string) {
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
				pmsg, ok := msg.(P2PMessage)
				if !ok {
					t.Errorf("casting failed")
				}
				p, ok := pmsg.Msg.Message.(*Ping)
				if !ok {
					t.Errorf("casting failed")
				}

				reply := Request{}

				reply.ctx, reply.cancel = context.WithTimeout(context.Background(), 5*time.Second)
				reply.id = client.remoteID
				reply.rType = 2
				reply.nonce = pmsg.RequestNonce
				errc := make(chan error)
				reply.errc = errc
				reply.msg = proto.Message(&Pong{Count: p.Count + 10})
				client.send(reply)
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
	go listen(t, listenerA, "server")

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
	go listen(t, listenerA, "server")

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
	wg.Add(1)
	for count = 0; count < 1; count++ {
		go func(count uint64) {
			defer wg.Done()

			callReq := Request{}
			callReq.ctx, callReq.cancel = context.WithTimeout(context.Background(), 35*time.Second)
			callReq.rType = 1
			callReq.id = client.remoteID
			callReq.reply = make(chan interface{})
			callReq.errc = make(chan error)
			defer close(callReq.reply)
			defer close(callReq.errc)
			callReq.msg = proto.Message(&Ping{Count: count})
			client.send(callReq)

			select {
			case r, ok := <-callReq.reply:
				if !ok {
					return
				}
				msg, ok := r.(P2PMessage)
				if ok {

				}
				p, ok := msg.Msg.Message.(*Pong)
				if !ok {
					t.Errorf("casting failed")
				}
				checkRoll[count] = p.Count
				return
			case e, ok := <-callReq.errc:
				if ok {
					err = e
					fmt.Println("reply err ", err)
					return
				}
			case <-callReq.ctx.Done():
				err = callReq.ctx.Err()
				t.Errorf("TestRequest, ctx Error %s", callReq.ctx.Err())

				return
			}
		}(count)
	}
	wg.Wait()
	for count = 0; count < 5; count++ {
		if checkRoll[count]-count != 10 {
			//t.Errorf("TestRequest ,Expected %d Actual %d", count+10, checkRoll[count])
		}
	}
}
