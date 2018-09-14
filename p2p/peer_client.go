package p2p

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/DOSNetwork/core/p2p/internal"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	"github.com/DOSNetwork/core/sign/bls"
	"github.com/dedis/kyber"
)

type PeerClient struct {
	p2pnet      *P2P
	conn        *net.Conn
	rw          *bufio.ReadWriter
	messageChan chan P2PMessage
	status      int
	id          string
	ip          string
	pubKey      kyber.Point
}

func (p *PeerClient) Dial(addr string) {

	log.Println("Dial " + addr)
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
			fmt.Println("PeerClient ", p.id, " end")
			return
		case err != nil:
			fmt.Println("PeerClient ", p.id, " end")
			return
		}
		ptr := p.decodePackage(buf)
		switch content := ptr.Message.(type) {
		case *internal.Hi:
			if p.id == "" {
				p.id = content.Id
				p.ip = content.Ip
				pub := suite.G2().Point()
				_ = pub.UnmarshalBinary(content.Pubkey)
				p.pubKey = pub
				p.p2pnet.peers.LoadOrStore(p.id, p)
				fmt.Println("Receive Hi id = ", p.id, " ip = ", p.id)
			} else {
				fmt.Println("Ignore Hi")
			}

		default:
			msg := P2PMessage{Msg: ptr, Sender: ""}
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

func (p *PeerClient) decodePackage(bytes []byte) ptypes.DynamicAny {
	pa := new(internal.Package)
	_ = proto.Unmarshal(bytes, pa)

	//Todo verify pa.Signature by public key
	pub := suite.G2().Point()
	_ = pub.UnmarshalBinary(pa.Pubkey)

	err := bls.Verify(p.p2pnet.suite, pub, pa.Anything.Value, pa.Signature)
	if err == nil {
		fmt.Println("bls verify okokok")
	}

	var ptr ptypes.DynamicAny
	_ = ptypes.UnmarshalAny(pa.Anything, &ptr)
	return ptr
}

func (p *PeerClient) SayHi() {
	fmt.Println("say hi")
	pub, _ := p.p2pnet.pubKey.MarshalBinary()
	pa := &internal.Hi{
		Id:     p.p2pnet.id,
		Ip:     p.p2pnet.ip,
		Pubkey: pub,
	}
	msg := proto.Message(pa)
	p.SendPackage(&msg)
}

func (p *PeerClient) SendPackage(msg *proto.Message) error {
	var bytes []byte
	pub, _ := p.p2pnet.pubKey.MarshalBinary()
	anything, _ := ptypes.MarshalAny(*msg)
	sig, _ := bls.Sign(p.p2pnet.suite, p.p2pnet.secKey, anything.Value)

	pa := &internal.Package{
		Anything:  anything,
		Pubkey:    pub,
		Signature: sig,
	}

	//Encode the package
	bytes, _ = proto.Marshal(pa)
	// Serialize size.
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(bytes)))

	bytes = append(prefix, bytes...)
	p.rw.Write(bytes)
	p.rw.Flush()
	return nil
}
