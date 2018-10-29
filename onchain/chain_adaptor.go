package onchain

import (
	"fmt"
	"math/big"

	"github.com/dedis/kyber"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ETH = "ETH"
)

// TODO: Move to configuration/config.go
type ChainConfig struct {
	// TODO: Refactor out of ChainConfig
	RemoteNodeType    string
	RemoteNodeAddress string
	// TODO: remove unrelated field
	AskMeAnythingAddress string

	ChainType               string
	DOSProxyAddress         string
	DOSAddressBridgeAddress string
	DOSOnChainSDKAddress    string
}

type ChainInterface interface {
	Init(autoReplenish bool, chainConfig *ChainConfig) (err error)
	SubscribeEvent(ch chan interface{}, subscribeType int) (err error)
	InitialWhiteList() (err error)
	GetWhitelist() (address common.Address, err error)
	UploadID() (err error)
	UploadPubKey(pubKey kyber.Point) (err error)
	GetId() (id []byte)
	GetBlockHashByNumber(blknum *big.Int) (hash common.Hash, err error)
	SetRandomNum(sig []byte) (err error)
	DataReturn(requestId *big.Int, trafficType uint8, data, sig []byte) (err error)
	SubscribeToAll(msgChan chan interface{}) (err error)
	//For test
	ResetNodeIDs() (err error)
	RandomNumberTimeOut() (err error)
}

func AdaptTo(chainName string, autoReplenish bool, chainConfig *ChainConfig) (conn ChainInterface, err error) {
	switch chainConfig.ChainType {
	case ETH:
		conn = &EthAdaptor{}
		err = conn.Init(autoReplenish, chainConfig)
	default:
		err = fmt.Errorf("Chain %s not supported error\n", chainName)
	}
	return conn, err
}
