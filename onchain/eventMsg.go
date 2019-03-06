package onchain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//TODO: MAKE IT A UNIVERSAL INTERFACE FOR ALL KIND OF CHAINS IN FUTURE
//TODO: NEED TO MOVE IT DIRECTLY INSIDE BLOCKCHAIN FOLDER
type DOSProxyLogUrl struct {
	QueryId           *big.Int
	Timeout           *big.Int
	DataSource        string
	Selector          string
	Randomness        *big.Int
	DispatchedGroupId *big.Int
	DispatchedGroup   [4]*big.Int
	Tx                string
	BlockN            uint64
	Removed           bool
}

type DOSProxyLogRequestUserRandom struct {
	RequestId            *big.Int
	LastSystemRandomness *big.Int
	UserSeed             *big.Int
	DispatchedGroupId    *big.Int
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
	LastRandomness    *big.Int
	DispatchedGroupId *big.Int
	DispatchedGroup   [4]*big.Int
	Tx                string
	BlockN            uint64
	Removed           bool
}

type DOSProxyLogValidationResult struct {
	TrafficType uint8
	TrafficId   *big.Int
	Message     []byte
	Signature   [2]*big.Int
	PubKey      [4]*big.Int
	GroupId     *big.Int
	Pass        bool
	Version     uint8
	Tx          string
	BlockN      uint64
	Removed     bool
}

type DOSProxyLogInsufficientWorkingGroup struct {
	NumWorkingGroups *big.Int
	Tx               string
	BlockN           uint64
	Removed          bool
}

type DOSProxyLogInsufficientPendingNode struct {
	NumPendingNodes *big.Int
	Tx              string
	BlockN          uint64
	Removed         bool
}

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
	Removed bool
	BlockN  uint64
	Tx      string
}

type DOSProxyLogPublicKeyAccepted struct {
	GroupId          *big.Int
	PubKey           [4]*big.Int
	WorkingGroupSize *big.Int
	Removed          bool
	BlockN           uint64
	Tx               string
}

type DOSProxyLogPublicKeyUploaded struct {
	GroupId   *big.Int
	PubKey    [4]*big.Int
	Count     *big.Int
	GroupSize *big.Int
	Removed   bool
	BlockN    uint64
	Tx        string
}

type DOSProxyLogGroupDismiss struct {
	GroupId *big.Int
	PubKey  [4]*big.Int
	Removed bool
	BlockN  uint64
	Tx      string
}
