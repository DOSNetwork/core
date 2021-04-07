package onchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	errors "golang.org/x/xerrors"
)

func (e *ethAdaptor) ReqLoop() {
	for {
		select {
		case req := <-e.reqQueue:
			e.handleReq(req)
		case <-e.ctx.Done():
			return
		}
	}
}

/*
Err: transaction failed
Err: insufficient funds for gas * price + value
->Fail request and send back error to requester to handle the error

Err: failed to retrieve account nonce
Err: use of closed network connection
->Switch to other geth to set again
*/
func (e *ethAdaptor) handleReq(req *request) {
	var tx *types.Transaction
	var err error
	var idx int
	var ctx context.Context
L:
	for idx, ctx = range e.ctxes {
		select {
		case <-req.opCtx.Done():
			return
		case <-ctx.Done():
			continue
		default:
			tx, err = req.f(ctx)
			if err != nil {
				if strings.Contains(err.Error(), "transaction failed") ||
					strings.Contains(err.Error(), "insufficient funds for gas * price + value") {
					break L
				}
				if strings.Contains(err.Error(), "failed to retrieve account nonce") ||
					strings.Contains(err.Error(), "use of closed network connection") {
					var oError *OnchainError
					if errors.As(err, &oError) {
						e.cancels[oError.Idx]()
					}
				}
				continue
			}
			break L
		}
	}

	resp := &response{idx, tx, err}
	go func() {
		defer close(req.reply)
		select {
		case req.reply <- resp:
		case <-req.opCtx.Done():
		}
	}()
}

func (e *ethAdaptor) waitForReply(req *request) (err error) {
	for {
		select {
		case e.reqQueue <- req:
		case <-e.ctx.Done():
			err = e.ctx.Err()
			return
		case <-req.opCtx.Done():
			err = req.opCtx.Err()
			return
		}

		select {
		case r := <-req.reply:
			err = r.err
			if err != nil {
				fmt.Println(" error ", r.err)
				return
			}
		case <-e.ctx.Done():
			err = e.ctx.Err()
			return
		case <-req.opCtx.Done():
			err = req.opCtx.Err()
			return
		}
		return
	}
}

func (e *ethAdaptor) SetGroupSize(g uint64) (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	proxies := e.proxies
	groupSize := new(big.Int).SetUint64(g)

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = proxies[idx].SetGroupSize(groupSize)
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {
				e.logger.Info(fmt.Sprintf("SetGroupSize %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}

func (e *ethAdaptor) UpdateRandomness(sign *vss.Signature) (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	proxies := e.proxies
	x, y := sign.ToBigInt()
	sig := [2]*big.Int{x, y}

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = proxies[idx].UpdateRandomness(sig)
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {
				e.logger.Info(fmt.Sprintf("UpdateRandomness %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}

func (e *ethAdaptor) DataReturn(sign *vss.Signature) (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	proxies := e.proxies
	requestId := new(big.Int).SetBytes(sign.RequestId)
	trafficType := uint8(sign.Index)
	result := sign.Content
	x, y := sign.ToBigInt()
	sig := [2]*big.Int{x, y}

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = proxies[idx].TriggerCallback(requestId, trafficType, result, sig)
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {
				e.logger.Info(fmt.Sprintf("DataReturn %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}

func (e *ethAdaptor) RegisterGroupPubKey(idPubkey [5]*big.Int) (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	proxies := e.proxies
	groupId := idPubkey[0]
	var pubKey [4]*big.Int
	copy(pubKey[:], idPubkey[1:])

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = proxies[idx].RegisterGroupPubKey(groupId, pubKey)
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {
				e.logger.Info(fmt.Sprintf("RegisterGroupPubKey %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}

func (e *ethAdaptor) RegisterNewNode() (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	proxies := e.proxies

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = proxies[idx].RegisterNewNode()
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {
				e.logger.Info(fmt.Sprintf("RegisterNewNode %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}

func (e *ethAdaptor) UnRegisterNode() (err error) {
	if !e.isConnecting(true) {
		err = errors.New("[ONCHAIN] not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	proxies := e.proxies

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = proxies[idx].UnregisterNode()
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {
				e.logger.Info(fmt.Sprintf("UnRegisterNode %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}

func (e *ethAdaptor) SignalRandom() (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}
	if len(e.proxies) != 0 && len(e.rpcClients) != 0 {
		tx, err := e.proxies[0].SignalRandom()
		if err != nil {
			return err
		}
		e.logger.Info(fmt.Sprintf("SignalRandom tx sent: %x, waiting for confirmation...", tx.Hash()))
		if err = CheckTransaction(e.rpcClients[0], e.blockTime, tx); err != nil {
			e.logger.Error(err)
		}
	}
	return
}

func (e *ethAdaptor) SignalGroupFormation() (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}
	if len(e.proxies) != 0 && len(e.rpcClients) != 0 {
		tx, err := e.proxies[0].SignalGroupFormation()
		if err != nil {
			return err
		}
		e.logger.Info(fmt.Sprintf("SignalGroupFormation tx sent: %x, waiting for confirmation...", tx.Hash()))
		if err = CheckTransaction(e.rpcClients[0], e.blockTime, tx); err != nil {
			e.logger.Error(err)
		}
	}
	return
}

func (e *ethAdaptor) SignalGroupDissolve() (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}
	if len(e.proxies) != 0 && len(e.rpcClients) != 0 {
		tx, err := e.proxies[0].SignalGroupDissolve()
		if err != nil {
			return err
		}
		e.logger.Info(fmt.Sprintf("SignalGroupDissolve tx sent: %x, waiting for confirmation...", tx.Hash()))
		if err = CheckTransaction(e.rpcClients[0], e.blockTime, tx); err != nil {
			e.logger.Error(err)
		}
	}
	return
}

func (e *ethAdaptor) SetGasLimit(gasLimit *big.Int) {
	if len(e.proxies) > 0 && len(e.crs) > 0 {
		e.proxies[0].TransactOpts.GasLimit = gasLimit.Uint64()
		e.crs[0].TransactOpts.GasLimit = gasLimit.Uint64()
	}
}

func (e *ethAdaptor) SetGasPrice(gasPrice *big.Int) {
	if len(e.proxies) > 0 && len(e.crs) > 0 {
		if gasPrice.Cmp(big.NewInt(0)) == 0 {
			e.proxies[0].TransactOpts.GasPrice = nil
			e.crs[0].TransactOpts.GasPrice = nil
		} else {
			e.proxies[0].TransactOpts.GasPrice = gasPrice
			e.crs[0].TransactOpts.GasPrice = gasPrice
		}
	}
}

func (e *ethAdaptor) SignalBootstrap(cid *big.Int) (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}
	if len(e.proxies) != 0 && len(e.rpcClients) != 0 {
		tx, err := e.proxies[0].SignalBootstrap(cid)
		if err != nil {
			return err
		}
		e.logger.Info(fmt.Sprintf("SignalBootstrap tx sent: %x, waiting for confirmation...", tx.Hash()))
		if err = CheckTransaction(e.rpcClients[0], e.blockTime, tx); err != nil {
			e.logger.Error(err)
		}
	}
	return
}

func (e *ethAdaptor) SignalUnregister(addr common.Address) (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	proxies := e.proxies

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = proxies[idx].SignalUnregister(addr)
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {

				e.logger.Info(fmt.Sprintf("SignalUnregister %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}

func (e *ethAdaptor) StartCommitReveal(startBlock int64, commitDuration int64, revealDuration int64, revealThreshold int64) (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	crs := e.crs

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = crs[idx].StartCommitReveal(big.NewInt(startBlock), big.NewInt(commitDuration), big.NewInt(revealDuration), big.NewInt(revealThreshold))
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {
				e.logger.Info(fmt.Sprintf("StartCommitReveal %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}

func (e *ethAdaptor) Commit(cid *big.Int, commitment [32]byte) (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	crs := e.crs

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = crs[idx].Commit(cid, commitment)
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {
				e.logger.Info(fmt.Sprintf("Commit %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}

func (e *ethAdaptor) Reveal(cid *big.Int, secret *big.Int) (err error) {
	if !e.isConnecting(true) {
		err = errors.New("not connecting to geth")
		return
	}

	// define how to parse parameters and execute proxy function
	opCtx, opCancel := context.WithTimeout(e.ctx, e.setTimeout)
	defer opCancel()
	crs := e.crs

	f := func(ctx context.Context) (tx *types.Transaction, err error) {
		idx := getIndex(ctx)
		if idx == -1 {
			err = errors.New("no client index in context")
		} else {
			tx, err = crs[idx].Reveal(cid, secret)
			if err != nil {
				err = &OnchainError{err: errors.Errorf(": %w", err), Idx: idx}
			} else {
				e.logger.Info(fmt.Sprintf("Reveal %x", tx.Hash()))
			}
		}
		if err != nil {
			e.logger.Error(err)
		}
		return
	}
	return e.waitForReply(&request{opCtx: opCtx, f: f, reply: make(chan *response)})
}
