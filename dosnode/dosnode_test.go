package dosnode

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/suites"
)

func TestPipeCheckURL(test *testing.T) {
	suite := suites.MustFind("bn256")
	nbParticipants := 1
	id := []byte{1, 2, 3}
	//1)Build a p2p network
	peerEvent := make(chan p2p.P2PMessage, 100)
	defer close(peerEvent)
	p, _ := p2p.CreateP2PNetwork(peerEvent, 0)
	p.SetId(id)
	p.Listen()

	peerEventForDKG := make(chan p2p.P2PMessage, 100)
	defer close(peerEventForDKG)
	p2pDkg, _ := dkg.CreateP2PDkg(p, suite, peerEventForDKG, nbParticipants)
	go p2pDkg.EventLoop()
	p2pDkg.RunDKG()

	d, _ := CreateDosNode(suite, nbParticipants, nil, nil, p2pDkg)
	chUrl := make(chan interface{}, 100)
	defer close(chUrl)
	out1 := d.PipeCheckURL(chUrl)
	QueryId := big.NewInt(1)
	url := "test"
	Timeout := big.NewInt(1)
	Randomness := big.NewInt(1)
	DispatchedGroup := [4]*big.Int{QueryId, QueryId, QueryId, QueryId}
	chUrl <- &onchain.DOSProxyLogUrl{
		QueryId:         QueryId,
		Url:             url,
		Timeout:         Timeout,
		Randomness:      Randomness,
		DispatchedGroup: DispatchedGroup,
	}
	fmt.Println(<-out1)
}
