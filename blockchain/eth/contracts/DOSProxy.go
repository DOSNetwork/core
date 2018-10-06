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

// DOSProxyABI is the input ABI used to generate the binding from.
const DOSProxyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBootstrapIp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRandomNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"query_id\",\"type\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"triggerCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"block_number\",\"type\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\"},{\"name\":\"query_type\",\"type\":\"string\"},{\"name\":\"query_path\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"group_id\",\"type\":\"uint256\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"toBytes\",\"outputs\":[{\"name\":\"numBytes\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ip\",\"type\":\"string\"}],\"name\":\"setBootstrapIp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group_id\",\"type\":\"uint256\"},{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"x2\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"},{\"name\":\"y2\",\"type\":\"uint256\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group_id\",\"type\":\"uint256\"},{\"name\":\"random\",\"type\":\"uint256\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"setRandomNum\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"grouping\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uploadNodeId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"groupId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"queryId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"LogUrl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"query_type\",\"type\":\"string\"}],\"name\":\"LogNonSupportedType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"}],\"name\":\"LogNonContractCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user_contract_addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"LogCallbackTriggeredFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogQueryFromNonExistentUC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogInvalidSignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogInsufficientGroupNumber\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"GroupId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"NodeId\",\"type\":\"uint256[]\"}],\"name\":\"LogGrouping\",\"type\":\"event\"}]"

// DOSProxyBin is the compiled bytecode used for deploying new contracts.
const DOSProxyBin = `0x608060405243600055600060015534801561001957600080fd5b50611745806100296000396000f3006080604052600436106100b95763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630322c2ad81146100bb57806325b814f41461014557806334ebda4d1461016c578063482edfaa146101d357806375d2c1ce1461027d578063775a8f5e146102bb578063872e8640146102d3578063b181a8fc1461032c578063c8a03ef214610341578063dfb6f0e514610365578063eab14fe614610386578063f89a15f71461039e575b005b3480156100c757600080fd5b506100d06103b6565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010a5781810151838201526020016100f2565b50505050905090810190601f1680156101375780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015157600080fd5b5061015a61044d565b60408051918252519081900360200190f35b34801561017857600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100b99583359536956044949193909101919081908401838280828437509497505084359550505060209092013591506104539050565b3480156101df57600080fd5b50604080516020601f6064356004818101359283018490048402850184019095528184526100b994600160a060020a03813516946024803595604435953695608494930191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061068f9650505050505050565b34801561028957600080fd5b506102956004356109ea565b604080519485526020850193909352838301919091526060830152519081900360800190f35b3480156102c757600080fd5b506100d0600435610a12565b3480156102df57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100b9943694929360249392840191908190840183828082843750949750610a3d9650505050505050565b34801561033857600080fd5b506100b9610a54565b34801561034d57600080fd5b506100b9600435602435604435606435608435610a88565b34801561037157600080fd5b506100b9600435602435604435606435610baa565b34801561039257600080fd5b506100b9600435610c20565b3480156103aa57600080fd5b506100b9600435610d9b565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104425780601f1061041757610100808354040283529160200191610442565b820191906000526020600020905b81548152906001019060200180831161042557829003601f168201915b505050505090505b90565b60005490565b61045b611530565b506040805180820182528381526020808201849052600087815260069091529190912054600160a060020a03168015156104bd576040517f158bff16635ac24f3d1acce162f0626cc6751bd434047538d76421366edf590690600090a1610687565b6104c8868684610dd0565b15156104fc576040517ff0cda705e46caa68e5854fa85a2635f77f3f6b5c927bd409ee7d935e4bb0322c90600090a1610687565b7fcd714230b213422971bdd48f3fa7f63c52e50f9fa7356f6aa42a191c12f046d081866040518083600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561057457818101518382015260200161055c565b50505050905090810190601f1680156105a15780820380516001836020036101000a031916815260200191505b50935050505060405180910390a16040517f4a40f5f9000000000000000000000000000000000000000000000000000000008152602060048201818152875160248401528751600160a060020a03851693634a40f5f9938a939283926044019185019080838360005b8381101561062257818101518382015260200161060a565b50505050905090810190601f16801561064f5780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15801561066e57600080fd5b505af1158015610682573d6000803e3d6000fd5b505050505b505050505050565b600080600061069d88610f45565b11156109a45786868686866040518086600160a060020a0316600160a060020a03166c0100000000000000000000000002815260140185815260200184815260200183805190602001908083835b6020831061070a5780518252601f1990920191602091820191016106eb565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106107525780518252601f199092019160209182019101610733565b6001836020036101000a038019825116818451168082178552505050505050905001955050505050506040518091039020600190049150866006600084815260200190815260200160002060006101000a815481600160a060020a030219169083600160a060020a03160217905550610800846040805190810160405280600381526020017f4150490000000000000000000000000000000000000000000000000000000000815250610f49565b156109055750600280546000818152600760209081526040808320868452600890925290912091929091906108389082908490611547565b5061084b60028281019084810190611547565b509050507fc24f120340e1beeac424b47e9e08f835661963ec3a5f7aa7edf537b71533b2ae818385886040518085815260200184815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156108c35781810151838201526020016108ab565b50505050905090810190601f1680156108f05780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a161099f565b7f70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41846040518080602001828103825283818151815260200191508051906020019080838360005b8381101561096457818101518382015260200161094c565b50505050905090810190601f1680156109915780820380516001836020036101000a031916815260200191505b509250505060405180910390a15b6109e1565b60408051600160a060020a038916815290517f6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb59181900360200190a15b50505050505050565b6000908152600760205260409020805460018201546002830154600390930154919390929190565b6040805160208082528183019092526060918082016104008038833950505060208101929092525090565b8051610a50906003906020840190611582565b5050565b60006001819055610a666004826115f0565b50604080516020810191829052600090819052610a8591600391611582565b50565b604080516002808252606080830184529283929190602083019080388339505060408051600280825260608201835293955092915060208301908038833901905050905085826000815181101515610adc57fe5b602090810290910101528151859083906001908110610af757fe5b602090810290910101528051849082906000908110610b1257fe5b602090810290910101528051839082906001908110610b2d57fe5b602090810290910181019190915260408051608081018252808201898152606082018990528152815180830183528781528084018790528184015260008a8152600790935291208151610b839082906002611619565b506020820151610b999060028084019190611619565b505050600296909655505050505050565b610bb2611530565b5060408051808201909152828152602081018290526060610bd285610a12565b9050610bdf868284611018565b1515610c13576040517ff0cda705e46caa68e5854fa85a2635f77f3f6b5c927bd409ee7d935e4bb0322c90600090a1610687565b6000859055505050505050565b606060008083604051908082528060200260200182016040528015610c4f578160200160208202803883390190505b50600454909350841115610c8b576040517f08a70ba288e836bee6c9b4aea7482ee5ff8f63c5ad9d2533d9cf0ced64adc26290600090a1610d95565b600091505b83821015610cec57600480546000198101908110610caa57fe5b90600052602060002001548383815181101515610cc357fe5b602090810290910101526004805490610ce09060001983016115f0565b50600190910190610c90565b506001805480820190915560008181526005602090815260409091208451610d1692860190611646565b507ffdb627147562eb968c80eba5de51c105b83bf375bf584d31caf63648087a0ed981846040518083815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015610d80578181015183820152602001610d68565b50505050905001935050505060405180910390a15b50505050565b600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0155565b6040805160028082526060828101909352600092918291816020015b610df4611530565b815260200190600190039081610dec57505060408051600280825260608201909252919350602082015b610e26611680565b815260200190600190039081610e1e57905050905083826000815181101515610e4b57fe5b60209081029091010152610e5e8561116a565b826001815181101515610e6d57fe5b60209081029091010152610e7f6111ef565b816000815181101515610e8e57fe5b602090810290910181019190915260008781526008909152604090819020815160808101835291829081018260028282826020028201915b815481526020019060010190808311610ec657505050918352505060408051808201918290526020909201919060028481019182845b815481526020019060010190808311610efc57505050505081525050816001815181101515610f2757fe5b60209081029091010152610f3b82826112af565b9695505050505050565b3b90565b8051825160009184918491849114610f64576000935061100f565b5060005b825181101561100a578181815181101515610f7f57fe5b90602001015160f860020a900460f860020a027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168382815181101515610fc257fe5b60209101015160f860020a90819004027fff000000000000000000000000000000000000000000000000000000000000001614611002576000935061100f565b600101610f68565b600193505b50505092915050565b6040805160028082526060828101909352600092918291816020015b61103c611530565b81526020019060019003908161103457505060408051600280825260608201909252919350602082015b61106e611680565b8152602001906001900390816110665790505090508382600081518110151561109357fe5b602090810290910101526110a68561116a565b8260018151811015156110b557fe5b602090810290910101526110c76111ef565b8160008151811015156110d657fe5b6020908102919091018101919091526000878152600790915260409081902081516080810180845282549382019384529092839183906002906001830160608601808311610ec657505050918352505060408051808201918290526002848101805483526020948501949293909260038701908501808311610efc57505050505081525050816001815181101515610f2757fe5b611172611530565b6000826040518082805190602001908083835b602083106111a45780518252601f199092019160209182019101611185565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206001900490506111e86111e26114bb565b826114dc565b9392505050565b6111f7611680565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b82527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa60208381019190915281019190915290565b6000806000606060006112c06116a6565b86518851600091146112d157600080fd5b8851955085600602945084604051908082528060200260200182016040528015611305578160200160208202803883390190505b509350600092505b8583101561147a57888381518110151561132357fe5b6020908102909101015151845185906006860290811061133f57fe5b60209081029091010152885189908490811061135757fe5b9060200190602002015160200151848460060260010181518110151561137957fe5b60209081029091010152875188908490811061139157fe5b6020908102919091010151515184518590600260068702019081106113b257fe5b6020908102909101015287518890849081106113ca57fe5b60209081029190910181015151015184518590600360068702019081106113ed57fe5b60209081029091010152875188908490811061140557fe5b602090810291909101810151015151845185906004600687020190811061142857fe5b60209081029091010152875188908490811061144057fe5b6020908102919091018101518101510151845185906005600687020190811061146557fe5b6020908102909101015260019092019161130d565b6020826020870260208701600060086107d05a03f1905080801561149d5761149f565bfe5b508015156114ac57600080fd5b50511515979650505050505050565b6114c3611530565b5060408051808201909152600181526002602082015290565b6114e4611530565b6114ec6116c5565b83518152602080850151908201526040810183905260006060836080848460076107d05a03f1905080801561149d575080151561152857600080fd5b505092915050565b604080518082019091526000808252602082015290565b8260028101928215611572579182015b82811115611572578254825591600101919060010190611557565b5061157e9291506116e4565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115c357805160ff1916838001178555611572565b82800160010185558215611572579182015b828111156115725782518255916020019190600101906115d5565b815481835581811115611614576000838152602090206116149181019083016116e4565b505050565b826002810192821561157257916020028201828111156115725782518255916020019190600101906115d5565b82805482825590600052602060002090810192821561157257916020028201828111156115725782518255916020019190600101906115d5565b6080604051908101604052806116946116fe565b81526020016116a16116fe565b905290565b6020604051908101604052806001906020820280388339509192915050565b6060604051908101604052806003906020820280388339509192915050565b61044a91905b8082111561157e57600081556001016116ea565b604080518082018252906002908290803883395091929150505600a165627a7a72305820e8a32d284497b45c4220312fdffac70437f288d66e663eb6977820e3f3f713370029`

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

// GetBootstrapIp is a free data retrieval call binding the contract method 0x0322c2ad.
//
// Solidity: function getBootstrapIp() constant returns(string)
func (_DOSProxy *DOSProxyCaller) GetBootstrapIp(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "getBootstrapIp")
	return *ret0, err
}

// GetBootstrapIp is a free data retrieval call binding the contract method 0x0322c2ad.
//
// Solidity: function getBootstrapIp() constant returns(string)
func (_DOSProxy *DOSProxySession) GetBootstrapIp() (string, error) {
	return _DOSProxy.Contract.GetBootstrapIp(&_DOSProxy.CallOpts)
}

// GetBootstrapIp is a free data retrieval call binding the contract method 0x0322c2ad.
//
// Solidity: function getBootstrapIp() constant returns(string)
func (_DOSProxy *DOSProxyCallerSession) GetBootstrapIp() (string, error) {
	return _DOSProxy.Contract.GetBootstrapIp(&_DOSProxy.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x75d2c1ce.
//
// Solidity: function getPublicKey(group_id uint256) constant returns(uint256, uint256, uint256, uint256)
func (_DOSProxy *DOSProxyCaller) GetPublicKey(opts *bind.CallOpts, group_id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _DOSProxy.contract.Call(opts, out, "getPublicKey", group_id)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPublicKey is a free data retrieval call binding the contract method 0x75d2c1ce.
//
// Solidity: function getPublicKey(group_id uint256) constant returns(uint256, uint256, uint256, uint256)
func (_DOSProxy *DOSProxySession) GetPublicKey(group_id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _DOSProxy.Contract.GetPublicKey(&_DOSProxy.CallOpts, group_id)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x75d2c1ce.
//
// Solidity: function getPublicKey(group_id uint256) constant returns(uint256, uint256, uint256, uint256)
func (_DOSProxy *DOSProxyCallerSession) GetPublicKey(group_id *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _DOSProxy.Contract.GetPublicKey(&_DOSProxy.CallOpts, group_id)
}

// GetRandomNum is a free data retrieval call binding the contract method 0x25b814f4.
//
// Solidity: function getRandomNum() constant returns(uint256)
func (_DOSProxy *DOSProxyCaller) GetRandomNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "getRandomNum")
	return *ret0, err
}

// GetRandomNum is a free data retrieval call binding the contract method 0x25b814f4.
//
// Solidity: function getRandomNum() constant returns(uint256)
func (_DOSProxy *DOSProxySession) GetRandomNum() (*big.Int, error) {
	return _DOSProxy.Contract.GetRandomNum(&_DOSProxy.CallOpts)
}

// GetRandomNum is a free data retrieval call binding the contract method 0x25b814f4.
//
// Solidity: function getRandomNum() constant returns(uint256)
func (_DOSProxy *DOSProxyCallerSession) GetRandomNum() (*big.Int, error) {
	return _DOSProxy.Contract.GetRandomNum(&_DOSProxy.CallOpts)
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
// Solidity: function query(from address, block_number uint256, timeout uint256, query_type string, query_path string) returns()
func (_DOSProxy *DOSProxyTransactor) Query(opts *bind.TransactOpts, from common.Address, block_number *big.Int, timeout *big.Int, query_type string, query_path string) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "query", from, block_number, timeout, query_type, query_path)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query(from address, block_number uint256, timeout uint256, query_type string, query_path string) returns()
func (_DOSProxy *DOSProxySession) Query(from common.Address, block_number *big.Int, timeout *big.Int, query_type string, query_path string) (*types.Transaction, error) {
	return _DOSProxy.Contract.Query(&_DOSProxy.TransactOpts, from, block_number, timeout, query_type, query_path)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query(from address, block_number uint256, timeout uint256, query_type string, query_path string) returns()
func (_DOSProxy *DOSProxyTransactorSession) Query(from common.Address, block_number *big.Int, timeout *big.Int, query_type string, query_path string) (*types.Transaction, error) {
	return _DOSProxy.Contract.Query(&_DOSProxy.TransactOpts, from, block_number, timeout, query_type, query_path)
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

// SetBootstrapIp is a paid mutator transaction binding the contract method 0x872e8640.
//
// Solidity: function setBootstrapIp(ip string) returns()
func (_DOSProxy *DOSProxyTransactor) SetBootstrapIp(opts *bind.TransactOpts, ip string) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "setBootstrapIp", ip)
}

// SetBootstrapIp is a paid mutator transaction binding the contract method 0x872e8640.
//
// Solidity: function setBootstrapIp(ip string) returns()
func (_DOSProxy *DOSProxySession) SetBootstrapIp(ip string) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetBootstrapIp(&_DOSProxy.TransactOpts, ip)
}

// SetBootstrapIp is a paid mutator transaction binding the contract method 0x872e8640.
//
// Solidity: function setBootstrapIp(ip string) returns()
func (_DOSProxy *DOSProxyTransactorSession) SetBootstrapIp(ip string) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetBootstrapIp(&_DOSProxy.TransactOpts, ip)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0xc8a03ef2.
//
// Solidity: function setPublicKey(group_id uint256, x1 uint256, x2 uint256, y1 uint256, y2 uint256) returns()
func (_DOSProxy *DOSProxyTransactor) SetPublicKey(opts *bind.TransactOpts, group_id *big.Int, x1 *big.Int, x2 *big.Int, y1 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "setPublicKey", group_id, x1, x2, y1, y2)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0xc8a03ef2.
//
// Solidity: function setPublicKey(group_id uint256, x1 uint256, x2 uint256, y1 uint256, y2 uint256) returns()
func (_DOSProxy *DOSProxySession) SetPublicKey(group_id *big.Int, x1 *big.Int, x2 *big.Int, y1 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetPublicKey(&_DOSProxy.TransactOpts, group_id, x1, x2, y1, y2)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0xc8a03ef2.
//
// Solidity: function setPublicKey(group_id uint256, x1 uint256, x2 uint256, y1 uint256, y2 uint256) returns()
func (_DOSProxy *DOSProxyTransactorSession) SetPublicKey(group_id *big.Int, x1 *big.Int, x2 *big.Int, y1 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetPublicKey(&_DOSProxy.TransactOpts, group_id, x1, x2, y1, y2)
}

// SetRandomNum is a paid mutator transaction binding the contract method 0xdfb6f0e5.
//
// Solidity: function setRandomNum(group_id uint256, random uint256, x uint256, y uint256) returns()
func (_DOSProxy *DOSProxyTransactor) SetRandomNum(opts *bind.TransactOpts, group_id *big.Int, random *big.Int, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "setRandomNum", group_id, random, x, y)
}

// SetRandomNum is a paid mutator transaction binding the contract method 0xdfb6f0e5.
//
// Solidity: function setRandomNum(group_id uint256, random uint256, x uint256, y uint256) returns()
func (_DOSProxy *DOSProxySession) SetRandomNum(group_id *big.Int, random *big.Int, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetRandomNum(&_DOSProxy.TransactOpts, group_id, random, x, y)
}

// SetRandomNum is a paid mutator transaction binding the contract method 0xdfb6f0e5.
//
// Solidity: function setRandomNum(group_id uint256, random uint256, x uint256, y uint256) returns()
func (_DOSProxy *DOSProxyTransactorSession) SetRandomNum(group_id *big.Int, random *big.Int, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetRandomNum(&_DOSProxy.TransactOpts, group_id, random, x, y)
}

// ToBytes is a paid mutator transaction binding the contract method 0x775a8f5e.
//
// Solidity: function toBytes(num uint256) returns(numBytes bytes)
func (_DOSProxy *DOSProxyTransactor) ToBytes(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "toBytes", num)
}

// ToBytes is a paid mutator transaction binding the contract method 0x775a8f5e.
//
// Solidity: function toBytes(num uint256) returns(numBytes bytes)
func (_DOSProxy *DOSProxySession) ToBytes(num *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.ToBytes(&_DOSProxy.TransactOpts, num)
}

// ToBytes is a paid mutator transaction binding the contract method 0x775a8f5e.
//
// Solidity: function toBytes(num uint256) returns(numBytes bytes)
func (_DOSProxy *DOSProxyTransactorSession) ToBytes(num *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.ToBytes(&_DOSProxy.TransactOpts, num)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x34ebda4d.
//
// Solidity: function triggerCallback(query_id uint256, result bytes, x uint256, y uint256) returns()
func (_DOSProxy *DOSProxyTransactor) TriggerCallback(opts *bind.TransactOpts, query_id *big.Int, result []byte, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "triggerCallback", query_id, result, x, y)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x34ebda4d.
//
// Solidity: function triggerCallback(query_id uint256, result bytes, x uint256, y uint256) returns()
func (_DOSProxy *DOSProxySession) TriggerCallback(query_id *big.Int, result []byte, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.TriggerCallback(&_DOSProxy.TransactOpts, query_id, result, x, y)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x34ebda4d.
//
// Solidity: function triggerCallback(query_id uint256, result bytes, x uint256, y uint256) returns()
func (_DOSProxy *DOSProxyTransactorSession) TriggerCallback(query_id *big.Int, result []byte, x *big.Int, y *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.TriggerCallback(&_DOSProxy.TransactOpts, query_id, result, x, y)
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
	UserContractAddr common.Address
	Result           []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogCallbackTriggeredFor is a free log retrieval operation binding the contract event 0xcd714230b213422971bdd48f3fa7f63c52e50f9fa7356f6aa42a191c12f046d0.
//
// Solidity: e LogCallbackTriggeredFor(user_contract_addr address, result bytes)
func (_DOSProxy *DOSProxyFilterer) FilterLogCallbackTriggeredFor(opts *bind.FilterOpts) (*DOSProxyLogCallbackTriggeredForIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogCallbackTriggeredFor")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogCallbackTriggeredForIterator{contract: _DOSProxy.contract, event: "LogCallbackTriggeredFor", logs: logs, sub: sub}, nil
}

// WatchLogCallbackTriggeredFor is a free log subscription operation binding the contract event 0xcd714230b213422971bdd48f3fa7f63c52e50f9fa7356f6aa42a191c12f046d0.
//
// Solidity: e LogCallbackTriggeredFor(user_contract_addr address, result bytes)
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
	GroupId *big.Int
	NodeId  []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogGrouping is a free log retrieval operation binding the contract event 0xfdb627147562eb968c80eba5de51c105b83bf375bf584d31caf63648087a0ed9.
//
// Solidity: e LogGrouping(GroupId uint256, NodeId uint256[])
func (_DOSProxy *DOSProxyFilterer) FilterLogGrouping(opts *bind.FilterOpts) (*DOSProxyLogGroupingIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogGrouping")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogGroupingIterator{contract: _DOSProxy.contract, event: "LogGrouping", logs: logs, sub: sub}, nil
}

// WatchLogGrouping is a free log subscription operation binding the contract event 0xfdb627147562eb968c80eba5de51c105b83bf375bf584d31caf63648087a0ed9.
//
// Solidity: e LogGrouping(GroupId uint256, NodeId uint256[])
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

// DOSProxyLogInvalidSignatureIterator is returned from FilterLogInvalidSignature and is used to iterate over the raw logs and unpacked data for LogInvalidSignature events raised by the DOSProxy contract.
type DOSProxyLogInvalidSignatureIterator struct {
	Event *DOSProxyLogInvalidSignature // Event containing the contract specifics and raw log

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
func (it *DOSProxyLogInvalidSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogInvalidSignature)
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
		it.Event = new(DOSProxyLogInvalidSignature)
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
func (it *DOSProxyLogInvalidSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogInvalidSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogInvalidSignature represents a LogInvalidSignature event raised by the DOSProxy contract.
type DOSProxyLogInvalidSignature struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogInvalidSignature is a free log retrieval operation binding the contract event 0xf0cda705e46caa68e5854fa85a2635f77f3f6b5c927bd409ee7d935e4bb0322c.
//
// Solidity: e LogInvalidSignature()
func (_DOSProxy *DOSProxyFilterer) FilterLogInvalidSignature(opts *bind.FilterOpts) (*DOSProxyLogInvalidSignatureIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogInvalidSignature")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogInvalidSignatureIterator{contract: _DOSProxy.contract, event: "LogInvalidSignature", logs: logs, sub: sub}, nil
}

// WatchLogInvalidSignature is a free log subscription operation binding the contract event 0xf0cda705e46caa68e5854fa85a2635f77f3f6b5c927bd409ee7d935e4bb0322c.
//
// Solidity: e LogInvalidSignature()
func (_DOSProxy *DOSProxyFilterer) WatchLogInvalidSignature(opts *bind.WatchOpts, sink chan<- *DOSProxyLogInvalidSignature) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogInvalidSignature")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogInvalidSignature)
				if err := _DOSProxy.contract.UnpackLog(event, "LogInvalidSignature", log); err != nil {
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
// Solidity: e LogNonSupportedType(query_type string)
func (_DOSProxy *DOSProxyFilterer) FilterLogNonSupportedType(opts *bind.FilterOpts) (*DOSProxyLogNonSupportedTypeIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogNonSupportedType")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogNonSupportedTypeIterator{contract: _DOSProxy.contract, event: "LogNonSupportedType", logs: logs, sub: sub}, nil
}

// WatchLogNonSupportedType is a free log subscription operation binding the contract event 0x70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41.
//
// Solidity: e LogNonSupportedType(query_type string)
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
	GroupId *big.Int
	QueryId *big.Int
	Url     string
	Timeout *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogUrl is a free log retrieval operation binding the contract event 0xc24f120340e1beeac424b47e9e08f835661963ec3a5f7aa7edf537b71533b2ae.
//
// Solidity: e LogUrl(groupId uint256, queryId uint256, url string, timeout uint256)
func (_DOSProxy *DOSProxyFilterer) FilterLogUrl(opts *bind.FilterOpts) (*DOSProxyLogUrlIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogUrl")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogUrlIterator{contract: _DOSProxy.contract, event: "LogUrl", logs: logs, sub: sub}, nil
}

// WatchLogUrl is a free log subscription operation binding the contract event 0xc24f120340e1beeac424b47e9e08f835661963ec3a5f7aa7edf537b71533b2ae.
//
// Solidity: e LogUrl(groupId uint256, queryId uint256, url string, timeout uint256)
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
const UserContractInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// Callback_ is a paid mutator transaction binding the contract method 0x4a40f5f9.
//
// Solidity: function __callback__( bytes) returns()
func (_UserContractInterface *UserContractInterfaceTransactor) Callback_(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _UserContractInterface.contract.Transact(opts, "__callback__", arg0)
}

// Callback_ is a paid mutator transaction binding the contract method 0x4a40f5f9.
//
// Solidity: function __callback__( bytes) returns()
func (_UserContractInterface *UserContractInterfaceSession) Callback_(arg0 []byte) (*types.Transaction, error) {
	return _UserContractInterface.Contract.Callback_(&_UserContractInterface.TransactOpts, arg0)
}

// Callback_ is a paid mutator transaction binding the contract method 0x4a40f5f9.
//
// Solidity: function __callback__( bytes) returns()
func (_UserContractInterface *UserContractInterfaceTransactorSession) Callback_(arg0 []byte) (*types.Transaction, error) {
	return _UserContractInterface.Contract.Callback_(&_UserContractInterface.TransactOpts, arg0)
}
