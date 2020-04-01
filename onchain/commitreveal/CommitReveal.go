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

// CommitrevealABI is the input ABI used to generate the binding from.
const CommitrevealABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"LogCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"random\",\"type\":\"uint256\"}],\"name\":\"LogRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commitNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealThreshold\",\"type\":\"uint256\"}],\"name\":\"LogRandomFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secret\",\"type\":\"uint256\"}],\"name\":\"LogReveal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealThreshold\",\"type\":\"uint256\"}],\"name\":\"LogStartCommitReveal\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressBridge\",\"outputs\":[{\"internalType\":\"contractDOSAddressBridgeInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaigns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commitNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"generatedRandom\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_secretHash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cid\",\"type\":\"uint256\"}],\"name\":\"getRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secret\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_revealDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_revealThreshold\",\"type\":\"uint256\"}],\"name\":\"startCommitReveal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CommitrevealBin is the compiled bytecode used for deploying new contracts.
const CommitrevealBin = `608060405234801561001057600080fd5b50604051610b10380380610b108339818101604052602081101561003357600080fd5b50516000805490610047906001830161007a565b50600280546001600160a01b03199081166001600160a01b039384161791829055600180549290931691161790556100f5565b8154818355818111156100a6576009028160090283600052602060002091820191016100a691906100ab565b505050565b6100f291905b808211156100ee576000808255600182018190556002820181905560038201819055600482018190556005820181905560068201556009016100b1565b5090565b90565b610a0c806101046000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80639348cef71161005b5780639348cef714610103578063b917b5a514610128578063cd4b691414610169578063f2f03877146101865761007d565b8063141961bc1461008257806376cffa53146100d757806391874ef7146100fb575b600080fd5b61009f6004803603602081101561009857600080fd5b50356101a9565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b6100df6101f8565b604080516001600160a01b039092168252519081900360200190f35b6100df610207565b6101266004803603604081101561011957600080fd5b5080359060200135610216565b005b6101576004803603608081101561013e57600080fd5b50803590602081013590604081013590606001356103c0565b60408051918252519081900360200190f35b6101576004803603602081101561017f57600080fd5b503561063c565b6101266004803603604081101561019c57600080fd5b50803590602001356107a6565b600081815481106101b657fe5b90600052602060002090600902016000915090508060000154908060010154908060020154908060030154908060040154908060050154908060060154905087565b6001546001600160a01b031681565b6002546001600160a01b031681565b81600080828154811061022557fe5b906000526020600020906009020190508160001415801561024d575060018101548154014310155b80156102655750600281015460018201548254010143105b6102ac576040805162461bcd60e51b81526020600482015260136024820152724e6f7420696e2072657665616c20706861736560681b604482015290519081900360640190fd5b60008085815481106102ba57fe5b6000918252602080832033845260076009909302019182019052604090912060028101549192509060ff16158015610314575060018101546040805160208082018990528251808303820181529183019092528051910120145b61034f5760405162461bcd60e51b815260040180806020018281038252602d8152602001806109ab602d913960400191505060405180910390fd5b84815560028101805460ff191660019081179091556005830180549091019055600682018054861890556040805187815233602082015280820187905290517f9141bfaedbc77aa7b8d9c989cd81909d95bb1677e556e34cfd45e50e0bea29679181900360600190a1505050505050565b600154604080516321d39ecd60e11b815290516000926001600160a01b0316916343a73d9a916004808301926020929190829003018186803b15801561040557600080fd5b505afa158015610419573d6000803e3d6000fd5b505050506040513d602081101561042f57600080fd5b50516001600160a01b0316331461048d576040805162461bcd60e51b815260206004820152601760248201527f4e6f742066726f6d2070726f787920636f6e7472616374000000000000000000604482015290519081900360640190fd5b600080546040805160e081018252888152602080820189815282840189815260608085018a815260808087018a815260a08089018c815260c08a018d815260018d018e559c8052985160098c027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56381019190915596517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56488015594517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56587015591517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56686015590517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56785015594517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56884015596517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5699092019190915583518581529182018b90528184018a9052948101889052908101869052905191927fbbfccb30e8cf1b5d88802741ceb4d63cf854fa8931eeeaec6b700f31f429dc0992918290030190a195945050505050565b600081600080828154811061064d57fe5b906000526020600020906009020190508160001415801561067b575060028101546001820154825401014310155b6106cc576040805162461bcd60e51b815260206004820152601e60248201527f436f6d6d69742052657665616c206e6f742066696e6973686564207965740000604482015290519081900360640190fd5b60008085815481106106da57fe5b90600052602060002090600902019050806003015481600501541061074457600681015460408051878152602081019290925280517fa34f42a90fadfe357ee14419d618438a057569fbb63bab0c5fbcca0fc2b11e8d9281900390910190a160060154925061079f565b600481015460058201546003830154604080518981526020810194909452838101929092526060830152517fe888e7582d0505bce81eef694dfa216179eaaa3c1bd96b7894de8b4370d8543e9181900360800190a160009350505b5050919050565b818160008083815481106107b657fe5b90600052602060002090600902019050826000141580156107d8575080544310155b80156107ea5750600181015481540143105b610831576040805162461bcd60e51b81526020600482015260136024820152724e6f7420696e20636f6d6d697420706861736560681b604482015290519081900360640190fd5b81610876576040805162461bcd60e51b815260206004820152601060248201526f115b5c1d1e4818dbdb5b5a5d1b595b9d60821b604482015290519081900360640190fd5b600082815260088201602052604090205460ff16156108d4576040805162461bcd60e51b8152602060048201526015602482015274111d5c1b1a58d85d19590818dbdb5b5a5d1b595b9d605a1b604482015290519081900360640190fd5b60008086815481106108e257fe5b60009182526020808320888452600860099093020191820181526040808420805460ff1990811660019081179092558251606080820185528782528186018d815282860189815233808b5260078a01895299879020935184559051838601555160029092018054921515929093169190911790915560048501805490920190915581518b815292830194909452818101899052519193507f918c00c65dd2a8dee4c6985d1d67f04aa8cd2c93e8d427d398a90444c7f7c75e92908290030190a150505050505056fe52657665616c65642073656372657420646f65736e2774206d61746368207769746820636f6d6d69746d656e74a265627a7a72315820ae83711c5b44dc09cde02c2f88ad0ff9f4842988f271bad1bcba506353088e8f64736f6c63430005100032`

// DeployCommitreveal deploys a new Ethereum contract, binding an instance of Commitreveal to it.
func DeployCommitreveal(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddr common.Address) (common.Address, *types.Transaction, *Commitreveal, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitrevealABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CommitrevealBin), backend, _bridgeAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Commitreveal{CommitrevealCaller: CommitrevealCaller{contract: contract}, CommitrevealTransactor: CommitrevealTransactor{contract: contract}, CommitrevealFilterer: CommitrevealFilterer{contract: contract}}, nil
}

// Commitreveal is an auto generated Go binding around an Ethereum contract.
type Commitreveal struct {
	CommitrevealCaller     // Read-only binding to the contract
	CommitrevealTransactor // Write-only binding to the contract
	CommitrevealFilterer   // Log filterer for contract events
}

// CommitrevealCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitrevealCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitrevealTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitrevealTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitrevealFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitrevealFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitrevealSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitrevealSession struct {
	Contract     *Commitreveal     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitrevealCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitrevealCallerSession struct {
	Contract *CommitrevealCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CommitrevealTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitrevealTransactorSession struct {
	Contract     *CommitrevealTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CommitrevealRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitrevealRaw struct {
	Contract *Commitreveal // Generic contract binding to access the raw methods on
}

// CommitrevealCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitrevealCallerRaw struct {
	Contract *CommitrevealCaller // Generic read-only contract binding to access the raw methods on
}

// CommitrevealTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitrevealTransactorRaw struct {
	Contract *CommitrevealTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitreveal creates a new instance of Commitreveal, bound to a specific deployed contract.
func NewCommitreveal(address common.Address, backend bind.ContractBackend) (*Commitreveal, error) {
	contract, err := bindCommitreveal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Commitreveal{CommitrevealCaller: CommitrevealCaller{contract: contract}, CommitrevealTransactor: CommitrevealTransactor{contract: contract}, CommitrevealFilterer: CommitrevealFilterer{contract: contract}}, nil
}

// NewCommitrevealCaller creates a new read-only instance of Commitreveal, bound to a specific deployed contract.
func NewCommitrevealCaller(address common.Address, caller bind.ContractCaller) (*CommitrevealCaller, error) {
	contract, err := bindCommitreveal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitrevealCaller{contract: contract}, nil
}

// NewCommitrevealTransactor creates a new write-only instance of Commitreveal, bound to a specific deployed contract.
func NewCommitrevealTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitrevealTransactor, error) {
	contract, err := bindCommitreveal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitrevealTransactor{contract: contract}, nil
}

// NewCommitrevealFilterer creates a new log filterer instance of Commitreveal, bound to a specific deployed contract.
func NewCommitrevealFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitrevealFilterer, error) {
	contract, err := bindCommitreveal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitrevealFilterer{contract: contract}, nil
}

// bindCommitreveal binds a generic wrapper to an already deployed contract.
func bindCommitreveal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitrevealABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Commitreveal *CommitrevealRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Commitreveal.Contract.CommitrevealCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Commitreveal *CommitrevealRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal.Contract.CommitrevealTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Commitreveal *CommitrevealRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Commitreveal.Contract.CommitrevealTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Commitreveal *CommitrevealCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Commitreveal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Commitreveal *CommitrevealTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Commitreveal *CommitrevealTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Commitreveal.Contract.contract.Transact(opts, method, params...)
}

// AddressBridge is a free data retrieval call binding the contract method 0x76cffa53.
//
// Solidity: function addressBridge() constant returns(address)
func (_Commitreveal *CommitrevealCaller) AddressBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Commitreveal.contract.Call(opts, out, "addressBridge")
	return *ret0, err
}

// AddressBridge is a free data retrieval call binding the contract method 0x76cffa53.
//
// Solidity: function addressBridge() constant returns(address)
func (_Commitreveal *CommitrevealSession) AddressBridge() (common.Address, error) {
	return _Commitreveal.Contract.AddressBridge(&_Commitreveal.CallOpts)
}

// AddressBridge is a free data retrieval call binding the contract method 0x76cffa53.
//
// Solidity: function addressBridge() constant returns(address)
func (_Commitreveal *CommitrevealCallerSession) AddressBridge() (common.Address, error) {
	return _Commitreveal.Contract.AddressBridge(&_Commitreveal.CallOpts)
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() constant returns(address)
func (_Commitreveal *CommitrevealCaller) BridgeAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Commitreveal.contract.Call(opts, out, "bridgeAddr")
	return *ret0, err
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() constant returns(address)
func (_Commitreveal *CommitrevealSession) BridgeAddr() (common.Address, error) {
	return _Commitreveal.Contract.BridgeAddr(&_Commitreveal.CallOpts)
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() constant returns(address)
func (_Commitreveal *CommitrevealCallerSession) BridgeAddr() (common.Address, error) {
	return _Commitreveal.Contract.BridgeAddr(&_Commitreveal.CallOpts)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns( uint256) constant returns(startBlock uint256, commitDuration uint256, revealDuration uint256, revealThreshold uint256, commitNum uint256, revealNum uint256, generatedRandom uint256)
func (_Commitreveal *CommitrevealCaller) Campaigns(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Commitreveal.contract.Call(opts, out, "campaigns", arg0)
	return *ret, err
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns( uint256) constant returns(startBlock uint256, commitDuration uint256, revealDuration uint256, revealThreshold uint256, commitNum uint256, revealNum uint256, generatedRandom uint256)
func (_Commitreveal *CommitrevealSession) Campaigns(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	GeneratedRandom *big.Int
}, error) {
	return _Commitreveal.Contract.Campaigns(&_Commitreveal.CallOpts, arg0)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns( uint256) constant returns(startBlock uint256, commitDuration uint256, revealDuration uint256, revealThreshold uint256, commitNum uint256, revealNum uint256, generatedRandom uint256)
func (_Commitreveal *CommitrevealCallerSession) Campaigns(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	GeneratedRandom *big.Int
}, error) {
	return _Commitreveal.Contract.Campaigns(&_Commitreveal.CallOpts, arg0)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(_cid uint256, _secretHash bytes32) returns()
func (_Commitreveal *CommitrevealTransactor) Commit(opts *bind.TransactOpts, _cid *big.Int, _secretHash [32]byte) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "commit", _cid, _secretHash)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(_cid uint256, _secretHash bytes32) returns()
func (_Commitreveal *CommitrevealSession) Commit(_cid *big.Int, _secretHash [32]byte) (*types.Transaction, error) {
	return _Commitreveal.Contract.Commit(&_Commitreveal.TransactOpts, _cid, _secretHash)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(_cid uint256, _secretHash bytes32) returns()
func (_Commitreveal *CommitrevealTransactorSession) Commit(_cid *big.Int, _secretHash [32]byte) (*types.Transaction, error) {
	return _Commitreveal.Contract.Commit(&_Commitreveal.TransactOpts, _cid, _secretHash)
}

// GetRandom is a paid mutator transaction binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(_cid uint256) returns(uint256)
func (_Commitreveal *CommitrevealTransactor) GetRandom(opts *bind.TransactOpts, _cid *big.Int) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "getRandom", _cid)
}

// GetRandom is a paid mutator transaction binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(_cid uint256) returns(uint256)
func (_Commitreveal *CommitrevealSession) GetRandom(_cid *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.GetRandom(&_Commitreveal.TransactOpts, _cid)
}

// GetRandom is a paid mutator transaction binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(_cid uint256) returns(uint256)
func (_Commitreveal *CommitrevealTransactorSession) GetRandom(_cid *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.GetRandom(&_Commitreveal.TransactOpts, _cid)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(_cid uint256, _secret uint256) returns()
func (_Commitreveal *CommitrevealTransactor) Reveal(opts *bind.TransactOpts, _cid *big.Int, _secret *big.Int) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "reveal", _cid, _secret)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(_cid uint256, _secret uint256) returns()
func (_Commitreveal *CommitrevealSession) Reveal(_cid *big.Int, _secret *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.Reveal(&_Commitreveal.TransactOpts, _cid, _secret)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(_cid uint256, _secret uint256) returns()
func (_Commitreveal *CommitrevealTransactorSession) Reveal(_cid *big.Int, _secret *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.Reveal(&_Commitreveal.TransactOpts, _cid, _secret)
}

// StartCommitReveal is a paid mutator transaction binding the contract method 0xb917b5a5.
//
// Solidity: function startCommitReveal(_startBlock uint256, _commitDuration uint256, _revealDuration uint256, _revealThreshold uint256) returns(uint256)
func (_Commitreveal *CommitrevealTransactor) StartCommitReveal(opts *bind.TransactOpts, _startBlock *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _revealThreshold *big.Int) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "startCommitReveal", _startBlock, _commitDuration, _revealDuration, _revealThreshold)
}

// StartCommitReveal is a paid mutator transaction binding the contract method 0xb917b5a5.
//
// Solidity: function startCommitReveal(_startBlock uint256, _commitDuration uint256, _revealDuration uint256, _revealThreshold uint256) returns(uint256)
func (_Commitreveal *CommitrevealSession) StartCommitReveal(_startBlock *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _revealThreshold *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.StartCommitReveal(&_Commitreveal.TransactOpts, _startBlock, _commitDuration, _revealDuration, _revealThreshold)
}

// StartCommitReveal is a paid mutator transaction binding the contract method 0xb917b5a5.
//
// Solidity: function startCommitReveal(_startBlock uint256, _commitDuration uint256, _revealDuration uint256, _revealThreshold uint256) returns(uint256)
func (_Commitreveal *CommitrevealTransactorSession) StartCommitReveal(_startBlock *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _revealThreshold *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.StartCommitReveal(&_Commitreveal.TransactOpts, _startBlock, _commitDuration, _revealDuration, _revealThreshold)
}

// CommitrevealLogCommitIterator is returned from FilterLogCommit and is used to iterate over the raw logs and unpacked data for LogCommit events raised by the Commitreveal contract.
type CommitrevealLogCommitIterator struct {
	Event *CommitrevealLogCommit // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogCommit)
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
		it.Event = new(CommitrevealLogCommit)
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
func (it *CommitrevealLogCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogCommit represents a LogCommit event raised by the Commitreveal contract.
type CommitrevealLogCommit struct {
	Cid        *big.Int
	From       common.Address
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogCommit is a free log retrieval operation binding the contract event 0x918c00c65dd2a8dee4c6985d1d67f04aa8cd2c93e8d427d398a90444c7f7c75e.
//
// Solidity: e LogCommit(cid uint256, from address, commitment bytes32)
func (_Commitreveal *CommitrevealFilterer) FilterLogCommit(opts *bind.FilterOpts) (*CommitrevealLogCommitIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogCommit")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogCommitIterator{contract: _Commitreveal.contract, event: "LogCommit", logs: logs, sub: sub}, nil
}

// WatchLogCommit is a free log subscription operation binding the contract event 0x918c00c65dd2a8dee4c6985d1d67f04aa8cd2c93e8d427d398a90444c7f7c75e.
//
// Solidity: e LogCommit(cid uint256, from address, commitment bytes32)
func (_Commitreveal *CommitrevealFilterer) WatchLogCommit(opts *bind.WatchOpts, sink chan<- *CommitrevealLogCommit) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogCommit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogCommit)
				if err := _Commitreveal.contract.UnpackLog(event, "LogCommit", log); err != nil {
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

// CommitrevealLogRandomIterator is returned from FilterLogRandom and is used to iterate over the raw logs and unpacked data for LogRandom events raised by the Commitreveal contract.
type CommitrevealLogRandomIterator struct {
	Event *CommitrevealLogRandom // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogRandomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogRandom)
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
		it.Event = new(CommitrevealLogRandom)
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
func (it *CommitrevealLogRandomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogRandomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogRandom represents a LogRandom event raised by the Commitreveal contract.
type CommitrevealLogRandom struct {
	Cid    *big.Int
	Random *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogRandom is a free log retrieval operation binding the contract event 0xa34f42a90fadfe357ee14419d618438a057569fbb63bab0c5fbcca0fc2b11e8d.
//
// Solidity: e LogRandom(cid uint256, random uint256)
func (_Commitreveal *CommitrevealFilterer) FilterLogRandom(opts *bind.FilterOpts) (*CommitrevealLogRandomIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogRandom")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogRandomIterator{contract: _Commitreveal.contract, event: "LogRandom", logs: logs, sub: sub}, nil
}

// WatchLogRandom is a free log subscription operation binding the contract event 0xa34f42a90fadfe357ee14419d618438a057569fbb63bab0c5fbcca0fc2b11e8d.
//
// Solidity: e LogRandom(cid uint256, random uint256)
func (_Commitreveal *CommitrevealFilterer) WatchLogRandom(opts *bind.WatchOpts, sink chan<- *CommitrevealLogRandom) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogRandom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogRandom)
				if err := _Commitreveal.contract.UnpackLog(event, "LogRandom", log); err != nil {
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

// CommitrevealLogRandomFailureIterator is returned from FilterLogRandomFailure and is used to iterate over the raw logs and unpacked data for LogRandomFailure events raised by the Commitreveal contract.
type CommitrevealLogRandomFailureIterator struct {
	Event *CommitrevealLogRandomFailure // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogRandomFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogRandomFailure)
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
		it.Event = new(CommitrevealLogRandomFailure)
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
func (it *CommitrevealLogRandomFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogRandomFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogRandomFailure represents a LogRandomFailure event raised by the Commitreveal contract.
type CommitrevealLogRandomFailure struct {
	Cid             *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	RevealThreshold *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogRandomFailure is a free log retrieval operation binding the contract event 0xe888e7582d0505bce81eef694dfa216179eaaa3c1bd96b7894de8b4370d8543e.
//
// Solidity: e LogRandomFailure(cid uint256, commitNum uint256, revealNum uint256, revealThreshold uint256)
func (_Commitreveal *CommitrevealFilterer) FilterLogRandomFailure(opts *bind.FilterOpts) (*CommitrevealLogRandomFailureIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogRandomFailure")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogRandomFailureIterator{contract: _Commitreveal.contract, event: "LogRandomFailure", logs: logs, sub: sub}, nil
}

// WatchLogRandomFailure is a free log subscription operation binding the contract event 0xe888e7582d0505bce81eef694dfa216179eaaa3c1bd96b7894de8b4370d8543e.
//
// Solidity: e LogRandomFailure(cid uint256, commitNum uint256, revealNum uint256, revealThreshold uint256)
func (_Commitreveal *CommitrevealFilterer) WatchLogRandomFailure(opts *bind.WatchOpts, sink chan<- *CommitrevealLogRandomFailure) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogRandomFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogRandomFailure)
				if err := _Commitreveal.contract.UnpackLog(event, "LogRandomFailure", log); err != nil {
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

// CommitrevealLogRevealIterator is returned from FilterLogReveal and is used to iterate over the raw logs and unpacked data for LogReveal events raised by the Commitreveal contract.
type CommitrevealLogRevealIterator struct {
	Event *CommitrevealLogReveal // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogReveal)
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
		it.Event = new(CommitrevealLogReveal)
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
func (it *CommitrevealLogRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogReveal represents a LogReveal event raised by the Commitreveal contract.
type CommitrevealLogReveal struct {
	Cid    *big.Int
	From   common.Address
	Secret *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogReveal is a free log retrieval operation binding the contract event 0x9141bfaedbc77aa7b8d9c989cd81909d95bb1677e556e34cfd45e50e0bea2967.
//
// Solidity: e LogReveal(cid uint256, from address, secret uint256)
func (_Commitreveal *CommitrevealFilterer) FilterLogReveal(opts *bind.FilterOpts) (*CommitrevealLogRevealIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogReveal")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogRevealIterator{contract: _Commitreveal.contract, event: "LogReveal", logs: logs, sub: sub}, nil
}

// WatchLogReveal is a free log subscription operation binding the contract event 0x9141bfaedbc77aa7b8d9c989cd81909d95bb1677e556e34cfd45e50e0bea2967.
//
// Solidity: e LogReveal(cid uint256, from address, secret uint256)
func (_Commitreveal *CommitrevealFilterer) WatchLogReveal(opts *bind.WatchOpts, sink chan<- *CommitrevealLogReveal) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogReveal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogReveal)
				if err := _Commitreveal.contract.UnpackLog(event, "LogReveal", log); err != nil {
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

// CommitrevealLogStartCommitRevealIterator is returned from FilterLogStartCommitReveal and is used to iterate over the raw logs and unpacked data for LogStartCommitReveal events raised by the Commitreveal contract.
type CommitrevealLogStartCommitRevealIterator struct {
	Event *CommitrevealLogStartCommitReveal // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogStartCommitRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogStartCommitReveal)
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
		it.Event = new(CommitrevealLogStartCommitReveal)
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
func (it *CommitrevealLogStartCommitRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogStartCommitRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogStartCommitReveal represents a LogStartCommitReveal event raised by the Commitreveal contract.
type CommitrevealLogStartCommitReveal struct {
	Cid             *big.Int
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogStartCommitReveal is a free log retrieval operation binding the contract event 0xbbfccb30e8cf1b5d88802741ceb4d63cf854fa8931eeeaec6b700f31f429dc09.
//
// Solidity: e LogStartCommitReveal(cid uint256, startBlock uint256, commitDuration uint256, revealDuration uint256, revealThreshold uint256)
func (_Commitreveal *CommitrevealFilterer) FilterLogStartCommitReveal(opts *bind.FilterOpts) (*CommitrevealLogStartCommitRevealIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogStartCommitReveal")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogStartCommitRevealIterator{contract: _Commitreveal.contract, event: "LogStartCommitReveal", logs: logs, sub: sub}, nil
}

// WatchLogStartCommitReveal is a free log subscription operation binding the contract event 0xbbfccb30e8cf1b5d88802741ceb4d63cf854fa8931eeeaec6b700f31f429dc09.
//
// Solidity: e LogStartCommitReveal(cid uint256, startBlock uint256, commitDuration uint256, revealDuration uint256, revealThreshold uint256)
func (_Commitreveal *CommitrevealFilterer) WatchLogStartCommitReveal(opts *bind.WatchOpts, sink chan<- *CommitrevealLogStartCommitReveal) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogStartCommitReveal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogStartCommitReveal)
				if err := _Commitreveal.contract.UnpackLog(event, "LogStartCommitReveal", log); err != nil {
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
