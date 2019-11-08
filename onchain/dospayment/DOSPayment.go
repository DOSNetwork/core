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

// DospaymentABI is the input ABI used to generate the binding from.
const DospaymentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"_defaultGuardianFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"nodeTokenAddres\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setNetworkToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"refundServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"_paymentMethods\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guardianAddr\",\"type\":\"address\"}],\"name\":\"claimGuardianReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"paymentInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"serviceType\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeTokenAddresLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"quo\",\"type\":\"uint256\"}],\"name\":\"setDropBurnMaxQuota\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fundsAddr\",\"type\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setGuardianFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_guardianFundsTokenAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultTokenAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultSubmitterRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"requestID\",\"type\":\"uint256\"},{\"name\":\"serviceType\",\"type\":\"uint256\"}],\"name\":\"chargeServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"uint256\"}],\"name\":\"withdrawAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_payments\",\"outputs\":[{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"serviceType\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultWorkerRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dropburnToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"consumer\",\"type\":\"address\"},{\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setPaymentMethod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"_feeList\",\"outputs\":[{\"name\":\"submitterRate\",\"type\":\"uint256\"},{\"name\":\"workerRate\",\"type\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\"},{\"name\":\"guardianFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"address\"}],\"name\":\"fromValidStakingNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultUserQueryFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultSystemRandomFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"submitterRate\",\"type\":\"uint256\"},{\"name\":\"workerRate\",\"type\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"setFeeDividend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_guardianFundsAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dropburnMaxQuota\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDropBurnToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setGuardianFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"nodeTokenAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_defaultUserRandomFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"requestID\",\"type\":\"uint256\"},{\"name\":\"submitter\",\"type\":\"address\"},{\"name\":\"workers\",\"type\":\"address[]\"}],\"name\":\"claimServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateNetworkTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"UpdateDropBurnTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"UpdateDropBurnMaxQuota\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogChargeServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"consumer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogRefundServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"feeForSubmitter\",\"type\":\"uint256\"}],\"name\":\"LogClaimServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"feeForSubmitter\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LogClaimGuardianFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// DospaymentBin is the compiled bytecode used for deploying new contracts.
const DospaymentBin = `6080604052600580546001600160a01b031990811673214e79c85744cd2ebbc64ddc0047131496871bee908117808455600693909355600360078190556002600855674563918244f400006009819055600a819055600b55670de0b6b3a7640000600c55600d80548416732a3b59ac638f90d82bdaf5e2da5d37c1a31b29f3179055600e805484166001600160a01b039590951694909417909355600f80548316738811ccb0d8128eee3bd48553a69f62205da1dea2179055601080548316909117905560118054909116739bfe8f5749d90eb4049ad94cc4de9b6c4c31f82217905561c3506012556013553480156100f757600080fd5b50600080546001600160a01b031916331781556005546001600160a01b031681526002602081815260408084206009548580529281905281852092909255600a54600180865282862091909155600b548486529185209190915560075490820155600854918101919091556006546003820155600c546004909101556121f190819061018390396000f3fe608060405234801561001057600080fd5b50600436106102535760003560e01c80638da5cb5b11610146578063d5fd1f0f116100c3578063eac051f911610087578063eac051f91461065e578063eebede8314610684578063f2fde38b146106b0578063f87059c1146106d6578063fe7b663514610702578063fedb8b6a1461070a57610253565b8063d5fd1f0f14610606578063de439e9a1461060e578063e2edf58c14610646578063e65449041461064e578063e8c3470c1461065657610253565b8063b26584b81161010a578063b26584b814610556578063b73a3f8f1461055e578063c5e7ffef1461058c578063c7e6a9bc146105d8578063d40f68d1146105fe57610253565b80638da5cb5b146104d25780638f32d59b146104da578063958e2d31146104e25780639ed713b4146104ff578063a387722b1461054e57610253565b80633ccfd60b116101d4578063635c971d11610198578063635c971d1461048057806365e46041146104885780636ca95a4e14610490578063715018a6146104985780637aa9181b146104a057610253565b80633ccfd60b146103ff5780633e698ad5146104075780633f3381e11461042d5780635a1fa5031461044a578063610c3b961461047857610253565b806323ff34cb1161021b57806323ff34cb14610325578063240028e81461034b5780633157f16d14610385578063375b3c0a146103c55780633939c401146103cd57610253565b806311bbe27614610258578063139bcedb1461027257806317107c49146102ba5780631efa5a98146102e257806322e60ea5146102ff575b600080fd5b6102606107c2565b60408051918252519081900360200190f35b61029e6004803603604081101561028857600080fd5b506001600160a01b0381351690602001356107c8565b604080516001600160a01b039092168252519081900360200190f35b6102e0600480360360208110156102d057600080fd5b50356001600160a01b0316610874565b005b6102e0600480360360208110156102f857600080fd5b50356108ef565b61029e6004803603602081101561031557600080fd5b50356001600160a01b0316610b6f565b6102e06004803603602081101561033b57600080fd5b50356001600160a01b0316610b8a565b6103716004803603602081101561036157600080fd5b50356001600160a01b0316610e06565b604080519115158252519081900360200190f35b6103a26004803603602081101561039b57600080fd5b5035610eae565b604080516001600160a01b03909316835260208301919091528051918290030190f35b610260610ed5565b6102e0600480360360608110156103e357600080fd5b506001600160a01b038135169060208101359060400135610edb565b6102e0610f6a565b6102606004803603602081101561041d57600080fd5b50356001600160a01b03166110d7565b6102e06004803603602081101561044357600080fd5b50356110f5565b6102e06004803603604081101561046057600080fd5b506001600160a01b0381358116916020013516611198565b61029e611236565b61029e611245565b610260611254565b61029e61125a565b6102e0611269565b6102e0600480360360608110156104b657600080fd5b506001600160a01b0381351690602081013590604001356112c2565b61029e611518565b610371611527565b6102e0600480360360208110156104f857600080fd5b5035611538565b61051c6004803603602081101561051557600080fd5b5035611618565b604080516001600160a01b03958616815293909416602084015282840191909152606082015290519081900360800190f35b61026061164e565b61029e611654565b6102e06004803603604081101561057457600080fd5b506001600160a01b0381358116916020013516611663565b6105b2600480360360208110156105a257600080fd5b50356001600160a01b03166116f0565b604080519485526020850193909352838301919091526060830152519081900360800190f35b610371600480360360208110156105ee57600080fd5b50356001600160a01b031661171b565b610260611974565b61026061197a565b6102e06004803603608081101561062457600080fd5b506001600160a01b038135169060208101359060408101359060600135611980565b61029e611a16565b610260611a25565b610260611a2b565b6102e06004803603602081101561067457600080fd5b50356001600160a01b0316611a31565b6102e06004803603604081101561069a57600080fd5b506001600160a01b038135169060200135611aac565b6102e0600480360360208110156106c657600080fd5b50356001600160a01b0316611b34565b610260600480360360408110156106ec57600080fd5b506001600160a01b038135169060200135611b51565b610260611c1a565b6102e06004803603606081101561072057600080fd5b8135916001600160a01b036020820135169181019060608101604082013564010000000081111561075057600080fd5b82018360208201111561076257600080fd5b8035906020019184602083028401116401000000008311171561078457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611c20945050505050565b600c5481565b6001600160a01b03821660009081526004602052604081206001015482106108315760408051600160e51b62461bcd0281526020600482015260116024820152600160781b704e6f20746f6b656e20616464726573732102604482015290519081900360640190fd5b6001600160a01b038316600090815260046020526040902060010180548390811061085857fe5b6000918252602090912001546001600160a01b03169392505050565b61087c611527565b61088557600080fd5b601054604080516001600160a01b039283168152918316602083015280517f4d27a2adceae86b92fb74fb7e8f96dc902d917e243fbff389b5a793c9040dafe9281900390910190a1601080546001600160a01b0319166001600160a01b0392909216919091179055565b6108f7611527565b61090057600080fd5b6000818152600360208190526040909120015481906109615760408051600160e51b62461bcd0281526020600482015260126024820152600160701b714e6f2066656520696e666f6d6174696f6e2102604482015290519081900360640190fd5b6000818152600360205260409020546001600160a01b03166109cd5760408051600160e51b62461bcd02815260206004820152601760248201527f4e6f20636f6e73756d657220696e666f6d6174696f6e21000000000000000000604482015290519081900360640190fd5b6000818152600360205260409020600101546001600160a01b0316610a3c5760408051600160e51b62461bcd02815260206004820152601860248201527f4e6f20746f6b656e4164647220696e666f6d6174696f6e210000000000000000604482015290519081900360640190fd5b6000828152600360208181526040808420928301805460028501805460018701805488546001600160a01b0319808216909a55988116909155918890559690925582516001600160a01b039586168082529590921693820184905281830188905260608201869052608082018190529151919493917fde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa9181900360a00190a1816001600160a01b031663a9059cbb82866040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015610b3b57600080fd5b505af1158015610b4f573d6000803e3d6000fd5b505050506040513d6020811015610b6557600080fd5b5050505050505050565b6001602052600090815260409020546001600160a01b031681565b600f60009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610bd857600080fd5b505afa158015610bec573d6000803e3d6000fd5b505050506040513d6020811015610c0257600080fd5b50516001600160a01b03163314610c5c5760408051600160e51b62461bcd0281526020600482015260136024820152600160681b724e6f742066726f6d20444f532070726f78792102604482015290519081900360640190fd5b600d546001600160a01b0316610ca657604051600160e51b62461bcd02815260040180806020018281038252602381526020018061217f6023913960400191505060405180910390fd5b600e546001600160a01b0316610d065760408051600160e51b62461bcd02815260206004820152601a60248201527f4e6f7420612076616c696420746f6b656e206164647265737321000000000000604482015290519081900360640190fd5b600e546001600160a01b039081166000818152600260209081526040918290206004015482519486168552908401929092528281018290523360608401525190917f47ad88344c408450ef0ccab93ed97dd83af7a27dedfaa205c0725cfc4ca819cd919081900360800190a1600e54600d5460408051600160e01b6323b872dd0281526001600160a01b039283166004820152858316602482015260448101859052905191909216916323b872dd9160648083019260209291908290030181600087803b158015610dd657600080fd5b505af1158015610dea573d6000803e3d6000fd5b505050506040513d6020811015610e0057600080fd5b50505050565b60006001600160a01b0382161580610e3f57506001600160a01b0382166000908152600260209081526040808320838052909152902054155b80610e6c57506001600160a01b038216600090815260026020908152604080832060018452909152902054155b80610e9857506001600160a01b0382166000908152600260208181526040808420928452919052902054155b15610ea557506000610ea9565b5060015b919050565b600090815260036020819052604090912060018101549101546001600160a01b0390911691565b60125481565b610ee3611527565b610eec57600080fd5b6001600160a01b038316610f445760408051600160e51b62461bcd0281526020600482015260146024820152600160601b734e6f7420612076616c696420616464726573732102604482015290519081900360640190fd5b6001600160a01b0390921660009081526002602090815260408083209383529290522055565b33600090815260046020526040902060010154610fc25760408051600160e51b62461bcd02815260206004820152600b6024820152600160a81b6a4e6f20726577617264732102604482015290519081900360640190fd5b3360009081526004602052604081206001810180549192916000198101908110610fe857fe5b6000918252602090912001546001830180546001600160a01b039092169250908061100f57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038316825283905260408120805491905580156110d25760408051600160e01b63a9059cbb0281523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b1580156110a557600080fd5b505af11580156110b9573d6000803e3d6000fd5b505050506040513d60208110156110cf57600080fd5b50505b505050565b6001600160a01b031660009081526004602052604090206001015490565b6110fd611527565b61110657600080fd5b60135481141580156111185750600a81105b61115657604051600160e51b62461bcd0281526004018080602001828103825260248152602001806121a26024913960400191505060405180910390fd5b601354604080519182526020820183905280517f0aee95cca46da64ee373e28dee5994361b4002c54035d92932c9825b76382e999281900390910190a1601355565b6111a0611527565b6111a957600080fd5b806111b381610e06565b6112075760408051600160e51b62461bcd02815260206004820152601e60248201527f4e6f74206120737570706f7274656420746f6b656e2061646472657373210000604482015290519081900360640190fd5b50600d80546001600160a01b039384166001600160a01b031991821617909155600e8054929093169116179055565b600e546001600160a01b031681565b6005546001600160a01b031681565b60075481565b6010546001600160a01b031681565b611271611527565b61127a57600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b600f60009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561131057600080fd5b505afa158015611324573d6000803e3d6000fd5b505050506040513d602081101561133a57600080fd5b50516001600160a01b031633146113945760408051600160e51b62461bcd0281526020600482015260136024820152600160681b724e6f742066726f6d20444f532070726f78792102604482015290519081900360640190fd5b6001600160a01b03808416600090815260016020526040902054166113b881610e06565b61140c5760408051600160e51b62461bcd02815260206004820152601960248201527f4e6f7420612076616c696420746f6b656e206164647265737300000000000000604482015290519081900360640190fd5b6001600160a01b0380821660008181526002602081815260408084208885528252808420548985526003808452948290208054978c166001600160a01b0319988916811782559481018a90556001810180549098168717909755938601849055805192835290820193909352808301879052606081018690526080810182905291519092917fa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba919081900360a00190a160408051600160e01b6323b872dd0281526001600160a01b038881166004830152306024830152604482018590529151918516916323b872dd916064808201926020929091908290030181600087803b158015610b3b57600080fd5b6000546001600160a01b031690565b6000546001600160a01b0316331490565b611540611527565b61154957600080fd5b60408051600160e01b6370a0823102815230600482015290516000916001600160a01b038416916370a0823191602480820192602092909190829003018186803b15801561159657600080fd5b505afa1580156115aa573d6000803e3d6000fd5b505050506040513d60208110156115c057600080fd5b505160408051600160e01b63a9059cbb0281523360048201526024810183905290519192506001600160a01b0384169163a9059cbb916044808201926020929091908290030181600087803b158015610dd657600080fd5b600360208190526000918252604090912080546001820154600283015492909301546001600160a01b0391821693909116919084565b60085481565b6011546001600160a01b031681565b8061166d81610e06565b6116c15760408051600160e51b62461bcd02815260206004820152601e60248201527f4e6f74206120737570706f7274656420746f6b656e2061646472657373210000604482015290519081900360640190fd5b506001600160a01b03918216600090815260016020526040902080546001600160a01b03191691909216179055565b6002602081905260009182526040909120600181015491810154600382015460049092015490919084565b60105460408051600160e01b6370a082310281526001600160a01b0384811660048301529151600093849316916370a08231916024808301926020929190829003018186803b15801561176d57600080fd5b505afa158015611781573d6000803e3d6000fd5b505050506040513d602081101561179757600080fd5b505160105460408051600160e01b63313ce56702815290519293506000926001600160a01b039092169163313ce56791600480820192602092909190829003018186803b1580156117e757600080fd5b505afa1580156117fb573d6000803e3d6000fd5b505050506040513d602081101561181157600080fd5b5051601254909150600a82900a028083106118325760019350505050610ea9565b6011546001600160a01b031661184e5760009350505050610ea9565b60115460408051600160e01b63313ce56702815290516000926001600160a01b03169163313ce567916004808301926020929190829003018186803b15801561189657600080fd5b505afa1580156118aa573d6000803e3d6000fd5b505050506040513d60208110156118c057600080fd5b505160115460408051600160e01b6370a082310281526001600160a01b038a811660048301529151600a9490940a9391909216916370a08231916024808301926020929190829003018186803b15801561191957600080fd5b505afa15801561192d573d6000803e3d6000fd5b505050506040513d602081101561194357600080fd5b50518161194c57fe5b04905060135481111561195e57506013545b600a818103830204841015945050505050610ea9565b600b5481565b60095481565b611988611527565b61199157600080fd5b6001600160a01b0384166119e95760408051600160e51b62461bcd0281526020600482015260146024820152600160601b734e6f7420612076616c696420616464726573732102604482015290519081900360640190fd5b6001600160a01b039093166000908152600260208190526040909120600181019390935582015560030155565b600d546001600160a01b031681565b60065481565b60135481565b611a39611527565b611a4257600080fd5b601154604080516001600160a01b039283168152918316602083015280517ffc8013dfb0c8d38f3bcab9239bd5712457c48919b272cdb109488549199a01739281900390910190a1601180546001600160a01b0319166001600160a01b0392909216919091179055565b611ab4611527565b611abd57600080fd5b6001600160a01b038216611b155760408051600160e51b62461bcd0281526020600482015260146024820152600160601b734e6f7420612076616c696420616464726573732102604482015290519081900360640190fd5b6001600160a01b03909116600090815260026020526040902060040155565b611b3c611527565b611b4557600080fd5b611b4e81612079565b50565b6001600160a01b0382166000908152600460205260408120600101548210611bba5760408051600160e51b62461bcd0281526020600482015260116024820152600160781b704e6f20746f6b656e20616464726573732102604482015290519081900360640190fd5b6001600160a01b0383166000908152600460205260408120600101805484908110611be157fe5b60009182526020808320909101546001600160a01b03878116845260048352604080852091909216845290915290205491505092915050565b600a5481565b600f60009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c6e57600080fd5b505afa158015611c82573d6000803e3d6000fd5b505050506040513d6020811015611c9857600080fd5b50516001600160a01b03163314611cf25760408051600160e51b62461bcd0281526020600482015260136024820152600160681b724e6f742066726f6d20444f532070726f78792102604482015290519081900360640190fd5b600083815260036020819052604090912001548390611d535760408051600160e51b62461bcd0281526020600482015260126024820152600160701b714e6f2066656520696e666f6d6174696f6e2102604482015290519081900360640190fd5b6000818152600360205260409020546001600160a01b0316611dbf5760408051600160e51b62461bcd02815260206004820152601760248201527f4e6f20636f6e73756d657220696e666f6d6174696f6e21000000000000000000604482015290519081900360640190fd5b6000818152600360205260409020600101546001600160a01b0316611e2e5760408051600160e51b62461bcd02815260206004820152601860248201527f4e6f20746f6b656e4164647220696e666f6d6174696f6e210000000000000000604482015290519081900360640190fd5b60008481526003602081905260408220600181018054928201805460028401805485546001600160a01b031990811690965594861690935591859055939093556001600160a01b039091169190611e83612156565b506001600160a01b03831660009081526002602081815260409283902083516080810185526001820154815292810154918301919091526003810154928201839052600401546060820152908381611ed757fe5b60208301518851929091049450840290600090600019018281611ef657fe5b6001600160a01b038b1660009081526004602052604090209190049150611f1e9087846120e7565b604080516001600160a01b03808c168252881660208201528082018c9052606081018690526080810184905290517fab9fe896064c2c9dfd31acebc7d522b311e5f2e7d1b4965ac0cfd5a4abec813a9181900360a00190a160005b885181101561206c57896001600160a01b0316898281518110611f9857fe5b60200260200101516001600160a01b03161461206457611ff0600460008b8481518110611fc157fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002088846120e7565b7fab9fe896064c2c9dfd31acebc7d522b311e5f2e7d1b4965ac0cfd5a4abec813a89828151811061201d57fe5b602090810291909101810151604080516001600160a01b039283168152918b16928201929092528082018e9052606081018890526080810186905290519081900360a00190a15b600101611f79565b5050505050505050505050565b6001600160a01b03811661208c57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03821660009081526020849052604090205480612135576001848101805491820181556000908152602090200180546001600160a01b0319166001600160a01b0385161790555b6001600160a01b039092166000908152602093909352604090922091019055565b604051806080016040528060008152602001600081526020016000815260200160008152509056fe4e6f7420612076616c696420677561726469616e2066756e647320616464726573732156616c69642064726f706275726e4d617851756f74612077697468696e203020746f2039a165627a7a7230582050fdfa82be2368132fdb7d4c1c34cbbe90f367f502c75582cf4856e6e95642380029`

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
// Solidity: function _feeList( address) constant returns(submitterRate uint256, workerRate uint256, denominator uint256, guardianFee uint256)
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
// Solidity: function _feeList( address) constant returns(submitterRate uint256, workerRate uint256, denominator uint256, guardianFee uint256)
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
// Solidity: function _feeList( address) constant returns(submitterRate uint256, workerRate uint256, denominator uint256, guardianFee uint256)
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
// Solidity: function _paymentMethods( address) constant returns(address)
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
// Solidity: function _paymentMethods( address) constant returns(address)
func (_Dospayment *DospaymentSession) PaymentMethods(arg0 common.Address) (common.Address, error) {
	return _Dospayment.Contract.PaymentMethods(&_Dospayment.CallOpts, arg0)
}

// PaymentMethods is a free data retrieval call binding the contract method 0x22e60ea5.
//
// Solidity: function _paymentMethods( address) constant returns(address)
func (_Dospayment *DospaymentCallerSession) PaymentMethods(arg0 common.Address) (common.Address, error) {
	return _Dospayment.Contract.PaymentMethods(&_Dospayment.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x9ed713b4.
//
// Solidity: function _payments( uint256) constant returns(consumer address, tokenAddr address, serviceType uint256, amount uint256)
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
// Solidity: function _payments( uint256) constant returns(consumer address, tokenAddr address, serviceType uint256, amount uint256)
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
// Solidity: function _payments( uint256) constant returns(consumer address, tokenAddr address, serviceType uint256, amount uint256)
func (_Dospayment *DospaymentCallerSession) Payments(arg0 *big.Int) (struct {
	Consumer    common.Address
	TokenAddr   common.Address
	ServiceType *big.Int
	Amount      *big.Int
}, error) {
	return _Dospayment.Contract.Payments(&_Dospayment.CallOpts, arg0)
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
// Solidity: function fromValidStakingNode(node address) constant returns(bool)
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
// Solidity: function fromValidStakingNode(node address) constant returns(bool)
func (_Dospayment *DospaymentSession) FromValidStakingNode(node common.Address) (bool, error) {
	return _Dospayment.Contract.FromValidStakingNode(&_Dospayment.CallOpts, node)
}

// FromValidStakingNode is a free data retrieval call binding the contract method 0xc7e6a9bc.
//
// Solidity: function fromValidStakingNode(node address) constant returns(bool)
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

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(tokenAddr address) constant returns(bool)
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
// Solidity: function isSupportedToken(tokenAddr address) constant returns(bool)
func (_Dospayment *DospaymentSession) IsSupportedToken(tokenAddr common.Address) (bool, error) {
	return _Dospayment.Contract.IsSupportedToken(&_Dospayment.CallOpts, tokenAddr)
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(tokenAddr address) constant returns(bool)
func (_Dospayment *DospaymentCallerSession) IsSupportedToken(tokenAddr common.Address) (bool, error) {
	return _Dospayment.Contract.IsSupportedToken(&_Dospayment.CallOpts, tokenAddr)
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

// NodeTokenAddres is a free data retrieval call binding the contract method 0x139bcedb.
//
// Solidity: function nodeTokenAddres(nodeAddr address, idx uint256) constant returns(address)
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
// Solidity: function nodeTokenAddres(nodeAddr address, idx uint256) constant returns(address)
func (_Dospayment *DospaymentSession) NodeTokenAddres(nodeAddr common.Address, idx *big.Int) (common.Address, error) {
	return _Dospayment.Contract.NodeTokenAddres(&_Dospayment.CallOpts, nodeAddr, idx)
}

// NodeTokenAddres is a free data retrieval call binding the contract method 0x139bcedb.
//
// Solidity: function nodeTokenAddres(nodeAddr address, idx uint256) constant returns(address)
func (_Dospayment *DospaymentCallerSession) NodeTokenAddres(nodeAddr common.Address, idx *big.Int) (common.Address, error) {
	return _Dospayment.Contract.NodeTokenAddres(&_Dospayment.CallOpts, nodeAddr, idx)
}

// NodeTokenAddresLength is a free data retrieval call binding the contract method 0x3e698ad5.
//
// Solidity: function nodeTokenAddresLength(nodeAddr address) constant returns(uint256)
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
// Solidity: function nodeTokenAddresLength(nodeAddr address) constant returns(uint256)
func (_Dospayment *DospaymentSession) NodeTokenAddresLength(nodeAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeTokenAddresLength(&_Dospayment.CallOpts, nodeAddr)
}

// NodeTokenAddresLength is a free data retrieval call binding the contract method 0x3e698ad5.
//
// Solidity: function nodeTokenAddresLength(nodeAddr address) constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) NodeTokenAddresLength(nodeAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeTokenAddresLength(&_Dospayment.CallOpts, nodeAddr)
}

// NodeTokenAmount is a free data retrieval call binding the contract method 0xf87059c1.
//
// Solidity: function nodeTokenAmount(nodeAddr address, idx uint256) constant returns(uint256)
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
// Solidity: function nodeTokenAmount(nodeAddr address, idx uint256) constant returns(uint256)
func (_Dospayment *DospaymentSession) NodeTokenAmount(nodeAddr common.Address, idx *big.Int) (*big.Int, error) {
	return _Dospayment.Contract.NodeTokenAmount(&_Dospayment.CallOpts, nodeAddr, idx)
}

// NodeTokenAmount is a free data retrieval call binding the contract method 0xf87059c1.
//
// Solidity: function nodeTokenAmount(nodeAddr address, idx uint256) constant returns(uint256)
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
// Solidity: function paymentInfo(requestID uint256) constant returns(address, uint256)
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
// Solidity: function paymentInfo(requestID uint256) constant returns(address, uint256)
func (_Dospayment *DospaymentSession) PaymentInfo(requestID *big.Int) (common.Address, *big.Int, error) {
	return _Dospayment.Contract.PaymentInfo(&_Dospayment.CallOpts, requestID)
}

// PaymentInfo is a free data retrieval call binding the contract method 0x3157f16d.
//
// Solidity: function paymentInfo(requestID uint256) constant returns(address, uint256)
func (_Dospayment *DospaymentCallerSession) PaymentInfo(requestID *big.Int) (common.Address, *big.Int, error) {
	return _Dospayment.Contract.PaymentInfo(&_Dospayment.CallOpts, requestID)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(consumer address, requestID uint256, serviceType uint256) returns()
func (_Dospayment *DospaymentTransactor) ChargeServiceFee(opts *bind.TransactOpts, consumer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "chargeServiceFee", consumer, requestID, serviceType)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(consumer address, requestID uint256, serviceType uint256) returns()
func (_Dospayment *DospaymentSession) ChargeServiceFee(consumer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.ChargeServiceFee(&_Dospayment.TransactOpts, consumer, requestID, serviceType)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(consumer address, requestID uint256, serviceType uint256) returns()
func (_Dospayment *DospaymentTransactorSession) ChargeServiceFee(consumer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.ChargeServiceFee(&_Dospayment.TransactOpts, consumer, requestID, serviceType)
}

// ClaimGuardianReward is a paid mutator transaction binding the contract method 0x23ff34cb.
//
// Solidity: function claimGuardianReward(guardianAddr address) returns()
func (_Dospayment *DospaymentTransactor) ClaimGuardianReward(opts *bind.TransactOpts, guardianAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "claimGuardianReward", guardianAddr)
}

// ClaimGuardianReward is a paid mutator transaction binding the contract method 0x23ff34cb.
//
// Solidity: function claimGuardianReward(guardianAddr address) returns()
func (_Dospayment *DospaymentSession) ClaimGuardianReward(guardianAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.ClaimGuardianReward(&_Dospayment.TransactOpts, guardianAddr)
}

// ClaimGuardianReward is a paid mutator transaction binding the contract method 0x23ff34cb.
//
// Solidity: function claimGuardianReward(guardianAddr address) returns()
func (_Dospayment *DospaymentTransactorSession) ClaimGuardianReward(guardianAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.ClaimGuardianReward(&_Dospayment.TransactOpts, guardianAddr)
}

// ClaimServiceFee is a paid mutator transaction binding the contract method 0xfedb8b6a.
//
// Solidity: function claimServiceFee(requestID uint256, submitter address, workers address[]) returns()
func (_Dospayment *DospaymentTransactor) ClaimServiceFee(opts *bind.TransactOpts, requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "claimServiceFee", requestID, submitter, workers)
}

// ClaimServiceFee is a paid mutator transaction binding the contract method 0xfedb8b6a.
//
// Solidity: function claimServiceFee(requestID uint256, submitter address, workers address[]) returns()
func (_Dospayment *DospaymentSession) ClaimServiceFee(requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.ClaimServiceFee(&_Dospayment.TransactOpts, requestID, submitter, workers)
}

// ClaimServiceFee is a paid mutator transaction binding the contract method 0xfedb8b6a.
//
// Solidity: function claimServiceFee(requestID uint256, submitter address, workers address[]) returns()
func (_Dospayment *DospaymentTransactorSession) ClaimServiceFee(requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.ClaimServiceFee(&_Dospayment.TransactOpts, requestID, submitter, workers)
}

// RefundServiceFee is a paid mutator transaction binding the contract method 0x1efa5a98.
//
// Solidity: function refundServiceFee(requestID uint256) returns()
func (_Dospayment *DospaymentTransactor) RefundServiceFee(opts *bind.TransactOpts, requestID *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "refundServiceFee", requestID)
}

// RefundServiceFee is a paid mutator transaction binding the contract method 0x1efa5a98.
//
// Solidity: function refundServiceFee(requestID uint256) returns()
func (_Dospayment *DospaymentSession) RefundServiceFee(requestID *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.RefundServiceFee(&_Dospayment.TransactOpts, requestID)
}

// RefundServiceFee is a paid mutator transaction binding the contract method 0x1efa5a98.
//
// Solidity: function refundServiceFee(requestID uint256) returns()
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

// SetDropBurnMaxQuota is a paid mutator transaction binding the contract method 0x3f3381e1.
//
// Solidity: function setDropBurnMaxQuota(quo uint256) returns()
func (_Dospayment *DospaymentTransactor) SetDropBurnMaxQuota(opts *bind.TransactOpts, quo *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setDropBurnMaxQuota", quo)
}

// SetDropBurnMaxQuota is a paid mutator transaction binding the contract method 0x3f3381e1.
//
// Solidity: function setDropBurnMaxQuota(quo uint256) returns()
func (_Dospayment *DospaymentSession) SetDropBurnMaxQuota(quo *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDropBurnMaxQuota(&_Dospayment.TransactOpts, quo)
}

// SetDropBurnMaxQuota is a paid mutator transaction binding the contract method 0x3f3381e1.
//
// Solidity: function setDropBurnMaxQuota(quo uint256) returns()
func (_Dospayment *DospaymentTransactorSession) SetDropBurnMaxQuota(quo *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDropBurnMaxQuota(&_Dospayment.TransactOpts, quo)
}

// SetDropBurnToken is a paid mutator transaction binding the contract method 0xeac051f9.
//
// Solidity: function setDropBurnToken(addr address) returns()
func (_Dospayment *DospaymentTransactor) SetDropBurnToken(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setDropBurnToken", addr)
}

// SetDropBurnToken is a paid mutator transaction binding the contract method 0xeac051f9.
//
// Solidity: function setDropBurnToken(addr address) returns()
func (_Dospayment *DospaymentSession) SetDropBurnToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDropBurnToken(&_Dospayment.TransactOpts, addr)
}

// SetDropBurnToken is a paid mutator transaction binding the contract method 0xeac051f9.
//
// Solidity: function setDropBurnToken(addr address) returns()
func (_Dospayment *DospaymentTransactorSession) SetDropBurnToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetDropBurnToken(&_Dospayment.TransactOpts, addr)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0xde439e9a.
//
// Solidity: function setFeeDividend(tokenAddr address, submitterRate uint256, workerRate uint256, denominator uint256) returns()
func (_Dospayment *DospaymentTransactor) SetFeeDividend(opts *bind.TransactOpts, tokenAddr common.Address, submitterRate *big.Int, workerRate *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setFeeDividend", tokenAddr, submitterRate, workerRate, denominator)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0xde439e9a.
//
// Solidity: function setFeeDividend(tokenAddr address, submitterRate uint256, workerRate uint256, denominator uint256) returns()
func (_Dospayment *DospaymentSession) SetFeeDividend(tokenAddr common.Address, submitterRate *big.Int, workerRate *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetFeeDividend(&_Dospayment.TransactOpts, tokenAddr, submitterRate, workerRate, denominator)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0xde439e9a.
//
// Solidity: function setFeeDividend(tokenAddr address, submitterRate uint256, workerRate uint256, denominator uint256) returns()
func (_Dospayment *DospaymentTransactorSession) SetFeeDividend(tokenAddr common.Address, submitterRate *big.Int, workerRate *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetFeeDividend(&_Dospayment.TransactOpts, tokenAddr, submitterRate, workerRate, denominator)
}

// SetGuardianFee is a paid mutator transaction binding the contract method 0xeebede83.
//
// Solidity: function setGuardianFee(tokenAddr address, fee uint256) returns()
func (_Dospayment *DospaymentTransactor) SetGuardianFee(opts *bind.TransactOpts, tokenAddr common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setGuardianFee", tokenAddr, fee)
}

// SetGuardianFee is a paid mutator transaction binding the contract method 0xeebede83.
//
// Solidity: function setGuardianFee(tokenAddr address, fee uint256) returns()
func (_Dospayment *DospaymentSession) SetGuardianFee(tokenAddr common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetGuardianFee(&_Dospayment.TransactOpts, tokenAddr, fee)
}

// SetGuardianFee is a paid mutator transaction binding the contract method 0xeebede83.
//
// Solidity: function setGuardianFee(tokenAddr address, fee uint256) returns()
func (_Dospayment *DospaymentTransactorSession) SetGuardianFee(tokenAddr common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetGuardianFee(&_Dospayment.TransactOpts, tokenAddr, fee)
}

// SetGuardianFunds is a paid mutator transaction binding the contract method 0x5a1fa503.
//
// Solidity: function setGuardianFunds(fundsAddr address, tokenAddr address) returns()
func (_Dospayment *DospaymentTransactor) SetGuardianFunds(opts *bind.TransactOpts, fundsAddr common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setGuardianFunds", fundsAddr, tokenAddr)
}

// SetGuardianFunds is a paid mutator transaction binding the contract method 0x5a1fa503.
//
// Solidity: function setGuardianFunds(fundsAddr address, tokenAddr address) returns()
func (_Dospayment *DospaymentSession) SetGuardianFunds(fundsAddr common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetGuardianFunds(&_Dospayment.TransactOpts, fundsAddr, tokenAddr)
}

// SetGuardianFunds is a paid mutator transaction binding the contract method 0x5a1fa503.
//
// Solidity: function setGuardianFunds(fundsAddr address, tokenAddr address) returns()
func (_Dospayment *DospaymentTransactorSession) SetGuardianFunds(fundsAddr common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetGuardianFunds(&_Dospayment.TransactOpts, fundsAddr, tokenAddr)
}

// SetNetworkToken is a paid mutator transaction binding the contract method 0x17107c49.
//
// Solidity: function setNetworkToken(addr address) returns()
func (_Dospayment *DospaymentTransactor) SetNetworkToken(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setNetworkToken", addr)
}

// SetNetworkToken is a paid mutator transaction binding the contract method 0x17107c49.
//
// Solidity: function setNetworkToken(addr address) returns()
func (_Dospayment *DospaymentSession) SetNetworkToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetNetworkToken(&_Dospayment.TransactOpts, addr)
}

// SetNetworkToken is a paid mutator transaction binding the contract method 0x17107c49.
//
// Solidity: function setNetworkToken(addr address) returns()
func (_Dospayment *DospaymentTransactorSession) SetNetworkToken(addr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetNetworkToken(&_Dospayment.TransactOpts, addr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(consumer address, tokenAddr address) returns()
func (_Dospayment *DospaymentTransactor) SetPaymentMethod(opts *bind.TransactOpts, consumer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setPaymentMethod", consumer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(consumer address, tokenAddr address) returns()
func (_Dospayment *DospaymentSession) SetPaymentMethod(consumer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetPaymentMethod(&_Dospayment.TransactOpts, consumer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(consumer address, tokenAddr address) returns()
func (_Dospayment *DospaymentTransactorSession) SetPaymentMethod(consumer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetPaymentMethod(&_Dospayment.TransactOpts, consumer, tokenAddr)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x3939c401.
//
// Solidity: function setServiceFee(tokenAddr address, serviceType uint256, fee uint256) returns()
func (_Dospayment *DospaymentTransactor) SetServiceFee(opts *bind.TransactOpts, tokenAddr common.Address, serviceType *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setServiceFee", tokenAddr, serviceType, fee)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x3939c401.
//
// Solidity: function setServiceFee(tokenAddr address, serviceType uint256, fee uint256) returns()
func (_Dospayment *DospaymentSession) SetServiceFee(tokenAddr common.Address, serviceType *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetServiceFee(&_Dospayment.TransactOpts, tokenAddr, serviceType, fee)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x3939c401.
//
// Solidity: function setServiceFee(tokenAddr address, serviceType uint256, fee uint256) returns()
func (_Dospayment *DospaymentTransactorSession) SetServiceFee(tokenAddr common.Address, serviceType *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetServiceFee(&_Dospayment.TransactOpts, tokenAddr, serviceType, fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dospayment *DospaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Dospayment *DospaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.TransferOwnership(&_Dospayment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
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
// Solidity: function withdrawAll(tokenAddr uint256) returns()
func (_Dospayment *DospaymentTransactor) WithdrawAll(opts *bind.TransactOpts, tokenAddr *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "withdrawAll", tokenAddr)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(tokenAddr uint256) returns()
func (_Dospayment *DospaymentSession) WithdrawAll(tokenAddr *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.WithdrawAll(&_Dospayment.TransactOpts, tokenAddr)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(tokenAddr uint256) returns()
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
// Solidity: e LogChargeServiceFee(consumer address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256)
func (_Dospayment *DospaymentFilterer) FilterLogChargeServiceFee(opts *bind.FilterOpts) (*DospaymentLogChargeServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogChargeServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogChargeServiceFeeIterator{contract: _Dospayment.contract, event: "LogChargeServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogChargeServiceFee is a free log subscription operation binding the contract event 0xa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba.
//
// Solidity: e LogChargeServiceFee(consumer address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256)
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
// Solidity: e LogClaimGuardianFee(nodeAddr address, tokenAddr address, feeForSubmitter uint256, sender address)
func (_Dospayment *DospaymentFilterer) FilterLogClaimGuardianFee(opts *bind.FilterOpts) (*DospaymentLogClaimGuardianFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogClaimGuardianFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogClaimGuardianFeeIterator{contract: _Dospayment.contract, event: "LogClaimGuardianFee", logs: logs, sub: sub}, nil
}

// WatchLogClaimGuardianFee is a free log subscription operation binding the contract event 0x47ad88344c408450ef0ccab93ed97dd83af7a27dedfaa205c0725cfc4ca819cd.
//
// Solidity: e LogClaimGuardianFee(nodeAddr address, tokenAddr address, feeForSubmitter uint256, sender address)
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
// Solidity: e LogClaimServiceFee(nodeAddr address, tokenAddr address, requestID uint256, serviceType uint256, feeForSubmitter uint256)
func (_Dospayment *DospaymentFilterer) FilterLogClaimServiceFee(opts *bind.FilterOpts) (*DospaymentLogClaimServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogClaimServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogClaimServiceFeeIterator{contract: _Dospayment.contract, event: "LogClaimServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogClaimServiceFee is a free log subscription operation binding the contract event 0xab9fe896064c2c9dfd31acebc7d522b311e5f2e7d1b4965ac0cfd5a4abec813a.
//
// Solidity: e LogClaimServiceFee(nodeAddr address, tokenAddr address, requestID uint256, serviceType uint256, feeForSubmitter uint256)
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
// Solidity: e LogRefundServiceFee(consumer address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256)
func (_Dospayment *DospaymentFilterer) FilterLogRefundServiceFee(opts *bind.FilterOpts) (*DospaymentLogRefundServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogRefundServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogRefundServiceFeeIterator{contract: _Dospayment.contract, event: "LogRefundServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogRefundServiceFee is a free log subscription operation binding the contract event 0xde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa.
//
// Solidity: e LogRefundServiceFee(consumer address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256)
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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e UpdateDropBurnMaxQuota(oldQuota uint256, newQuota uint256)
func (_Dospayment *DospaymentFilterer) FilterUpdateDropBurnMaxQuota(opts *bind.FilterOpts) (*DospaymentUpdateDropBurnMaxQuotaIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "UpdateDropBurnMaxQuota")
	if err != nil {
		return nil, err
	}
	return &DospaymentUpdateDropBurnMaxQuotaIterator{contract: _Dospayment.contract, event: "UpdateDropBurnMaxQuota", logs: logs, sub: sub}, nil
}

// WatchUpdateDropBurnMaxQuota is a free log subscription operation binding the contract event 0x0aee95cca46da64ee373e28dee5994361b4002c54035d92932c9825b76382e99.
//
// Solidity: e UpdateDropBurnMaxQuota(oldQuota uint256, newQuota uint256)
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
// Solidity: e UpdateDropBurnTokenAddress(oldAddress address, newAddress address)
func (_Dospayment *DospaymentFilterer) FilterUpdateDropBurnTokenAddress(opts *bind.FilterOpts) (*DospaymentUpdateDropBurnTokenAddressIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "UpdateDropBurnTokenAddress")
	if err != nil {
		return nil, err
	}
	return &DospaymentUpdateDropBurnTokenAddressIterator{contract: _Dospayment.contract, event: "UpdateDropBurnTokenAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateDropBurnTokenAddress is a free log subscription operation binding the contract event 0xfc8013dfb0c8d38f3bcab9239bd5712457c48919b272cdb109488549199a0173.
//
// Solidity: e UpdateDropBurnTokenAddress(oldAddress address, newAddress address)
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
// Solidity: e UpdateNetworkTokenAddress(oldAddress address, newAddress address)
func (_Dospayment *DospaymentFilterer) FilterUpdateNetworkTokenAddress(opts *bind.FilterOpts) (*DospaymentUpdateNetworkTokenAddressIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "UpdateNetworkTokenAddress")
	if err != nil {
		return nil, err
	}
	return &DospaymentUpdateNetworkTokenAddressIterator{contract: _Dospayment.contract, event: "UpdateNetworkTokenAddress", logs: logs, sub: sub}, nil
}

// WatchUpdateNetworkTokenAddress is a free log subscription operation binding the contract event 0x4d27a2adceae86b92fb74fb7e8f96dc902d917e243fbff389b5a793c9040dafe.
//
// Solidity: e UpdateNetworkTokenAddress(oldAddress address, newAddress address)
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
