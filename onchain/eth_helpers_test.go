package onchain

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ID = 4
)

func TestReadEthKey(t *testing.T) {
	_, err := ReadEthKey(credentialPath, passphrase)
	if err != nil {
		t.Errorf("ReadEthKey Error: %s.", err.Error())
	}
}

func TestDialToEth(t *testing.T) {
	urls = []string{"http://123.123.123.123", "http://18.236.117.126:8545"}
	var clients []*ethclient.Client
	ctx := context.Background()
	outc := DialToEth(ctx, urls)
	for client := range outc {
		clients = append(clients, client)
	}
	if len(clients) != len(urls) {
		t.Errorf("Dial success count, got: %d, want: %d.", len(clients), len(urls)-1)
	}

	//Check to see if client is really working by getting networkID
	count := 0
	for _, client := range clients {
		id, err := client.NetworkID(ctx)
		//i/o timeout
		if err != nil {
			fmt.Println("NetworkID err ", err)
			client.Close()
		} else {
			if ID != id.Uint64() {
				t.Errorf("ID incorrect, got: %d, want: %d.", id.Uint64(), ID)
			}
			count++
		}
	}
	if count != len(urls)-1 {
		t.Errorf("Dial count, got: %d, want: %d.", count, len(urls)-1)
	}
	fmt.Println("TestDialToEth pass")
}

func TestDialWS(t *testing.T) {
	urls = []string{"ws://18.236.117.126:8546", "ws:18.236.115.126:8546"}
	var clients []*ethclient.Client
	ctx := context.Background()
	outc := DialToEth(ctx, urls)
	for client := range outc {
		clients = append(clients, client)
	}
	if len(clients) != len(urls)-1 {
		t.Errorf("Dial success count, got: %d, want: %d.", len(clients), len(urls)-1)
	}

	fmt.Println("TestDialToEth pass")
}

func TestDialToEthDeadline(t *testing.T) {
	urls = []string{"ws://18.236.117.126:8546", "ws:18.236.115.126:8546"}
	d := time.Now().Add(1 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	defer cancelFunc()
	clients := DialToEth(ctx, urls)
	time.Sleep(2 * time.Second)

	for range clients {
		t.Errorf("Should not receive any client")
	}
}

func TestDialToEthErrHandling(t *testing.T) {
	var tUrls []string
	//It doesn't cause any error when dialing
	tUrls = append(tUrls, "http://123.123.123.123")
	//It cause an error (dial unix: missing address)
	tUrls = append(tUrls, "")
	tUrls = append(tUrls, "ws://123.123.123.123:8546")
	//tUrls = append(tUrls, "ws://51.15.0.157:8546")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	//ctx := context.Background()

	clients := DialToEth(ctx, tUrls)

	count := 0
	for range clients {
		count++
	}
	if count != 0 {
		t.Errorf("Dial success count, got: %d, want: %d.", count, 1)
	}
}
