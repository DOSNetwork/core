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
const DospaymentABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardianFundsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogChargeServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeForSubmitter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"LogClaimGuardianFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSubmitter\",\"type\":\"bool\"}],\"name\":\"LogRecordServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"LogRefundServiceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"bridgeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"}],\"name\":\"chargeServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardianAddr\",\"type\":\"address\"}],\"name\":\"claimGuardianReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultGuardianFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultSubmitterCut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultSystemRandomFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultUserQueryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultUserRandomFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeLists\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"submitterCut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guardianFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"guardianFundsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"guardianFundsTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"}],\"name\":\"hasServiceFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_guardianFundsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nodeClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"nodeClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"nodeFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"nodeFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"paymentInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentMethods\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"workers\",\"type\":\"address[]\"}],\"name\":\"recordServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"refundServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"submitterCut\",\"type\":\"uint256\"}],\"name\":\"setFeeDividend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setGuardianFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fundsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setGuardianFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"setPaymentMethod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"serviceType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DospaymentBin is the compiled bytecode used for deploying new contracts.
const DospaymentBin = `608060405234801561001057600080fd5b50604051611d22380380611d228339818101604052606081101561003357600080fd5b5080516020820151604090920151600080546001600160a01b0319163317905590919061006a8383836001600160e01b0361007216565b505050610177565b6007546001600160a01b031615801561009457506008546001600160a01b0316155b6100e5576040805162461bcd60e51b815260206004820152601360248201527f616c72656164792d696e697469616c697a656400000000000000000000000000604482015290519081900360640190fd5b600780546001600160a01b03199081166001600160a01b0395861617909155600580548216938516939093179092556006805483169190931690811790925560088054909116821790556000908152600260208181526040808420848052918290528084206802b5e3af16b1880000908190556001808652828620829055848652919094208490556004908201550155565b611b9c806101866000396000f3fe608060405234801561001057600080fd5b50600436106101fb5760003560e01c80638403f7dc1161011a578063c0f14e46116100ad578063e36503661161007c578063e365036614610619578063eebede8314610645578063f2fde38b14610671578063f39a19bf14610697578063fa2c775e146106bd576101fb565b8063c0f14e46146105ac578063c60be4fd14610341578063cb7ca88c14610341578063d95eaa7a146105da576101fb565b806391874ef7116100e957806391874ef714610518578063a93eb11014610520578063b73a3f8f14610546578063c0c53b8b14610574576101fb565b80638403f7dc1461040157806387d81789146104b95780638da5cb5b146105085780638f32d59b14610510576101fb565b8063571028e3116101925780636dfa72b0116101615780636dfa72b014610341578063715018a6146103bf578063746c73fd146103c75780637aa9181b146103cf576101fb565b8063571028e31461035b5780635a1fa503146103635780636059775a14610391578063694732c614610399576101fb565b80632c097993116101ce5780632c097993146102a35780633157f16d146102cf5780633939c4011461030f5780634a0a382f14610341576101fb565b806302b8b587146102005780631efa5a981461022457806323ff34cb14610243578063240028e814610269575b600080fd5b6102086106c5565b604080516001600160a01b039092168252519081900360200190f35b6102416004803603602081101561023a57600080fd5b50356106d4565b005b6102416004803603602081101561025957600080fd5b50356001600160a01b0316610927565b61028f6004803603602081101561027f57600080fd5b50356001600160a01b0316610ba9565b604080519115158252519081900360200190f35b610241600480360360408110156102b957600080fd5b506001600160a01b038135169060200135610c5b565b6102ec600480360360208110156102e557600080fd5b5035610cdd565b604080516001600160a01b03909316835260208301919091528051918290030190f35b6102416004803603606081101561032557600080fd5b506001600160a01b038135169060208101359060400135610d04565b610349610d8d565b60408051918252519081900360200190f35b610349610d9a565b6102416004803603604081101561037957600080fd5b506001600160a01b0381358116916020013516610d9f565b610208610e35565b610208600480360360208110156103af57600080fd5b50356001600160a01b0316610e44565b610241610e5f565b610349610eb8565b610241600480360360608110156103e557600080fd5b506001600160a01b038135169060208101359060400135610ed8565b6102416004803603606081101561041757600080fd5b8135916001600160a01b036020820135169181019060608101604082013564010000000081111561044757600080fd5b82018360208201111561045957600080fd5b8035906020019184602083028401116401000000008311171561047b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506111dc945050505050565b6104d6600480360360208110156104cf57600080fd5b50356115dd565b604080516001600160a01b03958616815293909416602084015282840191909152606082015290519081900360800190f35b610208611613565b61028f611622565b610208611633565b6103496004803603602081101561053657600080fd5b50356001600160a01b0316611642565b6102416004803603604081101561055c57600080fd5b506001600160a01b0381358116916020013516611662565b6102416004803603606081101561058a57600080fd5b506001600160a01b0381358116916020810135821691604090910135166116e7565b610349600480360360408110156105c257600080fd5b506001600160a01b03813581169160200135166117e2565b610600600480360360208110156105f057600080fd5b50356001600160a01b031661180d565b6040805192835260208301919091528051918290030190f35b61028f6004803603604081101561062f57600080fd5b506001600160a01b038135169060200135611829565b6102416004803603604081101561065b57600080fd5b506001600160a01b03813516906020013561195e565b6102416004803603602081101561068757600080fd5b50356001600160a01b03166119e1565b610349600480360360208110156106ad57600080fd5b50356001600160a01b03166119fe565b610208611a19565b6008546001600160a01b031681565b6106dc611622565b6106e557600080fd5b60008181526003602081905260409091200154819061073b576040805162461bcd60e51b815260206004820152600d60248201526c1b9bcb5999594b585b5bdd5b9d609a1b604482015290519081900360640190fd5b6000818152600360205260409020546001600160a01b0316610794576040805162461bcd60e51b815260206004820152600d60248201526c6e6f2d70617965722d696e666f60981b604482015290519081900360640190fd5b6000818152600360205260409020600101546001600160a01b03166107f4576040805162461bcd60e51b81526020600482015260116024820152706e6f2d6665652d746f6b656e2d696e666f60781b604482015290519081900360640190fd5b6000828152600360208181526040808420928301805460028501805460018701805488546001600160a01b0319808216909a55988116909155918890559690925582516001600160a01b039586168082529590921693820184905281830188905260608201869052608082018190529151919493917fde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa9181900360a00190a1816001600160a01b031663a9059cbb82866040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156108f357600080fd5b505af1158015610907573d6000803e3d6000fd5b505050506040513d602081101561091d57600080fd5b5050505050505050565b600760009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561097557600080fd5b505afa158015610989573d6000803e3d6000fd5b505050506040513d602081101561099f57600080fd5b50516001600160a01b031633146109f2576040805162461bcd60e51b81526020600482015260126024820152716e6f742d66726f6d2d646f732d70726f787960701b604482015290519081900360640190fd5b6005546001600160a01b0316610a4f576040805162461bcd60e51b815260206004820152601c60248201527f6e6f742d76616c69642d677561726469616e2d66756e642d6164647200000000604482015290519081900360640190fd5b6006546001600160a01b0316610aac576040805162461bcd60e51b815260206004820152601d60248201527f6e6f742d76616c69642d677561726469616e2d746f6b656e2d61646472000000604482015290519081900360640190fd5b6006546001600160a01b039081166000818152600260208181526040928390209091015482519486168552908401929092528281018290523360608401525190917f47ad88344c408450ef0ccab93ed97dd83af7a27dedfaa205c0725cfc4ca819cd919081900360800190a1600654600554604080516323b872dd60e01b81526001600160a01b039283166004820152858316602482015260448101859052905191909216916323b872dd9160648083019260209291908290030181600087803b158015610b7957600080fd5b505af1158015610b8d573d6000803e3d6000fd5b505050506040513d6020811015610ba357600080fd5b50505050565b60006001600160a01b038216610bc157506000610c56565b6001600160a01b0382166000908152600260209081526040808320838052909152902054610bf157506000610c56565b6001600160a01b038216600090815260026020908152604080832060018452909152902054610c2257506000610c56565b6001600160a01b0382166000908152600260208181526040808420928452919052902054610c5257506000610c56565b5060015b919050565b610c63611622565b610c6c57600080fd5b6001600160a01b038216610cbe576040805162461bcd60e51b81526020600482015260146024820152733737ba16bb30b634b216ba37b5b2b716b0b2323960611b604482015290519081900360640190fd5b6001600160a01b03909116600090815260026020526040902060010155565b600090815260036020819052604090912060018101549101546001600160a01b0390911691565b610d0c611622565b610d1557600080fd5b6001600160a01b038316610d67576040805162461bcd60e51b81526020600482015260146024820152733737ba16bb30b634b216ba37b5b2b716b0b2323960611b604482015290519081900360640190fd5b6001600160a01b0390921660009081526002602090815260408083209383529290522055565b6802b5e3af16b188000081565b600481565b610da7611622565b610db057600080fd5b80610dba81610ba9565b610e06576040805162461bcd60e51b81526020600482015260186024820152773737ba16b9bab83837b93a32b216ba37b5b2b716b0b2323960411b604482015290519081900360640190fd5b50600580546001600160a01b039384166001600160a01b03199182161790915560068054929093169116179055565b6005546001600160a01b031681565b6001602052600090815260409020546001600160a01b031681565b610e67611622565b610e7057600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b600854600090610ed39033906001600160a01b031681611a28565b905090565b600760009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f2657600080fd5b505afa158015610f3a573d6000803e3d6000fd5b505050506040513d6020811015610f5057600080fd5b50516001600160a01b03163314610fa3576040805162461bcd60e51b81526020600482015260126024820152716e6f742d66726f6d2d646f732d70726f787960701b604482015290519081900360640190fd5b600060016000856001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03169050600060026000836001600160a01b03166001600160a01b0316815260200190815260200160002060000160008481526020019081526020016000205490506040518060800160405280866001600160a01b03168152602001836001600160a01b03168152602001848152602001828152506003600086815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160020155606082015181600301559050507fa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba858386868560405180866001600160a01b03166001600160a01b03168152602001856001600160a01b03166001600160a01b031681526020018481526020018381526020018281526020019550505050505060405180910390a1604080516323b872dd60e01b81526001600160a01b038781166004830152306024830152604482018490529151918416916323b872dd916064808201926020929091908290030181600087803b1580156111a957600080fd5b505af11580156111bd573d6000803e3d6000fd5b505050506040513d60208110156111d357600080fd5b50505050505050565b600760009054906101000a90046001600160a01b03166001600160a01b03166343a73d9a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561122a57600080fd5b505afa15801561123e573d6000803e3d6000fd5b505050506040513d602081101561125457600080fd5b50516001600160a01b031633146112a7576040805162461bcd60e51b81526020600482015260126024820152716e6f742d66726f6d2d646f732d70726f787960701b604482015290519081900360640190fd5b6000838152600360208190526040909120015483906112fd576040805162461bcd60e51b815260206004820152600d60248201526c1b9bcb5999594b585b5bdd5b9d609a1b604482015290519081900360640190fd5b6000818152600360205260409020546001600160a01b0316611356576040805162461bcd60e51b815260206004820152600d60248201526c6e6f2d70617965722d696e666f60981b604482015290519081900360640190fd5b6000818152600360205260409020600101546001600160a01b03166113b6576040805162461bcd60e51b81526020600482015260116024820152706e6f2d6665652d746f6b656e2d696e666f60781b604482015290519081900360640190fd5b60008481526003602081815260408084206001808201805495830180546002808601805487546001600160a01b0319908116909855968a1690945592899055908890556001600160a01b0396871680895291865284882080840154978c16808a5260048852868a20848b528852988690208054600a909304988902928301905585519889529588018290528785018c9052606088018490526080880181905260a08801929092529251929591939290917f4758b94d44e129dcef9dc829628a55e921926b4383f3261f968ee8f9373571e7919081900360c00190a1600060018851038360010154600a038602816114a957fe5b04905060005b88518110156115d057896001600160a01b03168982815181106114ce57fe5b60200260200101516001600160a01b0316146115c85781600460008b84815181106114f557fe5b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206000896001600160a01b03166001600160a01b03168152602001908152602001600020600082825401925050819055507f4758b94d44e129dcef9dc829628a55e921926b4383f3261f968ee8f9373571e789828151811061157a57fe5b602090810291909101810151604080516001600160a01b039283168152918b16928201929092528082018e90526060810188905260808101859052600060a082015290519081900360c00190a15b6001016114af565b5050505050505050505050565b600360208190526000918252604090912080546001820154600283015492909301546001600160a01b0391821693909116919084565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b6007546001600160a01b031681565b60085460009061165c9083906001600160a01b03166117e2565b92915050565b8061166c81610ba9565b6116b8576040805162461bcd60e51b81526020600482015260186024820152773737ba16b9bab83837b93a32b216ba37b5b2b716b0b2323960411b604482015290519081900360640190fd5b506001600160a01b03918216600090815260016020526040902080546001600160a01b03191691909216179055565b6007546001600160a01b031615801561170957506008546001600160a01b0316155b611750576040805162461bcd60e51b8152602060048201526013602482015272185b1c9958591e4b5a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b600780546001600160a01b03199081166001600160a01b0395861617909155600580548216938516939093179092556006805483169190931690811790925560088054909116821790556000908152600260208181526040808420848052918290528084206802b5e3af16b1880000908190556001808652828620829055848652919094208490556004908201550155565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205490565b6002602081905260009182526040909120600181015491015482565b600754604080516321d39ecd60e11b815290516000926001600160a01b0316916343a73d9a916004808301926020929190829003018186803b15801561186e57600080fd5b505afa158015611882573d6000803e3d6000fd5b505050506040513d602081101561189857600080fd5b50516001600160a01b03848116911614156118b55750600161165c565b6001600160a01b03808416600081815260016020908152604080832054909416808352600282528483208784528252918490205484516370a0823160e01b815260048101949094529351919392839285926370a082319260248082019391829003018186803b15801561192757600080fd5b505afa15801561193b573d6000803e3d6000fd5b505050506040513d602081101561195157600080fd5b5051101595945050505050565b611966611622565b61196f57600080fd5b6001600160a01b0382166119c1576040805162461bcd60e51b81526020600482015260146024820152733737ba16bb30b634b216ba37b5b2b716b0b2323960611b604482015290519081900360640190fd5b6001600160a01b0390911660009081526002602081905260409091200155565b6119e9611622565b6119f257600080fd5b6119fb81611af9565b50565b60085460009061165c9033906001600160a01b031684611a28565b6006546001600160a01b031681565b6001600160a01b0380841660009081526004602090815260408083209386168352929052908120548015611af1576001600160a01b038086166000908152600460208181526040808420898616808652908352818520859055815163a9059cbb60e01b8152958916938601939093526024850186905251919363a9059cbb93604480830194928390030190829087803b158015611ac457600080fd5b505af1158015611ad8573d6000803e3d6000fd5b505050506040513d6020811015611aee57600080fd5b50505b949350505050565b6001600160a01b038116611b0c57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a72315820a9682a362144a51c1a0b3df9fbf5804eb2db6fae6644637ca27681783b44829164736f6c63430005110032`

// DeployDospayment deploys a new Ethereum contract, binding an instance of Dospayment to it.
func DeployDospayment(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddr common.Address, _guardianFundsAddr common.Address, _tokenAddr common.Address) (common.Address, *types.Transaction, *Dospayment, error) {
	parsed, err := abi.JSON(strings.NewReader(DospaymentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DospaymentBin), backend, _bridgeAddr, _guardianFundsAddr, _tokenAddr)
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

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() constant returns(address)
func (_Dospayment *DospaymentCaller) BridgeAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "bridgeAddr")
	return *ret0, err
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() constant returns(address)
func (_Dospayment *DospaymentSession) BridgeAddr() (common.Address, error) {
	return _Dospayment.Contract.BridgeAddr(&_Dospayment.CallOpts)
}

// BridgeAddr is a free data retrieval call binding the contract method 0x91874ef7.
//
// Solidity: function bridgeAddr() constant returns(address)
func (_Dospayment *DospaymentCallerSession) BridgeAddr() (common.Address, error) {
	return _Dospayment.Contract.BridgeAddr(&_Dospayment.CallOpts)
}

// DefaultGuardianFee is a free data retrieval call binding the contract method 0x6dfa72b0.
//
// Solidity: function defaultGuardianFee() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultGuardianFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "defaultGuardianFee")
	return *ret0, err
}

// DefaultGuardianFee is a free data retrieval call binding the contract method 0x6dfa72b0.
//
// Solidity: function defaultGuardianFee() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultGuardianFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultGuardianFee(&_Dospayment.CallOpts)
}

// DefaultGuardianFee is a free data retrieval call binding the contract method 0x6dfa72b0.
//
// Solidity: function defaultGuardianFee() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultGuardianFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultGuardianFee(&_Dospayment.CallOpts)
}

// DefaultSubmitterCut is a free data retrieval call binding the contract method 0x571028e3.
//
// Solidity: function defaultSubmitterCut() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultSubmitterCut(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "defaultSubmitterCut")
	return *ret0, err
}

// DefaultSubmitterCut is a free data retrieval call binding the contract method 0x571028e3.
//
// Solidity: function defaultSubmitterCut() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultSubmitterCut() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSubmitterCut(&_Dospayment.CallOpts)
}

// DefaultSubmitterCut is a free data retrieval call binding the contract method 0x571028e3.
//
// Solidity: function defaultSubmitterCut() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultSubmitterCut() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSubmitterCut(&_Dospayment.CallOpts)
}

// DefaultSystemRandomFee is a free data retrieval call binding the contract method 0xcb7ca88c.
//
// Solidity: function defaultSystemRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultSystemRandomFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "defaultSystemRandomFee")
	return *ret0, err
}

// DefaultSystemRandomFee is a free data retrieval call binding the contract method 0xcb7ca88c.
//
// Solidity: function defaultSystemRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultSystemRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSystemRandomFee(&_Dospayment.CallOpts)
}

// DefaultSystemRandomFee is a free data retrieval call binding the contract method 0xcb7ca88c.
//
// Solidity: function defaultSystemRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultSystemRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultSystemRandomFee(&_Dospayment.CallOpts)
}

// DefaultTokenAddr is a free data retrieval call binding the contract method 0x02b8b587.
//
// Solidity: function defaultTokenAddr() constant returns(address)
func (_Dospayment *DospaymentCaller) DefaultTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "defaultTokenAddr")
	return *ret0, err
}

// DefaultTokenAddr is a free data retrieval call binding the contract method 0x02b8b587.
//
// Solidity: function defaultTokenAddr() constant returns(address)
func (_Dospayment *DospaymentSession) DefaultTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.DefaultTokenAddr(&_Dospayment.CallOpts)
}

// DefaultTokenAddr is a free data retrieval call binding the contract method 0x02b8b587.
//
// Solidity: function defaultTokenAddr() constant returns(address)
func (_Dospayment *DospaymentCallerSession) DefaultTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.DefaultTokenAddr(&_Dospayment.CallOpts)
}

// DefaultUserQueryFee is a free data retrieval call binding the contract method 0xc60be4fd.
//
// Solidity: function defaultUserQueryFee() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultUserQueryFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "defaultUserQueryFee")
	return *ret0, err
}

// DefaultUserQueryFee is a free data retrieval call binding the contract method 0xc60be4fd.
//
// Solidity: function defaultUserQueryFee() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultUserQueryFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserQueryFee(&_Dospayment.CallOpts)
}

// DefaultUserQueryFee is a free data retrieval call binding the contract method 0xc60be4fd.
//
// Solidity: function defaultUserQueryFee() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultUserQueryFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserQueryFee(&_Dospayment.CallOpts)
}

// DefaultUserRandomFee is a free data retrieval call binding the contract method 0x4a0a382f.
//
// Solidity: function defaultUserRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentCaller) DefaultUserRandomFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "defaultUserRandomFee")
	return *ret0, err
}

// DefaultUserRandomFee is a free data retrieval call binding the contract method 0x4a0a382f.
//
// Solidity: function defaultUserRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentSession) DefaultUserRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserRandomFee(&_Dospayment.CallOpts)
}

// DefaultUserRandomFee is a free data retrieval call binding the contract method 0x4a0a382f.
//
// Solidity: function defaultUserRandomFee() constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) DefaultUserRandomFee() (*big.Int, error) {
	return _Dospayment.Contract.DefaultUserRandomFee(&_Dospayment.CallOpts)
}

// FeeLists is a free data retrieval call binding the contract method 0xd95eaa7a.
//
// Solidity: function feeLists( address) constant returns(submitterCut uint256, guardianFee uint256)
func (_Dospayment *DospaymentCaller) FeeLists(opts *bind.CallOpts, arg0 common.Address) (struct {
	SubmitterCut *big.Int
	GuardianFee  *big.Int
}, error) {
	ret := new(struct {
		SubmitterCut *big.Int
		GuardianFee  *big.Int
	})
	out := ret
	err := _Dospayment.contract.Call(opts, out, "feeLists", arg0)
	return *ret, err
}

// FeeLists is a free data retrieval call binding the contract method 0xd95eaa7a.
//
// Solidity: function feeLists( address) constant returns(submitterCut uint256, guardianFee uint256)
func (_Dospayment *DospaymentSession) FeeLists(arg0 common.Address) (struct {
	SubmitterCut *big.Int
	GuardianFee  *big.Int
}, error) {
	return _Dospayment.Contract.FeeLists(&_Dospayment.CallOpts, arg0)
}

// FeeLists is a free data retrieval call binding the contract method 0xd95eaa7a.
//
// Solidity: function feeLists( address) constant returns(submitterCut uint256, guardianFee uint256)
func (_Dospayment *DospaymentCallerSession) FeeLists(arg0 common.Address) (struct {
	SubmitterCut *big.Int
	GuardianFee  *big.Int
}, error) {
	return _Dospayment.Contract.FeeLists(&_Dospayment.CallOpts, arg0)
}

// GuardianFundsAddr is a free data retrieval call binding the contract method 0x6059775a.
//
// Solidity: function guardianFundsAddr() constant returns(address)
func (_Dospayment *DospaymentCaller) GuardianFundsAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "guardianFundsAddr")
	return *ret0, err
}

// GuardianFundsAddr is a free data retrieval call binding the contract method 0x6059775a.
//
// Solidity: function guardianFundsAddr() constant returns(address)
func (_Dospayment *DospaymentSession) GuardianFundsAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsAddr(&_Dospayment.CallOpts)
}

// GuardianFundsAddr is a free data retrieval call binding the contract method 0x6059775a.
//
// Solidity: function guardianFundsAddr() constant returns(address)
func (_Dospayment *DospaymentCallerSession) GuardianFundsAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsAddr(&_Dospayment.CallOpts)
}

// GuardianFundsTokenAddr is a free data retrieval call binding the contract method 0xfa2c775e.
//
// Solidity: function guardianFundsTokenAddr() constant returns(address)
func (_Dospayment *DospaymentCaller) GuardianFundsTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "guardianFundsTokenAddr")
	return *ret0, err
}

// GuardianFundsTokenAddr is a free data retrieval call binding the contract method 0xfa2c775e.
//
// Solidity: function guardianFundsTokenAddr() constant returns(address)
func (_Dospayment *DospaymentSession) GuardianFundsTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsTokenAddr(&_Dospayment.CallOpts)
}

// GuardianFundsTokenAddr is a free data retrieval call binding the contract method 0xfa2c775e.
//
// Solidity: function guardianFundsTokenAddr() constant returns(address)
func (_Dospayment *DospaymentCallerSession) GuardianFundsTokenAddr() (common.Address, error) {
	return _Dospayment.Contract.GuardianFundsTokenAddr(&_Dospayment.CallOpts)
}

// HasServiceFee is a free data retrieval call binding the contract method 0xe3650366.
//
// Solidity: function hasServiceFee(payer address, serviceType uint256) constant returns(bool)
func (_Dospayment *DospaymentCaller) HasServiceFee(opts *bind.CallOpts, payer common.Address, serviceType *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "hasServiceFee", payer, serviceType)
	return *ret0, err
}

// HasServiceFee is a free data retrieval call binding the contract method 0xe3650366.
//
// Solidity: function hasServiceFee(payer address, serviceType uint256) constant returns(bool)
func (_Dospayment *DospaymentSession) HasServiceFee(payer common.Address, serviceType *big.Int) (bool, error) {
	return _Dospayment.Contract.HasServiceFee(&_Dospayment.CallOpts, payer, serviceType)
}

// HasServiceFee is a free data retrieval call binding the contract method 0xe3650366.
//
// Solidity: function hasServiceFee(payer address, serviceType uint256) constant returns(bool)
func (_Dospayment *DospaymentCallerSession) HasServiceFee(payer common.Address, serviceType *big.Int) (bool, error) {
	return _Dospayment.Contract.HasServiceFee(&_Dospayment.CallOpts, payer, serviceType)
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

// NodeFeeBalance is a free data retrieval call binding the contract method 0xc0f14e46.
//
// Solidity: function nodeFeeBalance(nodeAddr address, tokenAddr address) constant returns(uint256)
func (_Dospayment *DospaymentCaller) NodeFeeBalance(opts *bind.CallOpts, nodeAddr common.Address, tokenAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "nodeFeeBalance", nodeAddr, tokenAddr)
	return *ret0, err
}

// NodeFeeBalance is a free data retrieval call binding the contract method 0xc0f14e46.
//
// Solidity: function nodeFeeBalance(nodeAddr address, tokenAddr address) constant returns(uint256)
func (_Dospayment *DospaymentSession) NodeFeeBalance(nodeAddr common.Address, tokenAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeFeeBalance(&_Dospayment.CallOpts, nodeAddr, tokenAddr)
}

// NodeFeeBalance is a free data retrieval call binding the contract method 0xc0f14e46.
//
// Solidity: function nodeFeeBalance(nodeAddr address, tokenAddr address) constant returns(uint256)
func (_Dospayment *DospaymentCallerSession) NodeFeeBalance(nodeAddr common.Address, tokenAddr common.Address) (*big.Int, error) {
	return _Dospayment.Contract.NodeFeeBalance(&_Dospayment.CallOpts, nodeAddr, tokenAddr)
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

// PaymentMethods is a free data retrieval call binding the contract method 0x694732c6.
//
// Solidity: function paymentMethods( address) constant returns(address)
func (_Dospayment *DospaymentCaller) PaymentMethods(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dospayment.contract.Call(opts, out, "paymentMethods", arg0)
	return *ret0, err
}

// PaymentMethods is a free data retrieval call binding the contract method 0x694732c6.
//
// Solidity: function paymentMethods( address) constant returns(address)
func (_Dospayment *DospaymentSession) PaymentMethods(arg0 common.Address) (common.Address, error) {
	return _Dospayment.Contract.PaymentMethods(&_Dospayment.CallOpts, arg0)
}

// PaymentMethods is a free data retrieval call binding the contract method 0x694732c6.
//
// Solidity: function paymentMethods( address) constant returns(address)
func (_Dospayment *DospaymentCallerSession) PaymentMethods(arg0 common.Address) (common.Address, error) {
	return _Dospayment.Contract.PaymentMethods(&_Dospayment.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments( uint256) constant returns(payer address, tokenAddr address, serviceType uint256, amount uint256)
func (_Dospayment *DospaymentCaller) Payments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Payer       common.Address
	TokenAddr   common.Address
	ServiceType *big.Int
	Amount      *big.Int
}, error) {
	ret := new(struct {
		Payer       common.Address
		TokenAddr   common.Address
		ServiceType *big.Int
		Amount      *big.Int
	})
	out := ret
	err := _Dospayment.contract.Call(opts, out, "payments", arg0)
	return *ret, err
}

// Payments is a free data retrieval call binding the contract method 0x87d81789.
//
// Solidity: function payments( uint256) constant returns(payer address, tokenAddr address, serviceType uint256, amount uint256)
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
// Solidity: function payments( uint256) constant returns(payer address, tokenAddr address, serviceType uint256, amount uint256)
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
// Solidity: function chargeServiceFee(payer address, requestID uint256, serviceType uint256) returns()
func (_Dospayment *DospaymentTransactor) ChargeServiceFee(opts *bind.TransactOpts, payer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "chargeServiceFee", payer, requestID, serviceType)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(payer address, requestID uint256, serviceType uint256) returns()
func (_Dospayment *DospaymentSession) ChargeServiceFee(payer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.ChargeServiceFee(&_Dospayment.TransactOpts, payer, requestID, serviceType)
}

// ChargeServiceFee is a paid mutator transaction binding the contract method 0x7aa9181b.
//
// Solidity: function chargeServiceFee(payer address, requestID uint256, serviceType uint256) returns()
func (_Dospayment *DospaymentTransactorSession) ChargeServiceFee(payer common.Address, requestID *big.Int, serviceType *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.ChargeServiceFee(&_Dospayment.TransactOpts, payer, requestID, serviceType)
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

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(_bridgeAddr address, _guardianFundsAddr address, _tokenAddr address) returns()
func (_Dospayment *DospaymentTransactor) Initialize(opts *bind.TransactOpts, _bridgeAddr common.Address, _guardianFundsAddr common.Address, _tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "initialize", _bridgeAddr, _guardianFundsAddr, _tokenAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(_bridgeAddr address, _guardianFundsAddr address, _tokenAddr address) returns()
func (_Dospayment *DospaymentSession) Initialize(_bridgeAddr common.Address, _guardianFundsAddr common.Address, _tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.Initialize(&_Dospayment.TransactOpts, _bridgeAddr, _guardianFundsAddr, _tokenAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(_bridgeAddr address, _guardianFundsAddr address, _tokenAddr address) returns()
func (_Dospayment *DospaymentTransactorSession) Initialize(_bridgeAddr common.Address, _guardianFundsAddr common.Address, _tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.Initialize(&_Dospayment.TransactOpts, _bridgeAddr, _guardianFundsAddr, _tokenAddr)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xf39a19bf.
//
// Solidity: function nodeClaim(to address) returns(uint256)
func (_Dospayment *DospaymentTransactor) NodeClaim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "nodeClaim", to)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xf39a19bf.
//
// Solidity: function nodeClaim(to address) returns(uint256)
func (_Dospayment *DospaymentSession) NodeClaim(to common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.NodeClaim(&_Dospayment.TransactOpts, to)
}

// NodeClaim is a paid mutator transaction binding the contract method 0xf39a19bf.
//
// Solidity: function nodeClaim(to address) returns(uint256)
func (_Dospayment *DospaymentTransactorSession) NodeClaim(to common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.NodeClaim(&_Dospayment.TransactOpts, to)
}

// RecordServiceFee is a paid mutator transaction binding the contract method 0x8403f7dc.
//
// Solidity: function recordServiceFee(requestID uint256, submitter address, workers address[]) returns()
func (_Dospayment *DospaymentTransactor) RecordServiceFee(opts *bind.TransactOpts, requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "recordServiceFee", requestID, submitter, workers)
}

// RecordServiceFee is a paid mutator transaction binding the contract method 0x8403f7dc.
//
// Solidity: function recordServiceFee(requestID uint256, submitter address, workers address[]) returns()
func (_Dospayment *DospaymentSession) RecordServiceFee(requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.RecordServiceFee(&_Dospayment.TransactOpts, requestID, submitter, workers)
}

// RecordServiceFee is a paid mutator transaction binding the contract method 0x8403f7dc.
//
// Solidity: function recordServiceFee(requestID uint256, submitter address, workers address[]) returns()
func (_Dospayment *DospaymentTransactorSession) RecordServiceFee(requestID *big.Int, submitter common.Address, workers []common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.RecordServiceFee(&_Dospayment.TransactOpts, requestID, submitter, workers)
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

// SetFeeDividend is a paid mutator transaction binding the contract method 0x2c097993.
//
// Solidity: function setFeeDividend(tokenAddr address, submitterCut uint256) returns()
func (_Dospayment *DospaymentTransactor) SetFeeDividend(opts *bind.TransactOpts, tokenAddr common.Address, submitterCut *big.Int) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setFeeDividend", tokenAddr, submitterCut)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0x2c097993.
//
// Solidity: function setFeeDividend(tokenAddr address, submitterCut uint256) returns()
func (_Dospayment *DospaymentSession) SetFeeDividend(tokenAddr common.Address, submitterCut *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetFeeDividend(&_Dospayment.TransactOpts, tokenAddr, submitterCut)
}

// SetFeeDividend is a paid mutator transaction binding the contract method 0x2c097993.
//
// Solidity: function setFeeDividend(tokenAddr address, submitterCut uint256) returns()
func (_Dospayment *DospaymentTransactorSession) SetFeeDividend(tokenAddr common.Address, submitterCut *big.Int) (*types.Transaction, error) {
	return _Dospayment.Contract.SetFeeDividend(&_Dospayment.TransactOpts, tokenAddr, submitterCut)
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

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(payer address, tokenAddr address) returns()
func (_Dospayment *DospaymentTransactor) SetPaymentMethod(opts *bind.TransactOpts, payer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.contract.Transact(opts, "setPaymentMethod", payer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(payer address, tokenAddr address) returns()
func (_Dospayment *DospaymentSession) SetPaymentMethod(payer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetPaymentMethod(&_Dospayment.TransactOpts, payer, tokenAddr)
}

// SetPaymentMethod is a paid mutator transaction binding the contract method 0xb73a3f8f.
//
// Solidity: function setPaymentMethod(payer address, tokenAddr address) returns()
func (_Dospayment *DospaymentTransactorSession) SetPaymentMethod(payer common.Address, tokenAddr common.Address) (*types.Transaction, error) {
	return _Dospayment.Contract.SetPaymentMethod(&_Dospayment.TransactOpts, payer, tokenAddr)
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
// Solidity: e LogChargeServiceFee(payer address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256)
func (_Dospayment *DospaymentFilterer) FilterLogChargeServiceFee(opts *bind.FilterOpts) (*DospaymentLogChargeServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogChargeServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogChargeServiceFeeIterator{contract: _Dospayment.contract, event: "LogChargeServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogChargeServiceFee is a free log subscription operation binding the contract event 0xa94e9ce5d0a7b76275efad947367b7999d9900f23bec1377d98f522ecad1b7ba.
//
// Solidity: e LogChargeServiceFee(payer address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256)
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
// Solidity: e LogRecordServiceFee(nodeAddr address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256, isSubmitter bool)
func (_Dospayment *DospaymentFilterer) FilterLogRecordServiceFee(opts *bind.FilterOpts) (*DospaymentLogRecordServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogRecordServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogRecordServiceFeeIterator{contract: _Dospayment.contract, event: "LogRecordServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogRecordServiceFee is a free log subscription operation binding the contract event 0x4758b94d44e129dcef9dc829628a55e921926b4383f3261f968ee8f9373571e7.
//
// Solidity: e LogRecordServiceFee(nodeAddr address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256, isSubmitter bool)
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
// Solidity: e LogRefundServiceFee(payer address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256)
func (_Dospayment *DospaymentFilterer) FilterLogRefundServiceFee(opts *bind.FilterOpts) (*DospaymentLogRefundServiceFeeIterator, error) {

	logs, sub, err := _Dospayment.contract.FilterLogs(opts, "LogRefundServiceFee")
	if err != nil {
		return nil, err
	}
	return &DospaymentLogRefundServiceFeeIterator{contract: _Dospayment.contract, event: "LogRefundServiceFee", logs: logs, sub: sub}, nil
}

// WatchLogRefundServiceFee is a free log subscription operation binding the contract event 0xde0a5183bfc8c743f7b95ecaaf7815e8f82d8ae05ca1ade1eac1ff9d961a2eaa.
//
// Solidity: e LogRefundServiceFee(payer address, tokenAddr address, requestID uint256, serviceType uint256, fee uint256)
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
