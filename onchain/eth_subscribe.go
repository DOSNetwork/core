package onchain

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"
	"time"

	"github.com/DOSNetwork/core/onchain/commitreveal"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	errors "golang.org/x/xerrors"
)

const (
	//SubscribeLogUpdateRandom is a log type to subscribe the event LogUpdateRandom
	SubscribeLogUpdateRandom = iota
	//SubscribeLogRequestUserRandom is a log type to subscribe the event LogRequestUserRandom
	SubscribeLogRequestUserRandom
	//SubscribeLogUrl is a log type to subscribe the event LogUrl
	SubscribeLogUrl
	//SubscribeLogValidationResult is a log type to subscribe the event LogValidationResult
	SubscribeLogValidationResult
	//SubscribeLogGrouping is a log type to subscribe the event LogGrouping
	SubscribeLogGrouping
	//SubscribeLogPublicKeyAccepted is a log type to subscribe the event LogPublicKeyAccepted
	SubscribeLogPublicKeyAccepted
	//SubscribeLogPublicKeySuggested is a log type to subscribe the event LogPublicKeySuggested
	SubscribeLogPublicKeySuggested
	//SubscribeLogGroupDissolve is a log type to subscribe the event LogGroupDissolve
	SubscribeLogGroupDissolve
	//SubscribeLogInsufficientPendingNode is a log type to subscribe the event LogInsufficientPendingNode
	SubscribeLogInsufficientPendingNode
	//SubscribeLogInsufficientWorkingGroup is a log type to subscribe the event LogInsufficientWorkingGroup
	SubscribeLogInsufficientWorkingGroup
	//SubscribeLogNoWorkingGroup is a log type to subscribe the event LogNoWorkingGroup
	SubscribeLogNoWorkingGroup
	//SubscribeLogGroupingInitiated is a log type to subscribe the event GroupingInitiated
	SubscribeLogGroupingInitiated
	//SubscribeDosproxyUpdateBootstrapGroups is a log type to subscribe the event UpdateGroupToPick
	SubscribeDosproxyUpdateBootstrapGroups
	//SubscribeDosproxyUpdateGroupSize is a log type to subscribe the event UpdateGroupSize
	SubscribeDosproxyUpdateGroupSize
	//SubscribeCommitrevealLogStartCommitreveal is a log type to subscribe the event StartCommitreveal
	SubscribeCommitrevealLogStartCommitreveal
	//SubscribeCommitrevealLogCommit is a log type to subscribe the event LogCommit
	SubscribeCommitrevealLogCommit
	//SubscribeCommitrevealLogReveal is a log type to subscribe the event LogReveal
	SubscribeCommitrevealLogReveal
	//SubscribeCommitrevealLogRandom is a log type to subscribe the event LogRandom
	SubscribeCommitrevealLogRandom
)

func firstEvent(ctx context.Context, source chan interface{}) (out chan interface{}) {
	out = make(chan interface{})

	go func() {
		defer close(out)
		visited := make(map[string]uint64)
		for {
			select {
			case <-ctx.Done():
				return
			case event, ok := <-source:
				if !ok {
					return
				}
				if content, ok := event.(*LogCommon); ok {
					if content.Removed {
						continue
					}
					var bytes []byte
					bytes = append(bytes, content.Raw.Data...)
					bytes = append(bytes, new(big.Int).SetUint64(content.BlockN).Bytes()...)
					nHash := sha256.Sum256(bytes)

					identity := string(nHash[:])
					if visited[identity] == 0 {
						visited[identity] = content.BlockN
						select {
						case out <- content.log:
						case <-ctx.Done():
						}
						go func(identity string) {
							select {
							case <-ctx.Done():
							case <-time.After(100 * 15 * time.Second):
								delete(visited, identity)
							}
						}(identity)
					}
				}
			}
		}
	}()

	return
}

// SubscribeEvent is a log subscription operation
func (e *ethAdaptor) SubscribeEvent(subscribeTypes []int) (chan interface{}, chan error) {
	var eventList []chan interface{}
	var errcs []chan error
	fmt.Println("Subscribe proxies ", len(e.proxies), " crs ", len(e.crs))
	for _, subscribeType := range subscribeTypes {
		if subscribeType >= SubscribeCommitrevealLogStartCommitreveal {
			for i := 0; i < len(e.crs); i++ {
				fmt.Println("Subscribe CR Event ", i)
				if e.crs[i] == nil ||
					e.ctxes[i] == nil {
					continue
				}
				select {
				case <-e.ctxes[i].Done():
					continue
				default:
				}
				out, errc := crTable[subscribeType](e.ctxes[i], e.crs[i])
				eventList = append(eventList, out)
				errcs = append(errcs, errc)
			}
		} else {
			for i := 0; i < len(e.proxies); i++ {
				fmt.Println("SubscribeEvent ", i, subscribeType)
				if e.proxies[i] == nil ||
					e.ctxes[i] == nil {
					continue
				}
				select {
				case <-e.ctxes[i].Done():
					continue
				default:
				}
				out, errc := proxyTable[subscribeType](e.ctxes[i], e.proxies[i])
				eventList = append(eventList, out)
				errcs = append(errcs, errc)
			}
		}
	}
	return firstEvent(e.ctx, merge(e.ctx, eventList...)), mergeError(e.ctx, errcs...)
}

func getIndex(ctx context.Context) (idx int) {
	if v := ctx.Value("index"); v != nil {
		if i, ok := v.(int); ok {
			idx = i
		} else {
			idx = -1
		}
	}
	return
}
func replyError(ctx context.Context, errc chan error, err error) {
	select {
	case <-ctx.Done():
	case errc <- err:
	}
}

var proxyTable = []func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error){
	SubscribeLogUpdateRandom: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogUpdateRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogUpdateRandom(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})
					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogUpdateRandom{
						LastRandomness:    i.LastRandomness,
						DispatchedGroupId: i.DispatchedGroupId,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogUrl: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogUrl)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.Contract.WatchLogUrl(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})
					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogUrl{
						QueryId:           i.QueryId,
						Timeout:           i.Timeout,
						DataSource:        i.DataSource,
						Selector:          i.Selector,
						Randomness:        i.Randomness,
						DispatchedGroupId: i.DispatchedGroupId,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogRequestUserRandom: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogRequestUserRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.Contract.WatchLogRequestUserRandom(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogRequestUserRandom{
						RequestId:            i.RequestId,
						LastSystemRandomness: i.LastSystemRandomness,
						UserSeed:             i.UserSeed,
						DispatchedGroupId:    i.DispatchedGroupId,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogValidationResult: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogValidationResult)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.Contract.WatchLogValidationResult(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogValidationResult{
						TrafficType: i.TrafficType,
						TrafficId:   i.TrafficId,
						Message:     i.Message,
						Signature:   i.Signature,
						PubKey:      i.PubKey,
						Pass:        i.Pass,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogInsufficientPendingNode: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogInsufficientPendingNode)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogInsufficientPendingNode(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogInsufficientPendingNode{
						NumPendingNodes: i.NumPendingNodes,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogInsufficientWorkingGroup: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogInsufficientWorkingGroup)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogInsufficientWorkingGroup(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogInsufficientWorkingGroup{
						NumWorkingGroups: i.NumWorkingGroups,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogGroupingInitiated: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogGroupingInitiated)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogGroupingInitiated(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogGroupingInitiated{}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeDosproxyUpdateGroupSize: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyUpdateGroupSize)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchUpdateGroupSize(opt, transitChan)
			if err != nil {

				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogUpdateGroupSize{
						OldSize: i.OldSize,
						NewSize: i.NewSize,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogGrouping: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogGrouping)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogGrouping(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					fmt.Println("SubscribeEvent ctx done")
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					var participants [][]byte
					for _, p := range i.NodeId {
						id := p.Bytes()
						participants = append(participants, id)
					}
					l := &LogGrouping{
						GroupId: i.GroupId,
						NodeId:  participants,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogPublicKeyAccepted: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogPublicKeyAccepted)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogPublicKeyAccepted(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogPublicKeyAccepted{
						GroupId:          i.GroupId,
						WorkingGroupSize: i.NumWorkingGroups,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogPublicKeySuggested: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogPublicKeySuggested)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogPublicKeySuggested(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogPublicKeySuggested{
						GroupId: i.GroupId,
						Count:   i.PubKeyCount,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeLogGroupDissolve: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogGroupDissolve)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogGroupDissolve(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogGroupDissolve{
						GroupId: i.GroupId,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
}
var crTable = []func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan error){
	SubscribeCommitrevealLogStartCommitreveal: func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogStartCommitReveal)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.Contract.WatchLogStartCommitReveal(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogStartCommitReveal{
						Cid:             i.Cid,
						StartBlock:      i.StartBlock,
						CommitDuration:  i.CommitDuration,
						RevealDuration:  i.RevealDuration,
						RevealThreshold: i.RevealThreshold,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeCommitrevealLogCommit: func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogCommit)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.Contract.WatchLogCommit(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogCommit{
						Cid:        i.Cid,
						From:       i.From,
						Commitment: i.Commitment,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeCommitrevealLogReveal: func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogReveal)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.Contract.WatchLogReveal(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogReveal{
						Cid:    i.Cid,
						From:   i.From,
						Secret: i.Secret,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
	SubscribeCommitrevealLogRandom: func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.Contract.WatchLogRandom(opt, transitChan)
			if err != nil {
				replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					replyError(ctx, errc, &OnchainError{err: errors.Errorf("SubscribeEvent err : %w", err), Idx: getIndex(ctx)})

					if err == io.EOF {
						return
					}
					continue
				case i := <-transitChan:
					l := &LogRandom{
						Cid:    i.Cid,
						Random: i.Random,
					}
					log = &LogCommon{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
						log:     l,
					}
				}
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case out <- log:
				}
			}
		}()
		return out, errc
	},
}
