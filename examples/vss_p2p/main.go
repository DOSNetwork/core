package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/examples/vss_p2p/internalMsg"
	"github.com/DOSNetwork/core/p2p"

	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"

	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/golang/protobuf/proto"

	"github.com/looplab/fsm"
)

var suite = suites.MustFind("bn256")
var startTime time.Time
var endTime time.Time

func genPair() (kyber.Scalar, vss.PublicKey) {

	secret := suite.Scalar().Pick(suite.RandomStream())
	public := vss.PublicKey{}
	err := public.SetPoint(suite, suite.Point().Mul(secret, nil))
	if err != nil {
		log.Fatal(err)
	}
	return secret, public
}

type Dealer struct {
	secret         kyber.Scalar
	pubKey         vss.PublicKey
	priKey         kyber.Scalar
	responses      []*vss.Response
	sigShares      [][]byte
	verifiersPub   []kyber.Point
	nbVerifiers    int
	vssThreshold   int
	totalResponses int
	currentResp    uint32
	vssDealer      *vss.Dealer
	pubPoly        *share.PubPoly
	network        *p2p.P2PInterface
	FSM            *fsm.FSM
	signMap        *sync.Map
	contentMap     *sync.Map
}

type Verifier struct {
	dealerPubKey vss.PublicKey
	pubKey       vss.PublicKey
	priKey       kyber.Scalar
	verifiersPub []kyber.Point
	vssVerifier  *vss.Verifier
	network      *p2p.P2PInterface
	FSM          *fsm.FSM
}

func dataFetch(url string) ([]byte, error) {
	sTime := time.Now()
	client := &http.Client{Timeout: 60 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body.Close()

	fmt.Println("fetched data: ", string(body))
	log.Println("dataFetch End:	took", time.Now().Sub(sTime))

	return body, nil

}

func (dealer *Dealer) afterReceivePubkey(e *fsm.Event) {
	//Wait for all verifiers
	if len(dealer.verifiersPub) == dealer.nbVerifiers {
		//Send dealer public key to all verifiers
		(*dealer.network).Broadcast(&dealer.pubKey)

		//Send all of verifier's pubkey to all verifiers
		publicKeys := vss.PublicKeys{}
		err := publicKeys.SetPointArray(suite, dealer.verifiersPub)
		if err != nil {
			log.Fatal(err)
		}
		(*dealer.network).Broadcast(&publicKeys)
		//Build a new dealer
		dealer.vssDealer, _ = vss.NewDealer(suite, dealer.priKey, dealer.secret, dealer.verifiersPub, dealer.vssThreshold)

		// 1. dispatch encrypted deals to all verifiers
		encDeals, err := dealer.vssDealer.EncryptedDeals()
		if err != nil {
			log.Fatal(err)
		}
		deals := vss.EncryptedDeals{}
		deals.SetEncryptedDealsArray(encDeals)
		(*dealer.network).Broadcast(&deals)
		dealer.totalResponses = 0
		dealer.currentResp = 0
	}
}
func (dealer *Dealer) afterReceiveResponse(e *fsm.Event) {
	fmt.Printf("EVENT: afterReceiveResponse\n")
	resp := dealer.responses[dealer.currentResp]
	_, err := dealer.vssDealer.ProcessResponse(resp)
	if err != nil {
		fmt.Println("dealer ProcessResponse err ", err)
	}

	if dealer.vssDealer.DealCertified() {
		fmt.Println("dealer EnoughApprovals ", dealer.vssDealer.EnoughApprovals(), " DealCertified ", dealer.vssDealer.DealCertified())
		dealer.pubPoly = share.NewPubPoly(suite, suite.Point().Base(), dealer.vssDealer.Commits())
		responses := vss.Responses{}
		responses.SetResponseArray(suite, dealer.responses)
		(*dealer.network).Broadcast(&responses)
	}
}
func (dealer *Dealer) enterVerified(e *fsm.Event) {
	fmt.Printf("STATE : enterVerified\n")
}

func (dealer *Dealer) afterReceiveSignature(e *fsm.Event) {
	fmt.Printf("EVENT: afterReceiveSignature \n")
	(*dealer.signMap).Range(func(key, value interface{}) bool {
		sigShares := value.([][]byte)
		if len(sigShares) == dealer.nbVerifiers {
			result, _ := (*dealer.contentMap).Load(key)
			content, ok := result.([]byte)
			if !ok {
				fmt.Println("afterReceiveSignature value not found for key: ", key)
			}

			sig, _ := tbls.Recover(suite, dealer.pubPoly, content, sigShares, dealer.nbVerifiers/2+1, dealer.nbVerifiers)
			err := bls.Verify(suite, dealer.pubPoly.Commit(), content, sig)
			if err == nil {
				endTime = time.Now()
				log.Println("End:	took", endTime.Sub(startTime))
				fmt.Println("Dealer result = ", string(content), " verify success")
			} else {
				fmt.Println("Dealer result = ", content, " verify failed")
			}
			(*dealer.contentMap).Delete(key)
			(*dealer.signMap).Delete(key)
		}
		return true
	})
}

// main
func main() {
	dealerFlag := flag.String("dealerAddr", "", "dealer address")
	roleFlag := flag.String("role", "dealer", "")
	nbVerifiersFlag := flag.Int("nbVerifiers", 3, "Number of verifiers")
	portFlag := flag.Int("port", 0, "port number")

	flag.Parse()
	dealerAddr := *dealerFlag
	role := *roleFlag
	nbVerifiers := *nbVerifiersFlag
	port := *portFlag
	fmt.Println("role: ", role)

	//1)Build a p2p network
	fmt.Println(*roleFlag)
	tunnel := make(chan p2p.P2PMessage)
	p, _ := p2p.InitClient(tunnel, port)
	defer close(tunnel)
	//2)Start to listen incoming connection
	if err := p.Listen(); err != nil {
		log.Fatal(err)
	}

	//3)
	var verifier *Verifier
	var dealer *Dealer
	if strings.TrimRight(role, "\n") == "dealer" {
		dealerSec, dealerPub := genPair()
		dealer = &Dealer{
			secret:         suite.Scalar().Pick(suite.RandomStream()),
			pubKey:         dealerPub,
			priKey:         dealerSec,
			verifiersPub:   make([]kyber.Point, 0),
			responses:      make([]*vss.Response, nbVerifiers),
			sigShares:      make([][]byte, 0),
			nbVerifiers:    nbVerifiers,
			vssThreshold:   (nbVerifiers + 1) / 2,
			totalResponses: 0,
			currentResp:    0,
			network:        &p,
			signMap:        new(sync.Map),
			contentMap:     new(sync.Map),
		}
		dealer.FSM = fsm.NewFSM(
			"initial",
			fsm.Events{
				{Name: "receivePubkey", Src: []string{"initial"}, Dst: "WaitingPubKeys"},
				{Name: "receivePubkey", Src: []string{"WaitingPubKeys"}, Dst: "WaitingPubKeys"},
				{Name: "receiveResponse", Src: []string{"WaitingPubKeys"}, Dst: "NotVerified"},
				{Name: "receiveResponse", Src: []string{"NotVerified"}, Dst: "NotVerified"},
				{Name: "enoughApproval", Src: []string{"NotVerified"}, Dst: "Verified"},
				{Name: "receiveSignature", Src: []string{"Verified"}, Dst: "Verified"},
			},
			fsm.Callbacks{
				"after_receivePubkey":    func(e *fsm.Event) { dealer.afterReceivePubkey(e) },
				"after_receiveResponse":  func(e *fsm.Event) { dealer.afterReceiveResponse(e) },
				"after_receiveSignature": func(e *fsm.Event) { dealer.afterReceiveSignature(e) },
			},
		)
	} else {
		verifierSec, verifierPub := genPair()
		verifier = &Verifier{
			pubKey:  verifierPub,
			priKey:  verifierSec,
			network: &p,
		}
		p.CreatePeer(dealerAddr)
		p.Broadcast(&verifierPub)
	}
	//
	//4)Handle message from peer
	go func() {
		for {
			select {
			//event from peer
			case msg := <-tunnel:
				switch content := msg.Msg.Message.(type) {
				case *vss.PublicKey:
					if dealer != nil {
						p, err := content.GetPoint(suite)
						if err != nil {
							log.Fatal(err)
						}
						dealer.verifiersPub = append(dealer.verifiersPub, p)
						err = dealer.FSM.Event("receivePubkey")
						if err != nil {
							fmt.Println(err)
						}
					} else {
						if verifier != nil {
							verifier.dealerPubKey = *content
						}
					}
				case *vss.PublicKeys:
					if verifier != nil {
						verifiersPub, err := content.GetPointArray(suite)
						if err != nil {
							log.Fatal(err)
						}
						p, err := verifier.dealerPubKey.GetPoint(suite)
						if err != nil {
							log.Fatal(err)
						}
						verifier.vssVerifier, err = vss.NewVerifier(suite, verifier.priKey, p, verifiersPub)
						if err != nil {
							log.Fatal(err)
						}
					}
				case *vss.EncryptedDeals:
					deals := content.Deals
					for i, deal := range deals {
						res, err := verifier.vssVerifier.ProcessEncryptedDeal(deal)
						if err != nil {
							fmt.Println("ProcessEncryptedDeal err ", err)
						} else {
							fmt.Println("ProcessEncryptedDeal SUCCESS ", i)
							p.Broadcast(res)
						}
					}
				case *vss.Response:
					if dealer != nil {
						if uint32(cap(dealer.responses)) > content.Index {
							dealer.currentResp = content.Index
							dealer.responses[dealer.currentResp] = content
							err := dealer.FSM.Event("receiveResponse")
							if err != nil {
								fmt.Println(err)
							}
							if dealer.vssDealer.DealCertified() {
								err := dealer.FSM.Event("enoughApproval")
								if err != nil {
									fmt.Println(err)
								}
							}
						} else {
							fmt.Printf("resp.Index %d out of range of dealer.responses %d\n", content.Index, cap(dealer.responses))
						}

					}
				case *vss.Responses:
					//fmt.Println("vssVerifier receive vss.Responses", content)

					resps := content.Responses
					for _, r := range resps {
						verifier.vssVerifier.ProcessResponse(r)
					}
					fmt.Println("vssVerifier ", verifier.vssVerifier.Index(), "dealCetified ", verifier.vssVerifier.DealCertified())

				case *vss.Signature:
					if dealer != nil {
						var sigShares [][]byte
						result, ok := (*dealer.signMap).Load(content.QueryId)
						if ok {
							fmt.Println("FOund content.QueryId ", content.QueryId)
							sigShares, ok = result.([][]byte)
							if !ok {
								//sigShares = make([][]byte, 0)
								fmt.Println("cast fail ", content.QueryId)
							}
						} else {
							fmt.Println("value not found for key: ", content.QueryId)
							sigShares = make([][]byte, 0)
						}
						sigShares = append(sigShares, content.Signature)
						(*dealer.signMap).Store(content.QueryId, sigShares)
						(*dealer.contentMap).Store(content.QueryId, content.Content)

						err := dealer.FSM.Event("receiveSignature")
						if err != nil {
							fmt.Println(err)
						}
					}
				case *internalMsg.Cmd:
					fmt.Println("receive cmd ", verifier.vssVerifier.DealCertified())
					if verifier.vssVerifier.DealCertified() {
						url := content.Args
						result, err := dataFetch(url)
						if err != nil {
							log.Fatal(err)
						}
						sig, err := tbls.Sign(suite, verifier.vssVerifier.Deal().SecShare, result)
						if err != nil {
							log.Fatal(err)
						}
						sign := &vss.Signature{
							Index:     uint32(verifier.vssVerifier.Index()),
							QueryId:   url,
							Content:   result,
							Signature: sig,
						}
						p.Broadcast(sign)
					}
				default:
					fmt.Println("unknown")
				}
			}
		}
	}()

	//5)Broadcast message to peers
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')

		// skip blank lines
		if len(strings.TrimSpace(input)) == 0 {
			continue
		}
		if strings.TrimRight(input, "\n") == "end" {
			fmt.Println("Stop()")
			break
		}

		if strings.Contains(input, "checkURL") {
			startTime = time.Now()
			dealer.sigShares = make([][]byte, 0)
			words := strings.Fields(input)
			fmt.Println(words, len(words))
			fmt.Println(words[1])
			if dealer.vssDealer.EnoughApprovals() {
				cmd := &internalMsg.Cmd{
					Ctype: internalMsg.Cmd_CHECKURL,
					Args:  words[1],
				}
				pb := proto.Message(cmd)
				(*dealer.network).Broadcast(pb)
			}
		}
	}
	fmt.Println("finish)")
}
