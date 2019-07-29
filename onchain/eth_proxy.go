package onchain

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math"
	"math/big"
	//	"runtime/debug"
	"strings"
	"time"

	"github.com/go-stack/stack"

	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain/commitreveal"
	"github.com/DOSNetwork/core/onchain/dosbridge"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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
	//SubscribeDosproxyUpdateGroupToPick is a log type to subscribe the event UpdateGroupToPick
	SubscribeDosproxyUpdateGroupToPick
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
const (
	//TrafficSystemRandom is a request type to build a corresponding pipeline
	TrafficSystemRandom = iota // 0
	//TrafficUserRandom is a request type to build a corresponding pipeline
	TrafficUserRandom
	//TrafficUserQuery is a request type to build a corresponding pipeline
	TrafficUserQuery
)

type getFunc func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{})
type setFunc func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error)

type request struct {
	ctx    context.Context
	idx    int
	proxy  *dosproxy.DosproxySession
	cr     *commitreveal.CommitrevealSession
	f      setFunc
	params []interface{}
	reply  chan *response
}

type response struct {
	idx int
	tx  *types.Transaction
	err error
}

type ethAdaptor struct {
	bridgeAddr       common.Address
	proxyAddr        common.Address
	commitRevealAddr common.Address
	httpUrls         []string
	wsUrls           []string
	key              *keystore.Key
	auth             *bind.TransactOpts

	proxies    []*dosproxy.DosproxySession
	crs        []*commitreveal.CommitrevealSession
	clients    []*ethclient.Client
	ctx        context.Context
	cancelFunc context.CancelFunc
	reqQueue   chan *request
	logger     log.Logger
}

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

//NewEthAdaptor creates an eth implemention of ProxyAdapter
func NewEthAdaptor(key *keystore.Key, bridgeAddr string, urls []string) (adaptor *ethAdaptor, err error) {
	var httpUrls []string
	var wsUrls []string
	for _, url := range urls {
		if strings.Contains(url, "http") {
			httpUrls = append(httpUrls, url)
		} else if strings.Contains(url, "ws") {
			wsUrls = append(wsUrls, url)
		}
	}
	fmt.Println("httpUrls ", httpUrls)
	fmt.Println("wsUrls ", wsUrls)

	if !common.IsHexAddress(bridgeAddr) {
		return nil, errors.New("bridge address is not a valid hex address")
	}
	adaptor = &ethAdaptor{}
	adaptor.httpUrls = httpUrls
	adaptor.wsUrls = wsUrls
	adaptor.bridgeAddr = common.HexToAddress(bridgeAddr)
	adaptor.key = key
	adaptor.logger = log.New("module", "EthProxy")
	//
	adaptor.ctx, adaptor.cancelFunc = context.WithCancel(context.Background())
	adaptor.auth = bind.NewKeyedTransactor(adaptor.key.PrivateKey)
	adaptor.auth.GasPrice = big.NewInt(20000000000) //1 Gwei
	adaptor.auth.GasLimit = uint64(6000000)
	adaptor.auth.Context = adaptor.ctx
	return
}

//End close the connection to eth and release all resources
func (e *ethAdaptor) Close() {
	e.cancelFunc()
	e.clients = nil
	e.proxies = nil
	e.crs = nil
	e.reqQueue = nil
	return
}

func (e *ethAdaptor) Connect(ctx context.Context) (err error) {
	clients := DialToEth(ctx, e.wsUrls)
	if len(clients) == 0 {
		return errors.New("No reachable geth client")
	}

	var bridge *dosbridge.Dosbridge
	for client := range clients {
		var err error
		if bridge == nil {
			bridge, err = dosbridge.NewDosbridge(e.bridgeAddr, client)
			if err != nil {
				e.logger.Error(err)
				continue
			}
			e.proxyAddr, err = bridge.GetProxyAddress(&bind.CallOpts{Context: ctx})
			if err != nil {
				e.logger.Error(err)
				continue
			}
			e.commitRevealAddr, err = bridge.GetCommitRevealAddress(&bind.CallOpts{Context: ctx})
			if err != nil {
				e.logger.Error(err)
				continue
			}
		}

		p, er := dosproxy.NewDosproxy(e.proxyAddr, client)
		if er != nil {
			fmt.Println("NewDosproxy err ", er)
			e.logger.Error(er)
			continue
		}
		c, er := commitreveal.NewCommitreveal(e.commitRevealAddr, client)
		if er != nil {
			e.logger.Error(er)
			err = er
			continue
		}
		e.clients = append(e.clients, client)
		e.proxies = append(e.proxies, &dosproxy.DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: e.ctx}, TransactOpts: *e.auth})
		e.crs = append(e.crs, &commitreveal.CommitrevealSession{Contract: c, CallOpts: bind.CallOpts{Context: e.ctx}, TransactOpts: *e.auth})
	}

	if len(e.proxies) == 0 || len(e.crs) == 0 {
		return errors.New("No reachable proxy or cr client")
	}
	e.reqLoop()
	return
}

func (e *ethAdaptor) GetTimeoutCtx(t time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(e.ctx, t)
}

func (e *ethAdaptor) UpdateWsUrls(urls []string) {
	e.wsUrls = nil
	e.wsUrls = urls
}

func (e *ethAdaptor) reqLoop() {
	go func() {
		defer fmt.Println("reqLoop exit")
		e.reqQueue = make(chan *request)
		defer close(e.reqQueue)

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
					close(req.reply)
				}(req, resp)
			case <-e.ctx.Done():
				return
			}
		}
	}()
}

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
	for {
		select {
		case val, ok := <-outc:
			if !ok {
				return nil, nil
			}
			return val, nil
		case err, ok := <-errc:
			if !ok {
				continue
			}
			fmt.Println("get err", err, " stack ", stack.Trace().TrimRuntime())
			e.logger.Error(err.(error))
		case <-ctx.Done():
			return nil, errors.New("Timeout")
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
				case resp := <-pre:
					//Request has been fulfulled by previous sendRequest or
					//transaction failed so delete the whole requestSend chain
					if resp.err == nil ||
						strings.Contains(resp.err.Error(), "transaction failed") {
						select {
						case out <- resp:
						case <-ctx.Done():
						}
						return
					}
					fmt.Println("Switch to ", idx, " Client to handle request because of e ,", resp.err)
				}
			}
			r.reply = make(chan *response)
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

// AddToWhitelist is a wrap function that build a pipeline to set groupToPick
func (e *ethAdaptor) AddToWhitelist(ctx context.Context, addr common.Address) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		if len(p) != 1 {
			err = errors.New("Invalid parameter")
			return
		}
		if addr, ok := p[0].(common.Address); ok {
			tx, err = cr.AddToWhitelist(addr)
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, addr)
	reply := e.set(ctx, params, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("AddToWhitelist response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("AddToWhitelist error ", r.err)
			}
		}
	case <-ctx.Done():
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
	reply := e.set(ctx, params, f)

	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("StartCommitReveal response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("StartCommitReveal error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return
}

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

	reply := e.set(ctx, params, f)

	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("SeGroupToPick response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("SeGroupToPick error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return

}

// RegisterNewNode is a wrap function that build a pipeline to call RegisterNewNode
func (e *ethAdaptor) RegisterNewNode(ctx context.Context) (err error) {
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.RegisterNewNode()
		return
	}
	defer e.logger.TimeTrack(time.Now(), "RegisterNewNode", nil)
	reply := e.set(ctx, nil, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("RegisterNewNode response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("RegisterNewNode error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return
}

// SignalRandom is a wrap function that build a pipeline to call SignalRandom
func (e *ethAdaptor) SignalRandom(ctx context.Context) (err error) {
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.SignalRandom()
		return
	}

	reply := e.set(ctx, nil, f)

	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("SignalRandom response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("SignalRandom error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return
}

// SignalGroupFormation is a wrap function that build a pipeline to call SignalGroupFormation
func (e *ethAdaptor) SignalGroupFormation(ctx context.Context) (err error) {
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.SignalGroupFormation()
		return
	}

	reply := e.set(ctx, nil, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("SignalGroupFormation response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("SignalGroupFormation error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return
}

// SignalGroupDissolve is a wrap function that build a pipeline to call SignalGroupDissolve
func (e *ethAdaptor) SignalGroupDissolve(ctx context.Context) (err error) {
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.SignalGroupDissolve()
		return
	}

	reply := e.set(ctx, nil, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("SignalGroupDissolve response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("SignalGroupDissolve error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return

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

	reply := e.set(ctx, params, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("SignalBootstrap response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("SignalBootstrap error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return

}

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

	reply := e.set(ctx, params, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("SetGroupSize response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("SetGroupSize error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return
}

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

	reply := e.set(ctx, params, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("SetGroupSize response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("SetGroupSize error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return
}

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

	reply := e.set(ctx, params, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("SetGroupingThreshold response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("SetGroupingThreshold error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return
}

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
				if err != nil {
					fmt.Println("inter Commit err ", err)
				} else {
					fmt.Println("inter Commit tx ", fmt.Sprintf("%x", tx.Hash()))
				}
			}
		}
		return
	}
	// define parameters
	var params []interface{}
	params = append(params, cid)
	params = append(params, commitment)
	reply := e.set(ctx, params, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("Commit response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("Commit error ", r.err)
			}
		}
	case <-ctx.Done():
		fmt.Println("Commit ctx.Done")
	}
	return
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
	reply := e.set(ctx, params, f)
	select {
	case r, ok := <-reply:
		if ok {
			err = r.err
			if r.err == nil {
				fmt.Println("Reveal response ", fmt.Sprintf("%x", r.tx.Hash()))
			} else {
				fmt.Println("Reveal error ", r.err)
			}
		}
	case <-ctx.Done():
	}
	return
}

// RegisterGroupPubKey is a wrap function that build a pipeline to call RegisterGroupPubKey
func (e *ethAdaptor) RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc chan error) {
	fmt.Println("RegisterGroupPubKey")
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case idPubkey, ok := <-IdWithPubKeys:
			if !ok {
				return
			}
			defer e.logger.TimeTrack(time.Now(), "RegisterGroupPubKey", map[string]interface{}{"GroupID": fmt.Sprintf("%x", idPubkey[0])})
			//			defer logger.TimeTrack(time.Now(), "askMembers", nil)
			fmt.Println("RegisterGroupPubKey got pubkey")

			// define how to parse parameters and execute proxy function
			f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
				err = errors.New("Invalid parameter")
				if len(p) != 1 {
					return
				}
				fmt.Println("RegisterGroupPubKey func")

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
			params = append(params, idPubkey)
			reply := e.set(ctx, params, f)

			select {
			case r, ok := <-reply:
				if ok {
					if r.err == nil {
						fmt.Println("RegisterGroupPubKey response ", fmt.Sprintf("%x", r.tx.Hash()))
					} else {
						fmt.Println("RegisterGroupPubKey error ", r.err)
						select {
						case errc <- r.err:
						case <-ctx.Done():
						}
					}
				}
			case <-ctx.Done():
			}
			return
		case <-ctx.Done():
			return
		}
	}()
	return
}

// SetRandomNum is a wrap function that build a pipeline to call SetRandomNum
func (e *ethAdaptor) SetRandomNum(ctx context.Context, signatures chan *vss.Signature) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
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
						tx, err = proxy.UpdateRandomness(sig, 0)
					case <-ctx.Done():
						err = ctx.Err()
					}
				}
				return
			}
			var params []interface{}
			params = append(params, signature)
			reply := e.set(ctx, params, f)
			for {
				select {
				case r, ok := <-reply:
					if ok {
						if r.err == nil {
							fmt.Println("RegisterGroupPubKey response ", fmt.Sprintf("%x", r.tx.Hash()))
						} else {
							fmt.Println("RegisterGroupPubKey error ", r.err)
							select {
							case errc <- r.err:
							case <-ctx.Done():
							}
						}
					}
					return
				case <-ctx.Done():
					return
				}
			}
		case <-ctx.Done():
			return
		}
	}()
	return
}

// DataReturn is a wrap function that build a pipeline to call DataReturn
func (e *ethAdaptor) DataReturn(ctx context.Context, signatures chan *vss.Signature) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			// define how to parse parameters and execute proxy function
			f := func(ctx context.Context, proxy *dosproxy.DosproxySession, cr *commitreveal.CommitrevealSession, p []interface{}) (tx *types.Transaction, err error) {
				err = errors.New("Invalid parameter")
				if len(p) != 1 {
					return
				}
				if sign, ok := p[0].(*vss.Signature); ok {
					select {
					default:
						requestId := new(big.Int).SetBytes(signature.RequestId)
						trafficType := uint8(signature.Index)
						result := signature.Content
						x, y := sign.ToBigInt()
						sig := [2]*big.Int{x, y}
						tx, err = proxy.TriggerCallback(requestId, trafficType, result, sig, 0)
					case <-ctx.Done():
						err = ctx.Err()
					}
				}
				return
			}
			var params []interface{}
			params = append(params, signature)
			reply := e.set(ctx, params, f)

			select {
			case r, ok := <-reply:
				if ok {
					if r.err == nil {
						fmt.Println("DataReturn response ", fmt.Sprintf("%x", r.tx.Hash()))
					} else {
						fmt.Println("DataReturn error ", r.err)
						select {
						case errc <- r.err:
						case <-ctx.Done():
						}
					}
				}
				return
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
			return
		}
	}()
	return
}

// SubscribeEvent is a log subscription operation
func (e *ethAdaptor) SubscribeEvent(subscribeTypes []int) (chan interface{}, chan error) {
	var eventList []chan interface{}
	var errcs []chan interface{}
	for _, subscribeType := range subscribeTypes {
		if subscribeType >= SubscribeCommitrevealLogStartCommitreveal {
			for i := 0; i < len(e.proxies); i++ {
				fmt.Println("Subscribe CR Event ", i)
				cr := e.crs[i]
				if cr == nil {
					continue
				}
				ctx := e.ctx
				if ctx == nil {
					continue
				}
				out, errc := crTable[subscribeType](ctx, cr)
				eventList = append(eventList, out)
				errcs = append(errcs, errc)
			}
		} else {
			for i := 0; i < len(e.proxies); i++ {
				fmt.Println("SubscribeEvent ", i, subscribeType)
				proxy := e.proxies[i]
				if proxy == nil {
					continue
				}
				ctx := e.ctx
				if ctx == nil {
					continue
				}
				out, errc := proxyTable[subscribeType](ctx, proxy)
				eventList = append(eventList, out)
				errcs = append(errcs, errc)
			}
		}
	}
	return firstEvent(e.ctx, merge(e.ctx, eventList...)), convertToError(e.ctx, merge(e.ctx, errcs...))
}

// LastRandomness return the last system random number
func (e *ethAdaptor) LastRandomness(ctx context.Context) (result *big.Int, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			val, err := proxy.LastRandomness()
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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
			val, err := proxy.GroupSize()
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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
			val, err := proxy.GetWorkingGroupSize()
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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
			val, err := proxy.NumPendingGroups()
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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
			val, err := proxy.RefreshSystemRandomHardLimit()
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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
			val, err := proxy.NumPendingNodes()
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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
			val, err := proxy.GetExpiredWorkingGroupSize()
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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
			val, err := proxy.GroupToPick()
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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
			val, err := proxy.LastUpdatedBlock()
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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
				val, err := proxy.PendingNodeList(id)
				if err != nil {
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				}
				select {
				case <-ctx.Done():
				case outc <- val:
				}
			} else {
				fmt.Printf("Type %T ", p)
				select {
				case <-ctx.Done():
				case errc <- errors.New("cast error"):
				}
			}
		}()
		return outc, errc
	}
	addr := common.Address{}

	addr.SetBytes(id)
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
				val, err := proxy.GetGroupPubKey(idx)
				if err != nil {
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				}
				select {
				case <-ctx.Done():
				case outc <- val:
				}
			} else {
				select {
				case <-ctx.Done():
				case errc <- errors.New("cast error"):
				}
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
				for {
					val, err := client.PendingNonceAt(ctx, address)
					if err != nil {
						fmt.Println("PendingNonce err ", err)
						select {
						case <-ctx.Done():
							return
						case errc <- err:
						}
					} else {
						fmt.Println("PendingNonce ", val)
						select {
						case <-ctx.Done():
						case outc <- val:
						}
						return
					}
				}
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, e.key.Address)
	fmt.Println("PendingNonce ", vr, ve)
	if v, ok := vr.(*big.Int); ok {
		result = v.Uint64()
	}
	if v, ok := ve.(error); ok {
		err = v
	}
	return

}

// Balance returns the wei balance of the node account.
func (e *ethAdaptor) Balance(ctx context.Context, id []byte) (result *big.Float, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if address, ok := p.(common.Address); ok {
				wei, err := client.BalanceAt(context.Background(), address, nil)
				balance := new(big.Float)
				balance.SetString(wei.String())
				balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))
				if err != nil {
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				}
				select {
				case <-ctx.Done():
				case outc <- balance:
				}
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, common.BytesToAddress(id))
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
			val, err := client.HeaderByNumber(ctx, nil)
			if err != nil {
				select {
				case <-ctx.Done():
				case errc <- err:
				}
				return
			}
			select {
			case <-ctx.Done():
			case outc <- val:
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

// Address gets the string representation of the underlying address.
func (e *ethAdaptor) Address() (addr common.Address) {
	return e.key.Address
}

func convertToError(ctx context.Context, i chan interface{}) (out chan error) {
	out = make(chan error)
	go func() {
		defer close(out)
		for {
			select {
			case e, ok := <-i:
				if !ok {
					return
				}
				if err, ok := e.(error); ok {
					select {
					case out <- err:
					case <-ctx.Done():
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

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
