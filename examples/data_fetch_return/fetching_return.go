package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/DOSNetwork/core/blockchain"
	"github.com/DOSNetwork/core/blockchain/eth"
	"github.com/DOSNetwork/core/examples/random_number_generator"
	"github.com/DOSNetwork/core/group/bn256"
	"github.com/DOSNetwork/core/suites"
)

var groupId = big.NewInt(123)

func dataPrepare(signers *example.RandomNumberGenerator, url string) (data []byte, x *big.Int, y *big.Int, err error) {
	data, err = dataFetch(url)
	if err != nil {
		return
	}

	sig, err := dataSign(signers, data)
	if err != nil {
		return
	}

	x, y = negate(sig)
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

func negate(sig []byte) (*big.Int, *big.Int) {
	x := big.NewInt(0)
	y := big.NewInt(0)
	x.SetBytes(sig[0:32])
	y.SetBytes(sig[32:])

	if x.Cmp(big.NewInt(0)) == 0 && y.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), big.NewInt(0)
	}

	return x, big.NewInt(0).Sub(bn256.P, big.NewInt(0).Mod(y, bn256.P))
}

func groupSetup(nbParticipants int) (*example.RandomNumberGenerator, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	suite := suites.MustFind("bn256")
	signers, err := example.InitialRandomNumberGenerator(suite, nbParticipants)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	x0, x1, y0, y1, err := signers.GetPubKey()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return signers, x0, x1, y0, y1, nil
}

func main() {
	signers, x0, x1, y0, y1, err := groupSetup(7)
	if err != nil {
		log.Fatal(err)
	}

	chainConn, err := blockchain.AdaptTo("ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start")

	chPubKey := make(chan interface{})
	go func() {
		chPubKey <- &eth.DOSProxyLogSuccPubKeySub{}
	}()
	chainConn.SubscribeEvent(chPubKey)

	err = chainConn.UploadPubKey(groupId, x0, x1, y0, y1)
	if err != nil {
		log.Fatal(err)
	}

	<-chPubKey

	chUrl := make(chan interface{})
	go func() {
		chUrl <- &eth.DOSProxyLogUrl{}
	}()
	chainConn.SubscribeEvent(chUrl)

	fmt.Println("Group set-up finished, start listening to query...")

	for i := range chUrl {
		go func() {
			switch i.(type) {
			case *eth.DOSProxyLogUrl:
				fmt.Printf("Query-ID: %v \n", i.(*eth.DOSProxyLogUrl).QueryId)
				fmt.Println("Query Url: ", i.(*eth.DOSProxyLogUrl).Url)
				data, x, y, err := dataPrepare(signers, i.(*eth.DOSProxyLogUrl).Url)
				if err != nil {
					log.Fatal(err)
				}
				chainConn.DataReturn(i.(*eth.DOSProxyLogUrl).QueryId, data, x, y)
			default:
				fmt.Println("type mismatch")
			}
		}()
	}
}
