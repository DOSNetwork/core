package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/DOSNetwork/core/examples/data_fetch_return/contract" // for demo
	"github.com/DOSNetwork/core/examples/random_number_generator"
	"github.com/DOSNetwork/core/group/bn256"
	"github.com/DOSNetwork/core/suites"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = ``
const passphrase = ``

var ethReomteNode = "wss://rinkeby.infura.io/ws"
var contractAddressHex = "0xbD5784b224D40213df1F9eeb572961E2a859Cb80"
var stopKeyWord = "https://api.coinbase.com/v2/prices/ETH-USD/spot"
var groupId = big.NewInt(123)


func subscribePubKey(proxy *dosproxy.DOSProxy, ch chan *dosproxy.DOSProxyLogSuccPubKeySub, done chan bool) {

	fmt.Println("Establishing listen channel to group public key...")

	connected := false

	retried := false

	go func() {
		time.Sleep(3 * time.Second)
		if !connected {
			retried = true
			fmt.Println("retry")


			client, err := ethclient.Dial(ethReomteNode)
			if err != nil {
				log.Fatal(err)
			}

			contractAddress := common.HexToAddress(contractAddressHex)
			proxy, err = dosproxy.NewDOSProxy(contractAddress, client)
			if err != nil {
				log.Fatal(err)
			}

			go subscribePubKey(proxy, ch, done)
		}
	}()

	opt := &bind.WatchOpts{}
	sub, err := proxy.DOSProxyFilterer.WatchLogSuccPubKeySub(opt, ch)
	if err != nil {
		log.Fatal(err)
	}

	if retried {
		return
	}

	connected = true

	done <- true

	fmt.Println("Channel established.")

	for range ch {
		fmt.Println("Group public key submitted event Caught.")
		done <- true
		break
	}
	sub.Unsubscribe()
	close(ch)
}


func subscribeEvent(client *ethclient.Client, proxy *dosproxy.DOSProxy, ch chan *dosproxy.DOSProxyLogUrl, signers *example.RandomNumberGenerator) {
	connected := false

	retried := false

	go func() {
		time.Sleep(3 * time.Second)
		if !connected {
			retried = true
			fmt.Println("retry")

			client, err := ethclient.Dial(ethReomteNode)
			if err != nil {
				log.Fatal(err)
			}

			contractAddress := common.HexToAddress(contractAddressHex)
			proxy, err = dosproxy.NewDOSProxy(contractAddress, client)
			if err != nil {
				log.Fatal(err)
			}

			go subscribeEvent(client, proxy, ch, signers)
		}
	}()
	opt := &bind.WatchOpts{}
	sub, err := proxy.DOSProxyFilterer.WatchLogUrl(opt, ch)
	if err != nil {
		log.Fatal(err)
	}

	if retried {
		return
	}

	connected = true

	for i := range ch {
		fmt.Printf("Query-ID: %v \n", i.QueryId)
		fmt.Println("Query Url: ", i.Url)
		go queryFulfill(client, proxy, signers, i.QueryId, i.Url)
	}
	sub.Unsubscribe()
	close(ch)
}

func queryFulfill(client *ethclient.Client, proxy *dosproxy.DOSProxy, signers *example.RandomNumberGenerator, queryId *big.Int, url string) {
	data, err := dataFetch(url)
	if err != nil {
		log.Fatal(err)
	}

	sig, err := dataSign(signers, data)
	if err != nil {
		log.Fatal(err)
	}

	x, y := negate(sig)

	err = dataReturn(client, proxy, queryId, data, x, y)
	if err != nil {
		log.Fatal(err)
	}
}

func dataFetch(url string) ([]byte, error){
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body.Close()

	fmt.Println("Detched data: ", string(body))

	return body, nil

}

func dataSign(signers *example.RandomNumberGenerator, data []byte) ([]byte, error){

	sig, err := signers.TBlsSign(data)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func dataReturn(client *ethclient.Client, proxy *dosproxy.DOSProxy, queryId *big.Int, data []byte, x, y *big.Int) error{
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		return err
	}

	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice

	tx, err := proxy.TriggerCallback(auth, queryId, data, x, y)
	if err != nil {
		return err
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Printf("Query_ID %v request fulfilled \n", queryId)

	return nil
}


func negate(sig []byte) (*big.Int, *big.Int) {
	x := big.NewInt(0)
	y := big.NewInt(0)
	x.SetBytes(sig[0:32])
	y.SetBytes(sig[32:])

	if x.Cmp(big.NewInt(0)) == 0 && y.Cmp(big.NewInt(0)) == 0{
		return big.NewInt(0), big.NewInt(0)
	}

	return x, big.NewInt(0).Sub(bn256.P,big.NewInt(0).Mod(y, bn256.P))
}

func groupSetup(nbParticipants int) (*example.RandomNumberGenerator, *big.Int, *big.Int, *big.Int, *big.Int, error){
	suite := suites.MustFind("bn256")
	signers, err := example.InitialRandomNumberGenerator(suite, nbParticipants)
	if err != nil {
		return nil,nil,nil,nil,nil,err
	}

	x0, x1, y0, y1, err := signers.GetPubKey()
	if err != nil {
		return nil,nil,nil,nil,nil,err
	}

	return signers, x0, x1, y0, y1, nil
}

func uploadPubKey(client *ethclient.Client, proxy *dosproxy.DOSProxy, groupId, x0, x1, y0, y1 *big.Int) error{

	fmt.Println("Starting submitting group public key...")

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

	tx, err := proxy.SetPublicKey(auth, groupId, x0, x1, y0, y1)

	if err != nil {
		return err
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("groupId: ", groupId)
	fmt.Println("x0: ", x0)
	fmt.Println("x1: ", x1)
	fmt.Println("y0: ", y0)
	fmt.Println("y1: ", y1)
	fmt.Println("Group public key submitted, waiting for confirmation...")

	return nil
}


func main() {

	signers, x0, x1, y0, y1, err := groupSetup(7)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(ethReomteNode)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(contractAddressHex)
	proxy, err := dosproxy.NewDOSProxy(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start")

	chPubKey := make(chan *dosproxy.DOSProxyLogSuccPubKeySub)
	done := make(chan bool)
	go subscribePubKey(proxy, chPubKey, done)

	<-done

	err = uploadPubKey(client, proxy, groupId, x0, x1, y0, y1)
	if err != nil {
		log.Fatal(err)
	}

	<- done
	fmt.Println("Group set-up finished, start listening to query...")

	chUrl := make(chan *dosproxy.DOSProxyLogUrl)
	go subscribeEvent(client, proxy, chUrl, signers)
	<- done
}