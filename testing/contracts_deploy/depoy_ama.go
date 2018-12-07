package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/testing/dosUser/contract"
	"github.com/DOSNetwork/core/testing/dosUser/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

func deployAMAContract(ethComm *onchain.EthCommon) string {
	var tx *types.Transaction
	var address common.Address
	auth, err := ethComm.GetAuth()

	fmt.Println("Starting deploy AskMeAnyThing.sol...")
	address, tx, _, err = dosUser.DeployAskMeAnything(auth, ethComm.Client)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		fmt.Println(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		address, tx, _, err = dosUser.DeployAskMeAnything(auth, ethComm.Client)
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

func main() {
	credentialPathFlag := flag.String("credentialPath", "./testAccounts/bootCredential/", "credentialPath")
	amaConfigPathFlag := flag.String("AMAPath", "ama.json", "ama.json path")
	flag.Parse()
	credentialPath := *credentialPathFlag
	amaConfigPath := *amaConfigPathFlag

	os.Setenv("CONFIGPATH", "../")

	config := configuration.OnChainConfig{}
	config.LoadConfig()
	chainConfig := config.GetChainConfig()
	conn := &onchain.EthCommon{}
	_ = conn.Init(credentialPath, &chainConfig)
	amaConfig := eth.AMAConfig{}
	fmt.Println("Starting deploy AskMeAnyThing.sol...")
	amaConfig.AskMeAnythingAddress = deployAMAContract(conn)
	configuration.UpdateConfig(amaConfigPath, amaConfig)
}
