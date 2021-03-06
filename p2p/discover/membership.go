package discover

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"github.com/hashicorp/serf/serf"
)

// Membership represents the p2p network
type Membership interface {
	Join(Ip []string) (num int, err error)
	Leave()
	Listen(ctx context.Context, outch chan P2PEvent)
	Lookup(id []byte) (addr string)
	NumOfPeers() int
	IsAlive() bool
	MembersIP() (addr []net.IP)
	MembersID() (list [][]byte)
}

type P2PEvent struct {
	EventType string
	NodeID    string
	Addr      net.IP
}

//NewSerfNet creates a Serf implementation
func NewSerfNet(Addr net.IP, id, port string) (Membership, error) {
	var err error
	serfNet := &serfNet{}
	conf := serf.DefaultConfig()
	conf.Init()
	conf.LogOutput = ioutil.Discard
	conf.MemberlistConfig.LogOutput = ioutil.Discard
	conf.MemberlistConfig.AdvertiseAddr = Addr.String()
	conf.NodeName = id + port
	conf.MemberlistConfig.IndirectChecks = 21             // Use 3 nodes for the indirect ping
	conf.MemberlistConfig.ProbeTimeout = 1 * time.Second  // Reasonable RTT time for LAN
	conf.MemberlistConfig.ProbeInterval = 5 * time.Second // Failure check every second
	conf.MemberlistConfig.DisableTcpPings = false         // TCP pings are safe, even with mixed versions
	conf.MemberlistConfig.AwarenessMaxMultiplier = 16     // Probe interval backs off to 8 seconds

	eventCh := make(chan serf.Event, 4)
	conf.EventCh = eventCh
	serfNet.eventch = eventCh
	serfNet._serf, err = serf.Create(conf)
	return serfNet, err
}

type serfNet struct {
	_serf   *serf.Serf
	eventch chan serf.Event
}

//NumOfPeers return the length of members
func (s *serfNet) Listen(ctx context.Context, outch chan P2PEvent) {
	defer fmt.Println("[P2P] End MemberList Listen")

	for event := range s.eventch {
		if e, ok := event.(serf.MemberEvent); ok {

			for _, member := range e.Members {
				nodeId := member.Name[:20]
				select {
				case <-ctx.Done():
				case outch <- P2PEvent{EventType: event.String(), NodeID: nodeId, Addr: member.Addr}:
				}
			}
		}
		//if event.EventType() == serf.EventQuery {
		//			query := event.(*serf.Query)
		//fmt.Println("[Gossip ] query", query)
		//}
	}
}

func (s *serfNet) IsAlive() bool {
	state := s._serf.State()
	if state != serf.SerfAlive {
		fmt.Println("[Gossip ] status ", state)
		return false
	}
	return true
}

//NumOfPeers return the length of members
func (s *serfNet) NumOfPeers() int {
	members := s._serf.Members()
	n := 0
	for i := 0; i < len(members); i++ {
		if members[i].Status == serf.StatusAlive && members[i].Name != s._serf.LocalMember().Name {
			n++
		}

	}
	return n
}

// Join joins an existing Serf cluster. Returns the number of nodes
// successfully contacted.
func (s *serfNet) Join(bootstrapIp []string) (num int, err error) {
	num, err = s._serf.Join(bootstrapIp, true)
	return
}

// Join leaves an Serf cluster.
func (s *serfNet) Leave() {
	s._serf.Leave()
	close(s.eventch)
	return
}

// Lookup return the IP address of the given ID
func (s *serfNet) Lookup(id []byte) (addr string) {
	members := s._serf.Members()
	var nodeId, port string
	for i := 0; i < len(members); i++ {
		if len(members[i].Name) < 20 {
			continue
		}
		nodeId = members[i].Name[:20]
		if len(members[i].Name) == 20 {
			port = "9501"
			continue
		} else {
			port = members[i].Name[20:]
		}
		if nodeId == string(id) && members[i].Status == serf.StatusAlive {
			fmt.Println("ID ", []byte(nodeId), " ", members[i].Addr.String()+":"+port)
			return members[i].Addr.String() + ":" + port
		}
	}
	return
}

// MembersIP return the all IP address of an existing cluster
func (s *serfNet) MembersIP() (addr []net.IP) {

	members := s._serf.Members()
	for i := 0; i < len(members); i++ {
		//fmt.Println("localMember ", []byte(s.serf.LocalMember().Name), "members[i].Name ", []byte(members[i].Name), " status ", members[i].Status, " addr ", members[i].Addr)
		if members[i].Name != s._serf.LocalMember().Name {
			addr = append(addr, members[i].Addr)
		}
	}
	return
}

// MembersID return the all ID of an existing cluster
func (s *serfNet) MembersID() (list [][]byte) {
	members := s._serf.Members()
	for i := 0; i < len(members); i++ {
		var port string
		if len(members[i].Name) < 20 {
			continue
		}
		if len(members[i].Name) == 20 {
			port = "9501"
			continue
		} else {
			port = members[i].Name[20:]
		}
		fmt.Println(members[i].Addr.String() + ":" + port)
		fmt.Println("len ", len(members[i].Name))
		if members[i].Status == serf.StatusAlive {
			list = append(list, []byte(members[i].Name))
		}

	}
	return
}
