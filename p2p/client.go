package p2p

import (
	"context"
	"encoding/binary"
	"io"
	"net"
	//	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	errors "golang.org/x/xerrors"

	"crypto/aes"
	"crypto/cipher"

	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/suites"
	"github.com/DOSNetwork/core/utils"
	"github.com/dedis/kyber"
	///"github.com/felixge/tcpkeepalive"
)

const (
	msgSizeLimit = 1024 * 1024
	bufSize      = 1024
	headerSize   = 4
	idleTimeout  = 60 * time.Second
)

type signFunc func(msg []byte) (sig []byte, err error)
type verifyFunc func(msg, sig []byte) (err error)

type client struct {
	addr     string
	conn     net.Conn
	inBound  bool
	localID  []byte
	remoteID []byte

	ctx      context.Context
	cancel   context.CancelFunc
	peerSend chan p2pRequest
	peerFeed chan P2PMessage
	errc     chan error

	//TODO : Move to other module
	suite        suites.Suite
	localSecKey  kyber.Scalar
	localPubKey  kyber.Point
	remotePubKey kyber.Point
	dhKey        []byte
	dhNonce      []byte
}

func newClient(localID []byte, conn net.Conn, peerFeed chan P2PMessage, inBound bool) (c *client) {
	/*
		kaConn, _ := tcpkeepalive.EnableKeepAlive(conn)
		kaConn.SetKeepAliveIdle(10 * time.Second)
		kaConn.SetKeepAliveCount(5)
		kaConn.SetKeepAliveInterval(10 * time.Second)
	*/
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		//error handle
	}
	tcpConn.SetKeepAlive(true)
	tcpConn.SetKeepAlivePeriod(time.Second * 1)
	c = &client{
		localID: localID,
		conn:    tcpConn,
		inBound: inBound,
		errc:    make(chan error),
	}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.peerSend = make(chan p2pRequest, 21)
	c.peerFeed = peerFeed
	//TODO : Move to other module
	c.suite = suites.MustFind("bn256")
	c.localSecKey = c.suite.Scalar().Pick(c.suite.RandomStream())
	c.localPubKey = c.suite.Point().Mul(c.localSecKey, nil)
	return
}

func (c *client) handShake(ctx context.Context) (err chan error) {
	return utils.MergeErrors(ctx, c.sendID(ctx), c.receiveID(ctx))
}

func (c *client) receiveID(ctx context.Context) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)

		buffer, err := readFrom(c.conn)
		if err != nil {
			utils.ReportError(ctx, errc, errors.Errorf("readFrom %s : %w", c.conn.RemoteAddr().String(), err))
			return
		}

		_, ptr, err := decodeBytes(buffer, nil)
		if err != nil {
			utils.ReportError(ctx, errc, errors.Errorf("decodeBytes: %w", err))
			return
		}

		id, ok := ptr.Message.(*ID)
		if !ok {
			err = errors.Errorf("ID casting: %w", ErrCasting)
			utils.ReportError(ctx, errc, err)
			return
		}
		//ID should not be nil or the same with local ID
		c.remoteID = id.GetId()
		if string(c.remoteID) == string(c.localID) {
			err = errors.Errorf("remoteID %b != localID %b: %w",
				c.remoteID, c.localID, ErrDuplicateID)
			utils.ReportError(ctx, errc, errors.Errorf("client : %w", err))
		}
		if c.remoteID == nil {
			err = errors.Errorf("remoteID is nil: %w", ErrNoRemoteID)
		}

		//TODO: Move to other module
		pub := c.suite.G2().Point()
		if err = pub.UnmarshalBinary(id.GetPublicKey()); err != nil {
			utils.ReportError(ctx, errc, errors.Errorf("UnmarshalBinary: %w", err))
			return
		}
		c.remotePubKey = pub

		var dhBytes []byte
		dhKey := c.suite.Point().Mul(c.localSecKey, c.remotePubKey)
		if dhBytes, err = dhKey.MarshalBinary(); err != nil {
			utils.ReportError(ctx, errc, errors.Errorf("MarshalBinary: %w", err))
			return
		}
		c.dhKey = dhBytes[0:32]
		c.dhNonce = dhBytes[32:44]

		return
	}()
	return
}

func (c *client) sendID(ctx context.Context) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)

		var err error
		var bytes, pubKeyBytes []byte
		if pubKeyBytes, err = c.localPubKey.MarshalBinary(); err != nil {
			utils.ReportError(ctx, errc, errors.Errorf("MarshalBinary: %w", err))
			return
		}

		pID := &ID{
			PublicKey: pubKeyBytes,
			Id:        c.localID,
		}

		if bytes, err = encodeProto(pID, c.localID, nil, 0, false); err != nil {
			utils.ReportError(ctx, errc, errors.Errorf("encodeProto: %w", err))
		}

		if err = writeTo(bytes, c.conn); err != nil {
			utils.ReportError(ctx, errc, errors.Errorf("writeTo: %w", err))
		}
		return
	}()
	return
}

func (c *client) run() (err error) {
	c.sendPipe(c.encryptPipe(c.packPipe(c.dispatch(c.decodePipe(c.decryptPipe(c.readPipe()))))))
	for {
		var ok bool
		select {
		case <-c.ctx.Done():
			return
		case err, ok = <-c.errc:
			if ok {
				if errors.Is(err, io.EOF) {
					err = nil
					return
				}
				return
			}
		}
	}
	return
}

func (c *client) close() (err error) {
	c.cancel()
	err = c.conn.Close()
	if err != nil {
		err = errors.Errorf("client close: %w", err)
	}
	<-c.ctx.Done()
	return
}

func (c *client) send(req p2pRequest) {
	go func() {
		select {
		case <-c.ctx.Done():
		default:
			if err := req.sendReq(c.peerSend); err != nil {
				req.replyResult(nil, errors.Errorf("client send: %w", err))
			}
		}
	}()
}

func (c *client) encryptPipe(plaintext chan []byte) (out chan []byte) {
	out = make(chan []byte)

	go func() {
		defer close(out)
		for {
			var result []byte
			select {
			case <-c.ctx.Done():
				return
			case text, ok := <-plaintext:
				if ok {
					var err error
					var block cipher.Block
					var aesgcm cipher.AEAD

					if block, err = aes.NewCipher(c.dhKey); err != nil {
						c.reportError(errors.Errorf("client encryptPipe: %w", err))
						continue
					}

					if aesgcm, err = cipher.NewGCM(block); err != nil {
						c.reportError(errors.Errorf("client encryptPipe: %w", err))
						continue
					}
					result = aesgcm.Seal(nil, c.dhNonce, text, nil)
				}
			}
			select {
			case <-c.ctx.Done():
			case out <- result:
			}
		}
	}()
	return out
}

func (c *client) decryptPipe(ciphertext chan []byte) (out chan []byte) {
	out = make(chan []byte)

	go func() {
		defer close(out)
		for {
			var result []byte
			select {
			case <-c.ctx.Done():
				return
			case text, ok := <-ciphertext:
				if ok {
					var block cipher.Block
					var aesgcm cipher.AEAD
					var err error
					if block, err = aes.NewCipher(c.dhKey); err != nil {
						c.reportError(errors.Errorf("client decryptPipe: %w", err))
						continue
					}

					if aesgcm, err = cipher.NewGCM(block); err != nil {
						c.reportError(errors.Errorf("client decryptPipe: %w", err))
						continue
					}

					if result, err = aesgcm.Open(nil, c.dhNonce, text, nil); err != nil {
						c.reportError(errors.Errorf("client decryptPipe: %w", err))
						continue
					}
					select {
					case out <- result:
					case <-c.ctx.Done():
					}
				}
			}
		}
	}()
	return out
}

func (c *client) dispatch(replyMsg, receivedMsg chan P2PMessage) (out chan p2pRequest) {
	out = make(chan p2pRequest)
	requests := make(map[uint64]*p2pRequest)
	var nonce uint64
	idleTimer := time.NewTimer(idleTimeout)
	var readTimer bool
	go func() {
		defer close(out)
		for {
			var (
				ok  bool
				req p2pRequest
				msg P2PMessage
			)
			select {
			case <-c.ctx.Done():
				for _, req := range requests {
					req.replyResult(nil, errors.Errorf("client dispatch: %w", c.ctx.Err()))
				}
				if !idleTimer.Stop() && !readTimer {
					<-idleTimer.C
				}
				return
			case <-idleTimer.C:
				readTimer = true
				err := c.close()
				if err != nil {
					errors.Errorf("client close: %w", err)
				}
			case req, ok = <-c.peerSend:
				if ok {
					if !idleTimer.Stop() {
						<-idleTimer.C
					}
					idleTimer.Reset(idleTimeout)
					if req.rType != replyReq {
						req.nonce = nonce
						requests[nonce] = &req
						nonce++
					}
					select {
					case <-c.ctx.Done():
					case out <- req:
					}
				}
			case msg, ok = <-replyMsg:
				if ok {
					if !idleTimer.Stop() {
						<-idleTimer.C
					}
					idleTimer.Reset(idleTimeout)
					p2pRequest := requests[msg.RequestNonce]
					if p2pRequest != nil {
						delete(requests, msg.RequestNonce)
						select {
						case <-p2pRequest.ctx.Done():
						default:
							p2pRequest.replyResult(msg, nil)
						}
					}
				}
			case msg, ok = <-receivedMsg:
				if ok {
					if !idleTimer.Stop() {
						<-idleTimer.C
					}
					idleTimer.Reset(idleTimeout)
					c.reportMsg(msg)
				}
			}
		}
	}()
	return
}

func (c *client) packPipe(reqC chan p2pRequest) (out chan []byte) {
	out = make(chan []byte)
	go func() {
		defer close(out)
		for {
			var err error
			var bytes []byte
			select {
			case <-c.ctx.Done():
				return
			case req, ok := <-reqC:
				if ok {

					select {
					case <-req.ctx.Done():
					default:
					}
					replyFlag := false
					if req.rType == replyReq {
						req.replyResult(nil, nil)
						replyFlag = true
					}
					bytes, err = encodeProto(req.msg, c.localID, c.signFn, req.nonce, replyFlag)
					if err != nil {
						c.reportError(errors.Errorf("client packPipe: %w", err))
					}
				}
			}
			select {
			case <-c.ctx.Done():
			case out <- bytes:
			}
		}
	}()
	return
}

func (c *client) sendPipe(bytesC chan []byte) {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			case bytes, ok := <-bytesC:
				if ok {
					err := writeTo(bytes, c.conn)
					if err != nil {
						c.reportError(errors.Errorf("client sendPipe: %w", err))
					}
				}
			}
		}
	}()
	return
}

func (c *client) decodePipe(bytesC chan []byte) (replyMsg, receivedMsg chan P2PMessage) {
	replyMsg = make(chan P2PMessage)
	receivedMsg = make(chan P2PMessage)
	go func() {
		defer close(replyMsg)
		defer close(receivedMsg)
		for {
			var msg P2PMessage
			select {
			case <-c.ctx.Done():
				return
			case bytes, ok := <-bytesC:
				if ok {
					if len(bytes) == 0 {
						continue
					}
					pa, ptr, err := decodeBytes(bytes, c.verifyFn)
					if err != nil {
						c.reportError(errors.Errorf("client decodePipe: %w", err))
						continue
					}
					//TODO: Move to other module
					if err := bls.Verify(c.suite, c.remotePubKey, pa.GetAnything().Value, pa.GetSignature()); err != nil {
						c.reportError(errors.Errorf("client decodePipe: %w", err))
						continue
					}
					msg = P2PMessage{Msg: ptr, Sender: pa.GetSender(), RequestNonce: pa.GetRequestNonce()}
					if pa.GetReplyFlag() {
						select {
						case <-c.ctx.Done():
						case replyMsg <- msg:
						}
					} else {
						select {
						case <-c.ctx.Done():
						case receivedMsg <- msg:
						}
					}
				}
			}
		}
	}()
	return
}

func (c *client) readPipe() (out chan []byte) {
	out = make(chan []byte, 10)
	go func() {
		defer close(out)
		for {
			var buffer []byte
			var err error
			select {
			case <-c.ctx.Done():
				return
			default:
				buffer, err = readFrom(c.conn)
				if err != nil {
					c.reportError(errors.Errorf("readPipe: %w", err))
					return
				}
			}
			select {
			case <-c.ctx.Done():
			case out <- buffer:
			}
		}
	}()
	return out
}

func encodeProto(msg proto.Message, sender []byte, signFn signFunc, nonce uint64, replyFlag bool) (bytes []byte, err error) {
	var anything *any.Any
	var sign []byte
	if anything, err = ptypes.MarshalAny(msg); err != nil {
		err = errors.Errorf("MarshalAny: %w", err)
		return
	}
	if signFn != nil {
		if sign, err = signFn(anything.Value); err != nil {
			err = errors.Errorf("signFn: %w", err)
			return
		}
	}
	p := &Package{
		Anything:     anything,
		Sender:       sender,
		Signature:    sign,
		RequestNonce: nonce,
		ReplyFlag:    replyFlag,
	}
	if bytes, err = proto.Marshal(p); err != nil {
		err = errors.Errorf("Marshal: %w", err)
	}
	return
}

func decodeBytes(bytes []byte, veifyfn verifyFunc) (pa *Package, ptr ptypes.DynamicAny, err error) {
	pa = &Package{}
	//TestPoint
	//bytes = []byte{}
	if err = proto.Unmarshal(bytes, pa); err != nil {
		err = errors.Errorf("Unmarshal: %w", err)
		return
	}
	if veifyfn != nil {
		if err = veifyfn(pa.GetAnything().Value, pa.GetSignature()); err != nil {
			err = errors.Errorf("veifyfn: %w", err)
			return
		}
	}
	if err = ptypes.UnmarshalAny(pa.GetAnything(), &ptr); err != nil {
		err = errors.Errorf("UnmarshalAny: %w", err)
	}
	return
}

func writeTo(bytes []byte, conn net.Conn) (err error) {
	prefix := make([]byte, headerSize)
	bytesWrite, totalBytesWrtie := 0, 0

	size := len(bytes)
	//TestPoint
	//size = msgSizeLimit + 1
	if size > msgSizeLimit {
		err = errors.Errorf("SizeLimit %d size %d: %w",
			msgSizeLimit, size, ErrMsgOverSize)
		return
	}
	binary.BigEndian.PutUint32(prefix, uint32(size))
	bytes = append(prefix, bytes...)
	//TestPoint
	//conn.Close()
	for totalBytesWrtie < len(bytes) && err == nil {
		if bytesWrite, err = conn.Write(bytes[totalBytesWrtie:]); err != nil {
			err = errors.Errorf("conn write: %w", err)
			return
		}
		totalBytesWrtie += bytesWrite
	}
	return
}

//read shold not be called in multiple go routine
func readFrom(conn net.Conn) (buffer []byte, err error) {
	// Read until all header bytes have been read.
	header := make([]byte, headerSize)
	// Read until all header bytes have been read.
	bytesRead, totalBytesRead := 0, 0
	//TestPoint
	//conn.Close()
	for totalBytesRead < headerSize && err == nil {
		if bytesRead, err = conn.Read(header[totalBytesRead:]); err != nil {
			err = errors.Errorf("conn read header: %w", err)
			return
		}
		totalBytesRead += bytesRead
	}

	// Decode message size.
	size := binary.BigEndian.Uint32(header)
	header = nil
	//TestPoint
	//size = msgSizeLimit + 1
	if size > msgSizeLimit || size <= 0 {
		err = errors.Errorf("SizeLimit %d size %d: %w", msgSizeLimit, size, ErrMsgOverSize)
		return
	}

	// Read until all message bytes have been read.
	buffer = make([]byte, size)
	contentBytesRead, totalContentBytesRead := 0, 0
	//TestPoint
	//conn.Close()
	for totalContentBytesRead < int(size) && err == nil {
		if contentBytesRead, err = conn.Read(buffer[totalContentBytesRead:]); err != nil {
			err = errors.Errorf("conn read content: %w", err)
			return
		}
		totalContentBytesRead += contentBytesRead
	}
	return
}

func (c *client) reportError(err error) {
	select {
	case <-c.ctx.Done():
	case c.errc <- err:
	}
}

func (c *client) reportMsg(msg P2PMessage) {
	select {
	case <-c.ctx.Done():
	case c.peerFeed <- msg:
	}
}

//TODO : Move to other module
func (c *client) signFn(msg []byte) (sig []byte, err error) {
	if sig, err = bls.Sign(c.suite, c.localSecKey, msg); err != nil {
		err = errors.Errorf("Sign: %w", err)
	}
	return
}
func (c *client) verifyFn(msg, sig []byte) (err error) {
	if err = bls.Verify(c.suite, c.remotePubKey, msg, sig); err != nil {
		err = errors.Errorf("Verify: %w", err)
	}
	return
}
