package eth

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"os"
	"sync"
	"time"

	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/onchain/eth/contracts/userContract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: Instead of hardcode, read from DOSAddressBridge.go
const (
	SubscribeAskMeAnythingSetTimeout = iota
	SubscribeAskMeAnythingQueryResponseReady
	SubscribeAskMeAnythingRequestSent
	SubscribeAskMeAnythingRandomReady
)

const balanceCheckInterval = 3

var workingDir string

type EthUserAdaptor struct {
	key       *keystore.Key
	client    *ethclient.Client
	proxy     *userContract.AskMeAnything
	lock      *sync.Mutex
	logFilter *sync.Map
	ethNonce  uint64
	config    *onchain.ChainConfig
}

func (e *EthUserAdaptor) Init(autoReplenish bool, config *onchain.ChainConfig) (err error) {

	path := os.Getenv("CONFIGPATH")
	if path != "" {
		workingDir = path
	} else {
		workingDir, err = os.Getwd()
		if err != nil {
			return
		}
	}

	if err != nil {
		return
	}

	e.config = config

	fmt.Println("start initial onChainConn...")

	e.logFilter = new(sync.Map)
	go e.logMapTimeout()

	e.lock = new(sync.Mutex)

	e.dialToEth()

	if err = e.setAccount(autoReplenish); err != nil {
		return
	}

	fmt.Println("onChainConn initialization finished.")
	return
}

func (e *EthUserAdaptor) dialToEth() (err error) {
	e.lock.Lock()
	fmt.Println("dialing...")
	e.client, err = ethclient.Dial(e.config.RemoteNodeAddress)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Cannot connect to the network, retrying...")
		e.client, err = ethclient.Dial(e.config.RemoteNodeAddress)
	}
	addr := common.HexToAddress(e.config.AskMeAnythingAddress)
	e.proxy, err = userContract.NewAskMeAnything(addr, e.client)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Connot Create new proxy, retrying...")
		e.proxy, err = userContract.NewAskMeAnything(addr, e.client)
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
		transitChan := make(chan *userContract.AskMeAnythingSetTimeout)
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
		transitChan := make(chan *userContract.AskMeAnythingQueryResponseReady)
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
				if !e.filterLog(i.Raw) {
					ch <- &AskMeAnythingQueryResponseReady{
						QueryId: i.QueryId,
						Result:  i.Result,
					}
				}
			}
		}
	case SubscribeAskMeAnythingRequestSent:
		fmt.Println("subscribing AskMeAnythingRequestSent event...")
		transitChan := make(chan *userContract.AskMeAnythingRequestSent)
		sub, err := e.proxy.AskMeAnythingFilterer.WatchRequestSent(opt, transitChan)
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
				if !e.filterLog(i.Raw) {
					ch <- &AskMeAnythingRequestSent{
						Succ:      i.Succ,
						RequestId: i.RequestId,
					}
				}
			}
		}
	case SubscribeAskMeAnythingRandomReady:
		fmt.Println("subscribing AskMeAnythingRandomReady event...")
		transitChan := make(chan *userContract.AskMeAnythingRandomReady)
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
				if !e.filterLog(i.Raw) {
					ch <- &AskMeAnythingRandomReady{
						GeneratedRandom: i.GeneratedRandom,
						RequestId:       i.RequestId,
					}
				}
			}
		}
	}
}

func (e *EthUserAdaptor) Query(url, selector string) (err error) {
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.AMA(auth, url, selector)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("Querying ", url, "selector", selector, "waiting for confirmation...")

	//err = e.checkTransaction(tx)

	return
}

func (e *EthUserAdaptor) GetSafeRandom() (err error) {
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.RequestSafeRandom(auth)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("RequestSafeRandom ", " waiting for confirmation...")

	//err = e.checkTransaction(tx)

	return
}

func (e *EthUserAdaptor) GetFastRandom() (err error) {
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.RequestFastRandom(auth)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("RequestSafeRandom ", " waiting for confirmation...")

	err = e.checkTransaction(tx)

	return
}

func (e *EthUserAdaptor) getAuth() (auth *bind.TransactOpts, err error) {
	auth = bind.NewKeyedTransactor(e.key.PrivateKey)
	if err != nil {
		return
	}

	e.lock.Lock()
	e.ethNonce++
	automatedNonce, err := e.client.PendingNonceAt(context.Background(), e.key.Address)
	if err != nil {
		return
	}
	if automatedNonce > e.ethNonce {
		e.ethNonce = automatedNonce
	}
	auth.Nonce = big.NewInt(int64(e.ethNonce))
	e.lock.Unlock()

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth.GasLimit = uint64(6000000) // in units
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(3))

	return
}

func (e *EthUserAdaptor) setAccount(autoReplenish bool) (err error) {
	credentialPath := workingDir + "/credential"
	fmt.Println("credentialPath: ", credentialPath)

	passPhraseBytes, err := ioutil.ReadFile(credentialPath + "/boot/passPhrase")
	if err != nil {
		return
	}

	passPhrase := string(passPhraseBytes)

	newKeyStore := keystore.NewKeyStore(credentialPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if len(newKeyStore.Accounts()) < 1 {
		_, err = newKeyStore.NewAccount(passPhrase)
		if err != nil {
			return
		}
	}

	usrKeyPath := newKeyStore.Accounts()[0].URL.Path
	usrKeyJson, err := ioutil.ReadFile(usrKeyPath)
	if err != nil {
		return
	}

	usrKey, err := keystore.DecryptKey(usrKeyJson, passPhrase)
	if err != nil {
		return
	}

	e.key = usrKey
	e.ethNonce, err = e.client.PendingNonceAt(context.Background(), e.key.Address)
	if err != nil {
		return
	}
	//for correctness of the first call of getAuth, because getAuth always ++,
	e.ethNonce--

	if autoReplenish {
		var rootKeyPath string
		var rootKeyJson []byte
		var rootKey *keystore.Key

		rootKeyPath = credentialPath + "/boot/rootKey"
		rootKeyJson, err = ioutil.ReadFile(rootKeyPath)
		if err != nil {
			return
		}

		rootKey, err = keystore.DecryptKey(rootKeyJson, passPhrase)
		if err != nil {
			return
		}

		err = e.balanceMaintain(usrKey, rootKey)

		go func() {
			ticker := time.NewTicker(balanceCheckInterval * time.Hour)
			for range ticker.C {
				err = e.balanceMaintain(usrKey, rootKey)
				if err != nil {
					fmt.Print("Fail to replenish.")
				}
			}
		}()
	}

	return
}

func (e *EthUserAdaptor) balanceMaintain(usrKey, rootKey *keystore.Key) (err error) {
	usrKeyBalance, err := e.getBalance(usrKey)
	if err != nil {
		return
	}

	if usrKeyBalance.Cmp(big.NewFloat(0.7)) == -1 {
		fmt.Println("userKey account replenishing...")
		if err = e.transferEth(rootKey, usrKey); err == nil {
			fmt.Println("userKey account replenished.")
		}
	}

	return
}

func (e *EthUserAdaptor) getBalance(key *keystore.Key) (balance *big.Float, err error) {
	wei, err := e.client.BalanceAt(context.Background(), key.Address, nil)
	if err != nil {
		return
	}

	balance = new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))
	return
}

func (e *EthUserAdaptor) transferEth(from, to *keystore.Key) (err error) {
	nonce, err := e.client.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		return
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	value := big.NewInt(800000000000000000) //0.8 Eth
	gasLimit := uint64(6000000)

	tx := types.NewTransaction(nonce, to.Address, value, gasLimit, gasPrice.Mul(gasPrice, big.NewInt(3)), nil)

	chainId, err := e.client.NetworkID(context.Background())
	if err != nil {
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), from.PrivateKey)
	if err != nil {
		return
	}

	err = e.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}
	fmt.Println("replenishing tx sent: ", signedTx.Hash().Hex(), ", waiting for confirmation...")

	err = e.checkTransaction(signedTx)

	return
}

func (e *EthUserAdaptor) checkTransaction(tx *types.Transaction) (err error) {
	receipt, err := e.client.TransactionReceipt(context.Background(), tx.Hash())
	for err == ethereum.NotFound {
		time.Sleep(1 * time.Second)
		receipt, err = e.client.TransactionReceipt(context.Background(), tx.Hash())
	}
	if err != nil {
		return
	}

	if receipt.Status == 1 {
		fmt.Println("transaction confirmed.")
	} else {
		err = errors.New("transaction failed")
	}

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
