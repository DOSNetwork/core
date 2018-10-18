package eth

import "github.com/ethereum/go-ethereum/common"

type netConfig struct {
	remoteNodeAddress    string
	proxyContractAddress common.Address
}

var (
	rinkebyNode        = initial("wss://rinkeby.infura.io/ws", "0xAd67eE265945fe895aEA482d09762011C9D67442")
	rinkebyPrivateNode = initial("ws://13.56.31.73:8546", "0xAd67eE265945fe895aEA482d09762011C9D67442")
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
