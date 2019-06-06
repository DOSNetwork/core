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
	SetRandomNum(ctx context.Context, signatures chan *vss.Signature) (errc chan error)
	DataReturn(ctx context.Context, signatures chan *vss.Signature) (errc chan error)
	RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc chan error)
	SubscribeEvent(subscribeType int) (chan interface{}, chan error)

	SetGroupingThreshold(ctx context.Context, threshold uint64) (errc error)
	SetGroupToPick(ctx context.Context, groupToPick uint64) (errc error)
	SetGroupSize(ctx context.Context, size uint64) (errc error)
	SetGroupMaturityPeriod(ctx context.Context, size uint64) (errc error)
	Commit(ctx context.Context, cid *big.Int, commitment [32]byte) (errc error)
	Reveal(ctx context.Context, cid *big.Int, secret *big.Int) (errc error)
	//Guardian node functions
	RegisterNewNode(ctx context.Context) (err error)
	SignalRandom(ctx context.Context) (errc error)
	SignalGroupFormation(ctx context.Context) (errc error)
	SignalGroupDissolve(ctx context.Context) (errc error)
	SignalBootstrap(ctx context.Context, cid uint64) (errc error)

	GetExpiredWorkingGroupSize(ctx context.Context) (r uint64, err error)
	GroupSize(ctx context.Context) (r uint64, err error)
	GetWorkingGroupSize(ctx context.Context) (r uint64, err error)
	GroupToPick(ctx context.Context) (r uint64, err error)
	LastUpdatedBlock(ctx context.Context) (r uint64, err error)
	NumPendingGroups(ctx context.Context) (r uint64, err error)
	NumPendingNodes(ctx context.Context) (r uint64, err error)
	Balance(ctx context.Context) (balance *big.Float, err error)
	Address() (addr []byte)
	CurrentBlock(ctx context.Context) (r uint64, err error)
	PendingNonce(ctx context.Context) (r uint64, err error)
	RefreshSystemRandomHardLimit(ctx context.Context) (limit uint64, err error)
	GroupPubKey(ctx context.Context, idx int) (groupPubKeys [4]*big.Int, err error)
	IsPendingNode(ctx context.Context, id []byte) (bool, error)
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
