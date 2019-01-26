package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/DOSNetwork/core/testing/dosUser/contract"
	"github.com/DOSNetwork/core/testing/dosUser/eth"
)

const (
	DOSAddressBridge = iota
	DOSProxy
	AMA
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
	path = path + "/DOSOnChainSDK.sol"
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

func updateBridge(ethComm *onchain.EthCommon, bridgeAddress, proxyAddress common.Address) (err error) {
	var auth *bind.TransactOpts
	fmt.Println("start to update proxy address to bridge...")
	auth, err = ethComm.GetAuth()
	if err != nil {
		return
	}
	bridge, err := dosproxy.NewDOSAddressBridge(bridgeAddress, ethComm.Client)
	if err != nil {
		return
	}

	tx, err := bridge.SetProxyAddress(auth, proxyAddress)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = bridge.SetProxyAddress(auth, proxyAddress)
	}
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("proxy address updated in bridge")

	err = ethComm.CheckTransaction(tx)

	return
}

func deployContract(ethComm *onchain.EthCommon, contractName int) string {
	var tx *types.Transaction
	var address common.Address
	auth, err := ethComm.GetAuth()

	switch contractName {
	case DOSAddressBridge:
		fmt.Println("Starting deploy DOSAddressBridge.sol...")
		address, tx, _, err = dosproxy.DeployDOSAddressBridge(auth, ethComm.Client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = dosproxy.DeployDOSAddressBridge(auth, ethComm.Client)
		}
	case DOSProxy:
		fmt.Println("Starting deploy DOSProxy.sol...")
		address, tx, _, err = dosproxy.DeployDOSProxy(auth, ethComm.Client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = dosproxy.DeployDOSProxy(auth, ethComm.Client)
		}
	case AMA:
		fmt.Println("Starting deploy AskMeAnyThing.sol...")
		address, tx, _, err = dosUser.DeployAskMeAnything(auth, ethComm.Client)
		for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
			fmt.Println(err)
			time.Sleep(time.Second)
			fmt.Println("transaction retry...")
			address, tx, _, err = dosUser.DeployAskMeAnything(auth, ethComm.Client)
		}
	}
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println("contract Address: ", address.Hex())
	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("contract deployed, waiting for confirmation...")

	err = ethComm.CheckTransaction(tx)
	if err != nil {
		fmt.Println(err)
	}
	return address.Hex()
}

func main() {
	credentialPathPtr := flag.String("credentialPath", "./testAccounts/bootCredential/", "credential path")
	contractPathPtr := flag.String("contractPath", "./contracts", "Contract file path")
	contractPtr := flag.String("contract", "DOSProxy", "DOSProxy or AMA or SimpleDice")
	configPathPtr := flag.String("configPath", "", "config path")

	flag.Parse()
	credentialPath := *credentialPathPtr
	contractPath := *contractPathPtr
	contract := *contractPtr
	configPath := *configPathPtr

	os.Setenv("CONFIGPATH", "..")
	os.Setenv("CREDENTIALPATH", credentialPath)
	//Get
	onChainConfig := configuration.OnChainConfig{}
	onChainConfig.LoadConfig()
	chainConfig := onChainConfig.GetChainConfig()
	fmt.Println(" address ", chainConfig.RemoteNodeAddress)

	//Dial to blockchain
	conn := &onchain.EthCommon{}
	conn.SetAccount(onChainConfig.CredentialPath)

	_ = conn.Init(chainConfig)

	if contract == "DOSProxy" {
		chainConfig.DOSProxyAddress = deployContract(conn, DOSProxy)
		chainConfig.DOSAddressBridgeAddress = deployContract(conn, DOSAddressBridge)
		onChainConfig.UpdateConfig(chainConfig)

		updateSDK(contractPath, chainConfig.DOSAddressBridgeAddress)

		bridgeAddress := common.HexToAddress(chainConfig.DOSAddressBridgeAddress)
		proxyAddress := common.HexToAddress(chainConfig.DOSProxyAddress)
		err := updateBridge(conn, bridgeAddress, proxyAddress)
		if err != nil {
			log.Fatal(err)
		}
	} else if contract == "AMA" {
		amaConfig := eth.AMAConfig{}
		amaConfig.AskMeAnythingAddress = deployContract(conn, AMA)
		configuration.UpdateConfig(configPath, amaConfig)
	}
}
