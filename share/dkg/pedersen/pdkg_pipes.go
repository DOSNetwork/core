package dkg

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	errors "golang.org/x/xerrors"
)

func genPub(ctx context.Context, logger log.Logger, suite suites.Suite, id []byte, groupIds [][]byte, sessionID string) (out chan interface{}, secrc chan kyber.Scalar, errc chan error) {
	out = make(chan interface{})
	secrc = make(chan kyber.Scalar)
	errc = make(chan error)
	go func() {
		defer fmt.Println("1) Close genPub")
		defer close(out)
		defer close(errc)

		defer logger.TimeTrack(time.Now(), "genPub", map[string]interface{}{"GroupID": sessionID})
		//Index pub key
		defer fmt.Println("1) Start genPub Pipe ")

		index := -1
		for i, groupId := range groupIds {
			if r := bytes.Compare(id, groupId); r == 0 {
				index = i
				break
			}
		}
		if index == -1 {
			reportErr(ctx, errc, errors.New("Can't find id in group IDs"))
		}
		//Generate secret and public key
		sec := suite.Scalar().Pick(suite.RandomStream())
		select {
		case secrc <- sec:
		case <-ctx.Done():
		}
		pub := suite.Point().Mul(sec, nil)
		if bin, err := pub.MarshalBinary(); err != nil {
			reportErr(ctx, errc, err)
		} else {
			pubkey := &PublicKey{SessionId: sessionID, Index: uint32(index), Publickey: &vss.PublicKey{Binary: bin}}
			select {
			case out <- pubkey:
			case <-ctx.Done():
			}
			return
		}
	}()
	return
}

func exchangePub(ctx context.Context, logger log.Logger, selfPubc chan interface{}, peerPubc chan []interface{}, p p2p.P2PInterface, groupIds [][]byte, sessionID string) (out chan []*PublicKey, errc chan error) {
	out = make(chan []*PublicKey)
	errc = make(chan error)
	go func() {
		defer logger.TimeTrack(time.Now(), "exchangePub", map[string]interface{}{"GroupID": sessionID})
		defer fmt.Println("2) Close exchangePub Pipe ")
		defer close(out)
		defer close(errc)
		var partPubs []*PublicKey
		fmt.Println("2)Start exchangePub")
		select {
		case <-ctx.Done():
		case resp, ok := <-selfPubc:
			if ok {
				fmt.Println("exchangePub selfPubc")
				logger.TimeTrack(time.Now(), "exchangePubselfPubc", map[string]interface{}{"GroupID": sessionID})
				if pubkey, ok := resp.(*PublicKey); ok {
					partPubs = append(partPubs, pubkey)
				}
			}
		}
		for {
			select {
			case <-ctx.Done():
				return
			case resps, ok := <-peerPubc:
				if !ok {
					return
				}
				fmt.Println("3)exchangePub peerPubc ", len(resps))
				logger.TimeTrack(time.Now(), "exchangePubpeerPubc", map[string]interface{}{"GroupID": sessionID})
				for _, resp := range resps {
					if pubkey, ok := resp.(*PublicKey); ok {
						partPubs = append(partPubs, pubkey)
					}
				}
			}
			if len(partPubs) == len(groupIds) {
				select {
				case <-ctx.Done():
				case out <- partPubs:
				}
				return
			}
		}
	}()
	return
}

func sendToMembers(ctx context.Context, logger log.Logger, msgc chan interface{}, p p2p.P2PInterface, groupIds [][]byte, sessionID string) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer fmt.Println("3) Close sendToMembers Pipe ")
		defer close(errc)
		select {
		case <-ctx.Done():
		case msg, ok := <-msgc:
			if ok {
				defer logger.TimeTrack(time.Now(), "sendToMembers", map[string]interface{}{"GroupID": sessionID})
				fmt.Println("3) Start sendToMembers ")
				if m, ok := msg.(proto.Message); ok {
					var wg sync.WaitGroup
					wg.Add(len(groupIds) - 1)
					for i, id := range groupIds {
						if r := bytes.Compare(p.GetID(), id); r != 0 {
							fmt.Println("sendToMembers ", id)
							go func(i int, id []byte) {
								defer wg.Done()
								defer fmt.Println("sendToMembers ", id, " ", i, " done ")
								for {
									select {
									case <-ctx.Done():
									default:
										if _, err := p.Request(ctx, id, m); err != nil {
											fmt.Println("sendToMembers ", id, " err ", err)
											reportErr(ctx, errc, err)
										} else {

											return
										}
									}
								}
							}(i, id)
						}
					}
					wg.Wait()
				}
			}
		}
	}()
	return
}

func askMembers(ctx context.Context, logger log.Logger, bufToNode chan interface{}, numOfResp, reqTpe int, sessionID string) (out chan []interface{}) {
	out = make(chan []interface{})
	go func() {
		defer fmt.Println("4) Close askMembers Pipe ")
		defer logger.TimeTrack(time.Now(), "askMembers", map[string]interface{}{"GroupID": sessionID})

		req := request{ctx: ctx, reqType: reqTpe, sessionID: sessionID, numOfResps: numOfResp, reply: out}
		select {
		case <-ctx.Done():
		case bufToNode <- req:
		}
	}()
	return
}

func genDistKeyGenerator(ctx context.Context, logger log.Logger, secrc chan kyber.Scalar, partPubs chan []*PublicKey, numOfPubkeys int, suite suites.Suite, sessionID string) (out chan *DistKeyGenerator, errc chan error) {
	out = make(chan *DistKeyGenerator)
	errc = make(chan error)
	go func() {
		defer fmt.Println("5) Close genDistKeyGenerator Pipe ")
		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case sec, ok := <-secrc:
			if ok {
				select {
				case <-ctx.Done():
					fmt.Println(" genDistKeyGenerator ctx.Done ")
				case pubs, ok := <-partPubs:
					if ok {
						fmt.Println("5) Start genDistKeyGenerator")

						defer logger.TimeTrack(time.Now(), "genDistKeyGenerator", map[string]interface{}{"GroupID": sessionID})
						pubPoints := make([]kyber.Point, numOfPubkeys)
						for _, pubkey := range pubs {
							if pubPoints[pubkey.Index] != nil {
								fmt.Println("!!!Duplicate Index", pubkey.Index)
								return
							}
							fmt.Println(pubkey.Index)
							pubPoints[pubkey.Index] = suite.Point()
							if err := pubPoints[pubkey.Index].UnmarshalBinary(pubkey.Publickey.Binary); err != nil {
								reportErr(ctx, errc, err)
								return
							}
						}
						fmt.Println("!!!!! pubPoints", len(pubPoints))

						dkg, err := NewDistKeyGenerator(suite, sec, pubPoints, numOfPubkeys/2+1)
						if err != nil {
							reportErr(ctx, errc, err)
						} else {
							select {
							case <-ctx.Done():
							case out <- dkg:
							}
							return
						}
					}
				}
			}

		}
	}()
	return
}

func genDealsAndSend(ctx context.Context, logger log.Logger, dkgc chan *DistKeyGenerator, p p2p.P2PInterface, groupIds [][]byte, sessionID string) (out chan *DistKeyGenerator, errc chan error) {
	out = make(chan *DistKeyGenerator)
	errc = make(chan error)
	go func() {
		defer fmt.Println("6) Close genDealsAndSend Pipe ")

		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case dkg, ok := <-dkgc:
			if ok {
				fmt.Println("6) Start genDealsAndSend")
				defer logger.TimeTrack(time.Now(), "genDealsAndSend", map[string]interface{}{"GroupID": sessionID})
				if deals, err := dkg.Deals(); err == nil {
					var wg sync.WaitGroup
					wg.Add(len(deals))
					fmt.Println("6) Start genDealsAndSend add wg ", len(groupIds)-1, " deals ", len(deals))

					for i, d := range deals {
						d.SessionId = sessionID
						go func(id []byte, d *Deal) {
							defer wg.Done()
							defer fmt.Println("6) Start genDealsAndSend wg.Done()")

							for {
								select {
								case <-ctx.Done():
								default:
									fmt.Println("genDealsAndSend Request")

									if _, err := p.Request(ctx, id, d); err != nil {
										fmt.Println("genDealsAndSend Request err", err)
										reportErr(ctx, errc, err)
									} else {
										fmt.Println("genDealsAndSend Request done")
										return
									}
								}
							}
						}(groupIds[i], d)
					}
					wg.Wait()
					fmt.Println("6) Start genDealsAndSend done waiting")

					select {
					case <-ctx.Done():
					case out <- dkg:
					}
					fmt.Println("6) Start genDealsAndSend end")

				} else {
					reportErr(ctx, errc, err)
				}
			} else {
				fmt.Println("genDealsAndSend dkgc end")
			}
		}
	}()
	return
}
func getAndProcessDeals(ctx context.Context, logger log.Logger, dkgc chan *DistKeyGenerator, dealsc chan []interface{}, sessionID string) (dkgOut chan *DistKeyGenerator, out chan interface{}, errc chan error) {
	dkgOut = make(chan *DistKeyGenerator)
	out = make(chan interface{})
	errc = make(chan error)
	go func() {
		var dkg *DistKeyGenerator
		var ok bool
		defer fmt.Println("7) Close getAndProcessDeals Pipe ")

		defer close(dkgOut)
		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case dkg, ok = <-dkgc:
			if !ok {
				fmt.Println("getAndProcessDeals dkgc end")
				return
			}
		}
		fmt.Println("7) getAndProcessDeals got dkg")
		defer logger.TimeTrack(time.Now(), "getAndProcessDeals", map[string]interface{}{"GroupID": sessionID})
		select {
		case <-ctx.Done():
		case deals, ok := <-dealsc:
			if ok {
				fmt.Println("Start getAndProcessDeals")
				var resps []*Response
				for _, d := range deals {
					if deal, ok := d.(*Deal); ok {
						if resp, err := dkg.ProcessDeal(deal); err == nil {
							resp.SessionId = sessionID
							if vss.StatusApproval == resp.Response.Status {
								resps = append(resps, resp)
							} else {
								reportErr(ctx, errc, errors.New("resp StatusNotApproval"))
							}
						} else {
							reportErr(ctx, errc, err)
						}
					}
				}

				select {
				case <-ctx.Done():
					return
				case out <- &Responses{SessionId: sessionID, Response: resps}:
				}

				select {
				case <-ctx.Done():
					return
				case dkgOut <- dkg:
				}
			} else {
				fmt.Println("getAndProcessDeals dealsc end")
			}
		}
	}()
	return
}

func getAndProcessResponses(ctx context.Context, logger log.Logger, dkgc chan *DistKeyGenerator, respsc chan []interface{}, sessionID string) (out chan *DistKeyGenerator, errc chan error) {
	out = make(chan *DistKeyGenerator)
	errc = make(chan error)
	go func() {
		var dkg *DistKeyGenerator
		var ok bool
		defer fmt.Println("8) Close getAndProcessResponses Pipe ")

		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case dkg, ok = <-dkgc:
			if !ok {
				return
			}
		}
		fmt.Println("8) Start getAndProcessResponses")
		defer logger.TimeTrack(time.Now(), "getAndProcessResponses", map[string]interface{}{"GroupID": sessionID})
		select {
		case <-ctx.Done():
		case resps, ok := <-respsc:
			if ok {
				fmt.Println("getAndProcessResponses got resps")

				for _, r := range resps {
					if resp, ok := r.(*Response); ok {
						if _, err := dkg.ProcessResponse(resp); err != nil {
							reportErr(ctx, errc, err)
						}
					} else {
						reportErr(ctx, errc, errors.New("Response cast error"))
					}
				}

				select {
				case <-ctx.Done():
				case out <- dkg:
				}
			} else {
				fmt.Println("getAndProcessResponses respsc end")
			}
		}
	}()
	return
}
func genGroup(ctx context.Context, logger log.Logger, group *group, suite suites.Suite, dkgc <-chan *DistKeyGenerator, sessionID string) (out chan [5]*big.Int, errc chan error) {
	out = make(chan [5]*big.Int)
	errc = make(chan error)
	go func() {
		defer fmt.Println("9) Close genGroup Pipe ")

		defer close(out)
		defer close(errc)
		select {
		case <-ctx.Done():
		case dkg, ok := <-dkgc:
			fmt.Println("9) Start genGroup Pipe ")
			if ok {
				defer logger.TimeTrack(time.Now(), "genGroup", map[string]interface{}{"GroupID": sessionID})
				if !dkg.Certified() {
					reportErr(ctx, errc, errors.New("dkg is not certified"))
				}
				if secShare, err := dkg.DistKeyShare(); err == nil {
					group.secShare = secShare
					group.pubPoly = share.NewPubPoly(suite, suite.Point().Base(), group.secShare.Commitments())
					pubKey := group.pubPoly.Commit()
					if pubKeyCoor, err := decodePubKey(pubKey); err == nil {
						if groupId, ok := new(big.Int).SetString(sessionID, 16); ok {
							dataReturn := [5]*big.Int{groupId}
							copy(dataReturn[1:], pubKeyCoor[:])
							select {
							case <-ctx.Done():
							case out <- dataReturn:
							}
						} else {
							reportErr(ctx, errc, errors.New("sessionID cast error "))
						}
					} else {
						reportErr(ctx, errc, err)
					}
				} else {
					reportErr(ctx, errc, err)
				}
			} else {
				fmt.Println("genGroup dkgc end")

			}
		}
	}()
	return
}
