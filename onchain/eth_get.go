package onchain

import (
	"context"
	"math/big"

	"github.com/DOSNetwork/core/onchain/dosproxy"
	"github.com/DOSNetwork/core/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	errors "golang.org/x/xerrors"
)

type getFunc func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{})

func (e *ethAdaptor) get(ctx context.Context, f getFunc, p interface{}) (interface{}, interface{}) {
	var valList []chan interface{}
	var errList []chan interface{}
	for i, client := range e.clients {
		outc, errc := f(ctx, client, e.proxies[i], p)
		valList = append(valList, outc)
		errList = append(errList, errc)
	}
	outc := first(ctx, merge(ctx, valList...))
	errc := merge(ctx, errList...)
	var (
		err    interface{}
		result interface{}
		ok     bool
	)
	for {
		select {
		case result, ok = <-outc:
			if !ok {
				return nil, err
			}
			return result, nil
		case err, ok = <-errc:
			if ok {
				e.logger.Error(err.(error))
			}
			continue
		case <-ctx.Done():
			return nil, errors.Errorf(" : %w", ctx.Err())
		}
	}
}

// LastRandomness return the last system random number
func (e *ethAdaptor) LastRandomness(ctx context.Context) (result *big.Int, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := proxy.LastRandomness(); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("LastRandomness failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}
	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Int); ok {
		result = v
	}
	if v, ok := ve.(error); ok {
		err = v
	}

	return
}

// GroupSize returns the GroupSize value
func (e *ethAdaptor) GroupSize(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := proxy.GroupSize(); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("GroupSize failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}

	return
}

// GetWorkingGroupSize returns the number of working groups
func (e *ethAdaptor) GetWorkingGroupSize(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := proxy.GetWorkingGroupSize(); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("GetWorkingGroupSize failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// NumPendingGroups returns the number of pending groups
func (e *ethAdaptor) NumPendingGroups(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := proxy.NumPendingGroups(); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("NumPendingGroups failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// RefreshSystemRandomHardLimit returns the RefreshSystemRandomHardLimit value
func (e *ethAdaptor) RefreshSystemRandomHardLimit(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := proxy.RefreshSystemRandomHardLimit(); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("RefreshSystemRandomHardLimit failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// NumPendingNodes returns the number of pending nodes
func (e *ethAdaptor) NumPendingNodes(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := proxy.NumPendingNodes(); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("RefreshSystemRandomHardLimit failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// GetExpiredWorkingGroupSize returns the expired working group size
func (e *ethAdaptor) GetExpiredWorkingGroupSize(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := proxy.GetExpiredWorkingGroupSize(); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("RefreshSystemRandomHardLimit failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// GroupToPick returns the groupToPick value
func (e *ethAdaptor) GroupToPick(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := proxy.GroupToPick(); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("GroupToPick failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// LastUpdatedBlock returns the block number of the last updated system random number
func (e *ethAdaptor) LastUpdatedBlock(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := proxy.LastUpdatedBlock(); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("GroupToPick failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// IsPendingNode checks to see if the node account is a pending node
func (e *ethAdaptor) IsPendingNode(ctx context.Context, id []byte) (result bool, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if id, ok := p.(common.Address); ok {
				if val, err := proxy.PendingNodeList(id); err != nil {
					utils.ReportResult(ctx, errc, errors.Errorf("GroupToPick failed : %w", err))
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			} else {
				utils.ReportResult(ctx, errc, errors.Errorf("PendingNodeList failed Type %T : %w", p, ErrCast))
			}
		}()
		return outc, errc
	}
	addr := common.Address{}
	b := []byte(id)
	addr.SetBytes(b)
	vr, ve := e.get(ctx, f, addr)
	if v, ok := vr.(common.Address); ok {
		if v.Big().Cmp(big.NewInt(0)) == 0 {
			result = false
		} else {
			result = true
		}
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// GroupPubKey returns the group public key of the given index
func (e *ethAdaptor) GroupPubKey(ctx context.Context, idx int) (result [4]*big.Int, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if idx, ok := p.(*big.Int); ok {
				if val, err := proxy.GetGroupPubKey(idx); err != nil {
					utils.ReportResult(ctx, errc, errors.Errorf("GetGroupPubKey failed : %w", err))
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			} else {
				utils.ReportResult(ctx, errc, errors.Errorf("GetGroupPubKey failed Type %T : %w", p, ErrCast))
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, big.NewInt(int64(idx)))
	if v, ok := vr.([4]*big.Int); ok {
		result = v
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// PendingNonce returns the account nonce of the node account in the pending state.
// This is the nonce that should be used for the next transaction.
func (e *ethAdaptor) PendingNonce(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if address, ok := p.(common.Address); ok {
				if val, err := client.PendingNonceAt(ctx, address); err != nil {
					utils.ReportResult(ctx, errc, errors.Errorf("PendingNonce failed : %w", err))
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			} else {
				utils.ReportResult(ctx, errc, errors.Errorf("PendingNonce failed Type %T : %w", p, ErrCast))
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, e.key.Address)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return

}

// Balance returns the wei balance of the node account.
func (e *ethAdaptor) Balance(ctx context.Context) (result *big.Float, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if balance, err := GetBalance(client, e.key); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("GetBalance failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, balance)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*big.Float); ok {
		result = v
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}

// CurrentBlock return the block number of the latest known header
func (e *ethAdaptor) CurrentBlock(ctx context.Context) (result uint64, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if val, err := client.HeaderByNumber(ctx, nil); err != nil {
				utils.ReportResult(ctx, errc, errors.Errorf("PendingNonce failed : %w", err))
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, nil)
	if v, ok := vr.(*types.Header); ok {
		result = v.Number.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return
}
