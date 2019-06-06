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

var (
	urls           = []string{"ws://18.236.117.126:8546", "ws://18.237.179.193:8546", "ws://18.236.80.51:8546"}
	proxyAddr      = "0xE41546987bf5d4737961291C3eA7F3eDd3e2fbC7"
	credentialPath = "/Users/chenhaonien/go/src/github.com/DOSNetwork/core/testAccounts/bootCredential/fundKey/"
	passphrase     = "123"
)

func TestGetPendingNonce(t *testing.T) {
	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, "", urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an Error : %s.", err.Error())
		return
	}
	d := time.Now().Add(1 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), d)
	defer cancelFunc()

	r, e := adaptor.PendingNonce(ctx)
	if e != nil {
		t.Errorf("TestGetPendingNonce Failed , got an error: %s.", err.Error())
	}
	fmt.Println(r)
}

func TestLastUpdatedBlock(t *testing.T) {
	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, "", urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an Error : %s.", err.Error())
		return
	}

	val, e := adaptor.LastUpdatedBlock(context.Background())
	if e != nil {
		t.Errorf("LastUpdatedBlock Failed , got an error: %s.", e.Error())
	}
	fmt.Println("LastUpdatedBlock ", val)

	val, e = adaptor.GetWorkingGroupSize(context.Background())
	if e != nil {
		t.Errorf("GetWorkingGroupSize Failed , got an error: %s.", e.Error())
	}
	fmt.Println("GetWorkingGroupSize ", val)

	rand, e := adaptor.LastRandomness(context.Background())
	if e != nil {
		t.Errorf("LastRandomness() Failed , got an error: %s.", e.Error())
	}
	fmt.Println("LastRandomness()	 ", rand)
}

func TestConcurrentSend(t *testing.T) {
	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, "", urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an error : %s.", err.Error())
		return
	}

	sink, errc := adaptor.SubscribeEvent(SubscribeDosproxyUpdateGroupToPick)

	ctx := context.Background()
	for i := 3; i < 8; i++ {
		go func(i int) {
			err = adaptor.SetGroupToPick(ctx, uint64(i))
			fmt.Println("TestConcurrentSend err ", err)
		}(i)
	}
	result := 0
L:
	for {
		select {
		case event, ok := <-sink:
			if ok {
				switch content := event.(type) {
				case *LogUpdateGroupToPick:
					fmt.Println("DOSProxyUpdateGroupToPick ", int(content.NewNum.Uint64()), content.Removed)
					if content.Removed != true {
						result = result + int(content.NewNum.Uint64())
						if result == 25 {
							break L
						}
					}
				}
			}
		case e, ok := <-errc:
			if ok {
				err = e
				fmt.Println("TestConcurrentSend event err ", err)
			}
		}
	}
}
func TestSetErrorHandling(t *testing.T) {
	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, "", urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an error : %s.", err.Error())
		return
	}

	sink, errc := adaptor.SubscribeEvent(SubscribeDosproxyUpdateGroupToPick)

	ctx := context.Background()
	for i := 3; i < 8; i++ {
		err = adaptor.SetGroupToPick(ctx, uint64(i))
		fmt.Println("TestConcurrentSend err ", err)
		fmt.Println("Stop geth client to test")
		time.Sleep(3 * time.Second)
	}
	result := 0
L:
	for {
		select {
		case event, ok := <-sink:
			if ok {
				switch content := event.(type) {
				case *LogUpdateGroupToPick:
					fmt.Println("DOSProxyUpdateGroupToPick ", int(content.NewNum.Uint64()), content.Removed)
					if content.Removed != true {
						result = result + int(content.NewNum.Uint64())
						if result == 25 {
							break L
						}
					}
				}
			}
		case e, ok := <-errc:
			if ok {
				err = e
				fmt.Println("TestConcurrentSend event err ", err)
			}
		}
	}
}

func TestReconnect(t *testing.T) {
S:
	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, "", urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an error : %s.", err.Error())
		return
	}

	var errcList []chan error
	sink, errc := adaptor.SubscribeEvent(SubscribeDosproxyUpdateGroupToPick)
	errcList = append(errcList, errc)

	ctx := context.Background()
	for i := 3; i < 8; i++ {
		//go func(i int) {
		err = adaptor.SetGroupToPick(ctx, uint64(i))
		time.Sleep(2 * time.Second)
		fmt.Println(" err ", err)
		//}(i)
	}
	errc = errcList[0]
	result := 0
L:
	for {
		select {
		case event, ok := <-sink:
			if ok {
				switch content := event.(type) {
				case *LogUpdateGroupToPick:
					fmt.Println("DOSProxyUpdateGroupToPick ", int(content.NewNum.Uint64()), content.Removed)
					if content.Removed != true {
						result = result + int(content.NewNum.Uint64())
						if result == 25 {
							time.Sleep(15 * time.Second)
							break L
						}
					}
				}
			} else {
				fmt.Println("sink event !ok err ", err)
				break L
			}
		case e, ok := <-errc:
			if ok {
				err = e
				fmt.Println("errc event err ", err)
			} else {
				fmt.Println("errc event !ok err ", err)
				break L
			}
		}
	}
	time.Sleep(5 * time.Second)
	goto S
}

/*
func TestSendRequest(t *testing.T) {
	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, "", urls)
	if err != nil {
		t.Errorf("ReqSetInt Failed, got an error : %s.", err.Error())
		return
	}

	var errcList []<-chan error
	sink, errc := adaptor.PollLogs(SubscribeDosproxyUpdateGroupToPick, 0, 0)
	errcList = append(errcList, errc)

	fmt.Println("!!!!!!!!stop geth")
	time.Sleep(5 * time.Second)
	ctx := context.Background()
	count := 13
	for i := 8; i < count; i++ {
		_ = adaptor.SetGroupToPick(ctx, uint64(i))
		time.Sleep(1 * time.Second)
	}
	fmt.Println("!!!!!!!!restore geth")

	errc = MergeErrors(ctx, errcList...)
	result := 0
L:
	for {
		select {
		case event := <-sink:
			switch content := event.(type) {
			case *LogUpdateGroupToPick:
				fmt.Println("DOSProxyUpdateGroupToPick ", int(content.NewNum.Uint64()), content.Removed, result)
				if content.Removed != true {
					result = result + int(content.NewNum.Uint64())
					if result == 50 {
						break L
					}
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
