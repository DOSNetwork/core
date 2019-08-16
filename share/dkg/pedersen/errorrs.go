package dkg

import (
	"fmt"
	"time"

	errors "golang.org/x/xerrors"
)

var (
	ErrDupPubKey       = errors.New("duplicated share public key")
	ErrCanNotFindID    = errors.New("can't find id in group IDs")
	ErrCasting         = errors.New("casting failed")
	ErrRespNotApproval = errors.New("response not approval")
	ErrNotCertified    = errors.New("not certified")
	ErrCanNotLoadSec   = errors.New("can't load sec")
)

type DKGError struct {
	frame errors.Frame
	t     time.Time // the time when the error happened
	err   error     // the wrapped error
}

func (e *DKGError) Format(f fmt.State, c rune) {
	errors.FormatError(e, f, c)
}

func (e *DKGError) FormatError(p errors.Printer) error {
	p.Printf("Ocurred at: %s", e.t)
	e.frame.Format(p)
	return e.err
}

func (e *DKGError) Error() string {
	return fmt.Sprint(e)
}

func (e *DKGError) Unwrap() error {
	return e.err
}
