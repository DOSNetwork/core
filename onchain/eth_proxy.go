package onchain

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SubscribeDOSProxyLogUpdateRandom = iota
	SubscribeDOSProxyLogRequestUserRandom
	SubscribeDOSProxyLogUrl
	SubscribeDOSProxyLogValidationResult
	SubscribeDOSProxyLogGrouping
	SubscribeDOSProxyLogPublicKeyAccepted
	SubscribeDOSProxyLogPublicKeySuggested
	SubscribeDOSProxyLogGroupDismiss
	SubscribeDOSProxyLogInsufficientPendingNode
	SubscribeDOSProxyLogInsufficientWorkingGroup
	SubscribeDOSProxyLogNoWorkingGroup
)

var logger log.Logger

// TODO: Move constants to some unified places.
const (
	TrafficSystemRandom = iota // 0
	TrafficUserRandom
	TrafficUserQuery
)

const (
	LogBlockDiff        = 1
	LogCheckingInterval = 15 //in second
	SubscribeTimeout    = 60 //in second
)

type EthAdaptor struct {
	wProxy      *dosproxy.DOSProxySession
	wClient     *ethclient.Client
	key         *keystore.Key
	wCtx        context.Context
	wCancelFunc context.CancelFunc
	wReqQueue   chan interface{}

	//read only DOSProxy
	rProxies     []*dosproxy.DOSProxy
	rClients     []*ethclient.Client
	rCtxs        []context.Context
	rCancelFuncs []context.CancelFunc
}

type Request struct {
	ctx    context.Context
	f      func() (*types.Transaction, error)
	result chan Reply
}

type ReqGrouping struct {
	ctx        context.Context
	candidates []common.Address
	size       *big.Int
	f          func([]common.Address, *big.Int) (*types.Transaction, error)
	result     chan Reply
}

type ReqSetRandomNum struct {
	ctx     context.Context
	sig     [2]*big.Int
	version uint8
	f       func([2]*big.Int, uint8) (*types.Transaction, error)
	result  chan Reply
}

type ReqSetPublicKey struct {
	ctx     context.Context
	groupId *big.Int
	pubKey  [4]*big.Int
	f       func(*big.Int, [4]*big.Int) (*types.Transaction, error)
	result  chan Reply
}

type ReqTriggerCallback struct {
	ctx         context.Context
	requestId   *big.Int
	trafficType uint8
	content     []byte
	sig         [2]*big.Int
	version     uint8
	f           func(*big.Int, uint8, []byte, [2]*big.Int, uint8) (*types.Transaction, error)
	result      chan Reply
}

type Reply struct {
	tx    *types.Transaction
	nonce uint64
	err   error
}

func NewEthAdaptor(credentialPath, passphrase, proxyAddr string, gethUrls []string) (adaptor *EthAdaptor, err error) {
	var ok bool
	var proxy *dosproxy.DOSProxy
	key, err := ReadEthKey(credentialPath, passphrase)
	if err != nil {
		logger.Error(err)
		return
	}
	log.Init(key.Address.Bytes()[:])

	//Init log module with nodeID that is an onchain account address
	if logger == nil {
		logger = log.New("module", "EthProxy")
	}

	clients, errs := DialToEth(context.Background(), gethUrls)
	go func() {
		for err := range errs {
			logger.Error(err)
		}
	}()

	adaptor = &EthAdaptor{}
	adaptor.key = key
	//Use first client as sender
	adaptor.wClient, ok = <-clients
	if !ok {
		err = errors.New("No any working eth client")
		return
	}
	proxy, err = dosproxy.NewDOSProxy(common.HexToAddress(proxyAddr), adaptor.wClient)
	if err != nil {
		adaptor.wClient.Close()
		logger.Error(err)
		return
	}
	adaptor.wCtx, adaptor.wCancelFunc = context.WithCancel(context.Background())

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(GASLIMIT)
	auth.Context = adaptor.wCtx
	nonce, err := adaptor.wClient.PendingNonceAt(adaptor.wCtx, key.Address)
	if err == nil {
		auth.Nonce = new(big.Int).SetUint64(nonce)
	}
	adaptor.wProxy = &dosproxy.DOSProxySession{Contract: proxy, CallOpts: bind.CallOpts{Context: adaptor.wCtx}, TransactOpts: *auth}
	adaptor.wReqQueue = make(chan interface{})

	for client := range clients {
		p, err := dosproxy.NewDOSProxy(common.HexToAddress(proxyAddr), client)
		if err != nil {
			logger.Error(err)
			continue
		}
		ctx, cancelFunc := context.WithCancel(context.Background())
		adaptor.rClients = append(adaptor.rClients, client)
		adaptor.rProxies = append(adaptor.rProxies, p)
		adaptor.rCtxs = append(adaptor.rCtxs, ctx)
		adaptor.rCancelFuncs = append(adaptor.rCancelFuncs, cancelFunc)
	}
	if len(adaptor.rProxies) == 0 {
		adaptor.wClient.Close()
		err = errors.New("No any working eth client for event tracking")
		//return
	}

	adaptor.reqLoop()

	return
}

func (e *EthAdaptor) End() {
	//e.cancel()
	//e.c.Close()
	return
}

func (e *EthAdaptor) reqLoop() {
	go func() {
		defer fmt.Println("reqLoop exit")
		defer close(e.wReqQueue)

		for {
			select {
			case req := <-e.wReqQueue:
				var tx *types.Transaction
				var err error
				var resultC chan Reply
				var ctx context.Context
				reply := Reply{}

				//Compare with latest pending nonce
				nonce, err := e.wClient.PendingNonceAt(e.wCtx, e.key.Address)
				if err == nil {
					reply.err = err
					nonceBig := new(big.Int).SetUint64(nonce)
					if e.wProxy.TransactOpts.Nonce.Cmp(nonceBig) == -1 {
						e.wProxy.TransactOpts.Nonce = nonceBig
					}
				}

				switch content := req.(type) {
				case *Request:
					tx, err = content.f()
					resultC = content.result
					ctx = content.ctx
				case *ReqGrouping:
					tx, err = content.f(content.candidates, content.size)
					resultC = content.result
					ctx = content.ctx
				case *ReqSetRandomNum:
					tx, err = content.f(content.sig, content.version)
					resultC = content.result
					ctx = content.ctx
				case *ReqTriggerCallback:
					tx, err = content.f(content.requestId, content.trafficType, content.content, content.sig, content.version)
					resultC = content.result
					ctx = content.ctx
				case *ReqSetPublicKey:
					tx, err = content.f(content.groupId, content.pubKey)
					resultC = content.result
					ctx = content.ctx
				}
				if err != nil {
					if err.Error() == "replacement transaction underpriced" ||
						strings.Contains(err.Error(), "known transaction") {
						e.wProxy.TransactOpts.Nonce = e.wProxy.TransactOpts.Nonce.Add(e.wProxy.TransactOpts.Nonce, big.NewInt(1))
					}
					reply.err = err
					resultC <- reply
					continue
				}
				//fmt.Println(reqType, " nonce ", e.s.TransactOpts.Nonce, " size ", size, " tx ", fmt.Sprintf("%x", tx.Hash()))
				reply.tx = tx
				reply.nonce = e.wProxy.TransactOpts.Nonce.Uint64()
				e.wProxy.TransactOpts.Nonce = e.wProxy.TransactOpts.Nonce.Add(e.wProxy.TransactOpts.Nonce, big.NewInt(1))
				select {
				case resultC <- reply:
				case <-ctx.Done():
				}
			case <-e.wCtx.Done():
				return
			}
		}
	}()
}

func (e *EthAdaptor) sendRequest(ctx context.Context, request interface{}, result chan Reply) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		defer close(result)
		//The same account might be used outside of system that causes a nonce conflict.
		//Need to check if tx is confirmed to avoid lost transaction because of nonce conflict
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			e.wReqQueue <- request

			select {
			case reply := <-result:
				err := reply.err
				tx := reply.tx
				//nonce := reply.nonce
				if err != nil {
					if err.Error() == "client is closed" ||
						err.Error() == "EOF" ||
						strings.Contains(err.Error(), "connection refused") ||
						strings.Contains(err.Error(), "use of closed network connection") {
						errc <- err
						return
					} else {
						logger.Error(err)
						f := map[string]interface{}{
							"ErrMsg": err.Error(),
							"Time":   time.Now()}
						logger.Event("SendRequestFail", f)
					}
					continue
				}
				defer logger.TimeTrack(time.Now(), "SendRequest", map[string]interface{}{"RequestId": ctx.Value("RequestID"), "Tx": fmt.Sprintf("%x", tx.Hash())})
				//fmt.Println("Tx", fmt.Sprintf("%x", tx.Hash()))
				err = CheckTransaction(e.wClient, tx)
				if err != nil {
					logger.Error(err)
					if err.Error() == "client is closed" ||
						err.Error() == "EOF" ||
						strings.Contains(err.Error(), "connection refused") ||
						strings.Contains(err.Error(), "use of closed network connection") ||
						strings.Contains(err.Error(), "transaction failed") {
						errc <- err
						return
					} else {
						logger.Error(err)
						f := map[string]interface{}{
							"Tx":     tx,
							"ErrMsg": err.Error(),
							"Time":   time.Now()}
						logger.Event("TransactionFail", f)
					}
					continue
				}
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return errc
}

func (e *EthAdaptor) RegisterNewNode(ctx context.Context) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "RegisterNewNode", nil)
	result := make(chan Reply)
	request := &Request{ctx, e.wProxy.RegisterNewNode, result}
	return e.sendRequest(ctx, request, result)
}

func (e *EthAdaptor) RandomNumberTimeOut(ctx context.Context) (errc <-chan error) {
	result := make(chan Reply)
	request := &Request{ctx, e.wProxy.HandleTimeout, result}
	return e.sendRequest(ctx, request, result)
}

func (e *EthAdaptor) RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case IdWithPubKey, ok := <-IdWithPubKeys:
			if !ok {
				return
			}
			result := make(chan Reply)
			groupId := IdWithPubKey[0]
			var pubKey [4]*big.Int
			copy(pubKey[:], IdWithPubKey[1:])
			request := &ReqSetPublicKey{ctx, groupId, pubKey, e.wProxy.RegisterGroupPubKey, result}
			errChan := e.sendRequest(ctx, request, result)
			err := <-errChan
			errc <- err
			f := map[string]interface{}{
				"DispatchedGroupId": fmt.Sprintf("%x", groupId.Bytes()),
				"DispatchedGroup_1": fmt.Sprintf("%x", pubKey[0].Bytes()),
				"DispatchedGroup_2": fmt.Sprintf("%x", pubKey[1].Bytes()),
				"DispatchedGroup_3": fmt.Sprintf("%x", pubKey[2].Bytes()),
				"DispatchedGroup_4": fmt.Sprintf("%x", pubKey[3].Bytes()),
				"WorkingGroup":      e.GetWorkingGroupSize(),
				"Time":              time.Now()}
			logger.Event("DOS_RegisterGroupPubKey", f)
			return
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

func (e *EthAdaptor) SetRandomNum(ctx context.Context, signatures <-chan *vss.Signature) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			x, y := DecodeSig(signature.Signature)
			result := make(chan Reply)
			request := &ReqSetRandomNum{ctx, [2]*big.Int{x, y}, 0, e.wProxy.UpdateRandomness, result}
			errChan := e.sendRequest(ctx, request, result)
			err := <-errChan
			errc <- err
			return
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

func (e *EthAdaptor) DataReturn(ctx context.Context, signatures <-chan *vss.Signature) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "DataReturn", map[string]interface{}{"RequestId": ctx.Value("RequestID")})

			x, y := DecodeSig(signature.Signature)
			requestId := new(big.Int).SetBytes(signature.RequestId)

			result := make(chan Reply)
			request := &ReqTriggerCallback{ctx, requestId, uint8(signature.Index), signature.Content, [2]*big.Int{x, y}, 0, e.wProxy.TriggerCallback, result}
			errChan := e.sendRequest(ctx, request, result)
			err := <-errChan
			errc <- err
			return
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

func mergeEvents(ctx context.Context, cs ...chan interface{}) chan interface{} {
	var wg sync.WaitGroup
	// We must ensure that the output channel has the capacity to
	// hold as many errors
	// as there are error channels.
	// This will ensure that it never blocks, even
	// if WaitForPipeline returns early.
	out := make(chan interface{}, len(cs))
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls
	// wg.Done.
	output := func(c <-chan interface{}) {
		for n := range c {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	// Start a goroutine to close out once all the output goroutines
	// are done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func (e *EthAdaptor) first(ctx context.Context, sink, source chan interface{}) {
	go func() {
		visited := make(map[string]uint64)
		for {
			var bytes []byte
			var event interface{}
			var blkNum uint64
			select {
			case event = <-source:
				switch content := event.(type) {
				case *DOSProxyLogUpdateRandom:
					bytes = append(bytes, content.Raw.Data...)
					blkNum = content.BlockN
					bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
				case *DOSProxyLogRequestUserRandom:
					bytes = append(bytes, content.Raw.Data...)
					blkNum = content.BlockN
					bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
				case *DOSProxyLogUrl:
					bytes = append(bytes, content.Raw.Data...)
					blkNum = content.BlockN
					bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
				case *DOSProxyLogValidationResult:
					bytes = append(bytes, content.Raw.Data...)
					blkNum = content.BlockN
					bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
				case *DOSProxyLogNoWorkingGroup:
					bytes = append(bytes, content.Raw.Data...)
					blkNum = content.BlockN
					bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
				}
			}
			nHash := sha256.Sum256(bytes)
			identity := string(nHash[:])
			curBlk, err := e.CurrentBlock()
			if err != nil {
				continue
			}

			if visited[identity] == 0 {
				visited[identity] = blkNum
				select {
				case sink <- event:
				case <-ctx.Done():
				}
			} else {
				for k, blkN := range visited {
					if curBlk >= blkN+100 {
						delete(visited, k)
					}
				}
			}
		}
	}()
}

func (e *EthAdaptor) SubscribeEvent(subscribeType int, sink chan interface{}) {

	var eventList []chan interface{}
	for i := 0; i < len(e.rProxies); i++ {
		fmt.Println("SubscribeEvent ", i)
		proxy := e.rProxies[i]
		if proxy == nil {
			continue
		}
		ctx := e.rCtxs[i]
		if ctx == nil {
			continue
		}
		out, _ := subscribeEvent(ctx, proxy, subscribeType)
		eventList = append(eventList, out)
	}
	out, _ := e.PollLogs(subscribeType, 0, 0)
	eventList = append(eventList, out)
	e.first(e.wCtx, sink, mergeEvents(e.wCtx, eventList...))
}

func subscribeEvent(ctx context.Context, proxy *dosproxy.DOSProxy, subscribeType int) (out chan interface{}, errc chan error) {
	out = make(chan interface{})
	errc = make(chan error)
	opt := &bind.WatchOpts{}

	switch subscribeType {
	case SubscribeDOSProxyLogUpdateRandom:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogUpdateRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DOSProxyFilterer.WatchLogUpdateRandom(opt, transitChan)
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
					out <- &DOSProxyLogUpdateRandom{
						LastRandomness:    i.LastRandomness,
						DispatchedGroupId: i.DispatchedGroupId,
						DispatchedGroup:   i.DispatchedGroup,
						Tx:                i.Raw.TxHash.Hex(),
						BlockN:            i.Raw.BlockNumber,
						Removed:           i.Raw.Removed,
						Raw:               i.Raw,
					}
				}
			}
		}()
	case SubscribeDOSProxyLogUrl:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogUrl)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.DOSProxyFilterer.WatchLogUrl(opt, transitChan)
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
					out <- &DOSProxyLogUrl{
						QueryId:           i.QueryId,
						Timeout:           i.Timeout,
						DataSource:        i.DataSource,
						Selector:          i.Selector,
						Randomness:        i.Randomness,
						DispatchedGroupId: i.DispatchedGroupId,
						DispatchedGroup:   i.DispatchedGroup,
						Tx:                i.Raw.TxHash.Hex(),
						BlockN:            i.Raw.BlockNumber,
						Removed:           i.Raw.Removed,
						Raw:               i.Raw,
					}
				}
			}
		}()
	case SubscribeDOSProxyLogRequestUserRandom:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogRequestUserRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.DOSProxyFilterer.WatchLogRequestUserRandom(opt, transitChan)
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
					out <- &DOSProxyLogRequestUserRandom{
						RequestId:            i.RequestId,
						LastSystemRandomness: i.LastSystemRandomness,
						UserSeed:             i.UserSeed,
						DispatchedGroupId:    i.DispatchedGroupId,
						DispatchedGroup:      i.DispatchedGroup,
						Tx:                   i.Raw.TxHash.Hex(),
						BlockN:               i.Raw.BlockNumber,
						Removed:              i.Raw.Removed,
						Raw:                  i.Raw,
					}
				}
			}
		}()
	case SubscribeDOSProxyLogValidationResult:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogValidationResult)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.DOSProxyFilterer.WatchLogValidationResult(opt, transitChan)
			if err != nil {
				fmt.Println("SubscribeDOSProxyLogValidationResult err", err)
				return
			}
			for {
				select {
				case <-ctx.Done():
					fmt.Println("SubscribeDOSProxyLogValidationResult Done")

					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					fmt.Println("SubscribeDOSProxyLogValidationResult err", err)

					errc <- err
					return
				case i := <-transitChan:
					out <- &DOSProxyLogValidationResult{
						TrafficType: i.TrafficType,
						TrafficId:   i.TrafficId,
						Message:     i.Message,
						Signature:   i.Signature,
						PubKey:      i.PubKey,
						GroupId:     i.GroupId,
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
	case SubscribeDOSProxyLogInsufficientPendingNode:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogInsufficientPendingNode)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DOSProxyFilterer.WatchLogInsufficientPendingNode(opt, transitChan)
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
					out <- &DOSProxyLogInsufficientPendingNode{
						NumPendingNodes: i.NumPendingNodes,
						Tx:              i.Raw.TxHash.Hex(),
						BlockN:          i.Raw.BlockNumber,
						Removed:         i.Raw.Removed,
					}
				}
			}
		}()
		/*
			case SubscribeDOSProxyLogNoWorkingGroup:
				go func() {
					transitChan := make(chan *dosproxy.DOSProxyLogNoWorkingGroup)
					defer close(transitChan)
					defer close(errc)
					defer close(out)
					sub, err := proxy.DOSProxyFilterer.WatchLogNoWorkingGroup(opt, transitChan)
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
							out <- &DOSProxyLogNoWorkingGroup{
								Raw:     i.Raw,
								Tx:      i.Raw.TxHash.Hex(),
								BlockN:  i.Raw.BlockNumber,
								Removed: i.Raw.Removed,
							}
						}
					}
				}()
		*/
	}
	return
}

func (e *EthAdaptor) PollLogs(subscribeType int, LogBlockDiff, preBlockBuf uint64) (chan interface{}, <-chan error) {
	errc := make(chan error)
	sink := make(chan interface{})
	go func() {
		defer close(errc)
		defer close(sink)
		targetBlockN, err := GetCurrentBlock(e.wClient)
		if err != nil {
			errc <- err
			return
		}
		targetBlockN -= preBlockBuf
		timer := time.NewTimer(LogCheckingInterval * time.Second)
		for {
			select {
			case <-timer.C:
				currentBlockN, err := GetCurrentBlock(e.wClient)
				if err != nil {
					errc <- err
					continue
				}
				for ; currentBlockN-LogBlockDiff >= targetBlockN; targetBlockN++ {
					switch subscribeType {
					case SubscribeDOSProxyLogGrouping:
						logs, err := e.wProxy.Contract.DOSProxyFilterer.FilterLogGrouping(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.wCtx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DOSProxyLogGrouping{
								GroupId: logs.Event.GroupId,
								NodeId:  logs.Event.NodeId,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
								Raw:     logs.Event.Raw,
							}
						}
					case SubscribeDOSProxyLogGroupDismiss:
						logs, err := e.wProxy.Contract.DOSProxyFilterer.FilterLogGroupDismiss(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.wCtx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DOSProxyLogGroupDismiss{
								PubKey:  logs.Event.PubKey,
								GroupId: logs.Event.GroupId,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
								Raw:     logs.Event.Raw,
							}
						}
					case SubscribeDOSProxyLogUpdateRandom:
						logs, err := e.wProxy.Contract.DOSProxyFilterer.FilterLogUpdateRandom(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.wCtx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DOSProxyLogUpdateRandom{
								LastRandomness:    logs.Event.LastRandomness,
								DispatchedGroupId: logs.Event.DispatchedGroupId,
								DispatchedGroup:   logs.Event.DispatchedGroup,
								Tx:                logs.Event.Raw.TxHash.Hex(),
								BlockN:            logs.Event.Raw.BlockNumber,
								Removed:           logs.Event.Raw.Removed,
								Raw:               logs.Event.Raw,
							}
						}
					case SubscribeDOSProxyLogRequestUserRandom:
						logs, err := e.wProxy.Contract.DOSProxyFilterer.FilterLogRequestUserRandom(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.wCtx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DOSProxyLogRequestUserRandom{
								RequestId:            logs.Event.RequestId,
								LastSystemRandomness: logs.Event.LastSystemRandomness,
								UserSeed:             logs.Event.UserSeed,
								DispatchedGroupId:    logs.Event.DispatchedGroupId,
								DispatchedGroup:      logs.Event.DispatchedGroup,
								Tx:                   logs.Event.Raw.TxHash.Hex(),
								BlockN:               logs.Event.Raw.BlockNumber,
								Removed:              logs.Event.Raw.Removed,
								Raw:                  logs.Event.Raw,
							}
						}
					case SubscribeDOSProxyLogUrl:
						logs, err := e.wProxy.Contract.DOSProxyFilterer.FilterLogUrl(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.wCtx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DOSProxyLogUrl{
								QueryId:           logs.Event.QueryId,
								Timeout:           logs.Event.Timeout,
								DataSource:        logs.Event.DataSource,
								Selector:          logs.Event.Selector,
								Randomness:        logs.Event.Randomness,
								DispatchedGroupId: logs.Event.DispatchedGroupId,
								DispatchedGroup:   logs.Event.DispatchedGroup,
								Tx:                logs.Event.Raw.TxHash.Hex(),
								BlockN:            logs.Event.Raw.BlockNumber,
								Removed:           logs.Event.Raw.Removed,
								Raw:               logs.Event.Raw,
							}
						}
					case SubscribeDOSProxyLogValidationResult:
						logs, err := e.wProxy.Contract.DOSProxyFilterer.FilterLogValidationResult(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.wCtx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DOSProxyLogValidationResult{
								TrafficType: logs.Event.TrafficType,
								TrafficId:   logs.Event.TrafficId,
								Message:     logs.Event.Message,
								Signature:   logs.Event.Signature,
								PubKey:      logs.Event.PubKey,
								GroupId:     logs.Event.GroupId,
								Pass:        logs.Event.Pass,
								Version:     logs.Event.Version,
								Tx:          logs.Event.Raw.TxHash.Hex(),
								BlockN:      logs.Event.Raw.BlockNumber,
								Removed:     logs.Event.Raw.Removed,
								Raw:         logs.Event.Raw,
							}
						}
					}
				}
				timer.Reset(LogCheckingInterval * time.Second)
			case <-e.wCtx.Done():
				return
			}
		}
	}()
	return sink, errc
}

func (e *EthAdaptor) LastRandomness() (rand *big.Int, err error) {
	rand, err = e.wProxy.LastRandomness()
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

func (e *EthAdaptor) GetWorkingGroupSize() (size uint64) {
	sizeBig, err := e.wProxy.GetWorkingGroupSize()
	if err != nil {
		logger.Error(err)
		return
	}
	size = sizeBig.Uint64()
	return
}

func (e *EthAdaptor) LastUpdatedBlock() (blknum uint64, err error) {
	lastBlk, err := e.wProxy.LastUpdatedBlock()
	if err != nil {
		logger.Error(err)
		return
	}
	blknum = lastBlk.Uint64()
	return
}

func (e *EthAdaptor) GroupPubKey(idx int) (groupPubKeys [4]*big.Int, err error) {
	return e.wProxy.GetGroupPubKey(big.NewInt(int64(idx)))
}

func (e *EthAdaptor) BootStrap() error {
	result := make(chan Reply)
	request := &Request{e.wCtx, e.wProxy.BootStrap, result}
	errc := e.sendRequest(e.wCtx, request, result)
	return <-errc
}

func (e *EthAdaptor) ResetContract() error {
	result := make(chan Reply)
	request := &Request{e.wCtx, e.wProxy.ResetContract, result}
	errc := e.sendRequest(e.wCtx, request, result)
	return <-errc
}

func (e *EthAdaptor) CurrentBlock() (blknum uint64, err error) {
	var header *types.Header
	header, err = e.wClient.HeaderByNumber(e.wCtx, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	blknum = header.Number.Uint64()
	return
}

func (e *EthAdaptor) GetBalance() (balance *big.Float) {
	return GetBalance(e.wClient, e.key)
}

func (e *EthAdaptor) Address() (addr []byte) {
	return e.key.Address.Bytes()
}

func DecodeSig(sig []byte) (x, y *big.Int) {
	x = new(big.Int)
	y = new(big.Int)
	x.SetBytes(sig[0:32])
	y.SetBytes(sig[32:])
	return
}
