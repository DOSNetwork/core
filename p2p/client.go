package p2p

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
	//	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/suites"

	"github.com/dedis/kyber"

	"crypto/aes"
	"crypto/cipher"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

const (
	MAXMESSAGESIZE = 1024 * 1024
	BUFFERSIZE     = 1024
	HEADERSIZE     = 4
)

type Client struct {
	conn         net.Conn
	incomingConn bool

	suite        suites.Suite
	localID      []byte
	localSecKey  kyber.Scalar
	localPubKey  kyber.Point
	remoteID     []byte
	remotePubKey kyber.Point
	dhKey        []byte
	dhNonce      []byte

	ctx      context.Context
	cancel   context.CancelFunc
	sender   chan Request
	receiver chan interface{}
	errc     chan error
}

func MergeErrors(ctx context.Context, cs ...<-chan error) chan error {
	var wg sync.WaitGroup
	// We must ensure that the output channel has the capacity to
	// hold as many errors
	// as there are error channels.
	// This will ensure that it never blocks, even
	// if WaitForPipeline returns early.
	out := make(chan error, len(cs))
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls
	// wg.Done.
	output := func(c <-chan error) {
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
	// Start a goroutine to close out once all the output goroutines
	// are done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func NewClient(suite suites.Suite, secKey kyber.Scalar, localPubKey kyber.Point, localID []byte, conn net.Conn, incomingConn bool) (client *Client, err error) {
	client = &Client{
		suite:        suite,
		localSecKey:  secKey,
		localPubKey:  localPubKey,
		localID:      localID,
		conn:         conn,
		incomingConn: incomingConn,
	}
	client.ctx, client.cancel = context.WithCancel(context.Background())
	client.sender = make(chan Request, 21)
	//Wait for exchanging ID complete
	err = client.exchangeID()
	if client.remoteID == nil {
		err = errors.New("exchangeID failed")
		return
	}
	if err != nil {
		logger.Error(err)
		return
	}

	var errs []<-chan error
	var packByte chan []byte
	//Build a secure pipeline
	bytes, errc := readBytes(client.ctx, conn, client.localID, client.remoteID, incomingConn)
	errs = append(errs, errc)
	decrypted, errc := decrypt(client.ctx, client.dhKey, client.dhNonce, bytes)
	errs = append(errs, errc)
	client.receiver, packByte, errc = dispatch(client.localID, client.remoteID, incomingConn, client.ctx, client.suite, client.remotePubKey, localID, client.sender, decrypted)
	errs = append(errs, errc)
	encrypted, errc := encrypt(client.ctx, client.dhKey, client.dhNonce, packByte)
	errs = append(errs, errc)
	errc = sendBytes(client.ctx, client.conn, encrypted, client.localID, client.remoteID, incomingConn)
	errs = append(errs, errc)
	client.errc = MergeErrors(client.ctx, errs...)
	return
}

func (c *Client) Close() {
	c.conn.Close()
	c.cancel()

}
func (c *Client) ErrorHandling(errc <-chan error) {
	go func() {
		for err := range errc {
			if err.Error() != "cipher: message authentication failed" {
			}
			if err.Error() == "EOF" {
				c.Close()
			}
		}
	}()
}

func (c *Client) exchangeID() (err error) {
	var wg sync.WaitGroup
	ch := make(chan struct{})
	wg.Add(2)
	go c.sendID(&wg)
	go c.receiveID(&wg)
	go func() {
		wg.Wait()
		ch <- struct{}{}
	}()
	timeout := time.Duration(10) * time.Second
	select {
	case <-ch:
		f := map[string]interface{}{
			"localID":  c.localID,
			"remoteID": c.remoteID}
		if logger != nil {
			logger.Event("exchangeIDSuccess", f)
		}
	case <-time.After(timeout):
		f := map[string]interface{}{
			"localID": c.localID}
		logger.Event("exchangeIDFail", f)
	}
	return
}

func (c *Client) receiveID(wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	// Read until all header bytes have been read.
	header := make([]byte, HEADERSIZE)
	// Read until all header bytes have been read.
	bytesRead, totalBytesRead := 0, 0
	for totalBytesRead < HEADERSIZE && err == nil {
		if bytesRead, err = c.conn.Read(header[totalBytesRead:]); err != nil {
			if err.Error() == "EOF" {
				c.Close()
			}
			return
		}
		totalBytesRead += bytesRead
	}

	// Decode message size.
	size := binary.BigEndian.Uint32(header)

	header = nil
	if size > MAXMESSAGESIZE {
		err = errors.New("p2p message size is too big " + strconv.Itoa(int(size)))
		return
	}
	// Read until all message bytes have been read.
	buffer := make([]byte, size)
	if size == 0 {
		return
	}
	bytesRead, totalBytesRead = 0, 0

	for totalBytesRead < int(size) && err == nil {
		if bytesRead, err = c.conn.Read(buffer[totalBytesRead:]); err != nil {
			return
		}
		totalBytesRead += bytesRead
	}
	pa := new(Package)
	if err = proto.Unmarshal(buffer, pa); err != nil {
		logger.Error(err)
		return
	}
	buffer = nil

	var ptr ptypes.DynamicAny
	if err = ptypes.UnmarshalAny(pa.GetAnything(), &ptr); err != nil {
		logger.Error(err)
		return
	}
	id, ok := ptr.Message.(*ID)
	if !ok {
		return
	}

	c.remoteID = id.GetId()

	if string(c.remoteID) == string(c.localID) {
		os.Exit(2)
	}
	pub := c.suite.G2().Point()
	if err = pub.UnmarshalBinary(id.GetPublicKey()); err != nil {
		logger.Error(err)
		return
	}
	c.remotePubKey = pub

	var dhBytes []byte
	dhKey := c.suite.Point().Mul(c.localSecKey, c.remotePubKey)
	if dhBytes, err = dhKey.MarshalBinary(); err != nil {
		logger.Error(err)
		return
	}
	c.dhKey = dhBytes[0:32]
	c.dhNonce = dhBytes[32:44]
	return
}

func (c *Client) sendID(wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	var anything *any.Any
	var bytes, pubKeyBytes []byte
	if pubKeyBytes, err = c.localPubKey.MarshalBinary(); err != nil {
		logger.Error(err)
		return
	}

	pID := &ID{
		PublicKey: pubKeyBytes,
		Id:        c.localID,
	}

	if anything, err = ptypes.MarshalAny(pID); err != nil {
		logger.Error(err)
		return
	}
	pa := &Package{
		Anything: anything,
	}
	if bytes, err = proto.Marshal(pa); err != nil {
		logger.Error(err)
		return
	}
	if len(bytes) > MAXMESSAGESIZE {
		err = errors.New("p2p message size is too big " + strconv.Itoa(int(len(bytes))))
		return
	}
	prefix := make([]byte, HEADERSIZE)
	binary.BigEndian.PutUint32(prefix, uint32(len(bytes)))

	bytes = append(prefix, bytes...)

	bytesRead, totalBytesRead := 0, 0
	for totalBytesRead < HEADERSIZE && err == nil {
		if bytesRead, err = c.conn.Write(bytes); err != nil {
			c.Close()
			return
		}
		totalBytesRead += bytesRead
	}

	return
}

func encrypt(ctx context.Context, dhKey, dhNonce []byte, plaintext chan []byte) (chan []byte, chan error) {
	out := make(chan []byte)
	errc := make(chan error)

	go func() {
		defer close(out)
		defer close(errc)
		for {
			select {
			case c, ok := <-plaintext:
				if !ok {
					return
				}

				var err error
				var block cipher.Block
				var aesgcm cipher.AEAD

				if block, err = aes.NewCipher(dhKey); err != nil {
					select {
					case errc <- err:
					case <-ctx.Done():
					}
					continue
				}

				if aesgcm, err = cipher.NewGCM(block); err != nil {
					select {
					case errc <- err:
					case <-ctx.Done():
					}
					continue
				}
				c = aesgcm.Seal(nil, dhNonce, c, nil)

				select {
				case out <- c:
				case <-ctx.Done():
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, errc
}

func decrypt(ctx context.Context, dhKey, dhNonce []byte, ciphertext chan []byte) (chan []byte, chan error) {
	out := make(chan []byte)
	errc := make(chan error)

	go func() {
		defer close(out)
		defer close(errc)

		for {
			select {
			case c, ok := <-ciphertext:
				if !ok {
					return
				}

				var block cipher.Block
				var aesgcm cipher.AEAD
				var err error

				if block, err = aes.NewCipher(dhKey); err != nil {
					select {
					case errc <- err:
					case <-ctx.Done():
					}
					continue
				}

				if aesgcm, err = cipher.NewGCM(block); err != nil {
					select {
					case errc <- err:
					case <-ctx.Done():
					}
					continue
				}

				if c, err = aesgcm.Open(nil, dhNonce, c, nil); err != nil {
					select {
					case errc <- err:
					case <-ctx.Done():
					}
					continue
				}

				select {
				case out <- c:
				case <-ctx.Done():
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, errc
}

func dispatch(localID, remoteID []byte, incomingConn bool, ctx context.Context, suite suites.Suite, pub kyber.Point, Id []byte, sendMsg chan Request, receiveBytes chan []byte) (chan interface{}, chan []byte, chan error) {
	receiver := make(chan interface{}, 21)
	sender := make(chan []byte, 21)
	errc := make(chan error)
	requests := make(map[uint64]Request)
	var nonce uint64
	go func() {
		defer close(receiver)
		defer close(sender)
		defer close(errc)

		for {
			select {
			case req, ok := <-sendMsg:
				if !ok {
					return
				}

				if !req.p.ReplyFlag {
					if req.ctx == nil && req.reply == nil {
						continue
					}
					req.p.RequestNonce = nonce
				}
				//Encode the package
				var bytes []byte
				var err error
				if bytes, err = proto.Marshal(req.p); err != nil {
					go func() {
						select {
						case req.errc <- err:
						case <-req.ctx.Done():
						}
						close(req.errc)
						if req.reply != nil {
							close(req.reply)
						}
					}()
					continue
				}
				if !req.p.ReplyFlag {
					requests[nonce] = req
					nonce++
				}
				select {
				case sender <- bytes:
					if !req.p.ReplyFlag {
					} else {
						//req.cancel()
					}
				case <-req.ctx.Done():
					close(req.errc)
					if req.reply != nil {
						close(req.reply)
					}
				}
			case bytes, ok := <-receiveBytes:
				if !ok {
					return
				}

				pa := new(Package)
				if err := proto.Unmarshal(bytes, pa); err != nil {
					select {
					case errc <- err:
					case <-ctx.Done():
					}
					continue
				}

				replyErrc := errc
				replyCtx := ctx
				replyRecivier := receiver
				if pa.GetReplyFlag() {
					request := requests[pa.RequestNonce]
					if request.ctx == nil {

						continue
					}
					delete(requests, pa.RequestNonce)

					replyErrc = request.errc
					replyCtx = request.ctx
					replyRecivier = request.reply
				}
				if pub != nil {
					if err := bls.Verify(suite, pub, pa.GetAnything().Value, pa.GetSignature()); err != nil {

						select {
						case replyErrc <- err:
						case <-replyCtx.Done():
						}
						continue
					}
				}

				var ptr ptypes.DynamicAny
				if err := ptypes.UnmarshalAny(pa.GetAnything(), &ptr); err != nil {

					select {
					case replyErrc <- err:
					case <-replyCtx.Done():
					}
					continue
				}
				msg := P2PMessage{Msg: ptr, Sender: pa.GetSender(), RequestNonce: pa.GetRequestNonce()}

				select {
				case replyRecivier <- msg:
				case <-replyCtx.Done():
				}

				continue
			case <-ctx.Done():
				continue
			}
		}
	}()
	return receiver, sender, errc
}

func sendBytes(ctx context.Context, c net.Conn, bytesC chan []byte, localID, remoteID []byte, incomingConn bool) chan error {
	errc := make(chan error)
	prefix := make([]byte, HEADERSIZE)
	go func() {
		for {
			select {
			case bytes, ok := <-bytesC:
				if !ok {
					return
				}

				var err error
				bytesWrite, totalBytesWrtie := 0, 0

				size := uint32(len(bytes))
				binary.BigEndian.PutUint32(prefix, size)
				bytes = append(prefix, bytes...)
				for totalBytesWrtie < len(bytes) && err == nil {
					if bytesWrite, err = c.Write(bytes[totalBytesWrtie:]); err != nil {
						select {
						case errc <- err:
						case <-ctx.Done():
							return
						}
					}
					totalBytesWrtie += bytesWrite
				}
			case <-ctx.Done():

				return
			}
		}
	}()
	return errc
}

func readBytes(ctx context.Context, c net.Conn, localID, remoteID []byte, incomingConn bool) (chan []byte, chan error) {
	out := make(chan []byte, 10)
	errc := make(chan error)
	header := make([]byte, HEADERSIZE)
	go func() {
		defer close(out)
		defer close(errc)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				var err error
				// Read until all header bytes have been read.
				bytesRead, totalBytesRead := 0, 0
				for totalBytesRead < HEADERSIZE && err == nil {
					if bytesRead, err = c.Read(header[totalBytesRead:]); err != nil {
						fmt.Println("readBytes err ", err)
						if err.Error() == "EOF" {
							c.Close()
						}

						return
					}
					totalBytesRead += bytesRead
				}

				// Decode message size and check size to avoid OOM
				size := binary.BigEndian.Uint32(header)

				if size > MAXMESSAGESIZE {
					err = errors.New("p2p message size is too big " + strconv.Itoa(int(size)))
					select {
					case errc <- err:
					case <-ctx.Done():
					}
					return
				}
				if size == 0 {
					return
				}

				// Read until all message bytes have been read.
				buffer := make([]byte, size)
				bytesRead, totalBytesRead = 0, 0
				for totalBytesRead < int(size) && err == nil {
					if bytesRead, err = c.Read(buffer[totalBytesRead:]); err != nil {
						select {
						case errc <- err:
						case <-ctx.Done():
						}
						return
					}
					totalBytesRead += bytesRead
				}
				select {
				case out <- buffer:
				case <-ctx.Done():
				}

				buffer = nil

			}
		}
	}()
	return out, errc
}

func (c *Client) send(req Request) error {
	if req.msg == nil {
		//return errors.New("Request msg is nil")
	}
	if req.ctx == nil {
		return errors.New("Request msg is nil")
	}
	go func(req Request) {
		var anything *any.Any
		var err error
		var sig []byte

		if anything, err = ptypes.MarshalAny(req.msg); err != nil {
			select {
			case req.errc <- err:
			case <-req.ctx.Done():
			}
			close(req.errc)
			if req.reply != nil {
				close(req.reply)
			}
			return
		}
		if sig, err = bls.Sign(c.suite, c.localSecKey, anything.Value); err != nil {
			select {
			case req.errc <- err:
			case <-req.ctx.Done():
			}
			close(req.errc)
			if req.reply != nil {
				close(req.reply)
			}
			return
		}

		req.p = &Package{
			Sender:    c.localID,
			Anything:  anything,
			Signature: sig,
		}
		if req.rType == 2 {
			req.p.RequestNonce = req.nonce
			req.p.ReplyFlag = true
		}

		select {
		case c.sender <- req:
		case <-req.ctx.Done():
			err := errors.New("Request Timeout")
			logger.Error(err)
			close(req.errc)
			if req.reply != nil {
				close(req.reply)
			}
		}
		return
	}(req)
	return nil
}
