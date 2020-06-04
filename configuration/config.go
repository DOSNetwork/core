package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

const (
	envChainType     = "CHAINTYPE"
	envConfigPath    = "CONFIGPATH"
	envBootStrapIP   = "BOOTSTRAPIP"
	envNodeIP        = "NODEIP"
	envNodePort      = "NODEPORT"
	envGethPool      = "GETHPOOL"
	defaultChainType = "ETH"
	defaultPort      = "9501"
)

var path string

// Config is the configuration for creating a DOS client instance.
type Config struct {
	VERSION                 string
	NodeIP                  string
	NodePort                string
	ChainType               string
	DOSAddressBridgeAddress string
	BootStrapIPs            []string
	ChainNodePool           []string
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

// LoadConfig loads configuration file from path.
func (c *Config) LoadConfig() (err error) {
	path = os.Getenv(envConfigPath)
	if path == "" {
		path, err = os.Getwd()
		if err != nil {
			return
		}
		if path == "/" {
			path = "."
		}
	}
	path = path + "/config.json"

	err = LoadConfig(path, c)
	if err != nil {
		fmt.Println("LoadConfig  err", err)
		return
	}
	err = c.overWrite()

	return
}

func (c *Config) overWrite() (err error) {
	ip := os.Getenv(envNodeIP)
	if ip != "" {
		testInput := net.ParseIP(ip)
		if testInput.To4() == nil {
			fmt.Printf("%v is not a valid IPv4 address\n", testInput)
			return errors.New("not a valid IPv4 address")
		} else {
			c.NodeIP = ip
		}
	}
	port := os.Getenv(envNodePort)
	if port != "" {
		c.NodePort = port
	} else if c.NodePort == "" {
		c.NodePort = defaultPort
	}

	chainType := os.Getenv(envChainType)
	if chainType != "" {
		c.ChainType = chainType
	} else if c.ChainType == "" {
		c.ChainType = defaultChainType
	}

	gethIP := os.Getenv(envGethPool)
	if gethIP != "" {
		nodeList := make(map[string]bool)
		for _, ethNode := range c.ChainNodePool {
			nodeList[ethNode] = true
		}
		ipPool := strings.Split(gethIP, ",")
		for _, ip := range ipPool {
			if !nodeList[ip] {
				c.ChainNodePool = append(c.ChainNodePool, ip)
			}
		}
	}
	return
}

// UpdateConfig saves configuration to a file.
func (c *Config) UpdateConfig() (err error) {
	configJson, _ := json.MarshalIndent(c, "", "    ")
	err = ioutil.WriteFile(path, configJson, 0644)
	return
}
