package blockchain

import (
	"fmt"
	"math/big"

	"github.com/dedis/kyber"
	"github.com/ethereum/go-ethereum/common"

	"github.com/DOSNetwork/core/blockchain/eth"
)

const (
	ETH = iota
)

type ChainInterface interface {
	Init(autoReplenish bool, netType int) (err error)
	SubscribeEvent(ch chan interface{}, subscribeType int) (err error)
	UploadID() (err error)
	UploadPubKey(pubKey kyber.Point) (err error)
	GetId() (id []byte)
	GetBlockHashByNumber(blknum *big.Int) (hash common.Hash, err error)
	SetRandomNum(content, sig []byte) (err error)
	DataReturn(queryId *big.Int, data, sig []byte) (err error)
	DeployContract(contractName int) (address common.Address, err error)
	DeployAll() (proxyAddress, bridgeAddress, askAddress common.Address, err error)
	SubscribeToAll(msgChan chan interface{}) (err error)
	//For test
	ResetNodeIDs() (err error)
	RandomNumberTimeOut() (err error)
}

func AdaptTo(chainName int, autoReplenish bool, netType int) (conn ChainInterface, err error) {
	switch chainName {
	case ETH:
		conn = &eth.EthAdaptor{}
		err = conn.Init(autoReplenish, netType)
	default:
		err = fmt.Errorf("Chain %s not supported error\n", chainName)
	}
	return conn, err
}
