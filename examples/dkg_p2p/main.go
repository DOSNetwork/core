package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"sync"
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
	signMap        *sync.Map
	contentMap     *sync.Map
	nbParticipants int
	groupPubPoly   share.PubPoly
	shareSec       share.PriShare
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
func (d *dosNode) CheckURL() {
	url := <-d.chURL
	result, err := d.dataFetch(url)
	if err != nil {
		fmt.Println(err)
	}

	sig, _ := tbls.Sign(d.suite, &d.shareSec, result)
	sign := &vss.Signature{
		Index:     uint32(0),
		QueryId:   url,
		Content:   result,
		Signature: sig,
	}

	var sigShares [][]byte
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
	(*d.signMap).Store(sign.QueryId, sigShares)
	(*d.contentMap).Store(sign.QueryId, sign.Content)

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
	(*d.signMap).Store(sign.QueryId, sigShares)
	(*d.contentMap).Store(sign.QueryId, sign.Content)

	fmt.Println("Event afterReceiveSignature ", len(sigShares))
	if len(sigShares) == d.nbParticipants {
		result, _ := (*d.contentMap).Load(sign.QueryId)
		content, ok := result.([]byte)
		if !ok {
			fmt.Println("afterReceiveSignature value not found for sign.QueryId: ", sign.QueryId)
		}

		sig, _ := tbls.Recover(d.suite, &d.groupPubPoly, content, sigShares, d.nbParticipants/2+1, d.nbParticipants)
		err := bls.Verify(d.suite, d.groupPubPoly.Commit(), content, sig)
		if err == nil {
			fmt.Println("checkURL result = ", string(content), " verify success")
			//fmt.Println(time.Since(d.checkURLStart))
		} else {
			fmt.Println("checkURL result = ", string(content), " verify failed ", err)
		}
		(*d.contentMap).Delete(sign.QueryId)
		(*d.signMap).Delete(sign.QueryId)
	}
}

// main
func main() {
	seedFlag := flag.String("seedAddr", "", "seed address")
	nbParticipantsFlag := flag.Int("nbVerifiers", 20, "Number of Participants")
	portFlag := flag.Int("port", 0, "port number")
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
		signMap:        new(sync.Map),
		contentMap:     new(sync.Map),
		chSignature:    make(chan vss.Signature, 100),
		chURL:          make(chan string, 100),
		nbParticipants: nbParticipants,
		network:        &p,
	}

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
	for i := 1; i <= node.nbParticipants; i++ {
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
				go p.SendMessageById(member, pb)
			}

			cmd = &internalMsg.Cmd{
				Ctype: internalMsg.Cmd_CHECKURL,
				Args:  "https://api.weather.gov/",
			}
			pb = proto.Message(cmd)
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
