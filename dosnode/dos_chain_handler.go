package dosnode

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"

	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/share"
	errors "golang.org/x/xerrors"
)

func (d *DosNode) onchainLoop() {
	defer d.logger.Info("[DOS] End onchainLoop")
	var watchdogInterval int
	var currentBlockNumber uint64
	randSeed, _ := new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	inactiveNodes := make(map[string]time.Time)
	reconn := 0

	_, membersEvent, err := d.p.SubscribeEvent()
	if err != nil {
		d.logger.Error(err)
		return
	}
	defer func() {
		for _ = range membersEvent {
		}
	}()

	_ = d.chain.RegisterNewNode()

	for {
		//var onchainEvent chan interface{}
		var onchainErrc chan error
		subescriptions := []int{onchain.SubscribeLogGrouping, onchain.SubscribeLogGroupDissolve, onchain.SubscribeLogUrl,
			onchain.SubscribeLogUpdateRandom, onchain.SubscribeLogRequestUserRandom,
			onchain.SubscribeLogPublicKeyAccepted, onchain.SubscribeCommitrevealLogStartCommitreveal}
		d.onchainEvent, onchainErrc = d.chain.SubscribeEvent(subescriptions)

		checkBlkNumberPeriod := 0
		watchdogInterval = 15
		watchdog := time.NewTicker(time.Duration(watchdogInterval) * time.Second)
		reconn = 0
	L:
		for {
			select {
			case <-d.ctx.Done():
				d.logger.Info("[DOS] ctx.Done")
				d.chain.DisconnectAll()
				break L
			case event, ok := <-membersEvent:
				if !ok {
					d.logger.Info("[DOS] End membersEvent")
					d.End()
					continue
				}
				if d.isAdmin {
					switch event.EventType {
					case "member-join":
						if !inactiveNodes[event.NodeID].IsZero() {
							inactiveNodes[event.NodeID] = time.Time{}
						}
					case "member-failed":
						inactiveNodes[event.NodeID] = time.Now()
					}
				}
				d.logger.Event("peersUpdate", map[string]interface{}{"numOfPeers": d.p.NumOfMembers()})
			case <-watchdog.C:
				checkBlkNumberPeriod--
				if checkBlkNumberPeriod <= 0 {
					currentBlockNumber, err = d.chain.CurrentBlock()
					if err != nil {
						d.logger.Error(err)
						continue
					}
					checkBlkNumberPeriod = 50
				} else {
					currentBlockNumber++
				}
				if balance, err := d.chain.Balance(); err != nil {
					d.logger.Error(err)
					continue
				} else {
					if balance.Cmp(big.NewFloat(0.5)) == -1 {
						d.logger.Error(fmt.Errorf("No enough balance %f", balance))
						d.End()
						continue
					}
				}

				if d.isAdmin {
					now := time.Now()
					for nodeID, inactiveTime := range inactiveNodes {
						if !inactiveTime.IsZero() {
							diff := now.Sub(inactiveTime)
							mins := int(diff.Minutes())
							if mins >= 5 {
								d.logger.Debug(fmt.Sprintf("[DOS] Difference in Minutes over 5: %d Minutes %x", mins, nodeID))
								inactiveNodes[nodeID] = time.Time{}
								addr := common.Address{}
								b := []byte(nodeID)
								addr.SetBytes(b)
								d.chain.SignalUnregister(addr)
							}
						}
					}
				}

				if d.isGuardian {
					d.handleRandom(currentBlockNumber)
					d.handleGroupFormation()
					d.handleBootstrap(currentBlockNumber)
				}

			case err, ok := <-onchainErrc:
				if !ok {
					d.logger.Info("[DOS] End onchainErrc")
					break L
				}
				var oError *onchain.OnchainError
				if errors.As(err, &oError) {
					d.chain.Disconnect(oError.Idx)
				}
				d.logger.Error(err)
			case event, ok := <-d.onchainEvent:
				if !ok {
					d.logger.Info("[DOS] End onchainEvent")
					break L
				}
				switch content := event.(type) {
				case *onchain.LogGrouping:
					groupID := fmt.Sprintf("%x", content.GroupId)
					if d.isGuardian {
						go func() {
							select {
							case <-d.ctx.Done():
							case <-time.After(15 * 20 * time.Second):
								d.handleGroupDissolve()
							}
						}()
					}
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
					f := map[string]interface{}{
						"Cid":        content.Cid,
						"StartBlock": content.StartBlock.String(),
					}
					d.logger.Event("StartCR", f)
					d.logger.Info(fmt.Sprintf("startBlock %s commitDur %s revealDur %s", content.StartBlock.String(), content.CommitDuration.String(), content.RevealDuration.String()))
					go d.handleCR(content, randSeed)
				}
			}
		}
		d.logger.Info("[DOS] Rest onchainLoop")
		watchdog.Stop()
		//Drain the events out of the channel
		for _ = range d.onchainEvent {
		}
		d.logger.Info("[DOS] End Drain onchainEvent")

		for err = range onchainErrc {
			d.logger.Error(fmt.Errorf("[DOS] Drain onchainErrc %+v \n", err))
		}
		d.logger.Info("[DOS] End Drain onchainErrc")
		d.chain.DisconnectAll()
		select {
		case <-d.ctx.Done():
			return
		default:
		}
		d.logger.Info("[DOS] Reconnect to geth")
		//Connect to geth
		for {
			reconn++
			if reconn >= 10 {
				d.logger.Error(errors.New("Can't connect to geth"))
				d.End()
				return
			}
			//TODO : Add more geth from other sources
			t := time.Now().Add(90 * time.Second)
			if err := d.chain.Connect(d.config.ChainNodePool, t); err != nil {
				d.logger.Error(err)
				time.Sleep(10 * time.Second)
				d.logger.Info("[DOS] Reconnecting to geth")
				continue
			}
			break
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
	d.logger.Info("[DOS] Grouping start")
	d.logger.Event("GroupingStart", map[string]interface{}{"GroupID": groupID, "Topic": "Grouping"})
	defer d.logger.TimeTrack(time.Now(), "GroupingDone", map[string]interface{}{"GroupID": groupID, "Topic": "Grouping"})
	defer d.logger.Info(fmt.Sprintf("Grouping Done %x", groupID))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(20*15*time.Second))
	defer cancel()

	var errcList []chan error
	outFromDkg, errc, err := d.dkg.Grouping(ctx, groupID, participants)
	if err != nil {
		d.logger.Error(err)
		return
	}
	errcList = append(errcList, errc)
	errcList = append(errcList, registerGroup(ctx, d.chain, outFromDkg))
	allErrc := mergeErrors(ctx, errcList...)
	var ok bool
	for {
		select {
		case err, ok = <-allErrc:
			if !ok {
				return
			}
			d.logger.Error(err)
		case <-ctx.Done():
			return
		}
	}
	if err == nil {
		d.logger.Event("GroupingSucc", map[string]interface{}{"GroupID": groupID, "Topic": "Grouping"})
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

	// Generate random numbers in range [0..randSeed]
	if randSeed.Cmp(big.NewInt(1)) == -1 {
		randSeed, _ = new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	}
	sec, err := rand.Int(rand.Reader, randSeed)
	if err != nil {
		d.logger.Error(err)
		return
	}
	h := sha3.NewKeccak256()
	h.Write(abi.U256(sec))
	b := h.Sum(nil)
	hash := byte32(b)
	currentBlockNumber, err := d.chain.CurrentBlock()
	if err != nil {
		d.logger.Error(err)
		return
	}

	cid := cr.Cid
	waitCommit := cr.StartBlock.Uint64() - currentBlockNumber + 1
	waitReveal := cr.CommitDuration.Uint64() + 1
	waitRandom := cr.RevealDuration.Uint64() + 1
	if waitCommit < 0 {
		waitReveal = waitReveal - waitCommit
		waitRandom = waitRandom - waitCommit
		waitCommit = 0
	}

	time.Sleep(time.Duration(waitCommit*15) * time.Second)
	d.logger.Info("[DOS] Commit")
	d.logger.Event("Commit", map[string]interface{}{"CID": fmt.Sprintf("%x", cid)})
	if err := d.chain.Commit(cid, *hash); err != nil {
		d.logger.Error(err)
	}
	<-time.After(time.Duration(waitReveal*15) * time.Second)

	d.logger.Info("[DOS] Reveal")
	d.logger.Event("Reveal", map[string]interface{}{"CID": fmt.Sprintf("%x", cid)})
	if err := d.chain.Reveal(cid, sec); err != nil {
		d.logger.Error(err)
	}
}

func byte32(s []byte) (a *[32]byte) {
	if len(a) <= len(s) {
		a = (*[len(a)]byte)(unsafe.Pointer(&s[0]))
	}
	return a
}

func (d *DosNode) handleGroupFormation() {
	groupSize, err := d.chain.GroupSize()
	if err != nil {
		d.logger.Error(err)
		return
	}
	pendingNodeSize, err := d.chain.NumPendingNodes()
	if err != nil {
		d.logger.Error(err)
		return
	}
	expiredGroupSize, err := d.chain.GetExpiredWorkingGroupSize()
	if err != nil {
		d.logger.Error(err)
		return
	}
	if pendingNodeSize < groupSize && expiredGroupSize > 0 {
		d.chain.SignalGroupFormation()
		return
	}

	workingGroup, err := d.chain.GetWorkingGroupSize()
	if err != nil {
		d.logger.Error(err)
		return
	}
	bootstrapThreshold, err := d.chain.BootstrapStartThreshold()
	if err != nil {
		d.logger.Error(err)
		return
	}
	if workingGroup == 0 && pendingNodeSize < bootstrapThreshold && expiredGroupSize > 0 {
		d.chain.SignalGroupFormation()
		return
	}

	if pendingNodeSize < groupSize {
		d.logger.Debug(fmt.Sprintf("[DOS] Not enough pendingNodes (%v) vs groupSize (%v), skipping group formation ...", pendingNodeSize, groupSize))
		return
	}

	groupToPick, err := d.chain.GroupToPick()
	if err != nil {
		d.logger.Error(err)
		return
	}
	if workingGroup > 0 {
		if expiredGroupSize >= groupToPick {
			lastGrpFormReqId, err := d.chain.LastGroupFormationRequestId()
			if err != nil {
				d.logger.Error(err)
				return
			}
			if lastGrpFormReqId != 0 {
				d.logger.Debug("[DOS] Already in Group Formation Stage, skipping ...")
				return
			}
			d.logger.Debug("[DOS] Signaling new group formation ...")
			d.chain.SignalGroupFormation()
			return
		}
	} else if pendingNodeSize >= bootstrapThreshold {
		cid, err := d.chain.BootstrapRound()
		if err != nil {
			d.logger.Error(err)
			return
		}
		if cid == 0 {
			d.logger.Debug("[DOS] Bootstrap condition matches, signaling ...")
			d.chain.SignalGroupFormation()
			return
		}
	}
}

func (d *DosNode) handleRandom(currentBlockNumber uint64) {
	groupSize, err := d.chain.GetWorkingGroupSize()
	if err != nil {
		d.logger.Error(err)
		return
	}
	if groupSize == 0 {
		d.logger.Debug("[DOS] No live working group, skipping...")
		return
	}
	lastUpdatedBlock, err := d.chain.LastUpdatedBlock()
	if err != nil {
		d.logger.Error(err)
		return
	}
	sysrandInterval, err := d.chain.RefreshSystemRandomHardLimit()
	if err != nil {
		d.logger.Error(err)
		return
	}
	diff := currentBlockNumber - lastUpdatedBlock
	if diff > sysrandInterval {
		d.logger.Debug("[DOS] Signaling system randomness refresh...")
		d.chain.SignalRandom()
	}
}

func (d *DosNode) handleBootstrap(currentBlockNumber uint64) {
	cid, err := d.chain.BootstrapRound()
	if err != nil {
		d.logger.Error(err)
		return
	}
	if cid == 0 {
		d.logger.Debug("[DOS] Not in bootstrap phase ...")
		return
	}
	bootstrapEndBlk, err := d.chain.BootstrapEndBlk()
	if err != nil {
		d.logger.Error(err)
		return
	}
	if currentBlockNumber <= bootstrapEndBlk {
		d.logger.Debug("[DOS] Waiting for bootstrap to end before next step ...")
		return
	}
	pendingNodeSize, err := d.chain.NumPendingNodes()
	if err != nil {
		d.logger.Error(err)
		return
	}
	bootstrapThreshold, err := d.chain.BootstrapStartThreshold()
	if err != nil {
		d.logger.Error(err)
		return
	}
	if pendingNodeSize < bootstrapThreshold {
		d.logger.Debug(
			fmt.Sprintf(
				"[DOS] Not enough registered pendingNodes (%v) vs minimum bootstrap threshold (%v), skipping bootstrap ...",
				pendingNodeSize, bootstrapThreshold))
		return
	}

	d.logger.Debug("[DOS] Signaling system to bootstrap ...")
	if err := d.chain.SignalBootstrap(big.NewInt(int64(cid))); err != nil {
		d.logger.Error(err)
	}
}

func (d *DosNode) handleGroupDissolve() {
	pendingGroupSize, err := d.chain.NumPendingGroups()
	if err != nil {
		d.logger.Error(err)
		return
	}

	if pendingGroupSize > 0 {
		d.chain.SignalGroupDissolve()
	}
}

func (d *DosNode) isMember(groupID string) bool {
	return d.dkg.GetShareSecurity(groupID) != nil
}
