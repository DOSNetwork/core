package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const ENVCHAINTYPE = "CHAINTYPE"
const ENVCHAINNODE = "CHAINNODE"
const ENVCONFIGPATH = "CONFIGPATH"
const ENVNODEROLE = "NODEROLE"
const ENVBOOTSTRAPIP = "BOOTSTRAPIP"
const ENVNODEPORT = "NODEPORT"
const CONFIGMODE = "TESTMODE"
const ENVGROUPSIZE = "GROUPSIZE"
const ENVCREDENTIALPATH = "CREDENTIALPATH"

type OffChainConfig struct {
	NodeRole        string
	BootStrapIp     string
	Port            int
	RandomGroupSize int
	QueryGroupSize  int
}

type OnChainConfig struct {
	ChainConfigs   []ChainConfig
	CredentialPath string
	path           string
	currentType    string
	currentNode    string
}

type ChainConfig struct {
	// TODO: Refactor out of ChainConfig
	RemoteNodeType    string
	RemoteNodeAddress string

	ChainType               string
	DOSProxyAddress         string
	DOSAddressBridgeAddress string
}

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

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		fmt.Println("Unmarshal error ", err)
		return
	}
	return
}

func UpdateConfig(path string, c interface{}) (err error) {
	configsJson, _ := json.MarshalIndent(c, "", "    ")
	err = ioutil.WriteFile(path, configsJson, 0644)
	return
}

func (c *OffChainConfig) LoadConfig() (err error) {
	var workingDir string
	path := os.Getenv(ENVCONFIGPATH)
	if path != "" {
		workingDir = path
	} else {
		workingDir, err = os.Getwd()
		if err != nil {
			return
		}
	}
	_ = workingDir
	err = LoadConfig("./offChain.json", c)
	if err != nil {
		fmt.Println("LoadConfig  err", err)
		return
	}

	err = c.overWrite()
	//---------------------------------
	if c.NodeRole == "testNode" {
		var credential []byte
		var resp *http.Response
		fmt.Println("This is a test node : ", c.BootStrapIp)
		s := strings.Split(c.BootStrapIp, ":")
		ip, _ := s[0], s[1]
		tServer := "http://" + ip + ":8080/getCredential"
		resp, err = http.Get(tServer)
		for err != nil {
			time.Sleep(1 * time.Second)
			resp, err = http.Get(tServer)
		}

		credential, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		fmt.Println("credential : ", string(credential))
		credentialPath := workingDir + "/credential/usrKey"
		err = ioutil.WriteFile(credentialPath, []byte(credential), 0644)
		resp.Body.Close()
		if err != nil {
			return
		}
	}
	//---------------------------------
	return
}

func (c *OffChainConfig) overWrite() (err error) {
	bootStrapIP := os.Getenv(ENVBOOTSTRAPIP)
	if bootStrapIP != "" {
		c.BootStrapIp = bootStrapIP
	}

	envSize := os.Getenv(ENVGROUPSIZE)
	if envSize != "" {
		var size int
		size, err = strconv.Atoi(envSize)
		if err != nil {
			return
		}
		c.RandomGroupSize = size
		c.QueryGroupSize = size
	}

	nodeRole := os.Getenv(ENVNODEROLE)
	if nodeRole != "" {
		c.NodeRole = nodeRole
	}

	port := os.Getenv(ENVNODEPORT)
	if port != "" {
		i, err := strconv.Atoi(port)
		if err == nil {
			c.Port = i
		}
	}
	return
}

func (c *OnChainConfig) LoadConfig() (err error) {
	var workingDir string

	path := os.Getenv(ENVCONFIGPATH)
	if path != "" {
		workingDir = path
	} else {
		workingDir, err = os.Getwd()
		if err != nil {
			return
		}
	}
	_ = workingDir
	c.path = "./onChain.json"
	chainType := os.Getenv(ENVCHAINTYPE)
	if chainType == "" {
		fmt.Println("No CHAINTYPE Environment variable.")
		chainType = "ETH"
	}
	c.currentType = chainType

	chainNode := os.Getenv(ENVCHAINNODE)
	if chainNode == "" {
		fmt.Println("No CHAINNODE Environment variable.")
		chainNode = "rinkebyPrivateNode"
	}
	c.currentNode = chainNode

	credentialPath := os.Getenv(ENVCREDENTIALPATH)
	if chainNode == "" {
		fmt.Println("No ENVCREDENTIALPATH Environment variable.")
		credentialPath = "."
	}
	c.CredentialPath = credentialPath

	err = LoadConfig(c.path, c)
	if err != nil {
		fmt.Println("read config err ", err)
		return
	}
	return
}

func (c *OnChainConfig) GetChainConfig() (config ChainConfig) {
	for _, config := range c.ChainConfigs {
		if c.currentNode == config.RemoteNodeType &&
			c.currentType == config.ChainType {
			fmt.Println("Use : ", config)
			return config
		}
	}
	return
}

func (c *OnChainConfig) UpdateConfig(updated ChainConfig) (err error) {
	for i, config := range c.ChainConfigs {
		if c.currentNode == config.RemoteNodeType &&
			c.currentType == config.ChainType {
			c.ChainConfigs[i] = updated
		}
	}
	configsJson, _ := json.MarshalIndent(c, "", "    ")
	err = ioutil.WriteFile(c.path, configsJson, 0644)
	return nil
}
