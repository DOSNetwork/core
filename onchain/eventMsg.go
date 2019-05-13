package onchain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//TODO: MAKE IT A UNIVERSAL INTERFACE FOR ALL KIND OF CHAINS IN FUTURE
//TODO: NEED TO MOVE IT DIRECTLY INSIDE BLOCKCHAIN FOLDER

type DosproxyLogUrl struct {
	QueryId           *big.Int
	Timeout           *big.Int
	DataSource        string
	Selector          string
	Randomness        *big.Int
	DispatchedGroupId *big.Int
	Tx                string
	BlockN            uint64
	Removed           bool
	Raw               types.Log
}

type DosproxyLogRequestUserRandom struct {
	RequestId            *big.Int
	LastSystemRandomness *big.Int
	UserSeed             *big.Int
	DispatchedGroupId    *big.Int
	Tx                   string
	BlockN               uint64
	Removed              bool
	Raw                  types.Log
}

type DosproxyLogNonSupportedType struct {
	InvalidSelector string
}

type DosproxyLogNonContractCall struct {
	From common.Address
}

type DosproxyLogCallbackTriggeredFor struct {
	CallbackAddr common.Address
}

type DosproxyLogRequestFromNonExistentUC struct{}

type DosproxyLogUpdateRandom struct {
	LastRandomness    *big.Int
	DispatchedGroupId *big.Int
	Tx                string
	BlockN            uint64
	Removed           bool
	Raw               types.Log
}

type DosproxyLogValidationResult struct {
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

type DosproxyLogGroupingInitiated struct {
	NumPendingNodes   *big.Int
	GroupSize         *big.Int
	GroupingThreshold *big.Int
	Tx                string
	BlockN            uint64
	Removed           bool
	Raw               types.Log
}

type DosproxyLogInsufficientWorkingGroup struct {
	NumWorkingGroups *big.Int
	Tx               string
	BlockN           uint64
	Removed          bool
	Raw              types.Log
}

type DosproxyLogInsufficientPendingNode struct {
	NumPendingNodes *big.Int
	Tx              string
	BlockN          uint64
	Removed         bool
	Raw             types.Log
}

type DosproxyLogGrouping struct {
	GroupId *big.Int
	NodeId  []common.Address
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DosproxyLogDuplicatePubKey struct {
	GroupId *big.Int
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DosproxyLogAddressNotFound struct {
	GroupId *big.Int
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DosproxyLogPublicKeyAccepted struct {
	GroupId          *big.Int
	WorkingGroupSize *big.Int
	Tx               string
	BlockN           uint64
	Removed          bool
	Raw              types.Log
}

type DosproxyLogPublicKeySuggested struct {
	GroupId   *big.Int
	Count     *big.Int
	GroupSize *big.Int
	Tx        string
	BlockN    uint64
	Removed   bool
	Raw       types.Log
}

type DosproxyLogGroupDissolve struct {
	GroupId *big.Int
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type DosproxyLogNoWorkingGroup struct {
	Raw     types.Log
	Removed bool
	BlockN  uint64
	Tx      string
}

type DosproxyUpdateGroupToPick struct {
	OldNum  *big.Int
	NewNum  *big.Int
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type LogStartCommitReveal struct {
	Cid             *big.Int
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	Tx              string
	BlockN          uint64
	Removed         bool
	Raw             types.Log
}

type LogCommit struct {
	Cid        *big.Int
	From       common.Address
	Commitment [32]byte
	Tx         string
	BlockN     uint64
	Removed    bool
	Raw        types.Log
}

type LogReveal struct {
	Cid     *big.Int
	From    common.Address
	Secret  *big.Int
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}

type LogRandom struct {
	Cid     *big.Int
	Random  *big.Int
	Tx      string
	BlockN  uint64
	Removed bool
	Raw     types.Log
}
