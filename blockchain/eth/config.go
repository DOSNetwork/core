package eth

import "github.com/ethereum/go-ethereum/common"

type netConfig struct {
	remoteNodeAddress  string
	contractAddressHex string
	contractAddress    common.Address
}

var (
	rinkebyNode = initial("https://rinkeby.infura.io", "0x566B0eA1247870C3204f4e3587d84D635b8De1aa")
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
