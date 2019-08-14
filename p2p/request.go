package p2p

import (
	"context"

	"github.com/golang/protobuf/proto"
)

type request struct {
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
	reply chan interface{}
	errc  chan error
}

func (r *request) sendReq(ch chan request) (err error) {
	select {
	case ch <- *r:
	case <-r.ctx.Done():
		err = r.ctx.Err()
	}
	return
}

func (r *request) waitForResult() (result interface{}, err error) {
	defer close(r.reply)
	defer close(r.errc)
	select {
	case result = <-r.reply:
	case err = <-r.errc:
	case <-r.ctx.Done():
		err = r.ctx.Err()
	}
	return
}
func (r *request) waitForError() (err error) {
	defer close(r.errc)
	select {
	case err = <-r.errc:
	case <-r.ctx.Done():
	}
	return
}

func (r *request) replyResult(result interface{}) {
	defer r.cancel()
	select {
	case <-r.ctx.Done():
	case r.reply <- result:
	}
}

func (r *request) replyError(err error) {
	defer r.cancel()
	select {
	case <-r.ctx.Done():
	case r.errc <- err:
	}
}
