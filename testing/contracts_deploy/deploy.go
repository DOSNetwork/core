package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/onchain/commitreveal"
	"github.com/DOSNetwork/core/onchain/dosbridge"
	"github.com/DOSNetwork/core/onchain/dospayment"
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
	proxyAddressType = iota
	crAddressType
	paymentAddressType
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

func updateBridge(client *ethclient.Client, key *keystore.Key, addrType int, bridgeAddress, addr common.Address) (err error) {
	var auth *bind.TransactOpts
	fmt.Println("start to update proxy address to bridge...")
	auth = bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(6000000)
	auth.GasPrice = big.NewInt(30000000000)
	bridge, err := dosbridge.NewDosbridge(bridgeAddress, client)
	if err != nil {
		return
	}
	var tx *types.Transaction

	switch addrType {
	case proxyAddressType:
		tx, err = bridge.SetProxyAddress(auth, addr)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			tx, err = bridge.SetProxyAddress(auth, addr)
		}
		if err == nil {
			fmt.Println("tx sent: ", tx.Hash().Hex())
			fmt.Println("proxy address updated in bridge")
		}
	case crAddressType:
		tx, err = bridge.SetCommitRevealAddress(auth, addr)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			tx, err = bridge.SetCommitRevealAddress(auth, addr)
		}

		if err == nil {
			fmt.Println("tx sent: ", tx.Hash().Hex())
			fmt.Println("commitReveal address updated in bridge")
		}
	case paymentAddressType:
		tx, err = bridge.SetPaymentAddress(auth, addr)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			tx, err = bridge.SetPaymentAddress(auth, addr)
		}
		if err == nil {
			fmt.Println("tx sent: ", tx.Hash().Hex())
			fmt.Println("payment address updated in bridge")
		}
	default:
	}

	err = onchain.CheckTransaction(client, tx)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func addProxyToCRWhiteList(client *ethclient.Client, key *keystore.Key, bridgeAddress common.Address) (err error) {
	var auth *bind.TransactOpts
	fmt.Println("start to update proxy address to bridge...")
	auth = bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(6000000)
	auth.GasPrice = big.NewInt(30000000000)
	bridge, err := dosbridge.NewDosbridge(bridgeAddress, client)
	if err != nil {
		return
	}
	proxyAddress, err := bridge.GetProxyAddress(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return
	}
	crAddress, err := bridge.GetCommitRevealAddress(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		return

	}
	cr, err := commitreveal.NewCommitreveal(crAddress, client)
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

func deployContract(contractPath, targetTask, configPath string, config configuration.Config, client *ethclient.Client, key *keystore.Key) {
	var tx *types.Transaction
	var address common.Address
	auth := bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(6000000)
	auth.GasPrice = big.NewInt(30000000000)
	var err error
	if targetTask == "deployBridge" {
		fmt.Println("Starting deploy DOSAddressBridge.sol...")
		address, tx, _, err = dosbridge.DeployDosbridge(auth, client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = dosbridge.DeployDosbridge(auth, client)
		}
		fmt.Println("contract Address: ", address.Hex())
		fmt.Println("tx sent: ", tx.Hash().Hex())
		fmt.Println("contract deployed, waiting for confirmation...")
		err = onchain.CheckTransaction(client, tx)
		if err != nil {
			fmt.Println(err)
		}
		config.DOSAddressBridgeAddress = address.Hex()
		if err := config.UpdateConfig(); err != nil {
			log.Fatal(err)
		}
		if err := updateBridgeAddr(contractPath+"/DOSOnChainSDK.sol", address.Hex()); err != nil {
			log.Fatal(err)
		}
		if err := updateBridgeAddr(contractPath+"/DOSProxy.sol", address.Hex()); err != nil {
			log.Fatal(err)
		}
		return
	}

	if !common.IsHexAddress(config.DOSAddressBridgeAddress) {
		return
	}
	bridgeAddr := common.HexToAddress(config.DOSAddressBridgeAddress)

	if targetTask == "deployPayment" {
		fmt.Println("Starting deploy DOSPayment.sol...")
		address, tx, _, err = dospayment.DeployDospayment(auth, client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = dospayment.DeployDospayment(auth, client)
		}
		fmt.Println("contract Address: ", address.Hex())
		fmt.Println("tx sent: ", tx.Hash().Hex())
		fmt.Println("contract deployed, waiting for confirmation...")
		err = onchain.CheckTransaction(client, tx)
		if err != nil {
			fmt.Println(err)
		}
		err := updateBridge(client, key, paymentAddressType, bridgeAddr, address)
		if err != nil {
			log.Fatal(err)
		}

	} else if targetTask == "deployCommitReveal" {
		fmt.Println("Starting deploy Commitreveal.sol...")
		address, tx, _, err = commitreveal.DeployCommitreveal(auth, client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = commitreveal.DeployCommitreveal(auth, client)
		}
		fmt.Println("contract Address: ", address.Hex())
		fmt.Println("tx sent: ", tx.Hash().Hex())
		fmt.Println("contract deployed, waiting for confirmation...")
		err = onchain.CheckTransaction(client, tx)
		if err != nil {
			fmt.Println(err)
		}
		err := updateBridge(client, key, crAddressType, bridgeAddr, address)
		if err != nil {
			log.Fatal(err)
		}

	} else if targetTask == "deployProxy" {
		fmt.Println("Starting deploy DOSProxy.sol...")

		address, tx, _, err = dosproxy.DeployDosproxy(auth, client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = dosproxy.DeployDosproxy(auth, client)
		}
		fmt.Println("contract Address: ", address.Hex())
		fmt.Println("tx sent: ", tx.Hash().Hex())
		fmt.Println("contract deployed, waiting for confirmation...")
		err = onchain.CheckTransaction(client, tx)
		if err != nil {
			fmt.Println(err)
		}

		err := updateBridge(client, key, proxyAddressType, bridgeAddr, address)
		if err != nil {
			log.Fatal(err)
		}

		err = addProxyToCRWhiteList(client, key, bridgeAddr)
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

	}
}

func main() {
	credentialPathPtr := flag.String("credentialPath", "./testAccounts/fundKey/fundKey", "credential path")
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

	password := os.Getenv("PASSPHRASE")
	for password == "" {
		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(0)
		if err == nil {
			fmt.Println("\nPassword typed: ***")
		}
		password = strings.TrimSpace(string(bytePassword))
	}

	//Dial to blockchain
	key, err := onchain.ReadEthKey(credentialPath, password)
	if err != nil {
		fmt.Println("NewETHProxySession ", err)
		return
	}
	var c *ethclient.Client
	results := onchain.DialToEth(context.Background(), []string{"ws://51.15.0.157:8546"})
	for result := range results {
		if result.Err != nil {
			continue
		}
		c = result.Client
		break
	}
	if c == nil {
		fmt.Println("NewETHProxySession ", err)
		return
	}

	deployContract(contractPath, target, configPath, config, c, key)
}
