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
	CHANTIMEOUT    = 1  //IN SECOND
	REQUESTTIMEOUT = 60 //IN SECOND
)

type P2PDkgInterface interface {
	GetGroupPublicPoly() *share.PubPoly
	GetShareSecurity() *share.PriShare
	IsCertified() bool
	Start(dkgSession *DkgSession) (chan bool, <-chan error)
}

type P2PDkg struct {
	groupId        []byte
	suite          suites.Suite
	groupCmd       chan *DkgSession
	network        *p2p.P2PInterface
	logger         log.Logger
	currentSession *DkgSession
}

type DkgSession struct {
	SessionId      string
	GroupIds       [][]byte
	certified      bool
	partSec        kyber.Scalar
	partPub        kyber.Point
	partDkg        *DistKeyGenerator
	pubKeys        Pubkeys
	pubkeyIdMap    map[string]string
	deals          []Deal
	partDks        *DistKeyShare
	groupPubPoly   *share.PubPoly
	subscribeEvent chan bool
	ctx            context.Context
	cancel         context.CancelFunc
	errc           chan error
	groupingStart  time.Time
}

func CreateP2PDkg(p p2p.P2PInterface, suite suites.Suite) (P2PDkgInterface, error) {
	d := &P2PDkg{
		groupId:  p.GetID(),
		suite:    suite,
		groupCmd: make(chan *DkgSession),
		network:  &p,
		logger:   log.New("module", "dkg"),
	}
	return d, d.eventLoop()
}

func (d *P2PDkg) Start(newSession *DkgSession) (chan bool, <-chan error) {
	newSession.partSec = d.suite.Scalar().Pick(d.suite.RandomStream())
	newSession.partPub = d.suite.Point().Mul(newSession.partSec, nil)
	newSession.ctx, newSession.cancel = context.WithCancel(context.Background())
	newSession.subscribeEvent = make(chan bool)
	newSession.errc = make(chan error)
	d.groupCmd <- newSession
	return newSession.subscribeEvent, newSession.errc
}

func (d *P2PDkg) IsCertified() bool {
	return d.currentSession.certified
}

func (d *P2PDkg) GetGroupPublicPoly() *share.PubPoly {
	if d.currentSession.certified {
		return d.currentSession.groupPubPoly
	}
	return nil
}

func (d *P2PDkg) GetShareSecurity() *share.PriShare {
	if d.currentSession.certified {
		return d.currentSession.partDks.Share
	}
	return nil
}

func (d *P2PDkg) eventLoop() (err error) {
	peerEvent, err := (*d.network).SubscribeEvent(100, ReqPublicKey{}, ReqDeal{}, ReqResponses{})
	if err == nil {
		go func() {
			defer (*d.network).UnSubscribeEvent(ReqPublicKey{}, ReqDeal{}, ReqResponses{})
			selfPubKey := vss.PublicKey{}
			selfDeals := make(map[string]*Deal)
			selfResps := Responses{}
			pubKeyResult := make(chan vss.PublicKey)
			dealResult := make(chan map[string]*Deal)
			responseResult := make(chan Responses)
			dkgResultChan := make(chan *DkgSession)
			for {
				select {
				case newSession := <-d.groupCmd:
					if d.currentSession != nil {
						d.currentSession.cancel()
						for range d.currentSession.errc {
						}
					}
					selfPubKey = vss.PublicKey{}
					selfDeals = make(map[string]*Deal)
					selfResps = Responses{}
					var errcList []<-chan error
					outForDeal, pubKeyErrc := d.pipeExchangePubKey(newSession, pubKeyResult)
					errcList = append(errcList, pubKeyErrc)
					outForResponse, genErrc := d.pipeNewDistKeyGenerator(outForDeal, dealResult)
					errcList = append(errcList, genErrc)
					RespErrc := d.pipeProcessDealAndResponses(outForResponse, responseResult, dkgResultChan)
					errcList = append(errcList, RespErrc)
					merge(newSession.errc, errcList...)
					d.currentSession = newSession
					newSession.groupingStart = time.Now()
				case msg := <-peerEvent:
					switch content := msg.Msg.Message.(type) {
					case *ReqPublicKey:
						if d.currentSession != nil && d.currentSession.SessionId == content.SessionId {
							(*d.network).Reply(msg.Sender, msg.RequestNonce, &selfPubKey)
						} else {
							(*d.network).Reply(msg.Sender, msg.RequestNonce, &vss.PublicKey{})
						}
					case *ReqDeal:
						if d.currentSession != nil && d.currentSession.SessionId == content.SessionId {
							(*d.network).Reply(msg.Sender, msg.RequestNonce, selfDeals[string(msg.Sender)])
						} else {
							(*d.network).Reply(msg.Sender, msg.RequestNonce, &Deal{})
						}
					case *ReqResponses:
						if d.currentSession != nil && d.currentSession.SessionId == content.SessionId {
							(*d.network).Reply(msg.Sender, msg.RequestNonce, &selfResps)
						} else {
							(*d.network).Reply(msg.Sender, msg.RequestNonce, &Responses{})
						}
					default:
					}
				case finishedSession := <-dkgResultChan:
					finishedSession.certified = true
					timeCost := time.Since(finishedSession.groupingStart).Seconds()
					fmt.Println("DistKeyShare SUCCESS ", timeCost)
					if finishedSession.subscribeEvent != nil {
						finishedSession.subscribeEvent <- finishedSession.certified
					}
				case selfPubKey = <-pubKeyResult:
				case selfDeals = <-dealResult:
				case selfResps = <-responseResult:
				}
			}
		}()
	}
	return
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

func (d *P2PDkg) pipeExchangePubKey(newSession *DkgSession, outToEventloop chan<- vss.PublicKey) (<-chan *DkgSession, <-chan error) {
	out := make(chan *DkgSession)
	errc := make(chan error)

	go func() {
		defer close(errc)
		defer close(out)
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
		select {
		case outToEventloop <- public:
		case <-newSession.ctx.Done():
			fmt.Println("pipeExchangePubKey closed")
			return
		}

		groupPeers := make(map[string][]byte)
		for _, id := range newSession.GroupIds {
			if bytes.Compare(id, d.groupId) != 0 {
				groupPeers[string(id)] = id
			}
		}
		timer := time.NewTimer(REQUESTTIMEOUT * time.Second)
		for len(groupPeers) > 0 {
			for key, id := range groupPeers {
				select {
				case <-newSession.ctx.Done():
					fmt.Println("pipeExchangePubKey closed")
					return
				default:
				}
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
				default:
				}
			}
		}
		select {
		case out <- newSession:
		case <-newSession.ctx.Done():
			fmt.Println("pipeExchangePubKey closed")
			return
		}
		fmt.Println("pipeExchangePubKey closed")
	}()

	return out, errc
}

func (d *P2PDkg) pipeNewDistKeyGenerator(dkgSession <-chan *DkgSession, outToEventloop chan<- map[string]*Deal) (<-chan *DkgSession, <-chan error) {
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
			select {
			case outToEventloop <- idDealMap:
			case <-newSession.ctx.Done():
				fmt.Println("pipeExchangePubKey closed")
				return
			}

			groupPeers := make(map[string][]byte)
			for _, id := range newSession.GroupIds {
				if bytes.Compare(id, d.groupId) != 0 {
					groupPeers[string(id)] = id
				}
			}
			timer := time.NewTimer(REQUESTTIMEOUT * time.Second)
			for len(groupPeers) > 0 {
				for key, id := range groupPeers {
					select {
					case <-newSession.ctx.Done():
						fmt.Println("pipeNewDistKeyGenerator closed")
						return
					default:
					}
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
					default:
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

func (d *P2PDkg) pipeProcessDealAndResponses(dkgSession <-chan *DkgSession, RespsToEventloop chan<- Responses, finishToEventloop chan<- *DkgSession) <-chan error {
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
			select {
			case RespsToEventloop <- Responses{Response: resps}:
			case <-newSession.ctx.Done():
				fmt.Println("pipeExchangePubKey closed")
				return
			}

			groupPeers := make(map[string][]byte)
			for _, id := range newSession.GroupIds {
				if bytes.Compare(id, d.groupId) != 0 {
					groupPeers[string(id)] = id
				}
			}
			timer := time.NewTimer(REQUESTTIMEOUT * time.Second)
			for len(groupPeers) > 0 {
				for key, id := range groupPeers {
					select {
					case <-newSession.ctx.Done():
						fmt.Println("pipeNewDistKeyGenerator closed")
						return
					default:
					}
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
					default:
					}
				}
			}
			if (*newSession.partDkg).Certified() {
				var err error
				newSession.partDks, err = newSession.partDkg.DistKeyShare()
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
				newSession.groupPubPoly = share.NewPubPoly(d.suite, d.suite.Point().Base(), newSession.partDks.Commitments())
				select {
				case finishToEventloop <- newSession:
				case <-newSession.ctx.Done():
					fmt.Println("pipeNewDistKeyGenerator closed")
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
