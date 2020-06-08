// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dosproxy

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DosproxyABI is the input ABI used to generate the binding from.
const DosproxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyFundsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyFundsTokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blkNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callbackAddr\",\"type\":\"address\"}],\"name\":\"LogCallbackTriggeredFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"LogGroupDissolve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nodeId\",\"type\":\"address[]\"}],\"name\":\"LogGrouping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingNodePool\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupsize\",\"type\":\"uint256\"}],\"name\":\"LogGroupingInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numPendingNodes\",\"type\":\"uint256\"}],\"name\":\"LogInsufficientPendingNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numWorkingGroups\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numPendingGroups\",\"type\":\"uint256\"}],\"name\":\"LogInsufficientWorkingGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"LogMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"LogNoPendingGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"LogNonContractCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"invalidSelector\",\"type\":\"string\"}],\"name\":\"LogNonSupportedType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"LogPendingGroupRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"pubKey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numWorkingGroups\",\"type\":\"uint256\"}],\"name\":\"LogPublicKeyAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubKeyCount\",\"type\":\"uint256\"}],\"name\":\"LogPublicKeySuggested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"LogRegisteredNewPendingNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogRequestFromNonExistentUC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastSystemRandomness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userSeed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dispatchedGroupId\",\"type\":\"uint256\"}],\"name\":\"LogRequestUserRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"unregisterFrom\",\"type\":\"uint8\"}],\"name\":\"LogUnRegisteredNewPendingNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastRandomness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dispatchedGroupId\",\"type\":\"uint256\"}],\"name\":\"LogUpdateRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"queryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dataSource\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dispatchedGroupId\",\"type\":\"uint256\"}],\"name\":\"LogUrl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"trafficType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"trafficId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"pubKey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pass\",\"type\":\"bool\"}],\"name\":\"LogValidationResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"UpdateBootstrapCommitDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSize\",\"type\":\"uint256\"}],\"name\":\"UpdateBootstrapGroups\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"UpdateBootstrapRevealDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"UpdateGroupMaturityPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSize\",\"type\":\"uint256\"}],\"name\":\"UpdateGroupSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lifeDiversity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDiversity\",\"type\":\"uint256\"}],\"name\":\"UpdateLifeDiversity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLifeBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLifeBlocks\",\"type\":\"uint256\"}],\"name\":\"UpdatePendingGroupMaxLife\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFundAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFundAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTokenAddr\",\"type\":\"address\"}],\"name\":\"UpdateProxyFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateSystemRandomHardLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"UpdatebootstrapStartThreshold\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rndSeed\",\"type\":\"uint256\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addToGuardianList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressBridge\",\"outputs\":[{\"internalType\":\"contractDOSAddressBridgeInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapCommitDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapEndBlk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapRevealDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapStartThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkExpireLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"expiredWorkingGroupIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"expiredWorkingGroupIdsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExpiredWorkingGroupSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getGroupPubKey\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastHandledGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"getWorkingGroupById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWorkingGroupSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"groupMaturityPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"groupSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"groupToPick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"guardianListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initBlkN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lifeDiversity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeToGroupIdList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numPendingGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numPendingNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingGroupList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingGroupMaxLife\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingGroupTail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlkNum\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingNodeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingNodeTail\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxyFundsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxyFundsTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dataSource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"refreshSystemRandomHardLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"suggestedPubKey\",\"type\":\"uint256[4]\"}],\"name\":\"registerGroupPubKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerNewNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"removeFromGuardianList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userSeed\",\"type\":\"uint256\"}],\"name\":\"requestRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"setBootstrapStartThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"setGroupMaturityPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSize\",\"type\":\"uint256\"}],\"name\":\"setGroupSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDiversity\",\"type\":\"uint256\"}],\"name\":\"setLifeDiversity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLife\",\"type\":\"uint256\"}],\"name\":\"setPendingGroupMaxLife\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFund\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFundToken\",\"type\":\"address\"}],\"name\":\"setProxyFund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"setSystemRandomHardLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cid\",\"type\":\"uint256\"}],\"name\":\"signalBootstrap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signalGroupDissolve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signalGroupFormation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signalRandom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"signalUnregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"trafficType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"sig\",\"type\":\"uint256[2]\"}],\"name\":\"triggerCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unregisterNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"sig\",\"type\":\"uint256[2]\"}],\"name\":\"updateRandomness\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"workingGroupIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"workingGroupIdsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DosproxyBin is the compiled bytecode used for deploying new contracts.
const DosproxyBin = `60806040526105a0600455619d8060055561438060065560326007556003600881905560026009908155600a9182556028600b819055600c55600d556000600e819055600f819055601b91909155601e553480156200005d57600080fd5b506040516200623d3803806200623d833981810160405260608110156200008357600080fd5b50805160208083015160409384015160008054336001600160a01b0319918216178255436001908155918290527f4155c2f711f2cdd34f8262ab8fb9b7020a700fe7b6948222152f7670d1fdf34d805482168317905560148054821683179055601c85527f6de76108811faf2f94afbe5ac6c98e8393206cd093932de1fbfd61bbeec43a02829055601d919091556010805482166001600160a01b038088169190911791829055601180548416828716179055601280549093168185161790925586516313a4cbcb60e31b81529651959693959294911692639d265e589260048281019392829003018186803b1580156200017d57600080fd5b505afa15801562000192573d6000803e3d6000fd5b505050506040513d6020811015620001a957600080fd5b50516011546012546040805163b73a3f8f60e01b81526001600160a01b039384166004820152918316602483015251919092169163b73a3f8f91604480830192600092919082900301818387803b1580156200020457600080fd5b505af115801562000219573d6000803e3d6000fd5b5050505050505061600d80620002306000396000f3fe608060405234801561001057600080fd5b50600436106103c55760003560e01c8063830687c4116101ff578063b836ccea1161011a578063df37c617116100ad578063f2fde38b1161007c578063f2fde38b14610a5c578063f41a158714610a82578063f90ce5ba14610a8a578063fc84dde414610a92576103c5565b8063df37c61714610a44578063ef112dfc14610a4c578063efde068b14610756578063f2a3072d14610a54576103c5565b8063c7c3f4a5116100e9578063c7c3f4a5146109b4578063d18c81b7146109e0578063d79351b214610a16578063dd6ceddf14610a3c576103c5565b8063b836ccea14610947578063b9424b3514610969578063c457aa8f14610971578063c58ebe1c1461098e576103c5565b8063962ba8a411610192578063aeb3da7311610161578063aeb3da731461084b578063b45ef79d14610853578063b53722641461057d578063b7fb8fd714610870576103c5565b8063962ba8a4146107f857806399ca2d3014610800578063a54fb00e14610808578063a60b007d14610825576103c5565b80638f32d59b116101ce5780638f32d59b14610776578063920216531461077e578063925fc6c9146107d357806395071cf6146107f0576103c5565b8063830687c41461075657806385ed42231461075e578063863bc0a1146107665780638da5cb5b1461076e576103c5565b8063372a53cc116102ef5780635d3812041161028257806374ad3a061161025157806374ad3a06146106a257806376cffa53146107205780637c1cf083146107285780637c48d1a01461074e576103c5565b80635d3812041461066d57806363b635ea1461068a5780636e5454d314610692578063715018a61461069a576103c5565b80634a4b52b4116102be5780634a4b52b414610623578063559ea9de1461062b5780635be6c3af146106485780635c0e159f14610650576103c5565b8063372a53cc146105c85780633d385cf5146105d057806340e4a5af146105d85780634a28a74d14610606576103c5565b8063100063ec1161036757806318a1908d1161033657806318a1908d1461058d578063190ca29e146105b057806319717203146105b857806331bf6464146105c0576103c5565b8063100063ec1461055857806311bc53111461057557806311db65741461057d578063155fa82c14610585576103c5565b8063094c3612116103a3578063094c3612146104cd57806309ac86d3146104f15780630dfc09cb1461050f5780630eeee5c11461052c576103c5565b806302957d53146103ca5780630434ccd21461047957806309011cb914610493575b600080fd5b6103e7600480360360208110156103e057600080fd5b5035610a9a565b6040518581526020810185608080838360005b838110156104125781810151838201526020016103fa565b5050505090500184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015610461578181015183820152602001610449565b50505050905001965050505050505060405180910390f35b610481610b4f565b60408051918252519081900360200190f35b6104b9600480360360208110156104a957600080fd5b50356001600160a01b0316610b55565b604080519115158252519081900360200190f35b6104d5610b6a565b604080516001600160a01b039092168252519081900360200190f35b61050d6004803603604081101561050757600080fd5b50610b79565b005b61050d6004803603602081101561052557600080fd5b5035610eb6565b6104816004803603604081101561054257600080fd5b506001600160a01b038135169060200135610f5c565b61050d6004803603602081101561056e57600080fd5b5035610f79565b610481611011565b610481611017565b61050d61101e565b61050d600480360360408110156105a357600080fd5b5080359060200135611179565b610481611427565b61048161142d565b610481611433565b610481611439565b6104b961143f565b61050d600480360360408110156105ee57600080fd5b506001600160a01b038135811691602001351661157a565b61050d6004803603602081101561061c57600080fd5b503561179e565b6103e7611841565b61050d6004803603602081101561064157600080fd5b50356118d0565b61050d611973565b61050d6004803603602081101561066657600080fd5b5035611b18565b6104816004803603602081101561068357600080fd5b5035611f09565b610481611f27565b610481611f2d565b61050d611f33565b61050d600480360360a08110156106b857600080fd5b81359160ff602082013516918101906060810160408201356401000000008111156106e257600080fd5b8201836020820111156106f457600080fd5b8035906020019184600183028401116401000000008311171561071657600080fd5b9193509150611f8c565b6104d561268f565b61050d6004803603602081101561073e57600080fd5b50356001600160a01b031661269e565b6104816127cf565b6104816127d5565b6104816127db565b6104816127e1565b6104d56127e7565b6104b96127f6565b61079b6004803603602081101561079457600080fd5b5035612807565b6040518082608080838360005b838110156107c05781810151838201526020016107a8565b5050505090500191505060405180910390f35b61050d600480360360208110156107e957600080fd5b503561288c565b61048161292f565b610481612935565b6104d561293b565b6104816004803603602081101561081e57600080fd5b503561294a565b6104d56004803603602081101561083b57600080fd5b50356001600160a01b031661295c565b61050d612977565b6104816004803603602081101561086957600080fd5b5035612c32565b6104816004803603608081101561088657600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156108b657600080fd5b8201836020820111156108c857600080fd5b803590602001918460018302840111640100000000831117156108ea57600080fd5b91939092909160208101903564010000000081111561090857600080fd5b82018360208201111561091a57600080fd5b8035906020019184600183028401116401000000008311171561093c57600080fd5b509092509050612c3f565b61050d600480360360a081101561095d57600080fd5b5080359060200161330d565b61050d613881565b61050d6004803603602081101561098757600080fd5b50356139e9565b61050d600480360360208110156109a457600080fd5b50356001600160a01b0316613a8c565b610481600480360360408110156109ca57600080fd5b506001600160a01b038135169060200135613abe565b6109fd600480360360208110156109f657600080fd5b5035613f6c565b6040805192835260208301919091528051918290030190f35b61050d60048036036020811015610a2c57600080fd5b50356001600160a01b0316613f85565b610481613fba565b6104d5613fc0565b610481613fcf565b610481613fd5565b61050d60048036036020811015610a7257600080fd5b50356001600160a01b0316613fdb565b610481613ff5565b610481613ffb565b610481614001565b6000610aa4615bd2565b6000838152601760205260408120548190606090610ac187612807565b6000888152601760209081526040918290206005810154600682015460079092018054855181860281018601909652808652919492939092918391830182828015610b3557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610b17575b505050505090509450945094509450945091939590929450565b60095481565b60296020526000908152604090205460ff1681565b6014546001600160a01b031681565b601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610bc757600080fd5b505afa158015610bdb573d6000803e3d6000fd5b505050506040513d6020811015610bf157600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b158015610c3b57600080fd5b505afa158015610c4f573d6000803e3d6000fd5b505050506040513d6020811015610c6557600080fd5b5051610caf576040805162461bcd60e51b8152602060048201526014602482015273496e76616c6964207374616b696e67206e6f646560601b604482015290519081900360640190fd5b610d4c6000602054610cc081614007565b604080518082018252863581526020808801359082015281516080810180845291929091602291839190820190839060029082845b815481526020019060010190808311610cf557505050918352505060408051808201918290526020909201919060028481019182845b815481526020019060010190808311610d2b57505050505081525050614031565b610d5557610eb3565b602080546040805184358185015284840135818301528151808203830181526060820180845281519186019190912085556010546313a4cbcb60e31b909152915192936001600160a01b0390921692639d265e5892606480840193919291829003018186803b158015610dc757600080fd5b505afa158015610ddb573d6000803e3d6000fd5b505050506040513d6020811015610df157600080fd5b5051604051632100fdf760e21b815260048101838152336024830181905260606044840190815260288054606486018190526001600160a01b0390961695638403f7dc95889592939160849091019084908015610e7757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610e59575b5050945050505050600060405180830381600087803b158015610e9957600080fd5b505af1158015610ead573d6000803e3d6000fd5b50505050505b50565b610ebe6127f6565b610ec757600080fd5b600a548114158015610edb57506002810615155b610f1a576040805162461bcd60e51b81526020600482015260156024820152600080516020615f2e833981519152604482015290519081900360640190fd5b600a54604080519182526020820183905280517f28eb4f48ae7c6c17a714b104832bdd949ebd0a984d37f4893d6cb91f92a8ae579281900390910190a1600a55565b601660209081526000928352604080842090915290825290205481565b610f816127f6565b610f8a57600080fd5b600d54811415610fcf576040805162461bcd60e51b81526020600482015260156024820152600080516020615f2e833981519152604482015290519081900360640190fd5b600d54604080519182526020820183905280517f1fa02fb08d726e79971d6de0ee1e2f637f068fed6d3fb859a1765e666bb193079281900390910190a1600d55565b600d5481565b6018545b90565b611026614358565b1561113357604080514381523360208201528151600080516020615e57833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b1580156110a157600080fd5b505afa1580156110b5573d6000803e3d6000fd5b505050506040513d60208110156110cb57600080fd5b5051604080516323ff34cb60e01b815233600482015290516001600160a01b03909216916323ff34cb9160248082019260009290919082900301818387803b15801561111657600080fd5b505af115801561112a573d6000803e3d6000fd5b50505050611177565b6040805160208082526012908201527127379033b937bab8103337b936b0ba34b7b760711b818301529051600080516020615e778339815191529181900360600190a15b565b3330146111cd576040805162461bcd60e51b815260206004820152601860248201527f556e61757468656e7469636174656420726573706f6e73650000000000000000604482015290519081900360640190fd5b6009546019541015611226576040805162461bcd60e51b815260206004820152601f60248201527f4e6f20656e6f756768206578706972656420776f726b696e672067726f757000604482015290519081900360640190fd5b600a5460155410156112695760405162461bcd60e51b8152600401808060200182810382526021815260200180615e976021913960400191505060405180910390fd5b6000600954600101600a540290506060816040519080825280602002602001820160405280156112a3578160200160208202803883390190505b50905060005b6009548110156113f4576019546040805160208082018890528183018990526060808301869052835180840390910181526080909201909252805191012060009190816112f257fe5b0690506000601760006019848154811061130857fe5b90600052602060002001548152602001908152602001600020905060008090505b600a548110156113925781600701818154811061134257fe5b9060005260206000200160009054906101000a90046001600160a01b03168582600a548702018151811061137257fe5b6001600160a01b0390921660209283029190910190910152600101611329565b5080546113a0906000614671565b6019805460001981019081106113b257fe5b9060005260206000200154601983815481106113ca57fe5b60009182526020909120015560198054906113e9906000198301615bf0565b5050506001016112a9565b50611408600a54600954600a5402836147b2565b6114128184614871565b61142181600954600101614978565b50505050565b601d5481565b600f5481565b60085481565b600b5481565b60105460408051630e9ed68b60e01b815290516000926001600160a01b031691630e9ed68b916004808301926020929190829003018186803b15801561148457600080fd5b505afa158015611498573d6000803e3d6000fd5b505050506040513d60208110156114ae57600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b1580156114f857600080fd5b505afa15801561150c573d6000803e3d6000fd5b505050506040513d602081101561152257600080fd5b505161156c576040805162461bcd60e51b8152602060048201526014602482015273496e76616c6964207374616b696e67206e6f646560601b604482015290519081900360640190fd5b61157533614c72565b905090565b6115826127f6565b61158b57600080fd5b6011546001600160a01b038381169116148015906115b157506001600160a01b03821615155b6115f0576040805162461bcd60e51b81526020600482015260156024820152600080516020615f2e833981519152604482015290519081900360640190fd5b6012546001600160a01b0382811691161480159061161657506001600160a01b03811615155b611655576040805162461bcd60e51b81526020600482015260156024820152600080516020615f2e833981519152604482015290519081900360640190fd5b601154601254604080516001600160a01b039384168152858416602082015291831682820152918316606082015290517f2ae8e7023c1978c8540df9af00881f2f942d6e7233463a3f0def2b6e57e6dd609181900360800190a1601180546001600160a01b038085166001600160a01b031992831617909255601280548484169216919091179055601054604080516313a4cbcb60e31b815290519190921691639d265e58916004808301926020929190829003018186803b15801561171a57600080fd5b505afa15801561172e573d6000803e3d6000fd5b505050506040513d602081101561174457600080fd5b50516011546012546040805163b73a3f8f60e01b81526001600160a01b039384166004820152918316602483015251919092169163b73a3f8f91604480830192600092919082900301818387803b158015610e9957600080fd5b6117a66127f6565b6117af57600080fd5b601b5481141580156117c057508015155b6117ff576040805162461bcd60e51b81526020600482015260156024820152600080516020615f2e833981519152604482015290519081900360640190fd5b601b54604080519182526020820183905280517ffc644126d2177f897a0e09f46bf2678f9577840113d685f4a56bd9e4d48d012c9281900390910190a1601b55565b600061184b615bd2565b602154600090819060609061185f81612807565b60265460275460288054604080516020808402820181019092528281529183918301828280156118b857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161189a575b50505050509050945094509450945094509091929394565b6118d86127f6565b6118e157600080fd5b60065481141580156118f257508015155b611931576040805162461bcd60e51b81526020600482015260156024820152600080516020615f2e833981519152604482015290519081900360640190fd5b600654604080519182526020820183905280517f978a29592cb150802d04222f78a83519049bde42bb2e86e17250efde5820c6879281900390910190a1600655565b60016000819052601c6020527f6de76108811faf2f94afbe5ac6c98e8393206cd093932de1fbfd61bbeec43a02549081148015906119c65750601b546000828152601a6020526040902060010154439101105b15611adc576119d481614fd8565b604080514381523360208201528151600080516020615e57833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015611a4a57600080fd5b505afa158015611a5e573d6000803e3d6000fd5b505050506040513d6020811015611a7457600080fd5b5051604080516323ff34cb60e01b815233600482015290516001600160a01b03909216916323ff34cb9160248082019260009290919082900301818387803b158015611abf57600080fd5b505af1158015611ad3573d6000803e3d6000fd5b50505050610eb3565b600080516020615e77833981519152604051808060200182810382526024815260200180615fb56024913960400191505060405180910390a150565b80600e5414611b67576040805162461bcd60e51b81526020600482015260166024820152754e6f7420696e20626f6f74737472617020706861736560501b604482015290519081900360640190fd5b600f544311611bc357604080516020808252601f908201527f57616974696e6720746f20636f6c6c656374206d6f726520656e74726f707900818301529051600080516020615e778339815191529181900360600190a1610eb3565b600d546015541015611c2257604080516020808252601d908201527f4e6f7420656e6f756768206e6f64657320746f20626f6f747374726170000000818301529051600080516020615e778339815191529181900360600190a1610eb3565b6000600e819055600f819055601054604080516306b810cf60e21b815290516001600160a01b0390921691631ae0433c91600480820192602092909190829003018186803b158015611c7357600080fd5b505afa158015611c87573d6000803e3d6000fd5b505050506040513d6020811015611c9d57600080fd5b505160408051633352da4560e21b81526004810185905290516001600160a01b039092169163cd4b6914916024808201926020929091908290030181600087803b158015611cea57600080fd5b505af1158015611cfe573d6000803e3d6000fd5b505050506040513d6020811015611d1457600080fd5b5051905080611d5c57600080516020615e7783398151915260405180806020018281038252602a815260200180615f04602a913960400191505060405180910390a150610eb3565b60208054604080518084019290925281810184905280518083038201815260609092019052805190820120905543601f55600a54600d5460009190819081611da057fe5b04029050606081604051908082528060200260200182016040528015611dd0578160200160208202803883390190505b509050611ddf826000836147b2565b611deb81602054614871565b611e0081600a548481611dfa57fe5b04614978565b604080514381523360208201528151600080516020615e57833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015611e7657600080fd5b505afa158015611e8a573d6000803e3d6000fd5b505050506040513d6020811015611ea057600080fd5b5051604080516323ff34cb60e01b815233600482015290516001600160a01b03909216916323ff34cb9160248082019260009290919082900301818387803b158015611eeb57600080fd5b505af1158015611eff573d6000803e3d6000fd5b5050505050505050565b60188181548110611f1657fe5b600091825260209091200154905081565b600a5481565b60075481565b611f3b6127f6565b611f4457600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b158015611fda57600080fd5b505afa158015611fee573d6000803e3d6000fd5b505050506040513d602081101561200457600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b15801561204e57600080fd5b505afa158015612062573d6000803e3d6000fd5b505050506040513d602081101561207857600080fd5b50516120c2576040805162461bcd60e51b8152602060048201526014602482015273496e76616c6964207374616b696e67206e6f646560601b604482015290519081900360640190fd5b6000858152600360205260409020600601546001600160a01b031680612111576040517f40d87958cd48e8b698a94f35390a9020a9127528227647da089cc6bfd7931f4290600090a150612688565b6121ef858786868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040805180820190915288358152915050602081018760016020908102919091013590915260008c81526003918290526040908190208151608081018084526002808401805495840195865292959294869490938693910160608601808311610cf557505050918352505060408051808201918290526002848101805483526020948501949293909260038701908501808311610d2b57505050505081525050614031565b6121f95750612688565b604080516001600160a01b038316815290517f065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf09181900360200190a160008681526003602052604081208181556001810182905590600282018161225d8282615c14565b61226b600283016000615c14565b50505060060180546001600160a01b031916905560ff8516600214156123225760408051636d11297760e01b81526004810188815260248201928352604482018690526001600160a01b03841692636d112977928a9289928992606401848480828437600081840152601f19601f820116905080830192505050945050505050600060405180830381600087803b15801561230557600080fd5b505af1158015612319573d6000803e3d6000fd5b5050505061242e565b60ff8516600114156123e157604080516020808252600a90820152695573657252616e646f6d60b01b818301529051600080516020615e778339815191529181900360600190a1604080518335602082810191909152808501358284015282518083038401815260608301808552815191909201206318a1908d60e01b90915260648201899052608482015290516001600160a01b038316916318a1908d9160a480830192600092919082900301818387803b15801561230557600080fd5b6040805162461bcd60e51b815260206004820152601860248201527f556e737570706f72746564207472616666696320747970650000000000000000604482015290519081900360640190fd5b612436615c22565b600087815260036020908152604080832060019081015484526017835292819020815160a081018352815481528251608081018085529195929486019390928501918391820190839060029082845b81548152602001906001019080831161248557505050918352505060408051808201918290526020909201919060028481019182845b8154815260200190600101908083116124bb57505050505081525050815260200160058201548152602001600682015481526020016007820180548060200260200160405190810160405280929190818152602001828054801561254857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161252a575b5050505050815250509050601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b1580156125a157600080fd5b505afa1580156125b5573d6000803e3d6000fd5b505050506040513d60208110156125cb57600080fd5b50516080820151604051632100fdf760e21b8152600481018a815233602483018190526060604484019081528451606485015284516001600160a01b0390961695638403f7dc958e95939490939092916084909101906020858101910280838360005b8381101561264657818101518382015260200161262e565b50505050905001945050505050600060405180830381600087803b15801561266d57600080fd5b505af1158015612681573d6000803e3d6000fd5b5050505050505b5050505050565b6010546001600160a01b031681565b3360009081526029602052604090205460ff16612702576040805162461bcd60e51b815260206004820152601960248201527f4e6f742077686974656c697374656420677561726469616e2100000000000000604482015290519081900360640190fd5b61270b81614c72565b1561278657604080514381523360208201528151600080516020615e57833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015611a4a57600080fd5b604080516020808252601590820152742737ba3434b733903a37903ab73932b3b4b9ba32b960591b818301529051600080516020615e778339815191529181900360600190a150565b60055481565b60195490565b600e5481565b601e5481565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b61280f615bd2565b6000601760006018858154811061282257fe5b90600052602060002001548152602001908152602001600020600101905060405180608001604052808260000160006002811061285b57fe5b015481526020018260010154815260200160028301600001548152602001600283016001015490529150505b919050565b6128946127f6565b61289d57600080fd5b60055481141580156128ae57508015155b6128ed576040805162461bcd60e51b81526020600482015260156024820152600080516020615f2e833981519152604482015290519081900360640190fd5b600554604080519182526020820183905280517f96a027b03aa3233feda42c74f270026db98f223e64b4df4b81231da93bac04b39281900390910190a1600555565b60015481565b60045481565b6012546001600160a01b031681565b601c6020526000908152604090205481565b6013602052600090815260409020546001600160a01b031681565b601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156129c557600080fd5b505afa1580156129d9573d6000803e3d6000fd5b505050506040513d60208110156129ef57600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b158015612a3957600080fd5b505afa158015612a4d573d6000803e3d6000fd5b505050506040513d6020811015612a6357600080fd5b5051612aad576040805162461bcd60e51b8152602060048201526014602482015273496e76616c6964207374616b696e67206e6f646560601b604482015290519081900360640190fd5b336000908152601360205260409020546001600160a01b031615612ad057611177565b3360009081526016602090815260408083206001845290915290205415612af657611177565b3360008181526016602090815260408083206001808552925290912055612b1c90615116565b6040805133815290517f707a6d64786780aac9cd0c5813ea04241eb135ddd2280c06eea6719afb9131519181900360200190a1601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b158015612b9d57600080fd5b505afa158015612bb1573d6000803e3d6000fd5b505050506040513d6020811015612bc757600080fd5b505160408051634c542d3d60e01b815233600482015290516001600160a01b0390921691634c542d3d9160248082019260009290919082900301818387803b158015612c1257600080fd5b505af1158015612c26573d6000803e3d6000fd5b50505050610eb3614358565b60198181548110611f1657fe5b6000866002601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015612c9257600080fd5b505afa158015612ca6573d6000803e3d6000fd5b505050506040513d6020811015612cbc57600080fd5b5051604080516371b281b360e11b81526001600160a01b038581166004830152602482018590529151919092169163e3650366916044808301926020929190829003018186803b158015612d0f57600080fd5b505afa158015612d23573d6000803e3d6000fd5b505050506040513d6020811015612d3957600080fd5b5051612d8c576040805162461bcd60e51b815260206004820181905260248201527f4e6f7420656e6f7567682066656520746f2072657175657374206f7261636c65604482015290519081900360640190fd5b6000612d978a615178565b11156132c057606085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250508251929350509015905080612e07575080600081518110612df057fe5b6020910101516001600160f81b031916600960fa1b145b80612e31575080600081518110612e1a57fe5b6020910101516001600160f81b031916602f60f81b145b156132565760006002600081546001019190508190558b8b8b8b8b8b60405160200180888152602001876001600160a01b03166001600160a01b0316815260200186815260200180602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f82011690508083019250505099505050505050505050506040516020818303038152906040528051906020012060001c90506000612f0160028361517c565b9050600019811415612f5257600080516020615e77833981519152604051808060200182810382526024815260200180615eb86024913960400191505060405180910390a160009550505050613301565b60006017600060188481548110612f6557fe5b6000918252602080832090910154835282810193909352604091820190208151608080820184528782528254948201949094528251938401835290935091828201916001850190829081018260028282826020028201915b815481526020019060010190808311612fbd57505050918352505060408051808201918290526020909201919060028481019182845b815481526020019060010190808311612ff3575050509190925250505081526001600160a01b038f1660209182015260008581526003825260409081902083518155918301516001830155820151805160028084019161305591839190615c57565b50602082015161306b9060028084019190615c57565b50505060608201518160060160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050507f05e1614af4efb13caeba2369a57a05ee5830f33364f82e2c899fd5710cb56ef3838d8d8d8d8d60205488600001546040518089815260200188815260200180602001806020018581526020018481526020018381038352898982818152602001925080828437600083820152601f01601f191690910184810383528781526020019050878780828437600083820152604051601f909101601f19169092018290039c50909a5050505050505050505050a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b15801561319e57600080fd5b505afa1580156131b2573d6000803e3d6000fd5b505050506040513d60208110156131c857600080fd5b50516001600160a01b0316637aa9181b8e8560026040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b031681526020018381526020018281526020019350505050600060405180830381600087803b15801561323257600080fd5b505af1158015613246573d6000803e3d6000fd5b5050505082965050505050613301565b7f70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41868660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a16000935050613301565b604080516001600160a01b038b16815290517f6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb59181900360200190a1600092505b50509695505050505050565b601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561335b57600080fd5b505afa15801561336f573d6000803e3d6000fd5b505050506040513d602081101561338557600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b1580156133cf57600080fd5b505afa1580156133e3573d6000803e3d6000fd5b505050506040513d60208110156133f957600080fd5b5051613443576040805162461bcd60e51b8152602060048201526014602482015273496e76616c6964207374616b696e67206e6f646560601b604482015290519081900360640190fd5b6000828152601a602052604090208054613490576040805184815290517f71047c0893a51085656a2894bba10bc6ef51a654f25e1ead1929b076487a95699181900360200190a15061387d565b3360009081526003820160205260409020546001600160a01b03166134fc576040805162461bcd60e51b815260206004820181905260248201527f4e6f742066726f6d20617574686f72697a65642067726f7570206d656d626572604482015290519081900360640190fd5b6040805183356020808301919091528085013582840152848301356060808401919091528501356080808401919091528351808403909101815260a08301808552815191830191909120600081815260028701909352918490208054600101908190559087905260c083015291517f717e526bce26f8e67908004294b35133bbe2a9c7f611384cb0f484aca9223d019181900360e00190a16002600a54816135a057fe5b60008381526002850160205260409020549190041015611421576060600a546040519080825280602002602001820160405280156135e8578160200160208202803883390190505b5060016000908152600385016020526040812054919250906001600160a01b03165b6001600160a01b038116600114613685578083838060010194508151811061362e57fe5b6001600160a01b03928316602091820292909201810191909152908216600090815260169091526040902061366390886151a6565b6001600160a01b0390811660009081526003860160205260409020541661360a565b6018805460018082019092557fb13d2d76d1f4b7be834882e410b3e3a8afaf69f83600ae24db354391d2378d2e018890556040805160a0810182528981528151608080820184528a358285019081526020808d0135606080860191909152918452855180870187528d8701358152828e01358183015281850152808501938452600654601e540285870152439185019190915290830188905260008c8152601790915292909220815181559151805191939091908301906137499082906002615c57565b50602082015161375f9060028084019190615c57565b505050604082015160058201556060820151600682015560808201518051613791916007840191602090910190615c95565b509050506000806137a3601c8a6151ca565b915091508080156137b5575088601d54145b156137c057601d8290555b6000898152601a6020908152604080832083815560010192909255601e805460001901905581518b815291517f156927b06a61046135669011768a03b2592ee3667374525502b62b0aef4cbd219281900390910190a16018546040518a81527f9ad0d0bfff7f0fc653b03785685d46101b09e3cb1f50081915bc8101662e4a88918b918b91906020810183608080828437600083820152601f01601f191690910192835250506040519081900360200192509050a1505050505050505b5050565b43600454601f540111156138e257604080516020808252601c908201527f53797374656d52616e646f6d206e6f7420657870697265642079657400000000818301529051600080516020615e778339815191529181900360600190a1611177565b6138ea615253565b604080514381523360208201528151600080516020615e57833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b15801561396057600080fd5b505afa158015613974573d6000803e3d6000fd5b505050506040513d602081101561398a57600080fd5b5051604080516323ff34cb60e01b815233600482015290516001600160a01b03909216916323ff34cb9160248082019260009290919082900301818387803b1580156139d557600080fd5b505af1158015611421573d6000803e3d6000fd5b6139f16127f6565b6139fa57600080fd5b6004548114158015613a0b57508015155b613a4a576040805162461bcd60e51b81526020600482015260156024820152600080516020615f2e833981519152604482015290519081900360640190fd5b600454604080519182526020820183905280517fdb95a2fbbee34de5822459ce9edd234f70f321a1cc2395b2dc902b69e1f8093b9281900390910190a1600455565b613a946127f6565b613a9d57600080fd5b6001600160a01b03166000908152602960205260409020805460ff19169055565b6000826001601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015613b1157600080fd5b505afa158015613b25573d6000803e3d6000fd5b505050506040513d6020811015613b3b57600080fd5b5051604080516371b281b360e11b81526001600160a01b038581166004830152602482018590529151919092169163e3650366916044808301926020929190829003018186803b158015613b8e57600080fd5b505afa158015613ba2573d6000803e3d6000fd5b505050506040513d6020811015613bb857600080fd5b5051613c0b576040805162461bcd60e51b815260206004820181905260248201527f4e6f7420656e6f7567682066656520746f2072657175657374206f7261636c65604482015290519081900360640190fd5b60028054600190810191829055604080516020808201949094526001600160a01b038916818301526060808201899052825180830390910181526080909101909152805192019190912090600090613c63908361517c565b9050600019811415613cb357600080516020615e7783398151915260405180806020018281038252602d815260200180615f88602d913960400191505060405180910390a1600094505050613f64565b60006017600060188481548110613cc657fe5b6000918252602080832090910154835282810193909352604091820190208151608080820184528782528254948201949094528251938401835290935091828201916001850190829081018260028282826020028201915b815481526020019060010190808311613d1e57505050918352505060408051808201918290526020909201919060028481019182845b815481526020019060010190808311613d54575050509190925250505081526001600160a01b038a16602091820152600085815260038252604090819020835181559183015160018301558201518051600280840191613db691839190615c57565b506020820151613dcc9060028084019190615c57565b505050606091820151600690910180546001600160a01b039092166001600160a01b031990921691909117905560208054835460408051888152938401929092528282018b90529282019290925290517fd587179d80d44e74955fa5d651db2f31b5470fcee8f9736f07ae3b24456a2cf09181900360800190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015613e9457600080fd5b505afa158015613ea8573d6000803e3d6000fd5b505050506040513d6020811015613ebe57600080fd5b50516001600160a01b0390811690637aa9181b908a163014613ee05789613eed565b6011546001600160a01b03165b8560016040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b031681526020018381526020018281526020019350505050600060405180830381600087803b158015613f4657600080fd5b505af1158015613f5a573d6000803e3d6000fd5b5094975050505050505b505092915050565b601a602052600090815260409020805460019091015482565b613f8d6127f6565b613f9657600080fd5b6001600160a01b03166000908152602960205260409020805460ff19166001179055565b60065481565b6011546001600160a01b031681565b600c5481565b60205481565b613fe36127f6565b613fec57600080fd5b610eb381615468565b60155481565b601f5481565b601b5481565b60408051602080825281830190925260609160208201818038833950505060208101929092525090565b6000606084336040516020018083805190602001908083835b602083106140695780518252601f19909201916020918201910161404a565b5181516020939093036101000a6000190180199091169216919091179052606094851b6bffffffffffffffffffffffff191692019182525060408051808303600b19018152600260148401818152607485019093529096509394509291506034015b6140d3615cf6565b8152602001906001900390816140cb57505060408051600280825260608083019093529293509091602082015b614108615d10565b815260200190600190039081614100579050509050614126866154d6565b8260008151811061413357fe5b60200260200101819052506141478361554f565b8260018151811061415457fe5b6020026020010181905250614167615576565b8160008151811061417457fe5b6020026020010181905250848160018151811061418d57fe5b602002602001018190525060006141a48383615636565b90507fd33c44f7ce166bcd2616c8f4d811261d4d24e1af815f78614683e0c7682c42a58a8a8660405180604001604052808c6000015181526020018c6020015181525060405180608001604052808c6000015160006002811061420357fe5b602002015181526020018c6000015160016002811061421e57fe5b602002015181526020018c6020015160006002811061423957fe5b602002015181526020018c6020015160016002811061425457fe5b602002015181525086604051808760ff1660ff1681526020018681526020018060200185600260200280838360005b8381101561429b578181015183820152602001614283565b5050505090500184600460200280838360005b838110156142c65781810151838201526020016142ae565b5050505090500183151515158152602001828103825286818151815260200191508051906020019080838360005b8381101561430c5781810151838201526020016142f4565b50505050905090810190601f1680156143395780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a19998505050505050505050565b6000600a54601554108061437a575060185415801561437a5750600d54601554105b156143f657601954156143f6576143aa601960008154811061439857fe5b90600052602060002001546001614671565b6019805460001981019081106143bc57fe5b906000526020600020015460196000815481106143d557fe5b60009182526020909120015560198054906143f4906000198301615bf0565b505b600a54601554101561443f5760155460408051918252517fc03848aa1689c7c291dcc68fa62d6109e13f16b25e89bdef5a6c8598b36b80199181900360200190a150600061101b565b601854156144e657600954601954106144a85761445c3043613abe565b507f5ee79c79dd5be1f25f86c7028ec8fb6849839e8a23f676afc907eade45946986601554600a54604051808381526020018281526020019250505060405180910390a150600161101b565b600080516020615e7783398151915260405180806020018281038252603a815260200180615f4e603a913960400191505060405180910390a161466b565b600d546015541061466b57600e5461461c57601060009054906101000a90046001600160a01b03166001600160a01b0316631ae0433c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561454657600080fd5b505afa15801561455a573d6000803e3d6000fd5b505050506040513d602081101561457057600080fd5b5051600b54600c54600d546040805163b917b5a560e01b8152436004820152602481019490945260448401929092526064830152516001600160a01b039092169163b917b5a5916084808201926020929091908290030181600087803b1580156145d957600080fd5b505af11580156145ed573d6000803e3d6000fd5b505050506040513d602081101561460357600080fd5b5051600e5550600c54600b54430101600f55600161101b565b604080516020808252601a908201527f416c726561647920696e20626f6f747374726170207068617365000000000000818301529051600080516020615e778339815191529181900360600190a15b50600090565b6000828152601760205260408120905b600782015481101561472657600082600701828154811061469e57fe5b60009182526020808320909101546001600160a01b0316808352601690915260408220855491935082916146d291906151ca565b915091508080156146e35750600182145b1561471b5785801561470d57506001600160a01b0383811660009081526013602052604090205416155b1561471b5761471b83615116565b505050600101614681565b5060008381526017602052604081208181559060018201816147488282615c14565b614756600283016000615c14565b5050600582016000905560068201600090556007820160006147789190615d35565b50506040805184815290517ff7377b41bdc770cc22a1bad318571f0ae6d65188371bdbe0cb660c0de57b31509181900360200190a1505050565b60005b838110156148485760136020527f4155c2f711f2cdd34f8262ab8fb9b7020a700fe7b6948222152f7670d1fdf34d80546001600160a01b0380821660008181526040902080549092166001600160a01b0319938416179093558054909116905582518190849084870190811061482757fe5b6001600160a01b0390921660209283029190910190910152506001016147b5565b50601580548490039081905561486c57601480546001600160a01b03191660011790555b505050565b8151600019015b801561486c57600081600101838386858151811061489257fe5b602002602001015160405160200180848152602001838152602001826001600160a01b03166001600160a01b031660601b815260140193505050506040516020818303038152906040528051906020012060001c816148ed57fe5b06905060008483815181106148fe57fe5b6020026020010151905084828151811061491457fe5b602002602001015185848151811061492857fe5b60200260200101906001600160a01b031690816001600160a01b0316815250508085838151811061495557fe5b6001600160a01b0390921660209283029190910190910152505060001901614878565b80600a54028251146149d1576040805162461bcd60e51b815260206004820152601a60248201527f63616e646964617465732e6c656e6774682069732077726f6e67000000000000604482015290519081900360640190fd5b6060600a546040519080825280602002602001820160405280156149ff578160200160208202803883390190505b5090506000805b838110156126885760009150815b600a54811015614ac1578581600a5484020181518110614a3057fe5b6020026020010151848281518110614a4457fe5b60200260200101906001600160a01b031690816001600160a01b03168152505082848281518110614a7157fe5b602090810291909101810151604080518084019490945260609190911b6bffffffffffffffffffffffff191683820152805180840360340181526054909301905281519101209250600101614a14565b506040805180820182528381524360208083019182526000868152601a825284812093518455915160018085019190915580835260039093019081905292812080546001600160a01b0319169092179091555b600a54811015614be257600160009081526020839052604081205486516001600160a01b03909116918491889085908110614b4b57fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b03160217905550848181518110614ba357fe5b6020908102919091018101516001600081815292859052604090922080546001600160a01b0319166001600160a01b0390921691909117905501614b14565b50614bec83615839565b7f78bf54a42d1b98e6c809c3e5904898c5b3304ede546b6f070e83d9a32e544d4f83856040518083815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015614c55578181015183820152602001614c3d565b50505050905001935050505060405180910390a150600101614a06565b6001600160a01b03811660009081526016602090815260408083206001845290915281205481808215801590614ca9575060018314155b15614df557614cb9836001614671565b60005b601854811015614d52578360188281548110614cd457fe5b90600052602060002001541415614d4a57601854600019018114614d2957601880546000198101908110614d0457fe5b906000526020600020015460188281548110614d1c57fe5b6000918252602090912001555b6018805490614d3c906000198301615bf0565b506001925090821790614d52565b600101614cbc565b5081614df55760005b601954811015614df3578360198281548110614d7357fe5b90600052602060002001541415614deb57601954600019018114614dc857601980546000198101908110614da357fe5b906000526020600020015460198281548110614dbb57fe5b6000918252602090912001555b6019805490614ddb906000198301615bf0565b5060019250600282179150614df3565b600101614d5b565b505b81158015614e075750614e0785615869565b15614e10576004175b6001600160a01b038581166000908152601360205260409020541615614ea8576000614e3d601387615901565b935090508215614ea657601580546000190190556001600160a01b038087166000818152601660209081526040808320600184529091528120556014549091161415614e9f57601480546001600160a01b0319166001600160a01b0383161790555b6008821791505b505b604080516001600160a01b038716815260ff8316602082015281517faa40dce54d678a8a16fc3cf106c1ddf0b34b66a43c7a365af3268c63bb95fead929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b158015614f3c57600080fd5b505afa158015614f50573d6000803e3d6000fd5b505050506040513d6020811015614f6657600080fd5b50516040805163c5375c2960e01b81526001600160a01b0388811660048301529151919092169163c5375c2991602480830192600092919082900301818387803b158015614fb357600080fd5b505af1158015614fc7573d6000803e3d6000fd5b5050505060ff161515949350505050565b6000818152601a602090815260408083206001845260038101909252909120546001600160a01b03165b6001600160a01b03811660011461508e576001600160a01b0381166000908152601660209081526040808320600180855292529091205414801561505e57506001600160a01b0381811660009081526013602052604090205416155b1561506c5761506c81615116565b6001600160a01b03908116600090815260038301602052604090205416615002565b60008061509c601c866151ca565b915091508080156150ae575084601d54145b156150b957601d8290555b6000858152601a6020908152604080832083815560010192909255601e8054600019019055815187815291517f156927b06a61046135669011768a03b2592ee3667374525502b62b0aef4cbd219281900390910190a15050505050565b601480546001600160a01b039081166000908152601360205260408082205494831680835281832080549685166001600160a01b0319978816179055845490931682529020805484168217905581549092169091179055601580546001019055565b3b90565b600043601f54600454011161519357615193615253565b61519d8383615969565b90505b92915050565b60016000818152602093909352604080842080548486529185209190915592529055565b6001600081815260208490526040812054909182915b600181141580156151f15750848114155b1561520d576000818152602087905260409020549091506151e0565b6001811415615225576001600093509350505061524c565b60008181526020879052604080822080548584529183209190915591815290559150600190505b9250929050565b600061526481600019430140615969565b90506000198114156152af57600080516020615e77833981519152604051808060200182810382526028815260200180615edc6028913960400191505060405180910390a150611177565b43601f8190555060176000601883815481106152c757fe5b600091825260208083209091015483528201929092526040019020805460219081556001820160226152fb81836002615d53565b5061530e60028281019084810190615d53565b5050506005820154816005015560068201548160060155600782018160070190805461533b929190615d7e565b505060208054602154604080519283529282015281517ffaa99731d2c5abc7ee76b2e31b6b7e293a30e1e2274f59396a7a59cabd5eadf393509081900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b1580156153cb57600080fd5b505afa1580156153df573d6000803e3d6000fd5b505050506040513d60208110156153f557600080fd5b505160115460205460408051637aa9181b60e01b81526001600160a01b039384166004820152602481019290925260006044830181905290519290931692637aa9181b9260648084019382900301818387803b15801561545457600080fd5b505af1158015612688573d6000803e3d6000fd5b6001600160a01b03811661547b57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6154de615cf6565b81511580156154ef57506020820151155b156154fb575080612887565b60007f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4790506040518060400160405280846000015181526020018285602001518161554257fe5b0690920390915292915050565b615557615cf6565b8151602083012061556f615569615acd565b82615aee565b9392505050565b61557e615d10565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b82527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa60208381019190915281019190915290565b6000815183511461564657600080fd5b8251604080516006830280825260c084028201602001909252606090828015615679578160200160208202803883390190505b50905060005b838110156157fe5786818151811061569357fe5b6020026020010151600001518282600602600001815181106156b157fe5b6020026020010181815250508681815181106156c957fe5b6020026020010151602001518282600602600101815181106156e757fe5b6020026020010181815250508581815181106156ff57fe5b60209081029190910101515151825183906002600685020190811061572057fe5b60200260200101818152505085818151811061573857fe5b6020908102919091010151516001602002015182826006026003018151811061575d57fe5b60200260200101818152505085818151811061577557fe5b60200260200101516020015160006002811061578d57fe5b60200201518282600602600401815181106157a457fe5b6020026020010181815250508581815181106157bc57fe5b6020026020010151602001516001600281106157d457fe5b60200201518282600602600501815181106157eb57fe5b602090810291909101015260010161567f565b50615807615dbe565b60006020826020860260208601600060086107d05a03f1905080801561582d5750815115155b98975050505050505050565b601d80546000908152601c6020526040808220548483528183205582548252902082905555601e80546001019055565b60016000818152601c6020527f6de76108811faf2f94afbe5ac6c98e8393206cd093932de1fbfd61bbeec43a02549091905b600181146158f7576000818152601a60205260408120906158bf6003830187615b32565b91505080156158de576158d183614fd8565b6001945050505050612887565b50506000818152601c602052604090205490915061589b565b5060009392505050565b6000806000806159118686615b32565b91509150801561595e576001600160a01b03858116600081815260208990526040808220805487861684529183208054929095166001600160a01b03199283161790945591905281541690555b909590945092505050565b6000805b60185461597f576000199150506151a0565b6018548110158061599257506007548110155b15615a08576000848460205443604051602001808560028111156159b257fe5b60ff1660f81b81526001018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c905060188054905081816159fe57fe5b06925050506151a0565b60006017600060188481548110615a1b57fe5b9060005260206000200154815260200190815260200160002090504381600501548260060154600554010111615ac457601960188381548110615a5a57fe5b60009182526020808320909101548354600181018555938352912090910155601880546000198101908110615a8b57fe5b906000526020600020015460188381548110615aa357fe5b6000918252602090912001556018805490615ac2906000198301615bf0565b505b5060010161596d565b615ad5615cf6565b5060408051808201909152600181526002602082015290565b615af6615cf6565b615afe615ddc565b8351815260208085015190820152604080820184905282606083600060076107d05a03f1615b2b57600080fd5b5092915050565b6001600081815260208490526040812054909182916001600160a01b03165b6001600160a01b038116600114801590615b7d5750846001600160a01b0316816001600160a01b031614155b15615ba5576001600160a01b0380821660009081526020889052604090205491925016615b51565b6001600160a01b03811660011415615bc6576001600093509350505061524c565b5091506001905061524c565b60405180608001604052806004906020820280388339509192915050565b81548183558181111561486c5760008381526020902061486c918101908301615dfa565b506000815560010160009055565b6040518060a0016040528060008152602001615c3c615d10565b81526020016000815260200160008152602001606081525090565b8260028101928215615c85579160200282015b82811115615c85578251825591602001919060010190615c6a565b50615c91929150615dfa565b5090565b828054828255906000526020600020908101928215615cea579160200282015b82811115615cea57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190615cb5565b50615c91929150615e14565b604051806040016040528060008152602001600081525090565b6040518060400160405280615d23615e38565b8152602001615d30615e38565b905290565b5080546000825590600052602060002090810190610eb39190615dfa565b8260028101928215615c85579182015b82811115615c85578254825591600101919060010190615d63565b828054828255906000526020600020908101928215615cea5760005260206000209182015b82811115615cea578254825591600101919060010190615da3565b60405180602001604052806001906020820280388339509192915050565b60405180606001604052806003906020820280388339509192915050565b61101b91905b80821115615c915760008155600101615e00565b61101b91905b80821115615c915780546001600160a01b0319168155600101615e1a565b6040518060400160405280600290602082028038833950919291505056fea60d55093b21f740878d9871e95e5031eaf5cf08a167c898ed3c62b1fb24f88796561394bac381230de4649200e8831afcab1f451881bbade9ef209f6dd304804e6f7420656e6f756768206e65776c792072656769737465726564206e6f6465734e6f206c69766520776f726b696e672067726f75702c20736b69707065642071756572794e6f206c69766520776f726b696e672067726f75702c207472696767657220626f6f747374726170436f6d6d697452657665616c206661696c7572652c20626f6f747374726170526f756e642072657365744e6f7420612076616c696420706172616d657465720000000000000000000000536b69707065642067726f757020666f726d6174696f6e2c206e6f7420656e6f756768206578706972656420776f726b696e672067726f75702e4e6f206c69766520776f726b696e672067726f75702c20736b69707065642072616e646f6d20726571756573744e6f20657870697265642070656e64696e672067726f757020746f20636c65616e207570a265627a7a723158200f4e40ae94afcdac9a5b1b6574bbc74e209c1a873460e408c5a9969e4832850764736f6c63430005100032`

// DeployDosproxy deploys a new Ethereum contract, binding an instance of Dosproxy to it.
func DeployDosproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddr common.Address, _proxyFundsAddr common.Address, _proxyFundsTokenAddr common.Address) (common.Address, *types.Transaction, *Dosproxy, error) {
	parsed, err := abi.JSON(strings.NewReader(DosproxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DosproxyBin), backend, _bridgeAddr, _proxyFundsAddr, _proxyFundsTokenAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dosproxy{DosproxyCaller: DosproxyCaller{contract: contract}, DosproxyTransactor: DosproxyTransactor{contract: contract}, DosproxyFilterer: DosproxyFilterer{contract: contract}}, nil
}

// Dosproxy is an auto generated Go binding around an Ethereum contract.
type Dosproxy struct {
	DosproxyCaller     // Read-only binding to the contract
	DosproxyTransactor // Write-only binding to the contract
	DosproxyFilterer   // Log filterer for contract events
}

// DosproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DosproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DosproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DosproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DosproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DosproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DosproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DosproxySession struct {
	Contract     *Dosproxy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DosproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DosproxyCallerSession struct {
	Contract *DosproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DosproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DosproxyTransactorSession struct {
	Contract     *DosproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DosproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DosproxyRaw struct {
	Contract *Dosproxy // Generic contract binding to access the raw methods on
}

// DosproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DosproxyCallerRaw struct {
	Contract *DosproxyCaller // Generic read-only contract binding to access the raw methods on
}

// DosproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DosproxyTransactorRaw struct {
	Contract *DosproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDosproxy creates a new instance of Dosproxy, bound to a specific deployed contract.
func NewDosproxy(address common.Address, backend bind.ContractBackend) (*Dosproxy, error) {
	contract, err := bindDosproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dosproxy{DosproxyCaller: DosproxyCaller{contract: contract}, DosproxyTransactor: DosproxyTransactor{contract: contract}, DosproxyFilterer: DosproxyFilterer{contract: contract}}, nil
}

// NewDosproxyCaller creates a new read-only instance of Dosproxy, bound to a specific deployed contract.
func NewDosproxyCaller(address common.Address, caller bind.ContractCaller) (*DosproxyCaller, error) {
	contract, err := bindDosproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DosproxyCaller{contract: contract}, nil
}

// NewDosproxyTransactor creates a new write-only instance of Dosproxy, bound to a specific deployed contract.
func NewDosproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*DosproxyTransactor, error) {
	contract, err := bindDosproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DosproxyTransactor{contract: contract}, nil
}

// NewDosproxyFilterer creates a new log filterer instance of Dosproxy, bound to a specific deployed contract.
func NewDosproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*DosproxyFilterer, error) {
	contract, err := bindDosproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DosproxyFilterer{contract: contract}, nil
}

// bindDosproxy binds a generic wrapper to an already deployed contract.
func bindDosproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DosproxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dosproxy *DosproxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dosproxy.Contract.DosproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dosproxy *DosproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosproxy.Contract.DosproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dosproxy *DosproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dosproxy.Contract.DosproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dosproxy *DosproxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dosproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dosproxy *DosproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dosproxy *DosproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dosproxy.Contract.contract.Transact(opts, method, params...)
}

// AddressBridge is a free data retrieval call binding the contract method 0x76cffa53.
//
// Solidity: function addressBridge() constant returns(address)
func (_Dosproxy *DosproxyCaller) AddressBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "addressBridge")
	return *ret0, err
}

// AddressBridge is a free data retrieval call binding the contract method 0x76cffa53.
//
// Solidity: function addressBridge() constant returns(address)
func (_Dosproxy *DosproxySession) AddressBridge() (common.Address, error) {
	return _Dosproxy.Contract.AddressBridge(&_Dosproxy.CallOpts)
}

// AddressBridge is a free data retrieval call binding the contract method 0x76cffa53.
//
// Solidity: function addressBridge() constant returns(address)
func (_Dosproxy *DosproxyCallerSession) AddressBridge() (common.Address, error) {
	return _Dosproxy.Contract.AddressBridge(&_Dosproxy.CallOpts)
}

// BootstrapCommitDuration is a free data retrieval call binding the contract method 0x372a53cc.
//
// Solidity: function bootstrapCommitDuration() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) BootstrapCommitDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "bootstrapCommitDuration")
	return *ret0, err
}

// BootstrapCommitDuration is a free data retrieval call binding the contract method 0x372a53cc.
//
// Solidity: function bootstrapCommitDuration() constant returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapCommitDuration() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapCommitDuration(&_Dosproxy.CallOpts)
}

// BootstrapCommitDuration is a free data retrieval call binding the contract method 0x372a53cc.
//
// Solidity: function bootstrapCommitDuration() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapCommitDuration() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapCommitDuration(&_Dosproxy.CallOpts)
}

// BootstrapEndBlk is a free data retrieval call binding the contract method 0x19717203.
//
// Solidity: function bootstrapEndBlk() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) BootstrapEndBlk(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "bootstrapEndBlk")
	return *ret0, err
}

// BootstrapEndBlk is a free data retrieval call binding the contract method 0x19717203.
//
// Solidity: function bootstrapEndBlk() constant returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapEndBlk() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapEndBlk(&_Dosproxy.CallOpts)
}

// BootstrapEndBlk is a free data retrieval call binding the contract method 0x19717203.
//
// Solidity: function bootstrapEndBlk() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapEndBlk() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapEndBlk(&_Dosproxy.CallOpts)
}

// BootstrapGroups is a free data retrieval call binding the contract method 0x31bf6464.
//
// Solidity: function bootstrapGroups() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) BootstrapGroups(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "bootstrapGroups")
	return *ret0, err
}

// BootstrapGroups is a free data retrieval call binding the contract method 0x31bf6464.
//
// Solidity: function bootstrapGroups() constant returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapGroups() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapGroups(&_Dosproxy.CallOpts)
}

// BootstrapGroups is a free data retrieval call binding the contract method 0x31bf6464.
//
// Solidity: function bootstrapGroups() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapGroups() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapGroups(&_Dosproxy.CallOpts)
}

// BootstrapRevealDuration is a free data retrieval call binding the contract method 0xef112dfc.
//
// Solidity: function bootstrapRevealDuration() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) BootstrapRevealDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "bootstrapRevealDuration")
	return *ret0, err
}

// BootstrapRevealDuration is a free data retrieval call binding the contract method 0xef112dfc.
//
// Solidity: function bootstrapRevealDuration() constant returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapRevealDuration() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapRevealDuration(&_Dosproxy.CallOpts)
}

// BootstrapRevealDuration is a free data retrieval call binding the contract method 0xef112dfc.
//
// Solidity: function bootstrapRevealDuration() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapRevealDuration() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapRevealDuration(&_Dosproxy.CallOpts)
}

// BootstrapRound is a free data retrieval call binding the contract method 0x85ed4223.
//
// Solidity: function bootstrapRound() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) BootstrapRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "bootstrapRound")
	return *ret0, err
}

// BootstrapRound is a free data retrieval call binding the contract method 0x85ed4223.
//
// Solidity: function bootstrapRound() constant returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapRound() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapRound(&_Dosproxy.CallOpts)
}

// BootstrapRound is a free data retrieval call binding the contract method 0x85ed4223.
//
// Solidity: function bootstrapRound() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapRound() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapRound(&_Dosproxy.CallOpts)
}

// BootstrapStartThreshold is a free data retrieval call binding the contract method 0x11bc5311.
//
// Solidity: function bootstrapStartThreshold() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) BootstrapStartThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "bootstrapStartThreshold")
	return *ret0, err
}

// BootstrapStartThreshold is a free data retrieval call binding the contract method 0x11bc5311.
//
// Solidity: function bootstrapStartThreshold() constant returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapStartThreshold() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapStartThreshold(&_Dosproxy.CallOpts)
}

// BootstrapStartThreshold is a free data retrieval call binding the contract method 0x11bc5311.
//
// Solidity: function bootstrapStartThreshold() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapStartThreshold() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapStartThreshold(&_Dosproxy.CallOpts)
}

// CheckExpireLimit is a free data retrieval call binding the contract method 0x6e5454d3.
//
// Solidity: function checkExpireLimit() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) CheckExpireLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "checkExpireLimit")
	return *ret0, err
}

// CheckExpireLimit is a free data retrieval call binding the contract method 0x6e5454d3.
//
// Solidity: function checkExpireLimit() constant returns(uint256)
func (_Dosproxy *DosproxySession) CheckExpireLimit() (*big.Int, error) {
	return _Dosproxy.Contract.CheckExpireLimit(&_Dosproxy.CallOpts)
}

// CheckExpireLimit is a free data retrieval call binding the contract method 0x6e5454d3.
//
// Solidity: function checkExpireLimit() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) CheckExpireLimit() (*big.Int, error) {
	return _Dosproxy.Contract.CheckExpireLimit(&_Dosproxy.CallOpts)
}

// ExpiredWorkingGroupIds is a free data retrieval call binding the contract method 0xb45ef79d.
//
// Solidity: function expiredWorkingGroupIds( uint256) constant returns(uint256)
func (_Dosproxy *DosproxyCaller) ExpiredWorkingGroupIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "expiredWorkingGroupIds", arg0)
	return *ret0, err
}

// ExpiredWorkingGroupIds is a free data retrieval call binding the contract method 0xb45ef79d.
//
// Solidity: function expiredWorkingGroupIds( uint256) constant returns(uint256)
func (_Dosproxy *DosproxySession) ExpiredWorkingGroupIds(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.ExpiredWorkingGroupIds(&_Dosproxy.CallOpts, arg0)
}

// ExpiredWorkingGroupIds is a free data retrieval call binding the contract method 0xb45ef79d.
//
// Solidity: function expiredWorkingGroupIds( uint256) constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) ExpiredWorkingGroupIds(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.ExpiredWorkingGroupIds(&_Dosproxy.CallOpts, arg0)
}

// ExpiredWorkingGroupIdsLength is a free data retrieval call binding the contract method 0x830687c4.
//
// Solidity: function expiredWorkingGroupIdsLength() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) ExpiredWorkingGroupIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "expiredWorkingGroupIdsLength")
	return *ret0, err
}

// ExpiredWorkingGroupIdsLength is a free data retrieval call binding the contract method 0x830687c4.
//
// Solidity: function expiredWorkingGroupIdsLength() constant returns(uint256)
func (_Dosproxy *DosproxySession) ExpiredWorkingGroupIdsLength() (*big.Int, error) {
	return _Dosproxy.Contract.ExpiredWorkingGroupIdsLength(&_Dosproxy.CallOpts)
}

// ExpiredWorkingGroupIdsLength is a free data retrieval call binding the contract method 0x830687c4.
//
// Solidity: function expiredWorkingGroupIdsLength() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) ExpiredWorkingGroupIdsLength() (*big.Int, error) {
	return _Dosproxy.Contract.ExpiredWorkingGroupIdsLength(&_Dosproxy.CallOpts)
}

// GetExpiredWorkingGroupSize is a free data retrieval call binding the contract method 0xefde068b.
//
// Solidity: function getExpiredWorkingGroupSize() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) GetExpiredWorkingGroupSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "getExpiredWorkingGroupSize")
	return *ret0, err
}

// GetExpiredWorkingGroupSize is a free data retrieval call binding the contract method 0xefde068b.
//
// Solidity: function getExpiredWorkingGroupSize() constant returns(uint256)
func (_Dosproxy *DosproxySession) GetExpiredWorkingGroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GetExpiredWorkingGroupSize(&_Dosproxy.CallOpts)
}

// GetExpiredWorkingGroupSize is a free data retrieval call binding the contract method 0xefde068b.
//
// Solidity: function getExpiredWorkingGroupSize() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GetExpiredWorkingGroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GetExpiredWorkingGroupSize(&_Dosproxy.CallOpts)
}

// GetGroupPubKey is a free data retrieval call binding the contract method 0x92021653.
//
// Solidity: function getGroupPubKey(idx uint256) constant returns(uint256[4])
func (_Dosproxy *DosproxyCaller) GetGroupPubKey(opts *bind.CallOpts, idx *big.Int) ([4]*big.Int, error) {
	var (
		ret0 = new([4]*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "getGroupPubKey", idx)
	return *ret0, err
}

// GetGroupPubKey is a free data retrieval call binding the contract method 0x92021653.
//
// Solidity: function getGroupPubKey(idx uint256) constant returns(uint256[4])
func (_Dosproxy *DosproxySession) GetGroupPubKey(idx *big.Int) ([4]*big.Int, error) {
	return _Dosproxy.Contract.GetGroupPubKey(&_Dosproxy.CallOpts, idx)
}

// GetGroupPubKey is a free data retrieval call binding the contract method 0x92021653.
//
// Solidity: function getGroupPubKey(idx uint256) constant returns(uint256[4])
func (_Dosproxy *DosproxyCallerSession) GetGroupPubKey(idx *big.Int) ([4]*big.Int, error) {
	return _Dosproxy.Contract.GetGroupPubKey(&_Dosproxy.CallOpts, idx)
}

// GetLastHandledGroup is a free data retrieval call binding the contract method 0x4a4b52b4.
//
// Solidity: function getLastHandledGroup() constant returns(uint256, uint256[4], uint256, uint256, address[])
func (_Dosproxy *DosproxyCaller) GetLastHandledGroup(opts *bind.CallOpts) (*big.Int, [4]*big.Int, *big.Int, *big.Int, []common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([4]*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Dosproxy.contract.Call(opts, out, "getLastHandledGroup")
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetLastHandledGroup is a free data retrieval call binding the contract method 0x4a4b52b4.
//
// Solidity: function getLastHandledGroup() constant returns(uint256, uint256[4], uint256, uint256, address[])
func (_Dosproxy *DosproxySession) GetLastHandledGroup() (*big.Int, [4]*big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Dosproxy.Contract.GetLastHandledGroup(&_Dosproxy.CallOpts)
}

// GetLastHandledGroup is a free data retrieval call binding the contract method 0x4a4b52b4.
//
// Solidity: function getLastHandledGroup() constant returns(uint256, uint256[4], uint256, uint256, address[])
func (_Dosproxy *DosproxyCallerSession) GetLastHandledGroup() (*big.Int, [4]*big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Dosproxy.Contract.GetLastHandledGroup(&_Dosproxy.CallOpts)
}

// GetWorkingGroupById is a free data retrieval call binding the contract method 0x02957d53.
//
// Solidity: function getWorkingGroupById(groupId uint256) constant returns(uint256, uint256[4], uint256, uint256, address[])
func (_Dosproxy *DosproxyCaller) GetWorkingGroupById(opts *bind.CallOpts, groupId *big.Int) (*big.Int, [4]*big.Int, *big.Int, *big.Int, []common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([4]*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Dosproxy.contract.Call(opts, out, "getWorkingGroupById", groupId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetWorkingGroupById is a free data retrieval call binding the contract method 0x02957d53.
//
// Solidity: function getWorkingGroupById(groupId uint256) constant returns(uint256, uint256[4], uint256, uint256, address[])
func (_Dosproxy *DosproxySession) GetWorkingGroupById(groupId *big.Int) (*big.Int, [4]*big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Dosproxy.Contract.GetWorkingGroupById(&_Dosproxy.CallOpts, groupId)
}

// GetWorkingGroupById is a free data retrieval call binding the contract method 0x02957d53.
//
// Solidity: function getWorkingGroupById(groupId uint256) constant returns(uint256, uint256[4], uint256, uint256, address[])
func (_Dosproxy *DosproxyCallerSession) GetWorkingGroupById(groupId *big.Int) (*big.Int, [4]*big.Int, *big.Int, *big.Int, []common.Address, error) {
	return _Dosproxy.Contract.GetWorkingGroupById(&_Dosproxy.CallOpts, groupId)
}

// GetWorkingGroupSize is a free data retrieval call binding the contract method 0xb5372264.
//
// Solidity: function getWorkingGroupSize() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) GetWorkingGroupSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "getWorkingGroupSize")
	return *ret0, err
}

// GetWorkingGroupSize is a free data retrieval call binding the contract method 0xb5372264.
//
// Solidity: function getWorkingGroupSize() constant returns(uint256)
func (_Dosproxy *DosproxySession) GetWorkingGroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GetWorkingGroupSize(&_Dosproxy.CallOpts)
}

// GetWorkingGroupSize is a free data retrieval call binding the contract method 0xb5372264.
//
// Solidity: function getWorkingGroupSize() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GetWorkingGroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GetWorkingGroupSize(&_Dosproxy.CallOpts)
}

// GroupMaturityPeriod is a free data retrieval call binding the contract method 0x7c48d1a0.
//
// Solidity: function groupMaturityPeriod() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) GroupMaturityPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "groupMaturityPeriod")
	return *ret0, err
}

// GroupMaturityPeriod is a free data retrieval call binding the contract method 0x7c48d1a0.
//
// Solidity: function groupMaturityPeriod() constant returns(uint256)
func (_Dosproxy *DosproxySession) GroupMaturityPeriod() (*big.Int, error) {
	return _Dosproxy.Contract.GroupMaturityPeriod(&_Dosproxy.CallOpts)
}

// GroupMaturityPeriod is a free data retrieval call binding the contract method 0x7c48d1a0.
//
// Solidity: function groupMaturityPeriod() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GroupMaturityPeriod() (*big.Int, error) {
	return _Dosproxy.Contract.GroupMaturityPeriod(&_Dosproxy.CallOpts)
}

// GroupSize is a free data retrieval call binding the contract method 0x63b635ea.
//
// Solidity: function groupSize() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) GroupSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "groupSize")
	return *ret0, err
}

// GroupSize is a free data retrieval call binding the contract method 0x63b635ea.
//
// Solidity: function groupSize() constant returns(uint256)
func (_Dosproxy *DosproxySession) GroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GroupSize(&_Dosproxy.CallOpts)
}

// GroupSize is a free data retrieval call binding the contract method 0x63b635ea.
//
// Solidity: function groupSize() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GroupSize(&_Dosproxy.CallOpts)
}

// GroupToPick is a free data retrieval call binding the contract method 0x0434ccd2.
//
// Solidity: function groupToPick() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) GroupToPick(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "groupToPick")
	return *ret0, err
}

// GroupToPick is a free data retrieval call binding the contract method 0x0434ccd2.
//
// Solidity: function groupToPick() constant returns(uint256)
func (_Dosproxy *DosproxySession) GroupToPick() (*big.Int, error) {
	return _Dosproxy.Contract.GroupToPick(&_Dosproxy.CallOpts)
}

// GroupToPick is a free data retrieval call binding the contract method 0x0434ccd2.
//
// Solidity: function groupToPick() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GroupToPick() (*big.Int, error) {
	return _Dosproxy.Contract.GroupToPick(&_Dosproxy.CallOpts)
}

// GuardianListed is a free data retrieval call binding the contract method 0x09011cb9.
//
// Solidity: function guardianListed( address) constant returns(bool)
func (_Dosproxy *DosproxyCaller) GuardianListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "guardianListed", arg0)
	return *ret0, err
}

// GuardianListed is a free data retrieval call binding the contract method 0x09011cb9.
//
// Solidity: function guardianListed( address) constant returns(bool)
func (_Dosproxy *DosproxySession) GuardianListed(arg0 common.Address) (bool, error) {
	return _Dosproxy.Contract.GuardianListed(&_Dosproxy.CallOpts, arg0)
}

// GuardianListed is a free data retrieval call binding the contract method 0x09011cb9.
//
// Solidity: function guardianListed( address) constant returns(bool)
func (_Dosproxy *DosproxyCallerSession) GuardianListed(arg0 common.Address) (bool, error) {
	return _Dosproxy.Contract.GuardianListed(&_Dosproxy.CallOpts, arg0)
}

// InitBlkN is a free data retrieval call binding the contract method 0x95071cf6.
//
// Solidity: function initBlkN() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) InitBlkN(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "initBlkN")
	return *ret0, err
}

// InitBlkN is a free data retrieval call binding the contract method 0x95071cf6.
//
// Solidity: function initBlkN() constant returns(uint256)
func (_Dosproxy *DosproxySession) InitBlkN() (*big.Int, error) {
	return _Dosproxy.Contract.InitBlkN(&_Dosproxy.CallOpts)
}

// InitBlkN is a free data retrieval call binding the contract method 0x95071cf6.
//
// Solidity: function initBlkN() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) InitBlkN() (*big.Int, error) {
	return _Dosproxy.Contract.InitBlkN(&_Dosproxy.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dosproxy *DosproxyCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dosproxy *DosproxySession) IsOwner() (bool, error) {
	return _Dosproxy.Contract.IsOwner(&_Dosproxy.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dosproxy *DosproxyCallerSession) IsOwner() (bool, error) {
	return _Dosproxy.Contract.IsOwner(&_Dosproxy.CallOpts)
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) LastRandomness(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "lastRandomness")
	return *ret0, err
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() constant returns(uint256)
func (_Dosproxy *DosproxySession) LastRandomness() (*big.Int, error) {
	return _Dosproxy.Contract.LastRandomness(&_Dosproxy.CallOpts)
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) LastRandomness() (*big.Int, error) {
	return _Dosproxy.Contract.LastRandomness(&_Dosproxy.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) LastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "lastUpdatedBlock")
	return *ret0, err
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() constant returns(uint256)
func (_Dosproxy *DosproxySession) LastUpdatedBlock() (*big.Int, error) {
	return _Dosproxy.Contract.LastUpdatedBlock(&_Dosproxy.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) LastUpdatedBlock() (*big.Int, error) {
	return _Dosproxy.Contract.LastUpdatedBlock(&_Dosproxy.CallOpts)
}

// LifeDiversity is a free data retrieval call binding the contract method 0xdd6ceddf.
//
// Solidity: function lifeDiversity() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) LifeDiversity(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "lifeDiversity")
	return *ret0, err
}

// LifeDiversity is a free data retrieval call binding the contract method 0xdd6ceddf.
//
// Solidity: function lifeDiversity() constant returns(uint256)
func (_Dosproxy *DosproxySession) LifeDiversity() (*big.Int, error) {
	return _Dosproxy.Contract.LifeDiversity(&_Dosproxy.CallOpts)
}

// LifeDiversity is a free data retrieval call binding the contract method 0xdd6ceddf.
//
// Solidity: function lifeDiversity() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) LifeDiversity() (*big.Int, error) {
	return _Dosproxy.Contract.LifeDiversity(&_Dosproxy.CallOpts)
}

// NodeToGroupIdList is a free data retrieval call binding the contract method 0x0eeee5c1.
//
// Solidity: function nodeToGroupIdList( address,  uint256) constant returns(uint256)
func (_Dosproxy *DosproxyCaller) NodeToGroupIdList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "nodeToGroupIdList", arg0, arg1)
	return *ret0, err
}

// NodeToGroupIdList is a free data retrieval call binding the contract method 0x0eeee5c1.
//
// Solidity: function nodeToGroupIdList( address,  uint256) constant returns(uint256)
func (_Dosproxy *DosproxySession) NodeToGroupIdList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.NodeToGroupIdList(&_Dosproxy.CallOpts, arg0, arg1)
}

// NodeToGroupIdList is a free data retrieval call binding the contract method 0x0eeee5c1.
//
// Solidity: function nodeToGroupIdList( address,  uint256) constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) NodeToGroupIdList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.NodeToGroupIdList(&_Dosproxy.CallOpts, arg0, arg1)
}

// NumPendingGroups is a free data retrieval call binding the contract method 0x863bc0a1.
//
// Solidity: function numPendingGroups() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) NumPendingGroups(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "numPendingGroups")
	return *ret0, err
}

// NumPendingGroups is a free data retrieval call binding the contract method 0x863bc0a1.
//
// Solidity: function numPendingGroups() constant returns(uint256)
func (_Dosproxy *DosproxySession) NumPendingGroups() (*big.Int, error) {
	return _Dosproxy.Contract.NumPendingGroups(&_Dosproxy.CallOpts)
}

// NumPendingGroups is a free data retrieval call binding the contract method 0x863bc0a1.
//
// Solidity: function numPendingGroups() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) NumPendingGroups() (*big.Int, error) {
	return _Dosproxy.Contract.NumPendingGroups(&_Dosproxy.CallOpts)
}

// NumPendingNodes is a free data retrieval call binding the contract method 0xf41a1587.
//
// Solidity: function numPendingNodes() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) NumPendingNodes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "numPendingNodes")
	return *ret0, err
}

// NumPendingNodes is a free data retrieval call binding the contract method 0xf41a1587.
//
// Solidity: function numPendingNodes() constant returns(uint256)
func (_Dosproxy *DosproxySession) NumPendingNodes() (*big.Int, error) {
	return _Dosproxy.Contract.NumPendingNodes(&_Dosproxy.CallOpts)
}

// NumPendingNodes is a free data retrieval call binding the contract method 0xf41a1587.
//
// Solidity: function numPendingNodes() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) NumPendingNodes() (*big.Int, error) {
	return _Dosproxy.Contract.NumPendingNodes(&_Dosproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dosproxy *DosproxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dosproxy *DosproxySession) Owner() (common.Address, error) {
	return _Dosproxy.Contract.Owner(&_Dosproxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dosproxy *DosproxyCallerSession) Owner() (common.Address, error) {
	return _Dosproxy.Contract.Owner(&_Dosproxy.CallOpts)
}

// PendingGroupList is a free data retrieval call binding the contract method 0xa54fb00e.
//
// Solidity: function pendingGroupList( uint256) constant returns(uint256)
func (_Dosproxy *DosproxyCaller) PendingGroupList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "pendingGroupList", arg0)
	return *ret0, err
}

// PendingGroupList is a free data retrieval call binding the contract method 0xa54fb00e.
//
// Solidity: function pendingGroupList( uint256) constant returns(uint256)
func (_Dosproxy *DosproxySession) PendingGroupList(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupList(&_Dosproxy.CallOpts, arg0)
}

// PendingGroupList is a free data retrieval call binding the contract method 0xa54fb00e.
//
// Solidity: function pendingGroupList( uint256) constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) PendingGroupList(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupList(&_Dosproxy.CallOpts, arg0)
}

// PendingGroupMaxLife is a free data retrieval call binding the contract method 0xfc84dde4.
//
// Solidity: function pendingGroupMaxLife() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) PendingGroupMaxLife(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "pendingGroupMaxLife")
	return *ret0, err
}

// PendingGroupMaxLife is a free data retrieval call binding the contract method 0xfc84dde4.
//
// Solidity: function pendingGroupMaxLife() constant returns(uint256)
func (_Dosproxy *DosproxySession) PendingGroupMaxLife() (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupMaxLife(&_Dosproxy.CallOpts)
}

// PendingGroupMaxLife is a free data retrieval call binding the contract method 0xfc84dde4.
//
// Solidity: function pendingGroupMaxLife() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) PendingGroupMaxLife() (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupMaxLife(&_Dosproxy.CallOpts)
}

// PendingGroupTail is a free data retrieval call binding the contract method 0x190ca29e.
//
// Solidity: function pendingGroupTail() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) PendingGroupTail(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "pendingGroupTail")
	return *ret0, err
}

// PendingGroupTail is a free data retrieval call binding the contract method 0x190ca29e.
//
// Solidity: function pendingGroupTail() constant returns(uint256)
func (_Dosproxy *DosproxySession) PendingGroupTail() (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupTail(&_Dosproxy.CallOpts)
}

// PendingGroupTail is a free data retrieval call binding the contract method 0x190ca29e.
//
// Solidity: function pendingGroupTail() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) PendingGroupTail() (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupTail(&_Dosproxy.CallOpts)
}

// PendingGroups is a free data retrieval call binding the contract method 0xd18c81b7.
//
// Solidity: function pendingGroups( uint256) constant returns(groupId uint256, startBlkNum uint256)
func (_Dosproxy *DosproxyCaller) PendingGroups(opts *bind.CallOpts, arg0 *big.Int) (struct {
	GroupId     *big.Int
	StartBlkNum *big.Int
}, error) {
	ret := new(struct {
		GroupId     *big.Int
		StartBlkNum *big.Int
	})
	out := ret
	err := _Dosproxy.contract.Call(opts, out, "pendingGroups", arg0)
	return *ret, err
}

// PendingGroups is a free data retrieval call binding the contract method 0xd18c81b7.
//
// Solidity: function pendingGroups( uint256) constant returns(groupId uint256, startBlkNum uint256)
func (_Dosproxy *DosproxySession) PendingGroups(arg0 *big.Int) (struct {
	GroupId     *big.Int
	StartBlkNum *big.Int
}, error) {
	return _Dosproxy.Contract.PendingGroups(&_Dosproxy.CallOpts, arg0)
}

// PendingGroups is a free data retrieval call binding the contract method 0xd18c81b7.
//
// Solidity: function pendingGroups( uint256) constant returns(groupId uint256, startBlkNum uint256)
func (_Dosproxy *DosproxyCallerSession) PendingGroups(arg0 *big.Int) (struct {
	GroupId     *big.Int
	StartBlkNum *big.Int
}, error) {
	return _Dosproxy.Contract.PendingGroups(&_Dosproxy.CallOpts, arg0)
}

// PendingNodeList is a free data retrieval call binding the contract method 0xa60b007d.
//
// Solidity: function pendingNodeList( address) constant returns(address)
func (_Dosproxy *DosproxyCaller) PendingNodeList(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "pendingNodeList", arg0)
	return *ret0, err
}

// PendingNodeList is a free data retrieval call binding the contract method 0xa60b007d.
//
// Solidity: function pendingNodeList( address) constant returns(address)
func (_Dosproxy *DosproxySession) PendingNodeList(arg0 common.Address) (common.Address, error) {
	return _Dosproxy.Contract.PendingNodeList(&_Dosproxy.CallOpts, arg0)
}

// PendingNodeList is a free data retrieval call binding the contract method 0xa60b007d.
//
// Solidity: function pendingNodeList( address) constant returns(address)
func (_Dosproxy *DosproxyCallerSession) PendingNodeList(arg0 common.Address) (common.Address, error) {
	return _Dosproxy.Contract.PendingNodeList(&_Dosproxy.CallOpts, arg0)
}

// PendingNodeTail is a free data retrieval call binding the contract method 0x094c3612.
//
// Solidity: function pendingNodeTail() constant returns(address)
func (_Dosproxy *DosproxyCaller) PendingNodeTail(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "pendingNodeTail")
	return *ret0, err
}

// PendingNodeTail is a free data retrieval call binding the contract method 0x094c3612.
//
// Solidity: function pendingNodeTail() constant returns(address)
func (_Dosproxy *DosproxySession) PendingNodeTail() (common.Address, error) {
	return _Dosproxy.Contract.PendingNodeTail(&_Dosproxy.CallOpts)
}

// PendingNodeTail is a free data retrieval call binding the contract method 0x094c3612.
//
// Solidity: function pendingNodeTail() constant returns(address)
func (_Dosproxy *DosproxyCallerSession) PendingNodeTail() (common.Address, error) {
	return _Dosproxy.Contract.PendingNodeTail(&_Dosproxy.CallOpts)
}

// ProxyFundsAddr is a free data retrieval call binding the contract method 0xdf37c617.
//
// Solidity: function proxyFundsAddr() constant returns(address)
func (_Dosproxy *DosproxyCaller) ProxyFundsAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "proxyFundsAddr")
	return *ret0, err
}

// ProxyFundsAddr is a free data retrieval call binding the contract method 0xdf37c617.
//
// Solidity: function proxyFundsAddr() constant returns(address)
func (_Dosproxy *DosproxySession) ProxyFundsAddr() (common.Address, error) {
	return _Dosproxy.Contract.ProxyFundsAddr(&_Dosproxy.CallOpts)
}

// ProxyFundsAddr is a free data retrieval call binding the contract method 0xdf37c617.
//
// Solidity: function proxyFundsAddr() constant returns(address)
func (_Dosproxy *DosproxyCallerSession) ProxyFundsAddr() (common.Address, error) {
	return _Dosproxy.Contract.ProxyFundsAddr(&_Dosproxy.CallOpts)
}

// ProxyFundsTokenAddr is a free data retrieval call binding the contract method 0x99ca2d30.
//
// Solidity: function proxyFundsTokenAddr() constant returns(address)
func (_Dosproxy *DosproxyCaller) ProxyFundsTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "proxyFundsTokenAddr")
	return *ret0, err
}

// ProxyFundsTokenAddr is a free data retrieval call binding the contract method 0x99ca2d30.
//
// Solidity: function proxyFundsTokenAddr() constant returns(address)
func (_Dosproxy *DosproxySession) ProxyFundsTokenAddr() (common.Address, error) {
	return _Dosproxy.Contract.ProxyFundsTokenAddr(&_Dosproxy.CallOpts)
}

// ProxyFundsTokenAddr is a free data retrieval call binding the contract method 0x99ca2d30.
//
// Solidity: function proxyFundsTokenAddr() constant returns(address)
func (_Dosproxy *DosproxyCallerSession) ProxyFundsTokenAddr() (common.Address, error) {
	return _Dosproxy.Contract.ProxyFundsTokenAddr(&_Dosproxy.CallOpts)
}

// RefreshSystemRandomHardLimit is a free data retrieval call binding the contract method 0x962ba8a4.
//
// Solidity: function refreshSystemRandomHardLimit() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) RefreshSystemRandomHardLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "refreshSystemRandomHardLimit")
	return *ret0, err
}

// RefreshSystemRandomHardLimit is a free data retrieval call binding the contract method 0x962ba8a4.
//
// Solidity: function refreshSystemRandomHardLimit() constant returns(uint256)
func (_Dosproxy *DosproxySession) RefreshSystemRandomHardLimit() (*big.Int, error) {
	return _Dosproxy.Contract.RefreshSystemRandomHardLimit(&_Dosproxy.CallOpts)
}

// RefreshSystemRandomHardLimit is a free data retrieval call binding the contract method 0x962ba8a4.
//
// Solidity: function refreshSystemRandomHardLimit() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) RefreshSystemRandomHardLimit() (*big.Int, error) {
	return _Dosproxy.Contract.RefreshSystemRandomHardLimit(&_Dosproxy.CallOpts)
}

// WorkingGroupIds is a free data retrieval call binding the contract method 0x5d381204.
//
// Solidity: function workingGroupIds( uint256) constant returns(uint256)
func (_Dosproxy *DosproxyCaller) WorkingGroupIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "workingGroupIds", arg0)
	return *ret0, err
}

// WorkingGroupIds is a free data retrieval call binding the contract method 0x5d381204.
//
// Solidity: function workingGroupIds( uint256) constant returns(uint256)
func (_Dosproxy *DosproxySession) WorkingGroupIds(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.WorkingGroupIds(&_Dosproxy.CallOpts, arg0)
}

// WorkingGroupIds is a free data retrieval call binding the contract method 0x5d381204.
//
// Solidity: function workingGroupIds( uint256) constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) WorkingGroupIds(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.WorkingGroupIds(&_Dosproxy.CallOpts, arg0)
}

// WorkingGroupIdsLength is a free data retrieval call binding the contract method 0x11db6574.
//
// Solidity: function workingGroupIdsLength() constant returns(uint256)
func (_Dosproxy *DosproxyCaller) WorkingGroupIdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "workingGroupIdsLength")
	return *ret0, err
}

// WorkingGroupIdsLength is a free data retrieval call binding the contract method 0x11db6574.
//
// Solidity: function workingGroupIdsLength() constant returns(uint256)
func (_Dosproxy *DosproxySession) WorkingGroupIdsLength() (*big.Int, error) {
	return _Dosproxy.Contract.WorkingGroupIdsLength(&_Dosproxy.CallOpts)
}

// WorkingGroupIdsLength is a free data retrieval call binding the contract method 0x11db6574.
//
// Solidity: function workingGroupIdsLength() constant returns(uint256)
func (_Dosproxy *DosproxyCallerSession) WorkingGroupIdsLength() (*big.Int, error) {
	return _Dosproxy.Contract.WorkingGroupIdsLength(&_Dosproxy.CallOpts)
}

// Callback_ is a paid mutator transaction binding the contract method 0x18a1908d.
//
// Solidity: function __callback__(requestId uint256, rndSeed uint256) returns()
func (_Dosproxy *DosproxyTransactor) Callback_(opts *bind.TransactOpts, requestId *big.Int, rndSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "__callback__", requestId, rndSeed)
}

// Callback_ is a paid mutator transaction binding the contract method 0x18a1908d.
//
// Solidity: function __callback__(requestId uint256, rndSeed uint256) returns()
func (_Dosproxy *DosproxySession) Callback_(requestId *big.Int, rndSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.Callback_(&_Dosproxy.TransactOpts, requestId, rndSeed)
}

// Callback_ is a paid mutator transaction binding the contract method 0x18a1908d.
//
// Solidity: function __callback__(requestId uint256, rndSeed uint256) returns()
func (_Dosproxy *DosproxyTransactorSession) Callback_(requestId *big.Int, rndSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.Callback_(&_Dosproxy.TransactOpts, requestId, rndSeed)
}

// AddToGuardianList is a paid mutator transaction binding the contract method 0xd79351b2.
//
// Solidity: function addToGuardianList(_addr address) returns()
func (_Dosproxy *DosproxyTransactor) AddToGuardianList(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "addToGuardianList", _addr)
}

// AddToGuardianList is a paid mutator transaction binding the contract method 0xd79351b2.
//
// Solidity: function addToGuardianList(_addr address) returns()
func (_Dosproxy *DosproxySession) AddToGuardianList(_addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.AddToGuardianList(&_Dosproxy.TransactOpts, _addr)
}

// AddToGuardianList is a paid mutator transaction binding the contract method 0xd79351b2.
//
// Solidity: function addToGuardianList(_addr address) returns()
func (_Dosproxy *DosproxyTransactorSession) AddToGuardianList(_addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.AddToGuardianList(&_Dosproxy.TransactOpts, _addr)
}

// Query is a paid mutator transaction binding the contract method 0xb7fb8fd7.
//
// Solidity: function query(from address, timeout uint256, dataSource string, selector string) returns(uint256)
func (_Dosproxy *DosproxyTransactor) Query(opts *bind.TransactOpts, from common.Address, timeout *big.Int, dataSource string, selector string) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "query", from, timeout, dataSource, selector)
}

// Query is a paid mutator transaction binding the contract method 0xb7fb8fd7.
//
// Solidity: function query(from address, timeout uint256, dataSource string, selector string) returns(uint256)
func (_Dosproxy *DosproxySession) Query(from common.Address, timeout *big.Int, dataSource string, selector string) (*types.Transaction, error) {
	return _Dosproxy.Contract.Query(&_Dosproxy.TransactOpts, from, timeout, dataSource, selector)
}

// Query is a paid mutator transaction binding the contract method 0xb7fb8fd7.
//
// Solidity: function query(from address, timeout uint256, dataSource string, selector string) returns(uint256)
func (_Dosproxy *DosproxyTransactorSession) Query(from common.Address, timeout *big.Int, dataSource string, selector string) (*types.Transaction, error) {
	return _Dosproxy.Contract.Query(&_Dosproxy.TransactOpts, from, timeout, dataSource, selector)
}

// RegisterGroupPubKey is a paid mutator transaction binding the contract method 0xb836ccea.
//
// Solidity: function registerGroupPubKey(groupId uint256, suggestedPubKey uint256[4]) returns()
func (_Dosproxy *DosproxyTransactor) RegisterGroupPubKey(opts *bind.TransactOpts, groupId *big.Int, suggestedPubKey [4]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "registerGroupPubKey", groupId, suggestedPubKey)
}

// RegisterGroupPubKey is a paid mutator transaction binding the contract method 0xb836ccea.
//
// Solidity: function registerGroupPubKey(groupId uint256, suggestedPubKey uint256[4]) returns()
func (_Dosproxy *DosproxySession) RegisterGroupPubKey(groupId *big.Int, suggestedPubKey [4]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.RegisterGroupPubKey(&_Dosproxy.TransactOpts, groupId, suggestedPubKey)
}

// RegisterGroupPubKey is a paid mutator transaction binding the contract method 0xb836ccea.
//
// Solidity: function registerGroupPubKey(groupId uint256, suggestedPubKey uint256[4]) returns()
func (_Dosproxy *DosproxyTransactorSession) RegisterGroupPubKey(groupId *big.Int, suggestedPubKey [4]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.RegisterGroupPubKey(&_Dosproxy.TransactOpts, groupId, suggestedPubKey)
}

// RegisterNewNode is a paid mutator transaction binding the contract method 0xaeb3da73.
//
// Solidity: function registerNewNode() returns()
func (_Dosproxy *DosproxyTransactor) RegisterNewNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "registerNewNode")
}

// RegisterNewNode is a paid mutator transaction binding the contract method 0xaeb3da73.
//
// Solidity: function registerNewNode() returns()
func (_Dosproxy *DosproxySession) RegisterNewNode() (*types.Transaction, error) {
	return _Dosproxy.Contract.RegisterNewNode(&_Dosproxy.TransactOpts)
}

// RegisterNewNode is a paid mutator transaction binding the contract method 0xaeb3da73.
//
// Solidity: function registerNewNode() returns()
func (_Dosproxy *DosproxyTransactorSession) RegisterNewNode() (*types.Transaction, error) {
	return _Dosproxy.Contract.RegisterNewNode(&_Dosproxy.TransactOpts)
}

// RemoveFromGuardianList is a paid mutator transaction binding the contract method 0xc58ebe1c.
//
// Solidity: function removeFromGuardianList(_addr address) returns()
func (_Dosproxy *DosproxyTransactor) RemoveFromGuardianList(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "removeFromGuardianList", _addr)
}

// RemoveFromGuardianList is a paid mutator transaction binding the contract method 0xc58ebe1c.
//
// Solidity: function removeFromGuardianList(_addr address) returns()
func (_Dosproxy *DosproxySession) RemoveFromGuardianList(_addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.RemoveFromGuardianList(&_Dosproxy.TransactOpts, _addr)
}

// RemoveFromGuardianList is a paid mutator transaction binding the contract method 0xc58ebe1c.
//
// Solidity: function removeFromGuardianList(_addr address) returns()
func (_Dosproxy *DosproxyTransactorSession) RemoveFromGuardianList(_addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.RemoveFromGuardianList(&_Dosproxy.TransactOpts, _addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dosproxy *DosproxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dosproxy *DosproxySession) RenounceOwnership() (*types.Transaction, error) {
	return _Dosproxy.Contract.RenounceOwnership(&_Dosproxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dosproxy *DosproxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dosproxy.Contract.RenounceOwnership(&_Dosproxy.TransactOpts)
}

// RequestRandom is a paid mutator transaction binding the contract method 0xc7c3f4a5.
//
// Solidity: function requestRandom(from address, userSeed uint256) returns(uint256)
func (_Dosproxy *DosproxyTransactor) RequestRandom(opts *bind.TransactOpts, from common.Address, userSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "requestRandom", from, userSeed)
}

// RequestRandom is a paid mutator transaction binding the contract method 0xc7c3f4a5.
//
// Solidity: function requestRandom(from address, userSeed uint256) returns(uint256)
func (_Dosproxy *DosproxySession) RequestRandom(from common.Address, userSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.RequestRandom(&_Dosproxy.TransactOpts, from, userSeed)
}

// RequestRandom is a paid mutator transaction binding the contract method 0xc7c3f4a5.
//
// Solidity: function requestRandom(from address, userSeed uint256) returns(uint256)
func (_Dosproxy *DosproxyTransactorSession) RequestRandom(from common.Address, userSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.RequestRandom(&_Dosproxy.TransactOpts, from, userSeed)
}

// SetBootstrapStartThreshold is a paid mutator transaction binding the contract method 0x100063ec.
//
// Solidity: function setBootstrapStartThreshold(newThreshold uint256) returns()
func (_Dosproxy *DosproxyTransactor) SetBootstrapStartThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setBootstrapStartThreshold", newThreshold)
}

// SetBootstrapStartThreshold is a paid mutator transaction binding the contract method 0x100063ec.
//
// Solidity: function setBootstrapStartThreshold(newThreshold uint256) returns()
func (_Dosproxy *DosproxySession) SetBootstrapStartThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetBootstrapStartThreshold(&_Dosproxy.TransactOpts, newThreshold)
}

// SetBootstrapStartThreshold is a paid mutator transaction binding the contract method 0x100063ec.
//
// Solidity: function setBootstrapStartThreshold(newThreshold uint256) returns()
func (_Dosproxy *DosproxyTransactorSession) SetBootstrapStartThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetBootstrapStartThreshold(&_Dosproxy.TransactOpts, newThreshold)
}

// SetGroupMaturityPeriod is a paid mutator transaction binding the contract method 0x925fc6c9.
//
// Solidity: function setGroupMaturityPeriod(newPeriod uint256) returns()
func (_Dosproxy *DosproxyTransactor) SetGroupMaturityPeriod(opts *bind.TransactOpts, newPeriod *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setGroupMaturityPeriod", newPeriod)
}

// SetGroupMaturityPeriod is a paid mutator transaction binding the contract method 0x925fc6c9.
//
// Solidity: function setGroupMaturityPeriod(newPeriod uint256) returns()
func (_Dosproxy *DosproxySession) SetGroupMaturityPeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetGroupMaturityPeriod(&_Dosproxy.TransactOpts, newPeriod)
}

// SetGroupMaturityPeriod is a paid mutator transaction binding the contract method 0x925fc6c9.
//
// Solidity: function setGroupMaturityPeriod(newPeriod uint256) returns()
func (_Dosproxy *DosproxyTransactorSession) SetGroupMaturityPeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetGroupMaturityPeriod(&_Dosproxy.TransactOpts, newPeriod)
}

// SetGroupSize is a paid mutator transaction binding the contract method 0x0dfc09cb.
//
// Solidity: function setGroupSize(newSize uint256) returns()
func (_Dosproxy *DosproxyTransactor) SetGroupSize(opts *bind.TransactOpts, newSize *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setGroupSize", newSize)
}

// SetGroupSize is a paid mutator transaction binding the contract method 0x0dfc09cb.
//
// Solidity: function setGroupSize(newSize uint256) returns()
func (_Dosproxy *DosproxySession) SetGroupSize(newSize *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetGroupSize(&_Dosproxy.TransactOpts, newSize)
}

// SetGroupSize is a paid mutator transaction binding the contract method 0x0dfc09cb.
//
// Solidity: function setGroupSize(newSize uint256) returns()
func (_Dosproxy *DosproxyTransactorSession) SetGroupSize(newSize *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetGroupSize(&_Dosproxy.TransactOpts, newSize)
}

// SetLifeDiversity is a paid mutator transaction binding the contract method 0x559ea9de.
//
// Solidity: function setLifeDiversity(newDiversity uint256) returns()
func (_Dosproxy *DosproxyTransactor) SetLifeDiversity(opts *bind.TransactOpts, newDiversity *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setLifeDiversity", newDiversity)
}

// SetLifeDiversity is a paid mutator transaction binding the contract method 0x559ea9de.
//
// Solidity: function setLifeDiversity(newDiversity uint256) returns()
func (_Dosproxy *DosproxySession) SetLifeDiversity(newDiversity *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetLifeDiversity(&_Dosproxy.TransactOpts, newDiversity)
}

// SetLifeDiversity is a paid mutator transaction binding the contract method 0x559ea9de.
//
// Solidity: function setLifeDiversity(newDiversity uint256) returns()
func (_Dosproxy *DosproxyTransactorSession) SetLifeDiversity(newDiversity *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetLifeDiversity(&_Dosproxy.TransactOpts, newDiversity)
}

// SetPendingGroupMaxLife is a paid mutator transaction binding the contract method 0x4a28a74d.
//
// Solidity: function setPendingGroupMaxLife(newLife uint256) returns()
func (_Dosproxy *DosproxyTransactor) SetPendingGroupMaxLife(opts *bind.TransactOpts, newLife *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setPendingGroupMaxLife", newLife)
}

// SetPendingGroupMaxLife is a paid mutator transaction binding the contract method 0x4a28a74d.
//
// Solidity: function setPendingGroupMaxLife(newLife uint256) returns()
func (_Dosproxy *DosproxySession) SetPendingGroupMaxLife(newLife *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetPendingGroupMaxLife(&_Dosproxy.TransactOpts, newLife)
}

// SetPendingGroupMaxLife is a paid mutator transaction binding the contract method 0x4a28a74d.
//
// Solidity: function setPendingGroupMaxLife(newLife uint256) returns()
func (_Dosproxy *DosproxyTransactorSession) SetPendingGroupMaxLife(newLife *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetPendingGroupMaxLife(&_Dosproxy.TransactOpts, newLife)
}

// SetProxyFund is a paid mutator transaction binding the contract method 0x40e4a5af.
//
// Solidity: function setProxyFund(newFund address, newFundToken address) returns()
func (_Dosproxy *DosproxyTransactor) SetProxyFund(opts *bind.TransactOpts, newFund common.Address, newFundToken common.Address) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setProxyFund", newFund, newFundToken)
}

// SetProxyFund is a paid mutator transaction binding the contract method 0x40e4a5af.
//
// Solidity: function setProxyFund(newFund address, newFundToken address) returns()
func (_Dosproxy *DosproxySession) SetProxyFund(newFund common.Address, newFundToken common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetProxyFund(&_Dosproxy.TransactOpts, newFund, newFundToken)
}

// SetProxyFund is a paid mutator transaction binding the contract method 0x40e4a5af.
//
// Solidity: function setProxyFund(newFund address, newFundToken address) returns()
func (_Dosproxy *DosproxyTransactorSession) SetProxyFund(newFund common.Address, newFundToken common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetProxyFund(&_Dosproxy.TransactOpts, newFund, newFundToken)
}

// SetSystemRandomHardLimit is a paid mutator transaction binding the contract method 0xc457aa8f.
//
// Solidity: function setSystemRandomHardLimit(newLimit uint256) returns()
func (_Dosproxy *DosproxyTransactor) SetSystemRandomHardLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setSystemRandomHardLimit", newLimit)
}

// SetSystemRandomHardLimit is a paid mutator transaction binding the contract method 0xc457aa8f.
//
// Solidity: function setSystemRandomHardLimit(newLimit uint256) returns()
func (_Dosproxy *DosproxySession) SetSystemRandomHardLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetSystemRandomHardLimit(&_Dosproxy.TransactOpts, newLimit)
}

// SetSystemRandomHardLimit is a paid mutator transaction binding the contract method 0xc457aa8f.
//
// Solidity: function setSystemRandomHardLimit(newLimit uint256) returns()
func (_Dosproxy *DosproxyTransactorSession) SetSystemRandomHardLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetSystemRandomHardLimit(&_Dosproxy.TransactOpts, newLimit)
}

// SignalBootstrap is a paid mutator transaction binding the contract method 0x5c0e159f.
//
// Solidity: function signalBootstrap(_cid uint256) returns()
func (_Dosproxy *DosproxyTransactor) SignalBootstrap(opts *bind.TransactOpts, _cid *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "signalBootstrap", _cid)
}

// SignalBootstrap is a paid mutator transaction binding the contract method 0x5c0e159f.
//
// Solidity: function signalBootstrap(_cid uint256) returns()
func (_Dosproxy *DosproxySession) SignalBootstrap(_cid *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalBootstrap(&_Dosproxy.TransactOpts, _cid)
}

// SignalBootstrap is a paid mutator transaction binding the contract method 0x5c0e159f.
//
// Solidity: function signalBootstrap(_cid uint256) returns()
func (_Dosproxy *DosproxyTransactorSession) SignalBootstrap(_cid *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalBootstrap(&_Dosproxy.TransactOpts, _cid)
}

// SignalGroupDissolve is a paid mutator transaction binding the contract method 0x5be6c3af.
//
// Solidity: function signalGroupDissolve() returns()
func (_Dosproxy *DosproxyTransactor) SignalGroupDissolve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "signalGroupDissolve")
}

// SignalGroupDissolve is a paid mutator transaction binding the contract method 0x5be6c3af.
//
// Solidity: function signalGroupDissolve() returns()
func (_Dosproxy *DosproxySession) SignalGroupDissolve() (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalGroupDissolve(&_Dosproxy.TransactOpts)
}

// SignalGroupDissolve is a paid mutator transaction binding the contract method 0x5be6c3af.
//
// Solidity: function signalGroupDissolve() returns()
func (_Dosproxy *DosproxyTransactorSession) SignalGroupDissolve() (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalGroupDissolve(&_Dosproxy.TransactOpts)
}

// SignalGroupFormation is a paid mutator transaction binding the contract method 0x155fa82c.
//
// Solidity: function signalGroupFormation() returns()
func (_Dosproxy *DosproxyTransactor) SignalGroupFormation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "signalGroupFormation")
}

// SignalGroupFormation is a paid mutator transaction binding the contract method 0x155fa82c.
//
// Solidity: function signalGroupFormation() returns()
func (_Dosproxy *DosproxySession) SignalGroupFormation() (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalGroupFormation(&_Dosproxy.TransactOpts)
}

// SignalGroupFormation is a paid mutator transaction binding the contract method 0x155fa82c.
//
// Solidity: function signalGroupFormation() returns()
func (_Dosproxy *DosproxyTransactorSession) SignalGroupFormation() (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalGroupFormation(&_Dosproxy.TransactOpts)
}

// SignalRandom is a paid mutator transaction binding the contract method 0xb9424b35.
//
// Solidity: function signalRandom() returns()
func (_Dosproxy *DosproxyTransactor) SignalRandom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "signalRandom")
}

// SignalRandom is a paid mutator transaction binding the contract method 0xb9424b35.
//
// Solidity: function signalRandom() returns()
func (_Dosproxy *DosproxySession) SignalRandom() (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalRandom(&_Dosproxy.TransactOpts)
}

// SignalRandom is a paid mutator transaction binding the contract method 0xb9424b35.
//
// Solidity: function signalRandom() returns()
func (_Dosproxy *DosproxyTransactorSession) SignalRandom() (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalRandom(&_Dosproxy.TransactOpts)
}

// SignalUnregister is a paid mutator transaction binding the contract method 0x7c1cf083.
//
// Solidity: function signalUnregister(member address) returns()
func (_Dosproxy *DosproxyTransactor) SignalUnregister(opts *bind.TransactOpts, member common.Address) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "signalUnregister", member)
}

// SignalUnregister is a paid mutator transaction binding the contract method 0x7c1cf083.
//
// Solidity: function signalUnregister(member address) returns()
func (_Dosproxy *DosproxySession) SignalUnregister(member common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalUnregister(&_Dosproxy.TransactOpts, member)
}

// SignalUnregister is a paid mutator transaction binding the contract method 0x7c1cf083.
//
// Solidity: function signalUnregister(member address) returns()
func (_Dosproxy *DosproxyTransactorSession) SignalUnregister(member common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalUnregister(&_Dosproxy.TransactOpts, member)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dosproxy *DosproxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dosproxy *DosproxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.TransferOwnership(&_Dosproxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dosproxy *DosproxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.TransferOwnership(&_Dosproxy.TransactOpts, newOwner)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x74ad3a06.
//
// Solidity: function triggerCallback(requestId uint256, trafficType uint8, result bytes, sig uint256[2]) returns()
func (_Dosproxy *DosproxyTransactor) TriggerCallback(opts *bind.TransactOpts, requestId *big.Int, trafficType uint8, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "triggerCallback", requestId, trafficType, result, sig)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x74ad3a06.
//
// Solidity: function triggerCallback(requestId uint256, trafficType uint8, result bytes, sig uint256[2]) returns()
func (_Dosproxy *DosproxySession) TriggerCallback(requestId *big.Int, trafficType uint8, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.TriggerCallback(&_Dosproxy.TransactOpts, requestId, trafficType, result, sig)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x74ad3a06.
//
// Solidity: function triggerCallback(requestId uint256, trafficType uint8, result bytes, sig uint256[2]) returns()
func (_Dosproxy *DosproxyTransactorSession) TriggerCallback(requestId *big.Int, trafficType uint8, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.TriggerCallback(&_Dosproxy.TransactOpts, requestId, trafficType, result, sig)
}

// UnregisterNode is a paid mutator transaction binding the contract method 0x3d385cf5.
//
// Solidity: function unregisterNode() returns(bool)
func (_Dosproxy *DosproxyTransactor) UnregisterNode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "unregisterNode")
}

// UnregisterNode is a paid mutator transaction binding the contract method 0x3d385cf5.
//
// Solidity: function unregisterNode() returns(bool)
func (_Dosproxy *DosproxySession) UnregisterNode() (*types.Transaction, error) {
	return _Dosproxy.Contract.UnregisterNode(&_Dosproxy.TransactOpts)
}

// UnregisterNode is a paid mutator transaction binding the contract method 0x3d385cf5.
//
// Solidity: function unregisterNode() returns(bool)
func (_Dosproxy *DosproxyTransactorSession) UnregisterNode() (*types.Transaction, error) {
	return _Dosproxy.Contract.UnregisterNode(&_Dosproxy.TransactOpts)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(sig uint256[2]) returns()
func (_Dosproxy *DosproxyTransactor) UpdateRandomness(opts *bind.TransactOpts, sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "updateRandomness", sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(sig uint256[2]) returns()
func (_Dosproxy *DosproxySession) UpdateRandomness(sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.UpdateRandomness(&_Dosproxy.TransactOpts, sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(sig uint256[2]) returns()
func (_Dosproxy *DosproxyTransactorSession) UpdateRandomness(sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.UpdateRandomness(&_Dosproxy.TransactOpts, sig)
}

// DosproxyGuardianRewardIterator is returned from FilterGuardianReward and is used to iterate over the raw logs and unpacked data for GuardianReward events raised by the Dosproxy contract.
type DosproxyGuardianRewardIterator struct {
	Event *DosproxyGuardianReward // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyGuardianRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyGuardianReward)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyGuardianReward)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyGuardianRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyGuardianRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyGuardianReward represents a GuardianReward event raised by the Dosproxy contract.
type DosproxyGuardianReward struct {
	BlkNum   *big.Int
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGuardianReward is a free log retrieval operation binding the contract event 0xa60d55093b21f740878d9871e95e5031eaf5cf08a167c898ed3c62b1fb24f887.
//
// Solidity: e GuardianReward(blkNum uint256, guardian address)
func (_Dosproxy *DosproxyFilterer) FilterGuardianReward(opts *bind.FilterOpts) (*DosproxyGuardianRewardIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "GuardianReward")
	if err != nil {
		return nil, err
	}
	return &DosproxyGuardianRewardIterator{contract: _Dosproxy.contract, event: "GuardianReward", logs: logs, sub: sub}, nil
}

// WatchGuardianReward is a free log subscription operation binding the contract event 0xa60d55093b21f740878d9871e95e5031eaf5cf08a167c898ed3c62b1fb24f887.
//
// Solidity: e GuardianReward(blkNum uint256, guardian address)
func (_Dosproxy *DosproxyFilterer) WatchGuardianReward(opts *bind.WatchOpts, sink chan<- *DosproxyGuardianReward) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "GuardianReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyGuardianReward)
				if err := _Dosproxy.contract.UnpackLog(event, "GuardianReward", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogCallbackTriggeredForIterator is returned from FilterLogCallbackTriggeredFor and is used to iterate over the raw logs and unpacked data for LogCallbackTriggeredFor events raised by the Dosproxy contract.
type DosproxyLogCallbackTriggeredForIterator struct {
	Event *DosproxyLogCallbackTriggeredFor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogCallbackTriggeredForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogCallbackTriggeredFor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogCallbackTriggeredFor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogCallbackTriggeredForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogCallbackTriggeredForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogCallbackTriggeredFor represents a LogCallbackTriggeredFor event raised by the Dosproxy contract.
type DosproxyLogCallbackTriggeredFor struct {
	CallbackAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogCallbackTriggeredFor is a free log retrieval operation binding the contract event 0x065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf0.
//
// Solidity: e LogCallbackTriggeredFor(callbackAddr address)
func (_Dosproxy *DosproxyFilterer) FilterLogCallbackTriggeredFor(opts *bind.FilterOpts) (*DosproxyLogCallbackTriggeredForIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogCallbackTriggeredFor")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogCallbackTriggeredForIterator{contract: _Dosproxy.contract, event: "LogCallbackTriggeredFor", logs: logs, sub: sub}, nil
}

// WatchLogCallbackTriggeredFor is a free log subscription operation binding the contract event 0x065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf0.
//
// Solidity: e LogCallbackTriggeredFor(callbackAddr address)
func (_Dosproxy *DosproxyFilterer) WatchLogCallbackTriggeredFor(opts *bind.WatchOpts, sink chan<- *DosproxyLogCallbackTriggeredFor) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogCallbackTriggeredFor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogCallbackTriggeredFor)
				if err := _Dosproxy.contract.UnpackLog(event, "LogCallbackTriggeredFor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogGroupDissolveIterator is returned from FilterLogGroupDissolve and is used to iterate over the raw logs and unpacked data for LogGroupDissolve events raised by the Dosproxy contract.
type DosproxyLogGroupDissolveIterator struct {
	Event *DosproxyLogGroupDissolve // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogGroupDissolveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogGroupDissolve)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogGroupDissolve)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogGroupDissolveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogGroupDissolveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogGroupDissolve represents a LogGroupDissolve event raised by the Dosproxy contract.
type DosproxyLogGroupDissolve struct {
	GroupId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogGroupDissolve is a free log retrieval operation binding the contract event 0xf7377b41bdc770cc22a1bad318571f0ae6d65188371bdbe0cb660c0de57b3150.
//
// Solidity: e LogGroupDissolve(groupId uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogGroupDissolve(opts *bind.FilterOpts) (*DosproxyLogGroupDissolveIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogGroupDissolve")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogGroupDissolveIterator{contract: _Dosproxy.contract, event: "LogGroupDissolve", logs: logs, sub: sub}, nil
}

// WatchLogGroupDissolve is a free log subscription operation binding the contract event 0xf7377b41bdc770cc22a1bad318571f0ae6d65188371bdbe0cb660c0de57b3150.
//
// Solidity: e LogGroupDissolve(groupId uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogGroupDissolve(opts *bind.WatchOpts, sink chan<- *DosproxyLogGroupDissolve) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogGroupDissolve")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogGroupDissolve)
				if err := _Dosproxy.contract.UnpackLog(event, "LogGroupDissolve", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogGroupingIterator is returned from FilterLogGrouping and is used to iterate over the raw logs and unpacked data for LogGrouping events raised by the Dosproxy contract.
type DosproxyLogGroupingIterator struct {
	Event *DosproxyLogGrouping // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogGroupingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogGrouping)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogGrouping)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogGroupingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogGroupingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogGrouping represents a LogGrouping event raised by the Dosproxy contract.
type DosproxyLogGrouping struct {
	GroupId *big.Int
	NodeId  []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogGrouping is a free log retrieval operation binding the contract event 0x78bf54a42d1b98e6c809c3e5904898c5b3304ede546b6f070e83d9a32e544d4f.
//
// Solidity: e LogGrouping(groupId uint256, nodeId address[])
func (_Dosproxy *DosproxyFilterer) FilterLogGrouping(opts *bind.FilterOpts) (*DosproxyLogGroupingIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogGrouping")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogGroupingIterator{contract: _Dosproxy.contract, event: "LogGrouping", logs: logs, sub: sub}, nil
}

// WatchLogGrouping is a free log subscription operation binding the contract event 0x78bf54a42d1b98e6c809c3e5904898c5b3304ede546b6f070e83d9a32e544d4f.
//
// Solidity: e LogGrouping(groupId uint256, nodeId address[])
func (_Dosproxy *DosproxyFilterer) WatchLogGrouping(opts *bind.WatchOpts, sink chan<- *DosproxyLogGrouping) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogGrouping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogGrouping)
				if err := _Dosproxy.contract.UnpackLog(event, "LogGrouping", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogGroupingInitiatedIterator is returned from FilterLogGroupingInitiated and is used to iterate over the raw logs and unpacked data for LogGroupingInitiated events raised by the Dosproxy contract.
type DosproxyLogGroupingInitiatedIterator struct {
	Event *DosproxyLogGroupingInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogGroupingInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogGroupingInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogGroupingInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogGroupingInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogGroupingInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogGroupingInitiated represents a LogGroupingInitiated event raised by the Dosproxy contract.
type DosproxyLogGroupingInitiated struct {
	PendingNodePool *big.Int
	Groupsize       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogGroupingInitiated is a free log retrieval operation binding the contract event 0x5ee79c79dd5be1f25f86c7028ec8fb6849839e8a23f676afc907eade45946986.
//
// Solidity: e LogGroupingInitiated(pendingNodePool uint256, groupsize uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogGroupingInitiated(opts *bind.FilterOpts) (*DosproxyLogGroupingInitiatedIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogGroupingInitiated")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogGroupingInitiatedIterator{contract: _Dosproxy.contract, event: "LogGroupingInitiated", logs: logs, sub: sub}, nil
}

// WatchLogGroupingInitiated is a free log subscription operation binding the contract event 0x5ee79c79dd5be1f25f86c7028ec8fb6849839e8a23f676afc907eade45946986.
//
// Solidity: e LogGroupingInitiated(pendingNodePool uint256, groupsize uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogGroupingInitiated(opts *bind.WatchOpts, sink chan<- *DosproxyLogGroupingInitiated) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogGroupingInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogGroupingInitiated)
				if err := _Dosproxy.contract.UnpackLog(event, "LogGroupingInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogInsufficientPendingNodeIterator is returned from FilterLogInsufficientPendingNode and is used to iterate over the raw logs and unpacked data for LogInsufficientPendingNode events raised by the Dosproxy contract.
type DosproxyLogInsufficientPendingNodeIterator struct {
	Event *DosproxyLogInsufficientPendingNode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogInsufficientPendingNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogInsufficientPendingNode)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogInsufficientPendingNode)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogInsufficientPendingNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogInsufficientPendingNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogInsufficientPendingNode represents a LogInsufficientPendingNode event raised by the Dosproxy contract.
type DosproxyLogInsufficientPendingNode struct {
	NumPendingNodes *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogInsufficientPendingNode is a free log retrieval operation binding the contract event 0xc03848aa1689c7c291dcc68fa62d6109e13f16b25e89bdef5a6c8598b36b8019.
//
// Solidity: e LogInsufficientPendingNode(numPendingNodes uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogInsufficientPendingNode(opts *bind.FilterOpts) (*DosproxyLogInsufficientPendingNodeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogInsufficientPendingNode")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogInsufficientPendingNodeIterator{contract: _Dosproxy.contract, event: "LogInsufficientPendingNode", logs: logs, sub: sub}, nil
}

// WatchLogInsufficientPendingNode is a free log subscription operation binding the contract event 0xc03848aa1689c7c291dcc68fa62d6109e13f16b25e89bdef5a6c8598b36b8019.
//
// Solidity: e LogInsufficientPendingNode(numPendingNodes uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogInsufficientPendingNode(opts *bind.WatchOpts, sink chan<- *DosproxyLogInsufficientPendingNode) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogInsufficientPendingNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogInsufficientPendingNode)
				if err := _Dosproxy.contract.UnpackLog(event, "LogInsufficientPendingNode", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogInsufficientWorkingGroupIterator is returned from FilterLogInsufficientWorkingGroup and is used to iterate over the raw logs and unpacked data for LogInsufficientWorkingGroup events raised by the Dosproxy contract.
type DosproxyLogInsufficientWorkingGroupIterator struct {
	Event *DosproxyLogInsufficientWorkingGroup // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogInsufficientWorkingGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogInsufficientWorkingGroup)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogInsufficientWorkingGroup)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogInsufficientWorkingGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogInsufficientWorkingGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogInsufficientWorkingGroup represents a LogInsufficientWorkingGroup event raised by the Dosproxy contract.
type DosproxyLogInsufficientWorkingGroup struct {
	NumWorkingGroups *big.Int
	NumPendingGroups *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogInsufficientWorkingGroup is a free log retrieval operation binding the contract event 0x1850da28de32299250accda835079ca1340fbd447032a1f6dac77381a77a26c8.
//
// Solidity: e LogInsufficientWorkingGroup(numWorkingGroups uint256, numPendingGroups uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogInsufficientWorkingGroup(opts *bind.FilterOpts) (*DosproxyLogInsufficientWorkingGroupIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogInsufficientWorkingGroup")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogInsufficientWorkingGroupIterator{contract: _Dosproxy.contract, event: "LogInsufficientWorkingGroup", logs: logs, sub: sub}, nil
}

// WatchLogInsufficientWorkingGroup is a free log subscription operation binding the contract event 0x1850da28de32299250accda835079ca1340fbd447032a1f6dac77381a77a26c8.
//
// Solidity: e LogInsufficientWorkingGroup(numWorkingGroups uint256, numPendingGroups uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogInsufficientWorkingGroup(opts *bind.WatchOpts, sink chan<- *DosproxyLogInsufficientWorkingGroup) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogInsufficientWorkingGroup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogInsufficientWorkingGroup)
				if err := _Dosproxy.contract.UnpackLog(event, "LogInsufficientWorkingGroup", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogMessageIterator is returned from FilterLogMessage and is used to iterate over the raw logs and unpacked data for LogMessage events raised by the Dosproxy contract.
type DosproxyLogMessageIterator struct {
	Event *DosproxyLogMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogMessage represents a LogMessage event raised by the Dosproxy contract.
type DosproxyLogMessage struct {
	Info string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogMessage is a free log retrieval operation binding the contract event 0x96561394bac381230de4649200e8831afcab1f451881bbade9ef209f6dd30480.
//
// Solidity: e LogMessage(info string)
func (_Dosproxy *DosproxyFilterer) FilterLogMessage(opts *bind.FilterOpts) (*DosproxyLogMessageIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogMessage")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogMessageIterator{contract: _Dosproxy.contract, event: "LogMessage", logs: logs, sub: sub}, nil
}

// WatchLogMessage is a free log subscription operation binding the contract event 0x96561394bac381230de4649200e8831afcab1f451881bbade9ef209f6dd30480.
//
// Solidity: e LogMessage(info string)
func (_Dosproxy *DosproxyFilterer) WatchLogMessage(opts *bind.WatchOpts, sink chan<- *DosproxyLogMessage) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogMessage)
				if err := _Dosproxy.contract.UnpackLog(event, "LogMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogNoPendingGroupIterator is returned from FilterLogNoPendingGroup and is used to iterate over the raw logs and unpacked data for LogNoPendingGroup events raised by the Dosproxy contract.
type DosproxyLogNoPendingGroupIterator struct {
	Event *DosproxyLogNoPendingGroup // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogNoPendingGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogNoPendingGroup)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogNoPendingGroup)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogNoPendingGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogNoPendingGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogNoPendingGroup represents a LogNoPendingGroup event raised by the Dosproxy contract.
type DosproxyLogNoPendingGroup struct {
	GroupId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogNoPendingGroup is a free log retrieval operation binding the contract event 0x71047c0893a51085656a2894bba10bc6ef51a654f25e1ead1929b076487a9569.
//
// Solidity: e LogNoPendingGroup(groupId uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogNoPendingGroup(opts *bind.FilterOpts) (*DosproxyLogNoPendingGroupIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogNoPendingGroup")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogNoPendingGroupIterator{contract: _Dosproxy.contract, event: "LogNoPendingGroup", logs: logs, sub: sub}, nil
}

// WatchLogNoPendingGroup is a free log subscription operation binding the contract event 0x71047c0893a51085656a2894bba10bc6ef51a654f25e1ead1929b076487a9569.
//
// Solidity: e LogNoPendingGroup(groupId uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogNoPendingGroup(opts *bind.WatchOpts, sink chan<- *DosproxyLogNoPendingGroup) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogNoPendingGroup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogNoPendingGroup)
				if err := _Dosproxy.contract.UnpackLog(event, "LogNoPendingGroup", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogNonContractCallIterator is returned from FilterLogNonContractCall and is used to iterate over the raw logs and unpacked data for LogNonContractCall events raised by the Dosproxy contract.
type DosproxyLogNonContractCallIterator struct {
	Event *DosproxyLogNonContractCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogNonContractCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogNonContractCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogNonContractCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogNonContractCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogNonContractCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogNonContractCall represents a LogNonContractCall event raised by the Dosproxy contract.
type DosproxyLogNonContractCall struct {
	From common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogNonContractCall is a free log retrieval operation binding the contract event 0x6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb5.
//
// Solidity: e LogNonContractCall(from address)
func (_Dosproxy *DosproxyFilterer) FilterLogNonContractCall(opts *bind.FilterOpts) (*DosproxyLogNonContractCallIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogNonContractCall")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogNonContractCallIterator{contract: _Dosproxy.contract, event: "LogNonContractCall", logs: logs, sub: sub}, nil
}

// WatchLogNonContractCall is a free log subscription operation binding the contract event 0x6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb5.
//
// Solidity: e LogNonContractCall(from address)
func (_Dosproxy *DosproxyFilterer) WatchLogNonContractCall(opts *bind.WatchOpts, sink chan<- *DosproxyLogNonContractCall) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogNonContractCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogNonContractCall)
				if err := _Dosproxy.contract.UnpackLog(event, "LogNonContractCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogNonSupportedTypeIterator is returned from FilterLogNonSupportedType and is used to iterate over the raw logs and unpacked data for LogNonSupportedType events raised by the Dosproxy contract.
type DosproxyLogNonSupportedTypeIterator struct {
	Event *DosproxyLogNonSupportedType // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogNonSupportedTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogNonSupportedType)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogNonSupportedType)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogNonSupportedTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogNonSupportedTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogNonSupportedType represents a LogNonSupportedType event raised by the Dosproxy contract.
type DosproxyLogNonSupportedType struct {
	InvalidSelector string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogNonSupportedType is a free log retrieval operation binding the contract event 0x70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41.
//
// Solidity: e LogNonSupportedType(invalidSelector string)
func (_Dosproxy *DosproxyFilterer) FilterLogNonSupportedType(opts *bind.FilterOpts) (*DosproxyLogNonSupportedTypeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogNonSupportedType")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogNonSupportedTypeIterator{contract: _Dosproxy.contract, event: "LogNonSupportedType", logs: logs, sub: sub}, nil
}

// WatchLogNonSupportedType is a free log subscription operation binding the contract event 0x70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41.
//
// Solidity: e LogNonSupportedType(invalidSelector string)
func (_Dosproxy *DosproxyFilterer) WatchLogNonSupportedType(opts *bind.WatchOpts, sink chan<- *DosproxyLogNonSupportedType) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogNonSupportedType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogNonSupportedType)
				if err := _Dosproxy.contract.UnpackLog(event, "LogNonSupportedType", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogPendingGroupRemovedIterator is returned from FilterLogPendingGroupRemoved and is used to iterate over the raw logs and unpacked data for LogPendingGroupRemoved events raised by the Dosproxy contract.
type DosproxyLogPendingGroupRemovedIterator struct {
	Event *DosproxyLogPendingGroupRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogPendingGroupRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogPendingGroupRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogPendingGroupRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogPendingGroupRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogPendingGroupRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogPendingGroupRemoved represents a LogPendingGroupRemoved event raised by the Dosproxy contract.
type DosproxyLogPendingGroupRemoved struct {
	GroupId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogPendingGroupRemoved is a free log retrieval operation binding the contract event 0x156927b06a61046135669011768a03b2592ee3667374525502b62b0aef4cbd21.
//
// Solidity: e LogPendingGroupRemoved(groupId uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogPendingGroupRemoved(opts *bind.FilterOpts) (*DosproxyLogPendingGroupRemovedIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogPendingGroupRemoved")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogPendingGroupRemovedIterator{contract: _Dosproxy.contract, event: "LogPendingGroupRemoved", logs: logs, sub: sub}, nil
}

// WatchLogPendingGroupRemoved is a free log subscription operation binding the contract event 0x156927b06a61046135669011768a03b2592ee3667374525502b62b0aef4cbd21.
//
// Solidity: e LogPendingGroupRemoved(groupId uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogPendingGroupRemoved(opts *bind.WatchOpts, sink chan<- *DosproxyLogPendingGroupRemoved) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogPendingGroupRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogPendingGroupRemoved)
				if err := _Dosproxy.contract.UnpackLog(event, "LogPendingGroupRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogPublicKeyAcceptedIterator is returned from FilterLogPublicKeyAccepted and is used to iterate over the raw logs and unpacked data for LogPublicKeyAccepted events raised by the Dosproxy contract.
type DosproxyLogPublicKeyAcceptedIterator struct {
	Event *DosproxyLogPublicKeyAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogPublicKeyAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogPublicKeyAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogPublicKeyAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogPublicKeyAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogPublicKeyAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogPublicKeyAccepted represents a LogPublicKeyAccepted event raised by the Dosproxy contract.
type DosproxyLogPublicKeyAccepted struct {
	GroupId          *big.Int
	PubKey           [4]*big.Int
	NumWorkingGroups *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogPublicKeyAccepted is a free log retrieval operation binding the contract event 0x9ad0d0bfff7f0fc653b03785685d46101b09e3cb1f50081915bc8101662e4a88.
//
// Solidity: e LogPublicKeyAccepted(groupId uint256, pubKey uint256[4], numWorkingGroups uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogPublicKeyAccepted(opts *bind.FilterOpts) (*DosproxyLogPublicKeyAcceptedIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogPublicKeyAccepted")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogPublicKeyAcceptedIterator{contract: _Dosproxy.contract, event: "LogPublicKeyAccepted", logs: logs, sub: sub}, nil
}

// WatchLogPublicKeyAccepted is a free log subscription operation binding the contract event 0x9ad0d0bfff7f0fc653b03785685d46101b09e3cb1f50081915bc8101662e4a88.
//
// Solidity: e LogPublicKeyAccepted(groupId uint256, pubKey uint256[4], numWorkingGroups uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogPublicKeyAccepted(opts *bind.WatchOpts, sink chan<- *DosproxyLogPublicKeyAccepted) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogPublicKeyAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogPublicKeyAccepted)
				if err := _Dosproxy.contract.UnpackLog(event, "LogPublicKeyAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogPublicKeySuggestedIterator is returned from FilterLogPublicKeySuggested and is used to iterate over the raw logs and unpacked data for LogPublicKeySuggested events raised by the Dosproxy contract.
type DosproxyLogPublicKeySuggestedIterator struct {
	Event *DosproxyLogPublicKeySuggested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogPublicKeySuggestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogPublicKeySuggested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogPublicKeySuggested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogPublicKeySuggestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogPublicKeySuggestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogPublicKeySuggested represents a LogPublicKeySuggested event raised by the Dosproxy contract.
type DosproxyLogPublicKeySuggested struct {
	GroupId     *big.Int
	PubKeyCount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogPublicKeySuggested is a free log retrieval operation binding the contract event 0x717e526bce26f8e67908004294b35133bbe2a9c7f611384cb0f484aca9223d01.
//
// Solidity: e LogPublicKeySuggested(groupId uint256, pubKeyCount uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogPublicKeySuggested(opts *bind.FilterOpts) (*DosproxyLogPublicKeySuggestedIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogPublicKeySuggested")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogPublicKeySuggestedIterator{contract: _Dosproxy.contract, event: "LogPublicKeySuggested", logs: logs, sub: sub}, nil
}

// WatchLogPublicKeySuggested is a free log subscription operation binding the contract event 0x717e526bce26f8e67908004294b35133bbe2a9c7f611384cb0f484aca9223d01.
//
// Solidity: e LogPublicKeySuggested(groupId uint256, pubKeyCount uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogPublicKeySuggested(opts *bind.WatchOpts, sink chan<- *DosproxyLogPublicKeySuggested) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogPublicKeySuggested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogPublicKeySuggested)
				if err := _Dosproxy.contract.UnpackLog(event, "LogPublicKeySuggested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogRegisteredNewPendingNodeIterator is returned from FilterLogRegisteredNewPendingNode and is used to iterate over the raw logs and unpacked data for LogRegisteredNewPendingNode events raised by the Dosproxy contract.
type DosproxyLogRegisteredNewPendingNodeIterator struct {
	Event *DosproxyLogRegisteredNewPendingNode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogRegisteredNewPendingNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogRegisteredNewPendingNode)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogRegisteredNewPendingNode)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogRegisteredNewPendingNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogRegisteredNewPendingNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogRegisteredNewPendingNode represents a LogRegisteredNewPendingNode event raised by the Dosproxy contract.
type DosproxyLogRegisteredNewPendingNode struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRegisteredNewPendingNode is a free log retrieval operation binding the contract event 0x707a6d64786780aac9cd0c5813ea04241eb135ddd2280c06eea6719afb913151.
//
// Solidity: e LogRegisteredNewPendingNode(node address)
func (_Dosproxy *DosproxyFilterer) FilterLogRegisteredNewPendingNode(opts *bind.FilterOpts) (*DosproxyLogRegisteredNewPendingNodeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogRegisteredNewPendingNode")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogRegisteredNewPendingNodeIterator{contract: _Dosproxy.contract, event: "LogRegisteredNewPendingNode", logs: logs, sub: sub}, nil
}

// WatchLogRegisteredNewPendingNode is a free log subscription operation binding the contract event 0x707a6d64786780aac9cd0c5813ea04241eb135ddd2280c06eea6719afb913151.
//
// Solidity: e LogRegisteredNewPendingNode(node address)
func (_Dosproxy *DosproxyFilterer) WatchLogRegisteredNewPendingNode(opts *bind.WatchOpts, sink chan<- *DosproxyLogRegisteredNewPendingNode) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogRegisteredNewPendingNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogRegisteredNewPendingNode)
				if err := _Dosproxy.contract.UnpackLog(event, "LogRegisteredNewPendingNode", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogRequestFromNonExistentUCIterator is returned from FilterLogRequestFromNonExistentUC and is used to iterate over the raw logs and unpacked data for LogRequestFromNonExistentUC events raised by the Dosproxy contract.
type DosproxyLogRequestFromNonExistentUCIterator struct {
	Event *DosproxyLogRequestFromNonExistentUC // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogRequestFromNonExistentUCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogRequestFromNonExistentUC)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogRequestFromNonExistentUC)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogRequestFromNonExistentUCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogRequestFromNonExistentUCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogRequestFromNonExistentUC represents a LogRequestFromNonExistentUC event raised by the Dosproxy contract.
type DosproxyLogRequestFromNonExistentUC struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogRequestFromNonExistentUC is a free log retrieval operation binding the contract event 0x40d87958cd48e8b698a94f35390a9020a9127528227647da089cc6bfd7931f42.
//
// Solidity: e LogRequestFromNonExistentUC()
func (_Dosproxy *DosproxyFilterer) FilterLogRequestFromNonExistentUC(opts *bind.FilterOpts) (*DosproxyLogRequestFromNonExistentUCIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogRequestFromNonExistentUC")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogRequestFromNonExistentUCIterator{contract: _Dosproxy.contract, event: "LogRequestFromNonExistentUC", logs: logs, sub: sub}, nil
}

// WatchLogRequestFromNonExistentUC is a free log subscription operation binding the contract event 0x40d87958cd48e8b698a94f35390a9020a9127528227647da089cc6bfd7931f42.
//
// Solidity: e LogRequestFromNonExistentUC()
func (_Dosproxy *DosproxyFilterer) WatchLogRequestFromNonExistentUC(opts *bind.WatchOpts, sink chan<- *DosproxyLogRequestFromNonExistentUC) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogRequestFromNonExistentUC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogRequestFromNonExistentUC)
				if err := _Dosproxy.contract.UnpackLog(event, "LogRequestFromNonExistentUC", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogRequestUserRandomIterator is returned from FilterLogRequestUserRandom and is used to iterate over the raw logs and unpacked data for LogRequestUserRandom events raised by the Dosproxy contract.
type DosproxyLogRequestUserRandomIterator struct {
	Event *DosproxyLogRequestUserRandom // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogRequestUserRandomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogRequestUserRandom)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogRequestUserRandom)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogRequestUserRandomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogRequestUserRandomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogRequestUserRandom represents a LogRequestUserRandom event raised by the Dosproxy contract.
type DosproxyLogRequestUserRandom struct {
	RequestId            *big.Int
	LastSystemRandomness *big.Int
	UserSeed             *big.Int
	DispatchedGroupId    *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLogRequestUserRandom is a free log retrieval operation binding the contract event 0xd587179d80d44e74955fa5d651db2f31b5470fcee8f9736f07ae3b24456a2cf0.
//
// Solidity: e LogRequestUserRandom(requestId uint256, lastSystemRandomness uint256, userSeed uint256, dispatchedGroupId uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogRequestUserRandom(opts *bind.FilterOpts) (*DosproxyLogRequestUserRandomIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogRequestUserRandom")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogRequestUserRandomIterator{contract: _Dosproxy.contract, event: "LogRequestUserRandom", logs: logs, sub: sub}, nil
}

// WatchLogRequestUserRandom is a free log subscription operation binding the contract event 0xd587179d80d44e74955fa5d651db2f31b5470fcee8f9736f07ae3b24456a2cf0.
//
// Solidity: e LogRequestUserRandom(requestId uint256, lastSystemRandomness uint256, userSeed uint256, dispatchedGroupId uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogRequestUserRandom(opts *bind.WatchOpts, sink chan<- *DosproxyLogRequestUserRandom) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogRequestUserRandom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogRequestUserRandom)
				if err := _Dosproxy.contract.UnpackLog(event, "LogRequestUserRandom", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogUnRegisteredNewPendingNodeIterator is returned from FilterLogUnRegisteredNewPendingNode and is used to iterate over the raw logs and unpacked data for LogUnRegisteredNewPendingNode events raised by the Dosproxy contract.
type DosproxyLogUnRegisteredNewPendingNodeIterator struct {
	Event *DosproxyLogUnRegisteredNewPendingNode // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogUnRegisteredNewPendingNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogUnRegisteredNewPendingNode)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogUnRegisteredNewPendingNode)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogUnRegisteredNewPendingNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogUnRegisteredNewPendingNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogUnRegisteredNewPendingNode represents a LogUnRegisteredNewPendingNode event raised by the Dosproxy contract.
type DosproxyLogUnRegisteredNewPendingNode struct {
	Node           common.Address
	UnregisterFrom uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogUnRegisteredNewPendingNode is a free log retrieval operation binding the contract event 0xaa40dce54d678a8a16fc3cf106c1ddf0b34b66a43c7a365af3268c63bb95fead.
//
// Solidity: e LogUnRegisteredNewPendingNode(node address, unregisterFrom uint8)
func (_Dosproxy *DosproxyFilterer) FilterLogUnRegisteredNewPendingNode(opts *bind.FilterOpts) (*DosproxyLogUnRegisteredNewPendingNodeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogUnRegisteredNewPendingNode")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogUnRegisteredNewPendingNodeIterator{contract: _Dosproxy.contract, event: "LogUnRegisteredNewPendingNode", logs: logs, sub: sub}, nil
}

// WatchLogUnRegisteredNewPendingNode is a free log subscription operation binding the contract event 0xaa40dce54d678a8a16fc3cf106c1ddf0b34b66a43c7a365af3268c63bb95fead.
//
// Solidity: e LogUnRegisteredNewPendingNode(node address, unregisterFrom uint8)
func (_Dosproxy *DosproxyFilterer) WatchLogUnRegisteredNewPendingNode(opts *bind.WatchOpts, sink chan<- *DosproxyLogUnRegisteredNewPendingNode) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogUnRegisteredNewPendingNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogUnRegisteredNewPendingNode)
				if err := _Dosproxy.contract.UnpackLog(event, "LogUnRegisteredNewPendingNode", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogUpdateRandomIterator is returned from FilterLogUpdateRandom and is used to iterate over the raw logs and unpacked data for LogUpdateRandom events raised by the Dosproxy contract.
type DosproxyLogUpdateRandomIterator struct {
	Event *DosproxyLogUpdateRandom // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogUpdateRandomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogUpdateRandom)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogUpdateRandom)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogUpdateRandomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogUpdateRandomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogUpdateRandom represents a LogUpdateRandom event raised by the Dosproxy contract.
type DosproxyLogUpdateRandom struct {
	LastRandomness    *big.Int
	DispatchedGroupId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogUpdateRandom is a free log retrieval operation binding the contract event 0xfaa99731d2c5abc7ee76b2e31b6b7e293a30e1e2274f59396a7a59cabd5eadf3.
//
// Solidity: e LogUpdateRandom(lastRandomness uint256, dispatchedGroupId uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogUpdateRandom(opts *bind.FilterOpts) (*DosproxyLogUpdateRandomIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogUpdateRandom")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogUpdateRandomIterator{contract: _Dosproxy.contract, event: "LogUpdateRandom", logs: logs, sub: sub}, nil
}

// WatchLogUpdateRandom is a free log subscription operation binding the contract event 0xfaa99731d2c5abc7ee76b2e31b6b7e293a30e1e2274f59396a7a59cabd5eadf3.
//
// Solidity: e LogUpdateRandom(lastRandomness uint256, dispatchedGroupId uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogUpdateRandom(opts *bind.WatchOpts, sink chan<- *DosproxyLogUpdateRandom) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogUpdateRandom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogUpdateRandom)
				if err := _Dosproxy.contract.UnpackLog(event, "LogUpdateRandom", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogUrlIterator is returned from FilterLogUrl and is used to iterate over the raw logs and unpacked data for LogUrl events raised by the Dosproxy contract.
type DosproxyLogUrlIterator struct {
	Event *DosproxyLogUrl // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogUrlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogUrl)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogUrl)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogUrlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogUrlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogUrl represents a LogUrl event raised by the Dosproxy contract.
type DosproxyLogUrl struct {
	QueryId           *big.Int
	Timeout           *big.Int
	DataSource        string
	Selector          string
	Randomness        *big.Int
	DispatchedGroupId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogUrl is a free log retrieval operation binding the contract event 0x05e1614af4efb13caeba2369a57a05ee5830f33364f82e2c899fd5710cb56ef3.
//
// Solidity: e LogUrl(queryId uint256, timeout uint256, dataSource string, selector string, randomness uint256, dispatchedGroupId uint256)
func (_Dosproxy *DosproxyFilterer) FilterLogUrl(opts *bind.FilterOpts) (*DosproxyLogUrlIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogUrl")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogUrlIterator{contract: _Dosproxy.contract, event: "LogUrl", logs: logs, sub: sub}, nil
}

// WatchLogUrl is a free log subscription operation binding the contract event 0x05e1614af4efb13caeba2369a57a05ee5830f33364f82e2c899fd5710cb56ef3.
//
// Solidity: e LogUrl(queryId uint256, timeout uint256, dataSource string, selector string, randomness uint256, dispatchedGroupId uint256)
func (_Dosproxy *DosproxyFilterer) WatchLogUrl(opts *bind.WatchOpts, sink chan<- *DosproxyLogUrl) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogUrl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogUrl)
				if err := _Dosproxy.contract.UnpackLog(event, "LogUrl", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyLogValidationResultIterator is returned from FilterLogValidationResult and is used to iterate over the raw logs and unpacked data for LogValidationResult events raised by the Dosproxy contract.
type DosproxyLogValidationResultIterator struct {
	Event *DosproxyLogValidationResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyLogValidationResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyLogValidationResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyLogValidationResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyLogValidationResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyLogValidationResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyLogValidationResult represents a LogValidationResult event raised by the Dosproxy contract.
type DosproxyLogValidationResult struct {
	TrafficType uint8
	TrafficId   *big.Int
	Message     []byte
	Signature   [2]*big.Int
	PubKey      [4]*big.Int
	Pass        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogValidationResult is a free log retrieval operation binding the contract event 0xd33c44f7ce166bcd2616c8f4d811261d4d24e1af815f78614683e0c7682c42a5.
//
// Solidity: e LogValidationResult(trafficType uint8, trafficId uint256, message bytes, signature uint256[2], pubKey uint256[4], pass bool)
func (_Dosproxy *DosproxyFilterer) FilterLogValidationResult(opts *bind.FilterOpts) (*DosproxyLogValidationResultIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogValidationResult")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogValidationResultIterator{contract: _Dosproxy.contract, event: "LogValidationResult", logs: logs, sub: sub}, nil
}

// WatchLogValidationResult is a free log subscription operation binding the contract event 0xd33c44f7ce166bcd2616c8f4d811261d4d24e1af815f78614683e0c7682c42a5.
//
// Solidity: e LogValidationResult(trafficType uint8, trafficId uint256, message bytes, signature uint256[2], pubKey uint256[4], pass bool)
func (_Dosproxy *DosproxyFilterer) WatchLogValidationResult(opts *bind.WatchOpts, sink chan<- *DosproxyLogValidationResult) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "LogValidationResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyLogValidationResult)
				if err := _Dosproxy.contract.UnpackLog(event, "LogValidationResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Dosproxy contract.
type DosproxyOwnershipRenouncedIterator struct {
	Event *DosproxyOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyOwnershipRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyOwnershipRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyOwnershipRenounced represents a OwnershipRenounced event raised by the Dosproxy contract.
type DosproxyOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Dosproxy *DosproxyFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DosproxyOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DosproxyOwnershipRenouncedIterator{contract: _Dosproxy.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Dosproxy *DosproxyFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DosproxyOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyOwnershipRenounced)
				if err := _Dosproxy.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dosproxy contract.
type DosproxyOwnershipTransferredIterator struct {
	Event *DosproxyOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyOwnershipTransferred represents a OwnershipTransferred event raised by the Dosproxy contract.
type DosproxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Dosproxy *DosproxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DosproxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DosproxyOwnershipTransferredIterator{contract: _Dosproxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Dosproxy *DosproxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DosproxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyOwnershipTransferred)
				if err := _Dosproxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdateBootstrapCommitDurationIterator is returned from FilterUpdateBootstrapCommitDuration and is used to iterate over the raw logs and unpacked data for UpdateBootstrapCommitDuration events raised by the Dosproxy contract.
type DosproxyUpdateBootstrapCommitDurationIterator struct {
	Event *DosproxyUpdateBootstrapCommitDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdateBootstrapCommitDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdateBootstrapCommitDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdateBootstrapCommitDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdateBootstrapCommitDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdateBootstrapCommitDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdateBootstrapCommitDuration represents a UpdateBootstrapCommitDuration event raised by the Dosproxy contract.
type DosproxyUpdateBootstrapCommitDuration struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateBootstrapCommitDuration is a free log retrieval operation binding the contract event 0xbdae601725e6f9108b15103969d6a682c09f7d87ec505e67ceee0baac2c550c8.
//
// Solidity: e UpdateBootstrapCommitDuration(oldDuration uint256, newDuration uint256)
func (_Dosproxy *DosproxyFilterer) FilterUpdateBootstrapCommitDuration(opts *bind.FilterOpts) (*DosproxyUpdateBootstrapCommitDurationIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdateBootstrapCommitDuration")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdateBootstrapCommitDurationIterator{contract: _Dosproxy.contract, event: "UpdateBootstrapCommitDuration", logs: logs, sub: sub}, nil
}

// WatchUpdateBootstrapCommitDuration is a free log subscription operation binding the contract event 0xbdae601725e6f9108b15103969d6a682c09f7d87ec505e67ceee0baac2c550c8.
//
// Solidity: e UpdateBootstrapCommitDuration(oldDuration uint256, newDuration uint256)
func (_Dosproxy *DosproxyFilterer) WatchUpdateBootstrapCommitDuration(opts *bind.WatchOpts, sink chan<- *DosproxyUpdateBootstrapCommitDuration) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdateBootstrapCommitDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdateBootstrapCommitDuration)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdateBootstrapCommitDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdateBootstrapGroupsIterator is returned from FilterUpdateBootstrapGroups and is used to iterate over the raw logs and unpacked data for UpdateBootstrapGroups events raised by the Dosproxy contract.
type DosproxyUpdateBootstrapGroupsIterator struct {
	Event *DosproxyUpdateBootstrapGroups // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdateBootstrapGroupsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdateBootstrapGroups)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdateBootstrapGroups)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdateBootstrapGroupsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdateBootstrapGroupsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdateBootstrapGroups represents a UpdateBootstrapGroups event raised by the Dosproxy contract.
type DosproxyUpdateBootstrapGroups struct {
	OldSize *big.Int
	NewSize *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateBootstrapGroups is a free log retrieval operation binding the contract event 0xf9da68cf2452df09a5c96de5099bed44a4f40947e5dfbac9fc0a0775be49675b.
//
// Solidity: e UpdateBootstrapGroups(oldSize uint256, newSize uint256)
func (_Dosproxy *DosproxyFilterer) FilterUpdateBootstrapGroups(opts *bind.FilterOpts) (*DosproxyUpdateBootstrapGroupsIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdateBootstrapGroups")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdateBootstrapGroupsIterator{contract: _Dosproxy.contract, event: "UpdateBootstrapGroups", logs: logs, sub: sub}, nil
}

// WatchUpdateBootstrapGroups is a free log subscription operation binding the contract event 0xf9da68cf2452df09a5c96de5099bed44a4f40947e5dfbac9fc0a0775be49675b.
//
// Solidity: e UpdateBootstrapGroups(oldSize uint256, newSize uint256)
func (_Dosproxy *DosproxyFilterer) WatchUpdateBootstrapGroups(opts *bind.WatchOpts, sink chan<- *DosproxyUpdateBootstrapGroups) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdateBootstrapGroups")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdateBootstrapGroups)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdateBootstrapGroups", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdateBootstrapRevealDurationIterator is returned from FilterUpdateBootstrapRevealDuration and is used to iterate over the raw logs and unpacked data for UpdateBootstrapRevealDuration events raised by the Dosproxy contract.
type DosproxyUpdateBootstrapRevealDurationIterator struct {
	Event *DosproxyUpdateBootstrapRevealDuration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdateBootstrapRevealDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdateBootstrapRevealDuration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdateBootstrapRevealDuration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdateBootstrapRevealDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdateBootstrapRevealDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdateBootstrapRevealDuration represents a UpdateBootstrapRevealDuration event raised by the Dosproxy contract.
type DosproxyUpdateBootstrapRevealDuration struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateBootstrapRevealDuration is a free log retrieval operation binding the contract event 0x2e2857fe2c7b1963919b23c68d0074aac750303e8f14d18d0115cc792668cdb6.
//
// Solidity: e UpdateBootstrapRevealDuration(oldDuration uint256, newDuration uint256)
func (_Dosproxy *DosproxyFilterer) FilterUpdateBootstrapRevealDuration(opts *bind.FilterOpts) (*DosproxyUpdateBootstrapRevealDurationIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdateBootstrapRevealDuration")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdateBootstrapRevealDurationIterator{contract: _Dosproxy.contract, event: "UpdateBootstrapRevealDuration", logs: logs, sub: sub}, nil
}

// WatchUpdateBootstrapRevealDuration is a free log subscription operation binding the contract event 0x2e2857fe2c7b1963919b23c68d0074aac750303e8f14d18d0115cc792668cdb6.
//
// Solidity: e UpdateBootstrapRevealDuration(oldDuration uint256, newDuration uint256)
func (_Dosproxy *DosproxyFilterer) WatchUpdateBootstrapRevealDuration(opts *bind.WatchOpts, sink chan<- *DosproxyUpdateBootstrapRevealDuration) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdateBootstrapRevealDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdateBootstrapRevealDuration)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdateBootstrapRevealDuration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdateGroupMaturityPeriodIterator is returned from FilterUpdateGroupMaturityPeriod and is used to iterate over the raw logs and unpacked data for UpdateGroupMaturityPeriod events raised by the Dosproxy contract.
type DosproxyUpdateGroupMaturityPeriodIterator struct {
	Event *DosproxyUpdateGroupMaturityPeriod // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdateGroupMaturityPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdateGroupMaturityPeriod)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdateGroupMaturityPeriod)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdateGroupMaturityPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdateGroupMaturityPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdateGroupMaturityPeriod represents a UpdateGroupMaturityPeriod event raised by the Dosproxy contract.
type DosproxyUpdateGroupMaturityPeriod struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateGroupMaturityPeriod is a free log retrieval operation binding the contract event 0x96a027b03aa3233feda42c74f270026db98f223e64b4df4b81231da93bac04b3.
//
// Solidity: e UpdateGroupMaturityPeriod(oldPeriod uint256, newPeriod uint256)
func (_Dosproxy *DosproxyFilterer) FilterUpdateGroupMaturityPeriod(opts *bind.FilterOpts) (*DosproxyUpdateGroupMaturityPeriodIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdateGroupMaturityPeriod")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdateGroupMaturityPeriodIterator{contract: _Dosproxy.contract, event: "UpdateGroupMaturityPeriod", logs: logs, sub: sub}, nil
}

// WatchUpdateGroupMaturityPeriod is a free log subscription operation binding the contract event 0x96a027b03aa3233feda42c74f270026db98f223e64b4df4b81231da93bac04b3.
//
// Solidity: e UpdateGroupMaturityPeriod(oldPeriod uint256, newPeriod uint256)
func (_Dosproxy *DosproxyFilterer) WatchUpdateGroupMaturityPeriod(opts *bind.WatchOpts, sink chan<- *DosproxyUpdateGroupMaturityPeriod) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdateGroupMaturityPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdateGroupMaturityPeriod)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdateGroupMaturityPeriod", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdateGroupSizeIterator is returned from FilterUpdateGroupSize and is used to iterate over the raw logs and unpacked data for UpdateGroupSize events raised by the Dosproxy contract.
type DosproxyUpdateGroupSizeIterator struct {
	Event *DosproxyUpdateGroupSize // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdateGroupSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdateGroupSize)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdateGroupSize)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdateGroupSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdateGroupSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdateGroupSize represents a UpdateGroupSize event raised by the Dosproxy contract.
type DosproxyUpdateGroupSize struct {
	OldSize *big.Int
	NewSize *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateGroupSize is a free log retrieval operation binding the contract event 0x28eb4f48ae7c6c17a714b104832bdd949ebd0a984d37f4893d6cb91f92a8ae57.
//
// Solidity: e UpdateGroupSize(oldSize uint256, newSize uint256)
func (_Dosproxy *DosproxyFilterer) FilterUpdateGroupSize(opts *bind.FilterOpts) (*DosproxyUpdateGroupSizeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdateGroupSize")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdateGroupSizeIterator{contract: _Dosproxy.contract, event: "UpdateGroupSize", logs: logs, sub: sub}, nil
}

// WatchUpdateGroupSize is a free log subscription operation binding the contract event 0x28eb4f48ae7c6c17a714b104832bdd949ebd0a984d37f4893d6cb91f92a8ae57.
//
// Solidity: e UpdateGroupSize(oldSize uint256, newSize uint256)
func (_Dosproxy *DosproxyFilterer) WatchUpdateGroupSize(opts *bind.WatchOpts, sink chan<- *DosproxyUpdateGroupSize) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdateGroupSize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdateGroupSize)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdateGroupSize", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdateLifeDiversityIterator is returned from FilterUpdateLifeDiversity and is used to iterate over the raw logs and unpacked data for UpdateLifeDiversity events raised by the Dosproxy contract.
type DosproxyUpdateLifeDiversityIterator struct {
	Event *DosproxyUpdateLifeDiversity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdateLifeDiversityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdateLifeDiversity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdateLifeDiversity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdateLifeDiversityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdateLifeDiversityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdateLifeDiversity represents a UpdateLifeDiversity event raised by the Dosproxy contract.
type DosproxyUpdateLifeDiversity struct {
	LifeDiversity *big.Int
	NewDiversity  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateLifeDiversity is a free log retrieval operation binding the contract event 0x978a29592cb150802d04222f78a83519049bde42bb2e86e17250efde5820c687.
//
// Solidity: e UpdateLifeDiversity(lifeDiversity uint256, newDiversity uint256)
func (_Dosproxy *DosproxyFilterer) FilterUpdateLifeDiversity(opts *bind.FilterOpts) (*DosproxyUpdateLifeDiversityIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdateLifeDiversity")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdateLifeDiversityIterator{contract: _Dosproxy.contract, event: "UpdateLifeDiversity", logs: logs, sub: sub}, nil
}

// WatchUpdateLifeDiversity is a free log subscription operation binding the contract event 0x978a29592cb150802d04222f78a83519049bde42bb2e86e17250efde5820c687.
//
// Solidity: e UpdateLifeDiversity(lifeDiversity uint256, newDiversity uint256)
func (_Dosproxy *DosproxyFilterer) WatchUpdateLifeDiversity(opts *bind.WatchOpts, sink chan<- *DosproxyUpdateLifeDiversity) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdateLifeDiversity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdateLifeDiversity)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdateLifeDiversity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdatePendingGroupMaxLifeIterator is returned from FilterUpdatePendingGroupMaxLife and is used to iterate over the raw logs and unpacked data for UpdatePendingGroupMaxLife events raised by the Dosproxy contract.
type DosproxyUpdatePendingGroupMaxLifeIterator struct {
	Event *DosproxyUpdatePendingGroupMaxLife // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdatePendingGroupMaxLifeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdatePendingGroupMaxLife)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdatePendingGroupMaxLife)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdatePendingGroupMaxLifeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdatePendingGroupMaxLifeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdatePendingGroupMaxLife represents a UpdatePendingGroupMaxLife event raised by the Dosproxy contract.
type DosproxyUpdatePendingGroupMaxLife struct {
	OldLifeBlocks *big.Int
	NewLifeBlocks *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatePendingGroupMaxLife is a free log retrieval operation binding the contract event 0xfc644126d2177f897a0e09f46bf2678f9577840113d685f4a56bd9e4d48d012c.
//
// Solidity: e UpdatePendingGroupMaxLife(oldLifeBlocks uint256, newLifeBlocks uint256)
func (_Dosproxy *DosproxyFilterer) FilterUpdatePendingGroupMaxLife(opts *bind.FilterOpts) (*DosproxyUpdatePendingGroupMaxLifeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdatePendingGroupMaxLife")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdatePendingGroupMaxLifeIterator{contract: _Dosproxy.contract, event: "UpdatePendingGroupMaxLife", logs: logs, sub: sub}, nil
}

// WatchUpdatePendingGroupMaxLife is a free log subscription operation binding the contract event 0xfc644126d2177f897a0e09f46bf2678f9577840113d685f4a56bd9e4d48d012c.
//
// Solidity: e UpdatePendingGroupMaxLife(oldLifeBlocks uint256, newLifeBlocks uint256)
func (_Dosproxy *DosproxyFilterer) WatchUpdatePendingGroupMaxLife(opts *bind.WatchOpts, sink chan<- *DosproxyUpdatePendingGroupMaxLife) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdatePendingGroupMaxLife")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdatePendingGroupMaxLife)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdatePendingGroupMaxLife", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdateProxyFundIterator is returned from FilterUpdateProxyFund and is used to iterate over the raw logs and unpacked data for UpdateProxyFund events raised by the Dosproxy contract.
type DosproxyUpdateProxyFundIterator struct {
	Event *DosproxyUpdateProxyFund // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdateProxyFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdateProxyFund)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdateProxyFund)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdateProxyFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdateProxyFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdateProxyFund represents a UpdateProxyFund event raised by the Dosproxy contract.
type DosproxyUpdateProxyFund struct {
	OldFundAddr  common.Address
	NewFundAddr  common.Address
	OldTokenAddr common.Address
	NewTokenAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateProxyFund is a free log retrieval operation binding the contract event 0x2ae8e7023c1978c8540df9af00881f2f942d6e7233463a3f0def2b6e57e6dd60.
//
// Solidity: e UpdateProxyFund(oldFundAddr address, newFundAddr address, oldTokenAddr address, newTokenAddr address)
func (_Dosproxy *DosproxyFilterer) FilterUpdateProxyFund(opts *bind.FilterOpts) (*DosproxyUpdateProxyFundIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdateProxyFund")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdateProxyFundIterator{contract: _Dosproxy.contract, event: "UpdateProxyFund", logs: logs, sub: sub}, nil
}

// WatchUpdateProxyFund is a free log subscription operation binding the contract event 0x2ae8e7023c1978c8540df9af00881f2f942d6e7233463a3f0def2b6e57e6dd60.
//
// Solidity: e UpdateProxyFund(oldFundAddr address, newFundAddr address, oldTokenAddr address, newTokenAddr address)
func (_Dosproxy *DosproxyFilterer) WatchUpdateProxyFund(opts *bind.WatchOpts, sink chan<- *DosproxyUpdateProxyFund) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdateProxyFund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdateProxyFund)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdateProxyFund", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdateSystemRandomHardLimitIterator is returned from FilterUpdateSystemRandomHardLimit and is used to iterate over the raw logs and unpacked data for UpdateSystemRandomHardLimit events raised by the Dosproxy contract.
type DosproxyUpdateSystemRandomHardLimitIterator struct {
	Event *DosproxyUpdateSystemRandomHardLimit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdateSystemRandomHardLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdateSystemRandomHardLimit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdateSystemRandomHardLimit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdateSystemRandomHardLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdateSystemRandomHardLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdateSystemRandomHardLimit represents a UpdateSystemRandomHardLimit event raised by the Dosproxy contract.
type DosproxyUpdateSystemRandomHardLimit struct {
	OldLimit *big.Int
	NewLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateSystemRandomHardLimit is a free log retrieval operation binding the contract event 0xdb95a2fbbee34de5822459ce9edd234f70f321a1cc2395b2dc902b69e1f8093b.
//
// Solidity: e UpdateSystemRandomHardLimit(oldLimit uint256, newLimit uint256)
func (_Dosproxy *DosproxyFilterer) FilterUpdateSystemRandomHardLimit(opts *bind.FilterOpts) (*DosproxyUpdateSystemRandomHardLimitIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdateSystemRandomHardLimit")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdateSystemRandomHardLimitIterator{contract: _Dosproxy.contract, event: "UpdateSystemRandomHardLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateSystemRandomHardLimit is a free log subscription operation binding the contract event 0xdb95a2fbbee34de5822459ce9edd234f70f321a1cc2395b2dc902b69e1f8093b.
//
// Solidity: e UpdateSystemRandomHardLimit(oldLimit uint256, newLimit uint256)
func (_Dosproxy *DosproxyFilterer) WatchUpdateSystemRandomHardLimit(opts *bind.WatchOpts, sink chan<- *DosproxyUpdateSystemRandomHardLimit) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdateSystemRandomHardLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdateSystemRandomHardLimit)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdateSystemRandomHardLimit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DosproxyUpdatebootstrapStartThresholdIterator is returned from FilterUpdatebootstrapStartThreshold and is used to iterate over the raw logs and unpacked data for UpdatebootstrapStartThreshold events raised by the Dosproxy contract.
type DosproxyUpdatebootstrapStartThresholdIterator struct {
	Event *DosproxyUpdatebootstrapStartThreshold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DosproxyUpdatebootstrapStartThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosproxyUpdatebootstrapStartThreshold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DosproxyUpdatebootstrapStartThreshold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DosproxyUpdatebootstrapStartThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosproxyUpdatebootstrapStartThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosproxyUpdatebootstrapStartThreshold represents a UpdatebootstrapStartThreshold event raised by the Dosproxy contract.
type DosproxyUpdatebootstrapStartThreshold struct {
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatebootstrapStartThreshold is a free log retrieval operation binding the contract event 0x1fa02fb08d726e79971d6de0ee1e2f637f068fed6d3fb859a1765e666bb19307.
//
// Solidity: e UpdatebootstrapStartThreshold(oldThreshold uint256, newThreshold uint256)
func (_Dosproxy *DosproxyFilterer) FilterUpdatebootstrapStartThreshold(opts *bind.FilterOpts) (*DosproxyUpdatebootstrapStartThresholdIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "UpdatebootstrapStartThreshold")
	if err != nil {
		return nil, err
	}
	return &DosproxyUpdatebootstrapStartThresholdIterator{contract: _Dosproxy.contract, event: "UpdatebootstrapStartThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdatebootstrapStartThreshold is a free log subscription operation binding the contract event 0x1fa02fb08d726e79971d6de0ee1e2f637f068fed6d3fb859a1765e666bb19307.
//
// Solidity: e UpdatebootstrapStartThreshold(oldThreshold uint256, newThreshold uint256)
func (_Dosproxy *DosproxyFilterer) WatchUpdatebootstrapStartThreshold(opts *bind.WatchOpts, sink chan<- *DosproxyUpdatebootstrapStartThreshold) (event.Subscription, error) {

	logs, sub, err := _Dosproxy.contract.WatchLogs(opts, "UpdatebootstrapStartThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosproxyUpdatebootstrapStartThreshold)
				if err := _Dosproxy.contract.UnpackLog(event, "UpdatebootstrapStartThreshold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
