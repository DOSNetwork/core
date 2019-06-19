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

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/light"
)

const (
	checkSyncInterval = 15
	syncBlockDiff     = 3
)

func first(ctx context.Context, source <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		first := true
		for val := range source {
			if first {
				if val != nil {
					first = false
					select {
					case <-ctx.Done():
						return
					case out <- val:
					}
					return
				}
			}
		}
	}()
	return out
}

func merge(ctx context.Context, cs ...chan interface{}) chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan interface{}) {
		for n := range c {
			out <- n
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func GenEthkey(credentialPath, passPhrase string) (err error) {
	newKeyStore := keystore.NewKeyStore(credentialPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if len(newKeyStore.Accounts()) >= 1 {
		fmt.Println("", newKeyStore.Accounts()[0].URL.Path)
		err = errors.New("Found an existing key in " + credentialPath)
		return
	}
	_, err = newKeyStore.NewAccount(passPhrase)

	return
}

//ReadEthKey is a utility function to read a keystore file
func ReadEthKey(credentialPath, passphrase string) (key *keystore.Key, err error) {
	newKeyStore := keystore.NewKeyStore(credentialPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if len(newKeyStore.Accounts()) < 1 {
		return nil, errors.New("No Account")
	}

	usrKeyPath := newKeyStore.Accounts()[0].URL.Path

	keyJson, err := ioutil.ReadFile(usrKeyPath)
	if err != nil {
		return
	}

	key, err = keystore.DecryptKey(keyJson, passphrase)

	return

}

//DialToEth is a utility function to dial to Ethereum
func DialToEth(ctx context.Context, urlPool []string) (out chan *ethclient.Client) {
	out = make(chan *ethclient.Client)
	var wg sync.WaitGroup

	multiplex := func(url string) {
		var e error
		var client *ethclient.Client
		defer wg.Done()

		client, e = ethclient.Dial(url)
		if e != nil {
			//ws connect: connection timed out
			fmt.Println(url, ":DialToEth err ", e)
			return
		}
		_, err := client.NetworkID(ctx)
		if err != nil {
			//Post http i/o timeout
			fmt.Println(url, "NetworkID err ", err)
			client.Close()
			return
		}
		fmt.Println(url, "DialToEthr got a client")
		select {
		case <-ctx.Done():
			client.Close()
			return
		case out <- client:
		}
	}

	// Select from all the channels
	wg.Add(len(urlPool))

	go func() {
		for i := 0; i < len(urlPool); i++ {
			go multiplex(urlPool[i])
		}
	}()

	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(out)
	}()

	return
}

//CheckSync is a utility function to check sync state
func CheckSync(ctx context.Context, mClient *ethclient.Client, cs chan *ethclient.Client) chan *ethclient.Client {
	out := make(chan *ethclient.Client)
	var wg sync.WaitGroup
	size := 0

	for client := range cs {
		size++
		go func(client *ethclient.Client) {
			//defer logger.TimeTrack(time.Now(), "CheckSync", nil)

			defer wg.Done()
			ticker := time.NewTicker(time.Second * time.Duration(checkSyncInterval))
			for range ticker.C {
				highestBlk, e := mClient.BlockByNumber(ctx, nil)
				if e != nil {
					fmt.Println(e)
					if e.Error() == light.ErrNoPeers.Error() {
						continue
					} else {
						return
					}
				}
				highestBlkN := highestBlk.NumberU64()
				currBlk, e := client.BlockByNumber(ctx, nil)
				if e != nil {
					fmt.Println(e)
					if e.Error() == light.ErrNoPeers.Error() {
						continue
					} else {
						return
					}
				}
				currBlkN := currBlk.NumberU64()
				fmt.Println("highestBlkN ", highestBlkN, "  currBlkN ", currBlkN)
				blockDiff := math.Abs(float64(highestBlkN) - float64(currBlkN))
				fmt.Println("block to Sync ", blockDiff)
				if blockDiff <= syncBlockDiff {
					ticker.Stop()
					out <- client
					return
				}
			}
		}(client)
	}
	wg.Add(size)
	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// GetCurrentBlock returns a block number of latest known header from the
// current canonical chain.
func GetCurrentBlock(client *ethclient.Client) (blknum uint64, err error) {
	var header *types.Header
	header, err = client.HeaderByNumber(context.Background(), nil)
	if err == nil {
		blknum = header.Number.Uint64()
	}
	return
}

// CheckTransaction return an error if the transaction is failed
func CheckTransaction(client *ethclient.Client, tx *types.Transaction) (err error) {
	start := time.Now()
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	for err == ethereum.NotFound {
		fmt.Println("ethereum.NotFound ", err)
		if time.Since(start).Minutes() > 30 {
			err = errors.New("transaction not found")
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
		fmt.Println("transaction Status != 0 .  tx ", fmt.Sprintf("%x", tx.Hash()))
	}

	return
}

// GetBalance returns the wei balance of the given account.
func GetBalance(client *ethclient.Client, key *keystore.Key) (balance *big.Float) {
	wei, err := client.BalanceAt(context.Background(), key.Address, nil)
	if err != nil {
		return
	}

	balance = new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))

	return balance
}
