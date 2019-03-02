package onchain

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"

	"errors"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	GASLIMIT            = 5000000
	REPLENISHTHRESHOLD  = 0.6
	REPLENISHAMOUNT     = 800000000000000000 //0.8 Eth
	STOPSUBMITTHRESHOLD = 0.1
	RETRTCOUNT          = 2
)

func SetEthKey(credentialPath, passphrase string) (key *keystore.Key, err error) {
	fmt.Println("credentialPath: ", credentialPath)
	newKeyStore := keystore.NewKeyStore(credentialPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if len(newKeyStore.Accounts()) < 1 {
		_, err = newKeyStore.NewAccount(passphrase)
		if err != nil {
			return
		}
	}

	usrKeyPath := newKeyStore.Accounts()[0].URL.Path

	keyJson, err := ioutil.ReadFile(usrKeyPath)
	if err != nil {
		return
	}

	key, err = keystore.DecryptKey(keyJson, passphrase)

	return
}

func DialToEth(ctx context.Context, urlPool []string) (out chan *ethclient.Client, err chan error) {
	out = make(chan *ethclient.Client)
	err = make(chan error)
	var wg sync.WaitGroup

	multiplex := func(url string) {
		defer wg.Done()
		client, e := ethclient.Dial(url)
		if e != nil {
			select {
			case <-ctx.Done():
				return
			case err <- e:
				return
			}
		}
		select {
		case <-ctx.Done():
			client.Close()
			return
		case out <- client:
		}
	}

	// Select from all the channels
	wg.Add(len(urlPool))

	for _, url := range urlPool {
		go multiplex(url)
	}

	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(out)
		close(err)
	}()

	return
}

func CurrentBlock(client *ethclient.Client) (blknum uint64, err error) {
	var header *types.Header
	header, err = client.HeaderByNumber(context.Background(), nil)
	if err == nil {
		blknum = header.Number.Uint64()
	}
	return
}

func CheckTransaction(client *ethclient.Client, tx *types.Transaction) (err error) {
	start := time.Now()
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	for err == ethereum.NotFound {

		if time.Since(start).Seconds() > 90 {
			err = errors.New("transaction failed")
			fmt.Println("transaction failed. tx ", fmt.Sprintf("%x", tx.Hash()))
			return
		}

		time.Sleep(1 * time.Second)
		receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	}
	if err != nil {
		fmt.Println("CheckTransaction ", err)
		return
	}

	if receipt.Status == 1 {
		fmt.Println("transaction confirmed. tx ", fmt.Sprintf("%x", tx.Hash()))
	} else {
		err = errors.New("transaction failed")
		fmt.Println("transaction failed.  tx ", fmt.Sprintf("%x", tx.Hash()))
	}

	return
}

func Balance(client *ethclient.Client, key *keystore.Key) (balance *big.Float) {
	wei, err := client.BalanceAt(context.Background(), key.Address, nil)
	if err != nil {
		return
	}

	balance = new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))

	return balance
}
