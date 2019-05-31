package eth

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"strings"
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

var logger log.Logger

//AMAConfig is a struct that saves AMA contract address
type AMAConfig struct {
	AskMeAnythingAddressPool []string
}

//EthUserAdaptor is an adaptor to the proxy contract on the ethereum
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

//NewAMAUserSession creates a dosproxy adaptor struct
func NewAMAUserSession(credentialPath, passphrase, addr string, gethUrls []string) (adaptor *EthUserAdaptor, err error) {
	var httpUrls []string
	var wsUrls []string
	for _, url := range gethUrls {
		if strings.Contains(url, "http") {
			httpUrls = append(httpUrls, url)
		} else if strings.Contains(url, "ws") {
			wsUrls = append(wsUrls, url)
		}
	}
	fmt.Println("gethUrls ", httpUrls)
	fmt.Println("eventUrls ", wsUrls)
	key, err := onchain.ReadEthKey(credentialPath, passphrase)
	if err != nil {
		fmt.Println("NewETHProxySession ", err)
		return
	}
	log.Init(key.Address.Bytes()[:])

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	clients := onchain.DialToEth(ctx, httpUrls, key)
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
	auth.GasPrice = big.NewInt(2000000000)

	auth.GasLimit = uint64(5000000)
	auth.Context = ctxD
	//nonce, _ := c.PendingNonceAt(ctxD, key.Address)
	//auth.Nonce = big.NewInt(0)
	//auth.Nonce = auth.Nonce.SetUint64(nonce)
	adaptor.s = &dosUser.AskMeAnythingSession{Contract: a, CallOpts: bind.CallOpts{Context: ctxD}, TransactOpts: *auth}
	fmt.Println("adaptor.s ", adaptor.s)

	adaptor.c = c
	adaptor.key = key
	adaptor.ctx = ctxD
	adaptor.cancel = cancelSession
	adaptor.reqQueue = make(chan interface{})

	clients = onchain.DialToEth(ctx, wsUrls, key)
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

// Address return a key address
func (e *EthUserAdaptor) Address() (addr common.Address) {
	return e.key.Address
}

//SubscribeEvent is a log subscription operation binding the contract event
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
	case SubscribeAskMeAnythingRandomReady:
		go func() {
			fmt.Println("AskMeAnythingRandomReady")

			transitChan := make(chan *dosUser.AskMeAnythingRandomReady)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.AskMeAnythingFilterer.WatchRandomReady(opt, transitChan)
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
					out <- &AskMeAnythingRandomReady{
						GeneratedRandom: i.GeneratedRandom,
						RequestId:       i.RequestId,
						Tx:              i.Raw.TxHash.Hex(),
						BlockN:          i.Raw.BlockNumber,
						Removed:         i.Raw.Removed,
						Raw:             i.Raw,
					}
				}
			}
		}()
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
				fmt.Println("SubscribeAskMeAnythingRequestSent err ", err)
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					fmt.Println("SubscribeAskMeAnythingRequestSent err ", err)

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

//MergeEvents is a fan in pattern to merge all events from all channel
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

//MergeEvents is a fan in pattern to merge all errors from all channel
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

//Query requests for external data
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

//GetSafeRandom requests for a random number that is generated for this requesting
func (e *EthUserAdaptor) GetSafeRandom(internalSerial uint8) (err error) {
	if e.s == nil {
		fmt.Println("e.s ==nil")
	}
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

//GetFastRandom requests for a random number that is updated periodically on dos proxy
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
