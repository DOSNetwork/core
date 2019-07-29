package dosnode

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
)

var (
	urls            = []string{"https://rinkeby.infura.io/v3/2a6901876ca54406960499e888e70439", "wss://rinkeby.infura.io/ws/v3/3a3e5d776961418e93a8b33fef2f6642"}
	proxyAddr       = "0x3b8Cb935bDdFAF59EFa11aFfDfc8760387624fa2"
	crAddr          = "0xE04B34A113BB707eCF8dc01D51f8A56213Bdcb81"
	credentialPath  = "/Users/chenhaonien/go/src/github.com/DOSNetwork/core/testing/testAccounts/1/credential/"
	credential2Path = "/Users/chenhaonien/go/src/github.com/DOSNetwork/core/testing/testAccounts/2/credential/"
	credential3Path = "/Users/chenhaonien/go/src/github.com/DOSNetwork/core/testing/testAccounts/3/credential/"

	passphrase = "123"
)

func buildComponents(port, credentialPath string, t *testing.T) (p2p.P2PInterface, onchain.ProxyAdapter) {
	key, err := onchain.ReadEthKey(credentialPath, passphrase)
	if err != nil {
		t.Errorf("TestChoseSubmitter Failed, got an error : %s.", err.Error())
	}
	adaptor, err := onchain.NewEthAdaptor(key, proxyAddr, crAddr, urls)
	if err != nil {
		t.Errorf("TestChoseSubmitter Failed, got an Error : %s.", err.Error())
	}
	adaptor.Start()
	nodeId := key.Address.Bytes()

	os.Setenv("PUBLICIP", "127.0.0.1")

	p, err := p2p.CreateP2PNetwork(nodeId, port, p2p.NoDiscover)
	if err != nil {
		t.Errorf("Error %s", err)
	}
	p.Listen()
	return p, adaptor
}

func TestChoseSubmitter(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	os.Setenv("LOGIP", "163.172.36.173:9500")
	log.Init([]byte{'a'})

	p1, e1 := buildComponents("9905", credentialPath, t)
	p2, e2 := buildComponents("9904", credential2Path, t)
	p3, _ := buildComponents("9904", credential3Path, t)

	var ids [][]byte
	ids = append(ids, p1.GetID())
	ids = append(ids, p2.GetID())
	ids = append(ids, p3.GetID())

	logger := log.New("module", "dostage")
	lastSysRand := big.NewInt(2)
	outCount := 1
	p1.SetPort("9904")
	_, _ = p1.ConnectTo(ctx, "127.0.0.1", nil)

	p2.SetPort("9905")
	_, _ = p2.ConnectTo(ctx, "127.0.0.1", nil)

	submitterc, errc := choseSubmitter(ctx, p1, e1, lastSysRand, ids, outCount, logger)
	select {
	case s := <-submitterc[0]:
		if string(s) == string(p3.GetID()) {
			t.Errorf("TestChoseSubmitter failed ")
		}
	case err := <-errc:
		fmt.Println("choseSubmitter err ", err)
	case <-ctx.Done():
		t.Errorf("Error %s", ctx.Err())
	}

	submitterc, errc = choseSubmitter(ctx, p2, e2, lastSysRand, ids, outCount, logger)
	select {
	case s := <-submitterc[0]:
		if string(s) == string(p3.GetID()) {
			t.Errorf("TestChoseSubmitter failed ")
		}
	case err := <-errc:
		fmt.Println("choseSubmitter err ", err)
	case <-ctx.Done():
		t.Errorf("Error %s", ctx.Err())
	}
}
