package p2p

type PeerArray []*PeerConn

func (pa PeerArray) Len() int { return len(pa) }
func (pa PeerArray) Less(i, j int) bool {  // 绑定less方法
	//return pa[i] < pa[j]  // 如果pa[i]<pa[j]生成的就是小根堆，如果pa[i]>pa[j]生成的就是大根堆
	return pa[i].lastusedtime.Unix() < pa[j].lastusedtime.Unix()
}
func (pa PeerArray) Swap(i, j int) {  // 绑定swap方法，交换两个元素位置
	pa[i], pa[j] = pa[j], pa[i]
}

func (pa *PeerArray) Pop() interface{} {
	old := *pa
	n := len(old)
	x := old[n-1]
	*pa = old[0 : n-1]
	return x
}

func (pa *PeerArray) Push(x interface{}) {  // 绑定push方法，插入新元素
	*pa = append(*pa, x.(*PeerConn))
}