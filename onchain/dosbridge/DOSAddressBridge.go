// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dosbridge

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DosbridgeABI is the input ABI used to generate the binding from.
const DosbridgeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCommitRevealAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setProxyAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setPaymentAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setCommitRevealAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPaymentAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setRegistryAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegistryAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousProxy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newProxy\",\"type\":\"address\"}],\"name\":\"ProxyAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"CommitRevealAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousPayment\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newPayment\",\"type\":\"address\"}],\"name\":\"PaymentAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousRegistry\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"RegistryAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DosbridgeBin is the compiled bytecode used for deploying new contracts.
const DosbridgeBin = `608060405260008054600160a060020a031916331790556105b2806100256000396000f3fe608060405234801561001057600080fd5b50600436106100d1576000357c0100000000000000000000000000000000000000000000000000000000900480638da5cb5b1161008e5780638da5cb5b1461017e5780638f32d59b146101865780639d265e58146101a2578063ab7b4993146101aa578063f21de1e8146101d0578063f2fde38b146101d8576100d1565b80631ae0433c146100d657806343a73d9a146100fa57806346a7dadc146101025780635e1e10041461012a578063715018a6146101505780637b08cd0314610158575b600080fd5b6100de6101fe565b60408051600160a060020a039092168252519081900360200190f35b6100de61020d565b6101286004803603602081101561011857600080fd5b5035600160a060020a031661021c565b005b6101286004803603602081101561014057600080fd5b5035600160a060020a03166102a6565b610128610330565b6101286004803603602081101561016e57600080fd5b5035600160a060020a0316610398565b6100de610422565b61018e610431565b604080519115158252519081900360200190f35b6100de610442565b610128600480360360208110156101c057600080fd5b5035600160a060020a0316610451565b6100de6104db565b610128600480360360208110156101ee57600080fd5b5035600160a060020a03166104ea565b600254600160a060020a031690565b600154600160a060020a031690565b610224610431565b151561022f57600080fd5b60015460408051600160a060020a039283168152918316602083015280517fafa5c16901af5d392255707d27b3e2687e79a18df187b9f1525e7f0fc2144f6f9281900390910190a16001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6102ae610431565b15156102b957600080fd5b60035460408051600160a060020a039283168152918316602083015280517fb3d3f832f05d764f8934189cba7879e2dd829dd3f92749ec959339fd5cd8b0be9281900390910190a16003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b610338610431565b151561034357600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6103a0610431565b15156103ab57600080fd5b60025460408051600160a060020a039283168152918316602083015280517f23b082fc42fcc9c7d42de567b56abef6a737aa2600b8036ee5c304086a2545c39281900390910190a16002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a031690565b600054600160a060020a0316331490565b600354600160a060020a031690565b610459610431565b151561046457600080fd5b60045460408051600160a060020a039283168152918316602083015280517f6144918c239a794463afd709d2affba8e0a35b21444f4d461c9d700a2d6bb5049281900390910190a16004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600454600160a060020a031690565b6104f2610431565b15156104fd57600080fd5b61050681610509565b50565b600160a060020a038116151561051e57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556fea165627a7a7230582074506d54623b60795bbfca0f552388d4c34561e544ef350dcd77e678a5fdfc7d0029`

// DeployDosbridge deploys a new Ethereum contract, binding an instance of Dosbridge to it.
func DeployDosbridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dosbridge, error) {
	parsed, err := abi.JSON(strings.NewReader(DosbridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DosbridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dosbridge{DosbridgeCaller: DosbridgeCaller{contract: contract}, DosbridgeTransactor: DosbridgeTransactor{contract: contract}, DosbridgeFilterer: DosbridgeFilterer{contract: contract}}, nil
}

// Dosbridge is an auto generated Go binding around an Ethereum contract.
type Dosbridge struct {
	DosbridgeCaller     // Read-only binding to the contract
	DosbridgeTransactor // Write-only binding to the contract
	DosbridgeFilterer   // Log filterer for contract events
}

// DosbridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DosbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DosbridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DosbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DosbridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DosbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DosbridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DosbridgeSession struct {
	Contract     *Dosbridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DosbridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DosbridgeCallerSession struct {
	Contract *DosbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DosbridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DosbridgeTransactorSession struct {
	Contract     *DosbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DosbridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DosbridgeRaw struct {
	Contract *Dosbridge // Generic contract binding to access the raw methods on
}

// DosbridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DosbridgeCallerRaw struct {
	Contract *DosbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// DosbridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DosbridgeTransactorRaw struct {
	Contract *DosbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDosbridge creates a new instance of Dosbridge, bound to a specific deployed contract.
func NewDosbridge(address common.Address, backend bind.ContractBackend) (*Dosbridge, error) {
	contract, err := bindDosbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dosbridge{DosbridgeCaller: DosbridgeCaller{contract: contract}, DosbridgeTransactor: DosbridgeTransactor{contract: contract}, DosbridgeFilterer: DosbridgeFilterer{contract: contract}}, nil
}

// NewDosbridgeCaller creates a new read-only instance of Dosbridge, bound to a specific deployed contract.
func NewDosbridgeCaller(address common.Address, caller bind.ContractCaller) (*DosbridgeCaller, error) {
	contract, err := bindDosbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DosbridgeCaller{contract: contract}, nil
}

// NewDosbridgeTransactor creates a new write-only instance of Dosbridge, bound to a specific deployed contract.
func NewDosbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*DosbridgeTransactor, error) {
	contract, err := bindDosbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DosbridgeTransactor{contract: contract}, nil
}

// NewDosbridgeFilterer creates a new log filterer instance of Dosbridge, bound to a specific deployed contract.
func NewDosbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*DosbridgeFilterer, error) {
	contract, err := bindDosbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DosbridgeFilterer{contract: contract}, nil
}

// bindDosbridge binds a generic wrapper to an already deployed contract.
func bindDosbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DosbridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dosbridge *DosbridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dosbridge.Contract.DosbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dosbridge *DosbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosbridge.Contract.DosbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dosbridge *DosbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dosbridge.Contract.DosbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dosbridge *DosbridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dosbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dosbridge *DosbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dosbridge *DosbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dosbridge.Contract.contract.Transact(opts, method, params...)
}

// GetCommitRevealAddress is a free data retrieval call binding the contract method 0x1ae0433c.
//
// Solidity: function getCommitRevealAddress() constant returns(address)
func (_Dosbridge *DosbridgeCaller) GetCommitRevealAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosbridge.contract.Call(opts, out, "getCommitRevealAddress")
	return *ret0, err
}

// GetCommitRevealAddress is a free data retrieval call binding the contract method 0x1ae0433c.
//
// Solidity: function getCommitRevealAddress() constant returns(address)
func (_Dosbridge *DosbridgeSession) GetCommitRevealAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetCommitRevealAddress(&_Dosbridge.CallOpts)
}

// GetCommitRevealAddress is a free data retrieval call binding the contract method 0x1ae0433c.
//
// Solidity: function getCommitRevealAddress() constant returns(address)
func (_Dosbridge *DosbridgeCallerSession) GetCommitRevealAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetCommitRevealAddress(&_Dosbridge.CallOpts)
}

// GetPaymentAddress is a free data retrieval call binding the contract method 0x9d265e58.
//
// Solidity: function getPaymentAddress() constant returns(address)
func (_Dosbridge *DosbridgeCaller) GetPaymentAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosbridge.contract.Call(opts, out, "getPaymentAddress")
	return *ret0, err
}

// GetPaymentAddress is a free data retrieval call binding the contract method 0x9d265e58.
//
// Solidity: function getPaymentAddress() constant returns(address)
func (_Dosbridge *DosbridgeSession) GetPaymentAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetPaymentAddress(&_Dosbridge.CallOpts)
}

// GetPaymentAddress is a free data retrieval call binding the contract method 0x9d265e58.
//
// Solidity: function getPaymentAddress() constant returns(address)
func (_Dosbridge *DosbridgeCallerSession) GetPaymentAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetPaymentAddress(&_Dosbridge.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_Dosbridge *DosbridgeCaller) GetProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosbridge.contract.Call(opts, out, "getProxyAddress")
	return *ret0, err
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_Dosbridge *DosbridgeSession) GetProxyAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetProxyAddress(&_Dosbridge.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_Dosbridge *DosbridgeCallerSession) GetProxyAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetProxyAddress(&_Dosbridge.CallOpts)
}

// GetRegistryAddress is a free data retrieval call binding the contract method 0xf21de1e8.
//
// Solidity: function getRegistryAddress() constant returns(address)
func (_Dosbridge *DosbridgeCaller) GetRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosbridge.contract.Call(opts, out, "getRegistryAddress")
	return *ret0, err
}

// GetRegistryAddress is a free data retrieval call binding the contract method 0xf21de1e8.
//
// Solidity: function getRegistryAddress() constant returns(address)
func (_Dosbridge *DosbridgeSession) GetRegistryAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetRegistryAddress(&_Dosbridge.CallOpts)
}

// GetRegistryAddress is a free data retrieval call binding the contract method 0xf21de1e8.
//
// Solidity: function getRegistryAddress() constant returns(address)
func (_Dosbridge *DosbridgeCallerSession) GetRegistryAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetRegistryAddress(&_Dosbridge.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dosbridge *DosbridgeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dosbridge.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dosbridge *DosbridgeSession) IsOwner() (bool, error) {
	return _Dosbridge.Contract.IsOwner(&_Dosbridge.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dosbridge *DosbridgeCallerSession) IsOwner() (bool, error) {
	return _Dosbridge.Contract.IsOwner(&_Dosbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dosbridge *DosbridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosbridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dosbridge *DosbridgeSession) Owner() (common.Address, error) {
	return _Dosbridge.Contract.Owner(&_Dosbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dosbridge *DosbridgeCallerSession) Owner() (common.Address, error) {
	return _Dosbridge.Contract.Owner(&_Dosbridge.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dosbridge *DosbridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dosbridge *DosbridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dosbridge.Contract.RenounceOwnership(&_Dosbridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dosbridge *DosbridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dosbridge.Contract.RenounceOwnership(&_Dosbridge.TransactOpts)
}

// SetCommitRevealAddress is a paid mutator transaction binding the contract method 0x7b08cd03.
//
// Solidity: function setCommitRevealAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeTransactor) SetCommitRevealAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "setCommitRevealAddress", newAddr)
}

// SetCommitRevealAddress is a paid mutator transaction binding the contract method 0x7b08cd03.
//
// Solidity: function setCommitRevealAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeSession) SetCommitRevealAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetCommitRevealAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetCommitRevealAddress is a paid mutator transaction binding the contract method 0x7b08cd03.
//
// Solidity: function setCommitRevealAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeTransactorSession) SetCommitRevealAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetCommitRevealAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetPaymentAddress is a paid mutator transaction binding the contract method 0x5e1e1004.
//
// Solidity: function setPaymentAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeTransactor) SetPaymentAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "setPaymentAddress", newAddr)
}

// SetPaymentAddress is a paid mutator transaction binding the contract method 0x5e1e1004.
//
// Solidity: function setPaymentAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeSession) SetPaymentAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetPaymentAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetPaymentAddress is a paid mutator transaction binding the contract method 0x5e1e1004.
//
// Solidity: function setPaymentAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeTransactorSession) SetPaymentAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetPaymentAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetProxyAddress is a paid mutator transaction binding the contract method 0x46a7dadc.
//
// Solidity: function setProxyAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeTransactor) SetProxyAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "setProxyAddress", newAddr)
}

// SetProxyAddress is a paid mutator transaction binding the contract method 0x46a7dadc.
//
// Solidity: function setProxyAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeSession) SetProxyAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetProxyAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetProxyAddress is a paid mutator transaction binding the contract method 0x46a7dadc.
//
// Solidity: function setProxyAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeTransactorSession) SetProxyAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetProxyAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetRegistryAddress is a paid mutator transaction binding the contract method 0xab7b4993.
//
// Solidity: function setRegistryAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeTransactor) SetRegistryAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "setRegistryAddress", newAddr)
}

// SetRegistryAddress is a paid mutator transaction binding the contract method 0xab7b4993.
//
// Solidity: function setRegistryAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeSession) SetRegistryAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetRegistryAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetRegistryAddress is a paid mutator transaction binding the contract method 0xab7b4993.
//
// Solidity: function setRegistryAddress(address newAddr) returns()
func (_Dosbridge *DosbridgeTransactorSession) SetRegistryAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetRegistryAddress(&_Dosbridge.TransactOpts, newAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dosbridge *DosbridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dosbridge *DosbridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.TransferOwnership(&_Dosbridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dosbridge *DosbridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.TransferOwnership(&_Dosbridge.TransactOpts, newOwner)
}

// DosbridgeCommitRevealAddressUpdatedIterator is returned from FilterCommitRevealAddressUpdated and is used to iterate over the raw logs and unpacked data for CommitRevealAddressUpdated events raised by the Dosbridge contract.
type DosbridgeCommitRevealAddressUpdatedIterator struct {
	Event *DosbridgeCommitRevealAddressUpdated // Event containing the contract specifics and raw log

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
func (it *DosbridgeCommitRevealAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosbridgeCommitRevealAddressUpdated)
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
		it.Event = new(DosbridgeCommitRevealAddressUpdated)
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
func (it *DosbridgeCommitRevealAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosbridgeCommitRevealAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosbridgeCommitRevealAddressUpdated represents a CommitRevealAddressUpdated event raised by the Dosbridge contract.
type DosbridgeCommitRevealAddressUpdated struct {
	PreviousAddr common.Address
	NewAddr      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCommitRevealAddressUpdated is a free log retrieval operation binding the contract event 0x23b082fc42fcc9c7d42de567b56abef6a737aa2600b8036ee5c304086a2545c3.
//
// Solidity: event CommitRevealAddressUpdated(address previousAddr, address newAddr)
func (_Dosbridge *DosbridgeFilterer) FilterCommitRevealAddressUpdated(opts *bind.FilterOpts) (*DosbridgeCommitRevealAddressUpdatedIterator, error) {

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "CommitRevealAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DosbridgeCommitRevealAddressUpdatedIterator{contract: _Dosbridge.contract, event: "CommitRevealAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitRevealAddressUpdated is a free log subscription operation binding the contract event 0x23b082fc42fcc9c7d42de567b56abef6a737aa2600b8036ee5c304086a2545c3.
//
// Solidity: event CommitRevealAddressUpdated(address previousAddr, address newAddr)
func (_Dosbridge *DosbridgeFilterer) WatchCommitRevealAddressUpdated(opts *bind.WatchOpts, sink chan<- *DosbridgeCommitRevealAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _Dosbridge.contract.WatchLogs(opts, "CommitRevealAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosbridgeCommitRevealAddressUpdated)
				if err := _Dosbridge.contract.UnpackLog(event, "CommitRevealAddressUpdated", log); err != nil {
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

// DosbridgeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Dosbridge contract.
type DosbridgeOwnershipRenouncedIterator struct {
	Event *DosbridgeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DosbridgeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosbridgeOwnershipRenounced)
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
		it.Event = new(DosbridgeOwnershipRenounced)
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
func (it *DosbridgeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosbridgeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosbridgeOwnershipRenounced represents a OwnershipRenounced event raised by the Dosbridge contract.
type DosbridgeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Dosbridge *DosbridgeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DosbridgeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DosbridgeOwnershipRenouncedIterator{contract: _Dosbridge.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Dosbridge *DosbridgeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DosbridgeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Dosbridge.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosbridgeOwnershipRenounced)
				if err := _Dosbridge.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DosbridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dosbridge contract.
type DosbridgeOwnershipTransferredIterator struct {
	Event *DosbridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DosbridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosbridgeOwnershipTransferred)
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
		it.Event = new(DosbridgeOwnershipTransferred)
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
func (it *DosbridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosbridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosbridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Dosbridge contract.
type DosbridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dosbridge *DosbridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DosbridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DosbridgeOwnershipTransferredIterator{contract: _Dosbridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dosbridge *DosbridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DosbridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dosbridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosbridgeOwnershipTransferred)
				if err := _Dosbridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// DosbridgePaymentAddressUpdatedIterator is returned from FilterPaymentAddressUpdated and is used to iterate over the raw logs and unpacked data for PaymentAddressUpdated events raised by the Dosbridge contract.
type DosbridgePaymentAddressUpdatedIterator struct {
	Event *DosbridgePaymentAddressUpdated // Event containing the contract specifics and raw log

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
func (it *DosbridgePaymentAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosbridgePaymentAddressUpdated)
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
		it.Event = new(DosbridgePaymentAddressUpdated)
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
func (it *DosbridgePaymentAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosbridgePaymentAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosbridgePaymentAddressUpdated represents a PaymentAddressUpdated event raised by the Dosbridge contract.
type DosbridgePaymentAddressUpdated struct {
	PreviousPayment common.Address
	NewPayment      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaymentAddressUpdated is a free log retrieval operation binding the contract event 0xb3d3f832f05d764f8934189cba7879e2dd829dd3f92749ec959339fd5cd8b0be.
//
// Solidity: event PaymentAddressUpdated(address previousPayment, address newPayment)
func (_Dosbridge *DosbridgeFilterer) FilterPaymentAddressUpdated(opts *bind.FilterOpts) (*DosbridgePaymentAddressUpdatedIterator, error) {

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "PaymentAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DosbridgePaymentAddressUpdatedIterator{contract: _Dosbridge.contract, event: "PaymentAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPaymentAddressUpdated is a free log subscription operation binding the contract event 0xb3d3f832f05d764f8934189cba7879e2dd829dd3f92749ec959339fd5cd8b0be.
//
// Solidity: event PaymentAddressUpdated(address previousPayment, address newPayment)
func (_Dosbridge *DosbridgeFilterer) WatchPaymentAddressUpdated(opts *bind.WatchOpts, sink chan<- *DosbridgePaymentAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _Dosbridge.contract.WatchLogs(opts, "PaymentAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosbridgePaymentAddressUpdated)
				if err := _Dosbridge.contract.UnpackLog(event, "PaymentAddressUpdated", log); err != nil {
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

// DosbridgeProxyAddressUpdatedIterator is returned from FilterProxyAddressUpdated and is used to iterate over the raw logs and unpacked data for ProxyAddressUpdated events raised by the Dosbridge contract.
type DosbridgeProxyAddressUpdatedIterator struct {
	Event *DosbridgeProxyAddressUpdated // Event containing the contract specifics and raw log

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
func (it *DosbridgeProxyAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosbridgeProxyAddressUpdated)
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
		it.Event = new(DosbridgeProxyAddressUpdated)
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
func (it *DosbridgeProxyAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosbridgeProxyAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosbridgeProxyAddressUpdated represents a ProxyAddressUpdated event raised by the Dosbridge contract.
type DosbridgeProxyAddressUpdated struct {
	PreviousProxy common.Address
	NewProxy      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProxyAddressUpdated is a free log retrieval operation binding the contract event 0xafa5c16901af5d392255707d27b3e2687e79a18df187b9f1525e7f0fc2144f6f.
//
// Solidity: event ProxyAddressUpdated(address previousProxy, address newProxy)
func (_Dosbridge *DosbridgeFilterer) FilterProxyAddressUpdated(opts *bind.FilterOpts) (*DosbridgeProxyAddressUpdatedIterator, error) {

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "ProxyAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DosbridgeProxyAddressUpdatedIterator{contract: _Dosbridge.contract, event: "ProxyAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyAddressUpdated is a free log subscription operation binding the contract event 0xafa5c16901af5d392255707d27b3e2687e79a18df187b9f1525e7f0fc2144f6f.
//
// Solidity: event ProxyAddressUpdated(address previousProxy, address newProxy)
func (_Dosbridge *DosbridgeFilterer) WatchProxyAddressUpdated(opts *bind.WatchOpts, sink chan<- *DosbridgeProxyAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _Dosbridge.contract.WatchLogs(opts, "ProxyAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosbridgeProxyAddressUpdated)
				if err := _Dosbridge.contract.UnpackLog(event, "ProxyAddressUpdated", log); err != nil {
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

// DosbridgeRegistryAddressUpdatedIterator is returned from FilterRegistryAddressUpdated and is used to iterate over the raw logs and unpacked data for RegistryAddressUpdated events raised by the Dosbridge contract.
type DosbridgeRegistryAddressUpdatedIterator struct {
	Event *DosbridgeRegistryAddressUpdated // Event containing the contract specifics and raw log

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
func (it *DosbridgeRegistryAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosbridgeRegistryAddressUpdated)
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
		it.Event = new(DosbridgeRegistryAddressUpdated)
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
func (it *DosbridgeRegistryAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosbridgeRegistryAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosbridgeRegistryAddressUpdated represents a RegistryAddressUpdated event raised by the Dosbridge contract.
type DosbridgeRegistryAddressUpdated struct {
	PreviousRegistry common.Address
	NewRegistry      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRegistryAddressUpdated is a free log retrieval operation binding the contract event 0x6144918c239a794463afd709d2affba8e0a35b21444f4d461c9d700a2d6bb504.
//
// Solidity: event RegistryAddressUpdated(address previousRegistry, address newRegistry)
func (_Dosbridge *DosbridgeFilterer) FilterRegistryAddressUpdated(opts *bind.FilterOpts) (*DosbridgeRegistryAddressUpdatedIterator, error) {

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "RegistryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DosbridgeRegistryAddressUpdatedIterator{contract: _Dosbridge.contract, event: "RegistryAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchRegistryAddressUpdated is a free log subscription operation binding the contract event 0x6144918c239a794463afd709d2affba8e0a35b21444f4d461c9d700a2d6bb504.
//
// Solidity: event RegistryAddressUpdated(address previousRegistry, address newRegistry)
func (_Dosbridge *DosbridgeFilterer) WatchRegistryAddressUpdated(opts *bind.WatchOpts, sink chan<- *DosbridgeRegistryAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _Dosbridge.contract.WatchLogs(opts, "RegistryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosbridgeRegistryAddressUpdated)
				if err := _Dosbridge.contract.UnpackLog(event, "RegistryAddressUpdated", log); err != nil {
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
