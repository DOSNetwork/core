package main

import (
	"fmt"
	"log"

	"github.com/DOSNetwork/core/examples/proxy-log-subscribe/contracts" // for demo
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"net/http"
	"time"
	"io/ioutil"
	"github.com/DOSNetwork/core/suites"
	"github.com/DOSNetwork/core/examples/random-number-generator"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"context"
	"math/big"
	"strings"
)

var ethReomteNode = "wss://rinkeby.infura.io/ws"
var contractAddressHex = "0xa479d7446da2d3b258383aea52f79cd25a613a1e"
var stopKeyWord = "https://api.coinbase.com/v2/prices/ETH-USD/spot"
const key = ``
const passphrase = ``


func subscribeEvent(client *ethclient.Client, proxy *dosproxy.Dosproxy, ch chan *dosproxy.DosproxyLogUrl, done chan bool) {
	opt := &bind.WatchOpts{}
	sub, err := proxy.DosproxyFilterer.WatchLogUrl(opt, ch)
	if err != nil {
		log.Fatal(err)
	}
	for i := range ch {
		fmt.Println(i.Url)
		go queryFulfill(client, proxy, big.NewInt(123), i.Url)
		//fmt.Println(i.QueryId)
		//fmt.Println(i.Timeout)
		//fmt.Println(i.Raw)
	}
	sub.Unsubscribe()
	close(ch)
	done <- true
}

func queryFulfill(client *ethclient.Client, proxy *dosproxy.Dosproxy, queryId *big.Int, url string) {
	data, err := dataFetch(url)
	if err != nil {
		log.Fatal(err)
	}
	sig, err := dataSign(data)
	if err != nil {
		log.Fatal(err)
	}
	data += ", sig: " + string(sig)
	err = dataReturn(client, proxy, queryId, data)
	if err != nil {
		log.Fatal(err)
	}
}

func dataFetch(url string) (string, error){
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	r.Body.Close()

	return string(body), nil

}

func dataSign(data string) ([]byte, error){
	suite := suites.MustFind("bn256")
	nbParticipants := 7

	signers, err := example.InitialRandomNumberGenerator(suite, nbParticipants)
	if err != nil {
		return nil, err
	}

	sig, err := signers.TBlsSign([]byte(data))
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func dataReturn(client *ethclient.Client, proxy *dosproxy.Dosproxy, queryId *big.Int, data string) error{
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		return err
	}

	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := proxy.TriggerCallback(auth, queryId, data)
	if err != nil {
		return err
	}

	fmt.Println("tx sent: %s", tx.Hash().Hex())

	return nil
}


func main() {
	client, err := ethclient.Dial(ethReomteNode)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(contractAddressHex)
	proxy, err := dosproxy.NewDosproxy(contractAddress, client)

	ch := make(chan *dosproxy.DosproxyLogUrl)
	done := make(chan bool)
	go subscribeEvent(client, proxy, ch, done)


	<-done
}