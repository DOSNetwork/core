package dkg

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"

	//"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	bls "github.com/DOSNetwork/core/sign/bls"
	tbls "github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
)

func buildConn(s, d p2p.P2PInterface, port string, t *testing.T) {
	s.SetPort(port)
	connected, err := s.ConnectTo("0.0.0.0", nil)
	if err != nil {
		t.Errorf("TestRequest ,err %s", err)
	}
	if !bytes.Equal(connected, d.GetID()) {
		t.Errorf("TestRequest ,Expected [% x] Actual [% x]", d.GetID(), connected)
	}
}

func TestPDKG(t *testing.T) {
	var groupIds [][]byte
	groupIds = append(groupIds, []byte("Participant0"))
	groupIds = append(groupIds, []byte("Participant1"))
	groupIds = append(groupIds, []byte("Participant2"))
	os.Setenv("PUBLICIP", "0.0.0.0")
	p1, err := p2p.CreateP2PNetwork(groupIds[0], "9905", 0)
	if err != nil {
		t.Errorf("CreateP2PNetwork err %s", err)
	}
	p1.Listen()
	p2, err := p2p.CreateP2PNetwork(groupIds[1], "9906", 0)
	if err != nil {
		t.Errorf("CreateP2PNetwork err %s", err)
	}
	p2.Listen()
	p3, err := p2p.CreateP2PNetwork(groupIds[2], "9907", 0)
	if err != nil {
		t.Errorf("CreateP2PNetwork err %s", err)
	}
	p3.Listen()
	//Setup p2p network
	buildConn(p1, p2, "9906", t)
	buildConn(p1, p3, "9907", t)
	buildConn(p2, p1, "9905", t)
	buildConn(p2, p3, "9907", t)
	buildConn(p3, p1, "9905", t)
	buildConn(p3, p2, "9906", t)

	suite := suites.MustFind("bn256")
	pdkg1 := NewPDKG(p1, suite)
	pdkg2 := NewPDKG(p2, suite)
	pdkg3 := NewPDKG(p3, suite)
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(groupID string) {
			defer wg.Done()
			ctx := context.Background()
			out1, errc1, err := pdkg1.Grouping(ctx, groupID, groupIds)
			if err != nil {
				t.Errorf("Grouping err %s", err)
			}
			out2, errc2, err := pdkg2.Grouping(ctx, groupID, groupIds)
			if err != nil {
				t.Errorf("Grouping err %s", err)
			}
			out3, errc3, err := pdkg3.Grouping(ctx, groupID, groupIds)
			if err != nil {
				t.Errorf("Grouping err %s", err)
			}
			_ = <-out1
			_ = <-out2
			_ = <-out3
			//fmt.Println("out1 ", <-out1)
			//fmt.Println("out2 ", <-out2)
			//fmt.Println("out3 ", <-out3)
			pp := pdkg1.GetGroupPublicPoly(groupID)
			pubKeyCoor, _ := decodePubKey(pp.Commit())
			fmt.Println("pubKeyCoor ", pubKeyCoor)

			content := []byte{'a', 'b'}
			var sigShares [][]byte
			var sig []byte
			sig, err = tbls.Sign(suite, pdkg1.GetShareSecurity(groupID), content)
			if err != nil {
				t.Errorf("Sign err %s", err)
			}
			sigShares = append(sigShares, sig)
			sig, err = tbls.Sign(suite, pdkg2.GetShareSecurity(groupID), content)
			if err != nil {
				t.Errorf("Sign err %s", err)
			}
			sigShares = append(sigShares, sig)
			sig, err = tbls.Sign(suite, pdkg3.GetShareSecurity(groupID), content)
			if err != nil {
				t.Errorf("Sign err %s", err)
			}
			sigShares = append(sigShares, sig)

			sig, err = tbls.Recover(
				suite,
				pdkg1.GetGroupPublicPoly(groupID),
				content,
				sigShares,
				2,
				3)
			if err != nil {
				t.Errorf("Recover err %s", err)
			}
			if err = bls.Verify(
				suite,
				pdkg3.GetGroupPublicPoly(groupID).Commit(),
				content,
				sig); err != nil {
				t.Errorf("Verify err %s", err)
			}
		L1:
			for {
				select {
				case err, ok := <-errc1:
					if !ok {
						break L1
					} else {
						fmt.Println("errc1 ", err)
					}
				}
			}
		L2:
			for {
				select {
				case err, ok := <-errc2:
					if !ok {
						break L2
					} else {
						fmt.Println("errc2 ", err)
					}
				}
			}
			fmt.Println("test done")
		L3:
			for {
				select {
				case err, ok := <-errc3:
					if !ok {
						break L3
					} else {
						fmt.Println("errc3 ", err)
					}
				}
			}
			fmt.Println("test done")
		}(strconv.Itoa(i))
	}
	wg.Wait()
}
