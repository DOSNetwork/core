package eth

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/DOSNetwork/core/configuration"
)

var userTestAdaptor *EthUserAdaptor
var wg sync.WaitGroup

func TestMain(m *testing.M) {
	os.Setenv("CONFIGPATH", "../../..")
	config := configuration.ReadConfig()
	chainConfig := configuration.GetOnChainConfig(config)
	userTestAdaptor = &EthUserAdaptor{}
	err := userTestAdaptor.Init(true, &chainConfig)
	if err != nil {
		fmt.Println(err)
	}
	events := make(chan interface{})
	userTestAdaptor.SubscribeToAll(events)
	go func() {
		for {
			select {
			case event := <-events:
				switch i := event.(type) {
				case *AskMeAnythingSetTimeout:
					fmt.Println("AskMeAnythingSetTimeout")
					fmt.Println("new timeout:", i.NewTimeout)
					fmt.Println("previous timeout:", i.PreviousTimeout)
					fmt.Println("____________________________________________")
				case *AskMeAnythingQueryResponseReady:
					fmt.Println("AskMeAnythingQueryResponseReady")
					fmt.Println("Callback Ready Query id:", i.QueryId)
					fmt.Println("result: ", i.Result)
					fmt.Println("initial new query...")
					fmt.Println("____________________________________________")
				case *AskMeAnythingRequestSent:
					fmt.Println("AskMeAnythingRequestSent")
					fmt.Println("succ:", i.Succ)
					fmt.Println("RequestId", i.RequestId)
					fmt.Println("____________________________________________")
				case *AskMeAnythingRandomReady:
					fmt.Println("AskMeAnythingRandomReady")
					fmt.Println("GeneratedRandom:", i.GeneratedRandom)
					fmt.Println("____________________________________________")
					wg.Done()
				default:
					fmt.Println("type mismatch")
				}
			}
		}
	}()
	m.Run()
	//os.Exit(m.Run())
	fmt.Println("TestMain TEAR DOWN!!!!!!!!!!!!!1")
}
func TestPipeCheckURL(test *testing.T) {
	fmt.Println("       TestPipeCheckURL")
}
func BenchmarkMathRand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		wg.Add(1)
		userTestAdaptor.GetSafeRandom()
		wg.Wait()
	}
}

/*
func TestPipeCheckURL(test *testing.T) {
	suite := suites.MustFind("bn256")
	nbParticipants := 1
	id := []byte{1, 2, 3}
	//1)Build a p2p network
	peerEvent := make(chan p2p.P2PMessage, 100)
	defer close(peerEvent)
	p, _ := p2p.CreateP2PNetwork(peerEvent, 0)
	p.SetId(id)
	p.Listen()

	peerEventForDKG := make(chan p2p.P2PMessage, 100)
	defer close(peerEventForDKG)
	p2pDkg, _ := dkg.CreateP2PDkg(p, suite, peerEventForDKG, nbParticipants)
	go p2pDkg.EventLoop()
	p2pDkg.RunDKG()

	d, _ := CreateDosNode(suite, nbParticipants, nil, nil, p2pDkg)
	chUrl := make(chan interface{}, 100)
	defer close(chUrl)
	out1 := d.PipeQueries(chUrl)
	QueryId := big.NewInt(1)
	url := "test"
	Timeout := big.NewInt(1)
	Randomness := big.NewInt(1)
	DispatchedGroup := [4]*big.Int{QueryId, QueryId, QueryId, QueryId}
	chUrl <- &onchain.DOSProxyLogUrl{
		QueryId:         QueryId,
		Url:             url,
		Timeout:         Timeout,
		Randomness:      Randomness,
		DispatchedGroup: DispatchedGroup,
	}
	fmt.Println(<-out1)
}
*/
