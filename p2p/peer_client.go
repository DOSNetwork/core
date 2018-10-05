package p2p

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/p2p/internal"
	"github.com/DOSNetwork/core/sign/bls"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	"github.com/dedis/kyber"
)

type PeerClient struct {
	p2pnet       *P2P
	conn         *net.Conn
	rw           *bufio.ReadWriter
	messageChan  chan P2PMessage
	status       int
	identity     dht.ID
	pubKey       kyber.Point
	wg           sync.WaitGroup
	RequestNonce uint64
	Requests     sync.Map
	mux          sync.Mutex
}

// RequestState represents a state of a request.
type RequestState struct {
	data        chan proto.Message
	closeSignal chan struct{}
}

func (p *PeerClient) Dial(addr string) {

	fmt.Println(p.p2pnet.identity.Address, "Dial", addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	p.conn = &conn
}

func (p *PeerClient) HandlePackages() {
	for {
		buf, err := p.receivePackage()
		switch {
		case err == io.EOF:
			fmt.Println("PeerClient ", p.identity.Id, " end")
			return
		case err != nil:
			fmt.Println("PeerClient ", p.identity.Id, " end")
			return
		}

		pa, ptr, err := p.decodePackage(buf)
		if err != nil {
			log.Fatal(err)
		}

		if pa.GetRequestNonce() > 0 && pa.GetReplyFlag() {
			if _state, exists := p.Requests.Load(pa.GetRequestNonce()); exists {
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
				p.p2pnet.peers.LoadOrStore(string(p.identity.Id), p)
				p.p2pnet.routingTable.Update(p.identity)
				fmt.Println(p.p2pnet.identity.Address, "Receive Hi id = ", p.identity.Id, "from", p.identity.Address)
				p.wg.Done()
			} else {
				fmt.Println("Ignore Hi")
			}
		case *internal.Ping:
			// Send pong to peer.
			err := p.Reply(pa.GetRequestNonce(), &internal.Pong{})
			if err != nil {
				log.Fatal(err)
			}
		case *internal.Pong:
			peers := p.p2pnet.findNode(p.identity, dht.BucketSize, 8)

			// Update routing table w/ closest peers to self.
			for _, peerID := range peers {
				p.p2pnet.routingTable.Update(peerID)
			}

		case *internal.LookupNodeRequest:
			// Prepare response.
			response := &internal.LookupNodeResponse{}

			// Respond back with closest peers to a provided target.
			for _, peerID := range p.p2pnet.routingTable.FindClosestPeers(dht.ID(*content.GetTarget()), dht.BucketSize) {
				id := internal.ID(peerID)
				response.Peers = append(response.Peers, &id)
			}

			err := p.Reply(pa.GetRequestNonce(), response)
			if err != nil {
				log.Fatal(err)
			}

		default:
			msg := P2PMessage{Msg: *ptr, Sender: p.identity.Id}
			p.messageChan <- msg
		}
		//fmt.Println("PeerClient ", p.id, " receive", string(buf))

	}
}

func (p *PeerClient) receivePackage() ([]byte, error) {
	var err error

	// Read until all header bytes have been read.
	buffer := make([]byte, 4)

	bytesRead, totalBytesRead := 0, 0
	c := *p.conn
	for totalBytesRead < 4 && err == nil {
		bytesRead, err = c.Read(buffer[totalBytesRead:])
		totalBytesRead += bytesRead
		if err != nil {
			return nil, err
		}
	}

	// Decode message size.
	size := binary.BigEndian.Uint32(buffer)
	if size == 0 {
		return nil, errors.New("received an empty message from a peer")
	}

	// Read until all message bytes have been read.
	buffer = make([]byte, size)

	bytesRead, totalBytesRead = 0, 0

	for totalBytesRead < int(size) && err == nil {
		bytesRead, err = c.Read(buffer[totalBytesRead:])
		totalBytesRead += bytesRead
	}
	return buffer, nil
}

func (p *PeerClient) decodePackage(bytes []byte) (*internal.Package, *ptypes.DynamicAny, error) {
	pa := new(internal.Package)
	if err := proto.Unmarshal(bytes, pa); err != nil {
		return nil, nil, err
	}

	//Todo verify pa.Signature by public key
	pub := suite.G2().Point()
	_ = pub.UnmarshalBinary(pa.GetPubkey())

	if err := bls.Verify(p.p2pnet.suite, pub, pa.GetAnything().Value, pa.GetSignature()); err != nil {
		return nil, nil, err
	}

	var ptr ptypes.DynamicAny
	if err := ptypes.UnmarshalAny(pa.GetAnything(), &ptr); err != nil {
		return nil, nil, err
	}
	return pa, &ptr, nil
}

func (p *PeerClient) SayHi() {
	fmt.Println(p.p2pnet.identity.Address, "say hi")
	pa := &internal.Hi{
		PublicKey: p.p2pnet.identity.PublicKey,
		Address:   p.p2pnet.identity.Address,
		Id:        p.p2pnet.identity.Id,
	}

	prepared, err := p.PrepareMessage(pa)
	if err != nil {
		log.Fatal(err)
	}

	prepared.RequestNonce = atomic.AddUint64(&p.RequestNonce, 1)
	if err := p.SendPackage(prepared); err != nil {
		log.Fatal(err)
	}

	p.wg.Wait()
}

func (p *PeerClient) SendPackage(msg proto.Message) error {
	if msg == nil {
		return errors.New("network: message is null")
	}
	//Encode the package
	bytes, _ := proto.Marshal(msg)
	// Serialize size.
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(bytes)))
	p.mux.Lock()
	defer p.mux.Unlock()
	bytes = append(prefix, bytes...)
	if _, err := p.rw.Write(bytes); err != nil {
		return err
	}
	if err := p.rw.Flush(); err != nil {
		return err
	}

	return nil
}

func (p *PeerClient) PrepareMessage(msg proto.Message) (*internal.Package, error) {
	if msg == nil {
		return nil, errors.New("network: message is null")
	}

	id := internal.ID(p.p2pnet.identity)
	anything, _ := ptypes.MarshalAny(msg)
	sig, _ := bls.Sign(p.p2pnet.suite, p.p2pnet.secKey, anything.Value)
	pub, _ := p.p2pnet.pubKey.MarshalBinary()

	pa := &internal.Package{
		Sender:    &id,
		Anything:  anything,
		Pubkey:    pub,
		Signature: sig,
	}

	return pa, nil
}

// Request requests for a response for a request sent to a given peer.
func (p *PeerClient) Request(req *Request) (proto.Message, error) {
	prepared, err := p.PrepareMessage(req.Message)
	if err != nil {
		return nil, err
	}

	prepared.RequestNonce = atomic.AddUint64(&p.RequestNonce, 1)
	if err := p.SendPackage(prepared); err != nil {
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
func (p *PeerClient) Reply(nonce uint64, message proto.Message) error {
	prepared, err := p.PrepareMessage(message)
	if err != nil {
		return err
	}

	// Set the nonce.
	prepared.RequestNonce = nonce
	prepared.ReplyFlag = true

	if err := p.SendPackage(prepared); err != nil {
		return err
	}

	return nil
}

func (p *PeerClient) FindMe(alpha int) (results []dht.ID) {
	lookup := lookupBucket{}
	lookup.queue = append(lookup.queue, p.identity)

	visited := new(sync.Map)
	visited.Store(p.identity.PublicKeyHex(), struct{}{})

	targetId := p.p2pnet.identity

	results = append(results, p.identity)
	results = append(results, lookup.performLookup(p.p2pnet, targetId, alpha, visited)...)

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
		p.p2pnet.routingTable.Update(result)
		fmt.Println("Update peer: ", result.Address)
	}
	return
}
