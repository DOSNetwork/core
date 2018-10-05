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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/DOSNetwork/core/blockchain/eth/contracts"
)

const ethRemoteNode = "wss://rinkeby.infura.io/ws"
const contractAddressHex = "0x6eD7fe957305070f2fB5f5476Eb2dce8f8285cC1"

var contractAddress = common.HexToAddress(contractAddressHex)

type EthAdaptor struct {
	id     *big.Int
	key    *keystore.Key
	client *ethclient.Client
	proxy  *dosproxy.DOSProxy
	lock   *sync.Mutex
}

func (e *EthAdaptor) Init(autoReplenish bool) (err error) {
	fmt.Println("start initial onChainConn...")
	e.lock = new(sync.Mutex)
	e.lock.Lock()
	fmt.Println("dialing...")
	e.client, err = ethclient.Dial(ethRemoteNode)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Cannot connect to the network, retrying...")
		e.client, err = ethclient.Dial(ethRemoteNode)
	}

	e.proxy, err = dosproxy.NewDOSProxy(contractAddress, e.client)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Connot Create new proxy, retrying...")
		e.proxy, err = dosproxy.NewDOSProxy(contractAddress, e.client)
	}
	e.lock.Unlock()

	dir, err := os.Getwd()
	if err != nil {
		return
	}

	if err = e.setAccount(dir, autoReplenish); err != nil {
		return
	}
	fmt.Println("nodeId: ", e.id)

	fmt.Println("onChainConn initialization finished.")
	return
}

func (e *EthAdaptor) SubscribeEvent(ch chan interface{}) (err error) {
	opt := &bind.WatchOpts{}
	identity := <-ch
	done := make(chan bool)

	go e.subscribeEventAttempt(ch, opt, identity, done)

	for {
		select {
		case <-done:
			fmt.Println("subscribing done.")
			return
		case <-time.After(3 * time.Second):
			fmt.Println("retry...")
			e.lock.Lock()
			e.client, err = ethclient.Dial(ethRemoteNode)
			for err != nil {
				fmt.Println(err)
				fmt.Println("Cannot connect to the network, retrying...")
				e.client, err = ethclient.Dial(ethRemoteNode)
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

func (e *EthAdaptor) subscribeEventAttempt(ch chan interface{}, opt *bind.WatchOpts, identity interface{}, done chan bool) {
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
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogGrouping{
					GroupId: i.GroupId,
					NodeId:  i.NodeId,
				}
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
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogUrl{
					QueryId: i.QueryId,
					Url:     i.Url,
					Timeout: i.Timeout,
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

func (e *EthAdaptor) UploadID() (err error) {
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
	fmt.Println("NodeId submitted, waiting for confirmation...")

	err = e.checkTransaction(tx)

	return
}

func (e *EthAdaptor) GetId() (id []byte) {
	return e.id.Bytes()
}

func (e *EthAdaptor) GetBootstrapIp() (ip string, err error) {
	return e.proxy.GetBootstrapIp(&bind.CallOpts{})
}

func (e *EthAdaptor) SetBootstrapIp(ip string) (err error) {
	fmt.Println("Starting submitting bootstrapIp...")
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.SetBootstrapIp(auth, ip)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("bootstrapIp ", ip, " submitted, waiting for confirmation...")

	err = e.checkTransaction(tx)

	return

}

func (e *EthAdaptor) UploadPubKey(groupId, x0, x1, y0, y1 *big.Int) (err error) {
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

	err = e.checkTransaction(tx)

	return
}

func (e *EthAdaptor) DataReturn(queryId *big.Int, data []byte, x, y *big.Int) (err error) {
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.TriggerCallback(auth, queryId, data, x, y)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("Query_ID request fulfilled ", queryId, " waiting for confirmation...")

	err = e.checkTransaction(tx)

	return
}

func (e *EthAdaptor) getAuth() (auth *bind.TransactOpts, err error) {
	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth = bind.NewKeyedTransactor(e.key.PrivateKey)
	if err != nil {
		return
	}

	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice

	return
}

func (e *EthAdaptor) setAccount(path string, autoReplenish bool) (err error) {
	credentialPath := path + "/credential"
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
	e.id = new(big.Int)
	e.id.SetBytes([]byte(e.key.Id.String()))

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
	}

	return
}

func (e *EthAdaptor) balanceMaintain(usrKey, rootKey *keystore.Key) (err error) {
	usrKeyBalance, err := e.getBalance(usrKey)
	if err != nil {
		return
	}

	if usrKeyBalance.Cmp(big.NewFloat(0.2)) == -1 {
		fmt.Println("userKey account replenishing...")
		if err = e.transferEth(rootKey, usrKey); err == nil {
			fmt.Println("userKey account replenished.")
		}
	}

	return
}

func (e *EthAdaptor) getBalance(key *keystore.Key) (balance *big.Float, err error) {
	wei, err := e.client.BalanceAt(context.Background(), key.Address, nil)
	if err != nil {
		return
	}

	balance = new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))
	return
}

func (e *EthAdaptor) transferEth(from, to *keystore.Key) (err error) {
	nonce, err := e.client.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		return
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	value := big.NewInt(1000000000000000000)
	gasLimit := uint64(5000000)
	tx := types.NewTransaction(nonce, to.Address, value, gasLimit, gasPrice, nil)

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

func (e *EthAdaptor) checkTransaction(tx *types.Transaction) (err error) {
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
