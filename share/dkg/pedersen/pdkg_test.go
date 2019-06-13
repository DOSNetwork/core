package dkg

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/DOSNetwork/core/log"

	//"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	bls "github.com/DOSNetwork/core/sign/bls"
	tbls "github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
)

func buildConn(s, d p2p.P2PInterface, port string, t *testing.T) {
	//Println(string(s.GetID()), " connect ", string(d.GetID()), " port ", port)
	oldPort := s.GetPort()
	s.SetPort(port)
	connected, err := s.ConnectTo("0.0.0.0", nil)
	if err != nil {
		t.Errorf("TestRequest ,err %s", err)
	}
	if !bytes.Equal(connected, d.GetID()) {
		t.Errorf("TestRequest ,Expected [% x] Actual [% x]", d.GetID(), connected)
	}
	s.SetPort(oldPort)
}
func setUpP2P(id []byte, port string, t *testing.T) (p p2p.P2PInterface) {
	var err error
	p, err = p2p.CreateP2PNetwork(id, port, 0)
	if err != nil {
		t.Errorf("CreateP2PNetwork err %s", err)
	}
	p.Listen()
	return
}

func buildPdkg(size int, t *testing.T) ([]PDKGInterface, [][]byte) {
	var p []p2p.P2PInterface
	var d []PDKGInterface
	var groupIds [][]byte
	portStart := 9905
	id := "Participant"
	suite := suites.MustFind("bn256")
	for i := 0; i < size; i++ {
		nodePort := strconv.Itoa(portStart + i)
		nodeId := []byte(id + strconv.Itoa(i))
		groupIds = append(groupIds, nodeId)
		pi := setUpP2P(nodeId, nodePort, t)
		p = append(p, pi)
		d = append(d, NewPDKG(pi, suite))
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i != j {
				buildConn(p[i], p[j], p[j].GetPort(), t)
			}
		}
	}
	return d, groupIds
}

func TestPDKG(t *testing.T) {

	os.Setenv("APPSESSION", "test")
	os.Setenv("LOGIP", "163.172.36.173:9500")
	os.Setenv("PUBLICIP", "0.0.0.0")
	id := []byte("Participant0")
	log.Init(id[:])

	pdkgs, groupIds := buildPdkg(11, t)

	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 1; i++ {
		go func(groupID string, pdkgs []PDKGInterface, groupIds [][]byte) {
			defer wg.Done()
			ctx := context.Background()
			var wgPdkg sync.WaitGroup
			wgPdkg.Add(len(pdkgs))
			for _, pdkg := range pdkgs {
				go func(pdkg PDKGInterface) {
					defer wgPdkg.Done()
					out, errc, err := pdkg.Grouping(ctx, groupID, groupIds)
					if err != nil {
						t.Errorf("Grouping err %s", err)
					}
					select {
					case <-out:
					case err := <-errc:
						t.Errorf("Grouping err %s", err)
					}
				}(pdkg)
			}
			wgPdkg.Wait()

			content := []byte{'a', 'b'}
			var sigShares [][]byte
			var sig []byte
			var err error
			for _, pdkg := range pdkgs {
				sig, err = tbls.Sign(suite, pdkg.GetShareSecurity(groupID), content)
				if err != nil {
					t.Errorf("Sign err %s", err)
				}
				sigShares = append(sigShares, sig)
			}

			sig, err = tbls.Recover(suite, pdkgs[0].GetGroupPublicPoly(groupID), content, sigShares, len(pdkgs)/2+1, len(pdkgs))
			if err != nil {
				t.Errorf("Recover err %s", err)
			}
			if err = bls.Verify(suite, pdkgs[1].GetGroupPublicPoly(groupID).Commit(), content, sig); err != nil {
				t.Errorf("Verify err %s", err)
			}

			fmt.Println("test done")
		}(strconv.Itoa(i), pdkgs, groupIds)
	}
	wg.Wait()
}
