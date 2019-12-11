package dosnode

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/DOSNetwork/core/onchain"
)

func (d *DosNode) startRESTServer() (err error) {
	defer fmt.Println("[DOS] End RESTServer")
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
	s := http.Server{Addr: ":8080", Handler: mux}
	go func() {
		<-d.ctx.Done()
		s.Shutdown(context.Background())
	}()
	err = s.ListenAndServe()
	return
}

func (d *DosNode) status(w http.ResponseWriter, r *http.Request) {
	isPendingNode, err := d.chain.IsPendingNode(d.id)
	if err != nil {
		html := "err : " + err.Error() + "\n|"
		w.Write([]byte(html))
		return
	}
	html := "=================================================" + "\n|"
	html = html + "Version         : " + d.config.VERSION
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
	balance, err := d.chain.Balance()
	if err != nil {
		html = html + "Balance  : " + err.Error() + "\n|"
	} else {
		html = html + "Balance  : " + balance.String() + "\n|"
	}
	workingGroupNum, err := d.chain.GetWorkingGroupSize()
	if err != nil {
		html = html + "WorkingGroupSize  : " + err.Error() + "\n|"
	} else {
		html = html + "WorkingGroupSize  : " + strconv.FormatUint(workingGroupNum, 10) + "\n|"
	}
	expiredGroupNum, err := d.chain.GetExpiredWorkingGroupSize()
	if err != nil {
		html = html + "ExpiredGroupSize  : " + err.Error() + "\n|"
	} else {
		html = html + "ExpiredGroupSize  : " + strconv.FormatUint(expiredGroupNum, 10) + "\n|"
	}
	pendingGroupNum, err := d.chain.NumPendingGroups()
	if err != nil {
		html = html + "PendingGroupSize  : " + err.Error() + "\n|"
	} else {
		html = html + "PendingGroupSize  : " + strconv.FormatUint(pendingGroupNum, 10) + "\n|"
	}
	pendingNodeNum, err := d.chain.NumPendingNodes()
	if err != nil {
		html = html + "PendingNodeSize   : " + err.Error() + "\n|"
	} else {
		html = html + "PendingNodeSize   : " + strconv.FormatUint(pendingNodeNum, 10) + "\n|"
	}

	curBlk, err := d.chain.CurrentBlock()
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
	result, err := d.chain.Balance()
	if err != nil {
		html = html + err.Error()
	} else {
		html = html + result.String()
	}
	w.Write([]byte(html))
}

func (d *DosNode) enableAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DOS] isAdmin")
	d.isAdmin = true
}

func (d *DosNode) enableGuardian(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DOS] enableGuardian")
	d.isGuardian = true
}

func (d *DosNode) signalBootstrap(w http.ResponseWriter, r *http.Request) {
	cid := -1
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("[DOS]  %s: %s\n", k, v)
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
		d.chain.SignalBootstrap(big.NewInt(int64(cid)))
	}
}

func (d *DosNode) signalRandom(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalRandom()
}

func (d *DosNode) signalGroupFormation(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalGroupFormation()
}

func (d *DosNode) signalGroupDissolve(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalGroupDissolve()
}

func (d *DosNode) p2pTest(w http.ResponseWriter, r *http.Request) {
	d.End()
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
	start := -1
	end := -1
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("[DOS]  %s: %s\n", k, v)
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
		fmt.Println("[DOS] readLines err ", err)
		return
	}
	if start >= 0 && end >= 0 {
		if len(members) < (end - start) {
			fmt.Println("[DOS] members size not enough:", len(members))
			return
		}
		var participants [][]byte

		for i := start; i < end; i++ {
			participants = append(participants, []byte(members[i]))
		}

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
			fmt.Printf("[DOS]  %s: %s\n", k, v)
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
