package p2p

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/p2p/internal"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"

	"crypto/aes"
	"crypto/cipher"

	"github.com/dedis/kyber"
)

const (
	HEARTBEATINTERVAL = 60 //In seconds
	HEARTBEATMAXWAIT  = 10 //In seconds
	HEARTBEATMAXCOUNT = 6
	CONNIDLETIMEOUT   = 5
	HITIMEOUT         = 60 //In seconds
)

type PeerConn struct {
	p2pnet          *P2P
	rxMessage       chan P2PMessage
	pubKey          kyber.Point
	conn            *net.Conn
	rw              *bufio.ReadWriter
	waitForHi       chan bool
	done            chan bool
	identity        internal.ID
	RequestNonce    uint64
	Requests        sync.Map
	mux             sync.Mutex
	readWriteCount  uint64
	idelPeriodCount uint8
	lastusedtime    time.Time
	logger          log.Logger
	dhkey           []byte
	nonce           []byte
}

// RequestState represents a state of a request.
type RequestState struct {
	data        chan proto.Message
	closeSignal chan struct{}
}

func NewPeerConn(p2pnet *P2P, conn *net.Conn, rxMessage chan P2PMessage) (peer *PeerConn, err error) {
	peer = &PeerConn{
		p2pnet:       p2pnet,
		conn:         conn,
		rxMessage:    rxMessage,
		waitForHi:    make(chan bool, 2),
		done:         make(chan bool, 1),
		rw:           bufio.NewReadWriter(bufio.NewReader(*conn), bufio.NewWriter(*conn)),
		lastusedtime: time.Now(),
		logger:       p2pnet.logger,
	}
	if err = peer.Start(); err != nil {
		p2pnet.logger.Error(err)
		return
	}

	go peer.heartBeat()
	return
}

func (p *PeerConn) Start() (err error) {
	go p.receiveLoop()
	return
}

func (p *PeerConn) SendMessage(msg proto.Message) (err error) {
	p.lastusedtime = time.Now()
	var prepared *internal.Package
	if prepared, err = p.prepareMessage(msg); err != nil {
		p.logger.Error(err)
		return
	}

	prepared.RequestNonce = atomic.AddUint64(&p.RequestNonce, 1)
	if err = p.SendPackage(prepared); err != nil {
		return
	}

	atomic.AddUint64(&p.readWriteCount, 1)
	return
}

func (p *PeerConn) receiveLoop() {
	var err error
	var buf []byte
	var pa *internal.Package
	var ptr *ptypes.DynamicAny

	for {
		if buf, err = p.receivePackage(); err != nil {
			if err != io.EOF {
				p.logger.Error(err)
			}
			break
		}

		if pa, ptr, err = p.decodePackage(buf); err != nil {
			p.logger.Error(err)
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
					p.logger.Error(err)
				}
				p.pubKey = pub
				p.logger.Debug("Conn Established")
				p.waitForHi <- true
				response := &internal.Hi{
					PublicKey: p.p2pnet.identity.PublicKey,
					Address:   p.p2pnet.identity.Address,
					Id:        p.p2pnet.identity.Id,
				}
				if err := p.Reply(pa.GetRequestNonce(), response); err != nil {
					p.logger.Error(err)
				}
				if err := p.getShareKeyAndNonce(); err != nil {
					p.logger.Error(err)
				}
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

			if err := p.Reply(pa.GetRequestNonce(), response); err != nil {
				p.logger.Error(err)
			}
		case *internal.Ping:
			response := &internal.Pong{}
			if err := p.Reply(pa.GetRequestNonce(), response); err != nil {
				p.logger.Error(err)
			}
		case *internal.Pong:
		default:
			p.lastusedtime = time.Now()
			msg := P2PMessage{Msg: *ptr, Sender: p.identity.Id}
			go func() { p.rxMessage <- msg }()
			response := &internal.Pong{}
			if err := p.Reply(pa.GetRequestNonce(), response); err != nil {
				p.logger.Error(err)
			}
		}
	}
	p.p2pnet.peers.DeletePeer(string(p.identity.Id))

	if err := (*p.conn).Close(); err != nil {
		p.logger.Error(err)
	}
	close(p.waitForHi)
	p.logger.Debug("Conn End")
}

func (p *PeerConn) End() {
	if p.conn == nil {
		return
	}
	if err := (*p.conn).Close(); err != nil {
		p.logger.Error(err)
		p.p2pnet.peers.DeletePeer(string(p.identity.Id))
	}
}

func (p *PeerConn) receivePackage() ([]byte, error) {
	var err error

	// Read until all header bytes have been read.
	buffer := make([]byte, 4)

	bytesRead, totalBytesRead := 0, 0
	c := *p.conn
	for totalBytesRead < 4 && err == nil {
		if bytesRead, err = c.Read(buffer[totalBytesRead:]); err != nil {
			return nil, err
		}
		totalBytesRead += bytesRead
	}

	// Decode message size.
	size := binary.BigEndian.Uint32(buffer)
	if size == 0 {
		err := errors.New("received an empty message from a peer")
		return nil, err
	}

	// Read until all message bytes have been read.
	buffer = make([]byte, size)

	bytesRead, totalBytesRead = 0, 0

	for totalBytesRead < int(size) && err == nil {
		if bytesRead, err = c.Read(buffer[totalBytesRead:]); err != nil {
			return nil, err
		}
		totalBytesRead += bytesRead
	}

	return buffer, nil
}

func (p *PeerConn) decodePackage(bytes []byte) (*internal.Package, *ptypes.DynamicAny, error) {
	var content []byte
	var err error
	if content, err = p.decrypt(bytes); err != nil {
		p.logger.Error(err)
		return nil, nil, err
	}
	pa := new(internal.Package)
	if err = proto.Unmarshal(content, pa); err != nil {
		p.logger.Error(err)
		return nil, nil, err
	}

	pub := suite.G2().Point()
	if err = pub.UnmarshalBinary(pa.GetPubkey()); err != nil {
		p.logger.Error(err)
		return nil, nil, err
	}

	if err = bls.Verify(p.p2pnet.suite, pub, pa.GetAnything().Value, pa.GetSignature()); err != nil {
		p.logger.Error(err)
		return nil, nil, err
	}

	var ptr ptypes.DynamicAny
	if err = ptypes.UnmarshalAny(pa.GetAnything(), &ptr); err != nil {
		p.logger.Error(err)
		return nil, nil, err
	}

	switch ptr.Message.(type) {
	case *internal.Ping:
	case *internal.Pong:
	default:
		atomic.AddUint64(&p.readWriteCount, 1)
	}

	return pa, &ptr, nil
}

func (p *PeerConn) SayHi() (err error) {
	var response proto.Message
	var content *internal.Hi
	var ok bool

	request := new(Request)
	request.SetMessage(&internal.Hi{
		PublicKey: p.p2pnet.identity.PublicKey,
		Address:   p.p2pnet.identity.Address,
		Id:        p.p2pnet.identity.Id,
	})
	request.SetTimeout(HITIMEOUT * time.Second)

	if response, err = p.Request(request); err != nil {
		return err
	}

	if content, ok = response.(*internal.Hi); !ok {
		err := errors.New("Not a Hi message")
		return err
	}

	if len(p.identity.Id) == 0 {
		p.identity.Id = content.GetId()
		p.identity.Address = content.GetAddress()
		p.identity.PublicKey = content.GetPublicKey()
		pub := suite.G2().Point()
		if err = pub.UnmarshalBinary(content.GetPublicKey()); err != nil {
			return err
		}
		p.pubKey = pub
		if err = p.getShareKeyAndNonce(); err != nil {
			return err
		}
		p.logger.Debug("Conn Established")

	}
	return nil
}

func (p *PeerConn) getShareKeyAndNonce() (err error) {
	var dhBytes []byte
	dhKey := suite.Point().Mul(p.p2pnet.secKey, p.pubKey)
	if dhBytes, err = dhKey.MarshalBinary(); err != nil {
		return
	}
	p.dhkey = dhBytes[0:32]
	p.nonce = dhBytes[32:44]
	return
}

func (p *PeerConn) encrypt(plaintext []byte) (c []byte, err error) {
	var block cipher.Block
	var aesgcm cipher.AEAD
	c = plaintext
	if len(p.dhkey) == 0 {
		return
	}

	if block, err = aes.NewCipher(p.dhkey); err != nil {
		p.logger.Error(err)
		return
	}

	if aesgcm, err = cipher.NewGCM(block); err != nil {
		p.logger.Error(err)
		return
	}
	c = aesgcm.Seal(nil, p.nonce, plaintext, nil)
	return
}

func (p *PeerConn) decrypt(ciphertext []byte) (c []byte, err error) {
	var block cipher.Block
	var aesgcm cipher.AEAD
	c = ciphertext
	if len(p.dhkey) == 0 {
		return
	}
	if block, err = aes.NewCipher(p.dhkey); err != nil {
		p.logger.Error(err)
		return
	}

	if aesgcm, err = cipher.NewGCM(block); err != nil {
		p.logger.Error(err)
		return
	}

	if c, err = aesgcm.Open(nil, p.nonce, ciphertext, nil); err != nil {
		p.logger.Error(err)
	}
	return
}

//Add a timer to avoid wait for Hi forever
func (p *PeerConn) HeardHi() (err error) {
	timer := time.NewTimer(HITIMEOUT * time.Second)

	select {
	case <-timer.C:
		err = errors.New("Info Exchange Timeout")
	case <-p.waitForHi:
		timer.Stop()
	}

	return
}

func (p *PeerConn) SendPackage(msg proto.Message) (err error) {
	var bytes []byte
	if msg == nil {
		err := errors.New("Message is nil")
		p.logger.Error(err)
		return err
	}
	//Encode the package
	if bytes, err = proto.Marshal(msg); err != nil {
		p.logger.Error(err)
		return err
	}
	// Serialize size.
	if bytes, err = p.encrypt(bytes); err != nil {
		p.logger.Error(err)
		return err
	}
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(bytes)))
	bytes = append(prefix, bytes...)
	p.mux.Lock()
	defer p.mux.Unlock()

	if _, err := p.rw.Write(bytes); err != nil {
		p.logger.Error(err)
		return err
	}
	if err := p.rw.Flush(); err != nil {
		p.logger.Error(err)
		return err
	}

	return nil
}

func (p *PeerConn) prepareMessage(msg proto.Message) (pa *internal.Package, err error) {
	var anything *any.Any
	var sig, pub []byte
	if msg == nil {
		err = errors.New("network: message is null")
		p.logger.Error(err)
		return
	}

	id := internal.ID(p.p2pnet.identity)
	if anything, err = ptypes.MarshalAny(msg); err != nil {
		p.logger.Error(err)
		return
	}
	//TODO:change to AES256-GCM
	if sig, err = bls.Sign(p.p2pnet.suite, p.p2pnet.secKey, anything.Value); err != nil {
		p.logger.Error(err)
		return
	}
	if pub, err = p.p2pnet.pubKey.MarshalBinary(); err != nil {
		p.logger.Error(err)
		return
	}

	pa = &internal.Package{
		Sender:    &id,
		Anything:  anything,
		Pubkey:    pub,
		Signature: sig,
	}

	return
}

// Request requests for a response for a request sent to a given peer.
func (p *PeerConn) Request(req *Request) (proto.Message, error) {
	prepared, err := p.prepareMessage(req.Message)
	if err != nil {
		return nil, err
	}

	prepared.RequestNonce = atomic.AddUint64(&p.RequestNonce, 1)

	switch req.Message.(type) {
	case *internal.Ping:
	default:
		atomic.AddUint64(&p.readWriteCount, 1)
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

	if err := p.SendPackage(prepared); err != nil {
		return nil, err
	}

	select {
	case res := <-channel:
		return res, nil
	case <-time.After(req.Timeout):
		err := errors.New("request timed out")
		p.logger.Error(err)
		return nil, err
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
		return err
	}

	switch message.(type) {
	case *internal.Pong:
	default:
		atomic.AddUint64(&p.readWriteCount, 1)
	}

	return nil
}

func (p *PeerConn) heartBeat() {
	ticker := time.NewTicker(HEARTBEATINTERVAL * time.Second)
	for {
		select {
		case <-ticker.C:
			if p.readWriteCount == 0 {
				p.idelPeriodCount++
			} else {
				p.idelPeriodCount = 0
			}

			if p.idelPeriodCount == CONNIDLETIMEOUT {
				if err := (*p.conn).Close(); err != nil {
					p.logger.Error(err)
				}
				return
			}

			timeout := false
			for i := 0; i < HEARTBEATMAXCOUNT; i++ {
				request := new(Request)
				request.SetMessage(&internal.Ping{})
				request.SetTimeout(HEARTBEATMAXWAIT * time.Second)
				if _, err := p.Request(request); err != nil {
					p.logger.Error(err)
					timeout = true
				} else {
					timeout = false
					break
				}
			}
			if timeout {
				if err := (*p.conn).Close(); err != nil {
					p.logger.Error(err)
				}
				return
			} else {
				atomic.StoreUint64(&p.readWriteCount, 0)
			}
		}
	}
}
