package dosproxy

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
)

type DOSAddressBridgeTestSuite struct {
	suite.Suite
	auth    *bind.TransactOpts
	address common.Address
	gAlloc  core.GenesisAlloc
	sim     *backends.SimulatedBackend
	bridge  *DOSAddressBridge
}

func TestRunDOSAddressBridgeSuite(t *testing.T) {
	suite.Run(t, new(DOSAddressBridgeTestSuite))
}

func (s *DOSAddressBridgeTestSuite) SetupTest() {
	key, _ := crypto.GenerateKey()
	s.auth = bind.NewKeyedTransactor(key)

	s.address = s.auth.From
	s.gAlloc = map[common.Address]core.GenesisAccount{
		s.address: {Balance: big.NewInt(10000000000)},
	}

	s.sim = backends.NewSimulatedBackend(s.gAlloc, uint64(8000000))

	_, _, contract, e := DeployDOSAddressBridge(s.auth, s.sim)
	s.bridge = contract
	s.Nil(e)
	s.sim.Commit()

}

func (s *DOSAddressBridgeTestSuite) TestProxyAddress() {
	var z1 big.Int
	z1.SetUint64(123) // z1 := 123
	z1.SetString("Ffb6A23dE33eB7efE126A844882F40411a02A21b", 16)
	toAddr := common.BigToAddress(&z1)
	_, err := s.bridge.SetProxyAddress(&bind.TransactOpts{
		From:   s.auth.From,
		Signer: s.auth.Signer,
		Value:  nil,
	}, toAddr)
	s.Nil(err)
	s.sim.Commit()
	got, _ := s.bridge.GetProxyAddress(nil)
	s.Exactly(got.String(), toAddr.String())
}
