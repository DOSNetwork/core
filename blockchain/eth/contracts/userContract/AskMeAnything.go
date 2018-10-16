// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package userContract

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

// AskMeAnythingABI is the input ABI used to generate the binding from.
const AskMeAnythingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"last_queried_url\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_mode\",\"type\":\"bool\"}],\"name\":\"setQueryMode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"response\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"url\",\"type\":\"string\"}],\"name\":\"AMA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"query_id\",\"type\":\"uint256\"},{\"name\":\"result\",\"type\":\"string\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repeated_call\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_timeout\",\"type\":\"uint256\"}],\"name\":\"setTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previous_timeout\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"new_timeout\",\"type\":\"uint256\"}],\"name\":\"SetTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"query_id\",\"type\":\"uint256\"}],\"name\":\"CallbackReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// AskMeAnythingBin is the compiled bytecode used for deploying new contracts.
const AskMeAnythingBin = `0x608060405260028054600160a060020a031990811673593bce0faf2d3d0863324fffb1a1c988cd22d5e5179091556005805460ff19169055601c6006556000805490911633179055610c6b806100566000396000f3006080604052600436106100b95763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633c4f7a0b81146100be57806357ae678b1461014857806370dea79a14610164578063715018a61461018b5780637a7f01a7146101a05780638da5cb5b146101b55780638f32d59b146101e6578063aa6a1a4d1461020f578063b30f87ec14610268578063be17a5821461028c578063c58a34cc146102a1578063f2fde38b146102b9575b600080fd5b3480156100ca57600080fd5b506100d36102da565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010d5781810151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015457600080fd5b506101626004351515610368565b005b34801561017057600080fd5b5061017961038e565b60408051918252519081900360200190f35b34801561019757600080fd5b50610162610394565b3480156101ac57600080fd5b506100d36103fc565b3480156101c157600080fd5b506101ca610457565b60408051600160a060020a039092168252519081900360200190f35b3480156101f257600080fd5b506101fb610467565b604080519115158252519081900360200190f35b34801561021b57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101629436949293602493928401919081908401838280828437509497506104789650505050505050565b34801561027457600080fd5b50610162600480359060248035908101910135610562565b34801561029857600080fd5b506101fb610782565b3480156102ad57600080fd5b5061016260043561078b565b3480156102c557600080fd5b50610162600160a060020a03600435166107e0565b6007805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103605780601f1061033557610100808354040283529160200191610360565b820191906000526020600020905b81548152906001019060200180831161034357829003601f168201915b505050505081565b610370610467565b151561037b57600080fd5b6005805460ff1916911515919091179055565b60065481565b61039c610467565b15156103a757600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103605780601f1061033557610100808354040283529160200191610360565b600054600160a060020a03165b90565b600054600160a060020a0316331490565b805160009061048e906007906020850190610b39565b506104d16006546040805190810160405280600381526020017f4150490000000000000000000000000000000000000000000000000000000000815250846107ff565b905080156104f7576000818152600460205260409020805460ff1916600117905561055e565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f496e76616c69642071756572792069642e000000000000000000000000000000604482015290519081900360640190fd5b5050565b61056a610a2c565b600160a060020a0316331461060657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f556e61757468656e7469636174656420726573706f6e73652066726f6d206e6f60448201527f6e2d444f532e0000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008381526004602052604090205460ff16151561068557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f526573706f6e7365207769746820696e76616c69642071756572792069642100604482015290519081900360640190fd5b6040805184815290517f69e02e44616e2da5b8f7ec0cad80ae44a7d1cb069604051d73b6cc3a24330c8d9181900360200190a16106c460038383610bb7565b506000838152600460205260409020805460ff1916905560055460ff161561077d5760078054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815261077d93909290918301828280156107735780601f1061074857610100808354040283529160200191610773565b820191906000526020600020905b81548152906001019060200180831161075657829003601f168201915b5050505050610478565b505050565b60055460ff1681565b610793610467565b151561079e57600080fd5b600654604080519182526020820183905280517f9aa0de0157c9133b911d2d811f590159622cea28cefe31505c203c828799da589281900390910190a1600655565b6107e8610467565b15156107f357600080fd5b6107fc81610abc565b50565b600254604080517f43a73d9a0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a0316916343a73d9a91600480830192602092919082900301818787803b15801561085e57600080fd5b505af1158015610872573d6000803e3d6000fd5b505050506040513d602081101561088857600080fd5b50516001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392831617908190556040517f482edfaa00000000000000000000000000000000000000000000000000000000815230600482018181524360248401819052604484018a905260a060648501908152895160a48601528951959096169563482edfaa95939491938b938b938b939192608482019160c40190602087019080838360005b8381101561094757818101518382015260200161092f565b50505050905090810190601f1680156109745780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156109a757818101518382015260200161098f565b50505050905090810190601f1680156109d45780820380516001836020036101000a031916815260200191505b50975050505050505050602060405180830381600087803b1580156109f857600080fd5b505af1158015610a0c573d6000803e3d6000fd5b505050506040513d6020811015610a2257600080fd5b5051949350505050565b600254604080517f43a73d9a0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a0316916343a73d9a91600480830192602092919082900301818787803b158015610a8b57600080fd5b505af1158015610a9f573d6000803e3d6000fd5b505050506040513d6020811015610ab557600080fd5b5051905090565b600160a060020a0381161515610ad157600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b7a57805160ff1916838001178555610ba7565b82800160010185558215610ba7579182015b82811115610ba7578251825591602001919060010190610b8c565b50610bb3929150610c25565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610bf85782800160ff19823516178555610ba7565b82800160010185558215610ba7579182015b82811115610ba7578235825591602001919060010190610c0a565b61046491905b80821115610bb35760008155600101610c2b5600a165627a7a72305820cd53766c6547087e3cc933e8964240a1143cc38402a7a91b8a9096c5ce05d69a0029`

// DeployAskMeAnything deploys a new Ethereum contract, binding an instance of AskMeAnything to it.
func DeployAskMeAnything(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AskMeAnything, error) {
	parsed, err := abi.JSON(strings.NewReader(AskMeAnythingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AskMeAnythingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AskMeAnything{AskMeAnythingCaller: AskMeAnythingCaller{contract: contract}, AskMeAnythingTransactor: AskMeAnythingTransactor{contract: contract}, AskMeAnythingFilterer: AskMeAnythingFilterer{contract: contract}}, nil
}

// AskMeAnything is an auto generated Go binding around an Ethereum contract.
type AskMeAnything struct {
	AskMeAnythingCaller     // Read-only binding to the contract
	AskMeAnythingTransactor // Write-only binding to the contract
	AskMeAnythingFilterer   // Log filterer for contract events
}

// AskMeAnythingCaller is an auto generated read-only Go binding around an Ethereum contract.
type AskMeAnythingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AskMeAnythingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AskMeAnythingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AskMeAnythingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AskMeAnythingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AskMeAnythingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AskMeAnythingSession struct {
	Contract     *AskMeAnything    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AskMeAnythingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AskMeAnythingCallerSession struct {
	Contract *AskMeAnythingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AskMeAnythingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AskMeAnythingTransactorSession struct {
	Contract     *AskMeAnythingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AskMeAnythingRaw is an auto generated low-level Go binding around an Ethereum contract.
type AskMeAnythingRaw struct {
	Contract *AskMeAnything // Generic contract binding to access the raw methods on
}

// AskMeAnythingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AskMeAnythingCallerRaw struct {
	Contract *AskMeAnythingCaller // Generic read-only contract binding to access the raw methods on
}

// AskMeAnythingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AskMeAnythingTransactorRaw struct {
	Contract *AskMeAnythingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAskMeAnything creates a new instance of AskMeAnything, bound to a specific deployed contract.
func NewAskMeAnything(address common.Address, backend bind.ContractBackend) (*AskMeAnything, error) {
	contract, err := bindAskMeAnything(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AskMeAnything{AskMeAnythingCaller: AskMeAnythingCaller{contract: contract}, AskMeAnythingTransactor: AskMeAnythingTransactor{contract: contract}, AskMeAnythingFilterer: AskMeAnythingFilterer{contract: contract}}, nil
}

// NewAskMeAnythingCaller creates a new read-only instance of AskMeAnything, bound to a specific deployed contract.
func NewAskMeAnythingCaller(address common.Address, caller bind.ContractCaller) (*AskMeAnythingCaller, error) {
	contract, err := bindAskMeAnything(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingCaller{contract: contract}, nil
}

// NewAskMeAnythingTransactor creates a new write-only instance of AskMeAnything, bound to a specific deployed contract.
func NewAskMeAnythingTransactor(address common.Address, transactor bind.ContractTransactor) (*AskMeAnythingTransactor, error) {
	contract, err := bindAskMeAnything(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingTransactor{contract: contract}, nil
}

// NewAskMeAnythingFilterer creates a new log filterer instance of AskMeAnything, bound to a specific deployed contract.
func NewAskMeAnythingFilterer(address common.Address, filterer bind.ContractFilterer) (*AskMeAnythingFilterer, error) {
	contract, err := bindAskMeAnything(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingFilterer{contract: contract}, nil
}

// bindAskMeAnything binds a generic wrapper to an already deployed contract.
func bindAskMeAnything(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AskMeAnythingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AskMeAnything *AskMeAnythingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AskMeAnything.Contract.AskMeAnythingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AskMeAnything *AskMeAnythingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AskMeAnything.Contract.AskMeAnythingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AskMeAnything *AskMeAnythingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AskMeAnything.Contract.AskMeAnythingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AskMeAnything *AskMeAnythingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AskMeAnything.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AskMeAnything *AskMeAnythingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AskMeAnything.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AskMeAnything *AskMeAnythingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AskMeAnything.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AskMeAnything *AskMeAnythingCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AskMeAnything *AskMeAnythingSession) IsOwner() (bool, error) {
	return _AskMeAnything.Contract.IsOwner(&_AskMeAnything.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AskMeAnything *AskMeAnythingCallerSession) IsOwner() (bool, error) {
	return _AskMeAnything.Contract.IsOwner(&_AskMeAnything.CallOpts)
}

// LastQueriedUrl is a free data retrieval call binding the contract method 0x3c4f7a0b.
//
// Solidity: function last_queried_url() constant returns(string)
func (_AskMeAnything *AskMeAnythingCaller) LastQueriedUrl(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "last_queried_url")
	return *ret0, err
}

// LastQueriedUrl is a free data retrieval call binding the contract method 0x3c4f7a0b.
//
// Solidity: function last_queried_url() constant returns(string)
func (_AskMeAnything *AskMeAnythingSession) LastQueriedUrl() (string, error) {
	return _AskMeAnything.Contract.LastQueriedUrl(&_AskMeAnything.CallOpts)
}

// LastQueriedUrl is a free data retrieval call binding the contract method 0x3c4f7a0b.
//
// Solidity: function last_queried_url() constant returns(string)
func (_AskMeAnything *AskMeAnythingCallerSession) LastQueriedUrl() (string, error) {
	return _AskMeAnything.Contract.LastQueriedUrl(&_AskMeAnything.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AskMeAnything *AskMeAnythingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AskMeAnything *AskMeAnythingSession) Owner() (common.Address, error) {
	return _AskMeAnything.Contract.Owner(&_AskMeAnything.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AskMeAnything *AskMeAnythingCallerSession) Owner() (common.Address, error) {
	return _AskMeAnything.Contract.Owner(&_AskMeAnything.CallOpts)
}

// RepeatedCall is a free data retrieval call binding the contract method 0xbe17a582.
//
// Solidity: function repeated_call() constant returns(bool)
func (_AskMeAnything *AskMeAnythingCaller) RepeatedCall(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "repeated_call")
	return *ret0, err
}

// RepeatedCall is a free data retrieval call binding the contract method 0xbe17a582.
//
// Solidity: function repeated_call() constant returns(bool)
func (_AskMeAnything *AskMeAnythingSession) RepeatedCall() (bool, error) {
	return _AskMeAnything.Contract.RepeatedCall(&_AskMeAnything.CallOpts)
}

// RepeatedCall is a free data retrieval call binding the contract method 0xbe17a582.
//
// Solidity: function repeated_call() constant returns(bool)
func (_AskMeAnything *AskMeAnythingCallerSession) RepeatedCall() (bool, error) {
	return _AskMeAnything.Contract.RepeatedCall(&_AskMeAnything.CallOpts)
}

// Response is a free data retrieval call binding the contract method 0x7a7f01a7.
//
// Solidity: function response() constant returns(string)
func (_AskMeAnything *AskMeAnythingCaller) Response(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "response")
	return *ret0, err
}

// Response is a free data retrieval call binding the contract method 0x7a7f01a7.
//
// Solidity: function response() constant returns(string)
func (_AskMeAnything *AskMeAnythingSession) Response() (string, error) {
	return _AskMeAnything.Contract.Response(&_AskMeAnything.CallOpts)
}

// Response is a free data retrieval call binding the contract method 0x7a7f01a7.
//
// Solidity: function response() constant returns(string)
func (_AskMeAnything *AskMeAnythingCallerSession) Response() (string, error) {
	return _AskMeAnything.Contract.Response(&_AskMeAnything.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_AskMeAnything *AskMeAnythingCaller) Timeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "timeout")
	return *ret0, err
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_AskMeAnything *AskMeAnythingSession) Timeout() (*big.Int, error) {
	return _AskMeAnything.Contract.Timeout(&_AskMeAnything.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_AskMeAnything *AskMeAnythingCallerSession) Timeout() (*big.Int, error) {
	return _AskMeAnything.Contract.Timeout(&_AskMeAnything.CallOpts)
}

// AMA is a paid mutator transaction binding the contract method 0xaa6a1a4d.
//
// Solidity: function AMA(url string) returns()
func (_AskMeAnything *AskMeAnythingTransactor) AMA(opts *bind.TransactOpts, url string) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "AMA", url)
}

// AMA is a paid mutator transaction binding the contract method 0xaa6a1a4d.
//
// Solidity: function AMA(url string) returns()
func (_AskMeAnything *AskMeAnythingSession) AMA(url string) (*types.Transaction, error) {
	return _AskMeAnything.Contract.AMA(&_AskMeAnything.TransactOpts, url)
}

// AMA is a paid mutator transaction binding the contract method 0xaa6a1a4d.
//
// Solidity: function AMA(url string) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) AMA(url string) (*types.Transaction, error) {
	return _AskMeAnything.Contract.AMA(&_AskMeAnything.TransactOpts, url)
}

// Callback_ is a paid mutator transaction binding the contract method 0xb30f87ec.
//
// Solidity: function __callback__(query_id uint256, result string) returns()
func (_AskMeAnything *AskMeAnythingTransactor) Callback_(opts *bind.TransactOpts, query_id *big.Int, result string) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "__callback__", query_id, result)
}

// Callback_ is a paid mutator transaction binding the contract method 0xb30f87ec.
//
// Solidity: function __callback__(query_id uint256, result string) returns()
func (_AskMeAnything *AskMeAnythingSession) Callback_(query_id *big.Int, result string) (*types.Transaction, error) {
	return _AskMeAnything.Contract.Callback_(&_AskMeAnything.TransactOpts, query_id, result)
}

// Callback_ is a paid mutator transaction binding the contract method 0xb30f87ec.
//
// Solidity: function __callback__(query_id uint256, result string) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) Callback_(query_id *big.Int, result string) (*types.Transaction, error) {
	return _AskMeAnything.Contract.Callback_(&_AskMeAnything.TransactOpts, query_id, result)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AskMeAnything *AskMeAnythingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AskMeAnything *AskMeAnythingSession) RenounceOwnership() (*types.Transaction, error) {
	return _AskMeAnything.Contract.RenounceOwnership(&_AskMeAnything.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AskMeAnything.Contract.RenounceOwnership(&_AskMeAnything.TransactOpts)
}

// SetQueryMode is a paid mutator transaction binding the contract method 0x57ae678b.
//
// Solidity: function setQueryMode(new_mode bool) returns()
func (_AskMeAnything *AskMeAnythingTransactor) SetQueryMode(opts *bind.TransactOpts, new_mode bool) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "setQueryMode", new_mode)
}

// SetQueryMode is a paid mutator transaction binding the contract method 0x57ae678b.
//
// Solidity: function setQueryMode(new_mode bool) returns()
func (_AskMeAnything *AskMeAnythingSession) SetQueryMode(new_mode bool) (*types.Transaction, error) {
	return _AskMeAnything.Contract.SetQueryMode(&_AskMeAnything.TransactOpts, new_mode)
}

// SetQueryMode is a paid mutator transaction binding the contract method 0x57ae678b.
//
// Solidity: function setQueryMode(new_mode bool) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) SetQueryMode(new_mode bool) (*types.Transaction, error) {
	return _AskMeAnything.Contract.SetQueryMode(&_AskMeAnything.TransactOpts, new_mode)
}

// SetTimeout is a paid mutator transaction binding the contract method 0xc58a34cc.
//
// Solidity: function setTimeout(new_timeout uint256) returns()
func (_AskMeAnything *AskMeAnythingTransactor) SetTimeout(opts *bind.TransactOpts, new_timeout *big.Int) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "setTimeout", new_timeout)
}

// SetTimeout is a paid mutator transaction binding the contract method 0xc58a34cc.
//
// Solidity: function setTimeout(new_timeout uint256) returns()
func (_AskMeAnything *AskMeAnythingSession) SetTimeout(new_timeout *big.Int) (*types.Transaction, error) {
	return _AskMeAnything.Contract.SetTimeout(&_AskMeAnything.TransactOpts, new_timeout)
}

// SetTimeout is a paid mutator transaction binding the contract method 0xc58a34cc.
//
// Solidity: function setTimeout(new_timeout uint256) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) SetTimeout(new_timeout *big.Int) (*types.Transaction, error) {
	return _AskMeAnything.Contract.SetTimeout(&_AskMeAnything.TransactOpts, new_timeout)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AskMeAnything *AskMeAnythingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AskMeAnything *AskMeAnythingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AskMeAnything.Contract.TransferOwnership(&_AskMeAnything.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AskMeAnything.Contract.TransferOwnership(&_AskMeAnything.TransactOpts, newOwner)
}

// AskMeAnythingCallbackReadyIterator is returned from FilterCallbackReady and is used to iterate over the raw logs and unpacked data for CallbackReady events raised by the AskMeAnything contract.
type AskMeAnythingCallbackReadyIterator struct {
	Event *AskMeAnythingCallbackReady // Event containing the contract specifics and raw log

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
func (it *AskMeAnythingCallbackReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AskMeAnythingCallbackReady)
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
		it.Event = new(AskMeAnythingCallbackReady)
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
func (it *AskMeAnythingCallbackReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AskMeAnythingCallbackReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AskMeAnythingCallbackReady represents a CallbackReady event raised by the AskMeAnything contract.
type AskMeAnythingCallbackReady struct {
	QueryId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCallbackReady is a free log retrieval operation binding the contract event 0x69e02e44616e2da5b8f7ec0cad80ae44a7d1cb069604051d73b6cc3a24330c8d.
//
// Solidity: e CallbackReady(query_id uint256)
func (_AskMeAnything *AskMeAnythingFilterer) FilterCallbackReady(opts *bind.FilterOpts) (*AskMeAnythingCallbackReadyIterator, error) {

	logs, sub, err := _AskMeAnything.contract.FilterLogs(opts, "CallbackReady")
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingCallbackReadyIterator{contract: _AskMeAnything.contract, event: "CallbackReady", logs: logs, sub: sub}, nil
}

// WatchCallbackReady is a free log subscription operation binding the contract event 0x69e02e44616e2da5b8f7ec0cad80ae44a7d1cb069604051d73b6cc3a24330c8d.
//
// Solidity: e CallbackReady(query_id uint256)
func (_AskMeAnything *AskMeAnythingFilterer) WatchCallbackReady(opts *bind.WatchOpts, sink chan<- *AskMeAnythingCallbackReady) (event.Subscription, error) {

	logs, sub, err := _AskMeAnything.contract.WatchLogs(opts, "CallbackReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AskMeAnythingCallbackReady)
				if err := _AskMeAnything.contract.UnpackLog(event, "CallbackReady", log); err != nil {
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

// AskMeAnythingOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the AskMeAnything contract.
type AskMeAnythingOwnershipRenouncedIterator struct {
	Event *AskMeAnythingOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *AskMeAnythingOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AskMeAnythingOwnershipRenounced)
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
		it.Event = new(AskMeAnythingOwnershipRenounced)
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
func (it *AskMeAnythingOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AskMeAnythingOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AskMeAnythingOwnershipRenounced represents a OwnershipRenounced event raised by the AskMeAnything contract.
type AskMeAnythingOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_AskMeAnything *AskMeAnythingFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*AskMeAnythingOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _AskMeAnything.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingOwnershipRenouncedIterator{contract: _AskMeAnything.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_AskMeAnything *AskMeAnythingFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *AskMeAnythingOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _AskMeAnything.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AskMeAnythingOwnershipRenounced)
				if err := _AskMeAnything.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// AskMeAnythingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AskMeAnything contract.
type AskMeAnythingOwnershipTransferredIterator struct {
	Event *AskMeAnythingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AskMeAnythingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AskMeAnythingOwnershipTransferred)
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
		it.Event = new(AskMeAnythingOwnershipTransferred)
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
func (it *AskMeAnythingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AskMeAnythingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AskMeAnythingOwnershipTransferred represents a OwnershipTransferred event raised by the AskMeAnything contract.
type AskMeAnythingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AskMeAnything *AskMeAnythingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AskMeAnythingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AskMeAnything.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingOwnershipTransferredIterator{contract: _AskMeAnything.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AskMeAnything *AskMeAnythingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AskMeAnythingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AskMeAnything.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AskMeAnythingOwnershipTransferred)
				if err := _AskMeAnything.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// AskMeAnythingSetTimeoutIterator is returned from FilterSetTimeout and is used to iterate over the raw logs and unpacked data for SetTimeout events raised by the AskMeAnything contract.
type AskMeAnythingSetTimeoutIterator struct {
	Event *AskMeAnythingSetTimeout // Event containing the contract specifics and raw log

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
func (it *AskMeAnythingSetTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AskMeAnythingSetTimeout)
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
		it.Event = new(AskMeAnythingSetTimeout)
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
func (it *AskMeAnythingSetTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AskMeAnythingSetTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AskMeAnythingSetTimeout represents a SetTimeout event raised by the AskMeAnything contract.
type AskMeAnythingSetTimeout struct {
	PreviousTimeout *big.Int
	NewTimeout      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetTimeout is a free log retrieval operation binding the contract event 0x9aa0de0157c9133b911d2d811f590159622cea28cefe31505c203c828799da58.
//
// Solidity: e SetTimeout(previous_timeout uint256, new_timeout uint256)
func (_AskMeAnything *AskMeAnythingFilterer) FilterSetTimeout(opts *bind.FilterOpts) (*AskMeAnythingSetTimeoutIterator, error) {

	logs, sub, err := _AskMeAnything.contract.FilterLogs(opts, "SetTimeout")
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingSetTimeoutIterator{contract: _AskMeAnything.contract, event: "SetTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTimeout is a free log subscription operation binding the contract event 0x9aa0de0157c9133b911d2d811f590159622cea28cefe31505c203c828799da58.
//
// Solidity: e SetTimeout(previous_timeout uint256, new_timeout uint256)
func (_AskMeAnything *AskMeAnythingFilterer) WatchSetTimeout(opts *bind.WatchOpts, sink chan<- *AskMeAnythingSetTimeout) (event.Subscription, error) {

	logs, sub, err := _AskMeAnything.contract.WatchLogs(opts, "SetTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AskMeAnythingSetTimeout)
				if err := _AskMeAnything.contract.UnpackLog(event, "SetTimeout", log); err != nil {
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

// DOSAddressBridgeInterfaceABI is the input ABI used to generate the binding from.
const DOSAddressBridgeInterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DOSAddressBridgeInterfaceBin is the compiled bytecode used for deploying new contracts.
const DOSAddressBridgeInterfaceBin = `0x`

// DeployDOSAddressBridgeInterface deploys a new Ethereum contract, binding an instance of DOSAddressBridgeInterface to it.
func DeployDOSAddressBridgeInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DOSAddressBridgeInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSAddressBridgeInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DOSAddressBridgeInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DOSAddressBridgeInterface{DOSAddressBridgeInterfaceCaller: DOSAddressBridgeInterfaceCaller{contract: contract}, DOSAddressBridgeInterfaceTransactor: DOSAddressBridgeInterfaceTransactor{contract: contract}, DOSAddressBridgeInterfaceFilterer: DOSAddressBridgeInterfaceFilterer{contract: contract}}, nil
}

// DOSAddressBridgeInterface is an auto generated Go binding around an Ethereum contract.
type DOSAddressBridgeInterface struct {
	DOSAddressBridgeInterfaceCaller     // Read-only binding to the contract
	DOSAddressBridgeInterfaceTransactor // Write-only binding to the contract
	DOSAddressBridgeInterfaceFilterer   // Log filterer for contract events
}

// DOSAddressBridgeInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DOSAddressBridgeInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSAddressBridgeInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DOSAddressBridgeInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSAddressBridgeInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DOSAddressBridgeInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSAddressBridgeInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DOSAddressBridgeInterfaceSession struct {
	Contract     *DOSAddressBridgeInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DOSAddressBridgeInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DOSAddressBridgeInterfaceCallerSession struct {
	Contract *DOSAddressBridgeInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// DOSAddressBridgeInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DOSAddressBridgeInterfaceTransactorSession struct {
	Contract     *DOSAddressBridgeInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// DOSAddressBridgeInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DOSAddressBridgeInterfaceRaw struct {
	Contract *DOSAddressBridgeInterface // Generic contract binding to access the raw methods on
}

// DOSAddressBridgeInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DOSAddressBridgeInterfaceCallerRaw struct {
	Contract *DOSAddressBridgeInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DOSAddressBridgeInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DOSAddressBridgeInterfaceTransactorRaw struct {
	Contract *DOSAddressBridgeInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDOSAddressBridgeInterface creates a new instance of DOSAddressBridgeInterface, bound to a specific deployed contract.
func NewDOSAddressBridgeInterface(address common.Address, backend bind.ContractBackend) (*DOSAddressBridgeInterface, error) {
	contract, err := bindDOSAddressBridgeInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DOSAddressBridgeInterface{DOSAddressBridgeInterfaceCaller: DOSAddressBridgeInterfaceCaller{contract: contract}, DOSAddressBridgeInterfaceTransactor: DOSAddressBridgeInterfaceTransactor{contract: contract}, DOSAddressBridgeInterfaceFilterer: DOSAddressBridgeInterfaceFilterer{contract: contract}}, nil
}

// NewDOSAddressBridgeInterfaceCaller creates a new read-only instance of DOSAddressBridgeInterface, bound to a specific deployed contract.
func NewDOSAddressBridgeInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DOSAddressBridgeInterfaceCaller, error) {
	contract, err := bindDOSAddressBridgeInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DOSAddressBridgeInterfaceCaller{contract: contract}, nil
}

// NewDOSAddressBridgeInterfaceTransactor creates a new write-only instance of DOSAddressBridgeInterface, bound to a specific deployed contract.
func NewDOSAddressBridgeInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DOSAddressBridgeInterfaceTransactor, error) {
	contract, err := bindDOSAddressBridgeInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DOSAddressBridgeInterfaceTransactor{contract: contract}, nil
}

// NewDOSAddressBridgeInterfaceFilterer creates a new log filterer instance of DOSAddressBridgeInterface, bound to a specific deployed contract.
func NewDOSAddressBridgeInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DOSAddressBridgeInterfaceFilterer, error) {
	contract, err := bindDOSAddressBridgeInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DOSAddressBridgeInterfaceFilterer{contract: contract}, nil
}

// bindDOSAddressBridgeInterface binds a generic wrapper to an already deployed contract.
func bindDOSAddressBridgeInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSAddressBridgeInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSAddressBridgeInterface.Contract.DOSAddressBridgeInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSAddressBridgeInterface.Contract.DOSAddressBridgeInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSAddressBridgeInterface.Contract.DOSAddressBridgeInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSAddressBridgeInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSAddressBridgeInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSAddressBridgeInterface.Contract.contract.Transact(opts, method, params...)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceCaller) GetProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DOSAddressBridgeInterface.contract.Call(opts, out, "getProxyAddress")
	return *ret0, err
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceSession) GetProxyAddress() (common.Address, error) {
	return _DOSAddressBridgeInterface.Contract.GetProxyAddress(&_DOSAddressBridgeInterface.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() constant returns(address)
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceCallerSession) GetProxyAddress() (common.Address, error) {
	return _DOSAddressBridgeInterface.Contract.GetProxyAddress(&_DOSAddressBridgeInterface.CallOpts)
}

// DOSOnChainSDKABI is the input ABI used to generate the binding from.
const DOSOnChainSDKABI = "[]"

// DOSOnChainSDKBin is the compiled bytecode used for deploying new contracts.
const DOSOnChainSDKBin = `0x608060405260018054600160a060020a03191673593bce0faf2d3d0863324fffb1a1c988cd22d5e5179055348015603557600080fd5b5060358060436000396000f3006080604052600080fd00a165627a7a72305820adbca4a9308ab55021e4b841aaecb0f0abaab89341827889571487670715b6590029`

// DeployDOSOnChainSDK deploys a new Ethereum contract, binding an instance of DOSOnChainSDK to it.
func DeployDOSOnChainSDK(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DOSOnChainSDK, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSOnChainSDKABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DOSOnChainSDKBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DOSOnChainSDK{DOSOnChainSDKCaller: DOSOnChainSDKCaller{contract: contract}, DOSOnChainSDKTransactor: DOSOnChainSDKTransactor{contract: contract}, DOSOnChainSDKFilterer: DOSOnChainSDKFilterer{contract: contract}}, nil
}

// DOSOnChainSDK is an auto generated Go binding around an Ethereum contract.
type DOSOnChainSDK struct {
	DOSOnChainSDKCaller     // Read-only binding to the contract
	DOSOnChainSDKTransactor // Write-only binding to the contract
	DOSOnChainSDKFilterer   // Log filterer for contract events
}

// DOSOnChainSDKCaller is an auto generated read-only Go binding around an Ethereum contract.
type DOSOnChainSDKCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSOnChainSDKTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DOSOnChainSDKTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSOnChainSDKFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DOSOnChainSDKFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSOnChainSDKSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DOSOnChainSDKSession struct {
	Contract     *DOSOnChainSDK    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DOSOnChainSDKCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DOSOnChainSDKCallerSession struct {
	Contract *DOSOnChainSDKCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DOSOnChainSDKTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DOSOnChainSDKTransactorSession struct {
	Contract     *DOSOnChainSDKTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DOSOnChainSDKRaw is an auto generated low-level Go binding around an Ethereum contract.
type DOSOnChainSDKRaw struct {
	Contract *DOSOnChainSDK // Generic contract binding to access the raw methods on
}

// DOSOnChainSDKCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DOSOnChainSDKCallerRaw struct {
	Contract *DOSOnChainSDKCaller // Generic read-only contract binding to access the raw methods on
}

// DOSOnChainSDKTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DOSOnChainSDKTransactorRaw struct {
	Contract *DOSOnChainSDKTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDOSOnChainSDK creates a new instance of DOSOnChainSDK, bound to a specific deployed contract.
func NewDOSOnChainSDK(address common.Address, backend bind.ContractBackend) (*DOSOnChainSDK, error) {
	contract, err := bindDOSOnChainSDK(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DOSOnChainSDK{DOSOnChainSDKCaller: DOSOnChainSDKCaller{contract: contract}, DOSOnChainSDKTransactor: DOSOnChainSDKTransactor{contract: contract}, DOSOnChainSDKFilterer: DOSOnChainSDKFilterer{contract: contract}}, nil
}

// NewDOSOnChainSDKCaller creates a new read-only instance of DOSOnChainSDK, bound to a specific deployed contract.
func NewDOSOnChainSDKCaller(address common.Address, caller bind.ContractCaller) (*DOSOnChainSDKCaller, error) {
	contract, err := bindDOSOnChainSDK(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DOSOnChainSDKCaller{contract: contract}, nil
}

// NewDOSOnChainSDKTransactor creates a new write-only instance of DOSOnChainSDK, bound to a specific deployed contract.
func NewDOSOnChainSDKTransactor(address common.Address, transactor bind.ContractTransactor) (*DOSOnChainSDKTransactor, error) {
	contract, err := bindDOSOnChainSDK(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DOSOnChainSDKTransactor{contract: contract}, nil
}

// NewDOSOnChainSDKFilterer creates a new log filterer instance of DOSOnChainSDK, bound to a specific deployed contract.
func NewDOSOnChainSDKFilterer(address common.Address, filterer bind.ContractFilterer) (*DOSOnChainSDKFilterer, error) {
	contract, err := bindDOSOnChainSDK(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DOSOnChainSDKFilterer{contract: contract}, nil
}

// bindDOSOnChainSDK binds a generic wrapper to an already deployed contract.
func bindDOSOnChainSDK(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSOnChainSDKABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSOnChainSDK *DOSOnChainSDKRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSOnChainSDK.Contract.DOSOnChainSDKCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSOnChainSDK *DOSOnChainSDKRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.DOSOnChainSDKTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSOnChainSDK *DOSOnChainSDKRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.DOSOnChainSDKTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSOnChainSDK *DOSOnChainSDKCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSOnChainSDK.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSOnChainSDK *DOSOnChainSDKTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSOnChainSDK *DOSOnChainSDKTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.contract.Transact(opts, method, params...)
}

// DOSProxyInterfaceABI is the input ABI used to generate the binding from.
const DOSProxyInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DOSProxyInterfaceBin is the compiled bytecode used for deploying new contracts.
const DOSProxyInterfaceBin = `0x`

// DeployDOSProxyInterface deploys a new Ethereum contract, binding an instance of DOSProxyInterface to it.
func DeployDOSProxyInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DOSProxyInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSProxyInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DOSProxyInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DOSProxyInterface{DOSProxyInterfaceCaller: DOSProxyInterfaceCaller{contract: contract}, DOSProxyInterfaceTransactor: DOSProxyInterfaceTransactor{contract: contract}, DOSProxyInterfaceFilterer: DOSProxyInterfaceFilterer{contract: contract}}, nil
}

// DOSProxyInterface is an auto generated Go binding around an Ethereum contract.
type DOSProxyInterface struct {
	DOSProxyInterfaceCaller     // Read-only binding to the contract
	DOSProxyInterfaceTransactor // Write-only binding to the contract
	DOSProxyInterfaceFilterer   // Log filterer for contract events
}

// DOSProxyInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DOSProxyInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSProxyInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DOSProxyInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSProxyInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DOSProxyInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSProxyInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DOSProxyInterfaceSession struct {
	Contract     *DOSProxyInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DOSProxyInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DOSProxyInterfaceCallerSession struct {
	Contract *DOSProxyInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DOSProxyInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DOSProxyInterfaceTransactorSession struct {
	Contract     *DOSProxyInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DOSProxyInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DOSProxyInterfaceRaw struct {
	Contract *DOSProxyInterface // Generic contract binding to access the raw methods on
}

// DOSProxyInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DOSProxyInterfaceCallerRaw struct {
	Contract *DOSProxyInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DOSProxyInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DOSProxyInterfaceTransactorRaw struct {
	Contract *DOSProxyInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDOSProxyInterface creates a new instance of DOSProxyInterface, bound to a specific deployed contract.
func NewDOSProxyInterface(address common.Address, backend bind.ContractBackend) (*DOSProxyInterface, error) {
	contract, err := bindDOSProxyInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DOSProxyInterface{DOSProxyInterfaceCaller: DOSProxyInterfaceCaller{contract: contract}, DOSProxyInterfaceTransactor: DOSProxyInterfaceTransactor{contract: contract}, DOSProxyInterfaceFilterer: DOSProxyInterfaceFilterer{contract: contract}}, nil
}

// NewDOSProxyInterfaceCaller creates a new read-only instance of DOSProxyInterface, bound to a specific deployed contract.
func NewDOSProxyInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DOSProxyInterfaceCaller, error) {
	contract, err := bindDOSProxyInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DOSProxyInterfaceCaller{contract: contract}, nil
}

// NewDOSProxyInterfaceTransactor creates a new write-only instance of DOSProxyInterface, bound to a specific deployed contract.
func NewDOSProxyInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DOSProxyInterfaceTransactor, error) {
	contract, err := bindDOSProxyInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DOSProxyInterfaceTransactor{contract: contract}, nil
}

// NewDOSProxyInterfaceFilterer creates a new log filterer instance of DOSProxyInterface, bound to a specific deployed contract.
func NewDOSProxyInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DOSProxyInterfaceFilterer, error) {
	contract, err := bindDOSProxyInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DOSProxyInterfaceFilterer{contract: contract}, nil
}

// bindDOSProxyInterface binds a generic wrapper to an already deployed contract.
func bindDOSProxyInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSProxyInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSProxyInterface *DOSProxyInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSProxyInterface.Contract.DOSProxyInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSProxyInterface *DOSProxyInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.DOSProxyInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSProxyInterface *DOSProxyInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.DOSProxyInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSProxyInterface *DOSProxyInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSProxyInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSProxyInterface *DOSProxyInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSProxyInterface *DOSProxyInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.contract.Transact(opts, method, params...)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query( address,  uint256,  uint256,  string,  string) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceTransactor) Query(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 string, arg4 string) (*types.Transaction, error) {
	return _DOSProxyInterface.contract.Transact(opts, "query", arg0, arg1, arg2, arg3, arg4)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query( address,  uint256,  uint256,  string,  string) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceSession) Query(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 string, arg4 string) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.Query(&_DOSProxyInterface.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query( address,  uint256,  uint256,  string,  string) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceTransactorSession) Query(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 string, arg4 string) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.Query(&_DOSProxyInterface.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610248806100326000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a681146100665780638da5cb5b1461007d5780638f32d59b146100ae578063f2fde38b146100d7575b600080fd5b34801561007257600080fd5b5061007b6100f8565b005b34801561008957600080fd5b50610092610160565b60408051600160a060020a039092168252519081900360200190f35b3480156100ba57600080fd5b506100c361016f565b604080519115158252519081900360200190f35b3480156100e357600080fd5b5061007b600160a060020a0360043516610180565b61010061016f565b151561010b57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031690565b600054600160a060020a0316331490565b61018861016f565b151561019357600080fd5b61019c8161019f565b50565b600160a060020a03811615156101b457600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058206a6ad552b7b356387ae1152ba958dae8c25fb447073478e70967911d38aaef410029`

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
