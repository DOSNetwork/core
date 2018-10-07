package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

//TODO: MAKE IT A UNIVERSAL INTERFACE FOR ALL KIND OF CHAINS IN FUTURE
//TODO: NEED TO MOVE IT DIRECTLY INSIDE BLOCKCHAIN FOLDER

type DOSProxyLogCallbackTriggeredFor struct {
	UserContractAddr common.Address
	Result           []byte
}

type DOSProxyLogInvalidSignature struct{}

type DOSProxyLogNonContractCall struct {
	From common.Address
}

type DOSProxyLogNonSupportedType struct {
	QueryType string
}

type DOSProxyLogQueryFromNonExistentUC struct{}

type DOSProxyLogUrl struct {
	QueryId *big.Int
	Url     string
	Timeout *big.Int
}

type DOSProxyLogGrouping struct {
	GroupId *big.Int
	NodeId  []*big.Int
}

type DOSProxyLogUpdateRandom struct {
	randomId 		*big.Int
	groupId  		*big.Int
	preRandomNumber *big.Int
}

type DOSProxyLogInsufficientGroupNumber struct{}
