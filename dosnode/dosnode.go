package dosnode

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto/sha3"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
)

const (
	WATCHDOGINTERVAL = 20 //In minutes
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
	cPipCancel  chan []byte
	id          []byte
}

func NewDosNode(credentialPath, passphrase string) (dosNode *DosNode, err error) {
	if passphrase == "" {
		passphrase = os.Getenv(configuration.ENVPASSPHRASE)
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

	port := config.Port
	bootstrapIp := config.BootStrapIp

	workingDir, err := os.Getwd()
	if err != nil {
		return
	}
	if workingDir == "/" {
		workingDir = "."
	}

	chainConfig := config.GetChainConfig()
	if config.NodeRole == "testNode" {
		var rspBytes []byte
		var resp *http.Response
		ip := config.BootStrapIp
		tServer := "http://" + ip + ":8080/getCredential"
		resp, err = http.Get(tServer)
		for err != nil {
			time.Sleep(1 * time.Second)
			resp, err = http.Get(tServer)
		}

		rspBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		if err = resp.Body.Close(); err != nil {
			return
		}

		respArray := strings.Split(string(rspBytes), ",")
		if len(respArray) < 1 {
			return nil, errors.New("cannot get credential from bootNode")
		}

		credentialPath = workingDir + "/testAccounts/" + respArray[0] + "/credential"
		chainConfig.RemoteNodeAddressPool = chainConfig.RemoteNodeAddressPool[:1]
		for _, address := range respArray[1:] {
			chainConfig.RemoteNodeAddressPool = append(chainConfig.RemoteNodeAddressPool, address)
		}
		fmt.Println("RemoteNodeAddressPool", chainConfig.RemoteNodeAddressPool)
	} else if credentialPath == "" {
		credentialPath = workingDir + "/credential"
	}

	//Set up an onchain adapter
	chainConn, err := onchain.NewProxyAdapter(config.GetCurrentType(), credentialPath, passphrase, chainConfig.DOSProxyAddress, chainConfig.DOSCommitReveal, chainConfig.RemoteNodeAddressPool)
	if err != nil {
		if err.Error() != "No any working eth client for event tracking" {
			fmt.Println("NewDosNode failed ", err)
			return
		}
	}
	id := chainConn.Address()

	//Build a p2p network
	p, err := p2p.CreateP2PNetwork(id, port, p2p.SWIM)
	if err != nil {
		fmt.Println("CreateP2PNetwork err ", err)
		return
	}

	err = p.Listen()
	if err != nil {
		fmt.Println("Listen() err ", err)
		return
	}

	//Bootstrapping p2p network
	fmt.Println("Join :", bootstrapIp)
	err = p.Join([]string{bootstrapIp})
	if err != nil {
		fmt.Println("Join ", err)
	}

	//Build a p2pDKG
	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(p, suite)
	if err != nil {
		fmt.Println("p2pDKG ", err)
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
		cPipCancel:  make(chan []byte),
		id:          id,
	}
	return dosNode, nil
}

func (d *DosNode) Start() (err error) {
	if err = d.listen(); err != nil {
		fmt.Println("listen err ", err)
		logger.Error(err)
		return
	}
	ctx, _ := context.WithCancel(context.Background())
	_ = d.chain.RegisterNewNode(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/", d.status)
	mux.HandleFunc("/balance", d.balance)
	mux.HandleFunc("/groupSize", d.groupSize)
	mux.HandleFunc("/proxy", d.proxy)
	mux.HandleFunc("/guardian", d.guardian)
	mux.HandleFunc("/p2p", d.p2p)

	http.ListenAndServe("localhost:8080", mux)
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

func (d *DosNode) waitForRequestDone(ctx context.Context, requestID []byte, errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "WaitForRequestDone", map[string]interface{}{"RequestId": ctx.Value("RequestID")})
	for err := range errc {
		if err != nil {
			logger.Error(err)
		}
	}
	d.cPipCancel <- requestID

}

func (d *DosNode) buildPipeline(valueCtx context.Context, pubkey [4]*big.Int, groupID, requestID, lastRand, useSeed *big.Int, url, selector string, pType uint32) {
	defer logger.TimeTrack(time.Now(), "BuildPipeline", map[string]interface{}{"RequestId": valueCtx.Value("RequestID")})
	var signShares []<-chan *vss.Signature
	var errcList []<-chan error
	var cSubmitter []chan []byte
	var cErr <-chan error
	var cSign <-chan *vss.Signature
	var cContent <-chan []byte
	var nonce []byte
	ids := d.dkg.GetGroupIDs(pubkey)
	pubPoly := d.dkg.GetGroupPublicPoly(pubkey)

	//Generate an unique id
	switch pType {
	case onchain.TrafficSystemRandom:
		var bytes []byte
		bytes = append(bytes, groupID.Bytes()...)
		bytes = append(bytes, requestID.Bytes()...)
		bytes = append(bytes, lastRand.Bytes()...)
		nHash := sha256.Sum256(bytes)
		nonce = nHash[:]
	case onchain.TrafficUserRandom:
		var bytes []byte
		bytes = append(bytes, groupID.Bytes()...)
		bytes = append(bytes, requestID.Bytes()...)
		bytes = append(bytes, lastRand.Bytes()...)
		bytes = append(bytes, useSeed.Bytes()...)
		nHash := sha256.Sum256(bytes)
		nonce = nHash[:]
	case onchain.TrafficUserQuery:
		var bytes []byte
		bytes = append(bytes, groupID.Bytes()...)
		bytes = append(bytes, requestID.Bytes()...)
		bytes = append(bytes, lastRand.Bytes()...)
		bytes = append(bytes, []byte(url)...)
		bytes = append(bytes, []byte(selector)...)
		nHash := sha256.Sum256(bytes)
		nonce = nHash[:]
	}

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

	cSign, cErr = genSign(valueCtx, cContent, d.cSignToPeer, d.dkg, d.suite, d.id, pubkey, requestID.Bytes(), pType, nonce)
	errcList = append(errcList, cErr)

	signShares = append(signShares, cSign)
	idx := 1
	for _, id := range ids {
		if r := bytes.Compare(d.id, id); r != 0 {
			cSign, cErr = requestSign(valueCtx, cSubmitter[idx], cContent, d.p, d.id, requestID.Bytes(), pType, id, nonce)
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

	go d.waitForRequestDone(valueCtx, requestID.Bytes(), errc)

}

func (d *DosNode) listen() (err error) {
	keyUploaded := make(chan interface{})

	var errcList []<-chan error
	eventGrouping, errc := d.chain.PollLogs(onchain.SubscribeDOSProxyLogGrouping, 5, 0)
	errcList = append(errcList, errc)
	eventGroupDismiss, errc := d.chain.PollLogs(onchain.SubscribeDOSProxyLogGroupDismiss, 30, 0)
	errcList = append(errcList, errc)

	chUrl, errc := d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogUrl)
	errcList = append(errcList, errc)
	chRandom, errc := d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogUpdateRandom)
	errcList = append(errcList, errc)
	chUsrRandom, errc := d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogRequestUserRandom)
	errcList = append(errcList, errc)
	eventValidation, errc := d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogValidationResult)
	errcList = append(errcList, errc)
	keyAccepted, errc := d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogPublicKeyAccepted)
	errcList = append(errcList, errc)
	noworkinggroup, errc := d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogNoWorkingGroup)
	errcList = append(errcList, errc)
	/*
		chInsufficientWG, errc := d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogInsufficientWorkingGroup)
		errcList = append(errcList, errc)
		groupInitated, errc := d.chain.SubscribeEvent(onchain.SubscribeDOSProxyLogGroupingInitiated)
		errcList = append(errcList, errc)
	*/
	commitRevealStart, errc := d.chain.PollLogs(onchain.SubscribeDOSCommitRevealLogStartCommitReveal, 0, 0)
	errcList = append(errcList, errc)

	peerEvent, err := d.p.SubscribeEvent(50, vss.Signature{})
	errcList = append(errcList, errc)
	errc = mergeErrors(context.Background(), errcList...)

	go func() {
		pipeCancel := make(map[string]context.CancelFunc)
		peerSignMap := make(map[string]*vss.Signature)
		watchdog := time.NewTicker(WATCHDOGINTERVAL * time.Minute)

		defer watchdog.Stop()
		defer d.p.UnSubscribeEvent(vss.Signature{})

		for {

			select {
			/*
				case _, ok := <-chInsufficientWG:
					if !ok {
						continue
					}

					if d.dkg.GetGroupNumber() == 0 {
						currentBlockNumber, err := d.chain.CurrentBlock()
						if err != nil {
							logger.Error(err)
						}
						commitRevealTargetBlk, err := d.chain.CommitRevealTargetBlk()
						if err != nil {
							continue
						}
						if currentBlockNumber > commitRevealTargetBlk {
							f := map[string]interface{}{
								"Time": time.Now()}
							logger.Event("InsufficientWorkingGroup", f)
							d.chain.SignalGroupFormation(context.Background())
						}
					}

				case msg, ok := <-groupInitated:
					if !ok {
						continue
					}
					content, ok := msg.(*onchain.DOSProxyLogGroupingInitiated)
					if !ok {
						log.Error(err)
						continue
					}
					if d.dkg.GetGroupNumber() == 0 {
						f := map[string]interface{}{
							"NumPendingNodes":   content.NumPendingNodes,
							"GroupSize":         content.GroupSize,
							"GroupingThreshold": content.GroupingThreshold,
							"Time":              time.Now()}
						logger.Event("GroupingInitiated", f)
					}*/
			case msg, ok := <-commitRevealStart:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DOSCommitRevealLogStartCommitReveal)
				if !ok {
					log.Error(err)
					continue
				}
				go func(blkNumUint uint64, commitDurUint uint64, revealDurUint uint64) {
					f := map[string]interface{}{
						"blkNumUint":    blkNumUint,
						"commitDurUint": commitDurUint,
						"revealDurUint": revealDurUint,
						"Time":          time.Now()}
					logger.Event("commitRevealStart", f)
					var hash *[32]byte

					var prime1 *big.Int
					// Generate random numbers in range [0..prime1]
					prime1, ok = new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
					if !ok {
						return
					}
					// Don't use this code to generate secret keys that protect important stuff!
					sec, err := rand.Int(rand.Reader, prime1)
					if err != nil {
						sec.SetInt64(0)
						return
					}

					h := sha3.NewKeccak256()
					h.Write(abi.U256(sec))
					b := h.Sum(nil)
					hash = byte32(b)
					for {
						cur, err := d.chain.CurrentBlock()
						if err != nil {
							logger.Error(err)
							return
						}
						fmt.Println("Waiting for commit ", cur, blkNumUint-revealDurUint-commitDurUint)
						if cur >= blkNumUint-revealDurUint-commitDurUint {
							break
						}
						time.Sleep(15 * time.Second)
					}
					errc := d.chain.Commit(context.Background(), *hash)
					err = <-errc
					if err != nil {
						return
					}
					for {
						cur, err := d.chain.CurrentBlock()
						if err != nil {
							logger.Error(err)
							return
						}
						fmt.Println("Waiting for Reveal ", cur, blkNumUint-revealDurUint)
						if cur > blkNumUint-revealDurUint {
							break
						}
						time.Sleep(15 * time.Second)
					}
					d.chain.Reveal(context.Background(), sec)
					for {
						cur, err := d.chain.CurrentBlock()
						if err != nil {
							logger.Error(err)
							return
						}
						fmt.Println("Waiting for random ", cur, blkNumUint)
						if cur > blkNumUint {
							break
						}
						time.Sleep(15 * time.Second)
					}
					d.chain.SignalGroupFormation(context.Background())
				}(content.TargetBlkNum.Uint64(), content.CommitDuration.Uint64(), content.RevealDuration.Uint64())

			case err, ok := <-errc:
				if ok && err.Error() == "EOF" {
					fmt.Println("!!!dosnode err ", err)
				}
				logger.Error(err)
				/*
					TODO: subscribe again and push err to errc,
					TODO: otherwise possibly closed errc will keep consuming "select",
				*/
			case <-watchdog.C:
				//Let pending node as a guardian
				if d.dkg.GetGroupNumber() == 0 {
					fmt.Println("watchdog")
					ids := d.dkg.GetAnyGroupIDs()
					workingGroup, err := d.chain.GetWorkingGroupSize()
					if err != nil {
						continue
					}
					groupToPick, err := d.chain.GetGroupToPick()
					if err != nil {
						continue
					}
					commitRevealTargetBlk, err := d.chain.CommitRevealTargetBlk()
					if err != nil {
						continue
					}
					pendingNodeSize, err := d.chain.GetPengindNodeSize()
					if err != nil {
						continue
					}
					pendingGrouSize, err := d.chain.NumPendingGroups()
					if err != nil {
						continue
					}
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					lastUpdatedBlock, err := d.chain.LastUpdatedBlock()
					if err != nil {
						logger.Error(err)
						continue
					}
					f := map[string]interface{}{
						"WorkingGroupSize":      workingGroup,
						"GroupToPick":           groupToPick,
						"currentBlockNumber":    currentBlockNumber,
						"commitRevealTargetBlk": commitRevealTargetBlk,
						"PendingNodeSize":       pendingNodeSize,
						"PendingGrouSize":       pendingGrouSize,
						"Members":               d.p.Members(),
						"Time":                  time.Now()}
					logger.Event("DOS_Watchdog", f)

					if len(ids) != 0 {
						diff := currentBlockNumber - lastUpdatedBlock
						if diff > SYSRANDOMNTERVAL {
							ctx, _ := context.WithCancel(context.Background())
							d.chain.SignalRandom(ctx)
						}
					}

					if workingGroup < groupToPick {
						if currentBlockNumber > commitRevealTargetBlk {
							d.chain.SignalGroupFormation(context.Background())
						}
					}
				}
			case requestID := <-d.cPipCancel:
				_ = requestID
			/*
				if pipeCancel[requestID] != nil {
					a, _ := new(big.Int).SetString(requestID, 10)
					f := map[string]interface{}{"RequestId": fmt.Sprintf("%x", a)}
					logger.TimeTrack(time.Now(), "EndPipeline", f)

					pipeCancel[requestID]()
					delete(pipeCancel, requestID)
				}*/
			case msg, ok := <-d.cSignToPeer:
				if !ok {
					continue
				}
				peerSignMap[string(msg.Nonce)] = msg
			case msg := <-peerEvent:
				switch content := msg.Msg.Message.(type) {
				case *vss.Signature:
					if peerSignMap[string(content.Nonce)] != nil {
						fmt.Println("Got Sign ", peerSignMap[string(content.Nonce)].RequestId)
					}
					d.p.Reply(msg.Sender, msg.RequestNonce, peerSignMap[string(content.Nonce)])
				}
			case msg, ok := <-eventGrouping:
				if !ok {
					continue
				}
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
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
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
						errc = d.chain.RegisterGroupPubKey(valueCtx, outFromDkg)
						errcList = append(errcList, errc)
						errc = mergeErrors(valueCtx, errcList...)
						go d.waitForGrouping(valueCtx, errc)
					}
				}
			case msg, ok := <-eventGroupDismiss:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DOSProxyLogGroupDismiss)
				if !ok {
					e, ok := msg.(error)
					if ok {
						logger.Error(e)
					}
					continue
				}

				if d.isMember(content.PubKey) && d.dkg.GroupDismiss(content.PubKey) {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					f := map[string]interface{}{
						"SessionID":         fmt.Sprintf("%x", content.GroupId),
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
					go func() {

						errc := d.chain.RegisterNewNode(ctx)
						err := <-errc
						if err != nil {
							fmt.Println("RegisterNewNode err ", err)
						}
					}()

				}
			case msg, ok := <-chRandom:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DOSProxyLogUpdateRandom)
				if !ok {
					log.Error(err)
				}
				if d.isMember(content.DispatchedGroup) {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					f := map[string]interface{}{
						"LastSystemRandomness": fmt.Sprintf("%x", content.LastRandomness),
						"DispatchedGroupId":    fmt.Sprintf("%x", content.DispatchedGroupId.Bytes()),
						"DispatchedGroup_1":    fmt.Sprintf("%x", content.DispatchedGroup[0].Bytes()),
						"DispatchedGroup_2":    fmt.Sprintf("%x", content.DispatchedGroup[1].Bytes()),
						"DispatchedGroup_3":    fmt.Sprintf("%x", content.DispatchedGroup[2].Bytes()),
						"DispatchedGroup_4":    fmt.Sprintf("%x", content.DispatchedGroup[3].Bytes()),
						"Removed":              content.Removed,
						"IsMember":             d.isMember(content.DispatchedGroup),
						"Tx":                   content.Tx,
						"CurBlkN":              currentBlockNumber,
						"BlockN":               content.BlockN,
						"Time":                 time.Now()}
					logger.Event("DOS_QuerySysRandom", f)
					fmt.Println("EthLog : systemRandom", d.isMember(content.DispatchedGroup))
					if !content.Removed {
						requestID := content.LastRandomness.String()
						//if pipeCancel[requestID] != nil {
						//	pipeCancel[requestID]()
						//	delete(pipeCancel, requestID)
						//}
						ctx, cancelFunc := context.WithCancel(context.Background())
						valueCtx := context.WithValue(ctx, "RequestID", fmt.Sprintf("%x", content.LastRandomness))
						pipeCancel[requestID] = cancelFunc
						d.buildPipeline(valueCtx, content.DispatchedGroup, content.DispatchedGroupId, content.LastRandomness, content.LastRandomness, nil, "", "", uint32(onchain.TrafficSystemRandom))
					}
				}
			case msg, ok := <-chUsrRandom:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DOSProxyLogRequestUserRandom)
				if !ok {
					log.Error(err)
					continue
				}
				if d.isMember(content.DispatchedGroup) {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					f := map[string]interface{}{
						"RequestId":            fmt.Sprintf("%x", content.RequestId),
						"LastSystemRandomness": fmt.Sprintf("%x", content.LastSystemRandomness),
						"DispatchedGroupId":    fmt.Sprintf("%x", content.DispatchedGroupId),
						"DispatchedGroup_1":    fmt.Sprintf("%x", content.DispatchedGroup[0].Bytes()),
						"DispatchedGroup_2":    fmt.Sprintf("%x", content.DispatchedGroup[1].Bytes()),
						"DispatchedGroup_3":    fmt.Sprintf("%x", content.DispatchedGroup[2].Bytes()),
						"DispatchedGroup_4":    fmt.Sprintf("%x", content.DispatchedGroup[3].Bytes()),
						"Removed":              content.Removed,
						"IsMember":             d.isMember(content.DispatchedGroup),
						"Tx":                   content.Tx,
						"CurBlkN":              currentBlockNumber,
						"BlockN":               content.BlockN,
						"Time":                 time.Now()}
					logger.Event("DOS_QueryUserRandom", f)
					fmt.Println("EthLog : userRandom ", d.isMember(content.DispatchedGroup))
					if !content.Removed {
						requestID := content.RequestId.String()
						//if pipeCancel[requestID] != nil {
						//	pipeCancel[requestID]()
						//	delete(pipeCancel, requestID)
						//}
						ctx, cancelFunc := context.WithCancel(context.Background())
						valueCtx := context.WithValue(ctx, "RequestID", fmt.Sprintf("%x", content.RequestId))
						pipeCancel[requestID] = cancelFunc
						d.buildPipeline(valueCtx, content.DispatchedGroup, content.DispatchedGroupId, content.RequestId, content.LastSystemRandomness, content.UserSeed, "", "", uint32(onchain.TrafficUserRandom))
					}
				}
			case msg, ok := <-chUrl:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DOSProxyLogUrl)
				if !ok {
					log.Error(err)
					continue
				}
				if d.isMember(content.DispatchedGroup) {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					startTime := time.Now()
					fmt.Println("set up time ", startTime)
					f := map[string]interface{}{
						"RequestId":         fmt.Sprintf("%x", content.QueryId),
						"Randomness":        fmt.Sprintf("%x", content.Randomness),
						"DataSource":        fmt.Sprintf("%x", content.DataSource),
						"DispatchedGroupId": fmt.Sprintf("%x", content.DispatchedGroupId),
						"DispatchedGroup_1": fmt.Sprintf("%x", content.DispatchedGroup[0].Bytes()),
						"DispatchedGroup_2": fmt.Sprintf("%x", content.DispatchedGroup[1].Bytes()),
						"DispatchedGroup_3": fmt.Sprintf("%x", content.DispatchedGroup[2].Bytes()),
						"DispatchedGroup_4": fmt.Sprintf("%x", content.DispatchedGroup[3].Bytes()),
						"Removed":           content.Removed,
						"IsMember":          d.isMember(content.DispatchedGroup),
						"Tx":                content.Tx,
						"CurBlkN":           currentBlockNumber,
						"BlockN":            content.BlockN,
						"Time":              time.Now()}
					logger.Event("DOS_QueryURL", f)
					fmt.Println("EthLog : queryLog ", d.isMember(content.DispatchedGroup))
					if !content.Removed {
						requestID := content.QueryId.String()
						//if pipeCancel[requestID] != nil {
						//	pipeCancel[requestID]()
						//	delete(pipeCancel, requestID)
						//}
						ctx, cancelFunc := context.WithCancel(context.Background())
						valueCtx := context.WithValue(ctx, "RequestID", fmt.Sprintf("%x", content.QueryId))
						pipeCancel[requestID] = cancelFunc
						d.buildPipeline(valueCtx, content.DispatchedGroup, content.DispatchedGroupId, content.QueryId, content.Randomness, nil, content.DataSource, content.Selector, uint32(onchain.TrafficUserQuery))
					}
				}
			case msg, ok := <-eventValidation:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DOSProxyLogValidationResult)
				if !ok {
					log.Error(err)
					continue
				}
				if d.isMember(content.PubKey) {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					fmt.Println("EthLog : validationResult", content.TrafficType, content.Pass)
					if content.TrafficType == onchain.TrafficUserQuery {
						z := new(big.Int)
						z.SetBytes(content.Message)
						f := map[string]interface{}{
							"RequestId":         fmt.Sprintf("%x", content.TrafficId),
							"QueryResult":       string(content.Message),
							"ValidationPass":    content.Pass,
							"GroupId":           fmt.Sprintf("%x", content.GroupId),
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
							"GroupId":           fmt.Sprintf("%x", content.GroupId),
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
							"GeneratedRandom": fmt.Sprintf("0x%x", z),
							"GroupId":         fmt.Sprintf("%x", content.GroupId),
							"ValidationPass":  content.Pass,
							"Removed":         content.Removed,
							"Tx":              content.Tx,
							"CurBlkN":         currentBlockNumber,
							"BlockN":          content.BlockN,
							"Time":            time.Now()}
						logger.Event("DOS_SysRandomResult", f)
					}
				}
			case msg, ok := <-keyUploaded:
				if !ok {
					continue
				}
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
			case msg, ok := <-keyAccepted:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DOSProxyLogPublicKeyAccepted)
				if !ok {
					e, ok := msg.(error)
					if ok {
						logger.Error(e)
					}
					continue
				}
				if d.isMember(content.PubKey) {
					logger.TimeTrack(time.Now(), "keyAccepted", map[string]interface{}{
						"workingGroupSize": content.WorkingGroupSize,
						"Removed":          content.Removed,
						"Tx":               content.Tx,
						"BlockN":           content.BlockN,
					})
				}
			case msg, ok := <-noworkinggroup:
				if !ok {
					continue
				}
				currentBlockNumber, err := d.chain.CurrentBlock()
				if err != nil {
					logger.Error(err)
				}
				content, ok := msg.(*onchain.DOSProxyLogNoWorkingGroup)
				if !ok {
					e, ok := msg.(error)
					if ok {
						logger.Error(e)
					}
					continue
				}
				f := map[string]interface{}{
					"Removed": content.Removed,
					"Tx":      content.Tx,
					"CurBlkN": currentBlockNumber,
					"BlockN":  content.BlockN,
					"Time":    time.Now()}
				logger.Event("DOS_NOWORKINGGROUP", f)
			case <-d.done:
				return
			}
		}
	}()
	return
}

func (d *DosNode) isMember(pubkey [4]*big.Int) bool {
	return d.dkg.GetShareSecurity(pubkey) != nil
}

func byte32(s []byte) (a *[32]byte) {
	if len(a) <= len(s) {
		a = (*[len(a)]byte)(unsafe.Pointer(&s[0]))
	}
	return a
}
