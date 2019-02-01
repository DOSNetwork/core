package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ENVCHAINTYPE      = "CHAINTYPE"
	ENVCHAINNODE      = "CHAINNODE"
	ENVRANDOMCONNECT  = "RANDOMCONNECT"
	ENVCONFIGPATH     = "CONFIGPATH"
	ENVNODEROLE       = "NODEROLE"
	ENVBOOTSTRAPIP    = "BOOTSTRAPIP"
	ENVNODEPORT       = "NODEPORT"
	CONFIGMODE        = "TESTMODE"
	ENVGROUPSIZE      = "GROUPSIZE"
	ENVCREDENTIALPATH = "CREDENTIALPATH"
)

type OffChainConfig struct {
	NodeRole        string
	BootStrapIp     string
	Port            int
	RandomGroupSize int
	QueryGroupSize  int
}

type OnChainConfig struct {
	ChainConfigs   map[string]map[string]ChainConfig
	credentialPath string
	path           string
	currentType    string
	currentNode    string
}

type ChainConfig struct {
	DOSProxyAddress         string
	DOSAddressBridgeAddress string
	RemoteNodeAddressPool   []string
	RemoteNodeAddress       string
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
		if workingDir == "/" {
			workingDir = "."
		}
	}
	err = LoadConfig(workingDir+"/offChain.json", c)
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

func (c *OnChainConfig) GetCredentialPath() string {
	return c.credentialPath
}

func (c *OnChainConfig) GetCurrentType() string {
	return c.currentType
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
		if workingDir == "/" {
			workingDir = "."
		}
	}
	c.path = workingDir + "/onChain.json"
	err = LoadConfig(c.path, c)
	if err != nil {
		fmt.Println("read config err ", err)
		return
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
		chainNode = "rinkebyPrivateNode"
	}
	c.currentNode = chainNode

	if config, loaded := c.ChainConfigs[c.currentType][c.currentNode]; loaded {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		if os.Getenv(ENVRANDOMCONNECT) == "true" {
			config.RemoteNodeAddress = config.RemoteNodeAddressPool[random.Intn(len(config.RemoteNodeAddressPool))]
		} else {
			config.RemoteNodeAddress = config.RemoteNodeAddressPool[0]
		}
		c.ChainConfigs[c.currentType][c.currentNode] = config
	}

	credentialPath := os.Getenv(ENVCREDENTIALPATH)
	if credentialPath == "" {
		fmt.Println("No ENVCREDENTIALPATH Environment variable.")
		credentialPath = "./credential"
	}
	c.credentialPath = credentialPath
	fmt.Println("c.credentialPath ", c.credentialPath, " ", credentialPath)

	return
}

func (c *OnChainConfig) GetChainConfig() (config ChainConfig) {
	return c.ChainConfigs[c.currentType][c.currentNode]
}

func (c *OnChainConfig) UpdateConfig(updated ChainConfig) (err error) {
	if _, loaded := c.ChainConfigs[c.currentType][c.currentNode]; loaded {
		c.ChainConfigs[c.currentType][c.currentNode] = updated
	}
	return UpdateConfig(c.path, c)
}
