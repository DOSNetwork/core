package eth

import "github.com/ethereum/go-ethereum/common"

type netConfig struct {
	remoteNodeAddress    string
	proxyContractAddress common.Address
}

var (
	rinkebyNode        = initial("wss://rinkeby.infura.io/ws", "0x8B510350D6cC2E067dfAd3f1E7F5E36863ac36a8")
	rinkebyPrivateNode = initial("ws://13.56.31.73:8546", "0x8B510350D6cC2E067dfAd3f1E7F5E36863ac36a8")
	privateNode        = initial("ws://13.52.16.14:8546", "0xA5a1A5F848B13Aa64ba3bdF5777fF787e4105424")
)

func initial(remoteNodeAddress, proxyContractAddressHex string) (node *netConfig) {
	proxyContractAddress := common.HexToAddress(proxyContractAddressHex)
	node = &netConfig{
		remoteNodeAddress:    remoteNodeAddress,
		proxyContractAddress: proxyContractAddress,
	}
	return
}
