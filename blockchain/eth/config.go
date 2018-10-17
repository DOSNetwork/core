package eth

import "github.com/ethereum/go-ethereum/common"

type netConfig struct {
	remoteNodeAddress    string
	proxyContractAddress common.Address
}

var (
	rinkebyNode        = initial("wss://rinkeby.infura.io/ws", "0xe1F4F37E193F4c993E754672768E688A477E2b2e")
	rinkebyPrivateNode = initial("ws://13.56.31.73:8546", "0xe1F4F37E193F4c993E754672768E688A477E2b2e")
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
