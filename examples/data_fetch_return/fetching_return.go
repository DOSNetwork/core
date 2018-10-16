package main

import (
	"fmt"
	"github.com/dedis/kyber"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/DOSNetwork/core/blockchain"
	"github.com/DOSNetwork/core/blockchain/eth"
	"github.com/DOSNetwork/core/examples/random_number_generator"
	"github.com/DOSNetwork/core/suites"
)

var groupId = big.NewInt(123)

func dataPrepare(signers *example.RandomNumberGenerator, url string) (data, sig []byte, err error) {
	data, err = dataFetch(url)
	if err != nil {
		return
	}

	sig, err = dataSign(signers, data)
	if err != nil {
		return
	}

	return
}

func dataFetch(url string) ([]byte, error) {
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

	fmt.Println("Fetched data: ", string(body))

	return body, nil

}

func dataSign(signers *example.RandomNumberGenerator, data []byte) ([]byte, error) {

	sig, err := signers.TBlsSign(data)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func groupSetup(nbParticipants int) (signers *example.RandomNumberGenerator, pubKey kyber.Point, err error) {
	suite := suites.MustFind("bn256")
	signers, err = example.InitialRandomNumberGenerator(suite, nbParticipants)
	if err != nil {
		return nil, nil, err
	}

	pubKey = signers.GetPubKey()

	return
}

func main() {
	signers, pubKey, err := groupSetup(7)
	if err != nil {
		log.Fatal(err)
	}

	chainConn, err := blockchain.AdaptTo(blockchain.ETH, true, eth.Private)
	if err != nil {
		log.Fatal(err)
	}

	err = chainConn.UploadID()
	if err != nil {
		log.Fatal(err)
	}

	err = chainConn.UploadPubKey(pubKey)
	if err != nil {
		log.Fatal(err)
	}

	chUrl := make(chan interface{})
	chainConn.SubscribeEvent(chUrl, eth.SubscribeDOSProxyLogUrl)

	fmt.Println("Group set-up finished, start listening to query...")

	for i := range chUrl {
		go func() {
			switch i.(type) {
			case *eth.DOSProxyLogUrl:
				fmt.Printf("Query-ID: %v \n", i.(*eth.DOSProxyLogUrl).QueryId)
				fmt.Println("Query Url: ", i.(*eth.DOSProxyLogUrl).Url)
				data, sig, err := dataPrepare(signers, i.(*eth.DOSProxyLogUrl).Url)
				if err != nil {
					log.Fatal(err)
				}
				chainConn.DataReturn(i.(*eth.DOSProxyLogUrl).QueryId, data, sig)
			default:
				fmt.Println("type mismatch")
			}
		}()
	}
}
