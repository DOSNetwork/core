package onchain

import (
	"context"
	//	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/ethereum/go-ethereum/accounts/keystore"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SubscribeDOSProxyLogUpdateRandom = iota
	SubscribeDOSProxyLogRequestUserRandom
	SubscribeDOSProxyLogUrl
	SubscribeDOSProxyLogValidationResult
	SubscribeDOSProxyLogGrouping
	SubscribeDOSProxyLogPublicKeyAccepted
	SubscribeDOSProxyLogPublicKeySuggested
	SubscribeDOSProxyLogGroupDismiss
	SubscribeDOSProxyLogInsufficientPendingNode
	SubscribeDOSProxyLogInsufficientWorkingGroup
)

var logger log.Logger

// TODO: Move constants to some unified places.
const (
	TrafficSystemRandom = iota // 0
	TrafficUserRandom
	TrafficUserQuery
)

const (
	LogBlockDiff        = 2
	LogCheckingInterval = 15 //in second
	SubscribeTimeout    = 60 //in second
)

type EthAdaptor struct {
	s        *dosproxy.DOSProxySession
	c        *ethclient.Client
	key      *keystore.Key
	ctx      context.Context
	cancel   context.CancelFunc
	wOpts    *bind.WatchOpts
	reqQueue chan interface{}
}

type Request struct {
	ctx    context.Context
	f      func() (*types.Transaction, error)
	result chan Reply
}

type ReqGrouping struct {
	ctx        context.Context
	candidates []common.Address
	size       *big.Int
	f          func([]common.Address, *big.Int) (*types.Transaction, error)
	result     chan Reply
}

type ReqSetRandomNum struct {
	ctx     context.Context
	sig     [2]*big.Int
	version uint8
	f       func([2]*big.Int, uint8) (*types.Transaction, error)
	result  chan Reply
}

type ReqSetPublicKey struct {
	ctx     context.Context
	groupId *big.Int
	pubKey  [4]*big.Int
	f       func(*big.Int, [4]*big.Int) (*types.Transaction, error)
	result  chan Reply
}

type ReqTriggerCallback struct {
	ctx         context.Context
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

func NewETHProxySession(credentialPath, passphrase, proxyAddr string, gethUrls []string) (adaptor *EthAdaptor, err error) {

	if logger == nil {
		logger = log.New("module", "EthProxy")
	}

	key, err := ReadEthKey(credentialPath, passphrase)
	if err != nil {
		logger.Error(err)
		return
	}

	clients, errs := DialToEth(context.Background(), gethUrls)
	go func() {
		for err := range errs {
			logger.Error(err)
		}
	}()

	//Use first client
	c, ok := <-clients
	if !ok {
		err = errors.New("No any working eth client")
		return
	}
	for client := range clients {
		client.Close()
	}

	p, err := dosproxy.NewDOSProxy(common.HexToAddress(proxyAddr), c)
	if err != nil {
		logger.Error(err)
		return
	}

	adaptor = &EthAdaptor{}
	ctx, cancelSession := context.WithCancel(context.Background())
	auth := bind.NewKeyedTransactor(key.PrivateKey)
	auth.GasLimit = uint64(GASLIMIT)
	auth.Context = ctx
	nonce, _ := c.PendingNonceAt(ctx, key.Address)
	auth.Nonce = big.NewInt(0)
	auth.Nonce = auth.Nonce.SetUint64(nonce)
	adaptor.wOpts = &bind.WatchOpts{Context: ctx}
	adaptor.s = &dosproxy.DOSProxySession{Contract: p, CallOpts: bind.CallOpts{Context: ctx}, TransactOpts: *auth}
	adaptor.c = c
	adaptor.key = key
	adaptor.ctx = ctx
	adaptor.cancel = cancelSession
	adaptor.reqQueue = make(chan interface{})

	adaptor.reqHandle()

	return
}

func (e *EthAdaptor) reqHandle() {
	go func() {
		defer fmt.Println("reqHandle exit")
		defer close(e.reqQueue)

		for {
			select {
			case req := <-e.reqQueue:
				var tx *types.Transaction
				var err error
				var resultC chan Reply
				var ctx context.Context
				reply := Reply{}

				//Compare with latest pending nonce
				nonce, err := e.c.PendingNonceAt(e.ctx, e.key.Address)
				if err == nil {
					reply.err = err
					nonceBig := big.NewInt(0)
					nonceBig = nonceBig.SetUint64(nonce)
					if e.s.TransactOpts.Nonce.Cmp(nonceBig) == -1 {
						e.s.TransactOpts.Nonce = nonceBig
					}
				}

				switch content := req.(type) {
				case *Request:
					tx, err = content.f()
					resultC = content.result
					ctx = content.ctx
				case *ReqGrouping:
					tx, err = content.f(content.candidates, content.size)
					resultC = content.result
					ctx = content.ctx
				case *ReqSetRandomNum:
					tx, err = content.f(content.sig, content.version)
					resultC = content.result
					ctx = content.ctx
				case *ReqTriggerCallback:
					tx, err = content.f(content.requestId, content.trafficType, content.content, content.sig, content.version)
					resultC = content.result
					ctx = content.ctx
				case *ReqSetPublicKey:
					tx, err = content.f(content.groupId, content.pubKey)
					resultC = content.result
					ctx = content.ctx
				}
				if err != nil {
					if err.Error() == "replacement transaction underpriced" ||
						strings.Contains(err.Error(), "known transaction") {
						e.s.TransactOpts.Nonce = e.s.TransactOpts.Nonce.Add(e.s.TransactOpts.Nonce, big.NewInt(1))
					}
					reply.err = err
					resultC <- reply
					continue
				}
				//fmt.Println(reqType, " nonce ", e.s.TransactOpts.Nonce, " size ", size, " tx ", fmt.Sprintf("%x", tx.Hash()))
				reply.tx = tx
				reply.nonce = e.s.TransactOpts.Nonce.Uint64()
				e.s.TransactOpts.Nonce = e.s.TransactOpts.Nonce.Add(e.s.TransactOpts.Nonce, big.NewInt(1))
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

func (e *EthAdaptor) End() {
	e.cancel()
	e.c.Close()

	return
}
func (e *EthAdaptor) sendRequest(ctx context.Context, request interface{}, result chan Reply) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		defer close(result)
		//The same account might be used outside of system that causes a nonce conflict.
		//Need to check if tx is confirmed to avoid lost transaction because of nonce conflict
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
					if err.Error() == "client is closed" ||
						err.Error() == "EOF" ||
						strings.Contains(err.Error(), "connection refused") ||
						strings.Contains(err.Error(), "use of closed network connection") {
						errc <- err
						return
					}
					continue
				}
				err = CheckTransaction(e.c, tx)
				if err != nil {
					if err.Error() == "client is closed" ||
						err.Error() == "EOF" ||
						strings.Contains(err.Error(), "connection refused") ||
						strings.Contains(err.Error(), "use of closed network connection") ||
						strings.Contains(err.Error(), "transaction failed") {
						errc <- err
						return
					}
					continue
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
	result := make(chan Reply)
	request := &Request{ctx, e.s.RegisterNewNode, result}
	return e.sendRequest(ctx, request, result)
}

/*
func (e *EthAdaptor) ResetNodeIDs(ctx context.Context) (errc <-chan error) {
	result := make(chan Reply)
	request := &Request{ctx, e.s.ResetContract, result}
	return e.sendRequest(ctx, request, result)
}
*/
func (e *EthAdaptor) RandomNumberTimeOut(ctx context.Context) (errc <-chan error) {
	result := make(chan Reply)
	request := &Request{ctx, e.s.HandleTimeout, result}
	return e.sendRequest(ctx, request, result)
}

func (e *EthAdaptor) RegisterGroupPubKey(ctx context.Context, IdWithPubKeys chan [5]*big.Int) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case IdWithPubKey, ok := <-IdWithPubKeys:
			if !ok {
				return
			}
			result := make(chan Reply)
			groupId := IdWithPubKey[0]
			var pubKey [4]*big.Int
			copy(pubKey[:], IdWithPubKey[1:])
			request := &ReqSetPublicKey{ctx, groupId, pubKey, e.s.RegisterGroupPubKey, result}
			errChan := e.sendRequest(ctx, request, result)
			err := <-errChan
			errc <- err
			return
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

func (e *EthAdaptor) SetRandomNum(ctx context.Context, signatures <-chan *vss.Signature) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			x, y := DecodeSig(signature.Signature)
			result := make(chan Reply)
			request := &ReqSetRandomNum{ctx, [2]*big.Int{x, y}, 0, e.s.UpdateRandomness, result}
			errChan := e.sendRequest(ctx, request, result)
			err := <-errChan
			errc <- err
			return
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

func (e *EthAdaptor) DataReturn(ctx context.Context, signatures <-chan *vss.Signature) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			x, y := DecodeSig(signature.Signature)
			requestId := new(big.Int).SetBytes(signature.RequestId)

			result := make(chan Reply)
			request := &ReqTriggerCallback{ctx, requestId, uint8(signature.Index), signature.Content, [2]*big.Int{x, y}, 0, e.s.TriggerCallback, result}
			errChan := e.sendRequest(ctx, request, result)
			err := <-errChan
			errc <- err
			return
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

/*
func (e *EthAdaptor) Grouping(ctx context.Context, size int) <-chan error {
	result := make(chan Reply)
	request := &ReqGrouping{ctx, nil, big.NewInt(int64(size)), e.s.Grouping, result}
	return e.sendRequest(ctx, request, result)
}
*/
func (e *EthAdaptor) PollLogs(subscribeType int, sink chan interface{}) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		targetBlockN, err := GetCurrentBlock(e.c)
		if err != nil {
			errc <- err
			return
		}
		timer := time.NewTimer(LogCheckingInterval * time.Second)
		for {
			select {
			case <-timer.C:
				currentBlockN, err := GetCurrentBlock(e.c)
				if err != nil {
					errc <- err
					return
				}
				for ; currentBlockN-LogBlockDiff >= targetBlockN; targetBlockN++ {
					switch subscribeType {
					case SubscribeDOSProxyLogGrouping:
						logs, err := e.s.Contract.DOSProxyFilterer.FilterLogGrouping(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.ctx,
						})
						if err != nil {
							errc <- err
							return
						}
						for logs.Next() {
							sink <- &DOSProxyLogGrouping{
								GroupId: logs.Event.GroupId,
								NodeId:  logs.Event.NodeId,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
							}
						}
					case SubscribeDOSProxyLogGroupDismiss:
						logs, err := e.s.Contract.DOSProxyFilterer.FilterLogGroupDismiss(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: e.ctx,
						})
						if err != nil {
							errc <- err
							return
						}
						for logs.Next() {
							sink <- &DOSProxyLogGroupDismiss{
								PubKey:  logs.Event.PubKey,
								GroupId: logs.Event.GroupId,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
							}
						}
					}
				}
				timer.Reset(LogCheckingInterval * time.Second)
			case <-e.ctx.Done():
				return
			}
		}
	}()
	return errc
}

func (e *EthAdaptor) SubscribeEvent(subscribeType int, sink chan interface{}) chan error {
	errc := make(chan error)
	switch subscribeType {
	case SubscribeDOSProxyLogUpdateRandom:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogUpdateRandom)
			defer close(transitChan)
			defer close(errc)
			sub, err := e.s.Contract.DOSProxyFilterer.WatchLogUpdateRandom(e.wOpts, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-e.ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					sink <- &DOSProxyLogUpdateRandom{
						LastRandomness:    i.LastRandomness,
						DispatchedGroupId: i.DispatchedGroupId,
						DispatchedGroup:   i.DispatchedGroup,
						Tx:                i.Raw.TxHash.Hex(),
						BlockN:            i.Raw.BlockNumber,
						Removed:           i.Raw.Removed,
					}
				}
			}
		}()
	case SubscribeDOSProxyLogUrl:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogUrl)
			defer close(transitChan)
			defer close(errc)
			sub, err := e.s.Contract.DOSProxyFilterer.WatchLogUrl(e.wOpts, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-e.ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					sink <- &DOSProxyLogUrl{
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
					}
				}
			}
		}()
	case SubscribeDOSProxyLogRequestUserRandom:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogRequestUserRandom)
			defer close(transitChan)
			defer close(errc)
			sub, err := e.s.Contract.DOSProxyFilterer.WatchLogRequestUserRandom(e.wOpts, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-e.ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					sink <- &DOSProxyLogRequestUserRandom{
						RequestId:            i.RequestId,
						LastSystemRandomness: i.LastSystemRandomness,
						UserSeed:             i.UserSeed,
						DispatchedGroupId:    i.DispatchedGroupId,
						DispatchedGroup:      i.DispatchedGroup,
						Tx:                   i.Raw.TxHash.Hex(),
						BlockN:               i.Raw.BlockNumber,
						Removed:              i.Raw.Removed,
					}
				}
			}
		}()
	case SubscribeDOSProxyLogValidationResult:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogValidationResult)
			defer close(transitChan)
			defer close(errc)
			sub, err := e.s.Contract.DOSProxyFilterer.WatchLogValidationResult(e.wOpts, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-e.ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					sink <- &DOSProxyLogValidationResult{
						TrafficType: i.TrafficType,
						TrafficId:   i.TrafficId,
						Message:     i.Message,
						Signature:   i.Signature,
						PubKey:      i.PubKey,
						GroupId:     i.GroupId,
						Pass:        i.Pass,
						Version:     i.Version,
						Tx:          i.Raw.TxHash.Hex(),
						BlockN:      i.Raw.BlockNumber,
						Removed:     i.Raw.Removed,
					}
				}
			}
		}()
	case SubscribeDOSProxyLogInsufficientPendingNode:
		go func() {
			transitChan := make(chan *dosproxy.DOSProxyLogInsufficientPendingNode)
			defer close(transitChan)
			defer close(errc)
			sub, err := e.s.Contract.DOSProxyFilterer.WatchLogInsufficientPendingNode(e.wOpts, transitChan)
			if err != nil {
				return
			}
			for {
				select {
				case <-e.ctx.Done():
					sub.Unsubscribe()
					return
				case err := <-sub.Err():
					errc <- err
					return
				case i := <-transitChan:
					sink <- &DOSProxyLogInsufficientPendingNode{
						NumPendingNodes: i.NumPendingNodes,
						Tx:              i.Raw.TxHash.Hex(),
						BlockN:          i.Raw.BlockNumber,
						Removed:         i.Raw.Removed,
					}
				}
			}
		}()
	case SubscribeDOSProxyLogInsufficientWorkingGroup:
	}
	return errc
}

func (e *EthAdaptor) LastRandomness() (rand *big.Int, err error) {
	rand, err = e.s.LastRandomness()
	return
}

func (e *EthAdaptor) LastUpdatedBlock() (blknum uint64, err error) {
	lastBlk, err := e.s.LastUpdatedBlock()
	blknum = lastBlk.Uint64()
	return
}

func (e *EthAdaptor) GroupPubKey(idx int) (groupPubKeys [4]*big.Int, err error) {
	return e.s.GetGroupPubKey(big.NewInt(int64(idx)))
}

func (e *EthAdaptor) CurrentBlock() (blknum uint64, err error) {
	var header *types.Header
	header, err = e.c.HeaderByNumber(e.ctx, nil)
	if err == nil {
		blknum = header.Number.Uint64()
	}
	return
}

func (e *EthAdaptor) GetBalance() (balance *big.Float) {
	return GetBalance(e.c, e.key)
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
