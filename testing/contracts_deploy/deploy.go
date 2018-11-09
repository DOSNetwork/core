package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/onchain/dosproxy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	DOSAddressBridge = iota
	DOSProxy
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

func deployContract(ethComm *onchain.EthCommon, contractName int) string {
	var tx *types.Transaction
	var address common.Address
	auth, err := ethComm.GetAuth()

	switch contractName {
	case DOSAddressBridge:
		fmt.Println("Starting deploy DOSAddressBridge.sol...")
		address, tx, _, err = dosproxy.DeployDOSAddressBridge(auth, ethComm.Client)
	case DOSProxy:
		fmt.Println("Starting deploy DOSProxy.sol...")
		address, tx, _, err = dosproxy.DeployDOSProxy(auth, ethComm.Client)
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
		//		return ""
	}
	return address.Hex()
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
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("proxy address updated in bridge")

	err = ethComm.CheckTransaction(tx)

	return
}

func main() {
	credentialPathFlag := flag.String("credentialPath", "./testAccounts/bootCredential/", "credentialPath")
	contractPathFlag := flag.String("contractPath", "./contracts", "Contract file path")
	stepFlag := flag.String("step", "ProxyAndBridge", "ProxyAndBridge or SDKAndAMA")
	flag.Parse()
	credentialPath := *credentialPathFlag
	contractPath := *contractPathFlag
	contractPath = contractPath + "/DOSOnChainSDK.sol"
	step := *stepFlag

	os.Setenv("CONFIGPATH", "../")

	config := configuration.OnChainConfig{}
	config.LoadConfig()
	chainConfig := config.GetChainConfig()
	fmt.Println(" address ", chainConfig.RemoteNodeAddress)

	conn := &onchain.EthCommon{}
	_ = conn.Init(credentialPath, &chainConfig)


	if step == "ProxyAndBridge" {
		chainConfig.DOSProxyAddress = deployContract(conn, DOSProxy)
		chainConfig.DOSAddressBridgeAddress = deployContract(conn, DOSAddressBridge)
		config.UpdateConfig(chainConfig)
		updateSDK(contractPath, chainConfig.DOSAddressBridgeAddress)
	} else if step == "SetProxyAddress" {
		bridgeAddress := common.HexToAddress(chainConfig.DOSAddressBridgeAddress)
		proxyAddress := common.HexToAddress(chainConfig.DOSProxyAddress)
		err := updateBridge(conn, bridgeAddress, proxyAddress)
		if err != nil {
			log.Fatal(err)
		}
	}
}
