package onchain

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"sync"
	"time"

	errors "golang.org/x/xerrors"

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

type DialResult struct {
	Client *ethclient.Client
	Url    string
	Err    error
}

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

func NumOfAccounts(credentialPath string) (n int) {
	newKeyStore := keystore.NewKeyStore(credentialPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if newKeyStore != nil {
		n = len(newKeyStore.Accounts())
	}
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
func reporeResult(ctx context.Context, out chan DialResult, result DialResult) {
	select {
	case <-ctx.Done():
	case out <- result:
	}
}

//DialToEth is a utility function to dial to Ethereum
func DialToEth(ctx context.Context, urlPool []string) (out chan DialResult) {
	out = make(chan DialResult)
	var wg sync.WaitGroup

	multiplex := func(url string) {
		r := DialResult{}
		var err error
		defer wg.Done()

		client, err := ethclient.Dial(url)
		if err != nil {
			//ws connect: connection timed out
			fmt.Println(url, ":DialToEth err ", err)
			r.Err = errors.Errorf("DialToEth: %w", err)
			reporeResult(ctx, out, r)
			return
		}

		id, err := client.NetworkID(ctx)
		if err != nil {
			fmt.Println("NetworkID err ", err)
			//Post http i/o timeout
			r.Err = errors.Errorf("DialToEth: %w", err)
			reporeResult(ctx, out, r)
			client.Close()
			return
		}
		blk, err := client.BlockByNumber(ctx, nil)
		if err != nil {
			fmt.Println("BlockByNumber error ", err, url)
			return
		}
		fmt.Println("highestBlkN ", blk.NumberU64())
		progress, err := client.SyncProgress(ctx)
		if err != nil {
			fmt.Println("SyncProgress err ", err, url)
			return
		}
		fmt.Println("progress ", progress)
		fmt.Println(url, "DialToEthr got a client ", id)
		r.Client = client
		r.Url = url
		reporeResult(ctx, out, r)
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
func CheckSync(ctx context.Context, cs chan *ethclient.Client) chan *ethclient.Client {
	out := make(chan *ethclient.Client)
	var wg sync.WaitGroup
	size := 0

	for client := range cs {
		size++
		go func(client *ethclient.Client) {
			defer wg.Done()
			var blk *types.Block
			var progress *ethereum.SyncProgress
			var err error
			if blk, err = client.BlockByNumber(ctx, nil); err != nil {
				fmt.Println("CheckSync error ", err)
				if err.Error() == light.ErrNoPeers.Error() {
					fmt.Println("CheckSync error ErrNoPeers ", err)
				}
				return
			}
			highestBlkN := blk.NumberU64()

			fmt.Println("highestBlkN ", highestBlkN)
			if progress, err = client.SyncProgress(ctx); err != nil {
				fmt.Println("SyncProgress err ", err)
				return
			}
			fmt.Println("progress ", progress)
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
func GetBalance(client *ethclient.Client, key *keystore.Key) (balance *big.Float, err error) {
	wei, err := client.BalanceAt(context.Background(), key.Address, nil)
	if err != nil {
		return
	}

	balance = new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))

	return
}
