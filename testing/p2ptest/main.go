package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	"github.com/golang/protobuf/proto"
	"github.com/manifoldco/promptui"
)

var messages map[string]p2p.P2PMessage

func receiveEvent(node p2p.P2PInterface) {
	defer fmt.Println("end receiveEvent")
	messages = make(map[string]p2p.P2PMessage)
	go node.Listen()
	events, _ := node.SubscribeEvent(1, p2p.Ping{})
	for msg := range events {
		fmt.Println("receiveEvent")
		messages[strconv.FormatUint(msg.RequestNonce, 10)] = msg
	}
}

func main() {
	id := os.Args[1]
	nodePort := os.Args[2]
	destid := os.Args[3]
	if id == "" || nodePort == "" || destid == "" {
		return
	}
	log.Init([]byte(id))
	var node p2p.P2PInterface

	//	var remoteNode []byte
	validate := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}
	for {
		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		switch result {
		case "p":
			node, _ = p2p.CreateP2PNetwork([]byte(id), "0.0.0.0", nodePort, p2p.NoDiscover)
			go receiveEvent(node)
		case "l":
			if node != nil {
				node.Leave()
			}
		case "c":
			if node != nil {
				node.DisConnectTo([]byte(destid))
			}
		case "s":
			if node != nil {

				for i := 0; i <= 10; i++ {
					go func() {
						ctx := context.Background()
						cmd := &p2p.Ping{Count: 1}
						reply, _ := node.Request(ctx, []byte(destid), proto.Message(cmd))
						pong, _ := reply.Msg.Message.(*p2p.Pong)
						fmt.Println("pong ", pong)
					}()
				}
			}
		case "r":
			if node != nil {
				for key, msg := range messages {
					r, ok := msg.Msg.Message.(*p2p.Ping)
					if !ok {
						fmt.Println("not ok.")
						return
					}
					node.Reply(context.Background(), msg.Sender, msg.RequestNonce, proto.Message(&p2p.Pong{Count: r.Count + 10}))
					delete(messages, key)

				}
			}
		case "exit":
			fmt.Println("exit.")
			os.Exit(0)
		default:
			//fmt.Printf("Not supported command \n")
		}
	}
}
