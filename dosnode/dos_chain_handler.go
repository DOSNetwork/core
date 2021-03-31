package dosnode

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/share"
	"golang.org/x/crypto/sha3"
	errors "golang.org/x/xerrors"
)

func (d *DosNode) onchainLoop() {
	defer d.logger.Info("End onchainLoop")
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

		watchdog := time.NewTicker(time.Duration(d.chain.GetBlockTime()) * time.Second)
		reconn = 0
	L:
		for {
			select {
			case <-d.ctx.Done():
				d.logger.Info("ctx.Done")
				d.chain.DisconnectAll()
				break L
			case event, ok := <-membersEvent:
				if !ok {
					d.logger.Info("End membersEvent")
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
				if balance, err := d.chain.Balance(); err != nil {
					d.logger.Error(err)
					continue
				} else {
					if balance.Cmp(big.NewFloat(0.1)) == -1 {
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
								d.logger.Debug(fmt.Sprintf("Difference in Minutes over 5: %d Minutes %x", mins, nodeID))
								inactiveNodes[nodeID] = time.Time{}
								addr := common.Address{}
								b := []byte(nodeID)
								addr.SetBytes(b)
								d.chain.SignalUnregister(addr)
							}
						}
					}
				}
			case err, ok := <-onchainErrc:
				if !ok {
					d.logger.Info("End onchainErrc")
					break L
				}
				var oError *onchain.OnchainError
				if errors.As(err, &oError) {
					d.chain.Disconnect(oError.Idx)
				}
				d.logger.Error(err)
			case event, ok := <-d.onchainEvent:
				if !ok {
					d.logger.Info("End onchainEvent")
					break L
				}
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
		d.logger.Info("Rest onchainLoop")
		watchdog.Stop()
		//Drain the events out of the channel
		for _ = range d.onchainEvent {
		}
		d.logger.Info("End Drain onchainEvent")

		for err = range onchainErrc {
			d.logger.Error(fmt.Errorf("Drain onchainErrc %+v \n", err))
		}
		d.logger.Info("End Drain onchainErrc")
		d.chain.DisconnectAll()
		select {
		case <-d.ctx.Done():
			return
		default:
		}
		d.logger.Info("Reconnect to geth")
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
				d.logger.Info("Reconnecting to geth")
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
	d.logger.Info("Grouping start")
	d.logger.Event("GroupingStart", map[string]interface{}{"GroupID": groupID, "Topic": "Grouping"})
	defer d.logger.TimeTrack(time.Now(), "GroupingDone", map[string]interface{}{"GroupID": groupID, "Topic": "Grouping"})
	defer d.logger.Info(fmt.Sprintf("Grouping Done %x", groupID))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(20*d.chain.GetBlockTime())*time.Second)
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
	h := sha3.NewLegacyKeccak256()
	h.Write(math.U256Bytes(sec))
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

	time.Sleep(time.Duration(waitCommit*d.chain.GetBlockTime()) * time.Second)
	d.logger.Info("Commit")
	d.logger.Event("Commit", map[string]interface{}{"CID": fmt.Sprintf("%x", cid)})
	if err := d.chain.Commit(cid, *hash); err != nil {
		d.logger.Error(err)
	}
	<-time.After(time.Duration(waitReveal*d.chain.GetBlockTime()) * time.Second)

	d.logger.Info("Reveal")
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

func (d *DosNode) handleGroupFormation() bool {
	groupSize, err := d.chain.GroupSize()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	pendingNodeSize, err := d.chain.NumPendingNodes()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	expiredGroupSize, err := d.chain.GetExpiredWorkingGroupSize()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if pendingNodeSize < groupSize && expiredGroupSize > 0 {
		d.chain.SignalGroupFormation()
		return true
	}

	workingGroup, err := d.chain.GetWorkingGroupSize()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	bootstrapThreshold, err := d.chain.BootstrapStartThreshold()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if workingGroup == 0 && pendingNodeSize < bootstrapThreshold && expiredGroupSize > 0 {
		d.chain.SignalGroupFormation()
		return true
	}

	if pendingNodeSize < groupSize {
		d.logger.Debug(fmt.Sprintf("Not enough pendingNodes (%v) vs groupSize (%v), skipping group formation ...", pendingNodeSize, groupSize))
		return false
	}

	groupToPick, err := d.chain.GroupToPick()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if workingGroup > 0 {
		if expiredGroupSize >= groupToPick {
			lastGrpFormReqId, err := d.chain.LastGroupFormationRequestId()
			if err != nil {
				d.logger.Error(err)
				return false
			}
			if lastGrpFormReqId != 0 {
				d.logger.Debug("Already in Group Formation Stage, skipping ...")
				return false
			}
			d.logger.Debug("Signaling new group formation ...")
			d.chain.SignalGroupFormation()
			return true
		}
	} else if pendingNodeSize >= bootstrapThreshold {
		cid, err := d.chain.BootstrapRound()
		if err != nil {
			d.logger.Error(err)
			return false
		}
		if cid == 0 {
			d.logger.Debug("Bootstrap condition matches, signaling ...")
			d.chain.SignalGroupFormation()
			return true
		}
	}
	return false
}

func (d *DosNode) handleRandom(currentBlockNumber uint64) bool {
	groupSize, err := d.chain.GetWorkingGroupSize()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if groupSize == 0 {
		d.logger.Debug("No live working group, skipping update randomness ...")
		return false
	}
	lastUpdatedBlock, err := d.chain.LastUpdatedBlock()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	sysRandInterval, err := d.chain.RefreshSystemRandomHardLimit()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if lastUpdatedBlock+sysRandInterval > currentBlockNumber {
		d.logger.Debug("System random not expired yet, skipping update randomness ...")
		return false
	}
	cachedUpdatedBlock, err := d.chain.CachedUpdatedBlock()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	relayRespondLimit, err := d.chain.RelayRespondLimit()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if cachedUpdatedBlock+relayRespondLimit > currentBlockNumber {
		d.logger.Debug("System random is under updating, waiting for the settlment ...")
		return false
	}
	d.logger.Debug("Signaling system randomness refresh ...")
	if err := d.chain.SignalRandom(); err != nil {
		d.logger.Error(err)
		return false
	}
	return true
}

func (d *DosNode) handleBootstrap(currentBlockNumber uint64) bool {
	cid, err := d.chain.BootstrapRound()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if cid == 0 {
		d.logger.Debug("Not in bootstrap phase ...")
		return false
	}
	bootstrapEndBlk, err := d.chain.BootstrapEndBlk()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if currentBlockNumber <= bootstrapEndBlk {
		d.logger.Debug("Waiting for bootstrap to end before next step ...")
		return false
	}
	pendingNodeSize, err := d.chain.NumPendingNodes()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	bootstrapThreshold, err := d.chain.BootstrapStartThreshold()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if pendingNodeSize < bootstrapThreshold {
		d.logger.Debug(
			fmt.Sprintf(
				"Not enough registered pendingNodes (%v) vs minimum bootstrap threshold (%v), skipping bootstrap ...",
				pendingNodeSize, bootstrapThreshold))
		return false
	}

	d.logger.Debug("Signaling system to bootstrap ...")
	if err := d.chain.SignalBootstrap(big.NewInt(int64(cid))); err != nil {
		d.logger.Error(err)
		return false
	}
	return true
}

func (d *DosNode) handleGroupDissolve(currentBlockNumber uint64) bool {
	gid, err := d.chain.FirstPendingGroupId()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if gid.Cmp(big.NewInt(1)) == 0 {
		d.logger.Debug("Empty Pending Group List, skipping group dissolve ...")
		return false
	} else if gid.Cmp(big.NewInt(0)) == 0 {
		d.logger.Debug("Invalid groupId, skipping group dissolve ...")
		return false
	}
	startBlock, err := d.chain.PendingGroupStartBlock(gid)
	if err != nil {
		d.logger.Error(err)
		return false
	}
	pgMaxLife, err := d.chain.PendingGroupMaxLife()
	if err != nil {
		d.logger.Error(err)
		return false
	}
	if startBlock+pgMaxLife < currentBlockNumber {
		d.logger.Info("Signaling group dissolve ...")
		d.chain.SignalGroupDissolve()
		return true
	}
	return false
}

func (d *DosNode) isMember(groupID string) bool {
	return d.dkg.GetShareSecurity(groupID) != nil
}
