package dosnode

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"

	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
)

func (d *DosNode) queryLoop() {
	defer fmt.Println("End queryLoop")
	bufSign := make(map[string][]*vss.Signature)
	reqSign := make(map[string]request)
	peerMsg, _ := d.p.SubscribeMsg(50, vss.Signature{})
	defer d.p.UnSubscribeMsg(vss.Signature{})
	watchdog := time.NewTicker(watchdogInterval * time.Minute)
	defer watchdog.Stop()
	for {
		select {
		case <-d.ctx.Done():
			return
		case <-watchdog.C:
			for _, req := range reqSign {
				select {
				case <-req.ctx.Done():
					close(req.reply)
					delete(bufSign, req.requestID)
					delete(reqSign, req.requestID)
				default:
				}
			}
		case msg, ok := <-peerMsg:
			if ok {
				if content, ok := msg.Msg.Message.(*vss.Signature); ok {
					requestID := string(content.RequestId)
					if req := reqSign[requestID]; req.requestID == requestID {
						select {
						case <-req.ctx.Done():
						case req.reply <- content:
						}
					} else {
						bufSign[requestID] = append(bufSign[requestID], content)
					}
				}
			}
		case req, ok := <-d.reqSignc:
			if ok {
				//1)Check buf to see if it has enough signatures
				reqSign[req.requestID] = req
				if signs := bufSign[req.requestID]; len(signs) >= 0 {
					for _, sign := range signs {
						select {
						case <-req.ctx.Done():
						case req.reply <- sign:
						}
					}
					bufSign[req.requestID] = nil
				}
			}
		}
	}
}

func (d *DosNode) handleQuery(ids [][]byte, pubPoly *share.PubPoly, sec *share.PriShare, groupID string, requestID, lastRand, useSeed *big.Int, url, selector string, pType uint32) {
	queryCtx, cancel := context.WithTimeout(context.Background(), time.Duration(60*15*time.Second))
	defer cancel()
	queryCtxWithValue := context.WithValue(context.WithValue(queryCtx, ctxKey("RequestID"), fmt.Sprintf("%x", requestID)), ctxKey("GroupID"), groupID)
	d.logger.Event("HandleQuery", map[string]interface{}{"GroupID": groupID, "RequestID": fmt.Sprintf("%x", requestID)})
	defer d.logger.TimeTrack(time.Now(), "TimeHandleQuery", map[string]interface{}{"GroupID": groupID, "RequestID": fmt.Sprintf("%x", requestID)})
	defer cancel()
	var nonce []byte
	//Generate an unique id
	switch pType {
	case onchain.TrafficSystemRandom:
		var bytes []byte
		bytes = append(bytes, []byte(groupID)...)
		bytes = append(bytes, requestID.Bytes()...)
		bytes = append(bytes, lastRand.Bytes()...)
		nHash := sha256.Sum256(bytes)
		nonce = nHash[:]
	case onchain.TrafficUserRandom:
		var bytes []byte
		bytes = append(bytes, []byte(groupID)...)
		bytes = append(bytes, requestID.Bytes()...)
		bytes = append(bytes, lastRand.Bytes()...)
		bytes = append(bytes, useSeed.Bytes()...)
		nHash := sha256.Sum256(bytes)
		nonce = nHash[:]
	case onchain.TrafficUserQuery:
		var bytes []byte
		bytes = append(bytes, []byte(groupID)...)
		bytes = append(bytes, requestID.Bytes()...)
		bytes = append(bytes, lastRand.Bytes()...)
		bytes = append(bytes, []byte(url)...)
		bytes = append(bytes, []byte(selector)...)
		nHash := sha256.Sum256(bytes)
		nonce = nHash[:]
	}
	sign := &vss.Signature{
		Index:     pType,
		RequestId: requestID.Bytes(),
		Nonce:     nonce,
	}
	//Build a pipeline
	var errcList []chan error

	submitterc, errc := choseSubmitter(queryCtxWithValue, d.p, d.chain, lastRand, ids, 2, d.logger)
	errcList = append(errcList, errc)

	var contentc chan []byte
	switch pType {
	case onchain.TrafficSystemRandom:
		contentc = genSysRandom(queryCtxWithValue, submitterc[0], lastRand.Bytes(), d.logger)
	case onchain.TrafficUserRandom:
		contentc = genUserRandom(queryCtxWithValue, submitterc[0], requestID.Bytes(), lastRand.Bytes(), useSeed.Bytes(), d.logger)
	case onchain.TrafficUserQuery:
		contentc, errc = genQueryResult(queryCtxWithValue, submitterc[0], url, selector, d.logger)
		errcList = append(errcList, errc)
	}

	signc, errc := genSign(queryCtxWithValue, contentc, sec, d.suite, sign, d.logger)
	errcList = append(errcList, errc)
	signAllc := dispatchSign(queryCtxWithValue, submitterc[1], signc, d.reqSignc, d.p, requestID.Bytes(), (len(ids)/2 + 1), d.logger)
	errcList = append(errcList, errc)
	recoveredSignc, errc := recoverSign(queryCtxWithValue, signAllc, d.suite, pubPoly, (len(ids)/2 + 1), len(ids), d.logger)
	errcList = append(errcList, errc)

	switch pType {
	case onchain.TrafficSystemRandom:
		errc := d.chain.SetRandomNum(queryCtxWithValue, recoveredSignc)
		errcList = append(errcList, errc)
	default:
		errc := d.chain.DataReturn(queryCtxWithValue, recoveredSignc)
		errcList = append(errcList, errc)
	}
	allErrc := mergeErrors(queryCtxWithValue, errcList...)
	for {
		select {
		case err, ok := <-allErrc:
			if !ok {
				return
			}
			d.logger.Event("handleQueryError", map[string]interface{}{"Error": err.Error(), "GroupID": groupID})
		case <-queryCtxWithValue.Done():
			d.logger.Event("handleQueryError", map[string]interface{}{"Error": queryCtxWithValue.Err(), "GroupID": groupID})
			return
		}
	}
}
