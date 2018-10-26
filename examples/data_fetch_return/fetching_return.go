package main

import (
	"encoding/json"
	"fmt"
	"github.com/dedis/kyber"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/DOSNetwork/core/examples/random_number_generator"
	"github.com/DOSNetwork/core/onchain"
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

type NetCofigs struct {
	NetCofigs []onchain.NetConfig
}

func readConfig() (node *onchain.NetConfig) {

	var configs NetCofigs
	// Open our jsonFile
	jsonFile, err := os.Open("./config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened NetCofigs json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &configs)
	if err != nil {
		log.Fatal(err)
	}

	targetNode := os.Getenv("TargetNode")
	if targetNode == "" {
		fmt.Println("No TargetNode Environment variable.")
		targetNode = "rinkebyPrivateNode"
	}

	for _, config := range configs.NetCofigs {
		if targetNode == config.RemoteNodeType {
			fmt.Println("Use : ", config)
			return &config
		}
	}
	return nil
}

func main() {
	signers, pubKey, err := groupSetup(7)
	if err != nil {
		log.Fatal(err)
	}

	config := readConfig()
	chainConn, err := onchain.AdaptTo(onchain.ETH, true, config)
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
