package dosnode

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/DOSNetwork/core/onchain"
)

func (d *DosNode) p2p(w http.ResponseWriter, r *http.Request) {
	members := d.p.MemberList()
	fmt.Println("p2p test ", len(members))
	for i := 0; i < len(members); i++ {
		fmt.Println("p2p ConnectTo ", members[i])
		id, err := d.p.ConnectTo("", members[i])
		if err != nil {
			fmt.Println("ConnectTo ", id, " err ", err)
		}
	}
}

func (d *DosNode) grouping(w http.ResponseWriter, r *http.Request) {
	groupId := big.NewInt(0)
	members := d.p.MemberList()
	start := -1
	end := -1
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
			if k == "start" {
				i, err := strconv.Atoi(v[0])
				if err == nil && i >= 0 {
					start = i
				}
			} else if k == "end" {
				i, err := strconv.Atoi(v[0])
				if err == nil && i >= 0 {
					end = i
				}
			} else if k == "gid" {
				i, err := strconv.Atoi(v[0])
				if err == nil && i >= 0 {
					groupId = groupId.SetInt64(int64(i))
				}
			}
		}
	case "POST":
		/*
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", reqBody)
			w.Write([]byte("Received a POST request\n"))*/
	}
	if start >= 0 && end >= 0 {
		if len(members) < (end - start) {
			fmt.Println("members size not enough:", len(members))
			return
		}
		fmt.Println("members size:", len(members))
		sort.SliceStable(members, func(i, j int) bool {
			a := big.NewInt(0)
			a = a.SetBytes(members[i])
			b := big.NewInt(0)
			b = b.SetBytes(members[j])
			r := a.Cmp(b)
			if r == -1 {
				return true
			} else {
				return false
			}
		})
		var participants [][]byte
		for i := start; i < end; i++ {
			participants = append(participants, members[i])
		}

		d.eventGrouping <- &onchain.LogGrouping{
			GroupId: groupId,
			NodeId:  participants,
		}
	}
}
func (d *DosNode) status(w http.ResponseWriter, r *http.Request) {
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancelFunc()
	isPendingNode, err := d.chain.IsPendingNode(ctx, d.id)
	if err != nil {
		html := "err" + err.Error() + "\n|"
		w.Write([]byte(html))
		return
	}
	html := "=================================================" + "\n|"
	html = html + "StartTime      :" + d.startTime.Format("2006-01-02T15:04:05.999999-07:00") + "\n|"
	html = html + "Address          : " + fmt.Sprintf("%x", d.p.GetID()) + "\n|"
	html = html + "IP          : " + fmt.Sprintf("%s", d.p.GetIP()) + "\n|"
	html = html + "NumOfMembers     : " + strconv.Itoa(d.p.NumOfMembers()) + "\n|"
	html = html + "State          : " + d.state + "\n|"
	html = html + "IsPendingnNode : " + strconv.FormatBool(isPendingNode) + "\n|"
	html = html + "TotalQuery     : " + strconv.Itoa(d.totalQuery) + "\n|"
	html = html + "FulfilledQuery : " + strconv.Itoa(d.fulfilledQuery) + "\n|"
	html = html + "Group Number   : " + strconv.Itoa(d.numOfworkingGroup) + "\n|"
	html = html + "Group Number   : " + strconv.Itoa(d.numOfworkingGroup) + "\n"
	html = html + "=================================================" + "\n"
	w.Write([]byte(html))
}

func (d *DosNode) balance(w http.ResponseWriter, r *http.Request) {
	html := "Balance :"
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancelFunc()
	result, err := d.chain.Balance(ctx)
	if err != nil {
		html = html + result.String()
	} else {
		html = html + err.Error()
	}
	w.Write([]byte(html))
}

func (d *DosNode) groupSize(w http.ResponseWriter, r *http.Request) {
	html := "Group Size :"

	w.Write([]byte(html))
}

func (d *DosNode) guardian(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalGroupFormation(context.Background())
	d.chain.SignalGroupDissolve(context.Background())
	cid := -1
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
			if k == "cid" {
				i, err := strconv.Atoi(v[0])
				if err == nil && i >= 0 {
					cid = i
				}
			}
		}
	default:
	}
	if cid >= 0 {
		d.chain.SignalBootstrap(context.Background(), uint64(cid))
	}
}

func (d *DosNode) signalRandom(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalRandom(context.Background())
}

func (d *DosNode) proxy(w http.ResponseWriter, r *http.Request) {
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancelFunc()
	html := "=================================================" + "\n|"
	//	result := d.dkg.GetGroupNumber()
	workingGroupNum, err := d.chain.GetWorkingGroupSize(ctx)
	if err != nil {
		html = html + "WorkingGroupSize :" + err.Error() + "\n|"
	} else {
		html = html + "WorkingGroupSize :" + strconv.FormatUint(workingGroupNum, 10) + "\n|"
	}
	expiredGroupNum, err := d.chain.GetExpiredWorkingGroupSize(ctx)
	if err != nil {
		html = html + "ExpiredGroupSize :" + err.Error() + "\n|"
	} else {
		html = html + "ExpiredGroupSize :" + strconv.FormatUint(expiredGroupNum, 10) + "\n|"
	}
	pendingGroupNum, err := d.chain.NumPendingGroups(ctx)
	if err != nil {
		html = html + "PendingGroupSize :" + err.Error() + "\n|"
	} else {
		html = html + "PendingGroupSize :" + strconv.FormatUint(pendingGroupNum, 10) + "\n|"
	}
	pendingNodeNum, err := d.chain.NumPendingNodes(ctx)
	if err != nil {
		html = html + "PendingNodeSize  :" + err.Error() + "\n|"
	} else {
		html = html + "PendingNodeSize  :" + strconv.FormatUint(pendingNodeNum, 10) + "\n|"
	}

	curBlk, err := d.chain.CurrentBlock(ctx)
	if err != nil {
		html = html + "CurrentBlock     :" + err.Error() + "\n"
	} else {
		html = html + "CurrentBlock     :" + strconv.FormatUint(curBlk, 10) + "\n"
	}
	html = html + "=================================================" + "\n"

	w.Write([]byte(html))
}
