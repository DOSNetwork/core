package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/testing/dhtNode/internalMsg"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/suites"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

var lock sync.Mutex
var count int
var dhtSize int
var nodeID []byte
var members [][]byte
var p p2p.P2PInterface

// main
func main() {
	var ip string
	var err error
	members = [][]byte{}
	count = 0

	//1)Load config
	offChainConfig := configuration.OffChainConfig{}
	offChainConfig.LoadConfig()
	port := offChainConfig.Port
	//It also need to connect to bootstrape node to get crential
	bootStrapIP := os.Getenv("BOOTSTRAPIP")
	noderole := os.Getenv("NODEROLE")
	dhtSize, err = strconv.Atoi(os.Getenv("DHTSIZE"))
	if err != nil {
		fmt.Println(err)
	}

	//2)Build a p2p network
	_ = suites.MustFind("bn256")
	peerEvent := make(chan p2p.P2PMessage, 100)
	p, _ = p2p.CreateP2PNetwork(peerEvent, port)
	defer close(peerEvent)

	//boot node
	if noderole == "boot" {
		nodeID = []byte{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		p.SetId(nodeID[:])
		go p.Listen()
		//Declare a new router to handle REST API call
		r := mux.NewRouter()
		r.HandleFunc("/getID", getID).Methods("GET")
		go http.ListenAndServe(":8080", r)

	} else {
		s := strings.Split(bootStrapIP, ":")
		ip, _ = s[0], s[1]
		//Wait until bootstrap node assign an ID
		for {
			tServer := "http://" + ip + ":8080/getID"
			resp, err := http.Get(tServer)
			if err != nil {
				fmt.Println(err)
			}
			for err != nil {
				///	fmt.Println(err)
				time.Sleep(10 * time.Second)
				resp, err = http.Get(tServer)
			}

			r, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			} else {
				nodeID = r
				break
			}
		}
		p.SetId(nodeID[:])
		fmt.Println("nodeID = ", nodeID[:])
		//Start to listen incoming connection
		go p.Listen()
		//TODO : Fix race conditation between Listen and CreatePeer
		time.Sleep(2 * time.Second)

		p.CreatePeer(bootStrapIP, nil)
		results := p.FindNode(p.GetId(), dht.BucketSize, 20)
		for _, result := range results {
			p.GetRoutingTable().Update(result)
			fmt.Println(p.GetId().Address, "Update peer: ", result.Address)
		}
	}

	groupIds := [][]byte{}
	for {
		select {
		//event from peer
		case msg := <-peerEvent:
			switch content := msg.Msg.Message.(type) {
			case *internalMsg.Cmd:
				if content.Ctype == internalMsg.Cmd_FINDNODE {
					nodes := strings.Split(content.Args, ",")
					for _, node := range nodes {
						Id := []byte(node)
						groupIds = append(groupIds, Id[:])
					}
					fmt.Println(groupIds)

					cmd := &internalMsg.Cmd{
						Ctype: 2,
						Args:  "test",
					}
					pb := proto.Message(cmd)
					for _, member := range groupIds {
						if string(member[:]) != string(p.GetId().Id) {
							fmt.Println("send message to ", member[:])
							p.SendMessageById(member, pb)
						}
					}
				} else {
					fmt.Println("internalMsg.Cmd args ", content.Args)
				}
			default:
				fmt.Println(content)
			}
		}
	}
}

func getID(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	count++
	lock.Unlock()
	clientID := make([]byte, len(nodeID))
	copy(clientID, nodeID)
	clientID[0] = byte(count)
	fmt.Println(nodeID)
	fmt.Fprintf(w, string(clientID))
	members = append(members, clientID)

	fmt.Println("count ", count, " dhtSize", dhtSize, " members", members)
	if count == dhtSize {
		go requestGrouping()
	}
}

func requestGrouping() {
	fmt.Println("requestGrouping ", members)
	time.Sleep(5 * time.Second)
	s := []string{}
	for i := 0; i < len(members); i++ {
		s = append(s, string(members[i]))
	}
	Ids := strings.Join(s, ",")
	cmd := &internalMsg.Cmd{
		Ctype: internalMsg.Cmd_FINDNODE,
		Args:  Ids,
	}
	pb := proto.Message(cmd)
	for _, member := range members {
		p.SendMessageById(member, pb)
	}
}
