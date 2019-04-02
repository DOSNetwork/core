package onchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/DOSNetwork/core/share/vss/pedersen"
)

const (
	ETH = "ETH"
)

type ProxyAdapter interface {
	RegisterNewNode(ctx context.Context) (errc <-chan error)
	RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc <-chan error)
	SetRandomNum(ctx context.Context, signatures <-chan *vss.Signature) (errc <-chan error)
	DataReturn(ctx context.Context, signatures <-chan *vss.Signature) (errc <-chan error)
	SetGroupingThreshold(ctx context.Context, threshold uint64) (errc <-chan error)
	SetGroupToPick(ctx context.Context, groupToPick uint64) (errc <-chan error)
	SetGroupSize(ctx context.Context, size uint64) (errc <-chan error)
	SetGroupMaturityPeriod(ctx context.Context, size uint64) (errc <-chan error)
	Commit(ctx context.Context, commitment [32]byte) (errc <-chan error)
	Reveal(ctx context.Context, secret *big.Int) (errc <-chan error)

	//Guardian node functions
	SignalRandom(ctx context.Context) (errc <-chan error)
	SignalGroupFormation(ctx context.Context) (errc <-chan error)
	SignalDissolve(ctx context.Context, idx uint64) (errc <-chan error)

	SubscribeEvent(subscribeType int) (<-chan interface{}, <-chan error)
	PollLogs(subscribeType int, LogBlockDiff, preBlockBuf uint64) (<-chan interface{}, <-chan error)

	GetWorkingGroupSize() (size uint64, err error)
	GetGroupToPick() (size uint64, err error)
	LastUpdatedBlock() (blknum uint64, err error)
	CommitRevealTargetBlk() (blk uint64, err error)
	NumPendingGroups() (size uint64, err error)
	GetPengindNodeSize() (size uint64, err error)
	GetBalance() (balance *big.Float)
	Address() (addr []byte)
	CurrentBlock() (blknum uint64, err error)
	PendingNonce() (nonce uint64, err error)
	GroupPubKey(idx int) (groupPubKeys [4]*big.Int, err error)
}

func NewProxyAdapter(ChainType, credentialPath, passphrase, proxyAddr, crAddress string, urls []string) (ProxyAdapter, error) {
	switch ChainType {
	case ETH:
		adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, crAddress, urls)
		return adaptor, err
	default:
		err := fmt.Errorf("Chain %s not supported error\n", ChainType)
		return nil, err
	}
}
