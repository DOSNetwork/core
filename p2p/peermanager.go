package p2p

import (
	"fmt"
	"math"
	"sync"

	"github.com/sirupsen/logrus"
)

const MAXPEERCOUNT = 100000

type PeerConnManager struct {
	mu    sync.Mutex
	peers sync.Map
	count uint32
	log   *logrus.Entry
}

func (pm *PeerConnManager) FindLessUsedPeerConn() (pconn *PeerConn) {
	var lastusedtime int64
	lastusedtime = math.MaxInt64
	pconn = nil
	pm.peers.Range(func(key, value interface{}) bool {
		pc := value.(*PeerConn)
		if pc.lastusedtime.Unix() < lastusedtime {
			lastusedtime = pc.lastusedtime.Unix()
			pconn = pc
		}
		return true
	})
	return
}

func (pm *PeerConnManager) LoadOrStore(id string, peer *PeerConn) (actual *PeerConn, loaded bool) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	ac, l := pm.peers.LoadOrStore(id, peer)
	loaded = l
	if !l {
		pm.count++
		if pm.count > MAXPEERCOUNT {
			p := pm.FindLessUsedPeerConn()
			fmt.Println("Force delete ", p.identity.Id)
			p.End()
			pm.peers.Delete(string(p.identity.Id))
			pm.count--

		}
		pm.log.WithFields(logrus.Fields{
			"countConn": pm.count,
		}).Info()
		actual = peer
	} else {
		actual = ac.(*PeerConn)
		fmt.Println("Load Peerconn:" + actual.identity.Address)
	}
	return
}

func (pm *PeerConnManager) Range(f func(key, value interface{}) bool) {
	pm.peers.Range(f)
}

func (pm *PeerConnManager) GetPeerByID(id string) (value *PeerConn, ok bool) {
	v, exist := pm.peers.Load(id)
	ok = exist
	if ok {
		value = v.(*PeerConn)
	} else {
		value = nil
	}
	return
}

func (pm *PeerConnManager) DeletePeer(id string) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	_, exist := pm.GetPeerByID(id)
	if exist {
		pm.peers.Delete(id)
		pm.count--
		pm.log.WithFields(logrus.Fields{
			"countConn": pm.count,
		}).Info()
	}

}

func (pm *PeerConnManager) PeerConnNum() uint32 {
	return pm.count
}
