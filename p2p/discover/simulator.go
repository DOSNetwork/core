package discover

import (
	"context"
	"net"
)

//NewSerfNet creates a Serf implementation
func NewSimulator() (Membership, error) {
	return &simulator{}, nil
}

type simulator struct {
}

func (s *simulator) IsAlive() bool {
	return true
}

//NumOfPeers return the length of members
func (s *simulator) NumOfPeers() int {
	return 1
}

// Join joins an existing Serf cluster. Returns the number of nodes
// successfully contacted.
func (s *simulator) Join(bootstrapIp []string) (num int, err error) {
	return
}

// Join leaves an Serf cluster.
func (s *simulator) Leave() {
	return
}
func (s *simulator) Listen(ctx context.Context, outch chan P2PEvent) {
}

// Lookup return the IP address of the given ID
func (s *simulator) Lookup(id []byte) (addr string) {
	if string(id) == "b" {
		return "127.0.0.1:9502"
	} else if string(id) == "a" {
		return "127.0.0.1:9501"
	}

	return
}

// MembersIP return the all IP address of an existing cluster
func (s *simulator) MembersIP() (addr []net.IP) {

	return
}

// MembersID return the all ID of an existing cluster
func (s *simulator) MembersID() (list [][]byte) {

	return
}
