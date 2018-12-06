package p2p

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
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


type PeerConn struct {
	//TODO:remove *P2P and use sync map to manage the lifecycle of PeerConn
	p2pnet *P2P
	//TODO:
	rxMessage chan P2PMessage
	//TODO:
	pubKey    kyber.Point
	conn      *net.Conn
	rw        *bufio.ReadWriter
	txMessage chan proto.Message
	//TODO:Need to be refactored
	waitForHi    chan bool
	done         chan bool
	status       int
	identity     internal.ID
	RequestNonce uint64
	Requests     sync.Map
	mux          sync.Mutex
}

// RequestState represents a state of a request.
type RequestState struct {
	data        chan proto.Message
	closeSignal chan struct{}
}

func NewPeerConn(p2pnet *P2P, conn *net.Conn, rxMessage chan P2PMessage) (peer *PeerConn, err error) {
	peer = &PeerConn{
		p2pnet:    p2pnet,
		conn:      conn,
		rxMessage: rxMessage,
		txMessage: make(chan proto.Message, 100),
		waitForHi: make(chan bool, 2),
		done:      make(chan bool, 2),
		rw:        bufio.NewReadWriter(bufio.NewReader(*conn), bufio.NewWriter(*conn)),
	}
	err = peer.Start()
	if err != nil {
		peer = nil
	}
	return
}

func (p *PeerConn) Start() (err error) {
	go p.receiveLoop()
	err = p.SayHi()
	if err != nil {
		close(p.txMessage)
	}
	return
}

func (p *PeerConn) SendMessage(msg proto.Message) (err error) {
	var prepared *internal.Package
	prepared, err = p.prepareMessage(msg)
	if err != nil {
		fmt.Println("PeerConn sendLoop err ", err)
		return
	}

	prepared.RequestNonce = atomic.AddUint64(&p.RequestNonce, 1)
	if err = p.SendPackage(prepared); err != nil {
		fmt.Println("PeerConn sendLoop SendPackage ", err)
	}
	return
}

func (p *PeerConn) receiveLoop() {
	//time.Sleep(TIMEOUTFORHI * time.Second)
L:
	for {
		select {
		case <-p.done:
			break L
		default:
			buf, err := p.receivePackage()
			switch {
			case err == io.EOF:
				fmt.Println("PeerConn ", p.identity.Id, " EOF")
				break L
			case err != nil:
				fmt.Println("PeerConn ", p.identity.Id, " err ", err)
				break L
			}

			pa, ptr, err := p.decodePackage(buf)
			if err != nil {
				fmt.Println("PeerConn decodePackage err ", err)
				continue
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
			//TODO:Refactor this to use request/reply and move out of this loop
			case *internal.Hi:
				if len(p.identity.Id) == 0 {
					p.identity.Id = content.GetId()
					p.identity.Address = content.GetAddress()
					p.identity.PublicKey = content.GetPublicKey()
					pub := suite.G2().Point()
					if err = pub.UnmarshalBinary(content.GetPublicKey()); err != nil {
						fmt.Println("PeerConn UnmarshalBinary err ", err)
					}
					p.pubKey = pub
					p.waitForHi <- true
				} else {
					fmt.Println("Ignore Hi")
				}
				//TODO:move this to routing
			case *internal.LookupNodeRequest:
				// Prepare response.
				response := &internal.LookupNodeResponse{}

				// Respond back with closest peers to a provided target.
				for _, peerID := range p.p2pnet.routingTable.FindClosestPeers(internal.ID(*content.GetTarget()), dht.BucketSize) {
					id := internal.ID(peerID)
					response.Peers = append(response.Peers, &id)
				}

				err := p.Reply(pa.GetRequestNonce(), response)
				if err != nil {
					log.Fatal(err)
				}

			default:
				//TODO
				msg := P2PMessage{Msg: *ptr, Sender: p.identity.Id}
				p.rxMessage <- msg
			}
		}
	}
	_, ok := p.p2pnet.peers.Load(string(p.identity.Id))
	if ok {
		p.p2pnet.peers.Delete(string(p.identity.Id))
	}
	//(*p.conn).Close()
	close(p.done)
	close(p.waitForHi)
	fmt.Println("PeerConn receiveLoop done")
}

func (p *PeerConn) receivePackage() ([]byte, error) {
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

func (p *PeerConn) decodePackage(bytes []byte) (*internal.Package, *ptypes.DynamicAny, error) {
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

func (p *PeerConn) SayHi() (err error) {
	pa := &internal.Hi{
		PublicKey: p.p2pnet.identity.PublicKey,
		Address:   p.p2pnet.identity.Address,
		Id:        p.p2pnet.identity.Id,
	}

	err = p.SendMessage(pa)

	//Add a timer to avoid wait for Hi forever
	timer := time.NewTimer(60 * time.Second)
L:
	for {
		select {
		case <-timer.C:
			p.done <- true
			fmt.Println("Time expire")
			err = errors.New("PeerConn: Time expire")
			break L
		case <-p.waitForHi:
			_ = timer.Stop()
			break L
		}
	}

	return
}

func (p *PeerConn) SendPackage(msg proto.Message) error {
	if msg == nil {
		return errors.New("network: message is null")
	}
	//Encode the package
	bytes, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("SendPackage Marshal err ", err)
	}
	// Serialize size.
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(bytes)))
	p.mux.Lock()
	defer p.mux.Unlock()
	bytes = append(prefix, bytes...)

	if _, err := p.rw.Write(bytes); err != nil {
		fmt.Println("SendPackage Write err ", err)
		return err
	}
	if err := p.rw.Flush(); err != nil {
		fmt.Println("!!!!SendPackage Flush err ", err)
		return err
	}

	return nil
}

func (p *PeerConn) prepareMessage(msg proto.Message) (*internal.Package, error) {
	var err error
	if msg == nil {
		return nil, errors.New("network: message is null")
	}

	id := internal.ID(p.p2pnet.identity)
	anything, err := ptypes.MarshalAny(msg)
	if err != nil {
		fmt.Println("ptypes.MarshalAny ", err)
		return nil, err
	}
	//TODO:change to AES256-GCM
	sig, err := bls.Sign(p.p2pnet.suite, p.p2pnet.secKey, anything.Value)
	if err != nil {
		fmt.Println("prepareMessage ", err)
	}
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
func (p *PeerConn) Request(req *Request) (proto.Message, error) {
	prepared, err := p.prepareMessage(req.Message)
	if err != nil {
		return nil, err
	}

	prepared.RequestNonce = atomic.AddUint64(&p.RequestNonce, 1)
	if err := p.SendPackage(prepared); err != nil {
		fmt.Println("PeerConn Request err ", err)
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
func (p *PeerConn) Reply(nonce uint64, message proto.Message) error {
	prepared, err := p.prepareMessage(message)
	if err != nil {
		return err
	}

	// Set the nonce.
	prepared.RequestNonce = nonce
	prepared.ReplyFlag = true

	if err := p.SendPackage(prepared); err != nil {
		fmt.Println("PeerConn Reply err ", err)
		return err
	}

	return nil
}