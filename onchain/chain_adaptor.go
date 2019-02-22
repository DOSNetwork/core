package onchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/share/vss/pedersen"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ETH = "ETH"
)

type ChainInterface interface {
	SetAccount(credentialPath, passphrase string) (err error)
	Init(chainConfig configuration.ChainConfig) (err error)
	SubscribeEvent(ch chan interface{}, subscribeType int) (err error)
	InitialWhiteList() (err error)
	GetWhitelist() (address common.Address, err error)
	UploadID() (err error)
	UploadPubKey(ctx context.Context, pubKey chan [5]*big.Int) <-chan error
	GetId() (id []byte)
	GetBlockHashByNumber(blknum *big.Int) (hash common.Hash, err error)
	GetLastRandomness() (*big.Int, error)
	GetLastUpdatedBlock() (uint64, error)
	GetCurrentBlock() (uint64, error)
	SetRandomNum(ctx context.Context, signatures <-chan *vss.Signature) <-chan error
	DataReturn(ctx context.Context, signatures <-chan *vss.Signature) <-chan error
	SubscribeToAll(msgChan chan interface{}) (err error)
	BalanceMaintain(rootCredentialPath string) (err error)
	//For test
	ResetNodeIDs() (err error)
	RandomNumberTimeOut() (err error)
	EnoughBalance(address common.Address) (isEnough bool)
	GetBalance() (balance *big.Float)
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
