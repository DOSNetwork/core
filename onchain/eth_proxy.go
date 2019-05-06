package onchain

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain/commitreveal"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SubscribeDosproxyLogUpdateRandom = iota
	SubscribeDosproxyLogRequestUserRandom
	SubscribeDosproxyLogUrl
	SubscribeDosproxyLogValidationResult
	SubscribeDosproxyLogGrouping
	SubscribeDosproxyLogPublicKeyAccepted
	SubscribeDosproxyLogPublicKeySuggested
	SubscribeDosproxyLogGroupDissolve
	SubscribeDosproxyLogInsufficientPendingNode
	SubscribeDosproxyLogInsufficientWorkingGroup
	SubscribeDosproxyLogNoWorkingGroup
	SubscribeDosproxyLogGroupingInitiated
	SubscribeDosproxyUpdateGroupToPick
	//For bootstraping
	SubscribeCommitrevealLogStartCommitreveal
	SubscribeCommitrevealLogCommit
	SubscribeCommitrevealLogReveal
	SubscribeCommitrevealLogRandom
	LastRandomness
	WorkingGroupSize
	LastUpdatedBlock
	PendingNonce
	GroupToPick
	NumPendingGroups
	PengindNodeSize
	GroupSize
)

// TODO: Move constants to some unified places.
const (
	TrafficSystemRandom = iota // 0
	TrafficUserRandom
	TrafficUserQuery
)

const (
	LogBlockDiff        = 1
	LogCheckingInterval = 15 //in second
	SubscribeTimeout    = 60 //in second
)

var logger log.Logger

type Request struct {
	ctx    context.Context
	tOpts  *bind.TransactOpts
	f      func() (*types.Transaction, error)
	result chan Reply
}

type ReqSetInt struct {
	ctx       context.Context
	tOpts     *bind.TransactOpts
	parameter *big.Int
	f         func(*big.Int) (*types.Transaction, error)
	result    chan Reply
}

type ReqCommit struct {
	ctx        context.Context
	tOpts      *bind.TransactOpts
	cid        *big.Int
	secretHash [32]byte
	f          func(*big.Int, [32]byte) (*types.Transaction, error)
	result     chan Reply
}

type ReqReveal struct {
	ctx    context.Context
	tOpts  *bind.TransactOpts
	cid    *big.Int
	secret *big.Int
	f      func(*big.Int, *big.Int) (*types.Transaction, error)
	result chan Reply
}

type ReqGrouping struct {
	ctx        context.Context
	tOpts      *bind.TransactOpts
	candidates []common.Address
	size       *big.Int
	f          func([]common.Address, *big.Int) (*types.Transaction, error)
	result     chan Reply
}

type ReqSetRandomNum struct {
	ctx     context.Context
	tOpts   *bind.TransactOpts
	sig     [2]*big.Int
	version uint8
	f       func([2]*big.Int, uint8) (*types.Transaction, error)
	result  chan Reply
}

type ReqSetPublicKey struct {
	ctx     context.Context
	tOpts   *bind.TransactOpts
	groupId *big.Int
	pubKey  [4]*big.Int
	f       func(*big.Int, [4]*big.Int) (*types.Transaction, error)
	result  chan Reply
}

type ReqTriggerCallback struct {
	ctx         context.Context
	tOpts       *bind.TransactOpts
	requestId   *big.Int
	trafficType uint8
	content     []byte
	sig         [2]*big.Int
	version     uint8
	f           func(*big.Int, uint8, []byte, [2]*big.Int, uint8) (*types.Transaction, error)
	result      chan Reply
}

type Reply struct {
	tx    *types.Transaction
	nonce uint64
	err   error
}

type EthAdaptor struct {
	proxyAddr        string
	commitRevealAddr string
	gethUrls         []string
	eventUrls        []string
	key              *keystore.Key
	auth             *bind.TransactOpts
	//rpc connection over http/https
	proxies    []*dosproxy.DosproxySession
	crs        []*commitreveal.CommitrevealSession
	clients    []*ethclient.Client
	ctx        context.Context
	cancelFunc context.CancelFunc
	reqQueue   chan interface{}

	//rpc connection over WebSockets for event notification
	eProxies    []*dosproxy.Dosproxy
	eCRs        []*commitreveal.Commitreveal
	eClients    []*ethclient.Client
	eCtx        context.Context
	eCancelFunc context.CancelFunc
}

func NewEthAdaptor(credentialPath, passphrase, proxyAddr, commitRevealAddr string, urls []string) (adaptor *EthAdaptor, err error) {
	var httpUrls []string
	var wsUrls []string
	for _, url := range urls {
		if strings.Contains(url, "http") {
			httpUrls = append(httpUrls, url)
		} else if strings.Contains(url, "ws") {
			wsUrls = append(wsUrls, url)
		}
	}
	fmt.Println("gethUrls ", httpUrls)
	fmt.Println("eventUrls ", wsUrls)

	adaptor = &EthAdaptor{}
	adaptor.gethUrls = httpUrls
	adaptor.eventUrls = wsUrls
	adaptor.proxyAddr = proxyAddr

	//Read Ethereum keystore
	key, err := ReadEthKey(credentialPath, passphrase)
	if err != nil {
		logger.Error(err)
		return
	}
	adaptor.key = key
	debug.FreeOSMemory()

	//Use account address as ID to init log module
	log.Init(key.Address.Bytes()[:])
	if logger == nil {
		logger = log.New("module", "EthProxy")
		fmt.Println("logger init EthProxy")
	}

	adaptor.ctx, adaptor.cancelFunc = context.WithCancel(context.Background())
	adaptor.auth = bind.NewKeyedTransactor(key.PrivateKey)
	adaptor.auth.GasLimit = uint64(GASLIMIT)
	adaptor.auth.Context = adaptor.ctx

	//
	clients := DialToEth(context.Background(), httpUrls, key)
	for client := range clients {
		p, e := dosproxy.NewDosproxy(common.HexToAddress(proxyAddr), client)
		if e != nil {
			fmt.Println("NewDosproxy e ", e)
			logger.Error(e)
			err = e
			continue
		}
		c, e := commitreveal.NewCommitreveal(common.HexToAddress(commitRevealAddr), client)
		if e != nil {
			logger.Error(e)
			err = e
			continue
		}
		adaptor.clients = append(adaptor.clients, client)
		adaptor.proxies = append(adaptor.proxies, &dosproxy.DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: adaptor.ctx}, TransactOpts: *adaptor.auth})
		adaptor.crs = append(adaptor.crs, &commitreveal.CommitrevealSession{Contract: c, CallOpts: bind.CallOpts{Context: adaptor.ctx}, TransactOpts: *adaptor.auth})
	}
	if len(adaptor.proxies) == 0 {
		fmt.Println("No any working eth client ", len(adaptor.clients), len(adaptor.proxies))
		adaptor = nil
		return
	}
	adaptor.reqQueue = make(chan interface{})
	fmt.Println("working eth http client ", len(adaptor.clients), "Balance ", adaptor.GetBalance())

	adaptor.eCtx, adaptor.eCancelFunc = context.WithCancel(context.Background())
	syncClients := CheckSync(adaptor.eCtx, adaptor.clients[0], DialToEth(context.Background(), wsUrls, key))
	for client := range syncClients {
		fmt.Println("syncClients")
		p, err := dosproxy.NewDosproxy(common.HexToAddress(proxyAddr), client)
		if err != nil {
			logger.Error(err)
			continue
		}
		c, err := commitreveal.NewCommitreveal(common.HexToAddress(commitRevealAddr), client)
		if err != nil {
			logger.Error(err)
			continue
		}
		adaptor.eCRs = append(adaptor.eCRs, c)
		adaptor.eClients = append(adaptor.eClients, client)
		adaptor.eProxies = append(adaptor.eProxies, p)
	}
	if len(adaptor.eProxies) == 0 {
		err = errors.New("No any working eth client for event tracking")
		adaptor = nil
		return
	}
	fmt.Println("working eth ws client ", len(adaptor.eClients))

	adaptor.reqLoop()

	return
}

func (e *EthAdaptor) AddEventNode() (err error) {
	/*
		e.eCtx, e.eCancelFunc = context.WithCancel(context.Background())
		syncClients := CheckSync(e.eCtx, e.clients[0], DialToEth(context.Background(), e.eventUrls, e.key))
		for client := range syncClients {
			fmt.Println("syncClients")
			p, err := dosproxy.NewDosproxy(common.HexToAddress(e.proxyAddr), client)
			if err != nil {
				logger.Error(err)
				continue
			}
			c, err := commitreveal.NewCommitreveal(common.HexToAddress(e.commitRevealAddr), client)
			if err != nil {
				logger.Error(err)
				continue
			}
			e.eCRs = append(e.eCRs, c)
			e.eClients = append(e.eClients, client)
			e.eProxies = append(e.eProxies, p)
		}
		if len(e.eProxies) == 0 {
			err = errors.New("No any working eth client for event tracking")
		}
		fmt.Println("working eth ws client ", len(e.eClients))
	*/
	return
}

func (e *EthAdaptor) End() {
	//e.cancel()
	//e.c.Close()
	return
}

func (e *EthAdaptor) reqLoop() {
	go func() {
		defer fmt.Println("reqLoop exit")
		defer close(e.reqQueue)

		for {
			select {
			case req := <-e.reqQueue:
				var tx *types.Transaction
				var err error
				var resultC chan Reply
				var ctx context.Context

				switch content := req.(type) {
				case *Request:
					resultC = content.result
				case *ReqGrouping:
					resultC = content.result
				case *ReqSetInt:
					resultC = content.result
				case *ReqCommit:
					resultC = content.result
				case *ReqReveal:
					resultC = content.result
				case *ReqSetRandomNum:
					resultC = content.result
				case *ReqTriggerCallback:
					resultC = content.result
				case *ReqSetPublicKey:
					resultC = content.result
				}
				reply := Reply{}

				//TODO:Change this to Fan In and save nonce to e.auth

				nonce, err := e.PendingNonce()
				if err != nil {
					reply.err = err
					resultC <- reply
					continue
				}

				nonceBig := new(big.Int).SetUint64(nonce)
				if e.auth.Nonce == nil {
					e.auth.Nonce = nonceBig
				} else if e.auth.Nonce.Cmp(nonceBig) == -1 {
					e.auth.Nonce = nonceBig
				}
				fmt.Println("Got a request nonce , ", e.auth.Nonce)

				switch content := req.(type) {
				case *Request:
					content.tOpts.Nonce = e.auth.Nonce
					tx, err = content.f()
					resultC = content.result
					ctx = content.ctx
				case *ReqGrouping:
					content.tOpts.Nonce = e.auth.Nonce
					tx, err = content.f(content.candidates, content.size)
					resultC = content.result
					ctx = content.ctx
				case *ReqSetInt:
					content.tOpts.Nonce = e.auth.Nonce
					tx, err = content.f(content.parameter)
					resultC = content.result
					ctx = content.ctx
				case *ReqCommit:
					content.tOpts.Nonce = e.auth.Nonce
					tx, err = content.f(content.cid, content.secretHash)
					resultC = content.result
					ctx = content.ctx
				case *ReqReveal:
					content.tOpts.Nonce = e.auth.Nonce
					tx, err = content.f(content.cid, content.secret)
					resultC = content.result
					ctx = content.ctx
				case *ReqSetRandomNum:
					content.tOpts.Nonce = e.auth.Nonce
					tx, err = content.f(content.sig, content.version)
					resultC = content.result
					ctx = content.ctx
				case *ReqTriggerCallback:
					content.tOpts.Nonce = e.auth.Nonce
					tx, err = content.f(content.requestId, content.trafficType, content.content, content.sig, content.version)
					resultC = content.result
					ctx = content.ctx
				case *ReqSetPublicKey:
					content.tOpts.Nonce = e.auth.Nonce
					tx, err = content.f(content.groupId, content.pubKey)
					resultC = content.result
					ctx = content.ctx
				}
				if err != nil {
					if err.Error() == "replacement transaction underpriced" ||
						strings.Contains(err.Error(), "known transaction") ||
						strings.Contains(err.Error(), "nonce too low") {
						e.auth.Nonce = e.auth.Nonce.Add(e.auth.Nonce, big.NewInt(1))
					}
					reply.err = err
					resultC <- reply
					continue
				}

				reply.tx = tx
				reply.nonce = e.auth.Nonce.Uint64()
				e.auth.Nonce = e.auth.Nonce.Add(e.auth.Nonce, big.NewInt(1))
				select {
				case resultC <- reply:
				case <-ctx.Done():
				}
			case <-e.ctx.Done():
				return
			}
		}
	}()
}

func (e *EthAdaptor) sendRequest(ctx context.Context, c *ethclient.Client, pre <-chan error, request interface{}, result chan Reply) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		retry := 10
		if pre != nil {
			select {
			case <-ctx.Done():
				return
			case err := <-pre:
				//Request has been fulfulled by previous sendRequest
				if err == nil {
					fmt.Println("Request has been fulfulled by previous sendRequest")
					return
				} else {
					if strings.Contains(err.Error(), "transaction failed") {
						fmt.Println("Transaction failed so delete the whole requestSend chain")
						return
					}
					fmt.Println("Switch Client to handle request beacuse of e ,", err)
				}
			}
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			e.reqQueue <- request

			select {
			case reply := <-result:
				err := reply.err
				tx := reply.tx
				//nonce := reply.nonce
				if err != nil {
					if err.Error() == "replacement transaction underpriced" ||
						strings.Contains(err.Error(), "known transaction") ||
						strings.Contains(err.Error(), "nonce too low") {
						continue
					}
					if retry == 0 {
						fmt.Println("Reply err ", err)
						select {
						case errc <- err:
						case <-ctx.Done():
						}
						return
					} else {
						fmt.Println("Retry :Reply err ", err, " Balance ", e.GetBalance())

						time.Sleep(5 * time.Second)
						retry--
						continue
					}
				}
				//TODO : move this out of sendRequest
				err = CheckTransaction(c, tx)
				if err != nil {
					if strings.Contains(err.Error(), "transaction not found") {
						//Resend the request
						continue
					}
					logger.Error(err)
					f := map[string]interface{}{
						"Tx":     tx,
						"ErrMsg": err.Error()}
					logger.Event("TransactionFail", f)
					fmt.Println("TransactionFail err ", err)
					//Don't return err to errc to delete the whole sendRequest chain
					select {
					case errc <- err:
					case <-ctx.Done():
					}
					return
				}
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return errc
}

func (e *EthAdaptor) RegisterNewNode(ctx context.Context) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "RegisterNewNode", nil)

	result := make(chan Reply)
	for i, proxy := range e.proxies {
		request := &Request{ctx, &proxy.TransactOpts, proxy.RegisterNewNode, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func (e *EthAdaptor) SignalRandom(ctx context.Context) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "SignalRandom", nil)
	result := make(chan Reply)
	for i, proxy := range e.proxies {
		request := &Request{ctx, &proxy.TransactOpts, proxy.SignalRandom, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func (e *EthAdaptor) SignalGroupFormation(ctx context.Context) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "SignalGroupFormation", nil)
	result := make(chan Reply)
	for i, proxy := range e.proxies {
		request := &Request{ctx, &proxy.TransactOpts, proxy.SignalGroupFormation, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func (e *EthAdaptor) SignalDissolve(ctx context.Context, idx uint64) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "SignalDissolve", nil)
	result := make(chan Reply)
	x := new(big.Int)
	x.SetUint64(idx)
	for i, proxy := range e.proxies {
		request := &ReqSetInt{ctx, &proxy.TransactOpts, x, proxy.SignalDissolve, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func (e *EthAdaptor) SignalBootstrap(ctx context.Context, cid uint64) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "SignalBootstrap", nil)
	result := make(chan Reply)
	x := new(big.Int)
	x.SetUint64(cid)
	for i, proxy := range e.proxies {
		request := &ReqSetInt{ctx, &proxy.TransactOpts, x, proxy.SignalBootstrap, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func (e *EthAdaptor) Commit(ctx context.Context, cid *big.Int, commitment [32]byte) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "Commit", nil)
	result := make(chan Reply)
	for i, cr := range e.crs {
		request := &ReqCommit{ctx, &cr.TransactOpts, cid, commitment, cr.Commit, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func (e *EthAdaptor) Reveal(ctx context.Context, cid *big.Int, secret *big.Int) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "Reveal", nil)
	result := make(chan Reply)
	for i, cr := range e.crs {
		request := &ReqReveal{ctx, &cr.TransactOpts, cid, secret, cr.Reveal, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func (e *EthAdaptor) RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc <-chan error) {
	go func() {
		select {
		case IdWithPubKey, ok := <-IdWithPubKeys:
			if !ok {
				return
			}
			result := make(chan Reply)
			groupId := IdWithPubKey[0]
			var pubKey [4]*big.Int
			copy(pubKey[:], IdWithPubKey[1:])

			for i, proxy := range e.proxies {
				request := &ReqSetPublicKey{ctx, &proxy.TransactOpts, groupId, pubKey, proxy.RegisterGroupPubKey, result}
				errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
			}

			f := map[string]interface{}{
				"SessionID":         fmt.Sprintf("%x", groupId.Bytes()),
				"DispatchedGroup_1": fmt.Sprintf("%x", pubKey[0].Bytes()),
				"DispatchedGroup_2": fmt.Sprintf("%x", pubKey[1].Bytes()),
				"DispatchedGroup_3": fmt.Sprintf("%x", pubKey[2].Bytes()),
				"DispatchedGroup_4": fmt.Sprintf("%x", pubKey[3].Bytes())}
			logger.Event("DOS_RegisterGroupPubKey", f)
			return
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

func (e *EthAdaptor) SetRandomNum(ctx context.Context, signatures <-chan *vss.Signature) (errc <-chan error) {
	go func() {
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			x, y := DecodeSig(signature.Signature)
			result := make(chan Reply)
			for i, proxy := range e.proxies {
				request := &ReqSetRandomNum{ctx, &proxy.TransactOpts, [2]*big.Int{x, y}, 0, proxy.UpdateRandomness, result}
				errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
			}
			return
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

func (e *EthAdaptor) DataReturn(ctx context.Context, signatures <-chan *vss.Signature) (errc <-chan error) {
	go func() {
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "DataReturn", map[string]interface{}{"RequestId": ctx.Value("RequestID")})

			x, y := DecodeSig(signature.Signature)
			requestId := new(big.Int).SetBytes(signature.RequestId)

			result := make(chan Reply)
			for i, proxy := range e.proxies {
				request := &ReqTriggerCallback{ctx, &proxy.TransactOpts, requestId, uint8(signature.Index), signature.Content, [2]*big.Int{x, y}, 0, proxy.TriggerCallback, result}
				errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
			}
			return
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

func (e *EthAdaptor) SetGroupSize(ctx context.Context, size uint64) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "SetGroupSize", nil)
	result := make(chan Reply)
	x := new(big.Int)
	x.SetUint64(size)
	for i, proxy := range e.proxies {
		request := &ReqSetInt{ctx, &proxy.TransactOpts, x, proxy.SetGroupSize, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}
func (e *EthAdaptor) SetGroupMaturityPeriod(ctx context.Context, period uint64) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "setGroupMaturityPeriod", nil)
	result := make(chan Reply)
	x := new(big.Int)
	x.SetUint64(period)
	for i, proxy := range e.proxies {
		request := &ReqSetInt{ctx, &proxy.TransactOpts, x, proxy.SetGroupMaturityPeriod, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func (e *EthAdaptor) SetGroupToPick(ctx context.Context, groupToPick uint64) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "SetGroupToPick", nil)
	result := make(chan Reply)
	x := new(big.Int)
	x.SetUint64(groupToPick)
	for i, proxy := range e.proxies {
		request := &ReqSetInt{ctx, &proxy.TransactOpts, x, proxy.SetGroupToPick, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func (e *EthAdaptor) SetGroupingThreshold(ctx context.Context, threshold uint64) (errc <-chan error) {
	defer logger.TimeTrack(time.Now(), "SetGroupToPick", nil)
	result := make(chan Reply)
	x := new(big.Int)
	x.SetUint64(threshold)
	for i, proxy := range e.proxies {
		request := &ReqSetInt{ctx, &proxy.TransactOpts, x, proxy.SetGroupingThreshold, result}
		errc = e.sendRequest(ctx, e.clients[i], errc, request, result)
	}
	return
}

func MergeEvents(ctx context.Context, cs ...<-chan interface{}) chan interface{} {
	var wg sync.WaitGroup
	// We must ensure that the output channel has the capacity to
	// hold as many errors
	// as there are error channels.
	// This will ensure that it never blocks, even
	// if WaitForPipeline returns early.
	out := make(chan interface{}, len(cs))
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls
	// wg.Done.
	output := func(c <-chan interface{}) {
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

func MergeErrors(ctx context.Context, cs ...<-chan error) <-chan error {
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

func (e *EthAdaptor) firstEvent(ctx context.Context, source chan interface{}) <-chan interface{} {
	sink := make(chan interface{})

	go func() {
		defer close(sink)
		visited := make(map[string]uint64)
		for {
			var bytes []byte
			var blkNum uint64
			var event interface{}
			var ok bool
			select {
			case event, ok = <-source:
				if ok {
					switch content := event.(type) {
					case *DosproxyLogUpdateRandom:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *DosproxyLogRequestUserRandom:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *DosproxyLogUrl:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *DosproxyLogValidationResult:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *DosproxyLogNoWorkingGroup:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *DosproxyLogGrouping:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *DosproxyLogPublicKeyAccepted:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *DosproxyLogPublicKeySuggested:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *DosproxyLogGroupDissolve:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *DosproxyUpdateGroupToPick:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogStartCommitReveal:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogCommit:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogReveal:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					case *LogRandom:
						bytes = append(bytes, content.Raw.Data...)
						blkNum = content.BlockN
						bytes = append(bytes, new(big.Int).SetUint64(blkNum).Bytes()...)
					}
				} else {
					return
				}
			}
			nHash := sha256.Sum256(bytes)
			identity := string(nHash[:])
			curBlk, err := e.CurrentBlock()
			if err != nil {
				continue
			}

			if visited[identity] == 0 {
				visited[identity] = blkNum
				select {
				case sink <- event:
				case <-ctx.Done():
				}
			} else {
				for k, blkN := range visited {
					if curBlk >= blkN+100 {
						delete(visited, k)
					}
				}
			}
		}
	}()

	return sink
}

func (e *EthAdaptor) SubscribeEvent(subscribeType int) (<-chan interface{}, <-chan error) {
	var eventList []<-chan interface{}
	var errcs []<-chan error
	if subscribeType == SubscribeCommitrevealLogStartCommitreveal ||
		subscribeType == SubscribeCommitrevealLogCommit ||
		subscribeType == SubscribeCommitrevealLogReveal ||
		subscribeType == SubscribeCommitrevealLogRandom {
		for i := 0; i < len(e.eProxies); i++ {
			fmt.Println("SubscribeEvent ", i)
			cr := e.eCRs[i]
			if cr == nil {
				continue
			}
			ctx := e.eCtx
			if ctx == nil {
				continue
			}
			out, errc := subscribeCREvent(ctx, cr, subscribeType)
			eventList = append(eventList, out)
			errcs = append(errcs, errc)
		}
	} else {
		for i := 0; i < len(e.eProxies); i++ {
			fmt.Println("SubscribeEvent ", i)
			proxy := e.eProxies[i]
			if proxy == nil {
				continue
			}
			ctx := e.eCtx
			if ctx == nil {
				continue
			}
			out, errc := subscribeEvent(ctx, proxy, subscribeType)
			eventList = append(eventList, out)
			errcs = append(errcs, errc)
		}
	}
	return e.firstEvent(e.ctx, MergeEvents(e.ctx, eventList...)), MergeErrors(e.ctx, errcs...)
}
func subscribeCREvent(ctx context.Context, cr *commitreveal.Commitreveal, subscribeType int) (<-chan interface{}, <-chan error) {
	out := make(chan interface{})
	errc := make(chan error)
	opt := &bind.WatchOpts{}

	switch subscribeType {
	case SubscribeCommitrevealLogStartCommitreveal:
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogStartCommitReveal)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.CommitrevealFilterer.WatchLogStartCommitReveal(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogStartCommitReveal{
						Cid:             i.Cid,
						StartBlock:      i.StartBlock,
						CommitDuration:  i.CommitDuration,
						RevealDuration:  i.RevealDuration,
						RevealThreshold: i.RevealThreshold,
						Tx:              i.Raw.TxHash.Hex(),
						BlockN:          i.Raw.BlockNumber,
						Removed:         i.Raw.Removed,
						Raw:             i.Raw,
					}
				}
			}
		}()
	case SubscribeCommitrevealLogCommit:
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogCommit)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.CommitrevealFilterer.WatchLogCommit(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogCommit{
						Cid:        i.Cid,
						From:       i.From,
						Commitment: i.Commitment,
						Tx:         i.Raw.TxHash.Hex(),
						BlockN:     i.Raw.BlockNumber,
						Removed:    i.Raw.Removed,
						Raw:        i.Raw,
					}
				}
			}
		}()
	case SubscribeCommitrevealLogReveal:
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogReveal)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.CommitrevealFilterer.WatchLogReveal(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogReveal{
						Cid:     i.Cid,
						From:    i.From,
						Secret:  i.Secret,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeCommitrevealLogRandom:
		go func() {
			transitChan := make(chan *commitreveal.CommitrevealLogRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := cr.CommitrevealFilterer.WatchLogRandom(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &LogRandom{
						Cid:     i.Cid,
						Random:  i.Random,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	}
	return out, errc
}

func subscribeEvent(ctx context.Context, proxy *dosproxy.Dosproxy, subscribeType int) (<-chan interface{}, <-chan error) {
	out := make(chan interface{})
	errc := make(chan error)
	opt := &bind.WatchOpts{}

	switch subscribeType {
	case SubscribeDosproxyLogUpdateRandom:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogUpdateRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DosproxyFilterer.WatchLogUpdateRandom(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogUpdateRandom{
						LastRandomness:    i.LastRandomness,
						DispatchedGroupId: i.DispatchedGroupId,
						DispatchedGroup:   i.DispatchedGroup,
						Tx:                i.Raw.TxHash.Hex(),
						BlockN:            i.Raw.BlockNumber,
						Removed:           i.Raw.Removed,
						Raw:               i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogUrl:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogUrl)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.DosproxyFilterer.WatchLogUrl(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogUrl{
						QueryId:           i.QueryId,
						Timeout:           i.Timeout,
						DataSource:        i.DataSource,
						Selector:          i.Selector,
						Randomness:        i.Randomness,
						DispatchedGroupId: i.DispatchedGroupId,
						DispatchedGroup:   i.DispatchedGroup,
						Tx:                i.Raw.TxHash.Hex(),
						BlockN:            i.Raw.BlockNumber,
						Removed:           i.Raw.Removed,
						Raw:               i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogRequestUserRandom:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogRequestUserRandom)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.DosproxyFilterer.WatchLogRequestUserRandom(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogRequestUserRandom{
						RequestId:            i.RequestId,
						LastSystemRandomness: i.LastSystemRandomness,
						UserSeed:             i.UserSeed,
						DispatchedGroupId:    i.DispatchedGroupId,
						DispatchedGroup:      i.DispatchedGroup,
						Tx:                   i.Raw.TxHash.Hex(),
						BlockN:               i.Raw.BlockNumber,
						Removed:              i.Raw.Removed,
						Raw:                  i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogValidationResult:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogValidationResult)
			defer close(transitChan)
			defer close(errc)
			defer close(out)

			sub, err := proxy.DosproxyFilterer.WatchLogValidationResult(opt, transitChan)
			if err != nil {
				fmt.Println("SubscribeDosproxyLogValidationResult err", err)
				return
			}
			for {
				select {
				case <-ctx.Done():
					fmt.Println("SubscribeDosproxyLogValidationResult Done")

					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					fmt.Println("SubscribeDosproxyLogValidationResult err", err)

					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogValidationResult{
						TrafficType: i.TrafficType,
						TrafficId:   i.TrafficId,
						Message:     i.Message,
						Signature:   i.Signature,
						PubKey:      i.PubKey,
						Pass:        i.Pass,
						Version:     i.Version,
						Tx:          i.Raw.TxHash.Hex(),
						BlockN:      i.Raw.BlockNumber,
						Removed:     i.Raw.Removed,
						Raw:         i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogInsufficientPendingNode:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogInsufficientPendingNode)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DosproxyFilterer.WatchLogInsufficientPendingNode(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogInsufficientPendingNode{
						NumPendingNodes: i.NumPendingNodes,
						Tx:              i.Raw.TxHash.Hex(),
						BlockN:          i.Raw.BlockNumber,
						Removed:         i.Raw.Removed,
						Raw:             i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogInsufficientWorkingGroup:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogInsufficientWorkingGroup)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DosproxyFilterer.WatchLogInsufficientWorkingGroup(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogInsufficientWorkingGroup{
						NumWorkingGroups: i.NumWorkingGroups,
						Tx:               i.Raw.TxHash.Hex(),
						BlockN:           i.Raw.BlockNumber,
						Removed:          i.Raw.Removed,
						Raw:              i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogGroupingInitiated:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogGroupingInitiated)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DosproxyFilterer.WatchLogGroupingInitiated(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogGroupingInitiated{
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyUpdateGroupToPick:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyUpdateGroupToPick)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DosproxyFilterer.WatchUpdateGroupToPick(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyUpdateGroupToPick{
						Tx:      i.Raw.TxHash.Hex(),
						OldNum:  i.OldNum,
						NewNum:  i.NewNum,
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogGrouping:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogGrouping)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DosproxyFilterer.WatchLogGrouping(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogGrouping{
						GroupId: i.GroupId,
						NodeId:  i.NodeId,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogPublicKeyAccepted:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogPublicKeyAccepted)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DosproxyFilterer.WatchLogPublicKeyAccepted(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogPublicKeyAccepted{
						GroupId:          i.GroupId,
						PubKey:           i.PubKey,
						WorkingGroupSize: i.WorkingGroupSize,
						Tx:               i.Raw.TxHash.Hex(),
						BlockN:           i.Raw.BlockNumber,
						Removed:          i.Raw.Removed,
						Raw:              i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogPublicKeySuggested:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogPublicKeySuggested)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DosproxyFilterer.WatchLogPublicKeySuggested(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogPublicKeySuggested{
						GroupId:   i.GroupId,
						PubKey:    i.PubKey,
						Count:     i.Count,
						GroupSize: i.GroupSize,
						Tx:        i.Raw.TxHash.Hex(),
						BlockN:    i.Raw.BlockNumber,
						Removed:   i.Raw.Removed,
						Raw:       i.Raw,
					}
				}
			}
		}()
	case SubscribeDosproxyLogGroupDissolve:
		go func() {
			transitChan := make(chan *dosproxy.DosproxyLogGroupDissolve)
			defer close(transitChan)
			defer close(errc)
			defer close(out)
			sub, err := proxy.DosproxyFilterer.WatchLogGroupDissolve(opt, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					out <- &DosproxyLogGroupDissolve{
						PubKey:  i.PubKey,
						GroupId: i.GroupId,
						Tx:      i.Raw.TxHash.Hex(),
						BlockN:  i.Raw.BlockNumber,
						Removed: i.Raw.Removed,
						Raw:     i.Raw,
					}
				}
			}
		}()
	}
	return out, errc
}

func (e *EthAdaptor) PollLogs(subscribeType int, logBlockDiff, preBlockBuf uint64) (<-chan interface{}, <-chan error) {
	var errcs []<-chan error
	var sinks []<-chan interface{}
	var wg sync.WaitGroup

	multiplex := func(client *ethclient.Client, proxyFilter *dosproxy.DosproxyFilterer, ctx context.Context) {
		errc := make(chan error)
		errcs = append(errcs, errc)
		sink := make(chan interface{})
		sinks = append(sinks, sink)
		wg.Done()
		defer close(errc)
		defer close(sink)
		targetBlockN, err := GetCurrentBlock(client)
		if err != nil {
			return
		}
		targetBlockN -= preBlockBuf
		timer := time.NewTimer(LogCheckingInterval * time.Second)
		for {
			select {
			case <-timer.C:
				currentBlockN, err := GetCurrentBlock(client)
				if err != nil {
					timer.Reset(LogCheckingInterval * time.Second)
					continue
				}

				for ; currentBlockN-logBlockDiff >= targetBlockN; targetBlockN++ {
					switch subscribeType {
					case SubscribeDosproxyLogGrouping:
						logs, err := proxyFilter.FilterLogGrouping(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DosproxyLogGrouping{
								GroupId: logs.Event.GroupId,
								NodeId:  logs.Event.NodeId,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
								Raw:     logs.Event.Raw,
							}
						}
					case SubscribeDosproxyLogGroupDissolve:
						logs, err := proxyFilter.FilterLogGroupDissolve(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DosproxyLogGroupDissolve{
								PubKey:  logs.Event.PubKey,
								GroupId: logs.Event.GroupId,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
								Raw:     logs.Event.Raw,
							}
						}
					case SubscribeDosproxyLogUpdateRandom:
						logs, err := proxyFilter.FilterLogUpdateRandom(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DosproxyLogUpdateRandom{
								LastRandomness:    logs.Event.LastRandomness,
								DispatchedGroupId: logs.Event.DispatchedGroupId,
								DispatchedGroup:   logs.Event.DispatchedGroup,
								Tx:                logs.Event.Raw.TxHash.Hex(),
								BlockN:            logs.Event.Raw.BlockNumber,
								Removed:           logs.Event.Raw.Removed,
								Raw:               logs.Event.Raw,
							}
						}
					case SubscribeDosproxyLogRequestUserRandom:
						logs, err := proxyFilter.FilterLogRequestUserRandom(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DosproxyLogRequestUserRandom{
								RequestId:            logs.Event.RequestId,
								LastSystemRandomness: logs.Event.LastSystemRandomness,
								UserSeed:             logs.Event.UserSeed,
								DispatchedGroupId:    logs.Event.DispatchedGroupId,
								DispatchedGroup:      logs.Event.DispatchedGroup,
								Tx:                   logs.Event.Raw.TxHash.Hex(),
								BlockN:               logs.Event.Raw.BlockNumber,
								Removed:              logs.Event.Raw.Removed,
								Raw:                  logs.Event.Raw,
							}
						}
					case SubscribeDosproxyLogUrl:
						logs, err := proxyFilter.FilterLogUrl(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DosproxyLogUrl{
								QueryId:           logs.Event.QueryId,
								Timeout:           logs.Event.Timeout,
								DataSource:        logs.Event.DataSource,
								Selector:          logs.Event.Selector,
								Randomness:        logs.Event.Randomness,
								DispatchedGroupId: logs.Event.DispatchedGroupId,
								DispatchedGroup:   logs.Event.DispatchedGroup,
								Tx:                logs.Event.Raw.TxHash.Hex(),
								BlockN:            logs.Event.Raw.BlockNumber,
								Removed:           logs.Event.Raw.Removed,
								Raw:               logs.Event.Raw,
							}
						}
					case SubscribeDosproxyLogValidationResult:
						logs, err := proxyFilter.FilterLogValidationResult(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}
						for logs.Next() {
							sink <- &DosproxyLogValidationResult{
								TrafficType: logs.Event.TrafficType,
								TrafficId:   logs.Event.TrafficId,
								Message:     logs.Event.Message,
								Signature:   logs.Event.Signature,
								PubKey:      logs.Event.PubKey,
								Pass:        logs.Event.Pass,
								Version:     logs.Event.Version,
								Tx:          logs.Event.Raw.TxHash.Hex(),
								BlockN:      logs.Event.Raw.BlockNumber,
								Removed:     logs.Event.Raw.Removed,
								Raw:         logs.Event.Raw,
							}
						}
					case SubscribeDosproxyLogInsufficientPendingNode:
						logs, err := proxyFilter.FilterLogInsufficientPendingNode(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}

						for logs.Next() {
							sink <- &DosproxyLogInsufficientPendingNode{
								NumPendingNodes: logs.Event.NumPendingNodes,
								Tx:              logs.Event.Raw.TxHash.Hex(),
								BlockN:          logs.Event.Raw.BlockNumber,
								Removed:         logs.Event.Raw.Removed,
								Raw:             logs.Event.Raw,
							}
						}
					case SubscribeDosproxyLogInsufficientWorkingGroup:
						logs, err := proxyFilter.FilterLogInsufficientWorkingGroup(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}

						for logs.Next() {
							sink <- &DosproxyLogInsufficientWorkingGroup{
								NumWorkingGroups: logs.Event.NumWorkingGroups,
								Tx:               logs.Event.Raw.TxHash.Hex(),
								BlockN:           logs.Event.Raw.BlockNumber,
								Removed:          logs.Event.Raw.Removed,
								Raw:              logs.Event.Raw,
							}
						}
					case SubscribeDosproxyLogGroupingInitiated:
						logs, err := proxyFilter.FilterLogGroupingInitiated(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}

						for logs.Next() {
							sink <- &DosproxyLogGroupingInitiated{
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
								Raw:     logs.Event.Raw,
							}
						}
					case SubscribeDosproxyUpdateGroupToPick:
						logs, err := proxyFilter.FilterUpdateGroupToPick(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							errc <- err
							continue
						}

						for logs.Next() {
							sink <- &DosproxyUpdateGroupToPick{
								Tx:      logs.Event.Raw.TxHash.Hex(),
								OldNum:  logs.Event.OldNum,
								NewNum:  logs.Event.NewNum,
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
								Raw:     logs.Event.Raw,
							}
						}
					}
				}
				timer.Reset(LogCheckingInterval * time.Second)
			case <-ctx.Done():
				return
			}
		}

	}

	wg.Add(len(e.proxies))
	for i := 0; i < len(e.clients); i++ {
		go multiplex(e.clients[i], &e.proxies[i].Contract.DosproxyFilterer, e.ctx)
	}
	wg.Wait()

	return e.firstEvent(e.ctx, MergeEvents(e.ctx, sinks...)), MergeErrors(e.ctx, errcs...)
}

func proxyGet(proxy *dosproxy.DosproxySession, vType int) chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)
		var val *big.Int
		var err error
		switch vType {
		case LastRandomness:
			val, err = proxy.LastRandomness()
		case GroupSize:
			val, err = proxy.GroupSize()
		case WorkingGroupSize:
			val, err = proxy.GetWorkingGroupSize()
		case PengindNodeSize:
			val, err = proxy.NumPendingNodes()
		case NumPendingGroups:
			val, err = proxy.NumPendingGroups()
		case GroupToPick:
			val, err = proxy.GroupToPick()
		case LastUpdatedBlock:
			val, err = proxy.LastUpdatedBlock()
		}
		if err != nil {
			logger.Error(err)
			return
		}
		out <- val
	}()
	return out
}

func (e *EthAdaptor) LastRandomness() (rand *big.Int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var valList []chan interface{}
	for _, proxy := range e.proxies {
		valList = append(valList, proxyGet(proxy, LastRandomness))
	}
	out := first(ctx, merge(ctx, valList...))
	select {
	case val := <-out:
		var ok bool
		rand, ok = val.(*big.Int)
		if !ok {
			err = errors.New("type error")
			return
		}
	case <-ctx.Done():
		err = errors.New("Timeout")
	}
	return
}

func (e *EthAdaptor) GroupSize() (size uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var valList []chan interface{}
	for _, proxy := range e.proxies {
		valList = append(valList, proxyGet(proxy, GroupSize))
	}
	out := first(ctx, merge(ctx, valList...))
	select {
	case val, ok := <-out:
		sizeBig, ok := val.(*big.Int)
		if !ok {
			err = errors.New("type error")
			return
		}
		size = sizeBig.Uint64()
	case <-ctx.Done():
		err = errors.New("Timeout")
	}
	return
}

func (e *EthAdaptor) GetWorkingGroupSize() (size uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var valList []chan interface{}
	for _, proxy := range e.proxies {
		valList = append(valList, proxyGet(proxy, WorkingGroupSize))
	}
	out := first(ctx, merge(ctx, valList...))
	select {
	case val, ok := <-out:
		sizeBig, ok := val.(*big.Int)
		if !ok {
			err = errors.New("type error")
			return
		}
		size = sizeBig.Uint64()
	case <-ctx.Done():
		err = errors.New("Timeout")
	}
	return
}

func (e *EthAdaptor) NumPendingGroups() (size uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var valList []chan interface{}
	for _, proxy := range e.proxies {
		valList = append(valList, proxyGet(proxy, NumPendingGroups))
	}
	out := first(ctx, merge(ctx, valList...))
	select {
	case val := <-out:
		sizeBig, ok := val.(*big.Int)
		if !ok {
			err = errors.New("type error")
			return
		}
		size = sizeBig.Uint64()
	case <-ctx.Done():
		err = errors.New("Timeout")
	}
	return
}

func (e *EthAdaptor) GetPengindNodeSize() (size uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var valList []chan interface{}
	for _, proxy := range e.proxies {
		valList = append(valList, proxyGet(proxy, PengindNodeSize))
	}
	out := first(ctx, merge(ctx, valList...))
	select {
	case val := <-out:
		sizeBig, ok := val.(*big.Int)
		if !ok {
			err = errors.New("type error")
			return
		}
		size = sizeBig.Uint64()
	case <-ctx.Done():
		err = errors.New("Timeout")
	}
	return
}

func (e *EthAdaptor) GetGroupToPick() (size uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var valList []chan interface{}
	for _, proxy := range e.proxies {
		valList = append(valList, proxyGet(proxy, GroupToPick))
	}
	out := first(ctx, merge(ctx, valList...))
	select {
	case val := <-out:
		sizeBig, ok := val.(*big.Int)
		if !ok {
			err = errors.New("type error")
			return
		}
		size = sizeBig.Uint64()
	case <-ctx.Done():
		err = errors.New("Timeout")
	}
	return
}

func (e *EthAdaptor) LastUpdatedBlock() (blknum uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var valList []chan interface{}
	for _, proxy := range e.proxies {
		valList = append(valList, proxyGet(proxy, LastUpdatedBlock))
	}
	out := first(ctx, merge(ctx, valList...))
	select {
	case val := <-out:
		blknumBig, ok := val.(*big.Int)
		if !ok {
			err = errors.New("type error")
			return
		}
		blknum = blknumBig.Uint64()
	case <-ctx.Done():
		err = errors.New("Timeout")
	}
	return
}

func (e *EthAdaptor) clientGet(ctx context.Context, client *ethclient.Client, vType int) chan interface{} {
	out := make(chan interface{})
	go func(client *ethclient.Client) {
		defer close(out)
		var val uint64
		var err error
		switch vType {
		case PendingNonce:
			val, err = client.PendingNonceAt(ctx, e.key.Address)
		}
		if err != nil {
			logger.Error(err)
			return
		}
		select {
		case out <- val:
		case <-ctx.Done():
		}
	}(client)
	return out
}

func (e *EthAdaptor) PendingNonce() (nonce uint64, err error) {
	var ok bool
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	var valList []chan interface{}
	for _, client := range e.clients {
		valList = append(valList, e.clientGet(ctx, client, PendingNonce))
	}
	out := first(ctx, merge(ctx, valList...))
	select {
	case val := <-out:
		nonce, ok = val.(uint64)
		if !ok {
			err = errors.New("type error")
		}
	case <-ctx.Done():
		err = errors.New("Timeout")
	}
	return
}

//TODO move this to eth_helper and add First/Merge/proxyGet in here
func (e *EthAdaptor) CurrentBlock() (blknum uint64, err error) {
	var header *types.Header
	header, err = e.clients[0].HeaderByNumber(e.ctx, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	blknum = header.Number.Uint64()
	return
}

func (e *EthAdaptor) GroupPubKey(idx int) (groupPubKeys [4]*big.Int, err error) {
	return e.proxies[0].GetGroupPubKey(big.NewInt(int64(idx)))
}

//TODO move this to eth_helper and add First/Merge/proxyGet in here
func (e *EthAdaptor) GetBalance() (balance *big.Float) {
	return GetBalance(e.clients[0], e.key)
}

func (e *EthAdaptor) Address() (addr []byte) {
	return e.key.Address.Bytes()
}

func DecodeSig(sig []byte) (x, y *big.Int) {
	x = new(big.Int)
	y = new(big.Int)
	x.SetBytes(sig[0:32])
	y.SetBytes(sig[32:])
	return
}
