package p2p

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/p2p/internal"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

// The node itself
type Client struct {
	identity dht.ID
	//Map of connection addresses (string) <-> *p2p.Peer
	peers *sync.Map
	// Channels are thread safe
	messageChan  chan P2PMessage
	suite        suites.Suite
	port         int
	secKey       kyber.Scalar
	pubKey       kyber.Point
	routingTable *dht.RoutingTable
}

// Structure holding information of each connected peer.
type Peer struct {
	p2pnet       *Client
	conn         *net.Conn
	rw           *bufio.ReadWriter
	messageChan  chan P2PMessage
	status       int
	identity     dht.ID
	pubKey       kyber.Point
	wg           sync.WaitGroup
	RequestNonce uint64
	Requests     sync.Map
}

// RequestState represents a state of a request.
type RequestState struct {
	data        chan proto.Message
	closeSignal chan struct{}
}

func (n *Client) GetId() dht.ID {
	return n.identity
}

func (n *Client) Listen() error {
	ip, err := getLocalIp()
	if err != nil {
		log.Fatal(err)
	}

	p := fmt.Sprintf(":%d", n.port)

	listener, err := net.Listen("tcp", p)
	if err != nil {
		return err
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

	n.port = listener.Addr().(*net.TCPAddr).Port
	ip = ip + ":" + strconv.Itoa(n.port)

	n.secKey, n.pubKey = genPair()
	pubKeyBytes, err := n.pubKey.MarshalBinary()
	if err != nil {
		return err
	}

	n.identity = dht.CreateID(ip, pubKeyBytes)
	n.routingTable = dht.CreateRoutingTable(n.identity)

	fmt.Println("listen on: ", ip, " id: ", n.identity.Id)
	// Handle new clients.

	go func() {
		for {
			if conn, err := listener.Accept(); err == nil {
				if err != nil {
					log.Fatal(err)
				}
				// connect to peer based on accepted connection
				n.AcceptConnectionFromPeer(&conn)
			} else {
				fmt.Println("Failed accepting a connection request:", err)
				log.Fatal(err)
			}
		}
	}()

	return nil
}

// Prepare message going to be sent to peers.
func (n *Client) PrepareMessage(msg proto.Message) (*internal.Package, error) {
	if msg == nil {
		return nil, errors.New("network: message is null")
	}

	id := internal.ID(n.identity)
	anything, _ := ptypes.MarshalAny(msg)
	sig, _ := bls.Sign(n.suite, n.secKey, anything.Value)
	pub, _ := n.pubKey.MarshalBinary()

	pkg := &internal.Package{
		Sender:    &id,
		Anything:  anything,
		Pubkey:    pub,
		Signature: sig,
	}

	return pkg, nil
}

// Send prepared message to peer p.
func (n *Client) SendPackage(p *Peer, msg proto.Message) error {
	if msg == nil {
		return errors.New("network: message is null")
	}
	//Encode the package
	bytes, _ := proto.Marshal(msg)
	// Serialize size.
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(bytes)))

	bytes = append(prefix, bytes...)
	if _, err := p.rw.Write(bytes); err != nil {
		return err
	}
	if err := p.rw.Flush(); err != nil {
		return err
	}

	return nil
}

// Client pings peer for first time.
func (n *Client) SayHi(p *Peer) {
	fmt.Printf("%v say hi to %v\n", n.identity.Address, p.identity.Address)
	msg := &internal.Hi{
		PublicKey: n.identity.PublicKey,
		Address:   n.identity.Address,
		Id:        n.identity.Id,
	}

	prepared, err := n.PrepareMessage(msg)
	if err != nil {
		log.Fatal(err)
	}

	prepared.RequestNonce = atomic.AddUint64(&p.RequestNonce, 1)
	if err := n.SendPackage(p, prepared); err != nil {
		log.Fatal(err)
	}

	p.wg.Wait()
}

// Handle packages from peer p
func (n *Client) HandlePackages(p *Peer) {
	for {
		buf, err := n.receivePackage(p)
		if err != nil {
			fmt.Println("Peer ", p.identity.Id, " end")
			return
		}

		pkg, ptr, err := n.decodePackage(buf)
		if err != nil {
			log.Fatal(err)
		}

		if pkg.GetRequestNonce() > 0 && pkg.GetReplyFlag() {
			if _state, exists := p.Requests.Load(pkg.GetRequestNonce()); exists {
				state := _state.(*RequestState)
				select {
				case state.data <- ptr.Message:
				case <-state.closeSignal:
				}
				continue
			}
		}

		switch content := ptr.Message.(type) {
		case *internal.Hi:
			if len(p.identity.Id) == 0 {
				p.identity.Id = content.GetId()
				p.identity.Address = content.GetAddress()
				p.identity.PublicKey = content.GetPublicKey()
				pub := suite.G2().Point()
				if err = pub.UnmarshalBinary(content.GetPublicKey()); err != nil {
					log.Fatal(err)
				}
				p.pubKey = pub
				n.peers.LoadOrStore(string(p.identity.Id), p)
				n.routingTable.Update(p.identity)
				fmt.Println(n.identity.Address, "Receive Hi id = ", p.identity.Id, "from", p.identity.Address)
				p.wg.Done()
			} else {
				fmt.Println("Ignore Hi")
			}
		case *internal.Ping:
			// Send pong to peer.
			err := n.Reply(p, pkg.GetRequestNonce(), &internal.Pong{})
			if err != nil {
				log.Fatal(err)
			}
		case *internal.Pong:
			peers := n.FindNode(p.identity, dht.BucketSize, 8)

			// Update routing table w/ closest peers to self.
			for _, peerID := range peers {
				n.routingTable.Update(peerID)
			}

			glog.Infof("bootstrapped w/ peer(s): %s.", strings.Join(n.routingTable.GetPeerAddresses(), ", "))
		case *internal.LookupNodeRequest:
			// Prepare response.
			response := &internal.LookupNodeResponse{}

			// Respond back with closest peers to a provided target.
			for _, peerID := range n.routingTable.FindClosestPeers(dht.ID(*content.GetTarget()), dht.BucketSize) {
				id := internal.ID(peerID)
				response.Peers = append(response.Peers, &id)
			}

			err := n.Reply(p, pkg.GetRequestNonce(), response)
			if err != nil {
				log.Fatal(err)
			}

			glog.Infof("connected peers: %s.", strings.Join(n.routingTable.GetPeerAddresses(), ", "))
		default:
			msg := P2PMessage{Msg: *ptr, Sender: p.identity.Id}
			p.messageChan <- msg
		}
		//fmt.Println("Peer ", p.id, " receive", string(buf))

	}
}

// Receive package from peer p.
func (n *Client) receivePackage(p *Peer) ([]byte, error) {
	buf, err := ioutil.ReadAll(*p.conn)
	if err != nil {
		return nil, err
	} else if len(buf) < 4 {
		return nil, errors.New("received truncated package with size less than 4!")
	}

	// Decode message size.
	size := binary.BigEndian.Uint32(buf[:4])
	if size == 0 {
		return nil, errors.New("received an empty message from a peer")
	} else if size+4 > uint32(len(buf)) {
		return buf[4:], errors.New("received truncated message")
	} else {
		return buf[4:], nil
	}
}

func (n *Client) decodePackage(bytes []byte) (*internal.Package, *ptypes.DynamicAny, error) {
	pkg := new(internal.Package)
	if err := proto.Unmarshal(bytes, pkg); err != nil {
		return nil, nil, err
	}

	//Todo verify pkg.Signature by public key
	pub := suite.G2().Point()
	_ = pub.UnmarshalBinary(pkg.GetPubkey())

	if err := bls.Verify(n.suite, pub, pkg.GetAnything().Value, pkg.GetSignature()); err != nil {
		return nil, nil, err
	}

	var ptr ptypes.DynamicAny
	if err := ptypes.UnmarshalAny(pkg.GetAnything(), &ptr); err != nil {
		return nil, nil, err
	}
	return pkg, &ptr, nil
}

// Request requests for a response for a request sent to a given peer p.
func (n *Client) Request(p *Peer, req *Request) (proto.Message, error) {
	prepared, err := n.PrepareMessage(req.Message)
	if err != nil {
		return nil, err
	}

	prepared.RequestNonce = atomic.AddUint64(&p.RequestNonce, 1)
	if err := n.SendPackage(p, prepared); err != nil {
		return nil, err
	}

	// Start tracking the request.
	channel := make(chan proto.Message, 1)
	closeSignal := make(chan struct{})

	p.Requests.Store(prepared.GetRequestNonce(), &RequestState{
		data:        channel,
		closeSignal: closeSignal,
	})

	// Stop tracking the request.
	defer close(closeSignal)
	defer p.Requests.Delete(prepared.GetRequestNonce())

	select {
	case res := <-channel:
		return res, nil
	case <-time.After(req.Timeout):
		return nil, errors.New("request timed out")
	}
}

// Reply is equivalent to Write() with an appended nonce to signal a reply.
func (n *Client) Reply(p *Peer, nonce uint64, message proto.Message) error {
	prepared, err := n.PrepareMessage(message)
	if err != nil {
		return err
	}

	// Set the nonce.
	prepared.RequestNonce = nonce
	prepared.ReplyFlag = true

	if err := n.SendPackage(p, prepared); err != nil {
		return err
	}

	return nil
}

// Broadcast m to all peers.
func (n *Client) Broadcast(m proto.Message) {
	prepared, err := n.PrepareMessage(m)
	if err != nil {
		log.Fatal(err)
	}

	n.peers.Range(func(key, value interface{}) bool {
		peer := value.(*Peer)
		prepared.RequestNonce = atomic.AddUint64(&peer.RequestNonce, 1)
		if err := n.SendPackage(peer, prepared); err != nil {
			// TODO: Error handlding, shouldn't fail/exit when not reachable.
			log.Fatal(err)
		}
		return true
	})
}

// Send msg m to peer selected by id.
func (n *Client) SendMessageById(id []byte, m proto.Message) {
	value, loaded := n.peers.Load(string(id))
	if loaded {
		prepared, err := n.PrepareMessage(m)
		if err != nil {
			log.Fatal(err)
		}

		peer := value.(*Peer)
		prepared.RequestNonce = atomic.AddUint64(&peer.RequestNonce, 1)
		if err := n.SendPackage(peer, prepared); err != nil {
			log.Fatal(err)
		}
	}
}

func (n *Client) GetTunnel() chan P2PMessage {
	return n.messageChan
}

func (n *Client) GetRoutingTable() *dht.RoutingTable {
	return n.routingTable
}

func (n *Client) InitPeerDataStructure(p *Peer) {
	p.rw = bufio.NewReadWriter(bufio.NewReader(*p.conn), bufio.NewWriter(*p.conn))
	//n.peers.LoadOrStore(p.id, p)
	p.messageChan = n.messageChan

	p.wg.Add(1)
	//fmt.Println("InitClient id ", p.id)
	go n.HandlePackages(p)
	go n.SayHi(p)
}

// Client creates a connection to a peer specified by addr.
func (n *Client) CreatePeer(addr string) error {
	fmt.Println(n.identity.Address, " creates connection to peer ", addr)
	peer := &Peer{
		p2pnet: n,
	}
	fmt.Println(n.identity.Address, " dials to ", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	peer.conn = &conn
	n.InitPeerDataStructure(peer)
	return nil
}

// Client accepts a connection from a peer and initialize data structure
// correspondingly.
func (n *Client) AcceptConnectionFromPeer(c *net.Conn) {
	fmt.Println(n.identity.Address, " accepts connection from peer")
	peer := &Peer{
		p2pnet: n,
		conn:   c,
	}
	n.InitPeerDataStructure(peer)
}

func (n *Client) FindNodeById(id []byte) []dht.ID {
	targetId := dht.ID{
		Id: id,
	}
	return n.FindNode(targetId, dht.BucketSize, 8)
}

type lookupBucket struct {
	pending int
	queue   []dht.ID
}

// FindNode queries all peers this current node acknowledges for the closest peers
// to a specified target ID.
//
// All lookups are done under a number of disjoint lookups in parallel.
//
// Queries at most #ALPHA nodes at a time per lookup, and returns all peer IDs closest to a target peer ID.
func (n *Client) FindNode(targetID dht.ID, alpha int, disjointPaths int) (results []dht.ID) {
	visited := new(sync.Map)

	var lookups []*lookupBucket

	// Start searching for target from #ALPHA peers closest to target by queuing
	// them up and marking them as visited.
	for i, peerID := range n.routingTable.FindClosestPeers(targetID, alpha) {

		visited.Store(peerID.PublicKeyHex(), struct{}{})

		if len(lookups) < disjointPaths {
			lookups = append(lookups, new(lookupBucket))
		}

		lookup := lookups[i%disjointPaths]
		lookup.queue = append(lookup.queue, peerID)

		results = append(results, peerID)
	}

	wait, mutex := &sync.WaitGroup{}, &sync.Mutex{}

	for _, lookup := range lookups {
		go func(lookup *lookupBucket) {
			mutex.Lock()
			results = append(results, lookup.performLookup(n, targetID, alpha, visited)...)
			mutex.Unlock()

			wait.Done()
		}(lookup)

		wait.Add(1)
	}

	// Wait until all #D parallel lookups have been completed.
	wait.Wait()

	// Sort resulting peers by XOR distance.
	sort.Slice(results, func(i, j int) bool {
		left := results[i].Xor(targetID)
		right := results[j].Xor(targetID)
		return left.Less(right)
	})

	// Cut off list of results to only have the routing table focus on the
	// #dht.BucketSize closest peers to the current node.
	if len(results) > dht.BucketSize {
		results = results[:dht.BucketSize]
	}
	return
}

// TODO: What's this used for?
func (n *Client) FindMe(p *Peer, alpha int) (results []dht.ID) {
	lookup := lookupBucket{}
	lookup.queue = append(lookup.queue, p.identity)

	visited := new(sync.Map)
	visited.Store(p.identity.PublicKeyHex(), struct{}{})

	targetId := n.identity

	results = append(results, p.identity)
	results = append(results, lookup.performLookup(n, targetId, alpha, visited)...)

	// Sort resulting peers by XOR distance.
	sort.Slice(results, func(i, j int) bool {
		left := results[i].Xor(targetId)
		right := results[j].Xor(targetId)
		return left.Less(right)
	})

	// Cut off list of results to only have the routing table focus on the
	// #dht.BucketSize closest peers to the current node.
	if len(results) > dht.BucketSize {
		results = results[:dht.BucketSize]
	}

	for _, result := range results {
		n.routingTable.Update(result)
		fmt.Println("Update peer: ", result.Address)
	}
	return
}

func (lookup *lookupBucket) performLookup(n *Client, targetID dht.ID, alpha int, visited *sync.Map) (results []dht.ID) {
	responses := make(chan []*internal.ID)

	// Go through every peer in the entire queue and queue up what peers believe
	// is closest to a target ID.

	for ; lookup.pending < alpha && len(lookup.queue) > 0; lookup.pending++ {
		go n.queryPeerByID(lookup.queue[0], targetID, responses)

		lookup.queue = lookup.queue[1:]
	}

	// Empty queue.
	lookup.queue = lookup.queue[:0]

	// Asynchronous breadth-first search.
	for lookup.pending > 0 {
		response := <-responses

		lookup.pending--

		// Expand responses containing a peer's belief on the closest peers to target ID.
		for _, id := range response {
			peerID := dht.ID(*id)

			if _, seen := visited.LoadOrStore(peerID.PublicKeyHex(), struct{}{}); !seen && targetID.PublicKeyHex() != peerID.PublicKeyHex() {
				// Append new peer to be queued by the routing table.
				results = append(results, peerID)
				lookup.queue = append(lookup.queue, peerID)
			}
		}

		// Queue and request for #ALPHA closest peers to target ID from expanded results.
		for ; lookup.pending < alpha && len(lookup.queue) > 0; lookup.pending++ {
			go n.queryPeerByID(lookup.queue[0], targetID, responses)
			lookup.queue = lookup.queue[1:]
		}

		// Empty queue.
		lookup.queue = lookup.queue[:0]
	}

	return
}

func (n *Client) queryPeerByID(peerID dht.ID, targetID dht.ID, responses chan []*internal.ID) {
	_, loaded := n.peers.Load(string(peerID.Id))
	if !loaded {
		err := n.CreatePeer(peerID.Address)
		// TODO: Error Handling
		if err != nil {
			log.Fatal(err)
		}
	}
	value, _ := n.peers.Load(string(peerID.Id))
	p := value.(*Peer)

	p.wg.Wait()

	targetProtoID := internal.ID(targetID)

	request := new(Request)
	request.SetMessage(&internal.LookupNodeRequest{Target: &targetProtoID})
	request.SetTimeout(30 * time.Second)

	response, err := n.Request(p, request)

	if err != nil {
		responses <- []*internal.ID{}
		return
	}

	if response, ok := response.(*internal.LookupNodeResponse); ok {
		responses <- response.Peers
	} else {
		responses <- []*internal.ID{}
	}
}
