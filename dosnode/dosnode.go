package dosnode

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
)

const (
	watchdogInterval = 10 //In minutes
	envPassPhrase    = "PASSPHRASE"
)

type ctxKey string

// DosNode is a strcut that represents a offchain dos client
type DosNode struct {
	ctx          context.Context
	cancel       context.CancelFunc
	suite        suites.Suite
	chain        onchain.ProxyAdapter
	dkg          dkg.PDKGInterface
	p            p2p.P2PInterface
	bootStrapIPs []string
	done         chan interface{}
	reqSignc     chan request
	cRequestDone chan [4]*big.Int
	onchainEvent chan interface{}
	id           []byte
	logger       log.Logger
	isGuardian   bool

	//For REST API
	startTime         time.Time
	state             string
	totalQuery        int
	fulfilledQuery    int
	numOfworkingGroup int
}

type request struct {
	ctx       context.Context
	requestID string
	threshold int
	reply     chan *vss.Signature
}

type crDurations struct {
	cid        *big.Int
	startBlock *big.Int
	commitDur  *big.Int
	revealDur  *big.Int
	sec        *big.Int
}

//NewDosNode creates a DosNode struct
func NewDosNode(key *keystore.Key, config configuration.Config) (dosNode *DosNode, err error) {
	id := key.Address
	l := log.New("module", "dosclient")

	//Set up an onchain adapter
	chainConn, err := onchain.NewProxyAdapter(config.ChainType, key, config.DOSAddressBridgeAddress, config.ChainNodePool)
	if err != nil {
		if err.Error() != "No any working eth client for event tracking" {
			fmt.Println("NewDosNode failed ", err)
			return
		}
	}
	//Build a p2p network
	p, err := p2p.CreateP2PNetwork(id.Bytes(), config.NodeIP, config.NodePort, p2p.GossipDiscover)
	if err != nil {
		fmt.Println("CreateP2PNetwork err ", err)
		return
	}

	//Build a p2pDKG
	suite := suites.MustFind("bn256")
	p2pDkg := dkg.NewPDKG(p, suite)

	ctx, cancel := context.WithCancel(context.Background())
	dosNode = &DosNode{
		ctx:               ctx,
		cancel:            cancel,
		suite:             suite,
		p:                 p,
		bootStrapIPs:      config.BootStrapIPs,
		chain:             chainConn,
		dkg:               p2pDkg,
		done:              make(chan interface{}),
		reqSignc:          make(chan request, 21),
		cRequestDone:      make(chan [4]*big.Int),
		id:                id.Bytes(),
		logger:            l,
		startTime:         time.Now(),
		state:             "Init Done",
		totalQuery:        0,
		fulfilledQuery:    0,
		numOfworkingGroup: 0,
	}

	return dosNode, nil
}

//End is an operation that does a graceful shutdown
func (d *DosNode) End() {
	d.cancel()
	<-d.ctx.Done()
	fmt.Println("End")
}

func (d *DosNode) Start() {
	defer fmt.Println("End Start")
	d.state = "Working"
	d.startRESTServer()
	go d.chain.ReqLoop()
	go d.onchainLoop()

	go func() {
		if err := d.p.Listen(); err != nil {
			fmt.Println("Listen() err ", err)
			return
		}
	}()
	go d.queryLoop()
	go d.dkg.Loop()
	//Bootstrapping p2p network

	fmt.Println("Join :", d.bootStrapIPs)
	retry, num := 0, 0
	var err error
	for {
		num, err = d.p.Join(d.bootStrapIPs)
		if err != nil || num == 0 {
			fmt.Println("Join ", err, num)

			if retry == 10 {
				return
			}
			retry++
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	fmt.Println("Join : num of peer ", num)
	peerChecker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-d.ctx.Done():
		case <-peerChecker.C:
			if num != d.p.NumOfMembers() {
				num = d.p.NumOfMembers()
				d.logger.Event("peersUpdate", map[string]interface{}{"numOfPeers": num})
			}
		}
	}

}
