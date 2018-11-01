package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/test/userTest/eth"
)

type querySet struct {
	url      string
	selector string
}

var querySets = []querySet{
	{"https://api.coinbase.com/v2/prices/ETH-USD/spot", ""},
	{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "$"},
	{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "$.data"},
	{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "$.data.base"},
	{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "$.data.currency"},
	{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "$.data.amount"},
	{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "$.data.NOTVALID"},
	{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "NOTVALID"},
	{"https://api.coinmarketcap.com/v1/global/", ""},
	{"https://api.coinmarketcap.com/v1/global/", "$"},
	{"https://api.coinmarketcap.com/v1/global/", "$.total_market_cap_usd"},
	{"https://api.coinmarketcap.com/v1/global/", "$.total_24h_volume_usd"},
	{"https://api.coinmarketcap.com/v1/global/", "$.bitcoin_percentage_of_market_cap"},
	{"https://api.coinmarketcap.com/v1/global/", "$.active_currencies"},
	{"https://api.coinmarketcap.com/v1/global/", "$.active_assets"},
	{"https://api.coinmarketcap.com/v1/global/", "$.active_markets"},
	{"https://api.coinmarketcap.com/v1/global/", "$.last_updated"},
	{"https://api.coinmarketcap.com/v1/global/", "$.NOTVALID"},
	{"https://api.coinmarketcap.com/v1/global/", "NOTVALID"},
}

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
	events := make(chan interface{}, 50)
	userTestAdaptor.SubscribeToAll(events)

	ticker := time.NewTicker(3 * time.Minute)
	lastQuery := time.Now()

	lottery := rand.Intn(len(querySets))
	userTestAdaptor.Query(querySets[lottery].url, querySets[lottery].selector)

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
				lottery = rand.Intn(len(querySets))
				userTestAdaptor.Query(querySets[lottery].url, querySets[lottery].selector)
				lastQuery = time.Now()
			default:
				fmt.Println("type mismatch")
			}
		case <-ticker.C:
			if time.Since(lastQuery).Minutes() > 3 {
				lottery = rand.Intn(len(querySets))
				userTestAdaptor.Query(querySets[lottery].url, querySets[lottery].selector)
				lastQuery = time.Now()
			}
		}
	}
}
