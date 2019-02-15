package onchain

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/DOSNetwork/core/share/vss/pedersen"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	SubscribeDOSProxyLogUrl = iota
	SubscribeDOSProxyLogRequestUserRandom
	SubscribeDOSProxyLogNonSupportedType
	SubscribeDOSProxyLogNonContractCall
	SubscribeDOSProxyLogCallbackTriggeredFor
	SubscribeDOSProxyLogQueryFromNonExistentUC
	SubscribeDOSProxyLogUpdateRandom
	SubscribeDOSProxyLogValidationResult
	SubscribeDOSProxyLogInsufficientGroupNumber
	SubscribeDOSProxyLogGrouping
	SubscribeDOSProxyLogDuplicatePubKey
	SubscribeDOSProxyLogAddressNotFound
	SubscribeDOSProxyLogPublicKeyAccepted
	SubscribeDOSProxyLogGroupDismiss
	SubscribeDOSProxyWhitelistAddressTransferred
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

type EthAdaptor struct {
	EthCommon
	proxy     *dosproxy.DOSProxy
	logFilter *sync.Map
	logger    log.Logger
}

func (e *EthAdaptor) Init(config configuration.ChainConfig) (err error) {
	e.logger = log.New("module", "EthProxy")
	if err = e.EthCommon.Init(config); err != nil {
		e.logger.Error(err)
		return
	}

	e.logFilter = new(sync.Map)
	go e.logMapTimeout()

	fmt.Println("onChainConn initialization finished.")
	err = e.dialToProxy()
	return
}

func (e *EthAdaptor) SubscribeEvent(ch chan interface{}, subscribeType int) (err error) {
	opt := &bind.WatchOpts{}
	var cancel context.CancelFunc
	opt.Context, cancel = context.WithCancel(context.Background())
	done := make(chan bool)
	timer := time.NewTimer(SubscribeTimeout * time.Second)

	go e.subscribeEventAttempt(ch, opt, subscribeType, done)

	for {
		select {
		case succ := <-done:
			if succ {
				fmt.Println("subscribe done")
				return
			} else {
				fmt.Println("retry...")
				if err = e.resetOnChainConn(); err != nil {
					return
				} else {
					go e.subscribeEventAttempt(ch, opt, subscribeType, done)
				}
			}
		case <-timer.C:
			cancel()
			fmt.Println("subscribe timeout")
			return
		}
	}
}

func (e *EthAdaptor) dialToProxy() (err error) {
	addr := common.HexToAddress(e.config.DOSProxyAddress)
	e.proxy, err = dosproxy.NewDOSProxy(addr, e.Client)
	for err != nil {
		e.logger.Error(err)
		e.proxy, err = dosproxy.NewDOSProxy(addr, e.Client)
	}
	return
}

func (e *EthAdaptor) resetOnChainConn() (err error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if err = e.EthCommon.DialToEth(); err != nil {
		e.logger.Error(err)
		return
	}
	err = e.dialToProxy()
	return
}

func (e *EthAdaptor) fetchMatureLogs(ctx context.Context, subscribeType int, ch chan interface{}) (err error) {
	targetBlockN, err := e.GetCurrentBlock()
	if err != nil {
		return
	}

	duplicates := make(map[string]struct{})

	timer := time.NewTimer(LogCheckingInterval * time.Second)
	go func() {
		for {
			select {
			case <-timer.C:
				currentBlockN, err := e.GetCurrentBlock()
				if err != nil {
					e.logger.Error(err)
				}
				for ; currentBlockN-LogBlockDiff >= targetBlockN; targetBlockN++ {
					switch subscribeType {
					case SubscribeDOSProxyLogGrouping:
						logs, err := e.proxy.DOSProxyFilterer.FilterLogGrouping(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							e.logger.Error(err)
						}
						for logs.Next() {
							ch <- &DOSProxyLogGrouping{
								NodeId:  logs.Event.NodeId,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
							}

							f := map[string]interface{}{
								"logName": "Grouping",
								"Removed": logs.Event.Raw.Removed,
								"Tx":      logs.Event.Raw.TxHash.String(),
								"BlockN":  logs.Event.Raw.BlockNumber}

							var toHash []byte
							for _, add := range logs.Event.NodeId {
								toHash = append(toHash, add.Bytes()...)
							}
							toHash = append(toHash, logs.Event.Raw.TxHash.Bytes()...)
							toHash = append(toHash, logs.Event.Raw.Address.Bytes()...)
							toHash = append(toHash, logs.Event.Raw.Data...)
							toHash = append(toHash, logs.Event.Raw.BlockHash.Bytes()...)

							hashBytes := sha256.Sum256(toHash)

							if _, loaded := duplicates[string(hashBytes[:])]; loaded {
								f["duplicate"] = true
							} else {
								duplicates[string(hashBytes[:])] = struct{}{}
								f["duplicate"] = false
							}
							e.logger.Event("FetchMatureLog", f)
						}
						if logs.Error() != nil {
							e.logger.Error(logs.Error())
						}
						if logs.Close() != nil {
							e.logger.Error(logs.Error())
						}
					case SubscribeDOSProxyLogGroupDismiss:
						logs, err := e.proxy.DOSProxyFilterer.FilterLogGroupDismiss(&bind.FilterOpts{
							Start:   targetBlockN,
							End:     &targetBlockN,
							Context: ctx,
						})
						if err != nil {
							e.logger.Error(err)
						}
						for logs.Next() {
							ch <- &DOSProxyLogGroupDismiss{
								PubKey:  logs.Event.PubKey,
								Tx:      logs.Event.Raw.TxHash.Hex(),
								BlockN:  logs.Event.Raw.BlockNumber,
								Removed: logs.Event.Raw.Removed,
							}

							f := map[string]interface{}{
								"logName": "Dismissing",
								"Removed": logs.Event.Raw.Removed,
								"Tx":      logs.Event.Raw.TxHash.String(),
								"BlockN":  logs.Event.Raw.BlockNumber}

							var toHash []byte
							for _, add := range logs.Event.PubKey {
								toHash = append(toHash, add.Bytes()...)
							}
							toHash = append(toHash, logs.Event.Raw.TxHash.Bytes()...)
							toHash = append(toHash, logs.Event.Raw.Address.Bytes()...)
							toHash = append(toHash, logs.Event.Raw.Data...)
							toHash = append(toHash, logs.Event.Raw.BlockHash.Bytes()...)

							hashBytes := sha256.Sum256(toHash)

							if _, loaded := duplicates[string(hashBytes[:])]; loaded {
								f["duplicate"] = true
							} else {
								duplicates[string(hashBytes[:])] = struct{}{}
								f["duplicate"] = false
							}
							e.logger.Event("FetchMatureLog", f)
						}
						if logs.Error() != nil {
							e.logger.Error(logs.Error())
						}
						if logs.Close() != nil {
							e.logger.Error(logs.Error())
						}
					}
				}
				timer.Reset(LogCheckingInterval * time.Second)
			case <-ctx.Done():
				return
			}
		}
	}()
	return
}

func (e *EthAdaptor) subscribeEventAttempt(ch chan interface{}, opt *bind.WatchOpts, subscribeType int, done chan bool) {
	fmt.Println("attempt to subscribe event...")
	switch subscribeType {
	case SubscribeDOSProxyLogGrouping:
		if err := e.fetchMatureLogs(opt.Context, SubscribeDOSProxyLogGrouping, ch); err != nil {
			done <- false
			e.logger.Error(err)
			fmt.Println("Network fail, will retry shortly")
			return
		} else {
			done <- true
		}
	case SubscribeDOSProxyLogUrl:
		fmt.Println("subscribing DOSProxyLogUrl event...")
		transitChan := make(chan *dosproxy.DOSProxyLogUrl)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogUrl(opt, transitChan)
		if err != nil {
			done <- false
			e.logger.Error(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogUrl event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				e.logger.Error(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogUrl{
					QueryId:         i.QueryId,
					Timeout:         i.Timeout,
					DataSource:      i.DataSource,
					Selector:        i.Selector,
					Randomness:      i.Randomness,
					DispatchedGroup: i.DispatchedGroup,
					Tx:              i.Raw.TxHash.Hex(),
					BlockN:          i.Raw.BlockNumber,
					Removed:         i.Raw.Removed,
				}

				ch <- &DOSProxyLogUrl{
					QueryId:         i.QueryId,
					Timeout:         i.Timeout,
					DataSource:      i.DataSource,
					Selector:        i.Selector,
					Randomness:      i.Randomness,
					DispatchedGroup: i.DispatchedGroup,
					Tx:              i.Raw.TxHash.Hex(),
					BlockN:          i.Raw.BlockNumber,
					Removed:         true,
				}

				ch <- &DOSProxyLogUrl{
					QueryId:         i.QueryId,
					Timeout:         i.Timeout,
					DataSource:      i.DataSource,
					Selector:        i.Selector,
					Randomness:      i.Randomness,
					DispatchedGroup: i.DispatchedGroup,
					Tx:              i.Raw.TxHash.Hex(),
					BlockN:          i.Raw.BlockNumber,
					Removed:         i.Raw.Removed,
				}
			}
		}

	case SubscribeDOSProxyLogRequestUserRandom:
		fmt.Println("subscribing DOSProxyLogRequestUserRandom event...")
		transitChan := make(chan *dosproxy.DOSProxyLogRequestUserRandom)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogRequestUserRandom(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchLogRequestUserRandom").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogRequestUserRandom event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogRequestUserRandom{
					RequestId:            i.RequestId,
					LastSystemRandomness: i.LastSystemRandomness,
					UserSeed:             i.UserSeed,
					DispatchedGroup:      i.DispatchedGroup,
					Tx:                   i.Raw.TxHash.Hex(),
					BlockN:               i.Raw.BlockNumber,
					Removed:              i.Raw.Removed,
				}

				ch <- &DOSProxyLogRequestUserRandom{
					RequestId:            i.RequestId,
					LastSystemRandomness: i.LastSystemRandomness,
					UserSeed:             i.UserSeed,
					DispatchedGroup:      i.DispatchedGroup,
					Tx:                   i.Raw.TxHash.Hex(),
					BlockN:               i.Raw.BlockNumber,
					Removed:              true,
				}

				ch <- &DOSProxyLogRequestUserRandom{
					RequestId:            i.RequestId,
					LastSystemRandomness: i.LastSystemRandomness,
					UserSeed:             i.UserSeed,
					DispatchedGroup:      i.DispatchedGroup,
					Tx:                   i.Raw.TxHash.Hex(),
					BlockN:               i.Raw.BlockNumber,
					Removed:              i.Raw.Removed,
				}
			}
		}
	case SubscribeDOSProxyLogUpdateRandom:
		fmt.Println("subscribing DOSProxyLogUpdateRandom event...")
		transitChan := make(chan *dosproxy.DOSProxyLogUpdateRandom)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogUpdateRandom(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchLogRequestUserRandom").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogUpdateRandom event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				ch <- &DOSProxyLogUpdateRandom{
					LastRandomness:  i.LastRandomness,
					DispatchedGroup: i.DispatchedGroup,
					Tx:              i.Raw.TxHash.Hex(),
					BlockN:          i.Raw.BlockNumber,
					Removed:         i.Raw.Removed,
				}

				ch <- &DOSProxyLogUpdateRandom{
					LastRandomness:  i.LastRandomness,
					DispatchedGroup: i.DispatchedGroup,
					Tx:              i.Raw.TxHash.Hex(),
					BlockN:          i.Raw.BlockNumber,
					Removed:         true,
				}

				ch <- &DOSProxyLogUpdateRandom{
					LastRandomness:  i.LastRandomness,
					DispatchedGroup: i.DispatchedGroup,
					Tx:              i.Raw.TxHash.Hex(),
					BlockN:          i.Raw.BlockNumber,
					Removed:         i.Raw.Removed,
				}
			}
		}
	case SubscribeDOSProxyLogValidationResult:
		fmt.Println("subscribing SubscribeDOSProxyLogValidationResult event...")
		transitChan := make(chan *dosproxy.DOSProxyLogValidationResult)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogValidationResult(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchLogValidationResult").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("SubscribeDOSProxyLogValidationResult event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				//if i.Raw.Removed == false {
				ch <- &DOSProxyLogValidationResult{
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
				}
				//}
			}
		}
	case SubscribeDOSProxyLogNonSupportedType:
		fmt.Println("subscribing DOSProxyLogNonSupportedType event...")
		transitChan := make(chan *dosproxy.DOSProxyLogNonSupportedType)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogNonSupportedType(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchLogNonSupportedType").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogNonSupportedType event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogNonSupportedType{
						InvalidSelector: i.InvalidSelector,
					}
				}
			}
		}
	case SubscribeDOSProxyLogNonContractCall:
		fmt.Println("subscribing DOSProxyLogNonContractCall event...")
		transitChan := make(chan *dosproxy.DOSProxyLogNonContractCall)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogNonContractCall(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchLogNonContractCall").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogNonContractCall event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogNonContractCall{
						From: i.From,
					}
				}
			}
		}
	case SubscribeDOSProxyLogCallbackTriggeredFor:
		fmt.Println("subscribing DOSProxyLogCallbackTriggeredFor event...")
		transitChan := make(chan *dosproxy.DOSProxyLogCallbackTriggeredFor)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogCallbackTriggeredFor(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchLogCallbackTriggeredFor").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogCallbackTriggeredFor event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogCallbackTriggeredFor{
						CallbackAddr: i.CallbackAddr,
					}
				}
			}
		}
	case SubscribeDOSProxyLogQueryFromNonExistentUC:
		fmt.Println("subscribing DOSProxyLogQueryFromNonExistentUC event...")
		transitChan := make(chan *dosproxy.DOSProxyLogRequestFromNonExistentUC)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogRequestFromNonExistentUC(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchLogRequestFromNonExistentUC").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogQueryFromNonExistentUC event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogRequestFromNonExistentUC{}
				}
			}
		}
	case SubscribeDOSProxyLogInsufficientGroupNumber:
		fmt.Println("subscribing DOSProxyLogInsufficientGroupNumber event...")
		transitChan := make(chan *dosproxy.DOSProxyLogInsufficientGroupNumber)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogInsufficientGroupNumber(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchLogInsufficientGroupNumber").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogInsufficientGroupNumber event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogInsufficientGroupNumber{}
				}
			}
		}
	case SubscribeDOSProxyLogDuplicatePubKey:
		fmt.Println("subscribing DOSProxyLogDuplicatePubKey event...")
		transitChan := make(chan *dosproxy.DOSProxyLogDuplicatePubKey)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogDuplicatePubKey(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "WatchLogDuplicatePubKey").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogDuplicatePubKey event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogDuplicatePubKey{
						PubKey: i.PubKey,
					}
				}
			}
		}
	case SubscribeDOSProxyLogAddressNotFound:
		fmt.Println("subscribing DOSProxyLogAddressNotFound event...")
		transitChan := make(chan *dosproxy.DOSProxyLogAddressNotFound)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogAddressNotFound(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "WatchLogAddressNotFound").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogAddressNotFound event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogAddressNotFound{
						PubKey: i.PubKey,
					}
				}
			}
		}
	case SubscribeDOSProxyLogPublicKeyAccepted:
		fmt.Println("subscribing DOSProxyLogPublicKeyAccepted event...")
		transitChan := make(chan *dosproxy.DOSProxyLogPublicKeyAccepted)
		sub, err := e.proxy.DOSProxyFilterer.WatchLogPublicKeyAccepted(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchLogPublicKeyAccepted").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyLogPublicKeyAccepted event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyLogPublicKeyAccepted{
						PubKey: i.PubKey,
					}
				}
			}
		}
	case SubscribeDOSProxyLogGroupDismiss:
		if err := e.fetchMatureLogs(opt.Context, SubscribeDOSProxyLogGroupDismiss, ch); err != nil {
			done <- false
			e.logger.Error(err)
			fmt.Println("Network fail, will retry shortly")
			return
		} else {
			done <- true
		}
	case SubscribeDOSProxyWhitelistAddressTransferred:
		fmt.Println("subscribing DOSProxyWhitelistAddressTransferred event...")
		transitChan := make(chan *dosproxy.DOSProxyWhitelistAddressTransferred)
		sub, err := e.proxy.DOSProxyFilterer.WatchWhitelistAddressTransferred(opt, transitChan)
		if err != nil {
			done <- false
			//log.WithField("function", "watchWhitelistAddressTransferred").Warn(err)
			fmt.Println("Network fail, will retry shortly")
			return
		}

		fmt.Println("DOSProxyWhitelistAddressTransferred event subscribed")
		done <- true
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case i := <-transitChan:
				if !e.filterLog(i.Raw) {
					ch <- &DOSProxyWhitelistAddressTransferred{
						Previous: i.Previous,
						Curr:     i.Curr,
					}
				}
			}
		}
	}

}

func (e *EthAdaptor) InitialWhiteList() (err error) {
	fmt.Println("Starting initialing WhiteList...")
	auth, err := e.GetAuth()
	if err != nil {
		return
	}

	addresses := [21]common.Address{
		common.HexToAddress("0xaec1f213677de24842a96e72fe6efbcbc2b77ca5"),
		common.HexToAddress("0x5a3582d7c8fc97c194168342f9a347e8d41b4038"),
		common.HexToAddress("0xe38497ec6d9442413b16a2a7055649ac91106d79"),
		common.HexToAddress("0xb8c228a842390595751b1a09a2c4ffaccf12c44b"),
		common.HexToAddress("0xec96405d35c3cadfd3e0bda90c66b74b159d4156"),
		common.HexToAddress("0x0a65d5dccc87ecb21bc4b24c16a3c01e0cdd42ac"),
		common.HexToAddress("0x3ebe227e9fd42bb97b9a950e4a731d8975263812"),
		common.HexToAddress("0x6ca3ee1386f7c05d886211a6378f49bdf9c7ee88"),
		common.HexToAddress("0x234ae33713afde52de9aa9203fb6696531edc74f"),
		common.HexToAddress("0xe610c52cbeb14dc722a9668e59bad46c5e464e55"),
		common.HexToAddress("0xcb04aae925218094863809ec0289a8fdccfd68cf"),
		common.HexToAddress("0x921f2cf348b8b45d6cd5eaf139d30303e6b9646f"),
		common.HexToAddress("0x69ba6867602d650fc433fd62eabaf17f11fd5132"),
		common.HexToAddress("0x4ed2814cd63e83504221424215d0655c6db0b674"),
		common.HexToAddress("0xac3e1e84e3b7a0a83d36feba984a836c768fcb72"),
		common.HexToAddress("0x3609aa202ab8b96499e379da7226145e5697695e"),
		common.HexToAddress("0xc96c0f2d346a3f0f4bbcfbefe70a75c710d23369"),
		common.HexToAddress("0xc1a3aab78a6dc3e5cbeec059a2727a66f4bc7088"),
		common.HexToAddress("0xc98a4df797f0b8155ef38c4701588cbed21a1b26"),
		common.HexToAddress("0x34d950db8e9345a638ba0bee9945d56c9f7728ee"),
		common.HexToAddress("0x96272c390ae674d3a3e3f1d636f3ae4128afd688")}

	tx, err := e.proxy.InitWhitelist(auth, addresses)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		//log.WithField("function", "initWhitelist").Warn(err)
		time.Sleep(time.Second)
		tx, err = e.proxy.InitWhitelist(auth, addresses)
	}
	if err != nil {
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("Whitelist initialized")

	err = e.CheckTransaction(tx)

	return

}

func (e *EthAdaptor) WhitelistInitialized() (initialized bool, err error) {
	return e.proxy.WhitelistInitialized(&bind.CallOpts{})
}

func (e *EthAdaptor) GetGroupPubKey(idx int) (groupPubKeys [4]*big.Int, err error) {
	return e.proxy.GetGroupPubKey(&bind.CallOpts{}, big.NewInt(int64(idx)))
}

func (e *EthAdaptor) Grouping(size int) (err error) {
	fmt.Println("!!!!!!Starting Grouping ", size)
	auth, err := e.GetAuth()
	if err != nil {
		fmt.Println("!!! err ", err)
		return
	}

	tx, err := e.proxy.Grouping(auth, nil, big.NewInt(int64(size)))
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		//log.WithField("function", "grouping").Warn(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...,err ", err)
		tx, err = e.proxy.Grouping(auth, nil, big.NewInt(int64(size)))
	}
	if err != nil {
		fmt.Println("!!! err2 ", err)

		return
	}

	err = e.CheckTransaction(tx)
	return
}

func (e *EthAdaptor) FireRandom() (err error) {
	fmt.Println("!!!!!!FireRandom ")
	auth, err := e.GetAuth()
	if err != nil {
		fmt.Println("!!! err ", err)
		return
	}

	tx, err := e.proxy.FireRandom(auth)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		//log.WithField("function", "grouping").Warn(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...,err ", err)
		tx, err = e.proxy.FireRandom(auth)
	}
	if err != nil {
		fmt.Println("!!! err2 ", err)

		return
	}

	err = e.CheckTransaction(tx)
	return
}

func (e *EthAdaptor) GetWhitelist() (address common.Address, err error) {
	return e.proxy.GetWhitelistAddress(&bind.CallOpts{}, big.NewInt(1))
}

func (e *EthAdaptor) UploadID() (err error) {
	fmt.Println("Starting submitting nodeId...")
	auth, err := e.GetAuth()
	if err != nil {
		fmt.Println("GetAuth() error")
		e.logger.Error(err)
		return
	}

	tx, err := e.proxy.UploadNodeId(auth)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		//log.WithField("function", "uploadNodeId").Warn(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = e.proxy.UploadNodeId(auth)
	}
	if err != nil {
		fmt.Println("UploadNodeId error", err)
		e.logger.Error(err)
		return
	}

	fmt.Println("tx sent: ", tx.Hash().Hex())
	fmt.Println("NodeId submitted")

	//err = e.CheckTransaction(tx)

	return
}

func (e *EthAdaptor) GetId() (id []byte) {
	return e.GetAddress().Bytes()
}

func (e *EthAdaptor) GetBlockHashByNumber(blknum *big.Int) (hash common.Hash, err error) {
	block, err := e.Client.BlockByNumber(context.Background(), blknum)
	if err != nil {
		return
	}

	hash = block.Hash()
	return
}

func (e *EthAdaptor) GetLastRandomness() (rand *big.Int, err error) {
	rand, err = e.proxy.LastRandomness(nil)
	return
}

func (e *EthAdaptor) GetLastUpdatedBlock() (blknum uint64, err error) {
	lastBlk, err := e.proxy.LastUpdatedBlock(nil)
	blknum = lastBlk.Uint64()
	return
}

func (e *EthAdaptor) GetCurrentBlock() (blknum uint64, err error) {
	var header *types.Header
	header, err = e.Client.HeaderByNumber(context.Background(), nil)
	if err == nil {
		blknum = header.Number.Uint64()
	}
	return
}

func (e *EthAdaptor) SetRandomNum(ctx context.Context, signatures <-chan *vss.Signature) <-chan error {
	errc := make(chan error, 1)
	go func() {
		defer close(errc)
		fmt.Println("Starting submitting random number...")
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			defer e.logger.TimeTrack(time.Now(), "SetRandomNum", map[string]interface{}{"RequestId": ctx.Value("RequestID")})
			auth, err := e.GetAuth()
			if err != nil {
				fmt.Println("GetAuth() error")
				e.logger.Error(err)
				errc <- err
				return
			}
			x, y := DecodeSig(signature.Signature)
			tx, err := e.proxy.UpdateRandomness(auth, [2]*big.Int{x, y}, 0)
			if err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
				timer := time.NewTimer(1 * time.Second)
				retryCount := 0
			l:
				for {
					select {
					case <-timer.C:
						fmt.Println("transaction retry...")
						retryCount++
						auth, err = e.GetAuth()
						if err != nil {
							fmt.Println("GetAuth() error")
							e.logger.Error(err)
							errc <- err
							timer.Stop()
							return
						}
						tx, err = e.proxy.UpdateRandomness(auth, [2]*big.Int{x, y}, 0)
						if err == nil || (err.Error() != core.ErrNonceTooLow.Error() && err.Error() != core.ErrReplaceUnderpriced.Error()) {
							timer.Stop()
							break l
						}
						timer.Reset(1 * time.Second)
					case <-ctx.Done():
						timer.Stop()
						return
					}
				}
			}
			if err != nil {
				fmt.Println("SetRandomNum error", err)
				e.logger.Error(err)
				errc <- err
				return
			}

			fmt.Println("tx sent: ", tx.Hash().Hex())
			fmt.Println("SetRandomNum success")
			//err = e.CheckTransaction(tx)
		case <-ctx.Done():
			return
		}
	}()

	return errc
}

func (e *EthAdaptor) UploadPubKey(ctx context.Context, pubKeys chan [4]*big.Int) <-chan error {
	errc := make(chan error, 1)
	go func() {
		defer close(errc)
		fmt.Println("Starting UploadPubKey...")
		select {
		case pubKey := <-pubKeys:
			defer e.logger.TimeTrack(time.Now(), "UploadPubKey", map[string]interface{}{"SessionID": ctx.Value("SessionID")})

			auth, err := e.GetAuth()
			if err != nil {
				fmt.Println("GetAuth() error")
				e.logger.Error(err)
				errc <- err
				return
			}
			tx, err := e.proxy.SetPublicKey(auth, pubKey)
			if err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
				timer := time.NewTimer(1 * time.Second)
				retryCount := 0
			l:
				for {
					select {
					case <-timer.C:
						fmt.Println("transaction retry...")
						retryCount++
						auth, err = e.GetAuth()
						if err != nil {
							fmt.Println("GetAuth() error")
							e.logger.Error(err)
							errc <- err
							timer.Stop()
							return
						}
						tx, err = e.proxy.SetPublicKey(auth, pubKey)
						if err == nil || (err.Error() != core.ErrNonceTooLow.Error() && err.Error() != core.ErrReplaceUnderpriced.Error()) {
							timer.Stop()
							break l
						}
						timer.Reset(1 * time.Second)
					case <-ctx.Done():
						timer.Stop()
						return
					}
				}
			}
			if err != nil {
				fmt.Println("UploadPubKey error", err)
				e.logger.Error(err)
				errc <- err
				return
			}

			fmt.Println("tx sent: ", tx.Hash().Hex())
			fmt.Println("UploadPubKey success")
			//err = e.CheckTransaction(tx)
		case <-ctx.Done():
			return
		}
	}()
	return errc
}

func (e *EthAdaptor) ResetNodeIDs() (err error) {
	fmt.Println("Starting ResetNodeIDs...")
	auth, err := e.GetAuth()
	if err != nil {
		return
	}

	tx, err := e.proxy.ResetContract(auth)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		//log.WithField("function", "resetContract").Warn(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		tx, err = e.proxy.ResetContract(auth)
	}
	if err != nil {
		return
	}

	err = e.CheckTransaction(tx)
	return
}

func (e *EthAdaptor) RandomNumberTimeOut() (err error) {
	fmt.Println("Starting RandomNumberTimeOut...")
	auth, err := e.GetAuth()
	if err != nil {
		return
	}

	_, err = e.proxy.HandleTimeout(auth)
	for err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
		//log.WithField("function", "handleTimeout").Warn(err)
		time.Sleep(time.Second)
		fmt.Println("transaction retry...")
		_, err = e.proxy.HandleTimeout(auth)
	}
	if err != nil {
		return
	}

	//err = e.CheckTransaction(tx)
	return
}

func (e *EthAdaptor) DataReturn(ctx context.Context, signatures <-chan *vss.Signature) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		fmt.Println("Starting DataReturn...")
		select {
		case signature, ok := <-signatures:
			if !ok {
				return
			}
			defer e.logger.TimeTrack(time.Now(), "DataReturn", map[string]interface{}{"RequestId": ctx.Value("RequestID")})

			x, y := DecodeSig(signature.Signature)
			requestId, _ := new(big.Int).SetString(signature.QueryId, 10)

			auth, err := e.GetAuth()
			if err != nil {
				fmt.Println("GetAuth() error")
				e.logger.Error(err)
				errc <- err
				return
			}

			tx, err := e.proxy.TriggerCallback(auth, requestId, uint8(signature.Index), signature.Content, [2]*big.Int{x, y}, 0)
			if err != nil && (err.Error() == core.ErrNonceTooLow.Error() || err.Error() == core.ErrReplaceUnderpriced.Error()) {
				timer := time.NewTimer(1 * time.Second)
				retryCount := 0
			l:
				for {
					select {
					case <-timer.C:
						fmt.Println("transaction retry...")
						retryCount++
						auth, err = e.GetAuth()
						if err != nil {
							fmt.Println("GetAuth() error")
							e.logger.Error(err)
							errc <- err
							timer.Stop()
							return
						}
						tx, err = e.proxy.TriggerCallback(auth, requestId, uint8(signature.Index), signature.Content, [2]*big.Int{x, y}, 0)
						if err == nil || (err.Error() != core.ErrNonceTooLow.Error() && err.Error() != core.ErrReplaceUnderpriced.Error()) {
							timer.Stop()
							break l
						}
						timer.Reset(1 * time.Second)
					case <-ctx.Done():
						timer.Stop()
						return
					}
				}
			}
			if err != nil {
				fmt.Println("DataReturn error", err)
				e.logger.Error(err)
				errc <- err
				return
			}

			fmt.Println("tx sent: ", tx.Hash().Hex())
			fmt.Println("DataReturn success")
			//err = e.CheckTransaction(tx)
		case <-ctx.Done():
			return
		}
	}()

	return errc
}

func DecodeSig(sig []byte) (x, y *big.Int) {
	x = new(big.Int)
	y = new(big.Int)
	x.SetBytes(sig[0:32])
	y.SetBytes(sig[32:])
	return
}

func (e *EthAdaptor) SubscribeToAll(msgChan chan interface{}) (err error) {
	for i := SubscribeDOSProxyLogUrl; i <= SubscribeDOSProxyWhitelistAddressTransferred; i++ {
		err = e.SubscribeEvent(msgChan, i)
	}
	return
}

type logRecord struct {
	content       types.Log
	currTimeStamp time.Time
}

func (e *EthAdaptor) filterLog(raw types.Log) (duplicates bool) {
	fmt.Println("check duplicates")
	identityBytes := append(raw.Address.Bytes(), raw.Topics[0].Bytes()...)
	identityBytes = append(identityBytes, raw.Data...)
	identity := new(big.Int).SetBytes(identityBytes).String()

	if _, duplicates = e.logFilter.Load(identity); duplicates {

	}
	e.logFilter.Store(identity, logRecord{raw, time.Now()})

	return
}

func (e *EthAdaptor) logMapTimeout() {
	timer := time.NewTimer(10 * time.Minute)
	for range timer.C {
		e.logFilter.Range(e.checkTime)
		timer.Reset(10 * time.Minute)
	}

}

func (e *EthAdaptor) checkTime(log, deliverTime interface{}) (okToDelete bool) {
	switch t := deliverTime.(type) {
	case logRecord:
		if time.Now().Sub(t.currTimeStamp).Seconds() > 60*10 {
			e.logFilter.Delete(log)
		}
	}
	return true
}
