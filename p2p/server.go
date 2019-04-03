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

	//	"reflect"
	//	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p/network"
	"github.com/DOSNetwork/core/suites"
)

const DefaultModelName = "Undefined"

type Server struct {
	id     []byte
	suite  suites.Suite
	secKey kyber.Scalar
	pubKey kyber.Point

	addr     net.IP
	port     string
	listener net.Listener

	//Client lookup
	network   network.Network
	calling   chan Request
	replying  chan Request
	incoming  chan *Client
	registerC chan *Client

	//Event
	messages  chan P2PMessage
	subscribe chan Subscription
	unscribe  chan Subscription

	ctx    context.Context
	cancel context.CancelFunc
}

type Request struct {
	rType  int
	ctx    context.Context
	cancel context.CancelFunc
	addr   net.IP
	id     []byte
	//Client signs and packs msg into Package
	msg proto.Message
	p   *Package
	//
	nonce uint64
	reply chan interface{}
	errc  chan error
}

type Subscription struct {
	eventType string
	message   chan P2PMessage
}

func (n *Server) Join(bootstrapIp []string) (err error) {
	n.network.Join(bootstrapIp)
	return
}

func (n *Server) Members() int {
	return n.network.NumPeers()
}

func (n *Server) ConnectToAll() (memNum, connNum int) {
	addrs := n.network.GetOtherMembersIP()
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

func (n *Server) SetID(id []byte) {
	n.id = id
}

func (n *Server) SetPort(port string) {
	n.port = port
}

func (n *Server) GetID() []byte {
	return n.id
}

func (n *Server) GetIP() net.IP {
	return n.addr
}

func (n *Server) Listen() (err error) {
	var ip net.IP
	n.receiveHandler()
	n.callHandler()
	go n.messageDispatch(context.Background())

	if ip, err = GetLocalIP(); err != nil {
		//fmt.Println("GetLocalIP err", err)

		logger.Error(err)
		return
	}

	p := fmt.Sprintf(":%s", n.port)
	if n.listener, err = net.Listen("tcp", p); err != nil {
		//fmt.Println("listener err", err)

		logger.Error(err)
		return
	}

	//NAT discover

	//isPrivateIp, err := nat.IsPrivateIp()
	//if err != nil {
	//	return err
	//}
	//
	//if isPrivateIp {
	//	externalPort := nat.RandomPort()
	//	nat, err := nat.SetMapping("tcp", externalPort, listener.Addr().(*net.TCPAddr).Port, "DosNode")
	//	if err != nil {
	//		return err
	//	}
	//
	//	externalIp, err := nat.GetExternalAddress()
	//	if err != nil {
	//		return err
	//	}
	//
	//	n.port = externalPort
	//	ip = externalIp.String() + ":" + strconv.Itoa(n.port)
	//} else {
	//	n.port = listener.Addr().(*net.TCPAddr).Port
	//	ip = ip + ":" + strconv.Itoa(n.port)
	//}

	n.addr = ip
	fmt.Println("Listen to ", ip, " ", n.port)
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
				client, err := NewClient(n.suite, n.secKey, n.pubKey, n.id, conn, true)
				if err != nil {
					//fmt.Println("listen to client err", err)
					logger.Error(err)
					return
				}
				go func(client *Client, messages chan P2PMessage) {
					//defer //fmt.Println("connect to client over")
					for {
						select {
						case pa, ok := <-client.receiver:
							if !ok {
								return
							}
							if m, ok := pa.(P2PMessage); ok {
								messages <- m
							}
						case err, ok := <-client.errc:
							if !ok {
								return
							}
							//fmt.Println(client.localID, " err ", err)
							if err.Error() == "EOF" {
								client.Close()
								return
							}
						case <-client.ctx.Done():
							//fmt.Println(client.localID, " Over")
							return
						}
					}
				}(client, n.messages)
				n.incoming <- client
			}(conn, start)
		}
	}()

	return nil
}
func (n *Server) receiveHandler() {
	n.incoming = make(chan *Client, 100)
	n.replying = make(chan Request)
	clients := make(map[string]*Client)

	go func() {
		for {
			select {
			case c, ok := <-n.incoming:
				if !ok {
					return
				}
				//fmt.Println(time.Now(), "receiveHandler incoming ", c.remoteID)

				clients[string(c.remoteID)] = c

			case req, ok := <-n.replying:
				if !ok || req.id == nil {
					//fmt.Println(time.Now(), "receiveHandler close")

					return
				}

				client := clients[string(req.id)]
				if client == nil {
					//fmt.Println(time.Now(), "receiveHandler client is nil")

					select {
					case n.replying <- req:
						//fmt.Println(time.Now(), "receiveHandler retry late")

						if req.ctx == nil || req.reply == nil {
							//fmt.Println(time.Now(), "receiveHandler  req is nil ")
						}
					case <-req.ctx.Done():
					}
				} else {
					//fmt.Println(time.Now(), "receiveHandler found client")
					if req.ctx == nil {
						//fmt.Println(time.Now(), "receiveHandler  req is nil ")
					}
					client.send(req)
				}
			}
		}
	}()
}
func (n *Server) callHandler() {
	n.calling = make(chan Request, 100)
	hangup := make(chan string)
	addrToid := make(map[string][]byte)
	idTostatus := make(map[string][]byte)

	clients := make(map[string]*Client)
	n.registerC = make(chan *Client)
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
							go func(req Request) {
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
						//TODO:ASK Client to send request here
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
						req.addr = n.network.Lookups(req.id)

						if req.addr == nil {
							//Retry later
							go func(req Request) {
								select {
								case n.calling <- req:
								case <-req.ctx.Done():
								}
							}(req)
							continue
						}
					}
					idTostatus[string(req.id)] = []byte{'p', 'e', 'n', 'd', 'i', 'n', 'g'}
					addrToid[req.addr.String()+":"+n.port] = []byte{'p', 'e', 'n', 'd', 'i', 'n', 'g'}

					go func(req Request, start time.Time) {
						var conn net.Conn
						var client *Client
						if conn, err = net.Dial("tcp", req.addr.String()+":"+n.port); err != nil {
							logger.Error(err)
							select {
							case req.errc <- err:
							case <-req.ctx.Done():
							}
							//Retry later
							go func(req Request) {
								select {
								case n.calling <- req:
								case <-req.ctx.Done():
								}
							}(req)
							return
						}

						if client, err = NewClient(n.suite, n.secKey, n.pubKey, n.id, conn, false); err != nil {
							logger.Error(err)
							select {
							case req.errc <- err:
							case <-req.ctx.Done():
							}
							conn.Close()
							//Retry later
							go func(req Request) {
								select {
								case n.calling <- req:
								case <-req.ctx.Done():
								}
							}(req)
							return
						}

						go func() {
							for {
								select {
								case pa, ok := <-client.receiver:
									if !ok {
										return
									}
									if m, ok := pa.(P2PMessage); ok {
										n.messages <- m
									}
								case err := <-client.errc:
									if err.Error() == "EOF" {
										client.Close()
										return
									}
								case <-client.ctx.Done():
									return
								}
							}
						}()
						//TODO:ASK Client to send request her
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
						select {
						case n.registerC <- client:
						case <-req.ctx.Done():
						}
					}(req, start)
				}
			case client, ok := <-n.registerC:
				if !ok {
					return
				}

				clients[string(client.remoteID)] = client
				addrToid[client.conn.RemoteAddr().String()] = client.remoteID

				delete(idTostatus, string(client.remoteID))
				f := map[string]interface{}{
					"localID":    client.localID,
					"remoteID":   client.remoteID,
					"RemoteAddr": client.conn.RemoteAddr().String(),
					"Time":       time.Now()}
				logger.Event("registerClient", f)
			case _, _ = <-hangup:
			}
		}
	}()
	return
}

func (n *Server) Leave() {
	err := n.listener.Close()
	if err != nil {
	}
	n.cancel()

	if n.network != nil {
		n.network.Leave()
	}

	return
}

/*
This is a block call
*/
func (n *Server) ConnectTo(addr string, id []byte) ([]byte, error) {
	var err error
	callReq := Request{}
	callReq.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
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
		client, ok := r.(*Client)
		if ok {
			id = client.remoteID
		}
		return id, nil

	case err = <-callReq.errc:
		return nil, err
	case <-callReq.ctx.Done():
		return nil, callReq.ctx.Err()
	}

	return nil, err
}

func (n *Server) Request(id []byte, m proto.Message) (msg P2PMessage, err error) {
	defer logger.TimeTrack(time.Now(), "Request", nil)
	callReq := Request{}
	callReq.ctx, callReq.cancel = context.WithTimeout(context.Background(), 5*time.Second)
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

func (n *Server) Reply(id []byte, nonce uint64, response proto.Message) (err error) {
	callReq := Request{}

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
	/*
		select {
		case e, ok := <-callReq.errc:
			if ok {
				err = e
				//fmt.Println(time.Now(), "Reply err ", e)
				return
			} else {
				//fmt.Println(time.Now(), "Server reply  done")

			}
		case <-callReq.ctx.Done():
			err = callReq.ctx.Err()
			//if strings.Contains(callReq.ctx.Err(), "deadline exceeded") {
			//fmt.Println(time.Now(), "Reply ctx err ", callReq.ctx.Err())
			//}
			return

		}*/
	return
}
func (n *Server) messageDispatch(ctx context.Context) {
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

func (n *Server) SubscribeEvent(chanBuffer int, messages ...interface{}) (outch chan P2PMessage, err error) {
	if chanBuffer > 0 {
		outch = make(chan P2PMessage, chanBuffer)
	} else {
		outch = make(chan P2PMessage)
	}
	for _, m := range messages {
		n.subscribe <- Subscription{reflect.TypeOf(m).String(), outch}
	}
	return
}

func (n *Server) UnSubscribeEvent(messages ...interface{}) {
	for _, m := range messages {
		n.unscribe <- Subscription{reflect.TypeOf(m).String(), nil}

	}
	return
}
