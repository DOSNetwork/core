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

		defer logger.TimeTrack(time.Now(), "genPub", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
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
			err := &DKGError{err: errors.Errorf("index id failed for GID %s : %w", sessionID, ErrCanNotFindID)}
			reportErr(ctx, errc, err)
			return
		}
		//Generate secret and public key
		sec := suite.Scalar().Pick(suite.RandomStream())
		select {
		case secrc <- sec:
		case <-ctx.Done():
			return
		}
		pub := suite.Point().Mul(sec, nil)
		bin, err := pub.MarshalBinary()
		if err != nil {
			err := &DKGError{err: errors.Errorf("MarshalBinary failed for GID %s : %w", sessionID, err)}
			reportErr(ctx, errc, err)
			return
		}
		pubkey := &PublicKey{SessionId: sessionID, Index: uint32(index), Publickey: &vss.PublicKey{Binary: bin}}
		select {
		case out <- pubkey:
		case <-ctx.Done():
		}
		return
	}()
	return
}

func exchangePub(ctx context.Context, logger log.Logger, selfPubc chan interface{}, peerPubc chan []interface{}, p p2p.P2PInterface, groupIds [][]byte, sessionID string) (out chan []*PublicKey, errc chan error) {
	out = make(chan []*PublicKey)
	errc = make(chan error)
	go func() {
		defer logger.TimeTrack(time.Now(), "exchangePub", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
		defer fmt.Println("2) Close exchangePub Pipe ")
		defer close(out)
		defer close(errc)
		var partPubs []*PublicKey
		fmt.Println("2)Start exchangePub")
		select {
		case <-ctx.Done():
			return
		case resp, ok := <-selfPubc:
			if !ok {
				return
			}
			fmt.Println("exchangePub selfPubc")
			logger.TimeTrack(time.Now(), "exchangePubselfPubc", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
			pubkey, ok := resp.(*PublicKey)
			if !ok {
				err := &DKGError{err: errors.Errorf("casting PublicKey failed for GID %s : %w", sessionID, ErrCasting)}
				reportErr(ctx, errc, err)
				return
			}
			partPubs = append(partPubs, pubkey)

		}
		for {
			select {
			case <-ctx.Done():
				return
			case resps, ok := <-peerPubc:
				if !ok {
					return
				}
				logger.TimeTrack(time.Now(), "exchangePubpeerPubc", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
				for _, resp := range resps {
					pubkey, ok := resp.(*PublicKey)
					if !ok {
						err := &DKGError{err: errors.Errorf("casting PublicKey failed for GID %s : %w", sessionID, ErrCasting)}
						reportErr(ctx, errc, err)
						return
					}
					partPubs = append(partPubs, pubkey)
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
				defer logger.TimeTrack(time.Now(), "sendToMembers", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
				fmt.Println("3) Start sendToMembers ")
				if m, ok := msg.(proto.Message); ok {
					var wg sync.WaitGroup
					wg.Add(len(groupIds) - 1)
					for i, id := range groupIds {
						if r := bytes.Compare(p.GetID(), id); r != 0 {
							go func(i int, id []byte) {
								defer wg.Done()
								for {
									select {
									case <-ctx.Done():
										return
									default:
										if _, err := p.Request(ctx, id, m); err != nil {
											err := &DKGError{err: errors.Errorf("sendToMembers failed for GID %s : %w", sessionID, err)}
											reportErr(ctx, errc, err)
											time.Sleep(100 * time.Millisecond)
											continue
										}
										return
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
		defer logger.TimeTrack(time.Now(), "askMembers", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})

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
						defer logger.TimeTrack(time.Now(), "genDistKeyGenerator", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
						pubPoints := make([]kyber.Point, numOfPubkeys)
						for _, pubkey := range pubs {
							if pubPoints[pubkey.Index] != nil {
								err := &DKGError{err: errors.Errorf("genDistKeyGenerator failed for GID %s : %w", sessionID, ErrDupPubKeyIndex)}
								reportErr(ctx, errc, err)
								return
							}
							pubPoints[pubkey.Index] = suite.Point()
							if err := pubPoints[pubkey.Index].UnmarshalBinary(pubkey.Publickey.Binary); err != nil {
								err := &DKGError{err: errors.Errorf("UnmarshalBinary failed for GID %s : %w", sessionID, err)}
								reportErr(ctx, errc, err)
								return
							}
						}

						dkg, err := NewDistKeyGenerator(suite, sec, pubPoints, numOfPubkeys/2+1)
						if err != nil {
							err := &DKGError{err: errors.Errorf("NewDistKeyGenerator failed for GID %s : %w", sessionID, err)}
							reportErr(ctx, errc, err)
							return
						}
						select {
						case <-ctx.Done():
						case out <- dkg:
						}
						return
					}
				}
			}

		}
	}()
	return
}

//err = &P2PError{err: errors.Errorf("runClient (%t) err: %w", inBound, err), t: time.Now()}
//n.logger.Error(err)
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
				defer logger.TimeTrack(time.Now(), "genDealsAndSend", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
				deals, err := dkg.Deals()
				if err != nil {
					err := &DKGError{err: errors.Errorf("genDealsAndSend failed for GID %s : %w", sessionID, err)}
					reportErr(ctx, errc, err)
					return
				}
				var wg sync.WaitGroup
				wg.Add(len(deals))
				fmt.Println("6) Start genDealsAndSend add wg ", len(groupIds)-1, " deals ", len(deals))

				for i, d := range deals {
					d.SessionId = sessionID
					go func(id []byte, d *Deal) {
						defer wg.Done()
						for {
							select {
							case <-ctx.Done():
								return
							default:
								fmt.Println("genDealsAndSend Request")
								if _, err := p.Request(ctx, id, d); err != nil {
									err := &DKGError{err: errors.Errorf("genDealsAndSend failed for GID %s : %w", sessionID, err)}
									reportErr(ctx, errc, err)
									time.Sleep(100 * time.Millisecond)
									continue
								}
								return
							}
						}
					}(groupIds[i], d)
				}
				wg.Wait()

				select {
				case <-ctx.Done():
				case out <- dkg:
				}
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
		defer logger.TimeTrack(time.Now(), "getAndProcessDeals", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
		select {
		case <-ctx.Done():
		case deals, ok := <-dealsc:
			if ok {
				fmt.Println("Start getAndProcessDeals")
				var resps []*Response
				for _, d := range deals {
					deal, ok := d.(*Deal)
					if !ok {
						err := &DKGError{err: errors.Errorf("Casting Deal failed for GID %s : %w", sessionID, ErrCasting)}
						reportErr(ctx, errc, err)
						return
					}
					resp, err := dkg.ProcessDeal(deal)
					if err != nil {
						err = &DKGError{err: errors.Errorf("ProcessDeal failed for GID %s : %w", sessionID, err)}
						reportErr(ctx, errc, err)
						return
					}
					resp.SessionId = sessionID
					if vss.StatusApproval != resp.Response.Status {
						err = &DKGError{err: errors.Errorf("ProcessDeal failed for GID %s : %w", sessionID, ErrResponseNoApproval)}
						reportErr(ctx, errc, err)
						return
					}
					resps = append(resps, resp)
				}

				select {
				case <-ctx.Done():
					return
				case out <- &Responses{SessionId: sessionID, Response: resps}:
				}

				select {
				case <-ctx.Done():
				case dkgOut <- dkg:
				}
			}
		}
	}()
	return
}

func getAndProcessResponses(ctx context.Context, logger log.Logger, dkgc chan *DistKeyGenerator, respsc chan []interface{}, sessionID string) (out chan *DistKeyGenerator, errc chan error) {
	out = make(chan *DistKeyGenerator)
	errc = make(chan error)
	go func() {
		defer close(out)
		defer close(errc)
		var dkg *DistKeyGenerator
		var ok bool
		select {
		case <-ctx.Done():
		case dkg, ok = <-dkgc:
			if !ok {
				return
			}
		}
		defer logger.TimeTrack(time.Now(), "getAndProcessResponses", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
		select {
		case <-ctx.Done():
		case resps, ok := <-respsc:
			if ok {
				for _, r := range resps {
					resp, ok := r.(*Response)
					if !ok {
						err := &DKGError{err: errors.Errorf("getAndProcessResponses failed for GID %s : %w", sessionID, ErrCasting)}
						reportErr(ctx, errc, err)
						return
					}
					if _, err := dkg.ProcessResponse(resp); err != nil {
						err := &DKGError{err: errors.Errorf("ProcessResponse failed for GID %s : %w", sessionID, err)}
						reportErr(ctx, errc, err)
						return
					}
				}

				select {
				case <-ctx.Done():
				case out <- dkg:
				}
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
				defer logger.TimeTrack(time.Now(), "genGroup", map[string]interface{}{"GroupID": sessionID,"Topic": "Grouping"})
				if !dkg.Certified() {
					err := &DKGError{err: errors.Errorf("ProcessResponse failed for GID %s : %w", sessionID, ErrDKGNotCertified)}
					reportErr(ctx, errc, err)
					return
				}
				secShare, err := dkg.DistKeyShare()
				if err != nil {
					err := &DKGError{err: errors.Errorf("DistKeyShare failed for GID %s : %w", sessionID, err)}
					reportErr(ctx, errc, err)
					return
				}
				group.secShare = secShare
				group.pubPoly = share.NewPubPoly(suite, suite.Point().Base(), group.secShare.Commitments())
				pubKey := group.pubPoly.Commit()
				pubKeyCoor, err := decodePubKey(pubKey)
				if err != nil {
					err := &DKGError{err: errors.Errorf("decodePubKey failed for GID %s : %w", sessionID, err)}
					reportErr(ctx, errc, err)
					return
				}
				groupId, ok := new(big.Int).SetString(sessionID, 16)
				if !ok {
					err := &DKGError{err: errors.Errorf("genGroup failed for GID %s : %w", sessionID, ErrCasting)}
					reportErr(ctx, errc, err)
					return
				}
				dataReturn := [5]*big.Int{groupId}
				copy(dataReturn[1:], pubKeyCoor[:])
				select {
				case <-ctx.Done():
				case out <- dataReturn:
				}
			}
		}
	}()
	return
}
