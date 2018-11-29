package onchain

import (
	"fmt"
	"math/big"

	"github.com/DOSNetwork/core/configuration"

	"github.com/dedis/kyber"

	"github.com/ethereum/go-ethereum/common"

	"github.com/sirupsen/logrus"
)

const (
	ETH = "ETH"
)

var log *logrus.Logger

type ChainInterface interface {
	Init(chainConfig *configuration.ChainConfig) (err error)
	SubscribeEvent(ch chan interface{}, subscribeType int) (err error)
	InitialWhiteList() (err error)
	GetWhitelist() (address common.Address, err error)
	UploadID() (err error)
	UploadPubKey(pubKey kyber.Point) (err error)
	GetId() (id []byte)
	GetBlockHashByNumber(blknum *big.Int) (hash common.Hash, err error)
	SetRandomNum(sig []byte) (err error)
	DataReturn(requestId *big.Int, trafficType uint8, data, sig []byte, version uint8) (err error)
	SubscribeToAll(msgChan chan interface{}) (err error)
	//For test
	ResetNodeIDs() (err error)
	RandomNumberTimeOut() (err error)
}

func AdaptTo(chainName string, chainConfig *configuration.ChainConfig, logger *logrus.Logger) (conn ChainInterface, err error) {
	log = logger
	switch chainConfig.ChainType {
	case ETH:
		conn = &EthAdaptor{}
		err = conn.Init(chainConfig)
	default:
		err = fmt.Errorf("Chain %s not supported error\n", chainName)
	}
	return conn, err
}
