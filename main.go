package main

import (
	"bytes"
	"fmt"
	"math/big"
	"net"
	_ "net/http/pprof"
	"runtime/debug"
	"time"

	"github.com/bshuster-repo/logrus-logstash-hook"

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
	offChainConfig := configuration.OffChainConfig{}
	if err := offChainConfig.LoadConfig(); err != nil {
		log.Fatal(err)
	}
	role := offChainConfig.NodeRole
	nbParticipants := offChainConfig.RandomGroupSize
	port := offChainConfig.Port
	bootstrapIp := offChainConfig.BootStrapIp

	onChainConfig := configuration.OnChainConfig{}
	if err := onChainConfig.LoadConfig(); err != nil {
		log.Fatal(err)
	}
	chainConfig := onChainConfig.GetChainConfig()

	//0)initial log module
	log = logrus.New()

	//1)Connect to Eth
	chainConn, err := onchain.AdaptTo(chainConfig.ChainType, &chainConfig, log)
	if err != nil {
		log.Fatal(err)
	}

	//2)Build a p2p network
	p, peerEvent, err := p2p.CreateP2PNetwork(chainConn.GetId(), port, log.WithFields(logrus.Fields{}))
	if err != nil {
		log.Fatal(err)
	}
	if err := p.Listen(); err != nil {
		log.Error(err)
	}

	//2.5)Add hook to log module
	tcpAddr, err := net.ResolveTCPAddr("tcp", "163.172.36.173:9500")
	if err != nil {
		log.Error(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error(err)
	}

	if err = conn.SetKeepAlivePeriod(time.Minute); err != nil {
		log.Warn(err)
	}

	if err = conn.SetKeepAlive(true); err != nil {
		log.Warn(err)
	}

	hook, err := logrustash.NewHookWithFieldsAndConn(conn, "DOS_node", logrus.Fields{
		"DOS_node_ip": p.GetIP(),
		"Serial":      string(common.BytesToAddress(p.GetID()).String()),
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
		err = p.Join(bootstrapIp)
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

	//6)Set up a dosnode pipeline
	d := dos.CreateDosNode(suite, nbParticipants, p, chainConn, p2pDkg, log)
	d.PipeGrouping(eventGrouping)
	queriesReports := d.PipeQueries(chUrl, chUsrRandom, chRandom)
	outForRecover, outForValidate := d.PipeSignAndBroadcast(queriesReports)
	reportsToSubmit := d.PipeRecoverAndVerify(cSignatureFromPeer, outForRecover)
	submitForValidate := d.PipeSendToOnchain(reportsToSubmit)
	chRetrigerUrl := d.PipeCleanFinishMap(eventValidation, outForValidate, submitForValidate)

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
				if err := chainConn.UploadPubKey(p2pDkg.GetGroupPublicPoly().Commit()); err != nil {
					fmt.Println(err)
				}
			}
		//For re-trigger query
		case msg := <-chRetrigerUrl:
			chUrl <- msg
		}
	}
}
