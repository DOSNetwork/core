package main

import (
	"bytes"
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
	ticker         *time.Ticker
	quit           chan struct{}
	signMap        *sync.Map
	contentMap     *sync.Map
	signTypeMap    *sync.Map
	finishMap      *sync.Map
	nbParticipants int
	nbThreshold    int
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

// TODO: error handling and logging
// Note: Signed messages keep synced with on-chain contracts
func (d *dosNode) GenerateRandomNumber(preRandom *big.Int, preBlock *big.Int) {
	fmt.Printf("GenerateRandomNumber #%v \n", preRandom)
	//To avoid duplicate query
	_, ok := (*d.signMap).Load(preRandom.String())
	if !ok {
		hash, err := d.chainConn.GetBlockHashByNumber(preBlock)
		if err != nil {
			fmt.Printf("GetBlockHashByNumber #%v fails\n", preBlock)
			return
		}
		// message = concat(lastUpdatedBlockhash, lastRandomness)
		msg := append(hash[:], preRandom.Bytes()...)
		fmt.Printf("GenerateRandomNumber msg = %x\n", msg)
		d.signAndBroadcast(ForRandomNumber, preRandom, msg)
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
	for i, member := range d.groupIds {
		if string(member) != string((*d.network).GetId().Id) {
			//Todo:Need to check to see if it is thread safe
			fmt.Println(i, " : send to ", member)
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
			if len(sigShares) < d.nbParticipants {
				sigShares = append(sigShares, sign.Signature)
				(*d.signMap).Store(sign.QueryId, sigShares)
				(*d.contentMap).Store(sign.QueryId, sign.Content)
				(*d.signTypeMap).Store(sign.QueryId, sign.Index)
				fmt.Println("receiveSignature id ", sign.QueryId, " len ", len(sigShares), " nbParticipants ", d.nbParticipants)
				if len(sigShares) >= d.nbThreshold {
					d.toOnChainQueue <- sign.QueryId
				}
			} else {
				fmt.Println("Extra signature !!!!")
			}
		} else {
			//Todo:Need to ignore those finished signature
			_, ok := (*d.finishMap).Load(sign.QueryId)
			if !ok {
				//Other nodes has received eth event and send signature to other nodes
				//Node put back these signatures until it received eth event.
				d.chSignature <- sign
			}
		}
	}
}

func (d *dosNode) getReporter() int {
	/*
		randomNumber, err := d.chainConn.GetRandomNum()
		if err != nil {
			fmt.Println("getReporter err ", err)
		}

		x := int(randomNumber.Uint64())
		if x < 0 {
			x = 0 - x
		}
		y := int(d.nbParticipants)
		reporter := x % y
	*/
	return 0
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
			continue
		}
		result, ok = (*d.signTypeMap).Load(queryId)
		signType, ok = result.(uint32)
		if !ok {
			continue
		}

		sig, err := tbls.Recover(d.suite, d.p2pDkg.GetGroupPublicPoly(), content, sigShares, d.nbThreshold, d.nbParticipants)
		if err != nil {
			continue
		}
		err = bls.Verify(d.suite, d.p2pDkg.GetGroupPublicPoly().Commit(), content, sig)
		if err != nil {
			continue
		}
		repoter := d.getReporter()
		fmt.Println("reporter ", repoter)
		(*d.contentMap).Delete(queryId)
		(*d.signMap).Delete(queryId)
		(*d.signTypeMap).Delete(queryId)
		(*d.finishMap).Store(queryId, time.Now())
		//QueryId is 0 means that is for random number
		switch signType {
		case ForRandomNumber:
			randomN := new(big.Int)
			randomN.SetBytes(content)
			fmt.Println("Random Number result = ", randomN, " verify success")
			if d.p2pDkg.GetDKGIndex() == repoter {
				fmt.Println("Random Number result = ", randomN, " verify success")
				err = d.chainConn.SetRandomNum(sig)
				if err != nil {
					fmt.Println("SetRandomNum err ", err)
					err = d.chainConn.SetRandomNum(sig)
				}
			}
		default:
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
func (d *dosNode) isMember(dispatchedGroup [4]*big.Int) bool {
	if d.p2pDkg.IsCetified() {
		temp := append(dispatchedGroup[2].Bytes(), dispatchedGroup[3].Bytes()...)
		temp = append(dispatchedGroup[1].Bytes(), temp...)
		temp = append(dispatchedGroup[0].Bytes(), temp...)
		temp = append([]byte{0x01}, temp...)
		fmt.Println("isMember from eth : ", temp)
		groupPub, err := d.p2pDkg.GetGroupPublicPoly().Commit().MarshalBinary()
		fmt.Println("isMember : ", groupPub)
		if err != nil {
			fmt.Println(err)
			return false
		}
		r := bytes.Compare(groupPub, temp)
		if r == 0 {
			return true
		}
	}
	return false
}
func (d *dosNode) cleanFinishMap() {
	for {
		select {
		case <-d.ticker.C:
			d.finishMap.Range(func(key, value interface{}) bool {
				record := value.(time.Time)
				dur := time.Since(record)
				if dur.Minutes() >= 10 {
					fmt.Println("clean querID ", key, " time ", record)
					d.finishMap.Delete(key)
				}
				return true
			})
		case <-d.quit:
			d.ticker.Stop()
			return
		}
	}
}

// main
func main() {
	roleFlag := flag.String("role", "", "BootstrapNode or not")
	nbParticipantsFlag := flag.Int("nbVerifiers", 3, "Number of Participants")
	portFlag := flag.Int("port", 0, "port number")
	bootstrapIpFlag := flag.String("bootstrapIp", "67.207.98.117:42745", "bootstrapIp")

	flag.Parse()
	role := *roleFlag
	nbParticipants := *nbParticipantsFlag
	port := *portFlag
	bootstrapIp := *bootstrapIpFlag
	//1)Connect to Eth and Set node ID
	chainConn, err := blockchain.AdaptTo("ETH", true)
	if err != nil {
		log.Fatal(err)
	}
	if role == "bootstrape" {
		chainConn.ResetNodeIDs()
	}

	err = chainConn.UploadID()
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
	if role == "" {
		fmt.Println(bootstrapIp)
		p.CreatePeer(bootstrapIp, nil)
		results := p.FindNode(p.GetId(), dht.BucketSize, 20)
		for _, result := range results {
			p.GetRoutingTable().Update(result)
			fmt.Println(p.GetId().Address, "Update peer: ", result.Address)
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
		finishMap:      new(sync.Map),
		chSignature:    make(chan vss.Signature, 100),
		chURL:          make(chan string, 100),
		ticker:         time.NewTicker(10 * time.Minute),
		quit:           make(chan struct{}),
		nbParticipants: nbParticipants,
		nbThreshold:    nbParticipants/2 + 1,
		network:        &p,
		chainConn:      chainConn,
		p2pDkg:         p2pDkg,
		//randomNumber:   new(big.Int),
		toOnChainQueue: toOnChainQueue,
	}
	go d.sendToOnchain()
	go d.receiveSignature()
	go d.cleanFinishMap()
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
					d.chainConn.UploadPubKey(d.p2pDkg.GetGroupPublicPoly().Commit())
				}
			}
		case msg := <-chUrl:
			switch content := msg.(type) {
			case *eth.DOSProxyLogUrl:
				if d.isMember(content.DispatchedGroup) {
					d.CheckURL(content.QueryId, content.Url)
				}
			default:
				fmt.Println("type mismatch")
			}
		case msg := <-chRandom:
			switch content := msg.(type) {
			case *eth.DOSProxyLogUpdateRandom:
				fmt.Println("event DOSProxyLogUpdateRandom", d.isMember(content.DispatchedGroup))
				if d.isMember(content.DispatchedGroup) {
					//To avoid the err: SetRandomNum err  nonce too low
					timer := time.NewTimer(1 * time.Second)
					go func() {
						<-timer.C
						d.GenerateRandomNumber(content.LastRandomness, content.LastUpdatedBlock)
					}()
				}
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
				d.nbThreshold = d.nbParticipants/2 + 1
				d.p2pDkg.SetNbParticipants(d.nbParticipants)
				if isMember {
					d.groupIds = groupIds
					d.p2pDkg.SetGroupMembers(groupIds)
					d.p2pDkg.RunDKG()
				}

			default:
				fmt.Println("type mismatch")
			}
		}
	}
}
