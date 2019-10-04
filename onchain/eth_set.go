package onchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/DOSNetwork/core/onchain/commitreveal"
	"github.com/DOSNetwork/core/onchain/dosproxy"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	errors "golang.org/x/xerrors"
)

type setFunc func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error)

func (e *ethAdaptor) ReqLoop() (err error) {
	defer e.logger.Event("ReqLoopEnd", map[string]interface{}{"Topic": "LoopLife"})
	if err = e.isDone(); err != nil {
		return
	}
	for {
		select {
		case req := <-e.reqQueue:
			tx, err := req.f(req.ctx, req.proxy, req.cr, req.params)
			resp := &response{req.idx, tx, err}
			go func(req *request, resp *response) {
				select {
				case req.reply <- resp:
				case <-req.ctx.Done():
				}
			}(req, resp)
		case <-e.ctx.Done():
			err = e.ctx.Err()
			return
		}
	}
}

func (e *ethAdaptor) set(ctx context.Context, params []interface{}, setF setFunc) (reply chan *response) {
	f := func(ctx context.Context, idx int, pre chan *response, r *request) (out chan *response) {
		out = make(chan *response)
		go func() {
			defer close(out)
			if pre != nil {
				select {
				case <-ctx.Done():
					return
				case resp, ok := <-pre:
					if ok {
						//Request has been fulfulled by previous sendRequest or
						//transaction failed so delete the whole requestSend chain
						if resp.err == nil ||
							strings.Contains(resp.err.Error(), "transaction failed") ||
							strings.Contains(resp.err.Error(), "insufficient funds for gas * price + value") {
							select {
							case out <- resp:
							case <-ctx.Done():
							}
							return
						}
						fmt.Println("Switch to ", idx, " Client to handle request because of e ,", resp.err)
					}
				}
			}
			r.reply = make(chan *response)
			defer close(r.reply)
			select {
			case e.reqQueue <- r:
			case <-ctx.Done():
			}

			select {
			case resp, ok := <-r.reply:
				if ok {
					select {
					case out <- resp:
					case <-ctx.Done():
					}
				}
			case <-ctx.Done():
			}
		}()
		return
	}

	for i, proxy := range e.proxies {
		r := &request{ctx, i, proxy, e.crs[i], setF, params, nil}
		reply = f(ctx, i, reply, r)
	}

	return
}

func waitForReply(ctx context.Context,funcName string,reply chan *response)(err error){
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if err != nil {
				fmt.Println(funcName," error ", r.err)
				return
			}
			fmt.Println(funcName," tx:", fmt.Sprintf("%x", r.tx.Hash()))
		}
	case <-ctx.Done():
		err=ctx.Err()
	}
	return
}

// StartCommitReveal is a wrap function that build a pipeline to set groupToPick
func (e *ethAdaptor) StartCommitReveal(ctx context.Context, startBlock int64, commitDuration int64, revealDuration int64, revealThreshold int64) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		if len(p) != 4 {
			err = errors.New("Invalid parameter")
			return
		}
		if startBlock, ok := p[0].(*big.Int); ok {
			if commitDuration, ok := p[1].(*big.Int); ok {
				if revealDuration, ok := p[2].(*big.Int); ok {
					if revealThreshold, ok := p[3].(*big.Int); ok {
						tx, err = cr.StartCommitReveal(startBlock, commitDuration, revealDuration, revealThreshold)
					}
				}
			}
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, big.NewInt(startBlock))
	params = append(params, big.NewInt(commitDuration))
	params = append(params, big.NewInt(revealDuration))
	params = append(params, big.NewInt(revealThreshold))
	return waitForReply(ctx,"StartCommitReveal",e.set(ctx, params, f))
}
/*
// SetGroupToPick is a wrap function that build a pipeline to set groupToPick
func (e *ethAdaptor) SetGroupToPick(ctx context.Context, groupToPick uint64) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		if len(p) != 1 {
			err = errors.New("Invalid parameter")
			return
		}
		if groupToPick, ok := p[0].(*big.Int); ok {
			tx, err = proxy.SetGroupToPick(groupToPick)
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, big.NewInt(int64(groupToPick)))

	return waitForReply(ctx,"SetGroupToPick",e.set(ctx, params, f))
}
*/
// RegisterNewNode is a wrap function that build a pipeline to call RegisterNewNode
func (e *ethAdaptor) RegisterNewNode(ctx context.Context) (err error) {
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.RegisterNewNode()
		return
	}
	return waitForReply(ctx,"RegisterNewNode",e.set(ctx, nil, f))
}

// UnRegisterNode is a wrap function that build a pipeline to call RegisterNewNode
func (e *ethAdaptor) UnRegisterNode(ctx context.Context) (err error) {
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.UnregisterNode()
		return
	}
	return waitForReply(ctx,"UnregisterNode",e.set(ctx, nil, f))
}

// SignalRandom is a wrap function that build a pipeline to call SignalRandom
func (e *ethAdaptor) SignalRandom(ctx context.Context) (err error) {
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.SignalRandom()
		return
	}
	return waitForReply(ctx,"SignalRandom",e.set(ctx, nil, f))
}

// SignalGroupFormation is a wrap function that build a pipeline to call SignalGroupFormation
func (e *ethAdaptor) SignalGroupFormation(ctx context.Context) (err error) {
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.SignalGroupFormation()
		return
	}
	return waitForReply(ctx,"SignalGroupFormation",e.set(ctx, nil, f))
}

// SignalGroupDissolve is a wrap function that build a pipeline to call SignalGroupDissolve
func (e *ethAdaptor) SignalGroupDissolve(ctx context.Context) (err error) {
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.SignalGroupDissolve()
		return
	}
	return waitForReply(ctx,"SignalGroupFormation",e.set(ctx, nil, f))
}

// SignalBootstrap is a wrap function that build a pipeline to call SignalBootstrap
func (e *ethAdaptor) SignalUnregister(ctx context.Context, addr common.Address) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		if len(p) != 1 {
			err = errors.New("Invalid parameter")
			return
		}
		if addr, ok := p[0].(common.Address); ok {
			tx, err = proxy.SignalUnregister(addr)
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, addr)
	return waitForReply(ctx,"SignalGroupFormation",e.set(ctx, params, f))
}

// SignalBootstrap is a wrap function that build a pipeline to call SignalBootstrap
func (e *ethAdaptor) SignalBootstrap(ctx context.Context, cid *big.Int) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		if len(p) != 1 {
			err = errors.New("Invalid parameter")
			return
		}
		if cid, ok := p[0].(*big.Int); ok {
			tx, err = proxy.SignalBootstrap(cid)
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, cid)
	return waitForReply(ctx,"SignalGroupFormation",e.set(ctx, params, f))
}
/*
// SetGroupSize is a wrap function that build a pipeline to call SetGroupSize
func (e *ethAdaptor) SetGroupSize(ctx context.Context, size uint64) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		if len(p) != 1 {
			err = errors.New("Invalid parameter")
			return
		}
		if size, ok := p[0].(*big.Int); ok {
			tx, err = proxy.SetGroupSize(size)
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, big.NewInt(int64(size)))
	return waitForReply(ctx,"SetGroupSize",e.set(ctx, params, f))
}
*/
// SetGroupMaturityPeriod is a wrap function that build a pipeline to call SetGroupMaturityPeriod
func (e *ethAdaptor) SetGroupMaturityPeriod(ctx context.Context, period uint64) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		if len(p) != 1 {
			err = errors.New("Invalid parameter")
			return
		}
		if period, ok := p[0].(*big.Int); ok {
			tx, err = proxy.SetGroupMaturityPeriod(period)
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, big.NewInt(int64(period)))
	return waitForReply(ctx,"SetGroupMaturityPeriod",e.set(ctx, params, f))
}
/*
// SetGroupingThreshold is a wrap function that build a pipeline to call SetGroupingThreshold
func (e *ethAdaptor) SetGroupingThreshold(ctx context.Context, threshold uint64) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		if len(p) != 1 {
			err = errors.New("Invalid parameter")
			return
		}
		if threshold, ok := p[0].(*big.Int); ok {
			tx, err = proxy.SetGroupingThreshold(threshold)
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, big.NewInt(int64(threshold)))
	return waitForReply(ctx,"SetGroupingThreshold",e.set(ctx, params, f))
}
*/
// Commit is a wrap function that build a pipeline to call Commit
func (e *ethAdaptor) Commit(ctx context.Context, cid *big.Int, commitment [32]byte) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		fmt.Println("inter Commit")
		if len(p) != 2 {
			err = errors.New("Invalid parameter")
			return
		}
		if cid, ok := p[0].(*big.Int); ok {
			if commitment, ok := p[1].([32]byte); ok {
				tx, err = cr.Commit(cid, commitment)
			}
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, cid)
	params = append(params, commitment)
	return waitForReply(ctx,"Commit",e.set(ctx, params, f))

}

// Reveal is a wrap function that build a pipeline to call Reveal
func (e *ethAdaptor) Reveal(ctx context.Context, cid *big.Int, secret *big.Int) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		err = errors.New("Invalid parameter")
		if len(p) != 2 {
			return
		}
		if cid, ok := p[0].(*big.Int); ok {
			if secret, ok := p[1].(*big.Int); ok {
				tx, err = cr.Reveal(cid, secret)
			}
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, cid)
	params = append(params, secret)
	return waitForReply(ctx,"Reveal",e.set(ctx, params, f))

}

// RegisterGroupPubKey is a wrap function that build a pipeline to call RegisterGroupPubKey
func (e *ethAdaptor) RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc chan error) {
	fmt.Println("RegisterGroupPubKey")
	errc = make(chan error)
	go func() {
		defer close(errc)
		// define how to parse parameters and execute proxy function
		f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
			err = errors.New("Invalid parameter")
			if len(p) != 1 {
				return
			}
			if idPubkey, ok := p[0].([5]*big.Int); ok {
				groupId := idPubkey[0]
				var pubKey [4]*big.Int
				copy(pubKey[:], idPubkey[1:])
				select {
				default:
					tx, err = proxy.RegisterGroupPubKey(groupId, pubKey)
				case <-ctx.Done():
					err = ctx.Err()
				}
			}
			return
		}
		// define parameters
		var params []interface{}
		var groupID string
		var startTime time.Time
		select {
		case idPubkey, ok := <-IdWithPubKeys:
			if !ok {
				return
			}
			startTime = time.Now()
			groupID = fmt.Sprintf("%x", idPubkey[0])
			params = append(params, idPubkey)
		case <-ctx.Done():
			return
		}
		if err := waitForReply(ctx,"RegisterGroupPubKey",e.set(ctx, params, f)); err!=nil{
			select {
			case errc <- err:
			case <-ctx.Done():
			}
		}else{
			e.logger.TimeTrack(startTime, "RegisterGroupPubKey", map[string]interface{}{"GroupID": groupID,"Topic": "Grouping"})
		}
	}()
	return
}

// SetRandomNum is a wrap function that build a pipeline to call SetRandomNum
func (e *ethAdaptor) SetRandomNum(ctx context.Context, signatures chan *vss.Signature) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		// define how to parse parameters and execute proxy function
		f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
			err = errors.New("Invalid parameter")
			if len(p) != 1 {
				return
			}
			if sign, ok := p[0].(*vss.Signature); ok {
				select {
				default:
					x, y := sign.ToBigInt()
					sig := [2]*big.Int{x, y}
					tx, err = proxy.UpdateRandomness(sig)
				case <-ctx.Done():
					err = ctx.Err()
				}
			}
			return
		}
		var params []interface{}
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			params = append(params, signature)
		case <-ctx.Done():
			return
		}
		if err := waitForReply(ctx,"SetRandomNum",e.set(ctx, params, f)); err!=nil{
			select {
			case errc <- err:
			case <-ctx.Done():
			}
		}
	}()
	return
}

// DataReturn is a wrap function that build a pipeline to call DataReturn
func (e *ethAdaptor) DataReturn(ctx context.Context, signatures chan *vss.Signature) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		// define how to parse parameters and execute proxy function
		f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
			err = errors.New("Invalid parameter")
			if len(p) != 1 {
				return
			}
			if sign, ok := p[0].(*vss.Signature); ok {
				select {
					default:
						requestId := new(big.Int).SetBytes(sign.RequestId)
						trafficType := uint8(sign.Index)
						result := sign.Content
						x, y := sign.ToBigInt()
						sig := [2]*big.Int{x, y}
						tx, err = proxy.TriggerCallback(requestId, trafficType, result, sig)
					case <-ctx.Done():
						err = ctx.Err()
				}
			}
			return
		}
		var params []interface{}
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			params = append(params, signature)
		case <-ctx.Done():
			return
		}
		if err := waitForReply(ctx,"DataReturn",e.set(ctx, params, f)); err!=nil{
			select {
			case errc <- err:
			case <-ctx.Done():
			}
		}
	}()
	return
}
