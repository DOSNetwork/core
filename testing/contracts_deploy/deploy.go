package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/onchain/commitreveal"
	"github.com/DOSNetwork/core/onchain/dosbridge"
	"github.com/DOSNetwork/core/onchain/dosproxy"
	"github.com/DOSNetwork/core/testing/dosUser/contract"
	"github.com/DOSNetwork/core/testing/dosUser/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ProxyAddressType = iota
	CommitRevealAddressType
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

func updateBridgeAddr(path string, updated string) error {
	//path = path + "/DOSOnChainSDK.sol"
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

func updateBridge(client *ethclient.Client, key *keystore.Key, bridgeAddress common.Address, target common.Address) (err error) {
	var auth *bind.TransactOpts
	fmt.Println("start to update proxy address to bridge...")
	auth = bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(5000000)

	bridge, err := dosbridge.NewDOSAddressBridge(bridgeAddress, client)
	if err != nil {
		return
	}
	var tx *types.Transaction

	tx, err = bridge.SetProxyAddress(auth, target)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = bridge.SetProxyAddress(auth, target)
	}

	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("proxy address updated in bridge")

	err = onchain.CheckTransaction(client, tx)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func SetCommitrevealLibToProxy(client *ethclient.Client, key *keystore.Key, proxyAddress common.Address, crAddress common.Address) (err error) {
	var auth *bind.TransactOpts
	fmt.Println("start to update proxy address to bridge...")
	auth = bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(5000000)

	proxy, err := dosproxy.NewDOSProxy(proxyAddress, client)
	if err != nil {
		return
	}
	var tx *types.Transaction

	tx, err = proxy.SetCommitrevealLib(auth, crAddress)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = proxy.SetCommitrevealLib(auth, crAddress)
	}

	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("proxy address updated in bridge")

	err = onchain.CheckTransaction(client, tx)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func AddProxyToCRWhiteList(client *ethclient.Client, key *keystore.Key, crAddress common.Address, proxyAddress common.Address) (err error) {
	var auth *bind.TransactOpts
	fmt.Println("start to update proxy address to bridge...")
	auth = bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(5000000)

	cr, err := commitreveal.NewCommitReveal(crAddress, client)
	if err != nil {
		return
	}
	var tx *types.Transaction

	tx, err = cr.AddToWhitelist(auth, proxyAddress)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = cr.AddToWhitelist(auth, crAddress)
	}

	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("add proxy address to commitReveal whitelist")

	err = onchain.CheckTransaction(client, tx)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func deployContract(contractPath, targetTask, configPath string, config configuration.Config, chainConfig configuration.ChainConfig, client *ethclient.Client, key *keystore.Key) {
	var tx *types.Transaction
	var address common.Address
	auth := bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(5000000)
	var err error
	if targetTask == "deployBridge" {
		fmt.Println("Starting deploy DOSAddressBridge.sol...")
		address, tx, _, err = dosbridge.DeployDOSAddressBridge(auth, client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = dosbridge.DeployDOSAddressBridge(auth, client)
		}
		fmt.Println("contract Address: ", address.Hex())
		fmt.Println("tx sent: ", tx.Hash().Hex())
		fmt.Println("contract deployed, waiting for confirmation...")
		err = onchain.CheckTransaction(client, tx)
		if err != nil {
			fmt.Println(err)
		}
		chainConfig.DOSAddressBridgeAddress = address.Hex()
		if err := config.UpdateConfig(chainConfig); err != nil {
			log.Fatal(err)
		}
		if err := updateBridgeAddr(contractPath+"/DOSOnChainSDK.sol", chainConfig.DOSAddressBridgeAddress); err != nil {
			log.Fatal(err)
		}
	} else if targetTask == "deployCommitReveal" {
		fmt.Println("Starting deploy CommitReveal.sol...")
		address, tx, _, err = commitreveal.DeployCommitReveal(auth, client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = commitreveal.DeployCommitReveal(auth, client)
		}
		fmt.Println("contract Address: ", address.Hex())
		fmt.Println("tx sent: ", tx.Hash().Hex())
		fmt.Println("contract deployed, waiting for confirmation...")
		err = onchain.CheckTransaction(client, tx)
		if err != nil {
			fmt.Println(err)
		}
		chainConfig.DOSCommitReveal = address.Hex()
		if err := config.UpdateConfig(chainConfig); err != nil {
			log.Fatal(err)
		}

	} else if targetTask == "deployProxy" {
		fmt.Println("Starting deploy DOSProxy.sol...")
		address, tx, _, err = dosproxy.DeployDOSProxy(auth, client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = dosproxy.DeployDOSProxy(auth, client)
		}
		fmt.Println("contract Address: ", address.Hex())
		fmt.Println("tx sent: ", tx.Hash().Hex())
		fmt.Println("contract deployed, waiting for confirmation...")
		err = onchain.CheckTransaction(client, tx)
		if err != nil {
			fmt.Println(err)
		}
		chainConfig.DOSProxyAddress = address.Hex()
		if err := config.UpdateConfig(chainConfig); err != nil {
			log.Fatal(err)
		}

	} else if targetTask == "setUpDOSNetwork" {

		bridgeAddress := common.HexToAddress(chainConfig.DOSAddressBridgeAddress)
		proxyAddress := common.HexToAddress(chainConfig.DOSProxyAddress)
		crAddress := common.HexToAddress(chainConfig.DOSCommitReveal)

		err := updateBridge(client, key, bridgeAddress, proxyAddress)
		if err != nil {
			log.Fatal(err)
		}

		_ = SetCommitrevealLibToProxy(client, key, proxyAddress, crAddress)

		err = AddProxyToCRWhiteList(client, key, crAddress, proxyAddress)
		if err != nil {
			log.Fatal(err)
		}

	} else if targetTask == "deployAMA" {
		fmt.Println("Starting deploy AskMeAnyThing.sol...")
		amaConfig := eth.AMAConfig{}
		amaCount := os.Getenv("AMACOUNT")
		if amaCount == "" {
			fmt.Println("No AMACOUNT Environment variable.")
			amaCount = "1"
		}
		count, err := strconv.Atoi(amaCount)
		if err != nil {
			fmt.Println(err)
			return
		}

		var amaAdd []string
		for i := 0; i < count; i++ {
			address, tx, _, err = dosUser.DeployAskMeAnything(auth, client)
			for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
				fmt.Println(err)
				time.Sleep(time.Second)
				fmt.Println("transaction retry...")
				address, tx, _, err = dosUser.DeployAskMeAnything(auth, client)
			}
			fmt.Println("contract Address: ", address.Hex())
			fmt.Println("tx sent: ", tx.Hash().Hex())
			fmt.Println("contract deployed, waiting for confirmation...")
			err = onchain.CheckTransaction(client, tx)
			if err != nil {
				fmt.Println(err)
			}
			amaAdd = append(amaAdd, address.Hex())
		}
		amaConfig.AskMeAnythingAddressPool = amaAdd
		if err := configuration.UpdateConfig(configPath, amaConfig); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	credentialPathPtr := flag.String("credentialPath", "./testAccounts/bootCredential/fundKey", "credential path")
	contractPathPtr := flag.String("contractPath", "./contracts", "Contract file path")
	targetPtr := flag.String("target", "DOSProxy", "DOSProxy or AMA or SimpleDice")
	configPathPtr := flag.String("configPath", "", "config path")

	flag.Parse()
	credentialPath := *credentialPathPtr
	contractPath := *contractPathPtr
	target := *targetPtr
	configPath := *configPathPtr

	if err := os.Setenv("CONFIGPATH", ".."); err != nil {
		log.Fatal(err)
	}
	//Get
	config := configuration.Config{}
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	chainConfig := config.GetChainConfig()
	fmt.Println("Use : ", chainConfig)
	fmt.Println(" address ", chainConfig.RemoteNodeAddressPool[0])

	password := os.Getenv("PASSPHRASE")
	for password == "" {
		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(0)
		if err == nil {
			fmt.Println("\nPassword typed: ***\n")
		}
		password = strings.TrimSpace(string(bytePassword))
	}
	//Dial to blockchain
	key, err := onchain.ReadEthKey(credentialPath, password)
	if err != nil {
		fmt.Println("NewETHProxySession ", err)
		return
	}

	clients := onchain.DialToEth(context.Background(), chainConfig.RemoteNodeAddressPool[:1], key)

	//Use first client
	c, ok := <-clients
	if !ok {
		err = errors.New("No any working eth client")
		return
	}

	deployContract(contractPath, target, configPath, config, chainConfig, c, key)
}
