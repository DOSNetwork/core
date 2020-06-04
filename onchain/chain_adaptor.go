package onchain

import (
	///	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	errors "golang.org/x/xerrors"
)

const (
	//ETH represents the type of blockchain
	ETH = "ETH"
)

//ProxyAdapter represents an unified adapter interface for different blockchain
type ProxyAdapter interface {
	Connect(urls []string, deadline time.Time) (err error)
	DisconnectAll()
	Disconnect(idx int)
	SubscribeEvent(subscribeTypes []int) (chan interface{}, chan error)
	//Set functions
	SetGroupSize(g uint64) (err error)
	UpdateRandomness(signatures *vss.Signature) (err error)
	DataReturn(signatures *vss.Signature) (err error)
	RegisterGroupPubKey(IdWithPubKeys [5]*big.Int) (err error)
	RegisterNewNode() (err error)
	UnRegisterNode() (err error)
	SignalRandom() (err error)
	SignalGroupFormation() (err error)
	SignalGroupDissolve() (err error)
	SignalBootstrap(cid *big.Int) (err error)
	SignalUnregister(addr common.Address) (err error)
	StartCommitReveal(startBlock int64, commitDuration int64, revealDuration int64, revealThreshold int64) (err error)
	Commit(cid *big.Int, commitment [32]byte) (err error)
	Reveal(cid *big.Int, secret *big.Int) (err error)
	//Get functions
	GroupToPick() (result uint64, err error)
	PendingNonce() (result uint64, err error)
	GetExpiredWorkingGroupSize() (r uint64, err error)
	GroupSize() (r uint64, err error)
	GetWorkingGroupSize() (r uint64, err error)
	LastUpdatedBlock() (r uint64, err error)
	NumPendingGroups() (r uint64, err error)
	NumPendingNodes() (r uint64, err error)
	BootstrapEndBlk() (result uint64, err error)
	BootstrapRound() (result uint64, err error)
	Balance() (balance *big.Float, err error)
	Address() (addr common.Address)
	CurrentBlock() (r uint64, err error)
	RefreshSystemRandomHardLimit() (limit uint64, err error)
	BootstrapStartThreshold() (result uint64, err error)
	GroupPubKey(idx int) (groupPubKeys [4]*big.Int, err error)
	IsPendingNode(id []byte) (bool, error)
	BootStrapUrl() string
}

//NewProxyAdapter constructs a new ProxyAdapter with the given type of blockchain and contract addresses
func NewProxyAdapter(ChainType string, key *keystore.Key, bridgeAddr string) (ProxyAdapter, error) {
	switch ChainType {
	case ETH:
		l := log.New("module", "EthProxy")
		adaptor, err := NewEthAdaptor(key, bridgeAddr, l)
		if err != nil {
			l.Error(errors.Errorf("NewProxyAdapter : %w", err))
		}
		return adaptor, err
	default:
		err := fmt.Errorf("Chain %s not supported error\n", ChainType)
		return nil, err
	}
}
