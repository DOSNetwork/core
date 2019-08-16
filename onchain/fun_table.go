package onchain

import (
	"context"
	"fmt"

	"github.com/DOSNetwork/core/onchain/commitreveal"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type getFunc func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{})
type setFunc func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error)

var proxyTable = []func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}){
	SubscribeLogUpdateRandom: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogUpdateRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogUpdateRandom(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeLogUrl: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogUrl)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.Contract.WatchLogUrl(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeLogRequestUserRandom: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogRequestUserRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.Contract.WatchLogRequestUserRandom(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeLogValidationResult: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogValidationResult)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.Contract.WatchLogValidationResult(opt, transitChan)
			if err != nil {
				fmt.Println("SubscribeLogValidationResult err", err)
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					fmt.Println("SubscribeLogValidationResult Done")

					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					fmt.Println("SubscribeLogValidationResult err", err)

					errc <- err
					return
				case i := <-transitChan:
					l := &LogValidationResult{
						TrafficType: i.TrafficType,
						TrafficId:   i.TrafficId,
						Message:     i.Message,
						Signature:   i.Signature,
						PubKey:      i.PubKey,
						Pass:        i.Pass,
						Version:     i.Version,
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
	SubscribeLogInsufficientPendingNode: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogInsufficientPendingNode)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogInsufficientPendingNode(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeLogInsufficientWorkingGroup: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogInsufficientWorkingGroup)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogInsufficientWorkingGroup(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeLogGroupingInitiated: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogGroupingInitiated)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogGroupingInitiated(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeDosproxyUpdateGroupSize: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyUpdateGroupSize)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchUpdateGroupSize(opt, transitChan)
			if err != nil {
				fmt.Println("WatchUpdateGroupSize err ", err)
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeDosproxyUpdateGroupToPick: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyUpdateGroupToPick)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchUpdateGroupToPick(opt, transitChan)
			if err != nil {
				fmt.Println("WatchUpdateGroupToPick err ", err)

				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					l := &LogUpdateGroupToPick{
						OldNum: i.OldNum,
						NewNum: i.NewNum,
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
	SubscribeLogGrouping: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogGrouping)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogGrouping(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeLogPublicKeyAccepted: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogPublicKeyAccepted)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogPublicKeyAccepted(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeLogPublicKeySuggested: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogPublicKeySuggested)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogPublicKeySuggested(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeLogGroupDissolve: func(ctx context.Context, proxy *dosproxy.DosproxySession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogGroupDissolve)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.Contract.WatchLogGroupDissolve(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
var crTable = []func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan interface{}){
	SubscribeCommitrevealLogStartCommitreveal: func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogStartCommitReveal)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.Contract.WatchLogStartCommitReveal(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeCommitrevealLogCommit: func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogCommit)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.Contract.WatchLogCommit(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeCommitrevealLogReveal: func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogReveal)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.Contract.WatchLogReveal(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
	SubscribeCommitrevealLogRandom: func(ctx context.Context, cr *commitreveal.CommitrevealSession) (chan interface{}, chan interface{}) {
		out := make(chan interface{})
		errc := make(chan interface{})
		opt := &bind.WatchOpts{}
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.Contract.WatchLogRandom(opt, transitChan)
			if err != nil {
				return
			}
			for {
				var log *LogCommon
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
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
