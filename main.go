package main

import (
	"bytes"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"math/big"
	_ "net/http/pprof"
	"runtime/debug"

	"github.com/ethereum/go-ethereum/common"

	"github.com/sirupsen/logrus"

	"github.com/DOSNetwork/core/configuration"
	dos "github.com/DOSNetwork/core/dosnode"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
)

var log *logrus.Logger

// main
func main() {
	//0)initial log module
	log = logrus.New()
	mainLogger := log.WithField("module", "main")

	//0.5)Read Config
	offChainConfig := configuration.OffChainConfig{}
	if err := offChainConfig.LoadConfig(); err != nil {
		mainLogger.WithField("function", "loadConfig").Fatal(err)
	}
	role := offChainConfig.NodeRole
	nbParticipants := offChainConfig.RandomGroupSize
	port := offChainConfig.Port
	bootstrapIp := offChainConfig.BootStrapIp

	onChainConfig := configuration.OnChainConfig{}
	if err := onChainConfig.LoadConfig(); err != nil {
		mainLogger.WithField("function", "loadConfig").Fatal(err)
	}
	chainConfig := onChainConfig.GetChainConfig()

	//1)Connect to Eth
	chainConn, err := onchain.AdaptTo(chainConfig.ChainType, &chainConfig, log.WithField("module", "chainConn"))
	if err != nil {
		mainLogger.WithField("function", "adaptTo").Fatal(err)
	}

	//2)Build a p2p network
	p, peerEvent, err := p2p.CreateP2PNetwork(chainConn.GetId(), port)
	if err != nil {
		mainLogger.WithField("function", "createP2PNetwork").Fatal(err)
	}
	if err := p.Listen(); err != nil {
		mainLogger.WithField("function", "listen").Error(err)
	}

	//2.5)Add hook to log module
	hook, err := logrustash.NewHookWithFields("tcp", "163.172.36.173:9500", "DOS_node", logrus.Fields{
		"DOS_node_ip": p.GetIP(),
		"Serial":      string(common.BytesToAddress(p.GetID()).String()),
	})
	if err != nil {
		mainLogger.WithField("function", "newHookWithFields").Error(err)
	}

	log.Hooks.Add(hook)

	//3)Dial to peers to build peerClient and Set node ID
	if role == "BootstrapNode" {
		address, err := chainConn.GetWhitelist()
		if err != nil {
			mainLogger.WithField("function", "getWhitelist").Fatal(err)
		}

		if bytes.Compare(address.Bytes(), chainConn.GetId()) != 0 {
			if err = chainConn.InitialWhiteList(); err != nil {
				mainLogger.WithField("function", "initialWhiteList").Fatal(err)
			}
		}
	} else {
		if err = p.Join(bootstrapIp); err != nil {
			mainLogger.WithField("function", "join").Fatal(err)
		}
	}

	//4)Build a p2pDKG
	suite := suites.MustFind("bn256")
	peerEventForDKG := make(chan p2p.P2PMessage, 1)
	defer close(peerEventForDKG)
	p2pDkg := dkg.CreateP2PDkg(p, suite, peerEventForDKG)
	if err != nil {
		mainLogger.WithField("function", "createP2PDkg").Fatal(err)
	}

	//5)Subscribe Event from Eth
	eventGrouping := make(chan interface{}, 1)
	defer close(eventGrouping)
	chUrl := make(chan interface{}, 100)
	defer close(chUrl)
	chRandom := make(chan interface{}, 100)
	defer close(chRandom)
	chUsrRandom := make(chan interface{}, 100)
	defer close(chUsrRandom)
	cSignatureFromPeer := make(chan vss.Signature, 100)
	defer close(cSignatureFromPeer)
	eventValidation := make(chan interface{}, 20)
	defer close(eventValidation)
	if err = chainConn.SubscribeEvent(eventGrouping, onchain.SubscribeDOSProxyLogGrouping); err != nil {
		mainLogger.WithField("function", "subscribeEvent").Fatal(err)
	}

	if err = chainConn.SubscribeEvent(chUrl, onchain.SubscribeDOSProxyLogUrl); err != nil {
		mainLogger.WithField("function", "subscribeEvent").Fatal(err)
	}

	if err = chainConn.SubscribeEvent(chRandom, onchain.SubscribeDOSProxyLogUpdateRandom); err != nil {
		mainLogger.WithField("function", "subscribeEvent").Fatal(err)
	}

	if err = chainConn.SubscribeEvent(chUsrRandom, onchain.SubscribeDOSProxyLogRequestUserRandom); err != nil {
		mainLogger.WithField("function", "subscribeEvent").Fatal(err)
	}

	if err = chainConn.SubscribeEvent(eventValidation, onchain.SubscribeDOSProxyLogValidationResult); err != nil {
		mainLogger.WithField("function", "subscribeEvent").Fatal(err)
	}

	//6)Set up a dosnode pipeline
	d := dos.CreateDosNode(suite, nbParticipants, p, chainConn, p2pDkg, log.WithField("module", "dosNode"))
	d.PipeGrouping(eventGrouping)
	queriesReports := d.PipeQueries(chUrl, chUsrRandom, chRandom)
	outForRecover, outForValidate := d.PipeSignAndBroadcast(queriesReports)
	reportsToSubmit := d.PipeRecoverAndVerify(cSignatureFromPeer, outForRecover)
	submitForValidate := d.PipeSendToOnchain(reportsToSubmit)
	chRetrigerUrl := d.PipeCleanFinishMap(eventValidation, outForValidate, submitForValidate)

	if err = chainConn.UploadID(); err != nil {
		mainLogger.WithField("function", "uploadID").Fatal(err)
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
			case *dkg.Responses:
				peerEventForDKG <- msg
			case *vss.Signature:
				cSignatureFromPeer <- *content
				if err := msg.PeerConn.Reply(msg.RequestNonce, &vss.Signature{}); err != nil {
					mainLogger.WithField("function", "dispatchEvent").Error("signature reply", err)
				}
			default:
				mainLogger.WithField("function", "dispatchEvent").Warn("unknown", content)
			}
		case msg := <-p2pDkg.GetDkgEvent():
			if msg == dkg.VERIFIED {
				gId := new(big.Int)
				gId.SetBytes(p2pDkg.GetGroupId())
				if err := chainConn.UploadPubKey(p2pDkg.GetGroupPublicPoly().Commit()); err != nil {
					mainLogger.WithField("function", "uploadPubKey").Warn(err)
				}
			}
		//For re-trigger query
		case msg := <-chRetrigerUrl:
			chUrl <- msg
		}
	}
}
