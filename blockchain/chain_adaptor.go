package blockchain

import (
	"fmt"
	"math/big"

	"github.com/dedis/kyber"
	"github.com/ethereum/go-ethereum/common"

	"github.com/DOSNetwork/core/blockchain/eth"
)

type ChainInterface interface {
	Init(autoReplenish bool) (err error)
	SubscribeEvent(ch chan interface{}) (err error)
	UploadID() (err error)
	UploadPubKey(pubKey kyber.Point) (err error)
	GetId() (id []byte)
	GetBlockHashByNumber(blknum *big.Int) (hash common.Hash, err error)
	SetRandomNum(sig []byte) (err error)
	DataReturn(queryId *big.Int, data, sig []byte) (err error)
	//For test
	ResetNodeIDs()
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
