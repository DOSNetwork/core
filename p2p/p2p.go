package p2p

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/dedis/kyber"

	"github.com/golang/protobuf/proto"

	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/p2p/internal"
	"github.com/DOSNetwork/core/suites"
	"github.com/sirupsen/logrus"
)

type P2P struct {
	identity internal.ID
	//Map of ID (string) <-> *p2p.PeerConn
	peers *sync.Map
	// Todo:Add a subscrive fucntion to send message to application
	messages     chan P2PMessage
	suite        suites.Suite
	port         int
	secKey       kyber.Scalar
	pubKey       kyber.Point
	routingTable *dht.RoutingTable
	log          *logrus.Logger
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

func (n *P2P) Listen() error {
	ip, err := GetLocalIp()
	if err != nil {
		log.Fatal(err)
	}

	p := fmt.Sprintf(":%d", n.port)

	listener, err := net.Listen("tcp", p)
	if err != nil {
		return err
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
	pubKeyBytes, err := n.pubKey.MarshalBinary()
	if err != nil {
		return err
	}

	if n.identity.Id == nil {
		n.identity = dht.CreateID(ip, pubKeyBytes)
	} else {
		n.identity.Address = ip
		n.identity.PublicKey = pubKeyBytes
	}
	n.routingTable = dht.CreateRoutingTable(n.identity)

	fmt.Println("listen on: ", ip, " id: ", n.identity.Id)
	// Handle new clients.

	go func() {
		for {
			if conn, err := listener.Accept(); err == nil {
				//Create a peer client
				var peer *PeerConn
				peer, err = NewPeerConn(n, &conn, n.messages)
				if err != nil {
					log.Fatal(err)
				}
				n.peers.LoadOrStore(string(peer.identity.Id), peer)
				n.routingTable.Update(peer.identity)

			} else {
				fmt.Println("Failed accepting a connection request:", err)
				log.Fatal(err)
			}
		}
	}()

	return nil
}

func (n *P2P) Join(bootstrapIp string) (err error) {
	//it inserts the value of some known node c into the appropriate bucket as its first contact
	_, err = n.NewPeer(bootstrapIp)
	//it does an iterativeFindNode for self
	results := n.FindNodeById(n.GetID())
	for _, result := range results {
		n.GetRoutingTable().Update(result)
		//fmt.Println(n.GetIP(), "Update peer: ", result.Address)
	}
	return
}

func (n *P2P) Broadcast(m proto.Message) {
	n.peers.Range(func(key, value interface{}) bool {
		client := value.(*PeerConn)
		err := client.SendMessage(m)
		if err != nil {
			fmt.Println("P2P Broadcast ", err)
		}
		return true
	})
}

func (n *P2P) CheckPeers() {
	n.peers.Range(func(key, value interface{}) bool {
		client := value.(*PeerConn)
		fmt.Println("CheckPeers:", client.identity.Address)
		return true
	})
}

func (n *P2P) SendMessage(id []byte, m proto.Message) (err error) {
	var sendResult bool
	var tSendMessage float64
	var tFindNode float64
	start := time.Now()
	value, loaded := n.peers.Load(string(id))
	if !loaded {
		tFindNodeStart := time.Now()
		n.FindNodeById(id)
		tFindNode = time.Since(tFindNodeStart).Seconds()
	}
	value, loaded = n.peers.Load(string(id))
	if loaded {
		client := value.(*PeerConn)
		err = client.SendMessage(m)
		sendResult = true
	} else {
		err = fmt.Errorf("can't find node %s", string(id))
		sendResult = false
		contactsAfter := n.routingTable.GetPeers()
		fmt.Println("[ERROR] can't find node", id)
		for k, v := range contactsAfter {
			fmt.Println("[ERROR] ", []byte(k), " - ", v)
		}
	}
	tSendMessage = time.Since(start).Seconds()

	a := new(big.Int).SetBytes(n.GetID()).String()
	b := new(big.Int).SetBytes(id).String()

	n.log.WithFields(logrus.Fields{
		"time-sendMessage": tSendMessage,
		"time-findNode":    tFindNode,
		"targetID":         b,
		"nodeID":           a,
		"send-result":      sendResult,
	}).Info()
	return
}

func (n *P2P) GetRoutingTable() *dht.RoutingTable {
	return n.routingTable
}

/*
This is a block call
*/

func (n *P2P) NewPeer(addr string) (id []byte, err error) {
	var peer *PeerConn
	var conn net.Conn
	for retry := 0; retry < 10; retry++ {
		//1)Check if this address has been in peers map
		existing := false
		n.peers.Range(func(key, value interface{}) bool {
			client := value.(*PeerConn)
			if client.identity.Address == addr {

				existing = true
				id = make([]byte, len(client.identity.Id))
				copy(id, client.identity.Id)
			}
			return true
		})
		if existing {
			fmt.Println("Existing peer")
			return
		}

		//2)Dial to peer to get a connection
		conn, err = net.Dial("tcp", addr)
		if err != nil {
			fmt.Println("Dial err ", err)
			time.Sleep(1 * time.Second)
			continue
		}

		peer, err = NewPeerConn(n, &conn, n.messages)
		if err != nil {
			fmt.Println("NewPeerConn err ", err)
			continue
		}
		n.peers.LoadOrStore(string(peer.identity.Id), peer)
		n.routingTable.Update(peer.identity)

		n.peers.Range(func(key, value interface{}) bool {
			client := value.(*PeerConn)
			if client.identity.Address == addr {
				existing = true
				id = make([]byte, len(client.identity.Id))
				copy(id, client.identity.Id)
			}
			return true
		})
		return
	}
	err = errors.New("Peer : retried over 10 times")

	return
}

func (n *P2P) FindNodeById(id []byte) []internal.ID {
	targetId := internal.ID{
		Id: id,
	}
	return n.findNode(targetId, dht.BucketSize, 20)
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
		go func(lookup *lookupBucket) {
			mutex.Lock()
			results = append(results, lookup.performLookup(n, targetID, alpha, visited)...)
			mutex.Unlock()

			wait.Done()
		}(lookup)

		wait.Add(1)
	}

	// Wait until all #D parallel lookups have been completed.
	wait.Wait()

	// Sort resulting peers by XOR distance.
	sort.Slice(results, func(i, j int) bool {
		left := dht.Xor(results[i], targetID)
		right := dht.Xor(results[j], targetID)
		return dht.Less(left, right)
	})

	// Cut off list of results to only have the routing table focus on the
	// #dht.BucketSize closest peers to the current node.
	if len(results) > dht.BucketSize {
		results = results[:dht.BucketSize]
	}
	/*
		fmt.Println("===================================================")
		tFindNode := time.Since(start).Seconds()
		fmt.Println("FINDNODE n.GetId().Id", n.GetId().Id, "  targetID = ", targetID, " tFindNode = ", tFindNode, " Second")
		fmt.Println("===================================================")

		a := new(big.Int).SetBytes(n.GetId().Id).String()
		b := new(big.Int).SetBytes(targetID.Id).String()

		n.log.WithFields(logrus.Fields{
			"time-findnode": tFindNode,
			"targetID":      b,
			"nodeID":        a,
		}).Info()*/

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
	_, loaded := n.peers.Load(string(peerID.Id))
	if !loaded {
		_, err := n.NewPeer(peerID.Address)
		if err != nil {
			responses <- []*internal.ID{}
			return
		}
	}
	//TODO: Need to check if it can just return
	//panic: interface conversion: interface {} is nil, not *p2p.PeerConn
	value, ok := n.peers.Load(string(peerID.Id))
	if !ok {
		fmt.Println("value not found for key: peerID.Id", peerID.Id)
		responses <- []*internal.ID{}
		return
	}
	client = value.(*PeerConn)

	targetProtoID := internal.ID(targetID)

	request := new(Request)
	request.SetMessage(&internal.LookupNodeRequest{Target: &targetProtoID})
	request.SetTimeout(30 * time.Second)

	response, err := client.Request(request)

	if err != nil {
		responses <- []*internal.ID{}
		return
	}

	if response, ok := response.(*internal.LookupNodeResponse); ok {
		responses <- response.Peers
	} else {
		responses <- []*internal.ID{}
	}
}
