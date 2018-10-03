package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

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

type DOSProxyLogSuccPubKeySub struct{}

type DOSProxyLogUrl struct {
	QueryId *big.Int
	Url     string
	Timeout *big.Int
}
