package dosnode

import (
	"context"
	"net/http"
	"strconv"
)

func (d *DosNode) status(w http.ResponseWriter, r *http.Request) {
	html := "=================================================" + "\n|"
	html = html + "StartTime      :" + d.startTime.Format("2006-01-02T15:04:05.999999-07:00") + "\n|"
	html = html + "Balance        : " + d.chain.GetBalance().String() + "\n|"
	html = html + "State          : " + d.state + "\n|"
	html = html + "TotalQuery     : " + strconv.Itoa(d.totalQuery) + "\n|"
	html = html + "FulfilledQuery : " + strconv.Itoa(d.fulfilledQuery) + "\n|"
	html = html + "Group Number   : " + strconv.Itoa(d.numOfworkingGroup) + "\n|"
	html = html + "Group Number   : " + strconv.Itoa(d.numOfworkingGroup) + "\n"
	html = html + "=================================================" + "\n"
	w.Write([]byte(html))
}

func (d *DosNode) balance(w http.ResponseWriter, r *http.Request) {
	html := "Balance :"
	result := d.chain.GetBalance()
	html = html + result.String()
	w.Write([]byte(html))
}

func (d *DosNode) groupSize(w http.ResponseWriter, r *http.Request) {
	html := "Group Size :"

	w.Write([]byte(html))
}

func (d *DosNode) guardian(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalGroupFormation(context.Background())
}
func (d *DosNode) signalRandom(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalRandom(context.Background())
}
func (d *DosNode) p2p(w http.ResponseWriter, r *http.Request) {
	memNum, connNum := d.p.ConnectToAll()
	html := "members length : " + strconv.Itoa(memNum) + "\n connection length : " + strconv.Itoa(connNum)
	w.Write([]byte(html))
}
func (d *DosNode) proxy(w http.ResponseWriter, r *http.Request) {
	html := "=================================================" + "\n|"
	//	result := d.dkg.GetGroupNumber()
	workingGroupNum, err := d.chain.GetWorkingGroupSize()
	if err != nil {
		html = html + "WorkingGroupSize :" + err.Error() + "\n|"
	} else {
		html = html + "WorkingGroupSize :" + strconv.FormatUint(workingGroupNum, 10) + "\n|"
	}
	pendingGroupNum, err := d.chain.NumPendingGroups()
	if err != nil {
		html = html + "PendingGroupSize :" + err.Error() + "\n|"
	} else {
		html = html + "PendingGroupSize :" + strconv.FormatUint(pendingGroupNum, 10) + "\n|"
	}
	pendingNodeNum, err := d.chain.GetPengindNodeSize()
	if err != nil {
		html = html + "PendingNodeSize  :" + err.Error() + "\n|"
	} else {
		html = html + "PendingNodeSize  :" + strconv.FormatUint(pendingNodeNum, 10) + "\n|"
	}

	curBlk, err := d.chain.CurrentBlock()
	if err != nil {
		html = html + "CurrentBlock     :" + err.Error() + "\n"
	} else {
		html = html + "CurrentBlock     :" + strconv.FormatUint(curBlk, 10) + "\n"
	}
	html = html + "=================================================" + "\n"

	w.Write([]byte(html))
}
