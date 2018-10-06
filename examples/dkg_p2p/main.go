package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/DOSNetwork/core/examples/dkg_p2p/internalMsg"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/blake2b"
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
	p2pDkg         dkg.P2PDkgInterface
	network        *p2p.P2PInterface
	groupIds       [][]byte
	ticker         *time.Ticker
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
func (d *dosNode) CheckURL() {
	url := <-d.chURL
	result, err := d.dataFetch(url)
	if err != nil {
		fmt.Println(err)
	}
	QueryId := new(big.Int)
	QueryId.SetInt64(36)
	d.signAndBroadcast(QueryId, result)
}
func (d *dosNode) GenerateRandomNumber() {
	var seed []byte
	//i, r  j = sign(last blockhash || ri-1, Gi, sk  j)
	QueryId := new(big.Int)
	QueryId.SetInt64(-1)
	//get last blockhash
	//get last random number
	//generate a seed by combine last blockhash and last random number
	seed = []byte("test")
	d.signAndBroadcast(QueryId, seed)
}

func (d *dosNode) signAndBroadcast(QueryId *big.Int, content []byte) {
	sig, _ := tbls.Sign(d.suite, d.p2pDkg.GetShareSecuirty(), content)
	sign := &vss.Signature{
		Index:     uint32(0),
		QueryId:   QueryId.String(),
		Content:   content,
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

func (d *dosNode) receiveSignature() {
	sign := <-d.chSignature
	fmt.Println("receiveSignature !!!!!")
	var sigShares [][]byte
	sigShares = d.signMap[sign.QueryId]
	sigShares = append(sigShares, sign.Signature)
	d.signMap[sign.QueryId] = sigShares
	d.contentMap[sign.QueryId] = sign.Content

	fmt.Println("Event afterReceiveSignature ", len(sigShares))
	if len(sigShares) == d.nbParticipants {
		content := d.contentMap[sign.QueryId]
		sig, _ := tbls.Recover(d.suite, d.p2pDkg.GetGroupPublicPoly(), content, sigShares, d.nbParticipants/2+1, d.nbParticipants)
		//QueryId is 0 means that is for random number
		if sign.QueryId == "-1" {
			//
			hashSig := sha256.Sum256(sig)
			randomNum := hashSig[:]
			//sign this new random number  again
			QueryId := new(big.Int)
			QueryId.SetInt64(0)
			randomN := new(big.Int)
			randomN.SetBytes(randomNum)
			fmt.Println("For random number", randomN)
			d.signAndBroadcast(QueryId, randomNum)
		} else {
			err := bls.Verify(d.suite, d.p2pDkg.GetGroupPublicPoly().Commit(), content, sig)
			if err == nil {
				if sign.QueryId == "0" {
					randomN := new(big.Int)
					randomN.SetBytes(content)
					x := new(big.Int)
					x.SetString(sign.QueryId, 10)
					y := big.NewInt(int64(d.nbParticipants))
					z := big.NewInt(0)
					z = z.Mod(x, y)
					if d.p2pDkg.GetDKGIndex() == int(z.Int64()) {
						fmt.Println("Random Number result = ", randomN, " verify success")
					}
				} else {
					fmt.Println("checkURL result = ", string(content), " verify success")
				}

			} else {
				fmt.Println("checkURL result = ", string(content), " verify failed ", err)
			}

		}
		d.contentMap[sign.QueryId] = nil
		d.signMap[sign.QueryId] = nil
	}
}

// main
func main() {
	seedFlag := flag.String("seedAddr", "", "seed address")
	nbParticipantsFlag := flag.Int("nbVerifiers", 3, "Number of Participants")
	portFlag := flag.Int("port", 44460, "port number")
	idFlag := flag.String("id", "nodeHost", "id")

	flag.Parse()
	seedAddr := *seedFlag
	nbParticipants := *nbParticipantsFlag
	port := *portFlag
	id := *idFlag

	//1)Build a p2p network
	peerEvent := make(chan p2p.P2PMessage, 100)
	p, _ := p2p.CreateP2PNetwork(peerEvent, port)
	defer close(peerEvent)

	//2)Set node ID
	hashId5 := blake2b.Sum256([]byte(id))
	p.SetId(hashId5[:])

	if id == "node1" {
		go http.ListenAndServe(":8080", nil)
	}
	//2)Start to listen incoming connection
	p.Listen()

	//3)Dial to peers to build peerClient
	if seedAddr != "" {
		p.CreatePeer(seedAddr, nil)
		results := p.FindNode(p.GetId(), dht.BucketSize, 20)
		for _, result := range results {
			p.GetRoutingTable().Update(result)
			fmt.Println(p.GetId().Address, "Update peer: ", result.Address)
		}
	}

	suite := suites.MustFind("bn256")
	peerEventForDKG := make(chan p2p.P2PMessage, 100)
	defer close(peerEventForDKG)
	dosDkg, _ := dkg.CreateP2PDkg(p, suite, peerEventForDKG, nbParticipants)
	go dosDkg.EventLoop()

	node := &dosNode{
		suite:          suite,
		signMap:        make(map[string][][]byte),
		contentMap:     make(map[string][]byte),
		chSignature:    make(chan vss.Signature, 100),
		chURL:          make(chan string, 100),
		nbParticipants: nbParticipants,
		network:        &p,
		p2pDkg:         dosDkg,
	}
	//node.ticker = time.NewTicker(30 * time.Second)
	go func() {
		groupIds := [][]byte{}
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
					node.chSignature <- *content
					node.receiveSignature()
				case *internalMsg.Cmd:
					fmt.Println("!!!!!!!!!!!!receive cmd ", content)
					if content.Ctype == internalMsg.Cmd_GROUPING {
						groupIds = nil
						//dosDkg.groupingStart = time.Now()
						nodes := strings.Split(content.Args, ",")
						isMember := false
						for _, node := range nodes {
							hashId := blake2b.Sum256([]byte(node))
							if string(hashId[:]) == string(p.GetId().Id) {
								isMember = true
							}
							groupIds = append(groupIds, hashId[:])
						}
						if isMember {
							node.nbParticipants = len(groupIds)
							node.p2pDkg.SetNbParticipants(len(groupIds))
							node.groupIds = groupIds
							dosDkg.SetGroupMembers(groupIds)
							dosDkg.RunDKG()
						}
					}

					if content.Ctype == internalMsg.Cmd_CHECKURL {
						//dosDkg.checkURLStart = time.Now()
						url := content.Args
						node.chURL <- url
						if dosDkg.IsCetified() {
							node.shareSec = *dosDkg.GetShareSecuirty()
							node.groupPubPoly = *dosDkg.GetGroupPublicPoly()
							node.CheckURL()
							node.GenerateRandomNumber()
						}
					}
					if content.Ctype == internalMsg.Cmd_STOPALL {
						os.Exit(1)
					}
					if content.Ctype == internalMsg.Cmd_RESET {
						//dosDkg.Event("reset")
						dosDkg.Reset()
					}
				default:
					fmt.Println("unknown", content)
				}
			}
		}
	}()

	//5)Broadcast message to peers
	//Generate node member IDs
	members := [][]byte{}
	s := []string{}
	for i := 1; i <= nbParticipants; i++ {
		output := fmt.Sprintf("node%d", i)
		hashId := blake2b.Sum256([]byte(output))
		members = append(members, hashId[:])
		s = append(s, output)
	}
	Ids := strings.Join(s, ",")

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

		if strings.TrimRight(input, "\n") == "grouping" {
			members = [][]byte{}
			s = []string{}
			for i := 1; i <= nbParticipants; i++ {
				output := fmt.Sprintf("node%d", i)
				hashId := blake2b.Sum256([]byte(output))
				members = append(members, hashId[:])
				s = append(s, output)
			}
			Ids = strings.Join(s, ",")
			cmd := &internalMsg.Cmd{
				Ctype: internalMsg.Cmd_GROUPING,
				Args:  Ids,
			}
			pb := proto.Message(cmd)
			for _, member := range members {
				p.SendMessageById(member, pb)
			}
			continue
		}
		/*
			if strings.TrimRight(input, "\n") == "grouping-partial" {
				node.nbParticipants = 5
				dosDkg.Reset()
				dosDkg.SetNbParticipants(5)

				members = [][]byte{}
				s = []string{}
				for i := 1; i <= node.nbParticipants; i++ {
					output := fmt.Sprintf("node%d", i)
					hashId := blake2b.Sum256([]byte(output))
					members = append(members, hashId[:])
					s = append(s, output)
				}
				Ids = strings.Join(s, ",")
				cmd := &internalMsg.Cmd{
					Ctype: internalMsg.Cmd_GROUPING,
					Args:  Ids,
				}
				pb := proto.Message(cmd)
				for _, member := range members {
					p.SendMessageById(member, pb)
				}
				continue
			}*/
		if strings.TrimRight(input, "\n") == "checkURL" {

			cmd := &internalMsg.Cmd{
				Ctype: internalMsg.Cmd_CHECKURL,
				Args:  "https://api.coinmarketcap.com/v1/global/",
			}
			pb := proto.Message(cmd)
			for _, member := range members {
				p.SendMessageById(member, pb)
			}

			continue
		}
		if strings.TrimRight(input, "\n") == "stopall" {
			cmd := &internalMsg.Cmd{
				Ctype: internalMsg.Cmd_STOPALL,
				Args:  "",
			}
			pb := proto.Message(cmd)
			for _, member := range members {
				p.SendMessageById(member, pb)
			}
			continue
		}
		if strings.TrimRight(input, "\n") == "reset" {
			cmd := &internalMsg.Cmd{
				Ctype: internalMsg.Cmd_RESET,
				Args:  "",
			}
			pb := proto.Message(cmd)
			for _, member := range members {
				p.SendMessageById(member, pb)
			}
			continue
		}
	}
	fmt.Println("finish)")
}
