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
	Removed         bool
}

type DOSProxyLogRequestUserRandom struct {
	RequestId            *big.Int
	LastSystemRandomness *big.Int
	UserSeed             *big.Int
	DispatchedGroup      [4]*big.Int
	Tx                   string
	BlockN               uint64
	Removed              bool
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
	Removed         bool
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
	Removed     bool
}

type DOSProxyLogInsufficientGroupNumber struct{}

type DOSProxyLogGrouping struct {
	GroupId *big.Int
	NodeId  []common.Address
	Removed bool
	BlockN  uint64
	Tx      string
}

type DOSProxyLogDuplicatePubKey struct {
	GroupId *big.Int
	PubKey  [4]*big.Int
}

type DOSProxyLogAddressNotFound struct {
	GroupId *big.Int
	PubKey  [4]*big.Int
}

type DOSProxyLogPublicKeyAccepted struct {
	GroupId *big.Int
	PubKey  [4]*big.Int
}

type DOSProxyLogPublicKeyUploaded struct {
	GroupId *big.Int
	PubKey  [4]*big.Int
}

type DOSProxyLogGroupDismiss struct {
	PubKey  [4]*big.Int
	Removed bool
	BlockN  uint64
	Tx      string
}

type DOSProxyWhitelistAddressTransferred struct {
	Previous common.Address
	Curr     common.Address
}
