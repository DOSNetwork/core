package onchain

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
	//"github.com/DOSNetwork/core/share/vss/pedersen"
)

const (
	ONEFULLNODE = iota // 0
	TWOFULLNODES_1
	TWOFULLNODES_2
)

func mergeErrors(ctx context.Context, cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	// We must ensure that the output channel has the capacity to
	// hold as many errors
	// as there are error channels.
	// This will ensure that it never blocks, even
	// if WaitForPipeline returns early.
	out := make(chan error, len(cs))
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls
	// wg.Done.
	output := func(c <-chan error) {
		for n := range c {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	// Start a goroutine to close out once all the output goroutines
	// are done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func testProxy(setting int) (err error) {
	/*
		proxyAddr := "0x25d1b6aD154C84a998a836663c867853A916d4bD"
		urls := []string{"ws://51.15.0.157:8546", "ws://51.159.4.51:8546", "wss://rinkeby.infura.io/ws/8e609c76fce442f8a1735fbea9999747"}
		if setting == TWOFULLNODES_1 {
			//urls = []string{"wss://rinkeby.infura.io/ws/8e609c76fce442f8a1735fbea9999747"}
			urls = []string{"ws://51.15.0.157:8546"}
		} else if setting == TWOFULLNODES_2 {
			urls = []string{"ws://51.159.4.51:8546"}
		}
	*/
	proxyAddr := "0x4562daE9d912638413a5B290D657e28724E28e06"
	urls := []string{"ws://163.172.36.173:8546"}

	credentialPath := ""
	passphrase := ""
	adaptor, err := NewETHProxySession(credentialPath, passphrase, proxyAddr, urls)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if adaptor.s == nil {
		fmt.Println("adaptor is  nil ")
	}

	sink := make(chan interface{}, 1)
	if setting == TWOFULLNODES_1 {
		_ = adaptor.SubscribeEvent(SubscribeDOSProxyLogUpdateRandom, sink)
	} else if setting == TWOFULLNODES_2 {
		_ = adaptor.SubscribeEvent(SubscribeDOSProxyLogInsufficientGroupNumber, sink)
	} else {
		_ = adaptor.SubscribeEvent(SubscribeDOSProxyLogUpdateRandom, sink)
		_ = adaptor.SubscribeEvent(SubscribeDOSProxyLogInsufficientGroupNumber, sink)
		_ = adaptor.SubscribeEvent(SubscribeDOSProxyLogValidationResult, sink)
		//_ = adaptor.PollLogs(SubscribeDOSProxyLogInsufficientGroupNumber, sink)
	}

	ctx, _ := context.WithCancel(context.Background())
	var errcList []<-chan error
	//sigC := make(chan *vss.Signature)
	for i := 1; i <= 3; i++ {

		if setting == TWOFULLNODES_1 {
			//err := adaptor.FireRandom(ctx)
			err := adaptor.Grouping(ctx, i+3)
			errcList = append(errcList, err)
		} else if setting == TWOFULLNODES_2 {
			//err := adaptor.FireRandom(ctx)
			err := adaptor.Grouping(ctx, i)
			errcList = append(errcList, err)
		} else {
			err := adaptor.Grouping(ctx, i)
			errcList = append(errcList, err)
			/*
				err := adaptor.SetRandomNum(ctx, sigC)
				errcList = append(errcList, err)
				b := []byte{'g', 'o', 'l', 'a', 'n', 'g', 'g', 'o', 'l', 'a',
					'g', 'o', 'l', 'a', 'n', 'g', 'g', 'o', 'l', 'a',
					'g', 'o', 'l', 'a', 'n', 'g', 'g', 'o', 'l', 'a',
					'g', 'o', 'l', 'a', 'n', 'g', 'g', 'o', 'l', 'a',
					'g', 'o', 'l', 'a', 'n', 'g', 'g', 'o', 'l', 'a',
					'g', 'o', 'l', 'a', 'n', 'g', 'g', 'o', 'l', 'a',
					'g', 'o', 'l', 'a'}
				sig := &vss.Signature{Index: 1, QueryId: "test", Content: b, Signature: b}
				sigC <- sig
			*/
		}

	}
	errc := mergeErrors(ctx, errcList...)

	//count := 0
L:
	for {
		select {
		case event := <-sink:
			switch content := event.(type) {
			case *DOSProxyLogInsufficientGroupNumber:
				fmt.Println(setting, " Event DOSProxyLogInsufficientGroupNumber")
				//count++

			case *DOSProxyLogUpdateRandom:
				if !content.Removed {
					fmt.Println(setting, " Event DOSProxyLogUpdateRandom last: ", fmt.Sprintf("%x", content.LastRandomness))
					fmt.Println(setting, " Event DOSProxyLogUpdateRandom tx : ", content.Tx)
					//count++

				}
			case *DOSProxyLogValidationResult:
				if !content.Removed {
					fmt.Println(setting, " Event DOSProxyLogValidationResult ", content.Pass)
					//count++
				}
			}
		case _, ok := <-errc:
			if !ok {
				fmt.Println("errc done")
				adaptor.End()
				break L
			}
		}
	}
	close(sink)
	return
}

func TestConcurrentSend(t *testing.T) {
	testProxy(ONEFULLNODE)
	err := testProxy(ONEFULLNODE)
	for err != nil {
		time.Sleep(3 * time.Second)
		err = testProxy(ONEFULLNODE)
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
