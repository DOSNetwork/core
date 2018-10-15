package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

//TODO: MAKE IT A UNIVERSAL INTERFACE FOR ALL KIND OF CHAINS IN FUTURE
//TODO: NEED TO MOVE IT DIRECTLY INSIDE BLOCKCHAIN FOLDER
type DOSProxyLogUrl struct {
	QueryId         *big.Int
	Url             string
	Timeout         *big.Int
	Randomness      *big.Int
	DispatchedGroup [4]*big.Int
}

type DOSProxyLogNonSupportedType struct {
	QueryType string
}

type DOSProxyLogNonContractCall struct {
	From common.Address
}

type DOSProxyLogCallbackTriggeredFor struct {
	CallbackAddr common.Address
}

type DOSProxyLogQueryFromNonExistentUC struct{}

type DOSProxyLogUpdateRandom struct {
	LastRandomness   *big.Int
	LastUpdatedBlock *big.Int
	DispatchedGroup  [4]*big.Int
}

type DOSProxyLogValidationResult struct {
	TrafficType uint8
	TrafficId   *big.Int
	data        []byte
	Signature   [2]*big.Int
	PubKey      [4]*big.Int
	Pass        bool
}

type DOSProxyLogInsufficientGroupNumber struct{}

type DOSProxyLogGrouping struct {
	NodeId []*big.Int
}

type DOSProxyLogPublicKeyAccepted struct {
	X1 *big.Int
	X2 *big.Int
	Y1 *big.Int
	Y2 *big.Int
}
