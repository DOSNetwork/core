package onchain

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/DOSNetwork/core/configuration"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	GASLIMIT            = 3000000
	REPLENISHTHRESHOLD  = 0.7
	REPLENISHAMOUNT     = 800000000000000000 //0.8 Eth
	STOPSUBMITTHRESHOLD = 0.1
)

type EthCommon struct {
	key    *keystore.Key
	Client *ethclient.Client
	lock   *sync.Mutex
	//ethNonce uint64
	config *configuration.ChainConfig
}

func (e *EthCommon) DialToEth() (err error) {
	fmt.Println("dialing...")
	e.Client, err = ethclient.Dial(e.config.RemoteNodeAddress)
	for err != nil {
		log.WithField("function", "dialToEth").Warn(err)
		fmt.Println("Cannot connect to the network, retrying...", e.config.RemoteNodeAddress)
		e.Client, err = ethclient.Dial(e.config.RemoteNodeAddress)
	}
	return
}

func (e *EthCommon) Init(credentialPath string, config *configuration.ChainConfig) (err error) {
	e.config = config
	e.lock = new(sync.Mutex)

	fmt.Println("start initial onChainConn...", config.DOSProxyAddress)

	if err = e.DialToEth(); err != nil {
		log.WithField("function", "dialToEth").Warn(err)
	}
	if err = e.setAccount(credentialPath); err != nil {
		log.WithField("function", "setAccount").Warn(err)
	}
	return
}

func (e *EthCommon) setAccount(credentialPath string) (err error) {
	fmt.Println("credentialPath: ", credentialPath)

	passPhraseBytes, err := ioutil.ReadFile(credentialPath + "/passPhrase")
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
	//e.ethNonce, err = e.Client.PendingNonceAt(context.Background(), e.key.Address)
	//if err != nil {
	//	return
	//}
	////for correctness of the first call of GetAuth, because GetAuth always ++,
	//e.ethNonce--

	return
}

func (e *EthCommon) GetAuth() (auth *bind.TransactOpts, err error) {
	auth = bind.NewKeyedTransactor(e.key.PrivateKey)
	if err != nil {
		return
	}

	//e.lock.Lock()
	//e.ethNonce++
	//automatedNonce, err := e.Client.PendingNonceAt(context.Background(), e.key.Address)
	//if err != nil {
	//	return
	//}
	//
	//if automatedNonce > e.ethNonce {
	//	e.ethNonce = automatedNonce
	//}
	//
	//auth.Nonce = big.NewInt(int64(e.ethNonce))
	//fmt.Println(e.ethNonce)
	//e.lock.Unlock()

	auth.GasLimit = uint64(GASLIMIT)
	return
}

func (e *EthCommon) getKey(keyPath, passPhrase string) (key *keystore.Key, err error) {
	var keyJson []byte
	keyJson, err = ioutil.ReadFile(keyPath)
	if err != nil {
		return
	}

	key, err = keystore.DecryptKey(keyJson, passPhrase)
	if err != nil {
		return
	}
	return
}

func (e *EthCommon) GetAddress() (key common.Address) {
	return e.key.Address
}

func (e *EthCommon) BalanceMaintain(rootKeyPath, usrKeyPath, rPassPhrase, uPassPhrase string) (err error) {
	fmt.Println("EthCommon BalanceMaintain")

	rootKey, err := e.getKey(rootKeyPath, rPassPhrase)
	if err != nil {
		return
	}
	usrKey, err := e.getKey(usrKeyPath, uPassPhrase)
	if err != nil {
		return
	}
	err = e.balanceMaintain(usrKey, rootKey)
	return
}

func (e *EthCommon) balanceMaintain(usrKey, rootKey *keystore.Key) (err error) {
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

	if usrKeyBalance.Cmp(big.NewFloat(REPLENISHTHRESHOLD)) == -1 {
		fmt.Println("userKey account replenishing...")
		if err = e.transferEth(rootKey, usrKey); err == nil {
			fmt.Println("userKey account replenished.")
		}
	}

	return
}

func (e *EthCommon) EnoughBalance(address common.Address) (isEnough bool) {
	balance, err := e.getBalance(&keystore.Key{Address: address})
	if err != nil {
		log.Warn(err)
		return
	}

	return balance.Cmp(big.NewFloat(STOPSUBMITTHRESHOLD)) != -1
}

func (e *EthCommon) getBalance(key *keystore.Key) (balance *big.Float, err error) {
	wei, err := e.Client.BalanceAt(context.Background(), key.Address, nil)
	if err != nil {
		return
	}

	balance = new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))
	return
}

func (e *EthCommon) transferEth(from, to *keystore.Key) (err error) {
	nonce, err := e.Client.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		return
	}

	gasPrice, err := e.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	value := big.NewInt(REPLENISHAMOUNT)
	gasLimit := uint64(GASLIMIT)

	tx := types.NewTransaction(nonce, to.Address, value, gasLimit, gasPrice, nil)

	chainId, err := e.Client.NetworkID(context.Background())
	if err != nil {
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), from.PrivateKey)
	if err != nil {
		return
	}

	err = e.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}
	fmt.Println("replenishing tx sent: ", signedTx.Hash().Hex(), ", waiting for confirmation...")

	err = e.CheckTransaction(signedTx)

	return
}

func (e *EthCommon) CheckTransaction(tx *types.Transaction) (err error) {
	//start := time.Now()
	receipt, err := e.Client.TransactionReceipt(context.Background(), tx.Hash())
	for err == ethereum.NotFound {
		/*
			if time.Since(start).Seconds() > 30 {
				fmt.Println("no receipt yet, set to successful")
				return
			}
		*/
		time.Sleep(1 * time.Second)
		receipt, err = e.Client.TransactionReceipt(context.Background(), tx.Hash())
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
