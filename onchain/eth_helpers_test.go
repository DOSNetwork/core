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
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	clients := DialToEth(ctx, urls, nil)
	count := 0
	for client := range clients {
		id, err := client.NetworkID(ctx)
		if err != nil {
			t.Errorf("err %s.", err.Error())
		}
		if ID != id.Uint64() {
			t.Errorf("ID incorrect, got: %d, want: %d.", id.Uint64(), ID)
		}
		client.Close()
		count++
	}
	if count != len(urls) {
		t.Errorf("Dial success count, got: %d, want: %d.", count, len(urls))
	}
	fmt.Println("TestDialToEth pass")
}

func TestCheckSync(t *testing.T) {
	var mClient *ethclient.Client
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	clients := DialToEth(ctx, urls[:1], nil)
	mClient = <-clients
	id, err := mClient.NetworkID(ctx)
	if err != nil {
		t.Errorf("err %s.", err.Error())
	}
	if ID != id.Uint64() {
		t.Errorf("ID incorrect, got: %d, want: %d.", id.Uint64(), ID)
	}

	syncClients := CheckSync(context.Background(), mClient, DialToEth(context.Background(), urls[1:], nil))
	count := 0
	for client := range syncClients {
		id, err := client.NetworkID(ctx)
		if err != nil {
			t.Errorf("err %s.", err.Error())
		}
		if ID != id.Uint64() {
			t.Errorf("ID incorrect, got: %d, want: %d.", id.Uint64(), ID)
		}
		client.Close()
		count++
	}
	if count != len(urls)-1 {
		t.Errorf("Dial success count, got: %d, want: %d.", count, len(urls[1:]))
	}
	fmt.Println("TestDialToEth pass")
}

func TestDialToEthDeadline(t *testing.T) {
	d := time.Now().Add(1 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	defer cancelFunc()
	clients := DialToEth(ctx, urls, nil)
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
	key, err := ReadEthKey(credentialPath, passphrase)
	if err != nil {
		return
	}
	clients := DialToEth(ctx, tUrls, key)

	count := 0
	for range clients {
		count++
	}
	if count != 0 {
		t.Errorf("Dial success count, got: %d, want: %d.", count, 1)
	}
}
