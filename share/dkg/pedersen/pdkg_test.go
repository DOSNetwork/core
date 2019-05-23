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

func initP2P(id []byte, port string, t *testing.T) (p2p.P2PInterface, chan interface{}, error) {
	os.Setenv("PUBLICIP", "0.0.0.0")
	p, err := p2p.CreateP2PNetwork(id, port, 0)
	p.Listen()
	fmt.Println("initP2P", string(id))
	peersToBuf, _ := p.SubscribeEvent(1, PublicKey{}, Deal{}, Responses{})

	bufToNode := make(chan interface{})
	go func(id []byte) {
		sessionPubKeys := map[string][]*PublicKey{}
		sessionDeals := map[string][]*Deal{}
		sessionResps := map[string][]*Response{}
		sessionReqPubs := map[string]*ReqPubs{}
		sessionReqDeals := map[string]*ReqDeals{}
		sessionReResps := map[string]*ReResps{}

		for {
			select {
			case msg, ok := <-peersToBuf:
				if !ok {
					return
				}
				switch content := msg.Msg.Message.(type) {
				case *PublicKey:
					sessionPubKeys[content.SessionId] = append(sessionPubKeys[content.SessionId], content)
					fmt.Println(string(id), " Got sessionPubKeys ", len(sessionPubKeys[content.SessionId]))
					p.Reply(msg.Sender, msg.RequestNonce, &PublicKey{})

					logger.Event("PublicKeyFromPeer", map[string]interface{}{"GroupID": content.SessionId})
					if sessionReqPubs[content.SessionId] != nil {
						if len(sessionPubKeys[content.SessionId]) == sessionReqPubs[content.SessionId].numOfPubs {
							pubkeys := sessionPubKeys[content.SessionId]
							sessionPubKeys[content.SessionId] = nil
							sessionReqPubs[content.SessionId].Pubkeys <- pubkeys

							close(sessionReqPubs[content.SessionId].Pubkeys)
							sessionReqPubs[content.SessionId] = nil
						}
					}
				case *Deal:
					sessionDeals[content.SessionId] = append(sessionDeals[content.SessionId], content)
					fmt.Println(string(id), "Got Deal ", len(sessionDeals[content.SessionId]), content.SessionId)
					p.Reply(msg.Sender, msg.RequestNonce, &Deal{})

					if sessionReqDeals[content.SessionId] != nil {
						if len(sessionDeals[content.SessionId]) == sessionReqDeals[content.SessionId].numOfDeals {
							deals := sessionDeals[content.SessionId]
							sessionDeals[content.SessionId] = nil
							sessionReqDeals[content.SessionId].Reply <- deals
							close(sessionReqDeals[content.SessionId].Reply)
							sessionReqDeals[content.SessionId] = nil
						}
					}
				case *Responses:
					sessionResps[content.SessionId] = append(sessionResps[content.SessionId], content.Response...)
					fmt.Println(string(id), "Got Deal ", len(sessionResps[content.SessionId]), content.SessionId)
					p.Reply(msg.Sender, msg.RequestNonce, &Responses{})

					if sessionReResps[content.SessionId] != nil {
						if len(sessionResps[content.SessionId]) == sessionReResps[content.SessionId].numOfResps {
							resps := sessionResps[content.SessionId]
							sessionResps[content.SessionId] = nil
							sessionReResps[content.SessionId].Reply <- resps
							close(sessionReResps[content.SessionId].Reply)
							sessionReResps[content.SessionId] = nil
						}
					}
				}
			case msg, ok := <-bufToNode:
				if !ok {
					return
				}
				switch content := msg.(type) {
				case *ReqPubs:
					sessionReqPubs[content.SessionId] = content

					if len(sessionPubKeys[content.SessionId]) == content.numOfPubs {
						pubkeys := sessionPubKeys[content.SessionId]
						sessionPubKeys[content.SessionId] = nil
						sessionReqPubs[content.SessionId].Pubkeys <- pubkeys
						close(content.Pubkeys)
					}
				case *ReqDeals:
					sessionReqDeals[content.SessionId] = content

					if len(sessionDeals[content.SessionId]) == content.numOfDeals {
						deals := sessionDeals[content.SessionId]
						sessionDeals[content.SessionId] = nil
						sessionReqDeals[content.SessionId].Reply <- deals
						close(content.Reply)
					}
				case *ReResps:
					sessionReResps[content.SessionId] = content

					if len(sessionResps[content.SessionId]) == content.numOfResps {
						deals := sessionResps[content.SessionId]
						sessionResps[content.SessionId] = nil
						sessionReResps[content.SessionId].Reply <- deals
						close(content.Reply)
					}
				}
			}
		}
	}(id)
	return p, bufToNode, err
}

func TestExchange(t *testing.T) {
	log.Init([]byte("test"))
	logger = log.New("module", "dkg")
	var groupIds [][]byte
	groupIds = append(groupIds, []byte("Participant0"))
	groupIds = append(groupIds, []byte("Participant1"))
	groupIds = append(groupIds, []byte("Participant2"))

	var nodes []p2p.P2PInterface
	var toPeersPubs []chan interface{}
	p1, toPeersPub1, _ := initP2P(groupIds[0], "9901", t)
	nodes = append(nodes, p1)
	toPeersPubs = append(toPeersPubs, toPeersPub1)
	p2, toPeersPub2, _ := initP2P(groupIds[1], "9902", t)
	nodes = append(nodes, p2)
	toPeersPubs = append(toPeersPubs, toPeersPub2)
	p3, toPeersPub3, _ := initP2P(groupIds[2], "9903", t)
	nodes = append(nodes, p3)
	toPeersPubs = append(toPeersPubs, toPeersPub3)
	suite := suites.MustFind("bn256")
	//Setup p2p network
	buildConn(p1, p2, "9902", t)
	buildConn(p1, p3, "9903", t)
	buildConn(p2, p1, "9901", t)
	buildConn(p2, p3, "9903", t)
	buildConn(p3, p1, "9901", t)
	buildConn(p3, p2, "9902", t)

	//Exchange pubkey
	var wg sync.WaitGroup
	wg.Add(len(nodes))

	for i, network := range nodes {
		go func(network p2p.P2PInterface, index int) {
			defer wg.Done()
			ctx := context.Background()
			dkgc, errc := exchangePub(ctx, suite, toPeersPubs[index], groupIds, network, "session1")
			dkgCetifiedc, errc2 := processDeal(ctx, dkgc, toPeersPubs[index], groupIds, network, "session1")

			for {
				select {
				case <-ctx.Done():
				case dkg, ok := <-dkgCetifiedc:
					if !ok {
						fmt.Println("resp !ok")
						return
					}
					if !dkg.Certified() {
						t.Errorf("Not Certified")
					} else {
						fmt.Println("dkg.Certified")
					}
				case err, ok := <-errc:
					if ok {
						fmt.Println("errc1 ", err)
						t.Errorf("exchangePub err %s!!!!!!!!!!!!!!!!", err)
					}
				case err, ok := <-errc2:
					if ok {
						fmt.Println("errc2 ", err)
						t.Errorf("exchangePub err %s!!!!!!!!!!!!!!!!", err)
					}
				}
			}
		}(network, i)
	}
	wg.Wait()
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
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(groupID string) {
			defer wg.Done()
			ctx := context.Background()
			out1, _, err := pdkg1.Grouping(ctx, groupID, groupIds)
			if err != nil {
				t.Errorf("Grouping err %s", err)
			}
			out2, _, err := pdkg2.Grouping(ctx, groupID, groupIds)
			if err != nil {
				t.Errorf("Grouping err %s", err)
			}
			out3, _, err := pdkg3.Grouping(ctx, groupID, groupIds)
			if err != nil {
				t.Errorf("Grouping err %s", err)
			}
			fmt.Println("out1 ", <-out1)
			fmt.Println("out2 ", <-out2)
			fmt.Println("out3 ", <-out3)
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
			fmt.Println("test done")
		}(strconv.Itoa(i))
	}
	wg.Wait()
}
