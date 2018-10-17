package eth

import "github.com/ethereum/go-ethereum/common"

type netConfig struct {
	remoteNodeAddress    string
	proxyContractAddress common.Address
}

var (
	rinkebyNode        = initial("wss://rinkeby.infura.io/ws", "0xAd67eE265945fe895aEA482d09762011C9D67442")
	rinkebyPrivateNode = initial("ws://13.56.31.73:8546", "0xAd67eE265945fe895aEA482d09762011C9D67442")
	privateNode        = initial("ws://13.52.16.14:8546", "0x19002E8b5D11076cC27bA26670f239B41443E60d")
)

func initial(remoteNodeAddress, proxyContractAddressHex string) (node *netConfig) {
	proxyContractAddress := common.HexToAddress(proxyContractAddressHex)
	node = &netConfig{
		remoteNodeAddress:    remoteNodeAddress,
		proxyContractAddress: proxyContractAddress,
	}
	return
}
