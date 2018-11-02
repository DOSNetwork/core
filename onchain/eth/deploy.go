package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/onchain/eth/contracts"
	"github.com/DOSNetwork/core/onchain/eth/contracts/userContract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = `{"address":"2a3b59ac638f90d82bdaf5e2da5d37c1a31b29f3","crypto":{"cipher":"aes-128-ctr","ciphertext":"35000234819a4c42415922281ca492eb27eceb09d12fdd15b145a8b9cbc585b4","cipherparams":{"iv":"d65304cd8d8098a4e286ec589d746c1c"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"dbd946ab7d2e4ae9646c2abdb7260d64b1c217b99b699dcd30ef46d24daf4f48"},"mac":"476f013db27f7fd37fb45c6224e80833771bbd764e42032bf1d02a4de58783e4"},"id":"e9cd78c8-feb8-4f6e-9512-22dd5187b72e","version":3}`
const passPhrase = "123"
const (
	AskMeAnyThing = iota
	DOSAddressBridge
	DOSProxy
	DOSOnChainSDK
)

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func updateSDK(path string, updated string) error {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "DOSAddressBridgeInterface(") {
			out := between(lines[i], "(", ")")
			result := strings.Replace(lines[i], out, updated, -1)
			lines[i] = result
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func getAuth(client *ethclient.Client) (auth *bind.TransactOpts, err error) {

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err = bind.NewTransactor(strings.NewReader(key), passPhrase)
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func checkTransaction(client *ethclient.Client, tx *types.Transaction) (err error) {
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	for err == ethereum.NotFound {
		time.Sleep(1 * time.Second)
		receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	if receipt.Status == 1 {
		fmt.Println("transaction confirmed.")
	} else {
		err = errors.New("transaction failed")
	}

	return
}

func deployContract(client *ethclient.Client, contractName int) string {
	var tx *types.Transaction
	var address common.Address
	auth, err := getAuth(client)

	switch contractName {
	case DOSAddressBridge:
		fmt.Println("Starting deploy DOSAddressBridge.sol...")
		address, tx, _, err = dosproxy.DeployDOSAddressBridge(auth, client)
	case DOSProxy:
		fmt.Println("Starting deploy DOSProxy.sol...")
		address, tx, _, err = dosproxy.DeployDOSProxy(auth, client)
	case DOSOnChainSDK:
		fmt.Println("Starting deploy DeployDOSOnChainSDK.sol...")
		address, tx, _, err = dosproxy.DeployDOSOnChainSDK(auth, client)
	case AskMeAnyThing:
		fmt.Println("Starting deploy AskMeAnyThing.sol...")
		address, tx, _, err = userContract.DeployAskMeAnything(auth, client)
	}
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println("contract Address: ", address.Hex())
	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("contract deployed, waiting for confirmation...")

	err = checkTransaction(client, tx)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return address.Hex()
}

func updateBridge(client *ethclient.Client, bridgeAddress, proxyAddress common.Address) (err error) {
	var auth *bind.TransactOpts
	fmt.Println("start to update proxy address to bridge...")
	auth, err = getAuth(client)
	if err != nil {
		return
	}
	bridge, err := dosproxy.NewDOSAddressBridge(bridgeAddress, client)
	if err != nil {
		return
	}

	tx, err := bridge.SetProxyAddress(auth, proxyAddress)
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("proxy address updated in bridge")

	err = checkTransaction(client, tx)

	return
}

func main() {
	contractPathFlag := flag.String("contractPath", "./contracts", "Contract file path")
	configPathFlag := flag.String("configPath", "./config.json", "Configuration file path")
	stepFlag := flag.String("step", "ProxyAndBridge", "ProxyAndBridge or SDKAndAMA")
	flag.Parse()
	contractPath := *contractPathFlag
	configPath := *configPathFlag
	contractPath = contractPath + "/DOSOnChainSDK.sol"
	step := *stepFlag

	config := configuration.ReadConfig()
	chainConfig := configuration.GetOnChainConfig(config)
	fmt.Println("dial to ", chainConfig.RemoteNodeAddress)
	client, err := ethclient.Dial(chainConfig.RemoteNodeAddress)
	if err != nil {
		log.Fatal(err)
	}
	if step == "ProxyAndBridge" {
		chainConfig.DOSProxyAddress = deployContract(client, DOSProxy)
		chainConfig.DOSAddressBridgeAddress = deployContract(client, DOSAddressBridge)
		configuration.UpdateOnChainConfig(configPath, config, chainConfig)
		updateSDK(contractPath, chainConfig.DOSAddressBridgeAddress)
	} else if step == "SDKAndAMA" {
		chainConfig.DOSOnChainSDKAddress = deployContract(client, DOSOnChainSDK)
		chainConfig.AskMeAnythingAddress = deployContract(client, AskMeAnyThing)
		configuration.UpdateOnChainConfig(configPath, config, chainConfig)
	} else if step == "SetProxyAddress" {
		bridgeAddress := common.HexToAddress(chainConfig.DOSAddressBridgeAddress)
		proxyAddress := common.HexToAddress(chainConfig.DOSProxyAddress)
		err := updateBridge(client, bridgeAddress, proxyAddress)
		if err != nil {
			log.Fatal(err)
		}
	}
}
