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
	"strings"
	"sync"
	"time"

	"github.com/dedis/kyber"

	"github.com/DOSNetwork/core/blockchain/eth/contracts"
	"github.com/DOSNetwork/core/blockchain/eth/contracts/userContract"
	"github.com/DOSNetwork/core/group/bn256"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: Instead of hardcode, read from DOSAddressBridge.go

const (
	Rinkeby = iota
	RinkebyPrivate
	Private
)

const (
	AskMeAnyThing = iota
	DOSAddressBridge
	DOSProxy
)

const (
	SubscribeDOSProxyLogUrl = iota
	SubscribeDOSProxyLogNonSupportedType
	SubscribeDOSProxyLogNonContractCall
	SubscribeDOSProxyLogCallbackTriggeredFor
	SubscribeDOSProxyLogQueryFromNonExistentUC
	SubscribeDOSProxyLogUpdateRandom
	SubscribeDOSProxyLogValidationResult
	SubscribeDOSProxyLogInsufficientGroupNumber
	SubscribeDOSProxyLogGrouping
	SubscribeDOSProxyLogPublicKeyAccepted
)

var ethRemoteNode = new(netConfig)
var workingDir string

type EthAdaptor struct {
	id     *big.Int
	key    *keystore.Key
	client *ethclient.Client
	proxy  *dosproxy.DOSProxy
	lock   *sync.Mutex
}

func (e *EthAdaptor) Init(autoReplenish bool, netType int) (err error) {
	workingDir, err = os.Getwd()
	if err != nil {
		return
	}

	fmt.Println("start initial onChainConn...")
	e.lock = new(sync.Mutex)
	e.lock.Lock()

	switch netType {
	case Rinkeby:
		ethRemoteNode = rinkebyNode
	case RinkebyPrivate:
		ethRemoteNode = rinkebyPrivateNode
	case Private:
		ethRemoteNode = privateNode
	}

	fmt.Println("dialing...")
	e.client, err = ethclient.Dial(ethRemoteNode.remoteNodeAddress)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Cannot connect to the network, retrying...")
		e.client, err = ethclient.Dial(ethRemoteNode.remoteNodeAddress)
	}

	e.proxy, err = dosproxy.NewDOSProxy(ethRemoteNode.proxyContractAddress, e.client)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Connot Create new proxy, retrying...")
		e.proxy, err = dosproxy.NewDOSProxy(ethRemoteNode.proxyContractAddress, e.client)
	}
	e.lock.Unlock()

	if err = e.setAccount(autoReplenish); err != nil {
		return
	}
	fmt.Println("nodeId: ", e.id)

	fmt.Println("onChainConn initialization finished.")
	return
}

func (e *EthAdaptor) SubscribeEvent(ch chan interface{}, subscribeType int) (err error) {
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
			e.lock.Lock()
			e.client, err = ethclient.Dial(ethRemoteNode.remoteNodeAddress)
			for err != nil {
				fmt.Println(err)
				fmt.Println("Cannot connect to the network, retrying...")
				e.client, err = ethclient.Dial(ethRemoteNode.remoteNodeAddress)
			}

			e.proxy, err = dosproxy.NewDOSProxy(ethRemoteNode.proxyContractAddress, e.client)
			for err != nil {
				fmt.Println(err)
				fmt.Println("Connot Create new proxy, retrying...")
				e.proxy, err = dosproxy.NewDOSProxy(ethRemoteNode.proxyContractAddress, e.client)
			}

			e.lock.Unlock()
			go e.subscribeEventAttempt(ch, opt, subscribeType, done)

		}
	}
	return nil
}

func (e *EthAdaptor) subscribeEventAttempt(ch chan interface{}, opt *bind.WatchOpts, subscribeType int, done chan bool) {
	fmt.Println("attempt to subscribe event...")
	switch subscribeType {
	case SubscribeDOSProxyLogGrouping:
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
					NodeId: i.NodeId,
				}
			}
		}
	case SubscribeDOSProxyLogUrl:
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
					QueryId:         i.QueryId,
					Url:             i.Url,
					Timeout:         i.Timeout,
					Randomness:      i.Randomness,
					DispatchedGroup: i.DispatchedGroup,
				}
			}
		}
	case SubscribeDOSProxyLogUpdateRandom:
		fmt.Println("subscribing DOSProxyLogUpdateRandom event...")
		transitChan := make(chan *dosproxy.DOSProxyLogUpdateRandom)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogUpdateRandom(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogUpdateRandom event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogUpdateRandom{
					LastRandomness:   i.LastRandomness,
					LastUpdatedBlock: i.LastUpdatedBlock,
					DispatchedGroup:  i.DispatchedGroup,
				}
			}
		}
	case SubscribeDOSProxyLogValidationResult:
		fmt.Println("subscribing SubscribeDOSProxyLogValidationResult event...")
		transitChan := make(chan *dosproxy.DOSProxyLogValidationResult)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogValidationResult(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogValidationResult event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogValidationResult{
					TrafficType: i.TrafficType,
					TrafficId:   i.TrafficId,
					Message:     i.Message,
					Signature:   i.Signature,
					PubKey:      i.PubKey,
					Pass:        i.Pass,
				}
			}
		}
	case SubscribeDOSProxyLogNonSupportedType:
		fmt.Println("subscribing DOSProxyLogNonSupportedType event...")
		transitChan := make(chan *dosproxy.DOSProxyLogNonSupportedType)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogNonSupportedType(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogNonSupportedType event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogNonSupportedType{
					QueryType: i.QueryType,
				}
			}
		}
	case SubscribeDOSProxyLogNonContractCall:
		fmt.Println("subscribing DOSProxyLogNonContractCall event...")
		transitChan := make(chan *dosproxy.DOSProxyLogNonContractCall)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogNonContractCall(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogNonContractCall event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogNonContractCall{
					From: i.From,
				}
			}
		}
	case SubscribeDOSProxyLogCallbackTriggeredFor:
		fmt.Println("subscribing DOSProxyLogCallbackTriggeredFor event...")
		transitChan := make(chan *dosproxy.DOSProxyLogCallbackTriggeredFor)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogCallbackTriggeredFor(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogCallbackTriggeredFor event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogCallbackTriggeredFor{
					CallbackAddr: i.CallbackAddr,
				}
			}
		}
	case SubscribeDOSProxyLogQueryFromNonExistentUC:
		fmt.Println("subscribing DOSProxyLogQueryFromNonExistentUC event...")
		transitChan := make(chan *dosproxy.DOSProxyLogQueryFromNonExistentUC)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogQueryFromNonExistentUC(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogQueryFromNonExistentUC event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case _ = <-transitChan:
				ch <- &DOSProxyLogQueryFromNonExistentUC{}
			}
		}
	case SubscribeDOSProxyLogInsufficientGroupNumber:
		fmt.Println("subscribing DOSProxyLogInsufficientGroupNumber event...")
		transitChan := make(chan *dosproxy.DOSProxyLogInsufficientGroupNumber)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogInsufficientGroupNumber(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogInsufficientGroupNumber event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case _ = <-transitChan:
				ch <- &DOSProxyLogInsufficientGroupNumber{}
			}
		}
	case SubscribeDOSProxyLogPublicKeyAccepted:
		fmt.Println("subscribing DOSProxyLogPublicKeyAccepted event...")
		transitChan := make(chan *dosproxy.DOSProxyLogPublicKeyAccepted)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogPublicKeyAccepted(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogPublicKeyAccepted event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogPublicKeyAccepted{
					X1: i.X1,
					X2: i.X2,
					Y1: i.Y1,
					Y2: i.Y2,
				}
			}
		}
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

func (e *EthAdaptor) GetBlockHashByNumber(blknum *big.Int) (hash common.Hash, err error) {
	block, err := e.client.BlockByNumber(context.Background(), blknum)
	if err != nil {
		return
	}

	hash = block.Hash()
	return
}

func (e *EthAdaptor) SetRandomNum(sig []byte) (err error) {
	fmt.Println("Starting submitting random number...")

	auth, err := e.getAuth()
	if err != nil {
		return
	}

	x, y := decodeSig(sig)

	tx, err := e.proxy.UpdateRandomness(auth, [2]*big.Int{x, y})
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("new random number submitted, waiting for confirmation...")

	err = e.checkTransaction(tx)

	return
}

func (e *EthAdaptor) UploadPubKey(pubKey kyber.Point) (err error) {
	fmt.Println("Starting submitting group public key...")
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	x0, x1, y0, y1, err := decodePubKey(pubKey)
	if err != nil {
		return
	}

	tx, err := e.proxy.SetPublicKey(auth, x0, x1, y0, y1)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("x0: ", x0)
	fmt.Println("x1: ", x1)
	fmt.Println("y0: ", y0)
	fmt.Println("y1: ", y1)
	fmt.Println("Group public key submitted, waiting for confirmation...")

	err = e.checkTransaction(tx)

	return
}

func (e *EthAdaptor) ResetNodeIDs() (err error) {
	fmt.Println("Starting ResetNodeIDs...")
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.ResetContract(auth)
	if err != nil {
		return
	}

	err = e.checkTransaction(tx)
	return
}

func decodePubKey(pubKey kyber.Point) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	pubKeyMar, err := pubKey.MarshalBinary()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	x0 := new(big.Int)
	x1 := new(big.Int)
	y0 := new(big.Int)
	y1 := new(big.Int)
	x0.SetBytes(pubKeyMar[1:33])
	x1.SetBytes(pubKeyMar[33:65])
	y0.SetBytes(pubKeyMar[65:97])
	y1.SetBytes(pubKeyMar[97:])
	return x0, x1, y0, y1, nil
}

func (e *EthAdaptor) DataReturn(queryId *big.Int, data, sig []byte) (err error) {
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	x, y := decodeSig(sig)

	tx, err := e.proxy.TriggerCallback(auth, queryId, data, [2]*big.Int{x, y})
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("Query_ID request fulfilled ", queryId, " waiting for confirmation...")

	err = e.checkTransaction(tx)

	return
}

func decodeSig(sig []byte) (x, y *big.Int) {
	x = new(big.Int)
	y = new(big.Int)
	x.SetBytes(sig[0:32])
	y.SetBytes(sig[32:])

	if x.Cmp(big.NewInt(0)) == 0 && y.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(0)
	}

	y = big.NewInt(0).Sub(bn256.P, big.NewInt(0).Mod(y, bn256.P))

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

	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = gasPrice

	return
}

func (e *EthAdaptor) setAccount(autoReplenish bool) (err error) {
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
	e.id = new(big.Int)
	e.id.SetBytes(e.key.Address.Bytes())

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
	gasLimit := uint64(4000000)
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

func (e *EthAdaptor) DeployContract(contractName int) (address common.Address, err error) {
	var tx *types.Transaction

	auth, err := e.getAuth()
	if err != nil {
		return
	}

	switch contractName {
	case AskMeAnyThing:
		fmt.Println("Starting deploy AskMeAnyThing.sol...")
		address, tx, _, err = userContract.DeployAskMeAnything(auth, e.client)
	case DOSAddressBridge:
		fmt.Println("Starting deploy DOSAddressBridge.sol...")
		address, tx, _, err = dosproxy.DeployDOSAddressBridge(auth, e.client)
	case DOSProxy:
		fmt.Println("Starting deploy DOSProxy.sol...")
		address, tx, _, err = dosproxy.DeployDOSProxy(auth, e.client)
	}
	if err != nil {
		return
	}

	fmt.Println("contract Address: ", address.Hex())
	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("contract deployed, waiting for confirmation...")

	err = e.checkTransaction(tx)

	return
}

func (e *EthAdaptor) DeployAll() (proxyAddress, bridgeAddress, askAddress common.Address, err error) {
	proxyAddress, err = e.DeployContract(DOSProxy)
	if err != nil {
		return
	}

	bridgeAddress, err = e.DeployContract(DOSAddressBridge)
	if err != nil {
		return
	}

	if err = updataSDK(bridgeAddress.Hex()); err != nil {
		return
	}

	askAddress, err = e.DeployContract(AskMeAnyThing)
	if err != nil {
		return
	}

	err = e.updateBridge(bridgeAddress, proxyAddress)
	return
}

func updataSDK(bridgeAddress string) (err error) {
	sdkPath := workingDir + "/blockchain/eth/contracts/DOSOnChainSDK.sol"
	fmt.Println("SDKPath: ", sdkPath)

	sdkContent, err := ioutil.ReadFile(sdkPath)
	if err != nil {
		return
	}

	lines := strings.Split(string(sdkContent), "\n")
	for _, line := range lines {
		if strings.Contains(line, "0x") {
			parts := strings.Split(line, "(")
			parts[len(parts)-1] = "(" + bridgeAddress + ");"
			line = strings.Join(parts, "")
		}
	}

	newSdk := strings.Join(lines, "\n")
	err = ioutil.WriteFile(sdkPath, []byte(newSdk), 0644)
	return
}

func (e *EthAdaptor) updateBridge(bridgeAddress, proxyAddress common.Address) (err error) {
	fmt.Println("start to update proxy address to bridge...")

	auth, err := e.getAuth()
	if err != nil {
		return
	}

	bridge, err := dosproxy.NewDOSAddressBridge(bridgeAddress, e.client)
	if err != nil {
		return
	}

	tx, err := bridge.SetProxyAddress(auth, proxyAddress)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("proxy address updated in bridge")

	err = e.checkTransaction(tx)

	return
}

func (e *EthAdaptor) SubscribeToAll() (err error) {
	msgChan := make(chan interface{})

	for i := 0; i < 10; i++ {
		err = e.SubscribeEvent(msgChan, i)
	}

	go func() {
		defer close(msgChan)
		for msg := range msgChan {
			switch content := msg.(type) {
			case *DOSProxyLogGrouping:
				fmt.Println("got DOSProxyLogGrouping event...")
				fmt.Println("NodeId: ", content.NodeId)
				fmt.Println("----------------------------------------------")
			case *DOSProxyLogUrl:
				fmt.Println("got DOSProxyLogUrl event...")
				fmt.Println("QueryId: ", content.QueryId)
				fmt.Println("Url: ", content.Url)
				fmt.Println("Timeout: ", content.Timeout)
				fmt.Println("Randomness: ", content.Randomness)
				fmt.Println("DispatchedGroup: ", content.DispatchedGroup)
				fmt.Println("----------------------------------------------")
			case *DOSProxyLogUpdateRandom:
				fmt.Println("got DOSProxyLogUpdateRandom event...")
				fmt.Println("LastRandomness: ", content.LastRandomness)
				fmt.Println("LastUpdatedBlock: ", content.LastUpdatedBlock)
				fmt.Println("DispatchedGroup: ", content.DispatchedGroup)
				fmt.Println("----------------------------------------------")
			case *DOSProxyLogValidationResult:
				fmt.Println("got DOSProxyLogInvalidSignature event...")
				fmt.Println("TrafficType: ", content.TrafficType)
				fmt.Println("TrafficId: ", content.TrafficId)
				fmt.Println("Message: ", content.Message)
				fmt.Println("Signature: ", content.Signature)
				fmt.Println("PubKey: ", content.PubKey)
				fmt.Println("Pass: ", content.Pass)
				fmt.Println("----------------------------------------------")
			case *DOSProxyLogNonSupportedType:
				fmt.Println("got DOSProxyLogNonSupportedType event...")
				fmt.Println("QueryType: ", content.QueryType)
				fmt.Println("----------------------------------------------")
			case *DOSProxyLogNonContractCall:
				fmt.Println("got DOSProxyLogNonContractCall event...")
				fmt.Println("From: ", content.From)
				fmt.Println("----------------------------------------------")
			case *DOSProxyLogCallbackTriggeredFor:
				fmt.Println("got DOSProxyLogCallbackTriggeredFor event...")
				fmt.Println("CallbackAddr: ", content.CallbackAddr)
				fmt.Println("----------------------------------------------")
			case *DOSProxyLogQueryFromNonExistentUC:
				fmt.Println("got DOSProxyLogQueryFromNonExistentUC event...")
				fmt.Println("----------------------------------------------")
			case *DOSProxyLogInsufficientGroupNumber:
				fmt.Println("got DOSProxyLogInsufficientGroupNumber event...")
				fmt.Println("----------------------------------------------")
			case *DOSProxyLogPublicKeyAccepted:
				fmt.Println("got DOSProxyLogPublicKeyAccepted event...")
				fmt.Println("X1: ", content.X1)
				fmt.Println("X2: ", content.X2)
				fmt.Println("Y1: ", content.Y1)
				fmt.Println("Y2: ", content.Y2)
			}
		}
	}()
	return
}
