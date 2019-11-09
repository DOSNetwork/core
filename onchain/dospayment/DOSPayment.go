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
const DospaymentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"_defaultGuardianFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"nodeTokenAddres\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_bridgeAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"refundServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"_paymentMethods\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guardianAddr\",\"type\":\"address\"}],\"name\":\"claimGuardianReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"paymentInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"serviceType\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeTokenAddresLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fundsAddr\",\"type\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setGuardianFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_guardianFundsTokenAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultTokenAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultSubmitterRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"requestID\",\"type\":\"uint256\"},{\"name\":\"serviceType\",\"type\":\"uint256\"}],\"name\":\"chargeServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"uint256\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_payments\",\"outputs\":[{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"serviceType\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultWorkerRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setPaymentMethod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"_feeList\",\"outputs\":[{\"name\":\"submitterRate\",\"type\":\"uint256\"},{\"name\":\"workerRate\",\"type\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\"},{\"name\":\"guardianFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultUserQueryFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultSystemRandomFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"submitterRate\",\"type\":\"uint256\"},{\"name\":\"workerRate\",\"type\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"setFeeDividend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_guardianFundsAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setGuardianFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_addressBridge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"nodeTokenAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultUserRandomFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestID\",\"type\":\"uint256\"},{\"name\":\"submitter\",\"type\":\"address\"},{\"name\":\"workers\",\"type\":\"address[]\"}],\"name\":\"claimServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"bridgeAddr\",\"type\":\"address\"},{\"name\":\"guardianFundsAddr\",\"type\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateNetworkTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateDropBurnTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogChargeServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogRefundServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"feeForSubmitter\",\"type\":\"uint256\"}],\"name\":\"LogClaimServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"feeForSubmitter\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LogClaimGuardianFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DospaymentBin is the compiled bytecode used for deploying new contracts.
const DospaymentBin = `60806040526005805560036006556002600755674563918244f40000600855674563918244f40000600955674563918244f40000600a55670de0b6b3a7640000600b5534801561004e57600080fd5b50604051606080611e518339810180604052606081101561006e57600080fd5b508051602080830151604093840151600080546001600160a01b031990811633178255600c80546001600160a01b03948516908316811791829055600d805496861696841696909617909555600e80548316909517909455928216815260028085528682206008548380529581905287832095909555600954600180845288842091909155600a548284529783209790975560065496850196909655600754958401959095556005546003840155600b546004909301929092556010805482169383169390931792839055600f8054939092169216919091179055611cf890819061015990396000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c80638da5cb5b11610125578063de439e9a116100ad578063f0612ffe1161007c578063f0612ffe146105c7578063f2fde38b146105cf578063f87059c1146105f5578063fe7b663514610621578063fedb8b6a1461062957610211565b8063de439e9a14610553578063e2edf58c1461058b578063e654490414610593578063eebede831461059b57610211565b8063a387722b116100f4578063a387722b146104c1578063b73a3f8f146104c9578063c5e7ffef146104f7578063d40f68d114610543578063d5fd1f0f1461054b57610211565b80638da5cb5b146104455780638f32d59b1461044d578063958e2d31146104555780639ed713b41461047257610211565b80633939c401116101a8578063610c3b9611610177578063610c3b96146103f3578063635c971d146103fb57806365e4604114610403578063715018a61461040b5780637aa9181b1461041357610211565b80633939c401146103655780633ccfd60b146103975780633e698ad51461039f5780635a1fa503146103c557610211565b806322e60ea5116101e457806322e60ea51461029f57806323ff34cb146102c5578063240028e8146102eb5780633157f16d1461032557610211565b806311bbe27614610216578063139bcedb146102305780631aab7941146102785780631efa5a9814610280575b600080fd5b61021e6106e1565b60408051918252519081900360200190f35b61025c6004803603604081101561024657600080fd5b506001600160a01b0381351690602001356106e7565b604080516001600160a01b039092168252519081900360200190f35b61025c610793565b61029d6004803603602081101561029657600080fd5b50356107a2565b005b61025c600480360360208110156102b557600080fd5b50356001600160a01b0316610a22565b61029d600480360360208110156102db57600080fd5b50356001600160a01b0316610a3d565b6103116004803603602081101561030157600080fd5b50356001600160a01b0316610cb9565b604080519115158252519081900360200190f35b6103426004803603602081101561033b57600080fd5b5035610d6b565b604080516001600160a01b03909316835260208301919091528051918290030190f35b61029d6004803603606081101561037b57600080fd5b506001600160a01b038135169060208101359060400135610d92565b61029d610e21565b61021e600480360360208110156103b557600080fd5b50356001600160a01b0316610f8e565b61029d600480360360408110156103db57600080fd5b506001600160a01b0381358116916020013516610fac565b61025c61104a565b61025c611059565b61021e611068565b61029d61106e565b61029d6004803603606081101561042957600080fd5b506001600160a01b0381351690602081013590604001356110c7565b61025c61131d565b61031161132c565b61029d6004803603602081101561046b57600080fd5b503561133d565b61048f6004803603602081101561048857600080fd5b503561141d565b604080516001600160a01b03958616815293909416602084015282840191909152606082015290519081900360800190f35b61021e611453565b61029d600480360360408110156104df57600080fd5b506001600160a01b0381358116916020013516611459565b61051d6004803603602081101561050d57600080fd5b50356001600160a01b03166114e6565b604080519485526020850193909352838301919091526060830152519081900360800190f35b61021e611511565b61021e611517565b61029d6004803603608081101561056957600080fd5b506001600160a01b03813516906020810135906040810135906060013561151d565b61025c6115b3565b61021e6115c2565b61029d600480360360408110156105b157600080fd5b506001600160a01b0381351690602001356115c8565b61025c611650565b61029d600480360360208110156105e557600080fd5b50356001600160a01b031661165f565b61021e6004803603604081101561060b57600080fd5b506001600160a01b03813516906020013561167c565b61021e611745565b61029d6004803603606081101561063f57600080fd5b8135916001600160a01b036020820135169181019060608101604082013564010000000081111561066f57600080fd5b82018360208201111561068157600080fd5b803590602001918460208302840111640100000000831117156106a357600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061174b945050505050565b600b5481565b6001600160a01b03821660009081526004602052604081206001015482106107505760408051600160e51b62461bcd0281526020600482015260116024820152600160781b704e6f20746f6b656e20616464726573732102604482015290519081900360640190fd5b6001600160a01b038316600090815260046020526040902060010180548390811061077757fe5b6000918252602090912001546001600160a01b03169392505050565b6010546001600160a01b031681565b6107aa61132c565b6107b357600080fd5b6000818152600360208190526040909120015481906108145760408051600160e51b62461bcd0281526020600482015260126024820152600160701b714e6f2066656520696e666f6d6174696f6e2102604482015290519081900360640190fd5b6000818152600360205260409020546001600160a01b03166108805760408051600160e51b62461bcd02815260206004820152601760248201527f4e6f20636f6e73756d657220696e666f6d6174696f6e21000000000000000000604482015290519081900360640190fd5b6000818152600360205260409020600101546001600160a01b03166108ef5760408051600160e51b62461bcd02815260206004820152601860248201527f4e6f20746f6b656e4164647220696e666f6d6174696f6e210000000000000000604482015290519081900360640190fd5b6000828152600360208181526040808420928301805460028501805460018701805488546001600160a01b0319808216909a55988116909155918890559690925582516001600160a01b039586168082529590921693820184905281830188905260608201869052608082018190529151919493917fde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa9181900360a00190a1816001600160a01b031663a9059cbb82866040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156109ee57600080fd5b505af1158015610a02573d6000803e3d6000fd5b505050506040513d6020811015610a1857600080fd5b5050505050505050565b6001602052600090815260409020546001600160a01b031681565b600f60009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a8b57600080fd5b505afa158015610a9f573d6000803e3d6000fd5b505050506040513d6020811015610ab557600080fd5b50516001600160a01b03163314610b0f5760408051600160e51b62461bcd0281526020600482015260136024820152600160681b724e6f742066726f6d20444f532070726f78792102604482015290519081900360640190fd5b600d546001600160a01b0316610b5957604051600160e51b62461bcd028152600401808060200182810382526023815260200180611caa6023913960400191505060405180910390fd5b600e546001600160a01b0316610bb95760408051600160e51b62461bcd02815260206004820152601a60248201527f4e6f7420612076616c696420746f6b656e206164647265737321000000000000604482015290519081900360640190fd5b600e546001600160a01b039081166000818152600260209081526040918290206004015482519486168552908401929092528281018290523360608401525190917f47ad88344c408450ef0ccab93ed97dd83af7a27dedfaa205c0725cfc4ca819cd919081900360800190a1600e54600d5460408051600160e01b6323b872dd0281526001600160a01b039283166004820152858316602482015260448101859052905191909216916323b872dd9160648083019260209291908290030181600087803b158015610c8957600080fd5b505af1158015610c9d573d6000803e3d6000fd5b505050506040513d6020811015610cb357600080fd5b50505050565b60006001600160a01b038216610cd157506000610d66565b6001600160a01b0382166000908152600260209081526040808320838052909152902054610d0157506000610d66565b6001600160a01b038216600090815260026020908152604080832060018452909152902054610d3257506000610d66565b6001600160a01b0382166000908152600260208181526040808420928452919052902054610d6257506000610d66565b5060015b919050565b600090815260036020819052604090912060018101549101546001600160a01b0390911691565b610d9a61132c565b610da357600080fd5b6001600160a01b038316610dfb5760408051600160e51b62461bcd0281526020600482015260146024820152600160601b734e6f7420612076616c696420616464726573732102604482015290519081900360640190fd5b6001600160a01b0390921660009081526002602090815260408083209383529290522055565b33600090815260046020526040902060010154610e795760408051600160e51b62461bcd02815260206004820152600b6024820152600160a81b6a4e6f20726577617264732102604482015290519081900360640190fd5b3360009081526004602052604081206001810180549192916000198101908110610e9f57fe5b6000918252602090912001546001830180546001600160a01b0390921692509080610ec657fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03831682528390526040812080549190558015610f895760408051600160e01b63a9059cbb0281523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b158015610f5c57600080fd5b505af1158015610f70573d6000803e3d6000fd5b505050506040513d6020811015610f8657600080fd5b50505b505050565b6001600160a01b031660009081526004602052604090206001015490565b610fb461132c565b610fbd57600080fd5b80610fc781610cb9565b61101b5760408051600160e51b62461bcd02815260206004820152601e60248201527f4e6f74206120737570706f7274656420746f6b656e2061646472657373210000604482015290519081900360640190fd5b50600d80546001600160a01b039384166001600160a01b031991821617909155600e8054929093169116179055565b600e546001600160a01b031681565b600c546001600160a01b031681565b60065481565b61107661132c565b61107f57600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b600f60009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561111557600080fd5b505afa158015611129573d6000803e3d6000fd5b505050506040513d602081101561113f57600080fd5b50516001600160a01b031633146111995760408051600160e51b62461bcd0281526020600482015260136024820152600160681b724e6f742066726f6d20444f532070726f78792102604482015290519081900360640190fd5b6001600160a01b03808416600090815260016020526040902054166111bd81610cb9565b6112115760408051600160e51b62461bcd02815260206004820152601960248201527f4e6f7420612076616c696420746f6b656e206164647265737300000000000000604482015290519081900360640190fd5b6001600160a01b0380821660008181526002602081815260408084208885528252808420548985526003808452948290208054978c166001600160a01b0319988916811782559481018a90556001810180549098168717909755938601849055805192835290820193909352808301879052606081018690526080810182905291519092917fa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba919081900360a00190a160408051600160e01b6323b872dd0281526001600160a01b038881166004830152306024830152604482018590529151918516916323b872dd916064808201926020929091908290030181600087803b1580156109ee57600080fd5b6000546001600160a01b031690565b6000546001600160a01b0316331490565b61134561132c565b61134e57600080fd5b60408051600160e01b6370a0823102815230600482015290516000916001600160a01b038416916370a0823191602480820192602092909190829003018186803b15801561139b57600080fd5b505afa1580156113af573d6000803e3d6000fd5b505050506040513d60208110156113c557600080fd5b505160408051600160e01b63a9059cbb0281523360048201526024810183905290519192506001600160a01b0384169163a9059cbb916044808201926020929091908290030181600087803b158015610c8957600080fd5b600360208190526000918252604090912080546001820154600283015492909301546001600160a01b0391821693909116919084565b60075481565b8061146381610cb9565b6114b75760408051600160e51b62461bcd02815260206004820152601e60248201527f4e6f74206120737570706f7274656420746f6b656e2061646472657373210000604482015290519081900360640190fd5b506001600160a01b03918216600090815260016020526040902080546001600160a01b03191691909216179055565b6002602081905260009182526040909120600181015491810154600382015460049092015490919084565b600a5481565b60085481565b61152561132c565b61152e57600080fd5b6001600160a01b0384166115865760408051600160e51b62461bcd0281526020600482015260146024820152600160601b734e6f7420612076616c696420616464726573732102604482015290519081900360640190fd5b6001600160a01b039093166000908152600260208190526040909120600181019390935582015560030155565b600d546001600160a01b031681565b60055481565b6115d061132c565b6115d957600080fd5b6001600160a01b0382166116315760408051600160e51b62461bcd0281526020600482015260146024820152600160601b734e6f7420612076616c696420616464726573732102604482015290519081900360640190fd5b6001600160a01b03909116600090815260026020526040902060040155565b600f546001600160a01b031681565b61166761132c565b61167057600080fd5b61167981611ba4565b50565b6001600160a01b03821660009081526004602052604081206001015482106116e55760408051600160e51b62461bcd0281526020600482015260116024820152600160781b704e6f20746f6b656e20616464726573732102604482015290519081900360640190fd5b6001600160a01b038316600090815260046020526040812060010180548490811061170c57fe5b60009182526020808320909101546001600160a01b03878116845260048352604080852091909216845290915290205491505092915050565b60095481565b600f60009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561179957600080fd5b505afa1580156117ad573d6000803e3d6000fd5b505050506040513d60208110156117c357600080fd5b50516001600160a01b0316331461181d5760408051600160e51b62461bcd0281526020600482015260136024820152600160681b724e6f742066726f6d20444f532070726f78792102604482015290519081900360640190fd5b60008381526003602081905260409091200154839061187e5760408051600160e51b62461bcd0281526020600482015260126024820152600160701b714e6f2066656520696e666f6d6174696f6e2102604482015290519081900360640190fd5b6000818152600360205260409020546001600160a01b03166118ea5760408051600160e51b62461bcd02815260206004820152601760248201527f4e6f20636f6e73756d657220696e666f6d6174696f6e21000000000000000000604482015290519081900360640190fd5b6000818152600360205260409020600101546001600160a01b03166119595760408051600160e51b62461bcd02815260206004820152601860248201527f4e6f20746f6b656e4164647220696e666f6d6174696f6e210000000000000000604482015290519081900360640190fd5b60008481526003602081905260408220600181018054928201805460028401805485546001600160a01b031990811690965594861690935591859055939093556001600160a01b0390911691906119ae611c81565b506001600160a01b03831660009081526002602081815260409283902083516080810185526001820154815292810154918301919091526003810154928201839052600401546060820152908381611a0257fe5b60208301518851929091049450840290600090600019018281611a2157fe5b6001600160a01b038b1660009081526004602052604090209190049150611a49908784611c12565b604080516001600160a01b03808c168252881660208201528082018c9052606081018690526080810184905290517fab9fe896064c2c9dfd31acebc7d522b311e5f2e7d1b4965ac0cfd5a4abec813a9181900360a00190a160005b8851811015611b9757896001600160a01b0316898281518110611ac357fe5b60200260200101516001600160a01b031614611b8f57611b1b600460008b8481518110611aec57fe5b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000208884611c12565b7fab9fe896064c2c9dfd31acebc7d522b311e5f2e7d1b4965ac0cfd5a4abec813a898281518110611b4857fe5b602090810291909101810151604080516001600160a01b039283168152918b16928201929092528082018e9052606081018890526080810186905290519081900360a00190a15b600101611aa4565b5050505050505050505050565b6001600160a01b038116611bb757600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03821660009081526020849052604090205480611c60576001848101805491820181556000908152602090200180546001600160a01b0319166001600160a01b0385161790555b6001600160a01b039092166000908152602093909352604090922091019055565b604051806080016040528060008152602001600081526020016000815260200160008152509056fe4e6f7420612076616c696420677561726469616e2066756e6473206164647265737321a165627a7a723058203122a9c021eb48869467c350e1352c8b07ba82a806da4417ef3d173f109e18ef0029`

// DeployDospayment deploys a new Ethereum contract, binding an instance of Dospayment to it.
func DeployDospayment(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddr common.Address, guardianFundsAddr common.Address, tokenAddr common.Address) (common.Address, *types.Transaction, *Dospayment, error) {
	parsed, err := abi.JSON(strings.NewReader(DospaymentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DospaymentBin), backend, bridgeAddr, guardianFundsAddr, tokenAddr)
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

// AddressBridge is a free data retrieval call binding the contract method 0xf0612ffe.
//
// Solidity: function _addressBridge() constant returns(address)
func (_Dospayment *DospaymentCaller) AddressBridge(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_addressBridge")
	return *ret0, err
}

// AddressBridge is a free data retrieval call binding the contract method 0xf0612ffe.
//
// Solidity: function _addressBridge() constant returns(address)
func (_Dospayment *DospaymentSession) AddressBridge() (common.Address, error) {
	return _Dospayment.Contract.AddressBridge(&_Dospayment.CallOpts)
}

// AddressBridge is a free data retrieval call binding the contract method 0xf0612ffe.
//
// Solidity: function _addressBridge() constant returns(address)
func (_Dospayment *DospaymentCallerSession) AddressBridge() (common.Address, error) {
	return _Dospayment.Contract.AddressBridge(&_Dospayment.CallOpts)
}

// BridgeAddr is a free data retrieval call binding the contract method 0x1aab7941.
//
// Solidity: function _bridgeAddr() constant returns(address)
func (_Dospayment *DospaymentCaller) BridgeAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_bridgeAddr")
	return *ret0, err
}

// BridgeAddr is a free data retrieval call binding the contract method 0x1aab7941.
//
// Solidity: function _bridgeAddr() constant returns(address)
func (_Dospayment *DospaymentSession) BridgeAddr() (common.Address, error) {
	return _Dospayment.Contract.BridgeAddr(&_Dospayment.CallOpts)
}

// BridgeAddr is a free data retrieval call binding the contract method 0x1aab7941.
//
// Solidity: function _bridgeAddr() constant returns(address)
func (_Dospayment *DospaymentCallerSession) BridgeAddr() (common.Address, error) {
	return _Dospayment.Contract.BridgeAddr(&_Dospayment.CallOpts)
}

// DefaultDenominator is a free data retrieval call binding the contract method 0xe6544904.
//
// Solidity: function _defaultDenominator() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_defaultDenominator")
	return *ret0, err
}

// DefaultDenominator is a free data retrieval call binding the contract method 0xe6544904.
//
// Solidity: function _defaultDenominator() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultDenominator() (*big.Int, error) {
	return _Dospayment.Contract.DefaultDenominator(&_Dospayment.CallOpts)
}

// DefaultDenominator is a free data retrieval call binding the contract method 0xe6544904.
//
// Solidity: function _defaultDenominator() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultDenominator() (*big.Int, error) {
	return _Dospayment.Contract.DefaultDenominator(&_Dospayment.CallOpts)
}

// DefaultGuardianFee is a free data retrieval call binding the contract method 0x11bbe276.
//
// Solidity: function _defaultGuardianFee() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultGuardianFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_defaultGuardianFee")
	return *ret0, err
}

// DefaultGuardianFee is a free data retrieval call binding the contract method 0x11bbe276.
//
// Solidity: function _defaultGuardianFee() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultGuardianFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultGuardianFee(&_Dospayment.CallOpts)
}

// DefaultGuardianFee is a free data retrieval call binding the contract method 0x11bbe276.
//
// Solidity: function _defaultGuardianFee() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultGuardianFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultGuardianFee(&_Dospayment.CallOpts)
}

// DefaultSubmitterRate is a free data retrieval call binding the contract method 0x65e46041.
//
// Solidity: function _defaultSubmitterRate() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultSubmitterRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_defaultSubmitterRate")
	return *ret0, err
}

// DefaultSubmitterRate is a free data retrieval call binding the contract method 0x65e46041.
//
// Solidity: function _defaultSubmitterRate() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultSubmitterRate() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSubmitterRate(&_Dospayment.CallOpts)
}

// DefaultSubmitterRate is a free data retrieval call binding the contract method 0x65e46041.
//
// Solidity: function _defaultSubmitterRate() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultSubmitterRate() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSubmitterRate(&_Dospayment.CallOpts)
}

// DefaultSystemRandomFee is a free data retrieval call binding the contract method 0xd5fd1f0f.
//
// Solidity: function _defaultSystemRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultSystemRandomFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_defaultSystemRandomFee")
	return *ret0, err
}

// DefaultSystemRandomFee is a free data retrieval call binding the contract method 0xd5fd1f0f.
//
// Solidity: function _defaultSystemRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultSystemRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSystemRandomFee(&_Dospayment.CallOpts)
}

// DefaultSystemRandomFee is a free data retrieval call binding the contract method 0xd5fd1f0f.
//
// Solidity: function _defaultSystemRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultSystemRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSystemRandomFee(&_Dospayment.CallOpts)
}

// DefaultTokenAddr is a free data retrieval call binding the contract method 0x635c971d.
//
// Solidity: function _defaultTokenAddr() constant returns(address)
func (_Dospayment *DospaymentCaller) DefaultTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_defaultTokenAddr")
	return *ret0, err
}

// DefaultTokenAddr is a free data retrieval call binding the contract method 0x635c971d.
//
// Solidity: function _defaultTokenAddr() constant returns(address)
func (_Dospayment *DospaymentSession) DefaultTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.DefaultTokenAddr(&_Dospayment.CallOpts)
}

// DefaultTokenAddr is a free data retrieval call binding the contract method 0x635c971d.
//
// Solidity: function _defaultTokenAddr() constant returns(address)
func (_Dospayment *DospaymentCallerSession) DefaultTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.DefaultTokenAddr(&_Dospayment.CallOpts)
}

// DefaultUserQueryFee is a free data retrieval call binding the contract method 0xd40f68d1.
//
// Solidity: function _defaultUserQueryFee() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultUserQueryFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_defaultUserQueryFee")
	return *ret0, err
}

// DefaultUserQueryFee is a free data retrieval call binding the contract method 0xd40f68d1.
//
// Solidity: function _defaultUserQueryFee() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultUserQueryFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserQueryFee(&_Dospayment.CallOpts)
}

// DefaultUserQueryFee is a free data retrieval call binding the contract method 0xd40f68d1.
//
// Solidity: function _defaultUserQueryFee() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultUserQueryFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserQueryFee(&_Dospayment.CallOpts)
}

// DefaultUserRandomFee is a free data retrieval call binding the contract method 0xfe7b6635.
//
// Solidity: function _defaultUserRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultUserRandomFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_defaultUserRandomFee")
	return *ret0, err
}

// DefaultUserRandomFee is a free data retrieval call binding the contract method 0xfe7b6635.
//
// Solidity: function _defaultUserRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultUserRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserRandomFee(&_Dospayment.CallOpts)
}

// DefaultUserRandomFee is a free data retrieval call binding the contract method 0xfe7b6635.
//
// Solidity: function _defaultUserRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultUserRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserRandomFee(&_Dospayment.CallOpts)
}

// DefaultWorkerRate is a free data retrieval call binding the contract method 0xa387722b.
//
// Solidity: function _defaultWorkerRate() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultWorkerRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_defaultWorkerRate")
	return *ret0, err
}

// DefaultWorkerRate is a free data retrieval call binding the contract method 0xa387722b.
//
// Solidity: function _defaultWorkerRate() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultWorkerRate() (*big.Int, error) {
	return _Dospayment.Contract.DefaultWorkerRate(&_Dospayment.CallOpts)
}

// DefaultWorkerRate is a free data retrieval call binding the contract method 0xa387722b.
//
// Solidity: function _defaultWorkerRate() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultWorkerRate() (*big.Int, error) {
	return _Dospayment.Contract.DefaultWorkerRate(&_Dospayment.CallOpts)
}

// FeeList is a free data retrieval call binding the contract method 0xc5e7ffef.
//
// Solidity: function _feeList(address ) constant returns(uint256 submitterRate, uint256 workerRate, uint256 denominator, uint256 guardianFee)
func (_Dospayment *DospaymentCaller) FeeList(opts *bind.CallOpts, arg0 common.Address) (struct {
	SubmitterRate *big.Int
	WorkerRate    *big.Int
	Denominator   *big.Int
	GuardianFee   *big.Int
}, error) {
	ret := new(struct {
		SubmitterRate *big.Int
		WorkerRate    *big.Int
		Denominator   *big.Int
		GuardianFee   *big.Int
	})
	out := ret
	err := _Dospayment.contract.Call(opts, out, "_feeList", arg0)
	return *ret, err
}

// FeeList is a free data retrieval call binding the contract method 0xc5e7ffef.
//
// Solidity: function _feeList(address ) constant returns(uint256 submitterRate, uint256 workerRate, uint256 denominator, uint256 guardianFee)
func (_Dospayment *DospaymentSession) FeeList(arg0 common.Address) (struct {
	SubmitterRate *big.Int
	WorkerRate    *big.Int
	Denominator   *big.Int
	GuardianFee   *big.Int
}, error) {
	return _Dospayment.Contract.FeeList(&_Dospayment.CallOpts, arg0)
}

// FeeList is a free data retrieval call binding the contract method 0xc5e7ffef.
//
// Solidity: function _feeList(address ) constant returns(uint256 submitterRate, uint256 workerRate, uint256 denominator, uint256 guardianFee)
func (_Dospayment *DospaymentCallerSession) FeeList(arg0 common.Address) (struct {
	SubmitterRate *big.Int
	WorkerRate    *big.Int
	Denominator   *big.Int
	GuardianFee   *big.Int
}, error) {
	return _Dospayment.Contract.FeeList(&_Dospayment.CallOpts, arg0)
}

// GuardianFundsAddr is a free data retrieval call binding the contract method 0xe2edf58c.
//
// Solidity: function _guardianFundsAddr() constant returns(address)
func (_Dospayment *DospaymentCaller) GuardianFundsAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_guardianFundsAddr")
	return *ret0, err
}

// GuardianFundsAddr is a free data retrieval call binding the contract method 0xe2edf58c.
//
// Solidity: function _guardianFundsAddr() constant returns(address)
func (_Dospayment *DospaymentSession) GuardianFundsAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsAddr(&_Dospayment.CallOpts)
}

// GuardianFundsAddr is a free data retrieval call binding the contract method 0xe2edf58c.
//
// Solidity: function _guardianFundsAddr() constant returns(address)
func (_Dospayment *DospaymentCallerSession) GuardianFundsAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsAddr(&_Dospayment.CallOpts)
}

// GuardianFundsTokenAddr is a free data retrieval call binding the contract method 0x610c3b96.
//
// Solidity: function _guardianFundsTokenAddr() constant returns(address)
func (_Dospayment *DospaymentCaller) GuardianFundsTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_guardianFundsTokenAddr")
	return *ret0, err
}

// GuardianFundsTokenAddr is a free data retrieval call binding the contract method 0x610c3b96.
//
// Solidity: function _guardianFundsTokenAddr() constant returns(address)
func (_Dospayment *DospaymentSession) GuardianFundsTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsTokenAddr(&_Dospayment.CallOpts)
}

// GuardianFundsTokenAddr is a free data retrieval call binding the contract method 0x610c3b96.
//
// Solidity: function _guardianFundsTokenAddr() constant returns(address)
func (_Dospayment *DospaymentCallerSession) GuardianFundsTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsTokenAddr(&_Dospayment.CallOpts)
}

// PaymentMethods is a free data retrieval call binding the contract method 0x22e60ea5.
//
// Solidity: function _paymentMethods(address ) constant returns(address)
func (_Dospayment *DospaymentCaller) PaymentMethods(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "_paymentMethods", arg0)
	return *ret0, err
}

// PaymentMethods is a free data retrieval call binding the contract method 0x22e60ea5.
//
// Solidity: function _paymentMethods(address ) constant returns(address)
func (_Dospayment *DospaymentSession) PaymentMethods(arg0 common.Address) (common.Address, error) {
	return _Dospayment.Contract.PaymentMethods(&_Dospayment.CallOpts, arg0)
}

// PaymentMethods is a free data retrieval call binding the contract method 0x22e60ea5.
//
// Solidity: function _paymentMethods(address ) constant returns(address)
func (_Dospayment *DospaymentCallerSession) PaymentMethods(arg0 common.Address) (common.Address, error) {
	return _Dospayment.Contract.PaymentMethods(&_Dospayment.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x9ed713b4.
//
// Solidity: function _payments(uint256 ) constant returns(address consumer, address tokenAddr, uint256 serviceType, uint256 amount)
func (_Dospayment *DospaymentCaller) Payments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Consumer    common.Address
	TokenAddr   common.Address
	ServiceType *big.Int
	Amount      *big.Int
}, error) {
	ret := new(struct {
		Consumer    common.Address
		TokenAddr   common.Address
		ServiceType *big.Int
		Amount      *big.Int
	})
	out := ret
	err := _Dospayment.contract.Call(opts, out, "_payments", arg0)
	return *ret, err
}

// Payments is a free data retrieval call binding the contract method 0x9ed713b4.
//
// Solidity: function _payments(uint256 ) constant returns(address consumer, address tokenAddr, uint256 serviceType, uint256 amount)
func (_Dospayment *DospaymentSession) Payments(arg0 *big.Int) (struct {
	Consumer    common.Address
	TokenAddr   common.Address
	ServiceType *big.Int
	Amount      *big.Int
}, error) {
	return _Dospayment.Contract.Payments(&_Dospayment.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x9ed713b4.
//
// Solidity: function _payments(uint256 ) constant returns(address consumer, address tokenAddr, uint256 serviceType, uint256 amount)
func (_Dospayment *DospaymentCallerSession) Payments(arg0 *big.Int) (struct {
	Consumer    common.Address
	TokenAddr   common.Address
	ServiceType *big.Int
	Amount      *big.Int
}, error) {
	return _Dospayment.Contract.Payments(&_Dospayment.CallOpts, arg0)
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

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(address tokenAddr) constant returns(bool)
func (_Dospayment *DospaymentCaller) IsSupportedToken(opts *bind.CallOpts, tokenAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "isSupportedToken", tokenAddr)
	return *ret0, err
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(address tokenAddr) constant returns(bool)
func (_Dospayment *DospaymentSession) IsSupportedToken(tokenAddr common.Address) (bool, error) {
	return _Dospayment.Contract.IsSupportedToken(&_Dospayment.CallOpts, tokenAddr)
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(address tokenAddr) constant returns(bool)
func (_Dospayment *DospaymentCallerSession) IsSupportedToken(tokenAddr common.Address) (bool, error) {
	return _Dospayment.Contract.IsSupportedToken(&_Dospayment.CallOpts, tokenAddr)
}

// NodeTokenAddres is a free data retrieval call binding the contract method 0x139bcedb.
//
// Solidity: function nodeTokenAddres(address nodeAddr, uint256 idx) constant returns(address)
func (_Dospayment *DospaymentCaller) NodeTokenAddres(opts *bind.CallOpts, nodeAddr common.Address, idx *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "nodeTokenAddres", nodeAddr, idx)
	return *ret0, err
}

// NodeTokenAddres is a free data retrieval call binding the contract method 0x139bcedb.
//
// Solidity: function nodeTokenAddres(address nodeAddr, uint256 idx) constant returns(address)
func (_Dospayment *DospaymentSession) NodeTokenAddres(nodeAddr common.Address, idx *big.Int) (common.Address, error) {
	return _Dospayment.Contract.NodeTokenAddres(&_Dospayment.CallOpts, nodeAddr, idx)
}

// NodeTokenAddres is a free data retrieval call binding the contract method 0x139bcedb.
//
// Solidity: function nodeTokenAddres(address nodeAddr, uint256 idx) constant returns(address)
func (_Dospayment *DospaymentCallerSession) NodeTokenAddres(nodeAddr common.Address, idx *big.Int) (common.Address, error) {
	return _Dospayment.Contract.NodeTokenAddres(&_Dospayment.CallOpts, nodeAddr, idx)
}

// NodeTokenAddresLength is a free data retrieval call binding the contract method 0x3e698ad5.
//
// Solidity: function nodeTokenAddresLength(address nodeAddr) constant returns(uint256)
func (_Dospayment *DospaymentCaller) NodeTokenAddresLength(opts *bind.CallOpts, nodeAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "nodeTokenAddresLength", nodeAddr)
	return *ret0, err
}

// NodeTokenAddresLength is a free data retrieval call binding the contract method 0x3e698ad5.
//
// Solidity: function nodeTokenAddresLength(address nodeAddr) constant returns(uint256)
func (_Dospayment *DospaymentSession) NodeTokenAddresLength(nodeAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeTokenAddresLength(&_Dospayment.CallOpts, nodeAddr)
}

// NodeTokenAddresLength is a free data retrieval call binding the contract method 0x3e698ad5.
//
// Solidity: function nodeTokenAddresLength(address nodeAddr) constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) NodeTokenAddresLength(nodeAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeTokenAddresLength(&_Dospayment.CallOpts, nodeAddr)
}

// NodeTokenAmount is a free data retrieval call binding the contract method 0xf87059c1.
//
// Solidity: function nodeTokenAmount(address nodeAddr, uint256 idx) constant returns(uint256)
func (_Dospayment *DospaymentCaller) NodeTokenAmount(opts *bind.CallOpts, nodeAddr common.Address, idx *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "nodeTokenAmount", nodeAddr, idx)
	return *ret0, err
}

// NodeTokenAmount is a free data retrieval call binding the contract method 0xf87059c1.
//
// Solidity: function nodeTokenAmount(address nodeAddr, uint256 idx) constant returns(uint256)
func (_Dospayment *DospaymentSession) NodeTokenAmount(nodeAddr common.Address, idx *big.Int) (*big.Int, error) {
	return _Dospayment.Contract.NodeTokenAmount(&_Dospayment.CallOpts, nodeAddr, idx)
}

// NodeTokenAmount is a free data retrieval call binding the contract method 0xf87059c1.
//
// Solidity: function nodeTokenAmount(address nodeAddr, uint256 idx) constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) NodeTokenAmount(nodeAddr common.Address, idx *big.Int) (*big.Int, error) {
	return _Dospayment.Contract.NodeTokenAmount(&_Dospayment.CallOpts, nodeAddr, idx)
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

// PaymentInfo is a free data retrieval call binding the contract method 0x3157f16d.
//
// Solidity: function paymentInfo(uint256 requestID) constant returns(address, uint256)
func (_Dospayment *DospaymentCaller) PaymentInfo(opts *bind.CallOpts, requestID *big.Int) (common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Dospayment.contract.Call(opts, out, "paymentInfo", requestID)
	return *ret0, *ret1, err
}

// PaymentInfo is a free data retrieval call binding the contract method 0x3157f16d.
//
// Solidity: function paymentInfo(uint256 requestID) constant returns(address, uint256)
func (_Dospayment *DospaymentSession) PaymentInfo(requestID *big.Int) (common.Address, *big.Int, error) {
	return _Dospayment.Contract.PaymentInfo(&_Dospayment.CallOpts, requestID)
}

// PaymentInfo is a free data retrieval call binding the contract method 0x3157f16d.
//
// Solidity: function paymentInfo(uint256 requestID) constant returns(address, uint256)
func (_Dospayment *DospaymentCallerSession) PaymentInfo(requestID *big.Int) (common.Address, *big.Int, error) {
	return _Dospayment.Contract.PaymentInfo(&_Dospayment.CallOpts, requestID)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(address consumer, uint256 requestID, uint256 serviceType) returns()
func (_Dospayment *DospaymentTransactor) ChargeServiceFee(opts *bind.TransactOpts, consumer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "chargeServiceFee", consumer, requestID, serviceType)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(address consumer, uint256 requestID, uint256 serviceType) returns()
func (_Dospayment *DospaymentSession) ChargeServiceFee(consumer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.ChargeServiceFee(&_Dospayment.TransactOpts, consumer, requestID, serviceType)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(address consumer, uint256 requestID, uint256 serviceType) returns()
func (_Dospayment *DospaymentTransactorSession) ChargeServiceFee(consumer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.ChargeServiceFee(&_Dospayment.TransactOpts, consumer, requestID, serviceType)
}

// ClaimGuardianReward is a paid mutator transaction binding the contract method 0x23ff34cb.
//
// Solidity: function claimGuardianReward(address guardianAddr) returns()
func (_Dospayment *DospaymentTransactor) ClaimGuardianReward(opts *bind.TransactOpts, guardianAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "claimGuardianReward", guardianAddr)
}

// ClaimGuardianReward is a paid mutator transaction binding the contract method 0x23ff34cb.
//
// Solidity: function claimGuardianReward(address guardianAddr) returns()
func (_Dospayment *DospaymentSession) ClaimGuardianReward(guardianAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.ClaimGuardianReward(&_Dospayment.TransactOpts, guardianAddr)
}

// ClaimGuardianReward is a paid mutator transaction binding the contract method 0x23ff34cb.
//
// Solidity: function claimGuardianReward(address guardianAddr) returns()
func (_Dospayment *DospaymentTransactorSession) ClaimGuardianReward(guardianAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.ClaimGuardianReward(&_Dospayment.TransactOpts, guardianAddr)
}

// ClaimServiceFee is a paid mutator transaction binding the contract method 0xfedb8b6a.
//
// Solidity: function claimServiceFee(uint256 requestID, address submitter, address[] workers) returns()
func (_Dospayment *DospaymentTransactor) ClaimServiceFee(opts *bind.TransactOpts, requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "claimServiceFee", requestID, submitter, workers)
}

// ClaimServiceFee is a paid mutator transaction binding the contract method 0xfedb8b6a.
//
// Solidity: function claimServiceFee(uint256 requestID, address submitter, address[] workers) returns()
func (_Dospayment *DospaymentSession) ClaimServiceFee(requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.ClaimServiceFee(&_Dospayment.TransactOpts, requestID, submitter, workers)
}

// ClaimServiceFee is a paid mutator transaction binding the contract method 0xfedb8b6a.
//
// Solidity: function claimServiceFee(uint256 requestID, address submitter, address[] workers) returns()
func (_Dospayment *DospaymentTransactorSession) ClaimServiceFee(requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.ClaimServiceFee(&_Dospayment.TransactOpts, requestID, submitter, workers)
}

// RefundServiceFee is a paid mutator transaction binding the contract method 0x1efa5a98.
//
// Solidity: function refundServiceFee(uint256 requestID) returns()
func (_Dospayment *DospaymentTransactor) RefundServiceFee(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "refundServiceFee", requestID)
}

// RefundServiceFee is a paid mutator transaction binding the contract method 0x1efa5a98.
//
// Solidity: function refundServiceFee(uint256 requestID) returns()
func (_Dospayment *DospaymentSession) RefundServiceFee(requestID *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.RefundServiceFee(&_Dospayment.TransactOpts, requestID)
}

// RefundServiceFee is a paid mutator transaction binding the contract method 0x1efa5a98.
//
// Solidity: function refundServiceFee(uint256 requestID) returns()
func (_Dospayment *DospaymentTransactorSession) RefundServiceFee(requestID *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.RefundServiceFee(&_Dospayment.TransactOpts, requestID)
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

// SetFeeDividend is a paid mutator transaction binding the contract method 0xde439e9a.
//
// Solidity: function setFeeDividend(address tokenAddr, uint256 submitterRate, uint256 workerRate, uint256 denominator) returns()
func (_Dospayment *DospaymentTransactor) SetFeeDividend(opts *bind.TransactOpts, tokenAddr common.Address, submitterRate *big.Int, workerRate *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setFeeDividend", tokenAddr, submitterRate, workerRate, denominator)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0xde439e9a.
//
// Solidity: function setFeeDividend(address tokenAddr, uint256 submitterRate, uint256 workerRate, uint256 denominator) returns()
func (_Dospayment *DospaymentSession) SetFeeDividend(tokenAddr common.Address, submitterRate *big.Int, workerRate *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetFeeDividend(&_Dospayment.TransactOpts, tokenAddr, submitterRate, workerRate, denominator)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0xde439e9a.
//
// Solidity: function setFeeDividend(address tokenAddr, uint256 submitterRate, uint256 workerRate, uint256 denominator) returns()
func (_Dospayment *DospaymentTransactorSession) SetFeeDividend(tokenAddr common.Address, submitterRate *big.Int, workerRate *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetFeeDividend(&_Dospayment.TransactOpts, tokenAddr, submitterRate, workerRate, denominator)
}

// SetGuardianFee is a paid mutator transaction binding the contract method 0xeebede83.
//
// Solidity: function setGuardianFee(address tokenAddr, uint256 fee) returns()
func (_Dospayment *DospaymentTransactor) SetGuardianFee(opts *bind.TransactOpts, tokenAddr common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setGuardianFee", tokenAddr, fee)
}

// SetGuardianFee is a paid mutator transaction binding the contract method 0xeebede83.
//
// Solidity: function setGuardianFee(address tokenAddr, uint256 fee) returns()
func (_Dospayment *DospaymentSession) SetGuardianFee(tokenAddr common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetGuardianFee(&_Dospayment.TransactOpts, tokenAddr, fee)
}

// SetGuardianFee is a paid mutator transaction binding the contract method 0xeebede83.
//
// Solidity: function setGuardianFee(address tokenAddr, uint256 fee) returns()
func (_Dospayment *DospaymentTransactorSession) SetGuardianFee(tokenAddr common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetGuardianFee(&_Dospayment.TransactOpts, tokenAddr, fee)
}

// SetGuardianFunds is a paid mutator transaction binding the contract method 0x5a1fa503.
//
// Solidity: function setGuardianFunds(address fundsAddr, address tokenAddr) returns()
func (_Dospayment *DospaymentTransactor) SetGuardianFunds(opts *bind.TransactOpts, fundsAddr common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setGuardianFunds", fundsAddr, tokenAddr)
}

// SetGuardianFunds is a paid mutator transaction binding the contract method 0x5a1fa503.
//
// Solidity: function setGuardianFunds(address fundsAddr, address tokenAddr) returns()
func (_Dospayment *DospaymentSession) SetGuardianFunds(fundsAddr common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetGuardianFunds(&_Dospayment.TransactOpts, fundsAddr, tokenAddr)
}

// SetGuardianFunds is a paid mutator transaction binding the contract method 0x5a1fa503.
//
// Solidity: function setGuardianFunds(address fundsAddr, address tokenAddr) returns()
func (_Dospayment *DospaymentTransactorSession) SetGuardianFunds(fundsAddr common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetGuardianFunds(&_Dospayment.TransactOpts, fundsAddr, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(address consumer, address tokenAddr) returns()
func (_Dospayment *DospaymentTransactor) SetPaymentMethod(opts *bind.TransactOpts, consumer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setPaymentMethod", consumer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(address consumer, address tokenAddr) returns()
func (_Dospayment *DospaymentSession) SetPaymentMethod(consumer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetPaymentMethod(&_Dospayment.TransactOpts, consumer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(address consumer, address tokenAddr) returns()
func (_Dospayment *DospaymentTransactorSession) SetPaymentMethod(consumer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetPaymentMethod(&_Dospayment.TransactOpts, consumer, tokenAddr)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x3939c401.
//
// Solidity: function setServiceFee(address tokenAddr, uint256 serviceType, uint256 fee) returns()
func (_Dospayment *DospaymentTransactor) SetServiceFee(opts *bind.TransactOpts, tokenAddr common.Address, serviceType *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setServiceFee", tokenAddr, serviceType, fee)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x3939c401.
//
// Solidity: function setServiceFee(address tokenAddr, uint256 serviceType, uint256 fee) returns()
func (_Dospayment *DospaymentSession) SetServiceFee(tokenAddr common.Address, serviceType *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetServiceFee(&_Dospayment.TransactOpts, tokenAddr, serviceType, fee)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x3939c401.
//
// Solidity: function setServiceFee(address tokenAddr, uint256 serviceType, uint256 fee) returns()
func (_Dospayment *DospaymentTransactorSession) SetServiceFee(tokenAddr common.Address, serviceType *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetServiceFee(&_Dospayment.TransactOpts, tokenAddr, serviceType, fee)
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

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dospayment *DospaymentTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dospayment *DospaymentSession) Withdraw() (*types.Transaction, error) {
	return _Dospayment.Contract.Withdraw(&_Dospayment.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dospayment *DospaymentTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Dospayment.Contract.Withdraw(&_Dospayment.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 tokenAddr) returns()
func (_Dospayment *DospaymentTransactor) WithdrawAll(opts *bind.TransactOpts, tokenAddr *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "withdrawAll", tokenAddr)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 tokenAddr) returns()
func (_Dospayment *DospaymentSession) WithdrawAll(tokenAddr *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.WithdrawAll(&_Dospayment.TransactOpts, tokenAddr)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 tokenAddr) returns()
func (_Dospayment *DospaymentTransactorSession) WithdrawAll(tokenAddr *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.WithdrawAll(&_Dospayment.TransactOpts, tokenAddr)
}

// DospaymentLogChargeServiceFeeIterator is returned from FilterLogChargeServiceFee and is used to iterate over the raw logs and unpacked data for LogChargeServiceFee events raised by the Dospayment contract.
type DospaymentLogChargeServiceFeeIterator struct {
	Event *DospaymentLogChargeServiceFee // Event containing the contract specifics and raw log

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
func (it *DospaymentLogChargeServiceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentLogChargeServiceFee)
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
		it.Event = new(DospaymentLogChargeServiceFee)
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
func (it *DospaymentLogChargeServiceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentLogChargeServiceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentLogChargeServiceFee represents a LogChargeServiceFee event raised by the Dospayment contract.
type DospaymentLogChargeServiceFee struct {
	Consumer    common.Address
	TokenAddr   common.Address
	RequestID   *big.Int
	ServiceType *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogChargeServiceFee is a free log retrieval operation binding the contract event 0xa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba.
//
// Solidity: event LogChargeServiceFee(address consumer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
func (_Dospayment *DospaymentFilterer) FilterLogChargeServiceFee(opts *bind.FilterOpts) (*DospaymentLogChargeServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogChargeServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogChargeServiceFeeIterator{contract: _Dospayment.contract, event: "LogChargeServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogChargeServiceFee is a free log subscription operation binding the contract event 0xa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba.
//
// Solidity: event LogChargeServiceFee(address consumer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
func (_Dospayment *DospaymentFilterer) WatchLogChargeServiceFee(opts *bind.WatchOpts, sink chan<- *DospaymentLogChargeServiceFee) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "LogChargeServiceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentLogChargeServiceFee)
				if err := _Dospayment.contract.UnpackLog(event, "LogChargeServiceFee", log); err != nil {
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

// DospaymentLogClaimGuardianFeeIterator is returned from FilterLogClaimGuardianFee and is used to iterate over the raw logs and unpacked data for LogClaimGuardianFee events raised by the Dospayment contract.
type DospaymentLogClaimGuardianFeeIterator struct {
	Event *DospaymentLogClaimGuardianFee // Event containing the contract specifics and raw log

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
func (it *DospaymentLogClaimGuardianFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentLogClaimGuardianFee)
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
		it.Event = new(DospaymentLogClaimGuardianFee)
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
func (it *DospaymentLogClaimGuardianFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentLogClaimGuardianFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentLogClaimGuardianFee represents a LogClaimGuardianFee event raised by the Dospayment contract.
type DospaymentLogClaimGuardianFee struct {
	NodeAddr        common.Address
	TokenAddr       common.Address
	FeeForSubmitter *big.Int
	Sender          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogClaimGuardianFee is a free log retrieval operation binding the contract event 0x47ad88344c408450ef0ccab93ed97dd83af7a27dedfaa205c0725cfc4ca819cd.
//
// Solidity: event LogClaimGuardianFee(address nodeAddr, address tokenAddr, uint256 feeForSubmitter, address sender)
func (_Dospayment *DospaymentFilterer) FilterLogClaimGuardianFee(opts *bind.FilterOpts) (*DospaymentLogClaimGuardianFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogClaimGuardianFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogClaimGuardianFeeIterator{contract: _Dospayment.contract, event: "LogClaimGuardianFee", logs: logs, sub: sub}, nil
}

// WatchLogClaimGuardianFee is a free log subscription operation binding the contract event 0x47ad88344c408450ef0ccab93ed97dd83af7a27dedfaa205c0725cfc4ca819cd.
//
// Solidity: event LogClaimGuardianFee(address nodeAddr, address tokenAddr, uint256 feeForSubmitter, address sender)
func (_Dospayment *DospaymentFilterer) WatchLogClaimGuardianFee(opts *bind.WatchOpts, sink chan<- *DospaymentLogClaimGuardianFee) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "LogClaimGuardianFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentLogClaimGuardianFee)
				if err := _Dospayment.contract.UnpackLog(event, "LogClaimGuardianFee", log); err != nil {
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

// DospaymentLogClaimServiceFeeIterator is returned from FilterLogClaimServiceFee and is used to iterate over the raw logs and unpacked data for LogClaimServiceFee events raised by the Dospayment contract.
type DospaymentLogClaimServiceFeeIterator struct {
	Event *DospaymentLogClaimServiceFee // Event containing the contract specifics and raw log

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
func (it *DospaymentLogClaimServiceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentLogClaimServiceFee)
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
		it.Event = new(DospaymentLogClaimServiceFee)
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
func (it *DospaymentLogClaimServiceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentLogClaimServiceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentLogClaimServiceFee represents a LogClaimServiceFee event raised by the Dospayment contract.
type DospaymentLogClaimServiceFee struct {
	NodeAddr        common.Address
	TokenAddr       common.Address
	RequestID       *big.Int
	ServiceType     *big.Int
	FeeForSubmitter *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogClaimServiceFee is a free log retrieval operation binding the contract event 0xab9fe896064c2c9dfd31acebc7d522b311e5f2e7d1b4965ac0cfd5a4abec813a.
//
// Solidity: event LogClaimServiceFee(address nodeAddr, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 feeForSubmitter)
func (_Dospayment *DospaymentFilterer) FilterLogClaimServiceFee(opts *bind.FilterOpts) (*DospaymentLogClaimServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogClaimServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogClaimServiceFeeIterator{contract: _Dospayment.contract, event: "LogClaimServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogClaimServiceFee is a free log subscription operation binding the contract event 0xab9fe896064c2c9dfd31acebc7d522b311e5f2e7d1b4965ac0cfd5a4abec813a.
//
// Solidity: event LogClaimServiceFee(address nodeAddr, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 feeForSubmitter)
func (_Dospayment *DospaymentFilterer) WatchLogClaimServiceFee(opts *bind.WatchOpts, sink chan<- *DospaymentLogClaimServiceFee) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "LogClaimServiceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentLogClaimServiceFee)
				if err := _Dospayment.contract.UnpackLog(event, "LogClaimServiceFee", log); err != nil {
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

// DospaymentLogRefundServiceFeeIterator is returned from FilterLogRefundServiceFee and is used to iterate over the raw logs and unpacked data for LogRefundServiceFee events raised by the Dospayment contract.
type DospaymentLogRefundServiceFeeIterator struct {
	Event *DospaymentLogRefundServiceFee // Event containing the contract specifics and raw log

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
func (it *DospaymentLogRefundServiceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentLogRefundServiceFee)
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
		it.Event = new(DospaymentLogRefundServiceFee)
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
func (it *DospaymentLogRefundServiceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentLogRefundServiceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentLogRefundServiceFee represents a LogRefundServiceFee event raised by the Dospayment contract.
type DospaymentLogRefundServiceFee struct {
	Consumer    common.Address
	TokenAddr   common.Address
	RequestID   *big.Int
	ServiceType *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogRefundServiceFee is a free log retrieval operation binding the contract event 0xde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa.
//
// Solidity: event LogRefundServiceFee(address consumer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
func (_Dospayment *DospaymentFilterer) FilterLogRefundServiceFee(opts *bind.FilterOpts) (*DospaymentLogRefundServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogRefundServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogRefundServiceFeeIterator{contract: _Dospayment.contract, event: "LogRefundServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogRefundServiceFee is a free log subscription operation binding the contract event 0xde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa.
//
// Solidity: event LogRefundServiceFee(address consumer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
func (_Dospayment *DospaymentFilterer) WatchLogRefundServiceFee(opts *bind.WatchOpts, sink chan<- *DospaymentLogRefundServiceFee) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "LogRefundServiceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentLogRefundServiceFee)
				if err := _Dospayment.contract.UnpackLog(event, "LogRefundServiceFee", log); err != nil {
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
