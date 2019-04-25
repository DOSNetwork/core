package onchain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	Raw               types.Log
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
	Raw                  types.Log
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
	Raw               types.Log
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
	Raw         types.Log
}

type DOSProxyLogGroupingInitiated struct {
	NumPendingNodes   *big.Int
	GroupSize         *big.Int
	GroupingThreshold *big.Int
	Tx                string
	BlockN            uint64
	Removed           bool
	Raw               types.Log
}

type DOSProxyLogInsufficientWorkingGroup struct {
	NumWorkingGroups *big.Int
	Tx               string
	BlockN           uint64
	Removed          bool
	Raw              types.Log
}

type DOSProxyLogInsufficientPendingNode struct {
	NumPendingNodes *big.Int
	Tx              string
	BlockN          uint64
	Removed         bool
	Raw             types.Log
}

type DOSProxyLogGrouping struct {
	GroupId *big.Int
	NodeId  []common.Address
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DOSProxyLogDuplicatePubKey struct {
	GroupId *big.Int
	PubKey  [4]*big.Int
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DOSProxyLogAddressNotFound struct {
	GroupId *big.Int
	PubKey  [4]*big.Int
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DOSProxyLogPublicKeyAccepted struct {
	GroupId          *big.Int
	PubKey           [4]*big.Int
	WorkingGroupSize *big.Int
	Tx               string
	BlockN           uint64
	Removed          bool
	Raw              types.Log
}

type DOSProxyLogPublicKeySuggested struct {
	GroupId   *big.Int
	PubKey    [4]*big.Int
	Count     *big.Int
	GroupSize *big.Int
	Tx        string
	BlockN    uint64
	Removed   bool
	Raw       types.Log
}

type DOSProxyLogGroupDissolve struct {
	GroupId *big.Int
	PubKey  [4]*big.Int
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DOSProxyLogNoWorkingGroup struct {
	Raw     types.Log
	Removed bool
	BlockN  uint64
	Tx      string
}

type DOSProxyUpdateGroupToPick struct {
	OldNum  *big.Int
	NewNum  *big.Int
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DOSCommitRevealLogStartCommitReveal struct {
	Tx             string
	TargetBlkNum   *big.Int
	CommitDuration *big.Int
	RevealDuration *big.Int
	BlockN         uint64
	Removed        bool
	Raw            types.Log
}

type DOSCommitRevealLogCommit struct {
	Tx         string
	From       common.Address
	Commitment [32]byte
	BlockN     uint64
	Removed    bool
	Raw        types.Log
}

type DOSCommitRevealLogReveal struct {
	Tx      string
	From    common.Address
	Secret  *big.Int
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DOSCommitRevealLogRandom struct {
	Tx      string
	Random  *big.Int
	BlockN  uint64
	Removed bool
	Raw     types.Log
}
