package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/DOSNetwork/core/onchain"
)

const ENVCHAINTYPE = "CHAINTYPE"
const ENVCHAINNODE = "CHAINNODE"

type Config struct {
	NodeRole        string
	BootStrapIp     string
	Port            int
	RandomGroupSize int
	QueryGroupSize  int
	ChainConfigs    []onchain.ChainConfig
}

func ReadConfig(path string) (configs Config) {
	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened NetConfigs json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &configs)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func UpdateOnChainConfig(path string, config Config, updated onchain.ChainConfig) (err error) {
	chainType := os.Getenv(ENVCHAINTYPE)
	if chainType == "" {
		fmt.Println("No CHAINTYPE Environment variable.")
		chainType = "ETH"
	}
	chainNode := os.Getenv(ENVCHAINNODE)
	if chainNode == "" {
		fmt.Println("No CHAINNODE Environment variable.")
		chainNode = "rinkebyPrivateNode"
	}

	for i, c := range config.ChainConfigs {
		if chainNode == c.RemoteNodeType &&
			chainType == c.ChainType {
			config.ChainConfigs[i] = updated
		}
	}
	configsJson, _ := json.Marshal(config)
	err = ioutil.WriteFile(path, configsJson, 0644)
	return nil
}

func GetOnChainConfig(config Config) (node onchain.ChainConfig) {
	chainType := os.Getenv(ENVCHAINTYPE)
	if chainType == "" {
		fmt.Println("No CHAINTYPE Environment variable.")
		chainType = "ETH"
	}
	chainNode := os.Getenv(ENVCHAINNODE)
	if chainNode == "" {
		fmt.Println("No CHAINNODE Environment variable.")
		chainNode = "rinkebyPrivateNode"
	}

	for _, c := range config.ChainConfigs {
		if chainNode == c.RemoteNodeType &&
			chainType == c.ChainType {
			fmt.Println("Use : ", config)
			return c
		}
	}
	return onchain.ChainConfig{}
}
