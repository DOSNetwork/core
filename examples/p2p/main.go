package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/DOSNetwork/core/examples/p2p/internal"
	"github.com/DOSNetwork/core/p2p"

	"github.com/golang/protobuf/proto"
)

// main
func main() {
	connect := flag.String("connect", "", "IP address of process to join. If empty, go into listen mode.")
	flag.Parse()
	//1)Build a p2p network

	tunnel := make(chan p2p.P2PMessage)
	p, _ := p2p.CreateP2PNetwork(tunnel)
	defer close(tunnel)
	//2)Start to listen incoming connection
	p.Listen()

	//3)Dial to peers to build peerClient
	if *connect != "" {
		fmt.Println("Create peerclients")
		_ = p.CreatePeer(*connect, nil)
		//p.SendMessageById(*connect, []byte("hello"))
	}

	//4)Handle message from peer
	go func() {
		for {
			select {
			//event from peer
			case msg := <-tunnel:
				switch content := msg.Msg.Message.(type) {
				case *internal.Person:
					fmt.Println("receive messages.Person", content.Name, " from ", msg.Sender)
				case *internal.Company:
					fmt.Println("receive messages.Company", content.Id, " from ", msg.Sender)
				case *internal.Chat:
					fmt.Println("receive messages.Chat", content.Content, " from ", msg.Sender)
				default:
					fmt.Println("unknown")
				}
			}
		}
	}()

	//5)Broadcast message to peers
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')

		// skip blank lines
		if len(strings.TrimSpace(input)) == 0 {
			continue
		}
		if strings.TrimRight(input, "\n") == "end" {
			fmt.Println("Stop()")
			break
		}
		if strings.TrimRight(input, "\n") == "person" {
			pb := proto.Message(&internal.Person{Name: "Eric"})
			//raw, _ := ptypes.MarshalAny(pb)
			p.Broadcast(&pb)
			continue
		}
		if strings.TrimRight(input, "\n") == "company" {
			//raw, _ := ptypes.MarshalAny(&internal.Company{Id: 2})
			pb := proto.Message(&internal.Company{Id: 2})
			p.Broadcast(&pb)
			continue
		}
		pb := proto.Message(&internal.Chat{Content: input})
		p.Broadcast(&pb)
	}
	fmt.Println("finish)")
}
