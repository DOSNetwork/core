package onchain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//TODO: MAKE IT A UNIVERSAL INTERFACE FOR ALL KIND OF CHAINS IN FUTURE
//TODO: NEED TO MOVE IT DIRECTLY INSIDE BLOCKCHAIN FOLDER
type DOSProxyLogUrl struct {
	QueryId         *big.Int
	Timeout         *big.Int
	DataSource      string
	Selector        string
	Randomness      *big.Int
	DispatchedGroup [4]*big.Int
	Tx              string
	BlockN          uint64
}

type DOSProxyLogRequestUserRandom struct {
	RequestId            *big.Int
	LastSystemRandomness *big.Int
	UserSeed             *big.Int
	DispatchedGroup      [4]*big.Int
	Tx                   string
	BlockN               uint64
}

type DOSProxyLogNonSupportedType struct {
	InvalidSelector string
}

type DOSProxyLogNonContractCall struct {
	From common.Address
}

type DOSProxyLogCallbackTriggeredFor struct {
	CallbackAddr common.Address
}

type DOSProxyLogRequestFromNonExistentUC struct{}

type DOSProxyLogUpdateRandom struct {
	LastRandomness  *big.Int
	DispatchedGroup [4]*big.Int
	Tx              string
	BlockN          uint64
}

type DOSProxyLogValidationResult struct {
	TrafficType uint8
	TrafficId   *big.Int
	Message     []byte
	Signature   [2]*big.Int
	PubKey      [4]*big.Int
	Pass        bool
	Version     uint8
	Tx          string
	BlockN      uint64
}

type DOSProxyLogInsufficientGroupNumber struct{}

type DOSProxyLogGrouping struct {
	NodeId []common.Address
}

type DOSProxyLogDuplicatePubKey struct {
	PubKey [4]*big.Int
}

type DOSProxyLogAddressNotFound struct {
	PubKey [4]*big.Int
}

type DOSProxyLogPublicKeyAccepted struct {
	PubKey [4]*big.Int
}

type DOSProxyWhitelistAddressTransferred struct {
	Previous common.Address
	Curr     common.Address
}
