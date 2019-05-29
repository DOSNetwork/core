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
	watchdogInterval    = 10 //In minutes
	ContextKeyGroupID   = contextKey("GroupID")
	ContextKeyRequestID = contextKey("RequestID")
)

type contextKey string

func (c contextKey) String() string {
	return "dosnode " + string(c)
}

type DosNode struct {
	suite        suites.Suite
	chain        onchain.ProxyAdapter
	dkg          dkg.PDKGInterface
	p            p2p.P2PInterface
	done         chan interface{}
	cSignToPeer  chan *vss.Signature
	cRequestDone chan [4]*big.Int
	id           []byte
	logger       log.Logger
	//For REST API
	startTime         time.Time
	state             string
	totalQuery        int
	fulfilledQuery    int
	numOfworkingGroup int
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
	bootstrapIP := config.BootStrapIp

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
	fmt.Println("Join :", bootstrapIP)
	retry, num := 0, 0
	for {
		num, err = p.Join(bootstrapIP)
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
	p2pDkg := dkg.NewPDKG(p, suite)

	dosNode = &DosNode{
		suite:             suite,
		p:                 p,
		chain:             chainConn,
		dkg:               p2pDkg,
		done:              make(chan interface{}),
		cSignToPeer:       make(chan *vss.Signature, 21),
		cRequestDone:      make(chan [4]*big.Int),
		id:                id,
		logger:            log.New("module", "dosclient"),
		startTime:         time.Now(),
		state:             "Init Done",
		totalQuery:        0,
		fulfilledQuery:    0,
		numOfworkingGroup: 0,
	}

	return dosNode, nil
}

func (d *DosNode) Start() (err error) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", d.status)
	mux.HandleFunc("/balance", d.balance)
	mux.HandleFunc("/groupSize", d.groupSize)
	mux.HandleFunc("/proxy", d.proxy)
	mux.HandleFunc("/guardian", d.guardian)
	mux.HandleFunc("/p2p", d.p2p)
	mux.HandleFunc("/signalRandom", d.signalRandom)

	go http.ListenAndServe(":8080", mux)

	d.state = "connecting and syncing geth node"
	d.chain.AddEventNode()

	errc := d.chain.RegisterNewNode(context.Background())
	err = <-errc
	if err != nil {
		d.state = "RegisterNewNode failed"
		fmt.Println("RegisterNewNode failed ,", err.Error())
		d.logger.Error(err)
		return
	}

	if err = d.listen(); err != nil {
		fmt.Println("listen err ", err)
		d.logger.Error(err)
		return
	}
	d.state = "running"

	return
}

func (d *DosNode) End() {
	close(d.done)
}

func (d *DosNode) waitForGrouping(ctx context.Context, groupID string, errc <-chan error) {
	defer d.logger.TimeTrack(time.Now(), "waitForGrouping", map[string]interface{}{"GroupID": groupID})

	for err := range errc {
		if err != nil {
			d.logger.Event("waitForGroupingError", map[string]interface{}{"Error": err.Error(), "GroupID": groupID})

		}
	}
}

func (d *DosNode) waitForRequestDone(ctx context.Context, groupID string, requestID *big.Int, errc <-chan error) {
	defer d.logger.TimeTrack(time.Now(), "WaitForRequestDone", map[string]interface{}{"GroupID": groupID, "RequestID": ctx.Value("RequestID")})
	for err := range errc {
		if err != nil {
			d.logger.Event("waitForRequestError", map[string]interface{}{"Error": err.Error(), "GroupID": groupID, "RequestID": ctx.Value("RequestID")})
		}
	}
}

func (d *DosNode) buildPipeline(valueCtx context.Context, groupID string, requestID, lastRand, useSeed *big.Int, url, selector string, pType uint32) {
	defer d.logger.TimeTrack(time.Now(), "BuildPipeline", map[string]interface{}{"GroupID": groupID, "RequestID": valueCtx.Value("RequestID")})
	d.totalQuery++

	var nonce []byte
	ids := d.dkg.GetGroupIDs(groupID)
	fmt.Println("GetGroupIDs ", len(ids), " ids = ", ids)
	if len(ids) == 0 {
		d.logger.Event("EuildPipeError", map[string]interface{}{"GroupID": groupID, "RequestId": fmt.Sprintf("%x", requestID)})
		return
	}
	pubPoly := d.dkg.GetGroupPublicPoly(groupID)

	//Generate an unique id
	switch pType {
	case onchain.TrafficSystemRandom:
		var bytes []byte
		bytes = append(bytes, []byte(groupID)...)
		bytes = append(bytes, requestID.Bytes()...)
		bytes = append(bytes, lastRand.Bytes()...)
		nHash := sha256.Sum256(bytes)
		nonce = nHash[:]
	case onchain.TrafficUserRandom:
		var bytes []byte
		bytes = append(bytes, []byte(groupID)...)
		bytes = append(bytes, requestID.Bytes()...)
		bytes = append(bytes, lastRand.Bytes()...)
		bytes = append(bytes, useSeed.Bytes()...)
		nHash := sha256.Sum256(bytes)
		nonce = nHash[:]
	case onchain.TrafficUserQuery:
		var bytes []byte
		bytes = append(bytes, []byte(groupID)...)
		bytes = append(bytes, requestID.Bytes()...)
		bytes = append(bytes, lastRand.Bytes()...)
		bytes = append(bytes, []byte(url)...)
		bytes = append(bytes, []byte(selector)...)
		nHash := sha256.Sum256(bytes)
		nonce = nHash[:]
	}

	//Build a pipeline
	var signShares []<-chan *vss.Signature
	var errcList []<-chan error

	submitterc, errc := choseSubmitter(valueCtx, d.p, lastRand, ids, len(ids), d.logger)
	if len(submitterc) != len(ids) || len(ids) == 0 {
		d.logger.Event("EuildPipeError2", map[string]interface{}{"GroupID": groupID, "RequestId": fmt.Sprintf("%x", requestID), "lenSubmitter": len(submitterc)})
		return
	}
	errcList = append(errcList, errc)

	var contentc <-chan []byte
	switch pType {
	case onchain.TrafficSystemRandom:
		contentc = genSysRandom(valueCtx, submitterc[0], lastRand.Bytes(), d.logger)
	case onchain.TrafficUserRandom:
		contentc = genUserRandom(valueCtx, submitterc[0], requestID.Bytes(), lastRand.Bytes(), useSeed.Bytes(), d.logger)
	case onchain.TrafficUserQuery:
		contentc, errc = genQueryResult(valueCtx, submitterc[0], url, selector, d.logger)
		errcList = append(errcList, errc)
	}

	signc, errc := genSign(valueCtx, contentc, d.cSignToPeer, d.dkg.GetShareSecurity(groupID), d.suite, d.id, groupID, requestID.Bytes(), pType, nonce, d.logger)
	errcList = append(errcList, errc)
	signShares = append(signShares, signc)

	idx := 1
	for _, id := range ids {
		if r := bytes.Compare(d.id, id); r != 0 {
			signc, errc := requestSign(valueCtx, submitterc[idx], contentc, d.p, d.id, requestID.Bytes(), pType, id, nonce, d.logger)
			signShares = append(signShares, signc)
			errcList = append(errcList, errc)
			idx++
		}
	}

	recoveredSignc, errc := recoverSign(valueCtx, fanIn(valueCtx, signShares...), d.suite, pubPoly, (len(ids)/2 + 1), len(ids), d.logger)
	errcList = append(errcList, errc)

	switch pType {
	case onchain.TrafficSystemRandom:
		errc := d.chain.SetRandomNum(valueCtx, recoveredSignc)
		errcList = append(errcList, errc)
	default:
		errc := d.chain.DataReturn(valueCtx, recoveredSignc)
		errcList = append(errcList, errc)
	}

	go d.waitForRequestDone(valueCtx, groupID, requestID, mergeErrors(valueCtx, errcList...))
}

func (d *DosNode) listen() (err error) {

	var errcList []<-chan error
	eventGrouping, errc := d.chain.SubscribeEvent(onchain.SubscribeLogGrouping)
	errcList = append(errcList, errc)
	eventGroupDissolve, errc := d.chain.SubscribeEvent(onchain.SubscribeLogGroupDissolve)
	errcList = append(errcList, errc)
	chURL, errc := d.chain.SubscribeEvent(onchain.SubscribeLogUrl)
	errcList = append(errcList, errc)
	chRandom, errc := d.chain.SubscribeEvent(onchain.SubscribeLogUpdateRandom)
	errcList = append(errcList, errc)
	chUsrRandom, errc := d.chain.SubscribeEvent(onchain.SubscribeLogRequestUserRandom)
	errcList = append(errcList, errc)
	eventValidation, errc := d.chain.SubscribeEvent(onchain.SubscribeLogValidationResult)
	errcList = append(errcList, errc)
	keyAccepted, errc := d.chain.SubscribeEvent(onchain.SubscribeLogPublicKeyAccepted)
	errcList = append(errcList, errc)
	keySuggested, errc := d.chain.SubscribeEvent(onchain.SubscribeLogPublicKeySuggested)
	errcList = append(errcList, errc)
	commitRevealStart, errc := d.chain.SubscribeEvent(onchain.SubscribeCommitrevealLogStartCommitreveal)
	errcList = append(errcList, errc)

	peerEvent, err := d.p.SubscribeEvent(50, vss.Signature{})
	errcList = append(errcList, errc)
	errc = mergeErrors(context.Background(), errcList...)

	peerSignMap := make(map[string]*vss.Signature)
	watchdog := time.NewTicker(watchdogInterval * time.Minute)
	//	latestRandm := big.NewInt(0)
	defer watchdog.Stop()
	defer d.p.UnSubscribeEvent(vss.Signature{})

	for {
		select {
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
				d.logger.Event("commitRevealStart", f)

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
						d.logger.Error(err)
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
					fmt.Println("Waiting for commit err", err)
					return
				}
				for {
					cur, err := d.chain.CurrentBlock()
					if err != nil {
						fmt.Println("CurrentBlock err", err)

						d.logger.Error(err)
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
						fmt.Println("CurrentBlock err", err)

						d.logger.Error(err)
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
			d.logger.Error(err)
			/*
				TODO: subscribe again and push err to errc,
				TODO: otherwise possibly closed errc will keep consuming "select",
			*/
		case <-watchdog.C:

			//Let pending node as a guardian
			isPendingNode, err := d.chain.IsPendingNode(d.id)
			if err != nil {
				continue
			}
			if isPendingNode {
				fmt.Println("watchdog")
				currentBlockNumber, err := d.chain.CurrentBlock()
				if err != nil {
					d.logger.Error(err)
				}

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

				lastUpdatedBlock, err := d.chain.LastUpdatedBlock()
				if err != nil {
					d.logger.Error(err)
					continue
				}
				groupSize, err := d.chain.GroupSize()
				if err != nil {
					d.logger.Error(err)
					continue
				}
				expiredWGSize, err := d.chain.GetExpiredWorkingGroupSize()
				if err != nil {
					d.logger.Error(err)
					continue
				}
				sysrandInterval, err := d.chain.RefreshSystemRandomHardLimit()
				if err != nil {
					d.logger.Error(err)
					continue
				}
				f := map[string]interface{}{
					"WorkingGroupSize":    workingGroup,
					"GroupToPick":         groupToPick,
					"currentBlockNumber":  currentBlockNumber,
					"PendingNodeSize":     pendingNodeSize,
					"PendingGrouSize":     pendingGrouSize,
					"expiredWorkingGSize": expiredWGSize,
					"Members":             d.p.Members(),
					"sysrandInterval":     sysrandInterval}
				d.logger.Event("DWatchdog", f)
				//Let pendingNode be a guardian
				//Signal random if there are enough working groups
				if workingGroup >= groupToPick {
					diff := currentBlockNumber - lastUpdatedBlock
					if diff > sysrandInterval {
						d.chain.SignalRandom(context.Background())
					}
				}

				if pendingNodeSize >= groupSize+(groupSize/2) {
					d.chain.SignalGroupFormation(context.Background())
				}

				if expiredWGSize > 0 || pendingGrouSize > 0 {
					d.chain.SignalGroupDissolve(context.Background())
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
			content, ok := msg.(*onchain.LogGrouping)
			if !ok {
				log.Error(err)
				continue
			}
			isMember := false
			var participants [][]byte
			for _, p := range content.NodeId {
				id := p.Bytes()
				if r := bytes.Compare(d.id, id); r == 0 {
					isMember = true
				}
				participants = append(participants, id)
			}
			groupID := fmt.Sprintf("%x", content.GroupId)
			f := map[string]interface{}{
				"GroupID": groupID,
				"Removed": content.Removed,
				"Tx":      content.Tx}
			d.logger.Event("DGrouping1", f)
			if isMember {
				d.logger.Event("DGrouping2", f)
				//if !content.Removed {
				valueCtx := context.WithValue(context.Background(), ContextKeyGroupID, groupID)

				outFromDkg, errc, err := d.dkg.Grouping(valueCtx, groupID, participants)
				if err != nil {
					d.logger.Error(err)
					continue
				}
				//errcList = append(errcList, errc)
				_ = d.chain.RegisterGroupPubKey(valueCtx, outFromDkg)
				//errcList = append(errcList, errc)
				//errc = mergeErrors(valueCtx, errcList...)
				go d.waitForGrouping(valueCtx, groupID, errc)
				//}
			}
		case msg, ok := <-eventGroupDissolve:
			if !ok {
				continue
			}
			content, ok := msg.(*onchain.LogGroupDissolve)
			if !ok {
				e, ok := msg.(error)
				if ok {
					d.logger.Error(e)
				}
				continue
			}
			groupID := fmt.Sprintf("%x", content.GroupId)
			f := map[string]interface{}{
				"GroupID": fmt.Sprintf("%x", content.GroupId)}
			d.logger.Event("DGroupDismiss1", f)
			if d.isMember(groupID) {
				d.dkg.GroupDissolve(groupID)
				d.logger.Event("DGroupDismiss2", f)
			}
		case msg, ok := <-chRandom:
			if !ok {
				continue
			}
			content, ok := msg.(*onchain.LogUpdateRandom)
			if !ok {
				log.Error(err)
			}
			//			latestRandm = content.LastRandomness
			if d.isMember(fmt.Sprintf("%x", content.DispatchedGroupId)) {
				//	if !content.Removed {
				currentBlockNumber, err := d.chain.CurrentBlock()
				if err != nil {
					d.logger.Error(err)
				}
				requestID := fmt.Sprintf("%x", content.LastRandomness)
				groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
				lastRand := fmt.Sprintf("%x", content.LastRandomness)
				f := map[string]interface{}{
					"RequestId":            requestID,
					"GroupID":              groupID,
					"LastSystemRandomness": lastRand,
					"Tx":                   content.Tx,
					"CurBlkN":              currentBlockNumber,
					"BlockN":               content.BlockN}
				d.logger.Event("DOS_QuerySysRandom", f)

				valueCtx := context.WithValue(context.Background(), ContextKeyRequestID, requestID)
				valueCtx = context.WithValue(valueCtx, ContextKeyGroupID, groupID)

				d.buildPipeline(valueCtx, groupID, content.LastRandomness, content.LastRandomness, nil, "", "", uint32(onchain.TrafficSystemRandom))
				//}
			}
		case msg, ok := <-chUsrRandom:
			if !ok {
				continue
			}
			content, ok := msg.(*onchain.LogRequestUserRandom)
			if !ok {
				log.Error(err)
				continue
			}
			if d.isMember(fmt.Sprintf("%x", content.DispatchedGroupId)) {
				currentBlockNumber, err := d.chain.CurrentBlock()
				if err != nil {
					d.logger.Error(err)
				}
				requestID := fmt.Sprintf("%x", content.RequestId)
				groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
				lastRand := fmt.Sprintf("%x", content.LastSystemRandomness)
				f := map[string]interface{}{
					"RequestId":            requestID,
					"GroupID":              groupID,
					"LastSystemRandomness": lastRand,
					"Tx":                   content.Tx,
					"CurBlkN":              currentBlockNumber,
					"BlockN":               content.BlockN}
				d.logger.Event("DOS_QueryUserRandom", f)
				valueCtx := context.WithValue(context.Background(), ContextKeyRequestID, requestID)
				valueCtx = context.WithValue(valueCtx, ContextKeyGroupID, groupID)

				d.buildPipeline(valueCtx, groupID, content.RequestId, content.LastSystemRandomness, content.UserSeed, "", "", uint32(onchain.TrafficUserRandom))

			}
		case msg, ok := <-chURL:
			if !ok {
				continue
			}
			content, ok := msg.(*onchain.LogUrl)
			if !ok {
				log.Error(err)
				continue
			}
			if d.isMember(fmt.Sprintf("%x", content.DispatchedGroupId)) {
				//if !content.Removed {
				currentBlockNumber, err := d.chain.CurrentBlock()
				if err != nil {
					d.logger.Error(err)
				}
				startTime := time.Now()
				fmt.Println("set up time ", startTime)
				requestID := fmt.Sprintf("%x", content.QueryId)
				groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
				rand := fmt.Sprintf("%x", content.Randomness)
				f := map[string]interface{}{
					"RequestId":  requestID,
					"Randomness": rand,
					"DataSource": content.DataSource,
					"GroupID":    groupID,
					"Tx":         content.Tx,
					"CurBlkN":    currentBlockNumber,
					"BlockN":     content.BlockN}
				d.logger.Event("DOS_QueryURL", f)
				valueCtx := context.WithValue(context.Background(), ContextKeyRequestID, requestID)
				valueCtx = context.WithValue(valueCtx, ContextKeyGroupID, groupID)

				d.buildPipeline(valueCtx, groupID, content.QueryId, content.Randomness, nil, content.DataSource, content.Selector, uint32(onchain.TrafficUserQuery))
				//}
			}
		case _, ok := <-eventValidation:
			if !ok {
				continue
			}
		case msg, ok := <-keySuggested:
			if !ok {
				continue
			}
			content, ok := msg.(*onchain.LogPublicKeySuggested)
			if !ok {
				e, ok := msg.(error)
				if ok {
					d.logger.Error(e)
				}
				continue
			}
			if d.isMember(fmt.Sprintf("%x", content.GroupId)) {
				d.logger.TimeTrack(time.Now(), "keySuggested", map[string]interface{}{
					"GroupID": fmt.Sprintf("%x", content.GroupId.Bytes()),
					"Count":   content.Count,
					"Removed": content.Removed,
					"Tx":      content.Tx,
					"BlockN":  content.BlockN,
				})
			}
		case msg, ok := <-keyAccepted:
			if !ok {
				continue
			}
			content, ok := msg.(*onchain.LogPublicKeyAccepted)
			if !ok {
				e, ok := msg.(error)
				if ok {
					d.logger.Error(e)
				}
				continue
			}
			if d.isMember(fmt.Sprintf("%x", content.GroupId)) {
				d.logger.TimeTrack(time.Now(), "keyAccepted", map[string]interface{}{
					"GroupId":          fmt.Sprintf("%x", content.GroupId.Bytes()),
					"workingGroupSize": content.WorkingGroupSize,
					"Removed":          content.Removed,
					"Tx":               content.Tx,
					"BlockN":           content.BlockN,
				})
			}
		case <-d.done:
			return
		}
	}
}

func (d *DosNode) isMember(groupID string) bool {
	return d.dkg.GetShareSecurity(groupID) != nil
}

func byte32(s []byte) (a *[32]byte) {
	if len(a) <= len(s) {
		a = (*[len(a)]byte)(unsafe.Pointer(&s[0]))
	}
	return a
}
