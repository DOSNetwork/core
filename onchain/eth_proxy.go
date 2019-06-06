package onchain

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"runtime/debug"
	"strings"
	"time"

	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain/commitreveal"
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
type setFunc func(ctx context.Context, proxy *dosproxy.DosproxySession, p []interface{}) (tx *types.Transaction, err error)

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
	proxyAddr        string
	commitRevealAddr string
	gethUrls         []string
	eventUrls        []string
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

//NewEthAdaptor creates an eth implemention of ProxyAdapter
func NewEthAdaptor(credentialPath, passphrase, proxyAddr, commitRevealAddr string, urls []string) (adaptor *ethAdaptor, err error) {
	var httpUrls []string
	var wsUrls []string
	for _, url := range urls {
		if strings.Contains(url, "http") {
			httpUrls = append(httpUrls, url)
		} else if strings.Contains(url, "ws") {
			wsUrls = append(wsUrls, url)
		}
	}
	fmt.Println("gethUrls ", httpUrls)
	fmt.Println("eventUrls ", wsUrls)

	adaptor = &ethAdaptor{}
	adaptor.gethUrls = httpUrls
	adaptor.eventUrls = wsUrls
	adaptor.proxyAddr = proxyAddr
	debug.FreeOSMemory()
	//Read Ethereum keystore
	key, err := ReadEthKey(credentialPath, passphrase)
	if err != nil {
		return
	}
	adaptor.key = key
	debug.FreeOSMemory()

	//Use account address as ID to init log module
	log.Init(key.Address.Bytes()[:])
	adaptor.logger = log.New("module", "EthProxy")

	adaptor.ctx, adaptor.cancelFunc = context.WithCancel(context.Background())
	adaptor.auth = bind.NewKeyedTransactor(key.PrivateKey)
	adaptor.auth.GasPrice = big.NewInt(2000000000) //0.2 Gwei
	adaptor.auth.GasLimit = uint64(6000000)
	adaptor.auth.Context = adaptor.ctx
	adaptor.reqQueue = make(chan *request)
	adaptor.start()
	return
}

//End close the connection to eth and release all resources
func (e *ethAdaptor) End() {
	//e.cancel()
	//e.c.Close()
	return
}

func (e *ethAdaptor) start() (err error) {
	//
	clients := DialToEth(context.Background(), e.eventUrls)
	for client := range clients {
		p, er := dosproxy.NewDosproxy(common.HexToAddress(e.proxyAddr), client)
		if er != nil {
			fmt.Println("NewDosproxy err ", er)
			e.logger.Error(er)
			continue
		}
		c, er := commitreveal.NewCommitreveal(common.HexToAddress(e.commitRevealAddr), client)
		if er != nil {
			e.logger.Error(er)
			err = er
			continue
		}
		e.clients = append(e.clients, client)
		e.proxies = append(e.proxies, &dosproxy.DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: e.ctx}, TransactOpts: *e.auth})
		e.crs = append(e.crs, &commitreveal.CommitrevealSession{Contract: c, CallOpts: bind.CallOpts{Context: e.ctx}, TransactOpts: *e.auth})
	}

	if len(e.proxies) == 0 {
		fmt.Println("No any working eth client ", len(e.clients), len(e.proxies))
		return
	}
	e.reqLoop()
	return
}

func (e *ethAdaptor) reqLoop() {
	go func() {
		defer fmt.Println("reqLoop exit")
		defer close(e.reqQueue)

		for {
			select {
			case req := <-e.reqQueue:
				fmt.Println("reqLoop got req i", req.idx)
				tx, err := req.f(req.ctx, req.proxy, req.params)
				resp := &response{req.idx, tx, err}
				select {
				case req.reply <- resp:
				case <-req.ctx.Done():
				}
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
				fmt.Println("get !ok")
				return nil, nil
			}
			fmt.Println("get ", val)
			return val, nil
		case err, ok := <-errc:
			if !ok {
				fmt.Println("get errc !ok")
				continue
			}
			fmt.Println("get err", err)
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
						fmt.Println("set err e ,", resp.err)
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
		r := &request{ctx, i, proxy, nil, setF, params, nil}
		reply = f(ctx, i, reply, r)
	}

	return
}

// SetGroupToPick is a wrap function that build a pipeline to set groupToPick
func (e *ethAdaptor) SetGroupToPick(ctx context.Context, groupToPick uint64) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, p []interface{}) (tx *types.Transaction, err error) {
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
	for {
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
			return
		case <-ctx.Done():
			return
		}
	}
}

// RegisterNewNode is a wrap function that build a pipeline to call RegisterNewNode
func (e *ethAdaptor) RegisterNewNode(ctx context.Context) (err error) {
	// define how to parse parameters and execute proxy function
	f := func(ctx context.Context, proxy *dosproxy.DosproxySession, p []interface{}) (tx *types.Transaction, err error) {
		tx, err = proxy.RegisterNewNode()
		return
	}

	reply := e.set(ctx, nil, f)
	for {
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
			return
		case <-ctx.Done():
			return
		}
	}
}

// SignalRandom is a wrap function that build a pipeline to call SignalRandom
func (e *ethAdaptor) SignalRandom(ctx context.Context) (errc chan error) {

	return
}

// SignalGroupFormation is a wrap function that build a pipeline to call SignalGroupFormation
func (e *ethAdaptor) SignalGroupFormation(ctx context.Context) (errc chan error) {

	return
}

// SignalGroupDissolve is a wrap function that build a pipeline to call SignalGroupDissolve
func (e *ethAdaptor) SignalGroupDissolve(ctx context.Context) (errc chan error) {

	return
}

// SignalBootstrap is a wrap function that build a pipeline to call SignalBootstrap
func (e *ethAdaptor) SignalBootstrap(ctx context.Context, cid uint64) (errc chan error) {

	return
}

// Commit is a wrap function that build a pipeline to call Commit
func (e *ethAdaptor) Commit(ctx context.Context, cid *big.Int, commitment [32]byte) (errc chan error) {

	return
}

// Reveal is a wrap function that build a pipeline to call Reveal
func (e *ethAdaptor) Reveal(ctx context.Context, cid *big.Int, secret *big.Int) (errc chan error) {

	return
}

// RegisterGroupPubKey is a wrap function that build a pipeline to call RegisterGroupPubKey
func (e *ethAdaptor) RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc chan error) {
	return
}

// SetRandomNum is a wrap function that build a pipeline to call SetRandomNum
func (e *ethAdaptor) SetRandomNum(ctx context.Context, signatures chan *vss.Signature) (errc chan error) {
	return
}

// DataReturn is a wrap function that build a pipeline to call DataReturn
func (e *ethAdaptor) DataReturn(ctx context.Context, signatures chan *vss.Signature) (errc chan error) {
	return
}

// SetGroupSize is a wrap function that build a pipeline to call SetGroupSize
func (e *ethAdaptor) SetGroupSize(ctx context.Context, size uint64) (errc chan error) {

	return
}

// SetGroupMaturityPeriod is a wrap function that build a pipeline to call SetGroupMaturityPeriod
func (e *ethAdaptor) SetGroupMaturityPeriod(ctx context.Context, period uint64) (errc chan error) {
	return
}

// SetGroupingThreshold is a wrap function that build a pipeline to call SetGroupingThreshold
func (e *ethAdaptor) SetGroupingThreshold(ctx context.Context, threshold uint64) (errc chan error) {

	return
}

// SubscribeEvent is a log subscription operation
func (e *ethAdaptor) SubscribeEvent(subscribeType int) (chan interface{}, chan error) {
	var eventList []chan interface{}
	var errcs []chan interface{}
	if subscribeType == SubscribeCommitrevealLogStartCommitreveal ||
		subscribeType == SubscribeCommitrevealLogCommit ||
		subscribeType == SubscribeCommitrevealLogReveal ||
		subscribeType == SubscribeCommitrevealLogRandom {
		for i := 0; i < len(e.proxies); i++ {
			fmt.Println("SubscribeEvent ", i)
			cr := e.crs[i]
			if cr == nil {
				continue
			}
			ctx := e.ctx
			if ctx == nil {
				continue
			}
			out, errc := subscribeCREvent(ctx, cr, subscribeType)
			eventList = append(eventList, out)
			errcs = append(errcs, errc)
		}
	} else {
		for i := 0; i < len(e.proxies); i++ {
			fmt.Println("SubscribeEvent ", i)
			proxy := e.proxies[i]
			if proxy == nil {
				continue
			}
			ctx := e.ctx
			if ctx == nil {
				continue
			}
			out, errc := subscribeEvent(ctx, proxy, subscribeType)
			eventList = append(eventList, out)
			errcs = append(errcs, errc)
		}
	}
	return firstEvent(e.ctx, merge(e.ctx, eventList...)), convertToError(e.ctx, merge(e.ctx, errcs...))
}

func subscribeCREvent(ctx context.Context, cr *commitreveal.CommitrevealSession, subscribeType int) (chan interface{}, chan interface{}) {
	out := make(chan interface{})
	errc := make(chan interface{})
	opt := &bind.WatchOpts{}

	switch subscribeType {
	case SubscribeCommitrevealLogStartCommitreveal:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogStartCommitReveal{
						Cid:             i.Cid,
						StartBlock:      i.StartBlock,
						CommitDuration:  i.CommitDuration,
						RevealDuration:  i.RevealDuration,
						RevealThreshold: i.RevealThreshold,
						Tx:              i.Raw.TxHash.Hex(),
						BlockN:          i.Raw.BlockNumber,
						Removed:         i.Raw.Removed,
						Raw:             i.Raw,
					}
				}
			}
		}()
	case SubscribeCommitrevealLogCommit:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogCommit{
						Cid:        i.Cid,
						From:       i.From,
						Commitment: i.Commitment,
						Tx:         i.Raw.TxHash.Hex(),
						BlockN:     i.Raw.BlockNumber,
						Removed:    i.Raw.Removed,
						Raw:        i.Raw,
					}
				}
			}
		}()
	case SubscribeCommitrevealLogReveal:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogReveal{
						Cid:     i.Cid,
						From:    i.From,
						Secret:  i.Secret,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeCommitrevealLogRandom:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogRandom{
						Cid:     i.Cid,
						Random:  i.Random,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	}
	return out, errc
}

func subscribeEvent(ctx context.Context, proxy *dosproxy.DosproxySession, subscribeType int) (chan interface{}, chan interface{}) {
	out := make(chan interface{})
	errc := make(chan interface{})
	opt := &bind.WatchOpts{}

	switch subscribeType {
	case SubscribeLogUpdateRandom:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogUpdateRandom{
						LastRandomness:    i.LastRandomness,
						DispatchedGroupId: i.DispatchedGroupId,
						Tx:                i.Raw.TxHash.Hex(),
						BlockN:            i.Raw.BlockNumber,
						Removed:           i.Raw.Removed,
						Raw:               i.Raw,
					}
				}
			}
		}()
	case SubscribeLogUrl:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogUrl{
						QueryId:           i.QueryId,
						Timeout:           i.Timeout,
						DataSource:        i.DataSource,
						Selector:          i.Selector,
						Randomness:        i.Randomness,
						DispatchedGroupId: i.DispatchedGroupId,
						Tx:                i.Raw.TxHash.Hex(),
						BlockN:            i.Raw.BlockNumber,
						Removed:           i.Raw.Removed,
						Raw:               i.Raw,
					}
				}
			}
		}()
	case SubscribeLogRequestUserRandom:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogRequestUserRandom{
						RequestId:            i.RequestId,
						LastSystemRandomness: i.LastSystemRandomness,
						UserSeed:             i.UserSeed,
						DispatchedGroupId:    i.DispatchedGroupId,
						Tx:                   i.Raw.TxHash.Hex(),
						BlockN:               i.Raw.BlockNumber,
						Removed:              i.Raw.Removed,
						Raw:                  i.Raw,
					}
				}
			}
		}()
	case SubscribeLogValidationResult:
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
					out <- &LogValidationResult{
						TrafficType: i.TrafficType,
						TrafficId:   i.TrafficId,
						Message:     i.Message,
						Signature:   i.Signature,
						PubKey:      i.PubKey,
						Pass:        i.Pass,
						Version:     i.Version,
						Tx:          i.Raw.TxHash.Hex(),
						BlockN:      i.Raw.BlockNumber,
						Removed:     i.Raw.Removed,
						Raw:         i.Raw,
					}
				}
			}
		}()
	case SubscribeLogInsufficientPendingNode:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogInsufficientPendingNode{
						NumPendingNodes: i.NumPendingNodes,
						Tx:              i.Raw.TxHash.Hex(),
						BlockN:          i.Raw.BlockNumber,
						Removed:         i.Raw.Removed,
						Raw:             i.Raw,
					}
				}
			}
		}()
	case SubscribeLogInsufficientWorkingGroup:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogInsufficientWorkingGroup{
						NumWorkingGroups: i.NumWorkingGroups,
						Tx:               i.Raw.TxHash.Hex(),
						BlockN:           i.Raw.BlockNumber,
						Removed:          i.Raw.Removed,
						Raw:              i.Raw,
					}
				}
			}
		}()
	case SubscribeLogGroupingInitiated:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogGroupingInitiated{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyUpdateGroupToPick:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					fmt.Println("SubscribeDosproxyUpdateGroupToPick err ", err)
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogUpdateGroupToPick{
						Tx:      i.Raw.TxHash.Hex(),
						OldNum:  i.OldNum,
						NewNum:  i.NewNum,
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeLogGrouping:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogGrouping{
						GroupId: i.GroupId,
						NodeId:  i.NodeId,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeLogPublicKeyAccepted:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogPublicKeyAccepted{
						GroupId:          i.GroupId,
						WorkingGroupSize: i.NumWorkingGroups,
						Tx:               i.Raw.TxHash.Hex(),
						BlockN:           i.Raw.BlockNumber,
						Removed:          i.Raw.Removed,
						Raw:              i.Raw,
					}
				}
			}
		}()
	case SubscribeLogPublicKeySuggested:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogPublicKeySuggested{
						GroupId: i.GroupId,
						Count:   i.PubKeyCount,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeLogGroupDissolve:
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
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogGroupDissolve{
						GroupId: i.GroupId,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	}
	return out, errc
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
				select {
				case <-ctx.Done():
				case errc <- errors.New("cast error"):
				}
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, id)
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
						b, err := client.BalanceAt(ctx, e.key.Address, nil)
						if err != nil {
							//Post http i/o timeout
							fmt.Println("BalanceAt err ", err)
						} else {
							fmt.Println("BalanceAt id ", b)
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
func (e *ethAdaptor) Balance(ctx context.Context) (result *big.Float, err error) {
	f := func(ctx context.Context, client *ethclient.Client, proxy *dosproxy.DosproxySession, p interface{}) (chan interface{}, chan interface{}) {
		outc := make(chan interface{})
		errc := make(chan interface{})
		go func() {
			defer close(outc)
			defer close(errc)
			if address, ok := p.(common.Address); ok {
				val, err := client.BalanceAt(context.Background(), address, nil)
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
			}
		}()
		return outc, errc
	}

	vr, ve := e.get(ctx, f, e.key.Address)
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
func (e *ethAdaptor) Address() (addr []byte) {
	return e.key.Address.Bytes()
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
			var bytes []byte
			var blkNum uint64
			var event interface{}
			var ok bool
			var removed bool
			select {
			case event, ok = <-source:
				if ok {
					switch content := event.(type) {
					case *LogUpdateRandom:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogRequestUserRandom:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogUrl:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogValidationResult:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogGrouping:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogPublicKeyAccepted:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogPublicKeySuggested:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogGroupDissolve:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogUpdateGroupToPick:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogStartCommitReveal:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogCommit:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogReveal:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogRandom:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						removed = content.Removed
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					}
				} else {
					return
				}
			}
			if removed {
				continue
			}
			nHash := sha256.Sum256(bytes)
			identity := string(nHash[:])

			if visited[identity] == 0 {
				visited[identity] = blkNum
				select {
				case out <- event:
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
	}()

	return
}
