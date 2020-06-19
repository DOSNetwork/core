// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dosbridge

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DosbridgeABI is the input ABI used to generate the binding from.
const DosbridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"previousURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURL\",\"type\":\"string\"}],\"name\":\"BootStrapUrlUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"CommitRevealAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousPayment\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPayment\",\"type\":\"address\"}],\"name\":\"PaymentAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxy\",\"type\":\"address\"}],\"name\":\"ProxyAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousStaking\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStaking\",\"type\":\"address\"}],\"name\":\"StakingAddressUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBootStrapUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommitRevealAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPaymentAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"setBootStrapUrl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setCommitRevealAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setPaymentAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setProxyAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"setStakingAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DosbridgeBin is the compiled bytecode used for deploying new contracts.
const DosbridgeBin = `6080604052600080546001600160a01b03191633179055610907806100256000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637b08cd031161008c5780639d265e58116100665780639d265e5814610269578063b7e982be14610271578063f2fde38b146102ee578063f4e0d9ac14610314576100ea565b80637b08cd031461021f5780638da5cb5b146102455780638f32d59b1461024d576100ea565b80634400bc07116100c85780634400bc071461012357806346a7dadc146101cb5780635e1e1004146101f1578063715018a614610217576100ea565b80630e9ed68b146100ef5780631ae0433c1461011357806343a73d9a1461011b575b600080fd5b6100f761033a565b604080516001600160a01b039092168252519081900360200190f35b6100f761034a565b6100f7610359565b6101c96004803603602081101561013957600080fd5b81019060208101813564010000000081111561015457600080fd5b82018360208201111561016657600080fd5b8035906020019184600183028401116401000000008311171561018857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610368945050505050565b005b6101c9600480360360208110156101e157600080fd5b50356001600160a01b03166104a5565b6101c96004803603602081101561020757600080fd5b50356001600160a01b0316610520565b6101c961059b565b6101c96004803603602081101561023557600080fd5b50356001600160a01b03166105f4565b6100f761066f565b61025561067e565b604080519115158252519081900360200190f35b6100f761068f565b61027961069e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102b357818101518382015260200161029b565b50505050905090810190601f1680156102e05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101c96004803603602081101561030457600080fd5b50356001600160a01b0316610734565b6101c96004803603602081101561032a57600080fd5b50356001600160a01b0316610751565b6004546001600160a01b03165b90565b6002546001600160a01b031690565b6001546001600160a01b031690565b61037061067e565b61037957600080fd5b6040805181815260058054600260001961010060018416150201909116049282018390527fc2194dd450e596fc07061b41e1cb9e4d38bd372ed38c6f909979d464f71cde7c92909184918190602082019060608301908690801561041e5780601f106103f35761010080835404028352916020019161041e565b820191906000526020600020905b81548152906001019060200180831161040157829003601f168201915b5050838103825284518152845160209182019186019080838360005b8381101561045257818101518382015260200161043a565b50505050905090810190601f16801561047f5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a180516104a190600590602084019061083a565b5050565b6104ad61067e565b6104b657600080fd5b600154604080516001600160a01b039283168152918316602083015280517fafa5c16901af5d392255707d27b3e2687e79a18df187b9f1525e7f0fc2144f6f9281900390910190a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b61052861067e565b61053157600080fd5b600354604080516001600160a01b039283168152918316602083015280517fb3d3f832f05d764f8934189cba7879e2dd829dd3f92749ec959339fd5cd8b0be9281900390910190a1600380546001600160a01b0319166001600160a01b0392909216919091179055565b6105a361067e565b6105ac57600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6105fc61067e565b61060557600080fd5b600254604080516001600160a01b039283168152918316602083015280517f23b082fc42fcc9c7d42de567b56abef6a737aa2600b8036ee5c304086a2545c39281900390910190a1600280546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b6003546001600160a01b031690565b60058054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561072a5780601f106106ff5761010080835404028352916020019161072a565b820191906000526020600020905b81548152906001019060200180831161070d57829003601f168201915b5050505050905090565b61073c61067e565b61074557600080fd5b61074e816107cc565b50565b61075961067e565b61076257600080fd5b600454604080516001600160a01b039283168152918316602083015280517f03fbfa1263b46c684780f3c24be11a2e189a59bedf0e316a7eae861cc769eb4f9281900390910190a1600480546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0381166107df57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061087b57805160ff19168380011785556108a8565b828001600101855582156108a8579182015b828111156108a857825182559160200191906001019061088d565b506108b49291506108b8565b5090565b61034791905b808211156108b457600081556001016108be56fea265627a7a72315820df4dfc812716bbcc980d3e18688b9bb09820906604fccac5a60ef6891323ecbe64736f6c63430005110032`

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

// GetBootStrapUrl is a free data retrieval call binding the contract method 0xb7e982be.
//
// Solidity: function getBootStrapUrl() constant returns(string)
func (_Dosbridge *DosbridgeCaller) GetBootStrapUrl(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Dosbridge.contract.Call(opts, out, "getBootStrapUrl")
	return *ret0, err
}

// GetBootStrapUrl is a free data retrieval call binding the contract method 0xb7e982be.
//
// Solidity: function getBootStrapUrl() constant returns(string)
func (_Dosbridge *DosbridgeSession) GetBootStrapUrl() (string, error) {
	return _Dosbridge.Contract.GetBootStrapUrl(&_Dosbridge.CallOpts)
}

// GetBootStrapUrl is a free data retrieval call binding the contract method 0xb7e982be.
//
// Solidity: function getBootStrapUrl() constant returns(string)
func (_Dosbridge *DosbridgeCallerSession) GetBootStrapUrl() (string, error) {
	return _Dosbridge.Contract.GetBootStrapUrl(&_Dosbridge.CallOpts)
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

// GetStakingAddress is a free data retrieval call binding the contract method 0x0e9ed68b.
//
// Solidity: function getStakingAddress() constant returns(address)
func (_Dosbridge *DosbridgeCaller) GetStakingAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dosbridge.contract.Call(opts, out, "getStakingAddress")
	return *ret0, err
}

// GetStakingAddress is a free data retrieval call binding the contract method 0x0e9ed68b.
//
// Solidity: function getStakingAddress() constant returns(address)
func (_Dosbridge *DosbridgeSession) GetStakingAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetStakingAddress(&_Dosbridge.CallOpts)
}

// GetStakingAddress is a free data retrieval call binding the contract method 0x0e9ed68b.
//
// Solidity: function getStakingAddress() constant returns(address)
func (_Dosbridge *DosbridgeCallerSession) GetStakingAddress() (common.Address, error) {
	return _Dosbridge.Contract.GetStakingAddress(&_Dosbridge.CallOpts)
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

// SetBootStrapUrl is a paid mutator transaction binding the contract method 0x4400bc07.
//
// Solidity: function setBootStrapUrl(url string) returns()
func (_Dosbridge *DosbridgeTransactor) SetBootStrapUrl(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "setBootStrapUrl", url)
}

// SetBootStrapUrl is a paid mutator transaction binding the contract method 0x4400bc07.
//
// Solidity: function setBootStrapUrl(url string) returns()
func (_Dosbridge *DosbridgeSession) SetBootStrapUrl(url string) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetBootStrapUrl(&_Dosbridge.TransactOpts, url)
}

// SetBootStrapUrl is a paid mutator transaction binding the contract method 0x4400bc07.
//
// Solidity: function setBootStrapUrl(url string) returns()
func (_Dosbridge *DosbridgeTransactorSession) SetBootStrapUrl(url string) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetBootStrapUrl(&_Dosbridge.TransactOpts, url)
}

// SetCommitRevealAddress is a paid mutator transaction binding the contract method 0x7b08cd03.
//
// Solidity: function setCommitRevealAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeTransactor) SetCommitRevealAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "setCommitRevealAddress", newAddr)
}

// SetCommitRevealAddress is a paid mutator transaction binding the contract method 0x7b08cd03.
//
// Solidity: function setCommitRevealAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeSession) SetCommitRevealAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetCommitRevealAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetCommitRevealAddress is a paid mutator transaction binding the contract method 0x7b08cd03.
//
// Solidity: function setCommitRevealAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeTransactorSession) SetCommitRevealAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetCommitRevealAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetPaymentAddress is a paid mutator transaction binding the contract method 0x5e1e1004.
//
// Solidity: function setPaymentAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeTransactor) SetPaymentAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "setPaymentAddress", newAddr)
}

// SetPaymentAddress is a paid mutator transaction binding the contract method 0x5e1e1004.
//
// Solidity: function setPaymentAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeSession) SetPaymentAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetPaymentAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetPaymentAddress is a paid mutator transaction binding the contract method 0x5e1e1004.
//
// Solidity: function setPaymentAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeTransactorSession) SetPaymentAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetPaymentAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetProxyAddress is a paid mutator transaction binding the contract method 0x46a7dadc.
//
// Solidity: function setProxyAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeTransactor) SetProxyAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "setProxyAddress", newAddr)
}

// SetProxyAddress is a paid mutator transaction binding the contract method 0x46a7dadc.
//
// Solidity: function setProxyAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeSession) SetProxyAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetProxyAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetProxyAddress is a paid mutator transaction binding the contract method 0x46a7dadc.
//
// Solidity: function setProxyAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeTransactorSession) SetProxyAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetProxyAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetStakingAddress is a paid mutator transaction binding the contract method 0xf4e0d9ac.
//
// Solidity: function setStakingAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeTransactor) SetStakingAddress(opts *bind.TransactOpts, newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "setStakingAddress", newAddr)
}

// SetStakingAddress is a paid mutator transaction binding the contract method 0xf4e0d9ac.
//
// Solidity: function setStakingAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeSession) SetStakingAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetStakingAddress(&_Dosbridge.TransactOpts, newAddr)
}

// SetStakingAddress is a paid mutator transaction binding the contract method 0xf4e0d9ac.
//
// Solidity: function setStakingAddress(newAddr address) returns()
func (_Dosbridge *DosbridgeTransactorSession) SetStakingAddress(newAddr common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.SetStakingAddress(&_Dosbridge.TransactOpts, newAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dosbridge *DosbridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dosbridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dosbridge *DosbridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.TransferOwnership(&_Dosbridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dosbridge *DosbridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dosbridge.Contract.TransferOwnership(&_Dosbridge.TransactOpts, newOwner)
}

// DosbridgeBootStrapUrlUpdatedIterator is returned from FilterBootStrapUrlUpdated and is used to iterate over the raw logs and unpacked data for BootStrapUrlUpdated events raised by the Dosbridge contract.
type DosbridgeBootStrapUrlUpdatedIterator struct {
	Event *DosbridgeBootStrapUrlUpdated // Event containing the contract specifics and raw log

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
func (it *DosbridgeBootStrapUrlUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosbridgeBootStrapUrlUpdated)
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
		it.Event = new(DosbridgeBootStrapUrlUpdated)
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
func (it *DosbridgeBootStrapUrlUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosbridgeBootStrapUrlUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosbridgeBootStrapUrlUpdated represents a BootStrapUrlUpdated event raised by the Dosbridge contract.
type DosbridgeBootStrapUrlUpdated struct {
	PreviousURL string
	NewURL      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBootStrapUrlUpdated is a free log retrieval operation binding the contract event 0xc2194dd450e596fc07061b41e1cb9e4d38bd372ed38c6f909979d464f71cde7c.
//
// Solidity: e BootStrapUrlUpdated(previousURL string, newURL string)
func (_Dosbridge *DosbridgeFilterer) FilterBootStrapUrlUpdated(opts *bind.FilterOpts) (*DosbridgeBootStrapUrlUpdatedIterator, error) {

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "BootStrapUrlUpdated")
	if err != nil {
		return nil, err
	}
	return &DosbridgeBootStrapUrlUpdatedIterator{contract: _Dosbridge.contract, event: "BootStrapUrlUpdated", logs: logs, sub: sub}, nil
}

// WatchBootStrapUrlUpdated is a free log subscription operation binding the contract event 0xc2194dd450e596fc07061b41e1cb9e4d38bd372ed38c6f909979d464f71cde7c.
//
// Solidity: e BootStrapUrlUpdated(previousURL string, newURL string)
func (_Dosbridge *DosbridgeFilterer) WatchBootStrapUrlUpdated(opts *bind.WatchOpts, sink chan<- *DosbridgeBootStrapUrlUpdated) (event.Subscription, error) {

	logs, sub, err := _Dosbridge.contract.WatchLogs(opts, "BootStrapUrlUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosbridgeBootStrapUrlUpdated)
				if err := _Dosbridge.contract.UnpackLog(event, "BootStrapUrlUpdated", log); err != nil {
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
// Solidity: e CommitRevealAddressUpdated(previousAddr address, newAddr address)
func (_Dosbridge *DosbridgeFilterer) FilterCommitRevealAddressUpdated(opts *bind.FilterOpts) (*DosbridgeCommitRevealAddressUpdatedIterator, error) {

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "CommitRevealAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DosbridgeCommitRevealAddressUpdatedIterator{contract: _Dosbridge.contract, event: "CommitRevealAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitRevealAddressUpdated is a free log subscription operation binding the contract event 0x23b082fc42fcc9c7d42de567b56abef6a737aa2600b8036ee5c304086a2545c3.
//
// Solidity: e CommitRevealAddressUpdated(previousAddr address, newAddr address)
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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e PaymentAddressUpdated(previousPayment address, newPayment address)
func (_Dosbridge *DosbridgeFilterer) FilterPaymentAddressUpdated(opts *bind.FilterOpts) (*DosbridgePaymentAddressUpdatedIterator, error) {

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "PaymentAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DosbridgePaymentAddressUpdatedIterator{contract: _Dosbridge.contract, event: "PaymentAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPaymentAddressUpdated is a free log subscription operation binding the contract event 0xb3d3f832f05d764f8934189cba7879e2dd829dd3f92749ec959339fd5cd8b0be.
//
// Solidity: e PaymentAddressUpdated(previousPayment address, newPayment address)
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
// Solidity: e ProxyAddressUpdated(previousProxy address, newProxy address)
func (_Dosbridge *DosbridgeFilterer) FilterProxyAddressUpdated(opts *bind.FilterOpts) (*DosbridgeProxyAddressUpdatedIterator, error) {

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "ProxyAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DosbridgeProxyAddressUpdatedIterator{contract: _Dosbridge.contract, event: "ProxyAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyAddressUpdated is a free log subscription operation binding the contract event 0xafa5c16901af5d392255707d27b3e2687e79a18df187b9f1525e7f0fc2144f6f.
//
// Solidity: e ProxyAddressUpdated(previousProxy address, newProxy address)
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

// DosbridgeStakingAddressUpdatedIterator is returned from FilterStakingAddressUpdated and is used to iterate over the raw logs and unpacked data for StakingAddressUpdated events raised by the Dosbridge contract.
type DosbridgeStakingAddressUpdatedIterator struct {
	Event *DosbridgeStakingAddressUpdated // Event containing the contract specifics and raw log

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
func (it *DosbridgeStakingAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DosbridgeStakingAddressUpdated)
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
		it.Event = new(DosbridgeStakingAddressUpdated)
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
func (it *DosbridgeStakingAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DosbridgeStakingAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DosbridgeStakingAddressUpdated represents a StakingAddressUpdated event raised by the Dosbridge contract.
type DosbridgeStakingAddressUpdated struct {
	PreviousStaking common.Address
	NewStaking      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingAddressUpdated is a free log retrieval operation binding the contract event 0x03fbfa1263b46c684780f3c24be11a2e189a59bedf0e316a7eae861cc769eb4f.
//
// Solidity: e StakingAddressUpdated(previousStaking address, newStaking address)
func (_Dosbridge *DosbridgeFilterer) FilterStakingAddressUpdated(opts *bind.FilterOpts) (*DosbridgeStakingAddressUpdatedIterator, error) {

	logs, sub, err := _Dosbridge.contract.FilterLogs(opts, "StakingAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DosbridgeStakingAddressUpdatedIterator{contract: _Dosbridge.contract, event: "StakingAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchStakingAddressUpdated is a free log subscription operation binding the contract event 0x03fbfa1263b46c684780f3c24be11a2e189a59bedf0e316a7eae861cc769eb4f.
//
// Solidity: e StakingAddressUpdated(previousStaking address, newStaking address)
func (_Dosbridge *DosbridgeFilterer) WatchStakingAddressUpdated(opts *bind.WatchOpts, sink chan<- *DosbridgeStakingAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _Dosbridge.contract.WatchLogs(opts, "StakingAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DosbridgeStakingAddressUpdated)
				if err := _Dosbridge.contract.UnpackLog(event, "StakingAddressUpdated", log); err != nil {
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
