package onchain

import (
	///	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	errors "golang.org/x/xerrors"
)

const (
	//ETH represents the type of blockchain
	ETH  = "ETH"
	HECO = "Heco"
	BSC  = "BSC"
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
	SetGasPrice(gasPrice *big.Int)
	SetGasLimit(gasLimit *big.Int)
	//Get functions
	GetBlockTime() (result uint64)
	GetGasPrice() (result uint64)
	GetGasLimit() (result uint64)
	GroupToPick() (result uint64, err error)
	PendingNonce() (result uint64, err error)
	GetExpiredWorkingGroupSize() (r uint64, err error)
	GroupSize() (r uint64, err error)
	GetWorkingGroupSize() (r uint64, err error)
	LastGroupFormationRequestId() (r uint64, err error)
	LastUpdatedBlock() (r uint64, err error)
	PendingGroupStartBlock(gId *big.Int) (r uint64, err error)
	PendingGroupMaxLife() (r uint64, err error)
	FirstPendingGroupId() (gId *big.Int, err error)
	NumPendingGroups() (r uint64, err error)
	NumPendingNodes() (r uint64, err error)
	BootstrapEndBlk() (result uint64, err error)
	BootstrapRound() (result uint64, err error)
	Balance() (balance *big.Float, err error)
	Address() (addr common.Address)
	CurrentBlock() (r uint64, err error)
	RefreshSystemRandomHardLimit() (limit uint64, err error)
	CachedUpdatedBlock() (blkNum uint64, err error)
	RelayRespondLimit() (limit uint64, err error)
	BootstrapStartThreshold() (result uint64, err error)
	IsPendingNode(id []byte) (bool, error)
	BootStrapUrl() string
}

//NewProxyAdapter constructs a new ProxyAdapter with the given type of blockchain and contract addresses
func NewProxyAdapter(key *keystore.Key, config *configuration.Config) (ProxyAdapter, error) {
	switch config.ChainType {
	case ETH, HECO, BSC:
		l := log.New("module", "EthProxy")
		adaptor, err := NewEthAdaptor(key, config, l)
		if err != nil {
			l.Error(errors.Errorf("NewProxyAdapter : %w", err))
		}
		return adaptor, err
	default:
		err := fmt.Errorf("Chain %s not supported error\n", config.ChainType)
		return nil, err
	}
}
