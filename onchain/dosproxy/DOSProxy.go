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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DosproxyABI is the input ABI used to generate the binding from.
const DosproxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyFundsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyFundsTokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blkNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"callbackAddr\",\"type\":\"address\"}],\"name\":\"LogCallbackTriggeredFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"LogGroupDissolve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nodeId\",\"type\":\"address[]\"}],\"name\":\"LogGrouping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingNodePool\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupsize\",\"type\":\"uint256\"}],\"name\":\"LogGroupingInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numPendingNodes\",\"type\":\"uint256\"}],\"name\":\"LogInsufficientPendingNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numWorkingGroups\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numPendingGroups\",\"type\":\"uint256\"}],\"name\":\"LogInsufficientWorkingGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"}],\"name\":\"LogMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"LogNoPendingGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"LogNonContractCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"invalidSelector\",\"type\":\"string\"}],\"name\":\"LogNonSupportedType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"}],\"name\":\"LogPendingGroupRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"pubKey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numWorkingGroups\",\"type\":\"uint256\"}],\"name\":\"LogPublicKeyAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubKeyCount\",\"type\":\"uint256\"}],\"name\":\"LogPublicKeySuggested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"LogRegisteredNewPendingNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogRequestFromNonExistentUC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastSystemRandomness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userSeed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dispatchedGroupId\",\"type\":\"uint256\"}],\"name\":\"LogRequestUserRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"unregisterFrom\",\"type\":\"uint8\"}],\"name\":\"LogUnRegisteredNewPendingNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastRandomness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dispatchedGroupId\",\"type\":\"uint256\"}],\"name\":\"LogUpdateRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"queryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dataSource\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dispatchedGroupId\",\"type\":\"uint256\"}],\"name\":\"LogUrl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"trafficType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"trafficId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"pubKey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pass\",\"type\":\"bool\"}],\"name\":\"LogValidationResult\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rndSeed\",\"type\":\"uint256\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addToGuardianList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressBridge\",\"outputs\":[{\"internalType\":\"contractDOSAddressBridgeInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapCommitDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapEndBlk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapRevealDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bootstrapStartThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cachedUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"expiredWorkingGroupIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExpiredWorkingGroupSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWorkingGroupSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"groupMaturityPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"groupSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"groupToPick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"guardianListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initBlkN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastFormGrpReqId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRandomness\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lifeDiversity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loopLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodeToGroupIdList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numPendingGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numPendingNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingGroupList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingGroupMaxLife\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingGroupTail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingGroups\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlkNum\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingNodeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingNodeTail\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxyFundsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxyFundsTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dataSource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"refreshSystemRandomHardLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"groupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"suggestedPubKey\",\"type\":\"uint256[4]\"}],\"name\":\"registerGroupPubKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerNewNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayRespondLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"removeFromGuardianList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userSeed\",\"type\":\"uint256\"}],\"name\":\"requestRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"resetOnRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"setBootstrapCommitDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"setBootstrapRevealDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"setBootstrapStartThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"setGroupMaturityPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSize\",\"type\":\"uint256\"}],\"name\":\"setGroupSize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDiversity\",\"type\":\"uint256\"}],\"name\":\"setLifeDiversity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLife\",\"type\":\"uint256\"}],\"name\":\"setPendingGroupMaxLife\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFund\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFundToken\",\"type\":\"address\"}],\"name\":\"setProxyFund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"setRelayRespondLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"setSystemRandomHardLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cid\",\"type\":\"uint256\"}],\"name\":\"signalBootstrap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signalGroupDissolve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signalGroupFormation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"signalRandom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"signalUnregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"trafficType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"sig\",\"type\":\"uint256[2]\"}],\"name\":\"triggerCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unregisterNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"sig\",\"type\":\"uint256[2]\"}],\"name\":\"updateRandomness\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"workingGroupIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DosproxyBin is the compiled bytecode used for deploying new contracts.
var DosproxyBin = "0x6080604052606460039081556105a06004908155619d8060055561438060065560326007556008556002600955600a556028600b819055600c908155600d556014601b553480156200005057600080fd5b506040516200629a3803806200629a833981810160405260608110156200007657600080fd5b50805160208083015160409384015160008054336001600160a01b0319918216178255436001908155918290527f4155c2f711f2cdd34f8262ab8fb9b7020a700fe7b6948222152f7670d1fdf34d805482168317905560148054821683179055601c85527f6de76108811faf2f94afbe5ac6c98e8393206cd093932de1fbfd61bbeec43a02829055601d919091556010805482166001600160a01b038088169190911791829055601180548416828716179055601280549093168185161790925586516313a4cbcb60e31b81529651959693959294911692639d265e589260048281019392829003018186803b1580156200017057600080fd5b505afa15801562000185573d6000803e3d6000fd5b505050506040513d60208110156200019c57600080fd5b50516011546012546040805163b73a3f8f60e01b81526001600160a01b039384166004820152918316602483015251919092169163b73a3f8f91604480830192600092919082900301818387803b158015620001f757600080fd5b505af11580156200020c573d6000803e3d6000fd5b5050505050505061607780620002236000396000f3fe608060405234801561001057600080fd5b50600436106103af5760003560e01c8063863bc0a1116101f4578063ba419d0c1161011a578063df37c617116100ad578063f2a3072d1161007c578063f2a3072d146109a6578063f41a1587146109ae578063f90ce5ba146109b6578063fc84dde4146109be576103af565b8063df37c61714610971578063e1ff791614610979578063ef112dfc14610996578063efde068b1461099e576103af565b8063ca2aabcc116100e9578063ca2aabcc146108f0578063d18c81b71461090d578063d79351b214610943578063dd6ceddf14610969576103af565b8063ba419d0c14610879578063c457aa8f14610881578063c58ebe1c1461089e578063c7c3f4a5146108c4576103af565b8063a60b007d11610192578063b7fb8fd711610161578063b7fb8fd71461075b578063b836ccea14610832578063b8fa82e014610854578063b9424b3514610871576103af565b8063a60b007d14610708578063aeb3da731461072e578063b45ef79d14610736578063b537226414610753576103af565b8063962ba8a4116101ce578063962ba8a4146106be57806399ca2d30146106c65780639e718573146106ce578063a54fb00e146106eb576103af565b8063863bc0a114610691578063925fc6c91461069957806395071cf6146106b6576103af565b806340e4a5af116102d95780636c460899116102775780637c1cf083116102465780637c1cf083146106535780637c48d1a0146106795780637e55acb41461068157806385ed422314610689576103af565b80636c460899146105bd57806374ad3a06146105c557806376cffa531461064357806377f101921461064b576103af565b80635be6c3af116102b35780635be6c3af146105735780635c0e159f1461057b5780635d3812041461059857806363b635ea146105b5576103af565b806340e4a5af1461050b5780634a28a74d14610539578063559ea9de14610556576103af565b806311bc531111610351578063197172031161032057806319717203146104eb57806331bf6464146104f3578063372a53cc146104fb5780633d385cf514610503576103af565b806311bc5311146104b0578063155fa82c146104b857806318a1908d146104c0578063190ca29e146104e3576103af565b806309ac86d31161038d57806309ac86d31461042c5780630dfc09cb1461044a5780630eeee5c114610467578063100063ec14610493576103af565b80630434ccd2146103b457806309011cb9146103ce578063094c361214610408575b600080fd5b6103bc6109c6565b60408051918252519081900360200190f35b6103f4600480360360208110156103e457600080fd5b50356001600160a01b03166109cc565b604080519115158252519081900360200190f35b6104106109e1565b604080516001600160a01b039092168252519081900360200190f35b6104486004803603604081101561044257600080fd5b506109f0565b005b6104486004803603602081101561046057600080fd5b5035610d36565b6103bc6004803603604081101561047d57600080fd5b506001600160a01b038135169060200135610dab565b610448600480360360208110156104a957600080fd5b5035610dc8565b6103bc610e2f565b610448610e35565b610448600480360360408110156104d657600080fd5b5080359060200135610f8e565b6103bc611253565b6103bc611259565b6103bc61125f565b6103bc611265565b6103f461126b565b6104486004803603604081101561052157600080fd5b506001600160a01b03813581169160200135166113a7565b6104486004803603602081101561054f57600080fd5b5035611583565b6104486004803603602081101561056c57600080fd5b50356115f5565b610448611667565b6104486004803603602081101561059157600080fd5b5035611811565b6103bc600480360360208110156105ae57600080fd5b5035611c18565b6103bc611c36565b6103bc611c3c565b610448600480360360a08110156105db57600080fd5b81359160ff6020820135169181019060608101604082013564010000000081111561060557600080fd5b82018360208201111561061757600080fd5b8035906020019184600183028401116401000000008311171561063957600080fd5b9193509150611c42565b610410612347565b6103bc612356565b6104486004803603602081101561066957600080fd5b50356001600160a01b031661235c565b6103bc61247c565b6103bc612482565b6103bc612488565b6103bc61248e565b610448600480360360208110156106af57600080fd5b5035612494565b6103bc612506565b6103bc61250c565b610410612512565b610448600480360360208110156106e457600080fd5b5035612521565b6103bc6004803603602081101561070157600080fd5b5035612593565b6104106004803603602081101561071e57600080fd5b50356001600160a01b03166125a5565b6104486125c0565b6103bc6004803603602081101561074c57600080fd5b503561287b565b6103bc612888565b6103bc6004803603608081101561077157600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156107a157600080fd5b8201836020820111156107b357600080fd5b803590602001918460018302840111640100000000831117156107d557600080fd5b9193909290916020810190356401000000008111156107f357600080fd5b82018360208201111561080557600080fd5b8035906020019184600183028401116401000000008311171561082757600080fd5b50909250905061288e565b610448600480360360a081101561084857600080fd5b50803590602001612f60565b6104486004803603602081101561086a57600080fd5b50356134d4565b610448613546565b6103bc613612565b6104486004803603602081101561089757600080fd5b5035613618565b610448600480360360208110156108b457600080fd5b50356001600160a01b031661368a565b6103bc600480360360408110156108da57600080fd5b506001600160a01b0381351690602001356136c2565b6104486004803603602081101561090657600080fd5b5035613b6f565b61092a6004803603602081101561092357600080fd5b5035613bd6565b6040805192835260208301919091528051918290030190f35b6104486004803603602081101561095957600080fd5b50356001600160a01b0316613bef565b6103bc613c2a565b610410613c30565b6104486004803603602081101561098f57600080fd5b5035613c3f565b6103bc614071565b6103bc614077565b6103bc61407d565b6103bc614083565b6103bc614089565b6103bc61408f565b60095481565b602b6020526000908152604090205460ff1681565b6014546001600160a01b031681565b601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a3e57600080fd5b505afa158015610a52573d6000803e3d6000fd5b505050506040513d6020811015610a6857600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b158015610ab257600080fd5b505afa158015610ac6573d6000803e3d6000fd5b505050506040513d6020811015610adc57600080fd5b5051610b26576040805162461bcd60e51b8152602060048201526014602482015273696e76616c69642d7374616b696e672d6e6f646560601b604482015290519081900360640190fd5b610bc36000602154610b3781614095565b604080518082018252863581526020808801359082015281516080810180845291929091602491839190820190839060029082845b815481526020019060010190808311610b6c57505050918352505060408051808201918290526020909201919060028481019182845b815481526020019060010190808311610ba2575050505050815250506140bf565b610bcc57610d33565b6000601f5543602090815560218054604080518535818601528585013581830152815180820383018152606082018084528151918701919091209094556010546313a4cbcb60e31b909452905191936001600160a01b0390931692639d265e5892606480840193829003018186803b158015610c4757600080fd5b505afa158015610c5b573d6000803e3d6000fd5b505050506040513d6020811015610c7157600080fd5b5051604051632100fdf760e21b8152600481018381523360248301819052606060448401908152602a8054606486018190526001600160a01b0390961695638403f7dc95889592939160849091019084908015610cf757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610cd9575b5050945050505050600060405180830381600087803b158015610d1957600080fd5b505af1158015610d2d573d6000803e3d6000fd5b50505050505b50565b6000546001600160a01b03163314610d4d57600080fd5b600a548114158015610d6157506002810615155b610da6576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b600a55565b601660209081526000928352604080842090915290825290205481565b6000546001600160a01b03163314610ddf57600080fd5b600d54811415610e2a576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b600d55565b600d5481565b610e3d6143e6565b15610f4a57604080514381523360208201528151600080516020615fba833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015610eb857600080fd5b505afa158015610ecc573d6000803e3d6000fd5b505050506040513d6020811015610ee257600080fd5b5051604080516323ff34cb60e01b815233600482015290516001600160a01b03909216916323ff34cb9160248082019260009290919082900301818387803b158015610f2d57600080fd5b505af1158015610f41573d6000803e3d6000fd5b50505050610f8c565b6040805160208082526010908201526f373796b3b93816b337b936b0ba34b7b760811b8183015290516000805160206160238339815191529181900360600190a15b565b333014610fd9576040805162461bcd60e51b81526020600482015260146024820152730756e61757468656e746963617465642d726573760641b604482015290519081900360640190fd5b6009546019541015611032576040805162461bcd60e51b815260206004820152601f60248201527f726567726f75702d6e6f742d656e6f7567682d657870697265642d7767727000604482015290519081900360640190fd5b600a54601554101561108b576040805162461bcd60e51b815260206004820152601960248201527f726567726f75702d6e6f742d656e6f7567682d702d6e6f646500000000000000604482015290519081900360640190fd5b60006022819055506000600954600101600a540290506060816040519080825280602002602001820160405280156110cd578160200160208202803883390190505b50905060005b6009548110156112205760195460408051602080820188905281830189905260608083018690528351808403909101815260809092019092528051910120600091908161111c57fe5b0690506000601760006019848154811061113257fe5b90600052602060002001548152602001908152602001600020905060008090505b600a548110156111bc5781600701818154811061116c57fe5b9060005260206000200160009054906101000a90046001600160a01b03168582600a548702018151811061119c57fe5b6001600160a01b0390921660209283029190910190910152600101611153565b5080546111cc906000600161475b565b6019805460001981019081106111de57fe5b9060005260206000200154601983815481106111f657fe5b6000918252602090912001556019805490611215906000198301615d33565b5050506001016110d3565b50611234600a54600954600a5402836148bb565b61123e8184614968565b61124d81600954600101614a6f565b50505050565b601d5481565b600f5481565b60085481565b600b5481565b60105460408051630e9ed68b60e01b815290516000926001600160a01b031691630e9ed68b916004808301926020929190829003018186803b1580156112b057600080fd5b505afa1580156112c4573d6000803e3d6000fd5b505050506040513d60208110156112da57600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b15801561132457600080fd5b505afa158015611338573d6000803e3d6000fd5b505050506040513d602081101561134e57600080fd5b5051611398576040805162461bcd60e51b8152602060048201526014602482015273696e76616c69642d7374616b696e672d6e6f646560601b604482015290519081900360640190fd5b6113a133614d93565b90505b90565b6000546001600160a01b031633146113be57600080fd5b6011546001600160a01b038381169116148015906113e457506001600160a01b03821615155b611429576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b6012546001600160a01b0382811691161480159061144f57506001600160a01b03811615155b611494576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b601180546001600160a01b038085166001600160a01b031992831617909255601280548484169216919091179055601054604080516313a4cbcb60e31b815290519190921691639d265e58916004808301926020929190829003018186803b1580156114ff57600080fd5b505afa158015611513573d6000803e3d6000fd5b505050506040513d602081101561152957600080fd5b50516011546012546040805163b73a3f8f60e01b81526001600160a01b039384166004820152918316602483015251919092169163b73a3f8f91604480830192600092919082900301818387803b158015610d1957600080fd5b6000546001600160a01b0316331461159a57600080fd5b601b5481141580156115ab57508015155b6115f0576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b601b55565b6000546001600160a01b0316331461160c57600080fd5b600654811415801561161d57508015155b611662576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b600655565b60016000819052601c602052600080516020615f9a833981519152549081148015906116a85750601b546000828152601a6020526040902060010154439101105b156117c0576116b88160016150fc565b604080514381523360208201528151600080516020615fba833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b15801561172e57600080fd5b505afa158015611742573d6000803e3d6000fd5b505050506040513d602081101561175857600080fd5b5051604080516323ff34cb60e01b815233600482015290516001600160a01b03909216916323ff34cb9160248082019260009290919082900301818387803b1580156117a357600080fd5b505af11580156117b7573d6000803e3d6000fd5b50505050610d33565b6040805160208082526018908201527f6e6f2d657870697265642d706772702d746f2d636c65616e00000000000000008183015290516000805160206160238339815191529181900360600190a150565b80600e541461185a576040805162461bcd60e51b815260206004820152601060248201526f06e6f742d696e2d626f6f7473747261760841b604482015290519081900360640190fd5b600f5443116118b657604080516020808252601c908201527f776169742d746f2d636f6c6c6563742d6d6f72652d656e74726f7079000000008183015290516000805160206160238339815191529181900360600190a1610d33565b600d54601554101561191557604080516020808252601e908201527f6e6f742d656e6f7567682d702d6e6f64652d746f2d626f6f74737472617000008183015290516000805160206160238339815191529181900360600190a1610d33565b6000600e819055600f819055601054604080516306b810cf60e21b815290516001600160a01b0390921691631ae0433c91600480820192602092909190829003018186803b15801561196657600080fd5b505afa15801561197a573d6000803e3d6000fd5b505050506040513d602081101561199057600080fd5b505160408051633352da4560e21b81526004810185905290516001600160a01b039092169163cd4b6914916024808201926020929091908290030181600087803b1580156119dd57600080fd5b505af11580156119f1573d6000803e3d6000fd5b505050506040513d6020811015611a0757600080fd5b5051905080611a6457604080516020808252601f908201527f626f6f7473747261702d636f6d6d69742d72657665616c2d6661696c757265008183015290516000805160206160238339815191529181900360600190a150610d33565b6000601f8190554360209081556021805460408051808501929092528181018690528051808303820181526060909201905280519201919091209055600a54600d54819081611aaf57fe5b04029050606081604051908082528060200260200182016040528015611adf578160200160208202803883390190505b509050611aee826000836148bb565b611afa81602154614968565b611b0f81600a548481611b0957fe5b04614a6f565b604080514381523360208201528151600080516020615fba833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015611b8557600080fd5b505afa158015611b99573d6000803e3d6000fd5b505050506040513d6020811015611baf57600080fd5b5051604080516323ff34cb60e01b815233600482015290516001600160a01b03909216916323ff34cb9160248082019260009290919082900301818387803b158015611bfa57600080fd5b505af1158015611c0e573d6000803e3d6000fd5b5050505050505050565b60188181548110611c2557fe5b600091825260209091200154905081565b600a5481565b60035481565b601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c9057600080fd5b505afa158015611ca4573d6000803e3d6000fd5b505050506040513d6020811015611cba57600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b158015611d0457600080fd5b505afa158015611d18573d6000803e3d6000fd5b505050506040513d6020811015611d2e57600080fd5b5051611d78576040805162461bcd60e51b8152602060048201526014602482015273696e76616c69642d7374616b696e672d6e6f646560601b604482015290519081900360640190fd5b6000858152600260205260409020600601546001600160a01b031680611dc7576040517f40d87958cd48e8b698a94f35390a9020a9127528227647da089cc6bfd7931f4290600090a150612340565b611ea6858786868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040805180820190915288358152915050602081018760016020908102919091013590915260008c8152600291829052604090819020815160808101808452828501805494830194855291949193859390928592909160030160608601808311610b6c57505050918352505060408051808201918290526002848101805483526020948501949293909260038701908501808311610ba2575050505050815250506140bf565b611eb05750612340565b604080516001600160a01b038316815290517f065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf09181900360200190a160008681526002602081905260408220828155600181018390559190820181611f158282615d57565b611f23600283016000615d57565b50505060060180546001600160a01b031916905560ff851660021415611fda5760408051636d11297760e01b81526004810188815260248201928352604482018690526001600160a01b03841692636d112977928a9289928992606401848480828437600081840152601f19601f820116905080830192505050945050505050600060405180830381600087803b158015611fbd57600080fd5b505af1158015611fd1573d6000803e3d6000fd5b505050506120e6565b60ff85166001141561209957604080516020808252600a90820152695573657252616e646f6d60b01b8183015290516000805160206160238339815191529181900360600190a1604080518335602082810191909152808501358284015282518083038401815260608301808552815191909201206318a1908d60e01b90915260648201899052608482015290516001600160a01b038316916318a1908d9160a480830192600092919082900301818387803b158015611fbd57600080fd5b6040805162461bcd60e51b815260206004820152601860248201527f756e737570706f727465642d747261666669632d747970650000000000000000604482015290519081900360640190fd5b6120ee615d65565b600087815260026020818152604080842060019081015485526017835293819020815160a081018352815481528251608081018085529196929594870194909392860192849290830191849182845b81548152602001906001019080831161213d57505050918352505060408051808201918290526020909201919060028481019182845b81548152602001906001019080831161217357505050505081525050815260200160058201548152602001600682015481526020016007820180548060200260200160405190810160405280929190818152602001828054801561220057602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116121e2575b5050505050815250509050601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b15801561225957600080fd5b505afa15801561226d573d6000803e3d6000fd5b505050506040513d602081101561228357600080fd5b50516080820151604051632100fdf760e21b8152600481018a815233602483018190526060604484019081528451606485015284516001600160a01b0390961695638403f7dc958e95939490939092916084909101906020858101910280838360005b838110156122fe5781810151838201526020016122e6565b50505050905001945050505050600060405180830381600087803b15801561232557600080fd5b505af1158015612339573d6000803e3d6000fd5b5050505050505b5050505050565b6010546001600160a01b031681565b60225481565b336000908152602b602052604090205460ff166123af576040805162461bcd60e51b815260206004820152600c60248201526b3737ba16b3bab0b93234b0b760a11b604482015290519081900360640190fd5b6123b881614d93565b1561243357604080514381523360208201528151600080516020615fba833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b15801561172e57600080fd5b604080516020808252601590820152743737ba3434b73396ba3796bab73932b3b4b9ba32b960591b8183015290516000805160206160238339815191529181900360600190a150565b60055481565b60075481565b600e5481565b601e5481565b6000546001600160a01b031633146124ab57600080fd5b60055481141580156124bc57508015155b612501576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b600555565b60015481565b60045481565b6012546001600160a01b031681565b6000546001600160a01b0316331461253857600080fd5b600b54811415801561254957508015155b61258e576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b600b55565b601c6020526000908152604090205481565b6013602052600090815260409020546001600160a01b031681565b601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561260e57600080fd5b505afa158015612622573d6000803e3d6000fd5b505050506040513d602081101561263857600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b15801561268257600080fd5b505afa158015612696573d6000803e3d6000fd5b505050506040513d60208110156126ac57600080fd5b50516126f6576040805162461bcd60e51b8152602060048201526014602482015273696e76616c69642d7374616b696e672d6e6f646560601b604482015290519081900360640190fd5b336000908152601360205260409020546001600160a01b03161561271957610f8c565b336000908152601660209081526040808320600184529091529020541561273f57610f8c565b336000818152601660209081526040808320600180855292529091205561276590615259565b6040805133815290517f707a6d64786780aac9cd0c5813ea04241eb135ddd2280c06eea6719afb9131519181900360200190a1601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156127e657600080fd5b505afa1580156127fa573d6000803e3d6000fd5b505050506040513d602081101561281057600080fd5b505160408051634c542d3d60e01b815233600482015290516001600160a01b0390921691634c542d3d9160248082019260009290919082900301818387803b15801561285b57600080fd5b505af115801561286f573d6000803e3d6000fd5b50505050610d336143e6565b60198181548110611c2557fe5b60185490565b6000866002601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b1580156128e157600080fd5b505afa1580156128f5573d6000803e3d6000fd5b505050506040513d602081101561290b57600080fd5b5051604080516371b281b360e11b81526001600160a01b038581166004830152602482018590529151919092169163e3650366916044808301926020929190829003018186803b15801561295e57600080fd5b505afa158015612972573d6000803e3d6000fd5b505050506040513d602081101561298857600080fd5b50516129d6576040805162461bcd60e51b81526020600482015260186024820152776e6f742d656e6f7567682d6665652d746f2d6f7261636c6560401b604482015290519081900360640190fd5b60006129e18a6152bb565b1115612f1357606085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250508251929350509015905080612a51575080600081518110612a3a57fe5b6020910101516001600160f81b031916600960fa1b145b80612a7b575080600081518110612a6457fe5b6020910101516001600160f81b031916602f60f81b145b15612ea9576000428b8b8b8b8b8b60405160200180888152602001876001600160a01b03166001600160a01b0316815260200186815260200180602001806020018381038352878782818152602001925080828437600083820152601f01601f191690910184810383528581526020019050858580828437600081840152601f19601f82011690508083019250505099505050505050505050506040516020818303038152906040528051906020012060001c90506000612b3d6002836152bf565b9050600019811415612ba357604080516020808252601f908201527f736b69707065642d757365722d71756572792d6e6f2d6c6976652d77677270008183015290516000805160206160238339815191529181900360600190a160009550505050612f54565b60006017600060188481548110612bb657fe5b6000918252602080832090910154835282810193909352604091820190208151608080820184528782528254948201949094528251938401835290935091828201916001850190829081018260028282826020028201915b815481526020019060010190808311612c0e57505050918352505060408051808201918290526020909201919060028481019182845b815481526020019060010190808311612c44575050509190925250505081526001600160a01b038f166020918201526000858152600280835260409182902084518155928401516001840155908301518051909183810191612ca891839190615d9a565b506020820151612cbe9060028084019190615d9a565b50505060608201518160060160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050507f05e1614af4efb13caeba2369a57a05ee5830f33364f82e2c899fd5710cb56ef3838d8d8d8d8d60215488600001546040518089815260200188815260200180602001806020018581526020018481526020018381038352898982818152602001925080828437600083820152601f01601f191690910184810383528781526020019050878780828437600083820152604051601f909101601f19169092018290039c50909a5050505050505050505050a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015612df157600080fd5b505afa158015612e05573d6000803e3d6000fd5b505050506040513d6020811015612e1b57600080fd5b50516001600160a01b0316637aa9181b8e8560026040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b031681526020018381526020018281526020019350505050600060405180830381600087803b158015612e8557600080fd5b505af1158015612e99573d6000803e3d6000fd5b5050505082965050505050612f54565b7f70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41868660405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a16000935050612f54565b604080516001600160a01b038b16815290517f6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb59181900360200190a1600092505b50509695505050505050565b601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b158015612fae57600080fd5b505afa158015612fc2573d6000803e3d6000fd5b505050506040513d6020811015612fd857600080fd5b50516040805163151d156760e31b815233600482015290516001600160a01b039092169163a8e8ab3891602480820192602092909190829003018186803b15801561302257600080fd5b505afa158015613036573d6000803e3d6000fd5b505050506040513d602081101561304c57600080fd5b5051613096576040805162461bcd60e51b8152602060048201526014602482015273696e76616c69642d7374616b696e672d6e6f646560601b604482015290519081900360640190fd5b6000828152601a6020526040902080546130e3576040805184815290517f71047c0893a51085656a2894bba10bc6ef51a654f25e1ead1929b076487a95699181900360200190a1506134d0565b3360009081526003820160205260409020546001600160a01b031661314f576040805162461bcd60e51b815260206004820152601e60248201527f6e6f742d66726f6d2d617574686f72697a65642d6772702d6d656d6265720000604482015290519081900360640190fd5b6040805183356020808301919091528085013582840152848301356060808401919091528501356080808401919091528351808403909101815260a08301808552815191830191909120600081815260028701909352918490208054600101908190559087905260c083015291517f717e526bce26f8e67908004294b35133bbe2a9c7f611384cb0f484aca9223d019181900360e00190a16002600a54816131f357fe5b6000838152600285016020526040902054919004101561124d576060600a5460405190808252806020026020018201604052801561323b578160200160208202803883390190505b5060016000908152600385016020526040812054919250906001600160a01b03165b6001600160a01b0381166001146132d8578083838060010194508151811061328157fe5b6001600160a01b0392831660209182029290920181019190915290821660009081526016909152604090206132b690886152dd565b6001600160a01b0390811660009081526003860160205260409020541661325d565b6018805460018082019092557fb13d2d76d1f4b7be834882e410b3e3a8afaf69f83600ae24db354391d2378d2e018890556040805160a0810182528981528151608080820184528a358285019081526020808d0135606080860191909152918452855180870187528d8701358152828e01358183015281850152808501938452600654601e540285870152439185019190915290830188905260008c81526017909152929092208151815591518051919390919083019061339c9082906002615d9a565b5060208201516133b29060028084019190615d9a565b5050506040820151600582015560608201516006820155608082015180516133e4916007840191602090910190615dd8565b509050506000806133f6601c8a615301565b91509150808015613408575088601d54145b1561341357601d8290555b6000898152601a6020908152604080832083815560010192909255601e805460001901905581518b815291517f156927b06a61046135669011768a03b2592ee3667374525502b62b0aef4cbd219281900390910190a16018546040518a81527f9ad0d0bfff7f0fc653b03785685d46101b09e3cb1f50081915bc8101662e4a88918b918b91906020810183608080828437600083820152601f01601f191690910192835250506040519081900360200192509050a1505050505050505b5050565b6000546001600160a01b031633146134eb57600080fd5b600c5481141580156134fc57508015155b613541576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b600c55565b61354e61538a565b156135c957604080514381523360208201528151600080516020615fba833981519152929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015610eb857600080fd5b604080516020808252601690820152751cde5ccb5c985b991bdb4b5b9bdd0b595e1c1a5c995960521b8183015290516000805160206160238339815191529181900360600190a1565b601f5481565b6000546001600160a01b0316331461362f57600080fd5b600454811415801561364057508015155b613685576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b600455565b6000546001600160a01b031633146136a157600080fd5b6001600160a01b03166000908152602b60205260409020805460ff19169055565b6000826001601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b15801561371557600080fd5b505afa158015613729573d6000803e3d6000fd5b505050506040513d602081101561373f57600080fd5b5051604080516371b281b360e11b81526001600160a01b038581166004830152602482018590529151919092169163e3650366916044808301926020929190829003018186803b15801561379257600080fd5b505afa1580156137a6573d6000803e3d6000fd5b505050506040513d60208110156137bc57600080fd5b505161380a576040805162461bcd60e51b81526020600482015260186024820152776e6f742d656e6f7567682d6665652d746f2d6f7261636c6560401b604482015290519081900360640190fd5b60408051426020808301919091526001600160a01b038816828401526060808301889052835180840390910181526080909201909252805191012060006138526001836152bf565b90506000198114156138b757604080516020808252601d908201527f736b69707065642d757365722d726e642d6e6f2d6c6976652d776772700000008183015290516000805160206160238339815191529181900360600190a1600094505050613b67565b600060176000601884815481106138ca57fe5b6000918252602080832090910154835282810193909352604091820190208151608080820184528782528254948201949094528251938401835290935091828201916001850190829081018260028282826020028201915b81548152602001906001019080831161392257505050918352505060408051808201918290526020909201919060028481019182845b815481526020019060010190808311613958575050509190925250505081526001600160a01b038a1660209182015260008581526002808352604091829020845181559284015160018401559083015180519091838101916139bc91839190615d9a565b5060208201516139d29060028084019190615d9a565b505050606091820151600690910180546001600160a01b039092166001600160a01b031990921691909117905560215482546040805187815260208101939093528281018b90529282015290517fd587179d80d44e74955fa5d651db2f31b5470fcee8f9736f07ae3b24456a2cf09181900360800190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b158015613a9757600080fd5b505afa158015613aab573d6000803e3d6000fd5b505050506040513d6020811015613ac157600080fd5b50516001600160a01b0390811690637aa9181b908a163014613ae35789613af0565b6011546001600160a01b03165b8560016040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b031681526020018381526020018281526020019350505050600060405180830381600087803b158015613b4957600080fd5b505af1158015613b5d573d6000803e3d6000fd5b5094975050505050505b505092915050565b6000546001600160a01b03163314613b8657600080fd5b600354811415613bd1576040805162461bcd60e51b815260206004820152601160248201527034b73b30b634b216b830b930b6b2ba32b960791b604482015290519081900360640190fd5b600355565b601a602052600090815260409020805460019091015482565b6000546001600160a01b03163314613c0657600080fd5b6001600160a01b03166000908152602b60205260409020805460ff19166001179055565b60065481565b6011546001600160a01b031681565b6000546001600160a01b03163314613c5657600080fd5b601954601854601e54600a54601554919092019092010201808210613c7b5780613c7d565b815b60016000526013602052600080516020616003833981519152549092506001600160a01b0316825b600084118015613cd85750600160008190526013602052600080516020616003833981519152546001600160a01b031614155b15613d1e576001600160a01b0391821660008181526016602090815260408083206001845282528083208390559282526013905220546000199094019390911690613ca5565b6015548482031015613d3257505050610d33565b506001600081815260008051602061600383398151915280546001600160a01b0319908116841790915560148054909116909217909155601555601c602052600080516020615f9a8339815191525483905b600085118015613d95575060018114155b15613e39576000818152601a602090815260408083206001845260038101909252909120546001600160a01b03165b600087118015613dde57506001600160a01b038116600114155b15613e23576001600160a01b03908116600081815260166020908152604080832060018452825280832083905592825260038501905220546000199097019616613dc4565b50506000908152601c6020526040902054613d84565b601e54600a54028583031015613e525750505050610d33565b60016000818152601c602052600080516020615f9a833981519152829055601d91909155601e8190558592505b600086118015613e90575060185481105b15613f3d5760006017600060188481548110613ea857fe5b90600052602060002001548152602001908152602001600020600701905060008090505b600088118015613edc5750815481105b15613f3357600060166000848481548110613ef357fe5b60009182526020808320909101546001600160a01b0316835282810193909352604091820181206001808352935220919091556000199098019701613ecc565b5050600101613e7f565b50601854600a54028583031015613f575750505050610d33565b8491506000613f67601882615d33565b5060005b600086118015613f7c575060195481105b156140295760006017600060198481548110613f9457fe5b90600052602060002001548152602001908152602001600020600701905060008090505b600088118015613fc85750815481105b1561401f57600060166000848481548110613fdf57fe5b60009182526020808320909101546001600160a01b0316835282810193909352604091820181206001808352935220919091556000199098019701613fb8565b5050600101613f6b565b50601954600a540285830310156140435750505050610d33565b6000614050601982615d33565b50506000600e819055601f8190556020819055602181905560225550505050565b600c5481565b60195490565b60215481565b60155481565b60205481565b601b5481565b60408051602080825281830190925260609160208201818038833950505060208101929092525090565b6000606084336040516020018083805190602001908083835b602083106140f75780518252601f1990920191602091820191016140d8565b5181516020939093036101000a6000190180199091169216919091179052606094851b6bffffffffffffffffffffffff191692019182525060408051808303600b19018152600260148401818152607485019093529096509394509291506034015b614161615e39565b81526020019060019003908161415957505060408051600280825260608083019093529293509091602082015b614196615e53565b81526020019060019003908161418e5790505090506141b4866155ed565b826000815181106141c157fe5b60200260200101819052506141d583615666565b826001815181106141e257fe5b60200260200101819052506141f561568d565b8160008151811061420257fe5b6020026020010181905250848160018151811061421b57fe5b60200260200101819052506000614232838361574d565b90507fd33c44f7ce166bcd2616c8f4d811261d4d24e1af815f78614683e0c7682c42a58a8a8660405180604001604052808c6000015181526020018c6020015181525060405180608001604052808c6000015160006002811061429157fe5b602002015181526020018c600001516001600281106142ac57fe5b602002015181526020018c602001516000600281106142c757fe5b602002015181526020018c602001516001600281106142e257fe5b602002015181525086604051808760ff1660ff1681526020018681526020018060200185600260200280838360005b83811015614329578181015183820152602001614311565b5050505090500184600460200280838360005b8381101561435457818101518382015260200161433c565b5050505090500183151515158152602001828103825286818151815260200191508051906020019080838360005b8381101561439a578181015183820152602001614382565b50505050905090810190601f1680156143c75780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a19998505050505050505050565b6000600a54601554108061440857506018541580156144085750600d54601554105b15614485576019541561448557614439601960008154811061442657fe5b906000526020600020015460018061475b565b60198054600019810190811061444b57fe5b9060005260206000200154601960008154811061446457fe5b6000918252602090912001556019805490614483906000198301615d33565b505b600a5460155410156144ce5760155460408051918252517fc03848aa1689c7c291dcc68fa62d6109e13f16b25e89bdef5a6c8598b36b80199181900360200190a15060006113a4565b601854156145d9576009546019541061459b5760225461454e576144f230436136c2565b6022819055614503575060006113a4565b7f5ee79c79dd5be1f25f86c7028ec8fb6849839e8a23f676afc907eade45946986601554600a54604051808381526020018281526020019250505060405180910390a15060016113a4565b6040805160208082526014908201527330b63932b0b23c96b4b716b337b936b0ba34b7b760611b8183015290516000805160206160238339815191529181900360600190a15060006113a4565b600080516020616023833981519152604051808060200182810382526029815260200180615fda6029913960400191505060405180910390a1614755565b600d546015541061475557600e5461470f57601060009054906101000a90046001600160a01b03166001600160a01b0316631ae0433c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561463957600080fd5b505afa15801561464d573d6000803e3d6000fd5b505050506040513d602081101561466357600080fd5b5051600b54600c54600d546040805163b917b5a560e01b8152436004820152602481019490945260448401929092526064830152516001600160a01b039092169163b917b5a5916084808201926020929091908290030181600087803b1580156146cc57600080fd5b505af11580156146e0573d6000803e3d6000fd5b505050506040513d60208110156146f657600080fd5b5051600e5550600c54600b54430101600f5560016113a4565b604080516020808252601490820152730616c72656164792d696e2d626f6f7473747261760641b8183015290516000805160206160238339815191529181900360600190a15b50600090565b6000838152601760205260408120905b600782015481101561482e57600082600701828154811061478857fe5b60009182526020808320909101546001600160a01b0316808352601690915260408220855491935082916147bc9190615301565b915091508080156147cd5750600182145b15614823578680156147f15750856001600160a01b0316836001600160a01b031614155b801561481557506001600160a01b0383811660009081526013602052604090205416155b156148235761482383615259565b50505060010161476b565b5060008481526017602052604081208181559060018201816148508282615d57565b61485e600283016000615d57565b5050600582016000905560068201600090556007820160006148809190615e78565b50506040805185815290517ff7377b41bdc770cc22a1bad318571f0ae6d65188371bdbe0cb660c0de57b31509181900360200190a150505050565b60005b8381101561493f57601360205260008051602061600383398151915280546001600160a01b0380821660008181526040902080549092166001600160a01b0319938416179093558054909116905582518190849084870190811061491e57fe5b6001600160a01b0390921660209283029190910190910152506001016148be565b50601580548490039081905561496357601480546001600160a01b03191660011790555b505050565b8151600019015b801561496357600081600101838386858151811061498957fe5b602002602001015160405160200180848152602001838152602001826001600160a01b03166001600160a01b031660601b815260140193505050506040516020818303038152906040528051906020012060001c816149e457fe5b06905060008483815181106149f557fe5b60200260200101519050848281518110614a0b57fe5b6020026020010151858481518110614a1f57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505080858381518110614a4c57fe5b6001600160a01b039092166020928302919091019091015250506000190161496f565b80600a5402825114614ac8576040805162461bcd60e51b815260206004820152601960248201527f63616e6469646174652d6c656e6774682d6d69736d6174636800000000000000604482015290519081900360640190fd5b6060600a54604051908082528060200260200182016040528015614af6578160200160208202803883390190505b5090506000805b838110156123405760009150815b600a54811015614bb8578581600a5484020181518110614b2757fe5b6020026020010151848281518110614b3b57fe5b60200260200101906001600160a01b031690816001600160a01b03168152505082848281518110614b6857fe5b602090810291909101810151604080518084019490945260609190911b6bffffffffffffffffffffffff191683820152805180840360340181526054909301905281519101209250600101614b0b565b506040805160208082019490945242818301528151808203830181526060820180845281519186019190912060a0830184528082524360809093019283526000818152601a875284812092518355925160018084019190915580845260039092019586905292822080546001600160a01b03191690911790559092905b600a54811015614d0357600160009081526020839052604081205486516001600160a01b03909116918491889085908110614c6c57fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b03160217905550848181518110614cc457fe5b6020908102919091018101516001600081815292859052604090922080546001600160a01b0319166001600160a01b0390921691909117905501614c35565b50614d0d83615950565b7f78bf54a42d1b98e6c809c3e5904898c5b3304ede546b6f070e83d9a32e544d4f83856040518083815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015614d76578181015183820152602001614d5e565b50505050905001935050505060405180910390a150600101614afd565b6001600160a01b03811660009081526016602090815260408083206001845290915281205481808215801590614dca575060018314155b15614f1757614ddb8360018761475b565b60005b601854811015614e74578360188281548110614df657fe5b90600052602060002001541415614e6c57601854600019018114614e4b57601880546000198101908110614e2657fe5b906000526020600020015460188281548110614e3e57fe5b6000918252602090912001555b6018805490614e5e906000198301615d33565b506001925090821790614e74565b600101614dde565b5081614f175760005b601954811015614f15578360198281548110614e9557fe5b90600052602060002001541415614f0d57601954600019018114614eea57601980546000198101908110614ec557fe5b906000526020600020015460198281548110614edd57fe5b6000918252602090912001555b6019805490614efd906000198301615d33565b5060019250600282179150614f15565b600101614e7d565b505b81158015614f295750614f2985615980565b15614f32576004175b6001600160a01b038581166000908152601360205260409020541615614fca576000614f5f601387615a07565b935090508215614fc857601580546000190190556001600160a01b038087166000818152601660209081526040808320600184529091528120556014549091161415614fc157601480546001600160a01b0319166001600160a01b0383161790555b6008821791505b505b604080516001600160a01b038716815260ff8316602082015281517faa40dce54d678a8a16fc3cf106c1ddf0b34b66a43c7a365af3268c63bb95fead929181900390910190a1601060009054906101000a90046001600160a01b03166001600160a01b0316630e9ed68b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561505e57600080fd5b505afa158015615072573d6000803e3d6000fd5b505050506040513d602081101561508857600080fd5b50516040805163c5375c2960e01b81526001600160a01b0388811660048301529151919092169163c5375c2991602480830192600092919082900301818387803b1580156150d557600080fd5b505af11580156150e9573d6000803e3d6000fd5b5050505060ff161515925050505b919050565b6000828152601a602090815260408083206001845260038101909252909120546001600160a01b03165b6001600160a01b0381166001146151d0576001600160a01b0381166000908152601660209081526040808320600180855292529091205414801561517c5750826001600160a01b0316816001600160a01b031614155b80156151a057506001600160a01b0381811660009081526013602052604090205416155b156151ae576151ae81615a6f565b6001600160a01b03908116600090815260038301602052604090205416615126565b6000806151de601c87615301565b915091508080156151f0575085601d54145b156151fb57601d8290555b6000868152601a6020908152604080832083815560010192909255601e8054600019019055815188815291517f156927b06a61046135669011768a03b2592ee3667374525502b62b0aef4cbd219281900390910190a1505050505050565b601480546001600160a01b039081166000908152601360205260408082205494831680835281832080549685166001600160a01b0319978816179055845490931682529020805484168217905581549092169091179055601580546001019055565b3b90565b60006152c961538a565b506152d48383615aca565b90505b92915050565b60016000818152602093909352604080842080548486529185209190915592529055565b6001600081815260208490526040812054909182915b600181141580156153285750848114155b1561534457600081815260208790526040902054909150615317565b600181141561535c5760016000935093505050615383565b60008181526020879052604080822080548584529183209190915591815290559150600190505b9250929050565b6000436004546020540111806153a5575043600354601f5401115b156153b2575060006113a4565b60006153c381600019430140615aca565b905060001981141561542757604080516020808252601a908201527f6e6f2d6c6976652d776772702c7472792d626f6f7473747261700000000000008183015290516000805160206160238339815191529181900360600190a160009150506113a4565b43601f81905550601760006018838154811061543f57fe5b6000918252602080832090910154835282019290925260400190208054602390815560018201602461547381836002615e96565b5061548660028281019084810190615e96565b505050600582015481600501556006820154816006015560078201816007019080546154b3929190615ec1565b505060215460235460408051928352602083019190915280517ffaa99731d2c5abc7ee76b2e31b6b7e293a30e1e2274f59396a7a59cabd5eadf39350918290030190a1601060009054906101000a90046001600160a01b03166001600160a01b0316639d265e586040518163ffffffff1660e01b815260040160206040518083038186803b15801561554457600080fd5b505afa158015615558573d6000803e3d6000fd5b505050506040513d602081101561556e57600080fd5b505160115460215460408051637aa9181b60e01b81526001600160a01b039384166004820152602481019290925260006044830181905290519290931692637aa9181b9260648084019382900301818387803b1580156155cd57600080fd5b505af11580156155e1573d6000803e3d6000fd5b50505050600191505090565b6155f5615e39565b815115801561560657506020820151155b156156125750806150f7565b60007f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4790506040518060400160405280846000015181526020018285602001518161565957fe5b0690920390915292915050565b61566e615e39565b81516020830120615686615680615c2e565b82615c4f565b9392505050565b615695615e53565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b82527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa60208381019190915281019190915290565b6000815183511461575d57600080fd5b8251604080516006830280825260c084028201602001909252606090828015615790578160200160208202803883390190505b50905060005b83811015615915578681815181106157aa57fe5b6020026020010151600001518282600602600001815181106157c857fe5b6020026020010181815250508681815181106157e057fe5b6020026020010151602001518282600602600101815181106157fe57fe5b60200260200101818152505085818151811061581657fe5b60209081029190910101515151825183906002600685020190811061583757fe5b60200260200101818152505085818151811061584f57fe5b6020908102919091010151516001602002015182826006026003018151811061587457fe5b60200260200101818152505085818151811061588c57fe5b6020026020010151602001516000600281106158a457fe5b60200201518282600602600401815181106158bb57fe5b6020026020010181815250508581815181106158d357fe5b6020026020010151602001516001600281106158eb57fe5b602002015182826006026005018151811061590257fe5b6020908102919091010152600101615796565b5061591e615f01565b60006020826020860260208601600060086107d05a03f190508080156159445750815115155b98975050505050505050565b601d80546000908152601c6020526040808220548483528183205582548252902082905555601e80546001019055565b60016000818152601c602052600080516020615f9a833981519152549091905b600181146159fd576000818152601a60205260408120906159c46003830187615c93565b91505080156159e4576159d783876150fc565b60019450505050506150f7565b50506000818152601c60205260409020549091506159a0565b5060009392505050565b600080600080615a178686615c93565b915091508015615a64576001600160a01b03858116600081815260208990526040808220805487861684529183208054929095166001600160a01b03199283161790945591905281541690555b909590945092505050565b601360205260008051602061600383398151915280546001600160a01b039283166000818152604081208054959093166001600160a01b03199586161790925560019182905282549093169092179055601580549091019055565b6000805b601854615ae0576000199150506152d7565b60185481101580615af357506007548110155b15615b6957600084846021544360405160200180856002811115615b1357fe5b60ff1660f81b81526001018481526020018381526020018281526020019450505050506040516020818303038152906040528051906020012060001c90506018805490508181615b5f57fe5b06925050506152d7565b60006017600060188481548110615b7c57fe5b9060005260206000200154815260200190815260200160002090504381600501548260060154600554010111615c2557601960188381548110615bbb57fe5b60009182526020808320909101548354600181018555938352912090910155601880546000198101908110615bec57fe5b906000526020600020015460188381548110615c0457fe5b6000918252602090912001556018805490615c23906000198301615d33565b505b50600101615ace565b615c36615e39565b5060408051808201909152600181526002602082015290565b615c57615e39565b615c5f615f1f565b8351815260208085015190820152604080820184905282606083600060076107d05a03f1615c8c57600080fd5b5092915050565b6001600081815260208490526040812054909182916001600160a01b03165b6001600160a01b038116600114801590615cde5750846001600160a01b0316816001600160a01b031614155b15615d06576001600160a01b0380821660009081526020889052604090205491925016615cb2565b6001600160a01b03811660011415615d275760016000935093505050615383565b50915060019050615383565b81548183558181111561496357600083815260209020614963918101908301615f3d565b506000815560010160009055565b6040518060a0016040528060008152602001615d7f615e53565b81526020016000815260200160008152602001606081525090565b8260028101928215615dc8579160200282015b82811115615dc8578251825591602001919060010190615dad565b50615dd4929150615f3d565b5090565b828054828255906000526020600020908101928215615e2d579160200282015b82811115615e2d57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190615df8565b50615dd4929150615f57565b604051806040016040528060008152602001600081525090565b6040518060400160405280615e66615f7b565b8152602001615e73615f7b565b905290565b5080546000825590600052602060002090810190610d339190615f3d565b8260028101928215615dc8579182015b82811115615dc8578254825591600101919060010190615ea6565b828054828255906000526020600020908101928215615e2d5760005260206000209182015b82811115615e2d578254825591600101919060010190615ee6565b60405180602001604052806001906020820280388339509192915050565b60405180606001604052806003906020820280388339509192915050565b6113a491905b80821115615dd45760008155600101615f43565b6113a491905b80821115615dd45780546001600160a01b0319168155600101615f5d565b6040518060400160405280600290602082028038833950919291505056fe6de76108811faf2f94afbe5ac6c98e8393206cd093932de1fbfd61bbeec43a02a60d55093b21f740878d9871e95e5031eaf5cf08a167c898ed3c62b1fb24f887736b69707065642d666f726d6174696f6e2d6e6f742d656e6f7567682d657870697265642d776772704155c2f711f2cdd34f8262ab8fb9b7020a700fe7b6948222152f7670d1fdf34d96561394bac381230de4649200e8831afcab1f451881bbade9ef209f6dd30480a265627a7a72315820baaaa967d2a0102993081418e6e0c7a35be2355b5dd870e0578690f36caa6d0064736f6c63430005110032"

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
// Solidity: function addressBridge() view returns(address)
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
// Solidity: function addressBridge() view returns(address)
func (_Dosproxy *DosproxySession) AddressBridge() (common.Address, error) {
	return _Dosproxy.Contract.AddressBridge(&_Dosproxy.CallOpts)
}

// AddressBridge is a free data retrieval call binding the contract method 0x76cffa53.
//
// Solidity: function addressBridge() view returns(address)
func (_Dosproxy *DosproxyCallerSession) AddressBridge() (common.Address, error) {
	return _Dosproxy.Contract.AddressBridge(&_Dosproxy.CallOpts)
}

// BootstrapCommitDuration is a free data retrieval call binding the contract method 0x372a53cc.
//
// Solidity: function bootstrapCommitDuration() view returns(uint256)
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
// Solidity: function bootstrapCommitDuration() view returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapCommitDuration() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapCommitDuration(&_Dosproxy.CallOpts)
}

// BootstrapCommitDuration is a free data retrieval call binding the contract method 0x372a53cc.
//
// Solidity: function bootstrapCommitDuration() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapCommitDuration() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapCommitDuration(&_Dosproxy.CallOpts)
}

// BootstrapEndBlk is a free data retrieval call binding the contract method 0x19717203.
//
// Solidity: function bootstrapEndBlk() view returns(uint256)
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
// Solidity: function bootstrapEndBlk() view returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapEndBlk() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapEndBlk(&_Dosproxy.CallOpts)
}

// BootstrapEndBlk is a free data retrieval call binding the contract method 0x19717203.
//
// Solidity: function bootstrapEndBlk() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapEndBlk() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapEndBlk(&_Dosproxy.CallOpts)
}

// BootstrapGroups is a free data retrieval call binding the contract method 0x31bf6464.
//
// Solidity: function bootstrapGroups() view returns(uint256)
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
// Solidity: function bootstrapGroups() view returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapGroups() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapGroups(&_Dosproxy.CallOpts)
}

// BootstrapGroups is a free data retrieval call binding the contract method 0x31bf6464.
//
// Solidity: function bootstrapGroups() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapGroups() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapGroups(&_Dosproxy.CallOpts)
}

// BootstrapRevealDuration is a free data retrieval call binding the contract method 0xef112dfc.
//
// Solidity: function bootstrapRevealDuration() view returns(uint256)
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
// Solidity: function bootstrapRevealDuration() view returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapRevealDuration() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapRevealDuration(&_Dosproxy.CallOpts)
}

// BootstrapRevealDuration is a free data retrieval call binding the contract method 0xef112dfc.
//
// Solidity: function bootstrapRevealDuration() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapRevealDuration() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapRevealDuration(&_Dosproxy.CallOpts)
}

// BootstrapRound is a free data retrieval call binding the contract method 0x85ed4223.
//
// Solidity: function bootstrapRound() view returns(uint256)
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
// Solidity: function bootstrapRound() view returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapRound() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapRound(&_Dosproxy.CallOpts)
}

// BootstrapRound is a free data retrieval call binding the contract method 0x85ed4223.
//
// Solidity: function bootstrapRound() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapRound() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapRound(&_Dosproxy.CallOpts)
}

// BootstrapStartThreshold is a free data retrieval call binding the contract method 0x11bc5311.
//
// Solidity: function bootstrapStartThreshold() view returns(uint256)
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
// Solidity: function bootstrapStartThreshold() view returns(uint256)
func (_Dosproxy *DosproxySession) BootstrapStartThreshold() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapStartThreshold(&_Dosproxy.CallOpts)
}

// BootstrapStartThreshold is a free data retrieval call binding the contract method 0x11bc5311.
//
// Solidity: function bootstrapStartThreshold() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) BootstrapStartThreshold() (*big.Int, error) {
	return _Dosproxy.Contract.BootstrapStartThreshold(&_Dosproxy.CallOpts)
}

// CachedUpdatedBlock is a free data retrieval call binding the contract method 0xba419d0c.
//
// Solidity: function cachedUpdatedBlock() view returns(uint256)
func (_Dosproxy *DosproxyCaller) CachedUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "cachedUpdatedBlock")
	return *ret0, err
}

// CachedUpdatedBlock is a free data retrieval call binding the contract method 0xba419d0c.
//
// Solidity: function cachedUpdatedBlock() view returns(uint256)
func (_Dosproxy *DosproxySession) CachedUpdatedBlock() (*big.Int, error) {
	return _Dosproxy.Contract.CachedUpdatedBlock(&_Dosproxy.CallOpts)
}

// CachedUpdatedBlock is a free data retrieval call binding the contract method 0xba419d0c.
//
// Solidity: function cachedUpdatedBlock() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) CachedUpdatedBlock() (*big.Int, error) {
	return _Dosproxy.Contract.CachedUpdatedBlock(&_Dosproxy.CallOpts)
}

// ExpiredWorkingGroupIds is a free data retrieval call binding the contract method 0xb45ef79d.
//
// Solidity: function expiredWorkingGroupIds(uint256 ) view returns(uint256)
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
// Solidity: function expiredWorkingGroupIds(uint256 ) view returns(uint256)
func (_Dosproxy *DosproxySession) ExpiredWorkingGroupIds(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.ExpiredWorkingGroupIds(&_Dosproxy.CallOpts, arg0)
}

// ExpiredWorkingGroupIds is a free data retrieval call binding the contract method 0xb45ef79d.
//
// Solidity: function expiredWorkingGroupIds(uint256 ) view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) ExpiredWorkingGroupIds(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.ExpiredWorkingGroupIds(&_Dosproxy.CallOpts, arg0)
}

// GetExpiredWorkingGroupSize is a free data retrieval call binding the contract method 0xefde068b.
//
// Solidity: function getExpiredWorkingGroupSize() view returns(uint256)
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
// Solidity: function getExpiredWorkingGroupSize() view returns(uint256)
func (_Dosproxy *DosproxySession) GetExpiredWorkingGroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GetExpiredWorkingGroupSize(&_Dosproxy.CallOpts)
}

// GetExpiredWorkingGroupSize is a free data retrieval call binding the contract method 0xefde068b.
//
// Solidity: function getExpiredWorkingGroupSize() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GetExpiredWorkingGroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GetExpiredWorkingGroupSize(&_Dosproxy.CallOpts)
}

// GetWorkingGroupSize is a free data retrieval call binding the contract method 0xb5372264.
//
// Solidity: function getWorkingGroupSize() view returns(uint256)
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
// Solidity: function getWorkingGroupSize() view returns(uint256)
func (_Dosproxy *DosproxySession) GetWorkingGroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GetWorkingGroupSize(&_Dosproxy.CallOpts)
}

// GetWorkingGroupSize is a free data retrieval call binding the contract method 0xb5372264.
//
// Solidity: function getWorkingGroupSize() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GetWorkingGroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GetWorkingGroupSize(&_Dosproxy.CallOpts)
}

// GroupMaturityPeriod is a free data retrieval call binding the contract method 0x7c48d1a0.
//
// Solidity: function groupMaturityPeriod() view returns(uint256)
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
// Solidity: function groupMaturityPeriod() view returns(uint256)
func (_Dosproxy *DosproxySession) GroupMaturityPeriod() (*big.Int, error) {
	return _Dosproxy.Contract.GroupMaturityPeriod(&_Dosproxy.CallOpts)
}

// GroupMaturityPeriod is a free data retrieval call binding the contract method 0x7c48d1a0.
//
// Solidity: function groupMaturityPeriod() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GroupMaturityPeriod() (*big.Int, error) {
	return _Dosproxy.Contract.GroupMaturityPeriod(&_Dosproxy.CallOpts)
}

// GroupSize is a free data retrieval call binding the contract method 0x63b635ea.
//
// Solidity: function groupSize() view returns(uint256)
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
// Solidity: function groupSize() view returns(uint256)
func (_Dosproxy *DosproxySession) GroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GroupSize(&_Dosproxy.CallOpts)
}

// GroupSize is a free data retrieval call binding the contract method 0x63b635ea.
//
// Solidity: function groupSize() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GroupSize() (*big.Int, error) {
	return _Dosproxy.Contract.GroupSize(&_Dosproxy.CallOpts)
}

// GroupToPick is a free data retrieval call binding the contract method 0x0434ccd2.
//
// Solidity: function groupToPick() view returns(uint256)
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
// Solidity: function groupToPick() view returns(uint256)
func (_Dosproxy *DosproxySession) GroupToPick() (*big.Int, error) {
	return _Dosproxy.Contract.GroupToPick(&_Dosproxy.CallOpts)
}

// GroupToPick is a free data retrieval call binding the contract method 0x0434ccd2.
//
// Solidity: function groupToPick() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) GroupToPick() (*big.Int, error) {
	return _Dosproxy.Contract.GroupToPick(&_Dosproxy.CallOpts)
}

// GuardianListed is a free data retrieval call binding the contract method 0x09011cb9.
//
// Solidity: function guardianListed(address ) view returns(bool)
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
// Solidity: function guardianListed(address ) view returns(bool)
func (_Dosproxy *DosproxySession) GuardianListed(arg0 common.Address) (bool, error) {
	return _Dosproxy.Contract.GuardianListed(&_Dosproxy.CallOpts, arg0)
}

// GuardianListed is a free data retrieval call binding the contract method 0x09011cb9.
//
// Solidity: function guardianListed(address ) view returns(bool)
func (_Dosproxy *DosproxyCallerSession) GuardianListed(arg0 common.Address) (bool, error) {
	return _Dosproxy.Contract.GuardianListed(&_Dosproxy.CallOpts, arg0)
}

// InitBlkN is a free data retrieval call binding the contract method 0x95071cf6.
//
// Solidity: function initBlkN() view returns(uint256)
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
// Solidity: function initBlkN() view returns(uint256)
func (_Dosproxy *DosproxySession) InitBlkN() (*big.Int, error) {
	return _Dosproxy.Contract.InitBlkN(&_Dosproxy.CallOpts)
}

// InitBlkN is a free data retrieval call binding the contract method 0x95071cf6.
//
// Solidity: function initBlkN() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) InitBlkN() (*big.Int, error) {
	return _Dosproxy.Contract.InitBlkN(&_Dosproxy.CallOpts)
}

// LastFormGrpReqId is a free data retrieval call binding the contract method 0x77f10192.
//
// Solidity: function lastFormGrpReqId() view returns(uint256)
func (_Dosproxy *DosproxyCaller) LastFormGrpReqId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "lastFormGrpReqId")
	return *ret0, err
}

// LastFormGrpReqId is a free data retrieval call binding the contract method 0x77f10192.
//
// Solidity: function lastFormGrpReqId() view returns(uint256)
func (_Dosproxy *DosproxySession) LastFormGrpReqId() (*big.Int, error) {
	return _Dosproxy.Contract.LastFormGrpReqId(&_Dosproxy.CallOpts)
}

// LastFormGrpReqId is a free data retrieval call binding the contract method 0x77f10192.
//
// Solidity: function lastFormGrpReqId() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) LastFormGrpReqId() (*big.Int, error) {
	return _Dosproxy.Contract.LastFormGrpReqId(&_Dosproxy.CallOpts)
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() view returns(uint256)
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
// Solidity: function lastRandomness() view returns(uint256)
func (_Dosproxy *DosproxySession) LastRandomness() (*big.Int, error) {
	return _Dosproxy.Contract.LastRandomness(&_Dosproxy.CallOpts)
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) LastRandomness() (*big.Int, error) {
	return _Dosproxy.Contract.LastRandomness(&_Dosproxy.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
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
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_Dosproxy *DosproxySession) LastUpdatedBlock() (*big.Int, error) {
	return _Dosproxy.Contract.LastUpdatedBlock(&_Dosproxy.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) LastUpdatedBlock() (*big.Int, error) {
	return _Dosproxy.Contract.LastUpdatedBlock(&_Dosproxy.CallOpts)
}

// LifeDiversity is a free data retrieval call binding the contract method 0xdd6ceddf.
//
// Solidity: function lifeDiversity() view returns(uint256)
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
// Solidity: function lifeDiversity() view returns(uint256)
func (_Dosproxy *DosproxySession) LifeDiversity() (*big.Int, error) {
	return _Dosproxy.Contract.LifeDiversity(&_Dosproxy.CallOpts)
}

// LifeDiversity is a free data retrieval call binding the contract method 0xdd6ceddf.
//
// Solidity: function lifeDiversity() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) LifeDiversity() (*big.Int, error) {
	return _Dosproxy.Contract.LifeDiversity(&_Dosproxy.CallOpts)
}

// LoopLimit is a free data retrieval call binding the contract method 0x7e55acb4.
//
// Solidity: function loopLimit() view returns(uint256)
func (_Dosproxy *DosproxyCaller) LoopLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "loopLimit")
	return *ret0, err
}

// LoopLimit is a free data retrieval call binding the contract method 0x7e55acb4.
//
// Solidity: function loopLimit() view returns(uint256)
func (_Dosproxy *DosproxySession) LoopLimit() (*big.Int, error) {
	return _Dosproxy.Contract.LoopLimit(&_Dosproxy.CallOpts)
}

// LoopLimit is a free data retrieval call binding the contract method 0x7e55acb4.
//
// Solidity: function loopLimit() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) LoopLimit() (*big.Int, error) {
	return _Dosproxy.Contract.LoopLimit(&_Dosproxy.CallOpts)
}

// NodeToGroupIdList is a free data retrieval call binding the contract method 0x0eeee5c1.
//
// Solidity: function nodeToGroupIdList(address , uint256 ) view returns(uint256)
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
// Solidity: function nodeToGroupIdList(address , uint256 ) view returns(uint256)
func (_Dosproxy *DosproxySession) NodeToGroupIdList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.NodeToGroupIdList(&_Dosproxy.CallOpts, arg0, arg1)
}

// NodeToGroupIdList is a free data retrieval call binding the contract method 0x0eeee5c1.
//
// Solidity: function nodeToGroupIdList(address , uint256 ) view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) NodeToGroupIdList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.NodeToGroupIdList(&_Dosproxy.CallOpts, arg0, arg1)
}

// NumPendingGroups is a free data retrieval call binding the contract method 0x863bc0a1.
//
// Solidity: function numPendingGroups() view returns(uint256)
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
// Solidity: function numPendingGroups() view returns(uint256)
func (_Dosproxy *DosproxySession) NumPendingGroups() (*big.Int, error) {
	return _Dosproxy.Contract.NumPendingGroups(&_Dosproxy.CallOpts)
}

// NumPendingGroups is a free data retrieval call binding the contract method 0x863bc0a1.
//
// Solidity: function numPendingGroups() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) NumPendingGroups() (*big.Int, error) {
	return _Dosproxy.Contract.NumPendingGroups(&_Dosproxy.CallOpts)
}

// NumPendingNodes is a free data retrieval call binding the contract method 0xf41a1587.
//
// Solidity: function numPendingNodes() view returns(uint256)
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
// Solidity: function numPendingNodes() view returns(uint256)
func (_Dosproxy *DosproxySession) NumPendingNodes() (*big.Int, error) {
	return _Dosproxy.Contract.NumPendingNodes(&_Dosproxy.CallOpts)
}

// NumPendingNodes is a free data retrieval call binding the contract method 0xf41a1587.
//
// Solidity: function numPendingNodes() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) NumPendingNodes() (*big.Int, error) {
	return _Dosproxy.Contract.NumPendingNodes(&_Dosproxy.CallOpts)
}

// PendingGroupList is a free data retrieval call binding the contract method 0xa54fb00e.
//
// Solidity: function pendingGroupList(uint256 ) view returns(uint256)
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
// Solidity: function pendingGroupList(uint256 ) view returns(uint256)
func (_Dosproxy *DosproxySession) PendingGroupList(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupList(&_Dosproxy.CallOpts, arg0)
}

// PendingGroupList is a free data retrieval call binding the contract method 0xa54fb00e.
//
// Solidity: function pendingGroupList(uint256 ) view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) PendingGroupList(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupList(&_Dosproxy.CallOpts, arg0)
}

// PendingGroupMaxLife is a free data retrieval call binding the contract method 0xfc84dde4.
//
// Solidity: function pendingGroupMaxLife() view returns(uint256)
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
// Solidity: function pendingGroupMaxLife() view returns(uint256)
func (_Dosproxy *DosproxySession) PendingGroupMaxLife() (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupMaxLife(&_Dosproxy.CallOpts)
}

// PendingGroupMaxLife is a free data retrieval call binding the contract method 0xfc84dde4.
//
// Solidity: function pendingGroupMaxLife() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) PendingGroupMaxLife() (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupMaxLife(&_Dosproxy.CallOpts)
}

// PendingGroupTail is a free data retrieval call binding the contract method 0x190ca29e.
//
// Solidity: function pendingGroupTail() view returns(uint256)
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
// Solidity: function pendingGroupTail() view returns(uint256)
func (_Dosproxy *DosproxySession) PendingGroupTail() (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupTail(&_Dosproxy.CallOpts)
}

// PendingGroupTail is a free data retrieval call binding the contract method 0x190ca29e.
//
// Solidity: function pendingGroupTail() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) PendingGroupTail() (*big.Int, error) {
	return _Dosproxy.Contract.PendingGroupTail(&_Dosproxy.CallOpts)
}

// PendingGroups is a free data retrieval call binding the contract method 0xd18c81b7.
//
// Solidity: function pendingGroups(uint256 ) view returns(uint256 groupId, uint256 startBlkNum)
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
// Solidity: function pendingGroups(uint256 ) view returns(uint256 groupId, uint256 startBlkNum)
func (_Dosproxy *DosproxySession) PendingGroups(arg0 *big.Int) (struct {
	GroupId     *big.Int
	StartBlkNum *big.Int
}, error) {
	return _Dosproxy.Contract.PendingGroups(&_Dosproxy.CallOpts, arg0)
}

// PendingGroups is a free data retrieval call binding the contract method 0xd18c81b7.
//
// Solidity: function pendingGroups(uint256 ) view returns(uint256 groupId, uint256 startBlkNum)
func (_Dosproxy *DosproxyCallerSession) PendingGroups(arg0 *big.Int) (struct {
	GroupId     *big.Int
	StartBlkNum *big.Int
}, error) {
	return _Dosproxy.Contract.PendingGroups(&_Dosproxy.CallOpts, arg0)
}

// PendingNodeList is a free data retrieval call binding the contract method 0xa60b007d.
//
// Solidity: function pendingNodeList(address ) view returns(address)
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
// Solidity: function pendingNodeList(address ) view returns(address)
func (_Dosproxy *DosproxySession) PendingNodeList(arg0 common.Address) (common.Address, error) {
	return _Dosproxy.Contract.PendingNodeList(&_Dosproxy.CallOpts, arg0)
}

// PendingNodeList is a free data retrieval call binding the contract method 0xa60b007d.
//
// Solidity: function pendingNodeList(address ) view returns(address)
func (_Dosproxy *DosproxyCallerSession) PendingNodeList(arg0 common.Address) (common.Address, error) {
	return _Dosproxy.Contract.PendingNodeList(&_Dosproxy.CallOpts, arg0)
}

// PendingNodeTail is a free data retrieval call binding the contract method 0x094c3612.
//
// Solidity: function pendingNodeTail() view returns(address)
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
// Solidity: function pendingNodeTail() view returns(address)
func (_Dosproxy *DosproxySession) PendingNodeTail() (common.Address, error) {
	return _Dosproxy.Contract.PendingNodeTail(&_Dosproxy.CallOpts)
}

// PendingNodeTail is a free data retrieval call binding the contract method 0x094c3612.
//
// Solidity: function pendingNodeTail() view returns(address)
func (_Dosproxy *DosproxyCallerSession) PendingNodeTail() (common.Address, error) {
	return _Dosproxy.Contract.PendingNodeTail(&_Dosproxy.CallOpts)
}

// ProxyFundsAddr is a free data retrieval call binding the contract method 0xdf37c617.
//
// Solidity: function proxyFundsAddr() view returns(address)
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
// Solidity: function proxyFundsAddr() view returns(address)
func (_Dosproxy *DosproxySession) ProxyFundsAddr() (common.Address, error) {
	return _Dosproxy.Contract.ProxyFundsAddr(&_Dosproxy.CallOpts)
}

// ProxyFundsAddr is a free data retrieval call binding the contract method 0xdf37c617.
//
// Solidity: function proxyFundsAddr() view returns(address)
func (_Dosproxy *DosproxyCallerSession) ProxyFundsAddr() (common.Address, error) {
	return _Dosproxy.Contract.ProxyFundsAddr(&_Dosproxy.CallOpts)
}

// ProxyFundsTokenAddr is a free data retrieval call binding the contract method 0x99ca2d30.
//
// Solidity: function proxyFundsTokenAddr() view returns(address)
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
// Solidity: function proxyFundsTokenAddr() view returns(address)
func (_Dosproxy *DosproxySession) ProxyFundsTokenAddr() (common.Address, error) {
	return _Dosproxy.Contract.ProxyFundsTokenAddr(&_Dosproxy.CallOpts)
}

// ProxyFundsTokenAddr is a free data retrieval call binding the contract method 0x99ca2d30.
//
// Solidity: function proxyFundsTokenAddr() view returns(address)
func (_Dosproxy *DosproxyCallerSession) ProxyFundsTokenAddr() (common.Address, error) {
	return _Dosproxy.Contract.ProxyFundsTokenAddr(&_Dosproxy.CallOpts)
}

// RefreshSystemRandomHardLimit is a free data retrieval call binding the contract method 0x962ba8a4.
//
// Solidity: function refreshSystemRandomHardLimit() view returns(uint256)
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
// Solidity: function refreshSystemRandomHardLimit() view returns(uint256)
func (_Dosproxy *DosproxySession) RefreshSystemRandomHardLimit() (*big.Int, error) {
	return _Dosproxy.Contract.RefreshSystemRandomHardLimit(&_Dosproxy.CallOpts)
}

// RefreshSystemRandomHardLimit is a free data retrieval call binding the contract method 0x962ba8a4.
//
// Solidity: function refreshSystemRandomHardLimit() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) RefreshSystemRandomHardLimit() (*big.Int, error) {
	return _Dosproxy.Contract.RefreshSystemRandomHardLimit(&_Dosproxy.CallOpts)
}

// RelayRespondLimit is a free data retrieval call binding the contract method 0x6c460899.
//
// Solidity: function relayRespondLimit() view returns(uint256)
func (_Dosproxy *DosproxyCaller) RelayRespondLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dosproxy.contract.Call(opts, out, "relayRespondLimit")
	return *ret0, err
}

// RelayRespondLimit is a free data retrieval call binding the contract method 0x6c460899.
//
// Solidity: function relayRespondLimit() view returns(uint256)
func (_Dosproxy *DosproxySession) RelayRespondLimit() (*big.Int, error) {
	return _Dosproxy.Contract.RelayRespondLimit(&_Dosproxy.CallOpts)
}

// RelayRespondLimit is a free data retrieval call binding the contract method 0x6c460899.
//
// Solidity: function relayRespondLimit() view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) RelayRespondLimit() (*big.Int, error) {
	return _Dosproxy.Contract.RelayRespondLimit(&_Dosproxy.CallOpts)
}

// WorkingGroupIds is a free data retrieval call binding the contract method 0x5d381204.
//
// Solidity: function workingGroupIds(uint256 ) view returns(uint256)
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
// Solidity: function workingGroupIds(uint256 ) view returns(uint256)
func (_Dosproxy *DosproxySession) WorkingGroupIds(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.WorkingGroupIds(&_Dosproxy.CallOpts, arg0)
}

// WorkingGroupIds is a free data retrieval call binding the contract method 0x5d381204.
//
// Solidity: function workingGroupIds(uint256 ) view returns(uint256)
func (_Dosproxy *DosproxyCallerSession) WorkingGroupIds(arg0 *big.Int) (*big.Int, error) {
	return _Dosproxy.Contract.WorkingGroupIds(&_Dosproxy.CallOpts, arg0)
}

// Callback is a paid mutator transaction binding the contract method 0x18a1908d.
//
// Solidity: function __callback__(uint256 requestId, uint256 rndSeed) returns()
func (_Dosproxy *DosproxyTransactor) Callback(opts *bind.TransactOpts, requestId *big.Int, rndSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "__callback__", requestId, rndSeed)
}

// Callback is a paid mutator transaction binding the contract method 0x18a1908d.
//
// Solidity: function __callback__(uint256 requestId, uint256 rndSeed) returns()
func (_Dosproxy *DosproxySession) Callback(requestId *big.Int, rndSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.Callback(&_Dosproxy.TransactOpts, requestId, rndSeed)
}

// Callback is a paid mutator transaction binding the contract method 0x18a1908d.
//
// Solidity: function __callback__(uint256 requestId, uint256 rndSeed) returns()
func (_Dosproxy *DosproxyTransactorSession) Callback(requestId *big.Int, rndSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.Callback(&_Dosproxy.TransactOpts, requestId, rndSeed)
}

// AddToGuardianList is a paid mutator transaction binding the contract method 0xd79351b2.
//
// Solidity: function addToGuardianList(address _addr) returns()
func (_Dosproxy *DosproxyTransactor) AddToGuardianList(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "addToGuardianList", _addr)
}

// AddToGuardianList is a paid mutator transaction binding the contract method 0xd79351b2.
//
// Solidity: function addToGuardianList(address _addr) returns()
func (_Dosproxy *DosproxySession) AddToGuardianList(_addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.AddToGuardianList(&_Dosproxy.TransactOpts, _addr)
}

// AddToGuardianList is a paid mutator transaction binding the contract method 0xd79351b2.
//
// Solidity: function addToGuardianList(address _addr) returns()
func (_Dosproxy *DosproxyTransactorSession) AddToGuardianList(_addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.AddToGuardianList(&_Dosproxy.TransactOpts, _addr)
}

// Query is a paid mutator transaction binding the contract method 0xb7fb8fd7.
//
// Solidity: function query(address from, uint256 timeout, string dataSource, string selector) returns(uint256)
func (_Dosproxy *DosproxyTransactor) Query(opts *bind.TransactOpts, from common.Address, timeout *big.Int, dataSource string, selector string) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "query", from, timeout, dataSource, selector)
}

// Query is a paid mutator transaction binding the contract method 0xb7fb8fd7.
//
// Solidity: function query(address from, uint256 timeout, string dataSource, string selector) returns(uint256)
func (_Dosproxy *DosproxySession) Query(from common.Address, timeout *big.Int, dataSource string, selector string) (*types.Transaction, error) {
	return _Dosproxy.Contract.Query(&_Dosproxy.TransactOpts, from, timeout, dataSource, selector)
}

// Query is a paid mutator transaction binding the contract method 0xb7fb8fd7.
//
// Solidity: function query(address from, uint256 timeout, string dataSource, string selector) returns(uint256)
func (_Dosproxy *DosproxyTransactorSession) Query(from common.Address, timeout *big.Int, dataSource string, selector string) (*types.Transaction, error) {
	return _Dosproxy.Contract.Query(&_Dosproxy.TransactOpts, from, timeout, dataSource, selector)
}

// RegisterGroupPubKey is a paid mutator transaction binding the contract method 0xb836ccea.
//
// Solidity: function registerGroupPubKey(uint256 groupId, uint256[4] suggestedPubKey) returns()
func (_Dosproxy *DosproxyTransactor) RegisterGroupPubKey(opts *bind.TransactOpts, groupId *big.Int, suggestedPubKey [4]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "registerGroupPubKey", groupId, suggestedPubKey)
}

// RegisterGroupPubKey is a paid mutator transaction binding the contract method 0xb836ccea.
//
// Solidity: function registerGroupPubKey(uint256 groupId, uint256[4] suggestedPubKey) returns()
func (_Dosproxy *DosproxySession) RegisterGroupPubKey(groupId *big.Int, suggestedPubKey [4]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.RegisterGroupPubKey(&_Dosproxy.TransactOpts, groupId, suggestedPubKey)
}

// RegisterGroupPubKey is a paid mutator transaction binding the contract method 0xb836ccea.
//
// Solidity: function registerGroupPubKey(uint256 groupId, uint256[4] suggestedPubKey) returns()
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
// Solidity: function removeFromGuardianList(address _addr) returns()
func (_Dosproxy *DosproxyTransactor) RemoveFromGuardianList(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "removeFromGuardianList", _addr)
}

// RemoveFromGuardianList is a paid mutator transaction binding the contract method 0xc58ebe1c.
//
// Solidity: function removeFromGuardianList(address _addr) returns()
func (_Dosproxy *DosproxySession) RemoveFromGuardianList(_addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.RemoveFromGuardianList(&_Dosproxy.TransactOpts, _addr)
}

// RemoveFromGuardianList is a paid mutator transaction binding the contract method 0xc58ebe1c.
//
// Solidity: function removeFromGuardianList(address _addr) returns()
func (_Dosproxy *DosproxyTransactorSession) RemoveFromGuardianList(_addr common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.RemoveFromGuardianList(&_Dosproxy.TransactOpts, _addr)
}

// RequestRandom is a paid mutator transaction binding the contract method 0xc7c3f4a5.
//
// Solidity: function requestRandom(address from, uint256 userSeed) returns(uint256)
func (_Dosproxy *DosproxyTransactor) RequestRandom(opts *bind.TransactOpts, from common.Address, userSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "requestRandom", from, userSeed)
}

// RequestRandom is a paid mutator transaction binding the contract method 0xc7c3f4a5.
//
// Solidity: function requestRandom(address from, uint256 userSeed) returns(uint256)
func (_Dosproxy *DosproxySession) RequestRandom(from common.Address, userSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.RequestRandom(&_Dosproxy.TransactOpts, from, userSeed)
}

// RequestRandom is a paid mutator transaction binding the contract method 0xc7c3f4a5.
//
// Solidity: function requestRandom(address from, uint256 userSeed) returns(uint256)
func (_Dosproxy *DosproxyTransactorSession) RequestRandom(from common.Address, userSeed *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.RequestRandom(&_Dosproxy.TransactOpts, from, userSeed)
}

// ResetOnRecovery is a paid mutator transaction binding the contract method 0xe1ff7916.
//
// Solidity: function resetOnRecovery(uint256 cap) returns()
func (_Dosproxy *DosproxyTransactor) ResetOnRecovery(opts *bind.TransactOpts, cap *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "resetOnRecovery", cap)
}

// ResetOnRecovery is a paid mutator transaction binding the contract method 0xe1ff7916.
//
// Solidity: function resetOnRecovery(uint256 cap) returns()
func (_Dosproxy *DosproxySession) ResetOnRecovery(cap *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.ResetOnRecovery(&_Dosproxy.TransactOpts, cap)
}

// ResetOnRecovery is a paid mutator transaction binding the contract method 0xe1ff7916.
//
// Solidity: function resetOnRecovery(uint256 cap) returns()
func (_Dosproxy *DosproxyTransactorSession) ResetOnRecovery(cap *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.ResetOnRecovery(&_Dosproxy.TransactOpts, cap)
}

// SetBootstrapCommitDuration is a paid mutator transaction binding the contract method 0x9e718573.
//
// Solidity: function setBootstrapCommitDuration(uint256 newDuration) returns()
func (_Dosproxy *DosproxyTransactor) SetBootstrapCommitDuration(opts *bind.TransactOpts, newDuration *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setBootstrapCommitDuration", newDuration)
}

// SetBootstrapCommitDuration is a paid mutator transaction binding the contract method 0x9e718573.
//
// Solidity: function setBootstrapCommitDuration(uint256 newDuration) returns()
func (_Dosproxy *DosproxySession) SetBootstrapCommitDuration(newDuration *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetBootstrapCommitDuration(&_Dosproxy.TransactOpts, newDuration)
}

// SetBootstrapCommitDuration is a paid mutator transaction binding the contract method 0x9e718573.
//
// Solidity: function setBootstrapCommitDuration(uint256 newDuration) returns()
func (_Dosproxy *DosproxyTransactorSession) SetBootstrapCommitDuration(newDuration *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetBootstrapCommitDuration(&_Dosproxy.TransactOpts, newDuration)
}

// SetBootstrapRevealDuration is a paid mutator transaction binding the contract method 0xb8fa82e0.
//
// Solidity: function setBootstrapRevealDuration(uint256 newDuration) returns()
func (_Dosproxy *DosproxyTransactor) SetBootstrapRevealDuration(opts *bind.TransactOpts, newDuration *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setBootstrapRevealDuration", newDuration)
}

// SetBootstrapRevealDuration is a paid mutator transaction binding the contract method 0xb8fa82e0.
//
// Solidity: function setBootstrapRevealDuration(uint256 newDuration) returns()
func (_Dosproxy *DosproxySession) SetBootstrapRevealDuration(newDuration *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetBootstrapRevealDuration(&_Dosproxy.TransactOpts, newDuration)
}

// SetBootstrapRevealDuration is a paid mutator transaction binding the contract method 0xb8fa82e0.
//
// Solidity: function setBootstrapRevealDuration(uint256 newDuration) returns()
func (_Dosproxy *DosproxyTransactorSession) SetBootstrapRevealDuration(newDuration *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetBootstrapRevealDuration(&_Dosproxy.TransactOpts, newDuration)
}

// SetBootstrapStartThreshold is a paid mutator transaction binding the contract method 0x100063ec.
//
// Solidity: function setBootstrapStartThreshold(uint256 newThreshold) returns()
func (_Dosproxy *DosproxyTransactor) SetBootstrapStartThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setBootstrapStartThreshold", newThreshold)
}

// SetBootstrapStartThreshold is a paid mutator transaction binding the contract method 0x100063ec.
//
// Solidity: function setBootstrapStartThreshold(uint256 newThreshold) returns()
func (_Dosproxy *DosproxySession) SetBootstrapStartThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetBootstrapStartThreshold(&_Dosproxy.TransactOpts, newThreshold)
}

// SetBootstrapStartThreshold is a paid mutator transaction binding the contract method 0x100063ec.
//
// Solidity: function setBootstrapStartThreshold(uint256 newThreshold) returns()
func (_Dosproxy *DosproxyTransactorSession) SetBootstrapStartThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetBootstrapStartThreshold(&_Dosproxy.TransactOpts, newThreshold)
}

// SetGroupMaturityPeriod is a paid mutator transaction binding the contract method 0x925fc6c9.
//
// Solidity: function setGroupMaturityPeriod(uint256 newPeriod) returns()
func (_Dosproxy *DosproxyTransactor) SetGroupMaturityPeriod(opts *bind.TransactOpts, newPeriod *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setGroupMaturityPeriod", newPeriod)
}

// SetGroupMaturityPeriod is a paid mutator transaction binding the contract method 0x925fc6c9.
//
// Solidity: function setGroupMaturityPeriod(uint256 newPeriod) returns()
func (_Dosproxy *DosproxySession) SetGroupMaturityPeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetGroupMaturityPeriod(&_Dosproxy.TransactOpts, newPeriod)
}

// SetGroupMaturityPeriod is a paid mutator transaction binding the contract method 0x925fc6c9.
//
// Solidity: function setGroupMaturityPeriod(uint256 newPeriod) returns()
func (_Dosproxy *DosproxyTransactorSession) SetGroupMaturityPeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetGroupMaturityPeriod(&_Dosproxy.TransactOpts, newPeriod)
}

// SetGroupSize is a paid mutator transaction binding the contract method 0x0dfc09cb.
//
// Solidity: function setGroupSize(uint256 newSize) returns()
func (_Dosproxy *DosproxyTransactor) SetGroupSize(opts *bind.TransactOpts, newSize *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setGroupSize", newSize)
}

// SetGroupSize is a paid mutator transaction binding the contract method 0x0dfc09cb.
//
// Solidity: function setGroupSize(uint256 newSize) returns()
func (_Dosproxy *DosproxySession) SetGroupSize(newSize *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetGroupSize(&_Dosproxy.TransactOpts, newSize)
}

// SetGroupSize is a paid mutator transaction binding the contract method 0x0dfc09cb.
//
// Solidity: function setGroupSize(uint256 newSize) returns()
func (_Dosproxy *DosproxyTransactorSession) SetGroupSize(newSize *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetGroupSize(&_Dosproxy.TransactOpts, newSize)
}

// SetLifeDiversity is a paid mutator transaction binding the contract method 0x559ea9de.
//
// Solidity: function setLifeDiversity(uint256 newDiversity) returns()
func (_Dosproxy *DosproxyTransactor) SetLifeDiversity(opts *bind.TransactOpts, newDiversity *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setLifeDiversity", newDiversity)
}

// SetLifeDiversity is a paid mutator transaction binding the contract method 0x559ea9de.
//
// Solidity: function setLifeDiversity(uint256 newDiversity) returns()
func (_Dosproxy *DosproxySession) SetLifeDiversity(newDiversity *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetLifeDiversity(&_Dosproxy.TransactOpts, newDiversity)
}

// SetLifeDiversity is a paid mutator transaction binding the contract method 0x559ea9de.
//
// Solidity: function setLifeDiversity(uint256 newDiversity) returns()
func (_Dosproxy *DosproxyTransactorSession) SetLifeDiversity(newDiversity *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetLifeDiversity(&_Dosproxy.TransactOpts, newDiversity)
}

// SetPendingGroupMaxLife is a paid mutator transaction binding the contract method 0x4a28a74d.
//
// Solidity: function setPendingGroupMaxLife(uint256 newLife) returns()
func (_Dosproxy *DosproxyTransactor) SetPendingGroupMaxLife(opts *bind.TransactOpts, newLife *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setPendingGroupMaxLife", newLife)
}

// SetPendingGroupMaxLife is a paid mutator transaction binding the contract method 0x4a28a74d.
//
// Solidity: function setPendingGroupMaxLife(uint256 newLife) returns()
func (_Dosproxy *DosproxySession) SetPendingGroupMaxLife(newLife *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetPendingGroupMaxLife(&_Dosproxy.TransactOpts, newLife)
}

// SetPendingGroupMaxLife is a paid mutator transaction binding the contract method 0x4a28a74d.
//
// Solidity: function setPendingGroupMaxLife(uint256 newLife) returns()
func (_Dosproxy *DosproxyTransactorSession) SetPendingGroupMaxLife(newLife *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetPendingGroupMaxLife(&_Dosproxy.TransactOpts, newLife)
}

// SetProxyFund is a paid mutator transaction binding the contract method 0x40e4a5af.
//
// Solidity: function setProxyFund(address newFund, address newFundToken) returns()
func (_Dosproxy *DosproxyTransactor) SetProxyFund(opts *bind.TransactOpts, newFund common.Address, newFundToken common.Address) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setProxyFund", newFund, newFundToken)
}

// SetProxyFund is a paid mutator transaction binding the contract method 0x40e4a5af.
//
// Solidity: function setProxyFund(address newFund, address newFundToken) returns()
func (_Dosproxy *DosproxySession) SetProxyFund(newFund common.Address, newFundToken common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetProxyFund(&_Dosproxy.TransactOpts, newFund, newFundToken)
}

// SetProxyFund is a paid mutator transaction binding the contract method 0x40e4a5af.
//
// Solidity: function setProxyFund(address newFund, address newFundToken) returns()
func (_Dosproxy *DosproxyTransactorSession) SetProxyFund(newFund common.Address, newFundToken common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetProxyFund(&_Dosproxy.TransactOpts, newFund, newFundToken)
}

// SetRelayRespondLimit is a paid mutator transaction binding the contract method 0xca2aabcc.
//
// Solidity: function setRelayRespondLimit(uint256 newLimit) returns()
func (_Dosproxy *DosproxyTransactor) SetRelayRespondLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setRelayRespondLimit", newLimit)
}

// SetRelayRespondLimit is a paid mutator transaction binding the contract method 0xca2aabcc.
//
// Solidity: function setRelayRespondLimit(uint256 newLimit) returns()
func (_Dosproxy *DosproxySession) SetRelayRespondLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetRelayRespondLimit(&_Dosproxy.TransactOpts, newLimit)
}

// SetRelayRespondLimit is a paid mutator transaction binding the contract method 0xca2aabcc.
//
// Solidity: function setRelayRespondLimit(uint256 newLimit) returns()
func (_Dosproxy *DosproxyTransactorSession) SetRelayRespondLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetRelayRespondLimit(&_Dosproxy.TransactOpts, newLimit)
}

// SetSystemRandomHardLimit is a paid mutator transaction binding the contract method 0xc457aa8f.
//
// Solidity: function setSystemRandomHardLimit(uint256 newLimit) returns()
func (_Dosproxy *DosproxyTransactor) SetSystemRandomHardLimit(opts *bind.TransactOpts, newLimit *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "setSystemRandomHardLimit", newLimit)
}

// SetSystemRandomHardLimit is a paid mutator transaction binding the contract method 0xc457aa8f.
//
// Solidity: function setSystemRandomHardLimit(uint256 newLimit) returns()
func (_Dosproxy *DosproxySession) SetSystemRandomHardLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetSystemRandomHardLimit(&_Dosproxy.TransactOpts, newLimit)
}

// SetSystemRandomHardLimit is a paid mutator transaction binding the contract method 0xc457aa8f.
//
// Solidity: function setSystemRandomHardLimit(uint256 newLimit) returns()
func (_Dosproxy *DosproxyTransactorSession) SetSystemRandomHardLimit(newLimit *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SetSystemRandomHardLimit(&_Dosproxy.TransactOpts, newLimit)
}

// SignalBootstrap is a paid mutator transaction binding the contract method 0x5c0e159f.
//
// Solidity: function signalBootstrap(uint256 _cid) returns()
func (_Dosproxy *DosproxyTransactor) SignalBootstrap(opts *bind.TransactOpts, _cid *big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "signalBootstrap", _cid)
}

// SignalBootstrap is a paid mutator transaction binding the contract method 0x5c0e159f.
//
// Solidity: function signalBootstrap(uint256 _cid) returns()
func (_Dosproxy *DosproxySession) SignalBootstrap(_cid *big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalBootstrap(&_Dosproxy.TransactOpts, _cid)
}

// SignalBootstrap is a paid mutator transaction binding the contract method 0x5c0e159f.
//
// Solidity: function signalBootstrap(uint256 _cid) returns()
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
// Solidity: function signalUnregister(address member) returns()
func (_Dosproxy *DosproxyTransactor) SignalUnregister(opts *bind.TransactOpts, member common.Address) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "signalUnregister", member)
}

// SignalUnregister is a paid mutator transaction binding the contract method 0x7c1cf083.
//
// Solidity: function signalUnregister(address member) returns()
func (_Dosproxy *DosproxySession) SignalUnregister(member common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalUnregister(&_Dosproxy.TransactOpts, member)
}

// SignalUnregister is a paid mutator transaction binding the contract method 0x7c1cf083.
//
// Solidity: function signalUnregister(address member) returns()
func (_Dosproxy *DosproxyTransactorSession) SignalUnregister(member common.Address) (*types.Transaction, error) {
	return _Dosproxy.Contract.SignalUnregister(&_Dosproxy.TransactOpts, member)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x74ad3a06.
//
// Solidity: function triggerCallback(uint256 requestId, uint8 trafficType, bytes result, uint256[2] sig) returns()
func (_Dosproxy *DosproxyTransactor) TriggerCallback(opts *bind.TransactOpts, requestId *big.Int, trafficType uint8, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "triggerCallback", requestId, trafficType, result, sig)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x74ad3a06.
//
// Solidity: function triggerCallback(uint256 requestId, uint8 trafficType, bytes result, uint256[2] sig) returns()
func (_Dosproxy *DosproxySession) TriggerCallback(requestId *big.Int, trafficType uint8, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.TriggerCallback(&_Dosproxy.TransactOpts, requestId, trafficType, result, sig)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x74ad3a06.
//
// Solidity: function triggerCallback(uint256 requestId, uint8 trafficType, bytes result, uint256[2] sig) returns()
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
// Solidity: function updateRandomness(uint256[2] sig) returns()
func (_Dosproxy *DosproxyTransactor) UpdateRandomness(opts *bind.TransactOpts, sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.contract.Transact(opts, "updateRandomness", sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(uint256[2] sig) returns()
func (_Dosproxy *DosproxySession) UpdateRandomness(sig [2]*big.Int) (*types.Transaction, error) {
	return _Dosproxy.Contract.UpdateRandomness(&_Dosproxy.TransactOpts, sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(uint256[2] sig) returns()
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
// Solidity: event GuardianReward(uint256 blkNum, address guardian)
func (_Dosproxy *DosproxyFilterer) FilterGuardianReward(opts *bind.FilterOpts) (*DosproxyGuardianRewardIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "GuardianReward")
	if err != nil {
		return nil, err
	}
	return &DosproxyGuardianRewardIterator{contract: _Dosproxy.contract, event: "GuardianReward", logs: logs, sub: sub}, nil
}

// WatchGuardianReward is a free log subscription operation binding the contract event 0xa60d55093b21f740878d9871e95e5031eaf5cf08a167c898ed3c62b1fb24f887.
//
// Solidity: event GuardianReward(uint256 blkNum, address guardian)
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

// ParseGuardianReward is a log parse operation binding the contract event 0xa60d55093b21f740878d9871e95e5031eaf5cf08a167c898ed3c62b1fb24f887.
//
// Solidity: event GuardianReward(uint256 blkNum, address guardian)
func (_Dosproxy *DosproxyFilterer) ParseGuardianReward(log types.Log) (*DosproxyGuardianReward, error) {
	event := new(DosproxyGuardianReward)
	if err := _Dosproxy.contract.UnpackLog(event, "GuardianReward", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogCallbackTriggeredFor(address callbackAddr)
func (_Dosproxy *DosproxyFilterer) FilterLogCallbackTriggeredFor(opts *bind.FilterOpts) (*DosproxyLogCallbackTriggeredForIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogCallbackTriggeredFor")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogCallbackTriggeredForIterator{contract: _Dosproxy.contract, event: "LogCallbackTriggeredFor", logs: logs, sub: sub}, nil
}

// WatchLogCallbackTriggeredFor is a free log subscription operation binding the contract event 0x065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf0.
//
// Solidity: event LogCallbackTriggeredFor(address callbackAddr)
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

// ParseLogCallbackTriggeredFor is a log parse operation binding the contract event 0x065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf0.
//
// Solidity: event LogCallbackTriggeredFor(address callbackAddr)
func (_Dosproxy *DosproxyFilterer) ParseLogCallbackTriggeredFor(log types.Log) (*DosproxyLogCallbackTriggeredFor, error) {
	event := new(DosproxyLogCallbackTriggeredFor)
	if err := _Dosproxy.contract.UnpackLog(event, "LogCallbackTriggeredFor", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogGroupDissolve(uint256 groupId)
func (_Dosproxy *DosproxyFilterer) FilterLogGroupDissolve(opts *bind.FilterOpts) (*DosproxyLogGroupDissolveIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogGroupDissolve")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogGroupDissolveIterator{contract: _Dosproxy.contract, event: "LogGroupDissolve", logs: logs, sub: sub}, nil
}

// WatchLogGroupDissolve is a free log subscription operation binding the contract event 0xf7377b41bdc770cc22a1bad318571f0ae6d65188371bdbe0cb660c0de57b3150.
//
// Solidity: event LogGroupDissolve(uint256 groupId)
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

// ParseLogGroupDissolve is a log parse operation binding the contract event 0xf7377b41bdc770cc22a1bad318571f0ae6d65188371bdbe0cb660c0de57b3150.
//
// Solidity: event LogGroupDissolve(uint256 groupId)
func (_Dosproxy *DosproxyFilterer) ParseLogGroupDissolve(log types.Log) (*DosproxyLogGroupDissolve, error) {
	event := new(DosproxyLogGroupDissolve)
	if err := _Dosproxy.contract.UnpackLog(event, "LogGroupDissolve", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogGrouping(uint256 groupId, address[] nodeId)
func (_Dosproxy *DosproxyFilterer) FilterLogGrouping(opts *bind.FilterOpts) (*DosproxyLogGroupingIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogGrouping")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogGroupingIterator{contract: _Dosproxy.contract, event: "LogGrouping", logs: logs, sub: sub}, nil
}

// WatchLogGrouping is a free log subscription operation binding the contract event 0x78bf54a42d1b98e6c809c3e5904898c5b3304ede546b6f070e83d9a32e544d4f.
//
// Solidity: event LogGrouping(uint256 groupId, address[] nodeId)
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

// ParseLogGrouping is a log parse operation binding the contract event 0x78bf54a42d1b98e6c809c3e5904898c5b3304ede546b6f070e83d9a32e544d4f.
//
// Solidity: event LogGrouping(uint256 groupId, address[] nodeId)
func (_Dosproxy *DosproxyFilterer) ParseLogGrouping(log types.Log) (*DosproxyLogGrouping, error) {
	event := new(DosproxyLogGrouping)
	if err := _Dosproxy.contract.UnpackLog(event, "LogGrouping", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogGroupingInitiated(uint256 pendingNodePool, uint256 groupsize)
func (_Dosproxy *DosproxyFilterer) FilterLogGroupingInitiated(opts *bind.FilterOpts) (*DosproxyLogGroupingInitiatedIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogGroupingInitiated")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogGroupingInitiatedIterator{contract: _Dosproxy.contract, event: "LogGroupingInitiated", logs: logs, sub: sub}, nil
}

// WatchLogGroupingInitiated is a free log subscription operation binding the contract event 0x5ee79c79dd5be1f25f86c7028ec8fb6849839e8a23f676afc907eade45946986.
//
// Solidity: event LogGroupingInitiated(uint256 pendingNodePool, uint256 groupsize)
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

// ParseLogGroupingInitiated is a log parse operation binding the contract event 0x5ee79c79dd5be1f25f86c7028ec8fb6849839e8a23f676afc907eade45946986.
//
// Solidity: event LogGroupingInitiated(uint256 pendingNodePool, uint256 groupsize)
func (_Dosproxy *DosproxyFilterer) ParseLogGroupingInitiated(log types.Log) (*DosproxyLogGroupingInitiated, error) {
	event := new(DosproxyLogGroupingInitiated)
	if err := _Dosproxy.contract.UnpackLog(event, "LogGroupingInitiated", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogInsufficientPendingNode(uint256 numPendingNodes)
func (_Dosproxy *DosproxyFilterer) FilterLogInsufficientPendingNode(opts *bind.FilterOpts) (*DosproxyLogInsufficientPendingNodeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogInsufficientPendingNode")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogInsufficientPendingNodeIterator{contract: _Dosproxy.contract, event: "LogInsufficientPendingNode", logs: logs, sub: sub}, nil
}

// WatchLogInsufficientPendingNode is a free log subscription operation binding the contract event 0xc03848aa1689c7c291dcc68fa62d6109e13f16b25e89bdef5a6c8598b36b8019.
//
// Solidity: event LogInsufficientPendingNode(uint256 numPendingNodes)
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

// ParseLogInsufficientPendingNode is a log parse operation binding the contract event 0xc03848aa1689c7c291dcc68fa62d6109e13f16b25e89bdef5a6c8598b36b8019.
//
// Solidity: event LogInsufficientPendingNode(uint256 numPendingNodes)
func (_Dosproxy *DosproxyFilterer) ParseLogInsufficientPendingNode(log types.Log) (*DosproxyLogInsufficientPendingNode, error) {
	event := new(DosproxyLogInsufficientPendingNode)
	if err := _Dosproxy.contract.UnpackLog(event, "LogInsufficientPendingNode", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogInsufficientWorkingGroup(uint256 numWorkingGroups, uint256 numPendingGroups)
func (_Dosproxy *DosproxyFilterer) FilterLogInsufficientWorkingGroup(opts *bind.FilterOpts) (*DosproxyLogInsufficientWorkingGroupIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogInsufficientWorkingGroup")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogInsufficientWorkingGroupIterator{contract: _Dosproxy.contract, event: "LogInsufficientWorkingGroup", logs: logs, sub: sub}, nil
}

// WatchLogInsufficientWorkingGroup is a free log subscription operation binding the contract event 0x1850da28de32299250accda835079ca1340fbd447032a1f6dac77381a77a26c8.
//
// Solidity: event LogInsufficientWorkingGroup(uint256 numWorkingGroups, uint256 numPendingGroups)
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

// ParseLogInsufficientWorkingGroup is a log parse operation binding the contract event 0x1850da28de32299250accda835079ca1340fbd447032a1f6dac77381a77a26c8.
//
// Solidity: event LogInsufficientWorkingGroup(uint256 numWorkingGroups, uint256 numPendingGroups)
func (_Dosproxy *DosproxyFilterer) ParseLogInsufficientWorkingGroup(log types.Log) (*DosproxyLogInsufficientWorkingGroup, error) {
	event := new(DosproxyLogInsufficientWorkingGroup)
	if err := _Dosproxy.contract.UnpackLog(event, "LogInsufficientWorkingGroup", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogMessage(string info)
func (_Dosproxy *DosproxyFilterer) FilterLogMessage(opts *bind.FilterOpts) (*DosproxyLogMessageIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogMessage")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogMessageIterator{contract: _Dosproxy.contract, event: "LogMessage", logs: logs, sub: sub}, nil
}

// WatchLogMessage is a free log subscription operation binding the contract event 0x96561394bac381230de4649200e8831afcab1f451881bbade9ef209f6dd30480.
//
// Solidity: event LogMessage(string info)
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

// ParseLogMessage is a log parse operation binding the contract event 0x96561394bac381230de4649200e8831afcab1f451881bbade9ef209f6dd30480.
//
// Solidity: event LogMessage(string info)
func (_Dosproxy *DosproxyFilterer) ParseLogMessage(log types.Log) (*DosproxyLogMessage, error) {
	event := new(DosproxyLogMessage)
	if err := _Dosproxy.contract.UnpackLog(event, "LogMessage", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogNoPendingGroup(uint256 groupId)
func (_Dosproxy *DosproxyFilterer) FilterLogNoPendingGroup(opts *bind.FilterOpts) (*DosproxyLogNoPendingGroupIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogNoPendingGroup")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogNoPendingGroupIterator{contract: _Dosproxy.contract, event: "LogNoPendingGroup", logs: logs, sub: sub}, nil
}

// WatchLogNoPendingGroup is a free log subscription operation binding the contract event 0x71047c0893a51085656a2894bba10bc6ef51a654f25e1ead1929b076487a9569.
//
// Solidity: event LogNoPendingGroup(uint256 groupId)
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

// ParseLogNoPendingGroup is a log parse operation binding the contract event 0x71047c0893a51085656a2894bba10bc6ef51a654f25e1ead1929b076487a9569.
//
// Solidity: event LogNoPendingGroup(uint256 groupId)
func (_Dosproxy *DosproxyFilterer) ParseLogNoPendingGroup(log types.Log) (*DosproxyLogNoPendingGroup, error) {
	event := new(DosproxyLogNoPendingGroup)
	if err := _Dosproxy.contract.UnpackLog(event, "LogNoPendingGroup", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogNonContractCall(address from)
func (_Dosproxy *DosproxyFilterer) FilterLogNonContractCall(opts *bind.FilterOpts) (*DosproxyLogNonContractCallIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogNonContractCall")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogNonContractCallIterator{contract: _Dosproxy.contract, event: "LogNonContractCall", logs: logs, sub: sub}, nil
}

// WatchLogNonContractCall is a free log subscription operation binding the contract event 0x6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb5.
//
// Solidity: event LogNonContractCall(address from)
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

// ParseLogNonContractCall is a log parse operation binding the contract event 0x6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb5.
//
// Solidity: event LogNonContractCall(address from)
func (_Dosproxy *DosproxyFilterer) ParseLogNonContractCall(log types.Log) (*DosproxyLogNonContractCall, error) {
	event := new(DosproxyLogNonContractCall)
	if err := _Dosproxy.contract.UnpackLog(event, "LogNonContractCall", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogNonSupportedType(string invalidSelector)
func (_Dosproxy *DosproxyFilterer) FilterLogNonSupportedType(opts *bind.FilterOpts) (*DosproxyLogNonSupportedTypeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogNonSupportedType")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogNonSupportedTypeIterator{contract: _Dosproxy.contract, event: "LogNonSupportedType", logs: logs, sub: sub}, nil
}

// WatchLogNonSupportedType is a free log subscription operation binding the contract event 0x70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41.
//
// Solidity: event LogNonSupportedType(string invalidSelector)
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

// ParseLogNonSupportedType is a log parse operation binding the contract event 0x70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41.
//
// Solidity: event LogNonSupportedType(string invalidSelector)
func (_Dosproxy *DosproxyFilterer) ParseLogNonSupportedType(log types.Log) (*DosproxyLogNonSupportedType, error) {
	event := new(DosproxyLogNonSupportedType)
	if err := _Dosproxy.contract.UnpackLog(event, "LogNonSupportedType", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogPendingGroupRemoved(uint256 groupId)
func (_Dosproxy *DosproxyFilterer) FilterLogPendingGroupRemoved(opts *bind.FilterOpts) (*DosproxyLogPendingGroupRemovedIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogPendingGroupRemoved")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogPendingGroupRemovedIterator{contract: _Dosproxy.contract, event: "LogPendingGroupRemoved", logs: logs, sub: sub}, nil
}

// WatchLogPendingGroupRemoved is a free log subscription operation binding the contract event 0x156927b06a61046135669011768a03b2592ee3667374525502b62b0aef4cbd21.
//
// Solidity: event LogPendingGroupRemoved(uint256 groupId)
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

// ParseLogPendingGroupRemoved is a log parse operation binding the contract event 0x156927b06a61046135669011768a03b2592ee3667374525502b62b0aef4cbd21.
//
// Solidity: event LogPendingGroupRemoved(uint256 groupId)
func (_Dosproxy *DosproxyFilterer) ParseLogPendingGroupRemoved(log types.Log) (*DosproxyLogPendingGroupRemoved, error) {
	event := new(DosproxyLogPendingGroupRemoved)
	if err := _Dosproxy.contract.UnpackLog(event, "LogPendingGroupRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogPublicKeyAccepted(uint256 groupId, uint256[4] pubKey, uint256 numWorkingGroups)
func (_Dosproxy *DosproxyFilterer) FilterLogPublicKeyAccepted(opts *bind.FilterOpts) (*DosproxyLogPublicKeyAcceptedIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogPublicKeyAccepted")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogPublicKeyAcceptedIterator{contract: _Dosproxy.contract, event: "LogPublicKeyAccepted", logs: logs, sub: sub}, nil
}

// WatchLogPublicKeyAccepted is a free log subscription operation binding the contract event 0x9ad0d0bfff7f0fc653b03785685d46101b09e3cb1f50081915bc8101662e4a88.
//
// Solidity: event LogPublicKeyAccepted(uint256 groupId, uint256[4] pubKey, uint256 numWorkingGroups)
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

// ParseLogPublicKeyAccepted is a log parse operation binding the contract event 0x9ad0d0bfff7f0fc653b03785685d46101b09e3cb1f50081915bc8101662e4a88.
//
// Solidity: event LogPublicKeyAccepted(uint256 groupId, uint256[4] pubKey, uint256 numWorkingGroups)
func (_Dosproxy *DosproxyFilterer) ParseLogPublicKeyAccepted(log types.Log) (*DosproxyLogPublicKeyAccepted, error) {
	event := new(DosproxyLogPublicKeyAccepted)
	if err := _Dosproxy.contract.UnpackLog(event, "LogPublicKeyAccepted", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogPublicKeySuggested(uint256 groupId, uint256 pubKeyCount)
func (_Dosproxy *DosproxyFilterer) FilterLogPublicKeySuggested(opts *bind.FilterOpts) (*DosproxyLogPublicKeySuggestedIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogPublicKeySuggested")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogPublicKeySuggestedIterator{contract: _Dosproxy.contract, event: "LogPublicKeySuggested", logs: logs, sub: sub}, nil
}

// WatchLogPublicKeySuggested is a free log subscription operation binding the contract event 0x717e526bce26f8e67908004294b35133bbe2a9c7f611384cb0f484aca9223d01.
//
// Solidity: event LogPublicKeySuggested(uint256 groupId, uint256 pubKeyCount)
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

// ParseLogPublicKeySuggested is a log parse operation binding the contract event 0x717e526bce26f8e67908004294b35133bbe2a9c7f611384cb0f484aca9223d01.
//
// Solidity: event LogPublicKeySuggested(uint256 groupId, uint256 pubKeyCount)
func (_Dosproxy *DosproxyFilterer) ParseLogPublicKeySuggested(log types.Log) (*DosproxyLogPublicKeySuggested, error) {
	event := new(DosproxyLogPublicKeySuggested)
	if err := _Dosproxy.contract.UnpackLog(event, "LogPublicKeySuggested", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogRegisteredNewPendingNode(address node)
func (_Dosproxy *DosproxyFilterer) FilterLogRegisteredNewPendingNode(opts *bind.FilterOpts) (*DosproxyLogRegisteredNewPendingNodeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogRegisteredNewPendingNode")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogRegisteredNewPendingNodeIterator{contract: _Dosproxy.contract, event: "LogRegisteredNewPendingNode", logs: logs, sub: sub}, nil
}

// WatchLogRegisteredNewPendingNode is a free log subscription operation binding the contract event 0x707a6d64786780aac9cd0c5813ea04241eb135ddd2280c06eea6719afb913151.
//
// Solidity: event LogRegisteredNewPendingNode(address node)
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

// ParseLogRegisteredNewPendingNode is a log parse operation binding the contract event 0x707a6d64786780aac9cd0c5813ea04241eb135ddd2280c06eea6719afb913151.
//
// Solidity: event LogRegisteredNewPendingNode(address node)
func (_Dosproxy *DosproxyFilterer) ParseLogRegisteredNewPendingNode(log types.Log) (*DosproxyLogRegisteredNewPendingNode, error) {
	event := new(DosproxyLogRegisteredNewPendingNode)
	if err := _Dosproxy.contract.UnpackLog(event, "LogRegisteredNewPendingNode", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogRequestFromNonExistentUC()
func (_Dosproxy *DosproxyFilterer) FilterLogRequestFromNonExistentUC(opts *bind.FilterOpts) (*DosproxyLogRequestFromNonExistentUCIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogRequestFromNonExistentUC")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogRequestFromNonExistentUCIterator{contract: _Dosproxy.contract, event: "LogRequestFromNonExistentUC", logs: logs, sub: sub}, nil
}

// WatchLogRequestFromNonExistentUC is a free log subscription operation binding the contract event 0x40d87958cd48e8b698a94f35390a9020a9127528227647da089cc6bfd7931f42.
//
// Solidity: event LogRequestFromNonExistentUC()
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

// ParseLogRequestFromNonExistentUC is a log parse operation binding the contract event 0x40d87958cd48e8b698a94f35390a9020a9127528227647da089cc6bfd7931f42.
//
// Solidity: event LogRequestFromNonExistentUC()
func (_Dosproxy *DosproxyFilterer) ParseLogRequestFromNonExistentUC(log types.Log) (*DosproxyLogRequestFromNonExistentUC, error) {
	event := new(DosproxyLogRequestFromNonExistentUC)
	if err := _Dosproxy.contract.UnpackLog(event, "LogRequestFromNonExistentUC", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogRequestUserRandom(uint256 requestId, uint256 lastSystemRandomness, uint256 userSeed, uint256 dispatchedGroupId)
func (_Dosproxy *DosproxyFilterer) FilterLogRequestUserRandom(opts *bind.FilterOpts) (*DosproxyLogRequestUserRandomIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogRequestUserRandom")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogRequestUserRandomIterator{contract: _Dosproxy.contract, event: "LogRequestUserRandom", logs: logs, sub: sub}, nil
}

// WatchLogRequestUserRandom is a free log subscription operation binding the contract event 0xd587179d80d44e74955fa5d651db2f31b5470fcee8f9736f07ae3b24456a2cf0.
//
// Solidity: event LogRequestUserRandom(uint256 requestId, uint256 lastSystemRandomness, uint256 userSeed, uint256 dispatchedGroupId)
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

// ParseLogRequestUserRandom is a log parse operation binding the contract event 0xd587179d80d44e74955fa5d651db2f31b5470fcee8f9736f07ae3b24456a2cf0.
//
// Solidity: event LogRequestUserRandom(uint256 requestId, uint256 lastSystemRandomness, uint256 userSeed, uint256 dispatchedGroupId)
func (_Dosproxy *DosproxyFilterer) ParseLogRequestUserRandom(log types.Log) (*DosproxyLogRequestUserRandom, error) {
	event := new(DosproxyLogRequestUserRandom)
	if err := _Dosproxy.contract.UnpackLog(event, "LogRequestUserRandom", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogUnRegisteredNewPendingNode(address node, uint8 unregisterFrom)
func (_Dosproxy *DosproxyFilterer) FilterLogUnRegisteredNewPendingNode(opts *bind.FilterOpts) (*DosproxyLogUnRegisteredNewPendingNodeIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogUnRegisteredNewPendingNode")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogUnRegisteredNewPendingNodeIterator{contract: _Dosproxy.contract, event: "LogUnRegisteredNewPendingNode", logs: logs, sub: sub}, nil
}

// WatchLogUnRegisteredNewPendingNode is a free log subscription operation binding the contract event 0xaa40dce54d678a8a16fc3cf106c1ddf0b34b66a43c7a365af3268c63bb95fead.
//
// Solidity: event LogUnRegisteredNewPendingNode(address node, uint8 unregisterFrom)
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

// ParseLogUnRegisteredNewPendingNode is a log parse operation binding the contract event 0xaa40dce54d678a8a16fc3cf106c1ddf0b34b66a43c7a365af3268c63bb95fead.
//
// Solidity: event LogUnRegisteredNewPendingNode(address node, uint8 unregisterFrom)
func (_Dosproxy *DosproxyFilterer) ParseLogUnRegisteredNewPendingNode(log types.Log) (*DosproxyLogUnRegisteredNewPendingNode, error) {
	event := new(DosproxyLogUnRegisteredNewPendingNode)
	if err := _Dosproxy.contract.UnpackLog(event, "LogUnRegisteredNewPendingNode", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogUpdateRandom(uint256 lastRandomness, uint256 dispatchedGroupId)
func (_Dosproxy *DosproxyFilterer) FilterLogUpdateRandom(opts *bind.FilterOpts) (*DosproxyLogUpdateRandomIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogUpdateRandom")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogUpdateRandomIterator{contract: _Dosproxy.contract, event: "LogUpdateRandom", logs: logs, sub: sub}, nil
}

// WatchLogUpdateRandom is a free log subscription operation binding the contract event 0xfaa99731d2c5abc7ee76b2e31b6b7e293a30e1e2274f59396a7a59cabd5eadf3.
//
// Solidity: event LogUpdateRandom(uint256 lastRandomness, uint256 dispatchedGroupId)
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

// ParseLogUpdateRandom is a log parse operation binding the contract event 0xfaa99731d2c5abc7ee76b2e31b6b7e293a30e1e2274f59396a7a59cabd5eadf3.
//
// Solidity: event LogUpdateRandom(uint256 lastRandomness, uint256 dispatchedGroupId)
func (_Dosproxy *DosproxyFilterer) ParseLogUpdateRandom(log types.Log) (*DosproxyLogUpdateRandom, error) {
	event := new(DosproxyLogUpdateRandom)
	if err := _Dosproxy.contract.UnpackLog(event, "LogUpdateRandom", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogUrl(uint256 queryId, uint256 timeout, string dataSource, string selector, uint256 randomness, uint256 dispatchedGroupId)
func (_Dosproxy *DosproxyFilterer) FilterLogUrl(opts *bind.FilterOpts) (*DosproxyLogUrlIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogUrl")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogUrlIterator{contract: _Dosproxy.contract, event: "LogUrl", logs: logs, sub: sub}, nil
}

// WatchLogUrl is a free log subscription operation binding the contract event 0x05e1614af4efb13caeba2369a57a05ee5830f33364f82e2c899fd5710cb56ef3.
//
// Solidity: event LogUrl(uint256 queryId, uint256 timeout, string dataSource, string selector, uint256 randomness, uint256 dispatchedGroupId)
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

// ParseLogUrl is a log parse operation binding the contract event 0x05e1614af4efb13caeba2369a57a05ee5830f33364f82e2c899fd5710cb56ef3.
//
// Solidity: event LogUrl(uint256 queryId, uint256 timeout, string dataSource, string selector, uint256 randomness, uint256 dispatchedGroupId)
func (_Dosproxy *DosproxyFilterer) ParseLogUrl(log types.Log) (*DosproxyLogUrl, error) {
	event := new(DosproxyLogUrl)
	if err := _Dosproxy.contract.UnpackLog(event, "LogUrl", log); err != nil {
		return nil, err
	}
	return event, nil
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
// Solidity: event LogValidationResult(uint8 trafficType, uint256 trafficId, bytes message, uint256[2] signature, uint256[4] pubKey, bool pass)
func (_Dosproxy *DosproxyFilterer) FilterLogValidationResult(opts *bind.FilterOpts) (*DosproxyLogValidationResultIterator, error) {

	logs, sub, err := _Dosproxy.contract.FilterLogs(opts, "LogValidationResult")
	if err != nil {
		return nil, err
	}
	return &DosproxyLogValidationResultIterator{contract: _Dosproxy.contract, event: "LogValidationResult", logs: logs, sub: sub}, nil
}

// WatchLogValidationResult is a free log subscription operation binding the contract event 0xd33c44f7ce166bcd2616c8f4d811261d4d24e1af815f78614683e0c7682c42a5.
//
// Solidity: event LogValidationResult(uint8 trafficType, uint256 trafficId, bytes message, uint256[2] signature, uint256[4] pubKey, bool pass)
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

// ParseLogValidationResult is a log parse operation binding the contract event 0xd33c44f7ce166bcd2616c8f4d811261d4d24e1af815f78614683e0c7682c42a5.
//
// Solidity: event LogValidationResult(uint8 trafficType, uint256 trafficId, bytes message, uint256[2] signature, uint256[4] pubKey, bool pass)
func (_Dosproxy *DosproxyFilterer) ParseLogValidationResult(log types.Log) (*DosproxyLogValidationResult, error) {
	event := new(DosproxyLogValidationResult)
	if err := _Dosproxy.contract.UnpackLog(event, "LogValidationResult", log); err != nil {
		return nil, err
	}
	return event, nil
}
