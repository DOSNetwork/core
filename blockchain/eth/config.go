package eth

import "github.com/ethereum/go-ethereum/common"

type netConfig struct {
	remoteNodeAddress  string
	contractAddressHex string
	contractAddress    common.Address
}

var (
	rinkebyNode        = initial("wss://rinkeby.infura.io/ws", "0xe1F4F37E193F4c993E754672768E688A477E2b2e")
	rinkebyPrivateNode = initial("ws://54.183.214.48:8546", "0xe1F4F37E193F4c993E754672768E688A477E2b2e")
	privateNode        = initial("ws://54.183.214.48:8546", "0xE40977C2D05B11628A7A16cAC4a0ABf616237fc1")
)

func initial(remoteNodeAddress, contractAddressHex string) (node *netConfig) {
	contractAddress := common.HexToAddress(contractAddressHex)
	node = &netConfig{
		remoteNodeAddress:  remoteNodeAddress,
		contractAddressHex: contractAddressHex,
		contractAddress:    contractAddress,
	}
	return
}
