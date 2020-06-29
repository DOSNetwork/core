package dosnode

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
	errors "golang.org/x/xerrors"
)

const (
	envPassPhrase = "PASSPHRASE"
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
	isAdmin      bool
	config       *configuration.Config
	m            sync.Mutex
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
func NewDosNode(key *keystore.Key, config *configuration.Config) (dosNode *DosNode, err error) {
	id := key.Address
	l := log.New("module", "dosclient")

	//Set up an onchain adapter
	chainConn, err := onchain.NewProxyAdapter(key, config)
	if err != nil {
		if err.Error() != "No any working eth client for event tracking" {
			l.Error(err)
			return
		}
	}
	//Build a p2p network
	p, err := p2p.CreateP2PNetwork(id.Bytes(), config.NodeIP, config.NodePort, p2p.GossipDiscover)
	if err != nil {
		l.Error(err)
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
		config:            config,
	}

	return dosNode, nil
}

//End is an operation that does a graceful shutdown
func (d *DosNode) End() {
	d.m.Lock()
	defer d.m.Unlock()
	select {
	case <-d.ctx.Done():
	default:
		d.chain.UnRegisterNode()
		d.p.Leave()
		d.cancel()
	}
}

func (d *DosNode) Start() {
	d.logger.Event("peersUpdate", map[string]interface{}{"numOfPeers": 0})
	d.state = "Working"
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		defer wg.Done()
		d.startRESTServer()
		d.logger.Info("[DOS] End RESTServer")
		d.End()
	}()
	t := time.Now().Add(60 * time.Second)
	if err := d.chain.Connect(d.config.ChainNodePool, t); err != nil {
		d.logger.Error(err)
		return
	}
	go func() {
		defer wg.Done()
		d.onchainLoop()
		d.logger.Info("[DOS] End ONCHAIN Loop")
		d.End()
	}()
	go func() {
		defer wg.Done()
		d.p.Listen()
		d.logger.Info("[DOS] End P2P Loop")
		d.End()
	}()
	go func() {
		defer wg.Done()
		d.queryLoop()
		d.logger.Info("[DOS] End Query Loop")
		d.End()
	}()
	go func() {
		defer wg.Done()
		d.dkg.Loop()
		d.logger.Info("[DOS] End DKG Loop")
		d.End()
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-d.ctx.Done():
				d.logger.Info("[DOS] ctx.Done")
				return
			default:
				if d.isGuardian {
					currentBlockNumber, err := d.chain.CurrentBlock()
					if err != nil {
						d.logger.Error(err)
						continue
					}
					handled := d.handleRandom(currentBlockNumber)
					if !handled {
						handled = d.handleGroupFormation()
					}
					if !handled {
						handled = d.handleGroupDissolve(currentBlockNumber)
					}
					if !handled {
						handled = d.handleBootstrap(currentBlockNumber)
					}
				}
				time.Sleep(15 * time.Second)
			}
		}
	}()

	retry, num := 0, 0
	var err error
	for {
		//Bootstrap p2p network
		if d.chain.BootStrapUrl() != "" {
			d.logger.Info(fmt.Sprintf("BootStrapUrl : %s", d.chain.BootStrapUrl()))

			ips := getBootIps(d.chain.BootStrapUrl())
			if len(ips) != 0 {
				d.bootStrapIPs = append(d.bootStrapIPs, ips...)
				d.bootStrapIPs = unique(d.bootStrapIPs)
			}
		}

		d.logger.Info("Bootstraping ...")
		num, err = d.p.Join(d.bootStrapIPs)
		if err != nil || num == 0 {
			if err != nil {
				d.logger.Error(errors.Errorf(" : %w", err))
			}
			if retry == 10 {
				d.logger.Error(errors.New("Can't join p2p network"))
				return
			}
			retry++
			time.Sleep(5 * time.Second)
		} else {
			d.config.BootStrapIPs = d.bootStrapIPs
			d.config.UpdateConfig()
			d.logger.Event("peersUpdate", map[string]interface{}{"numOfPeers": num})
			break
		}
	}
	d.logger.Info(fmt.Sprintf("[DOS] Join : num of peer %d\n", num))
	wg.Wait()
	d.logger.Info("[DOS] End")
}

func getBootIps(bootStrapUrl string) []string {
	req, err := http.NewRequest("GET", bootStrapUrl, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)

	str := string(r)
	strlist := strings.Split(str, ",")
	nodeIPs := make([]string, len(strlist)-1)
	for i := 0; i < len(strlist)-1; i++ {
		nodeIPs[i] = strlist[i]
	}
	return nodeIPs
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
