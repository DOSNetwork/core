package p2p

import (
	"sync"
	"container/heap"
	"fmt"
)
const MAXPEERCOUNT = 21

type PeerManager struct {
	mu sync.Mutex
	peers sync.Map
	parray PeerArray
}

func (pm *PeerManager) LoadOrStore(id string, peer *PeerConn) (actual *PeerConn, loaded bool){
	pm.mu.Lock()
	defer pm.mu.Unlock()
	ac, l := pm.peers.LoadOrStore(id, peer)
	loaded = l
	if !l {
		heap.Push(&pm.parray, peer)
		if len(pm.parray) > MAXPEERCOUNT {
			p := heap.Pop(&pm.parray).(*PeerConn)
			//p.End() //todo disconnet peer_conn
			pm.peers.Delete(string(p.identity.Id))

		}
		actual = peer
	}else {
		actual = ac.(*PeerConn)
		fmt.Println(actual.identity.Address)
	}
	return
}

func (pm *PeerManager) Range(f func(key, value interface{}) bool) {
	pm.peers.Range(f)
}

func (pm *PeerManager) GetPeerByID(id string) (value *PeerConn, ok bool) {
	v, exist := pm.peers.Load(id)
	ok = exist
	if ok {
		value = v.(*PeerConn)
	}else {
		value = nil
	}
	return
}

func (pm *PeerManager) DeletePeer(id string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	_, exist := pm.GetPeerByID(id)
	if exist {
		pm.peers.Delete(id)
		index := -1
		for i := 0; i<pm.parray.Len(); i++ {
			if pm.parray[i].identity.Address == id {
				index = i
				break
			}
		}
		if index != -1 {
			heap.Remove(&pm.parray, index)
		}
	}

}

func (pm *PeerManager) PeerNum() int {
	return len(pm.parray)
}
