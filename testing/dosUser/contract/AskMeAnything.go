// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dosUser

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

// AskMeAnythingABI is the input ABI used to generate the binding from.
const AskMeAnythingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"generatedRandom\",\"type\":\"uint256\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestFastRandom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastQueriedUrl\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMode\",\"type\":\"bool\"}],\"name\":\"setQueryMode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"random\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"queryId\",\"type\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"response\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repeatedCall\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRequestedRandom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"setTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestSafeRandom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastQueriedSelector\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"selector\",\"type\":\"string\"}],\"name\":\"AMA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousTimeout\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"SetTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"queryId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"string\"}],\"name\":\"QueryResponseReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"succ\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"generatedRandom\",\"type\":\"uint256\"}],\"name\":\"RandomReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// AskMeAnythingBin is the compiled bytecode used for deploying new contracts.
const AskMeAnythingBin = `0x608060405260028054600160a060020a03199081167314b03cf8ed30e4c58af3dbb6e0a3d1465df18506179091556006805460ff19169055601c6007556000805490911633179055611230806100566000396000f3006080604052600436106100fb5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318a1908d811461010057806335cd49e71461011d57806349b03ca01461013257806357ae678b146101bc5780635ec01e4d146101d65780636d112977146101fd57806370dea79a14610221578063715018a6146102365780637a7f01a71461024b5780638da5cb5b146102605780638f32d59b14610291578063ab8c1bad146102ba578063b8cf904e146102cf578063c58a34cc146102e4578063e1bcba05146102fc578063e3abd8b614610311578063f11ea3f814610326578063f2fde38b146103bd575b600080fd5b34801561010c57600080fd5b5061011b6004356024356103de565b005b34801561012957600080fd5b5061011b610554565b34801561013e57600080fd5b5061014761056a565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610181578181015183820152602001610169565b50505050905090810190601f1680156101ae5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101c857600080fd5b5061011b60043515156105f8565b3480156101e257600080fd5b506101eb61061e565b60408051918252519081900360200190f35b34801561020957600080fd5b5061011b600480359060248035908101910135610624565b34801561022d57600080fd5b506101eb61094e565b34801561024257600080fd5b5061011b610954565b34801561025757600080fd5b506101476109bc565b34801561026c57600080fd5b50610275610a17565b60408051600160a060020a039092168252519081900360200190f35b34801561029d57600080fd5b506102a6610a27565b604080519115158252519081900360200190f35b3480156102c657600080fd5b506102a6610a38565b3480156102db57600080fd5b506101eb610a41565b3480156102f057600080fd5b5061011b600435610a47565b34801561030857600080fd5b5061011b610a9c565b34801561031d57600080fd5b50610147610b0d565b34801561033257600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261011b94369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a999881019791965091820194509250829150840183828082843750949750610b689650505050505050565b3480156103c957600080fd5b5061011b600160a060020a0360043516610c5a565b816103e7610c79565b600160a060020a0316331461046c576040805160e560020a62461bcd02815260206004820152602660248201527f556e61757468656e7469636174656420726573706f6e73652066726f6d206e6f60448201527f6e2d444f532e0000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008181526005602052604090205460ff1615156104fa576040805160e560020a62461bcd02815260206004820152602160248201527f526573706f6e7365207769746820696e76616c6964207265717565737420696460448201527f2100000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6004829055604080518481526020810184905281517fd0ecc71f8b5af397da9123fd2bff63c544c04af5c6935935a7f81e14b84522f2929181900390910190a150506000908152600560205260409020805460ff19169055565b600454600a55610565600042610d09565b600455565b6008805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105f05780601f106105c5576101008083540402835291602001916105f0565b820191906000526020600020905b8154815290600101906020018083116105d357829003601f168201915b505050505081565b610600610a27565b151561060b57600080fd5b6006805460ff1916911515919091179055565b60045481565b8261062d610c79565b600160a060020a031633146106b2576040805160e560020a62461bcd02815260206004820152602660248201527f556e61757468656e7469636174656420726573706f6e73652066726f6d206e6f60448201527f6e2d444f532e0000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008181526005602052604090205460ff161515610740576040805160e560020a62461bcd02815260206004820152602160248201527f526573706f6e7365207769746820696e76616c6964207265717565737420696460448201527f2100000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b61074c600384846110fe565b50604080518581526020810182815260038054600260001961010060018416150201909116049383018490527ffe1788dc549f39fbcdb06fdab5937f2d19ce5fe5616fa3be9311928bd5dabccb9388939192906060830190849080156107f35780601f106107c8576101008083540402835291602001916107f3565b820191906000526020600020905b8154815290600101906020018083116107d657829003601f168201915b5050935050505060405180910390a16000848152600560205260409020805460ff1916905560065460ff16156109485760088054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815261094893909290918301828280156108b05780601f10610885576101008083540402835291602001916108b0565b820191906000526020600020905b81548152906001019060200180831161089357829003601f168201915b505060098054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281529550919350915083018282801561093e5780601f106109135761010080835404028352916020019161093e565b820191906000526020600020905b81548152906001019060200180831161092157829003601f168201915b5050505050610b68565b50505050565b60075481565b61095c610a27565b151561096757600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105f05780601f106105c5576101008083540402835291602001916105f0565b600054600160a060020a03165b90565b600054600160a060020a0316331490565b60065460ff1681565b600a5481565b610a4f610a27565b1515610a5a57600080fd5b600754604080519182526020820183905280517f9aa0de0157c9133b911d2d811f590159622cea28cefe31505c203c828799da589281900390910190a1600755565b600454600a556000610aaf600142610d09565b600081815260056020908152604091829020805460ff19166001908117909155825190815290810183905281519293507f4d0a0fbbf63d9b08d7e94ab32ae848f6fbf902858f874ac63a475db03268d982929081900390910190a150565b6009805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105f05780601f106105c5576101008083540402835291602001916105f0565b8151600090610b7e90600890602086019061117c565b508151610b9290600990602085019061117c565b50610ba06007548484610e5d565b90508015610c0557600081815260056020908152604091829020805460ff19166001908117909155825190815290810183905281517f4d0a0fbbf63d9b08d7e94ab32ae848f6fbf902858f874ac63a475db03268d982929181900390910190a1610c55565b6040805160e560020a62461bcd02815260206004820152601160248201527f496e76616c69642071756572792069642e000000000000000000000000000000604482015290519081900360640190fd5b505050565b610c62610a27565b1515610c6d57600080fd5b610c7681611081565b50565b600254604080517f43a73d9a0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a0316916343a73d9a91600480830192602092919082900301818787803b158015610cd857600080fd5b505af1158015610cec573d6000803e3d6000fd5b505050506040513d6020811015610d0257600080fd5b5051905090565b600254604080517f43a73d9a0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a0316916343a73d9a91600480830192602092919082900301818787803b158015610d6857600080fd5b505af1158015610d7c573d6000803e3d6000fd5b505050506040513d6020811015610d9257600080fd5b50516001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039283161790819055604080517f1bf8205700000000000000000000000000000000000000000000000000000000815230600482015260ff871660248201526044810186905290519190921691631bf820579160648083019260209291908290030181600087803b158015610e2a57600080fd5b505af1158015610e3e573d6000803e3d6000fd5b505050506040513d6020811015610e5457600080fd5b50519392505050565b600254604080517f43a73d9a0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a0316916343a73d9a91600480830192602092919082900301818787803b158015610ebc57600080fd5b505af1158015610ed0573d6000803e3d6000fd5b505050506040513d6020811015610ee657600080fd5b50516001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392831617908190556040517fb7fb8fd7000000000000000000000000000000000000000000000000000000008152306004820181815260248301899052608060448401908152885160848501528851949095169463b7fb8fd79492938a938a938a9390929091606482019160a40190602087019080838360005b83811015610f9d578181015183820152602001610f85565b50505050905090810190601f168015610fca5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610ffd578181015183820152602001610fe5565b50505050905090810190601f16801561102a5780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561104d57600080fd5b505af1158015611061573d6000803e3d6000fd5b505050506040513d602081101561107757600080fd5b5051949350505050565b600160a060020a038116151561109657600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061113f5782800160ff1982351617855561116c565b8280016001018555821561116c579182015b8281111561116c578235825591602001919060010190611151565b506111789291506111ea565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106111bd57805160ff191683800117855561116c565b8280016001018555821561116c579182015b8281111561116c5782518255916020019190600101906111cf565b610a2491905b8082111561117857600081556001016111f05600a165627a7a72305820031dd62c755d91e3ac8af3919ae02372a47b17402fe2c6020ccd7b785b7d903e0029`

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

// LastQueriedSelector is a free data retrieval call binding the contract method 0xe3abd8b6.
//
// Solidity: function lastQueriedSelector() constant returns(string)
func (_AskMeAnything *AskMeAnythingCaller) LastQueriedSelector(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "lastQueriedSelector")
	return *ret0, err
}

// LastQueriedSelector is a free data retrieval call binding the contract method 0xe3abd8b6.
//
// Solidity: function lastQueriedSelector() constant returns(string)
func (_AskMeAnything *AskMeAnythingSession) LastQueriedSelector() (string, error) {
	return _AskMeAnything.Contract.LastQueriedSelector(&_AskMeAnything.CallOpts)
}

// LastQueriedSelector is a free data retrieval call binding the contract method 0xe3abd8b6.
//
// Solidity: function lastQueriedSelector() constant returns(string)
func (_AskMeAnything *AskMeAnythingCallerSession) LastQueriedSelector() (string, error) {
	return _AskMeAnything.Contract.LastQueriedSelector(&_AskMeAnything.CallOpts)
}

// LastQueriedUrl is a free data retrieval call binding the contract method 0x49b03ca0.
//
// Solidity: function lastQueriedUrl() constant returns(string)
func (_AskMeAnything *AskMeAnythingCaller) LastQueriedUrl(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "lastQueriedUrl")
	return *ret0, err
}

// LastQueriedUrl is a free data retrieval call binding the contract method 0x49b03ca0.
//
// Solidity: function lastQueriedUrl() constant returns(string)
func (_AskMeAnything *AskMeAnythingSession) LastQueriedUrl() (string, error) {
	return _AskMeAnything.Contract.LastQueriedUrl(&_AskMeAnything.CallOpts)
}

// LastQueriedUrl is a free data retrieval call binding the contract method 0x49b03ca0.
//
// Solidity: function lastQueriedUrl() constant returns(string)
func (_AskMeAnything *AskMeAnythingCallerSession) LastQueriedUrl() (string, error) {
	return _AskMeAnything.Contract.LastQueriedUrl(&_AskMeAnything.CallOpts)
}

// LastRequestedRandom is a free data retrieval call binding the contract method 0xb8cf904e.
//
// Solidity: function lastRequestedRandom() constant returns(uint256)
func (_AskMeAnything *AskMeAnythingCaller) LastRequestedRandom(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "lastRequestedRandom")
	return *ret0, err
}

// LastRequestedRandom is a free data retrieval call binding the contract method 0xb8cf904e.
//
// Solidity: function lastRequestedRandom() constant returns(uint256)
func (_AskMeAnything *AskMeAnythingSession) LastRequestedRandom() (*big.Int, error) {
	return _AskMeAnything.Contract.LastRequestedRandom(&_AskMeAnything.CallOpts)
}

// LastRequestedRandom is a free data retrieval call binding the contract method 0xb8cf904e.
//
// Solidity: function lastRequestedRandom() constant returns(uint256)
func (_AskMeAnything *AskMeAnythingCallerSession) LastRequestedRandom() (*big.Int, error) {
	return _AskMeAnything.Contract.LastRequestedRandom(&_AskMeAnything.CallOpts)
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

// Random is a free data retrieval call binding the contract method 0x5ec01e4d.
//
// Solidity: function random() constant returns(uint256)
func (_AskMeAnything *AskMeAnythingCaller) Random(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "random")
	return *ret0, err
}

// Random is a free data retrieval call binding the contract method 0x5ec01e4d.
//
// Solidity: function random() constant returns(uint256)
func (_AskMeAnything *AskMeAnythingSession) Random() (*big.Int, error) {
	return _AskMeAnything.Contract.Random(&_AskMeAnything.CallOpts)
}

// Random is a free data retrieval call binding the contract method 0x5ec01e4d.
//
// Solidity: function random() constant returns(uint256)
func (_AskMeAnything *AskMeAnythingCallerSession) Random() (*big.Int, error) {
	return _AskMeAnything.Contract.Random(&_AskMeAnything.CallOpts)
}

// RepeatedCall is a free data retrieval call binding the contract method 0xab8c1bad.
//
// Solidity: function repeatedCall() constant returns(bool)
func (_AskMeAnything *AskMeAnythingCaller) RepeatedCall(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "repeatedCall")
	return *ret0, err
}

// RepeatedCall is a free data retrieval call binding the contract method 0xab8c1bad.
//
// Solidity: function repeatedCall() constant returns(bool)
func (_AskMeAnything *AskMeAnythingSession) RepeatedCall() (bool, error) {
	return _AskMeAnything.Contract.RepeatedCall(&_AskMeAnything.CallOpts)
}

// RepeatedCall is a free data retrieval call binding the contract method 0xab8c1bad.
//
// Solidity: function repeatedCall() constant returns(bool)
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

// AMA is a paid mutator transaction binding the contract method 0xf11ea3f8.
//
// Solidity: function AMA(url string, selector string) returns()
func (_AskMeAnything *AskMeAnythingTransactor) AMA(opts *bind.TransactOpts, url string, selector string) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "AMA", url, selector)
}

// AMA is a paid mutator transaction binding the contract method 0xf11ea3f8.
//
// Solidity: function AMA(url string, selector string) returns()
func (_AskMeAnything *AskMeAnythingSession) AMA(url string, selector string) (*types.Transaction, error) {
	return _AskMeAnything.Contract.AMA(&_AskMeAnything.TransactOpts, url, selector)
}

// AMA is a paid mutator transaction binding the contract method 0xf11ea3f8.
//
// Solidity: function AMA(url string, selector string) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) AMA(url string, selector string) (*types.Transaction, error) {
	return _AskMeAnything.Contract.AMA(&_AskMeAnything.TransactOpts, url, selector)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__(queryId uint256, result bytes) returns()
func (_AskMeAnything *AskMeAnythingTransactor) Callback_(opts *bind.TransactOpts, queryId *big.Int, result []byte) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "__callback__", queryId, result)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__(queryId uint256, result bytes) returns()
func (_AskMeAnything *AskMeAnythingSession) Callback_(queryId *big.Int, result []byte) (*types.Transaction, error) {
	return _AskMeAnything.Contract.Callback_(&_AskMeAnything.TransactOpts, queryId, result)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__(queryId uint256, result bytes) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) Callback_(queryId *big.Int, result []byte) (*types.Transaction, error) {
	return _AskMeAnything.Contract.Callback_(&_AskMeAnything.TransactOpts, queryId, result)
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

// RequestFastRandom is a paid mutator transaction binding the contract method 0x35cd49e7.
//
// Solidity: function requestFastRandom() returns()
func (_AskMeAnything *AskMeAnythingTransactor) RequestFastRandom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "requestFastRandom")
}

// RequestFastRandom is a paid mutator transaction binding the contract method 0x35cd49e7.
//
// Solidity: function requestFastRandom() returns()
func (_AskMeAnything *AskMeAnythingSession) RequestFastRandom() (*types.Transaction, error) {
	return _AskMeAnything.Contract.RequestFastRandom(&_AskMeAnything.TransactOpts)
}

// RequestFastRandom is a paid mutator transaction binding the contract method 0x35cd49e7.
//
// Solidity: function requestFastRandom() returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) RequestFastRandom() (*types.Transaction, error) {
	return _AskMeAnything.Contract.RequestFastRandom(&_AskMeAnything.TransactOpts)
}

// RequestSafeRandom is a paid mutator transaction binding the contract method 0xe1bcba05.
//
// Solidity: function requestSafeRandom() returns()
func (_AskMeAnything *AskMeAnythingTransactor) RequestSafeRandom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "requestSafeRandom")
}

// RequestSafeRandom is a paid mutator transaction binding the contract method 0xe1bcba05.
//
// Solidity: function requestSafeRandom() returns()
func (_AskMeAnything *AskMeAnythingSession) RequestSafeRandom() (*types.Transaction, error) {
	return _AskMeAnything.Contract.RequestSafeRandom(&_AskMeAnything.TransactOpts)
}

// RequestSafeRandom is a paid mutator transaction binding the contract method 0xe1bcba05.
//
// Solidity: function requestSafeRandom() returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) RequestSafeRandom() (*types.Transaction, error) {
	return _AskMeAnything.Contract.RequestSafeRandom(&_AskMeAnything.TransactOpts)
}

// SetQueryMode is a paid mutator transaction binding the contract method 0x57ae678b.
//
// Solidity: function setQueryMode(newMode bool) returns()
func (_AskMeAnything *AskMeAnythingTransactor) SetQueryMode(opts *bind.TransactOpts, newMode bool) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "setQueryMode", newMode)
}

// SetQueryMode is a paid mutator transaction binding the contract method 0x57ae678b.
//
// Solidity: function setQueryMode(newMode bool) returns()
func (_AskMeAnything *AskMeAnythingSession) SetQueryMode(newMode bool) (*types.Transaction, error) {
	return _AskMeAnything.Contract.SetQueryMode(&_AskMeAnything.TransactOpts, newMode)
}

// SetQueryMode is a paid mutator transaction binding the contract method 0x57ae678b.
//
// Solidity: function setQueryMode(newMode bool) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) SetQueryMode(newMode bool) (*types.Transaction, error) {
	return _AskMeAnything.Contract.SetQueryMode(&_AskMeAnything.TransactOpts, newMode)
}

// SetTimeout is a paid mutator transaction binding the contract method 0xc58a34cc.
//
// Solidity: function setTimeout(newTimeout uint256) returns()
func (_AskMeAnything *AskMeAnythingTransactor) SetTimeout(opts *bind.TransactOpts, newTimeout *big.Int) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "setTimeout", newTimeout)
}

// SetTimeout is a paid mutator transaction binding the contract method 0xc58a34cc.
//
// Solidity: function setTimeout(newTimeout uint256) returns()
func (_AskMeAnything *AskMeAnythingSession) SetTimeout(newTimeout *big.Int) (*types.Transaction, error) {
	return _AskMeAnything.Contract.SetTimeout(&_AskMeAnything.TransactOpts, newTimeout)
}

// SetTimeout is a paid mutator transaction binding the contract method 0xc58a34cc.
//
// Solidity: function setTimeout(newTimeout uint256) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) SetTimeout(newTimeout *big.Int) (*types.Transaction, error) {
	return _AskMeAnything.Contract.SetTimeout(&_AskMeAnything.TransactOpts, newTimeout)
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

// AskMeAnythingQueryResponseReadyIterator is returned from FilterQueryResponseReady and is used to iterate over the raw logs and unpacked data for QueryResponseReady events raised by the AskMeAnything contract.
type AskMeAnythingQueryResponseReadyIterator struct {
	Event *AskMeAnythingQueryResponseReady // Event containing the contract specifics and raw log

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
func (it *AskMeAnythingQueryResponseReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AskMeAnythingQueryResponseReady)
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
		it.Event = new(AskMeAnythingQueryResponseReady)
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
func (it *AskMeAnythingQueryResponseReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AskMeAnythingQueryResponseReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AskMeAnythingQueryResponseReady represents a QueryResponseReady event raised by the AskMeAnything contract.
type AskMeAnythingQueryResponseReady struct {
	QueryId *big.Int
	Result  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterQueryResponseReady is a free log retrieval operation binding the contract event 0xfe1788dc549f39fbcdb06fdab5937f2d19ce5fe5616fa3be9311928bd5dabccb.
//
// Solidity: e QueryResponseReady(queryId uint256, result string)
func (_AskMeAnything *AskMeAnythingFilterer) FilterQueryResponseReady(opts *bind.FilterOpts) (*AskMeAnythingQueryResponseReadyIterator, error) {

	logs, sub, err := _AskMeAnything.contract.FilterLogs(opts, "QueryResponseReady")
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingQueryResponseReadyIterator{contract: _AskMeAnything.contract, event: "QueryResponseReady", logs: logs, sub: sub}, nil
}

// WatchQueryResponseReady is a free log subscription operation binding the contract event 0xfe1788dc549f39fbcdb06fdab5937f2d19ce5fe5616fa3be9311928bd5dabccb.
//
// Solidity: e QueryResponseReady(queryId uint256, result string)
func (_AskMeAnything *AskMeAnythingFilterer) WatchQueryResponseReady(opts *bind.WatchOpts, sink chan<- *AskMeAnythingQueryResponseReady) (event.Subscription, error) {

	logs, sub, err := _AskMeAnything.contract.WatchLogs(opts, "QueryResponseReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AskMeAnythingQueryResponseReady)
				if err := _AskMeAnything.contract.UnpackLog(event, "QueryResponseReady", log); err != nil {
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

// AskMeAnythingRandomReadyIterator is returned from FilterRandomReady and is used to iterate over the raw logs and unpacked data for RandomReady events raised by the AskMeAnything contract.
type AskMeAnythingRandomReadyIterator struct {
	Event *AskMeAnythingRandomReady // Event containing the contract specifics and raw log

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
func (it *AskMeAnythingRandomReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AskMeAnythingRandomReady)
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
		it.Event = new(AskMeAnythingRandomReady)
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
func (it *AskMeAnythingRandomReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AskMeAnythingRandomReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AskMeAnythingRandomReady represents a RandomReady event raised by the AskMeAnything contract.
type AskMeAnythingRandomReady struct {
	RequestId       *big.Int
	GeneratedRandom *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRandomReady is a free log retrieval operation binding the contract event 0xd0ecc71f8b5af397da9123fd2bff63c544c04af5c6935935a7f81e14b84522f2.
//
// Solidity: e RandomReady(requestId uint256, generatedRandom uint256)
func (_AskMeAnything *AskMeAnythingFilterer) FilterRandomReady(opts *bind.FilterOpts) (*AskMeAnythingRandomReadyIterator, error) {

	logs, sub, err := _AskMeAnything.contract.FilterLogs(opts, "RandomReady")
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingRandomReadyIterator{contract: _AskMeAnything.contract, event: "RandomReady", logs: logs, sub: sub}, nil
}

// WatchRandomReady is a free log subscription operation binding the contract event 0xd0ecc71f8b5af397da9123fd2bff63c544c04af5c6935935a7f81e14b84522f2.
//
// Solidity: e RandomReady(requestId uint256, generatedRandom uint256)
func (_AskMeAnything *AskMeAnythingFilterer) WatchRandomReady(opts *bind.WatchOpts, sink chan<- *AskMeAnythingRandomReady) (event.Subscription, error) {

	logs, sub, err := _AskMeAnything.contract.WatchLogs(opts, "RandomReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AskMeAnythingRandomReady)
				if err := _AskMeAnything.contract.UnpackLog(event, "RandomReady", log); err != nil {
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

// AskMeAnythingRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the AskMeAnything contract.
type AskMeAnythingRequestSentIterator struct {
	Event *AskMeAnythingRequestSent // Event containing the contract specifics and raw log

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
func (it *AskMeAnythingRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AskMeAnythingRequestSent)
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
		it.Event = new(AskMeAnythingRequestSent)
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
func (it *AskMeAnythingRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AskMeAnythingRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AskMeAnythingRequestSent represents a RequestSent event raised by the AskMeAnything contract.
type AskMeAnythingRequestSent struct {
	Succ      bool
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x4d0a0fbbf63d9b08d7e94ab32ae848f6fbf902858f874ac63a475db03268d982.
//
// Solidity: e RequestSent(succ bool, requestId uint256)
func (_AskMeAnything *AskMeAnythingFilterer) FilterRequestSent(opts *bind.FilterOpts) (*AskMeAnythingRequestSentIterator, error) {

	logs, sub, err := _AskMeAnything.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingRequestSentIterator{contract: _AskMeAnything.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x4d0a0fbbf63d9b08d7e94ab32ae848f6fbf902858f874ac63a475db03268d982.
//
// Solidity: e RequestSent(succ bool, requestId uint256)
func (_AskMeAnything *AskMeAnythingFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *AskMeAnythingRequestSent) (event.Subscription, error) {

	logs, sub, err := _AskMeAnything.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AskMeAnythingRequestSent)
				if err := _AskMeAnything.contract.UnpackLog(event, "RequestSent", log); err != nil {
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
// Solidity: e SetTimeout(previousTimeout uint256, newTimeout uint256)
func (_AskMeAnything *AskMeAnythingFilterer) FilterSetTimeout(opts *bind.FilterOpts) (*AskMeAnythingSetTimeoutIterator, error) {

	logs, sub, err := _AskMeAnything.contract.FilterLogs(opts, "SetTimeout")
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingSetTimeoutIterator{contract: _AskMeAnything.contract, event: "SetTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTimeout is a free log subscription operation binding the contract event 0x9aa0de0157c9133b911d2d811f590159622cea28cefe31505c203c828799da58.
//
// Solidity: e SetTimeout(previousTimeout uint256, newTimeout uint256)
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
const DOSOnChainSDKABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"generatedRandom\",\"type\":\"uint256\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"queryId\",\"type\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DOSOnChainSDKBin is the compiled bytecode used for deploying new contracts.
const DOSOnChainSDKBin = `0x608060405260018054600160a060020a0319167314b03cf8ed30e4c58af3dbb6e0a3d1465df1850617905534801561003657600080fd5b5060bd806100456000396000f30060806040526004361060485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318a1908d8114604d5780636d112977146067575b600080fd5b348015605857600080fd5b5060656004356024356088565b005b348015607257600080fd5b506065600480359060248035908101910135608c565b5050565b5050505600a165627a7a72305820e05a3e339a8e459bb5115cffe3b41c80f43b1e6bb697e11da024603cfa16febe0029`

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

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__(queryId uint256, result bytes) returns()
func (_DOSOnChainSDK *DOSOnChainSDKTransactor) Callback_(opts *bind.TransactOpts, queryId *big.Int, result []byte) (*types.Transaction, error) {
	return _DOSOnChainSDK.contract.Transact(opts, "__callback__", queryId, result)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__(queryId uint256, result bytes) returns()
func (_DOSOnChainSDK *DOSOnChainSDKSession) Callback_(queryId *big.Int, result []byte) (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.Callback_(&_DOSOnChainSDK.TransactOpts, queryId, result)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__(queryId uint256, result bytes) returns()
func (_DOSOnChainSDK *DOSOnChainSDKTransactorSession) Callback_(queryId *big.Int, result []byte) (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.Callback_(&_DOSOnChainSDK.TransactOpts, queryId, result)
}

// DOSProxyInterfaceABI is the input ABI used to generate the binding from.
const DOSProxyInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestRandom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// Query is a paid mutator transaction binding the contract method 0xb7fb8fd7.
//
// Solidity: function query( address,  uint256,  string,  string) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceTransactor) Query(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 string, arg3 string) (*types.Transaction, error) {
	return _DOSProxyInterface.contract.Transact(opts, "query", arg0, arg1, arg2, arg3)
}

// Query is a paid mutator transaction binding the contract method 0xb7fb8fd7.
//
// Solidity: function query( address,  uint256,  string,  string) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceSession) Query(arg0 common.Address, arg1 *big.Int, arg2 string, arg3 string) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.Query(&_DOSProxyInterface.TransactOpts, arg0, arg1, arg2, arg3)
}

// Query is a paid mutator transaction binding the contract method 0xb7fb8fd7.
//
// Solidity: function query( address,  uint256,  string,  string) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceTransactorSession) Query(arg0 common.Address, arg1 *big.Int, arg2 string, arg3 string) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.Query(&_DOSProxyInterface.TransactOpts, arg0, arg1, arg2, arg3)
}

// RequestRandom is a paid mutator transaction binding the contract method 0x1bf82057.
//
// Solidity: function requestRandom( address,  uint8,  uint256) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceTransactor) RequestRandom(opts *bind.TransactOpts, arg0 common.Address, arg1 uint8, arg2 *big.Int) (*types.Transaction, error) {
	return _DOSProxyInterface.contract.Transact(opts, "requestRandom", arg0, arg1, arg2)
}

// RequestRandom is a paid mutator transaction binding the contract method 0x1bf82057.
//
// Solidity: function requestRandom( address,  uint8,  uint256) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceSession) RequestRandom(arg0 common.Address, arg1 uint8, arg2 *big.Int) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.RequestRandom(&_DOSProxyInterface.TransactOpts, arg0, arg1, arg2)
}

// RequestRandom is a paid mutator transaction binding the contract method 0x1bf82057.
//
// Solidity: function requestRandom( address,  uint8,  uint256) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceTransactorSession) RequestRandom(arg0 common.Address, arg1 uint8, arg2 *big.Int) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.RequestRandom(&_DOSProxyInterface.TransactOpts, arg0, arg1, arg2)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610248806100326000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663715018a681146100665780638da5cb5b1461007d5780638f32d59b146100ae578063f2fde38b146100d7575b600080fd5b34801561007257600080fd5b5061007b6100f8565b005b34801561008957600080fd5b50610092610160565b60408051600160a060020a039092168252519081900360200190f35b3480156100ba57600080fd5b506100c361016f565b604080519115158252519081900360200190f35b3480156100e357600080fd5b5061007b600160a060020a0360043516610180565b61010061016f565b151561010b57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031690565b600054600160a060020a0316331490565b61018861016f565b151561019357600080fd5b61019c8161019f565b50565b600160a060020a03811615156101b457600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820c1f5e23e3ca8c3136bf365d866b447dee65b2773a25e098889a6159da50a363f0029`

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

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
const UtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820c0fa8c03761786a8659d183240455e77007a2d402d9ddaf00733ea5b9b892a990029`

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}
