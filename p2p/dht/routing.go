package dht

import (
	"container/list"
	"sort"
	"sync"

	"github.com/DOSNetwork/core/p2p/internal"
)

// BucketSize defines the NodeID, Key, and routing table data structures.
const BucketSize = 16

// RoutingTable contains one bucket list for lookups.
type RoutingTable struct {
	// Current node's ID.
	self internal.ID

	buckets []*Bucket
}
type ID internal.ID

// Bucket holds a list of contacts of this node.
type Bucket struct {
	*list.List
	mutex *sync.RWMutex
}

// NewBucket is a Factory method of Bucket, contains an empty list.
func NewBucket() *Bucket {
	return &Bucket{
		List:  list.New(),
		mutex: &sync.RWMutex{},
	}
}

// CreateRoutingTable is a Factory method of RoutingTable containing empty buckets.
func CreateRoutingTable(id internal.ID) *RoutingTable {
	table := &RoutingTable{
		self:    id,
		buckets: make([]*Bucket, len(id.Id)*8),
	}
	for i := 0; i < len(id.Id)*8; i++ {
		table.buckets[i] = NewBucket()
	}

	//table.Update(id)

	return table
}

// Self returns the ID of the node hosting the current routing table instance.
func (t *RoutingTable) Self() internal.ID {
	return t.self
}

// Update moves a peer to the front of a bucket in the routing table.
func (t *RoutingTable) Update(target internal.ID) {
	if len(t.self.Id) != len(target.Id) {
		return
	}

	bucketID := PrefixLen(XorID(target, t.self))
	bucket := t.Bucket(bucketID)

	var element *list.Element

	// Find current node in bucket.
	bucket.mutex.Lock()

	for e := bucket.Front(); e != nil; e = e.Next() {
		if Equals(e.Value.(internal.ID), target) {
			element = e
			break
		}
	}

	if element == nil {
		// Populate bucket if its not full.
		if bucket.Len() <= BucketSize {
			bucket.PushFront(target)
		}
	} else {
		bucket.MoveToFront(element)
	}

	bucket.mutex.Unlock()
}

// GetPeers returns a randomly-ordered, unique list of all peers within the routing network (excluding itself).
func (t *RoutingTable) GetPeers() (peers []internal.ID) {
	visited := make(map[string]struct{})
	visited[PublicKeyHex(t.self)] = struct{}{}

	for _, bucket := range t.buckets {
		bucket.mutex.RLock()

		for e := bucket.Front(); e != nil; e = e.Next() {
			id := e.Value.(internal.ID)
			if _, seen := visited[PublicKeyHex(id)]; !seen {
				peers = append(peers, id)
				visited[PublicKeyHex(id)] = struct{}{}
			}
		}

		bucket.mutex.RUnlock()
	}

	return
}

// GetPeerAddresses returns a unique list of all peer addresses within the routing network.
func (t *RoutingTable) GetPeerAddresses() (peers []string) {
	visited := make(map[string]struct{})
	visited[PublicKeyHex(t.self)] = struct{}{}

	for _, bucket := range t.buckets {
		bucket.mutex.RLock()

		for e := bucket.Front(); e != nil; e = e.Next() {
			id := e.Value.(internal.ID)
			if _, seen := visited[PublicKeyHex(id)]; !seen {
				peers = append(peers, id.Address)
				visited[PublicKeyHex(id)] = struct{}{}
			}
		}

		bucket.mutex.RUnlock()
	}

	return
}

// RemovePeer removes a peer from the routing table with O(bucket_size) time complexity.
func (t *RoutingTable) RemovePeer(target internal.ID) bool {
	bucketID := PrefixLen(XorID(target, t.self))
	bucket := t.Bucket(bucketID)

	bucket.mutex.Lock()

	for e := bucket.Front(); e != nil; e = e.Next() {
		if Equals(e.Value.(internal.ID), target) {
			bucket.Remove(e)

			bucket.mutex.Unlock()
			return true
		}
	}

	bucket.mutex.Unlock()

	return false
}

// PeerExists checks if a peer exists in the routing table with O(bucket_size) time complexity.
func (t *RoutingTable) PeerExists(target internal.ID) bool {
	bucketID := PrefixLen(XorID(target, t.self))
	bucket := t.Bucket(bucketID)

	bucket.mutex.Lock()

	defer bucket.mutex.Unlock()

	for e := bucket.Front(); e != nil; e = e.Next() {
		if Equals(e.Value.(internal.ID), target) {
			return true
		}
	}

	return false
}

// FindClosestPeers returns a list of k(count) peers with smallest XorID distance.
func (t *RoutingTable) FindClosestPeers(target internal.ID, count int) (peers []internal.ID) {
	if len(t.self.Id) != len(target.Id) {
		return []internal.ID{}
	}

	bucketID := PrefixLen(XorID(target, t.self))
	bucket := t.Bucket(bucketID)

	bucket.mutex.RLock()

	for e := bucket.Front(); e != nil; e = e.Next() {
		peers = append(peers, e.Value.(internal.ID))
	}

	bucket.mutex.RUnlock()

	for i := 1; len(peers) < count && (bucketID-i >= 0 || bucketID+i < len(t.self.Id)*8); i++ {
		if bucketID-i >= 0 {
			other := t.Bucket(bucketID - i)
			other.mutex.RLock()
			for e := other.Front(); e != nil; e = e.Next() {
				peers = append(peers, e.Value.(internal.ID))
			}
			other.mutex.RUnlock()
		}

		if bucketID+i < len(t.self.Id)*8 {
			other := t.Bucket(bucketID + i)
			other.mutex.RLock()
			for e := other.Front(); e != nil; e = e.Next() {
				peers = append(peers, e.Value.(internal.ID))
			}
			other.mutex.RUnlock()
		}
	}

	// Sort peers by XorID distance.
	sort.Slice(peers, func(i, j int) bool {
		left := XorID(peers[i], target)
		right := XorID(peers[j], target)
		return Less(left, right)
	})

	if len(peers) > count {
		peers = peers[:count]
	}

	return peers
}

// Bucket returns a specific Bucket by ID.
func (t *RoutingTable) Bucket(id int) *Bucket {
	if id >= 0 && id < len(t.buckets) {
		return t.buckets[id]
	}
	return nil
}
