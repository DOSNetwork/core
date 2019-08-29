package p2p

import (
	"context"
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
)

type p2pRequest struct {
	rType  int
	ctx    context.Context
	cancel context.CancelFunc
	addr   string
	id     []byte
	//client signs and packs msg into Package
	msg proto.Message
	p   *Package
	//
	nonce uint64
	reply chan p2pResult
	once  sync.Once
}
type p2pResult struct {
	res interface{}
	err error
}

func NewP2pRequest(ctx context.Context, rType int, id []byte, addr string, msg proto.Message, nonce uint64) (req *p2pRequest) {
	rctx, cancel := context.WithCancel(ctx)
	req = &p2pRequest{ctx: rctx, cancel: cancel, id: id, msg: msg, rType: rType, nonce: nonce, reply: make(chan p2pResult)}
	return
}
func (r *p2pRequest) sendReq(ch chan p2pRequest) (err error) {
	select {
	case ch <- *r:
	case <-r.ctx.Done():
		err = r.ctx.Err()
		fmt.Println("p2pRequest sendReq err", err)
	}
	return
}

func (r *p2pRequest) waitForResult() (res interface{}, err error) {
	select {
	case <-r.ctx.Done():
		err = r.ctx.Err()
	case result := <-r.reply:
		res = result.res
		err = result.err
	}
	r.cancel()
	return
}

func (r *p2pRequest) replyResult(res interface{}, err error) {
	r.once.Do(func() {
		defer close(r.reply)
		select {
		case <-r.ctx.Done():
		case r.reply <- p2pResult{res: res, err: err}:
		}
	})
}
