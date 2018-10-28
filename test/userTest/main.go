package main

import (
	"fmt"
	"log"
	"math/big"
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

	randomTurn := true

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
			case *eth.AskMeAnythingCallbackReady:
				fmt.Println("AskMeAnythingCallbackReady")
				fmt.Println("Callback Ready Query id:", i.QueryId)
				fmt.Println("result: ", i.Result)
				fmt.Println("randomNumber", i.RandomNumber)
				fmt.Println("initial new query...")
				if randomTurn {
					userTestAdaptor.GetRandom(0, big.NewInt(0))
					userTestAdaptor.GetRandom(1, big.NewInt(0))
				} else {
					userTestAdaptor.Query(queryUrls[user])
				}
				randomTurn = !randomTurn
				lastQuery = time.Now()
				fmt.Println("____________________________________________")
			case *eth.AskMeAnythingQuerySent:
				fmt.Println("AskMeAnythingQuerySent")
				fmt.Println("TrafficType", i.TrafficType)
				fmt.Println("succ:", i.Succ)
				fmt.Println("Query Id", i.QueryId)
				fmt.Println("____________________________________________")
			default:
				fmt.Println("type mismatch")
			}
		case <-ticker.C:
			if time.Since(lastQuery).Minutes() > 3 {
				userTestAdaptor.Query(queryUrls[user])
				randomTurn = true
			}
		}
	}
}
