package dkg

import (
	"fmt"

	errors "golang.org/x/xerrors"
)

var (
	ErrDupPubKey          = errors.New("duplicated share public key")
	ErrDupPubKeyIndex     = errors.New("duplicated public key index")
	ErrCanNotFindID       = errors.New("can't find id in group IDs")
	ErrCasting            = errors.New("casting failed")
	ErrRespNotApproval    = errors.New("response not approval")
	ErrNotCertified       = errors.New("not certified")
	ErrCanNotLoadSec      = errors.New("can't load sec")
	ErrCanNotLoadGroup    = errors.New("can't load group")
	ErrResponseNoApproval = errors.New("response no approval")
	ErrDKGNotCertified    = errors.New("dkg is not certified")
)

type DKGError struct {
	frame errors.Frame
	err   error // the wrapped error
}

func (e *DKGError) Format(f fmt.State, c rune) {
	errors.FormatError(e, f, c)
}

func (e *DKGError) FormatError(p errors.Printer) error {
	e.frame.Format(p)
	return e.err
}

func (e *DKGError) Error() string {
	return fmt.Sprint(e)
}

func (e *DKGError) Unwrap() error {
	return e.err
}
