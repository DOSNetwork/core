package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	ENVCHAINTYPE     = "CHAINTYPE"
	ENVCHAINNODE     = "CHAINNODE"
	ENVRANDOMCONNECT = "RANDOMCONNECT"
	ENVCONFIGPATH    = "CONFIGPATH"
	ENVNODEROLE      = "NODEROLE"
	ENVBOOTSTRAPIP   = "BOOTSTRAPIP"
	ENVNODEPORT      = "NODEPORT"
	ENVGROUPSIZE     = "GROUPSIZE"
	ENVPASSPHRASE    = "PASSPHRASE"
	ENVGROUPTOPICK   = "GROUPTOPICK"
)

// Config is the configuration for creating a DOS client instance.
type Config struct {
	NodeRole        string
	BootStrapIp     []string
	Port            string
	ChainConfigs    map[string]map[string]ChainConfig
	randomGroupSize int
	queryGroupSize  int
	credentialPath  string
	path            string
	currentType     string
	currentNode     string
}

// ChainConfig is the configuration for connecting to onchan contracts.
type ChainConfig struct {
	DOSProxyAddress         string
	DOSPaymentAddress       string
	DOSAddressBridgeAddress string
	CommitReveal            string
	RemoteNodeAddressPool   []string
}

// LoadConfig loads configuration file from path.
func LoadConfig(path string, c interface{}) (err error) {
	var jsonFile *os.File
	var byteValue []byte

	fmt.Println("Path ", path)
	// Open our jsonFile
	jsonFile, err = os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("LoadConfig error ", err)
		return
	}
	fmt.Println("Successfully Opened json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("ReadAll error ", err)
		return
	}

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		fmt.Println("Unmarshal error ", err)
		return
	}
	return
}

// UpdateConfig saves configuration to a file.
func UpdateConfig(path string, c interface{}) (err error) {
	configsJson, _ := json.MarshalIndent(c, "", "    ")
	err = ioutil.WriteFile(path, configsJson, 0644)
	return
}

// LoadConfig loads configuration file from path.
func (c *Config) LoadConfig() (err error) {
	path := os.Getenv(ENVCONFIGPATH)
	if path == "" {
		path, err = os.Getwd()
		if err != nil {
			return
		}
		if path == "/" {
			path = "."
		}
	}
	c.path = path + "/config.json"
	err = LoadConfig(c.path, c)
	if err != nil {
		fmt.Println("LoadConfig  err", err)
		return
	}

	err = c.overWrite()

	return
}

func (c *Config) overWrite() (err error) {
	envSize := os.Getenv(ENVGROUPSIZE)
	if envSize != "" {
		var size int
		size, err = strconv.Atoi(envSize)
		if err != nil {
			return
		}
		c.randomGroupSize = size
		c.queryGroupSize = size
	}

	nodeRole := os.Getenv(ENVNODEROLE)
	if nodeRole != "" {
		c.NodeRole = nodeRole
	}

	port := os.Getenv(ENVNODEPORT)
	if port != "" {
		//TODO:add a check
		c.Port = port
	}

	chainType := os.Getenv(ENVCHAINTYPE)
	if chainType == "" {
		fmt.Println("No CHAINTYPE Environment variable.")
		chainType = "ETH"
	}
	c.currentType = chainType

	chainNode := os.Getenv(ENVCHAINNODE)
	if chainNode == "" {
		fmt.Println("No CHAINNODE Environment variable.")
		chainNode = "rinkeby"
	}
	c.currentNode = chainNode
	if config, loaded := c.ChainConfigs[c.currentType][c.currentNode]; loaded {
		gethIP := os.Getenv("GETHPOOL")
		if gethIP != "" {
			ipPool := strings.Split(gethIP, ";")
			for _, ip := range ipPool {
				config.RemoteNodeAddressPool = append(config.RemoteNodeAddressPool, ip)
			}
		}
		c.ChainConfigs[c.currentType][c.currentNode] = config
	}

	return
}

// GetCurrentType return a string to indicate the type of onchain.
func (c *Config) GetCurrentType() string {
	return c.currentType
}

// GetChainConfig return a ChainConfig struct that contain onchain information
func (c *Config) GetChainConfig() (config ChainConfig) {
	return c.ChainConfigs[c.currentType][c.currentNode]
}

// UpdateConfig saves configuration to a file.
func (c *Config) UpdateConfig(updated ChainConfig) (err error) {
	if _, loaded := c.ChainConfigs[c.currentType][c.currentNode]; loaded {
		c.ChainConfigs[c.currentType][c.currentNode] = updated
	}
	return UpdateConfig(c.path, c)
}
