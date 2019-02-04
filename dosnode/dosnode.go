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

var logger log.Logger

type DosNode struct {
	suite  suites.Suite
	chain  onchain.ChainInterface
	dkg    dkg.P2PDkgInterface
	p      p2p.P2PInterface
	logger log.Logger
	done   chan interface{}
	signc  chan *vss.Signature
}

func NewDosNode(suite suites.Suite, p p2p.P2PInterface, chain onchain.ChainInterface, dkg dkg.P2PDkgInterface) (dosNode *DosNode) {
	if logger == nil {
		logger = log.New("module", "dosclient")
	}
	return &DosNode{
		suite: suite,
		p:     p,
		chain: chain,
		dkg:   dkg,
		done:  make(chan interface{}),
		signc: make(chan *vss.Signature, 50),
	}
}

func (d *DosNode) Start() (err error) {
	if err = d.listen(); err != nil {
		logger.Error(err)
		return
	}
	//Register nodeID to onchain
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
			fmt.Println(err)
			return
		}
	}
}

func (d *DosNode) listen() (err error) {
	eventGrouping := make(chan interface{})
	chUrl := make(chan interface{})
	chRandom := make(chan interface{})
	chUsrRandom := make(chan interface{})
	eventValidation := make(chan interface{})
	if err = d.chain.SubscribeEvent(eventGrouping, onchain.SubscribeDOSProxyLogGrouping); err != nil {
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
		defer close(eventGrouping)
		defer close(chUrl)
		defer close(chRandom)
		defer close(chUsrRandom)
		defer close(eventValidation)
		defer d.p.UnSubscribeEvent(vss.Signature{})
		for {
			select {
			case msg := <-d.signc:
				peerSignMap[msg.QueryId] = msg
			case msg := <-peerEvent:
				switch content := msg.Msg.Message.(type) {
				case *vss.Signature:
					d.p.Reply(msg.Sender, msg.RequestNonce, peerSignMap[content.QueryId])
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
						"DispatchedGroup_1":    fmt.Sprintf("%x", content.DispatchedGroup[0]),
						"DispatchedGroup_2":    fmt.Sprintf("%x", content.DispatchedGroup[1]),
						"DispatchedGroup_3":    fmt.Sprintf("%x", content.DispatchedGroup[2]),
						"DispatchedGroup_4":    fmt.Sprintf("%x", content.DispatchedGroup[3]),
						"Removed":              content.Removed,
						"Tx":                   content.Tx,
						"BlockN":               content.BlockN,
						"Time":                 time.Now()}
					logger.Event("EthProxyRequestRandom", f)
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

						//Build a pipeline
						submitterc, errc := choseSubmitter(ctx, d.chain, lastRand, ids, len(ids)+1)
						errcList = append(errcList, errc)
						for i, id := range ids {
							var signc <-chan *vss.Signature
							var errc <-chan error
							if r := bytes.Compare(d.chain.GetId(), id); r != 0 {
								signc, errc = requestSign(ctx, submitterc[i], d.p, d.chain.GetId(), requestID.String(), uint32(onchain.TrafficUserRandom), id)
							} else {
								contentc := genUserRandom(ctx, submitterc[len(ids)], requestID.Bytes(), lastRand.Bytes(), useSeed.Bytes())
								signc, errc = genSign(ctx, submitterc[i], contentc, d.signc, d.dkg, d.suite, d.chain.GetId(), pubkey, requestID.String(), uint32(onchain.TrafficUserRandom))
							}
							signShares = append(signShares, signc)
							errcList = append(errcList, errc)
						}
						signc, errc := recoverSign(ctx, fanIn(ctx, signShares...), d.suite, pubPoly, (len(ids)/2 + 1), len(ids))
						errcList = append(errcList, errc)
						errc = d.chain.DataReturn(ctx, signc)
						errcList = append(errcList, errc)
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
					f := map[string]interface{}{
						"RequestId":         fmt.Sprintf("%x", content.QueryId),
						"Randomness":        fmt.Sprintf("%x", content.Randomness),
						"DataSource":        fmt.Sprintf("%x", content.DataSource),
						"DispatchedGroup_1": fmt.Sprintf("%x", content.DispatchedGroup[0]),
						"DispatchedGroup_2": fmt.Sprintf("%x", content.DispatchedGroup[1]),
						"DispatchedGroup_3": fmt.Sprintf("%x", content.DispatchedGroup[2]),
						"DispatchedGroup_4": fmt.Sprintf("%x", content.DispatchedGroup[3]),
						"Removed":           content.Removed,
						"Tx":                content.Tx,
						"BlockN":            content.BlockN,
						"Time":              time.Now()}
					logger.Event("EthProxyQueryURL", f)
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

						//Build a pipeline
						submitterc, errc := choseSubmitter(ctx, d.chain, lastRand, ids, len(ids)+1)
						errcList = append(errcList, errc)
						for i, id := range ids {
							var signc <-chan *vss.Signature
							var contentc <-chan []byte
							var errc <-chan error
							if r := bytes.Compare(d.chain.GetId(), id); r != 0 {
								signc, errc = requestSign(ctx, submitterc[i], d.p, d.chain.GetId(), requestID.String(), uint32(onchain.TrafficUserQuery), id)
							} else {
								contentc, errc = genQueryResult(ctx, submitterc[len(ids)], url, selector)
								errcList = append(errcList, errc)
								signc, errc = genSign(ctx, submitterc[i], contentc, d.signc, d.dkg, d.suite, d.chain.GetId(), pubkey, requestID.String(), uint32(onchain.TrafficUserQuery))
							}
							signShares = append(signShares, signc)
							errcList = append(errcList, errc)
						}
						signc, errc := recoverSign(ctx, fanIn(ctx, signShares...), d.suite, pubPoly, (len(ids)/2 + 1), len(ids))
						errcList = append(errcList, errc)
						errc = d.chain.DataReturn(ctx, signc)
						errcList = append(errcList, errc)
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
						"DispatchedGroup_1":    fmt.Sprintf("%x", content.DispatchedGroup[0]),
						"DispatchedGroup_2":    fmt.Sprintf("%x", content.DispatchedGroup[1]),
						"DispatchedGroup_3":    fmt.Sprintf("%x", content.DispatchedGroup[2]),
						"DispatchedGroup_4":    fmt.Sprintf("%x", content.DispatchedGroup[3]),
						"Removed":              content.Removed,
						"Tx":                   content.Tx,
						"BlockN":               content.BlockN,
						"Time":                 time.Now()}
					logger.Event("EthProxyUpdateSysRandom", f)
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

						//Build a pipeline
						submitterc, errc := choseSubmitter(ctx, d.chain, lastRand, ids, len(ids)+1)
						errcList = append(errcList, errc)
						for i, id := range ids {
							var signc <-chan *vss.Signature
							var contentc <-chan []byte
							var errc <-chan error
							if r := bytes.Compare(d.chain.GetId(), id); r != 0 {
								signc, errc = requestSign(ctx, submitterc[i], d.p, d.chain.GetId(), requestID, uint32(onchain.TrafficUserQuery), id)
							} else {
								contentc = genSysRandom(ctx, submitterc[len(ids)], lastRand.Bytes())
								errcList = append(errcList, errc)
								signc, errc = genSign(ctx, submitterc[i], contentc, d.signc, d.dkg, d.suite, d.chain.GetId(), pubkey, requestID, uint32(onchain.TrafficSystemRandom))
							}
							signShares = append(signShares, signc)
							errcList = append(errcList, errc)
						}
						signc, errc := recoverSign(ctx, fanIn(ctx, signShares...), d.suite, pubPoly, (len(ids)/2 + 1), len(ids))
						errcList = append(errcList, errc)
						errc = d.chain.SetRandomNum(ctx, signc)
						errcList = append(errcList, errc)
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
							"DispatchedGroup_1": fmt.Sprintf("%x", content.PubKey[0]),
							"DispatchedGroup_2": fmt.Sprintf("%x", content.PubKey[1]),
							"DispatchedGroup_3": fmt.Sprintf("%x", content.PubKey[2]),
							"DispatchedGroup_4": fmt.Sprintf("%x", content.PubKey[3]),
							"Removed":           content.Removed,
							"Tx":                content.Tx,
							"BlockN":            content.BlockN,
							"Time":              time.Now()}
						logger.Event("EthProxyQueryResult", f)
					} else if content.TrafficType == onchain.TrafficUserRandom {
						z := new(big.Int)
						z.SetBytes(content.Message)
						f := map[string]interface{}{
							"RequestId":         fmt.Sprintf("%x", content.TrafficId),
							"GeneratedRandom":   fmt.Sprintf("%x", z),
							"ValidationPass":    content.Pass,
							"DispatchedGroup_1": fmt.Sprintf("%x", content.PubKey[0]),
							"DispatchedGroup_2": fmt.Sprintf("%x", content.PubKey[1]),
							"DispatchedGroup_3": fmt.Sprintf("%x", content.PubKey[2]),
							"DispatchedGroup_4": fmt.Sprintf("%x", content.PubKey[3]),
							"Removed":           content.Removed,
							"Tx":                content.Tx,
							"BlockN":            content.BlockN,
							"Time":              time.Now()}
						logger.Event("EthProxyRandomResult", f)
					} else if content.TrafficType == onchain.TrafficSystemRandom {
						z := new(big.Int)
						z.SetBytes(content.Message)
						f := map[string]interface{}{
							"RequestId":       fmt.Sprintf("%x", content.TrafficId),
							"GeneratedRandom": fmt.Sprintf("%x", z),
							"ValidationPass":  content.Pass,
							"Removed":         content.Removed,
							"Tx":              content.Tx,
							"BlockN":          content.BlockN,
							"Time":            time.Now()}
						logger.Event("EthProxySysRandomResult", f)
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
