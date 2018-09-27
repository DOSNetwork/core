package main

import (
	"fmt"
	"log"
	"testing"

	dosproxy "github.com/DOSNetwork/core/examples/proxy_log_subscribe/contracts" // for demo
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

var proxy *dosproxy.Dosproxy
var ch chan *dosproxy.DosproxyLogUrl
var done chan bool

func init() {
	client, err := ethclient.Dial(ethReomteNode)
	if err != nil {
		log.Fatalf("Failed Dial: %v", err)
	}
	contractAddress := common.HexToAddress(contractAddressHex)
	proxy, err = dosproxy.NewDosproxy(contractAddress, client)

	ch = make(chan *dosproxy.DosproxyLogUrl)
	done = make(chan bool)
}

func TestReadEvents(t *testing.T) {
	results := readEvents(proxy)
	require.NotNil(t, results)
	testresult := false
	for i := range results {
		fmt.Println(i)
		fmt.Printf("%s\n", results[i])
		if results[i] == stopKeyWord {
			testresult = true
			break
		}
	}
	if !testresult {
		t.Fatalf("can't find = %s", stopKeyWord)
	}
}
