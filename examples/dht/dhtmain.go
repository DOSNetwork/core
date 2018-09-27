package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/p2p/dht"
)

func main() {
	finish := make(chan bool)
	target := flag.Bool("target", false, "set to true if it is the target of dht node finding")
	address := flag.String("address", "", "address to be connected")
	find := flag.Bool("find", false, "set to true if current node is going to find another node")
	port := flag.Int("port", 0, "port to listen to")
	flag.Parse()

	//1)Build a p2p network
	tunnel := make(chan p2p.P2PMessage)
	p, _ := p2p.CreateP2PNetwork(tunnel, *port)
	defer close(tunnel)

	//2)Start to listen incoming connection
	if err := p.Listen(); err != nil {
		log.Fatal(err)
	}
	id := p.GetId()

	//3)Dial to peers to build peerClient
	if *address != "" {
		p.CreatePeer(*address, nil)

		results := p.FindNode(p.GetId(), dht.BucketSize, 8)
		for _, result := range results {
			p.GetRoutingTable().Update(result)
			fmt.Println(p.GetId().Address, "Update peer: ", result.Address)
		}

		//fmt.Println("=======routing table of", p.GetId().Address, "========")
		//peers := p.GetRoutingTable().GetPeers()
		//for _, peer := range peers {
		//	fmt.Println("Peer Address: ", peer.Address, " Peer id: ", peer.Id)
		//}
		//fmt.Println("===end of routing table of", p.GetId().Address, "===")
	}

	if *target {
		file, err := os.OpenFile("id", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		length, err := file.Write(id.Id)
		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}

		fmt.Println("write id to file, write length: ", length)
		file.Close()
	}

	if *find {
		fileRead, err := os.OpenFile("id", os.O_RDONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		idToFound := make([]byte, 32)
		jobCounter := 0
		succCounter := 0
		for {
			_, err := fileRead.Read(idToFound)
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Println(p.GetId().Address, "is finding peer ", idToFound)
			jobCounter++

			results := p.FindNodeById(idToFound)
			for _, result := range results {
				if bytes.Equal(idToFound, result.Id) {
					fmt.Println(p.GetId().Address, "Found node:", result.Id, "at", result.Address)
					succCounter++
					break
				}
			}
			fmt.Println(p.GetId().Address, "total find nodes:", succCounter, "/", jobCounter)
		}
		fileRead.Close()
	}

	<- finish
}
