package main

import (
	"bytes"
	"fmt"
	"math/big"
	_ "net/http/pprof"
	"runtime/debug"

	"github.com/bshuster-repo/logrus-logstash-hook"

	"github.com/ethereum/go-ethereum/common"

	"github.com/sirupsen/logrus"

	"github.com/DOSNetwork/core/configuration"
	dos "github.com/DOSNetwork/core/dosnode"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
)

var log *logrus.Logger

// main
func main() {
	offChainConfig := configuration.OffChainConfig{}
	offChainConfig.LoadConfig()
	role := offChainConfig.NodeRole
	nbParticipants := offChainConfig.RandomGroupSize
	port := offChainConfig.Port
	bootstrapIp := offChainConfig.BootStrapIp

	onChainConfig := configuration.OnChainConfig{}
	onChainConfig.LoadConfig()
	chainConfig := onChainConfig.GetChainConfig()

	//0)initial log module
	log = logrus.New()

	//1)Connect to Eth
	chainConn, err := onchain.AdaptTo(chainConfig.ChainType, &chainConfig, log)
	if err != nil {
		log.Fatal(err)
	}

	//2)Build a p2p network
	peerEvent := make(chan p2p.P2PMessage, 100)
	defer close(peerEvent)
	p, _ := p2p.CreateP2PNetwork(peerEvent, port)
	p.SetId(chainConn.GetId())
	p.Listen()

	//2.5)Add hook to log module
	hook, err := logrustash.NewHookWithFields("udp", "13.52.16.14:9500", "DOS_node", logrus.Fields{
		"DOS_node_ip": p.GetId().Address,
		"Serial":      string(common.BytesToAddress(p.GetId().Id).String()),
	})
	if err != nil {
		log.Error(err)
	}

	log.Hooks.Add(hook)

	//3)Dial to peers to build peerClient and Set node ID
	if role == "BootstrapNode" {
		address, err := chainConn.GetWhitelist()
		if err != nil {
			log.Fatal(err)
		}

		if bytes.Compare(address.Bytes(), chainConn.GetId()) != 0 {
			err = chainConn.InitialWhiteList()
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Println(bootstrapIp)
		p.CreatePeer(bootstrapIp, nil)
		results := p.FindNode(p.GetId(), dht.BucketSize, 20)
		for _, result := range results {
			p.GetRoutingTable().Update(result)
			fmt.Println(p.GetId().Address, "Update peer: ", result.Address)
		}
	}

	//4)Build a p2pDKG
	suite := suites.MustFind("bn256")
	peerEventForDKG := make(chan p2p.P2PMessage, 1)
	defer close(peerEventForDKG)
	p2pDkg, _ := dkg.CreateP2PDkg(p, suite, peerEventForDKG, nbParticipants, log)
	go p2pDkg.EventLoop()
	dkgEvent := make(chan string, 1)
	p2pDkg.SubscribeEvent(dkgEvent)
	defer close(dkgEvent)

	//5)Subscribe Event from Eth
	eventGrouping := make(chan interface{}, 1)
	defer close(eventGrouping)
	chUrl := make(chan interface{}, 20)
	defer close(chUrl)
	chRandom := make(chan interface{}, 20)
	defer close(chRandom)
	chUsrRandom := make(chan interface{}, 20)
	defer close(chUsrRandom)
	cSignatureFromPeer := make(chan vss.Signature, 100)
	defer close(cSignatureFromPeer)
	eventValidation := make(chan interface{}, 20)
	defer close(eventValidation)
	err = chainConn.SubscribeEvent(eventGrouping, onchain.SubscribeDOSProxyLogGrouping)
	if err != nil {
		log.Fatal(err)
	}

	err = chainConn.SubscribeEvent(chUrl, onchain.SubscribeDOSProxyLogUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = chainConn.SubscribeEvent(chRandom, onchain.SubscribeDOSProxyLogUpdateRandom)
	if err != nil {
		log.Fatal(err)
	}

	err = chainConn.SubscribeEvent(chUsrRandom, onchain.SubscribeDOSProxyLogRequestUserRandom)
	if err != nil {
		log.Fatal(err)
	}

	err = chainConn.SubscribeEvent(eventValidation, onchain.SubscribeDOSProxyLogValidationResult)
	if err != nil {
		log.Fatal(err)
	}

	//6)Set up a dosnode pipeline
	d := dos.CreateDosNode(suite, nbParticipants, p, chainConn, p2pDkg, log)
	d.PipeGrouping(eventGrouping)
	queriesReports := d.PipeQueries(chUrl, chUsrRandom, chRandom)
	signedReports := d.PipeSignAndBroadcast(queriesReports)
	reportsToSubmit, reportToValidate := d.PipeRecoverAndVerify(cSignatureFromPeer, signedReports)
	d.PipeSendToOnchain(reportsToSubmit, reportToValidate)
	chRetrigerUrl := d.PipeCleanFinishMap(eventValidation, reportToValidate)

	err = chainConn.UploadID()
	if err != nil {
		log.Fatal(err)
	}
	//TODO:/crypto/scrypt: memory not released after hash is calculated
	//https://github.com/golang/go/issues/20000
	//This caused dosnode take over 500MB memory usage.
	//Need to check to see if there is other way to trigger GC
	debug.FreeOSMemory()
	//7)Dispatch events
	for {
		select {
		//event from peer
		case msg := <-peerEvent:
			switch content := msg.Msg.Message.(type) {
			case *vss.PublicKey:
				peerEventForDKG <- msg
			case *dkg.Deal:
				peerEventForDKG <- msg
			case *dkg.Response:
				peerEventForDKG <- msg
			case *vss.Signature:
				cSignatureFromPeer <- *content
			default:
				fmt.Println("unknown", content)
			}
		case msg := <-dkgEvent:
			if msg == "certified" {
				gId := new(big.Int)
				gId.SetBytes(p2pDkg.GetGroupId())
				chainConn.UploadPubKey(p2pDkg.GetGroupPublicPoly().Commit())
			}
		//For re-trigger query
		case msg := <-chRetrigerUrl:
			chUrl <- msg
		}
	}
}
