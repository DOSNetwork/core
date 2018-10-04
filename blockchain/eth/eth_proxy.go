package eth

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/blockchain/eth/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const ethReomteNode = "wss://rinkeby.infura.io/ws"
const contractAddressHex = "0x064515c9b2CaAb4c2e1E176106F8F2a4B7cFC297"
const passphrase = "123"

var contractAddress = common.HexToAddress(contractAddressHex)

type EthConn struct {
	id			*big.Int
	fundedKey 	string
	client		*ethclient.Client
	proxy		*dosproxy.DOSProxy
	lock		*sync.Mutex
}

func (e *EthConn) Init() (err error) {
	dir, err := os.Getwd()
	if err != nil {
		return
	}

	idString, err := getId(dir)
	if err != nil {
		return
	}

	e.id = big.NewInt(0)
	e.id.SetBytes([]byte(idString))

	fmt.Println("nodeId: ", e.id)

	e.fundedKey, err = getFundedKey(dir)
	if err != nil {
		return
	}

	fmt.Println("start initial onChainConn...")
	e.lock = new(sync.Mutex)
	e.lock.Lock()

	e.client, err = ethclient.Dial(ethReomteNode)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Cannot connect to the network, retrying...")
		e.client, err = ethclient.Dial(ethReomteNode)
	}

	e.proxy, err = dosproxy.NewDOSProxy(contractAddress, e.client)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Connot Create new proxy, retrying...")
		e.proxy, err = dosproxy.NewDOSProxy(contractAddress, e.client)
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
			for err != nil {
				fmt.Println(err)
				fmt.Println("Cannot connect to the network, retrying...")
				e.client, err = ethclient.Dial(ethReomteNode)
			}

			e.proxy, err = dosproxy.NewDOSProxy(contractAddress, e.client)
			for err != nil {
				fmt.Println(err)
				fmt.Println("Connot Create new proxy, retrying...")
				e.proxy, err = dosproxy.NewDOSProxy(contractAddress, e.client)
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
	case *DOSProxyLogGrouping:
		fmt.Println("subscribing DOSProxyLogGrouping event...")
		transitChan := make(chan *dosproxy.DOSProxyLogGrouping)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogGrouping(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogGrouping event subscribed")

		done <- true
		for {
			select {
			case err := <- sub.Err():
				log.Fatal(err)
			case i := <- transitChan:
				ch <- &DOSProxyLogGrouping{
					GroupId:	i.GroupId,
					NodeId:		i.NodeId,
				}
			}
		}
	case *DOSProxyLogBootstrapIp:
		fmt.Println("subscribing DOSProxyLogBootstrapIp event...")
		transitChan := make(chan *dosproxy.DOSProxyLogBootstrapIp)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogBootstrapIp(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogBootstrapIp event subscribed")

		done <- true
		for {
			select {
			case err := <- sub.Err():
				log.Fatal(err)
			case i := <- transitChan:
				ch <- &DOSProxyLogBootstrapIp{
					ip:	i.Ip,
				}
			}
		}
	case *DOSProxyLogSuccPubKeySub:
		fmt.Println("subscribing DOSProxyLogSuccPubKeySub event...")
		transitChan := make(chan *dosproxy.DOSProxyLogSuccPubKeySub)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogSuccPubKeySub(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogSuccPubKeySub event subscribed")

		done <- true
		for {
			select {
			case err := <- sub.Err():
				log.Fatal(err)
			case _ = <- transitChan:
				ch <- &DOSProxyLogSuccPubKeySub{}
			}
		}
	case *DOSProxyLogUrl:
		fmt.Println("subscribing DOSProxyLogUrl event...")
		transitChan := make(chan *dosproxy.DOSProxyLogUrl)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogUrl(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogUrl event subscribed")

		done <- true
		for {
			select {
			case err := <- sub.Err():
				log.Fatal(err)
			case i := <- transitChan:
				ch <- &DOSProxyLogUrl{
					QueryId:	i.QueryId,
					Url:		i.Url,
					Timeout:	i.Timeout,
				}
			}
		}
	case *DOSProxyLogCallbackTriggeredFor:
	case *DOSProxyLogInvalidSignature:
	case *DOSProxyLogNonContractCall:
	case *DOSProxyLogNonSupportedType:
	case *DOSProxyLogQueryFromNonExistentUC:
	case *DOSProxyLogInsufficientGroupNumber:
	}
}

func (e *EthConn) UploadID() (bootstrapIp string, err error) {

	ipChan := make(chan interface{})
	go func() {
		ipChan <- &DOSProxyLogBootstrapIp{}
	}()
	e.SubscribeEvent(ipChan)

	fmt.Println("Starting submitting nodeId...")
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.UploadNodeId(auth, e.id)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("NodeId submitted, waiting for bootstrapIp...")

	ipReceived := <- ipChan
	switch ipReceived.(type) {
	case *DOSProxyLogBootstrapIp:
		bootstrapIp = ipReceived.(*DOSProxyLogBootstrapIp).ip
	default:
		return "", errors.New("failed to get bootstrapIp")
	}

	return
}

func (e *EthConn) UploadPubKey(groupId, x0, x1, y0, y1 *big.Int) (err error) {
	fmt.Println("Starting submitting group public key...")
	auth, err := e.getAuth()
	if err != nil {
		return
	}

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
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.TriggerCallback(auth, queryId, data, x, y)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Printf("Query_ID %v request fulfilled \n", queryId)

	return
}

func (e *EthConn) getAuth() (auth *bind.TransactOpts,err error){
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth, err = bind.NewTransactor(strings.NewReader(e.fundedKey), passphrase)
	if err != nil {
		return
	}

	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice

	return
}

func getId(path string) (id string, err error)  {
	credentialPath := path + "/credential"
	fmt.Println("credentialPath: ", credentialPath)

	newKeyStore := keystore.NewKeyStore(credentialPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if len(newKeyStore.Accounts()) < 1 {
		_, err = newKeyStore.NewAccount(passphrase)
		if err != nil {
			return
		}
	}

	keyPath := newKeyStore.Accounts()[0].URL.Path
	keyJson, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return
	}

	key, err := keystore.DecryptKey(keyJson, passphrase)
	if err != nil {
		return
	}

	id = key.Id.String()

	return
}

func (e *EthConn) GetId() (id *big.Int) {
	return e.id
}

func getFundedKey(path string) (fundedKey string, err error)  {
	fundedKeyPath := path + "/credential/funded/fundedKey"
	fundedKeyBytes, err := ioutil.ReadFile(fundedKeyPath)
	if err != nil {
		return
	}

	fundedKey = string(fundedKeyBytes)
	return
}