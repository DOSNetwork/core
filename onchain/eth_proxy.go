package onchain

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/DOSNetwork/core/configuration"
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
	rpcUrls          []string
	wsUrls           []string
	key              *keystore.Key
	blockTime        uint64
	gasPrice         uint64
	gasLimit         uint64
	connTimeout      time.Duration
	getTimeout       time.Duration
	setTimeout       time.Duration
	proxies          []*dosproxy.DosproxySession
	wsProxies        []*dosproxy.DosproxySession
	crs              []*commitreveal.CommitrevealSession
	wsCrs            []*commitreveal.CommitrevealSession
	rpcClients       []*ethclient.Client
	wsClients        []*ethclient.Client
	ctxes            []context.Context
	wsCtxes          []context.Context
	cancels          []context.CancelFunc
	wsCancels        []context.CancelFunc
	ctx              context.Context
	cancelFunc       context.CancelFunc
	reqQueue         chan *request
}

//NewEthAdaptor creates an eth implemention of ProxyAdapter
func NewEthAdaptor(key *keystore.Key, config *configuration.Config, l logger) (adaptor *ethAdaptor, err error) {
	if key == nil {
		return nil, errors.New("no keystore")
	}
	if !common.IsHexAddress(config.DOSAddressBridgeAddress) {
		return nil, errors.New("not a valid hex address")
	}
	gasLimitInt, err := strconv.Atoi(config.EthGasLimit)
	if err != nil {
		return nil, err
	}
	blockTimeInt, err := strconv.Atoi(config.BlockTime)
	if err != nil {
		return nil, err
	}

	adaptor = &ethAdaptor{}
	adaptor.bridgeAddr = common.HexToAddress(config.DOSAddressBridgeAddress)
	adaptor.key = key
	adaptor.logger = l
	adaptor.reqQueue = make(chan *request)
	adaptor.gasLimit = uint64(gasLimitInt)
	adaptor.blockTime = uint64(blockTimeInt)
	adaptor.connTimeout = 60 * time.Second
	adaptor.getTimeout = 60 * time.Second
	adaptor.setTimeout = 60 * time.Second
	return
}

// Close RPC & WS connections to eth and release all resources
func (e *ethAdaptor) DisconnectAll() {
	e.cancelFunc()
	for {
		if !e.isConnecting() {
			return
		}
	}
	return
}

func (e *ethAdaptor) DisconnectWs(idx int) {
	if e.wsCancels[idx] != nil {
		e.wsCancels[idx]()
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
	rpcConn, wsConn := false, false
C1:
	for _, ctx := range e.ctxes {
		select {
		case <-ctx.Done():
		default:
			rpcConn = true
			break C1
		}
	}
C2:
	for _, wsCtx := range e.wsCtxes {
		select {
		case <-wsCtx.Done():
		default:
			wsConn = true
			break C2
		}
	}
	return (rpcConn && wsConn)
}

func (e *ethAdaptor) Connect(urls []string, t time.Time) (err error) {
	if e.isConnecting() {
		err = errors.Errorf("connection existing")
		return
	}

	var rpcUrls []string
	var wsUrls []string
	for _, url := range urls {
		if strings.Contains(url, "http") {
			rpcUrls = append(rpcUrls, url)
		} else if strings.Contains(url, "ws") {
			wsUrls = append(wsUrls, url)
		}
	}
	e.rpcUrls = rpcUrls
	if len(rpcUrls) == 0 {
		return errors.New("no rpc URL")
	}
	e.wsUrls = wsUrls
	if len(wsUrls) == 0 {
		return errors.New("no web socket URL")
	}

	e.ctx, e.cancelFunc = context.WithCancel(context.Background())
	e.rpcClients = nil
	e.wsClients = nil
	e.proxies = nil
	e.wsProxies = nil
	e.crs = nil
	e.wsCrs = nil
	e.ctxes = nil
	e.wsCtxes = nil
	e.cancels = nil
	e.wsCancels = nil

	dialCtx, dialCancel := context.WithDeadline(e.ctx, t)
	defer dialCancel()

	wsResults := DialToEth(dialCtx, e.wsUrls, false)
	rpcResults := DialToEth(dialCtx, e.rpcUrls, true)
	wsConnected := false

	for rpc := range rpcResults {
		if rpc.Err != nil {
			e.logger.Error(errors.Errorf(": %w", rpc.Err))
			continue
		}
		rpcClient := rpc.Client
		bridge, err := dosbridge.NewDosbridge(e.bridgeAddr, rpcClient)
		if err != nil {
			e.logger.Error(errors.Errorf(": %w", err))
			rpcClient.Close()
			continue
		}
		proxyAddr, err := bridge.GetProxyAddress(&bind.CallOpts{Context: dialCtx})
		if err != nil {
			e.logger.Error(errors.Errorf(": %w", err))
			rpcClient.Close()
			continue
		}
		commitRevealAddr, err := bridge.GetCommitRevealAddress(&bind.CallOpts{Context: dialCtx})
		if err != nil {
			e.logger.Error(errors.Errorf(": %w", err))
			rpcClient.Close()
			continue
		}

		bootStrapUrl, err := bridge.GetBootStrapUrl(&bind.CallOpts{Context: dialCtx})
		if err != nil {
			e.logger.Error(errors.Errorf(": %w", err))
			rpcClient.Close()
			continue
		}

		p, err := dosproxy.NewDosproxy(proxyAddr, rpcClient)
		if err != nil {
			err = errors.Errorf("NewDosproxy rpc failed: %w", err)
			e.logger.Error(err)
			rpcClient.Close()
			continue
		}

		cr, err := commitreveal.NewCommitreveal(commitRevealAddr, rpcClient)
		if err != nil {
			err = errors.Errorf("NewCommitreveal rpc failed : %w", err)
			e.logger.Error(err)
			rpcClient.Close()
			continue
		}
		e.logger.Event("RPC_ConnToOnchain", map[string]interface{}{"OnchainURL": rpc.Url})

		ctx, cancel := context.WithCancel(e.ctx)
		ctx = context.WithValue(ctx, "index", len(e.ctxes))
		auth := bind.NewKeyedTransactor(e.key.PrivateKey)
		auth.GasLimit = e.gasLimit
		auth.Context = ctx

		if bootStrapUrl != "" {
			e.bootStrapUrl = bootStrapUrl
		}
		e.rpcClients = append(e.rpcClients, rpcClient)
		e.proxies = append(e.proxies, &dosproxy.DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: ctx}, TransactOpts: *auth})
		e.crs = append(e.crs, &commitreveal.CommitrevealSession{Contract: cr, CallOpts: bind.CallOpts{Context: ctx}, TransactOpts: *auth})
		e.ctxes = append(e.ctxes, ctx)
		e.cancels = append(e.cancels, cancel)

		if !wsConnected {
			for ws := range wsResults {
				if ws.Err != nil {
					e.logger.Error(errors.Errorf(": %w", ws.Err))
					continue
				}
				wsClient := ws.Client
				ws_p, err := dosproxy.NewDosproxy(proxyAddr, wsClient)
				if err != nil {
					err = errors.Errorf("NewDosproxy ws failed: %w", err)
					e.logger.Error(err)
					wsClient.Close()
					continue
				}
				ws_cr, err := commitreveal.NewCommitreveal(commitRevealAddr, wsClient)
				if err != nil {
					err = errors.Errorf("NewCommitreveal ws failed : %w", err)
					e.logger.Error(err)
					wsClient.Close()
					continue
				}
				e.logger.Event("WS_ConnToOnchain", map[string]interface{}{"OnchainURL": ws.Url})

				ctx, cancel := context.WithCancel(e.ctx)
				ctx = context.WithValue(ctx, "wsIndex", len(e.wsCtxes))
				auth := bind.NewKeyedTransactor(e.key.PrivateKey)
				auth.GasLimit = e.gasLimit
				auth.Context = ctx

				e.wsClients = append(e.wsClients, wsClient)
				e.wsProxies = append(e.wsProxies, &dosproxy.DosproxySession{Contract: ws_p, CallOpts: bind.CallOpts{Context: ctx}, TransactOpts: *auth})
				e.wsCrs = append(e.wsCrs, &commitreveal.CommitrevealSession{Contract: ws_cr, CallOpts: bind.CallOpts{Context: ctx}, TransactOpts: *auth})
				e.wsCtxes = append(e.wsCtxes, ctx)
				e.wsCancels = append(e.wsCancels, cancel)
			}
			if len(e.wsCtxes) > 0 {
				wsConnected = true
			}
		}
	}

	if len(e.proxies) == 0 || len(e.wsProxies) == 0 {
		err = errors.New("no proxy instance")
		e.logger.Error(err)
		return
	}
	if len(e.crs) == 0 || len(e.wsCrs) == 0 {
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
