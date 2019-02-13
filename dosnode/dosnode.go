package dosnode

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"time"

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
	chain       onchain.ChainInterface
	dkg         dkg.P2PDkgInterface
	p           p2p.P2PInterface
	logger      log.Logger
	done        chan interface{}
	cSignToPeer chan *vss.Signature
	cPipCancel  chan string
}

func NewDosNode(suite suites.Suite, p p2p.P2PInterface, chain onchain.ChainInterface, dkg dkg.P2PDkgInterface) (dosNode *DosNode) {
	if logger == nil {
		logger = log.New("module", "dosclient")
	}
	return &DosNode{
		suite:       suite,
		p:           p,
		chain:       chain,
		dkg:         dkg,
		done:        make(chan interface{}),
		cSignToPeer: make(chan *vss.Signature, 50),
		cPipCancel:  make(chan string),
	}
}

func (d *DosNode) Start() (err error) {
	if err = d.listen(); err != nil {
		logger.Error(err)
		return
	}
	if err = d.chain.UploadID(); err != nil {
		logger.Error(err)
	}
	return
}

func (d *DosNode) End() {
	close(d.done)
}

func (d *DosNode) waitForGrouping(errs ...<-chan error) {
	errc := mergeErrors(errs...)
	for err := range errc {
		if err != nil {
			logger.Error(err)
			//sreturn
		}
	}
}

func (d *DosNode) waitForRequestDone(requestID string, errs ...<-chan error) {
	startTime := time.Now()
	errc := mergeErrors(errs...)
	for err := range errc {
		if err != nil {
			logger.Error(err)
		}
	}
	logger.Event("Close_waitForRequestDone", map[string]interface{}{
		"DOSEVENT":  "Close_waitForRequestDone",
		"TWaitTime": (time.Since(startTime).Nanoseconds() / 1000)})
	d.cPipCancel <- requestID

}

func (d *DosNode) listen() (err error) {
	eventGrouping := make(chan interface{})
	eventGroupDismiss := make(chan interface{})
	chUrl := make(chan interface{})
	chRandom := make(chan interface{})
	chUsrRandom := make(chan interface{}, 1)
	eventValidation := make(chan interface{})
	if err = d.chain.SubscribeEvent(eventGrouping, onchain.SubscribeDOSProxyLogGrouping); err != nil {
		return err
	}
	if err = d.chain.SubscribeEvent(eventGroupDismiss, onchain.SubscribeDOSProxyLogGroupDismiss); err != nil {
		return err
	}
	if err = d.chain.SubscribeEvent(chUrl, onchain.SubscribeDOSProxyLogUrl); err != nil {
		return err
	}
	if err = d.chain.SubscribeEvent(chRandom, onchain.SubscribeDOSProxyLogUpdateRandom); err != nil {
		return err
	}
	if err = d.chain.SubscribeEvent(chUsrRandom, onchain.SubscribeDOSProxyLogRequestUserRandom); err != nil {
		return err
	}
	if err = d.chain.SubscribeEvent(eventValidation, onchain.SubscribeDOSProxyLogValidationResult); err != nil {
		return err
	}
	peerEvent, err := d.p.SubscribeEvent(1, vss.Signature{})

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
			currentBlockNumber, err := d.chain.GetCurrentBlock()
			if err != nil {
				logger.Error(err)
			}
			select {
			case <-watchdog.C:
				ids := d.dkg.GetAnyGroupIDs()
				if len(ids) != 0 {
					lastUpdatedBlock, err := d.chain.GetLastUpdatedBlock()
					if err != nil {
						logger.Error(err)
						continue
					}
					diff := currentBlockNumber - lastUpdatedBlock
					if diff > SYSRANDOMNTERVAL {
						go d.chain.RandomNumberTimeOut()
					}
				}
			case requestID := <-d.cPipCancel:
				if pipeCancel[requestID] != nil {
					pipeCancel[requestID]()
					delete(pipeCancel, requestID)
					logger.Event("End_Pipeplne", map[string]interface{}{
						"DOSEVENT": "End_Pipeplne",
						"Time":     time.Now()})
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
					if r := bytes.Compare(d.chain.GetId(), id); r == 0 {
						isMember = true
					}
					groupIds = append(groupIds, id)
				}

				if isMember {
					sessionId := dkg.GIdsToSessionID(groupIds)
					f := map[string]interface{}{
						"Removed": content.Removed,
						"Tx":      content.Tx,
						"CurBlkN": currentBlockNumber,
						"BlockN":  content.BlockN,
						"Time":    time.Now()}
					logger.Event("DOS_Grouping", f)
					if !content.Removed {
						ctx, cancelFunc := context.WithCancel(context.Background())
						pipeCancel[sessionId] = cancelFunc
						var errcList []<-chan error
						outFromDkg, errc := d.dkg.Start(ctx, groupIds, sessionId)
						errcList = append(errcList, errc)
						errc = d.chain.UploadPubKey(ctx, outFromDkg)
						errcList = append(errcList, errc)
						go d.waitForGrouping(errcList...)
					} else { //if chain reorg then call
						if pipeCancel[sessionId] != nil {
							pipeCancel[sessionId]()
						}
					}
				}
			case msg := <-eventGroupDismiss:
				content, ok := msg.(*onchain.DOSProxyLogGroupDismiss)
				if !ok {
					log.Error(err)
					continue
				}

				if d.isMember(content.PubKey) && d.dkg.GroupDismiss(content.PubKey) {
					f := map[string]interface{}{
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
					if err = d.chain.UploadID(); err != nil {
						logger.Error(err)
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
						lastRand := content.LastSystemRandomness
						requestID := content.RequestId
						pubkey := content.DispatchedGroup
						useSeed := content.UserSeed
						ids := d.dkg.GetGroupIDs(content.DispatchedGroup)
						pubPoly := d.dkg.GetGroupPublicPoly(pubkey)
						ctx, cancelFunc := context.WithCancel(context.Background())
						pipeCancel[content.RequestId.String()] = cancelFunc
						var signShares []<-chan *vss.Signature
						var errcList []<-chan error
						var cSubmitter []chan []byte
						var cErr <-chan error
						var cSign <-chan *vss.Signature
						var cContent <-chan []byte
						//Build a pipeline
						cSubmitter, cErr = choseSubmitter(ctx, d.chain, lastRand, ids, len(ids))
						errcList = append(errcList, cErr)

						cContent = genUserRandom(ctx, cSubmitter[0], requestID.Bytes(), lastRand.Bytes(), useSeed.Bytes())
						errcList = append(errcList, cErr)
						cSign, cErr = genSign(ctx, cContent, d.cSignToPeer, d.dkg, d.suite, d.chain.GetId(), pubkey, requestID.String(), uint32(onchain.TrafficUserQuery))
						signShares = append(signShares, cSign)
						idx := 1
						for _, id := range ids {
							if r := bytes.Compare(d.chain.GetId(), id); r != 0 {
								cSign, cErr = requestSign(ctx, cSubmitter[idx], d.p, d.chain.GetId(), requestID.String(), uint32(onchain.TrafficUserQuery), id)
								signShares = append(signShares, cSign)
								errcList = append(errcList, cErr)
								idx++
							}
						}

						cSign, cErr = recoverSign(ctx, fanIn(ctx, signShares...), d.suite, pubPoly, (len(ids)/2 + 1), len(ids))
						errcList = append(errcList, cErr)
						cErr = d.chain.DataReturn(ctx, cSign)
						errcList = append(errcList, cErr)
						go d.waitForRequestDone(content.RequestId.String(), errcList...)
					} else {
						if pipeCancel[content.RequestId.String()] != nil {
							pipeCancel[content.RequestId.String()]()
						}
					}
				}
			case msg := <-chUrl:
				content, ok := msg.(*onchain.DOSProxyLogUrl)
				if !ok {
					log.Error(err)
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
						lastRand := content.Randomness
						url := content.DataSource
						selector := content.Selector
						requestID := content.QueryId
						pubkey := content.DispatchedGroup
						pubPoly := d.dkg.GetGroupPublicPoly(pubkey)
						ids := d.dkg.GetGroupIDs(pubkey)
						ctx, cancelFunc := context.WithCancel(context.Background())
						pipeCancel[content.QueryId.String()] = cancelFunc
						var signShares []<-chan *vss.Signature
						var errcList []<-chan error
						var cSubmitter []chan []byte
						var cErr <-chan error
						var cSign <-chan *vss.Signature
						var cContent <-chan []byte
						//Build a pipeline
						cSubmitter, cErr = choseSubmitter(ctx, d.chain, lastRand, ids, len(ids))
						errcList = append(errcList, cErr)

						cContent, cErr = genQueryResult(ctx, cSubmitter[0], url, selector)
						errcList = append(errcList, cErr)
						cSign, cErr = genSign(ctx, cContent, d.cSignToPeer, d.dkg, d.suite, d.chain.GetId(), pubkey, requestID.String(), uint32(onchain.TrafficUserQuery))
						errcList = append(errcList, cErr)
						signShares = append(signShares, cSign)
						idx := 1
						for _, id := range ids {
							if r := bytes.Compare(d.chain.GetId(), id); r != 0 {
								cSign, cErr = requestSign(ctx, cSubmitter[idx], d.p, d.chain.GetId(), requestID.String(), uint32(onchain.TrafficUserQuery), id)
								signShares = append(signShares, cSign)
								errcList = append(errcList, cErr)
								idx++
							}
						}
						cSign, cErr = recoverSign(ctx, fanIn(ctx, signShares...), d.suite, pubPoly, (len(ids)/2 + 1), len(ids))
						errcList = append(errcList, cErr)
						cErr = d.chain.DataReturn(ctx, cSign)
						errcList = append(errcList, cErr)
						go d.waitForRequestDone(requestID.String(), errcList...)
					} else {
						if pipeCancel[content.QueryId.String()] != nil {
							pipeCancel[content.QueryId.String()]()
						}
					}
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
						lastRand := content.LastRandomness
						requestID := lastRand.String()
						ids := d.dkg.GetGroupIDs(content.DispatchedGroup)
						pubkey := content.DispatchedGroup
						pubPoly := d.dkg.GetGroupPublicPoly(pubkey)
						ctx, cancelFunc := context.WithCancel(context.Background())
						pipeCancel[requestID] = cancelFunc
						var signShares []<-chan *vss.Signature
						var errcList []<-chan error
						var cSubmitter []chan []byte
						var cErr <-chan error
						var cSign <-chan *vss.Signature
						//Build a pipeline
						cSubmitter, cErr = choseSubmitter(ctx, d.chain, lastRand, ids, len(ids))
						errcList = append(errcList, cErr)
						cSign, cErr = genSign(ctx, genSysRandom(ctx, cSubmitter[0], lastRand.Bytes()), d.cSignToPeer, d.dkg, d.suite, d.chain.GetId(), pubkey, requestID, uint32(onchain.TrafficUserQuery))
						signShares = append(signShares, cSign)
						idx := 1
						for _, id := range ids {
							if r := bytes.Compare(d.chain.GetId(), id); r != 0 {
								cSign, cErr = requestSign(ctx, cSubmitter[idx], d.p, d.chain.GetId(), requestID, uint32(onchain.TrafficUserQuery), id)
								signShares = append(signShares, cSign)
								errcList = append(errcList, cErr)
								idx++
							}
						}

						cSign, cErr = recoverSign(ctx, fanIn(ctx, signShares...), d.suite, pubPoly, (len(ids)/2 + 1), len(ids))
						errcList = append(errcList, cErr)
						cErr = d.chain.SetRandomNum(ctx, cSign)
						errcList = append(errcList, cErr)
						go d.waitForRequestDone(requestID, errcList...)
					} else {
						if pipeCancel[content.LastRandomness.String()] != nil {
							pipeCancel[content.LastRandomness.String()]()
						}
					}
				}

			case msg := <-eventValidation:
				content, ok := msg.(*onchain.DOSProxyLogValidationResult)
				if !ok {
					log.Error(err)
				}
				_ = content
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
