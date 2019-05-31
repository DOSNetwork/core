package p2p

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"reflect"

	//	"strings"
	"time"

	"github.com/dedis/kyber"

	"github.com/golang/protobuf/proto"

	"github.com/DOSNetwork/core/p2p/discover"
	"github.com/DOSNetwork/core/suites"
)

type server struct {
	id     []byte
	suite  suites.Suite
	secKey kyber.Scalar
	pubKey kyber.Point

	addr     net.IP
	port     string
	listener net.Listener

	//client lookup
	members  discover.Membership
	calling  chan request
	replying chan request

	addIncomingC    chan *client
	removeIncomingC chan []byte
	addCallingC     chan *client
	removeCallingC  chan []byte
	incomingNum     int
	callingNum      int

	//Event
	messages  chan P2PMessage
	subscribe chan subscription
	unscribe  chan subscription

	ctx    context.Context
	cancel context.CancelFunc
}

type request struct {
	rType  int
	ctx    context.Context
	cancel context.CancelFunc
	addr   net.IP
	id     []byte
	//client signs and packs msg into Package
	msg proto.Message
	p   *Package
	//
	nonce uint64
	reply chan interface{}
	errc  chan error
}

type subscription struct {
	eventType string
	message   chan P2PMessage
}

func (n *server) Join(bootstrapIP []string) (num int, err error) {
	return n.members.Join(bootstrapIP)
}

func (n *server) Members() int {
	return n.members.NumOfPeers()
}

func (n *server) ConnectToAll() (memNum, connNum int) {
	addrs := n.members.PeersIP()
	memNum = len(addrs)
	for _, addr := range addrs {
		if _, err := n.ConnectTo(addr.String(), nil); err != nil {
			fmt.Println("ConnectTo ", addr, " fail", err)
		} else {
			connNum++
		}
	}

	return
}
func (n *server) numOfClient() (iNum, cNum int) {
	return n.incomingNum, n.callingNum
}

func (n *server) SetID(id []byte) {
	n.id = id
}

func (n *server) SetPort(port string) {
	n.port = port
}

func (n *server) GetID() []byte {
	return n.id
}

func (n *server) GetIP() net.IP {
	return n.addr
}

func (n *server) Listen() (err error) {
	n.receiveHandler()
	n.callHandler()
	go n.messageDispatch(context.Background())

	p := fmt.Sprintf(":%s", n.port)
	if n.listener, err = net.Listen("tcp", p); err != nil {
		logger.Error(err)
		return
	}
	fmt.Println("Listen to ", n.addr, " ", n.port)

	go func() {
		for {
			conn, err := n.listener.Accept()
			if err != nil {
				//fmt.Println("Accept err", err)
				logger.Error(err)
				return
			}
			start := time.Now()
			//fmt.Println("new conn ")
			go func(conn net.Conn, start time.Time) {
				fmt.Println(string(n.id), "Got a conn")
				c, err := newClient(n.suite, n.secKey, n.pubKey, n.id, conn, true)
				if err != nil {
					//fmt.Println("listen to client err", err)
					logger.Error(err)
					return
				}
				go func(c *client, messages chan P2PMessage) {
					defer fmt.Println(string(n.id), "Close a Incoming Client")
					defer func() { n.removeIncomingC <- c.remoteID }()
					for {
						select {
						case pa, ok := <-c.receiver:
							if !ok {
								return
							}
							if m, ok := pa.(P2PMessage); ok {
								messages <- m
							}
						case err, ok := <-c.errc:
							if !ok {
								return
							}
							fmt.Println(string(n.id), "err.Error() ", err.Error())
							if err.Error() == "EOF" || err.Error() == "use of closed network connection" {
								c.close()
								return
							}
						case <-c.ctx.Done():
							return
						}
					}
				}(c, n.messages)
				n.addIncomingC <- c
			}(conn, start)
		}
	}()

	return nil
}
func (n *server) receiveHandler() {
	n.addIncomingC = make(chan *client, 21)
	n.removeIncomingC = make(chan []byte)
	n.replying = make(chan request)
	clients := make(map[string]*client)

	go func() {
		for {
			select {
			case c, ok := <-n.addIncomingC:
				if !ok {
					return
				}
				clients[string(c.remoteID)] = c
				n.incomingNum = len(clients)
			case id, ok := <-n.removeIncomingC:
				if !ok {
					return
				}
				c := clients[string(id)]
				if c != nil {
					delete(clients, string(id))
				}
				n.incomingNum = len(clients)

			case req, ok := <-n.replying:
				if !ok || req.id == nil {
					return
				}

				client := clients[string(req.id)]
				if client == nil {
					select {
					case n.replying <- req:
						if req.ctx == nil || req.reply == nil {
						}
					case <-req.ctx.Done():
					}
				} else {
					if req.ctx == nil {
					}
					client.send(req)
				}
			}
		}
	}()
}

func (n *server) callHandler() {
	n.calling = make(chan request, 21)
	hangup := make(chan string)
	addrToid := make(map[string][]byte)
	idTostatus := make(map[string][]byte)

	clients := make(map[string]*client)
	n.addCallingC = make(chan *client)
	n.removeCallingC = make(chan []byte)

	go func() {
		for {
			select {
			case req, ok := <-n.calling:
				if !ok {
					return
				}
				start := time.Now()

				if req.id == nil {
					if req.addr == nil || req.ctx == nil {
						continue
					}

					id := addrToid[req.addr.String()+":"+n.port]
					if id != nil {
						if !bytes.Equal(id, []byte{'p', 'e', 'n', 'd', 'i', 'n', 'g'}) {
							if client := clients[string(id)]; client != nil {
								if req.rType == 1 {
									client.send(req)
								} else if req.rType == 0 {
									select {
									case req.reply <- client:
									case <-req.ctx.Done():
									}
									close(req.reply)
									close(req.errc)
								}
							}
						} else {
							go func(req request) {
								time.Sleep(1 * time.Second)
								select {
								case n.calling <- req:
								case <-req.ctx.Done():
								}
							}(req)
						}
						continue
					}
				} else {
					if client := clients[string(req.id)]; client != nil {
						//TODO:ASK client to send request here
						if req.rType == 1 {
							client.send(req)
						} else if req.rType == 0 {
							select {
							case req.reply <- client:
							case <-req.ctx.Done():
							}
							close(req.reply)
							close(req.errc)
						}
						continue
					}
				}

				var err error
				select {
				case <-req.ctx.Done():
					continue
				default:

					if req.addr == nil && req.id != nil {
						if bytes.Equal(idTostatus[string(req.id)], []byte{'p', 'e', 'n', 'd', 'i', 'n', 'g'}) {
							continue
						}
						// Find Peer from routing map
						if n.members != nil {
							req.addr = n.members.Lookup(req.id)
						}

						if req.addr == nil {
							//Retry later
							go func(req request) {
								select {
								case n.calling <- req:
								case <-req.ctx.Done():
								}
							}(req)
							continue
						}
					}
					if req.id != nil {
						idTostatus[string(req.id)] = []byte{'p', 'e', 'n', 'd', 'i', 'n', 'g'}
					}
					if req.addr != nil {
						addrToid[req.addr.String()+":"+n.port] = []byte{'p', 'e', 'n', 'd', 'i', 'n', 'g'}
					}
					go func(req request, start time.Time) {
						var conn net.Conn
						var c *client
						if conn, err = net.Dial("tcp", req.addr.String()+":"+n.port); err != nil {
							logger.Error(err)
							select {
							case req.errc <- err:
							case <-req.ctx.Done():
							}
							//Retry later
							go func(req request) {
								select {
								case n.calling <- req:
								case <-req.ctx.Done():
								}
							}(req)
							return
						}

						if c, err = newClient(n.suite, n.secKey, n.pubKey, n.id, conn, false); err != nil {
							logger.Error(err)
							select {
							case req.errc <- err:
							case <-req.ctx.Done():
							}
							conn.Close()
							//Retry later
							go func(req request) {
								select {
								case n.calling <- req:
								case <-req.ctx.Done():
								}
							}(req)
							return
						}

						go func() {
							defer fmt.Println(string(n.id), "close a calling client")
							defer func() { n.removeCallingC <- c.remoteID }()
							for {
								select {
								case pa, ok := <-c.receiver:
									if !ok {
										return
									}
									if m, ok := pa.(P2PMessage); ok {
										n.messages <- m
									}
								case err, ok := <-c.errc:
									if !ok {
										return
									}
									if err.Error() == "EOF" || err.Error() == "use of closed network connection" {
										c.close()
										return
									}
								case <-c.ctx.Done():
									return
								}
							}
						}()
						//TODO:ASK client to send request her
						if req.rType == 1 {
							c.send(req)
						} else if req.rType == 0 {
							select {
							case req.reply <- c:
							case <-req.ctx.Done():
							}
							close(req.reply)
							close(req.errc)
						}
						select {
						case n.addCallingC <- c:
						case <-req.ctx.Done():
						}
					}(req, start)
				}
			case c, ok := <-n.addCallingC:
				if !ok {
					return
				}

				clients[string(c.remoteID)] = c
				n.callingNum = len(clients)
				addrToid[c.conn.RemoteAddr().String()] = c.remoteID
				delete(idTostatus, string(c.remoteID))
			case id, ok := <-n.removeCallingC:
				if !ok {
					return
				}
				c := clients[string(id)]
				if c != nil {
					delete(addrToid, c.conn.RemoteAddr().String())
					delete(clients, string(id))
					c.close()
				}
				n.callingNum = len(clients)
			case _, _ = <-hangup:
			}
		}
	}()
	return
}

func (n *server) Leave() {
	err := n.listener.Close()
	if err != nil {
	}
	n.cancel()

	if n.members != nil {
		n.members.Leave()
	}

	return
}

func (n *server) DisConnectTo(id []byte) (err error) {
	n.removeCallingC <- id
	return
}

/*
This is a block call
*/
func (n *server) ConnectTo(addr string, id []byte) ([]byte, error) {
	var err error
	callReq := request{}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	callReq.ctx = ctx
	callReq.rType = 0
	callReq.id = id
	if addr != "" {
		callReq.addr = net.ParseIP(addr)
	}
	callReq.reply = make(chan interface{})
	callReq.errc = make(chan error)

	select {
	case n.calling <- callReq:
	case <-callReq.ctx.Done():
		return nil, callReq.ctx.Err()
	}

	select {
	case r := <-callReq.reply:
		client, ok := r.(*client)
		if ok {
			id = client.remoteID
		}
		return id, nil

	case err = <-callReq.errc:
		return nil, err
	case <-callReq.ctx.Done():
		return nil, callReq.ctx.Err()
	}
}

// Request sends a proto message to the specific node
func (n *server) Request(id []byte, m proto.Message) (msg P2PMessage, err error) {
	//defer logger.TimeTrack(time.Now(), "Request", nil)
	callReq := request{}
	callReq.ctx, callReq.cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer callReq.cancel()
	callReq.rType = 1
	callReq.id = id
	callReq.reply = make(chan interface{})
	callReq.errc = make(chan error)
	callReq.msg = m
	select {
	case n.calling <- callReq:
	case <-callReq.ctx.Done():
		return
	}

	select {
	case r, ok := <-callReq.reply:
		if !ok {
			return
		}
		msg, ok = r.(P2PMessage)
		if !ok {
			err = errors.New("Reply cast error")
		}
		return
	case e, ok := <-callReq.errc:
		if ok {
			err = e
			return
		}
	case <-callReq.ctx.Done():
		err = callReq.ctx.Err()
		go func() {
			select {
			case _ = <-callReq.reply:
			case <-time.After(5 * time.Second):
			}
		}()
		return
	}
	return
}

// Reply sends a reply to the specific node
func (n *server) Reply(id []byte, nonce uint64, response proto.Message) (err error) {
	callReq := request{}

	callReq.ctx, callReq.cancel = context.WithTimeout(context.Background(), 5*time.Second)
	callReq.id = id
	callReq.rType = 2
	callReq.nonce = nonce
	errc := make(chan error)
	callReq.errc = errc
	callReq.msg = response
	if callReq.ctx == nil {
	}
	select {
	case n.replying <- callReq:

	case <-callReq.ctx.Done():
		return
	}
	return
}

func (n *server) messageDispatch(ctx context.Context) {
	subscriptions := make(map[string]chan P2PMessage)
	go func() {
		for {
			select {
			case msg, ok := <-n.messages:
				if !ok {
					return
				}
				if msg.Msg.Message == nil {

					continue
				}
				messagetype := reflect.TypeOf(msg.Msg.Message).String()
				if len(messagetype) > 0 && messagetype[0] == '*' {
					messagetype = messagetype[1:]
				}
				out := subscriptions[messagetype]
				if out != nil {

					select {
					case out <- msg:
					}
				} else {
				}
			case sub, ok := <-n.subscribe:
				if !ok {
					return
				}
				subscriptions[sub.eventType] = sub.message
			case sub, ok := <-n.unscribe:
				if !ok {
					return
				}
				delete(subscriptions, sub.eventType)
			case <-ctx.Done():
			}
		}
	}()
}

// SubscribeEvent is a message subscription operation binding the P2PMessage
func (n *server) SubscribeEvent(chanBuffer int, messages ...interface{}) (outch chan P2PMessage, err error) {
	if chanBuffer > 0 {
		outch = make(chan P2PMessage, chanBuffer)
	} else {
		outch = make(chan P2PMessage)
	}
	for _, m := range messages {
		n.subscribe <- subscription{reflect.TypeOf(m).String(), outch}
	}
	return
}

// UnSubscribeEvent is a un-subscription operation
func (n *server) UnSubscribeEvent(messages ...interface{}) {
	for _, m := range messages {
		n.unscribe <- subscription{reflect.TypeOf(m).String(), nil}

	}
	return
}
