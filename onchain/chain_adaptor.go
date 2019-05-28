package onchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/DOSNetwork/core/share/vss/pedersen"
)

const (
	//ETH represents the type of blockchain
	ETH = "ETH"
)

//ProxyAdapter represents an unified adapter interface for different blockchain
type ProxyAdapter interface {
	RegisterNewNode(ctx context.Context) (errc <-chan error)
	AddEventNode() (err error)
	RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc <-chan error)
	SetRandomNum(ctx context.Context, signatures <-chan *vss.Signature) (errc <-chan error)
	DataReturn(ctx context.Context, signatures <-chan *vss.Signature) (errc <-chan error)
	SetGroupingThreshold(ctx context.Context, threshold uint64) (errc <-chan error)
	SetGroupToPick(ctx context.Context, groupToPick uint64) (errc <-chan error)
	SetGroupSize(ctx context.Context, size uint64) (errc <-chan error)
	SetGroupMaturityPeriod(ctx context.Context, size uint64) (errc <-chan error)
	Commit(ctx context.Context, cid *big.Int, commitment [32]byte) (errc <-chan error)
	Reveal(ctx context.Context, cid *big.Int, secret *big.Int) (errc <-chan error)

	//Guardian node functions
	SignalRandom(ctx context.Context) (errc <-chan error)
	SignalGroupFormation(ctx context.Context) (errc <-chan error)
	SignalDissolve(ctx context.Context) (errc <-chan error)
	SignalBootstrap(ctx context.Context, cid uint64) (errc <-chan error)

	SubscribeEvent(subscribeType int) (<-chan interface{}, <-chan error)
	PollLogs(subscribeType int, LogBlockDiff, preBlockBuf uint64) (<-chan interface{}, <-chan error)

	GetExpiredWorkingGroupSize() (size uint64, err error)
	GroupSize() (size uint64, err error)
	GetWorkingGroupSize() (size uint64, err error)
	GetGroupToPick() (size uint64, err error)
	LastUpdatedBlock() (blknum uint64, err error)
	NumPendingGroups() (size uint64, err error)
	GetPengindNodeSize() (size uint64, err error)
	GetBalance() (balance *big.Float)
	Address() (addr []byte)
	CurrentBlock() (blknum uint64, err error)
	PendingNonce() (nonce uint64, err error)
	GroupPubKey(idx int) (groupPubKeys [4]*big.Int, err error)
	IsPendingNode(id []byte) (bool, error)
}

//NewProxyAdapter constructs a new ProxyAdapter with the given type of blockchain and contract addresses
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
