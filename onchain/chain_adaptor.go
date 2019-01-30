package onchain

import (
	"fmt"
	"math/big"

	"github.com/DOSNetwork/core/configuration"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ETH = "ETH"
)

type ChainInterface interface {
	SetAccount(credentialPath string) (err error)
	Init(chainConfig configuration.ChainConfig) (err error)
	SubscribeEvent(ch chan interface{}, subscribeType int) (err error)
	InitialWhiteList() (err error)
	GetWhitelist() (address common.Address, err error)
	UploadID() (err error)
	UploadPubKey(pubKey [4]*big.Int) (err error)
	GetId() (id []byte)
	GetBlockHashByNumber(blknum *big.Int) (hash common.Hash, err error)
	SetRandomNum(sig []byte, version uint8) (err error)
	DataReturn(requestId *big.Int, trafficType uint8, data, sig []byte, version uint8) (err error)
	SubscribeToAll(msgChan chan interface{}) (err error)
	//For test
	ResetNodeIDs() (err error)
	RandomNumberTimeOut() (err error)
	EnoughBalance(address common.Address) (isEnough bool)
	WhitelistInitialized() (initialized bool, err error)
}

func AdaptTo(ChainType string) (conn ChainInterface, err error) {
	switch ChainType {
	case ETH:
		conn = &EthAdaptor{}
	default:
		err = fmt.Errorf("Chain %s not supported error\n", ChainType)
	}
	return conn, err
}
