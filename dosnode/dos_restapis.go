package dosnode

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DOSNetwork/core/onchain"
)

func (d *DosNode) startRESTServer() {
	fmt.Println("startRESTServer")
	mux := http.NewServeMux()
	mux.HandleFunc("/", d.status)
	mux.HandleFunc("/balance", d.balance)
	mux.HandleFunc("/enableAdmin", d.enableAdmin)
	mux.HandleFunc("/enableGuardian", d.enableGuardian)
	mux.HandleFunc("/signalGroupFormation", d.signalGroupFormation)
	mux.HandleFunc("/signalGroupDissolve", d.signalGroupDissolve)
	mux.HandleFunc("/signalBootstrap", d.signalBootstrap)
	mux.HandleFunc("/signalRandom", d.signalRandom)
	//TODO : Remove after beta .Only used for dev test
	mux.HandleFunc("/p2pTest", d.p2pTest)
	mux.HandleFunc("/dkgTest", d.dkgTest)
	mux.HandleFunc("/queryTest", d.queryTest)
	go func() {
		if err := http.ListenAndServe(":8080", mux); err != nil {
			fmt.Println("RESTServer err : ", err)
			os.Exit(1)
		}
	}()
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
	html = html + "StartTime         : " + d.startTime.Format("2006-01-02T15:04:05.999999-07:00") + "\n|"
	html = html + "Address           : " + fmt.Sprintf("%x", d.p.GetID()) + "\n|"
	html = html + "IP                : " + fmt.Sprintf("%s", d.p.GetIP()) + "\n|"
	html = html + "NumOfMembers      : " + strconv.Itoa(d.p.NumOfMembers()) + "\n|"
	html = html + "State             : " + d.state + "\n|"
	html = html + "IsPendingnNode    : " + strconv.FormatBool(isPendingNode) + "\n|"
	html = html + "TotalQuery        : " + strconv.Itoa(d.totalQuery) + "\n|"
	html = html + "FulfilledQuery    : " + strconv.Itoa(d.fulfilledQuery) + "\n|"
	html = html + "Group Number      : " + strconv.Itoa(d.numOfworkingGroup) + "\n|"
	html = html + "Group Number      : " + strconv.Itoa(d.numOfworkingGroup) + "\n"
	html = html + "=================================================" + "\n|"
	//	result := d.dkg.GetGroupNumber()
	balance, err := d.chain.Balance(ctx)
	if err != nil {
		html = html + "Balance  : " + err.Error() + "\n|"
	} else {
		html = html + "Balance  : " + balance.String() + "\n|"
	}
	workingGroupNum, err := d.chain.GetWorkingGroupSize(ctx)
	if err != nil {
		html = html + "WorkingGroupSize  : " + err.Error() + "\n|"
	} else {
		html = html + "WorkingGroupSize  : " + strconv.FormatUint(workingGroupNum, 10) + "\n|"
	}
	expiredGroupNum, err := d.chain.GetExpiredWorkingGroupSize(ctx)
	if err != nil {
		html = html + "ExpiredGroupSize  : " + err.Error() + "\n|"
	} else {
		html = html + "ExpiredGroupSize  : " + strconv.FormatUint(expiredGroupNum, 10) + "\n|"
	}
	pendingGroupNum, err := d.chain.NumPendingGroups(ctx)
	if err != nil {
		html = html + "PendingGroupSize  : " + err.Error() + "\n|"
	} else {
		html = html + "PendingGroupSize  : " + strconv.FormatUint(pendingGroupNum, 10) + "\n|"
	}
	pendingNodeNum, err := d.chain.NumPendingNodes(ctx)
	if err != nil {
		html = html + "PendingNodeSize   : " + err.Error() + "\n|"
	} else {
		html = html + "PendingNodeSize   : " + strconv.FormatUint(pendingNodeNum, 10) + "\n|"
	}

	curBlk, err := d.chain.CurrentBlock(ctx)
	if err != nil {
		html = html + "CurrentBlock      : " + err.Error() + "\n"
	} else {
		html = html + "CurrentBlock      : " + strconv.FormatUint(curBlk, 10) + "\n"
	}
	html = html + "=================================================" + "\n"
	w.Write([]byte(html))
}

func (d *DosNode) balance(w http.ResponseWriter, r *http.Request) {
	html := "Balance :"
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancelFunc()
	result, err := d.chain.Balance(ctx)
	if err != nil {
		html = html + err.Error()
	} else {
		html = html + result.String()
	}
	w.Write([]byte(html))
}

func (d *DosNode) enableAdmin(w http.ResponseWriter, r *http.Request) {
	d.isAdmin = true
}

func (d *DosNode) enableGuardian(w http.ResponseWriter, r *http.Request) {
	d.isGuardian = true
}

func (d *DosNode) signalBootstrap(w http.ResponseWriter, r *http.Request) {
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
		d.chain.SignalBootstrap(context.Background(), big.NewInt(int64(cid)))
	}
}

func (d *DosNode) signalRandom(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalRandom(context.Background())
}

func (d *DosNode) signalGroupFormation(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalGroupFormation(context.Background())
}

func (d *DosNode) signalGroupDissolve(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalGroupDissolve(context.Background())
}

func (d *DosNode) p2pTest(w http.ResponseWriter, r *http.Request) {
	members := d.p.MembersID()
	fmt.Println("p2p test ", len(members))
	for i := 0; i < len(members); i++ {
		fmt.Println("p2p test ", members[i])
	}
	/*
		for i := 0; i < len(members); i++ {
			id, err := d.p.ConnectTo(context.Background(), "", members[i])
			if err != nil {
				fmt.Println("ConnectTo err", id, " err ", err)
			} else {
				fmt.Println(i, "p2p ConnectTo ", id)
			}
			d.p.DisConnectTo(id)
		}*/
}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	buffer := make([]byte, 20)

	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}
		lines = append(lines, string(buffer[:bytesread]))
	}
	return lines, nil
}
func (d *DosNode) dkgTest(w http.ResponseWriter, r *http.Request) {
	groupID := big.NewInt(0)
	//members := d.p.MembersID()
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
					groupID = groupID.SetInt64(int64(i))
				}
			}
		}
	}
	members, err := readLines("/vault/members.txt")
	if err != nil {
		fmt.Println("readLines err ", err)
		return
	}
	fmt.Println("members len ", len(members))
	if start >= 0 && end >= 0 {
		if len(members) < (end - start) {
			fmt.Println("members size not enough:", len(members))
			return
		}
		var participants [][]byte

		for i := start; i < end; i++ {
			participants = append(participants, []byte(members[i]))
		}
		fmt.Println("members participants:", participants)

		d.onchainEvent <- &onchain.LogGrouping{
			GroupId: groupID,
			NodeId:  participants,
		}
	}
}
func (d *DosNode) queryTest(w http.ResponseWriter, r *http.Request) {
	groupId := big.NewInt(0)
	lastSys := big.NewInt(0)
	userSeed := big.NewInt(0)
	requestId := big.NewInt(0)
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
			if k == "gid" {
				i, err := strconv.Atoi(v[0])
				if err == nil && i >= 0 {
					groupId = groupId.SetInt64(int64(i))
				}
			} else if k == "rid" {
				i, err := strconv.Atoi(v[0])
				if err == nil && i >= 0 {
					requestId = requestId.SetInt64(int64(i))
				}
			}
		}
	}
	d.onchainEvent <- &onchain.LogRequestUserRandom{
		RequestId:            requestId,
		LastSystemRandomness: lastSys,
		UserSeed:             userSeed,
		DispatchedGroupId:    groupId,
	}
}
