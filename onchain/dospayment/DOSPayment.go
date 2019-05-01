// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dospayment

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

// DospaymentABI is the input ABI used to generate the binding from.
const DospaymentABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNetworkToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"droplockToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"droplockMaxQuota\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDroplockToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"address\"}],\"name\":\"fromValidStakingNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"quo\",\"type\":\"uint256\"}],\"name\":\"setDroplockMaxQuota\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateNetworkTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateDroplockTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"UpdateDroplockMaxQuota\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DospaymentBin is the compiled bytecode used for deploying new contracts.
const DospaymentBin = `608060405260018054600160a060020a031990811673214e79c85744cd2ebbc64ddc0047131496871bee1790915561c3506003908155600455600080549091163317905561081e806100526000396000f3fe608060405234801561001057600080fd5b50600436106100d1576000357c0100000000000000000000000000000000000000000000000000000000900480638da5cb5b1161008e5780638da5cb5b146101545780638f32d59b1461015c5780639cebe97b14610178578063c7e6a9bc1461019e578063e13c7e44146101c4578063f2fde38b146101e1576100d1565b806317107c49146100d6578063375b3c0a146100fe57806353212c46146101185780636ca95a4e1461013c578063715018a6146101445780637bfe52931461014c575b600080fd5b6100fc600480360360208110156100ec57600080fd5b5035600160a060020a0316610207565b005b610106610291565b60408051918252519081900360200190f35b610120610297565b60408051600160a060020a039092168252519081900360200190f35b6101206102a6565b6100fc6102b5565b61010661031d565b610120610323565b610164610332565b604080519115158252519081900360200190f35b6100fc6004803603602081101561018e57600080fd5b5035600160a060020a0316610343565b610164600480360360208110156101b457600080fd5b5035600160a060020a03166103cd565b6100fc600480360360208110156101da57600080fd5b5035610674565b6100fc600480360360208110156101f757600080fd5b5035600160a060020a0316610732565b61020f610332565b151561021a57600080fd5b60015460408051600160a060020a039283168152918316602083015280517f4d27a2adceae86b92fb74fb7e8f96dc902d917e243fbff389b5a793c9040dafe9281900390910190a16001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60035481565b600254600160a060020a031681565b600154600160a060020a031681565b6102bd610332565b15156102c857600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b60045481565b600054600160a060020a031690565b600054600160a060020a0316331490565b61034b610332565b151561035657600080fd5b60025460408051600160a060020a039283168152918316602083015280517f26d12c8278b5711a05f4e96bd6e91cac2e75b4143b3d1edd6a1af194d30bceef9281900390910190a16002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093849316916370a08231916024808301926020929190829003018186803b15801561043557600080fd5b505afa158015610449573d6000803e3d6000fd5b505050506040513d602081101561045f57600080fd5b5051600154604080517f313ce5670000000000000000000000000000000000000000000000000000000081529051929350600092600160a060020a039092169163313ce56791600480820192602092909190829003018186803b1580156104c557600080fd5b505afa1580156104d9573d6000803e3d6000fd5b505050506040513d60208110156104ef57600080fd5b50516003549091508102821061050a5760019250505061066f565b600254600160a060020a031615156105275760009250505061066f565b600254604080517f313ce5670000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163313ce567916004808301926020929190829003018186803b15801561058557600080fd5b505afa158015610599573d6000803e3d6000fd5b505050506040513d60208110156105af57600080fd5b5051600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038981166004830152915191909216916370a08231916024808301926020929190829003018186803b15801561061757600080fd5b505afa15801561062b573d6000803e3d6000fd5b505050506040513d602081101561064157600080fd5b505181151561064c57fe5b04905060045481111561065e57506004545b600354600a92029082030204111590505b919050565b61067c610332565b151561068757600080fd5b60045481141580156106995750600a81105b15156106f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806107cf6024913960400191505060405180910390fd5b600454604080519182526020820183905280517f53714c2a9a4391f82767aa3e422fba28700b1e85a8f283b94d609b313e417c559281900390910190a1600455565b61073a610332565b151561074557600080fd5b61074e81610751565b50565b600160a060020a038116151561076657600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556fe56616c69642064726f706c6f636b4d617851756f74612077697468696e203020746f2039a165627a7a723058205e26e836097940de28f0a819a01f925b630231a0d0dbc387c178ddf74ef3d7880029`

// DeployDospayment deploys a new Ethereum contract, binding an instance of Dospayment to it.
func DeployDospayment(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dospayment, error) {
	parsed, err := abi.JSON(strings.NewReader(DospaymentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DospaymentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dospayment{DospaymentCaller: DospaymentCaller{contract: contract}, DospaymentTransactor: DospaymentTransactor{contract: contract}, DospaymentFilterer: DospaymentFilterer{contract: contract}}, nil
}

// Dospayment is an auto generated Go binding around an Ethereum contract.
type Dospayment struct {
	DospaymentCaller     // Read-only binding to the contract
	DospaymentTransactor // Write-only binding to the contract
	DospaymentFilterer   // Log filterer for contract events
}

// DospaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type DospaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DospaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DospaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DospaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DospaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DospaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DospaymentSession struct {
	Contract     *Dospayment       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DospaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DospaymentCallerSession struct {
	Contract *DospaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DospaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DospaymentTransactorSession struct {
	Contract     *DospaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DospaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type DospaymentRaw struct {
	Contract *Dospayment // Generic contract binding to access the raw methods on
}

// DospaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DospaymentCallerRaw struct {
	Contract *DospaymentCaller // Generic read-only contract binding to access the raw methods on
}

// DospaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DospaymentTransactorRaw struct {
	Contract *DospaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDospayment creates a new instance of Dospayment, bound to a specific deployed contract.
func NewDospayment(address common.Address, backend bind.ContractBackend) (*Dospayment, error) {
	contract, err := bindDospayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dospayment{DospaymentCaller: DospaymentCaller{contract: contract}, DospaymentTransactor: DospaymentTransactor{contract: contract}, DospaymentFilterer: DospaymentFilterer{contract: contract}}, nil
}

// NewDospaymentCaller creates a new read-only instance of Dospayment, bound to a specific deployed contract.
func NewDospaymentCaller(address common.Address, caller bind.ContractCaller) (*DospaymentCaller, error) {
	contract, err := bindDospayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DospaymentCaller{contract: contract}, nil
}

// NewDospaymentTransactor creates a new write-only instance of Dospayment, bound to a specific deployed contract.
func NewDospaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*DospaymentTransactor, error) {
	contract, err := bindDospayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DospaymentTransactor{contract: contract}, nil
}

// NewDospaymentFilterer creates a new log filterer instance of Dospayment, bound to a specific deployed contract.
func NewDospaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*DospaymentFilterer, error) {
	contract, err := bindDospayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DospaymentFilterer{contract: contract}, nil
}

// bindDospayment binds a generic wrapper to an already deployed contract.
func bindDospayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DospaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dospayment *DospaymentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dospayment.Contract.DospaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dospayment *DospaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dospayment.Contract.DospaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dospayment *DospaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dospayment.Contract.DospaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dospayment *DospaymentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dospayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dospayment *DospaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dospayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dospayment *DospaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dospayment.Contract.contract.Transact(opts, method, params...)
}

// DroplockMaxQuota is a free data retrieval call binding the contract method 0x7bfe5293.
//
// Solidity: function droplockMaxQuota() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DroplockMaxQuota(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "droplockMaxQuota")
	return *ret0, err
}

// DroplockMaxQuota is a free data retrieval call binding the contract method 0x7bfe5293.
//
// Solidity: function droplockMaxQuota() constant returns(uint256)
func (_Dospayment *DospaymentSession) DroplockMaxQuota() (*big.Int, error) {
	return _Dospayment.Contract.DroplockMaxQuota(&_Dospayment.CallOpts)
}

// DroplockMaxQuota is a free data retrieval call binding the contract method 0x7bfe5293.
//
// Solidity: function droplockMaxQuota() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DroplockMaxQuota() (*big.Int, error) {
	return _Dospayment.Contract.DroplockMaxQuota(&_Dospayment.CallOpts)
}

// DroplockToken is a free data retrieval call binding the contract method 0x53212c46.
//
// Solidity: function droplockToken() constant returns(address)
func (_Dospayment *DospaymentCaller) DroplockToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "droplockToken")
	return *ret0, err
}

// DroplockToken is a free data retrieval call binding the contract method 0x53212c46.
//
// Solidity: function droplockToken() constant returns(address)
func (_Dospayment *DospaymentSession) DroplockToken() (common.Address, error) {
	return _Dospayment.Contract.DroplockToken(&_Dospayment.CallOpts)
}

// DroplockToken is a free data retrieval call binding the contract method 0x53212c46.
//
// Solidity: function droplockToken() constant returns(address)
func (_Dospayment *DospaymentCallerSession) DroplockToken() (common.Address, error) {
	return _Dospayment.Contract.DroplockToken(&_Dospayment.CallOpts)
}

// FromValidStakingNode is a free data retrieval call binding the contract method 0xc7e6a9bc.
//
// Solidity: function fromValidStakingNode(address node) constant returns(bool)
func (_Dospayment *DospaymentCaller) FromValidStakingNode(opts *bind.CallOpts, node common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "fromValidStakingNode", node)
	return *ret0, err
}

// FromValidStakingNode is a free data retrieval call binding the contract method 0xc7e6a9bc.
//
// Solidity: function fromValidStakingNode(address node) constant returns(bool)
func (_Dospayment *DospaymentSession) FromValidStakingNode(node common.Address) (bool, error) {
	return _Dospayment.Contract.FromValidStakingNode(&_Dospayment.CallOpts, node)
}

// FromValidStakingNode is a free data retrieval call binding the contract method 0xc7e6a9bc.
//
// Solidity: function fromValidStakingNode(address node) constant returns(bool)
func (_Dospayment *DospaymentCallerSession) FromValidStakingNode(node common.Address) (bool, error) {
	return _Dospayment.Contract.FromValidStakingNode(&_Dospayment.CallOpts, node)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dospayment *DospaymentCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dospayment *DospaymentSession) IsOwner() (bool, error) {
	return _Dospayment.Contract.IsOwner(&_Dospayment.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dospayment *DospaymentCallerSession) IsOwner() (bool, error) {
	return _Dospayment.Contract.IsOwner(&_Dospayment.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_Dospayment *DospaymentCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "minStake")
	return *ret0, err
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_Dospayment *DospaymentSession) MinStake() (*big.Int, error) {
	return _Dospayment.Contract.MinStake(&_Dospayment.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) MinStake() (*big.Int, error) {
	return _Dospayment.Contract.MinStake(&_Dospayment.CallOpts)
}

// NetworkToken is a free data retrieval call binding the contract method 0x6ca95a4e.
//
// Solidity: function networkToken() constant returns(address)
func (_Dospayment *DospaymentCaller) NetworkToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "networkToken")
	return *ret0, err
}

// NetworkToken is a free data retrieval call binding the contract method 0x6ca95a4e.
//
// Solidity: function networkToken() constant returns(address)
func (_Dospayment *DospaymentSession) NetworkToken() (common.Address, error) {
	return _Dospayment.Contract.NetworkToken(&_Dospayment.CallOpts)
}

// NetworkToken is a free data retrieval call binding the contract method 0x6ca95a4e.
//
// Solidity: function networkToken() constant returns(address)
func (_Dospayment *DospaymentCallerSession) NetworkToken() (common.Address, error) {
	return _Dospayment.Contract.NetworkToken(&_Dospayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dospayment *DospaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dospayment *DospaymentSession) Owner() (common.Address, error) {
	return _Dospayment.Contract.Owner(&_Dospayment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dospayment *DospaymentCallerSession) Owner() (common.Address, error) {
	return _Dospayment.Contract.Owner(&_Dospayment.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dospayment *DospaymentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dospayment *DospaymentSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dospayment.Contract.RenounceOwnership(&_Dospayment.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dospayment *DospaymentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dospayment.Contract.RenounceOwnership(&_Dospayment.TransactOpts)
}

// SetDroplockMaxQuota is a paid mutator transaction binding the contract method 0xe13c7e44.
//
// Solidity: function setDroplockMaxQuota(uint256 quo) returns()
func (_Dospayment *DospaymentTransactor) SetDroplockMaxQuota(opts *bind.TransactOpts, quo *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setDroplockMaxQuota", quo)
}

// SetDroplockMaxQuota is a paid mutator transaction binding the contract method 0xe13c7e44.
//
// Solidity: function setDroplockMaxQuota(uint256 quo) returns()
func (_Dospayment *DospaymentSession) SetDroplockMaxQuota(quo *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDroplockMaxQuota(&_Dospayment.TransactOpts, quo)
}

// SetDroplockMaxQuota is a paid mutator transaction binding the contract method 0xe13c7e44.
//
// Solidity: function setDroplockMaxQuota(uint256 quo) returns()
func (_Dospayment *DospaymentTransactorSession) SetDroplockMaxQuota(quo *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDroplockMaxQuota(&_Dospayment.TransactOpts, quo)
}

// SetDroplockToken is a paid mutator transaction binding the contract method 0x9cebe97b.
//
// Solidity: function setDroplockToken(address addr) returns()
func (_Dospayment *DospaymentTransactor) SetDroplockToken(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setDroplockToken", addr)
}

// SetDroplockToken is a paid mutator transaction binding the contract method 0x9cebe97b.
//
// Solidity: function setDroplockToken(address addr) returns()
func (_Dospayment *DospaymentSession) SetDroplockToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDroplockToken(&_Dospayment.TransactOpts, addr)
}

// SetDroplockToken is a paid mutator transaction binding the contract method 0x9cebe97b.
//
// Solidity: function setDroplockToken(address addr) returns()
func (_Dospayment *DospaymentTransactorSession) SetDroplockToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDroplockToken(&_Dospayment.TransactOpts, addr)
}

// SetNetworkToken is a paid mutator transaction binding the contract method 0x17107c49.
//
// Solidity: function setNetworkToken(address addr) returns()
func (_Dospayment *DospaymentTransactor) SetNetworkToken(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setNetworkToken", addr)
}

// SetNetworkToken is a paid mutator transaction binding the contract method 0x17107c49.
//
// Solidity: function setNetworkToken(address addr) returns()
func (_Dospayment *DospaymentSession) SetNetworkToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetNetworkToken(&_Dospayment.TransactOpts, addr)
}

// SetNetworkToken is a paid mutator transaction binding the contract method 0x17107c49.
//
// Solidity: function setNetworkToken(address addr) returns()
func (_Dospayment *DospaymentTransactorSession) SetNetworkToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetNetworkToken(&_Dospayment.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dospayment *DospaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dospayment *DospaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.TransferOwnership(&_Dospayment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dospayment *DospaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.TransferOwnership(&_Dospayment.TransactOpts, newOwner)
}

// DospaymentOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Dospayment contract.
type DospaymentOwnershipRenouncedIterator struct {
	Event *DospaymentOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DospaymentOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentOwnershipRenounced)
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
		it.Event = new(DospaymentOwnershipRenounced)
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
func (it *DospaymentOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentOwnershipRenounced represents a OwnershipRenounced event raised by the Dospayment contract.
type DospaymentOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Dospayment *DospaymentFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DospaymentOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DospaymentOwnershipRenouncedIterator{contract: _Dospayment.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Dospayment *DospaymentFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DospaymentOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentOwnershipRenounced)
				if err := _Dospayment.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DospaymentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dospayment contract.
type DospaymentOwnershipTransferredIterator struct {
	Event *DospaymentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DospaymentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentOwnershipTransferred)
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
		it.Event = new(DospaymentOwnershipTransferred)
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
func (it *DospaymentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentOwnershipTransferred represents a OwnershipTransferred event raised by the Dospayment contract.
type DospaymentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dospayment *DospaymentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DospaymentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DospaymentOwnershipTransferredIterator{contract: _Dospayment.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dospayment *DospaymentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DospaymentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentOwnershipTransferred)
				if err := _Dospayment.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// DospaymentUpdateDroplockMaxQuotaIterator is returned from FilterUpdateDroplockMaxQuota and is used to iterate over the raw logs and unpacked data for UpdateDroplockMaxQuota events raised by the Dospayment contract.
type DospaymentUpdateDroplockMaxQuotaIterator struct {
	Event *DospaymentUpdateDroplockMaxQuota // Event containing the contract specifics and raw log

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
func (it *DospaymentUpdateDroplockMaxQuotaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentUpdateDroplockMaxQuota)
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
		it.Event = new(DospaymentUpdateDroplockMaxQuota)
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
func (it *DospaymentUpdateDroplockMaxQuotaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentUpdateDroplockMaxQuotaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentUpdateDroplockMaxQuota represents a UpdateDroplockMaxQuota event raised by the Dospayment contract.
type DospaymentUpdateDroplockMaxQuota struct {
	OldQuota *big.Int
	NewQuota *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateDroplockMaxQuota is a free log retrieval operation binding the contract event 0x53714c2a9a4391f82767aa3e422fba28700b1e85a8f283b94d609b313e417c55.
//
// Solidity: event UpdateDroplockMaxQuota(uint256 oldQuota, uint256 newQuota)
func (_Dospayment *DospaymentFilterer) FilterUpdateDroplockMaxQuota(opts *bind.FilterOpts) (*DospaymentUpdateDroplockMaxQuotaIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "UpdateDroplockMaxQuota")
	if err != nil {
		return nil, err
	}
	return &DospaymentUpdateDroplockMaxQuotaIterator{contract: _Dospayment.contract, event: "UpdateDroplockMaxQuota", logs: logs, sub: sub}, nil
}

// WatchUpdateDroplockMaxQuota is a free log subscription operation binding the contract event 0x53714c2a9a4391f82767aa3e422fba28700b1e85a8f283b94d609b313e417c55.
//
// Solidity: event UpdateDroplockMaxQuota(uint256 oldQuota, uint256 newQuota)
func (_Dospayment *DospaymentFilterer) WatchUpdateDroplockMaxQuota(opts *bind.WatchOpts, sink chan<- *DospaymentUpdateDroplockMaxQuota) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "UpdateDroplockMaxQuota")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentUpdateDroplockMaxQuota)
				if err := _Dospayment.contract.UnpackLog(event, "UpdateDroplockMaxQuota", log); err != nil {
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

// DospaymentUpdateDroplockTokenAddressIterator is returned from FilterUpdateDroplockTokenAddress and is used to iterate over the raw logs and unpacked data for UpdateDroplockTokenAddress events raised by the Dospayment contract.
type DospaymentUpdateDroplockTokenAddressIterator struct {
	Event *DospaymentUpdateDroplockTokenAddress // Event containing the contract specifics and raw log

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
func (it *DospaymentUpdateDroplockTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentUpdateDroplockTokenAddress)
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
		it.Event = new(DospaymentUpdateDroplockTokenAddress)
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
func (it *DospaymentUpdateDroplockTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentUpdateDroplockTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentUpdateDroplockTokenAddress represents a UpdateDroplockTokenAddress event raised by the Dospayment contract.
type DospaymentUpdateDroplockTokenAddress struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateDroplockTokenAddress is a free log retrieval operation binding the contract event 0x26d12c8278b5711a05f4e96bd6e91cac2e75b4143b3d1edd6a1af194d30bceef.
//
// Solidity: event UpdateDroplockTokenAddress(address oldAddress, address newAddress)
func (_Dospayment *DospaymentFilterer) FilterUpdateDroplockTokenAddress(opts *bind.FilterOpts) (*DospaymentUpdateDroplockTokenAddressIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "UpdateDroplockTokenAddress")
	if err != nil {
		return nil, err
	}
	return &DospaymentUpdateDroplockTokenAddressIterator{contract: _Dospayment.contract, event: "UpdateDroplockTokenAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateDroplockTokenAddress is a free log subscription operation binding the contract event 0x26d12c8278b5711a05f4e96bd6e91cac2e75b4143b3d1edd6a1af194d30bceef.
//
// Solidity: event UpdateDroplockTokenAddress(address oldAddress, address newAddress)
func (_Dospayment *DospaymentFilterer) WatchUpdateDroplockTokenAddress(opts *bind.WatchOpts, sink chan<- *DospaymentUpdateDroplockTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "UpdateDroplockTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentUpdateDroplockTokenAddress)
				if err := _Dospayment.contract.UnpackLog(event, "UpdateDroplockTokenAddress", log); err != nil {
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

// DospaymentUpdateNetworkTokenAddressIterator is returned from FilterUpdateNetworkTokenAddress and is used to iterate over the raw logs and unpacked data for UpdateNetworkTokenAddress events raised by the Dospayment contract.
type DospaymentUpdateNetworkTokenAddressIterator struct {
	Event *DospaymentUpdateNetworkTokenAddress // Event containing the contract specifics and raw log

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
func (it *DospaymentUpdateNetworkTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentUpdateNetworkTokenAddress)
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
		it.Event = new(DospaymentUpdateNetworkTokenAddress)
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
func (it *DospaymentUpdateNetworkTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentUpdateNetworkTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentUpdateNetworkTokenAddress represents a UpdateNetworkTokenAddress event raised by the Dospayment contract.
type DospaymentUpdateNetworkTokenAddress struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateNetworkTokenAddress is a free log retrieval operation binding the contract event 0x4d27a2adceae86b92fb74fb7e8f96dc902d917e243fbff389b5a793c9040dafe.
//
// Solidity: event UpdateNetworkTokenAddress(address oldAddress, address newAddress)
func (_Dospayment *DospaymentFilterer) FilterUpdateNetworkTokenAddress(opts *bind.FilterOpts) (*DospaymentUpdateNetworkTokenAddressIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "UpdateNetworkTokenAddress")
	if err != nil {
		return nil, err
	}
	return &DospaymentUpdateNetworkTokenAddressIterator{contract: _Dospayment.contract, event: "UpdateNetworkTokenAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateNetworkTokenAddress is a free log subscription operation binding the contract event 0x4d27a2adceae86b92fb74fb7e8f96dc902d917e243fbff389b5a793c9040dafe.
//
// Solidity: event UpdateNetworkTokenAddress(address oldAddress, address newAddress)
func (_Dospayment *DospaymentFilterer) WatchUpdateNetworkTokenAddress(opts *bind.WatchOpts, sink chan<- *DospaymentUpdateNetworkTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "UpdateNetworkTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentUpdateNetworkTokenAddress)
				if err := _Dospayment.contract.UnpackLog(event, "UpdateNetworkTokenAddress", log); err != nil {
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
