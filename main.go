package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/DOSNetwork/core/blockchain"
	"github.com/DOSNetwork/core/blockchain/eth"
	"github.com/DOSNetwork/core/group/bn256"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
)

type dosNode struct {
	suite          suites.Suite
	chSignature    chan vss.Signature
	chURL          chan string
	signMap        map[string][][]byte
	contentMap     map[string][]byte
	nbParticipants int
	groupPubPoly   share.PubPoly
	shareSec       share.PriShare
	chainConn      blockchain.ChainInterface
	p2pDkg         dkg.P2PDkgInterface
	network        *p2p.P2PInterface
	groupIds       [][]byte
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

	sig, _ := tbls.Sign(d.suite, d.p2pDkg.GetShareSecuirty(), result)
	sign := &vss.Signature{
		Index:     uint32(0),
		QueryId:   QueryId.String(),
		Content:   result,
		Signature: sig,
	}

	var sigShares [][]byte
	sigShares = d.signMap[sign.QueryId]
	sigShares = append(sigShares, sig)
	d.signMap[sign.QueryId] = sigShares
	d.contentMap[sign.QueryId] = sign.Content

	for _, member := range d.groupIds {
		if string(member) != string((*d.network).GetId().Id) {
			go (*d.network).SendMessageById(member, sign)
		}
	}
}

func (d *dosNode) negate(sig []byte) (*big.Int, *big.Int) {
	x := big.NewInt(0)
	y := big.NewInt(0)
	x.SetBytes(sig[0:32])
	y.SetBytes(sig[32:])

	if x.Cmp(big.NewInt(0)) == 0 && y.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(0)
	}

	return x, big.NewInt(0).Sub(bn256.P, big.NewInt(0).Mod(y, bn256.P))
}
func (d *dosNode) receiveSignature() {
	sign := <-d.chSignature
	var sigShares [][]byte
	sigShares = d.signMap[sign.QueryId]
	sigShares = append(sigShares, sign.Signature)
	d.signMap[sign.QueryId] = sigShares
	d.contentMap[sign.QueryId] = sign.Content

	if len(sigShares) == d.nbParticipants {
		content := d.contentMap[sign.QueryId]
		sig, _ := tbls.Recover(d.suite, d.p2pDkg.GetGroupPublicPoly(), content, sigShares, d.nbParticipants/2+1, d.nbParticipants)
		err := bls.Verify(d.suite, d.p2pDkg.GetGroupPublicPoly().Commit(), content, sig)
		if err == nil {
			fmt.Println("checkURL result = ", string(content), " verify success")
			x, y := d.negate(sig)
			qID := big.NewInt(0)
			qID.SetString(sign.QueryId, 10)
			d.chainConn.DataReturn(qID, content, x, y)
		} else {
			fmt.Println("checkURL result = ", string(content), " verify failed ", err)
		}
		d.contentMap[sign.QueryId] = nil
		d.signMap[sign.QueryId] = nil
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

	d := &dosNode{
		suite:          suite,
		signMap:        make(map[string][][]byte),
		contentMap:     make(map[string][]byte),
		chSignature:    make(chan vss.Signature, 100),
		chURL:          make(chan string, 100),
		nbParticipants: nbParticipants,
		network:        &p,
		chainConn:      chainConn,
		p2pDkg:         p2pDkg,
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
		case msg := <-dkgEvent:
			if msg == "cetified" {
				pubKeyMar, err := d.p2pDkg.GetGroupPublicPoly().Commit().MarshalBinary()
				if err != nil {
					fmt.Println(err)
				}
				x0 := new(big.Int)
				x1 := new(big.Int)
				y0 := new(big.Int)
				y1 := new(big.Int)
				x0.SetBytes(pubKeyMar[1:33])
				x1.SetBytes(pubKeyMar[33:65])
				y0.SetBytes(pubKeyMar[65:97])
				y1.SetBytes(pubKeyMar[97:])
				gId := new(big.Int)
				gId.SetBytes(d.p2pDkg.GetGroupId())
				d.chainConn.UploadPubKey(gId, x0, x1, y0, y1)
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
