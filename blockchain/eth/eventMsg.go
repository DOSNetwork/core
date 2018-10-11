package eth

import (
	"math/big"
)

//TODO: MAKE IT A UNIVERSAL INTERFACE FOR ALL KIND OF CHAINS IN FUTURE
//TODO: NEED TO MOVE IT DIRECTLY INSIDE BLOCKCHAIN FOLDER
type DOSProxyLogUrl struct {
	QueryId         *big.Int
	Url             string
	Timeout         *big.Int
	DispatchedGroup [4]*big.Int
}

type DOSProxyLogUpdateRandom struct {
	LastRandomness  [32]byte
	LastBlknum      *big.Int
	DispatchedGroup [4]*big.Int
}

type DOSProxyLogGrouping struct {
	NodeId []*big.Int
}
