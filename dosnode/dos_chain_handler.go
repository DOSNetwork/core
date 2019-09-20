package dosnode

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"

	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/share"
)

const (
	watchdogInterval = 10 //In Minute
)

func (d *DosNode) onchainLoop() {
	defer fmt.Println("End onchainLoop")
	watchdog := time.NewTicker(watchdogInterval * time.Minute)
	defer watchdog.Stop()
	randSeed, _ := new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	inactiveNodes := make(map[string]time.Time)
	for {
		//Connect to geth
		for {
			ips := d.p.RandomPeerIP()
			var urls = []string{}
			urls = append(urls, "ws://"+d.p.GetIP().String()+":8546")
			for _, ip := range ips {
				urls = append(urls, "ws://"+ip+":8546")
				if len(urls) >= 5 {
					break
				}
			}
			for _, url := range d.config.ChainNodePool {
				urls = append(urls, url)
			}
			if len(urls) >= 5 {
				d.chain.UpdateWsUrls(urls)
				fmt.Println("UpdateWsUrls ", urls)
				ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(60*time.Second))
				if err := d.chain.Connect(ctx); err == nil {
					break
				} else {
					fmt.Println("Connect err  ", err)
				}
				cancel()
			}
		}
		fmt.Println("Done connect ")

		//TODO: Check to see if it is a valid stacking node first
		_ = d.chain.RegisterNewNode(context.Background())

		//var onchainEvent chan interface{}
		var errc chan error
		subescriptions := []int{onchain.SubscribeLogGrouping, onchain.SubscribeLogGroupDissolve, onchain.SubscribeLogUrl,
			onchain.SubscribeLogUpdateRandom, onchain.SubscribeLogRequestUserRandom,
			onchain.SubscribeLogPublicKeyAccepted, onchain.SubscribeCommitrevealLogStartCommitreveal}
		d.onchainEvent, errc = d.chain.SubscribeEvent(subescriptions)
		_, membersEvent, err := d.p.SubscribeEvent()
		if err != nil {
			fmt.Println("SubscribeEvent err ", err)
			return
		}
	L:
		for {
			select {
			case <-d.ctx.Done():
				//Drain the events out of the channel
				for _ = range d.onchainEvent {
				}
				for _ = range errc {
				}

				return
			case event := <-membersEvent:
				if d.isGuardian {
					//fmt.Println("", event.EventType, fmt.Sprintf("%x", event.NodeID), event.Addr)
					switch event.EventType {
					case "member-join":
						if !inactiveNodes[event.NodeID].IsZero() {
							inactiveNodes[event.NodeID] = time.Time{}
						}
					case "member-failed":
						inactiveNodes[event.NodeID] = time.Now()
					}
				}
			case <-watchdog.C:
				currentBlockNumber, err := d.chain.CurrentBlock(context.Background())
				if err != nil {
					d.logger.Error(err)
					fmt.Println("Dos node CurrentBlock err ", err)
					break L
				}
				if balance, err := d.chain.Balance(context.Background()); err != nil {
					fmt.Println("Dos node CurrentBlock err ", err)
				} else {
					if balance.Cmp(big.NewFloat(0.5)) == -1 {
						fmt.Println("Dos node no enough balance ", balance)
						go d.End()
					}
				}
				if d.isGuardian {
					now := time.Now()
					for nodeID, inactiveTime := range inactiveNodes {
						if !inactiveTime.IsZero() {
							diff := now.Sub(inactiveTime)
							mins := int(diff.Minutes())
							if mins >= 5 {
								fmt.Printf("Diffrence in Minutes over 5: %d Minutes %s\n", mins, fmt.Sprintf("%x", nodeID))
								inactiveNodes[nodeID] = time.Time{}
								addr := common.Address{}
								b := []byte(nodeID)
								addr.SetBytes(b)
								d.chain.SignalUnregister(context.Background(), addr)
							}
						}
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
			case err, ok := <-errc:
				if !ok {
					break L
				}
				fmt.Println("onchainLoop err ", err)
			case event, ok := <-d.onchainEvent:
				if !ok {
					break L
				}
				switch content := event.(type) {
				case *onchain.LogGrouping:
					groupID := fmt.Sprintf("%x", content.GroupId)
					fmt.Println("onchain.LogGrouping")

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
					go d.handleCR(content, randSeed)
				}
			}
		}
	}
}

func (d *DosNode) handleGrouping(participants [][]byte, groupID string) {
	isMember := false
	for _, id := range participants {
		//fmt.Println("Compare ", d.id, id)
		if r := bytes.Compare(d.id, id); r == 0 {
			isMember = true
			break
		}
	}
	if !isMember {
		return
	}
	fmt.Println("Grouping start")
	d.logger.Event("Grouping", map[string]interface{}{"GroupID": groupID})
	defer d.logger.TimeTrack(time.Now(), "TimeGrouping", map[string]interface{}{"GroupID": groupID})
	defer fmt.Println("!!!!!!!!Grouping Done ", groupID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60*60*time.Second))
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
	var ok bool
	for {
		select {
		case err, ok = <-allErrc:
			if !ok {
				return
			}
			d.logger.Error(err)
			//d.logger.Event("waitForGroupingError", map[string]interface{}{"Error": err.Error(), "GroupID": groupID})
		case <-ctx.Done():
			return
		}
	}
	if err == nil {
		d.logger.Event("GroupingSucc", map[string]interface{}{"GroupID": groupID})
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

	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(160*15*time.Second))
	defer cancel()

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
	currentBlockNumber, err := d.chain.CurrentBlock(ctx)
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
	fmt.Println("Commit", *hash)
	d.logger.Event("Commit", map[string]interface{}{"CID": fmt.Sprintf("%x", cid)})
	if err := d.chain.Commit(ctx, cid, *hash); err != nil {
		fmt.Println("Commit err ", err)
		d.logger.Error(err)
	}
	<-time.After(time.Duration(waitReveal*15) * time.Second)

	fmt.Println("Reveal", fmt.Sprintf("%x", sec))
	d.logger.Event("Reveal", map[string]interface{}{"CID": fmt.Sprintf("%x", cid)})
	if err := d.chain.Reveal(ctx, cid, sec); err != nil {
		fmt.Println("Reveal err ", err)
		d.logger.Error(err)
	}
	<-time.After(time.Duration(waitRandom*15) * time.Second)

	fmt.Println("SignalBootstrap")
	d.logger.Event("SignalBootstrap", map[string]interface{}{"CID": fmt.Sprintf("%x", cid)})
	if err := d.chain.SignalBootstrap(ctx, cid); err != nil {
		fmt.Println("SignalBootstrap err ", err)

		d.logger.Error(err)
	}
}

func byte32(s []byte) (a *[32]byte) {
	if len(a) <= len(s) {
		a = (*[len(a)]byte)(unsafe.Pointer(&s[0]))
	}
	return a
}

func (d *DosNode) handleGroupFormation(currentBlockNumber uint64) {
	groupToPick, err := d.chain.GroupToPick(context.Background())
	if err != nil {
		return
	}
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
	workingGroup, err := d.chain.GetWorkingGroupSize(context.Background())
	if err != nil {
		d.logger.Error(err)
		return
	}
	expiredGroupSize, err := d.chain.GetExpiredWorkingGroupSize(context.Background())
	if err != nil {
		d.logger.Error(err)
		return
	}
	if pendingNodeSize < groupSize+(groupSize/2) {
		if expiredGroupSize != 0 {
			d.chain.SignalGroupFormation(context.Background())
		}
		return
	}

	if expiredGroupSize >= groupToPick {
		d.chain.SignalGroupFormation(context.Background())
		return
	}
	if workingGroup == 0 {
		d.chain.SignalGroupFormation(context.Background())
		return
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

	if pendingGrouSize > 0 {
		d.chain.SignalGroupDissolve(context.Background())
	}
}

func (d *DosNode) isMember(groupID string) bool {
	return d.dkg.GetShareSecurity(groupID) != nil
}
