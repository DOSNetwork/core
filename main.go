package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
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

const (
	ForRandomNumber = uint32(0)
	ForCheckURL     = uint32(1)
)

type dosNode struct {
	suite          suites.Suite
	chSignature    chan vss.Signature
	chURL          chan string
	toOnChainQueue chan string
	signMap        *sync.Map
	contentMap     *sync.Map
	signTypeMap    *sync.Map
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
	//To avoid duplicate query
	_, ok := (*d.signMap).Load(QueryId.String())
	if !ok {
		fmt.Println("CheckURL!!", QueryId)
		result, err := d.dataFetch(url)
		if err != nil {
			fmt.Println(err)
		}
		d.signAndBroadcast(ForCheckURL, QueryId, result)
	}
}
func (d *dosNode) GenerateRandomNumber(QueryId *big.Int, preRendomum *big.Int) {
	//To avoid duplicate query
	_, ok := (*d.signMap).Load(QueryId.String())
	if !ok {
		fmt.Println("GenerateRandomNumber!!", QueryId)
		d.signAndBroadcast(ForRandomNumber, QueryId, preRendomum.Bytes())
	}
}

func (d *dosNode) signAndBroadcast(index uint32, QueryId *big.Int, content []byte) {
	sig, _ := tbls.Sign(d.suite, d.p2pDkg.GetShareSecuirty(), content)
	sign := &vss.Signature{
		Index:     uint32(index),
		QueryId:   QueryId.String(),
		Content:   content,
		Signature: sig,
	}
	//d.chSignature <- *sign
	sigShares := make([][]byte, 0)
	sigShares = append(sigShares, sig)
	(*d.signMap).Store(sign.QueryId, sigShares)
	(*d.contentMap).Store(sign.QueryId, sign.Content)
	(*d.signTypeMap).Store(sign.QueryId, sign.Index)
	for _, member := range d.groupIds {
		if string(member) != string((*d.network).GetId().Id) {
			//Todo:Need to check to see if it is thread safe
			(*d.network).SendMessageById(member, sign)
		}
	}
}

//receiveSignature is thread safe.
func (d *dosNode) receiveSignature() {
	for sign := range d.chSignature {

		var sigShares [][]byte

		result, ok := (*d.signMap).Load(sign.QueryId)
		if ok {
			sigShares, ok = result.([][]byte)
			if !ok {
				fmt.Println("cast failed ", sign.QueryId)
			}
			sigShares = append(sigShares, sign.Signature)
			(*d.signMap).Store(sign.QueryId, sigShares)
			(*d.contentMap).Store(sign.QueryId, sign.Content)
			(*d.signTypeMap).Store(sign.QueryId, sign.Index)
			fmt.Println("receiveSignature id ", sign.QueryId, " len ", len(sigShares), " nbParticipants ", d.nbParticipants)

			if len(sigShares) == d.nbParticipants {
				d.toOnChainQueue <- sign.QueryId
			}
		} else {
			//Other nodes has received eth event and send signature to other nodes
			//Node put back these signatures until it received eth event.
			d.chSignature <- sign
		}
	}
}

func (d *dosNode) getReporter() int {
	randomNumber, err := d.chainConn.GetRandomNum()
	if err != nil {
		fmt.Println("getReporter err ", err)
	}
	x := int(randomNumber.Int64())
	y := d.nbParticipants
	reporter := x % y
	return reporter
}
func (d *dosNode) sendToOnchain() {
	for queryId := range d.toOnChainQueue {
		var content []byte
		var sigShares [][]byte
		var signType uint32
		result, _ := (*d.contentMap).Load(queryId)
		content, ok := result.([]byte)
		if !ok {
			fmt.Println("sendToOnchain content not found for sign.QueryId: ", queryId)
		}
		result, ok = (*d.signMap).Load(queryId)
		sigShares, ok = result.([][]byte)
		if !ok {
			fmt.Println("sendToOnchain value sigShares not found for sign.QueryId: ", queryId)
		}
		result, ok = (*d.signTypeMap).Load(queryId)
		signType, ok = result.(uint32)
		if !ok {
			fmt.Println("sendToOnchain value signType notfound for sign.QueryId: ", queryId)
		}
		repoter := d.getReporter()
		sig, err := tbls.Recover(d.suite, d.p2pDkg.GetGroupPublicPoly(), content, sigShares, d.nbParticipants/2+1, d.nbParticipants)
		if err != nil {
			fmt.Println("Recover failed ", err)
		}
		(*d.contentMap).Delete(queryId)
		(*d.signMap).Delete(queryId)
		(*d.signTypeMap).Delete(queryId)
		//QueryId is 0 means that is for random number
		switch signType {
		case ForRandomNumber:
			err := bls.Verify(d.suite, d.p2pDkg.GetGroupPublicPoly().Commit(), content, sig)
			if err != nil {
				fmt.Println("!!!!!  Verify failed", err, " queryID ", queryId)
			}
			randomN := new(big.Int)
			randomN.SetBytes(content)
			if d.p2pDkg.GetDKGIndex() == repoter {
				groupId := new(big.Int)
				groupId.SetBytes(d.p2pDkg.GetGroupId())
				fmt.Println("Random Number result = ", randomN, " verify success")
				qID := big.NewInt(0)
				qID.SetString(queryId, 10)
				err = d.chainConn.SetRandomNum(qID, groupId, sig)
				if err != nil {
					fmt.Println("SetRandomNum err ", err)
					err = d.chainConn.SetRandomNum(qID, groupId, sig)
				}
			}
		default:
			err := bls.Verify(d.suite, d.p2pDkg.GetGroupPublicPoly().Commit(), content, sig)
			if err != nil {
				fmt.Println("!!!!!  Verify failed", err, " queryID ", queryId)
			}
			fmt.Println("checkURL result = ", string(content), " verify success")
			qID := big.NewInt(0)
			qID.SetString(queryId, 10)
			if d.p2pDkg.GetDKGIndex() == repoter {
				//Todo:Need to have a way to detemie who should send back the result
				err = d.chainConn.DataReturn(qID, content, sig)
				if err != nil {
					fmt.Println("DataReturn err ", err)
					err = d.chainConn.DataReturn(qID, content, sig)
				}
			}
		}
	}
}

// main
func main() {
	seedFlag := flag.String("seedAddr", "", "seed address")
	nbParticipantsFlag := flag.Int("nbVerifiers", 21, "Number of Participants")
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
	chRandom := make(chan interface{})
	go func() {
		chRandom <- &eth.DOSProxyLogUpdateRandom{}
	}()
	defer close(chGroup)
	chainConn.SubscribeEvent(chUrl)
	chainConn.SubscribeEvent(chGroup)
	chainConn.SubscribeEvent(chRandom)

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

	toOnChainQueue := make(chan string, 100)
	defer close(toOnChainQueue)

	d := &dosNode{
		suite:          suite,
		signMap:        new(sync.Map),
		signTypeMap:    new(sync.Map),
		contentMap:     new(sync.Map),
		chSignature:    make(chan vss.Signature, 100),
		chURL:          make(chan string, 100),
		nbParticipants: nbParticipants,
		network:        &p,
		chainConn:      chainConn,
		p2pDkg:         p2pDkg,
		//randomNumber:   new(big.Int),
		toOnChainQueue: toOnChainQueue,
	}
	go d.sendToOnchain()
	go d.receiveSignature()
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
				fmt.Println("signature form peer")
			default:
				fmt.Println("unknown", content)
			}
		case msg := <-dkgEvent:
			if msg == "cetified" {
				if d.p2pDkg.GetDKGIndex() == 0 {
					gId := new(big.Int)
					gId.SetBytes(d.p2pDkg.GetGroupId())
					d.chainConn.UploadPubKey(gId, d.p2pDkg.GetGroupPublicPoly().Commit())
				}
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
		case msg := <-chRandom:
			switch content := msg.(type) {
			case *eth.DOSProxyLogUpdateRandom:
				//if d.p2pDkg.GetGroupId() == content.GroupId {
				d.GenerateRandomNumber(content.RandomId, content.PreRandomNumber)
				//}
			default:
				fmt.Println("type mismatch")
			}
		case msg := <-chGroup:
			switch content := msg.(type) {
			case *eth.DOSProxyLogGrouping:

				isMember := false
				groupIds := [][]byte{}
				for i, node := range content.NodeId {
					id := node.Bytes()
					if string(id) == string(p.GetId().Id) {
						isMember = true
					}
					fmt.Println("DOSProxyLogGrouping member i= ", i, " id ", id, " ", isMember)
					groupIds = append(groupIds, id)
				}
				d.nbParticipants = len(groupIds)
				d.p2pDkg.SetNbParticipants(d.nbParticipants)
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
