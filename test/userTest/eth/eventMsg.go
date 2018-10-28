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

type AskMeAnythingCallbackReady struct {
	QueryId      *big.Int
	Result       string
	RandomNumber *big.Int
}

type AskMeAnythingQuerySent struct {
	TrafficType uint8
	Succ        bool
	QueryId     *big.Int
}
