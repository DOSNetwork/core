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

// AskMeAnythingABI is the input ABI used to generate the binding from.
const AskMeAnythingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"generatedRandom\",\"type\":\"uint256\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"internalSerial\",\"type\":\"uint8\"}],\"name\":\"requestSafeRandom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastQueriedUrl\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMode\",\"type\":\"bool\"}],\"name\":\"setQueryMode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"random\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"queryId\",\"type\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"response\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"DOSWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastQueryInternalSerial\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repeatedCall\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRequestedRandom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"setTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastQueriedSelector\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"internalSerial\",\"type\":\"uint8\"},{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"selector\",\"type\":\"string\"}],\"name\":\"AMA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previousTimeout\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"name\":\"SetTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"queryId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"string\"}],\"name\":\"QueryResponseReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"internalSerial\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"succ\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"generatedRandom\",\"type\":\"uint256\"}],\"name\":\"RandomReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// AskMeAnythingBin is the compiled bytecode used for deploying new contracts.
const AskMeAnythingBin = `0x6080604052600280546001600160a01b031990811673848c0bb953755293230b705464654f647967639a179091556003805490911673214e79c85744cd2ebbc64ddc0047131496871bee1790556007805460ff19169055601c60085534801561006757600080fd5b50600080546001600160a01b03191633178155600254604080517f9d265e5800000000000000000000000000000000000000000000000000000000815290516001600160a01b039290921691639d265e5891600480820192602092909190829003018186803b1580156100d957600080fd5b505afa1580156100ed573d6000803e3d6000fd5b505050506040513d602081101561010357600080fd5b5051600354604080517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0380851660048301526000196024830152915193945091169163095ea7b3916044808201926020929091908290030181600087803b15801561017757600080fd5b505af115801561018b573d6000803e3d6000fd5b505050506040513d60208110156101a157600080fd5b5050600254604080517f9d265e5800000000000000000000000000000000000000000000000000000000815290516000926001600160a01b031691639d265e58916004808301926020929190829003018186803b15801561020157600080fd5b505afa158015610215573d6000803e3d6000fd5b505050506040513d602081101561022b57600080fd5b5051600354604080517fb73a3f8f0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b03928316602482015290519293509083169163b73a3f8f9160448082019260009290919082900301818387803b15801561029d57600080fd5b505af11580156102b1573d6000803e3d6000fd5b5050505050506112cf80620002c76000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c80638d05c4ca116100ad578063b8cf904e11610071578063b8cf904e1461031e578063c58a34cc14610326578063e3abd8b614610343578063e6eaadeb1461034b578063f2fde38b1461048257610121565b80638d05c4ca146102b05780638da5cb5b146102b85780638f32d59b146102dc578063a3a9bf9e146102f8578063ab8c1bad1461031657610121565b80635ec01e4d116100f45780635ec01e4d146102075780636d1129771461022157806370dea79a14610298578063715018a6146102a05780637a7f01a7146102a857610121565b806318a1908d1461012657806343c655831461014b57806349b03ca01461016b57806357ae678b146101e8575b600080fd5b6101496004803603604081101561013c57600080fd5b50803590602001356104a8565b005b6101496004803603602081101561016157600080fd5b503560ff166105ad565b610173610625565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ad578181015183820152602001610195565b50505050905090810190601f1680156101da5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610149600480360360208110156101fe57600080fd5b503515156106b3565b61020f6106d7565b60408051918252519081900360200190f35b6101496004803603604081101561023757600080fd5b8135919081019060408101602082013564010000000081111561025957600080fd5b82018360208201111561026b57600080fd5b8035906020019184600183028401116401000000008311171561028d57600080fd5b5090925090506106dd565b61020f61099d565b6101496109a3565b6101736109fc565b610149610a57565b6102c0610b6b565b604080516001600160a01b039092168252519081900360200190f35b6102e4610b7b565b604080519115158252519081900360200190f35b610300610b8c565b6040805160ff9092168252519081900360200190f35b6102e4610b95565b61020f610b9e565b6101496004803603602081101561033c57600080fd5b5035610ba4565b610173610bf7565b6101496004803603606081101561036157600080fd5b60ff823516919081019060408101602082013564010000000081111561038657600080fd5b82018360208201111561039857600080fd5b803590602001918460018302840111640100000000831117156103ba57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561040d57600080fd5b82018360208201111561041f57600080fd5b8035906020019184600183028401116401000000008311171561044157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c52945050505050565b6101496004803603602081101561049857600080fd5b50356001600160a01b0316610d58565b816104b1610d75565b6001600160a01b0316336001600160a01b03161461050357604051600160e51b62461bcd02815260040180806020018281038252602681526020018061125d6026913960400191505060405180910390fd5b60008181526006602052604090205460ff1661055357604051600160e51b62461bcd0281526004018080602001828103825260218152602001806112836021913960400191505060405180910390fd5b6005829055604080518481526020810184905281517fd0ecc71f8b5af397da9123fd2bff63c544c04af5c6935935a7f81e14b84522f2929181900390910190a150506000908152600660205260409020805460ff19169055565b600554600b5560006105be42610dee565b600081815260066020908152604091829020805460ff19166001908117909155825160ff8716815291820152808201839052905191925033917f15d9fcfed93b33c00d0175575d9069196dbab062132475484ac3d44e09163c519181900360600190a25050565b6009805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106ab5780601f10610680576101008083540402835291602001916106ab565b820191906000526020600020905b81548152906001019060200180831161068e57829003601f168201915b505050505081565b6106bb610b7b565b6106c457600080fd5b6007805460ff1916911515919091179055565b60055481565b826106e6610d75565b6001600160a01b0316336001600160a01b03161461073857604051600160e51b62461bcd02815260040180806020018281038252602681526020018061125d6026913960400191505060405180910390fd5b60008181526006602052604090205460ff1661078857604051600160e51b62461bcd0281526004018080602001828103825260218152602001806112836021913960400191505060405180910390fd5b61079460048484611156565b50604080518581526020810182815260048054600260001961010060018416150201909116049383018490527ffe1788dc549f39fbcdb06fdab5937f2d19ce5fe5616fa3be9311928bd5dabccb93889391929060608301908490801561083b5780601f106108105761010080835404028352916020019161083b565b820191906000526020600020905b81548152906001019060200180831161081e57829003601f168201915b5050935050505060405180910390a16000848152600660205260409020805460ff1916905560075460ff161561099757600c5460098054604080516020601f600260001960018716156101000201909516949094049384018190048102820181019092528281526109979460ff1693909290918301828280156108ff5780601f106108d4576101008083540402835291602001916108ff565b820191906000526020600020905b8154815290600101906020018083116108e257829003601f168201915b5050600a8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281529550919350915083018282801561098d5780601f106109625761010080835404028352916020019161098d565b820191906000526020600020905b81548152906001019060200180831161097057829003601f168201915b5050505050610c52565b50505050565b60085481565b6109ab610b7b565b6109b457600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106ab5780601f10610680576101008083540402835291602001916106ab565b610a5f610b7b565b610a6857600080fd5b60035460408051600160e01b6370a0823102815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b158015610ab657600080fd5b505afa158015610aca573d6000803e3d6000fd5b505050506040513d6020811015610ae057600080fd5b505160035460408051600160e01b63a9059cbb0281523360048201526024810184905290519293506001600160a01b039091169163a9059cbb916044808201926020929091908290030181600087803b158015610b3c57600080fd5b505af1158015610b50573d6000803e3d6000fd5b505050506040513d6020811015610b6657600080fd5b505050565b6000546001600160a01b03165b90565b6000546001600160a01b0316331490565b600c5460ff1681565b60075460ff1681565b600b5481565b610bac610b7b565b610bb557600080fd5b600854604080519182526020820183905280517f9aa0de0157c9133b911d2d811f590159622cea28cefe31505c203c828799da589281900390910190a1600855565b600a805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106ab5780601f10610680576101008083540402835291602001916106ab565b8151610c659060099060208501906111d4565b508051610c7990600a9060208401906111d4565b50600c805460ff191660ff8516179055600854600090610c9a908484610efe565b90508015610d0857600081815260066020908152604091829020805460ff19166001908117909155825160ff8816815291820152808201839052905133917f15d9fcfed93b33c00d0175575d9069196dbab062132475484ac3d44e09163c51919081900360600190a2610997565b60408051600160e51b62461bcd02815260206004820152601160248201527f496e76616c69642071756572792069642e000000000000000000000000000000604482015290519081900360640190fd5b610d60610b7b565b610d6957600080fd5b610d72816110e8565b50565b60025460408051600160e11b6321d39ecd02815290516000926001600160a01b0316916343a73d9a916004808301926020929190829003018186803b158015610dbd57600080fd5b505afa158015610dd1573d6000803e3d6000fd5b505050506040513d6020811015610de757600080fd5b5051905090565b60025460408051600160e11b6321d39ecd02815290516000926001600160a01b0316916343a73d9a916004808301926020929190829003018186803b158015610e3657600080fd5b505afa158015610e4a573d6000803e3d6000fd5b505050506040513d6020811015610e6057600080fd5b5051600180546001600160a01b0319166001600160a01b03928316179081905560408051600160e01b63c7c3f4a5028152306004820152602481018690529051919092169163c7c3f4a59160448083019260209291908290030181600087803b158015610ecc57600080fd5b505af1158015610ee0573d6000803e3d6000fd5b505050506040513d6020811015610ef657600080fd5b505192915050565b60025460408051600160e11b6321d39ecd02815290516000926001600160a01b0316916343a73d9a916004808301926020929190829003018186803b158015610f4657600080fd5b505afa158015610f5a573d6000803e3d6000fd5b505050506040513d6020811015610f7057600080fd5b5051600180546001600160a01b0319166001600160a01b039283161790819055604051600160e01b63b7fb8fd7028152306004820181815260248301899052608060448401908152885160848501528851949095169463b7fb8fd79492938a938a938a9390929091606482019160a40190602087019080838360005b83811015611004578181015183820152602001610fec565b50505050905090810190601f1680156110315780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561106457818101518382015260200161104c565b50505050905090810190601f1680156110915780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b1580156110b457600080fd5b505af11580156110c8573d6000803e3d6000fd5b505050506040513d60208110156110de57600080fd5b5051949350505050565b6001600160a01b0381166110fb57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106111975782800160ff198235161785556111c4565b828001600101855582156111c4579182015b828111156111c45782358255916020019190600101906111a9565b506111d0929150611242565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061121557805160ff19168380011785556111c4565b828001600101855582156111c4579182015b828111156111c4578251825591602001919060010190611227565b610b7891905b808211156111d0576000815560010161124856fe556e61757468656e7469636174656420726573706f6e73652066726f6d206e6f6e2d444f532e526573706f6e7365207769746820696e76616c6964207265717565737420696421a165627a7a7230582027be8a1b2c4d262a45309bb2fa79f12d9f43c27aea037df2e267cb8091a6ec7f0029`

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

// LastQueryInternalSerial is a free data retrieval call binding the contract method 0xa3a9bf9e.
//
// Solidity: function lastQueryInternalSerial() constant returns(uint8)
func (_AskMeAnything *AskMeAnythingCaller) LastQueryInternalSerial(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AskMeAnything.contract.Call(opts, out, "lastQueryInternalSerial")
	return *ret0, err
}

// LastQueryInternalSerial is a free data retrieval call binding the contract method 0xa3a9bf9e.
//
// Solidity: function lastQueryInternalSerial() constant returns(uint8)
func (_AskMeAnything *AskMeAnythingSession) LastQueryInternalSerial() (uint8, error) {
	return _AskMeAnything.Contract.LastQueryInternalSerial(&_AskMeAnything.CallOpts)
}

// LastQueryInternalSerial is a free data retrieval call binding the contract method 0xa3a9bf9e.
//
// Solidity: function lastQueryInternalSerial() constant returns(uint8)
func (_AskMeAnything *AskMeAnythingCallerSession) LastQueryInternalSerial() (uint8, error) {
	return _AskMeAnything.Contract.LastQueryInternalSerial(&_AskMeAnything.CallOpts)
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

// AMA is a paid mutator transaction binding the contract method 0xe6eaadeb.
//
// Solidity: function AMA(internalSerial uint8, url string, selector string) returns()
func (_AskMeAnything *AskMeAnythingTransactor) AMA(opts *bind.TransactOpts, internalSerial uint8, url string, selector string) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "AMA", internalSerial, url, selector)
}

// AMA is a paid mutator transaction binding the contract method 0xe6eaadeb.
//
// Solidity: function AMA(internalSerial uint8, url string, selector string) returns()
func (_AskMeAnything *AskMeAnythingSession) AMA(internalSerial uint8, url string, selector string) (*types.Transaction, error) {
	return _AskMeAnything.Contract.AMA(&_AskMeAnything.TransactOpts, internalSerial, url, selector)
}

// AMA is a paid mutator transaction binding the contract method 0xe6eaadeb.
//
// Solidity: function AMA(internalSerial uint8, url string, selector string) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) AMA(internalSerial uint8, url string, selector string) (*types.Transaction, error) {
	return _AskMeAnything.Contract.AMA(&_AskMeAnything.TransactOpts, internalSerial, url, selector)
}

// DOSWithdraw is a paid mutator transaction binding the contract method 0x8d05c4ca.
//
// Solidity: function DOSWithdraw() returns()
func (_AskMeAnything *AskMeAnythingTransactor) DOSWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "DOSWithdraw")
}

// DOSWithdraw is a paid mutator transaction binding the contract method 0x8d05c4ca.
//
// Solidity: function DOSWithdraw() returns()
func (_AskMeAnything *AskMeAnythingSession) DOSWithdraw() (*types.Transaction, error) {
	return _AskMeAnything.Contract.DOSWithdraw(&_AskMeAnything.TransactOpts)
}

// DOSWithdraw is a paid mutator transaction binding the contract method 0x8d05c4ca.
//
// Solidity: function DOSWithdraw() returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) DOSWithdraw() (*types.Transaction, error) {
	return _AskMeAnything.Contract.DOSWithdraw(&_AskMeAnything.TransactOpts)
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

// RequestSafeRandom is a paid mutator transaction binding the contract method 0x43c65583.
//
// Solidity: function requestSafeRandom(internalSerial uint8) returns()
func (_AskMeAnything *AskMeAnythingTransactor) RequestSafeRandom(opts *bind.TransactOpts, internalSerial uint8) (*types.Transaction, error) {
	return _AskMeAnything.contract.Transact(opts, "requestSafeRandom", internalSerial)
}

// RequestSafeRandom is a paid mutator transaction binding the contract method 0x43c65583.
//
// Solidity: function requestSafeRandom(internalSerial uint8) returns()
func (_AskMeAnything *AskMeAnythingSession) RequestSafeRandom(internalSerial uint8) (*types.Transaction, error) {
	return _AskMeAnything.Contract.RequestSafeRandom(&_AskMeAnything.TransactOpts, internalSerial)
}

// RequestSafeRandom is a paid mutator transaction binding the contract method 0x43c65583.
//
// Solidity: function requestSafeRandom(internalSerial uint8) returns()
func (_AskMeAnything *AskMeAnythingTransactorSession) RequestSafeRandom(internalSerial uint8) (*types.Transaction, error) {
	return _AskMeAnything.Contract.RequestSafeRandom(&_AskMeAnything.TransactOpts, internalSerial)
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
	MsgSender      common.Address
	InternalSerial uint8
	Succ           bool
	RequestId      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x15d9fcfed93b33c00d0175575d9069196dbab062132475484ac3d44e09163c51.
//
// Solidity: e RequestSent(msgSender indexed address, internalSerial uint8, succ bool, requestId uint256)
func (_AskMeAnything *AskMeAnythingFilterer) FilterRequestSent(opts *bind.FilterOpts, msgSender []common.Address) (*AskMeAnythingRequestSentIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _AskMeAnything.contract.FilterLogs(opts, "RequestSent", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return &AskMeAnythingRequestSentIterator{contract: _AskMeAnything.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x15d9fcfed93b33c00d0175575d9069196dbab062132475484ac3d44e09163c51.
//
// Solidity: e RequestSent(msgSender indexed address, internalSerial uint8, succ bool, requestId uint256)
func (_AskMeAnything *AskMeAnythingFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *AskMeAnythingRequestSent, msgSender []common.Address) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _AskMeAnything.contract.WatchLogs(opts, "RequestSent", msgSenderRule)
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
const DOSAddressBridgeInterfaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPaymentAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// GetPaymentAddress is a free data retrieval call binding the contract method 0x9d265e58.
//
// Solidity: function getPaymentAddress() constant returns(address)
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceCaller) GetPaymentAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DOSAddressBridgeInterface.contract.Call(opts, out, "getPaymentAddress")
	return *ret0, err
}

// GetPaymentAddress is a free data retrieval call binding the contract method 0x9d265e58.
//
// Solidity: function getPaymentAddress() constant returns(address)
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceSession) GetPaymentAddress() (common.Address, error) {
	return _DOSAddressBridgeInterface.Contract.GetPaymentAddress(&_DOSAddressBridgeInterface.CallOpts)
}

// GetPaymentAddress is a free data retrieval call binding the contract method 0x9d265e58.
//
// Solidity: function getPaymentAddress() constant returns(address)
func (_DOSAddressBridgeInterface *DOSAddressBridgeInterfaceCallerSession) GetPaymentAddress() (common.Address, error) {
	return _DOSAddressBridgeInterface.Contract.GetPaymentAddress(&_DOSAddressBridgeInterface.CallOpts)
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
const DOSOnChainSDKABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"generatedRandom\",\"type\":\"uint256\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"queryId\",\"type\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"DOSWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DOSOnChainSDKBin is the compiled bytecode used for deploying new contracts.
const DOSOnChainSDKBin = `0x6080604052600280546001600160a01b031990811673848c0bb953755293230b705464654f647967639a179091556003805490911673214e79c85744cd2ebbc64ddc0047131496871bee17905534801561005857600080fd5b50600080546001600160a01b03191633178155600254604080517f9d265e5800000000000000000000000000000000000000000000000000000000815290516001600160a01b039290921691639d265e5891600480820192602092909190829003018186803b1580156100ca57600080fd5b505afa1580156100de573d6000803e3d6000fd5b505050506040513d60208110156100f457600080fd5b5051600354604080517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0380851660048301526000196024830152915193945091169163095ea7b3916044808201926020929091908290030181600087803b15801561016857600080fd5b505af115801561017c573d6000803e3d6000fd5b505050506040513d602081101561019257600080fd5b5050600254604080517f9d265e5800000000000000000000000000000000000000000000000000000000815290516000926001600160a01b031691639d265e58916004808301926020929190829003018186803b1580156101f257600080fd5b505afa158015610206573d6000803e3d6000fd5b505050506040513d602081101561021c57600080fd5b5051600354604080517fb73a3f8f0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b03928316602482015290519293509083169163b73a3f8f9160448082019260009290919082900301818387803b15801561028e57600080fd5b505af11580156102a2573d6000803e3d6000fd5b5050505050506103dc806102b76000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638d05c4ca1161005b5780638d05c4ca146101265780638da5cb5b1461012e5780638f32d59b14610152578063f2fde38b1461016e5761007d565b806318a1908d146100825780636d112977146100a7578063715018a61461011e575b600080fd5b6100a56004803603604081101561009857600080fd5b5080359060200135610194565b005b6100a5600480360360408110156100bd57600080fd5b813591908101906040810160208201356401000000008111156100df57600080fd5b8201836020820111156100f157600080fd5b8035906020019184600183028401116401000000008311171561011357600080fd5b509092509050610198565b6100a561019d565b6100a56101f6565b610136610305565b604080516001600160a01b039092168252519081900360200190f35b61015a610314565b604080519115158252519081900360200190f35b6100a56004803603602081101561018457600080fd5b50356001600160a01b0316610325565b5050565b505050565b6101a5610314565b6101ae57600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6101fe610314565b61020757600080fd5b60035460408051600160e01b6370a0823102815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b15801561025557600080fd5b505afa158015610269573d6000803e3d6000fd5b505050506040513d602081101561027f57600080fd5b505160035460408051600160e01b63a9059cbb0281523360048201526024810184905290519293506001600160a01b039091169163a9059cbb916044808201926020929091908290030181600087803b1580156102db57600080fd5b505af11580156102ef573d6000803e3d6000fd5b505050506040513d602081101561019857600080fd5b6000546001600160a01b031690565b6000546001600160a01b0316331490565b61032d610314565b61033657600080fd5b61033f81610342565b50565b6001600160a01b03811661035557600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea165627a7a72305820c431218f1ec13e8b0af34a00f6c5676e732700542ea5988a537460623ab8787e0029`

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

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DOSOnChainSDK *DOSOnChainSDKCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DOSOnChainSDK.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DOSOnChainSDK *DOSOnChainSDKSession) IsOwner() (bool, error) {
	return _DOSOnChainSDK.Contract.IsOwner(&_DOSOnChainSDK.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DOSOnChainSDK *DOSOnChainSDKCallerSession) IsOwner() (bool, error) {
	return _DOSOnChainSDK.Contract.IsOwner(&_DOSOnChainSDK.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DOSOnChainSDK *DOSOnChainSDKCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DOSOnChainSDK.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DOSOnChainSDK *DOSOnChainSDKSession) Owner() (common.Address, error) {
	return _DOSOnChainSDK.Contract.Owner(&_DOSOnChainSDK.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DOSOnChainSDK *DOSOnChainSDKCallerSession) Owner() (common.Address, error) {
	return _DOSOnChainSDK.Contract.Owner(&_DOSOnChainSDK.CallOpts)
}

// DOSWithdraw is a paid mutator transaction binding the contract method 0x8d05c4ca.
//
// Solidity: function DOSWithdraw() returns()
func (_DOSOnChainSDK *DOSOnChainSDKTransactor) DOSWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSOnChainSDK.contract.Transact(opts, "DOSWithdraw")
}

// DOSWithdraw is a paid mutator transaction binding the contract method 0x8d05c4ca.
//
// Solidity: function DOSWithdraw() returns()
func (_DOSOnChainSDK *DOSOnChainSDKSession) DOSWithdraw() (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.DOSWithdraw(&_DOSOnChainSDK.TransactOpts)
}

// DOSWithdraw is a paid mutator transaction binding the contract method 0x8d05c4ca.
//
// Solidity: function DOSWithdraw() returns()
func (_DOSOnChainSDK *DOSOnChainSDKTransactorSession) DOSWithdraw() (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.DOSWithdraw(&_DOSOnChainSDK.TransactOpts)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DOSOnChainSDK *DOSOnChainSDKTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSOnChainSDK.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DOSOnChainSDK *DOSOnChainSDKSession) RenounceOwnership() (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.RenounceOwnership(&_DOSOnChainSDK.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DOSOnChainSDK *DOSOnChainSDKTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.RenounceOwnership(&_DOSOnChainSDK.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DOSOnChainSDK *DOSOnChainSDKTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DOSOnChainSDK.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DOSOnChainSDK *DOSOnChainSDKSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.TransferOwnership(&_DOSOnChainSDK.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_DOSOnChainSDK *DOSOnChainSDKTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DOSOnChainSDK.Contract.TransferOwnership(&_DOSOnChainSDK.TransactOpts, newOwner)
}

// DOSOnChainSDKOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the DOSOnChainSDK contract.
type DOSOnChainSDKOwnershipRenouncedIterator struct {
	Event *DOSOnChainSDKOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DOSOnChainSDKOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSOnChainSDKOwnershipRenounced)
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
		it.Event = new(DOSOnChainSDKOwnershipRenounced)
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
func (it *DOSOnChainSDKOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSOnChainSDKOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSOnChainSDKOwnershipRenounced represents a OwnershipRenounced event raised by the DOSOnChainSDK contract.
type DOSOnChainSDKOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DOSOnChainSDK *DOSOnChainSDKFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DOSOnChainSDKOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DOSOnChainSDK.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DOSOnChainSDKOwnershipRenouncedIterator{contract: _DOSOnChainSDK.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_DOSOnChainSDK *DOSOnChainSDKFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DOSOnChainSDKOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _DOSOnChainSDK.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSOnChainSDKOwnershipRenounced)
				if err := _DOSOnChainSDK.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DOSOnChainSDKOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DOSOnChainSDK contract.
type DOSOnChainSDKOwnershipTransferredIterator struct {
	Event *DOSOnChainSDKOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DOSOnChainSDKOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSOnChainSDKOwnershipTransferred)
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
		it.Event = new(DOSOnChainSDKOwnershipTransferred)
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
func (it *DOSOnChainSDKOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSOnChainSDKOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSOnChainSDKOwnershipTransferred represents a OwnershipTransferred event raised by the DOSOnChainSDK contract.
type DOSOnChainSDKOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DOSOnChainSDK *DOSOnChainSDKFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DOSOnChainSDKOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DOSOnChainSDK.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DOSOnChainSDKOwnershipTransferredIterator{contract: _DOSOnChainSDK.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_DOSOnChainSDK *DOSOnChainSDKFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DOSOnChainSDKOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DOSOnChainSDK.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSOnChainSDKOwnershipTransferred)
				if err := _DOSOnChainSDK.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// DOSPaymentInterfaceABI is the input ABI used to generate the binding from.
const DOSPaymentInterfaceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setPaymentMethod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DOSPaymentInterfaceBin is the compiled bytecode used for deploying new contracts.
const DOSPaymentInterfaceBin = `0x`

// DeployDOSPaymentInterface deploys a new Ethereum contract, binding an instance of DOSPaymentInterface to it.
func DeployDOSPaymentInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DOSPaymentInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSPaymentInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DOSPaymentInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DOSPaymentInterface{DOSPaymentInterfaceCaller: DOSPaymentInterfaceCaller{contract: contract}, DOSPaymentInterfaceTransactor: DOSPaymentInterfaceTransactor{contract: contract}, DOSPaymentInterfaceFilterer: DOSPaymentInterfaceFilterer{contract: contract}}, nil
}

// DOSPaymentInterface is an auto generated Go binding around an Ethereum contract.
type DOSPaymentInterface struct {
	DOSPaymentInterfaceCaller     // Read-only binding to the contract
	DOSPaymentInterfaceTransactor // Write-only binding to the contract
	DOSPaymentInterfaceFilterer   // Log filterer for contract events
}

// DOSPaymentInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DOSPaymentInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSPaymentInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DOSPaymentInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSPaymentInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DOSPaymentInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSPaymentInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DOSPaymentInterfaceSession struct {
	Contract     *DOSPaymentInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DOSPaymentInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DOSPaymentInterfaceCallerSession struct {
	Contract *DOSPaymentInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DOSPaymentInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DOSPaymentInterfaceTransactorSession struct {
	Contract     *DOSPaymentInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DOSPaymentInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DOSPaymentInterfaceRaw struct {
	Contract *DOSPaymentInterface // Generic contract binding to access the raw methods on
}

// DOSPaymentInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DOSPaymentInterfaceCallerRaw struct {
	Contract *DOSPaymentInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// DOSPaymentInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DOSPaymentInterfaceTransactorRaw struct {
	Contract *DOSPaymentInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDOSPaymentInterface creates a new instance of DOSPaymentInterface, bound to a specific deployed contract.
func NewDOSPaymentInterface(address common.Address, backend bind.ContractBackend) (*DOSPaymentInterface, error) {
	contract, err := bindDOSPaymentInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DOSPaymentInterface{DOSPaymentInterfaceCaller: DOSPaymentInterfaceCaller{contract: contract}, DOSPaymentInterfaceTransactor: DOSPaymentInterfaceTransactor{contract: contract}, DOSPaymentInterfaceFilterer: DOSPaymentInterfaceFilterer{contract: contract}}, nil
}

// NewDOSPaymentInterfaceCaller creates a new read-only instance of DOSPaymentInterface, bound to a specific deployed contract.
func NewDOSPaymentInterfaceCaller(address common.Address, caller bind.ContractCaller) (*DOSPaymentInterfaceCaller, error) {
	contract, err := bindDOSPaymentInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DOSPaymentInterfaceCaller{contract: contract}, nil
}

// NewDOSPaymentInterfaceTransactor creates a new write-only instance of DOSPaymentInterface, bound to a specific deployed contract.
func NewDOSPaymentInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*DOSPaymentInterfaceTransactor, error) {
	contract, err := bindDOSPaymentInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DOSPaymentInterfaceTransactor{contract: contract}, nil
}

// NewDOSPaymentInterfaceFilterer creates a new log filterer instance of DOSPaymentInterface, bound to a specific deployed contract.
func NewDOSPaymentInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*DOSPaymentInterfaceFilterer, error) {
	contract, err := bindDOSPaymentInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DOSPaymentInterfaceFilterer{contract: contract}, nil
}

// bindDOSPaymentInterface binds a generic wrapper to an already deployed contract.
func bindDOSPaymentInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSPaymentInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSPaymentInterface *DOSPaymentInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSPaymentInterface.Contract.DOSPaymentInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSPaymentInterface *DOSPaymentInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSPaymentInterface.Contract.DOSPaymentInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSPaymentInterface *DOSPaymentInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSPaymentInterface.Contract.DOSPaymentInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSPaymentInterface *DOSPaymentInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSPaymentInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSPaymentInterface *DOSPaymentInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSPaymentInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSPaymentInterface *DOSPaymentInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSPaymentInterface.Contract.contract.Transact(opts, method, params...)
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken( address) constant returns(bool)
func (_DOSPaymentInterface *DOSPaymentInterfaceCaller) IsSupportedToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DOSPaymentInterface.contract.Call(opts, out, "isSupportedToken", arg0)
	return *ret0, err
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken( address) constant returns(bool)
func (_DOSPaymentInterface *DOSPaymentInterfaceSession) IsSupportedToken(arg0 common.Address) (bool, error) {
	return _DOSPaymentInterface.Contract.IsSupportedToken(&_DOSPaymentInterface.CallOpts, arg0)
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken( address) constant returns(bool)
func (_DOSPaymentInterface *DOSPaymentInterfaceCallerSession) IsSupportedToken(arg0 common.Address) (bool, error) {
	return _DOSPaymentInterface.Contract.IsSupportedToken(&_DOSPaymentInterface.CallOpts, arg0)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(consumer address, tokenAddr address) returns()
func (_DOSPaymentInterface *DOSPaymentInterfaceTransactor) SetPaymentMethod(opts *bind.TransactOpts, consumer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _DOSPaymentInterface.contract.Transact(opts, "setPaymentMethod", consumer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(consumer address, tokenAddr address) returns()
func (_DOSPaymentInterface *DOSPaymentInterfaceSession) SetPaymentMethod(consumer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _DOSPaymentInterface.Contract.SetPaymentMethod(&_DOSPaymentInterface.TransactOpts, consumer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(consumer address, tokenAddr address) returns()
func (_DOSPaymentInterface *DOSPaymentInterfaceTransactorSession) SetPaymentMethod(consumer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _DOSPaymentInterface.Contract.SetPaymentMethod(&_DOSPaymentInterface.TransactOpts, consumer, tokenAddr)
}

// DOSProxyInterfaceABI is the input ABI used to generate the binding from.
const DOSProxyInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestRandom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// RequestRandom is a paid mutator transaction binding the contract method 0xc7c3f4a5.
//
// Solidity: function requestRandom( address,  uint256) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceTransactor) RequestRandom(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _DOSProxyInterface.contract.Transact(opts, "requestRandom", arg0, arg1)
}

// RequestRandom is a paid mutator transaction binding the contract method 0xc7c3f4a5.
//
// Solidity: function requestRandom( address,  uint256) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceSession) RequestRandom(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.RequestRandom(&_DOSProxyInterface.TransactOpts, arg0, arg1)
}

// RequestRandom is a paid mutator transaction binding the contract method 0xc7c3f4a5.
//
// Solidity: function requestRandom( address,  uint256) returns(uint256)
func (_DOSProxyInterface *DOSProxyInterfaceTransactorSession) RequestRandom(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _DOSProxyInterface.Contract.RequestRandom(&_DOSProxyInterface.TransactOpts, arg0, arg1)
}

// ERC20IABI is the input ABI used to generate the binding from.
const ERC20IABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20IBin is the compiled bytecode used for deploying new contracts.
const ERC20IBin = `0x`

// DeployERC20I deploys a new Ethereum contract, binding an instance of ERC20I to it.
func DeployERC20I(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20I, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20IABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20IBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20I{ERC20ICaller: ERC20ICaller{contract: contract}, ERC20ITransactor: ERC20ITransactor{contract: contract}, ERC20IFilterer: ERC20IFilterer{contract: contract}}, nil
}

// ERC20I is an auto generated Go binding around an Ethereum contract.
type ERC20I struct {
	ERC20ICaller     // Read-only binding to the contract
	ERC20ITransactor // Write-only binding to the contract
	ERC20IFilterer   // Log filterer for contract events
}

// ERC20ICaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20ICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ITransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20ITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20IFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20IFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20ISession struct {
	Contract     *ERC20I           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20ICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20ICallerSession struct {
	Contract *ERC20ICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20ITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20ITransactorSession struct {
	Contract     *ERC20ITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20IRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20IRaw struct {
	Contract *ERC20I // Generic contract binding to access the raw methods on
}

// ERC20ICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20ICallerRaw struct {
	Contract *ERC20ICaller // Generic read-only contract binding to access the raw methods on
}

// ERC20ITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20ITransactorRaw struct {
	Contract *ERC20ITransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20I creates a new instance of ERC20I, bound to a specific deployed contract.
func NewERC20I(address common.Address, backend bind.ContractBackend) (*ERC20I, error) {
	contract, err := bindERC20I(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20I{ERC20ICaller: ERC20ICaller{contract: contract}, ERC20ITransactor: ERC20ITransactor{contract: contract}, ERC20IFilterer: ERC20IFilterer{contract: contract}}, nil
}

// NewERC20ICaller creates a new read-only instance of ERC20I, bound to a specific deployed contract.
func NewERC20ICaller(address common.Address, caller bind.ContractCaller) (*ERC20ICaller, error) {
	contract, err := bindERC20I(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ICaller{contract: contract}, nil
}

// NewERC20ITransactor creates a new write-only instance of ERC20I, bound to a specific deployed contract.
func NewERC20ITransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20ITransactor, error) {
	contract, err := bindERC20I(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ITransactor{contract: contract}, nil
}

// NewERC20IFilterer creates a new log filterer instance of ERC20I, bound to a specific deployed contract.
func NewERC20IFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20IFilterer, error) {
	contract, err := bindERC20I(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20IFilterer{contract: contract}, nil
}

// bindERC20I binds a generic wrapper to an already deployed contract.
func bindERC20I(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20IABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20I *ERC20IRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20I.Contract.ERC20ICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20I *ERC20IRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20I.Contract.ERC20ITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20I *ERC20IRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20I.Contract.ERC20ITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20I *ERC20ICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20I.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20I *ERC20ITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20I.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20I *ERC20ITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20I.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20I *ERC20ICaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20I.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20I *ERC20ISession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20I.Contract.BalanceOf(&_ERC20I.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20I *ERC20ICallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20I.Contract.BalanceOf(&_ERC20I.CallOpts, who)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20I *ERC20ITransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20I.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20I *ERC20ISession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20I.Contract.Approve(&_ERC20I.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20I *ERC20ITransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20I.Contract.Approve(&_ERC20I.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20I *ERC20ITransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20I.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20I *ERC20ISession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20I.Contract.Transfer(&_ERC20I.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20I *ERC20ITransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20I.Contract.Transfer(&_ERC20I.TransactOpts, to, value)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556101f1806100326000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063715018a6146100515780638da5cb5b1461005b5780638f32d59b1461007f578063f2fde38b1461009b575b600080fd5b6100596100c1565b005b61006361011a565b604080516001600160a01b039092168252519081900360200190f35b610087610129565b604080519115158252519081900360200190f35b610059600480360360208110156100b157600080fd5b50356001600160a01b031661013a565b6100c9610129565b6100d257600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b610142610129565b61014b57600080fd5b61015481610157565b50565b6001600160a01b03811661016a57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea165627a7a72305820dff661e60575220200efb44e916362b46d15d7e6717d1bcac8dc96c1350a68920029`

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
const UtilsBin = `0x604c6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a723058209b06828082d13c1c3b270b98bb5a58aaa04dd951ff02a1d2ff4e1aaad37ff5be0029`

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
