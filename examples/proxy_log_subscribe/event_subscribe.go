package main

import (
	"fmt"
	"log"

	dosproxy "github.com/DOSNetwork/core/examples/proxy_log_subscribe/contracts" // for demo
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ethReomteNode = "wss://rinkeby.infura.io/ws"
var contractAddressHex = "0xa479d7446da2d3b258383aea52f79cd25a613a1e"
var stopKeyWord = "https://api.coinbase.com/v2/prices/ETH-USD/spot"

func readEvents(proxy *dosproxy.Dosproxy) []string {
	s := make([]string, 0)
	opt := &bind.FilterOpts{}
	past, err := proxy.DosproxyFilterer.FilterLogUrl(opt)
	if err != nil {
		log.Fatalf("Failed FilterLogUrl: %v", err)
	}
	notEmpty := true
	for notEmpty {
		notEmpty = past.Next()
		if notEmpty {
			//fmt.Println("event log:", past.Event.QueryId)
			//fmt.Println("event log:", past.Event.Timeout)
			//fmt.Println("event log:", past.Event.Url)
			s = append(s, past.Event.Url)
		}
	}
	//fmt.Println("event log:", s)
	return s
}

func subscribeEvent(proxy *dosproxy.Dosproxy, ch chan *dosproxy.DosproxyLogUrl, done chan bool) {
	opt := &bind.WatchOpts{}
	sub, err := proxy.DosproxyFilterer.WatchLogUrl(opt, ch)
	if err != nil {
		log.Fatal(err)
	}
	for i := range ch {
		//fmt.Println(i.QueryId)
		//fmt.Println(i.Timeout)
		fmt.Println(i.Url)
		//fmt.Println(i.Raw)
		if i.Url == stopKeyWord {
			sub.Unsubscribe()
			close(ch)
			done <- true
			break
		}
	}
}

func main() {
	client, err := ethclient.Dial(ethReomteNode)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(contractAddressHex)
	proxy, err := dosproxy.NewDosproxy(contractAddress, client)

	readEvents(proxy)

	ch := make(chan *dosproxy.DosproxyLogUrl)
	done := make(chan bool)
	go subscribeEvent(proxy, ch, done)

	<-done
}
