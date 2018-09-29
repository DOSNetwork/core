package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/blockchain/eth/contracts"
	"github.com/DOSNetwork/core/blockchain/eventMsg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = ``
const passphrase = ``
const ethReomteNode = "wss://rinkeby.infura.io/ws"
const contractAddressHex = "0xbD5784b224D40213df1F9eeb572961E2a859Cb80"

var contractAddress = common.HexToAddress(contractAddressHex)

type EthConn struct {
	client	*ethclient.Client
	proxy	*dosproxy.DOSProxy
	lock	*sync.Mutex
}

func (e *EthConn) Init() (err error) {
	fmt.Println("start initial onChainConn...")
	e.lock = new(sync.Mutex)
	e.lock.Lock()

	e.client, err = ethclient.Dial(ethReomteNode)
	if err != nil {
		e.lock.Unlock()
		return
	}

	e.proxy, err = dosproxy.NewDOSProxy(contractAddress, e.client)
	if err != nil {
		e.lock.Unlock()
		return
	}

	e.lock.Unlock()
	fmt.Println("onChainConn initialization finished.")
	return
}

func (e *EthConn) SubscribeEvent(ch chan interface{}) (err error){
	opt := &bind.WatchOpts{}
	identity :=<- ch
	done := make(chan bool)

	go e.subscribeEventAttempt(ch, opt, identity, done)

	for {
		select {
		case <- done:
			fmt.Println("subscribing done.")
			return
		case <-time.After(3 * time.Second):
			fmt.Println("retry...")
			e.lock.Lock()
			e.client, err = ethclient.Dial(ethReomteNode)
			if err != nil {
				e.lock.Unlock()
				return
			}

			e.proxy, err = dosproxy.NewDOSProxy(contractAddress, e.client)
			if err != nil {
				e.lock.Unlock()
				return
			}

			e.lock.Unlock()
			go e.subscribeEventAttempt(ch, opt, identity, done)
		}
	}
	return nil
}

func (e *EthConn) subscribeEventAttempt(ch chan interface{}, opt *bind.WatchOpts, identity interface{}, done chan bool) {
	fmt.Println("attempt to subscribe event...")
	switch identity.(type) {
	case *eventMsg.DOSProxyLogCallbackTriggeredFor:
	case *eventMsg.DOSProxyLogInvalidSignature:
	case *eventMsg.DOSProxyLogNonContractCall:
	case *eventMsg.DOSProxyLogNonSupportedType:
	case *eventMsg.DOSProxyLogQueryFromNonExistentUC:
	case *eventMsg.DOSProxyLogSuccPubKeySub:
		fmt.Println("subscribing DOSProxyLogSuccPubKeySub event...")
		transitChan := make(chan *dosproxy.DOSProxyLogSuccPubKeySub)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogSuccPubKeySub(opt, transitChan)
		if err != nil {
			fmt.Println("Network fail, will retry shortly")
		}
		fmt.Println("DOSProxyLogSuccPubKeySub event subscribed")

		done <- true
		for {
			select {
			case err := <- sub.Err():
				log.Fatal(err)
			case _ = <- transitChan:
				ch <- &eventMsg.DOSProxyLogSuccPubKeySub{}
			}
		}
	case *eventMsg.DOSProxyLogUrl:
		fmt.Println("subscribing DOSProxyLogUrl event...")
		transitChan := make(chan *dosproxy.DOSProxyLogUrl)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogUrl(opt, transitChan)
		if err != nil {
			fmt.Println("Network fail, will retry shortly")
		}
		fmt.Println("DOSProxyLogUrl event subscribed")

		done <- true
		for {
			select {
			case err := <- sub.Err():
				log.Fatal(err)
			case i := <- transitChan:
				ch <- &eventMsg.DOSProxyLogUrl{
					QueryId:	i.QueryId,
					Url:		i.Url,
					Timeout:	i.Timeout,
				}
			}
		}
	}
}

func (e *EthConn) UploadPubKey(groupId, x0, x1, y0, y1 *big.Int) (err error) {
	fmt.Println("Starting submitting group public key...")
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		return
	}

	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice

	tx, err := e.proxy.SetPublicKey(auth, groupId, x0, x1, y0, y1)

	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("groupId: ", groupId)
	fmt.Println("x0: ", x0)
	fmt.Println("x1: ", x1)
	fmt.Println("y0: ", y0)
	fmt.Println("y1: ", y1)
	fmt.Println("Group public key submitted, waiting for confirmation...")

	return
}

func (e *EthConn) DataReturn(queryId *big.Int, data []byte, x, y *big.Int) (err error) {
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		return
	}

	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice

	tx, err := e.proxy.TriggerCallback(auth, queryId, data, x, y)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Printf("Query_ID %v request fulfilled \n", queryId)

	return
}