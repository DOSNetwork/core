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
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

type P2P struct {
	identity dht.ID
	//Map of connection addresses (string) <-> *p2p.PeerClient
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

func (n *P2P) SetId(id []byte) {
	n.identity.Id = id
}

func (n *P2P) GetId() dht.ID {
	return n.identity
}

func (n *P2P) Listen() error {
	//init log
	n.log = logrus.New()
	hook, _ := logrustash.NewHookWithFields("udp", "13.52.16.14:9500", "DOS_node", logrus.Fields{
		"DOS_node_ip": n.identity.Address,
	})
	n.log.Hooks.Add(hook)
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
				NewPeerClient(n, &conn, n.messages)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Failed accepting a connection request:", err)
				log.Fatal(err)
			}
		}
	}()

	return nil
}

func (n *P2P) Broadcast(m proto.Message) {
	n.peers.Range(func(key, value interface{}) bool {
		client := value.(*PeerClient)
		err := client.SendMessage(m)
		if err != nil {
			fmt.Println("P2P Broadcast ", err)
		}
		return true
	})
}
func (n *P2P) CheckPeers() {
	n.peers.Range(func(key, value interface{}) bool {
		client := value.(*PeerClient)
		fmt.Println(client.identity.Address)
		return true
	})
}

func (n *P2P) SendMessageById(id []byte, m proto.Message) (err error) {
	value, loaded := n.peers.Load(string(id))
	if !loaded {
		//TODO : This is a sync call.It should be optimized to async call
		fmt.Println("Can't find node ", id)
		n.FindNodeById(id)
	}
	value, loaded = n.peers.Load(string(id))
	if loaded {
		client := value.(*PeerClient)
		err = client.SendMessage(m)
	} else {
		err = fmt.Errorf("can't find node %s", string(id))
	}
	return
}

func (n *P2P) GetRoutingTable() *dht.RoutingTable {
	return n.routingTable
}

/*
This is a block call
*/

func (n *P2P) NewPeer(addr string) (id []byte, err error) {
	var peer *PeerClient
	var conn net.Conn
	for retry := 0; retry < 10; retry++ {
		//1)Check if this address has been in peers map
		existing := false
		n.peers.Range(func(key, value interface{}) bool {
			client := value.(*PeerClient)
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

		peer, err = NewPeerClient(n, &conn, n.messages)
		if err != nil {
			fmt.Println("NewPeerClient err ", err)
			continue
		}
		n.peers.LoadOrStore(string(peer.identity.Id), peer)
		n.routingTable.Update(peer.identity)

		n.peers.Range(func(key, value interface{}) bool {
			client := value.(*PeerClient)
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

func (n *P2P) FindNodeById(id []byte) []dht.ID {
	targetId := dht.ID{
		Id: id,
	}
	return n.findNode(targetId, dht.BucketSize, 8)
}

func (n *P2P) FindNode(targetID dht.ID, alpha int, disjointPaths int) (results []dht.ID) {
	return n.findNode(targetID, alpha, disjointPaths)
}

type lookupBucket struct {
	pending int
	queue   []dht.ID
}

// FindNode queries all peers this current node acknowledges for the closest peers
// to a specified target ID.
//
// All lookups are done under a number of disjoint lookups in parallel.
//
// Queries at most #ALPHA nodes at a time per lookup, and returns all peer IDs closest to a target peer ID.
func (n *P2P) findNode(targetID dht.ID, alpha int, disjointPaths int) (results []dht.ID) {
	start := time.Now()
	visited := new(sync.Map)

	var lookups []*lookupBucket

	// Start searching for target from #ALPHA peers closest to target by queuing
	// them up and marking them as visited.
	for i, peerID := range n.routingTable.FindClosestPeers(targetID, alpha) {

		visited.Store(peerID.PublicKeyHex(), struct{}{})

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
		left := results[i].Xor(targetID)
		right := results[j].Xor(targetID)
		return left.Less(right)
	})

	// Cut off list of results to only have the routing table focus on the
	// #dht.BucketSize closest peers to the current node.
	if len(results) > dht.BucketSize {
		results = results[:dht.BucketSize]
	}
	fmt.Println("===================================================")
	tFindNode := time.Since(start).Seconds()
	fmt.Println("FINDNODE ", targetID, tFindNode)
	fmt.Println("===================================================")

	a := new(big.Int).SetBytes(n.GetId().Id).String()
	b := new(big.Int).SetBytes(targetID.Id).String()

	n.log.WithFields(logrus.Fields{
		"timecost-findnode": tFindNode,
		"targetID":          b,
		"nodeID":            a,
	}).Info()
	return
}

func (lookup *lookupBucket) performLookup(n *P2P, targetID dht.ID, alpha int, visited *sync.Map) (results []dht.ID) {
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
			peerID := dht.ID(*id)

			if _, seen := visited.LoadOrStore(peerID.PublicKeyHex(), struct{}{}); !seen && targetID.PublicKeyHex() != peerID.PublicKeyHex() {
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

func (n *P2P) queryPeerByID(peerID dht.ID, targetID dht.ID, responses chan []*internal.ID) {
	var client *PeerClient
	_, loaded := n.peers.Load(string(peerID.Id))
	if !loaded {
		_, err := n.NewPeer(peerID.Address)
		if err != nil {
			responses <- []*internal.ID{}
			return
		}
	}
	//TODO: Need to check if it can just return
	//panic: interface conversion: interface {} is nil, not *p2p.PeerClient
	value, ok := n.peers.Load(string(peerID.Id))
	if !ok {
		fmt.Println("value not found for key: peerID.Id", peerID.Id)
		responses <- []*internal.ID{}
		return
	}
	client = value.(*PeerClient)

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
