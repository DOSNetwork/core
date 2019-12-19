package onchain

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/DOSNetwork/core/onchain/commitreveal"
	"github.com/DOSNetwork/core/onchain/dosbridge"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	errors "golang.org/x/xerrors"
)

const (
	//TrafficSystemRandom is a request type to build a corresponding pipeline
	TrafficSystemRandom = iota // 0
	//TrafficUserRandom is a request type to build a corresponding pipeline
	TrafficUserRandom
	//TrafficUserQuery is a request type to build a corresponding pipeline
	TrafficUserQuery
)

type logger interface {
	Info(msg string)
	Error(err error)
	TimeTrack(start time.Time, e string, info map[string]interface{})
	Event(e string, info map[string]interface{})
}

type ethAdaptor struct {
	logger
	bridgeAddr       common.Address
	proxyAddr        common.Address
	commitRevealAddr common.Address
	bootStrapUrl     string
	httpUrls         []string
	wsUrls           []string
	key              *keystore.Key
	gasPrice         int64
	gasLimit         int64
	connTimeout      time.Duration
	getTimeout       time.Duration
	setTimeout       time.Duration
	proxies          []*dosproxy.DosproxySession
	crs              []*commitreveal.CommitrevealSession
	clients          []*ethclient.Client
	ctxes            []context.Context
	cancels          []context.CancelFunc
	ctx              context.Context
	cancelFunc       context.CancelFunc
	reqQueue         chan *request
}

//NewEthAdaptor creates an eth implemention of ProxyAdapter
func NewEthAdaptor(key *keystore.Key, bridgeAddr string, l logger) (adaptor *ethAdaptor, err error) {
	if key == nil {
		return nil, errors.New("no keystore")
	}
	if !common.IsHexAddress(bridgeAddr) {
		return nil, errors.New("not a valid hex address")
	}
	adaptor = &ethAdaptor{}
	adaptor.bridgeAddr = common.HexToAddress(bridgeAddr)
	adaptor.key = key
	adaptor.logger = l
	adaptor.reqQueue = make(chan *request)
	// Use SuggestGasPrice and EstimateGas instead of hard coding
	adaptor.gasPrice = 3000000000
	adaptor.gasLimit = 6000000
	adaptor.connTimeout = 60 * time.Second
	adaptor.getTimeout = 60 * time.Second
	adaptor.setTimeout = 60 * time.Second
	return
}

//End close the connection to eth and release all resources
func (e *ethAdaptor) DisconnectAll() {
	e.cancelFunc()
	for {
		if !e.isConnecting() {
			return
		}
	}
	return
}

func (e *ethAdaptor) Disconnect(idx int) {
	if e.cancels[idx] != nil {
		e.cancels[idx]()
	}
	return
}

func (e *ethAdaptor) isConnecting() (result bool) {
	if e.ctx == nil {
		return
	}
	select {
	case <-e.ctx.Done():
		return
	default:
	}
	for _, ctx := range e.ctxes {
		select {
		case <-ctx.Done():
		default:
			result = true
			return
		}
	}
	return
}

func (e *ethAdaptor) Connect(urls []string, t time.Time) (err error) {
	if e.isConnecting() {
		err = errors.Errorf("connection existing")
		return
	}

	var httpUrls []string
	var wsUrls []string
	for _, url := range urls {
		if strings.Contains(url, "http") {
			httpUrls = append(httpUrls, url)
		} else if strings.Contains(url, "ws") {
			wsUrls = append(wsUrls, url)
		}
	}
	e.wsUrls = wsUrls
	if len(wsUrls) == 0 {
		return errors.New("no web socket URL")
	}

	e.ctx, e.cancelFunc = context.WithCancel(context.Background())
	e.clients = nil
	e.proxies = nil
	e.crs = nil
	e.ctxes = nil
	e.cancels = nil

	dialCtx, dialCancel := context.WithDeadline(e.ctx, t)
	defer dialCancel()

	results := DialToEth(dialCtx, e.wsUrls)
	for result := range results {
		if result.Err != nil {
			e.logger.Error(errors.Errorf(": %w", result.Err))
			continue
		}
		client := result.Client
		bridge, err := dosbridge.NewDosbridge(e.bridgeAddr, client)
		if err != nil {
			e.logger.Error(errors.Errorf(": %w", err))
			client.Close()
			continue
		}
		proxyAddr, err := bridge.GetProxyAddress(&bind.CallOpts{Context: dialCtx})
		if err != nil {
			e.logger.Error(errors.Errorf(": %w", err))
			client.Close()
			continue
		}
		commitRevealAddr, err := bridge.GetCommitRevealAddress(&bind.CallOpts{Context: dialCtx})
		if err != nil {
			e.logger.Error(errors.Errorf(": %w", err))
			client.Close()
			continue
		}

		bootStrapUrl, err := bridge.GetBootStrapUrl(&bind.CallOpts{Context: dialCtx})
		if err != nil {
			e.logger.Error(errors.Errorf(": %w", err))
			client.Close()
			continue
		}

		p, err := dosproxy.NewDosproxy(proxyAddr, client)
		if err != nil {
			err = errors.Errorf("NewDosproxy failed: %w", err)
			e.logger.Error(err)
			client.Close()
			continue
		}

		cr, err := commitreveal.NewCommitreveal(commitRevealAddr, client)
		if err != nil {
			err = errors.Errorf("NewCommitreveal failed : %w", err)
			e.logger.Error(err)
			client.Close()
			continue
		}
		e.logger.Event("ConnToOnchain", map[string]interface{}{"OnchainURL": result.Url})

		ctx, cancel := context.WithCancel(e.ctx)
		ctx = context.WithValue(ctx, "index", len(e.ctxes))
		auth := bind.NewKeyedTransactor(e.key.PrivateKey)
		auth.GasPrice = big.NewInt(e.gasPrice) //1 Gwei
		auth.GasLimit = uint64(e.gasLimit)
		auth.Context = ctx

		if bootStrapUrl != "" {
			e.bootStrapUrl = bootStrapUrl
		}
		e.clients = append(e.clients, client)
		e.proxies = append(e.proxies, &dosproxy.DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: ctx}, TransactOpts: *auth})
		e.crs = append(e.crs, &commitreveal.CommitrevealSession{Contract: cr, CallOpts: bind.CallOpts{Context: ctx}, TransactOpts: *auth})
		e.ctxes = append(e.ctxes, ctx)
		e.cancels = append(e.cancels, cancel)
	}

	if len(e.proxies) == 0 {
		err = errors.New("no proxy instance")
		e.logger.Error(err)
		return
	}
	if len(e.crs) == 0 {
		err = errors.New("no commir reveal instance")
		e.logger.Error(err)
		return
	}
	go e.ReqLoop()
	return
}
func (e *ethAdaptor) BootStrapUrl() string {
	return e.bootStrapUrl
}

// Address gets the string representation of the underlying address.
func (e *ethAdaptor) Address() (addr common.Address) {
	return e.key.Address
}
