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

// BN256ABI is the input ABI used to generate the binding from.
const BN256ABI = "[]"

// BN256Bin is the compiled bytecode used for deploying new contracts.
const BN256Bin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058208e2fb6971c6afb724cf3836d0e3b07e5cc080c845b1c927bb37831339fe2d72d0029`

// DeployBN256 deploys a new Ethereum contract, binding an instance of BN256 to it.
func DeployBN256(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BN256, error) {
	parsed, err := abi.JSON(strings.NewReader(BN256ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BN256Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BN256{BN256Caller: BN256Caller{contract: contract}, BN256Transactor: BN256Transactor{contract: contract}, BN256Filterer: BN256Filterer{contract: contract}}, nil
}

// BN256 is an auto generated Go binding around an Ethereum contract.
type BN256 struct {
	BN256Caller     // Read-only binding to the contract
	BN256Transactor // Write-only binding to the contract
	BN256Filterer   // Log filterer for contract events
}

// BN256Caller is an auto generated read-only Go binding around an Ethereum contract.
type BN256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BN256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BN256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BN256Session struct {
	Contract     *BN256            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN256CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BN256CallerSession struct {
	Contract *BN256Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BN256TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BN256TransactorSession struct {
	Contract     *BN256Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN256Raw is an auto generated low-level Go binding around an Ethereum contract.
type BN256Raw struct {
	Contract *BN256 // Generic contract binding to access the raw methods on
}

// BN256CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BN256CallerRaw struct {
	Contract *BN256Caller // Generic read-only contract binding to access the raw methods on
}

// BN256TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BN256TransactorRaw struct {
	Contract *BN256Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBN256 creates a new instance of BN256, bound to a specific deployed contract.
func NewBN256(address common.Address, backend bind.ContractBackend) (*BN256, error) {
	contract, err := bindBN256(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BN256{BN256Caller: BN256Caller{contract: contract}, BN256Transactor: BN256Transactor{contract: contract}, BN256Filterer: BN256Filterer{contract: contract}}, nil
}

// NewBN256Caller creates a new read-only instance of BN256, bound to a specific deployed contract.
func NewBN256Caller(address common.Address, caller bind.ContractCaller) (*BN256Caller, error) {
	contract, err := bindBN256(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BN256Caller{contract: contract}, nil
}

// NewBN256Transactor creates a new write-only instance of BN256, bound to a specific deployed contract.
func NewBN256Transactor(address common.Address, transactor bind.ContractTransactor) (*BN256Transactor, error) {
	contract, err := bindBN256(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BN256Transactor{contract: contract}, nil
}

// NewBN256Filterer creates a new log filterer instance of BN256, bound to a specific deployed contract.
func NewBN256Filterer(address common.Address, filterer bind.ContractFilterer) (*BN256Filterer, error) {
	contract, err := bindBN256(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BN256Filterer{contract: contract}, nil
}

// bindBN256 binds a generic wrapper to an already deployed contract.
func bindBN256(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BN256ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN256 *BN256Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BN256.Contract.BN256Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN256 *BN256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN256.Contract.BN256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN256 *BN256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN256.Contract.BN256Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN256 *BN256CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BN256.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN256 *BN256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN256 *BN256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN256.Contract.contract.Transact(opts, method, params...)
}

// DOSProxyABI is the input ABI used to generate the binding from.
const DOSProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"sig\",\"type\":\"uint256[2]\"}],\"name\":\"updateRandomness\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"blkNum\",\"type\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\"},{\"name\":\"queryType\",\"type\":\"string\"},{\"name\":\"queryPath\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"x2\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"},{\"name\":\"y2\",\"type\":\"uint256\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getGroupPubKey\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"queryId\",\"type\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\"},{\"name\":\"sig\",\"type\":\"uint256[2]\"}],\"name\":\"triggerCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fireRandom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"grouping\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRandomness\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uploadNodeId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdatedBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"queryId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"timeout\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"randomness\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dispatchedGroup\",\"type\":\"uint256[4]\"}],\"name\":\"LogUrl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"queryType\",\"type\":\"string\"}],\"name\":\"LogNonSupportedType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"}],\"name\":\"LogNonContractCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"callbackAddr\",\"type\":\"address\"}],\"name\":\"LogCallbackTriggeredFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogQueryFromNonExistentUC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lastRandomness\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastUpdatedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dispatchedGroup\",\"type\":\"uint256[4]\"}],\"name\":\"LogUpdateRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"trafficType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"name\":\"pubKey\",\"type\":\"uint256[4]\"}],\"name\":\"LogInvalidSignatureFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogInsufficientGroupNumber\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"NodeId\",\"type\":\"uint256[]\"}],\"name\":\"LogGrouping\",\"type\":\"event\"}]"

// DOSProxyBin is the compiled bytecode used for deploying new contracts.
const DOSProxyBin = `0x608060405234801561001057600080fd5b506119c2806100206000396000f3006080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166309ac86d381146100b3578063482edfaa146100cc57806379a924c51461011f57806392021653146101405780639261cc7e14610190578063b181a8fc146101b6578063ea5cba13146101cb578063eab14fe6146101e0578063f2a3072d146101f8578063f89a15f71461020d578063f90ce5ba14610225575b600080fd5b3480156100bf57600080fd5b506100ca600461023a565b005b3480156100d857600080fd5b5061010d60048035600160a060020a031690602480359160443591606435808201929081013591608435908101910135610459565b60408051918252519081900360200190f35b34801561012b57600080fd5b506100ca600435602435604435606435610835565b34801561014c57600080fd5b506101586004356109fe565b6040518082608080838360005b8381101561017d578181015183820152602001610165565b5050505090500191505060405180910390f35b34801561019c57600080fd5b506100ca6004803590602480359081019101356044610b22565b3480156101c257600080fd5b506100ca610d77565b3480156101d757600080fd5b506100ca610d94565b3480156101ec57600080fd5b506100ca600435610eeb565b34801561020457600080fd5b5061010d61102d565b34801561021957600080fd5b506100ca600435611033565b34801561023157600080fd5b5061010d611065565b6040805180820190915260045440815260055460208201526000906102f1906001906102659061106b565b604080518082018252863581526020808801359082015281516080810180845291929091600691839190820190839060029082845b81548152602001906001019080831161029a57505050918352505060408051808201918290526020909201919060028481019182845b8154815260200190600101908083116102d0575050505050815250506110a5565b15156102fc57610455565b6040805183356020828101919091528085013582840152825180830384018152606090920192839052815191929182918401908083835b602083106103525780518252601f199092019160209182019101610333565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120600581905543600455600254935091505081151561039557fe5b0690506002818154811015156103a757fe5b6000918252602090912060049091020160066103c5818360026117cf565b506103d8600282810190848101906117cf565b509050507fb57d39816b5fafa35e9395c99afd19edb72160f491310ba0942ea8c58c8bc9f960055460045461040c846109fe565b6040518084815260200183815260200182600460200280838360005b83811015610440578181015183820152602001610428565b50505050905001935050505060405180910390a15b5050565b6000806000806104688b611353565b11156107e7576104dc87878080601f016020809104026020016040519081016040528093929190818152602001838380828437505060408051808201909152600381527f41504900000000000000000000000000000000000000000000000000000000006020820152935061135792505050565b1561078e57898989898989896040516020018088600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401878152602001868152602001858580828437820191505083838082843782019150509750505050505050506040516020818303038152906040526040518082805190602001908083835b6020831061057f5780518252601f199092019160209182019101610560565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912060025460055491965093509150508115156105bf57fe5b0690506060604051908101604052808381526020016002838154811015156105e357fe5b6000918252602090912060408051608081018083529093600402909201918391820190839060029082845b81548152602001906001019080831161060e57505050918352505060408051808201918290526020909201919060028481019182845b81548152602001906001019080831161064457505050919092525050508152600160a060020a038c1660209182015260008481526001808352604090912083518155918301518051909183019061069e908290600261180a565b5060208201516106b4906002808401919061180a565b50505060408201518160050160006101000a815481600160a060020a030219169083600160a060020a031602179055509050507f75b4fd18552b0d79989d6cc5de124544de30dc29b548e1f743a719ce5ad443298286868b600554610718876109fe565b604051808781526020018060200185815260200184815260200183600460200280838360005b8381101561075657818101518382015260200161073e565b5050505090500182810382528787828181526020019250808284376040519201829003995090975050505050505050a1819250610828565b7f70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b4187876040518080602001828103825284848281815260200192508082843760405192018290039550909350505050a160009250610828565b60408051600160a060020a038c16815290517f6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb59181900360200190a1600092505b5050979650505050505050565b600084848484604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040526040518082805190602001908083835b6020831061089e5780518252601f19909201916020918201910161087f565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008181526003909252929020549194505060ff1615915061094b905057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f67726f75702068617320616c7265616479207265676973746572656400000000604482015290519081900360640190fd5b604080516080810182528082018781526060820187905281528151808301909252848252602082810185905281019190915260028054600181018083556000839052835190939260049092027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01916109c591839161180a565b5060208201516109db906002808401919061180a565b505050600091825250600360205260409020805460ff1916600117905550505050565b610a06611838565b6002548210610a7657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f67726f757020696e646578206f7574206f662072616e67650000000000000000604482015290519081900360640190fd5b608060405190810160405280600284815481101515610a9157fe5b60009182526020822060049091020101548152602001600284815481101515610ab657fe5b60009182526020909120600490910201600101548152602001600284815481101515610ade57fe5b60009182526020822060026004909202010101548152602001600284815481101515610b0657fe5b6000918252602090912060036004909202010154905292915050565b6000610bee600085858080601f01602080910402602001604051908101604052809392919081815260200183838082843750506040805180820182528a3581526020808c01358183015260008f81526001918290528390208351608081018086529282018054958201958652939850965091945085935084916002919082016060860180831161029a575050509183525050604080518082019182905260028481018054835260209485019492939092600387019085018083116102d0575050505050815250506110a5565b1515610bf957610d70565b50600084815260016020526040902060050154600160a060020a0316801515610c4a576040517f158bff16635ac24f3d1acce162f0626cc6751bd434047538d76421366edf590690600090a1610d70565b60408051600160a060020a038316815290517f065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf09181900360200190a1604080517f6d112977000000000000000000000000000000000000000000000000000000008152600481018781526024820192835260448201869052600160a060020a03841692636d112977928992899289926064018484808284378201915050945050505050600060405180830381600087803b158015610d0757600080fd5b505af1158015610d1b573d6000803e3d6000fd5b505050600086815260016020819052604082208281559250820181610d408282611857565b610d4e600283016000611857565b505050600501805473ffffffffffffffffffffffffffffffffffffffff191690555b5050505050565b6000610d838180611865565b506000610d91600282611889565b50565b60408051600019430140602080830191909152825180830382018152918301928390528151600093918291908401908083835b60208310610de65780518252601f199092019160209182019101610dc7565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091206005819055436004556002549350915050811515610e2957fe5b069050600281815481101515610e3b57fe5b600091825260209091206004909102016006610e59818360026117cf565b50610e6c600282810190848101906117cf565b509050507fb57d39816b5fafa35e9395c99afd19edb72160f491310ba0942ea8c58c8bc9f9600554600454610ea0846109fe565b6040518084815260200183815260200182600460200280838360005b83811015610ed4578181015183820152602001610ebc565b50505050905001935050505060405180910390a150565b6060600082604051908082528060200260200182016040528015610f19578160200160208202803883390190505b50600054909250831115610f55576040517f08a70ba288e836bee6c9b4aea7482ee5ff8f63c5ad9d2533d9cf0ced64adc26290600090a1611028565b5060005b82811015610fb257600080546000198101908110610f7357fe5b90600052602060002001548282815181101515610f8c57fe5b602090810290910101526000805490610fa9906000198301611865565b50600101610f59565b7f5f30b698cceb472bcb5a80c4acc8c52ea45ea704f5aeeb2527d2d4c95f793dd7826040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015611014578181015183820152602001610ffc565b505050509050019250505060405180910390a15b505050565b60055481565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630155565b60045481565b6040805181815260608082018352918291906020820161080080388339505084516020808401919091529094015160408201529392505050565b600060608060608633600160a060020a03166001026040516020018083805190602001908083835b602083106110ec5780518252601f1990920191602091820191016110cd565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381526002928501838152608086018352909850945090920190505b6111366118b5565b81526020019060019003908161112e57505060408051600280825260608201909252919350602082015b6111686118cc565b8152602001906001900390816111605790505090508582600081518110151561118d57fe5b602090810290910101526111a083611426565b8260018151811015156111af57fe5b602090810290910101526111c16114ab565b8160008151811015156111d057fe5b6020908102909101015280518590829060019081106111eb57fe5b602090810290910101526111ff828261156c565b15156113485760408051808201825287518152602080890151818301528251608081018452885151815288518201518183015288820180515182860152518201516060820152835160ff8d1681527fc4a9adb51f42fa476b66c4c24c5553ba3a6dd39c7f9998d976f4e9164213e860948d948994909392919082019082810190859080838360005b8381101561129f578181015183820152602001611287565b5050505090500183600460200280838360005b838110156112ca5781810151838201526020016112b2565b50505050905001828103825285818151815260200191508051906020019080838360005b838110156113065781810151838201526020016112ee565b50505050905090810190601f1680156113335780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a1600093505b505050949350505050565b3b90565b8051825160009184918491849114611372576000935061141d565b5060005b825181101561141857818181518110151561138d57fe5b90602001015160f860020a900460f860020a027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191683828151811015156113d057fe5b60209101015160f860020a90819004027fff000000000000000000000000000000000000000000000000000000000000001614611410576000935061141d565b600101611376565b600193505b50505092915050565b61142e6118b5565b6000826040518082805190602001908083835b602083106114605780518252601f199092019160209182019101611441565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206001900490506114a461149e611768565b82611789565b9392505050565b6114b36118cc565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b82527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa6020838101919091528101919091525b90565b60008060006060600061157d6118f2565b865188516000911461158e57600080fd5b88519550856006029450846040519080825280602002602001820160405280156115c2578160200160208202803883390190505b509350600092505b858310156117375788838151811015156115e057fe5b602090810290910101515184518590600686029081106115fc57fe5b60209081029091010152885189908490811061161457fe5b9060200190602002015160200151848460060260010181518110151561163657fe5b60209081029091010152875188908490811061164e57fe5b60209081029190910101515151845185906002600687020190811061166f57fe5b60209081029091010152875188908490811061168757fe5b60209081029190910181015151015184518590600360068702019081106116aa57fe5b6020908102909101015287518890849081106116c257fe5b60209081029190910181015101515184518590600460068702019081106116e557fe5b6020908102909101015287518890849081106116fd57fe5b6020908102919091018101518101510151845185906005600687020190811061172257fe5b602090810290910101526001909201916115ca565b6020826020870260208701600060086107d05a03f1905080801561175b5750815115155b9998505050505050505050565b6117706118b5565b5060408051808201909152600181526002602082015290565b6117916118b5565b611799611911565b8351815260208085015190820152604080820184905282606083600060076107d05a03f115156117c857600080fd5b5092915050565b82600281019282156117fa579182015b828111156117fa5782548255916001019190600101906117df565b50611806929150611930565b5090565b82600281019282156117fa579160200282015b828111156117fa57825182559160200191906001019061181d565b6080604051908101604052806004906020820280388339509192915050565b506000815560010160009055565b81548183558181111561102857600083815260209020611028918101908301611930565b81548183558181111561102857600402816004028360005260206000209182019101611028919061194a565b604080518082019091526000808252602082015290565b6080604051908101604052806118e061197b565b81526020016118ed61197b565b905290565b6020604051908101604052806001906020820280388339509192915050565b6060604051908101604052806003906020820280388339509192915050565b61156991905b808211156118065760008155600101611936565b61156991905b808211156118065760006119648282611857565b611972600283016000611857565b50600401611950565b604080518082018252906002908290803883395091929150505600a165627a7a723058203ec3b7a451d6b07596dad9961008eea5ff818c3b8645e9df9882245207a05abd0029`

// DeployDOSProxy deploys a new Ethereum contract, binding an instance of DOSProxy to it.
func DeployDOSProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DOSProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DOSProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DOSProxy{DOSProxyCaller: DOSProxyCaller{contract: contract}, DOSProxyTransactor: DOSProxyTransactor{contract: contract}, DOSProxyFilterer: DOSProxyFilterer{contract: contract}}, nil
}

// DOSProxy is an auto generated Go binding around an Ethereum contract.
type DOSProxy struct {
	DOSProxyCaller     // Read-only binding to the contract
	DOSProxyTransactor // Write-only binding to the contract
	DOSProxyFilterer   // Log filterer for contract events
}

// DOSProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DOSProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DOSProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DOSProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DOSProxySession struct {
	Contract     *DOSProxy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DOSProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DOSProxyCallerSession struct {
	Contract *DOSProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DOSProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DOSProxyTransactorSession struct {
	Contract     *DOSProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DOSProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DOSProxyRaw struct {
	Contract *DOSProxy // Generic contract binding to access the raw methods on
}

// DOSProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DOSProxyCallerRaw struct {
	Contract *DOSProxyCaller // Generic read-only contract binding to access the raw methods on
}

// DOSProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DOSProxyTransactorRaw struct {
	Contract *DOSProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDOSProxy creates a new instance of DOSProxy, bound to a specific deployed contract.
func NewDOSProxy(address common.Address, backend bind.ContractBackend) (*DOSProxy, error) {
	contract, err := bindDOSProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DOSProxy{DOSProxyCaller: DOSProxyCaller{contract: contract}, DOSProxyTransactor: DOSProxyTransactor{contract: contract}, DOSProxyFilterer: DOSProxyFilterer{contract: contract}}, nil
}

// NewDOSProxyCaller creates a new read-only instance of DOSProxy, bound to a specific deployed contract.
func NewDOSProxyCaller(address common.Address, caller bind.ContractCaller) (*DOSProxyCaller, error) {
	contract, err := bindDOSProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DOSProxyCaller{contract: contract}, nil
}

// NewDOSProxyTransactor creates a new write-only instance of DOSProxy, bound to a specific deployed contract.
func NewDOSProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*DOSProxyTransactor, error) {
	contract, err := bindDOSProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DOSProxyTransactor{contract: contract}, nil
}

// NewDOSProxyFilterer creates a new log filterer instance of DOSProxy, bound to a specific deployed contract.
func NewDOSProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*DOSProxyFilterer, error) {
	contract, err := bindDOSProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DOSProxyFilterer{contract: contract}, nil
}

// bindDOSProxy binds a generic wrapper to an already deployed contract.
func bindDOSProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSProxy *DOSProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSProxy.Contract.DOSProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSProxy *DOSProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxy.Contract.DOSProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSProxy *DOSProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSProxy.Contract.DOSProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSProxy *DOSProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSProxy *DOSProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSProxy *DOSProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSProxy.Contract.contract.Transact(opts, method, params...)
}

// GetGroupPubKey is a free data retrieval call binding the contract method 0x92021653.
//
// Solidity: function getGroupPubKey(idx uint256) constant returns(uint256[4])
func (_DOSProxy *DOSProxyCaller) GetGroupPubKey(opts *bind.CallOpts, idx *big.Int) ([4]*big.Int, error) {
	var (
		ret0 = new([4]*big.Int)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "getGroupPubKey", idx)
	return *ret0, err
}

// GetGroupPubKey is a free data retrieval call binding the contract method 0x92021653.
//
// Solidity: function getGroupPubKey(idx uint256) constant returns(uint256[4])
func (_DOSProxy *DOSProxySession) GetGroupPubKey(idx *big.Int) ([4]*big.Int, error) {
	return _DOSProxy.Contract.GetGroupPubKey(&_DOSProxy.CallOpts, idx)
}

// GetGroupPubKey is a free data retrieval call binding the contract method 0x92021653.
//
// Solidity: function getGroupPubKey(idx uint256) constant returns(uint256[4])
func (_DOSProxy *DOSProxyCallerSession) GetGroupPubKey(idx *big.Int) ([4]*big.Int, error) {
	return _DOSProxy.Contract.GetGroupPubKey(&_DOSProxy.CallOpts, idx)
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() constant returns(uint256)
func (_DOSProxy *DOSProxyCaller) LastRandomness(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "lastRandomness")
	return *ret0, err
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() constant returns(uint256)
func (_DOSProxy *DOSProxySession) LastRandomness() (*big.Int, error) {
	return _DOSProxy.Contract.LastRandomness(&_DOSProxy.CallOpts)
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() constant returns(uint256)
func (_DOSProxy *DOSProxyCallerSession) LastRandomness() (*big.Int, error) {
	return _DOSProxy.Contract.LastRandomness(&_DOSProxy.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() constant returns(uint256)
func (_DOSProxy *DOSProxyCaller) LastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "lastUpdatedBlock")
	return *ret0, err
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() constant returns(uint256)
func (_DOSProxy *DOSProxySession) LastUpdatedBlock() (*big.Int, error) {
	return _DOSProxy.Contract.LastUpdatedBlock(&_DOSProxy.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() constant returns(uint256)
func (_DOSProxy *DOSProxyCallerSession) LastUpdatedBlock() (*big.Int, error) {
	return _DOSProxy.Contract.LastUpdatedBlock(&_DOSProxy.CallOpts)
}

// FireRandom is a paid mutator transaction binding the contract method 0xea5cba13.
//
// Solidity: function fireRandom() returns()
func (_DOSProxy *DOSProxyTransactor) FireRandom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "fireRandom")
}

// FireRandom is a paid mutator transaction binding the contract method 0xea5cba13.
//
// Solidity: function fireRandom() returns()
func (_DOSProxy *DOSProxySession) FireRandom() (*types.Transaction, error) {
	return _DOSProxy.Contract.FireRandom(&_DOSProxy.TransactOpts)
}

// FireRandom is a paid mutator transaction binding the contract method 0xea5cba13.
//
// Solidity: function fireRandom() returns()
func (_DOSProxy *DOSProxyTransactorSession) FireRandom() (*types.Transaction, error) {
	return _DOSProxy.Contract.FireRandom(&_DOSProxy.TransactOpts)
}

// Grouping is a paid mutator transaction binding the contract method 0xeab14fe6.
//
// Solidity: function grouping(size uint256) returns()
func (_DOSProxy *DOSProxyTransactor) Grouping(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "grouping", size)
}

// Grouping is a paid mutator transaction binding the contract method 0xeab14fe6.
//
// Solidity: function grouping(size uint256) returns()
func (_DOSProxy *DOSProxySession) Grouping(size *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.Grouping(&_DOSProxy.TransactOpts, size)
}

// Grouping is a paid mutator transaction binding the contract method 0xeab14fe6.
//
// Solidity: function grouping(size uint256) returns()
func (_DOSProxy *DOSProxyTransactorSession) Grouping(size *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.Grouping(&_DOSProxy.TransactOpts, size)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query(from address, blkNum uint256, timeout uint256, queryType string, queryPath string) returns(uint256)
func (_DOSProxy *DOSProxyTransactor) Query(opts *bind.TransactOpts, from common.Address, blkNum *big.Int, timeout *big.Int, queryType string, queryPath string) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "query", from, blkNum, timeout, queryType, queryPath)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query(from address, blkNum uint256, timeout uint256, queryType string, queryPath string) returns(uint256)
func (_DOSProxy *DOSProxySession) Query(from common.Address, blkNum *big.Int, timeout *big.Int, queryType string, queryPath string) (*types.Transaction, error) {
	return _DOSProxy.Contract.Query(&_DOSProxy.TransactOpts, from, blkNum, timeout, queryType, queryPath)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query(from address, blkNum uint256, timeout uint256, queryType string, queryPath string) returns(uint256)
func (_DOSProxy *DOSProxyTransactorSession) Query(from common.Address, blkNum *big.Int, timeout *big.Int, queryType string, queryPath string) (*types.Transaction, error) {
	return _DOSProxy.Contract.Query(&_DOSProxy.TransactOpts, from, blkNum, timeout, queryType, queryPath)
}

// ResetContract is a paid mutator transaction binding the contract method 0xb181a8fc.
//
// Solidity: function resetContract() returns()
func (_DOSProxy *DOSProxyTransactor) ResetContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "resetContract")
}

// ResetContract is a paid mutator transaction binding the contract method 0xb181a8fc.
//
// Solidity: function resetContract() returns()
func (_DOSProxy *DOSProxySession) ResetContract() (*types.Transaction, error) {
	return _DOSProxy.Contract.ResetContract(&_DOSProxy.TransactOpts)
}

// ResetContract is a paid mutator transaction binding the contract method 0xb181a8fc.
//
// Solidity: function resetContract() returns()
func (_DOSProxy *DOSProxyTransactorSession) ResetContract() (*types.Transaction, error) {
	return _DOSProxy.Contract.ResetContract(&_DOSProxy.TransactOpts)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x79a924c5.
//
// Solidity: function setPublicKey(x1 uint256, x2 uint256, y1 uint256, y2 uint256) returns()
func (_DOSProxy *DOSProxyTransactor) SetPublicKey(opts *bind.TransactOpts, x1 *big.Int, x2 *big.Int, y1 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "setPublicKey", x1, x2, y1, y2)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x79a924c5.
//
// Solidity: function setPublicKey(x1 uint256, x2 uint256, y1 uint256, y2 uint256) returns()
func (_DOSProxy *DOSProxySession) SetPublicKey(x1 *big.Int, x2 *big.Int, y1 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetPublicKey(&_DOSProxy.TransactOpts, x1, x2, y1, y2)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x79a924c5.
//
// Solidity: function setPublicKey(x1 uint256, x2 uint256, y1 uint256, y2 uint256) returns()
func (_DOSProxy *DOSProxyTransactorSession) SetPublicKey(x1 *big.Int, x2 *big.Int, y1 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetPublicKey(&_DOSProxy.TransactOpts, x1, x2, y1, y2)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x9261cc7e.
//
// Solidity: function triggerCallback(queryId uint256, result bytes, sig uint256[2]) returns()
func (_DOSProxy *DOSProxyTransactor) TriggerCallback(opts *bind.TransactOpts, queryId *big.Int, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "triggerCallback", queryId, result, sig)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x9261cc7e.
//
// Solidity: function triggerCallback(queryId uint256, result bytes, sig uint256[2]) returns()
func (_DOSProxy *DOSProxySession) TriggerCallback(queryId *big.Int, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.TriggerCallback(&_DOSProxy.TransactOpts, queryId, result, sig)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x9261cc7e.
//
// Solidity: function triggerCallback(queryId uint256, result bytes, sig uint256[2]) returns()
func (_DOSProxy *DOSProxyTransactorSession) TriggerCallback(queryId *big.Int, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.TriggerCallback(&_DOSProxy.TransactOpts, queryId, result, sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(sig uint256[2]) returns()
func (_DOSProxy *DOSProxyTransactor) UpdateRandomness(opts *bind.TransactOpts, sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "updateRandomness", sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(sig uint256[2]) returns()
func (_DOSProxy *DOSProxySession) UpdateRandomness(sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.UpdateRandomness(&_DOSProxy.TransactOpts, sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(sig uint256[2]) returns()
func (_DOSProxy *DOSProxyTransactorSession) UpdateRandomness(sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.UpdateRandomness(&_DOSProxy.TransactOpts, sig)
}

// UploadNodeId is a paid mutator transaction binding the contract method 0xf89a15f7.
//
// Solidity: function uploadNodeId(id uint256) returns()
func (_DOSProxy *DOSProxyTransactor) UploadNodeId(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "uploadNodeId", id)
}

// UploadNodeId is a paid mutator transaction binding the contract method 0xf89a15f7.
//
// Solidity: function uploadNodeId(id uint256) returns()
func (_DOSProxy *DOSProxySession) UploadNodeId(id *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.UploadNodeId(&_DOSProxy.TransactOpts, id)
}

// UploadNodeId is a paid mutator transaction binding the contract method 0xf89a15f7.
//
// Solidity: function uploadNodeId(id uint256) returns()
func (_DOSProxy *DOSProxyTransactorSession) UploadNodeId(id *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.UploadNodeId(&_DOSProxy.TransactOpts, id)
}

// DOSProxyLogCallbackTriggeredForIterator is returned from FilterLogCallbackTriggeredFor and is used to iterate over the raw logs and unpacked data for LogCallbackTriggeredFor events raised by the DOSProxy contract.
type DOSProxyLogCallbackTriggeredForIterator struct {
	Event *DOSProxyLogCallbackTriggeredFor // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogCallbackTriggeredForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogCallbackTriggeredFor)
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
		it.Event = new(DOSProxyLogCallbackTriggeredFor)
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
func (it *DOSProxyLogCallbackTriggeredForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogCallbackTriggeredForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogCallbackTriggeredFor represents a LogCallbackTriggeredFor event raised by the DOSProxy contract.
type DOSProxyLogCallbackTriggeredFor struct {
	CallbackAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogCallbackTriggeredFor is a free log retrieval operation binding the contract event 0x065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf0.
//
// Solidity: e LogCallbackTriggeredFor(callbackAddr address)
func (_DOSProxy *DOSProxyFilterer) FilterLogCallbackTriggeredFor(opts *bind.FilterOpts) (*DOSProxyLogCallbackTriggeredForIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogCallbackTriggeredFor")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogCallbackTriggeredForIterator{contract: _DOSProxy.contract, event: "LogCallbackTriggeredFor", logs: logs, sub: sub}, nil
}

// WatchLogCallbackTriggeredFor is a free log subscription operation binding the contract event 0x065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf0.
//
// Solidity: e LogCallbackTriggeredFor(callbackAddr address)
func (_DOSProxy *DOSProxyFilterer) WatchLogCallbackTriggeredFor(opts *bind.WatchOpts, sink chan<- *DOSProxyLogCallbackTriggeredFor) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogCallbackTriggeredFor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogCallbackTriggeredFor)
				if err := _DOSProxy.contract.UnpackLog(event, "LogCallbackTriggeredFor", log); err != nil {
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

// DOSProxyLogGroupingIterator is returned from FilterLogGrouping and is used to iterate over the raw logs and unpacked data for LogGrouping events raised by the DOSProxy contract.
type DOSProxyLogGroupingIterator struct {
	Event *DOSProxyLogGrouping // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogGroupingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogGrouping)
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
		it.Event = new(DOSProxyLogGrouping)
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
func (it *DOSProxyLogGroupingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogGroupingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogGrouping represents a LogGrouping event raised by the DOSProxy contract.
type DOSProxyLogGrouping struct {
	NodeId []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogGrouping is a free log retrieval operation binding the contract event 0x5f30b698cceb472bcb5a80c4acc8c52ea45ea704f5aeeb2527d2d4c95f793dd7.
//
// Solidity: e LogGrouping(NodeId uint256[])
func (_DOSProxy *DOSProxyFilterer) FilterLogGrouping(opts *bind.FilterOpts) (*DOSProxyLogGroupingIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogGrouping")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogGroupingIterator{contract: _DOSProxy.contract, event: "LogGrouping", logs: logs, sub: sub}, nil
}

// WatchLogGrouping is a free log subscription operation binding the contract event 0x5f30b698cceb472bcb5a80c4acc8c52ea45ea704f5aeeb2527d2d4c95f793dd7.
//
// Solidity: e LogGrouping(NodeId uint256[])
func (_DOSProxy *DOSProxyFilterer) WatchLogGrouping(opts *bind.WatchOpts, sink chan<- *DOSProxyLogGrouping) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogGrouping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogGrouping)
				if err := _DOSProxy.contract.UnpackLog(event, "LogGrouping", log); err != nil {
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

// DOSProxyLogInsufficientGroupNumberIterator is returned from FilterLogInsufficientGroupNumber and is used to iterate over the raw logs and unpacked data for LogInsufficientGroupNumber events raised by the DOSProxy contract.
type DOSProxyLogInsufficientGroupNumberIterator struct {
	Event *DOSProxyLogInsufficientGroupNumber // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogInsufficientGroupNumberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogInsufficientGroupNumber)
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
		it.Event = new(DOSProxyLogInsufficientGroupNumber)
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
func (it *DOSProxyLogInsufficientGroupNumberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogInsufficientGroupNumberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogInsufficientGroupNumber represents a LogInsufficientGroupNumber event raised by the DOSProxy contract.
type DOSProxyLogInsufficientGroupNumber struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogInsufficientGroupNumber is a free log retrieval operation binding the contract event 0x08a70ba288e836bee6c9b4aea7482ee5ff8f63c5ad9d2533d9cf0ced64adc262.
//
// Solidity: e LogInsufficientGroupNumber()
func (_DOSProxy *DOSProxyFilterer) FilterLogInsufficientGroupNumber(opts *bind.FilterOpts) (*DOSProxyLogInsufficientGroupNumberIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogInsufficientGroupNumber")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogInsufficientGroupNumberIterator{contract: _DOSProxy.contract, event: "LogInsufficientGroupNumber", logs: logs, sub: sub}, nil
}

// WatchLogInsufficientGroupNumber is a free log subscription operation binding the contract event 0x08a70ba288e836bee6c9b4aea7482ee5ff8f63c5ad9d2533d9cf0ced64adc262.
//
// Solidity: e LogInsufficientGroupNumber()
func (_DOSProxy *DOSProxyFilterer) WatchLogInsufficientGroupNumber(opts *bind.WatchOpts, sink chan<- *DOSProxyLogInsufficientGroupNumber) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogInsufficientGroupNumber")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogInsufficientGroupNumber)
				if err := _DOSProxy.contract.UnpackLog(event, "LogInsufficientGroupNumber", log); err != nil {
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

// DOSProxyLogInvalidSignatureForIterator is returned from FilterLogInvalidSignatureFor and is used to iterate over the raw logs and unpacked data for LogInvalidSignatureFor events raised by the DOSProxy contract.
type DOSProxyLogInvalidSignatureForIterator struct {
	Event *DOSProxyLogInvalidSignatureFor // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogInvalidSignatureForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogInvalidSignatureFor)
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
		it.Event = new(DOSProxyLogInvalidSignatureFor)
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
func (it *DOSProxyLogInvalidSignatureForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogInvalidSignatureForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogInvalidSignatureFor represents a LogInvalidSignatureFor event raised by the DOSProxy contract.
type DOSProxyLogInvalidSignatureFor struct {
	TrafficType uint8
	Message     []byte
	Signature   [2]*big.Int
	PubKey      [4]*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogInvalidSignatureFor is a free log retrieval operation binding the contract event 0xc4a9adb51f42fa476b66c4c24c5553ba3a6dd39c7f9998d976f4e9164213e860.
//
// Solidity: e LogInvalidSignatureFor(trafficType uint8, message bytes, signature uint256[2], pubKey uint256[4])
func (_DOSProxy *DOSProxyFilterer) FilterLogInvalidSignatureFor(opts *bind.FilterOpts) (*DOSProxyLogInvalidSignatureForIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogInvalidSignatureFor")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogInvalidSignatureForIterator{contract: _DOSProxy.contract, event: "LogInvalidSignatureFor", logs: logs, sub: sub}, nil
}

// WatchLogInvalidSignatureFor is a free log subscription operation binding the contract event 0xc4a9adb51f42fa476b66c4c24c5553ba3a6dd39c7f9998d976f4e9164213e860.
//
// Solidity: e LogInvalidSignatureFor(trafficType uint8, message bytes, signature uint256[2], pubKey uint256[4])
func (_DOSProxy *DOSProxyFilterer) WatchLogInvalidSignatureFor(opts *bind.WatchOpts, sink chan<- *DOSProxyLogInvalidSignatureFor) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogInvalidSignatureFor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogInvalidSignatureFor)
				if err := _DOSProxy.contract.UnpackLog(event, "LogInvalidSignatureFor", log); err != nil {
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

// DOSProxyLogNonContractCallIterator is returned from FilterLogNonContractCall and is used to iterate over the raw logs and unpacked data for LogNonContractCall events raised by the DOSProxy contract.
type DOSProxyLogNonContractCallIterator struct {
	Event *DOSProxyLogNonContractCall // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogNonContractCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogNonContractCall)
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
		it.Event = new(DOSProxyLogNonContractCall)
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
func (it *DOSProxyLogNonContractCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogNonContractCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogNonContractCall represents a LogNonContractCall event raised by the DOSProxy contract.
type DOSProxyLogNonContractCall struct {
	From common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogNonContractCall is a free log retrieval operation binding the contract event 0x6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb5.
//
// Solidity: e LogNonContractCall(from address)
func (_DOSProxy *DOSProxyFilterer) FilterLogNonContractCall(opts *bind.FilterOpts) (*DOSProxyLogNonContractCallIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogNonContractCall")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogNonContractCallIterator{contract: _DOSProxy.contract, event: "LogNonContractCall", logs: logs, sub: sub}, nil
}

// WatchLogNonContractCall is a free log subscription operation binding the contract event 0x6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb5.
//
// Solidity: e LogNonContractCall(from address)
func (_DOSProxy *DOSProxyFilterer) WatchLogNonContractCall(opts *bind.WatchOpts, sink chan<- *DOSProxyLogNonContractCall) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogNonContractCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogNonContractCall)
				if err := _DOSProxy.contract.UnpackLog(event, "LogNonContractCall", log); err != nil {
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

// DOSProxyLogNonSupportedTypeIterator is returned from FilterLogNonSupportedType and is used to iterate over the raw logs and unpacked data for LogNonSupportedType events raised by the DOSProxy contract.
type DOSProxyLogNonSupportedTypeIterator struct {
	Event *DOSProxyLogNonSupportedType // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogNonSupportedTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogNonSupportedType)
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
		it.Event = new(DOSProxyLogNonSupportedType)
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
func (it *DOSProxyLogNonSupportedTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogNonSupportedTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogNonSupportedType represents a LogNonSupportedType event raised by the DOSProxy contract.
type DOSProxyLogNonSupportedType struct {
	QueryType string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogNonSupportedType is a free log retrieval operation binding the contract event 0x70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41.
//
// Solidity: e LogNonSupportedType(queryType string)
func (_DOSProxy *DOSProxyFilterer) FilterLogNonSupportedType(opts *bind.FilterOpts) (*DOSProxyLogNonSupportedTypeIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogNonSupportedType")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogNonSupportedTypeIterator{contract: _DOSProxy.contract, event: "LogNonSupportedType", logs: logs, sub: sub}, nil
}

// WatchLogNonSupportedType is a free log subscription operation binding the contract event 0x70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41.
//
// Solidity: e LogNonSupportedType(queryType string)
func (_DOSProxy *DOSProxyFilterer) WatchLogNonSupportedType(opts *bind.WatchOpts, sink chan<- *DOSProxyLogNonSupportedType) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogNonSupportedType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogNonSupportedType)
				if err := _DOSProxy.contract.UnpackLog(event, "LogNonSupportedType", log); err != nil {
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

// DOSProxyLogQueryFromNonExistentUCIterator is returned from FilterLogQueryFromNonExistentUC and is used to iterate over the raw logs and unpacked data for LogQueryFromNonExistentUC events raised by the DOSProxy contract.
type DOSProxyLogQueryFromNonExistentUCIterator struct {
	Event *DOSProxyLogQueryFromNonExistentUC // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogQueryFromNonExistentUCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogQueryFromNonExistentUC)
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
		it.Event = new(DOSProxyLogQueryFromNonExistentUC)
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
func (it *DOSProxyLogQueryFromNonExistentUCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogQueryFromNonExistentUCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogQueryFromNonExistentUC represents a LogQueryFromNonExistentUC event raised by the DOSProxy contract.
type DOSProxyLogQueryFromNonExistentUC struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogQueryFromNonExistentUC is a free log retrieval operation binding the contract event 0x158bff16635ac24f3d1acce162f0626cc6751bd434047538d76421366edf5906.
//
// Solidity: e LogQueryFromNonExistentUC()
func (_DOSProxy *DOSProxyFilterer) FilterLogQueryFromNonExistentUC(opts *bind.FilterOpts) (*DOSProxyLogQueryFromNonExistentUCIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogQueryFromNonExistentUC")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogQueryFromNonExistentUCIterator{contract: _DOSProxy.contract, event: "LogQueryFromNonExistentUC", logs: logs, sub: sub}, nil
}

// WatchLogQueryFromNonExistentUC is a free log subscription operation binding the contract event 0x158bff16635ac24f3d1acce162f0626cc6751bd434047538d76421366edf5906.
//
// Solidity: e LogQueryFromNonExistentUC()
func (_DOSProxy *DOSProxyFilterer) WatchLogQueryFromNonExistentUC(opts *bind.WatchOpts, sink chan<- *DOSProxyLogQueryFromNonExistentUC) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogQueryFromNonExistentUC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogQueryFromNonExistentUC)
				if err := _DOSProxy.contract.UnpackLog(event, "LogQueryFromNonExistentUC", log); err != nil {
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

// DOSProxyLogUpdateRandomIterator is returned from FilterLogUpdateRandom and is used to iterate over the raw logs and unpacked data for LogUpdateRandom events raised by the DOSProxy contract.
type DOSProxyLogUpdateRandomIterator struct {
	Event *DOSProxyLogUpdateRandom // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogUpdateRandomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogUpdateRandom)
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
		it.Event = new(DOSProxyLogUpdateRandom)
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
func (it *DOSProxyLogUpdateRandomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogUpdateRandomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogUpdateRandom represents a LogUpdateRandom event raised by the DOSProxy contract.
type DOSProxyLogUpdateRandom struct {
	LastRandomness   *big.Int
	LastUpdatedBlock *big.Int
	DispatchedGroup  [4]*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogUpdateRandom is a free log retrieval operation binding the contract event 0xb57d39816b5fafa35e9395c99afd19edb72160f491310ba0942ea8c58c8bc9f9.
//
// Solidity: e LogUpdateRandom(lastRandomness uint256, lastUpdatedBlock uint256, dispatchedGroup uint256[4])
func (_DOSProxy *DOSProxyFilterer) FilterLogUpdateRandom(opts *bind.FilterOpts) (*DOSProxyLogUpdateRandomIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogUpdateRandom")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogUpdateRandomIterator{contract: _DOSProxy.contract, event: "LogUpdateRandom", logs: logs, sub: sub}, nil
}

// WatchLogUpdateRandom is a free log subscription operation binding the contract event 0xb57d39816b5fafa35e9395c99afd19edb72160f491310ba0942ea8c58c8bc9f9.
//
// Solidity: e LogUpdateRandom(lastRandomness uint256, lastUpdatedBlock uint256, dispatchedGroup uint256[4])
func (_DOSProxy *DOSProxyFilterer) WatchLogUpdateRandom(opts *bind.WatchOpts, sink chan<- *DOSProxyLogUpdateRandom) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogUpdateRandom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogUpdateRandom)
				if err := _DOSProxy.contract.UnpackLog(event, "LogUpdateRandom", log); err != nil {
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

// DOSProxyLogUrlIterator is returned from FilterLogUrl and is used to iterate over the raw logs and unpacked data for LogUrl events raised by the DOSProxy contract.
type DOSProxyLogUrlIterator struct {
	Event *DOSProxyLogUrl // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogUrlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogUrl)
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
		it.Event = new(DOSProxyLogUrl)
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
func (it *DOSProxyLogUrlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogUrlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogUrl represents a LogUrl event raised by the DOSProxy contract.
type DOSProxyLogUrl struct {
	QueryId         *big.Int
	Url             string
	Timeout         *big.Int
	Randomness      *big.Int
	DispatchedGroup [4]*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogUrl is a free log retrieval operation binding the contract event 0x75b4fd18552b0d79989d6cc5de124544de30dc29b548e1f743a719ce5ad44329.
//
// Solidity: e LogUrl(queryId uint256, url string, timeout uint256, randomness uint256, dispatchedGroup uint256[4])
func (_DOSProxy *DOSProxyFilterer) FilterLogUrl(opts *bind.FilterOpts) (*DOSProxyLogUrlIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogUrl")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogUrlIterator{contract: _DOSProxy.contract, event: "LogUrl", logs: logs, sub: sub}, nil
}

// WatchLogUrl is a free log subscription operation binding the contract event 0x75b4fd18552b0d79989d6cc5de124544de30dc29b548e1f743a719ce5ad44329.
//
// Solidity: e LogUrl(queryId uint256, url string, timeout uint256, randomness uint256, dispatchedGroup uint256[4])
func (_DOSProxy *DOSProxyFilterer) WatchLogUrl(opts *bind.WatchOpts, sink chan<- *DOSProxyLogUrl) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogUrl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogUrl)
				if err := _DOSProxy.contract.UnpackLog(event, "LogUrl", log); err != nil {
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

// UserContractInterfaceABI is the input ABI used to generate the binding from.
const UserContractInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UserContractInterfaceBin is the compiled bytecode used for deploying new contracts.
const UserContractInterfaceBin = `0x`

// DeployUserContractInterface deploys a new Ethereum contract, binding an instance of UserContractInterface to it.
func DeployUserContractInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserContractInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(UserContractInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserContractInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserContractInterface{UserContractInterfaceCaller: UserContractInterfaceCaller{contract: contract}, UserContractInterfaceTransactor: UserContractInterfaceTransactor{contract: contract}, UserContractInterfaceFilterer: UserContractInterfaceFilterer{contract: contract}}, nil
}

// UserContractInterface is an auto generated Go binding around an Ethereum contract.
type UserContractInterface struct {
	UserContractInterfaceCaller     // Read-only binding to the contract
	UserContractInterfaceTransactor // Write-only binding to the contract
	UserContractInterfaceFilterer   // Log filterer for contract events
}

// UserContractInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserContractInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserContractInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserContractInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserContractInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserContractInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserContractInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserContractInterfaceSession struct {
	Contract     *UserContractInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UserContractInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserContractInterfaceCallerSession struct {
	Contract *UserContractInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// UserContractInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserContractInterfaceTransactorSession struct {
	Contract     *UserContractInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// UserContractInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserContractInterfaceRaw struct {
	Contract *UserContractInterface // Generic contract binding to access the raw methods on
}

// UserContractInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserContractInterfaceCallerRaw struct {
	Contract *UserContractInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// UserContractInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserContractInterfaceTransactorRaw struct {
	Contract *UserContractInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserContractInterface creates a new instance of UserContractInterface, bound to a specific deployed contract.
func NewUserContractInterface(address common.Address, backend bind.ContractBackend) (*UserContractInterface, error) {
	contract, err := bindUserContractInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserContractInterface{UserContractInterfaceCaller: UserContractInterfaceCaller{contract: contract}, UserContractInterfaceTransactor: UserContractInterfaceTransactor{contract: contract}, UserContractInterfaceFilterer: UserContractInterfaceFilterer{contract: contract}}, nil
}

// NewUserContractInterfaceCaller creates a new read-only instance of UserContractInterface, bound to a specific deployed contract.
func NewUserContractInterfaceCaller(address common.Address, caller bind.ContractCaller) (*UserContractInterfaceCaller, error) {
	contract, err := bindUserContractInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserContractInterfaceCaller{contract: contract}, nil
}

// NewUserContractInterfaceTransactor creates a new write-only instance of UserContractInterface, bound to a specific deployed contract.
func NewUserContractInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*UserContractInterfaceTransactor, error) {
	contract, err := bindUserContractInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserContractInterfaceTransactor{contract: contract}, nil
}

// NewUserContractInterfaceFilterer creates a new log filterer instance of UserContractInterface, bound to a specific deployed contract.
func NewUserContractInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*UserContractInterfaceFilterer, error) {
	contract, err := bindUserContractInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserContractInterfaceFilterer{contract: contract}, nil
}

// bindUserContractInterface binds a generic wrapper to an already deployed contract.
func bindUserContractInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserContractInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserContractInterface *UserContractInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserContractInterface.Contract.UserContractInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserContractInterface *UserContractInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserContractInterface.Contract.UserContractInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserContractInterface *UserContractInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserContractInterface.Contract.UserContractInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserContractInterface *UserContractInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserContractInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserContractInterface *UserContractInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserContractInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserContractInterface *UserContractInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserContractInterface.Contract.contract.Transact(opts, method, params...)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__( uint256,  bytes) returns()
func (_UserContractInterface *UserContractInterfaceTransactor) Callback_(opts *bind.TransactOpts, arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _UserContractInterface.contract.Transact(opts, "__callback__", arg0, arg1)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__( uint256,  bytes) returns()
func (_UserContractInterface *UserContractInterfaceSession) Callback_(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _UserContractInterface.Contract.Callback_(&_UserContractInterface.TransactOpts, arg0, arg1)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__( uint256,  bytes) returns()
func (_UserContractInterface *UserContractInterfaceTransactorSession) Callback_(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _UserContractInterface.Contract.Callback_(&_UserContractInterface.TransactOpts, arg0, arg1)
}
