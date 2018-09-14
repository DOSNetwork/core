package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/DOSNetwork/core/examples/vss_p2p/internal"
	"github.com/DOSNetwork/core/p2p"

	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"

	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/golang/protobuf/proto"
)

var suite = suites.MustFind("bn256")

func genPair() (kyber.Scalar, kyber.Point) {

	secret := suite.Scalar().Pick(suite.RandomStream())
	public := suite.Point().Mul(secret, nil)
	return secret, public
}

type Dealer struct {
	secret          kyber.Scalar
	pubKey          kyber.Point
	priKey          kyber.Scalar
	responses       []*vss.Response
	sigShares       [][]byte
	verifiersPub    []kyber.Point
	nbVerifiers     int
	vssThreshold    int
	currentVerifier int
	vssDealer       *vss.Dealer
	pubPoly         *share.PubPoly
}
type Verifier struct {
	dealerPubKey kyber.Point
	pubKey       kyber.Point
	priKey       kyber.Scalar
	verifiersPub []kyber.Point
	vssVerifier  *vss.Verifier
}

// main
func main() {
	dealerFlag := flag.String("dealerAddr", "", "dealer address")
	roleFlag := flag.String("role", "dealer", "")
	nbVerifiersFlag := flag.Int("nbVerifiers", 3, "Number of verifiers")
	flag.Parse()
	dealerAddr := *dealerFlag
	role := *roleFlag
	nbVerifiers := *nbVerifiersFlag
	fmt.Println("role: ", role)

	//1)Build a p2p network
	fmt.Println(*roleFlag)
	tunnel := make(chan p2p.P2PMessage)
	p, _ := p2p.CreateP2PNetwork(tunnel)
	defer close(tunnel)
	//2)Start to listen incoming connection
	p.Listen()

	//3)
	var verifier *Verifier
	var dealer *Dealer
	if strings.TrimRight(role, "\n") == "dealer" {
		dealerSec, dealerPub := genPair()
		dealer = &Dealer{
			secret:          suite.Scalar().Pick(suite.RandomStream()),
			pubKey:          dealerPub,
			priKey:          dealerSec,
			verifiersPub:    make([]kyber.Point, nbVerifiers),
			responses:       make([]*vss.Response, nbVerifiers),
			sigShares:       make([][]byte, 0),
			nbVerifiers:     nbVerifiers,
			vssThreshold:    (nbVerifiers + 1) / 2,
			currentVerifier: 0,
		}
	} else {
		verifierSec, verifierPub := genPair()
		verifier = &Verifier{
			pubKey: verifierPub,
			priKey: verifierSec,
		}
		_ = p.CreatePeer(dealerAddr, nil)
		msg := packPublicKey(verifierPub)
		p.Broadcast(&msg)
	}
	//
	//4)Handle message from peer
	go func() {
		var packedMessage proto.Message
		for {
			select {
			//event from peer
			case msg := <-tunnel:
				switch content := msg.Msg.Message.(type) {
				case *internal.PublicKey:
					if dealer != nil {
						if dealer.currentVerifier < dealer.nbVerifiers {
							dealer.verifiersPub[dealer.currentVerifier] = unpackPublicKey(suite, content)
							dealer.currentVerifier++
							//Wait for all verifiers
							if dealer.currentVerifier == dealer.nbVerifiers {
								//Send dealer public key to all verifiers
								packedMessage := packPublicKey(dealer.pubKey)
								p.Broadcast(&packedMessage)
								//Send all of verifier's pubkey to all verifiers
								packedMessage = packPublicKeys(dealer.verifiersPub)
								p.Broadcast(&packedMessage)
								//Build a new dealer
								dealer.vssDealer, _ = vss.NewDealer(suite, dealer.priKey, dealer.secret, dealer.verifiersPub, dealer.vssThreshold)

								// 1. dispatch encrypted deals to all verifiers
								encDeals, _ := dealer.vssDealer.EncryptedDeals()
								packedMessage = packEncryptedDeals(encDeals)
								p.Broadcast(&packedMessage)
								dealer.currentVerifier = 0
							}
						}
					} else {
						if verifier != nil {
							verifier.dealerPubKey = unpackPublicKey(suite, content)
						}
					}
					//fmt.Println("receive PublicKey", unpackPublicKey(content).String())
				case *internal.PublicKeys:
					if verifier != nil {
						verifiersPub := unpackPublicKeys(suite, content)
						verifier.vssVerifier, _ = vss.NewVerifier(suite, verifier.priKey, verifier.dealerPubKey, verifiersPub)
					}
				case *internal.EncryptedDeals:
					deals := unpackEncryptedDeals(content)
					for _, deal := range deals {
						res, err := verifier.vssVerifier.ProcessEncryptedDeal(deal)
						if err != nil {
							fmt.Println("ProcessEncryptedDeal err ", err)
						} else {
							packedMessage = packResonse(res)
							p.Broadcast(&packedMessage)
						}
					}
				case *internal.Response:
					resp := unpackResonse(content)
					dealer.responses[resp.Index] = resp
					_, err := dealer.vssDealer.ProcessResponse(resp)
					if err != nil {
						fmt.Println("dealer ProcessResponse err ", err)
					}
					if dealer.vssDealer.EnoughApprovals() || dealer.vssDealer.DealCertified() {
						fmt.Println("dealer EnoughApprovals ", dealer.vssDealer.EnoughApprovals(), " DealCertified ", dealer.vssDealer.DealCertified())
						dealer.pubPoly = share.NewPubPoly(suite, suite.Point().Base(), dealer.vssDealer.Commits())
					}
					dealer.currentVerifier++
					if dealer.currentVerifier == dealer.nbVerifiers {
						packedMessage = packResonses(dealer.responses)
						p.Broadcast(&packedMessage)
					}
				case *internal.Responses:
					resps := unpackResonses(content)
					for _, r := range resps {
						verifier.vssVerifier.ProcessResponse(r)
					}
					fmt.Println("vssVerifier ", verifier.vssVerifier.Index(), "dealCetified ", verifier.vssVerifier.DealCertified())

					s := verifier.vssVerifier.Deal().SecShare
					sig, _ := tbls.Sign(suite, s, []byte("test"))

					packedMessage = packSignature(sig, verifier.vssVerifier.Index())
					p.Broadcast(&packedMessage)

				case *internal.Signature:
					dealer.sigShares = append(dealer.sigShares, content.Signature)
					fmt.Println("len sigShares ", len(dealer.sigShares))
					if len(dealer.sigShares) == dealer.nbVerifiers {
						fmt.Println("dealer EnoughApprovals ", dealer.vssDealer.EnoughApprovals(), " DealCertified ", dealer.vssDealer.DealCertified())
						pubPoly := share.NewPubPoly(suite, suite.Point().Base(), dealer.vssDealer.Commits())
						sig, _ := tbls.Recover(suite, pubPoly, []byte("test"), dealer.sigShares, nbVerifiers/2+1, nbVerifiers)
						err := bls.Verify(suite, pubPoly.Commit(), []byte("test"), sig)
						if err == nil {
							fmt.Println("Dealer bls.Verify success ")
						}
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
	}
	fmt.Println("finish)")
}
