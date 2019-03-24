package p2p

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	//	"io"
	//"encoding/hex"
	"net"
	"os"
	"sync"
	"time"
	//	"github.com/DOSNetwork/core/log"
	//"github.com/DOSNetwork/core/p2p/internal"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/suites"

	"github.com/dedis/kyber"

	"crypto/aes"
	"crypto/cipher"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
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
	receiver chan P2PMessage
	errc     chan error
}

type Request struct {
	ctx   context.Context
	p     *Package
	reply chan P2PMessage
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
	client.sender = make(chan Request, 100)
	fmt.Println(time.Now(), "!!!!!!NewClient , ", incomingConn, " ", client.localID, client.remoteID)

	//Wait for exchanging ID complete
	err = client.exchangeID()
	if err != nil {
		fmt.Println(time.Now(), "!!!!!!NewClient , exchangeID err", err, incomingConn, " ", client.localID, client.remoteID)
		return
	}
	fmt.Println(time.Now(), "!!!!!!NewClient , exchangeID ", incomingConn, " ", client.localID, client.remoteID)

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
	fmt.Println("Client Close")
	c.conn.Close()
	c.cancel()

}
func (c *Client) ErrorHandling(errc <-chan error) {
	go func() {
		for err := range errc {
			if err.Error() != "cipher: message authentication failed" {
				fmt.Println("ErrorHandling err ", err)
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
	//fmt.Printf("Wait for waitgroup (up to %s)\n", timeout)
	select {
	case <-ch:
		//fmt.Printf("Wait group finished\n")
	case <-time.After(timeout):
		fmt.Printf("Timed out waiting for wait group\n")
	}
	return
}
func (c *Client) receiveID(wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	buffer, err := readConn(c.conn, c.localID, c.remoteID, c.incomingConn)

	pa := new(Package)
	if err = proto.Unmarshal(buffer, pa); err != nil {
		fmt.Println(time.Now(), "!!!!!!receiveID Unmarshal err ", err.Error(), c.incomingConn, " ", c.localID, c.remoteID)
		return
	}
	var ptr ptypes.DynamicAny
	if err = ptypes.UnmarshalAny(pa.GetAnything(), &ptr); err != nil {
		fmt.Println(time.Now(), "!!!!!!receiveID UnmarshalAny err ", err.Error(), c.incomingConn, " ", c.localID, c.remoteID)
		return
	}
	id, ok := ptr.Message.(*ID)
	if !ok {
		fmt.Println(time.Now(), "!!!!!!receiveID cast to ID !OK err ", c.incomingConn, " ", c.localID, c.remoteID)
		return
	}
	c.remoteID = id.GetId()
	if string(c.remoteID) == string(c.localID) {
		os.Exit(2)
	}
	pub := c.suite.G2().Point()
	if err = pub.UnmarshalBinary(id.GetPublicKey()); err != nil {
		fmt.Println(time.Now(), "!!!!!!receiveID UnmarshalBinary err ", err.Error(), c.incomingConn, " ", c.localID, c.remoteID)
		return
	}
	c.remotePubKey = pub

	var dhBytes []byte
	dhKey := c.suite.Point().Mul(c.localSecKey, c.remotePubKey)
	if dhBytes, err = dhKey.MarshalBinary(); err != nil {
		fmt.Println(time.Now(), "!!!!!!receiveID MarshalBinary err ", err.Error(), c.incomingConn, " ", c.localID, c.remoteID)
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
		fmt.Println("prepareID pubKey.MarshalBinary ", err)
		return
	}

	pID := &ID{
		PublicKey: pubKeyBytes,
		Id:        c.localID,
	}

	if anything, err = ptypes.MarshalAny(pID); err != nil {
		fmt.Println("prepareID ptypes.MarshalAny ", err)
		return
	}
	pa := &Package{
		Anything: anything,
	}
	if bytes, err = proto.Marshal(pa); err != nil {
		fmt.Println("prepareID proto.Marshal(pa) ", err)
		return
	}
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(bytes)))
	bytes = append(prefix, bytes...)
	if _, E := c.conn.Write(bytes); E != nil {
		fmt.Println("prepareID sendBytes err ", E)
		return
	}
	//fmt.Println(time.Now(), "!!!!!!conn.Write ", c.incomingConn, " ", c.localID, c.remoteID)

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
					fmt.Println("NewCipher err ", err)
					select {
					case errc <- err:
					case <-ctx.Done():
					}
					continue
				}

				if aesgcm, err = cipher.NewGCM(block); err != nil {
					fmt.Println("NewGCM err ", err)
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

func dispatch(localID, remoteID []byte, incomingConn bool, ctx context.Context, suite suites.Suite, pub kyber.Point, Id []byte, sendMsg chan Request, receiveBytes chan []byte) (chan P2PMessage, chan []byte, chan error) {
	receiver := make(chan P2PMessage, 100)
	sender := make(chan []byte, 100)
	errc := make(chan error)
	requests := make(map[uint64]Request)
	var nonce uint64
	go func() {
		defer fmt.Println("close dispatch")
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
					fmt.Println(time.Now(), "!!!!!!Dispatch receive req to sent , ", nonce, " ", incomingConn, " ", localID, remoteID)
					if req.ctx == nil {
						fmt.Println(time.Now(), "!!!!!!Dispatch  not find request ", nonce, incomingConn, " ", localID, remoteID)
					}
					if req.reply == nil {
						fmt.Println(time.Now(), "!!!!!!Dispatch  not find reply ", nonce, incomingConn, " ", localID, remoteID)
					}
					if req.ctx == nil && req.reply == nil {
						continue
					}
					requests[nonce] = req
					req.p.RequestNonce = nonce
					nonce++

				} else {
					//fmt.Println(time.Now(), "!!!!!!Dispatch receive req to reply , ", incomingConn, " ", localID, remoteID)

				}
				//Encode the package
				var bytes []byte
				var err error
				if bytes, err = proto.Marshal(req.p); err != nil {
					fmt.Println(time.Now(), "!!!!!!Dispatch sent Marshal err , ", err, incomingConn, " ", localID, remoteID)

					select {
					case errc <- err:
					case <-ctx.Done():
					}
					continue
				}
				select {
				case sender <- bytes:
					fmt.Println(time.Now(), "!!!!!!Dispatch sent Done , ", incomingConn, " ", localID, remoteID)

				case <-ctx.Done():
					fmt.Println(time.Now(), "!!!!!!Dispatch sent ctx , ", ctx.Err(), incomingConn, " ", localID, remoteID)
				}
			case bytes, ok := <-receiveBytes:
				if !ok {
					fmt.Println("Dispatch receive byte !ok")
					return
				}
				//fmt.Println(time.Now(), "!!!!!!Dispatch receive byte , ", incomingConn, " ", localID, remoteID)

				pa := new(Package)
				if err := proto.Unmarshal(bytes, pa); err != nil {
					//fmt.Println(time.Now(), "!!!!!!Dispatch Unmarshal err ", err.Error(), incomingConn, " ", localID, remoteID)
					continue
				}

				if pub != nil {
					if err := bls.Verify(suite, pub, pa.GetAnything().Value, pa.GetSignature()); err != nil {
						//fmt.Println(time.Now(), "!!!!!!Dispatch Verify err ", err.Error(), incomingConn, " ", localID, remoteID)
						select {
						case errc <- err:
						case <-ctx.Done():
						}
						continue
					}
				}

				var ptr ptypes.DynamicAny
				if err := ptypes.UnmarshalAny(pa.GetAnything(), &ptr); err != nil {
					//fmt.Println(time.Now(), "!!!!!!Dispatch UnmarshalAny err ", err.Error(), incomingConn, " ", localID, remoteID)
					continue
				}

				msg := P2PMessage{Msg: ptr, Sender: pa.GetSender(), RequestNonce: pa.GetRequestNonce()}

				if pa.GetReplyFlag() {
					fmt.Println(time.Now(), "!!!!!!Dispatch Got a reply ", incomingConn, " ", localID, remoteID)

					request := requests[pa.RequestNonce]
					if request.ctx == nil {
						fmt.Println(time.Now(), "!!!!!!Dispatch  not find request ", pa.RequestNonce, incomingConn, " ", localID, remoteID)
					}
					if request.reply == nil {
						fmt.Println(time.Now(), "!!!!!!Dispatch  not find reply ", pa.RequestNonce, incomingConn, " ", localID, remoteID)
					}
					select {
					case <-request.ctx.Done():
						//fmt.Println(time.Now(), "!!!!!!Dispatch reply Ctx ", ctx.Err(), incomingConn, " ", localID, remoteID)

					case request.reply <- msg:
						//fmt.Println(time.Now(), "!!!!!!Dispatch reply Done", incomingConn, " ", localID, remoteID)

					}
					delete(requests, pa.RequestNonce)

					continue
				}

				select {
				case receiver <- msg:
					//fmt.Println(time.Now(), "!!!!!!Dispatch receive byte send to receiver Done", incomingConn, " ", localID, remoteID)

				case <-ctx.Done():
					//fmt.Println(time.Now(), "!!!!!!Dispatch receive byte send to receiverCtx ", ctx.Err(), incomingConn, " ", localID, remoteID)

				}
				//}()
			case <-ctx.Done():
				return
			}
		}
	}()
	return receiver, sender, errc
}

func sendBytes(ctx context.Context, c net.Conn, bytesC chan []byte, localID, remoteID []byte, incomingConn bool) chan error {
	errc := make(chan error)
	go func() {
		for {
			select {
			case bytes, ok := <-bytesC:
				if !ok {
					return
				}

				prefix := make([]byte, 4)
				binary.BigEndian.PutUint32(prefix, uint32(len(bytes)))
				bytes = append(prefix, bytes...)
				if _, err := c.Write(bytes); err != nil {
					//fmt.Println(time.Now(), "!!!!!!sendBytes , ", incomingConn, " ", localID, remoteID, " err", err)

					select {
					case errc <- err:
					case <-ctx.Done():
						fmt.Println("sendBytes ctx ", ctx.Err())
						return
					}
				} else {
					//fmt.Println(time.Now(), "!!!!!!sendBytes , ", incomingConn, " ", localID, remoteID)

				}

			case <-ctx.Done():
				fmt.Println("sendBytes ctx ", ctx.Err())

				return
			}
		}
	}()
	return errc
}

func readConn(c net.Conn, localID, remoteID []byte, incomingConn bool) (buffer []byte, err error) {
	// Read until all header bytes have been read.
	header := make([]byte, 4)
	// Read until all header bytes have been read.
	bytesRead, totalBytesRead := 0, 0
	for totalBytesRead < 4 && err == nil {
		if bytesRead, err = c.Read(header[totalBytesRead:]); err != nil {
			//fmt.Println(time.Now(), "!!!!!!ReadByte , err ", err.Error(), incomingConn, " ", localID, remoteID)
			if err.Error() == "EOF" {
				c.Close()
			}
			return
		}
		totalBytesRead += bytesRead
	}

	// Decode message size.
	size := binary.BigEndian.Uint32(header)
	//fmt.Println(time.Now(), "!!!!!!ReadByte , ", size, incomingConn, " ", localID, remoteID)

	// Read until all message bytes have been read.
	buffer = make([]byte, size)
	if size == 0 {
		//fmt.Println(time.Now(), "!!!!!!message is 0 , ", incomingConn, " ", localID, remoteID)
		return
	}
	bytesRead, totalBytesRead = 0, 0

	for totalBytesRead < int(size) && err == nil {
		if bytesRead, err = c.Read(buffer[totalBytesRead:]); err != nil {
			//fmt.Println(time.Now(), "!!!!!!ReadByte , err ", err.Error(), incomingConn, " ", localID, remoteID)
		}
		totalBytesRead += bytesRead
	}

	return
}

func readBytes(ctx context.Context, c net.Conn, localID, remoteID []byte, incomingConn bool) (chan []byte, chan error) {
	out := make(chan []byte, 10)
	errc := make(chan error)

	go func() {
		defer close(out)
		defer close(errc)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Read until all header bytes have been read.
				buffer, err := readConn(c, localID, remoteID, incomingConn)
				if err != nil {
					select {
					case errc <- err:
					case <-ctx.Done():
					}
				} else {
					select {
					case out <- buffer:
					case <-ctx.Done():
					}
				}
			}
		}
	}()
	return out, errc
}

func (c *Client) Request(msg proto.Message) (P2PMessage, error) {
	var anything *any.Any
	var err error
	var sig []byte
	fmt.Println(time.Now(), "!!!!!!!!!Request from ", c.localID, " to ", c.remoteID)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if anything, err = ptypes.MarshalAny(msg); err != nil {
		fmt.Println("Request MarshalAny err ", err.Error())
		return P2PMessage{}, err
	}
	if sig, err = bls.Sign(c.suite, c.localSecKey, anything.Value); err != nil {
		fmt.Println("Request bls.Sign err ", err.Error())
		return P2PMessage{}, err
	}
	pa := Package{
		Sender:    c.localID,
		Anything:  anything,
		Signature: sig,
	}

	reply := make(chan P2PMessage)
	req := Request{
		ctx:   ctx,
		p:     &pa,
		reply: reply,
	}

	select {
	case c.sender <- req:
		//fmt.Println(time.Now(), "!!!!!!!!Request has been sent to sender  ", c.localID, " to ", c.remoteID)

	case <-ctx.Done():
		//fmt.Println(time.Now(), "!!!!!!!!Request timeout  ", c.localID, " to ", c.remoteID, " err", ctx.Err())
		err := errors.New("Request Timeout")
		logger.Error(err)
	}

	select {
	case r := <-reply:
		//fmt.Println(time.Now(), "!!!!!!!!Request reply has been received  ", c.localID, " to ", c.remoteID)
		return r, err
	case <-ctx.Done():
		//fmt.Println(time.Now(), "!!!!!!!!Request wait for Reply timeout  ", c.localID, " to ", c.remoteID, " err", ctx.Err())
		err := errors.New("Request Timeout")
		logger.Error(err)
	}
	return P2PMessage{}, ctx.Err()
}

func (c *Client) Reply(nonce uint64, msg proto.Message) error {
	var anything *any.Any
	var err error
	var sig []byte
	////fmt.Println(time.Now(), "!!!!!!!!Reply  ", c.localID, " to ", c.remoteID)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx := context.Background()
	if anything, err = ptypes.MarshalAny(msg); err != nil {
		fmt.Println("Reply MarshalAny err ", err.Error())

		return err
	}
	if c.suite == nil {
		fmt.Println(" Reply err c.suite == nil", c.suite)
	}
	if c.suite == nil {
		fmt.Println(" Reply err c.localSecKey == nil", c.localSecKey)
	}
	if anything.Value == nil {
		fmt.Println(" Reply err anything.Value", anything.Value)
	}
	if sig, err = bls.Sign(c.suite, c.localSecKey, anything.Value); err != nil {
		fmt.Println("Reply bls.Sig err ", err.Error())

		return err
	}
	pa := Package{
		Sender:       c.localID,
		Anything:     anything,
		Signature:    sig,
		RequestNonce: nonce,
		ReplyFlag:    true,
	}

	req := Request{
		ctx: ctx,
		p:   &pa,
	}
	select {
	case c.sender <- req:
		//fmt.Println(time.Now(), "!!!!!!!!!Reply Success  ", c.localID, " to ", c.remoteID)

		return nil
	case <-ctx.Done():
		//fmt.Println(time.Now(), "!!!!!!!!!Reply Timeout  ", c.localID, " to ", c.remoteID, " err", ctx.Err())
		err := errors.New("Request Timeout")
		logger.Error(err)
	}
	return ctx.Err()
}
