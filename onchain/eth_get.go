package onchain

import (
	"context"
	"math/big"
	"strings"

	"github.com/DOSNetwork/core/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	errors "golang.org/x/xerrors"
)

type getFunc func(ctx context.Context) (chan interface{}, chan error)

func (e *ethAdaptor) get(f getFunc) (result interface{}, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
		return
	}
	var valList []chan interface{}
	var errList []chan error
	for _, ctx := range e.ctxes {
		select {
		case <-ctx.Done():
		default:
			opCtx, opCancel := context.WithTimeout(ctx, e.getTimeout)
			defer opCancel()
			outc, errc := f(opCtx)
			valList = append(valList, outc)
			errList = append(errList, errc)
		}
	}
	outc := first(e.ctx, merge(e.ctx, valList...))
	errc := mergeError(e.ctx, errList...)
	var ok bool
	select {
	case result = <-outc:
	case <-e.ctx.Done():
		return nil, errors.Errorf(" : %w", e.ctx.Err())
	}
	for {
		select {
		case err, ok = <-errc:
			if !ok {
				return
			}
			if strings.Contains(err.Error(), "use of closed network connection") {
				var oError *OnchainError
				if errors.As(err, &oError) {
					e.cancels[oError.Idx]()
				}
			}
			continue
		case <-e.ctx.Done():
			return nil, errors.Errorf(" : %w", e.ctx.Err())
		}
	}
}

// GroupToPick returns the groupToPick value
func (e *ethAdaptor) GroupToPick() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].GroupToPick(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}

// GetExpiredWorkingGroupSize returns the GetExpiredWorkingGroupSize value
func (e *ethAdaptor) GetExpiredWorkingGroupSize() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].GetExpiredWorkingGroupSize(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}

func (e *ethAdaptor) GroupSize() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].GroupSize(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}
func (e *ethAdaptor) GetWorkingGroupSize() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].GetWorkingGroupSize(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}

func (e *ethAdaptor) LastGroupFormationRequestId() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].LastFormGrpReqId(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}

func (e *ethAdaptor) LastUpdatedBlock() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].LastUpdatedBlock(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}

type PendingGroupT struct {
	GroupId     *big.Int
	StartBlkNum *big.Int
}

func (e *ethAdaptor) PendingGroupStartBlock(groupId *big.Int) (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].PendingGroups(groupId); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(struct {
		GroupId     *big.Int
		StartBlkNum *big.Int
	}); ok {
		result = v.StartBlkNum.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}

func (e *ethAdaptor) PendingGroupMaxLife() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].PendingGroupMaxLife(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}

func (e *ethAdaptor) FirstPendingGroupId() (result *big.Int, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].PendingGroupList(big.NewInt(1)); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v
	} else {
		err = errors.New("casting error")
	}
	return
}

func (e *ethAdaptor) NumPendingGroups() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].NumPendingGroups(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}

func (e *ethAdaptor) NumPendingNodes() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].NumPendingNodes(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}
func (e *ethAdaptor) BootstrapEndBlk() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].BootstrapEndBlk(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}
func (e *ethAdaptor) RefreshSystemRandomHardLimit() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].RefreshSystemRandomHardLimit(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}
func (e *ethAdaptor) BootstrapStartThreshold() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].BootstrapStartThreshold(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}

func (e *ethAdaptor) GetGasPrice() (result uint64) {
	if len(e.proxies) != 0 {
		if e.proxies[0].TransactOpts.GasPrice == nil {
			result = 0
		} else {
			result = e.proxies[0].TransactOpts.GasPrice.Uint64()
		}
	}
	return
}

func (e *ethAdaptor) GetGasLimit() (result uint64) {
	if len(e.proxies) != 0 {
		result = e.proxies[0].TransactOpts.GasLimit
	}
	return
}

func (e *ethAdaptor) GroupPubKey(gIdx int) (result [4]*big.Int, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].GetGroupPubKey(big.NewInt(int64(gIdx))); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.([4]*big.Int); ok {
		result = v
	} else {
		err = errors.New("casting error")
	}
	return
}
func (e *ethAdaptor) IsPendingNode(id []byte) (result bool, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	addr := common.Address{}
	addr.SetBytes(id)
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].PendingNodeList(addr); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(common.Address); ok {
		if v.Big().Cmp(big.NewInt(0)) == 0 {
			result = false
		} else {
			result = true
		}
	} else {
		err = errors.New("casting error")
	}
	return
}
func (e *ethAdaptor) BootstrapRound() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	proxies := e.proxies
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := proxies[idx].BootstrapRound(); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Int); ok {
		result = v.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}
func (e *ethAdaptor) Balance() (result *big.Float, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	clients := e.clients
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := GetBalance(clients[idx], e.key); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*big.Float); ok {
		result = v
	} else {
		err = errors.New("casting error")
	}
	return
}
func (e *ethAdaptor) CurrentBlock() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	clients := e.clients
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
			} else {
				if val, err := clients[idx].HeaderByNumber(ctx, nil); err != nil {
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
				} else {
					utils.ReportResult(ctx, outc, val)
				}
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(*types.Header); ok {
		result = v.Number.Uint64()
	} else {
		err = errors.New("casting error")
	}
	return
}
func (e *ethAdaptor) PendingNonce() (result uint64, err error) {
	if !e.isConnecting() {
		err = errors.New("not connecting to geth")
	}
	clients := e.clients
	f := func(ctx context.Context) (chan interface{}, chan error) {
		outc := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(outc)
			defer close(errc)
			idx := getIndex(ctx)
			if idx == -1 {
				err = errors.New("no client index in context")
				return
			}
			if val, err := clients[idx].PendingNonceAt(ctx, e.key.Address); err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("get err : %w", err), Idx: idx})
			} else {
				utils.ReportResult(ctx, outc, val)
			}
		}()
		return outc, errc
	}
	var r interface{}
	if r, err = e.get(f); err != nil {
		return
	}
	if v, ok := r.(uint64); ok {
		result = v
	} else {
		err = errors.New("casting error")
	}
	return

}
