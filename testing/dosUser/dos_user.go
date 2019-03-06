package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/log"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/testing/dosUser/eth"
)

const (
	ENVQUERYTIMES = "QUERYTIMES"
	ENVQUERYTYPE  = "QUERYTYPE"
)

const (
	INVALIDQUERYINDEX = 17
	CHECKINTERVAL     = 3
	FINALREPORTDUE    = 10
	SMALLLOGBLOCKDIFF = 1
	LARGELOGBLOCKDIFF = 3
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
	envTimes        = ""
	envTypes        = ""
	userTestAdaptor = &eth.EthUserAdaptor{}
	counter         = 10
	totalQuery      = 0
	invalidQuery    = 0
	rMap            = make(map[string]string)
	canceled        = make(chan struct{})
	done            = make(chan struct{})
)
var logger log.Logger
var wg sync.WaitGroup

func main() {
	var err error

	envTypes = os.Getenv(ENVQUERYTYPE)
	if envTypes == "" {
		envTypes = "random"
	}

	envTimes = os.Getenv(ENVQUERYTIMES)
	if envTimes != "" {
		counter, err = strconv.Atoi(envTimes)
		if err != nil {
			log.Error(err)
		}
	}

	//It also need to connect to bootstrape node to get crential
	bootStrapIP := os.Getenv("BOOTSTRAPIP")
	s := strings.Split(bootStrapIP, ":")
	ip, _ := s[0], s[1]

	//
	config := eth.AMAConfig{}
	err = configuration.LoadConfig("./ama.json", &config)
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

	if onChainConfig.NodeRole == "testNode" {
		var credential []byte
		var resp *http.Response
		s := strings.Split(onChainConfig.BootStrapIp, ":")
		ip, _ := s[0], s[1]
		tServer := "http://" + ip + ":8080/getCredential"
		resp, err = http.Get(tServer)
		for err != nil {
			time.Sleep(1 * time.Second)
			resp, err = http.Get(tServer)
		}

		credential, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		if err = resp.Body.Close(); err != nil {
			return
		}

		credentialPath = workingDir + "/testAccounts/" + string(credential) + "/credential"
	} else if credentialPath == "" {
		credentialPath = workingDir + "/credential"
	}
	chainConfig := onChainConfig.GetChainConfig()

	//Wait until contract has group public key
	for {
		tServer := "http://" + ip + ":8080/hasGroupPubkey"
		// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
		resp, err := http.Get(tServer)
		for err != nil {
			log.Error(err)
			time.Sleep(10 * time.Second)
			resp, err = http.Get(tServer)
		}

		r, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
		}

		if string(r) == "yes" {
			err = resp.Body.Close()
			if err != nil {
				log.Error(err)
			}
			break
		}
	}
	passphrase := os.Getenv(configuration.ENVPASSPHRASE)

	envSize := os.Getenv("AMA")
	index, err := strconv.Atoi(envSize)
	if err != nil {
		log.Fatal(err)
	}

	userTestAdaptor, err = eth.NewAMAUserSession(credentialPath, passphrase, config.AskMeAnythingAddressPool[index], chainConfig.RemoteNodeAddressPool)

	log.Init(userTestAdaptor.Address().Bytes()[:])

	if logger == nil {
		logger = log.New("module", "AMAUser")
	}
	events := make(chan interface{}, 5)
	userTestAdaptor.PollLogs(eth.SubscribeAskMeAnythingQueryResponseReady, events, LARGELOGBLOCKDIFF)
	userTestAdaptor.PollLogs(eth.SubscribeAskMeAnythingRequestSent, events, SMALLLOGBLOCKDIFF)
	userTestAdaptor.PollLogs(eth.SubscribeAskMeAnythingRandomReady, events, LARGELOGBLOCKDIFF)
	requestIdMap := make(map[string]bool)

	go func() {
		for {
			select {
			case event := <-events:
				switch i := event.(type) {
				case *eth.AskMeAnythingSetTimeout:

				case *eth.AskMeAnythingQueryResponseReady:
					requestId := fmt.Sprintf("%x", i.QueryId)
					f := map[string]interface{}{
						"RequestId": requestId,
						"Result":    i.Result,
						"Removed":   i.Removed}
					logger.Event("AMAResponseReady", f)
					if requestIdMap[requestId] {
						query(counter)
						counter--
					} else {
						logger.Event("RequestIdNotFound", map[string]interface{}{"RequestId": requestId})
					}
				case *eth.AskMeAnythingRequestSent:
					requestId := fmt.Sprintf("%x", i.RequestId)
					f := map[string]interface{}{
						"RequestId": requestId,
						"Removed":   i.Removed}
					logger.Event("AMARequestSent", f)
					requestIdMap[requestId] = true
				case *eth.AskMeAnythingRandomReady:
					requestId := fmt.Sprintf("%x", i.RequestId)
					f := map[string]interface{}{
						"RequestId":       fmt.Sprintf("%x", requestId),
						"GeneratedRandom": fmt.Sprintf("%x", i.GeneratedRandom),
						"Removed":         i.Removed}
					logger.Event("AMARandomReady", f)
					if requestIdMap[requestId] {
						query(counter)
						counter--
					} else {
						logger.Event("RequestIdNotFound", map[string]interface{}{"RequestId": requestId})
					}
				}
			}
		}
	}()

	query(counter)
	counter--

	finish := make(chan bool)
	<-finish
}

func query(counter int) {
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
		if lottery >= INVALIDQUERYINDEX {
			invalidQuery++
		}
	case "random":
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
			if lottery >= INVALIDQUERYINDEX {
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
