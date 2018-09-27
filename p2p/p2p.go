package p2p

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dedis/kyber"

	"github.com/golang/protobuf/proto"

	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/p2p/internal"
	"github.com/DOSNetwork/core/suites"
)

type P2P struct {
	identity		dht.ID
	//Map of connection addresses (string) <-> *p2p.PeerClient
	peers 			*sync.Map
	// Channels are thread safe
	messageChan	 	chan P2PMessage
	suite       	suites.Suite
	port        	int
	secKey      	kyber.Scalar
	pubKey			kyber.Point
	routingTable	*dht.RoutingTable
}

func (n *P2P) GetId() dht.ID {
	return n.identity
}

func (n *P2P) Listen() error {
	ip, err := getLocalIp()
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

	n.identity = dht.CreateID(ip, pubKeyBytes)
	n.routingTable = dht.CreateRoutingTable(n.identity)

	fmt.Println("listen on: ", ip, " id: ", n.identity.Id)
	// Handle new clients.

	go func() {
		for {
			if conn, err := listener.Accept(); err == nil {
				//Create a peer client
				n.CreatePeer("", &conn)
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

		prepared, err := client.PrepareMessage(m)
		if err != nil {
			log.Fatal(err)
		}

		prepared.RequestNonce = atomic.AddUint64(&client.RequestNonce, 1)
		if err := client.SendPackage(prepared); err != nil {
			log.Fatal(err)
		}

		client.SendPackage(prepared)
		return true
	})
}

func (n *P2P) SendMessageById(id []byte, m proto.Message) {
	value, loaded := n.peers.Load(string(id))
	if loaded {
		client := value.(*PeerClient)

		prepared, err := client.PrepareMessage(m)
		if err != nil {
			log.Fatal(err)
		}

		prepared.RequestNonce = atomic.AddUint64(&client.RequestNonce, 1)
		if err := client.SendPackage(prepared); err != nil {
			log.Fatal(err)
		}

		client.SendPackage(prepared)
	}
}

func (n *P2P) GetTunnel() chan P2PMessage {
	return n.messageChan
}

func (n *P2P) GetRoutingTable() *dht.RoutingTable {
	return n.routingTable
}

func (n *P2P) CreatePeer(addr string, c *net.Conn) {
	fmt.Println(n.identity.Address, "Create peer clients")
	peer := &PeerClient{
		conn:   c,
		p2pnet: n,
	}
	if addr != "" {
		peer.Dial(addr)
	}
	peer.rw = bufio.NewReadWriter(bufio.NewReader(*peer.conn), bufio.NewWriter(*peer.conn))
	//n.peers.LoadOrStore(peer.id, peer)
	peer.messageChan = n.messageChan

	peer.wg.Add(1)
	//fmt.Println("InitClient id ", peer.id)
	go peer.HandlePackages()
	peer.SayHi()
	return
}

func (n *P2P) FindNodeById(id []byte) []dht.ID {
	targetId := dht.ID{
		Id: id,
	}
	return n.findNode(targetId, dht.BucketSize, 8)
}

func (n *P2P) FindNode(targetID dht.ID, alpha int, disjointPaths int) (results []dht.ID) {
	return n.findNode(targetID,alpha,disjointPaths)
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
		n.CreatePeer(peerID.Address, nil)
	}
	value, _ := n.peers.Load(string(peerID.Id))
	client = value.(*PeerClient)

	client.wg.Wait()

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
