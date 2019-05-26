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
	urls           = []string{""}
	proxyAddr      = ""
	credentialPath = ""
	passphrase     = ""
)

func TestGetPendingNonce(t *testing.T) {
	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, "", urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an Error : %s.", err.Error())
		return
	}

	_, e := adaptor.PendingNonce()
	if e != nil {
		t.Errorf("TestGetPendingNonce Failed , got an error: %s.", err.Error())
	}
}

func TestLastUpdatedBlock(t *testing.T) {
	adaptor, err := NewEthAdaptor(credentialPath, passphrase, proxyAddr, "", urls)
	if err != nil {
		t.Errorf("TestConcurrentSend Failed, got an Error : %s.", err.Error())
		return
	}

	val, e := adaptor.LastUpdatedBlock()
	if e != nil {
		t.Errorf("LastUpdatedBlock Failed , got an error: %s.", e.Error())
	}
	fmt.Println("LastUpdatedBlock ", val)

	val, e = adaptor.GetWorkingGroupSize()
	if e != nil {
		t.Errorf("GetWorkingGroupSize Failed , got an error: %s.", e.Error())
	}
	fmt.Println("GetWorkingGroupSize ", val)

	rand, e := adaptor.LastRandomness()
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

	var errcList []<-chan error
	sink, errc := adaptor.SubscribeEvent(SubscribeDosproxyUpdateGroupToPick)
	errcList = append(errcList, errc)

	ctx := context.Background()

	for i := 3; i < 8; i++ {
		go func(i int) {
			_ = adaptor.SetGroupToPick(ctx, uint64(i))
		}(i)
	}
	errc = MergeErrors(ctx, errcList...)
	result := 0
L:
	for {
		select {
		case event := <-sink:
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
		case e, ok := <-errc:
			if ok {
				err = e
				fmt.Println(err)
			}
		}
	}
}

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
