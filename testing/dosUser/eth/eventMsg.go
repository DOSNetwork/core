package eth

import (
	"math/big"
)

//TODO: MAKE IT A UNIVERSAL INTERFACE FOR ALL KIND OF CHAINS IN FUTURE
//TODO: NEED TO MOVE IT DIRECTLY INSIDE BLOCKCHAIN FOLDER
type AskMeAnythingSetTimeout struct {
	PreviousTimeout *big.Int
	NewTimeout      *big.Int
}

type AskMeAnythingQueryResponseReady struct {
	QueryId *big.Int
	Result  string
	Tx      string
	BlockN  uint64
	Removed bool
}

type AskMeAnythingRequestSent struct {
	InternalSerial uint8
	Succ           bool
	RequestId      *big.Int
	Tx             string
	BlockN         uint64
	Removed        bool
}

type AskMeAnythingRandomReady struct {
	RequestId       *big.Int
	GeneratedRandom *big.Int
	Tx              string
	BlockN          uint64
	Removed         bool
}
