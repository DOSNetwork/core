package eth

import "github.com/ethereum/go-ethereum/common"

type netConfig struct {
	remoteNodeAddress  string
	contractAddressHex string
	contractAddress    common.Address
}

var (
	rinkebyNode = initial("wss://rinkeby.infura.io/ws", "0xe1F4F37E193F4c993E754672768E688A477E2b2e")
	privateNode = initial("ws://67.207.98.117:8546", "0xf59c7469b3668d0676DEDDb06B46E799ec2109ce")
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
