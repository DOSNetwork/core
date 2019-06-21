package dosnode

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto/sha3"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
)

const (
	watchdogInterval = 10 //In minutes
	envPassPhrase    = "PASSPHRASE"
)

type ctxKey string

// DosNode is a strcut that represents a offchain dos client
type DosNode struct {
	suite         suites.Suite
	chain         onchain.ProxyAdapter
	dkg           dkg.PDKGInterface
	p             p2p.P2PInterface
	done          chan interface{}
	cSignToPeer   chan *vss.Signature
	cRequestDone  chan [4]*big.Int
	eventGrouping chan interface{}
	id            []byte
	logger        log.Logger
	//For REST API
	startTime         time.Time
	state             string
	totalQuery        int
	fulfilledQuery    int
	numOfworkingGroup int
}
type crDurations struct {
	cid        *big.Int
	startBlock *big.Int
	commitDur  *big.Int
	revealDur  *big.Int
	sec        *big.Int
}

//NewDosNode creates a DosNode struct
func NewDosNode(key *keystore.Key) (dosNode *DosNode, err error) {

	//Read Configuration
	config := configuration.Config{}
	err = config.LoadConfig()
	if err != nil {
		return
	}

	port := config.Port
	bootstrapIP := config.BootStrapIp

	chainConfig := config.GetChainConfig()

	//Set up an onchain adapter
	chainConn, err := onchain.NewProxyAdapter(config.GetCurrentType(), key, chainConfig.DOSProxyAddress, chainConfig.CommitReveal, chainConfig.RemoteNodeAddressPool)
	if err != nil {
		if err.Error() != "No any working eth client for event tracking" {
			fmt.Println("NewDosNode failed ", err)
			return
		}
	}

	id := key.Address

	//Build a p2p network
	p, err := p2p.CreateP2PNetwork(id.Bytes(), port, p2p.GossipDiscover)
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
		id:                id.Bytes(),
		logger:            log.New("module", "dosclient"),
		startTime:         time.Now(),
		state:             "Init Done",
		totalQuery:        0,
		fulfilledQuery:    0,
		numOfworkingGroup: 0,
	}

	return dosNode, nil
}

// Start registers to onchain and listen to p2p events
func (d *DosNode) Start() (err error) {
	d.startRESTServer()

	d.state = "Working"

	d.listen()
	return
}

//End is an operation that does a graceful shutdown
func (d *DosNode) End() {
	close(d.done)
}

func (d *DosNode) handleQuery(ids [][]byte, pubPoly *share.PubPoly, sec *share.PriShare, groupID string, requestID, lastRand, useSeed *big.Int, url, selector string, pType uint32) {
	queryCtx, cancel := d.chain.GetTimeoutCtx(time.Duration(60 * 15 * time.Second))
	defer cancel()
	queryCtxWithValue := context.WithValue(context.WithValue(queryCtx, ctxKey("RequestID"), fmt.Sprintf("%x", requestID)), ctxKey("GroupID"), groupID)

	defer d.logger.TimeTrack(time.Now(), "TimeHandleQuery", map[string]interface{}{"GroupID": groupID, "RequestID": fmt.Sprintf("%x", requestID)})
	defer cancel()
	var nonce []byte
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
	var signShares []chan *vss.Signature
	var errcList []chan error

	submitterc, errc := choseSubmitter(queryCtxWithValue, d.p, lastRand, ids, len(ids), d.logger)
	if len(submitterc) != len(ids) || len(ids) == 0 {
		d.logger.Event("EuildPipeError2", map[string]interface{}{"GroupID": groupID, "RequestId": fmt.Sprintf("%x", requestID), "lenSubmitter": len(submitterc)})
		return
	}
	errcList = append(errcList, errc)

	var contentc chan []byte
	switch pType {
	case onchain.TrafficSystemRandom:
		contentc = genSysRandom(queryCtxWithValue, submitterc[0], lastRand.Bytes(), d.logger)
	case onchain.TrafficUserRandom:
		contentc = genUserRandom(queryCtxWithValue, submitterc[0], requestID.Bytes(), lastRand.Bytes(), useSeed.Bytes(), d.logger)
	case onchain.TrafficUserQuery:
		contentc, errc = genQueryResult(queryCtxWithValue, submitterc[0], url, selector, d.logger)
		errcList = append(errcList, errc)
	}

	signc, errc := genSign(queryCtxWithValue, contentc, d.cSignToPeer, sec, d.suite, d.id, groupID, requestID.Bytes(), pType, nonce, d.logger)
	errcList = append(errcList, errc)
	signShares = append(signShares, signc)

	idx := 1
	for _, id := range ids {
		if r := bytes.Compare(d.id, id); r != 0 {
			signc, errc := requestSign(queryCtxWithValue, submitterc[idx], contentc, d.p, d.id, requestID.Bytes(), pType, id, nonce, d.logger)
			signShares = append(signShares, signc)
			errcList = append(errcList, errc)
			idx++
		}
	}

	recoveredSignc, errc := recoverSign(queryCtxWithValue, fanIn(queryCtxWithValue, signShares...), d.suite, pubPoly, (len(ids)/2 + 1), len(ids), d.logger)
	errcList = append(errcList, errc)

	switch pType {
	case onchain.TrafficSystemRandom:
		errc := d.chain.SetRandomNum(queryCtxWithValue, recoveredSignc)
		errcList = append(errcList, errc)
	default:
		errc := d.chain.DataReturn(queryCtxWithValue, recoveredSignc)
		errcList = append(errcList, errc)
	}
	allErrc := mergeErrors(queryCtxWithValue, errcList...)
	for {
		select {
		case err, ok := <-allErrc:
			if !ok {
				return
			}
			d.logger.Event("handleQueryError", map[string]interface{}{"Error": err.Error(), "GroupID": groupID})
		case <-queryCtxWithValue.Done():
			d.logger.Event("handleQueryError", map[string]interface{}{"Error": queryCtxWithValue.Err(), "GroupID": groupID})
			return
		}
	}
}

func (d *DosNode) handleGrouping(participants [][]byte, groupID string) {
	isMember := false
	for _, id := range participants {
		if r := bytes.Compare(d.id, id); r == 0 {
			isMember = true
			break
		}
	}
	if !isMember {
		return
	}
	d.logger.Event("Grouping", map[string]interface{}{"GroupID": groupID})
	defer d.logger.TimeTrack(time.Now(), "TimeGrouping", map[string]interface{}{"GroupID": groupID})

	ctx, cancel := d.chain.GetTimeoutCtx(time.Duration(60 * 60 * time.Second))
	defer cancel()

	var errcList []chan error
	outFromDkg, errc, err := d.dkg.Grouping(ctx, groupID, participants)
	if err != nil {
		d.logger.Error(err)
		return
	}
	errcList = append(errcList, errc)
	errcList = append(errcList, d.chain.RegisterGroupPubKey(ctx, outFromDkg))
	allErrc := mergeErrors(ctx, errcList...)
	for {
		select {
		case err, ok := <-allErrc:
			if !ok {
				return
			}
			d.logger.Event("waitForGroupingError", map[string]interface{}{"Error": err.Error(), "GroupID": groupID})
		case <-ctx.Done():
			d.logger.Event("waitForGroupingError", map[string]interface{}{"Error": ctx.Err(), "GroupID": groupID})
			return
		}
	}
}

func (d *DosNode) groupInfo(groupID string) (ids [][]byte, pubPoly *share.PubPoly, sec *share.PriShare, err error) {
	//Get group members id
	ids = d.dkg.GetGroupIDs(groupID)
	//Get group pub key
	pubPoly = d.dkg.GetGroupPublicPoly(groupID)
	//Get group partial sec key
	sec = d.dkg.GetShareSecurity(groupID)
	if len(ids) == 0 || pubPoly == nil || sec == nil {
		err = errors.New("No Group info")
	}
	return
}

func (d *DosNode) handleCR(cr *onchain.LogStartCommitReveal, randSeed *big.Int) {

	ctx, cancel := d.chain.GetTimeoutCtx(time.Duration(160 * 15 * time.Second))
	defer cancel()

	// Generate random numbers in range [0..randSeed]
	sec, err := rand.Int(rand.Reader, randSeed)
	if err != nil {
		d.logger.Error(err)
		return
	}
	h := sha3.NewKeccak256()
	h.Write(abi.U256(sec))
	b := h.Sum(nil)
	hash := byte32(b)

	currentBlockNumber, err := d.chain.CurrentBlock(ctx)
	if err != nil {
		d.logger.Error(err)
		return
	}

	cid := cr.Cid
	waitCommit := cr.StartBlock.Uint64() - currentBlockNumber + 1
	if waitCommit < 0 {
		waitCommit = 0
	}
	go func(cid *big.Int, hash *[32]byte) {
		time.Sleep(time.Duration(waitCommit*15) * time.Second)
		if err := d.chain.Commit(ctx, cid, *hash); err != nil {
			d.logger.Error(err)
		}
	}(cid, hash)

	waitReveal := (cr.StartBlock.Uint64() + cr.CommitDuration.Uint64()) - currentBlockNumber + 1
	if waitReveal < 0 {
		waitReveal = 0
	}
	go func(cid *big.Int, sec *big.Int) {
		time.Sleep(time.Duration(waitReveal*15) * time.Second)
		if err := d.chain.Reveal(ctx, cid, sec); err != nil {
			d.logger.Error(err)
		}
	}(cid, sec)

	waitRandom := (cr.StartBlock.Uint64() + cr.CommitDuration.Uint64() + cr.RevealDuration.Uint64()) - currentBlockNumber + 1
	if waitRandom < 0 {
		waitRandom = 0
	}
	go func(cid *big.Int) {
		time.Sleep(time.Duration(waitRandom*15) * time.Second)
		if err := d.chain.SignalBootstrap(ctx, cid); err != nil {
			d.logger.Error(err)
		}
	}(cid)

}

func (d *DosNode) handleGroupFormation(currentBlockNumber uint64) {

	groupSize, err := d.chain.GroupSize(context.Background())
	if err != nil {
		d.logger.Error(err)
		return
	}
	pendingNodeSize, err := d.chain.NumPendingNodes(context.Background())
	if err != nil {
		d.logger.Error(err)
		return
	}

	if pendingNodeSize >= groupSize+(groupSize/2) {
		d.chain.SignalGroupFormation(context.Background())
	}
}

func (d *DosNode) handleRandom(currentBlockNumber uint64) {

	groupToPick, err := d.chain.GroupToPick(context.Background())
	if err != nil {
		return
	}
	workingGroup, err := d.chain.GetWorkingGroupSize(context.Background())
	if err != nil {
		d.logger.Error(err)
		return
	}
	lastUpdatedBlock, err := d.chain.LastUpdatedBlock(context.Background())
	if err != nil {
		d.logger.Error(err)
		return
	}
	sysrandInterval, err := d.chain.RefreshSystemRandomHardLimit(context.Background())
	if err != nil {
		d.logger.Error(err)
		return
	}
	if workingGroup >= groupToPick {
		diff := currentBlockNumber - lastUpdatedBlock
		if diff > sysrandInterval {
			d.chain.SignalRandom(context.Background())
		}
	}
}

func (d *DosNode) handleGroupDissolve() {
	pendingGrouSize, err := d.chain.NumPendingGroups(context.Background())
	if err != nil {
		d.logger.Error(err)
		return
	}
	expiredWGSize, err := d.chain.GetExpiredWorkingGroupSize(context.Background())
	if err != nil {
		d.logger.Error(err)
		return
	}
	if expiredWGSize > 0 || pendingGrouSize > 0 {
		d.chain.SignalGroupDissolve(context.Background())
	}
}

func (d *DosNode) listen() {
	watchdog := time.NewTicker(watchdogInterval * time.Minute)
	defer watchdog.Stop()
	peerEvent, _ := d.p.SubscribeEvent(50, vss.Signature{})
	peerSignMap := make(map[string]*vss.Signature)
	//	latestRandm := big.NewInt(0)
	defer d.p.UnSubscribeEvent(vss.Signature{})
	subescriptions := []int{onchain.SubscribeLogGrouping, onchain.SubscribeLogGroupDissolve, onchain.SubscribeLogUrl,
		onchain.SubscribeLogUpdateRandom, onchain.SubscribeLogRequestUserRandom,
		onchain.SubscribeLogPublicKeyAccepted, onchain.SubscribeCommitrevealLogStartCommitreveal}
	randSeed, _ := new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	d.chain.Start()
	//TODO: Check to see if it is a valid stacking node first
	_ = d.chain.RegisterNewNode(context.Background())
	fmt.Println("(d *DosNode) listen()")
S:
	sink, errc := d.chain.SubscribeEvent(subescriptions)
L:
	for {
		select {
		case <-watchdog.C:
			if isPendingNode, _ := d.chain.IsPendingNode(context.Background(), d.id); isPendingNode {
				//Let pending node as a guardian
				currentBlockNumber, err := d.chain.CurrentBlock(context.Background())
				if err != nil {
					d.logger.Error(err)
					return
				}
				switch index := currentBlockNumber % 3; index {
				case 0:
					d.handleRandom(currentBlockNumber)
				case 1:
					d.handleGroupFormation(currentBlockNumber)
				case 2:
					d.handleGroupDissolve()
				}
			}
		case event, ok := <-sink:
			if ok {
				switch content := event.(type) {
				case *onchain.LogGrouping:
					groupID := fmt.Sprintf("%x", content.GroupId)
					go d.handleGrouping(content.NodeId, groupID)
				case *onchain.LogGroupDissolve:
					groupID := fmt.Sprintf("%x", content.GroupId)
					if d.isMember(groupID) {
						d.logger.Event("DGroupDismiss", map[string]interface{}{"GroupID": groupID})
						d.dkg.GroupDissolve(groupID)
					}
				case *onchain.LogPublicKeyAccepted:
					groupID := fmt.Sprintf("%x", content.GroupId)
					if d.isMember(groupID) {
						d.logger.Event("keyAccepted", map[string]interface{}{"GroupID": groupID})
					}
				case *onchain.LogUpdateRandom:
					randSeed = content.LastRandomness
					groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
					if d.isMember(groupID) {
						requestID := fmt.Sprintf("%x", content.LastRandomness)
						groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
						lastRand := fmt.Sprintf("%x", content.LastRandomness)
						ids, pub, sec, err := d.groupInfo(groupID)
						if err != nil {
							d.logger.Error(err)
							continue
						}
						f := map[string]interface{}{
							"RequestId":            requestID,
							"GroupID":              groupID,
							"LastSystemRandomness": lastRand}
						d.logger.Event("LogUpdateRandom", f)
						go d.handleQuery(ids, pub, sec, groupID, content.LastRandomness, content.LastRandomness, nil, "", "", uint32(onchain.TrafficSystemRandom))
					}
				case *onchain.LogRequestUserRandom:
					randSeed = content.LastSystemRandomness
					groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
					if d.isMember(groupID) {
						requestID := fmt.Sprintf("%x", content.RequestId)
						groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
						lastRand := fmt.Sprintf("%x", content.LastSystemRandomness)
						ids, pub, sec, err := d.groupInfo(groupID)
						if err != nil {
							d.logger.Error(err)
							continue
						}
						f := map[string]interface{}{
							"RequestId":            requestID,
							"GroupID":              groupID,
							"LastSystemRandomness": lastRand}
						d.logger.Event("LogRequestUserRandom", f)
						go d.handleQuery(ids, pub, sec, groupID, content.RequestId, content.LastSystemRandomness, content.UserSeed, "", "", uint32(onchain.TrafficUserRandom))
					}
				case *onchain.LogUrl:
					randSeed = content.Randomness
					groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
					if d.isMember(groupID) {
						requestID := fmt.Sprintf("%x", content.QueryId)
						groupID := fmt.Sprintf("%x", content.DispatchedGroupId)
						rand := fmt.Sprintf("%x", content.Randomness)
						ids, pub, sec, err := d.groupInfo(groupID)
						if err != nil {
							d.logger.Error(err)
							continue
						}
						f := map[string]interface{}{
							"RequestId":  requestID,
							"Randomness": rand,
							"DataSource": content.DataSource,
							"GroupID":    groupID}
						d.logger.Event("LogUrl", f)
						go d.handleQuery(ids, pub, sec, groupID, content.QueryId, content.Randomness, nil, content.DataSource, content.Selector, uint32(onchain.TrafficUserQuery))
					}
				case *onchain.LogStartCommitReveal:
					fmt.Println("startBlock ", content.StartBlock.String(), " commitDur ", content.CommitDuration.String(), "revealDur", content.RevealDuration.String())
					d.handleCR(content, randSeed)
				}

			} else {
				break L
			}
		case e, ok := <-errc:
			if ok {
				fmt.Println("errc event err ", e)
			} else {
				break L
			}
		case msg, ok := <-d.cSignToPeer:
			if !ok {
				return
			}
			peerSignMap[string(msg.Nonce)] = msg
		case msg, ok := <-peerEvent:
			if !ok {
				return
			}
			switch content := msg.Msg.Message.(type) {
			case *vss.Signature:
				if peerSignMap[string(content.Nonce)] != nil {
					fmt.Println("Got Sign ", peerSignMap[string(content.Nonce)].RequestId)
				}
				d.p.Reply(msg.Sender, msg.RequestNonce, peerSignMap[string(content.Nonce)])
			}
		case <-d.done:
			return
		}
	}
	d.chain.End()
	ips := d.p.MembersIP()
	var urls = []string{}
	urls = append(urls, "wss://rinkeby.infura.io/ws/v3/db19cf9028054762865cb9ce883c6ab8")
	urls = append(urls, "wss://rinkeby.infura.io/ws/v3/3a3e5d776961418e93a8b33fef2f6642")
	for _, ip := range ips {
		urls = append(urls, "ws://"+ip.String()+":8546")
		if len(urls) >= 5 {
			break
		}
	}
	d.chain.UpdateWsUrls(urls)
	d.chain.Start()
	goto S
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
