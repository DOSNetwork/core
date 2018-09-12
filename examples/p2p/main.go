package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/DOSNetwork/core/p2p"
)

// main
func main() {
	connect := flag.String("connect", "", "IP address of process to join. If empty, go into listen mode.")
	portFlag := flag.String("port", "", "")
	flag.Parse()
	port, _ := strconv.Atoi(*portFlag)
	p := &p2p.P2P{
		Peers:       new(sync.Map),
		Port:        port,
		MessageChan: make(chan []byte),
	}
	//1)Build a p2p network
	//new(sync.Map),

	defer close(p.MessageChan)

	if *connect != "" {
		fmt.Println("Create peerclients")
		peer, _ := p.CreatePeer(*connect, nil)
		go peer.HandleMessages()
	}
	go p.Listen()

	go func() {
		for {
			select {
			//event from peer
			case buf := <-p.MessageChan:
				fmt.Println("messageChan receive value:", string(buf))
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')

		// skip blank lines
		if len(strings.TrimSpace(input)) == 0 {
			continue
		}
		p.Peers.Range(func(key, value interface{}) bool {
			ip := key.(string)
			client := value.(*p2p.PeerClient)
			fmt.Printf("key[%s]\n", ip)
			client.SendMessage([]byte(input))
			return true
		})
	}
}
