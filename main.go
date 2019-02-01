package main

import (
	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/dosnode"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/suites"
)

// main
func main() {
	//Read Configuration
	offChainConfig := configuration.OffChainConfig{}
	if err := offChainConfig.LoadConfig(); err != nil {
		log.Fatal(err)
	}
	role := offChainConfig.NodeRole
	port := offChainConfig.Port
	bootstrapIp := offChainConfig.BootStrapIp

	onChainConfig := configuration.OnChainConfig{}
	if err := onChainConfig.LoadConfig(); err != nil {
		log.Fatal(err)
	}
	chainConfig := onChainConfig.GetChainConfig()

	//Set up an onchain adapter
	chainConn, err := onchain.AdaptTo(onChainConfig.GetCurrentType())
	if err != nil {
		log.Fatal(err)
	}
	chainConn.SetAccount(onChainConfig.GetCredentialPath())
	//Init log module with nodeID that is an onchain account address
	log.Init(chainConn.GetId()[:])
	chainConn.Init(chainConfig)

	//Build a p2p network
	p, err := p2p.CreateP2PNetwork(chainConn.GetId(), port)
	if err != nil {
		log.Fatal(err)
	}
	if err := p.Listen(); err != nil {
		log.Fatal(err)
	}

	//Bootstrapping p2p network
	if role != "BootstrapNode" {
		if err = p.Join(bootstrapIp); err != nil {
			log.Fatal(err)
		}
	}

	//Build a p2pDKG
	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(p, suite)
	if err != nil {
		log.Fatal(err)
	}
	dosclient := dosnode.NewDosNode(suite, p, chainConn, p2pDkg)
	dosclient.Start()
	done := make(chan interface{})
	<-done
}
