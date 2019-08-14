package discover

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/hashicorp/serf/serf"
)

// Membership represents the p2p network
type Membership interface {
	Join(Ip []string) (num int, err error)
	Leave()
	Lookup(id []byte) (addr string)
	NumOfPeers() int
	MembersIP() (addr []net.IP)
	MembersID() (list [][]byte)
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
	serfNet.serf, err = serf.Create(conf)

	return serfNet, err
}

type serfNet struct {
	serf *serf.Serf
}

//NumOfPeers return the length of members
func (s *serfNet) NumOfPeers() int {
	return len(s.serf.Members())
}

// Join joins an existing Serf cluster. Returns the number of nodes
// successfully contacted.
func (s *serfNet) Join(bootstrapIp []string) (num int, err error) {
	num, err = s.serf.Join(bootstrapIp, true)
	return
}

// Join leaves an Serf cluster.
func (s *serfNet) Leave() {
	s.Leave()
	return
}

// Lookup return the IP address of the given ID
func (s *serfNet) Lookup(id []byte) (addr string) {
	members := s.serf.Members()
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
		if nodeId == string(id) {
			fmt.Println("ID ", []byte(nodeId), " ", members[i].Addr.String()+":"+port)

			return members[i].Addr.String() + ":" + port
		}
	}
	return
}

// MembersIP return the all IP address of an existing cluster
func (s *serfNet) MembersIP() (addr []net.IP) {
	//fmt.Println("members len ", len(s.serf.Members()))

	members := s.serf.Members()
	for i := 0; i < len(members); i++ {

		if members[i].Name != s.serf.LocalMember().Name {
			//fmt.Println("localMember ", []byte(s.serf.LocalMember().Name), "members[i].Name ", []byte(members[i].Name), " status ", members[i].Status, " addr ", members[i].Addr)

			addr = append(addr, members[i].Addr)
		}
	}
	return
}

// MembersID return the all ID of an existing cluster
func (s *serfNet) MembersID() (list [][]byte) {
	members := s.serf.Members()
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
