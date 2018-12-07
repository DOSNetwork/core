package eth

import (
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/testing/dosUser/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// TODO: Instead of hardcode, read from DOSAddressBridge.go
const (
	SubscribeAskMeAnythingSetTimeout = iota
	SubscribeAskMeAnythingQueryResponseReady
	SubscribeAskMeAnythingRequestSent
	SubscribeAskMeAnythingRandomReady
)

type AMAConfig struct {
	AskMeAnythingAddress string
}

type EthUserAdaptor struct {
	onchain.EthCommon
	proxy     *dosUser.AskMeAnything
	lock      *sync.Mutex
	logFilter *sync.Map
	address   string
}

func (e *EthUserAdaptor) Init(address string, config *configuration.ChainConfig) (err error) {
	e.EthCommon.Init("./credential", config)
	e.logFilter = new(sync.Map)
	go e.logMapTimeout()

	e.address = address
	fmt.Println("onChainConn initialization finished.")
	e.lock = new(sync.Mutex)
	e.dialToEth()
	return
}

func (e *EthUserAdaptor) dialToEth() (err error) {
	e.lock.Lock()
	fmt.Println("dialing...")
	e.EthCommon.DialToEth()
	addr := common.HexToAddress(e.address)
	e.proxy, err = dosUser.NewAskMeAnything(addr, e.Client)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Connot Create new proxy, retrying...")
		e.proxy, err = dosUser.NewAskMeAnything(addr, e.Client)
	}
	e.lock.Unlock()
	return
}

func (e *EthUserAdaptor) SubscribeEvent(ch chan interface{}, subscribeType int) (err error) {
	opt := &bind.WatchOpts{}
	done := make(chan bool)

	go e.subscribeEventAttempt(ch, opt, subscribeType, done)

	for {
		select {
		case <-done:
			fmt.Println("subscribing done.")
			return
			//Todo:This will cause multiple event from eth
		case <-time.After(60 * time.Second):
			fmt.Println("retry...")
			e.dialToEth()
			go e.subscribeEventAttempt(ch, opt, subscribeType, done)

		}
	}
}

func (e *EthUserAdaptor) subscribeEventAttempt(ch chan interface{}, opt *bind.WatchOpts, subscribeType int, done chan bool) {
	fmt.Println("attempt to subscribe event...")
	switch subscribeType {
	case SubscribeAskMeAnythingSetTimeout:
		fmt.Println("subscribing AskMeAnythingSetTimeout event...")
		transitChan := make(chan *dosUser.AskMeAnythingSetTimeout)
		sub, err := e.proxy.AskMeAnythingFilterer.WatchSetTimeout(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("AskMeAnythingSetTimeout event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &AskMeAnythingSetTimeout{
						PreviousTimeout: i.PreviousTimeout,
						NewTimeout:      i.NewTimeout,
					}
				}
			}
		}
	case SubscribeAskMeAnythingQueryResponseReady:
		fmt.Println("subscribing AskMeAnythingQueryResponseReady event...")
		transitChan := make(chan *dosUser.AskMeAnythingQueryResponseReady)
		sub, err := e.proxy.AskMeAnythingFilterer.WatchQueryResponseReady(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("AskMeAnythingQueryResponseReady event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !i.Raw.Removed {
					ch <- &AskMeAnythingQueryResponseReady{
						QueryId: i.QueryId,
						Result:  i.Result,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
					}
				}
			}
		}
	case SubscribeAskMeAnythingRequestSent:
		fmt.Println("subscribing AskMeAnythingRequestSent event...")
		transitChan := make(chan *dosUser.AskMeAnythingRequestSent)
		sub, err := e.proxy.AskMeAnythingFilterer.WatchRequestSent(opt, transitChan, []common.Address{e.GetAddress()})
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("AskMeAnythingRequestSent event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !i.Raw.Removed {
					ch <- &AskMeAnythingRequestSent{
						InternalSerial: i.InternalSerial,
						Succ:           i.Succ,
						RequestId:      i.RequestId,
						Tx:             i.Raw.TxHash.Hex(),
						BlockN:         i.Raw.BlockNumber,
					}
				}
			}
		}
	case SubscribeAskMeAnythingRandomReady:
		fmt.Println("subscribing AskMeAnythingRandomReady event...")
		transitChan := make(chan *dosUser.AskMeAnythingRandomReady)
		sub, err := e.proxy.AskMeAnythingFilterer.WatchRandomReady(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("AskMeAnythingRandomReady event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !i.Raw.Removed {
					ch <- &AskMeAnythingRandomReady{
						GeneratedRandom: i.GeneratedRandom,
						RequestId:       i.RequestId,
						Tx:              i.Raw.TxHash.Hex(),
						BlockN:          i.Raw.BlockNumber,
					}
				}
			}
		}
	}
}

func (e *EthUserAdaptor) Query(internalSerial uint8, url, selector string) (err error) {
	auth, err := e.GetAuth()
	if err != nil {
		fmt.Println(" Query GetAuth err ", err)
		return
	}

	tx, err := e.proxy.AMA(auth, internalSerial, url, selector)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = e.proxy.AMA(auth, internalSerial, url, selector)
	}
	if err != nil {
		fmt.Println(" Query AMAerr ", err)
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("Querying ", url, "selector", selector, "waiting for confirmation...")

	//err = e.CheckTransaction(tx)

	return
}

func (e *EthUserAdaptor) GetSafeRandom(internalSerial uint8) (err error) {
	auth, err := e.GetAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.RequestSafeRandom(auth, internalSerial)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = e.proxy.RequestSafeRandom(auth, internalSerial)
	}
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("RequestSafeRandom ", " waiting for confirmation...")

	//err = e.CheckTransaction(tx)

	return
}

func (e *EthUserAdaptor) GetFastRandom() (err error) {
	auth, err := e.GetAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.RequestFastRandom(auth)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = e.proxy.RequestFastRandom(auth)
	}
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("RequestSafeRandom ", " waiting for confirmation...")

	//err = e.CheckTransaction(tx)

	return
}

func (e *EthUserAdaptor) SubscribeToAll(msgChan chan interface{}) (err error) {
	for i := SubscribeAskMeAnythingSetTimeout; i <= SubscribeAskMeAnythingRandomReady; i++ {
		err = e.SubscribeEvent(msgChan, i)
	}
	return
}

type logRecord struct {
	content       types.Log
	currTimeStamp time.Time
}

func (e *EthUserAdaptor) filterLog(raw types.Log) (duplicates bool) {
	fmt.Println("check duplicates")
	identityBytes := append(raw.Address.Bytes(), raw.Topics[0].Bytes()...)
	identityBytes = append(identityBytes, raw.Data...)
	identity := new(big.Int).SetBytes(identityBytes).String()

	var record interface{}
	if record, duplicates = e.logFilter.Load(identity); duplicates {
		fmt.Println("got duplicate event", record, "\n", raw)
	}
	e.logFilter.Store(identity, logRecord{raw, time.Now()})

	return
}

func (e *EthUserAdaptor) logMapTimeout() {
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		e.logFilter.Range(e.checkTime)
	}

}

func (e *EthUserAdaptor) checkTime(log, deliverTime interface{}) (okToDelete bool) {
	switch t := deliverTime.(type) {
	case logRecord:
		if time.Now().Sub(t.currTimeStamp).Seconds() > 60*10 {
			e.logFilter.Delete(log)
		}
	}
	return true
}
