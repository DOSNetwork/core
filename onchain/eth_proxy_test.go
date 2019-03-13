package onchain

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const (
	ONEFULLNODE = iota // 0
	TWOFULLNODES_1
	TWOFULLNODES_2
)

func TestGetPendingNonce(t *testing.T) {
	urls := []string{""}
	proxyAddr := ""
	credentialPath := ""
	passphrase := ""

	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an Error : %s.", err.Error())
		return
	}

	_, e := adaptor.PendingNonce()
	if e != nil {
		t.Errorf("TestGetPendingNonce Failed , got an error: %s.", err.Error())
	}
}

func TestConcurrentSend(t *testing.T) {
	urls := []string{""}

	proxyAddr := ""
	credentialPath := ""
	passphrase := ""

	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an error : %s.", err.Error())
		return
	}

	var errcList []<-chan error
	sink, errc := adaptor.SubscribeEvent(SubscribeDOSProxyTestEvent)
	errcList = append(errcList, errc)

	ctx, _ := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(i int) {
			_ = adaptor.TestContract(ctx, uint64(i))
		}(i)
	}
	errc = mergeErrors(ctx, errcList...)
	result := 0
L:
	for {
		select {
		case event := <-sink:
			switch content := event.(type) {
			case *DOSProxyTestEvent:
				result = result + int(content.Parameter.Uint64())
				if result == 15 {
					break L
				}
			}
		case e, ok := <-errc:
			if ok {
				err = e
				fmt.Println(err)
			}
		}
	}
}

func TestSendRequest(t *testing.T) {
	urls := []string{""}

	proxyAddr := ""
	credentialPath := ""
	passphrase := ""

	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an error : %s.", err.Error())
		return
	}

	var errcList []<-chan error
	sink, errc := adaptor.SubscribeEvent(SubscribeDOSProxyTestEvent)
	errcList = append(errcList, errc)

	fmt.Println("!!!!!!!!stop geth")
	time.Sleep(5 * time.Second)
	ctx, _ := context.WithCancel(context.Background())
	count := 5
	for i := 0; i < count; i++ {
		_ = adaptor.TestContract(ctx, uint64(i))
		time.Sleep(1 * time.Second)
	}
	fmt.Println("!!!!!!!!restore geth")

	errc = mergeErrors(ctx, errcList...)
	result := 0
L:
	for {
		select {
		case event := <-sink:
			switch content := event.(type) {
			case *DOSProxyTestEvent:
				result = result + int(content.Parameter.Uint64())
				if result == 15 {
					break L
				}
			}
		case e, ok := <-errc:
			if ok {
				err = e
				fmt.Println(err)
				//t.Errorf("TestConcurrentSend Failed, got an error : %s.", err.Error())
				//break L
			}
		}
	}
}

/*
func TestNonceConflict(t *testing.T) {

	for i := 1; i < 2; i++ {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			testProxy(TWOFULLNODES_1)
		}()
		go func() {
			defer wg.Done()
			testProxy(TWOFULLNODES_2)
		}()

		wg.Wait()
	}
}
*/
