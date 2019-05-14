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
	dkg          dkg.PDKGInterface
	p            p2p.P2PInterface
	done         chan interface{}
	cSignToPeer  chan *vss.Signature
	cRequestDone chan [4]*big.Int
	id           []byte
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
	p2pDkg := dkg.NewPDKG(p, suite)

	//if logger == nil {
	logger = log.New("module", "dosclient")
	//}
	dosNode = &DosNode{
		suite:             suite,
		p:                 p,
		chain:             chainConn,
		dkg:               p2pDkg,
		done:              make(chan interface{}),
		cSignToPeer:       make(chan *vss.Signature, 21),
		cRequestDone:      make(chan [4]*big.Int),
		id:                id,
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

	ctx, _ := context.WithCancel(context.Background())
	d.state = "connecting and syncing geth node"
	d.chain.AddEventNode()

	_ = d.chain.RegisterNewNode(ctx)

	if err = d.listen(); err != nil {
		fmt.Println("listen err ", err)
		logger.Error(err)
		return
	}
	return
}

func (d *DosNode) End() {
	close(d.done)
}

func (d *DosNode) waitForGrouping(ctx context.Context, groupID string, errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "waitForGrouping", map[string]interface{}{"GroupID": groupID})

	for err := range errc {
		if err != nil {
			logger.Error(err)
		}
	}
}

func (d *DosNode) waitForRequestDone(ctx context.Context, groupID string, requestID *big.Int, errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "WaitForRequestDone", map[string]interface{}{"GroupID": groupID, "RequestID": ctx.Value("RequestID")})
	for err := range errc {
		if err != nil {
			logger.Error(err)
		}
	}
}

func (d *DosNode) buildPipeline(valueCtx context.Context, groupID string, requestID, lastRand, useSeed *big.Int, url, selector string, pType uint32) {
	defer logger.TimeTrack(time.Now(), "BuildPipeline", map[string]interface{}{"GroupID": groupID, "RequestID": valueCtx.Value("RequestID")})
	d.totalQuery++
	var signShares []<-chan *vss.Signature
	var errcList []<-chan error
	var cSubmitter []chan []byte
	var cErr <-chan error
	var cSign <-chan *vss.Signature
	var cContent <-chan []byte
	var nonce []byte
	ids := d.dkg.GetGroupIDs(groupID)
	fmt.Println("GetGroupIDs ", len(ids), " ids = ", ids)
	if len(ids) == 0 {
		logger.Event("EuildPipeError", map[string]interface{}{"GroupID": groupID, "RequestId": fmt.Sprintf("%x", requestID)})
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
	cSubmitter, cErr = choseSubmitter(valueCtx, d.chain, lastRand, ids, len(ids))
	if len(cSubmitter) != len(ids) || len(ids) == 0 {
		logger.Event("EuildPipeError2", map[string]interface{}{"GroupID": groupID, "RequestId": fmt.Sprintf("%x", requestID, "lenSubmitter", len(cSubmitter))})
		return
	}
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

	cSign, cErr = genSign(valueCtx, cContent, d.cSignToPeer, d.dkg.GetShareSecurity(groupID), d.suite, d.id, groupID, requestID.Bytes(), pType, nonce)
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

	go d.waitForRequestDone(valueCtx, groupID, requestID, errc)

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
	commitRevealStart, errc := d.chain.SubscribeEvent(onchain.SubscribeCommitrevealLogStartCommitreveal)
	errcList = append(errcList, errc)

	peerEvent, err := d.p.SubscribeEvent(50, vss.Signature{})
	errcList = append(errcList, errc)
	errc = mergeErrors(context.Background(), errcList...)

	pipeCancel := make(map[string]context.CancelFunc)
	peerSignMap := make(map[string]*vss.Signature)
	watchdog := time.NewTicker(WATCHDOGINTERVAL * time.Minute)
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
					fmt.Println("Waiting for commit err", err)
					return
				}
				for {
					cur, err := d.chain.CurrentBlock()
					if err != nil {
						fmt.Println("CurrentBlock err", err)

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
						fmt.Println("CurrentBlock err", err)

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
			if d.dkg.GetGroupNumber() == 0 {
				fmt.Println("watchdog")
				currentBlockNumber, err := d.chain.CurrentBlock()
				if err != nil {
					logger.Error(err)
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
				logger.Event("DWatchdog", f)
				//Let pendingNode be a guardian
				//Signal random if there are enough working groups
				if workingGroup >= groupToPick {
					diff := currentBlockNumber - lastUpdatedBlock
					if diff > SYSRANDOMNTERVAL {
						ctx, _ := context.WithCancel(context.Background())
						d.chain.SignalRandom(ctx)
					}
				}

				if pendingNodeSize >= groupSize+(groupSize/2) {
					d.chain.SignalGroupFormation(context.Background())
				}
				d.chain.SignalDissolve(context.Background())
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
				"Tx":      content.Tx}
			logger.Event("DGrouping1", f)
			if isMember {

				logger.Event("DGrouping2", f)
				if !content.Removed {
					ctx, cancelFunc := context.WithCancel(context.Background())
					valueCtx := context.WithValue(ctx, "GroupID", groupID)
					pipeCancel[groupID] = cancelFunc

					var errcList []<-chan error
					outFromDkg, errc, err := d.dkg.Grouping(valueCtx, groupID, participants)
					if err != nil {
						logger.Error(err)
						continue
					}
					errcList = append(errcList, errc)
					errc = d.chain.RegisterGroupPubKey(valueCtx, outFromDkg)
					errcList = append(errcList, errc)
					errc = mergeErrors(valueCtx, errcList...)
					go d.waitForGrouping(valueCtx, groupID, errc)
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
			groupID := fmt.Sprintf("%x", content.GroupId)
			f := map[string]interface{}{
				"GroupID": fmt.Sprintf("%x", content.GroupId)}
			logger.Event("DGroupDismiss1", f)
			if d.isMember(groupID) {
				d.dkg.GroupDissolve(groupID)
				logger.Event("DGroupDismiss2", f)
			}
		case msg, ok := <-chRandom:
			if !ok {
				continue
			}
			content, ok := msg.(*onchain.DosproxyLogUpdateRandom)
			if !ok {
				log.Error(err)
			}
			//			latestRandm = content.LastRandomness
			if d.isMember(fmt.Sprintf("%x", content.DispatchedGroupId)) {
				if !content.Removed {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					requestID := fmt.Sprintf("%x", content.LastRandomness)
					groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
					lastRand := fmt.Sprintf("%x", content.LastRandomness)
					f := map[string]interface{}{
						"RequestId":            requestID,
						"GroupID":              groupID,
						"LastSystemRandomness": lastRand,
						"Tx":      content.Tx,
						"CurBlkN": currentBlockNumber,
						"BlockN":  content.BlockN}
					logger.Event("DOS_QuerySysRandom", f)

					ctx, cancelFunc := context.WithCancel(context.Background())
					valueCtx := context.WithValue(ctx, "RequestID", requestID)
					valueCtx = context.WithValue(valueCtx, "GroupID", groupID)

					pipeCancel[requestID] = cancelFunc
					d.buildPipeline(valueCtx, groupID, content.LastRandomness, content.LastRandomness, nil, "", "", uint32(onchain.TrafficSystemRandom))
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
			if d.isMember(fmt.Sprintf("%x", content.DispatchedGroupId)) {
				if !content.Removed {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
					}
					requestID := fmt.Sprintf("%x", content.RequestId)
					groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
					lastRand := fmt.Sprintf("%x", content.LastSystemRandomness)
					f := map[string]interface{}{
						"RequestId":            requestID,
						"GroupID":              groupID,
						"LastSystemRandomness": lastRand,
						"Tx":      content.Tx,
						"CurBlkN": currentBlockNumber,
						"BlockN":  content.BlockN}
					logger.Event("DOS_QueryUserRandom", f)
					ctx, _ := context.WithCancel(context.Background())
					valueCtx := context.WithValue(ctx, "RequestID", requestID)
					valueCtx = context.WithValue(valueCtx, "GroupID", groupID)

					d.buildPipeline(valueCtx, groupID, content.RequestId, content.LastSystemRandomness, content.UserSeed, "", "", uint32(onchain.TrafficUserRandom))
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
			if d.isMember(fmt.Sprintf("%x", content.DispatchedGroupId)) {
				if !content.Removed {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						logger.Error(err)
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
					logger.Event("DOS_QueryURL", f)
					ctx, _ := context.WithCancel(context.Background())
					valueCtx := context.WithValue(ctx, "RequestID", requestID)
					valueCtx = context.WithValue(valueCtx, "GroupID", groupID)

					d.buildPipeline(valueCtx, groupID, content.QueryId, content.Randomness, nil, content.DataSource, content.Selector, uint32(onchain.TrafficUserQuery))
				}
			}
		case _, ok := <-eventValidation:
			if !ok {
				continue
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
			if d.isMember(fmt.Sprintf("%x", content.GroupId)) {
				logger.TimeTrack(time.Now(), "keySuggested", map[string]interface{}{
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
			content, ok := msg.(*onchain.DosproxyLogPublicKeyAccepted)
			if !ok {
				e, ok := msg.(error)
				if ok {
					logger.Error(e)
				}
				continue
			}
			if d.isMember(fmt.Sprintf("%x", content.GroupId)) {
				logger.TimeTrack(time.Now(), "keyAccepted", map[string]interface{}{
					"GroupId":          fmt.Sprintf("%x", content.GroupId.Bytes()),
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
	return
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
