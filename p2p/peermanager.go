package p2p

import (
	"fmt"
	"math"
	"sync"

	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var MAXPEERCONNCOUNT uint32 = 100000

func init() {
	temp, err := strconv.ParseUint(os.Getenv("MAXPEERCONNCOUNT"), 10, 32)
	if err == nil {
		MAXPEERCONNCOUNT = uint32(temp)
		fmt.Println("MAXPEERCONNCOUNT", MAXPEERCONNCOUNT)
	}
}

type PeerConnManager struct {
	mu    sync.RWMutex
	peers map[string]*PeerConn
	log   *logrus.Entry
}

func CreatePeerConnManager(log *logrus.Entry) (pconn *PeerConnManager) {
	pconn = &PeerConnManager{
		peers: make(map[string]*PeerConn),
		log:   log,
	}
	return pconn
}

func (pm *PeerConnManager) FindLessUsedPeerConn() (pconn *PeerConn) {
	var lastusedtime int64
	lastusedtime = math.MaxInt64
	pconn = nil
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	for _, value := range pm.peers {
		if value.lastusedtime.Unix() < lastusedtime {
			lastusedtime = value.lastusedtime.Unix()
			pconn = value
		}
	}
	return
}

func (pm *PeerConnManager) LoadOrStore(id string, peer *PeerConn) (actual *PeerConn, loaded bool) {
	if actual, loaded = pm.peers[id]; loaded {
		return
	}
	actual = peer
	loaded = false
	pm.mu.Lock()
	pm.peers[id] = peer
	pm.mu.Unlock()
	if pm.PeerConnNum() > MAXPEERCONNCOUNT {
		p := pm.FindLessUsedPeerConn()
		fmt.Println("Force delete ", p.identity.Id)
		p.End()
	}
	pm.log.WithFields(logrus.Fields{
		"countConn": pm.PeerConnNum(),
	}).Info()
	return
}

func (pm *PeerConnManager) Range(f func(key, value interface{}) bool) {
	for key, value := range pm.peers {
		if !f(key, value) {
			break
		}
	}
}

func (pm *PeerConnManager) GetPeerByID(id string) *PeerConn {
	value, ok := pm.peers[id]
	if ok {
		return value
	} else {
		return nil
	}
}

func (pm *PeerConnManager) DeletePeer(id string) {
	if pm.GetPeerByID(id) != nil {
		pm.mu.Lock()
		defer pm.mu.Unlock()
		delete(pm.peers, id)
		pm.log.WithFields(logrus.Fields{
			"countConn": pm.PeerConnNum,
		}).Info()
	}
}

func (pm *PeerConnManager) PeerConnNum() uint32 {
	return uint32(len(pm.peers))
}
