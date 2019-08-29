package onchain

import (
	"fmt"
	"time"

	errors "golang.org/x/xerrors"
)

var (
	ErrNotValidAddr = errors.New("not a valid hex address")
	ErrNoWSURL      = errors.New("no web socket URL")
	ErrNoKeystore   = errors.New("no keystore")
	ErrNoProxy      = errors.New("no proxy instance")
	ErrNoCR         = errors.New("no commir reveal instance")
	ErrCast         = errors.New("cast failed")
)

type OnchainError struct {
	frame errors.Frame
	t     time.Time // the time when the error happened
	err   error     // the wrapped error
}

func (e *OnchainError) Format(f fmt.State, c rune) {
	errors.FormatError(e, f, c)
}

func (e *OnchainError) FormatError(p errors.Printer) error {
	e.frame.Format(p)
	return e.err
}

func (e *OnchainError) Error() string {
	return fmt.Sprint(e)
}

func (e *OnchainError) Unwrap() error {
	return e.err
}
