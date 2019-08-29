package onchain

import (
	"context"
	"fmt"
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
	Error(err error)
	TimeTrack(start time.Time, e string, info map[string]interface{})
	Event(e string, info map[string]interface{})
}

type ethAdaptor struct {
	logger
	bridgeAddr       common.Address
	proxyAddr        common.Address
	commitRevealAddr common.Address
	httpUrls         []string
	wsUrls           []string
	key              *keystore.Key
	auth             *bind.TransactOpts

	proxies    []*dosproxy.DosproxySession
	crs        []*commitreveal.CommitrevealSession
	clients    []*ethclient.Client
	ctx        context.Context
	cancelFunc context.CancelFunc
	reqQueue   chan *request
}

//NewEthAdaptor creates an eth implemention of ProxyAdapter
func NewEthAdaptor(key *keystore.Key, bridgeAddr string, urls []string, l logger) (adaptor *ethAdaptor, err error) {
	var httpUrls []string
	var wsUrls []string
	for _, url := range urls {
		if strings.Contains(url, "http") {
			httpUrls = append(httpUrls, url)
		} else if strings.Contains(url, "ws") {
			wsUrls = append(wsUrls, url)
		}
	}
	fmt.Println("httpUrls ", httpUrls)
	fmt.Println("wsUrls ", wsUrls)
	if key == nil {
		return nil, errors.Errorf("NewEthAdaptor failed : %w", ErrNoKeystore)
	}
	if len(wsUrls) == 0 {
		return nil, errors.Errorf("NewEthAdaptor failed : %w", ErrNoWSURL)
	}
	if !common.IsHexAddress(bridgeAddr) {
		return nil, errors.Errorf("NewEthAdaptor failed : %w", ErrNotValidAddr)
	}
	adaptor = &ethAdaptor{}
	adaptor.httpUrls = httpUrls
	adaptor.wsUrls = wsUrls
	adaptor.bridgeAddr = common.HexToAddress(bridgeAddr)
	adaptor.key = key
	adaptor.logger = l
	adaptor.reqQueue = make(chan *request)
	//
	adaptor.ctx, adaptor.cancelFunc = context.WithCancel(context.Background())
	adaptor.auth = bind.NewKeyedTransactor(adaptor.key.PrivateKey)
	adaptor.auth.GasPrice = big.NewInt(20000000000) //1 Gwei
	adaptor.auth.GasLimit = uint64(6000000)
	adaptor.auth.Context = adaptor.ctx
	return
}

//End close the connection to eth and release all resources
func (e *ethAdaptor) Close() {
	e.cancelFunc()
	e.clients = nil
	e.proxies = nil
	e.crs = nil
	e.reqQueue = nil
	<-e.ctx.Done()
	return
}

func (e *ethAdaptor) isDone() (err error) {
	select {
	case <-e.ctx.Done():
		err = e.ctx.Err()
	default:
	}
	return
}

func (e *ethAdaptor) Connect(ctx context.Context) (err error) {
	if err = e.isDone(); err != nil {
		return
	}
	e.clients = nil
	e.proxies = nil
	e.crs = nil
	results := DialToEth(ctx, e.wsUrls)

	for result := range results {
		if result.err != nil {
			e.logger.Error(&OnchainError{err: result.err, t: time.Now()})
			continue
		}
		client := result.c

		bridge, err := dosbridge.NewDosbridge(e.bridgeAddr, client)
		if err != nil {
			err = errors.Errorf("NewDosbridge failed: %w", err)
			e.logger.Error(&OnchainError{err: err, t: time.Now()})
			client.Close()
			continue
		}
		proxyAddr, err := bridge.GetProxyAddress(&bind.CallOpts{Context: ctx})
		if err != nil {
			err = errors.Errorf("GetProxyAddress failed: %w", err)
			e.logger.Error(&OnchainError{err: err, t: time.Now()})
			client.Close()
			continue
		}
		commitRevealAddr, err := bridge.GetCommitRevealAddress(&bind.CallOpts{Context: ctx})
		if err != nil {
			err = errors.Errorf("GetCommitRevealAddress failed: %w", err)
			e.logger.Error(&OnchainError{err: err, t: time.Now()})
			client.Close()
			continue
		}

		p, err := dosproxy.NewDosproxy(proxyAddr, client)
		if err != nil {
			err = errors.Errorf("NewDosproxy failed: %w", err)
			e.logger.Error(&OnchainError{err: err, t: time.Now()})
			client.Close()
			continue
		}

		cr, err := commitreveal.NewCommitreveal(commitRevealAddr, client)
		if err != nil {
			err = errors.Errorf("NewCommitreveal failed : %w", err)
			e.logger.Error(&OnchainError{err: err, t: time.Now()})
			client.Close()
			continue
		}
		e.logger.Event("ConnToOnchain", map[string]interface{}{"OnchainURL": result.url})

		e.clients = append(e.clients, client)
		e.proxies = append(e.proxies, &dosproxy.DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: e.ctx}, TransactOpts: *e.auth})
		e.crs = append(e.crs, &commitreveal.CommitrevealSession{Contract: cr, CallOpts: bind.CallOpts{Context: e.ctx}, TransactOpts: *e.auth})
	}

	if len(e.proxies) == 0 {
		err = &OnchainError{err: errors.Errorf("Connect failed: %w", ErrNoProxy), t: time.Now()}
		e.logger.Error(err)
		return
	}
	if len(e.crs) == 0 {
		err = &OnchainError{err: errors.Errorf("Connect failed: %w", ErrNoCR), t: time.Now()}
		e.logger.Error(err)
		return
	}
	fmt.Println("Connect :", len(e.clients), len(e.proxies), len(e.crs), err)
	err = nil
	return
}

func (e *ethAdaptor) UpdateWsUrls(urls []string) {
	if err := e.isDone(); err != nil {
		return
	}
	e.wsUrls = urls
}

// Address gets the string representation of the underlying address.
func (e *ethAdaptor) Address() (addr common.Address) {
	if err := e.isDone(); err != nil {
		return
	}
	return e.key.Address
}
