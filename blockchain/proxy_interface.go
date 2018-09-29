package blockchain

import (
	"math/big"

	"github.com/DOSNetwork/core/blockchain/eth"
)

type onChainInterface interface {
	Init() (err error)
	SubscribeEvent(ch chan interface{}) (err error)
	UploadPubKey(groupId, x0, x1, y0, y1 *big.Int) (err error)
	DataReturn(queryId *big.Int, data []byte, x, y *big.Int) (err error)
}

func CreateOnChainConn(chainName string) (conn onChainInterface, err error) {
	if chainName == "Eth" {
		conn = &eth.EthConn{}
		conn.Init()
	}
	return
}