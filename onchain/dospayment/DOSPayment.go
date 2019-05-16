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
const DospaymentABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNetworkToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"quo\",\"type\":\"uint256\"}],\"name\":\"setDropBurnMaxQuota\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dropburnToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"address\"}],\"name\":\"fromValidStakingNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dropburnMaxQuota\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDropBurnToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateNetworkTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateDropBurnTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"UpdateDropBurnMaxQuota\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DospaymentBin is the compiled bytecode used for deploying new contracts.
const DospaymentBin = `6080604052600180546001600160a01b031990811673214e79c85744cd2ebbc64ddc0047131496871bee17909155600280548216739bfe8f5749d90eb4049ad94cc4de9b6c4c31f82217905561c3506003908155600455600080549091163317905561075b806100706000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638f32d59b116100715780638f32d59b1461014c578063b26584b814610168578063c7e6a9bc14610170578063e8c3470c14610196578063eac051f91461019e578063f2fde38b146101c4576100b4565b806317107c49146100b9578063375b3c0a146100e15780633f3381e1146100fb5780636ca95a4e14610118578063715018a61461013c5780638da5cb5b14610144575b600080fd5b6100df600480360360208110156100cf57600080fd5b50356001600160a01b03166101ea565b005b6100e9610265565b60408051918252519081900360200190f35b6100df6004803603602081101561011157600080fd5b503561026b565b61012061030e565b604080516001600160a01b039092168252519081900360200190f35b6100df61031d565b610120610376565b610154610385565b604080519115158252519081900360200190f35b610120610396565b6101546004803603602081101561018657600080fd5b50356001600160a01b03166103a5565b6100e96105ff565b6100df600480360360208110156101b457600080fd5b50356001600160a01b0316610605565b6100df600480360360208110156101da57600080fd5b50356001600160a01b0316610680565b6101f2610385565b6101fb57600080fd5b600154604080516001600160a01b039283168152918316602083015280517f4d27a2adceae86b92fb74fb7e8f96dc902d917e243fbff389b5a793c9040dafe9281900390910190a1600180546001600160a01b0319166001600160a01b0392909216919091179055565b60035481565b610273610385565b61027c57600080fd5b600454811415801561028e5750600a81105b6102cc57604051600160e51b62461bcd02815260040180806020018281038252602481526020018061070c6024913960400191505060405180910390fd5b600454604080519182526020820183905280517f0aee95cca46da64ee373e28dee5994361b4002c54035d92932c9825b76382e999281900390910190a1600455565b6001546001600160a01b031681565b610325610385565b61032e57600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b6002546001600160a01b031681565b60015460408051600160e01b6370a082310281526001600160a01b0384811660048301529151600093849316916370a08231916024808301926020929190829003018186803b1580156103f757600080fd5b505afa15801561040b573d6000803e3d6000fd5b505050506040513d602081101561042157600080fd5b505160015460408051600160e01b63313ce56702815290519293506000926001600160a01b039092169163313ce56791600480820192602092909190829003018186803b15801561047157600080fd5b505afa158015610485573d6000803e3d6000fd5b505050506040513d602081101561049b57600080fd5b5051600354909150600a82900a028083106104bc57600193505050506105fa565b6002546001600160a01b03166104d857600093505050506105fa565b60025460408051600160e01b63313ce56702815290516000926001600160a01b03169163313ce567916004808301926020929190829003018186803b15801561052057600080fd5b505afa158015610534573d6000803e3d6000fd5b505050506040513d602081101561054a57600080fd5b505160025460408051600160e01b6370a082310281526001600160a01b038a811660048301529151600a9490940a9391909216916370a08231916024808301926020929190829003018186803b1580156105a357600080fd5b505afa1580156105b7573d6000803e3d6000fd5b505050506040513d60208110156105cd57600080fd5b5051816105d657fe5b0490506004548111156105e857506004545b600a8181038302048410159450505050505b919050565b60045481565b61060d610385565b61061657600080fd5b600254604080516001600160a01b039283168152918316602083015280517ffc8013dfb0c8d38f3bcab9239bd5712457c48919b272cdb109488549199a01739281900390910190a1600280546001600160a01b0319166001600160a01b0392909216919091179055565b610688610385565b61069157600080fd5b61069a8161069d565b50565b6001600160a01b0381166106b057600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe56616c69642064726f706275726e4d617851756f74612077697468696e203020746f2039a165627a7a7230582064b4735d729d18fc19bb831f239b60033a5e8d3c475d6f21283c576102a14d520029`

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

// DropburnMaxQuota is a free data retrieval call binding the contract method 0xe8c3470c.
//
// Solidity: function dropburnMaxQuota() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DropburnMaxQuota(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "dropburnMaxQuota")
	return *ret0, err
}

// DropburnMaxQuota is a free data retrieval call binding the contract method 0xe8c3470c.
//
// Solidity: function dropburnMaxQuota() constant returns(uint256)
func (_Dospayment *DospaymentSession) DropburnMaxQuota() (*big.Int, error) {
	return _Dospayment.Contract.DropburnMaxQuota(&_Dospayment.CallOpts)
}

// DropburnMaxQuota is a free data retrieval call binding the contract method 0xe8c3470c.
//
// Solidity: function dropburnMaxQuota() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DropburnMaxQuota() (*big.Int, error) {
	return _Dospayment.Contract.DropburnMaxQuota(&_Dospayment.CallOpts)
}

// DropburnToken is a free data retrieval call binding the contract method 0xb26584b8.
//
// Solidity: function dropburnToken() constant returns(address)
func (_Dospayment *DospaymentCaller) DropburnToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "dropburnToken")
	return *ret0, err
}

// DropburnToken is a free data retrieval call binding the contract method 0xb26584b8.
//
// Solidity: function dropburnToken() constant returns(address)
func (_Dospayment *DospaymentSession) DropburnToken() (common.Address, error) {
	return _Dospayment.Contract.DropburnToken(&_Dospayment.CallOpts)
}

// DropburnToken is a free data retrieval call binding the contract method 0xb26584b8.
//
// Solidity: function dropburnToken() constant returns(address)
func (_Dospayment *DospaymentCallerSession) DropburnToken() (common.Address, error) {
	return _Dospayment.Contract.DropburnToken(&_Dospayment.CallOpts)
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

// SetDropBurnMaxQuota is a paid mutator transaction binding the contract method 0x3f3381e1.
//
// Solidity: function setDropBurnMaxQuota(uint256 quo) returns()
func (_Dospayment *DospaymentTransactor) SetDropBurnMaxQuota(opts *bind.TransactOpts, quo *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setDropBurnMaxQuota", quo)
}

// SetDropBurnMaxQuota is a paid mutator transaction binding the contract method 0x3f3381e1.
//
// Solidity: function setDropBurnMaxQuota(uint256 quo) returns()
func (_Dospayment *DospaymentSession) SetDropBurnMaxQuota(quo *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDropBurnMaxQuota(&_Dospayment.TransactOpts, quo)
}

// SetDropBurnMaxQuota is a paid mutator transaction binding the contract method 0x3f3381e1.
//
// Solidity: function setDropBurnMaxQuota(uint256 quo) returns()
func (_Dospayment *DospaymentTransactorSession) SetDropBurnMaxQuota(quo *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDropBurnMaxQuota(&_Dospayment.TransactOpts, quo)
}

// SetDropBurnToken is a paid mutator transaction binding the contract method 0xeac051f9.
//
// Solidity: function setDropBurnToken(address addr) returns()
func (_Dospayment *DospaymentTransactor) SetDropBurnToken(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setDropBurnToken", addr)
}

// SetDropBurnToken is a paid mutator transaction binding the contract method 0xeac051f9.
//
// Solidity: function setDropBurnToken(address addr) returns()
func (_Dospayment *DospaymentSession) SetDropBurnToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDropBurnToken(&_Dospayment.TransactOpts, addr)
}

// SetDropBurnToken is a paid mutator transaction binding the contract method 0xeac051f9.
//
// Solidity: function setDropBurnToken(address addr) returns()
func (_Dospayment *DospaymentTransactorSession) SetDropBurnToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDropBurnToken(&_Dospayment.TransactOpts, addr)
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

// DospaymentUpdateDropBurnMaxQuotaIterator is returned from FilterUpdateDropBurnMaxQuota and is used to iterate over the raw logs and unpacked data for UpdateDropBurnMaxQuota events raised by the Dospayment contract.
type DospaymentUpdateDropBurnMaxQuotaIterator struct {
	Event *DospaymentUpdateDropBurnMaxQuota // Event containing the contract specifics and raw log

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
func (it *DospaymentUpdateDropBurnMaxQuotaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentUpdateDropBurnMaxQuota)
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
		it.Event = new(DospaymentUpdateDropBurnMaxQuota)
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
func (it *DospaymentUpdateDropBurnMaxQuotaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentUpdateDropBurnMaxQuotaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentUpdateDropBurnMaxQuota represents a UpdateDropBurnMaxQuota event raised by the Dospayment contract.
type DospaymentUpdateDropBurnMaxQuota struct {
	OldQuota *big.Int
	NewQuota *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateDropBurnMaxQuota is a free log retrieval operation binding the contract event 0x0aee95cca46da64ee373e28dee5994361b4002c54035d92932c9825b76382e99.
//
// Solidity: event UpdateDropBurnMaxQuota(uint256 oldQuota, uint256 newQuota)
func (_Dospayment *DospaymentFilterer) FilterUpdateDropBurnMaxQuota(opts *bind.FilterOpts) (*DospaymentUpdateDropBurnMaxQuotaIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "UpdateDropBurnMaxQuota")
	if err != nil {
		return nil, err
	}
	return &DospaymentUpdateDropBurnMaxQuotaIterator{contract: _Dospayment.contract, event: "UpdateDropBurnMaxQuota", logs: logs, sub: sub}, nil
}

// WatchUpdateDropBurnMaxQuota is a free log subscription operation binding the contract event 0x0aee95cca46da64ee373e28dee5994361b4002c54035d92932c9825b76382e99.
//
// Solidity: event UpdateDropBurnMaxQuota(uint256 oldQuota, uint256 newQuota)
func (_Dospayment *DospaymentFilterer) WatchUpdateDropBurnMaxQuota(opts *bind.WatchOpts, sink chan<- *DospaymentUpdateDropBurnMaxQuota) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "UpdateDropBurnMaxQuota")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentUpdateDropBurnMaxQuota)
				if err := _Dospayment.contract.UnpackLog(event, "UpdateDropBurnMaxQuota", log); err != nil {
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

// DospaymentUpdateDropBurnTokenAddressIterator is returned from FilterUpdateDropBurnTokenAddress and is used to iterate over the raw logs and unpacked data for UpdateDropBurnTokenAddress events raised by the Dospayment contract.
type DospaymentUpdateDropBurnTokenAddressIterator struct {
	Event *DospaymentUpdateDropBurnTokenAddress // Event containing the contract specifics and raw log

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
func (it *DospaymentUpdateDropBurnTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentUpdateDropBurnTokenAddress)
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
		it.Event = new(DospaymentUpdateDropBurnTokenAddress)
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
func (it *DospaymentUpdateDropBurnTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentUpdateDropBurnTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentUpdateDropBurnTokenAddress represents a UpdateDropBurnTokenAddress event raised by the Dospayment contract.
type DospaymentUpdateDropBurnTokenAddress struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateDropBurnTokenAddress is a free log retrieval operation binding the contract event 0xfc8013dfb0c8d38f3bcab9239bd5712457c48919b272cdb109488549199a0173.
//
// Solidity: event UpdateDropBurnTokenAddress(address oldAddress, address newAddress)
func (_Dospayment *DospaymentFilterer) FilterUpdateDropBurnTokenAddress(opts *bind.FilterOpts) (*DospaymentUpdateDropBurnTokenAddressIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "UpdateDropBurnTokenAddress")
	if err != nil {
		return nil, err
	}
	return &DospaymentUpdateDropBurnTokenAddressIterator{contract: _Dospayment.contract, event: "UpdateDropBurnTokenAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateDropBurnTokenAddress is a free log subscription operation binding the contract event 0xfc8013dfb0c8d38f3bcab9239bd5712457c48919b272cdb109488549199a0173.
//
// Solidity: event UpdateDropBurnTokenAddress(address oldAddress, address newAddress)
func (_Dospayment *DospaymentFilterer) WatchUpdateDropBurnTokenAddress(opts *bind.WatchOpts, sink chan<- *DospaymentUpdateDropBurnTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "UpdateDropBurnTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentUpdateDropBurnTokenAddress)
				if err := _Dospayment.contract.UnpackLog(event, "UpdateDropBurnTokenAddress", log); err != nil {
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
