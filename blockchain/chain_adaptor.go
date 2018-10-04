package blockchain

import (
	"fmt"
	"math/big"

	"github.com/DOSNetwork/core/blockchain/eth"
)

type ChainInterface interface {
	Init(autoReplenish bool) (err error)
	SubscribeEvent(ch chan interface{}) (err error)
	UploadID() (err error)
	UploadPubKey(groupId, x0, x1, y0, y1 *big.Int) (err error)
	GetId() (id []byte)
	GetBootstrapIp() (ip string, err error)
	DataReturn(queryId *big.Int, data []byte, x, y *big.Int) (err error)
}

func AdaptTo(chainName string, autoReplenish bool) (conn ChainInterface, err error) {
	switch chainName {
	case "ETH":
		conn = &eth.EthAdaptor{}
		err = conn.Init(autoReplenish)
	default:
		err = fmt.Errorf("Chain %s not supported error\n", chainName)
	}
	return conn, err
}
