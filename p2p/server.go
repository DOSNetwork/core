package p2p

import (
	//	"bytes"
	"context"
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/dedis/kyber"

	"github.com/golang/protobuf/proto"

	"github.com/DOSNetwork/core/p2p/discover"
	"github.com/DOSNetwork/core/suites"
	errors "golang.org/x/xerrors"
)

const (
	dialReq = iota
	sendReq
	replyReq
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
	calling  chan p2pRequest
	replying chan p2pRequest

	addIncomingC    chan *client
	removeIncomingC chan []byte
	addCallingC     chan *client
	removeCallingC  chan []byte
	incomingNum     int
	callingNum      int

	//Event
	peersFeed chan P2PMessage
	subscribe chan subscription
	unscribe  chan subscription

	ctx    context.Context
	cancel context.CancelFunc
}

type subscription struct {
	eventType string
	message   chan P2PMessage
}

func (n *server) Listen() (err error) {
	defer fmt.Println("end Listen")
	go n.receiveHandler()
	go n.callHandler()
	go n.messageDispatch()

	p := fmt.Sprintf(":%s", n.port)
	if n.listener, err = net.Listen("tcp", p); err != nil {
		logger.Error(errors.Errorf("server Listen : %w", err))
		return
	}
	fmt.Println("Listen to ", n.addr, " ", n.port)

	for {
		select {
		case <-n.ctx.Done():
			return
		default:
			var fd net.Conn
			fd, err = n.listener.Accept()
			if err != nil {
				logger.Error(errors.Errorf("server Listen : %w", err))
				return
			}
			go func() {
				ctx, cancel := context.WithTimeout(n.ctx, 2*time.Second)
				defer cancel()
				c := newClient(n.id, fd, n.peersFeed, true)
				errc := c.handShake(ctx)
				for err = range errc {
					err = errors.Errorf("server Listen : %w", err)
					logger.Error(err)
				}
				if err != nil {
					c.close()
					return
				}
				select {
				case <-n.ctx.Done():
					return
				case n.addIncomingC <- c:
				}
				return
			}()
		}
	}

	return
}

func (n *server) receiveHandler() {
	defer fmt.Println("end receiveHandler")
	clients := make(map[string]*client)
	for {
		select {
		case <-n.ctx.Done():
			for _, client := range clients {
				client.close()
			}
			return
		case c := <-n.addIncomingC:
			if clients[string(c.remoteID)] != nil {
				c.close()
				continue
			}
			clients[string(c.remoteID)] = c
			n.incomingNum = len(clients)
			go n.runClient(c, true)
			fmt.Println("add inbound ", len(clients))
		case id := <-n.removeIncomingC:
			if c := clients[string(id)]; c != nil {
				delete(clients, string(id))
			}
			n.incomingNum = len(clients)
			fmt.Println("remove inbound ", len(clients))
		case req := <-n.replying:
			client := clients[string(req.id)]
			if client == nil {
				logger.Error(errors.Errorf("server reply: %w", ErrCanNotFindClient))
				req.cancel()
				go client.send(req)
			}
			go client.send(req)
		}
	}
}

func (n *server) callHandler() {
	defer fmt.Println("end callHandler")
	addrToid := make(map[string][]byte)
	clients := make(map[string]*client)

	for {
		select {
		case <-n.ctx.Done():
			for _, client := range clients {
				client.close()
			}
			return
		case id, ok := <-n.removeCallingC:
			if ok {
				c := clients[string(id)]
				if c != nil {
					delete(addrToid, c.conn.RemoteAddr().String())
					delete(clients, string(id))
					fmt.Println("outbound p2p remove ", len(clients))
				}
				n.callingNum = len(clients)
			}
		case req, ok := <-n.calling:
			if ok {
				var c *client
				if c = clients[string(req.id)]; c == nil {
					//TODO : performance bottleneck
					req.addr = n.members.Lookup(req.id)
					if c = n.handleCallReq(req); c == nil {
						continue
					}
					clients[string(req.id)] = c
					go n.runClient(c, false)
				}
				go c.send(req)
			}
		}
	}
	return
}

func (n *server) messageDispatch() {
	defer fmt.Println("end messageDispatch")
	subscriptions := make(map[string]chan P2PMessage)
	for {
		select {
		case msg, ok := <-n.peersFeed:
			if ok {
				if msg.Msg.Message == nil {
					continue
				}
				messagetype := reflect.TypeOf(msg.Msg.Message).String()
				if len(messagetype) > 0 && messagetype[0] == '*' {
					messagetype = messagetype[1:]
				}
				if out := subscriptions[messagetype]; out != nil {
					go func() {
						select {
						case <-n.ctx.Done():
						case out <- msg:
						}
					}()
				}
			}
		case sub, ok := <-n.subscribe:
			if ok {
				subscriptions[sub.eventType] = sub.message
			}
		case sub, ok := <-n.unscribe:
			if ok {
				delete(subscriptions, sub.eventType)
			}
		case <-n.ctx.Done():
			for _, outch := range subscriptions {
				close(outch)
			}
			return
		}
	}
}

func (n *server) runClient(c *client, inBound bool) {
	if err := c.run(); err != nil {
		fmt.Println("runClient err", err)
		logger.Error(err)
	}
	c.close()
	var delpeer chan []byte
	if inBound {
		delpeer = n.removeIncomingC
	} else {
		delpeer = n.removeCallingC
	}
	select {
	case <-n.ctx.Done():
		return
	case delpeer <- c.remoteID:
	}
}

func (n *server) handleCallReq(req p2pRequest) (c *client) {
	var fd net.Conn
	var err error

	if fd, err = net.Dial("tcp", req.addr); err != nil {
		req.replyResult(nil, errors.Errorf("server handleCallReq %s: %w", req.addr, err))
		return
	}

	c = newClient(n.id, fd, n.peersFeed, true)
	errc := c.handShake(req.ctx)
	for err = range errc {
		err = errors.Errorf("server handleCallReq : %w", err)
	}
	if err != nil {
		req.replyResult(nil, errors.Errorf("server handleCallReq : %w", err))
		c.close()
		c = nil
	}
	return
}

func (n *server) Leave() {
	n.cancel()
	err := n.listener.Close()
	if err != nil {
		logger.Error(err)
	}

	if n.members != nil {
		n.members.Leave()
	}
	return
}

func (n *server) DisConnectTo(id []byte) (err error) {
	select {
	case <-n.ctx.Done():
		return errors.Errorf("server DisConnectTo : %w", n.ctx.Err())
	case n.removeCallingC <- id:
	}
	return
}

// Request sends a proto message to the specific node
func (n *server) Request(ctx context.Context, id []byte, msg proto.Message) (p2pmsg P2PMessage, err error) {
	var (
		result interface{}
		ok     bool
	)
	req := NewP2pRequest(ctx, sendReq, id, "", msg, 0)
	if err = req.sendReq(n.calling); err != nil {
		err = errors.Errorf("server Request : %w", err)
		logger.Error(err)
		return
	}
	if result, err = req.waitForResult(); err != nil {
		err = errors.Errorf("server Request : %w", err)
		logger.Error(err)
		return
	}

	if p2pmsg, ok = result.(P2PMessage); !ok {
		err = errors.Errorf("server Request : %w", ErrCasting)
		logger.Error(err)
	}
	return
}

// Reply sends a reply to the specific node
func (n *server) Reply(ctx context.Context, id []byte, nonce uint64, msg proto.Message) (err error) {
	req := NewP2pRequest(ctx, replyReq, id, "", msg, nonce)
	if err = req.sendReq(n.replying); err != nil {
		err = errors.Errorf("server Reply : %w", err)
		logger.Error(err)
		return
	}
	if _, err = req.waitForResult(); err != nil {
		err = errors.Errorf("server Reply : %w", err)
		logger.Error(err)
	}
	return
}

// SubscribeEvent is a message subscription operation binding the P2PMessage
func (n *server) SubscribeEvent(chanBuffer int, peersFeed ...interface{}) (outch chan P2PMessage, err error) {
	if chanBuffer > 0 {
		outch = make(chan P2PMessage, chanBuffer)
	} else {
		outch = make(chan P2PMessage)
	}
	for _, m := range peersFeed {
		fmt.Println("SubscribeEvent ", reflect.TypeOf(m).String())
		select {
		case <-n.ctx.Done():
		case n.subscribe <- subscription{reflect.TypeOf(m).String(), outch}:
		}
	}
	return
}

// UnSubscribeEvent is a un-subscription operation
func (n *server) UnSubscribeEvent(peersFeed ...interface{}) {
	for _, m := range peersFeed {
		select {
		case <-n.ctx.Done():
		case n.unscribe <- subscription{reflect.TypeOf(m).String(), nil}:
		}
	}
	return
}

func (n *server) Join(bootstrapIP []string) (num int, err error) {
	select {
	case <-n.ctx.Done():
		err = n.ctx.Err()
	default:
		num, err = n.members.Join(bootstrapIP)
	}
	return
}

func (n *server) NumOfMembers() (num int) {
	select {
	case <-n.ctx.Done():
	default:
		num = n.members.NumOfPeers()
	}
	return
}

func (n *server) MembersIP() (ips []net.IP) {
	select {
	case <-n.ctx.Done():
	default:
		ips = n.members.MembersIP()
	}
	return
}

func (n *server) MembersID() (ids [][]byte) {
	select {
	case <-n.ctx.Done():
	default:
		ids = n.members.MembersID()
	}
	return
}

func (n *server) numOfClient() (iNum, cNum int) {
	select {
	case <-n.ctx.Done():
	default:
		iNum, cNum = n.incomingNum, n.callingNum
	}
	return
}

func (n *server) SetID(id []byte) {
	select {
	case <-n.ctx.Done():
	default:
		n.id = id
	}
	return
}

func (n *server) SetPort(port string) {
	select {
	case <-n.ctx.Done():
	default:
		n.port = port
	}
	return
}

func (n *server) GetPort() (port string) {
	select {
	case <-n.ctx.Done():
	default:
		port = n.port
	}
	return
}

func (n *server) GetID() (id []byte) {
	select {
	case <-n.ctx.Done():
	default:
		id = n.id
	}
	return
}

func (n *server) GetIP() (ip net.IP) {
	select {
	case <-n.ctx.Done():
	default:
		ip = n.addr
	}
	return
}
