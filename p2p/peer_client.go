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
)

type PeerClient struct {
	conn        *net.Conn
	rw          *bufio.ReadWriter
	messageChan chan P2PMessage
	id          string
	status      int
	//private key
	//public key
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
		//fmt.Println("PeerClient ", p.id, " receive", string(buf))
		p.messageChan <- p.decodePackage(buf)
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

func (p *PeerClient) decodePackage(bytes []byte) P2PMessage {
	pa := new(internal.Package)
	_ = proto.Unmarshal(bytes, pa)

	//Todo verify pa.Signature by public key

	var ptr ptypes.DynamicAny
	_ = ptypes.UnmarshalAny(pa.Anything, &ptr)

	return P2PMessage{Msg: ptr, Sender: p.id}
}

func (p *PeerClient) SendPackage(msg *proto.Message) error {
	anything, _ := ptypes.MarshalAny(*msg)
	pa := &internal.Package{
		Anything: anything,
	}
	//sign the package
	//pa.Signature =

	//Encode the package
	bytes, _ := proto.Marshal(pa)
	// Serialize size.
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(bytes)))

	bytes = append(prefix, bytes...)
	p.rw.Write(bytes)
	p.rw.Flush()
	return nil
}
