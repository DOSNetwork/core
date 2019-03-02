package eth

import (
	"context"
	"errors"
	"fmt"
	"math/big"
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
}

func NewAMAUserSession(credentialPath, passphrase, addr string, gethUrls []string) (adaptor *EthUserAdaptor, err error) {
	key, err := onchain.SetEthKey(credentialPath, passphrase)
	if err != nil {
		fmt.Println("NewETHProxySession ", err)
		return
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	clients, errs := onchain.DialToEth(ctx, gethUrls)
	go func() {
		for e := range errs {
			fmt.Println("NewETHProxySession ", e)
			cancelFunc()
		}
	}()

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
	auth.GasLimit = uint64(5000000)
	auth.Context = ctxD
	nonce, _ := c.PendingNonceAt(ctxD, key.Address)
	auth.Nonce = big.NewInt(0)
	auth.Nonce = auth.Nonce.SetUint64(nonce)
	adaptor.s = &dosUser.AskMeAnythingSession{Contract: a, CallOpts: bind.CallOpts{Context: ctxD}, TransactOpts: *auth}
	adaptor.c = c
	adaptor.key = key
	adaptor.ctx = ctxD
	adaptor.cancel = cancelSession
	adaptor.reqQueue = make(chan interface{})
	logger = log.New("module", "EthUser")
	return
}

func (e *EthUserAdaptor) Address() (addr common.Address) {
	return e.key.Address
}

func (e *EthUserAdaptor) PollLogs(subscribeType int, sink chan interface{}) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		targetBlockN, err := onchain.CurrentBlock(e.c)
		if err != nil {
			errc <- err
			return
		}
		timer := time.NewTimer(onchain.LogCheckingInterval * time.Second)

		for {
			select {
			case <-timer.C:
				currentBlockN, err := onchain.CurrentBlock(e.c)
				if err != nil {
					fmt.Println("PollLogs ", err)
					errc <- err
					return
				}
				for ; currentBlockN-onchain.LogBlockDiff >= targetBlockN; targetBlockN++ {
					switch subscribeType {
					case SubscribeAskMeAnythingQueryResponseReady:
						logs, err := e.s.Contract.AskMeAnythingFilterer.FilterQueryResponseReady(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.ctx,
						})
						if err != nil {
							fmt.Println("PollLogs ", err)
							errc <- err
							return
						}
						for logs.Next() {
							sink <- &AskMeAnythingQueryResponseReady{
								QueryId: logs.Event.QueryId,
								Result:  logs.Event.Result,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
							}
							f := map[string]interface{}{
								"RequestId": fmt.Sprintf("%x", logs.Event.QueryId),
								"Message":   logs.Event.Result,
								"Removed":   logs.Event.Raw.Removed,
								"Tx":        logs.Event.Raw.TxHash.Hex(),
								"BlockN":    logs.Event.Raw.BlockNumber,
								"Time":      time.Now()}
							logger.Event("EthUserQueryReady", f)
						}
					case SubscribeAskMeAnythingRequestSent:
						logs, err := e.s.Contract.AskMeAnythingFilterer.FilterRequestSent(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.ctx,
						}, []common.Address{e.Address()})
						if err != nil {
							errc <- err
							return
						}
						for logs.Next() {
							f := map[string]interface{}{
								"Event":     "EthUserRequestSent",
								"RequestId": fmt.Sprintf("%x", logs.Event.RequestId),
								"Succ":      logs.Event.Succ,
								"Removed":   logs.Event.Raw.Removed,
								"Tx":        logs.Event.Raw.TxHash.Hex(),
								"BlockN":    logs.Event.Raw.BlockNumber,
								"Time":      time.Now()}
							logger.Event("EthUserRequestSent", f)
							sink <- &AskMeAnythingRequestSent{
								InternalSerial: logs.Event.InternalSerial,
								Succ:           logs.Event.Succ,
								RequestId:      logs.Event.RequestId,
								Tx:             logs.Event.Raw.TxHash.Hex(),
								BlockN:         logs.Event.Raw.BlockNumber,
								Removed:        logs.Event.Raw.Removed,
							}
						}
					case SubscribeAskMeAnythingRandomReady:
						logs, err := e.s.Contract.AskMeAnythingFilterer.FilterRandomReady(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.ctx,
						})
						if err != nil {
							fmt.Println("PollLogs ", err)
							errc <- err
							return
						}
						for logs.Next() {
							sink <- &AskMeAnythingRandomReady{
								GeneratedRandom: logs.Event.GeneratedRandom,
								RequestId:       logs.Event.RequestId,
								Tx:              logs.Event.Raw.TxHash.Hex(),
								BlockN:          logs.Event.Raw.BlockNumber,
								Removed:         logs.Event.Raw.Removed,
							}
							f := map[string]interface{}{
								"RequestId":       fmt.Sprintf("%x", logs.Event.RequestId),
								"GeneratedRandom": fmt.Sprintf("%x", logs.Event.GeneratedRandom),
								"Removed":         logs.Event.Raw.Removed,
								"Tx":              logs.Event.Raw.TxHash.Hex(),
								"BlockN":          logs.Event.Raw.BlockNumber,
								"Time":            time.Now()}
							logger.Event("EthUserRandomReady", f)
						}
					}
				}
				timer.Reset(onchain.LogCheckingInterval * time.Second)
			case <-e.ctx.Done():
				return
			}
		}
	}()
	return errc
}

func (e *EthUserAdaptor) Query(internalSerial uint8, url, selector string) (err error) {
	tx, err := e.s.AMA(internalSerial, url, selector)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
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
