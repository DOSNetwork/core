package onchain

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

	"github.com/DOSNetwork/core/group/bn256"
	"github.com/DOSNetwork/core/onchain/eth/contracts"

	"github.com/dedis/kyber"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SubscribeDOSProxyLogUrl = iota
	SubscribeDOSProxyLogRequestUserRandom
	SubscribeDOSProxyLogNonSupportedType
	SubscribeDOSProxyLogNonContractCall
	SubscribeDOSProxyLogCallbackTriggeredFor
	SubscribeDOSProxyLogQueryFromNonExistentUC
	SubscribeDOSProxyLogUpdateRandom
	SubscribeDOSProxyLogValidationResult
	SubscribeDOSProxyLogInsufficientGroupNumber
	SubscribeDOSProxyLogGrouping
	SubscribeDOSProxyLogPublicKeyAccepted
	SubscribeDOSProxyWhitelistAddressTransferred
)

// TODO: Move constants to some unified places.
const (
	TrafficSystemRandom = iota // 0
	TrafficUserRandom
	TrafficUserQuery
)

const balanceCheckInterval = 3

var workingDir string

type EthAdaptor struct {
	id        *big.Int
	key       *keystore.Key
	client    *ethclient.Client
	proxy     *dosproxy.DOSProxy
	lock      *sync.Mutex
	logFilter *sync.Map
	ethNonce  uint64
	config    *ChainConfig
}

func (e *EthAdaptor) Init(autoReplenish bool, config *ChainConfig) (err error) {
	workingDir, err = os.Getwd()
	if err != nil {
		return
	}

	e.config = config

	fmt.Println("start initial onChainConn...", config.DOSProxyAddress)

	e.logFilter = new(sync.Map)
	go e.logMapTimeout()

	e.lock = new(sync.Mutex)

	e.dialToEth()

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
			e.dialToEth()
			go e.subscribeEventAttempt(ch, opt, subscribeType, done)

		}
	}
}

func (e *EthAdaptor) dialToEth() (err error) {
	e.lock.Lock()
	fmt.Println("dialing...")
	e.client, err = ethclient.Dial(e.config.RemoteNodeAddress)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Cannot connect to the network, retrying...")
		e.client, err = ethclient.Dial(e.config.RemoteNodeAddress)
	}
	addr := common.HexToAddress(e.config.DOSProxyAddress)
	e.proxy, err = dosproxy.NewDOSProxy(addr, e.client)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Connot Create new proxy, retrying...")
		e.proxy, err = dosproxy.NewDOSProxy(addr, e.client)
	}
	e.lock.Unlock()
	return
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
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogGrouping{
						NodeId: i.NodeId,
					}
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
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogUrl{
						QueryId:         i.QueryId,
						Timeout:         i.Timeout,
						DataSource:      i.DataSource,
						Selector:        i.Selector,
						Randomness:      i.Randomness,
						DispatchedGroup: i.DispatchedGroup,
					}
				}
			}
		}

	case SubscribeDOSProxyLogRequestUserRandom:
		fmt.Println("subscribing DOSProxyLogRequestUserRandom event...")
		transitChan := make(chan *dosproxy.DOSProxyLogRequestUserRandom)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogRequestUserRandom(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyLogRequestUserRandom event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogRequestUserRandom{
						RequestId:            i.RequestId,
						LastSystemRandomness: i.LastSystemRandomness,
						UserSeed:             i.UserSeed,
						DispatchedGroup:      i.DispatchedGroup,
					}
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
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogUpdateRandom{
						LastRandomness:  i.LastRandomness,
						DispatchedGroup: i.DispatchedGroup,
					}
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
				if !e.filterLog(i.Raw) {
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
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogNonSupportedType{
						InvalidSelector: i.InvalidSelector,
					}
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
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogNonContractCall{
						From: i.From,
					}
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
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogCallbackTriggeredFor{
						CallbackAddr: i.CallbackAddr,
					}
				}
			}
		}
	case SubscribeDOSProxyLogQueryFromNonExistentUC:
		fmt.Println("subscribing DOSProxyLogQueryFromNonExistentUC event...")
		transitChan := make(chan *dosproxy.DOSProxyLogRequestFromNonExistentUC)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogRequestFromNonExistentUC(opt, transitChan)
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
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogRequestFromNonExistentUC{}
				}
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
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogInsufficientGroupNumber{}
				}
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
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogPublicKeyAccepted{
						X1: i.X1,
						X2: i.X2,
						Y1: i.Y1,
						Y2: i.Y2,
					}
				}
			}
		}
	case SubscribeDOSProxyWhitelistAddressTransferred:
		fmt.Println("subscribing DOSProxyWhitelistAddressTransferred event...")
		transitChan := make(chan *dosproxy.DOSProxyWhitelistAddressTransferred)
		sub, err := e.proxy.DOSProxyFilterer.WatchWhitelistAddressTransferred(opt, transitChan)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}
		fmt.Println("DOSProxyWhitelistAddressTransferred event subscribed")

		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyWhitelistAddressTransferred{
						Previous: i.Previous,
						Curr:     i.Curr,
					}
				}
			}
		}
	}

}

func (e *EthAdaptor) InitialWhiteList() (err error) {
	fmt.Println("Starting initialing WhiteList...")
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	addresses := [21]common.Address{common.HexToAddress("0xaec1f213677de24842a96e72fe6efbcbc2b77ca5"),
		common.HexToAddress("0x5a3582d7c8fc97c194168342f9a347e8d41b4038"),
		common.HexToAddress("0xe38497ec6d9442413b16a2a7055649ac91106d79"),
		common.HexToAddress("0xb8c228a842390595751b1a09a2c4ffaccf12c44b"),
		common.HexToAddress("0xec96405d35c3cadfd3e0bda90c66b74b159d4156"),
		common.HexToAddress("0x0a65d5dccc87ecb21bc4b24c16a3c01e0cdd42ac"),
		common.HexToAddress("0x3ebe227e9fd42bb97b9a950e4a731d8975263812"),
		common.HexToAddress("0x6ca3ee1386f7c05d886211a6378f49bdf9c7ee88"),
		common.HexToAddress("0x234ae33713afde52de9aa9203fb6696531edc74f"),
		common.HexToAddress("0xe610c52cbeb14dc722a9668e59bad46c5e464e55"),
		common.HexToAddress("0xcb04aae925218094863809ec0289a8fdccfd68cf"),
		common.HexToAddress("0x921f2cf348b8b45d6cd5eaf139d30303e6b9646f"),
		common.HexToAddress("0x69ba6867602d650fc433fd62eabaf17f11fd5132"),
		common.HexToAddress("0x4ed2814cd63e83504221424215d0655c6db0b674"),
		common.HexToAddress("0xac3e1e84e3b7a0a83d36feba984a836c768fcb72"),
		common.HexToAddress("0x3609aa202ab8b96499e379da7226145e5697695e"),
		common.HexToAddress("0xc96c0f2d346a3f0f4bbcfbefe70a75c710d23369"),
		common.HexToAddress("0xc1a3aab78a6dc3e5cbeec059a2727a66f4bc7088"),
		common.HexToAddress("0xc98a4df797f0b8155ef38c4701588cbed21a1b26"),
		common.HexToAddress("0x34d950db8e9345a638ba0bee9945d56c9f7728ee"),
		common.HexToAddress("0x96272c390ae674d3a3e3f1d636f3ae4128afd688")}

	tx, err := e.proxy.InitWhitelist(auth, addresses)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("Whitelist initialized, waiting for confirmation...")

	err = e.checkTransaction(tx)

	return

}

func (e *EthAdaptor) GetGroupPubKey(idx int) (groupPubKeys [4]*big.Int, err error) {
	return e.proxy.GetGroupPubKey(&bind.CallOpts{}, big.NewInt(int64(idx)))
}

func (e *EthAdaptor) Grouping(size int) (err error) {
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.Grouping(auth, big.NewInt(int64(size)))
	if err != nil {
		return
	}

	err = e.checkTransaction(tx)
	return
}

func (e *EthAdaptor) GetWhitelist() (address common.Address, err error) {
	return e.proxy.GetWhitelistAddress(&bind.CallOpts{}, big.NewInt(1))
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

	x, y := DecodeSig(sig)

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

	x0, x1, y0, y1, err := DecodePubKey(pubKey)
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

func (e *EthAdaptor) RandomNumberTimeOut() (err error) {
	fmt.Println("Starting RandomNumberTimeOut...")
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.HandleTimeout(auth)
	if err != nil {
		return
	}

	err = e.checkTransaction(tx)
	return
}

func DecodePubKey(pubKey kyber.Point) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
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

func (e *EthAdaptor) DataReturn(requestId *big.Int, trafficType uint8, data, sig []byte) (err error) {
	auth, err := e.getAuth()
	if err != nil {
		return
	}

	x, y := DecodeSig(sig)

	tx, err := e.proxy.TriggerCallback(auth, requestId, trafficType, data, [2]*big.Int{x, y})
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Printf("Request with id(%v) type(%v) fulfilled, waiting for confirmation...\n", requestId, trafficType)

	err = e.checkTransaction(tx)

	return
}

func DecodeSig(sig []byte) (x, y *big.Int) {
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
	auth = bind.NewKeyedTransactor(e.key.PrivateKey)
	if err != nil {
		return
	}

	automatedNonce, err := e.client.PendingNonceAt(context.Background(), e.key.Address)
	if err != nil {
		return
	}

	e.lock.Lock()
	e.ethNonce++
	if automatedNonce > e.ethNonce {
		e.ethNonce = automatedNonce
	}
	auth.Nonce = big.NewInt(int64(e.ethNonce))
	e.lock.Unlock()

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(3))
	auth.GasLimit = uint64(1000000)

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

func (e *EthAdaptor) getKey(keyPath, passPrase string) (key *keystore.Key, err error) {
	var keyJson []byte
	keyJson, err = ioutil.ReadFile(keyPath)
	if err != nil {
		return
	}

	key, err = keystore.DecryptKey(keyJson, passPrase)
	if err != nil {
		return
	}
	return
}

func (e *EthAdaptor) BalanceMaintain(rootKeyPath, usrKeyPath, rPassPhrase, uPassPhrase string) (err error) {
	rootKey, err := e.getKey(rootKeyPath, rPassPhrase)
	if err != nil {
		return
	}
	userKey, err := e.getKey(usrKeyPath, uPassPhrase)
	if err != nil {
		return
	}
	err = e.balanceMaintain(rootKey, userKey)
	return
}

func (e *EthAdaptor) balanceMaintain(usrKey, rootKey *keystore.Key) (err error) {
	usrKeyBalance, err := e.getBalance(usrKey)
	if err != nil {
		return
	}
	fmt.Println("usrKeyBalance ", usrKeyBalance)

	rootKeyBalance, err := e.getBalance(rootKey)
	if err != nil {
		return
	}
	fmt.Println("rootKeyBalance ", rootKeyBalance)

	if usrKeyBalance.Cmp(big.NewFloat(0.7)) == -1 {
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

	value := big.NewInt(800000000000000000) //0.8 Eth
	gasLimit := uint64(1000000)

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

func (e *EthAdaptor) checkTransaction(tx *types.Transaction) (err error) {
	start := time.Now()
	receipt, err := e.client.TransactionReceipt(context.Background(), tx.Hash())
	for err == ethereum.NotFound {
		if time.Since(start).Seconds() > 30 {
			fmt.Println("no receipt yet, set to successful")
			return
		}

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

func (e *EthAdaptor) SubscribeToAll(msgChan chan interface{}) (err error) {
	for i := SubscribeDOSProxyLogUrl; i <= SubscribeDOSProxyWhitelistAddressTransferred; i++ {
		err = e.SubscribeEvent(msgChan, i)
	}
	return
}

type logRecord struct {
	content       types.Log
	currTimeStamp time.Time
}

func (e *EthAdaptor) filterLog(raw types.Log) (duplicates bool) {
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

func (e *EthAdaptor) logMapTimeout() {
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		e.logFilter.Range(e.checkTime)
	}

}

func (e *EthAdaptor) checkTime(log, deliverTime interface{}) (okToDelete bool) {
	switch t := deliverTime.(type) {
	case logRecord:
		if time.Now().Sub(t.currTimeStamp).Seconds() > 60*10 {
			e.logFilter.Delete(log)
		}
	}
	return true
}
