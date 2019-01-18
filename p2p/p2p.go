package p2p

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/dedis/kyber"

	"github.com/golang/protobuf/proto"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/p2p/internal"
	"github.com/DOSNetwork/core/suites"
)

type P2P struct {
	identity internal.ID
	//Map of ID (string) <-> *p2p.PeerConn
	peers        *PeerConnManager //*sync.Map
	messages     chan P2PMessage
	suite        suites.Suite
	port         int
	secKey       kyber.Scalar
	pubKey       kyber.Point
	routingTable *dht.RoutingTable
	logger       log.Logger
}

func (n *P2P) SetID(id []byte) {
	n.identity.Id = id
}

func (n *P2P) GetID() []byte {
	return n.identity.Id
}

func (n *P2P) GetIP() string {
	return n.identity.Address
}

func (n *P2P) GetPeerConnManager() *PeerConnManager {
	return n.peers
}

func (n *P2P) Listen() (err error) {
	var ip string
	var listener net.Listener
	var pubKeyBytes []byte
	if ip, err = GetLocalIP(); err != nil {
		n.logger.Error(err)
		return
	}

	p := fmt.Sprintf(":%d", n.port)
	if listener, err = net.Listen("tcp", p); err != nil {
		n.logger.Error(err)
		return
	}

	//NAT discover

	//isPrivateIp, err := nat.IsPrivateIp()
	//if err != nil {
	//	return err
	//}
	//
	//if isPrivateIp {
	//	externalPort := nat.RandomPort()
	//	nat, err := nat.SetMapping("tcp", externalPort, listener.Addr().(*net.TCPAddr).Port, "DosNode")
	//	if err != nil {
	//		return err
	//	}
	//
	//	externalIp, err := nat.GetExternalAddress()
	//	if err != nil {
	//		return err
	//	}
	//
	//	n.port = externalPort
	//	ip = externalIp.String() + ":" + strconv.Itoa(n.port)
	//} else {
	//	n.port = listener.Addr().(*net.TCPAddr).Port
	//	ip = ip + ":" + strconv.Itoa(n.port)
	//}

	n.port = listener.Addr().(*net.TCPAddr).Port
	ip = ip + ":" + strconv.Itoa(n.port)

	n.secKey, n.pubKey = genPair()
	if pubKeyBytes, err = n.pubKey.MarshalBinary(); err != nil {
		n.logger.Error(err)
		return err
	}

	if n.identity.Id == nil {
		n.identity = dht.CreateID(ip, pubKeyBytes)
	} else {
		n.identity.Address = ip
		n.identity.PublicKey = pubKeyBytes
	}
	n.routingTable = dht.CreateRoutingTable(n.identity)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				n.logger.Error(err)
				continue
			}

			go func() {
				peer, err := NewPeerConn(n, &conn, n.messages, true)
				if err != nil {
					peer.EndWithoutDelete()
					n.logger.Error(err)
					return
				}
				if err = peer.HeardHi(); err != nil {
					peer.EndWithoutDelete()
					n.logger.Error(err)
					return
				}
				n.routingTable.Update(peer.identity)
			}()
		}
	}()

	return nil
}

func (n *P2P) Join(bootstrapIp string) (err error) {
	//it inserts the value of some known node c into the appropriate bucket as its first contact
	if _, err = n.connectTo(bootstrapIp); err != nil {
		n.logger.Error(err)
		return
	}
	//it does an iterativeFindNode for self
	results := n.findNode(n.identity, dht.BucketSize, 20)
	for _, result := range results {
		n.routingTable.Update(result)
	}
	return
}

func (n *P2P) Leave() {
	n.peers.Range(func(key, value interface{}) bool {
		peer := value.(*PeerConn)
		peer.End()
		return true
	})
	return
}

func (n *P2P) findPeer(id []byte) (peer *PeerConn, found bool) {
	var err error

	startTime := time.Now()
	defer n.logger.Metrics(time.Since(startTime).Seconds() * 1000)

	// Find Peer from existing peerConn
	if peer = n.peers.GetPeerByID(string(id)); peer != nil {
		found = true
		n.logger.Debug("FromExisting")
		return
	}

	// Find Peer from routing map
	peers := n.routingTable.GetPeers()
	if ip, ok := peers[string(id)]; ok {
		if peer, err = n.connectTo(ip); err != nil {
			n.logger.Error(err)
			return
		}
		n.logger.Debug("FromRoutingTable")
		found = true
		return
	}

	// Updating Routing table to find peer
	results := n.findNode(internal.ID{Id: id}, dht.BucketSize, 20)
	for _, result := range results {
		n.routingTable.Update(result)
	}

	if len(results) > 0 && bytes.Equal(results[0].Id, id) {
		if peer, err = n.connectTo(results[0].Address); err != nil {
			n.logger.Error(err)
			return
		}
		found = true
		n.logger.Debug("FromPeers")
	}

	return
}

func (n *P2P) SendMessage(id []byte, m proto.Message) (err error) {
	var peer *PeerConn
	var found bool
	startTime := time.Now()
	defer n.logger.Metrics(time.Since(startTime).Seconds() * 1000)

	if peer, found = n.findPeer(id); found {
		request := new(Request)
		request.SetMessage(m)
		request.SetTimeout(1 * time.Second)

		if _, err = peer.Request(request); err != nil {
			n.logger.Error(err)
			return
		}
	}

	return
}

func (n *P2P) Request(id []byte, m proto.Message) (msg proto.Message, err error) {
	var peer *PeerConn
	var found bool
	startTime := time.Now()
	defer n.logger.Metrics(time.Since(startTime).Seconds() * 1000)

	if peer, found = n.findPeer(id); found {
		request := new(Request)
		request.SetMessage(m)
		request.SetTimeout(1 * time.Second)

		if msg, err = peer.Request(request); err != nil {
			n.logger.Error(err)
		}

		return
	} else {
		return nil, errors.New("p2pRequest: peer not found")
	}
}

/*
This is a block call
*/
func (n *P2P) ConnectTo(addr string) (id []byte, err error) {
	var peer *PeerConn

	if peer, err = n.connectTo(addr); err != nil {
		n.logger.Error(err)
		return
	}
	id = peer.identity.Id
	return
}

func (n *P2P) connectTo(addr string) (peer *PeerConn, err error) {
	var conn net.Conn

	for retry := 0; retry < 10; retry++ {
		if conn, err = net.Dial("tcp", addr); err != nil {
			n.logger.Error(err)
			time.Sleep(1 * time.Second)
			continue
		}

		if peer, err = NewPeerConn(n, &conn, n.messages, false); err != nil {
			peer.EndWithoutDelete()
			n.logger.Error(err)
			continue
		}

		if err = peer.SayHi(); err != nil {
			peer.EndWithoutDelete()
			n.logger.Error(err)
			continue
		}

		n.peers.LoadOrStore(string(peer.identity.Id), peer)
		n.routingTable.Update(peer.identity)
		return
	}
	err = errors.New("Retry connection over the limit")
	n.logger.Error(err)
	return
}

type lookupBucket struct {
	pending int
	queue   []internal.ID
}

// FindNode queries all peers this current node acknowledges for the closest peers
// to a specified target ID.
//
// All lookups are done under a number of disjoint lookups in parallel.
//
// Queries at most #ALPHA nodes at a time per lookup, and returns all peer IDs closest to a target peer ID.
func (n *P2P) findNode(targetID internal.ID, alpha int, disjointPaths int) (results []internal.ID) {
	startTime := time.Now()
	defer n.logger.Metrics(time.Since(startTime).Seconds() * 1000)
	visited := new(sync.Map)

	var lookups []*lookupBucket

	// Start searching for target from #ALPHA peers closest to target by queuing
	// them up and marking them as visited.
	for i, peerID := range n.routingTable.FindClosestPeers(targetID, alpha) {

		visited.Store(dht.PublicKeyHex(peerID), struct{}{})

		if len(lookups) < disjointPaths {
			lookups = append(lookups, new(lookupBucket))
		}

		lookup := lookups[i%disjointPaths]
		lookup.queue = append(lookup.queue, peerID)

		results = append(results, peerID)
	}

	wait, mutex := &sync.WaitGroup{}, &sync.Mutex{}

	for _, lookup := range lookups {
		wait.Add(1)
		go func(lookup *lookupBucket) {
			mutex.Lock()
			results = append(results, lookup.performLookup(n, targetID, alpha, visited)...)
			mutex.Unlock()

			wait.Done()
		}(lookup)

	}

	// Wait until all #D parallel lookups have been completed.
	wait.Wait()

	// Sort resulting peers by XOR distance.
	sort.Slice(results, func(i, j int) bool {
		left := dht.XorID(results[i], targetID)
		right := dht.XorID(results[j], targetID)
		return dht.Less(left, right)
	})

	// Cut off list of results to only have the routing table focus on the
	// #dht.BucketSize closest peers to the current node.
	if len(results) > dht.BucketSize {
		results = results[:dht.BucketSize]
	}

	return
}

func (lookup *lookupBucket) performLookup(n *P2P, targetID internal.ID, alpha int, visited *sync.Map) (results []internal.ID) {
	responses := make(chan []*internal.ID)

	// Go through every peer in the entire queue and queue up what peers believe
	// is closest to a target ID.

	for ; lookup.pending < alpha && len(lookup.queue) > 0; lookup.pending++ {
		go n.queryPeerByID(lookup.queue[0], targetID, responses)

		lookup.queue = lookup.queue[1:]
	}

	// Empty queue.
	lookup.queue = lookup.queue[:0]

	// Asynchronous breadth-first search.
	for lookup.pending > 0 {
		response := <-responses

		lookup.pending--

		// Expand responses containing a peer's belief on the closest peers to target ID.
		for _, id := range response {
			peerID := internal.ID(*id)

			if _, seen := visited.LoadOrStore(dht.PublicKeyHex(peerID), struct{}{}); !seen && dht.PublicKeyHex(targetID) != dht.PublicKeyHex(peerID) {
				// Append new peer to be queued by the routing table.
				results = append(results, peerID)
				lookup.queue = append(lookup.queue, peerID)
			}
		}

		// Queue and request for #ALPHA closest peers to target ID from expanded results.
		for ; lookup.pending < alpha && len(lookup.queue) > 0; lookup.pending++ {
			go n.queryPeerByID(lookup.queue[0], targetID, responses)
			lookup.queue = lookup.queue[1:]
		}

		// Empty queue.
		lookup.queue = lookup.queue[:0]
	}

	return
}

func (n *P2P) queryPeerByID(peerID internal.ID, targetID internal.ID, responses chan []*internal.ID) {
	var client *PeerConn
	var err error

	if client = n.peers.GetPeerByID(string(peerID.Id)); client == nil {
		if client, err = n.connectTo(peerID.Address); err != nil {
			n.logger.Error(err)
			responses <- []*internal.ID{}
			return
		}
	}

	targetProtoID := internal.ID(targetID)

	request := new(Request)
	request.SetMessage(&internal.LookupNodeRequest{Target: &targetProtoID})
	request.SetTimeout(1 * time.Second)

	response, err := client.Request(request)
	if err != nil {
		n.logger.Error(err)
		responses <- []*internal.ID{}
		return
	}

	if response, ok := response.(*internal.LookupNodeResponse); ok {
		responses <- response.Peers
	} else {
		err = errors.New("Cast Error")
		n.logger.Error(err)
		responses <- []*internal.ID{}
	}
}
