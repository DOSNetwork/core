package network

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/hashicorp/serf/serf"
)

type Network interface {
	Join(Ip []string) (err error)
	Leave()
	Lookups(id []byte) (addr net.IP)
	NumPeers() int
	GetOtherMembersIP() (addr []net.IP)
}

func NewSerfNet(Addr net.IP, id []byte) (n Network, err error) {
	serfNet := &SerfNet{}
	conf := serf.DefaultConfig()
	conf.Init()
	conf.LogOutput = ioutil.Discard
	conf.MemberlistConfig.AdvertiseAddr = Addr.String()
	conf.NodeName = string(id)
	serfNet.serf, err = serf.Create(conf)

	return serfNet, nil
}

type SerfNet struct {
	serf *serf.Serf
}

func (s *SerfNet) NumPeers() int {
	return len(s.serf.Members())
}

func (s *SerfNet) Join(bootstrapIp []string) (err error) {
	_, err = s.serf.Join(bootstrapIp, true)
	return
}

func (s *SerfNet) Leave() {
	s.Leave()
	return
}

func (s *SerfNet) Lookups(id []byte) (addr net.IP) {
	members := s.serf.Members()
	for i := 0; i < len(members); i++ {
		if members[i].Name == string(id) {
			return members[i].Addr
		}
	}
	return
}

func (s *SerfNet) GetOtherMembersIP() (addr []net.IP) {
	fmt.Println("members len ", len(s.serf.Members()))

	members := s.serf.Members()
	for i := 0; i < len(members); i++ {

		if members[i].Name != s.serf.LocalMember().Name && members[i].Status == serf.StatusAlive {
			fmt.Println("localMember ", []byte(s.serf.LocalMember().Name), "members[i].Name ", []byte(members[i].Name), " status ", members[i].Status, " addr ", members[i].Addr)

			addr = append(addr, members[i].Addr)
		}
	}
	return
}
