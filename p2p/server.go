package p2p

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"reflect"
	//	"strings"
	"time"

	"github.com/dedis/kyber"

	"github.com/golang/protobuf/proto"

	//	"reflect"
	//	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/suites"

	"github.com/hashicorp/serf/serf"
)

const DefaultModelName = "Undefined"

type Server struct {
	id     []byte
	suite  suites.Suite
	secKey kyber.Scalar
	pubKey kyber.Point

	address  string
	port     string
	listener net.Listener

	//Client lookup
	cluster   *serf.Serf
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
	addr   string
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

func SetupCluster(advertiseAddr string, id []byte) (*serf.Serf, error) {
	conf := serf.DefaultConfig()
	conf.Init()
	conf.LogOutput = ioutil.Discard
	conf.MemberlistConfig.AdvertiseAddr = advertiseAddr
	conf.NodeName = string(id)
	cluster, err := serf.Create(conf)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (n *Server) Join(bootstrapIp string) (err error) {
	_, err = n.cluster.Join([]string{bootstrapIp}, true)
	if err != nil {
		//fmt.Println(err)
	}
	return
}

func (n *Server) Members() int {
	return len(n.cluster.Members())
}

func (n *Server) SetID(id []byte) {
	n.id = id
}

func (n *Server) GetID() []byte {
	return n.id
}

func (n *Server) GetIP() string {
	return n.address
}

func (n *Server) Listen() (err error) {
	var ip string
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

	n.address = ip
	//fmt.Println("Listen to ", ip, " ", n.port)
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
				logger.TimeTrack(start, "ExchangeReceive", nil)

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
					if req.addr == "" || req.ctx == nil || req.reply == nil {
						continue
					}
					//fmt.Println(time.Now(), "callHandler req.id == nil ", len(clients), len(addrToid), addrToid[req.addr], req.addr)

					id := addrToid[req.addr]
					if id != nil {
						if !bytes.Equal(id, []byte{'p', 'e', 'n', 'd', 'i', 'n', 'g'}) {
							if client := clients[string(id)]; client != nil {
								client.send(req)
							}
						} else {
							//fmt.Println(time.Now(), "callHandler req.addr pending", req.addr, id)
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
						client.send(req)
						continue
					}
					//fmt.Println(time.Now(), "callHandler req.id not nil ", len(clients), len(addrToid), req.id, addrToid[req.addr], req.addr)
				}

				var err error
				select {
				case <-req.ctx.Done():
					continue
				default:

					if req.addr == "" && req.id != nil {
						if bytes.Equal(idTostatus[string(req.id)], []byte{'p', 'e', 'n', 'd', 'i', 'n', 'g'}) {
							continue
						}
						// Find Peer from routing map
						//fmt.Println(time.Now(), "callHandler Find Peer from routing map ", req.id, req.addr)

						members := n.cluster.Members()
						for i := 0; i < len(members); i++ {
							if members[i].Name == string(req.id) {
								req.addr = members[i].Addr.String()
								//fmt.Println(time.Now(), "callHandler Find client addr from routing map", req.id, req.addr)
								break
							}
						}
						if req.addr == "" {
							//Retry later
							//fmt.Println(time.Now(), "callHandler Retry later", req.id)
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
					addrToid[req.addr] = []byte{'p', 'e', 'n', 'd', 'i', 'n', 'g'}
					//fmt.Println(time.Now(), "callHandler Create Client ", len(clients), len(addrToid), req.id, req.addr)
					go func(req Request, start time.Time) {
						var conn net.Conn
						var client *Client
						if conn, err = net.Dial("tcp", req.addr+":"+n.port); err != nil {
							//fmt.Println(time.Now(), "callHandler Dial err", err)
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
						//fmt.Println(time.Now(), "callHandler Dial success", req.addr)

						if client, err = NewClient(n.suite, n.secKey, n.pubKey, n.id, conn, false); err != nil {
							//fmt.Println("connect to client err", err)
							logger.Error(err)
							select {
							case req.errc <- err:
							case <-req.ctx.Done():
							}
							conn.Close()
							//Retry later
							//fmt.Println(time.Now(), "callHandler Retry later")
							go func(req Request) {
								select {
								case n.calling <- req:
								case <-req.ctx.Done():
									//fmt.Println(time.Now(), "callHandler NewClient ctx ", req.ctx.Err())
								}
							}(req)
							return
						}
						//fmt.Println(time.Now(), "callHandler NewClient", client.localID, " - ", client.remoteID)

						go func() {
							//defer //fmt.Println("connect to client over")
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
						}()
						//TODO:ASK Client to send request here
						client.send(req)
						select {
						case n.registerC <- client:
						case <-req.ctx.Done():
							//fmt.Println(start, "callHandler NewClient ctx ", req.ctx.Err())
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
				//fmt.Println(time.Now(), "callHandler Register", client.conn.RemoteAddr().String(), client.localID, " - ", client.remoteID)
			case _, _ = <-hangup:
			}
		}
	}()
	return
}

func (n *Server) Leave() {
	//fmt.Println("server Leave")
	err := n.listener.Close()
	if err != nil {
		//fmt.Println("listener.Close err", err)
	}
	n.cancel()

	if n.cluster != nil {
		n.cluster.Leave()
	}

	return
}

/*
This is a block call
*/
func (n *Server) ConnectTo(addr string) (id []byte, err error) {
	callReq := Request{}
	callReq.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	callReq.addr = addr
	callReq.reply = make(chan interface{})
	callReq.errc = make(chan error)
	//fmt.Println("ConnectTo addr ", addr)

	select {
	case n.calling <- callReq:
	case <-callReq.ctx.Done():
		//fmt.Println("ConnectTo ctx err ", callReq.ctx.Err(), callReq.addr, " id ", callReq.id)
		return
	}

	select {
	case r := <-callReq.reply:
		client, ok := r.(*Client)
		if ok {
			id = client.remoteID
			//fmt.Println("ConnectTo addr ", callReq.addr, " id ", id)
		}
		return

	case err = <-callReq.errc:
		//fmt.Println("ConnectTo err ", err)
		return
	case <-callReq.ctx.Done():
		//fmt.Println("ConnectTo ctx err ", callReq.ctx.Err(), callReq.addr, " id ", callReq.id)
		return
	}

	return
}

func (n *Server) Request(id []byte, m proto.Message) (msg P2PMessage, err error) {
	defer logger.TimeTrack(time.Now(), "Request", nil)
	//fmt.Println(time.Now(), "Server call Request ", id)
	start := time.Now()
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
		//fmt.Println("Request ask for clinet ctx err ", callReq.ctx.Err())
		return
	}

	select {
	case r, ok := <-callReq.reply:
		//fmt.Println(time.Now(), "Request got reply ", ok)
		if !ok {
			return
		}
		msg, ok = r.(P2PMessage)
		if ok {
			logger.TimeTrack(start, "RequestSent", nil)
		} else {
			err = errors.New("Reply cast error")
			//fmt.Println(time.Now(), "Request Reply cast error", err)

		}
		close(callReq.reply)
		close(callReq.errc)
		return
	case e, ok := <-callReq.errc:
		if ok {
			err = e
			//fmt.Println(time.Now(), "Request err ", err)
			return
		}
	case <-callReq.ctx.Done():
		err = callReq.ctx.Err()
		go func() {
			select {
			case _ = <-callReq.reply:
			case <-time.After(5 * time.Second):
			}
			close(callReq.reply)
			close(callReq.errc)
		}()
		//fmt.Println(time.Now(), "Request ctx err ", callReq.ctx.Err())
		return
	}
	return
}

func (n *Server) Reply(id []byte, nonce uint64, response proto.Message) (err error) {
	//defer //fmt.Println("Server call Reply done")
	callReq := Request{}
	//fmt.Println(time.Now(), "Server call Reply ", nonce, id)

	callReq.ctx, callReq.cancel = context.WithTimeout(context.Background(), 5*time.Second)
	callReq.id = id
	callReq.rType = 2
	callReq.nonce = nonce
	errc := make(chan error)
	callReq.errc = errc
	defer close(callReq.errc)
	callReq.msg = response
	if callReq.ctx == nil {
		//fmt.Println(time.Now(), "Server  req is nil ")
	}
	select {
	case n.replying <- callReq:
		//fmt.Println(time.Now(), "Server has been send to replying")

	case <-callReq.ctx.Done():
		//fmt.Println(time.Now(), "Request ctx err ", callReq.ctx.Err())
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
			fmt.Println(time.Now(), "Reply ctx err ", callReq.ctx.Err())
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
					//fmt.Println("msg.Msg.Message is nil")

					continue
				}
				messagetype := reflect.TypeOf(msg.Msg.Message).String()
				if len(messagetype) > 0 && messagetype[0] == '*' {
					messagetype = messagetype[1:]
				}
				////fmt.Println("subscriptions ", len(subscriptions))
				out := subscriptions[messagetype]
				if out != nil {

					select {
					case out <- msg:
						////fmt.Println("dispatch message done")
					}
				} else {
					//fmt.Println("!!!!!no dispatch message err")
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
