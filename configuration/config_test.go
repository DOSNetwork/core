package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	config := OffChainConfig{}
	LoadConfig("../offChain.json", &config)
	require.Equal(t, "BootstrapNode", config.NodeRole)

	config2 := OnChainConfig{}
	LoadConfig("../onChain.json", &config2)
	require.Equal(t, "ETH", config2.ChainConfigs[0].ChainType)
}

func TestOffChainReadConfig(t *testing.T) {
	config := OffChainConfig{}
	os.Setenv("CONFIGPATH", "..")
	config.LoadConfig()
	assert.Equal(t, "BootstrapNode", config.NodeRole)
}

func TestOnChainReadConfig(t *testing.T) {
	os.Setenv("CONFIGPATH", "..")
	os.Setenv("CHAINNODE", "rinkebyPrivateNode")
	config := OnChainConfig{}
	config.LoadConfig()

	actualResult := config.GetChainConfig().RemoteNodeType
	expectedResult := "rinkebyPrivateNode"

	require.Equal(t, expectedResult, actualResult)
}

func TestOnChainUpdate(t *testing.T) {
	config := OnChainConfig{}
	os.Setenv("CONFIGPATH", "..")
	os.Setenv("CHAINNODE", "rinkebyPrivateNode")
	config.LoadConfig()
	chainConfig := config.GetChainConfig()
	chainConfig.DOSAddressBridgeAddress = "0x12345"
	config.UpdateConfig(chainConfig)

	config.LoadConfig()
	chainConfig = config.GetChainConfig()
	actualResult := config.GetChainConfig().DOSAddressBridgeAddress
	expectedResult := "0x12345"

	require.Equal(t, expectedResult, actualResult)
}
