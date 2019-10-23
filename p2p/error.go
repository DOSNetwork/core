package p2p

import (
	"fmt"
	"time"

	errors "golang.org/x/xerrors"
)

var (
	ErrNoRemoteID       = errors.New("remoteID is nil")
	ErrMsgOverSize      = errors.New("header size check failed")
	ErrCasting          = errors.New("casting failed")
	ErrDuplicateID      = errors.New("remote ID is the same with local ID")
	ErrCanNotFindClient = errors.New("can't find client")
	ErrIdleTimeout      = errors.New("idle timeout")
)

type P2PError struct {
	frame errors.Frame
	dest  string
	t     time.Time // the time when the error happened
	err   error     // the wrapped error
}

func (e *P2PError) Format(f fmt.State, c rune) {
	errors.FormatError(e, f, c)
}

func (e *P2PError) FormatError(p errors.Printer) error {
	p.Printf("IP : %s", e.dest)
	e.frame.Format(p)
	return e.err
}

func (e *P2PError) Error() string {
	return fmt.Sprint(e)
}

func (e *P2PError) Unwrap() error {
	return e.err
}
