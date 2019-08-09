package discover

import (
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
	if string(id) == "b" {
		return "127.0.0.1:9502"
	} else if string(id) == "a" {
		return "127.0.0.1:9501"
	}
	members := s.serf.Members()
	for i := 0; i < len(members); i++ {
		nodeId := members[i].Name[:42]
		port := members[i].Name[42:]
		if nodeId == string(id) {
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
		if members[i].Status == serf.StatusAlive {
			list = append(list, []byte(members[i].Name))
		}

	}
	return
}
