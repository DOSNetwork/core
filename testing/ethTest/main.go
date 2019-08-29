package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/manifoldco/promptui"
)

func receiveEvent(chain onchain.ProxyAdapter) {
S:
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(300*time.Second))
	err := chain.Connect(ctx)
	if err != nil {
		fmt.Println("chain.Connect err ", err)
		return
	}
	go chain.ReqLoop()

	subescriptions := []int{onchain.SubscribeDosproxyUpdateGroupToPick, onchain.SubscribeLogGrouping, onchain.SubscribeLogGroupDissolve, onchain.SubscribeLogUrl,
		onchain.SubscribeLogUpdateRandom, onchain.SubscribeLogRequestUserRandom,
		onchain.SubscribeLogPublicKeyAccepted, onchain.SubscribeCommitrevealLogStartCommitreveal}
	onchainEvent, _ := chain.SubscribeEvent(subescriptions)
L:
	for {
		select {
		case event, ok := <-onchainEvent:
			if ok {
				fmt.Println("event %+v", event)
				switch content := event.(type) {
				case *onchain.LogGrouping:

				case *onchain.LogGroupDissolve:

				case *onchain.LogPublicKeyAccepted:

				case *onchain.LogUpdateRandom:

				case *onchain.LogRequestUserRandom:

				case *onchain.LogUrl:

				case *onchain.LogStartCommitReveal:
					fmt.Println("startBlock ", content.StartBlock.String(), " commitDur ", content.CommitDuration.String(), "revealDur", content.RevealDuration.String())
				}
			} else {
				break L
			}
		}
	}
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	//chain.Close()
	chain.UpdateWsUrls([]string{"ws://34.220.51.148:8546", "ws://34.219.209.133:8546"})
	time.Sleep(5 * time.Second)
	goto S
}

func main() {
	credentialPath := "."
	password := "123"
	key, err := onchain.ReadEthKey(credentialPath, password)
	if err != nil {
		fmt.Println("ReadEthKey ", err)
		return
	}
	id := "ethtest"
	log.Init([]byte(id))
	//	var remoteNode []byte
	validate := func(input string) error {
		return nil
	}
	config := configuration.Config{}
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}
	var adaptor onchain.ProxyAdapter
	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}
	i := 4
	for {
		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch result {
		case "new":
			//Set up an onchain adapter
			var err error
			adaptor, err = onchain.NewProxyAdapter(config.ChainType, key, config.DOSAddressBridgeAddress, config.ChainNodePool)
			if err != nil {
				if err.Error() != "No any working eth client for event tracking" {
					fmt.Println("NewDosNode failed ", err)
					return
				}
			}
			go receiveEvent(adaptor)
		case "get":
			if adaptor != nil {
				ctx, cancel := context.WithCancel(context.Background())
				val, e := adaptor.GroupToPick(ctx)
				if e != nil {
					fmt.Println("GroupToPick err ", e)
					return
				}
				fmt.Println("GroupToPick ", val)
				_ = cancel
			}
		case "set":
			if err := adaptor.SetGroupToPick(context.Background(), uint64(i)); err != nil {
				fmt.Println("GroupToPick err ", err)
				continue
			}
			i++
		case "close":
			adaptor.Close()

		case "exit":
			fmt.Println("exit.")
			os.Exit(0)
		default:
			//fmt.Printf("Not supported command \n")
		}
	}
}
