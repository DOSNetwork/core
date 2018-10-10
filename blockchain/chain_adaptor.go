package blockchain

import (
	"fmt"
	"github.com/dedis/kyber"
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	"github.com/DOSNetwork/core/blockchain/eth"
)

type ChainInterface interface {
	Init(autoReplenish bool) (err error)
	SubscribeEvent(ch chan interface{}) (err error)
	UploadID() (err error)
	UploadPubKey(groupId *big.Int, pubKey kyber.Point) (err error)
	GetId() (id []byte)
	GetCurrBlockHash() (hash common.Hash, err error)
	GetBootstrapIp() (ip string, err error)
	SetBootstrapIp(ip string) (err error)
	GetRandomNum() (num *big.Int, err error)
	SetRandomNum(num *big.Int, sig []byte) (err error)
	DataReturn(queryId *big.Int, data, sig []byte) (err error)
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
