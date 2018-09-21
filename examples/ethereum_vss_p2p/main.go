package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/examples/ethereum_vss_p2p/contract"
	"github.com/DOSNetwork/core/group/bn256"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/DOSNetwork/core/examples/ethereum_vss_p2p/internalMsg"
	"github.com/DOSNetwork/core/examples/ethereum_vss_p2p/msgParser"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/golang/protobuf/proto"
)

const key = ``
const passphrase = ``

var ethReomteNode = "wss://rinkeby.infura.io/ws"
var contractAddressHex = "0xbD5784b224D40213df1F9eeb572961E2a859Cb80"
var groupId = big.NewInt(123)
var suite = suites.MustFind("bn256")
var mutex = &sync.Mutex{}

func genPair() (kyber.Scalar, kyber.Point) {

	secret := suite.Scalar().Pick(suite.RandomStream())
	public := suite.Point().Mul(secret, nil)
	return secret, public
}

type dataFeed struct {
	data []byte
	sig  [][]byte
}

type Dealer struct {
	secret          kyber.Scalar
	pubKey          kyber.Point
	priKey          kyber.Scalar
	responses       []*vss.Response
	sigShares       map[string]*dataFeed
	verifiersPub    []kyber.Point
	nbVerifiers     int
	vssThreshold    int
	currentVerifier int
	vssDealer       *vss.Dealer
	pubPoly         *share.PubPoly
	mux             sync.Mutex
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
	p, _ := p2p.CreateP2PNetwork(tunnel, 0)
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
			sigShares:       make(map[string]*dataFeed),
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
		msg := msgParser.PackPublicKey(verifierPub)
		p.Broadcast(msg)
	}

	//3.5) Connect to ethereum
	client, err := ethclient.Dial(ethReomteNode)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(contractAddressHex)
	proxy, err := dosproxy.NewDOSProxy(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	//4)Handle message from peer
	go func() {
		var packedMessage proto.Message
		for {
			select {
			//event from peer
			case msg := <-tunnel:
				switch content := msg.Msg.Message.(type) {
				case *internalMsg.PublicKey:
					if dealer != nil {
						if dealer.currentVerifier < dealer.nbVerifiers {
							dealer.verifiersPub[dealer.currentVerifier] = *msgParser.UnpackPublicKey(suite, content)
							dealer.currentVerifier++
							//Wait for all verifiers
							if dealer.currentVerifier == dealer.nbVerifiers {
								//Send dealer public key to all verifiers
								packedMessage := msgParser.PackPublicKey(dealer.pubKey)
								p.Broadcast(packedMessage)
								//Send all of verifier's pubkey to all verifiers
								packedMessage = msgParser.PackPublicKeys(dealer.verifiersPub)
								p.Broadcast(packedMessage)
								//Build a new dealer
								dealer.vssDealer, _ = vss.NewDealer(suite, dealer.priKey, dealer.secret, dealer.verifiersPub, dealer.vssThreshold)

								// 1. dispatch encrypted deals to all verifiers
								encDeals, _ := dealer.vssDealer.EncryptedDeals()
								packedMessage = msgParser.PackEncryptedDeals(encDeals)
								p.Broadcast(packedMessage)
								dealer.currentVerifier = 0
							}
						}
					} else {
						if verifier != nil {
							verifier.dealerPubKey = *msgParser.UnpackPublicKey(suite, content)
						}
					}
					//fmt.Println("receive PublicKey", unpackPublicKey(content).String())
				case *internalMsg.PublicKeys:
					if verifier != nil {
						verifiersPub := msgParser.UnpackPublicKeys(suite, content)
						verifier.vssVerifier, _ = vss.NewVerifier(suite, verifier.priKey, verifier.dealerPubKey, *verifiersPub)
					}
				case *internalMsg.EncryptedDeals:
					deals := msgParser.UnpackEncryptedDeals(content)
					for _, deal := range deals {
						res, err := verifier.vssVerifier.ProcessEncryptedDeal(deal)
						if err != nil {
							fmt.Println("ProcessEncryptedDeal err ", err)
						} else {
							packedMessage = msgParser.PackResponse(res)
							p.Broadcast(packedMessage)
						}
					}
				case *internalMsg.Response:
					resp := msgParser.UnpackResponse(content)
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
					fmt.Println(dealer.currentVerifier, "  ", dealer.nbVerifiers, "........")
					if dealer.currentVerifier == dealer.nbVerifiers {
						fmt.Println("Start")

						chPubKey := make(chan *dosproxy.DOSProxyLogSuccPubKeySub)
						done := make(chan bool)
						go subscribePubKey(proxy, chPubKey, done)

						<-done
						x0, x1, y0, y1, err := getPubKey(dealer.pubPoly)
						if err != nil {
							log.Fatal(err)
						}

						err = uploadPubKey(client, proxy, groupId, x0, x1, y0, y1)
						if err != nil {
							log.Fatal(err)
						}

						<-done
						fmt.Println("Group set-up finished, start listening to query...")

						packedMessage = msgParser.PackResponses(dealer.responses)
						p.Broadcast(packedMessage)
					}

				case *internalMsg.Responses:
					resps := msgParser.UnpackResponses(content)
					for _, r := range resps {
						verifier.vssVerifier.ProcessResponse(r)
					}
					fmt.Println("vssVerifier ", verifier.vssVerifier.Index(), "dealCetified ", verifier.vssVerifier.DealCertified())

					chUrl := make(chan *dosproxy.DOSProxyLogUrl)
					go subscribeEvent(p, verifier, proxy, chUrl)

				case *internalMsg.Signature:
					dealer.mux.Lock()
					if _, ok := dealer.sigShares[content.GetQueryId()]; !ok {
						newDataFeed := &dataFeed{
							data: content.GetContent(),
							sig:  make([][]byte, 0),
						}
						newDataFeed.sig = append(newDataFeed.sig, content.GetSignature())
						dealer.sigShares[content.GetQueryId()] = newDataFeed
					} else {
						dealer.sigShares[content.GetQueryId()].sig = append(dealer.sigShares[content.GetQueryId()].sig, content.GetSignature())
					}
					dealer.mux.Unlock()

					fmt.Println("len sigShares ", len(dealer.sigShares[content.GetQueryId()].sig))
					if len(dealer.sigShares[content.GetQueryId()].sig) == dealer.nbVerifiers {
						fmt.Println("dealer EnoughApprovals ", dealer.vssDealer.EnoughApprovals(), " DealCertified ", dealer.vssDealer.DealCertified())
						sig, _ := tbls.Recover(suite, dealer.pubPoly, dealer.sigShares[content.GetQueryId()].data, dealer.sigShares[content.GetQueryId()].sig, nbVerifiers/2+1, nbVerifiers)
						err := bls.Verify(suite, dealer.pubPoly.Commit(), dealer.sigShares[content.GetQueryId()].data, sig)
						if err == nil {
							fmt.Println("Dealer bls.Verify success ")
						}

						x, y := negate(sig)

						queryId := new(big.Int)
						queryId, ok := queryId.SetString(content.GetQueryId(), 10)
						if !ok {
							log.Fatal("SetString: error")
						}

						err = dataReturn(client, proxy, queryId, dealer.sigShares[content.GetQueryId()].data, x, y)
						if err != nil {
							log.Fatal(err)
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

func getPubKey(pubPoly *share.PubPoly) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	pubKey := pubPoly.Commit()
	pubKeyMar, err := pubKey.MarshalBinary()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	x0 := big.NewInt(0)
	x1 := big.NewInt(0)
	y0 := big.NewInt(0)
	y1 := big.NewInt(0)
	x0.SetBytes(pubKeyMar[1:33])
	x1.SetBytes(pubKeyMar[33:65])
	y0.SetBytes(pubKeyMar[65:97])
	y1.SetBytes(pubKeyMar[97:])
	return x0, x1, y0, y1, nil
}

func subscribePubKey(proxy *dosproxy.DOSProxy, ch chan *dosproxy.DOSProxyLogSuccPubKeySub, done chan bool) {
	fmt.Println("Establishing listen channel to group public key...")

	connected := false

	retried := false

	go func() {
		time.Sleep(3 * time.Second)
		if !connected {
			retried = true
			fmt.Println("retry")

			client, err := ethclient.Dial(ethReomteNode)
			if err != nil {
				log.Fatal(err)
			}

			contractAddress := common.HexToAddress(contractAddressHex)
			proxy, err = dosproxy.NewDOSProxy(contractAddress, client)
			if err != nil {
				log.Fatal(err)
			}

			go subscribePubKey(proxy, ch, done)
		}
	}()

	opt := &bind.WatchOpts{}
	sub, err := proxy.DOSProxyFilterer.WatchLogSuccPubKeySub(opt, ch)
	if err != nil {
		log.Fatal(err)
	}

	if retried {
		return
	}

	connected = true

	done <- true

	fmt.Println("Channel established.")

	for range ch {
		fmt.Println("Group public key submitted event Caught.")
		done <- true
		break
	}
	sub.Unsubscribe()
	close(ch)
}

func uploadPubKey(client *ethclient.Client, proxy *dosproxy.DOSProxy, groupId, x0, x1, y0, y1 *big.Int) error {
	fmt.Println("Starting submitting group public key...")

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		return err
	}

	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := proxy.SetPublicKey(auth, groupId, x0, x1, y0, y1)

	if err != nil {
		return err
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("groupId: ", groupId)
	fmt.Println("x0: ", x0)
	fmt.Println("x1: ", x1)
	fmt.Println("y0: ", y0)
	fmt.Println("y1: ", y1)
	fmt.Println("Group public key submitted, waiting for confirmation...")

	return nil
}

func subscribeEvent(p p2p.P2PInterface, verifier *Verifier, proxy *dosproxy.DOSProxy, ch chan *dosproxy.DOSProxyLogUrl) {

	connected := false

	retried := false

	go func() {
		time.Sleep(3 * time.Second)
		if !connected {
			retried = true
			fmt.Println("retry")

			client, err := ethclient.Dial(ethReomteNode)
			if err != nil {
				log.Fatal(err)
			}

			contractAddress := common.HexToAddress(contractAddressHex)
			proxy, err = dosproxy.NewDOSProxy(contractAddress, client)
			if err != nil {
				log.Fatal(err)
			}

			go subscribeEvent(p, verifier, proxy, ch)
		}
	}()

	opt := &bind.WatchOpts{}
	sub, err := proxy.DOSProxyFilterer.WatchLogUrl(opt, ch)
	if err != nil {
		log.Fatal(err)
	}

	if retried {
		return
	}

	connected = true

	fmt.Println("Group set-up finished, start listening to query...")
	for i := range ch {
		fmt.Printf("Query-ID: %v \n", i.QueryId)
		fmt.Println("Query Url: ", i.Url)
		go queryFulfill(p, verifier, i.QueryId, i.Url)
	}
	sub.Unsubscribe()
	close(ch)
}

func queryFulfill(p p2p.P2PInterface, verifier *Verifier, queryId *big.Int, url string) {
	data, err := dataFetch(url)
	if err != nil {
		log.Fatal(err)
	}

	sig, err := dataSign(verifier, data)
	if err != nil {
		log.Fatal(err)
	}

	packedMessage := msgParser.PackSignature(uint32(verifier.vssVerifier.Index()), queryId.String(), data, sig)
	p.Broadcast(packedMessage)
}

func dataFetch(url string) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body.Close()

	fmt.Println("Detched data: ", string(body))

	return body, nil

}

func dataSign(verifier *Verifier, data []byte) ([]byte, error) {
	s := verifier.vssVerifier.Deal().SecShare
	sig, err := tbls.Sign(suite, s, data)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func negate(sig []byte) (*big.Int, *big.Int) {
	x := big.NewInt(0)
	y := big.NewInt(0)
	x.SetBytes(sig[0:32])
	y.SetBytes(sig[32:])

	if x.Cmp(big.NewInt(0)) == 0 && y.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(0)
	}

	return x, big.NewInt(0).Sub(bn256.P, big.NewInt(0).Mod(y, bn256.P))
}

func dataReturn(client *ethclient.Client, proxy *dosproxy.DOSProxy, queryId *big.Int, data []byte, x, y *big.Int) error {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		return err
	}

	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice

	tx, err := proxy.TriggerCallback(auth, queryId, data, x, y)
	if err != nil {
		return err
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Printf("Query_ID %v request fulfilled \n", queryId)

	return nil
}
