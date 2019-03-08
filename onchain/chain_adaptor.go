package onchain

import (
	"context"
	"fmt"
	"math/big"

	//	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/share/vss/pedersen"
)

const (
	ETH = "ETH"
)

type ProxyAdapter interface {
	RegisterNewNode(ctx context.Context) (errc <-chan error)
	RandomNumberTimeOut(ctx context.Context) (errc <-chan error)
	RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc chan error)
	SetRandomNum(ctx context.Context, signatures <-chan *vss.Signature) (errc chan error)
	DataReturn(ctx context.Context, signatures <-chan *vss.Signature) (errc chan error)

	SubscribeEvent(subscribeType int, sink chan interface{})
	PollLogs(subscribeType int, LogBlockDiff, preBlockBuf uint64) (chan interface{}, <-chan error)

	GetWorkingGroupSize() (size uint64)
	LastUpdatedBlock() (blknum uint64, err error)
	GroupPubKey(idx int) (groupPubKeys [4]*big.Int, err error)

	GetBalance() (balance *big.Float)
	Address() (addr []byte)
	CurrentBlock() (blknum uint64, err error)
}

func NewProxyAdapter(ChainType, credentialPath, passphrase, proxyAddr string, urls []string) (ProxyAdapter, error) {
	switch ChainType {
	case ETH:
		adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, urls)
		return adaptor, err
	default:
		err := fmt.Errorf("Chain %s not supported error\n", ChainType)
		return nil, err
	}
}
