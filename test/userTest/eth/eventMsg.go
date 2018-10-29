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
}

type AskMeAnythingRequestSent struct {
	Succ      bool
	RequestId *big.Int
}

type AskMeAnythingRandomReady struct {
	GeneratedRandom *big.Int
}
