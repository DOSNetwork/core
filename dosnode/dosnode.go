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
	GROUPINITIATED   = 0
	GROUPDELETED     = 1
)

var logger log.Logger

type DosNode struct {
	suite        suites.Suite
	chain        onchain.ProxyAdapter
	dkg          dkg.P2PDkgInterface
	p            p2p.P2PInterface
	done         chan interface{}
	cSignToPeer  chan *vss.Signature
	cRequestDone chan [4]*big.Int
	id           []byte
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
		ip := config.BootStrapIp[0]
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
	chainConn, err := onchain.NewProxyAdapter(config.GetCurrentType(), credentialPath, passphrase, chainConfig.DOSProxyAddress, chainConfig.CommitReveal, chainConfig.RemoteNodeAddressPool)
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
	retry, num := 0, 0
	for {
		num, err = p.Join(bootstrapIp)
		if err != nil || num == 0 {
			fmt.Println("Join ", err, num)

			if retry == 10 {
				return
			}
			retry++
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	fmt.Println("Join : num of peer ", num)
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
		suite:        suite,
		p:            p,
		chain:        chainConn,
		dkg:          p2pDkg,
		done:         make(chan interface{}),
		cSignToPeer:  make(chan *vss.Signature, 21),
		cRequestDone: make(chan [4]*big.Int),
		id:           id,
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

	//http.ListenAndServe("localhost:8080", mux)
	http.ListenAndServe(":8080", mux)
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

func (d *DosNode) waitForRequestDone(ctx context.Context, pubKey [4]*big.Int, errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "WaitForRequestDone", map[string]interface{}{"RequestId": ctx.Value("RequestID")})
	for err := range errc {
		if err != nil {
			logger.Error(err)
		}
	}
	d.cRequestDone <- pubKey

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
	//errcList = append(errcList, cErr)
	errc := mergeErrors(valueCtx, errcList...)

	go d.waitForRequestDone(valueCtx, pubkey, errc)

}

func (d *DosNode) listen() (err error) {

	var errcList []<-chan error
	eventGrouping, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogGrouping)
	errcList = append(errcList, errc)
	eventGroupDissolve, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogGroupDissolve)
	errcList = append(errcList, errc)
	chUrl, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogUrl)
	errcList = append(errcList, errc)
	chRandom, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogUpdateRandom)
	errcList = append(errcList, errc)
	chUsrRandom, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogRequestUserRandom)
	errcList = append(errcList, errc)
	eventValidation, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogValidationResult)
	errcList = append(errcList, errc)
	keyAccepted, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogPublicKeyAccepted)
	errcList = append(errcList, errc)
	keySuggested, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogPublicKeySuggested)
	errcList = append(errcList, errc)
	noworkinggroup, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogNoWorkingGroup)
	errcList = append(errcList, errc)
	/*
		chInsufficientWG, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogInsufficientWorkingGroup)
		errcList = append(errcList, errc)

			groupInitated, errc := d.chain.SubscribeEvent(onchain.SubscribeDosproxyLogGroupingInitiated)
			errcList = append(errcList, errc)
	*/
	commitRevealStart, errc := d.chain.SubscribeEvent(onchain.SubscribeCommitrevealLogStartCommitreveal)
	errcList = append(errcList, errc)

	peerEvent, err := d.p.SubscribeEvent(50, vss.Signature{})
	errcList = append(errcList, errc)
	errc = mergeErrors(context.Background(), errcList...)

	go func() {
		pipeCancel := make(map[string]context.CancelFunc)
		peerSignMap := make(map[string]*vss.Signature)
		watchdog := time.NewTicker(WATCHDOGINTERVAL * time.Minute)
		requestTracking := map[string]map[string]int{}

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
			*/
			/*
				case msg, ok := <-groupInitated:
					if !ok {
						continue
					}
					content, ok := msg.(*onchain.DosproxyLogGroupingInitiated)
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
				content, ok := msg.(*onchain.LogStartCommitReveal)
				if !ok {
					log.Error(err)
					continue
				}
				go func(cid, StartBlock, CommitDuration, RevealDuration, RevealThreshold *big.Int) {
					f := map[string]interface{}{
						"cid": cid.Uint64()}
					logger.Event("commitRevealStart", f)

					var hash *[32]byte
					_ = hash
					var prime1 *big.Int
					// Generate random numbers in range [0..prime1]
					prime1, ok = new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
					if !ok {
						return
					}
					sec, err := rand.Int(rand.Reader, prime1)
					if err != nil {
						sec.SetInt64(0)
						return
					}

					h := sha3.NewKeccak256()
					h.Write(abi.U256(sec))
					b := h.Sum(nil)
					hash = byte32(b)
					startBlock := StartBlock.Uint64()
					commitDur := CommitDuration.Uint64()
					revealDur := RevealDuration.Uint64()
					fmt.Println("startBlock ", startBlock, " commitDur ", commitDur, "revealDur", revealDur)
					for {
						cur, err := d.chain.CurrentBlock()
						if err != nil {
							logger.Error(err)
							return
						}
						fmt.Println("Waiting for commit ", cur, startBlock)
						if cur >= startBlock {
							break
						}
						time.Sleep(15 * time.Second)
					}

					errc := d.chain.Commit(context.Background(), cid, *hash)
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
						fmt.Println("Waiting for Reveal ", cur, startBlock+commitDur)
						if cur > startBlock+commitDur {
							break
						}
						time.Sleep(15 * time.Second)
					}
					d.chain.Reveal(context.Background(), cid, sec)
					for {
						cur, err := d.chain.CurrentBlock()
						if err != nil {
							logger.Error(err)
							return
						}
						fmt.Println("Waiting for random ", cur, startBlock+commitDur+revealDur)
						if cur > startBlock+commitDur+revealDur {
							break
						}
						time.Sleep(15 * time.Second)
					}
					d.chain.SignalBootstrap(context.Background(), cid.Uint64())
				}(content.Cid, content.StartBlock, content.CommitDuration, content.RevealDuration, content.RevealThreshold)

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
				//if d.dkg.GetGroupNumber() == 0 {
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
				groupSize, err := d.chain.GroupSize()
				if err != nil {
					logger.Error(err)
					continue
				}

				f := map[string]interface{}{
					"WorkingGroupSize":   workingGroup,
					"GroupToPick":        groupToPick,
					"currentBlockNumber": currentBlockNumber,
					"PendingNodeSize":    pendingNodeSize,
					"PendingGrouSize":    pendingGrouSize,
					"Members":            d.p.Members()}
				logger.Event("DOS_Watchdog", f)
				fmt.Println("watchdog ", f, " len(ids) ", len(ids))
				//Let pendingNode be a guardian
				if len(ids) == 0 {
					//Signal random if there are enough working groups
					if workingGroup >= groupToPick {
						diff := currentBlockNumber - lastUpdatedBlock
						if diff > SYSRANDOMNTERVAL {
							ctx, _ := context.WithCancel(context.Background())
							d.chain.SignalRandom(ctx)
						}
					}
				}

				if pendingNodeSize >= groupSize+(groupSize/2) {
					d.chain.SignalGroupFormation(context.Background())
				}

			case pubkey := <-d.cRequestDone:
				gID := dkg.HashPoint(pubkey)
				fmt.Println("cRequestDone ", []byte(gID), requestTracking[gID]["count"])
				requestTracking[gID]["count"]--
				if requestTracking[gID]["count"] == 0 &&
					requestTracking[gID]["status"] == GROUPDELETED {
					f := map[string]interface{}{
						"DispatchedGroup_1": fmt.Sprintf("%x", pubkey[0].Bytes()),
						"DispatchedGroup_2": fmt.Sprintf("%x", pubkey[1].Bytes()),
						"DispatchedGroup_3": fmt.Sprintf("%x", pubkey[2].Bytes()),
						"DispatchedGroup_4": fmt.Sprintf("%x", pubkey[3].Bytes()),
					}
					logger.Event("DOS_GroupDeleted", f)
					if d.dkg.GroupDismiss(pubkey) {
						ids := d.dkg.GetGroupIDs(pubkey)
						for _, id := range ids {
							if r := bytes.Compare(d.id, id); r != 0 {
								d.p.DisConnectTo(id)
							}
						}
						ctx, _ := context.WithCancel(context.Background())
						go func() {
							errc := d.chain.RegisterNewNode(ctx)
							err := <-errc
							if err != nil {
								fmt.Println("RegisterNewNode err ", err)
							}
						}()
					}
				}
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
				content, ok := msg.(*onchain.DosproxyLogGrouping)
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

					if !content.Removed {
						f := map[string]interface{}{
							"SessionID": sessionId,
							"Removed":   content.Removed,
							"Tx":        content.Tx,
							"CurBlkN":   currentBlockNumber,
							"BlockN":    content.BlockN}
						logger.Event("DOS_Grouping", f)
						fmt.Println("eventGrouping groupid ", content.GroupId.String())

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
			case msg, ok := <-eventGroupDissolve:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DosproxyLogGroupDissolve)
				if !ok {
					e, ok := msg.(error)
					if ok {
						logger.Error(e)
					}
					continue
				}
				if d.isMember(content.PubKey) {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					gID := dkg.HashPoint(content.PubKey)
					fmt.Println("eventGroupDismiss ", []byte(gID), requestTracking[gID]["count"])

					f := map[string]interface{}{
						"SessionID":         fmt.Sprintf("%x", content.GroupId),
						"DispatchedGroup_1": fmt.Sprintf("%x", content.PubKey[0].Bytes()),
						"DispatchedGroup_2": fmt.Sprintf("%x", content.PubKey[1].Bytes()),
						"DispatchedGroup_3": fmt.Sprintf("%x", content.PubKey[2].Bytes()),
						"DispatchedGroup_4": fmt.Sprintf("%x", content.PubKey[3].Bytes()),
						"ReuqestCount":      requestTracking[gID]["count"],
						"Removed":           content.Removed,
						"Tx":                content.Tx,
						"CurBlkN":           currentBlockNumber,
						"BlockN":            content.BlockN}
					logger.Event("DOS_GroupDismiss", f)
					if requestTracking[gID]["count"] == 0 {
						if d.dkg.GroupDismiss(content.PubKey) {
							ctx, _ := context.WithCancel(context.Background())
							go func() {
								errc := d.chain.RegisterNewNode(ctx)
								err := <-errc
								if err != nil {
									fmt.Println("RegisterNewNode err ", err)
								}
							}()
						}
					} else {
						requestTracking[gID]["status"] = GROUPDELETED
					}
				}
			case msg, ok := <-chRandom:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DosproxyLogUpdateRandom)
				if !ok {
					log.Error(err)
				}
				if d.isMember(content.DispatchedGroup) {
					if !content.Removed {
						gID := dkg.HashPoint(content.DispatchedGroup)
						if requestTracking[gID] == nil {
							requestTracking[gID] = map[string]int{}
							requestTracking[gID]["status"] = GROUPINITIATED
							requestTracking[gID]["count"] = 0
						}
						requestTracking[gID]["count"]++
						fmt.Println("chRandom ", []byte(gID), requestTracking[gID]["count"])

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
							"ReuqestCount":         requestTracking[gID]["count"],
							"Tx":                   content.Tx,
							"CurBlkN":              currentBlockNumber,
							"BlockN":               content.BlockN}
						logger.Event("DOS_QuerySysRandom", f)

						requestID := content.LastRandomness.String()
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
				content, ok := msg.(*onchain.DosproxyLogRequestUserRandom)
				if !ok {
					log.Error(err)
					continue
				}
				if d.isMember(content.DispatchedGroup) {
					if !content.Removed {
						gID := dkg.HashPoint(content.DispatchedGroup)
						fmt.Println("chUsrRandom ", []byte(gID))
						if requestTracking[gID] == nil {
							requestTracking[gID] = map[string]int{}
							requestTracking[gID]["status"] = GROUPINITIATED
							requestTracking[gID]["count"] = 0
						}
						requestTracking[gID]["count"]++
						fmt.Println("chUsrRandom ", []byte(gID), requestTracking[gID]["count"])

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
							"ReuqestCount":         requestTracking[gID]["count"],
							"Tx":                   content.Tx,
							"CurBlkN":              currentBlockNumber,
							"BlockN":               content.BlockN}
						logger.Event("DOS_QueryUserRandom", f)
						requestID := content.RequestId.String()
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
				content, ok := msg.(*onchain.DosproxyLogUrl)
				if !ok {
					log.Error(err)
					continue
				}
				if d.isMember(content.DispatchedGroup) {
					if !content.Removed {
						gID := dkg.HashPoint(content.DispatchedGroup)
						fmt.Println("chUrl ", []byte(gID))
						if requestTracking[gID] == nil {
							requestTracking[gID] = map[string]int{}
							requestTracking[gID]["status"] = GROUPINITIATED
							requestTracking[gID]["count"] = 0
						}
						requestTracking[gID]["count"]++
						fmt.Println("chUrl ", []byte(gID), requestTracking[gID]["count"])

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
							"ReuqestCount":      requestTracking[gID]["count"],
							"Tx":                content.Tx,
							"CurBlkN":           currentBlockNumber,
							"BlockN":            content.BlockN}
						logger.Event("DOS_QueryURL", f)
						requestID := content.QueryId.String()
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
				content, ok := msg.(*onchain.DosproxyLogValidationResult)
				if !ok {
					log.Error(err)
					continue
				}
				if d.isMember(content.PubKey) {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					event := ""
					if content.TrafficType == onchain.TrafficUserQuery {
						event = "DOS_UrlResult"
					} else if content.TrafficType == onchain.TrafficUserRandom {
						event = "DOS_UserRandomResult"
					} else if content.TrafficType == onchain.TrafficSystemRandom {
						event = "DOS_SysRandomResult"
					}

					f := map[string]interface{}{
						"RequestId":      fmt.Sprintf("%x", content.TrafficId),
						"ValidationPass": content.Pass,
						"Message":        fmt.Sprintf("%x", content.Message),
						"Signature_1":    fmt.Sprintf("%x", content.Signature[0].Bytes()),
						"Signature_2":    fmt.Sprintf("%x", content.Signature[1].Bytes()),
						"PubKey_1":       fmt.Sprintf("%x", content.PubKey[0].Bytes()),
						"PubKey_2":       fmt.Sprintf("%x", content.PubKey[1].Bytes()),
						"PubKey_3":       fmt.Sprintf("%x", content.PubKey[2].Bytes()),
						"PubKey_4":       fmt.Sprintf("%x", content.PubKey[3].Bytes()),
						"Tx":             content.Tx,
						"CurBlkN":        currentBlockNumber,
						"BlockN":         content.BlockN}
					logger.Event(event, f)
				}
			case msg, ok := <-keySuggested:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DosproxyLogPublicKeySuggested)
				if !ok {
					e, ok := msg.(error)
					if ok {
						logger.Error(e)
					}
					continue
				}
				if d.isMember(content.PubKey) {
					logger.TimeTrack(time.Now(), "keySuggested", map[string]interface{}{
						"SessionID": d.dkg.GetSessionID(content.PubKey),
						"GroupSize": content.GroupSize,
						"Count":     content.Count,
						"Removed":   content.Removed,
						"Tx":        content.Tx,
						"BlockN":    content.BlockN,
					})
				}
			case msg, ok := <-keyAccepted:
				if !ok {
					continue
				}
				content, ok := msg.(*onchain.DosproxyLogPublicKeyAccepted)
				if !ok {
					e, ok := msg.(error)
					if ok {
						logger.Error(e)
					}
					continue
				}
				if d.isMember(content.PubKey) {
					logger.TimeTrack(time.Now(), "keyAccepted", map[string]interface{}{
						"SessionID":        d.dkg.GetSessionID(content.PubKey),
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
				content, ok := msg.(*onchain.DosproxyLogNoWorkingGroup)
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
					"BlockN":  content.BlockN}
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
