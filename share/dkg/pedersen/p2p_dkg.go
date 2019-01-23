package dkg

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"

	"github.com/dedis/kyber"
)

const (
	INIT = iota
	PUBLICKEYDONE
	DEALDONE
	RESPONSEDONE
	VERIFIED
)

const (
	CHANTIMEOUT    = 1  //IN SECOND
	REQUESTTIMEOUT = 60 //IN SECOND
)

type P2PDkgInterface interface {
	GetGroupPublicPoly() *share.PubPoly
	GetShareSecurity() *share.PriShare
	IsCertified() bool
	Start(dkgSession *DkgSession) (chan int, <-chan error)
}

type P2PDkg struct {
	groupId        []byte
	suite          suites.Suite
	partDks        *DistKeyShare
	groupPubPoly   *share.PubPoly
	currentState   int
	currentSession string
	groupCmd       chan *DkgSession
	network        *p2p.P2PInterface
	logger         log.Logger
	groupingStart  time.Time
}

type DkgSession struct {
	SessionId      string
	GroupIds       [][]byte
	partSec        kyber.Scalar
	partPub        kyber.Point
	partDkg        *DistKeyGenerator
	pubKeys        Pubkeys
	pubkeyIdMap    map[string]string
	deals          []Deal
	selfPubKey     vss.PublicKey
	selfDeals      map[string]*Deal
	selfResps      Responses
	subscribeEvent chan int
	ctx            context.Context
	cancel         context.CancelFunc
	err            chan error
}

func CreateP2PDkg(p p2p.P2PInterface, suite suites.Suite) P2PDkgInterface {
	d := &P2PDkg{
		groupId:      p.GetID(),
		suite:        suite,
		currentState: INIT,
		groupCmd:     make(chan *DkgSession),
		network:      &p,
		logger:       log.New("module", "dkg"),
	}
	go d.eventLoop()
	return d
}

func (d *P2PDkg) Start(newSession *DkgSession) (chan int, <-chan error) {
	newSession.partSec = d.suite.Scalar().Pick(d.suite.RandomStream())
	newSession.partPub = d.suite.Point().Mul(newSession.partSec, nil)
	newSession.ctx, newSession.cancel = context.WithCancel(context.Background())
	newSession.subscribeEvent = make(chan int)
	newSession.err = make(chan error)
	d.groupCmd <- newSession
	return newSession.subscribeEvent, newSession.err
}

func (d *P2PDkg) IsCertified() bool {
	return d.currentState == VERIFIED
}

func (d *P2PDkg) GetGroupPublicPoly() *share.PubPoly {
	if d.currentState == VERIFIED {
		return d.groupPubPoly
	}
	return nil
}

func (d *P2PDkg) GetShareSecurity() *share.PriShare {
	if d.currentState == VERIFIED {
		return d.partDks.Share
	}
	return nil
}

func (d *P2PDkg) eventLoop() {
	var preCancel func()
	var preErrCh <-chan error
	for newSession := range d.groupCmd {
		d.currentSession = newSession.SessionId
		if preCancel != nil {
			preCancel()
			for range preErrCh {
			}
		}
		var errcList []<-chan error
		outForPubKey, requestErrc := d.pipReplyContentRequest(newSession)
		errcList = append(errcList, requestErrc)
		outForDeal, pubKeyErrc := d.pipeExchangePubKey(outForPubKey)
		errcList = append(errcList, pubKeyErrc)
		outForResponse, genErrc := d.pipeNewDistKeyGenerator(outForDeal)
		errcList = append(errcList, genErrc)
		RespErrc := d.pipeProcessDealAndResponses(outForResponse)
		errcList = append(errcList, RespErrc)
		merge(newSession.err, errcList...)
		preErrCh = newSession.err
		preCancel = newSession.cancel
	}
}

func merge(out chan error, cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	output := func(c <-chan error) {
		for n := range c {
			timer := time.NewTimer(CHANTIMEOUT * time.Second)
			select {
			case out <- n:
			case <-timer.C:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func (d *P2PDkg) pipReplyContentRequest(newSession *DkgSession) (<-chan *DkgSession, <-chan error) {
	out := make(chan *DkgSession)
	errc := make(chan error)
	d.currentState = INIT
	d.groupingStart = time.Now()

	go func() {
		defer close(errc)
		defer close(out)
		defer (*d.network).UnSubscribeEvent(ReqPublicKey{}, ReqDeal{}, ReqResponses{})
		peerEvent, err := (*d.network).SubscribeEvent(100, ReqPublicKey{}, ReqDeal{}, ReqResponses{})
		if err != nil {
			select {
			case errc <- err:
				newSession.cancel()
			case <-newSession.ctx.Done():
			}
			fmt.Println("pipReplyContentRequest closed")
			return
		}
		out <- newSession
		for {
			select {
			//event from peer
			case msg := <-peerEvent:
				switch content := msg.Msg.Message.(type) {
				case *ReqPublicKey:
					if d.currentState >= PUBLICKEYDONE && newSession.SessionId == content.SessionId {
						(*d.network).Reply(msg.Sender, msg.RequestNonce, &newSession.selfPubKey)
					} else {
						(*d.network).Reply(msg.Sender, msg.RequestNonce, &vss.PublicKey{})
					}
				case *ReqDeal:
					if d.currentState >= DEALDONE && newSession.SessionId == content.SessionId {
						(*d.network).Reply(msg.Sender, msg.RequestNonce, newSession.selfDeals[string(msg.Sender)])
					} else {
						(*d.network).Reply(msg.Sender, msg.RequestNonce, &Deal{})
					}
				case *ReqResponses:
					if d.currentState >= RESPONSEDONE && newSession.SessionId == content.SessionId {
						(*d.network).Reply(msg.Sender, msg.RequestNonce, &newSession.selfResps)
					} else {
						(*d.network).Reply(msg.Sender, msg.RequestNonce, &Responses{})
					}
				default:
				}
			case <-newSession.ctx.Done():
				close(newSession.subscribeEvent)
				fmt.Println("pipReplyContentRequest closed")
				return
			}
		}
	}()

	return out, errc
}

func (d *P2PDkg) pipeExchangePubKey(dkgSession <-chan *DkgSession) (<-chan *DkgSession, <-chan error) {
	out := make(chan *DkgSession)
	errc := make(chan error)

	go func() {
		defer close(errc)
		defer close(out)
		for newSession := range dkgSession {
			newSession.pubkeyIdMap = make(map[string]string)
			newSession.pubkeyIdMap[newSession.partPub.String()] = string(d.groupId)
			newSession.pubKeys = append(newSession.pubKeys, newSession.partPub)
			public := vss.PublicKey{SenderId: d.groupId}
			if err := public.SetPoint(d.suite, newSession.partPub); err != nil {
				d.logger.Error(err)
				select {
				case errc <- err:
					newSession.cancel()
				case <-newSession.ctx.Done():
				}
				fmt.Println("pipeExchangePubKey closed")
				return
			}
			newSession.selfPubKey = public
			d.currentState = PUBLICKEYDONE

			groupPeers := make(map[string][]byte)
			for _, id := range newSession.GroupIds {
				if bytes.Compare(id, d.groupId) != 0 {
					groupPeers[string(id)] = id
				}
			}
			timer := time.NewTimer(REQUESTTIMEOUT * time.Second)
			for len(groupPeers) > 0 {
				for key, id := range groupPeers {
					if d.currentSession == newSession.SessionId {
						pubkey, err := (*d.network).Request(id, &ReqPublicKey{SessionId: newSession.SessionId})
						if err != nil {
							select {
							case <-timer.C:
								errc <- errors.New("pipeExchangePubKey request timeout")
								newSession.cancel()
								return
							case <-newSession.ctx.Done():
								fmt.Println("pipeExchangePubKey closed")
								return
							default:
								continue
							}
						}
						switch content := pubkey.(type) {
						case *vss.PublicKey:
							if content.GetBinary() != nil {
								p, err := content.GetPoint(d.suite)
								if err != nil {
									d.logger.Error(err)
									select {
									case errc <- err:
										newSession.cancel()
									case <-newSession.ctx.Done():
									}
									fmt.Println("pipeExchangePubKey closed")
									return
								}
								newSession.pubkeyIdMap[p.String()] = string(content.SenderId)
								newSession.pubKeys = append(newSession.pubKeys, p)
								delete(groupPeers, key)
							}
						}
					} else {
						newSession.cancel()
						fmt.Println("pipeExchangePubKey closed")
						return
					}
				}
			}
			select {
			case out <- newSession:
			case <-newSession.ctx.Done():
				fmt.Println("pipeExchangePubKey closed")
				return
			}
		}
		fmt.Println("pipeExchangePubKey closed")
	}()

	return out, errc
}

func (d *P2PDkg) pipeNewDistKeyGenerator(dkgSession <-chan *DkgSession) (<-chan *DkgSession, <-chan error) {
	out := make(chan *DkgSession)
	errc := make(chan error)

	go func() {
		defer close(errc)
		defer close(out)
		for newSession := range dkgSession {
			var err error
			sort.Sort(newSession.pubKeys)
			newSession.partDkg, err = NewDistKeyGenerator(d.suite, newSession.partSec, newSession.pubKeys, len(newSession.GroupIds)/2+1)
			if err != nil {
				d.logger.Error(err)
				select {
				case errc <- err:
					newSession.cancel()
				case <-newSession.ctx.Done():
				}
				fmt.Println("pipeNewDistKeyGenerator closed")
				return
			}

			idxDealMap, err := newSession.partDkg.Deals()
			if err != nil {
				d.logger.Error(err)
				select {
				case errc <- err:
					newSession.cancel()
				case <-newSession.ctx.Done():
				}
				fmt.Println("pipeNewDistKeyGenerator closed")
				return
			}

			idDealMap := make(map[string]*Deal)
			for i, pub := range newSession.pubKeys {
				idDealMap[newSession.pubkeyIdMap[pub.String()]] = idxDealMap[i]
			}
			newSession.selfDeals = idDealMap
			d.currentState = DEALDONE

			groupPeers := make(map[string][]byte)
			for _, id := range newSession.GroupIds {
				if bytes.Compare(id, d.groupId) != 0 {
					groupPeers[string(id)] = id
				}
			}
			timer := time.NewTimer(REQUESTTIMEOUT * time.Second)
			for len(groupPeers) > 0 {
				for key, id := range groupPeers {
					if d.currentSession == newSession.SessionId {
						deal, err := (*d.network).Request(id, &ReqDeal{SessionId: newSession.SessionId})
						if err != nil {
							select {
							case <-timer.C:
								errc <- errors.New("pipeNewDistKeyGenerator request timeout")
								newSession.cancel()
								return
							case <-newSession.ctx.Done():
								fmt.Println("pipeNewDistKeyGenerator closed")
								return
							default:
								continue
							}
						}
						switch content := deal.(type) {
						case *Deal:
							if content.GetDeal() != nil {
								newSession.deals = append(newSession.deals, *content)
								delete(groupPeers, key)
							}
						}
					} else {
						newSession.cancel()
						fmt.Println("pipeNewDistKeyGenerator closed")
						return
					}
				}
			}
			select {
			case out <- newSession:
			case <-newSession.ctx.Done():
				fmt.Println("pipeNewDistKeyGenerator closed")
				return
			}
		}
		fmt.Println("pipeNewDistKeyGenerator closed")
	}()

	return out, errc
}

func (d *P2PDkg) pipeProcessDealAndResponses(dkgSession <-chan *DkgSession) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)
		for newSession := range dkgSession {
			var resps []*Response
			for _, deal := range newSession.deals {
				resp, err := (*newSession.partDkg).ProcessDeal(&deal)
				if err != nil {
					d.logger.Error(err)
					select {
					case errc <- err:
						newSession.cancel()
					case <-newSession.ctx.Done():
					}
					fmt.Println("pipeProcessDealAndResponses closed")
					return
				} else {
					resps = append(resps, resp)
				}
			}
			newSession.selfResps = Responses{Response: resps}
			d.currentState = RESPONSEDONE

			groupPeers := make(map[string][]byte)
			for _, id := range newSession.GroupIds {
				if bytes.Compare(id, d.groupId) != 0 {
					groupPeers[string(id)] = id
				}
			}
			timer := time.NewTimer(REQUESTTIMEOUT * time.Second)
			for len(groupPeers) > 0 {
				for key, id := range groupPeers {
					if d.currentSession == newSession.SessionId {
						responses, err := (*d.network).Request(id, &ReqResponses{SessionId: newSession.SessionId})
						if err != nil {
							select {
							case <-timer.C:
								errc <- errors.New("pipeProcessDealAndResponses request timeout")
								newSession.cancel()
								return
							case <-newSession.ctx.Done():
								fmt.Println("pipeProcessDealAndResponses closed")
								return
							default:
								continue
							}
						}
						switch content := responses.(type) {
						case *Responses:
							if resps := content.GetResponse(); resps != nil {
								for _, r := range resps {
									if _, err := (*newSession.partDkg).ProcessResponse(r); err != nil {
										d.logger.Error(err)
										select {
										case errc <- err:
											newSession.cancel()
										case <-newSession.ctx.Done():
										}
										fmt.Println("pipeProcessDealAndResponses closed")
										return
									}
								}
								delete(groupPeers, key)
							}
						}
					} else {
						newSession.cancel()
						fmt.Println("pipeProcessDealAndResponses closed")
						return
					}
				}
			}

			if (*newSession.partDkg).Certified() && d.currentSession == newSession.SessionId {
				var err error
				d.partDks, err = newSession.partDkg.DistKeyShare()
				if err != nil {
					d.logger.Error(err)
					select {
					case errc <- err:
						newSession.cancel()
					case <-newSession.ctx.Done():
					}
					fmt.Println("pipeProcessDealAndResponses closed")
					return
				}
				d.groupPubPoly = share.NewPubPoly(d.suite, d.suite.Point().Base(), d.partDks.Commitments())
				d.currentState = VERIFIED
				timeCost := time.Since(d.groupingStart).Seconds()
				fmt.Println("DistKeyShare SUCCESS ", timeCost)
				if newSession.subscribeEvent != nil && d.currentSession == newSession.SessionId {
					select {
					case newSession.subscribeEvent <- VERIFIED:
					case <-newSession.ctx.Done():
						fmt.Println("pipeProcessDealAndResponses closed")
						return
					}
				} else {
					newSession.cancel()
					fmt.Println("pipeProcessDealAndResponses closed")
					return
				}
			} else {
				newSession.cancel()
				fmt.Println("pipeProcessDealAndResponses closed")
				return
			}
		}
		fmt.Println("pipeProcessDealAndResponses closed")
	}()

	return errc
}
