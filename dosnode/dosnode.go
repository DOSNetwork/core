package dosnode

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
)

const (
	WATCHDOGINTERVAL = 10 //In minutes
	SYSRANDOMNTERVAL = 5  //In block numbers
)

var logger log.Logger

type DosNode struct {
	suite       suites.Suite
	chain       onchain.ProxyAdapter
	dkg         dkg.P2PDkgInterface
	p           p2p.P2PInterface
	logger      log.Logger
	done        chan interface{}
	cSignToPeer chan *vss.Signature
	cPipCancel  chan string
	id          []byte
}

func NewDosNode(credentialPath, passphrase string) (dosNode *DosNode, err error) {
	if passphrase == "" {
		passphrase = os.Getenv("PASSPHRASE")
		if passphrase == "" {
			err = errors.New("No passphrase")
			return
		}
	}

	//Read Configuration
	config := configuration.Config{}
	err = config.LoadConfig()
	if err != nil {
		return
	}

	role := config.NodeRole
	port := config.Port
	bootstrapIp := config.BootStrapIp

	workingDir, err := os.Getwd()
	if err != nil {
		return
	}
	if workingDir == "/" {
		workingDir = "."
	}

	if config.NodeRole == "testNode" {
		var credential []byte
		var resp *http.Response
		s := strings.Split(config.BootStrapIp, ":")
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

		if err = resp.Body.Close(); err != nil {
			return
		}

		credentialPath = workingDir + "/testAccounts/" + string(credential) + "/credential"
	} else if credentialPath == "" {
		credentialPath = workingDir + "/credential"
	}

	//Set up an onchain adapter
	chainConfig := config.GetChainConfig()
	chainConn, err := onchain.NewProxyAdapter(config.GetCurrentType(), credentialPath, passphrase, chainConfig.DOSProxyAddress, chainConfig.RemoteNodeAddressPool)
	if err != nil {
		return
	}
	id := chainConn.Address()
	//Init log module with nodeID that is an onchain account address
	log.Init(id[:])

	//Build a p2p network
	p, err := p2p.CreateP2PNetwork(id, port)
	if err != nil {
		return
	}

	err = p.Listen()
	if err != nil {
		return
	}

	//Bootstrapping p2p network
	if role != "BootstrapNode" {
		err = p.Join(bootstrapIp)
		if err != nil {
			return
		}
	}

	//Build a p2pDKG
	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(p, suite)
	if err != nil {
		return
	}

	if logger == nil {
		logger = log.New("module", "dosclient")
	}

	dosNode = &DosNode{
		suite:       suite,
		p:           p,
		chain:       chainConn,
		dkg:         p2pDkg,
		done:        make(chan interface{}),
		cSignToPeer: make(chan *vss.Signature, 50),
		cPipCancel:  make(chan string),
		id:          id,
	}
	err = dosNode.Start()
	ctx, _ := context.WithCancel(context.Background())
	errc := dosNode.chain.UploadID(ctx)
	<-errc
	return
}

func (d *DosNode) Start() (err error) {
	if err = d.listen(); err != nil {
		logger.Error(err)
		return
	}

	return
}

func (d *DosNode) End() {
	close(d.done)
}

func (d *DosNode) waitForGrouping(ctx context.Context, errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "waitForGrouping", map[string]interface{}{"SessionID": ctx.Value("SessionID")})

	for err := range errc {
		if err != nil {
			logger.Error(err)
		}
	}
}

func (d *DosNode) waitForRequestDone(ctx context.Context, requestID string, errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "WaitForRequestDone", map[string]interface{}{"RequestId": ctx.Value("RequestID")})
	for err := range errc {
		if err != nil {
			logger.Error(err)
		}
	}
	d.cPipCancel <- requestID

}

func (d *DosNode) buildPipeline(valueCtx context.Context, pubkey [4]*big.Int, requestID, lastRand, useSeed *big.Int, url, selector string, pType uint32) {
	defer logger.TimeTrack(time.Now(), "BuildPipeline", map[string]interface{}{"RequestId": valueCtx.Value("RequestID")})
	var signShares []<-chan *vss.Signature
	var errcList []<-chan error
	var cSubmitter []chan []byte
	var cErr <-chan error
	var cSign <-chan *vss.Signature
	var cContent <-chan []byte

	ids := d.dkg.GetGroupIDs(pubkey)
	pubPoly := d.dkg.GetGroupPublicPoly(pubkey)

	//Build a pipeline
	cSubmitter, cErr = choseSubmitter(valueCtx, d.chain, lastRand, ids, len(ids))
	errcList = append(errcList, cErr)

	switch pType {
	case onchain.TrafficSystemRandom:
		cContent = genSysRandom(valueCtx, cSubmitter[0], lastRand.Bytes())
	case onchain.TrafficUserRandom:
		cContent = genUserRandom(valueCtx, cSubmitter[0], requestID.Bytes(), lastRand.Bytes(), useSeed.Bytes())
	case onchain.TrafficUserQuery:
		cContent, cErr = genQueryResult(valueCtx, cSubmitter[0], url, selector)
		errcList = append(errcList, cErr)
	}

	cSign, cErr = genSign(valueCtx, cContent, d.cSignToPeer, d.dkg, d.suite, d.id, pubkey, requestID.String(), pType)
	errcList = append(errcList, cErr)

	signShares = append(signShares, cSign)
	idx := 1
	for _, id := range ids {
		if r := bytes.Compare(d.id, id); r != 0 {
			cSign, cErr = requestSign(valueCtx, cSubmitter[idx], d.p, d.id, requestID.String(), pType, id)
			signShares = append(signShares, cSign)
			errcList = append(errcList, cErr)
			idx++
		}
	}

	cSign, cErr = recoverSign(valueCtx, fanIn(valueCtx, signShares...), d.suite, pubPoly, (len(ids)/2 + 1), len(ids))
	errcList = append(errcList, cErr)

	switch pType {
	case onchain.TrafficSystemRandom:
		cErr = d.chain.SetRandomNum(valueCtx, cSign)
	default:
		cErr = d.chain.DataReturn(valueCtx, cSign)
	}
	errcList = append(errcList, cErr)
	errc := mergeErrors(valueCtx, errcList...)

	go d.waitForRequestDone(valueCtx, requestID.String(), errc)

}

func (d *DosNode) listen() (err error) {
	eventGrouping := make(chan interface{})
	eventGroupDismiss := make(chan interface{})
	chUrl := make(chan interface{})
	chRandom := make(chan interface{})
	chUsrRandom := make(chan interface{})
	eventValidation := make(chan interface{})
	keyUploaded := make(chan interface{})
	keyAccepted := make(chan interface{})
	d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogGrouping, eventGrouping)
	d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogGroupDismiss, eventGroupDismiss)
	d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogUrl, chUrl)
	d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogUpdateRandom, chRandom)
	d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogRequestUserRandom, chUsrRandom)
	d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogValidationResult, eventValidation)
	d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogPublicKeyUploaded, keyUploaded)
	d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogPublicKeyAccepted, keyAccepted)
	peerEvent, err := d.p.SubscribeEvent(50, vss.Signature{})

	go func() {
		pipeCancel := make(map[string]context.CancelFunc)
		peerSignMap := make(map[string]*vss.Signature)
		watchdog := time.NewTicker(WATCHDOGINTERVAL * time.Minute)

		defer close(eventGrouping)
		defer close(chUrl)
		defer close(chRandom)
		defer close(chUsrRandom)
		defer close(eventValidation)
		defer watchdog.Stop()
		defer d.p.UnSubscribeEvent(vss.Signature{})
		for {
			currentBlockNumber, err := d.chain.CurrentBlock()
			if err != nil {
				logger.Error(err)
			}
			select {
			case <-watchdog.C:
				ids := d.dkg.GetAnyGroupIDs()
				if len(ids) != 0 {
					lastUpdatedBlock, err := d.chain.LastUpdatedBlock()
					if err != nil {
						logger.Error(err)
						continue
					}
					diff := currentBlockNumber - lastUpdatedBlock
					if diff > SYSRANDOMNTERVAL {
						ctx, _ := context.WithCancel(context.Background())
						d.chain.RandomNumberTimeOut(ctx)
					}
				}
			case requestID := <-d.cPipCancel:
				if pipeCancel[requestID] != nil {
					a, _ := new(big.Int).SetString(requestID, 10)
					f := map[string]interface{}{"RequestId": fmt.Sprintf("%x", a)}
					logger.TimeTrack(time.Now(), "EndPipeline", f)

					pipeCancel[requestID]()
					delete(pipeCancel, requestID)
				}
			case msg := <-d.cSignToPeer:
				peerSignMap[msg.QueryId] = msg
			case msg := <-peerEvent:
				switch content := msg.Msg.Message.(type) {
				case *vss.Signature:
					d.p.Reply(msg.Sender, msg.RequestNonce, peerSignMap[content.QueryId])
					fmt.Println("Reply 1!!!!!!!!!!")
				}
			case msg := <-eventGrouping:
				content, ok := msg.(*onchain.DOSProxyLogGrouping)
				if !ok {
					log.Error(err)
					continue
				}
				isMember := false
				var groupIds [][]byte
				for _, node := range content.NodeId {
					id := node.Bytes()
					if r := bytes.Compare(d.id, id); r == 0 {
						isMember = true
					}
					groupIds = append(groupIds, id)
				}

				if isMember {
					sessionId := fmt.Sprintf("%x", content.GroupId)
					f := map[string]interface{}{
						"SessionID": sessionId,
						"Removed":   content.Removed,
						"Tx":        content.Tx,
						"CurBlkN":   currentBlockNumber,
						"BlockN":    content.BlockN,
						"Time":      time.Now()}
					logger.Event("DOS_Grouping", f)
					if !content.Removed {
						ctx, cancelFunc := context.WithCancel(context.Background())
						valueCtx := context.WithValue(ctx, "SessionID", sessionId)
						pipeCancel[sessionId] = cancelFunc
						var errcList []<-chan error
						outFromDkg, errc := d.dkg.Start(valueCtx, groupIds, sessionId)
						errcList = append(errcList, errc)
						errc = d.chain.UploadPubKey(valueCtx, outFromDkg)
						errcList = append(errcList, errc)
						errc = mergeErrors(valueCtx, errcList...)
						go d.waitForGrouping(valueCtx, errc)
					}
				}
			case msg := <-eventGroupDismiss:
				content, ok := msg.(*onchain.DOSProxyLogGroupDismiss)
				if !ok {
					e, ok := msg.(error)
					if ok {
						logger.Error(e)
					}
					continue
				}

				if d.isMember(content.PubKey) && d.dkg.GroupDismiss(content.PubKey) {
					f := map[string]interface{}{
						"SessionID":         d.dkg.GetSessionID(content.PubKey),
						"DispatchedGroup_1": fmt.Sprintf("%x", content.PubKey[0].Bytes()),
						"DispatchedGroup_2": fmt.Sprintf("%x", content.PubKey[1].Bytes()),
						"DispatchedGroup_3": fmt.Sprintf("%x", content.PubKey[2].Bytes()),
						"DispatchedGroup_4": fmt.Sprintf("%x", content.PubKey[3].Bytes()),
						"Removed":           content.Removed,
						"Tx":                content.Tx,
						"CurBlkN":           currentBlockNumber,
						"BlockN":            content.BlockN,
						"Time":              time.Now()}
					logger.Event("DOS_GroupDismiss", f)
					ctx, _ := context.WithCancel(context.Background())
					d.chain.UploadID(ctx)

				}
			case msg := <-chRandom:
				content, ok := msg.(*onchain.DOSProxyLogUpdateRandom)
				if !ok {
					log.Error(err)
				}
				if d.isMember(content.DispatchedGroup) {
					f := map[string]interface{}{
						"LastSystemRandomness": fmt.Sprintf("%x", content.LastRandomness),
						"DispatchedGroup_1":    fmt.Sprintf("%x", content.DispatchedGroup[0].Bytes()),
						"DispatchedGroup_2":    fmt.Sprintf("%x", content.DispatchedGroup[1].Bytes()),
						"DispatchedGroup_3":    fmt.Sprintf("%x", content.DispatchedGroup[2].Bytes()),
						"DispatchedGroup_4":    fmt.Sprintf("%x", content.DispatchedGroup[3].Bytes()),
						"Removed":              content.Removed,
						"Tx":                   content.Tx,
						"CurBlkN":              currentBlockNumber,
						"BlockN":               content.BlockN,
						"Time":                 time.Now()}
					logger.Event("DOS_QuerySysRandom", f)
					if !content.Removed {
						fmt.Println("EthLog : systemRandom")
						requestID := content.LastRandomness.String()
						//if pipeCancel[requestID] != nil {
						//	pipeCancel[requestID]()
						//	delete(pipeCancel, requestID)
						//}
						ctx, cancelFunc := context.WithCancel(context.Background())
						valueCtx := context.WithValue(ctx, "RequestID", fmt.Sprintf("%x", content.LastRandomness))
						pipeCancel[requestID] = cancelFunc
						d.buildPipeline(valueCtx, content.DispatchedGroup, content.LastRandomness, content.LastRandomness, nil, "", "", uint32(onchain.TrafficSystemRandom))
					}
				}
			case msg := <-chUsrRandom:
				content, ok := msg.(*onchain.DOSProxyLogRequestUserRandom)
				if !ok {
					log.Error(err)
					continue
				}
				if d.isMember(content.DispatchedGroup) {
					f := map[string]interface{}{
						"RequestId":            fmt.Sprintf("%x", content.RequestId),
						"LastSystemRandomness": fmt.Sprintf("%x", content.LastSystemRandomness),
						"DispatchedGroup_1":    fmt.Sprintf("%x", content.DispatchedGroup[0].Bytes()),
						"DispatchedGroup_2":    fmt.Sprintf("%x", content.DispatchedGroup[1].Bytes()),
						"DispatchedGroup_3":    fmt.Sprintf("%x", content.DispatchedGroup[2].Bytes()),
						"DispatchedGroup_4":    fmt.Sprintf("%x", content.DispatchedGroup[3].Bytes()),
						"Removed":              content.Removed,
						"Tx":                   content.Tx,
						"CurBlkN":              currentBlockNumber,
						"BlockN":               content.BlockN,
						"Time":                 time.Now()}
					logger.Event("DOS_QueryUserRandom", f)
					if !content.Removed {
						fmt.Println("EthLog : userRandom")
						requestID := content.RequestId.String()
						//if pipeCancel[requestID] != nil {
						//	pipeCancel[requestID]()
						//	delete(pipeCancel, requestID)
						//}
						ctx, cancelFunc := context.WithCancel(context.Background())
						valueCtx := context.WithValue(ctx, "RequestID", fmt.Sprintf("%x", content.RequestId))
						pipeCancel[requestID] = cancelFunc
						d.buildPipeline(valueCtx, content.DispatchedGroup, content.RequestId, content.LastSystemRandomness, content.UserSeed, "", "", uint32(onchain.TrafficUserRandom))
					}
				}
			case msg := <-chUrl:
				content, ok := msg.(*onchain.DOSProxyLogUrl)
				if !ok {
					log.Error(err)
					continue
				}
				if d.isMember(content.DispatchedGroup) {
					startTime := time.Now()
					fmt.Println("set up time ", startTime)
					f := map[string]interface{}{
						"RequestId":         fmt.Sprintf("%x", content.QueryId),
						"Randomness":        fmt.Sprintf("%x", content.Randomness),
						"DataSource":        fmt.Sprintf("%x", content.DataSource),
						"DispatchedGroup_1": fmt.Sprintf("%x", content.DispatchedGroup[0].Bytes()),
						"DispatchedGroup_2": fmt.Sprintf("%x", content.DispatchedGroup[1].Bytes()),
						"DispatchedGroup_3": fmt.Sprintf("%x", content.DispatchedGroup[2].Bytes()),
						"DispatchedGroup_4": fmt.Sprintf("%x", content.DispatchedGroup[3].Bytes()),
						"Removed":           content.Removed,
						"Tx":                content.Tx,
						"CurBlkN":           currentBlockNumber,
						"BlockN":            content.BlockN,
						"Time":              time.Now()}
					logger.Event("DOS_QueryURL", f)
					if !content.Removed {
						fmt.Println("EthLog : queryLog")
						requestID := content.QueryId.String()
						//if pipeCancel[requestID] != nil {
						//	pipeCancel[requestID]()
						//	delete(pipeCancel, requestID)
						//}
						ctx, cancelFunc := context.WithCancel(context.Background())
						valueCtx := context.WithValue(ctx, "RequestID", fmt.Sprintf("%x", content.QueryId))
						pipeCancel[requestID] = cancelFunc
						d.buildPipeline(valueCtx, content.DispatchedGroup, content.QueryId, content.Randomness, nil, content.DataSource, content.Selector, uint32(onchain.TrafficUserQuery))
					}
				}
			case msg := <-eventValidation:
				content, ok := msg.(*onchain.DOSProxyLogValidationResult)
				if !ok {
					log.Error(err)
					continue
				}
				if d.isMember(content.PubKey) {
					fmt.Println("EthLog : validationResult", content.Pass)
					if content.TrafficType == onchain.TrafficUserQuery {
						z := new(big.Int)
						z.SetBytes(content.Message)
						f := map[string]interface{}{
							"RequestId":         fmt.Sprintf("%x", content.TrafficId),
							"QueryResult":       string(content.Message),
							"ValidationPass":    content.Pass,
							"DispatchedGroup_1": fmt.Sprintf("%x", content.PubKey[0].Bytes()),
							"DispatchedGroup_2": fmt.Sprintf("%x", content.PubKey[1].Bytes()),
							"DispatchedGroup_3": fmt.Sprintf("%x", content.PubKey[2].Bytes()),
							"DispatchedGroup_4": fmt.Sprintf("%x", content.PubKey[3].Bytes()),
							"Removed":           content.Removed,
							"Tx":                content.Tx,
							"CurBlkN":           currentBlockNumber,
							"BlockN":            content.BlockN,
							"Time":              time.Now()}
						logger.Event("DOS_UrlResult", f)
					} else if content.TrafficType == onchain.TrafficUserRandom {
						z := new(big.Int)
						z.SetBytes(content.Message)
						f := map[string]interface{}{
							"RequestId":         fmt.Sprintf("%x", content.TrafficId),
							"GeneratedRandom":   fmt.Sprintf("%x", z),
							"ValidationPass":    content.Pass,
							"DispatchedGroup_1": fmt.Sprintf("%x", content.PubKey[0].Bytes()),
							"DispatchedGroup_2": fmt.Sprintf("%x", content.PubKey[1].Bytes()),
							"DispatchedGroup_3": fmt.Sprintf("%x", content.PubKey[2].Bytes()),
							"DispatchedGroup_4": fmt.Sprintf("%x", content.PubKey[3].Bytes()),
							"Removed":           content.Removed,
							"Tx":                content.Tx,
							"CurBlkN":           currentBlockNumber,
							"BlockN":            content.BlockN,
							"Time":              time.Now()}
						logger.Event("DOS_UserRandomResult", f)
					} else if content.TrafficType == onchain.TrafficSystemRandom {
						z := new(big.Int)
						z.SetBytes(content.Message)
						f := map[string]interface{}{
							"RequestId":       fmt.Sprintf("%x", content.TrafficId),
							"GeneratedRandom": fmt.Sprintf("%x", z),
							"ValidationPass":  content.Pass,
							"Removed":         content.Removed,
							"Tx":              content.Tx,
							"CurBlkN":         currentBlockNumber,
							"BlockN":          content.BlockN,
							"Time":            time.Now()}
						logger.Event("DOS_SysRandomResult", f)
					}
				}
			case msg := <-keyUploaded:
				content, ok := msg.(*onchain.DOSProxyLogPublicKeyUploaded)
				if !ok {
					e, ok := msg.(error)
					if ok {
						logger.Error(e)
					}
					continue
				}
				if d.isMember(content.PubKey) {
					logger.TimeTrack(time.Now(), "keyUploaded", map[string]interface{}{
						"SessionID": d.dkg.GetSessionID(content.PubKey),
						"groupId":   fmt.Sprintf("%x", content.GroupId),
						"x0":        fmt.Sprintf("%x", content.PubKey[0]),
						"x1":        fmt.Sprintf("%x", content.PubKey[1]),
						"y0":        fmt.Sprintf("%x", content.PubKey[2]),
						"y1":        fmt.Sprintf("%x", content.PubKey[3]),
						"Count":     fmt.Sprintf("%d", content.Count),
						"GroupSize": fmt.Sprintf("%d", content.GroupSize),
						"Removed":   content.Removed,
						"Tx":        content.Tx,
						"BlockN":    content.BlockN,
					})
				}
			case msg := <-keyAccepted:
				content, ok := msg.(*onchain.DOSProxyLogPublicKeyAccepted)
				if !ok {
					e, ok := msg.(error)
					if ok {
						logger.Error(e)
						d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogPublicKeyAccepted, keyAccepted)
					}
					continue
				}
				if d.isMember(content.PubKey) {
					logger.TimeTrack(time.Now(), "keyAccepted", map[string]interface{}{
						"SessionID":        d.dkg.GetSessionID(content.PubKey),
						"groupId":          fmt.Sprintf("%x", content.GroupId),
						"x0":               fmt.Sprintf("%x", content.PubKey[0]),
						"x1":               fmt.Sprintf("%x", content.PubKey[1]),
						"y0":               fmt.Sprintf("%x", content.PubKey[2]),
						"y1":               fmt.Sprintf("%x", content.PubKey[3]),
						"workingGroupSize": content.WorkingGroupSize,
						"Removed":          content.Removed,
						"Tx":               content.Tx,
						"BlockN":           content.BlockN,
					})
				}
			case <-d.done:
				return
			default:
			}
		}
	}()
	return
}

func (d *DosNode) isMember(pubkey [4]*big.Int) bool {
	return d.dkg.GetShareSecurity(pubkey) != nil
}
