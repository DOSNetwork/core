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
	calling  chan request
	replying chan request

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

type request struct {
	rType  int
	ctx    context.Context
	cancel context.CancelFunc
	addr   string
	id     []byte
	//client signs and packs msg into Package
	msg proto.Message
	p   *Package
	//
	nonce uint64
	reply chan interface{}
	errc  chan error
}

func (r *request) sendReq(ch chan request) (err error) {
	select {
	case ch <- *r:
	case <-r.ctx.Done():
		err = r.ctx.Err()
	}
	return
}

func (r *request) waitForResult() (result interface{}, err error) {
	defer close(r.reply)
	defer close(r.errc)
	select {
	case result = <-r.reply:
	case err = <-r.errc:
	case <-r.ctx.Done():
		err = r.ctx.Err()
	}
	return
}
func (r *request) waitForError() (err error) {
	defer close(r.errc)
	select {
	case err = <-r.errc:
	case <-r.ctx.Done():
		err = r.ctx.Err()
	}
	return
}

func (r *request) replyResult(result interface{}) {
	defer r.cancel()
	select {
	case <-r.ctx.Done():
	case r.reply <- result:
	}
}

func (r *request) replyError(err error) {
	defer r.cancel()
	select {
	case <-r.ctx.Done():
	case r.errc <- err:
	}
}

type subscription struct {
	eventType string
	message   chan P2PMessage
}

func (n *server) Listen() (err error) {
	defer fmt.Println("Close Listen")
	n.receiveHandler()
	n.callHandler()
	go n.messageDispatch(context.Background())

	p := fmt.Sprintf(":%s", n.port)
	if n.listener, err = net.Listen("tcp", p); err != nil {
		err = &P2PError{
			err: errors.Errorf("Error in server Listen(): %w", err),
			t:   time.Now(),
		}
		logger.Error(err)
		return
	}
	fmt.Println("Listen to ", n.addr, " ", n.port)

	for {
		var fd net.Conn
		fd, err = n.listener.Accept()
		if err != nil {

			return
		}
		go func() {
			c, err := n.setupConn(fd, true)
			if err != nil {
				logger.Error(err)
				return
			}
			select {
			case n.addIncomingC <- c:
			case <-n.ctx.Done():
			}
		}()
	}

	return nil
}

// SetupConn runs the handshakes and attempts to add the connection
// as a peer. It returns when the connection has been added as a peer
// or the handshakes have failed.
func (n *server) setupConn(fd net.Conn, inOrOut bool) (c *client, err error) {
	c = newClient(n.id, fd, n.peersFeed, inOrOut)
	errc := c.handShake(context.Background())
	for err = range errc {
		c.close()
		break
	}
	return
}

func (n *server) runClient(c *client, inBound bool) {
	defer fmt.Println("Close runClient")
	_ = c.run()

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

func (n *server) receiveHandler() {
	n.addIncomingC = make(chan *client, 21)
	n.removeIncomingC = make(chan []byte)
	n.replying = make(chan request)
	clients := make(map[string]*client)

	go func() {
		defer fmt.Println("Close receiveHandler")
		for {
			select {
			case <-n.ctx.Done():
				return
			case c := <-n.addIncomingC:
				if clients[string(c.remoteID)] != nil {
					c.close()
					continue
				}
				clients[string(c.remoteID)] = c
				n.incomingNum = len(clients)
				fmt.Println("add inbound ", len(clients))
				go n.runClient(c, true)
			case id := <-n.removeIncomingC:
				if c := clients[string(id)]; c != nil {
					delete(clients, string(id))
				}
				n.incomingNum = len(clients)
				fmt.Println("remove inbound ", len(clients))
			case req := <-n.replying:
				fmt.Println("!!replying ", string(req.id))
				client := clients[string(req.id)]
				if client != nil {
					client.send(req)
				} else {
					err := &P2PError{
						err: errors.New("p2p reply can't find client"),
						t:   time.Now(),
					}
					logger.Error(err)
				}
			}
		}
	}()
}

func (n *server) callHandler() {
	n.calling = make(chan request)
	addrToid := make(map[string][]byte)
	clients := make(map[string]*client)
	sendCache := make(map[string][]request)
	n.removeCallingC = make(chan []byte)
	n.addCallingC = make(chan *client)

	go func() {
		defer fmt.Println("Close callHandler")
		for {
			select {
			case <-n.ctx.Done():
				return
			case c, ok := <-n.addCallingC:
				if !ok {
					return
				}
				go n.runClient(c, false)
				clients[string(c.remoteID)] = c
				n.callingNum = len(clients)
				addrToid[c.conn.RemoteAddr().String()] = c.remoteID
				fmt.Println("addCallingC ", len(clients), len(sendCache[string(c.remoteID)]))
				for _, req := range sendCache[string(c.remoteID)] {
					go c.send(req)
				}

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
				fmt.Println("outbound p2p remove ", len(clients))
			case req, ok := <-n.calling:
				if !ok {
					return
				}
				var c *client
				fmt.Println("calling ", string(req.id))
				if c = clients[string(req.id)]; c == nil {
					req.addr = n.members.Lookup(req.id)
					sendCache[string(req.id)] = append(sendCache[string(req.id)], req)
					go func() {
						fd, err := dialTo(req.ctx, req.addr)
						if err != nil {
							req.replyError(err)
							return
						}
						if c, err = n.setupConn(fd, true); err != nil {
							c.close()
							req.replyError(err)
							return
						}
						select {
						case <-req.ctx.Done():
						case n.addCallingC <- c:
						}
					}()
				} else {
					go c.send(req)
				}

			}
		}
	}()
	return
}

func dialTo(ctx context.Context, addr string) (fd net.Conn, err error) {
	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
			//if fd, err = net.DialTimeout("tcp", addr, 2*time.Second); err == nil {
			if fd, err = net.Dial("tcp", addr); err == nil {
				return
			} else {
				fmt.Println("dialTo fail err", err)
			}
		}
	}
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

// Request sends a proto message to the specific node
func (n *server) Request(ctx context.Context, id []byte, m proto.Message) (msg P2PMessage, err error) {
	var (
		result interface{}
		ok     bool
	)
	rctx, cancel := context.WithCancel(ctx)
	req := request{ctx: rctx, cancel: cancel, id: id, msg: m, rType: sendReq, reply: make(chan interface{}), errc: make(chan error)}
	if err = req.sendReq(n.calling); err != nil {
		return
	}
	if result, err = req.waitForResult(); err != nil {
		err = &P2PError{
			err: errors.New("Error in waitForResult error "),
			t:   time.Now(),
		}
		logger.Error(err)
		return
	}
	if msg, ok = result.(P2PMessage); !ok {
		err = &P2PError{
			err: errors.New("Error in Request casting error "),
			t:   time.Now(),
		}
		logger.Error(err)
	}
	return
}

// Reply sends a reply to the specific node
func (n *server) Reply(ctx context.Context, id []byte, nonce uint64, msg proto.Message) (err error) {
	fmt.Println("Server reply")
	rctx, cancel := context.WithCancel(ctx)
	req := request{ctx: rctx, cancel: cancel, id: id, rType: replyReq, nonce: nonce, msg: msg, errc: make(chan error)}
	if err = req.sendReq(n.replying); err != nil {
		fmt.Println("er ", err)
		return
	}

	return req.waitForError()
}

func (n *server) messageDispatch(ctx context.Context) {
	subscriptions := make(map[string]chan P2PMessage)
	go func() {
		for {
			select {
			case msg, ok := <-n.peersFeed:
				if !ok {
					return
				}

				if msg.Msg.Message == nil {

					continue
				}
				messagetype := reflect.TypeOf(msg.Msg.Message).String()
				fmt.Println("peersFeed ", messagetype)
				if len(messagetype) > 0 && messagetype[0] == '*' {
					messagetype = messagetype[1:]
				}
				out := subscriptions[messagetype]
				if out != nil {
					fmt.Println("peersFeed out", messagetype)
					go func(msg P2PMessage) {
						select {
						case out <- msg:
						}
					}(msg)
				}
			case sub, ok := <-n.subscribe:
				if !ok {
					return
				}
				fmt.Println("subscriptions out", sub.eventType)
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
func (n *server) SubscribeEvent(chanBuffer int, peersFeed ...interface{}) (outch chan P2PMessage, err error) {
	if chanBuffer > 0 {
		outch = make(chan P2PMessage, chanBuffer)
	} else {
		outch = make(chan P2PMessage)
	}
	for _, m := range peersFeed {
		fmt.Println("SubscribeEvent ", reflect.TypeOf(m).String())

		n.subscribe <- subscription{reflect.TypeOf(m).String(), outch}
	}
	return
}

// UnSubscribeEvent is a un-subscription operation
func (n *server) UnSubscribeEvent(peersFeed ...interface{}) {
	for _, m := range peersFeed {
		n.unscribe <- subscription{reflect.TypeOf(m).String(), nil}

	}
	return
}

func (n *server) Join(bootstrapIP []string) (num int, err error) {
	return n.members.Join(bootstrapIP)
}

func (n *server) NumOfMembers() int {
	return n.members.NumOfPeers()
}

func (n *server) MembersIP() []net.IP {
	return n.members.MembersIP()
}

func (n *server) MembersID() [][]byte {
	return n.members.MembersID()
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

func (n *server) GetPort() string {
	return n.port
}

func (n *server) GetID() []byte {
	return n.id
}

func (n *server) GetIP() net.IP {
	return n.addr
}
