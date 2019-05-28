package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"os"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
)

const (
	keyPath      = "KEYFOLDERPATH"
	scanRange    = "SCANRANGE"
	scanInterval = "SCANINTERVAL"
)

func balanceMaintain(client *ethclient.Client, usrKey, rootKey *keystore.Key) (err error) {
	usrKeyBalance, err := getBalance(client, usrKey)
	if err != nil {
		return
	}
	fmt.Println("usrKeyBalance ", usrKeyBalance)

	rootKeyBalance, err := getBalance(client, rootKey)
	if err != nil {
		return
	}
	fmt.Println("rootKeyBalance ", rootKeyBalance)

	if usrKeyBalance.Cmp(big.NewFloat(onchain.REPLENISHTHRESHOLD)) == -1 {
		fmt.Println("userKey account replenishing...")
		err = transferEth(client, rootKey, usrKey)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			err = transferEth(client, rootKey, usrKey)
		}
		if err != nil {
			return
		}
		fmt.Println("userKey account replenished.")
	}

	return
}

func getBalance(client *ethclient.Client, key *keystore.Key) (balance *big.Float, err error) {
	wei, err := client.BalanceAt(context.Background(), key.Address, nil)
	if err != nil {
		return
	}

	balance = new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))
	return
}

func transferEth(client *ethclient.Client, from, to *keystore.Key) (err error) {
	nonce, err := client.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	value := big.NewInt(onchain.REPLENISHAMOUNT)
	gasLimit := uint64(onchain.GASLIMIT)

	tx := types.NewTransaction(nonce, to.Address, value, gasLimit, gasPrice, nil)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), from.PrivateKey)
	if err != nil {
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}
	fmt.Println("replenishing tx sent: ", signedTx.Hash().Hex(), ", waiting for confirmation...")

	err = onchain.CheckTransaction(client, signedTx)

	return
}

func main() {
	fmt.Println("starting balance maintain process...")
	config := configuration.Config{}
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	chainConfig := config.GetChainConfig()

	keyFolderPath := os.Getenv(keyPath)
	if keyFolderPath == "" {
		fmt.Println("No KEYFOLDERPATH Environment variable.")
		keyFolderPath = "/testAccounts"
	}

	passPhrase := os.Getenv(configuration.ENVPASSPHRASE)
	rootKey, err := onchain.ReadEthKey(keyFolderPath+"/bootCredential/fundKey", passPhrase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rootKey loaded")

	scanRange := os.Getenv(scanRange)
	scanRangeInt, err := strconv.Atoi(scanRange)
	if err != nil {
		log.Fatal(err)
	}

	scanInterval := os.Getenv(scanInterval)
	scanIntervalInt, err := strconv.Atoi(scanInterval)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connecting to eth...")
	client, err := ethclient.Dial(chainConfig.RemoteNodeAddressPool[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")

	var keyArray []*keystore.Key
	for i := 1; i <= scanRangeInt; i++ {
		key, err := onchain.ReadEthKey(keyFolderPath+"/"+strconv.Itoa(i)+"/credential", passPhrase)
		if err != nil {
			log.Fatal(err)
		}
		keyArray = append(keyArray, key)
		fmt.Println(strconv.Itoa(i) + "/" + scanRange + " key loaded")
		debug.FreeOSMemory()
		if err = balanceMaintain(client, key, rootKey); err != nil {
			log.Fatal(err)
		}
	}

	ticker := time.NewTicker(time.Duration(scanIntervalInt) * time.Hour)
	for {
		fmt.Println("wait for another " + scanInterval + " hour(s)")
		select {
		case <-ticker.C:
			for i, key := range keyArray {
				fmt.Println("checking " + strconv.Itoa(i+1) + "/" + scanRange + " key")
				if err = balanceMaintain(client, key, rootKey); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
