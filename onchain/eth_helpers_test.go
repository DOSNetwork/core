package onchain

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const (
	ID = 4
)

func TestSetEthKey(t *testing.T) {
	credentialPath := "/Users/chenhaonien/go/src/github.com/DOSNetwork/core/credential"
	passphrase := "123"
	_, err := SetEthKey(credentialPath, passphrase)
	if err != nil {
		t.Errorf("SetEthKey Error: %s.", err.Error())
	}
}
func TestDialToEth(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	urls := []string{"ws://51.15.0.157:8546", "ws://51.159.4.51:8546", "wss://rinkeby.infura.io/ws/8e609c76fce442f8a1735fbea9999747"}
	clients, _ := DialToEth(ctx, urls)
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

func TestDialToEthDeadline(t *testing.T) {
	d := time.Now().Add(1 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	defer cancelFunc()
	urls := []string{"ws://163.172.36.173:8546", "ws://51.15.0.157:8546", "ws://51.159.4.51:8546", "wss://rinkeby.infura.io/ws/8e609c76fce442f8a1735fbea9999747"}
	clients, _ := DialToEth(ctx, urls)
	time.Sleep(2 * time.Second)

	for _ = range clients {
		t.Errorf("Should not receive any client")
	}
}

func TestDialToEthErrHandling(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	urls := []string{"ws://163.172.36.173:8545", "ws://51.15.0.157:8546"}
	clients, errs := DialToEth(ctx, urls)
	time.Sleep(2 * time.Second)

	go func() {
		for e := range errs {
			if e.Error() != "bad status" {
				t.Errorf("Err incorrect, got: %s, want: %s.", e.Error(), "bad status")
			}
		}
	}()

	count := 0
	for _ = range clients {
		count++
	}
	if count > 1 {
		t.Errorf("Dial success count, got: %d, want: %d.", count, 1)
	}
}
