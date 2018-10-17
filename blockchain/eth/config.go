package eth

import "github.com/ethereum/go-ethereum/common"

type netConfig struct {
	remoteNodeAddress    string
	proxyContractAddress common.Address
}

var (
	rinkebyNode        = initial("wss://rinkeby.infura.io/ws", "0xe1F4F37E193F4c993E754672768E688A477E2b2e")
	rinkebyPrivateNode = initial("ws://13.56.31.73:8546", "0xe1F4F37E193F4c993E754672768E688A477E2b2e")
	privateNode        = initial("ws://13.56.31.73:8546", "0xF75Fad8Deb6d1526FB7589bfa676db42574bd6Fe")
)

func initial(remoteNodeAddress, proxyContractAddressHex string) (node *netConfig) {
	proxyContractAddress := common.HexToAddress(proxyContractAddressHex)
	node = &netConfig{
		remoteNodeAddress:    remoteNodeAddress,
		proxyContractAddress: proxyContractAddress,
	}
	return
}
