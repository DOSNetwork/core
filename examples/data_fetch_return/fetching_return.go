package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/examples/random_number_generator"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
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

	config := configuration.ReadConfig("./config.json")
	chainConfig := configuration.GetOnChainConfig(config)
	chainConn, err := onchain.AdaptTo(onchain.ETH, true, &chainConfig)
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
	chainConn.SubscribeEvent(chUrl, onchain.SubscribeDOSProxyLogUrl)

	fmt.Println("Group set-up finished, start listening to query...")

	for i := range chUrl {
		go func() {
			switch i.(type) {
			case *onchain.DOSProxyLogUrl:
				fmt.Printf("Query-ID: %v \n", i.(*onchain.DOSProxyLogUrl).QueryId)
				fmt.Println("Query Url: ", i.(*onchain.DOSProxyLogUrl).Url)
				data, sig, err := dataPrepare(signers, i.(*onchain.DOSProxyLogUrl).Url)
				if err != nil {
					log.Fatal(err)
				}
				chainConn.DataReturn(i.(*onchain.DOSProxyLogUrl).QueryId, data, sig)
			default:
				fmt.Println("type mismatch")
			}
		}()
	}
}
