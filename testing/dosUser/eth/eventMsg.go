package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// AskMeAnythingSetTimeout represents a SetTimeout event raised by the AskMeAnything contract.
type AskMeAnythingSetTimeout struct {
	PreviousTimeout *big.Int
	NewTimeout      *big.Int
}

// AskMeAnythingQueryResponseReady represents a QueryResponseReady event raised by the AskMeAnything contract.
type AskMeAnythingQueryResponseReady struct {
	QueryId *big.Int
	Result  string
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

// AskMeAnythingRequestSent represents a RequestSent event raised by the AskMeAnything contract.types
type AskMeAnythingRequestSent struct {
	InternalSerial uint8
	Succ           bool
	RequestId      *big.Int
	Tx             string
	BlockN         uint64
	Removed        bool
	Raw            types.Log
}

// AskMeAnythingRandomReady represents a RandomReady event raised by the AskMeAnything contract.
type AskMeAnythingRandomReady struct {
	RequestId       *big.Int
	GeneratedRandom *big.Int
	Tx              string
	BlockN          uint64
	Removed         bool
	Raw             types.Log
}
