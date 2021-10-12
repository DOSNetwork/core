// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dospayment

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DospaymentMetaData contains all meta data concerning the Dospayment contract.
var DospaymentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardianFundsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogChargeServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeForSubmitter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LogClaimGuardianFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSubmitter\",\"type\":\"bool\"}],\"name\":\"LogRecordServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogRefundServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"UpdatePaymentAdmin\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"}],\"name\":\"chargeServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardianAddr\",\"type\":\"address\"}],\"name\":\"claimGuardianReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultGuardianFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultSubmitterCut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultSystemRandomFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultUserQueryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultUserRandomFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeLists\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"submitterCut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guardianFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"}],\"name\":\"getServiceTypeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"guardianFundsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"guardianFundsTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"}],\"name\":\"hasServiceFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardianFundsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nodeClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"nodeClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"nodeFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"paymentInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentMethods\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"workers\",\"type\":\"address[]\"}],\"name\":\"recordServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"refundServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"submitterCut\",\"type\":\"uint256\"}],\"name\":\"setFeeDividend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setGuardianFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fundsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setGuardianFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setPaymentMethod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405162001eed38038062001eed8339818101604052606081101561003557600080fd5b508051602082015160409092015190919061005a8383836001600160e01b0361006216565b50505061016d565b6007546001600160a01b031615801561008457506008546001600160a01b0316155b6100d5576040805162461bcd60e51b815260206004820152601360248201527f616c72656164792d696e697469616c697a656400000000000000000000000000604482015290519081900360640190fd5b60008054336001600160a01b03199182161782556007805482166001600160a01b039687161790556005805482169486169490941790935560068054841692909416918217909355600880549092168117909155815260026020818152604080842084805291829052808420674563918244f40000908190556001808652828620829055848652919094208490556004908201550155565b611d70806200017d6000396000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c80637aa9181b1161010f578063c60be4fd116100a2578063eebede8311610071578063eebede8314610674578063f39a19bf146106a0578063f851a440146106c6578063fa2c775e146106ce576101f0565b8063c60be4fd14610374578063cb7ca88c14610374578063d95eaa7a14610609578063e365036614610648576101f0565b8063a93eb110116100de578063a93eb1101461054f578063b73a3f8f14610575578063c0c53b8b146105a3578063c0f14e46146105db576101f0565b80637aa9181b1461040e5780638403f7dc1461044057806387d81789146104f857806391874ef714610547576101f0565b80634a0a382f11610187578063694732c611610156578063694732c6146103ba5780636dfa72b014610374578063704b6c02146103e0578063746c73fd14610406576101f0565b80634a0a382f14610374578063571028e31461037c5780635a1fa503146103845780636059775a146103b2576101f0565b80632c097993116101c35780632c09799314610298578063310ae824146102c45780633157f16d146103025780633939c40114610342576101f0565b806302b8b587146101f55780631efa5a981461021957806323ff34cb14610238578063240028e81461025e575b600080fd5b6101fd6106d6565b604080516001600160a01b039092168252519081900360200190f35b6102366004803603602081101561022f57600080fd5b50356106e5565b005b6102366004803603602081101561024e57600080fd5b50356001600160a01b0316610972565b6102846004803603602081101561027457600080fd5b50356001600160a01b0316610bf4565b604080519115158252519081900360200190f35b610236600480360360408110156102ae57600080fd5b506001600160a01b038135169060200135610ca6565b6102f0600480360360408110156102da57600080fd5b506001600160a01b038135169060200135610d62565b60408051918252519081900360200190f35b61031f6004803603602081101561031857600080fd5b5035610e31565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6102366004803603606081101561035857600080fd5b506001600160a01b038135169060208101359060400135610e58565b6102f0610f1b565b6102f0610f27565b6102366004803603604081101561039a57600080fd5b506001600160a01b0381358116916020013516610f2c565b6101fd610ffc565b6101fd600480360360208110156103d057600080fd5b50356001600160a01b031661100b565b610236600480360360208110156103f657600080fd5b50356001600160a01b0316611026565b6102f06110ee565b6102366004803603606081101561042457600080fd5b506001600160a01b03813516906020810135906040013561110e565b6102366004803603606081101561045657600080fd5b8135916001600160a01b036020820135169181019060608101604082013564010000000081111561048657600080fd5b82018360208201111561049857600080fd5b803590602001918460208302840111640100000000831117156104ba57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611412945050505050565b6105156004803603602081101561050e57600080fd5b5035611813565b604080516001600160a01b03958616815293909416602084015282840191909152606082015290519081900360800190f35b6101fd611849565b6102f06004803603602081101561056557600080fd5b50356001600160a01b0316611858565b6102366004803603604081101561058b57600080fd5b506001600160a01b0381358116916020013516611872565b610236600480360360608110156105b957600080fd5b506001600160a01b0381358116916020810135821691604090910135166118f7565b6102f0600480360360408110156105f157600080fd5b506001600160a01b03813581169160200135166119f8565b61062f6004803603602081101561061f57600080fd5b50356001600160a01b0316611a23565b6040805192835260208301919091528051918290030190f35b6102846004803603604081101561065e57600080fd5b506001600160a01b038135169060200135611a3f565b6102366004803603604081101561068a57600080fd5b506001600160a01b038135169060200135611b74565b6102f0600480360360208110156106b657600080fd5b50356001600160a01b0316611c31565b6101fd611c4c565b6101fd611c5b565b6008546001600160a01b031681565b6000546001600160a01b03163314610730576040805162461bcd60e51b815260206004820152600960248201526837b7363ca0b236b4b760b91b604482015290519081900360640190fd5b600081815260036020819052604090912001548190610786576040805162461bcd60e51b815260206004820152600d60248201526c1b9bcb5999594b585b5bdd5b9d609a1b604482015290519081900360640190fd5b6000818152600360205260409020546001600160a01b03166107df576040805162461bcd60e51b815260206004820152600d60248201526c6e6f2d70617965722d696e666f60981b604482015290519081900360640190fd5b6000818152600360205260409020600101546001600160a01b031661083f576040805162461bcd60e51b81526020600482015260116024820152706e6f2d6665652d746f6b656e2d696e666f60781b604482015290519081900360640190fd5b6000828152600360208181526040808420928301805460028501805460018701805488546001600160a01b0319808216909a55988116909155918890559690925582516001600160a01b039586168082529590921693820184905281830188905260608201869052608082018190529151919493917fde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa9181900360a00190a1816001600160a01b031663a9059cbb82866040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561093e57600080fd5b505af1158015610952573d6000803e3d6000fd5b505050506040513d602081101561096857600080fd5b5050505050505050565b600760009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156109c057600080fd5b505afa1580156109d4573d6000803e3d6000fd5b505050506040513d60208110156109ea57600080fd5b50516001600160a01b03163314610a3d576040805162461bcd60e51b81526020600482015260126024820152716e6f742d66726f6d2d646f732d70726f787960701b604482015290519081900360640190fd5b6005546001600160a01b0316610a9a576040805162461bcd60e51b815260206004820152601c60248201527f6e6f742d76616c69642d677561726469616e2d66756e642d6164647200000000604482015290519081900360640190fd5b6006546001600160a01b0316610af7576040805162461bcd60e51b815260206004820152601d60248201527f6e6f742d76616c69642d677561726469616e2d746f6b656e2d61646472000000604482015290519081900360640190fd5b6006546001600160a01b039081166000818152600260208181526040928390209091015482519486168552908401929092528281018290523360608401525190917f47ad88344c408450ef0ccab93ed97dd83af7a27dedfaa205c0725cfc4ca819cd919081900360800190a1600654600554604080516323b872dd60e01b81526001600160a01b039283166004820152858316602482015260448101859052905191909216916323b872dd9160648083019260209291908290030181600087803b158015610bc457600080fd5b505af1158015610bd8573d6000803e3d6000fd5b505050506040513d6020811015610bee57600080fd5b50505050565b60006001600160a01b038216610c0c57506000610ca1565b6001600160a01b0382166000908152600260209081526040808320838052909152902054610c3c57506000610ca1565b6001600160a01b038216600090815260026020908152604080832060018452909152902054610c6d57506000610ca1565b6001600160a01b0382166000908152600260208181526040808420928452919052902054610c9d57506000610ca1565b5060015b919050565b6000546001600160a01b03163314610cf1576040805162461bcd60e51b815260206004820152600960248201526837b7363ca0b236b4b760b91b604482015290519081900360640190fd5b6001600160a01b038216610d43576040805162461bcd60e51b81526020600482015260146024820152733737ba16bb30b634b216ba37b5b2b716b0b2323960611b604482015290519081900360640190fd5b6001600160a01b03909116600090815260026020526040902060010155565b60006001600160a01b03831615801590610d9757506001600160a01b0383166000908152600260208190526040909120015415155b8015610dbd57506001600160a01b03831660009081526002602052604090206001015415155b610e05576040805162461bcd60e51b81526020600482015260146024820152733737ba16bb30b634b216ba37b5b2b716b0b2323960611b604482015290519081900360640190fd5b506001600160a01b03821660009081526002602090815260408083208484529091529020545b92915050565b600090815260036020819052604090912060018101549101546001600160a01b0390911691565b6000546001600160a01b03163314610ea3576040805162461bcd60e51b815260206004820152600960248201526837b7363ca0b236b4b760b91b604482015290519081900360640190fd5b6001600160a01b038316610ef5576040805162461bcd60e51b81526020600482015260146024820152733737ba16bb30b634b216ba37b5b2b716b0b2323960611b604482015290519081900360640190fd5b6001600160a01b0390921660009081526002602090815260408083209383529290522055565b674563918244f4000081565b600481565b6000546001600160a01b03163314610f77576040805162461bcd60e51b815260206004820152600960248201526837b7363ca0b236b4b760b91b604482015290519081900360640190fd5b80610f8181610bf4565b610fcd576040805162461bcd60e51b81526020600482015260186024820152773737ba16b9bab83837b93a32b216ba37b5b2b716b0b2323960411b604482015290519081900360640190fd5b50600580546001600160a01b039384166001600160a01b03199182161790915560068054929093169116179055565b6005546001600160a01b031681565b6001602052600090815260409020546001600160a01b031681565b6000546001600160a01b03163314611071576040805162461bcd60e51b815260206004820152600960248201526837b7363ca0b236b4b760b91b604482015290519081900360640190fd5b6001600160a01b03811661108457600080fd5b600054604080516001600160a01b039283168152918316602083015280517fe0be0cc9c5ea7d7a7f3909ad261ab0e0fbc9aae3fe819864f60d81d91286dc309281900390910190a1600080546001600160a01b0319166001600160a01b0392909216919091179055565b6008546000906111099033906001600160a01b031681611c6a565b905090565b600760009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561115c57600080fd5b505afa158015611170573d6000803e3d6000fd5b505050506040513d602081101561118657600080fd5b50516001600160a01b031633146111d9576040805162461bcd60e51b81526020600482015260126024820152716e6f742d66726f6d2d646f732d70726f787960701b604482015290519081900360640190fd5b600060016000856001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03169050600060026000836001600160a01b03166001600160a01b0316815260200190815260200160002060000160008481526020019081526020016000205490506040518060800160405280866001600160a01b03168152602001836001600160a01b03168152602001848152602001828152506003600086815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160020155606082015181600301559050507fa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba858386868560405180866001600160a01b03166001600160a01b03168152602001856001600160a01b03166001600160a01b031681526020018481526020018381526020018281526020019550505050505060405180910390a1604080516323b872dd60e01b81526001600160a01b038781166004830152306024830152604482018490529151918416916323b872dd916064808201926020929091908290030181600087803b1580156113df57600080fd5b505af11580156113f3573d6000803e3d6000fd5b505050506040513d602081101561140957600080fd5b50505050505050565b600760009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561146057600080fd5b505afa158015611474573d6000803e3d6000fd5b505050506040513d602081101561148a57600080fd5b50516001600160a01b031633146114dd576040805162461bcd60e51b81526020600482015260126024820152716e6f742d66726f6d2d646f732d70726f787960701b604482015290519081900360640190fd5b600083815260036020819052604090912001548390611533576040805162461bcd60e51b815260206004820152600d60248201526c1b9bcb5999594b585b5bdd5b9d609a1b604482015290519081900360640190fd5b6000818152600360205260409020546001600160a01b031661158c576040805162461bcd60e51b815260206004820152600d60248201526c6e6f2d70617965722d696e666f60981b604482015290519081900360640190fd5b6000818152600360205260409020600101546001600160a01b03166115ec576040805162461bcd60e51b81526020600482015260116024820152706e6f2d6665652d746f6b656e2d696e666f60781b604482015290519081900360640190fd5b60008481526003602081815260408084206001808201805495830180546002808601805487546001600160a01b0319908116909855968a1690945592899055908890556001600160a01b0396871680895291865284882080840154978c16808a5260048852868a20848b528852988690208054600a909304988902928301905585519889529588018290528785018c9052606088018490526080880181905260a08801929092529251929591939290917f4758b94d44e129dcef9dc829628a55e921926b4383f3261f968ee8f9373571e7919081900360c00190a1600060018851038360010154600a038602816116df57fe5b04905060005b885181101561180657896001600160a01b031689828151811061170457fe5b60200260200101516001600160a01b0316146117fe5781600460008b848151811061172b57fe5b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206000896001600160a01b03166001600160a01b03168152602001908152602001600020600082825401925050819055507f4758b94d44e129dcef9dc829628a55e921926b4383f3261f968ee8f9373571e78982815181106117b057fe5b602090810291909101810151604080516001600160a01b039283168152918b16928201929092528082018e90526060810188905260808101859052600060a082015290519081900360c00190a15b6001016116e5565b5050505050505050505050565b600360208190526000918252604090912080546001820154600283015492909301546001600160a01b0391821693909116919084565b6007546001600160a01b031681565b600854600090610e2b9083906001600160a01b03166119f8565b8061187c81610bf4565b6118c8576040805162461bcd60e51b81526020600482015260186024820152773737ba16b9bab83837b93a32b216ba37b5b2b716b0b2323960411b604482015290519081900360640190fd5b506001600160a01b03918216600090815260016020526040902080546001600160a01b03191691909216179055565b6007546001600160a01b031615801561191957506008546001600160a01b0316155b611960576040805162461bcd60e51b8152602060048201526013602482015272185b1c9958591e4b5a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b60008054336001600160a01b03199182161782556007805482166001600160a01b039687161790556005805482169486169490941790935560068054841692909416918217909355600880549092168117909155815260026020818152604080842084805291829052808420674563918244f40000908190556001808652828620829055848652919094208490556004908201550155565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205490565b6002602081905260009182526040909120600181015491015482565b600754604080516321d39ecd60e11b815290516000926001600160a01b0316916343a73d9a916004808301926020929190829003018186803b158015611a8457600080fd5b505afa158015611a98573d6000803e3d6000fd5b505050506040513d6020811015611aae57600080fd5b50516001600160a01b0384811691161415611acb57506001610e2b565b6001600160a01b03808416600081815260016020908152604080832054909416808352600282528483208784528252918490205484516370a0823160e01b815260048101949094529351919392839285926370a082319260248082019391829003018186803b158015611b3d57600080fd5b505afa158015611b51573d6000803e3d6000fd5b505050506040513d6020811015611b6757600080fd5b5051101595945050505050565b6000546001600160a01b03163314611bbf576040805162461bcd60e51b815260206004820152600960248201526837b7363ca0b236b4b760b91b604482015290519081900360640190fd5b6001600160a01b038216611c11576040805162461bcd60e51b81526020600482015260146024820152733737ba16bb30b634b216ba37b5b2b716b0b2323960611b604482015290519081900360640190fd5b6001600160a01b0390911660009081526002602081905260409091200155565b600854600090610e2b9033906001600160a01b031684611c6a565b6000546001600160a01b031681565b6006546001600160a01b031681565b6001600160a01b0380841660009081526004602090815260408083209386168352929052908120548015611d33576001600160a01b038086166000908152600460208181526040808420898616808652908352818520859055815163a9059cbb60e01b8152958916938601939093526024850186905251919363a9059cbb93604480830194928390030190829087803b158015611d0657600080fd5b505af1158015611d1a573d6000803e3d6000fd5b505050506040513d6020811015611d3057600080fd5b50505b94935050505056fea265627a7a7231582088be7a92b146d4639c93b15382f33440a84888c1f510613fe892436ac6df06d564736f6c63430005110032",
}

// DospaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use DospaymentMetaData.ABI instead.
var DospaymentABI = DospaymentMetaData.ABI

// DospaymentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DospaymentMetaData.Bin instead.
var DospaymentBin = DospaymentMetaData.Bin

// DeployDospayment deploys a new Ethereum contract, binding an instance of Dospayment to it.
func DeployDospayment(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddr common.Address, _guardianFundsAddr common.Address, _tokenAddr common.Address) (common.Address, *types.Transaction, *Dospayment, error) {
	parsed, err := DospaymentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DospaymentBin), backend, _bridgeAddr, _guardianFundsAddr, _tokenAddr)
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
func (_Dospayment *DospaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Dospayment *DospaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Dospayment *DospaymentCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Dospayment *DospaymentSession) Admin() (common.Address, error) {
	return _Dospayment.Contract.Admin(&_Dospayment.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Dospayment *DospaymentCallerSession) Admin() (common.Address, error) {
	return _Dospayment.Contract.Admin(&_Dospayment.CallOpts)
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() view returns(address)
func (_Dospayment *DospaymentCaller) BridgeAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "bridgeAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() view returns(address)
func (_Dospayment *DospaymentSession) BridgeAddr() (common.Address, error) {
	return _Dospayment.Contract.BridgeAddr(&_Dospayment.CallOpts)
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() view returns(address)
func (_Dospayment *DospaymentCallerSession) BridgeAddr() (common.Address, error) {
	return _Dospayment.Contract.BridgeAddr(&_Dospayment.CallOpts)
}

// DefaultGuardianFee is a free data retrieval call binding the contract method 0x6dfa72b0.
//
// Solidity: function defaultGuardianFee() view returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultGuardianFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "defaultGuardianFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultGuardianFee is a free data retrieval call binding the contract method 0x6dfa72b0.
//
// Solidity: function defaultGuardianFee() view returns(uint256)
func (_Dospayment *DospaymentSession) DefaultGuardianFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultGuardianFee(&_Dospayment.CallOpts)
}

// DefaultGuardianFee is a free data retrieval call binding the contract method 0x6dfa72b0.
//
// Solidity: function defaultGuardianFee() view returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultGuardianFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultGuardianFee(&_Dospayment.CallOpts)
}

// DefaultSubmitterCut is a free data retrieval call binding the contract method 0x571028e3.
//
// Solidity: function defaultSubmitterCut() view returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultSubmitterCut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "defaultSubmitterCut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultSubmitterCut is a free data retrieval call binding the contract method 0x571028e3.
//
// Solidity: function defaultSubmitterCut() view returns(uint256)
func (_Dospayment *DospaymentSession) DefaultSubmitterCut() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSubmitterCut(&_Dospayment.CallOpts)
}

// DefaultSubmitterCut is a free data retrieval call binding the contract method 0x571028e3.
//
// Solidity: function defaultSubmitterCut() view returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultSubmitterCut() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSubmitterCut(&_Dospayment.CallOpts)
}

// DefaultSystemRandomFee is a free data retrieval call binding the contract method 0xcb7ca88c.
//
// Solidity: function defaultSystemRandomFee() view returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultSystemRandomFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "defaultSystemRandomFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultSystemRandomFee is a free data retrieval call binding the contract method 0xcb7ca88c.
//
// Solidity: function defaultSystemRandomFee() view returns(uint256)
func (_Dospayment *DospaymentSession) DefaultSystemRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSystemRandomFee(&_Dospayment.CallOpts)
}

// DefaultSystemRandomFee is a free data retrieval call binding the contract method 0xcb7ca88c.
//
// Solidity: function defaultSystemRandomFee() view returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultSystemRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSystemRandomFee(&_Dospayment.CallOpts)
}

// DefaultTokenAddr is a free data retrieval call binding the contract method 0x02b8b587.
//
// Solidity: function defaultTokenAddr() view returns(address)
func (_Dospayment *DospaymentCaller) DefaultTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "defaultTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultTokenAddr is a free data retrieval call binding the contract method 0x02b8b587.
//
// Solidity: function defaultTokenAddr() view returns(address)
func (_Dospayment *DospaymentSession) DefaultTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.DefaultTokenAddr(&_Dospayment.CallOpts)
}

// DefaultTokenAddr is a free data retrieval call binding the contract method 0x02b8b587.
//
// Solidity: function defaultTokenAddr() view returns(address)
func (_Dospayment *DospaymentCallerSession) DefaultTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.DefaultTokenAddr(&_Dospayment.CallOpts)
}

// DefaultUserQueryFee is a free data retrieval call binding the contract method 0xc60be4fd.
//
// Solidity: function defaultUserQueryFee() view returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultUserQueryFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "defaultUserQueryFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultUserQueryFee is a free data retrieval call binding the contract method 0xc60be4fd.
//
// Solidity: function defaultUserQueryFee() view returns(uint256)
func (_Dospayment *DospaymentSession) DefaultUserQueryFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserQueryFee(&_Dospayment.CallOpts)
}

// DefaultUserQueryFee is a free data retrieval call binding the contract method 0xc60be4fd.
//
// Solidity: function defaultUserQueryFee() view returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultUserQueryFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserQueryFee(&_Dospayment.CallOpts)
}

// DefaultUserRandomFee is a free data retrieval call binding the contract method 0x4a0a382f.
//
// Solidity: function defaultUserRandomFee() view returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultUserRandomFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "defaultUserRandomFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultUserRandomFee is a free data retrieval call binding the contract method 0x4a0a382f.
//
// Solidity: function defaultUserRandomFee() view returns(uint256)
func (_Dospayment *DospaymentSession) DefaultUserRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserRandomFee(&_Dospayment.CallOpts)
}

// DefaultUserRandomFee is a free data retrieval call binding the contract method 0x4a0a382f.
//
// Solidity: function defaultUserRandomFee() view returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultUserRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserRandomFee(&_Dospayment.CallOpts)
}

// FeeLists is a free data retrieval call binding the contract method 0xd95eaa7a.
//
// Solidity: function feeLists(address ) view returns(uint256 submitterCut, uint256 guardianFee)
func (_Dospayment *DospaymentCaller) FeeLists(opts *bind.CallOpts, arg0 common.Address) (struct {
	SubmitterCut *big.Int
	GuardianFee  *big.Int
}, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "feeLists", arg0)

	outstruct := new(struct {
		SubmitterCut *big.Int
		GuardianFee  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SubmitterCut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.GuardianFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FeeLists is a free data retrieval call binding the contract method 0xd95eaa7a.
//
// Solidity: function feeLists(address ) view returns(uint256 submitterCut, uint256 guardianFee)
func (_Dospayment *DospaymentSession) FeeLists(arg0 common.Address) (struct {
	SubmitterCut *big.Int
	GuardianFee  *big.Int
}, error) {
	return _Dospayment.Contract.FeeLists(&_Dospayment.CallOpts, arg0)
}

// FeeLists is a free data retrieval call binding the contract method 0xd95eaa7a.
//
// Solidity: function feeLists(address ) view returns(uint256 submitterCut, uint256 guardianFee)
func (_Dospayment *DospaymentCallerSession) FeeLists(arg0 common.Address) (struct {
	SubmitterCut *big.Int
	GuardianFee  *big.Int
}, error) {
	return _Dospayment.Contract.FeeLists(&_Dospayment.CallOpts, arg0)
}

// GetServiceTypeFee is a free data retrieval call binding the contract method 0x310ae824.
//
// Solidity: function getServiceTypeFee(address tokenAddr, uint256 serviceType) view returns(uint256)
func (_Dospayment *DospaymentCaller) GetServiceTypeFee(opts *bind.CallOpts, tokenAddr common.Address, serviceType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "getServiceTypeFee", tokenAddr, serviceType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetServiceTypeFee is a free data retrieval call binding the contract method 0x310ae824.
//
// Solidity: function getServiceTypeFee(address tokenAddr, uint256 serviceType) view returns(uint256)
func (_Dospayment *DospaymentSession) GetServiceTypeFee(tokenAddr common.Address, serviceType *big.Int) (*big.Int, error) {
	return _Dospayment.Contract.GetServiceTypeFee(&_Dospayment.CallOpts, tokenAddr, serviceType)
}

// GetServiceTypeFee is a free data retrieval call binding the contract method 0x310ae824.
//
// Solidity: function getServiceTypeFee(address tokenAddr, uint256 serviceType) view returns(uint256)
func (_Dospayment *DospaymentCallerSession) GetServiceTypeFee(tokenAddr common.Address, serviceType *big.Int) (*big.Int, error) {
	return _Dospayment.Contract.GetServiceTypeFee(&_Dospayment.CallOpts, tokenAddr, serviceType)
}

// GuardianFundsAddr is a free data retrieval call binding the contract method 0x6059775a.
//
// Solidity: function guardianFundsAddr() view returns(address)
func (_Dospayment *DospaymentCaller) GuardianFundsAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "guardianFundsAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuardianFundsAddr is a free data retrieval call binding the contract method 0x6059775a.
//
// Solidity: function guardianFundsAddr() view returns(address)
func (_Dospayment *DospaymentSession) GuardianFundsAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsAddr(&_Dospayment.CallOpts)
}

// GuardianFundsAddr is a free data retrieval call binding the contract method 0x6059775a.
//
// Solidity: function guardianFundsAddr() view returns(address)
func (_Dospayment *DospaymentCallerSession) GuardianFundsAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsAddr(&_Dospayment.CallOpts)
}

// GuardianFundsTokenAddr is a free data retrieval call binding the contract method 0xfa2c775e.
//
// Solidity: function guardianFundsTokenAddr() view returns(address)
func (_Dospayment *DospaymentCaller) GuardianFundsTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "guardianFundsTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GuardianFundsTokenAddr is a free data retrieval call binding the contract method 0xfa2c775e.
//
// Solidity: function guardianFundsTokenAddr() view returns(address)
func (_Dospayment *DospaymentSession) GuardianFundsTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsTokenAddr(&_Dospayment.CallOpts)
}

// GuardianFundsTokenAddr is a free data retrieval call binding the contract method 0xfa2c775e.
//
// Solidity: function guardianFundsTokenAddr() view returns(address)
func (_Dospayment *DospaymentCallerSession) GuardianFundsTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsTokenAddr(&_Dospayment.CallOpts)
}

// HasServiceFee is a free data retrieval call binding the contract method 0xe3650366.
//
// Solidity: function hasServiceFee(address payer, uint256 serviceType) view returns(bool)
func (_Dospayment *DospaymentCaller) HasServiceFee(opts *bind.CallOpts, payer common.Address, serviceType *big.Int) (bool, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "hasServiceFee", payer, serviceType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasServiceFee is a free data retrieval call binding the contract method 0xe3650366.
//
// Solidity: function hasServiceFee(address payer, uint256 serviceType) view returns(bool)
func (_Dospayment *DospaymentSession) HasServiceFee(payer common.Address, serviceType *big.Int) (bool, error) {
	return _Dospayment.Contract.HasServiceFee(&_Dospayment.CallOpts, payer, serviceType)
}

// HasServiceFee is a free data retrieval call binding the contract method 0xe3650366.
//
// Solidity: function hasServiceFee(address payer, uint256 serviceType) view returns(bool)
func (_Dospayment *DospaymentCallerSession) HasServiceFee(payer common.Address, serviceType *big.Int) (bool, error) {
	return _Dospayment.Contract.HasServiceFee(&_Dospayment.CallOpts, payer, serviceType)
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(address tokenAddr) view returns(bool)
func (_Dospayment *DospaymentCaller) IsSupportedToken(opts *bind.CallOpts, tokenAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "isSupportedToken", tokenAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(address tokenAddr) view returns(bool)
func (_Dospayment *DospaymentSession) IsSupportedToken(tokenAddr common.Address) (bool, error) {
	return _Dospayment.Contract.IsSupportedToken(&_Dospayment.CallOpts, tokenAddr)
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(address tokenAddr) view returns(bool)
func (_Dospayment *DospaymentCallerSession) IsSupportedToken(tokenAddr common.Address) (bool, error) {
	return _Dospayment.Contract.IsSupportedToken(&_Dospayment.CallOpts, tokenAddr)
}

// NodeFeeBalance is a free data retrieval call binding the contract method 0xa93eb110.
//
// Solidity: function nodeFeeBalance(address nodeAddr) view returns(uint256)
func (_Dospayment *DospaymentCaller) NodeFeeBalance(opts *bind.CallOpts, nodeAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "nodeFeeBalance", nodeAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeFeeBalance is a free data retrieval call binding the contract method 0xa93eb110.
//
// Solidity: function nodeFeeBalance(address nodeAddr) view returns(uint256)
func (_Dospayment *DospaymentSession) NodeFeeBalance(nodeAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeFeeBalance(&_Dospayment.CallOpts, nodeAddr)
}

// NodeFeeBalance is a free data retrieval call binding the contract method 0xa93eb110.
//
// Solidity: function nodeFeeBalance(address nodeAddr) view returns(uint256)
func (_Dospayment *DospaymentCallerSession) NodeFeeBalance(nodeAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeFeeBalance(&_Dospayment.CallOpts, nodeAddr)
}

// NodeFeeBalance0 is a free data retrieval call binding the contract method 0xc0f14e46.
//
// Solidity: function nodeFeeBalance(address nodeAddr, address tokenAddr) view returns(uint256)
func (_Dospayment *DospaymentCaller) NodeFeeBalance0(opts *bind.CallOpts, nodeAddr common.Address, tokenAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "nodeFeeBalance0", nodeAddr, tokenAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeFeeBalance0 is a free data retrieval call binding the contract method 0xc0f14e46.
//
// Solidity: function nodeFeeBalance(address nodeAddr, address tokenAddr) view returns(uint256)
func (_Dospayment *DospaymentSession) NodeFeeBalance0(nodeAddr common.Address, tokenAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeFeeBalance0(&_Dospayment.CallOpts, nodeAddr, tokenAddr)
}

// NodeFeeBalance0 is a free data retrieval call binding the contract method 0xc0f14e46.
//
// Solidity: function nodeFeeBalance(address nodeAddr, address tokenAddr) view returns(uint256)
func (_Dospayment *DospaymentCallerSession) NodeFeeBalance0(nodeAddr common.Address, tokenAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeFeeBalance0(&_Dospayment.CallOpts, nodeAddr, tokenAddr)
}

// PaymentInfo is a free data retrieval call binding the contract method 0x3157f16d.
//
// Solidity: function paymentInfo(uint256 requestID) view returns(address, uint256)
func (_Dospayment *DospaymentCaller) PaymentInfo(opts *bind.CallOpts, requestID *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "paymentInfo", requestID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PaymentInfo is a free data retrieval call binding the contract method 0x3157f16d.
//
// Solidity: function paymentInfo(uint256 requestID) view returns(address, uint256)
func (_Dospayment *DospaymentSession) PaymentInfo(requestID *big.Int) (common.Address, *big.Int, error) {
	return _Dospayment.Contract.PaymentInfo(&_Dospayment.CallOpts, requestID)
}

// PaymentInfo is a free data retrieval call binding the contract method 0x3157f16d.
//
// Solidity: function paymentInfo(uint256 requestID) view returns(address, uint256)
func (_Dospayment *DospaymentCallerSession) PaymentInfo(requestID *big.Int) (common.Address, *big.Int, error) {
	return _Dospayment.Contract.PaymentInfo(&_Dospayment.CallOpts, requestID)
}

// PaymentMethods is a free data retrieval call binding the contract method 0x694732c6.
//
// Solidity: function paymentMethods(address ) view returns(address)
func (_Dospayment *DospaymentCaller) PaymentMethods(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "paymentMethods", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentMethods is a free data retrieval call binding the contract method 0x694732c6.
//
// Solidity: function paymentMethods(address ) view returns(address)
func (_Dospayment *DospaymentSession) PaymentMethods(arg0 common.Address) (common.Address, error) {
	return _Dospayment.Contract.PaymentMethods(&_Dospayment.CallOpts, arg0)
}

// PaymentMethods is a free data retrieval call binding the contract method 0x694732c6.
//
// Solidity: function paymentMethods(address ) view returns(address)
func (_Dospayment *DospaymentCallerSession) PaymentMethods(arg0 common.Address) (common.Address, error) {
	return _Dospayment.Contract.PaymentMethods(&_Dospayment.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(address payer, address tokenAddr, uint256 serviceType, uint256 amount)
func (_Dospayment *DospaymentCaller) Payments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Payer       common.Address
	TokenAddr   common.Address
	ServiceType *big.Int
	Amount      *big.Int
}, error) {
	var out []interface{}
	err := _Dospayment.contract.Call(opts, &out, "payments", arg0)

	outstruct := new(struct {
		Payer       common.Address
		TokenAddr   common.Address
		ServiceType *big.Int
		Amount      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Payer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenAddr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ServiceType = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(address payer, address tokenAddr, uint256 serviceType, uint256 amount)
func (_Dospayment *DospaymentSession) Payments(arg0 *big.Int) (struct {
	Payer       common.Address
	TokenAddr   common.Address
	ServiceType *big.Int
	Amount      *big.Int
}, error) {
	return _Dospayment.Contract.Payments(&_Dospayment.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments(uint256 ) view returns(address payer, address tokenAddr, uint256 serviceType, uint256 amount)
func (_Dospayment *DospaymentCallerSession) Payments(arg0 *big.Int) (struct {
	Payer       common.Address
	TokenAddr   common.Address
	ServiceType *big.Int
	Amount      *big.Int
}, error) {
	return _Dospayment.Contract.Payments(&_Dospayment.CallOpts, arg0)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(address payer, uint256 requestID, uint256 serviceType) returns()
func (_Dospayment *DospaymentTransactor) ChargeServiceFee(opts *bind.TransactOpts, payer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "chargeServiceFee", payer, requestID, serviceType)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(address payer, uint256 requestID, uint256 serviceType) returns()
func (_Dospayment *DospaymentSession) ChargeServiceFee(payer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.ChargeServiceFee(&_Dospayment.TransactOpts, payer, requestID, serviceType)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(address payer, uint256 requestID, uint256 serviceType) returns()
func (_Dospayment *DospaymentTransactorSession) ChargeServiceFee(payer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.ChargeServiceFee(&_Dospayment.TransactOpts, payer, requestID, serviceType)
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

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _bridgeAddr, address _guardianFundsAddr, address _tokenAddr) returns()
func (_Dospayment *DospaymentTransactor) Initialize(opts *bind.TransactOpts, _bridgeAddr common.Address, _guardianFundsAddr common.Address, _tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "initialize", _bridgeAddr, _guardianFundsAddr, _tokenAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _bridgeAddr, address _guardianFundsAddr, address _tokenAddr) returns()
func (_Dospayment *DospaymentSession) Initialize(_bridgeAddr common.Address, _guardianFundsAddr common.Address, _tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.Initialize(&_Dospayment.TransactOpts, _bridgeAddr, _guardianFundsAddr, _tokenAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _bridgeAddr, address _guardianFundsAddr, address _tokenAddr) returns()
func (_Dospayment *DospaymentTransactorSession) Initialize(_bridgeAddr common.Address, _guardianFundsAddr common.Address, _tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.Initialize(&_Dospayment.TransactOpts, _bridgeAddr, _guardianFundsAddr, _tokenAddr)
}

// NodeClaim is a paid mutator transaction binding the contract method 0x746c73fd.
//
// Solidity: function nodeClaim() returns(uint256)
func (_Dospayment *DospaymentTransactor) NodeClaim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "nodeClaim")
}

// NodeClaim is a paid mutator transaction binding the contract method 0x746c73fd.
//
// Solidity: function nodeClaim() returns(uint256)
func (_Dospayment *DospaymentSession) NodeClaim() (*types.Transaction, error) {
	return _Dospayment.Contract.NodeClaim(&_Dospayment.TransactOpts)
}

// NodeClaim is a paid mutator transaction binding the contract method 0x746c73fd.
//
// Solidity: function nodeClaim() returns(uint256)
func (_Dospayment *DospaymentTransactorSession) NodeClaim() (*types.Transaction, error) {
	return _Dospayment.Contract.NodeClaim(&_Dospayment.TransactOpts)
}

// NodeClaim0 is a paid mutator transaction binding the contract method 0xf39a19bf.
//
// Solidity: function nodeClaim(address to) returns(uint256)
func (_Dospayment *DospaymentTransactor) NodeClaim0(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "nodeClaim0", to)
}

// NodeClaim0 is a paid mutator transaction binding the contract method 0xf39a19bf.
//
// Solidity: function nodeClaim(address to) returns(uint256)
func (_Dospayment *DospaymentSession) NodeClaim0(to common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.NodeClaim0(&_Dospayment.TransactOpts, to)
}

// NodeClaim0 is a paid mutator transaction binding the contract method 0xf39a19bf.
//
// Solidity: function nodeClaim(address to) returns(uint256)
func (_Dospayment *DospaymentTransactorSession) NodeClaim0(to common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.NodeClaim0(&_Dospayment.TransactOpts, to)
}

// RecordServiceFee is a paid mutator transaction binding the contract method 0x8403f7dc.
//
// Solidity: function recordServiceFee(uint256 requestID, address submitter, address[] workers) returns()
func (_Dospayment *DospaymentTransactor) RecordServiceFee(opts *bind.TransactOpts, requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "recordServiceFee", requestID, submitter, workers)
}

// RecordServiceFee is a paid mutator transaction binding the contract method 0x8403f7dc.
//
// Solidity: function recordServiceFee(uint256 requestID, address submitter, address[] workers) returns()
func (_Dospayment *DospaymentSession) RecordServiceFee(requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.RecordServiceFee(&_Dospayment.TransactOpts, requestID, submitter, workers)
}

// RecordServiceFee is a paid mutator transaction binding the contract method 0x8403f7dc.
//
// Solidity: function recordServiceFee(uint256 requestID, address submitter, address[] workers) returns()
func (_Dospayment *DospaymentTransactorSession) RecordServiceFee(requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.RecordServiceFee(&_Dospayment.TransactOpts, requestID, submitter, workers)
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

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Dospayment *DospaymentTransactor) SetAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setAdmin", newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Dospayment *DospaymentSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetAdmin(&_Dospayment.TransactOpts, newAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address newAdmin) returns()
func (_Dospayment *DospaymentTransactorSession) SetAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetAdmin(&_Dospayment.TransactOpts, newAdmin)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0x2c097993.
//
// Solidity: function setFeeDividend(address tokenAddr, uint256 submitterCut) returns()
func (_Dospayment *DospaymentTransactor) SetFeeDividend(opts *bind.TransactOpts, tokenAddr common.Address, submitterCut *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setFeeDividend", tokenAddr, submitterCut)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0x2c097993.
//
// Solidity: function setFeeDividend(address tokenAddr, uint256 submitterCut) returns()
func (_Dospayment *DospaymentSession) SetFeeDividend(tokenAddr common.Address, submitterCut *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetFeeDividend(&_Dospayment.TransactOpts, tokenAddr, submitterCut)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0x2c097993.
//
// Solidity: function setFeeDividend(address tokenAddr, uint256 submitterCut) returns()
func (_Dospayment *DospaymentTransactorSession) SetFeeDividend(tokenAddr common.Address, submitterCut *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetFeeDividend(&_Dospayment.TransactOpts, tokenAddr, submitterCut)
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
// Solidity: function setPaymentMethod(address payer, address tokenAddr) returns()
func (_Dospayment *DospaymentTransactor) SetPaymentMethod(opts *bind.TransactOpts, payer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setPaymentMethod", payer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(address payer, address tokenAddr) returns()
func (_Dospayment *DospaymentSession) SetPaymentMethod(payer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetPaymentMethod(&_Dospayment.TransactOpts, payer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(address payer, address tokenAddr) returns()
func (_Dospayment *DospaymentTransactorSession) SetPaymentMethod(payer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetPaymentMethod(&_Dospayment.TransactOpts, payer, tokenAddr)
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
	Payer       common.Address
	TokenAddr   common.Address
	RequestID   *big.Int
	ServiceType *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogChargeServiceFee is a free log retrieval operation binding the contract event 0xa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba.
//
// Solidity: event LogChargeServiceFee(address payer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
func (_Dospayment *DospaymentFilterer) FilterLogChargeServiceFee(opts *bind.FilterOpts) (*DospaymentLogChargeServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogChargeServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogChargeServiceFeeIterator{contract: _Dospayment.contract, event: "LogChargeServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogChargeServiceFee is a free log subscription operation binding the contract event 0xa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba.
//
// Solidity: event LogChargeServiceFee(address payer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
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

// ParseLogChargeServiceFee is a log parse operation binding the contract event 0xa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba.
//
// Solidity: event LogChargeServiceFee(address payer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
func (_Dospayment *DospaymentFilterer) ParseLogChargeServiceFee(log types.Log) (*DospaymentLogChargeServiceFee, error) {
	event := new(DospaymentLogChargeServiceFee)
	if err := _Dospayment.contract.UnpackLog(event, "LogChargeServiceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// ParseLogClaimGuardianFee is a log parse operation binding the contract event 0x47ad88344c408450ef0ccab93ed97dd83af7a27dedfaa205c0725cfc4ca819cd.
//
// Solidity: event LogClaimGuardianFee(address nodeAddr, address tokenAddr, uint256 feeForSubmitter, address sender)
func (_Dospayment *DospaymentFilterer) ParseLogClaimGuardianFee(log types.Log) (*DospaymentLogClaimGuardianFee, error) {
	event := new(DospaymentLogClaimGuardianFee)
	if err := _Dospayment.contract.UnpackLog(event, "LogClaimGuardianFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DospaymentLogRecordServiceFeeIterator is returned from FilterLogRecordServiceFee and is used to iterate over the raw logs and unpacked data for LogRecordServiceFee events raised by the Dospayment contract.
type DospaymentLogRecordServiceFeeIterator struct {
	Event *DospaymentLogRecordServiceFee // Event containing the contract specifics and raw log

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
func (it *DospaymentLogRecordServiceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentLogRecordServiceFee)
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
		it.Event = new(DospaymentLogRecordServiceFee)
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
func (it *DospaymentLogRecordServiceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentLogRecordServiceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentLogRecordServiceFee represents a LogRecordServiceFee event raised by the Dospayment contract.
type DospaymentLogRecordServiceFee struct {
	NodeAddr    common.Address
	TokenAddr   common.Address
	RequestID   *big.Int
	ServiceType *big.Int
	Fee         *big.Int
	IsSubmitter bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogRecordServiceFee is a free log retrieval operation binding the contract event 0x4758b94d44e129dcef9dc829628a55e921926b4383f3261f968ee8f9373571e7.
//
// Solidity: event LogRecordServiceFee(address nodeAddr, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee, bool isSubmitter)
func (_Dospayment *DospaymentFilterer) FilterLogRecordServiceFee(opts *bind.FilterOpts) (*DospaymentLogRecordServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogRecordServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogRecordServiceFeeIterator{contract: _Dospayment.contract, event: "LogRecordServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogRecordServiceFee is a free log subscription operation binding the contract event 0x4758b94d44e129dcef9dc829628a55e921926b4383f3261f968ee8f9373571e7.
//
// Solidity: event LogRecordServiceFee(address nodeAddr, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee, bool isSubmitter)
func (_Dospayment *DospaymentFilterer) WatchLogRecordServiceFee(opts *bind.WatchOpts, sink chan<- *DospaymentLogRecordServiceFee) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "LogRecordServiceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentLogRecordServiceFee)
				if err := _Dospayment.contract.UnpackLog(event, "LogRecordServiceFee", log); err != nil {
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

// ParseLogRecordServiceFee is a log parse operation binding the contract event 0x4758b94d44e129dcef9dc829628a55e921926b4383f3261f968ee8f9373571e7.
//
// Solidity: event LogRecordServiceFee(address nodeAddr, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee, bool isSubmitter)
func (_Dospayment *DospaymentFilterer) ParseLogRecordServiceFee(log types.Log) (*DospaymentLogRecordServiceFee, error) {
	event := new(DospaymentLogRecordServiceFee)
	if err := _Dospayment.contract.UnpackLog(event, "LogRecordServiceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Payer       common.Address
	TokenAddr   common.Address
	RequestID   *big.Int
	ServiceType *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogRefundServiceFee is a free log retrieval operation binding the contract event 0xde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa.
//
// Solidity: event LogRefundServiceFee(address payer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
func (_Dospayment *DospaymentFilterer) FilterLogRefundServiceFee(opts *bind.FilterOpts) (*DospaymentLogRefundServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogRefundServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogRefundServiceFeeIterator{contract: _Dospayment.contract, event: "LogRefundServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogRefundServiceFee is a free log subscription operation binding the contract event 0xde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa.
//
// Solidity: event LogRefundServiceFee(address payer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
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

// ParseLogRefundServiceFee is a log parse operation binding the contract event 0xde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa.
//
// Solidity: event LogRefundServiceFee(address payer, address tokenAddr, uint256 requestID, uint256 serviceType, uint256 fee)
func (_Dospayment *DospaymentFilterer) ParseLogRefundServiceFee(log types.Log) (*DospaymentLogRefundServiceFee, error) {
	event := new(DospaymentLogRefundServiceFee)
	if err := _Dospayment.contract.UnpackLog(event, "LogRefundServiceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DospaymentUpdatePaymentAdminIterator is returned from FilterUpdatePaymentAdmin and is used to iterate over the raw logs and unpacked data for UpdatePaymentAdmin events raised by the Dospayment contract.
type DospaymentUpdatePaymentAdminIterator struct {
	Event *DospaymentUpdatePaymentAdmin // Event containing the contract specifics and raw log

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
func (it *DospaymentUpdatePaymentAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DospaymentUpdatePaymentAdmin)
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
		it.Event = new(DospaymentUpdatePaymentAdmin)
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
func (it *DospaymentUpdatePaymentAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DospaymentUpdatePaymentAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DospaymentUpdatePaymentAdmin represents a UpdatePaymentAdmin event raised by the Dospayment contract.
type DospaymentUpdatePaymentAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatePaymentAdmin is a free log retrieval operation binding the contract event 0xe0be0cc9c5ea7d7a7f3909ad261ab0e0fbc9aae3fe819864f60d81d91286dc30.
//
// Solidity: event UpdatePaymentAdmin(address oldAdmin, address newAdmin)
func (_Dospayment *DospaymentFilterer) FilterUpdatePaymentAdmin(opts *bind.FilterOpts) (*DospaymentUpdatePaymentAdminIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "UpdatePaymentAdmin")
	if err != nil {
		return nil, err
	}
	return &DospaymentUpdatePaymentAdminIterator{contract: _Dospayment.contract, event: "UpdatePaymentAdmin", logs: logs, sub: sub}, nil
}

// WatchUpdatePaymentAdmin is a free log subscription operation binding the contract event 0xe0be0cc9c5ea7d7a7f3909ad261ab0e0fbc9aae3fe819864f60d81d91286dc30.
//
// Solidity: event UpdatePaymentAdmin(address oldAdmin, address newAdmin)
func (_Dospayment *DospaymentFilterer) WatchUpdatePaymentAdmin(opts *bind.WatchOpts, sink chan<- *DospaymentUpdatePaymentAdmin) (event.Subscription, error) {

	logs, sub, err := _Dospayment.contract.WatchLogs(opts, "UpdatePaymentAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DospaymentUpdatePaymentAdmin)
				if err := _Dospayment.contract.UnpackLog(event, "UpdatePaymentAdmin", log); err != nil {
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

// ParseUpdatePaymentAdmin is a log parse operation binding the contract event 0xe0be0cc9c5ea7d7a7f3909ad261ab0e0fbc9aae3fe819864f60d81d91286dc30.
//
// Solidity: event UpdatePaymentAdmin(address oldAdmin, address newAdmin)
func (_Dospayment *DospaymentFilterer) ParseUpdatePaymentAdmin(log types.Log) (*DospaymentUpdatePaymentAdmin, error) {
	event := new(DospaymentUpdatePaymentAdmin)
	if err := _Dospayment.contract.UnpackLog(event, "UpdatePaymentAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
