package eth

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/testing/dosUser/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: Instead of hardcode, read from DOSAddressBridge.go
const (
	SubscribeAskMeAnythingSetTimeout = iota
	SubscribeAskMeAnythingQueryResponseReady
	SubscribeAskMeAnythingRequestSent
	SubscribeAskMeAnythingRandomReady
)

type AMAConfig struct {
	AskMeAnythingAddressPool []string
}

var logger log.Logger

type EthUserAdaptor struct {
	s        *dosUser.AskMeAnythingSession
	c        *ethclient.Client
	key      *keystore.Key
	ctx      context.Context
	cancel   context.CancelFunc
	reqQueue chan interface{}

	//read only DOSProxy
	rProxies     []*dosUser.AskMeAnything
	rClients     []*ethclient.Client
	rCtxs        []context.Context
	rCancelFuncs []context.CancelFunc
}

func NewAMAUserSession(credentialPath, passphrase, addr string, gethUrls []string) (adaptor *EthUserAdaptor, err error) {
	key, err := onchain.ReadEthKey(credentialPath, passphrase)
	if err != nil {
		fmt.Println("NewETHProxySession ", err)
		return
	}
	log.Init(key.Address.Bytes()[:])

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	clients := onchain.DialToEth(ctx, gethUrls, key)
	//Use first client
	c, ok := <-clients
	if !ok {
		err = errors.New("No any working eth client")
		return
	}

	a, err := dosUser.NewAskMeAnything(common.HexToAddress(addr), c)
	if err != nil {
		fmt.Println("NewAskMeAnything ", err)
		return
	}
	adaptor = &EthUserAdaptor{}
	ctxD, cancelSession := context.WithCancel(context.Background())
	auth := bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(3000000)
	auth.Context = ctxD
	//nonce, _ := c.PendingNonceAt(ctxD, key.Address)
	//auth.Nonce = big.NewInt(0)
	//auth.Nonce = auth.Nonce.SetUint64(nonce)
	adaptor.s = &dosUser.AskMeAnythingSession{Contract: a, CallOpts: bind.CallOpts{Context: ctxD}, TransactOpts: *auth}
	adaptor.c = c
	adaptor.key = key
	adaptor.ctx = ctxD
	adaptor.cancel = cancelSession
	adaptor.reqQueue = make(chan interface{})

	for client := range clients {
		p, err := dosUser.NewAskMeAnything(common.HexToAddress(addr), client)
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

	return
}

func (e *EthUserAdaptor) Address() (addr common.Address) {
	return e.key.Address
}

func (e *EthUserAdaptor) PollLogs(subscribeType int, logBlockDiff, preBlockBuf uint64) (<-chan interface{}, <-chan error) {
	var errcs []<-chan error
	var sinks []<-chan interface{}
	var wg sync.WaitGroup

	multiplex := func(client *ethclient.Client, askMeAnythingFilterer *dosUser.AskMeAnythingFilterer, ctx context.Context) {
		errc := make(chan error)
		errcs = append(errcs, errc)
		sink := make(chan interface{})
		sinks = append(sinks, sink)
		wg.Done()
		defer close(errc)
		defer close(sink)
		targetBlockN, err := onchain.GetCurrentBlock(client)
		if err != nil {
			errc <- err
			return
		}
		targetBlockN -= preBlockBuf
		timer := time.NewTimer(onchain.LogCheckingInterval * time.Second)
		for {
			select {
			case <-timer.C:
				currentBlockN, err := onchain.GetCurrentBlock(client)
				if err != nil {
					errc <- err
					continue
				}
				for ; currentBlockN-logBlockDiff >= targetBlockN; targetBlockN++ {
					switch subscribeType {
					case SubscribeAskMeAnythingQueryResponseReady:
						logs, err := askMeAnythingFilterer.FilterQueryResponseReady(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &AskMeAnythingQueryResponseReady{
								QueryId: logs.Event.QueryId,
								Result:  logs.Event.Result,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
								Raw:     logs.Event.Raw,
							}
						}
					case SubscribeAskMeAnythingRequestSent:
						logs, err := askMeAnythingFilterer.FilterRequestSent(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &AskMeAnythingRequestSent{
								InternalSerial: logs.Event.InternalSerial,
								Succ:           logs.Event.Succ,
								RequestId:      logs.Event.RequestId,
								Tx:             logs.Event.Raw.TxHash.Hex(),
								BlockN:         logs.Event.Raw.BlockNumber,
								Removed:        logs.Event.Raw.Removed,
								Raw:            logs.Event.Raw,
							}
						}
					case SubscribeAskMeAnythingRandomReady:
						logs, err := askMeAnythingFilterer.FilterRandomReady(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &AskMeAnythingRandomReady{
								GeneratedRandom: logs.Event.GeneratedRandom,
								RequestId:       logs.Event.RequestId,
								Tx:              logs.Event.Raw.TxHash.Hex(),
								BlockN:          logs.Event.Raw.BlockNumber,
								Removed:         logs.Event.Raw.Removed,
								Raw:             logs.Event.Raw,
							}
						}
					}
				}
				timer.Reset(onchain.LogCheckingInterval * time.Second)
			case <-ctx.Done():
				return
			}
		}

	}

	wg.Add(len(e.rClients) + 1)
	go multiplex(e.c, &e.s.Contract.AskMeAnythingFilterer, e.ctx)
	for i := 0; i < len(e.rClients); i++ {
		go multiplex(e.rClients[i], &e.rProxies[i].AskMeAnythingFilterer, e.rCtxs[i])
	}

	wg.Wait()

	return e.first(e.ctx, MergeEvents(e.ctx, sinks...)), MergeErrors(e.ctx, errcs...)
}
func (e *EthUserAdaptor) SubscribeEvent(subscribeType int) (<-chan interface{}, <-chan error) {
	var eventList []<-chan interface{}
	var errcs []<-chan error
	for i := 0; i < len(e.rProxies); i++ {
		proxy := e.rProxies[i]
		if proxy == nil {
			continue
		}
		ctx := e.rCtxs[i]
		if ctx == nil {
			continue
		}
		out, errc := subscribeEvent(ctx, proxy, subscribeType)
		eventList = append(eventList, out)
		errcs = append(errcs, errc)
	}

	return e.first(e.ctx, MergeEvents(e.ctx, eventList...)), MergeErrors(e.ctx, errcs...)
}
func subscribeEvent(ctx context.Context, proxy *dosUser.AskMeAnything, subscribeType int) (<-chan interface{}, <-chan error) {
	out := make(chan interface{})
	errc := make(chan error)
	opt := &bind.WatchOpts{}

	switch subscribeType {
	case SubscribeAskMeAnythingQueryResponseReady:
		go func() {
			transitChan := make(chan *dosUser.AskMeAnythingQueryResponseReady)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.AskMeAnythingFilterer.WatchQueryResponseReady(opt, transitChan)
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
					out <- &AskMeAnythingQueryResponseReady{
						QueryId: i.QueryId,
						Result:  i.Result,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeAskMeAnythingRequestSent:
		go func() {
			transitChan := make(chan *dosUser.AskMeAnythingRequestSent)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.AskMeAnythingFilterer.WatchRequestSent(opt, transitChan)
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
					out <- &AskMeAnythingRequestSent{
						InternalSerial: i.InternalSerial,
						Succ:           i.Succ,
						RequestId:      i.RequestId,
						Tx:             i.Raw.TxHash.Hex(),
						BlockN:         i.Raw.BlockNumber,
						Removed:        i.Raw.Removed,
						Raw:            i.Raw,
					}
				}
			}
		}()
	}
	return out, errc
}

func (e *EthUserAdaptor) first(ctx context.Context, source chan interface{}) <-chan interface{} {
	sink := make(chan interface{})

	go func() {
		defer close(sink)
		visited := make(map[string]uint64)
		for {
			var bytes []byte
			var blkNum uint64
			var event interface{}
			var ok bool
			select {
			case event, ok = <-source:
				if ok {
					switch content := event.(type) {
					case *AskMeAnythingQueryResponseReady:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *AskMeAnythingRequestSent:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *AskMeAnythingRandomReady:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					}
				} else {
					return
				}
			}
			nHash := sha256.Sum256(bytes)
			identity := string(nHash[:])
			curBlk, err := onchain.GetCurrentBlock(e.c)
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

	return sink
}

func MergeEvents(ctx context.Context, cs ...<-chan interface{}) chan interface{} {
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

func MergeErrors(ctx context.Context, cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	// We must ensure that the output channel has the capacity to
	// hold as many errors
	// as there are error channels.
	// This will ensure that it never blocks, even
	// if WaitForPipeline returns early.
	out := make(chan error, len(cs))
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls
	// wg.Done.
	output := func(c <-chan error) {
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

func (e *EthUserAdaptor) Query(internalSerial uint8, url, selector string) (err error) {
	tx, err := e.s.AMA(internalSerial, url, selector)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error() || err.Error() == core.ErrInsufficientFunds.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = e.s.AMA(internalSerial, url, selector)
	}
	if err != nil {
		fmt.Println(" Query AMAerr ", err)
		return
	}
	fmt.Println("tx sent: ", tx.Hash().Hex())

	return
}

func (e *EthUserAdaptor) GetSafeRandom(internalSerial uint8) (err error) {
	tx, err := e.s.RequestSafeRandom(internalSerial)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = e.s.RequestSafeRandom(internalSerial)
	}
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("RequestSafeRandom ", " waiting for confirmation...")
	return
}

func (e *EthUserAdaptor) GetFastRandom() (err error) {
	tx, err := e.s.RequestFastRandom()
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = e.s.RequestFastRandom()
	}
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("RequestSafeRandom ", " waiting for confirmation...")

	//err = e.CheckTransaction(tx)
	return
}
