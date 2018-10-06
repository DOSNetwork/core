package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/DOSNetwork/core/blockchain"
	"github.com/DOSNetwork/core/blockchain/eth"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
)

var mux sync.Mutex

type dosNode struct {
	suite       suites.Suite
	chSignature chan vss.Signature
	chURL       chan string
	nodeEvent   chan string
	//signMap        map[string][][]byte
	//contentMap     map[string][]byte
	signMap        *sync.Map
	contentMap     *sync.Map
	nbParticipants int
	groupPubPoly   share.PubPoly
	shareSec       share.PriShare
	chainConn      blockchain.ChainInterface
	p2pDkg         dkg.P2PDkgInterface
	network        *p2p.P2PInterface
	groupIds       [][]byte
	randomNumber   *big.Int
}

func (d *dosNode) dataFetch(url string) ([]byte, error) {
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
	fmt.Println("dataFetch End:	took", time.Now().Sub(sTime))

	return body, nil

}

func (d *dosNode) CheckURL(QueryId *big.Int, url string) {

	result, err := d.dataFetch(url)
	if err != nil {
		fmt.Println(err)
	}
	d.signAndBroadcast(QueryId, result)
}
func (d *dosNode) GenerateRandomNumber() {

	//var err error
	//i, r  j = sign(last blockhash || ri-1, Gi, sk  j)
	QueryId := new(big.Int)
	QueryId.SetInt64(33)

	d.randomNumber.SetInt64(int64(rand.Intn(100)))
	//get last random number
	//d.randomNumber, err = d.chainConn.GetRandomNum()
	//if err != nil {
	//	fmt.Println("GetRandomNum() error!!!!!!!!!!!!!", err)
	//	d.randomNumber.SetInt64(1)
	//}
	//fmt.Println("GenerateRandomNumber!!", d.randomNumber.String())
	//Todo:generate a seed by combine last blockhash and last random number
	//get last blockhash
	//lastHash := d.chainConn.GetCurrBlockHash().Bytes()
	d.signAndBroadcast(QueryId, d.randomNumber.Bytes())
}

func (d *dosNode) signAndBroadcast(QueryId *big.Int, content []byte) {
	sig, _ := tbls.Sign(d.suite, d.p2pDkg.GetShareSecuirty(), content)
	sign := &vss.Signature{
		Index:     uint32(0),
		QueryId:   QueryId.String(),
		Content:   content,
		Signature: sig,
	}
	/*
		var sigShares [][]byte
		//sigShares = d.signMap[sign.QueryId]
		value, ok := (*d.signMap).Load(sign.QueryId)
		if ok {
			fmt.Println("Found content.QueryId ", sign.QueryId)
			sigShares, ok = value.([][]byte)
			if !ok {
				fmt.Println("cast failed ", sign.QueryId)
			}
		} else {
			fmt.Println("Not found for key: ", sign.QueryId)
			sigShares = make([][]byte, 0)
		}
		sigShares = append(sigShares, sig)
		//d.signMap[sign.QueryId] = sigShares
		//d.contentMap[sign.QueryId] = sign.Content
		(*d.signMap).Store(sign.QueryId, sigShares)
		(*d.contentMap).Store(sign.QueryId, sign.Content)
	*/
	d.chSignature <- *sign
	go d.receiveSignature()
	for _, member := range d.groupIds {
		if string(member) != string((*d.network).GetId().Id) {
			go (*d.network).SendMessageById(member, sign)
		}
	}
}

func (d *dosNode) receiveSignature() {
	sign := <-d.chSignature
	var sigShares [][]byte

	result, ok := (*d.signMap).Load(sign.QueryId)
	if ok {
		sigShares, ok = result.([][]byte)
		if !ok {
			fmt.Println("cast failed ", sign.QueryId)
		}
	} else {
		sigShares = make([][]byte, 0)
	}
	sigShares = append(sigShares, sign.Signature)
	//d.signMap[sign.QueryId] = sigShares
	//d.contentMap[sign.QueryId] = sign.Content
	(*d.signMap).Store(sign.QueryId, sigShares)
	(*d.contentMap).Store(sign.QueryId, sign.Content)
	fmt.Println("receiveSignature id ", sign.QueryId, " len ", len(sigShares))
	if len(sigShares) == d.nbParticipants {
		//content := d.contentMap[sign.QueryId]
		result, _ := (*d.contentMap).Load(sign.QueryId)
		content, ok := result.([]byte)
		if !ok {
			fmt.Println("afterReceiveSignature value not found for sign.QueryId: ", sign.QueryId)
		}
		sig, _ := tbls.Recover(d.suite, d.p2pDkg.GetGroupPublicPoly(), content, sigShares, d.nbParticipants/2+1, d.nbParticipants)
		//QueryId is 0 means that is for random number
		if sign.QueryId == "33" {
			//
			hashSig := sha256.Sum256(sig)
			randomNum := hashSig[:]
			d.randomNumber.SetBytes(randomNum)
			//sign this new random number  again
			QueryId := new(big.Int)
			QueryId.SetInt64(0)
			d.signAndBroadcast(QueryId, randomNum)
			(*d.contentMap).Delete(sign.QueryId)
			(*d.signMap).Delete(sign.QueryId)
		} else {

			err := bls.Verify(d.suite, d.p2pDkg.GetGroupPublicPoly().Commit(), content, sig)
			if err == nil {
				y := big.NewInt(int64(d.nbParticipants))
				z := big.NewInt(0)
				z = z.Mod(d.randomNumber, y)
				//Todo: Chose who should report to eth
				repoter := 0
				repoter = int(z.Int64())
				fmt.Println("repoter = ", repoter)
				repoter = 0
				if sign.QueryId != "0" {
					d.GenerateRandomNumber()
				}
				if d.p2pDkg.GetDKGIndex() == repoter {
					if sign.QueryId == "0" {
						randomN := new(big.Int)
						randomN.SetBytes(content)
						if d.p2pDkg.GetDKGIndex() == repoter {

							groupId := new(big.Int)
							groupId.SetBytes(d.p2pDkg.GetGroupId())
							mux.Lock()
							fmt.Println("Random Number result = ", randomN, " verify success")
							err = d.chainConn.SetRandomNum(groupId, d.randomNumber, sig)
							if err != nil {
								fmt.Println("SetRandomNum err ", err)
							}
							mux.Unlock()

						}
					} else {
						mux.Lock()
						fmt.Println("checkURL result = ", string(content), " verify success")
						qID := big.NewInt(0)
						qID.SetString(sign.QueryId, 10)
						//Todo:Need to have a way to detemie who should send back the result
						d.chainConn.DataReturn(qID, content, sig)
						mux.Unlock()
					}
				}
			} else {
				fmt.Println("checkURL result = ", string(content), " verify failed ", err)
			}

			(*d.contentMap).Delete(sign.QueryId)
			(*d.signMap).Delete(sign.QueryId)
		}
	}
}

// main
func main() {
	seedFlag := flag.String("seedAddr", "", "seed address")
	nbParticipantsFlag := flag.Int("nbVerifiers", 3, "Number of Participants")
	portFlag := flag.Int("port", 0, "port number")

	flag.Parse()
	seedAddr := *seedFlag
	_ = seedAddr
	nbParticipants := *nbParticipantsFlag
	port := *portFlag

	//1)Connect to Eth and Set node ID
	chainConn, err := blockchain.AdaptTo("ETH", true)
	if err != nil {
		log.Fatal(err)
	}

	err = chainConn.UploadID()
	if err != nil {
		log.Fatal(err)
	}

	bootstrapIp, err := chainConn.GetBootstrapIp()
	if err != nil {
		log.Fatal(err)
	}
	chUrl := make(chan interface{})
	go func() {
		chUrl <- &eth.DOSProxyLogUrl{}
	}()
	defer close(chUrl)

	chGroup := make(chan interface{})
	go func() {
		chGroup <- &eth.DOSProxyLogGrouping{}
	}()
	defer close(chGroup)

	chainConn.SubscribeEvent(chUrl)
	chainConn.SubscribeEvent(chGroup)

	//1)Build a p2p network
	peerEvent := make(chan p2p.P2PMessage, 100)
	defer close(peerEvent)
	p, _ := p2p.CreateP2PNetwork(peerEvent, port)
	p.SetId(chainConn.GetId())
	p.Listen()

	//3)Dial to peers to build peerClient
	if bootstrapIp != "" {
		fmt.Println(bootstrapIp)
		p.CreatePeer(bootstrapIp, nil)
		results := p.FindNode(p.GetId(), dht.BucketSize, 20)
		for _, result := range results {
			p.GetRoutingTable().Update(result)
			fmt.Println(p.GetId().Address, "Update peer: ", result.Address)
		}
	} else {
		err = chainConn.SetBootstrapIp(p.GetId().Address)
		if err != nil {
			fmt.Println(err)
		}
	}

	suite := suites.MustFind("bn256")
	peerEventForDKG := make(chan p2p.P2PMessage, 100)
	defer close(peerEventForDKG)
	p2pDkg, _ := dkg.CreateP2PDkg(p, suite, peerEventForDKG, nbParticipants)
	go p2pDkg.EventLoop()
	dkgEvent := make(chan string, 100)
	p2pDkg.SubscribeEvent(dkgEvent)
	defer close(dkgEvent)

	nodeEvent := make(chan string, 100)
	defer close(nodeEvent)

	d := &dosNode{
		suite:          suite,
		signMap:        new(sync.Map),
		contentMap:     new(sync.Map),
		chSignature:    make(chan vss.Signature, 100),
		chURL:          make(chan string, 100),
		nbParticipants: nbParticipants,
		network:        &p,
		chainConn:      chainConn,
		p2pDkg:         p2pDkg,
		randomNumber:   new(big.Int),
		nodeEvent:      nodeEvent,
	}

	for {
		select {
		//event from peer
		case msg := <-peerEvent:
			switch content := msg.Msg.Message.(type) {
			case *vss.PublicKey:
				peerEventForDKG <- msg
			case *dkg.Deal:
				peerEventForDKG <- msg
			case *dkg.Response:
				peerEventForDKG <- msg
			case *vss.Signature:
				d.chSignature <- *content
				go d.receiveSignature()
			default:
				fmt.Println("unknown", content)
			}
		case msg := <-chUrl:
			switch content := msg.(type) {
			case *eth.DOSProxyLogUrl:
				fmt.Printf("Query-ID: %v \n", content.QueryId)
				fmt.Println("Query Url: ", content.Url)
				d.CheckURL(content.QueryId, content.Url)
			default:
				fmt.Println("type mismatch")
			}
		case msg := <-d.nodeEvent:
			fmt.Println("nodeEvent", msg)
			//if msg == "random" {
			//d.GenerateRandomNumber()
			//}
		case msg := <-dkgEvent:
			if msg == "cetified" {
				if d.p2pDkg.GetDKGIndex() == 0 {
					gId := new(big.Int)
					gId.SetBytes(d.p2pDkg.GetGroupId())
					d.chainConn.UploadPubKey(gId, d.p2pDkg.GetGroupPublicPoly().Commit())
				}
			}

		case msg := <-chGroup:
			switch content := msg.(type) {
			case *eth.DOSProxyLogGrouping:
				fmt.Printf("DOSProxyLogGrouping \n")
				isMember := false
				groupIds := [][]byte{}
				for _, node := range content.NodeId {
					id := node.Bytes()
					if string(id) == string(p.GetId().Id) {
						isMember = true
					}
					groupIds = append(groupIds, id)
				}
				if isMember {
					d.groupIds = groupIds
					d.p2pDkg.SetGroupId(content.GroupId.Bytes())
					d.p2pDkg.SetGroupMembers(groupIds)
					d.p2pDkg.RunDKG()
				}

			default:
				fmt.Println("type mismatch")
			}
		}
	}
}
