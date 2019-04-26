// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commitreveal

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

// CommitRevealABI is the input ABI used to generate the binding from.
const CommitRevealABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaigns\",\"outputs\":[{\"name\":\"startBlock\",\"type\":\"uint256\"},{\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"name\":\"revealDuration\",\"type\":\"uint256\"},{\"name\":\"revealThreshold\",\"type\":\"uint256\"},{\"name\":\"commitNum\",\"type\":\"uint256\"},{\"name\":\"revealNum\",\"type\":\"uint256\"},{\"name\":\"generatedRandom\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cid\",\"type\":\"uint256\"},{\"name\":\"_secret\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"},{\"name\":\"_revealThreshold\",\"type\":\"uint256\"}],\"name\":\"startCommitReveal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_cid\",\"type\":\"uint256\"}],\"name\":\"getRandom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cid\",\"type\":\"uint256\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"campaignId\",\"type\":\"uint256\"}],\"name\":\"LogStartCommitReveal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"LogCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"secret\",\"type\":\"uint256\"}],\"name\":\"LogReveal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"random\",\"type\":\"uint256\"}],\"name\":\"LogRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// CommitRevealBin is the compiled bytecode used for deploying new contracts.
const CommitRevealBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317905560018054906100349082810161003a565b506100b5565b81548183558181111561006657600902816009028360005260206000209182019101610066919061006b565b505050565b6100b291905b808211156100ae57600080825560018201819055600282018190556003820181905560048201819055600582018190556006820155600901610071565b5090565b90565b610c36806100c46000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063b917b5a511610071578063b917b5a5146101a1578063cd4b6914146101e2578063d936547e146101ff578063e43252d714610225578063f2f038771461024b578063f2fde38b1461026e576100b4565b8063141961bc146100b9578063715018a61461010e5780638ab1d681146101185780638da5cb5b1461013e5780638f32d59b146101625780639348cef71461017e575b600080fd5b6100d6600480360360208110156100cf57600080fd5b5035610294565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b6101166102e3565b005b6101166004803603602081101561012e57600080fd5b50356001600160a01b031661033c565b61014661036e565b604080516001600160a01b039092168252519081900360200190f35b61016a61037d565b604080519115158252519081900360200190f35b6101166004803603604081101561019457600080fd5b508035906020013561038e565b6101d0600480360360808110156101b757600080fd5b508035906020810135906040810135906060013561054a565b60408051918252519081900360200190f35b6101d0600480360360208110156101f857600080fd5b503561079f565b61016a6004803603602081101561021557600080fd5b50356001600160a01b03166108ad565b6101166004803603602081101561023b57600080fd5b50356001600160a01b03166108c2565b6101166004803603604081101561026157600080fd5b50803590602001356108f7565b6101166004803603602081101561028457600080fd5b50356001600160a01b0316610b25565b600181815481106102a157fe5b90600052602060002090600902016000915090508060000154908060010154908060020154908060030154908060040154908060050154908060060154905087565b6102eb61037d565b6102f457600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b61034461037d565b61034d57600080fd5b6001600160a01b03166000908152600260205260409020805460ff19169055565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b8160006001828154811061039e57fe5b90600052602060002090600902019050816000141580156103c6575060018101548154014310155b80156103de5750600281015460018201548254010143105b6104325760408051600160e51b62461bcd02815260206004820152601360248201527f4e6f7420696e2072657665616c20706861736500000000000000000000000000604482015290519081900360640190fd5b60006001858154811061044157fe5b6000918252602080832033845260076009909302019182019052604090912060028101549192509060ff1615801561049b575060018101546040805160208082018990528251808303820181529183019092528051910120145b6104d957604051600160e51b62461bcd02815260040180806020018281038252602d815260200180610bde602d913960400191505060405180910390fd5b84815560028101805460ff191660019081179091556005830180549091019055600682018054861890556040805187815233602082015280820187905290517f9141bfaedbc77aa7b8d9c989cd81909d95bb1677e556e34cfd45e50e0bea29679181900360600190a1505050505050565b3360009081526002602052604081205460ff166105b15760408051600160e51b62461bcd02815260206004820152601060248201527f4e6f742077686974656c69737465642100000000000000000000000000000000604482015290519081900360640190fd5b844310156106095760408051600160e51b62461bcd02815260206004820152601d60248201527f636f6d6d69742d72657665616c206e6f74207374617274656420796574000000604482015290519081900360640190fd5b6040805160e08101825286815260208082018781528284018781526060840187815260006080860181815260a0870182815260c088018381526001805480820182559481905298517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf660099095029485015595517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf784015593517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf883015591517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf982015590517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfa82015590517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfb82015590517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfc909101559054825190815291517ffcff4e68c3586bff594e01bbb2197f3c799b56b9c6bfb1d0fde460616e4febe29281900390910190a150600154949350505050565b6000816000600182815481106107b157fe5b90600052602060002090600902019050816000141580156107df575060028101546001820154825401014310155b6108335760408051600160e51b62461bcd02815260206004820152601e60248201527f436f6d6d69742052657665616c206e6f742066696e6973686564207965740000604482015290519081900360640190fd5b60006001858154811061084257fe5b90600052602060002090600902019050806003015481600501541061086c576006015492506108a6565b604051600160e51b62461bcd02815260040180806020018281038252602d815260200180610bb1602d913960400191505060405180910390fd5b5050919050565b60026020526000908152604090205460ff1681565b6108ca61037d565b6108d357600080fd5b6001600160a01b03166000908152600260205260409020805460ff19166001179055565b818160006001838154811061090857fe5b906000526020600020906009020190508260001415801561092a575080544310155b801561093c5750600181015481540143105b6109905760408051600160e51b62461bcd02815260206004820152601360248201527f4e6f7420696e20636f6d6d697420706861736500000000000000000000000000604482015290519081900360640190fd5b816109e55760408051600160e51b62461bcd02815260206004820152601060248201527f456d70747920636f6d6d69746d656e7400000000000000000000000000000000604482015290519081900360640190fd5b600082815260088201602052604090205460ff1615610a4e5760408051600160e51b62461bcd02815260206004820152601560248201527f4475706c69636174656420636f6d6d69746d656e740000000000000000000000604482015290519081900360640190fd5b600060018681548110610a5d57fe5b60009182526020808320888452600860099093020191820181526040808420805460ff1990811660019081179092558251606080820185528782528186018d815282860189815233808b5260078a01895299879020935184559051838601555160029092018054921515929093169190911790915560048501805490920190915581518b815292830194909452818101899052519193507f918c00c65dd2a8dee4c6985d1d67f04aa8cd2c93e8d427d398a90444c7f7c75e92908290030190a1505050505050565b610b2d61037d565b610b3657600080fd5b610b3f81610b42565b50565b6001600160a01b038116610b5557600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe436f6d6d69742052657665616c2052616e646f6d204e756d6265722047656e65726174696f6e204661696c656452657665616c65642073656372657420646f65736e2774206d61746368207769746820636f6d6d69746d656e74a165627a7a723058200c0aee5cf134f71fba99ce431c60dd6eda7cab12fc819051c6156c92b31579320029`

// DeployCommitReveal deploys a new Ethereum contract, binding an instance of CommitReveal to it.
func DeployCommitReveal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CommitReveal, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitRevealABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CommitRevealBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitReveal{CommitRevealCaller: CommitRevealCaller{contract: contract}, CommitRevealTransactor: CommitRevealTransactor{contract: contract}, CommitRevealFilterer: CommitRevealFilterer{contract: contract}}, nil
}

// CommitReveal is an auto generated Go binding around an Ethereum contract.
type CommitReveal struct {
	CommitRevealCaller     // Read-only binding to the contract
	CommitRevealTransactor // Write-only binding to the contract
	CommitRevealFilterer   // Log filterer for contract events
}

// CommitRevealCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitRevealCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitRevealTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitRevealTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitRevealFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitRevealFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitRevealSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitRevealSession struct {
	Contract     *CommitReveal     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitRevealCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitRevealCallerSession struct {
	Contract *CommitRevealCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CommitRevealTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitRevealTransactorSession struct {
	Contract     *CommitRevealTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CommitRevealRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitRevealRaw struct {
	Contract *CommitReveal // Generic contract binding to access the raw methods on
}

// CommitRevealCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitRevealCallerRaw struct {
	Contract *CommitRevealCaller // Generic read-only contract binding to access the raw methods on
}

// CommitRevealTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitRevealTransactorRaw struct {
	Contract *CommitRevealTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitReveal creates a new instance of CommitReveal, bound to a specific deployed contract.
func NewCommitReveal(address common.Address, backend bind.ContractBackend) (*CommitReveal, error) {
	contract, err := bindCommitReveal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitReveal{CommitRevealCaller: CommitRevealCaller{contract: contract}, CommitRevealTransactor: CommitRevealTransactor{contract: contract}, CommitRevealFilterer: CommitRevealFilterer{contract: contract}}, nil
}

// NewCommitRevealCaller creates a new read-only instance of CommitReveal, bound to a specific deployed contract.
func NewCommitRevealCaller(address common.Address, caller bind.ContractCaller) (*CommitRevealCaller, error) {
	contract, err := bindCommitReveal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitRevealCaller{contract: contract}, nil
}

// NewCommitRevealTransactor creates a new write-only instance of CommitReveal, bound to a specific deployed contract.
func NewCommitRevealTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitRevealTransactor, error) {
	contract, err := bindCommitReveal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitRevealTransactor{contract: contract}, nil
}

// NewCommitRevealFilterer creates a new log filterer instance of CommitReveal, bound to a specific deployed contract.
func NewCommitRevealFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitRevealFilterer, error) {
	contract, err := bindCommitReveal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitRevealFilterer{contract: contract}, nil
}

// bindCommitReveal binds a generic wrapper to an already deployed contract.
func bindCommitReveal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitRevealABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitReveal *CommitRevealRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CommitReveal.Contract.CommitRevealCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitReveal *CommitRevealRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitReveal.Contract.CommitRevealTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitReveal *CommitRevealRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitReveal.Contract.CommitRevealTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitReveal *CommitRevealCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CommitReveal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitReveal *CommitRevealTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitReveal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitReveal *CommitRevealTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitReveal.Contract.contract.Transact(opts, method, params...)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns( uint256) constant returns(startBlock uint256, commitDuration uint256, revealDuration uint256, revealThreshold uint256, commitNum uint256, revealNum uint256, generatedRandom uint256)
func (_CommitReveal *CommitRevealCaller) Campaigns(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	GeneratedRandom *big.Int
}, error) {
	ret := new(struct {
		StartBlock      *big.Int
		CommitDuration  *big.Int
		RevealDuration  *big.Int
		RevealThreshold *big.Int
		CommitNum       *big.Int
		RevealNum       *big.Int
		GeneratedRandom *big.Int
	})
	out := ret
	err := _CommitReveal.contract.Call(opts, out, "campaigns", arg0)
	return *ret, err
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns( uint256) constant returns(startBlock uint256, commitDuration uint256, revealDuration uint256, revealThreshold uint256, commitNum uint256, revealNum uint256, generatedRandom uint256)
func (_CommitReveal *CommitRevealSession) Campaigns(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	GeneratedRandom *big.Int
}, error) {
	return _CommitReveal.Contract.Campaigns(&_CommitReveal.CallOpts, arg0)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns( uint256) constant returns(startBlock uint256, commitDuration uint256, revealDuration uint256, revealThreshold uint256, commitNum uint256, revealNum uint256, generatedRandom uint256)
func (_CommitReveal *CommitRevealCallerSession) Campaigns(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	GeneratedRandom *big.Int
}, error) {
	return _CommitReveal.Contract.Campaigns(&_CommitReveal.CallOpts, arg0)
}

// GetRandom is a free data retrieval call binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(_cid uint256) constant returns(uint256)
func (_CommitReveal *CommitRevealCaller) GetRandom(opts *bind.CallOpts, _cid *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CommitReveal.contract.Call(opts, out, "getRandom", _cid)
	return *ret0, err
}

// GetRandom is a free data retrieval call binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(_cid uint256) constant returns(uint256)
func (_CommitReveal *CommitRevealSession) GetRandom(_cid *big.Int) (*big.Int, error) {
	return _CommitReveal.Contract.GetRandom(&_CommitReveal.CallOpts, _cid)
}

// GetRandom is a free data retrieval call binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(_cid uint256) constant returns(uint256)
func (_CommitReveal *CommitRevealCallerSession) GetRandom(_cid *big.Int) (*big.Int, error) {
	return _CommitReveal.Contract.GetRandom(&_CommitReveal.CallOpts, _cid)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_CommitReveal *CommitRevealCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CommitReveal.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_CommitReveal *CommitRevealSession) IsOwner() (bool, error) {
	return _CommitReveal.Contract.IsOwner(&_CommitReveal.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_CommitReveal *CommitRevealCallerSession) IsOwner() (bool, error) {
	return _CommitReveal.Contract.IsOwner(&_CommitReveal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CommitReveal *CommitRevealCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CommitReveal.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CommitReveal *CommitRevealSession) Owner() (common.Address, error) {
	return _CommitReveal.Contract.Owner(&_CommitReveal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CommitReveal *CommitRevealCallerSession) Owner() (common.Address, error) {
	return _CommitReveal.Contract.Owner(&_CommitReveal.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted( address) constant returns(bool)
func (_CommitReveal *CommitRevealCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CommitReveal.contract.Call(opts, out, "whitelisted", arg0)
	return *ret0, err
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted( address) constant returns(bool)
func (_CommitReveal *CommitRevealSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _CommitReveal.Contract.Whitelisted(&_CommitReveal.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted( address) constant returns(bool)
func (_CommitReveal *CommitRevealCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _CommitReveal.Contract.Whitelisted(&_CommitReveal.CallOpts, arg0)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_addr address) returns()
func (_CommitReveal *CommitRevealTransactor) AddToWhitelist(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CommitReveal.contract.Transact(opts, "addToWhitelist", _addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_addr address) returns()
func (_CommitReveal *CommitRevealSession) AddToWhitelist(_addr common.Address) (*types.Transaction, error) {
	return _CommitReveal.Contract.AddToWhitelist(&_CommitReveal.TransactOpts, _addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(_addr address) returns()
func (_CommitReveal *CommitRevealTransactorSession) AddToWhitelist(_addr common.Address) (*types.Transaction, error) {
	return _CommitReveal.Contract.AddToWhitelist(&_CommitReveal.TransactOpts, _addr)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(_cid uint256, _secretHash bytes32) returns()
func (_CommitReveal *CommitRevealTransactor) Commit(opts *bind.TransactOpts, _cid *big.Int, _secretHash [32]byte) (*types.Transaction, error) {
	return _CommitReveal.contract.Transact(opts, "commit", _cid, _secretHash)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(_cid uint256, _secretHash bytes32) returns()
func (_CommitReveal *CommitRevealSession) Commit(_cid *big.Int, _secretHash [32]byte) (*types.Transaction, error) {
	return _CommitReveal.Contract.Commit(&_CommitReveal.TransactOpts, _cid, _secretHash)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(_cid uint256, _secretHash bytes32) returns()
func (_CommitReveal *CommitRevealTransactorSession) Commit(_cid *big.Int, _secretHash [32]byte) (*types.Transaction, error) {
	return _CommitReveal.Contract.Commit(&_CommitReveal.TransactOpts, _cid, _secretHash)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_addr address) returns()
func (_CommitReveal *CommitRevealTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CommitReveal.contract.Transact(opts, "removeFromWhitelist", _addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_addr address) returns()
func (_CommitReveal *CommitRevealSession) RemoveFromWhitelist(_addr common.Address) (*types.Transaction, error) {
	return _CommitReveal.Contract.RemoveFromWhitelist(&_CommitReveal.TransactOpts, _addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(_addr address) returns()
func (_CommitReveal *CommitRevealTransactorSession) RemoveFromWhitelist(_addr common.Address) (*types.Transaction, error) {
	return _CommitReveal.Contract.RemoveFromWhitelist(&_CommitReveal.TransactOpts, _addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CommitReveal *CommitRevealTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitReveal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CommitReveal *CommitRevealSession) RenounceOwnership() (*types.Transaction, error) {
	return _CommitReveal.Contract.RenounceOwnership(&_CommitReveal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CommitReveal *CommitRevealTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CommitReveal.Contract.RenounceOwnership(&_CommitReveal.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(_cid uint256, _secret uint256) returns()
func (_CommitReveal *CommitRevealTransactor) Reveal(opts *bind.TransactOpts, _cid *big.Int, _secret *big.Int) (*types.Transaction, error) {
	return _CommitReveal.contract.Transact(opts, "reveal", _cid, _secret)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(_cid uint256, _secret uint256) returns()
func (_CommitReveal *CommitRevealSession) Reveal(_cid *big.Int, _secret *big.Int) (*types.Transaction, error) {
	return _CommitReveal.Contract.Reveal(&_CommitReveal.TransactOpts, _cid, _secret)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(_cid uint256, _secret uint256) returns()
func (_CommitReveal *CommitRevealTransactorSession) Reveal(_cid *big.Int, _secret *big.Int) (*types.Transaction, error) {
	return _CommitReveal.Contract.Reveal(&_CommitReveal.TransactOpts, _cid, _secret)
}

// StartCommitReveal is a paid mutator transaction binding the contract method 0xb917b5a5.
//
// Solidity: function startCommitReveal(_startBlock uint256, _commitDuration uint256, _revealDuration uint256, _revealThreshold uint256) returns(uint256)
func (_CommitReveal *CommitRevealTransactor) StartCommitReveal(opts *bind.TransactOpts, _startBlock *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _revealThreshold *big.Int) (*types.Transaction, error) {
	return _CommitReveal.contract.Transact(opts, "startCommitReveal", _startBlock, _commitDuration, _revealDuration, _revealThreshold)
}

// StartCommitReveal is a paid mutator transaction binding the contract method 0xb917b5a5.
//
// Solidity: function startCommitReveal(_startBlock uint256, _commitDuration uint256, _revealDuration uint256, _revealThreshold uint256) returns(uint256)
func (_CommitReveal *CommitRevealSession) StartCommitReveal(_startBlock *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _revealThreshold *big.Int) (*types.Transaction, error) {
	return _CommitReveal.Contract.StartCommitReveal(&_CommitReveal.TransactOpts, _startBlock, _commitDuration, _revealDuration, _revealThreshold)
}

// StartCommitReveal is a paid mutator transaction binding the contract method 0xb917b5a5.
//
// Solidity: function startCommitReveal(_startBlock uint256, _commitDuration uint256, _revealDuration uint256, _revealThreshold uint256) returns(uint256)
func (_CommitReveal *CommitRevealTransactorSession) StartCommitReveal(_startBlock *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _revealThreshold *big.Int) (*types.Transaction, error) {
	return _CommitReveal.Contract.StartCommitReveal(&_CommitReveal.TransactOpts, _startBlock, _commitDuration, _revealDuration, _revealThreshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CommitReveal *CommitRevealTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CommitReveal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CommitReveal *CommitRevealSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CommitReveal.Contract.TransferOwnership(&_CommitReveal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CommitReveal *CommitRevealTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CommitReveal.Contract.TransferOwnership(&_CommitReveal.TransactOpts, newOwner)
}

// CommitRevealLogCommitIterator is returned from FilterLogCommit and is used to iterate over the raw logs and unpacked data for LogCommit events raised by the CommitReveal contract.
type CommitRevealLogCommitIterator struct {
	Event *CommitRevealLogCommit // Event containing the contract specifics and raw log

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
func (it *CommitRevealLogCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRevealLogCommit)
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
		it.Event = new(CommitRevealLogCommit)
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
func (it *CommitRevealLogCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitRevealLogCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitRevealLogCommit represents a LogCommit event raised by the CommitReveal contract.
type CommitRevealLogCommit struct {
	Cid        *big.Int
	From       common.Address
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogCommit is a free log retrieval operation binding the contract event 0x918c00c65dd2a8dee4c6985d1d67f04aa8cd2c93e8d427d398a90444c7f7c75e.
//
// Solidity: e LogCommit(cid uint256, from address, commitment bytes32)
func (_CommitReveal *CommitRevealFilterer) FilterLogCommit(opts *bind.FilterOpts) (*CommitRevealLogCommitIterator, error) {

	logs, sub, err := _CommitReveal.contract.FilterLogs(opts, "LogCommit")
	if err != nil {
		return nil, err
	}
	return &CommitRevealLogCommitIterator{contract: _CommitReveal.contract, event: "LogCommit", logs: logs, sub: sub}, nil
}

// WatchLogCommit is a free log subscription operation binding the contract event 0x918c00c65dd2a8dee4c6985d1d67f04aa8cd2c93e8d427d398a90444c7f7c75e.
//
// Solidity: e LogCommit(cid uint256, from address, commitment bytes32)
func (_CommitReveal *CommitRevealFilterer) WatchLogCommit(opts *bind.WatchOpts, sink chan<- *CommitRevealLogCommit) (event.Subscription, error) {

	logs, sub, err := _CommitReveal.contract.WatchLogs(opts, "LogCommit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitRevealLogCommit)
				if err := _CommitReveal.contract.UnpackLog(event, "LogCommit", log); err != nil {
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

// CommitRevealLogRandomIterator is returned from FilterLogRandom and is used to iterate over the raw logs and unpacked data for LogRandom events raised by the CommitReveal contract.
type CommitRevealLogRandomIterator struct {
	Event *CommitRevealLogRandom // Event containing the contract specifics and raw log

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
func (it *CommitRevealLogRandomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRevealLogRandom)
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
		it.Event = new(CommitRevealLogRandom)
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
func (it *CommitRevealLogRandomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitRevealLogRandomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitRevealLogRandom represents a LogRandom event raised by the CommitReveal contract.
type CommitRevealLogRandom struct {
	Cid    *big.Int
	Random *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogRandom is a free log retrieval operation binding the contract event 0xa34f42a90fadfe357ee14419d618438a057569fbb63bab0c5fbcca0fc2b11e8d.
//
// Solidity: e LogRandom(cid uint256, random uint256)
func (_CommitReveal *CommitRevealFilterer) FilterLogRandom(opts *bind.FilterOpts) (*CommitRevealLogRandomIterator, error) {

	logs, sub, err := _CommitReveal.contract.FilterLogs(opts, "LogRandom")
	if err != nil {
		return nil, err
	}
	return &CommitRevealLogRandomIterator{contract: _CommitReveal.contract, event: "LogRandom", logs: logs, sub: sub}, nil
}

// WatchLogRandom is a free log subscription operation binding the contract event 0xa34f42a90fadfe357ee14419d618438a057569fbb63bab0c5fbcca0fc2b11e8d.
//
// Solidity: e LogRandom(cid uint256, random uint256)
func (_CommitReveal *CommitRevealFilterer) WatchLogRandom(opts *bind.WatchOpts, sink chan<- *CommitRevealLogRandom) (event.Subscription, error) {

	logs, sub, err := _CommitReveal.contract.WatchLogs(opts, "LogRandom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitRevealLogRandom)
				if err := _CommitReveal.contract.UnpackLog(event, "LogRandom", log); err != nil {
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

// CommitRevealLogRevealIterator is returned from FilterLogReveal and is used to iterate over the raw logs and unpacked data for LogReveal events raised by the CommitReveal contract.
type CommitRevealLogRevealIterator struct {
	Event *CommitRevealLogReveal // Event containing the contract specifics and raw log

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
func (it *CommitRevealLogRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRevealLogReveal)
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
		it.Event = new(CommitRevealLogReveal)
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
func (it *CommitRevealLogRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitRevealLogRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitRevealLogReveal represents a LogReveal event raised by the CommitReveal contract.
type CommitRevealLogReveal struct {
	Cid    *big.Int
	From   common.Address
	Secret *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogReveal is a free log retrieval operation binding the contract event 0x9141bfaedbc77aa7b8d9c989cd81909d95bb1677e556e34cfd45e50e0bea2967.
//
// Solidity: e LogReveal(cid uint256, from address, secret uint256)
func (_CommitReveal *CommitRevealFilterer) FilterLogReveal(opts *bind.FilterOpts) (*CommitRevealLogRevealIterator, error) {

	logs, sub, err := _CommitReveal.contract.FilterLogs(opts, "LogReveal")
	if err != nil {
		return nil, err
	}
	return &CommitRevealLogRevealIterator{contract: _CommitReveal.contract, event: "LogReveal", logs: logs, sub: sub}, nil
}

// WatchLogReveal is a free log subscription operation binding the contract event 0x9141bfaedbc77aa7b8d9c989cd81909d95bb1677e556e34cfd45e50e0bea2967.
//
// Solidity: e LogReveal(cid uint256, from address, secret uint256)
func (_CommitReveal *CommitRevealFilterer) WatchLogReveal(opts *bind.WatchOpts, sink chan<- *CommitRevealLogReveal) (event.Subscription, error) {

	logs, sub, err := _CommitReveal.contract.WatchLogs(opts, "LogReveal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitRevealLogReveal)
				if err := _CommitReveal.contract.UnpackLog(event, "LogReveal", log); err != nil {
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

// CommitRevealLogStartCommitRevealIterator is returned from FilterLogStartCommitReveal and is used to iterate over the raw logs and unpacked data for LogStartCommitReveal events raised by the CommitReveal contract.
type CommitRevealLogStartCommitRevealIterator struct {
	Event *CommitRevealLogStartCommitReveal // Event containing the contract specifics and raw log

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
func (it *CommitRevealLogStartCommitRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRevealLogStartCommitReveal)
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
		it.Event = new(CommitRevealLogStartCommitReveal)
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
func (it *CommitRevealLogStartCommitRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitRevealLogStartCommitRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitRevealLogStartCommitReveal represents a LogStartCommitReveal event raised by the CommitReveal contract.
type CommitRevealLogStartCommitReveal struct {
	CampaignId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogStartCommitReveal is a free log retrieval operation binding the contract event 0xfcff4e68c3586bff594e01bbb2197f3c799b56b9c6bfb1d0fde460616e4febe2.
//
// Solidity: e LogStartCommitReveal(campaignId uint256)
func (_CommitReveal *CommitRevealFilterer) FilterLogStartCommitReveal(opts *bind.FilterOpts) (*CommitRevealLogStartCommitRevealIterator, error) {

	logs, sub, err := _CommitReveal.contract.FilterLogs(opts, "LogStartCommitReveal")
	if err != nil {
		return nil, err
	}
	return &CommitRevealLogStartCommitRevealIterator{contract: _CommitReveal.contract, event: "LogStartCommitReveal", logs: logs, sub: sub}, nil
}

// WatchLogStartCommitReveal is a free log subscription operation binding the contract event 0xfcff4e68c3586bff594e01bbb2197f3c799b56b9c6bfb1d0fde460616e4febe2.
//
// Solidity: e LogStartCommitReveal(campaignId uint256)
func (_CommitReveal *CommitRevealFilterer) WatchLogStartCommitReveal(opts *bind.WatchOpts, sink chan<- *CommitRevealLogStartCommitReveal) (event.Subscription, error) {

	logs, sub, err := _CommitReveal.contract.WatchLogs(opts, "LogStartCommitReveal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitRevealLogStartCommitReveal)
				if err := _CommitReveal.contract.UnpackLog(event, "LogStartCommitReveal", log); err != nil {
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

// CommitRevealOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the CommitReveal contract.
type CommitRevealOwnershipRenouncedIterator struct {
	Event *CommitRevealOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *CommitRevealOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRevealOwnershipRenounced)
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
		it.Event = new(CommitRevealOwnershipRenounced)
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
func (it *CommitRevealOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitRevealOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitRevealOwnershipRenounced represents a OwnershipRenounced event raised by the CommitReveal contract.
type CommitRevealOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_CommitReveal *CommitRevealFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*CommitRevealOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _CommitReveal.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CommitRevealOwnershipRenouncedIterator{contract: _CommitReveal.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_CommitReveal *CommitRevealFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *CommitRevealOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _CommitReveal.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitRevealOwnershipRenounced)
				if err := _CommitReveal.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// CommitRevealOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CommitReveal contract.
type CommitRevealOwnershipTransferredIterator struct {
	Event *CommitRevealOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CommitRevealOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitRevealOwnershipTransferred)
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
		it.Event = new(CommitRevealOwnershipTransferred)
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
func (it *CommitRevealOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitRevealOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitRevealOwnershipTransferred represents a OwnershipTransferred event raised by the CommitReveal contract.
type CommitRevealOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CommitReveal *CommitRevealFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CommitRevealOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CommitReveal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CommitRevealOwnershipTransferredIterator{contract: _CommitReveal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CommitReveal *CommitRevealFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitRevealOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CommitReveal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitRevealOwnershipTransferred)
				if err := _CommitReveal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556101f1806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063715018a6146100515780638da5cb5b1461005b5780638f32d59b1461007f578063f2fde38b1461009b575b600080fd5b6100596100c1565b005b61006361011a565b604080516001600160a01b039092168252519081900360200190f35b610087610129565b604080519115158252519081900360200190f35b610059600480360360208110156100b157600080fd5b50356001600160a01b031661013a565b6100c9610129565b6100d257600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b610142610129565b61014b57600080fd5b61015481610157565b50565b6001600160a01b03811661016a57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea165627a7a723058205f55ec58230293187624ace2c0e0b84f7af2af78d0f99af9f099482d958e6d730029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Ownable contract.
type OwnableOwnershipRenouncedIterator struct {
	Event *OwnableOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipRenounced)
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
		it.Event = new(OwnableOwnershipRenounced)
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
func (it *OwnableOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipRenounced represents a OwnershipRenounced event raised by the Ownable contract.
type OwnableOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*OwnableOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipRenouncedIterator{contract: _Ownable.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipRenounced)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
