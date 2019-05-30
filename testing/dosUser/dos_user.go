package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/DOSNetwork/core/log"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/testing/dosUser/eth"
)

const (
	queryTimes = "QUERYTIMES"
	queryType  = "QUERYTYPE"
)

const (
	invalidQueryIndex = 17
)

type record struct {
	start   time.Time
	end     time.Time
	version uint8
}

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
	//{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "$.data.NOTVALID"},
	{"https://api.coinmarketcap.com/v1/global/", ""},
	{"https://api.coinmarketcap.com/v1/global/", "$"},
	{"https://api.coinmarketcap.com/v1/global/", "$.total_market_cap_usd"},
	{"https://api.coinmarketcap.com/v1/global/", "$.total_24h_volume_usd"},
	{"https://api.coinmarketcap.com/v1/global/", "$.bitcoin_percentage_of_market_cap"},
	{"https://api.coinmarketcap.com/v1/global/", "$.active_currencies"},
	{"https://api.coinmarketcap.com/v1/global/", "$.active_assets"},
	{"https://api.coinmarketcap.com/v1/global/", "$.active_markets"},
	{"https://api.coinmarketcap.com/v1/global/", "$.last_updated"},
	//{"https://api.coinmarketcap.com/v1/global/", "$.NOTVALID"},

	//frequent-update queries
	//{"https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD,JPY,EUR", "$"},

	//invalid queries
	//{"https://api.coinbase.com/v2/prices/ETH-USD/spot", "NOTVALID"},
	//{"https://api.coinmarketcap.com/v1/global/", "NOTVALID"},
}

var (
	envTimes     = ""
	envTypes     = ""
	counter      = 10
	totalQuery   = 0
	invalidQuery = 0
	rMap         = make(map[string]string)
	canceled     = make(chan struct{})
	done         = make(chan struct{})
)
var logger log.Logger
var wg sync.WaitGroup

func main() {
	var err error
	envTypes = os.Getenv(queryType)
	if envTypes == "" {
		envTypes = "random"
	}

	envTimes = os.Getenv(queryTimes)
	if envTimes != "" {
		counter, err = strconv.Atoi(envTimes)
		if err != nil {
			log.Error(err)
		}
	}

	//It also need to connect to bootstrape node to get crential

	//
	c := eth.AMAConfig{}
	err = configuration.LoadConfig("./ama.json", &c)
	if err != nil {
		log.Fatal(err)
	}

	onChainConfig := configuration.Config{}
	if err = onChainConfig.LoadConfig(); err != nil {
		log.Fatal(err)
	}
	credentialPath := ""
	workingDir, err := os.Getwd()
	if err != nil {
		return
	}
	if workingDir == "/" {
		workingDir = "."
	}

	if credentialPath == "" {
		credentialPath = workingDir + "/credential"
	}
	chainConfig := onChainConfig.GetChainConfig()

	passphrase := os.Getenv(configuration.ENVPASSPHRASE)

	envSize := os.Getenv("AMA")
	index, err := strconv.Atoi(envSize)
	if err != nil {
		log.Fatal(err)
	}

	userTestAdaptor, err := eth.NewAMAUserSession(credentialPath, passphrase, c.AskMeAnythingAddressPool[index], chainConfig.RemoteNodeAddressPool)
	if err != nil {
		log.Fatal(err)
		fmt.Println("NewAMAUserSession err ", err)
	}
	fmt.Println("userTestAdaptor  ")
	log.Init(userTestAdaptor.Address().Bytes()[:])

	if logger == nil {
		logger = log.New("module", "AMAUser")
	}
	var events []<-chan interface{}
	var errcs []<-chan error
	event, errc := userTestAdaptor.SubscribeEvent(eth.SubscribeAskMeAnythingQueryResponseReady)
	events = append(events, event)
	errcs = append(errcs, errc)
	event, errc = userTestAdaptor.SubscribeEvent(eth.SubscribeAskMeAnythingRequestSent)
	events = append(events, event)
	errcs = append(errcs, errc)
	event, errc = userTestAdaptor.SubscribeEvent(eth.SubscribeAskMeAnythingRandomReady)
	events = append(events, event)
	errcs = append(errcs, errc)
	event = eth.MergeEvents(context.Background(), events...)
	errc = eth.MergeErrors(context.Background(), errcs...)

	requestIdMap := make(map[string]bool)

	go func() {
		for {
			select {
			case e := <-event:
				switch i := e.(type) {
				case *eth.AskMeAnythingSetTimeout:

				case *eth.AskMeAnythingQueryResponseReady:
					requestId := fmt.Sprintf("%x", i.QueryId)
					f := map[string]interface{}{
						"RequestId": requestId,
						"Result":    i.Result,
						"Removed":   i.Removed}
					logger.Event("AMAResponseReady", f)
					if requestIdMap[requestId] {
						query(counter, userTestAdaptor)
						counter--
					} else {
						logger.Event("RequestIdNotFound", map[string]interface{}{"RequestId": requestId})
					}
				case *eth.AskMeAnythingRequestSent:
					requestId := fmt.Sprintf("%x", i.RequestId)
					fmt.Println("RequestID ", requestId)
					f := map[string]interface{}{
						"RequestId": requestId,
						"Removed":   i.Removed}
					logger.Event("AMARequestSent", f)
					requestIdMap[requestId] = true
				case *eth.AskMeAnythingRandomReady:
					requestId := fmt.Sprintf("%x", i.RequestId)
					f := map[string]interface{}{
						"RequestId":       requestId,
						"GeneratedRandom": fmt.Sprintf("%x", i.GeneratedRandom),
						"Removed":         i.Removed}
					logger.Event("AMARandomReady", f)
					if requestIdMap[requestId] {
						query(counter, userTestAdaptor)
						counter--
					} else {
						logger.Event("RequestIdNotFound", map[string]interface{}{"RequestId": requestId})
					}
				}
			}
		}
	}()

	query(counter, userTestAdaptor)
	counter--

	for e := range errc {
		fmt.Println(e)
	}
}

func query(counter int, userTestAdaptor *eth.EthUserAdaptor) {
	if counter == 0 {
		os.Exit(0)
	}
	fmt.Println("query counter ", counter)
	f := map[string]interface{}{
		"Removed": false}
	logger.Event("AMAQueryCall", f)
	switch envTypes {
	case "url":
		lottery := rand.Intn(len(querySets))
		if err := userTestAdaptor.Query(uint8(counter), querySets[lottery].url, querySets[lottery].selector); err != nil {
			fmt.Println(err)
			return
		}
		if lottery >= invalidQueryIndex {
			invalidQuery++
		}
	case "random":
		if userTestAdaptor == nil {
			fmt.Println("userTestAdaptor is nil")
		}
		if err := userTestAdaptor.GetSafeRandom(uint8(counter)); err != nil {
			fmt.Println(err)
			return
		}
	default:
		if counter%2 == 0 {
			lottery := rand.Intn(len(querySets))
			if err := userTestAdaptor.Query(uint8(counter), querySets[lottery].url, querySets[lottery].selector); err != nil {
				fmt.Println(err)
				return
			}
			if lottery >= invalidQueryIndex {
				invalidQuery++
			}
		} else {
			if err := userTestAdaptor.GetSafeRandom(uint8(counter)); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	totalQuery++
}
