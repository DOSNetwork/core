package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gorilla/mux"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
)

var (
	lock                sync.Mutex
	credentialIndex     = 0
	addressToDistribute []string
)
var adaptor onchain.ProxyAdapter

// main
func main() {
	//1) Connect to Ethereum to reset contract
	passphrase := os.Getenv(configuration.ENVPASSPHRASE)
	if passphrase == "" {
		log.Fatal(errors.New("No passphrase"))
	}
	//Read Configuration
	config := configuration.Config{}
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return
	}
	if workingDir == "/" {
		workingDir = "."
	}
	credentialPath := workingDir + "/testAccounts/bootCredential"
	//Set up an onchain adapter
	chainConfig := config.GetChainConfig()
	fmt.Println("chainConfig.RemoteNodeAddressPool", chainConfig.RemoteNodeAddressPool)
	addressToDistribute = chainConfig.RemoteNodeAddressPool[1:]
	adaptor, err = onchain.NewProxyAdapter(config.GetCurrentType(), credentialPath, passphrase, chainConfig.DOSProxyAddress, chainConfig.RemoteNodeAddressPool)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := adaptor.Address()
	//Init log module with nodeID that is an onchain account address
	log.Init(id[:])
	ctx, _ := context.WithCancel(context.Background())

	var errcList []<-chan error
	var eventList []<-chan interface{}
	sink, errc := adaptor.PollLogs(onchain.SubscribeDOSProxyLogInsufficientPendingNode, 0, 0)
	eventList = append(eventList, sink)
	errcList = append(errcList, errc)
	sink, errc = adaptor.PollLogs(onchain.SubscribeDOSProxyLogInsufficientWorkingGroup, 0, 0)
	eventList = append(eventList, sink)
	errcList = append(errcList, errc)
	sink, errc = adaptor.PollLogs(onchain.SubscribeDOSProxyLogGroupingInitiated, 0, 0)
	eventList = append(eventList, sink)
	errcList = append(errcList, errc)
	errc = onchain.MergeErrors(ctx, errcList...)
	sink = onchain.MergeEvents(ctx, eventList...)

	groupToPick, err := strconv.Atoi(os.Getenv(configuration.ENVGROUPTOPICK))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("groupToPick: ", groupToPick)
	groupSize, err := strconv.Atoi(os.Getenv(configuration.ENVGROUPSIZE))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("groupSize: ", groupSize)

	go func() {
		for {
			select {
			case event := <-sink:
				switch content := event.(type) {
				case *onchain.DOSProxyLogInsufficientPendingNode:
					fmt.Println("DOSProxyLogInsufficientPendingNode ", content.NumPendingNodes)
				case *onchain.DOSProxyLogInsufficientWorkingGroup:
					fmt.Println("DOSProxyLogInsufficientWorkingGroup ", content.NumPendingNodes)
					if int(content.NumPendingNodes.Uint64()) >= groupSize*(groupToPick+1) {
						errc = adaptor.BootStrap()
						e := <-errc
						fmt.Println("BootStrap done ", e)
					}
				case *onchain.DOSProxyLogGroupingInitiated:
					fmt.Println("DOSProxyLogInsufficientWorkingGroup ", content.NumPendingNodes)
				}
			case e, ok := <-errc:
				if ok {
					err = e
					fmt.Println(err)
				}
			}
		}
	}()
	fmt.Println("ResetContract")
	errc = adaptor.ResetContract()
	fmt.Println("ResetContract", <-errc)

	fmt.Println("SetGroupSize")
	errc = adaptor.SetGroupSize(ctx, uint64(config.GetRandomGroupSize()))
	fmt.Println("ResetContract", <-errc)

	//2)Build a p2p network
	id = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	p, err := p2p.CreateP2PNetwork(id[:], config.Port)
	if err != nil {
		log.Fatal(err)
	}

	defer p.CloseMessagesChannel()

	//2-2)Start to listen incoming connection
	if err = p.Listen(); err != nil {
		log.Fatal(err)
	}

	//3) Declare a new router to handle REST API call
	r := mux.NewRouter()
	// This is where the router is useful, it allows us to declare methods that
	// this path will be valid for
	r.HandleFunc("/getCredential", getCredential).Methods("GET")
	r.HandleFunc("/hasGroupPubkey", hasGroupPubkey).Methods("GET")
	r.HandleFunc("/reset", reset).Methods("GET")

	if err = http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func getCredential(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	credentialIndex++
	response := strconv.Itoa(credentialIndex) + "," + addressToDistribute[credentialIndex%len(addressToDistribute)] + "," + addressToDistribute[(credentialIndex+1)%len(addressToDistribute)]
	fmt.Println("getCredential", response)
	if _, err := fmt.Fprintf(w, response); err != nil {
		fmt.Println(err)
	}
	/*
		go func() {
			if credentialIndex == 12 {
				fmt.Println("!!!!!!!!!!!!!!start BootStrap ")
				time.Sleep(300 * time.Second)

				err := adaptor.BootStrap()
				fmt.Println("BootStrap ", err)
			}
		}()
	*/
	lock.Unlock()

}

func hasGroupPubkey(w http.ResponseWriter, r *http.Request) {
	if adaptor == nil {
		fmt.Println("adaptor is nil")
	}
	key, err := adaptor.GroupPubKey(0)
	if err != nil {
		//TODO: Need to check why it has err : abi: improperly formatted output
	}
	if key[0] == nil {
		if _, err = fmt.Fprintf(w, "false"); err != nil {
			fmt.Println(err)
		}
	} else {
		if _, err = fmt.Fprintf(w, "yes"); err != nil {
			fmt.Println(err)
		}
		//if err = (*adaptor).FireRandom(); err != nil {
		//	fmt.Println(err)
		//}
	}

}

func reset(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	credentialIndex = 0
	lock.Unlock()
}
