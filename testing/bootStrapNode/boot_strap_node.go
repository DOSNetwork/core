package main

import (
	// Import the gorilla/mux library we just installed
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/DOSNetwork/core/p2p"
	"github.com/gorilla/mux"

	"github.com/DOSNetwork/core/configuration"

	"github.com/DOSNetwork/core/onchain"
)

var lock sync.Mutex
var credentialIndex int

var adaptor *onchain.EthAdaptor

// main
func main() {
	var err error
	credentialIndex = 0

	//1) Connect to Ethereum to reset contract
	config := configuration.ReadConfig()
	chainConfig := configuration.GetOnChainConfig(config)
	port := config.Port
	adaptor = &onchain.EthAdaptor{}
	err = adaptor.Init(false, &chainConfig)
	if err != nil {
		fmt.Println(err)
	}
	(*adaptor).ResetNodeIDs()
	(*adaptor).Grouping(config.RandomGroupSize)

	//2)Build a p2p network
	peerEvent := make(chan p2p.P2PMessage, 100)
	p, _ := p2p.CreateP2PNetwork(peerEvent, port)
	defer close(peerEvent)
	//2-1)Set node ID
	id := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	p.SetId(id[:])
	//2-2)Start to listen incoming connection
	go p.Listen()
	//2-3)To ignore peer event to avoid channel blocking
	go func() {
		for {
			select {
			//event from peer
			case _ = <-peerEvent:
			}
		}
	}()

	//3) Declare a new router to handle REST API call
	r := mux.NewRouter()
	// This is where the router is useful, it allows us to declare methods that
	// this path will be valid for
	r.HandleFunc("/getCredential", getCredential).Methods("GET")
	r.HandleFunc("/hasGroupPubkey", hasGroupPubkey).Methods("GET")
	r.HandleFunc("/reset", reset).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func getCredential(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	fmt.Println("getCredential")
	credentialIndex++
	if credentialIndex > 21 {
		credentialIndex = 1
	}
	credentialPath := "testAccounts/" + strconv.Itoa(credentialIndex) + "/credential"

	usrKeyPath := credentialPath + "/usrKey"
	rootKeyPath := credentialPath + "/boot/rootKey"

	passPhraseBytes, err := ioutil.ReadFile(credentialPath + "/boot/passPhrase")
	if err != nil {
		return
	}
	passPhrase := string(passPhraseBytes)
	(*adaptor).BalanceMaintain(rootKeyPath, usrKeyPath, passPhrase, passPhrase)

	credential, err := ioutil.ReadFile(usrKeyPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(credential))
	lock.Unlock()
}

func hasGroupPubkey(w http.ResponseWriter, r *http.Request) {
	key, err := (*adaptor).GetGroupPubKey(0)
	if err != nil {
		//TODO: Need to check why it has err : abi: improperly formatted output
	}
	if key[0] == nil {
		fmt.Fprintf(w, "false")
	} else {
		fmt.Fprintf(w, "yes")
	}
}

func reset(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	credentialIndex = 0
	lock.Unlock()
}
