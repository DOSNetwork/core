package p2p

import (
	//	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"sync"
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

type logger interface {
	Info(msg string)
	Error(err error)
	TimeTrack(start time.Time, e string, info map[string]interface{})
	Event(e string, info map[string]interface{})
}

type server struct {
	logger
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

	//Msg
	peersFeed    chan P2PMessage
	subscribeMsg chan *subscription
	unscribeMsg  chan string

	//Event
	peersEvent     chan discover.P2PEvent
	subscribeEvent chan *subscription
	unscribeEvent  chan int

	ctx    context.Context
	cancel context.CancelFunc
}

type subscription struct {
	subID   int
	msgType string
	msgCh   chan P2PMessage
	eventCh chan discover.P2PEvent
	replyCh chan *subscription
	err     error
}

func (n *server) Listen() (err error) {
	defer n.logger.Info("[P2P] End P2P ListenLoop")

	p := fmt.Sprintf(":%s", n.port)
	if n.listener, err = net.Listen("tcp", p); err != nil {
		err = &P2PError{err: errors.Errorf("server Listen failed: %w", err), t: time.Now()}
		n.logger.Error(err)
		return
	}
	n.logger.Info(fmt.Sprintf("Listen to %s:%s", n.addr, n.port))

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		n.receiveHandler()
		n.logger.Info("[P2P] End P2P receiveHandler")
	}()
	go func() {
		defer wg.Done()
		err = n.callHandler()
		n.logger.Info("[P2P] End P2P callHandler")
		n.listener.Close()
	}()
	go func() {
		defer wg.Done()
		n.messageDispatch()
		n.logger.Info("[P2P] End P2P messageDispatch")
	}()
	go func() {
		defer wg.Done()
		n.eventDispatch()
		n.logger.Info("[P2P] End P2P eventDispatch")
	}()
L:
	for {
		select {
		case <-n.ctx.Done():
			break L
		default:
			fd, errL := n.listener.Accept()
			if errL != nil {
				if err != nil {
					err = &P2PError{err: errors.Errorf("listener accept failed: %w", err), t: time.Now()}
				} else {
					err = &P2PError{err: errors.Errorf("listener accept failed: %w", errL), t: time.Now()}
				}
				n.logger.Error(err)
				n.cancel()
				break L
			}
			go func() {
				ctx, cancel := context.WithTimeout(n.ctx, 2*time.Second)
				defer cancel()
				c := newClient(n.id, fd, n.peersFeed, true)
				errc := c.handShake(ctx)
				for err = range errc {
					err = &P2PError{err: errors.Errorf("handShake failed: %w", err), t: time.Now()}
					n.logger.Error(err)
				}
				if err != nil {
					if err = c.close(); err != nil {
						err = &P2PError{err: errors.Errorf("conn closed failed: %w", err), t: time.Now()}
						n.logger.Error(err)
					}
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
	wg.Wait()
	return
}

func (n *server) receiveHandler() {
	clients := make(map[string]*client)
	for {
		select {
		case <-n.ctx.Done():
			for _, client := range clients {
				if err := client.close(); err != nil {
					err = &P2PError{err: errors.Errorf("conn close failed: %w", err), t: time.Now()}
					n.logger.Error(err)
				}
			}
			return
		case c := <-n.addIncomingC:
			if clients[string(c.remoteID)] != nil {
				if err := c.close(); err != nil {
					err = &P2PError{err: errors.Errorf("conn close failed: %w", err), t: time.Now()}
					n.logger.Error(err)
				}
				continue
			}
			clients[string(c.remoteID)] = c
			n.incomingNum = len(clients)
			go n.runClient(c, true)
		case id := <-n.removeIncomingC:
			if c := clients[string(id)]; c != nil {
				delete(clients, string(id))
			}
			n.incomingNum = len(clients)
		case req := <-n.replying:
			client := clients[string(req.id)]
			if client == nil {
				err := &P2PError{err: errors.Errorf("reply failed: %w", ErrCanNotFindClient), t: time.Now()}
				n.logger.Error(err)
				req.replyResult(nil, err)
				continue
			}
			go client.send(req)
		}
	}
}

func (n *server) callHandler() (err error) {
	addrToid := make(map[string][]byte)
	clients := make(map[string]*client)
	watchDog := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-n.ctx.Done():
			for _, client := range clients {
				if err := client.close(); err != nil {
					err = &P2PError{err: errors.Errorf("conn close failed: %w", err), t: time.Now()}
					n.logger.Error(err)
				}
			}
			err = n.ctx.Err()
			return
		case <-watchDog.C:
			if !n.members.IsAlive() {
				err = errors.New("p2p cluster status is not alive")
				n.logger.Error(err)
				return
			}
		case id, ok := <-n.removeCallingC:
			if ok {
				c := clients[string(id)]
				if c != nil {
					delete(addrToid, c.conn.RemoteAddr().String())
					delete(clients, string(id))
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

func (n *server) runClient(c *client, inBound bool) {
	if err := c.run(); err != nil {
		err = &P2PError{err: errors.Errorf("runClient (%t) err: %w", inBound, err), t: time.Now()}
		n.logger.Error(err)
	}

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
		err = &P2PError{err: errors.Errorf("dial %s failed: %w", req.addr, err), t: time.Now()}
		req.replyResult(nil, err)
		n.logger.Error(err)
		return
	}

	c = newClient(n.id, fd, n.peersFeed, true)
	errc := c.handShake(req.ctx)
	for err = range errc {
		err = &P2PError{err: errors.Errorf("server handleCallReq: %w", err), t: time.Now()}
		n.logger.Error(err)
	}
	if err != nil {
		req.replyResult(nil, err)
		err = c.close()
		if err != nil {
			err = &P2PError{err: errors.Errorf("close failed in handleCallReq: %w", err), t: time.Now()}
			n.logger.Error(err)
		}
		c = nil
	}
	return
}

func (n *server) Leave() {
	n.cancel()
	err := n.listener.Close()
	if err != nil {
		err = &P2PError{err: errors.Errorf("listener.Close failed: %w", err), t: time.Now()}
		n.logger.Error(err)
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
	addr := n.members.Lookup(id)
	if addr == "" {
		err = &P2PError{err: errors.New("No IP info"), dest: addr, t: time.Now()}
		n.logger.Error(err)
		return
	}
	defer n.logger.TimeTrack(time.Now(), "TimeRequest", nil)
	opCtx, opCancel := context.WithTimeout(ctx, 5*time.Second)
	defer opCancel()
	req := NewP2pRequest(opCtx, sendReq, id, "", msg, 0)
	if err = req.sendReq(n.calling); err != nil {
		err = &P2PError{err: errors.Errorf("Request sendReq failed: %w", err), dest: addr, t: time.Now()}
		n.logger.Error(err)
		return
	}
	if result, err = req.waitForResult(); err != nil {
		err = &P2PError{err: errors.Errorf("Request waitForResult: %w", err), dest: addr, t: time.Now()}
		n.logger.Error(err)
		return
	}

	if p2pmsg, ok = result.(P2PMessage); !ok {
		err = &P2PError{err: errors.Errorf("Request P2PMessage casting failed: %w", err), dest: addr, t: time.Now()}
		n.logger.Error(err)
	}
	return
}

// Reply sends a reply to the specific node
func (n *server) Reply(ctx context.Context, id []byte, nonce uint64, msg proto.Message) (err error) {
	opCtx, opCancel := context.WithTimeout(ctx, 15*time.Second)
	defer opCancel()
	req := NewP2pRequest(opCtx, replyReq, id, "", msg, nonce)
	if err = req.sendReq(n.replying); err != nil {
		err = &P2PError{err: errors.Errorf("Reply sendReq failed: %w", err), dest: req.addr, t: time.Now()}
		n.logger.Error(err)
		return
	}
	if _, err = req.waitForResult(); err != nil {
		err = &P2PError{err: errors.Errorf("Reply waitForResult failed: %w", err), dest: req.addr, t: time.Now()}
		n.logger.Error(err)
	}
	return
}

func (n *server) eventDispatch() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		n.members.Listen(n.ctx, n.peersEvent)
	}()

	subscriptions := make(map[int]chan discover.P2PEvent)
	subID := 1
L:
	for {
		select {
		case e, ok := <-n.peersEvent:
			if ok {
				for _, eventCh := range subscriptions {
					select {
					case <-n.ctx.Done():
					case eventCh <- e:
					}
				}
			}
		case sub, ok := <-n.subscribeEvent:
			if ok {
				sub.eventCh = make(chan discover.P2PEvent)
				sub.subID = subID
				subID++
				subscriptions[sub.subID] = sub.eventCh
				select {
				case <-n.ctx.Done():
				case sub.replyCh <- sub:
				}
			}
		case subID, ok := <-n.unscribeEvent:
			if ok {
				close(subscriptions[subID])
				delete(subscriptions, subID)
			}
		case <-n.ctx.Done():
			n.members.Leave()
			for _, eventCh := range subscriptions {
				close(eventCh)
			}
			break L
		}
	}
	wg.Wait()
	return
}

func (n *server) messageDispatch() {
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
					select {
					case <-n.ctx.Done():
					case out <- msg:
					}
				}
			}
		case sub, ok := <-n.subscribeMsg:
			if ok {
				subscriptions[sub.msgType] = sub.msgCh
			}
		case msgType, ok := <-n.unscribeMsg:
			if ok {
				delete(subscriptions, msgType)
			}
		case <-n.ctx.Done():
			for _, outch := range subscriptions {
				if outch != nil {
					close(outch)
				}
			}
			return
		}
	}
}

// SubscribeEvent is a message subscription operation binding the P2PMessage
func (n *server) SubscribeEvent() (subID int, outch chan discover.P2PEvent, err error) {
	reply := make(chan *subscription, 1)
	defer close(reply)
	select {
	case <-n.ctx.Done():
	case n.subscribeEvent <- &subscription{replyCh: reply}:
	}
	select {
	case <-n.ctx.Done():
		err = n.ctx.Err()
	case r := <-reply:
		subID = r.subID
		outch = r.eventCh
		err = r.err
	}
	return
}

// UnSubscribeEvent is a un-subscription operation
func (n *server) UnSubscribeEvent(subID int) {
	select {
	case <-n.ctx.Done():
	case n.unscribeEvent <- subID:
	}
	return
}

func merge(ctx context.Context, cs ...chan P2PMessage) chan P2PMessage {
	var wg sync.WaitGroup
	out := make(chan P2PMessage)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan P2PMessage) {
		for n := range c {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// SubscribeMsg is a message subscription operation binding the P2PMessage
func (n *server) SubscribeMsg(chanBuffer int, peersFeed ...interface{}) (outch chan P2PMessage, err error) {

	var eventList []chan P2PMessage
	for _, m := range peersFeed {
		if chanBuffer > 0 {
			outch = make(chan P2PMessage, chanBuffer)
		} else {
			outch = make(chan P2PMessage)
		}
		eventList = append(eventList, outch)
		select {
		case <-n.ctx.Done():
		case n.subscribeMsg <- &subscription{msgType: reflect.TypeOf(m).String(), msgCh: outch}:
		}
	}
	outch = merge(n.ctx, eventList...)
	return
}

// UnSubscribeEvent is a un-subscription operation
func (n *server) UnSubscribeMsg(peersFeed ...interface{}) {
	for _, m := range peersFeed {
		select {
		case <-n.ctx.Done():
		case n.unscribeMsg <- reflect.TypeOf(m).String():
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
		if err != nil {
			n.logger.Error(errors.Errorf(" : %w", err))
		}
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

func (n *server) RandomPeerIP() (ips []string) {
	select {
	case <-n.ctx.Done():
	default:
		slice := n.members.MembersIP()
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
		for _, ip := range slice {
			ips = append(ips, ip.String())
		}
	}
	return
}

func (n *server) MembersID() (ids [][]byte) {
	select {
	case <-n.ctx.Done():
	default:
		n.members.MembersIP()
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
