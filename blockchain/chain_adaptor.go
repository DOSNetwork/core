package blockchain

import (
	"fmt"
	"math/big"

	"github.com/DOSNetwork/core/blockchain/eth"
)

type ChainInterface interface {
	Init() (err error)
	SubscribeEvent(ch chan interface{}) (err error)
	UploadPubKey(groupId, x0, x1, y0, y1 *big.Int) (err error)
	DataReturn(queryId *big.Int, data []byte, x, y *big.Int) (err error)
}

func AdaptTo(chainName string) (conn ChainInterface, err error) {
	switch chainName {
	case "ETH":
		conn = &eth.EthConn{}
		err = conn.Init()
	default:
		err = fmt.Errorf("Chain %s not supported error\n", chainName)
	}
	return conn, err
}
