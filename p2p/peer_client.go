package p2p

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type PeerClient struct {
	conn        *net.Conn
	rw          *bufio.ReadWriter
	MessageChan chan []byte
	id          string
}

func (p *PeerClient) Dial(addr string) {
	// Dial the remote process.
	// Note that the local port is chosen on the fly. If the local port
	// must be a specific one, use DialTCP() instead.
	log.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	p.conn = &conn
}

func (p *PeerClient) HandleMessages() {
	for {
		buf, err := p.receiveMessage()
		fmt.Println("receive ", string(buf))
		switch {
		case err == io.EOF:
			return
		case err != nil:
			return
		}
		p.MessageChan <- buf
	}
}

func (p *PeerClient) receiveMessage() ([]byte, error) {
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
	fmt.Println("receiveMessage : size ", size)
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

func (p *PeerClient) SendMessage(buffer []byte) error {
	// Serialize size.
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(buffer)))

	buffer = append(prefix, buffer...)
	p.rw.Write(buffer)
	p.rw.Flush()
	return nil
}
