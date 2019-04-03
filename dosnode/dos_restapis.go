package dosnode

import (
	"context"
	"net/http"
	"strconv"
)

func (d *DosNode) status(w http.ResponseWriter, r *http.Request) {
	html := "Dos Client is working "
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
	result := d.dkg.GetGroupNumber()
	html = html + strconv.Itoa(result)
	w.Write([]byte(html))
}

func (d *DosNode) guardian(w http.ResponseWriter, r *http.Request) {
	d.chain.SignalGroupFormation(context.Background())
}
func (d *DosNode) p2p(w http.ResponseWriter, r *http.Request) {
	memNum, connNum := d.p.ConnectToAll()
	html := "members length : " + strconv.Itoa(memNum) + "\n connection length : " + strconv.Itoa(connNum)
	w.Write([]byte(html))
}
func (d *DosNode) proxy(w http.ResponseWriter, r *http.Request) {
	html := ""

	//	result := d.dkg.GetGroupNumber()
	workingGroupNum, err := d.chain.GetWorkingGroupSize()
	if err != nil {
		html = html + "WorkingGroupSize :" + err.Error() + "\n"
	} else {
		html = html + "WorkingGroupSize :" + strconv.FormatUint(workingGroupNum, 10) + "\n"
	}
	pendingGroupNum, err := d.chain.NumPendingGroups()
	if err != nil {
		html = html + "PendingGroupSize :" + err.Error() + "\n"
	} else {
		html = html + "PendingGroupSize :" + strconv.FormatUint(pendingGroupNum, 10) + "\n"
	}
	pendingNodeNum, err := d.chain.GetPengindNodeSize()
	if err != nil {
		html = html + "PendingNodeSize :" + err.Error() + "\n"
	} else {
		html = html + "PendingNodeSize :" + strconv.FormatUint(pendingNodeNum, 10) + "\n"
	}
	commitRevealTargetBlk, err := d.chain.CommitRevealTargetBlk()
	if err != nil {
		html = html + "commitRevealTargetBlk :" + err.Error() + "\n"
	} else {
		html = html + "commitRevealTargetBlk :" + strconv.FormatUint(commitRevealTargetBlk, 10) + "\n"
	}

	w.Write([]byte(html))
}
