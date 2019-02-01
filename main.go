package main

import (
	"bytes"
	"runtime/debug"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/dosnode"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
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

	if err != nil {
		log.Fatal(err)
	}

	//Dial to peers to build peerClient and Set node ID
	if role == "BootstrapNode" {
		address, err := chainConn.GetWhitelist()
		if err != nil {
			log.Fatal(err)
		}

		if bytes.Compare(address.Bytes(), chainConn.GetId()) != 0 {
			if err = chainConn.InitialWhiteList(); err != nil {
				log.Fatal(err)
			}
		}
	} else {
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

	//Subscribe Event from Eth
	eventGrouping := make(chan interface{}, 1)
	defer close(eventGrouping)
	chUrl := make(chan interface{}, 100)
	defer close(chUrl)
	chRandom := make(chan interface{}, 100)
	defer close(chRandom)
	chUsrRandom := make(chan interface{}, 100)
	defer close(chUsrRandom)
	cSignatureFromPeer, _ := p.SubscribeEvent(100, vss.Signature{})
	defer p.UnSubscribeEvent(vss.Signature{})
	eventValidation := make(chan interface{}, 20)
	defer close(eventValidation)
	if err = chainConn.SubscribeEvent(eventGrouping, onchain.SubscribeDOSProxyLogGrouping); err != nil {
		log.Fatal(err)
	}

	if err = chainConn.SubscribeEvent(chUrl, onchain.SubscribeDOSProxyLogUrl); err != nil {
		log.Fatal(err)
	}

	if err = chainConn.SubscribeEvent(chRandom, onchain.SubscribeDOSProxyLogUpdateRandom); err != nil {
		log.Fatal(err)
	}

	if err = chainConn.SubscribeEvent(chUsrRandom, onchain.SubscribeDOSProxyLogRequestUserRandom); err != nil {
		log.Fatal(err)
	}

	if err = chainConn.SubscribeEvent(eventValidation, onchain.SubscribeDOSProxyLogValidationResult); err != nil {
		log.Fatal(err)
	}

	//Set up a dosnode pipeline
	d := dosnode.CreateDosNode(suite, p, chainConn, p2pDkg)
	d.PipeGrouping(eventGrouping)
	queriesReports := d.PipeQueries(chUrl, chUsrRandom, chRandom)
	outForRecover, outForValidate := d.PipeSignAndBroadcast(queriesReports)
	reportsToSubmit := d.PipeRecoverAndVerify(cSignatureFromPeer, outForRecover)
	submitForValidate := d.PipeSendToOnchain(reportsToSubmit)
	chRetrigerUrl := d.PipeCleanFinishMap(eventValidation, outForValidate, submitForValidate)

	if err = chainConn.UploadID(); err != nil {
		log.Fatal(err)
	}
	//TODO:/crypto/scrypt: memory not released after hash is calculated
	//https://github.com/golang/go/issues/20000
	//This caused dosnode take over 500MB memory usage.
	//Need to check to see if there is other way to trigger GC
	debug.FreeOSMemory()

	//Dispatch events
	for {
		select {
		//For re-trigger query
		case msg := <-chRetrigerUrl:
			chUrl <- msg
		}
	}
}
