package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/test/userTest/eth"
)

var queryUrls = []string{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "https://api.coinmarketcap.com/v1/global/"}

func main() {
	user := 1
	startQuery(user)
}

func startQuery(user int) {
	config := configuration.ReadConfig("./config.json")
	chainConfig := configuration.GetOnChainConfig(config)
	userTestAdaptor := &eth.EthUserAdaptor{}
	err := userTestAdaptor.Init(true, &chainConfig)
	if err != nil {
		log.Fatal(err)
	}
	events := make(chan interface{})
	userTestAdaptor.SubscribeToAll(events)

	ticker := time.NewTicker(3 * time.Minute)
	lastQuery := time.Now()

	userTestAdaptor.Query(queryUrls[user])

	for {
		select {
		case event := <-events:
			fmt.Println("from user:", user)
			switch i := event.(type) {
			case *eth.AskMeAnythingSetTimeout:
				fmt.Println("AskMeAnythingSetTimeout")
				fmt.Println("new timeout:", i.NewTimeout)
				fmt.Println("previous timeout:", i.PreviousTimeout)
				fmt.Println("____________________________________________")
			case *eth.AskMeAnythingQueryResponseReady:
				fmt.Println("AskMeAnythingQueryResponseReady")
				fmt.Println("Callback Ready Query id:", i.QueryId)
				fmt.Println("result: ", i.Result)
				fmt.Println("initial new query...")
				fmt.Println("____________________________________________")
				userTestAdaptor.GetRandom()
				lastQuery = time.Now()
			case *eth.AskMeAnythingRequestSent:
				fmt.Println("AskMeAnythingRequestSent")
				fmt.Println("succ:", i.Succ)
				fmt.Println("RequestId", i.RequestId)
				fmt.Println("____________________________________________")
			case *eth.AskMeAnythingRandomReady:
				fmt.Println("AskMeAnythingRandomReady")
				fmt.Println("GeneratedRandom:", i.GeneratedRandom)
				fmt.Println("____________________________________________")
				userTestAdaptor.Query(queryUrls[user])
				lastQuery = time.Now()
			default:
				fmt.Println("type mismatch")
			}
		case <-ticker.C:
			if time.Since(lastQuery).Minutes() > 3 {
				userTestAdaptor.Query(queryUrls[user])
			}
		}
	}
}
