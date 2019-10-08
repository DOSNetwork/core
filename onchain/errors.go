package onchain

import (
	"fmt"

	errors "golang.org/x/xerrors"
)

type OnchainError struct {
	frame errors.Frame
	Idx   int
	err   error // the wrapped error
}

func (e *OnchainError) Format(f fmt.State, c rune) {
	errors.FormatError(e, f, c)
}

func (e *OnchainError) FormatError(p errors.Printer) error {
	p.Printf("Ocurred at: %d", e.Idx)
	e.frame.Format(p)
	return e.err
}

func (e *OnchainError) Error() string {
	return fmt.Sprint(e)
}

func (e *OnchainError) Unwrap() error {
	return e.err
}
