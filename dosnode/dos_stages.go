package dosnode

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
	"github.com/antchfx/xmlquery"
	"github.com/spyzhov/ajson"
)

const (
	randNumberSize = 32
	addrLen        = 20
)

func mergeErrors(ctx context.Context, cs ...chan error) chan error {
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

func fanIn(ctx context.Context, channels ...chan *vss.Signature) chan *vss.Signature {
	var wg sync.WaitGroup
	multiplexedStream := make(chan *vss.Signature)

	multiplex := func(c <-chan *vss.Signature) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedStream <- i:
			}
		}
	}

	// Select from all the channels
	wg.Add(len(channels))

	for _, c := range channels {
		go multiplex(c)
	}

	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

//choseSubmitter choses a submitter according to the last random number and check if the submitter is reachable
func choseSubmitter(ctx context.Context, p p2p.P2PInterface, e onchain.ProxyAdapter, lastSysRand *big.Int, ids [][]byte, outCount int, logger log.Logger) ([]chan []byte, chan error) {
	errc := make(chan error)
	var outs []chan []byte
	for i := 0; i < outCount; i++ {
		outs = append(outs, make(chan []byte, 1))
	}

	go func() {
		defer close(errc)
		start := time.Now()
		submitter := lastSysRand.Uint64() % uint64(len(ids))

		for _, out := range outs {
			select {
			case out <- ids[submitter]:
			case <-ctx.Done():
			}
		}
		for _, out := range outs {
			close(out)
		}
		logger.TimeTrack(start, "ChoseSubmitter", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})
		return
	}()
	return outs, errc
}

func genUserRandom(ctx context.Context, submitterc chan []byte, requestId []byte, lastSysRand []byte, userSeed []byte, logger log.Logger) chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)

		select {
		case submitter, ok := <-submitterc:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "GenUserRandom",
				map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})
			logger.Info("GenUserRandom ")

			// signed message: concat(requestId, lastSystemRandom, userSeed, submitter address)
			random := append(requestId, lastSysRand...)
			random = append(random, userSeed...)
			random = append(random, submitter...)
			select {
			case out <- random:
			case <-ctx.Done():
			}
			return

		case <-ctx.Done():
			return
		}
	}()
	return out
}

func genSysRandom(ctx context.Context, submitterc chan []byte, lastSysRand []byte, logger log.Logger) chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)

		select {
		case submitter, ok := <-submitterc:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "GenSysRandom",
				map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")),
					"RequestID": ctx.Value(ctxKey("RequestID"))})
			logger.Info("genSysRandom ")
			// signed message: concat(lastSystemRandom, submitter address)
			paddedLastSysRand := padOrTrim(lastSysRand, randNumberSize)
			random := append(paddedLastSysRand, submitter...)
			select {
			case out <- random:
			case <-ctx.Done():
			}
			return
		case <-ctx.Done():
			return
		}
	}()
	return out
}

func dataFetch(url string) (body []byte, err error) {
	client := &http.Client{Timeout: 60 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = r.Body.Close()
	return
}

func dataParse(rawMsg []byte, pathStr string) (msg []byte, err error) {
	if pathStr == "" {
		msg = rawMsg
	} else if strings.HasPrefix(pathStr, "$") {
		var nodes []*ajson.Node
		nodes, err = ajson.JSONPath(rawMsg, pathStr)
		if err != nil {
			return
		}
		results := make([]interface{}, 0)
		for _, node := range nodes {
			var value interface{}
			value, err = node.Unpack()
			if err != nil {
				return
			}
			results = append(results, value)
		}

		msg, err = json.Marshal(results)
	} else if strings.HasPrefix(pathStr, "/") {
		var rawMsgXml *xmlquery.Node
		if rawMsgXml, err = xmlquery.Parse(bytes.NewReader(rawMsg)); err != nil {
			return
		}

		xmlNodes := xmlquery.Find(rawMsgXml, pathStr)
		for _, xmlNode := range xmlNodes {
			msg = append(msg, []byte(xmlNode.OutputXML(false))...)
			msg = append(msg, "\n"...)
		}
	}
	return
}

func genQueryResult(ctx context.Context, submitterc chan []byte, url string, pathStr string, logger log.Logger) (chan []byte, chan error) {
	out := make(chan []byte)
	errc := make(chan error)
	go func() {
		startTime := time.Now()

		defer close(out)
		defer close(errc)

		rawMsg, err := dataFetch(url)
		if err != nil {
			logger.Error(err)
			errc <- err
			return
		}
		msgReturn, err := dataParse(rawMsg, pathStr)
		if err != nil {
			logger.Error(err)
			errc <- err
			return
		}
		logger.TimeTrack(startTime, "TFetch", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})
		select {
		case submitter, ok := <-submitterc:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "GenQueryResult", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})

			// signed message = concat(msgReturn, submitter address)
			msgReturn = append(msgReturn, submitter...)
			select {
			case out <- msgReturn:
			case <-ctx.Done():
			}
			return
		case <-ctx.Done():
			return
		}
	}()
	return out, errc
}

func genSign(ctx context.Context, contentc chan []byte, sec *share.PriShare, suite suites.Suite, sign *vss.Signature, logger log.Logger) (chan *vss.Signature, chan error) {
	out := make(chan *vss.Signature)
	errc := make(chan error)
	go func() {
		defer close(out)
		defer close(errc)

		select {
		case content, ok := <-contentc:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "GenSign", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})
			sign.Content = content
			sig, err := tbls.Sign(suite, sec, content)
			if err != nil {
				logger.Error(err)
				select {
				case errc <- err:
				case <-ctx.Done():
				}
				return
			}
			sign.Signature = sig
			select {
			case <-ctx.Done():
			case out <- sign:
			}
			return
		case <-ctx.Done():
			return
		}
	}()
	return out, errc
}

func dispatchSign(ctx context.Context, submitterc chan []byte, signc chan *vss.Signature, reqSignc chan request, p p2p.P2PInterface, requestID []byte, threshold int, logger log.Logger) chan *vss.Signature {
	out := make(chan *vss.Signature)
	go func() {
		select {
		case submitter, ok := <-submitterc:
			defer logger.TimeTrack(time.Now(), "dispatchSign", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})

			if !ok {
				return
			}
			if r := bytes.Compare(p.GetID(), submitter); r != 0 {
				logger.Event("dispatchSignToSubmitter", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})
				//Send signature to submitter
				logger.Info(fmt.Sprintf("dispatchSign to  %x", submitter))
				select {
				case <-ctx.Done():
				case sign := <-signc:
					logger.Event("dispatchSendToSubmitter", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})
					if _, err := p.Request(ctx, submitter, sign); err != nil {
						logger.Error(err)
					}
				}
				close(out)
				return
			}
		case <-ctx.Done():
			close(out)
			return
		}
		logger.Event("dispatchSignWaitForSign", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})

		//Bypass signature to next stage
		select {
		case <-ctx.Done():
			close(out)
		case sign := <-signc:
			select {
			case <-ctx.Done():
				close(out)
			case out <- sign:
			}
		}
		//Request for peer's signature
		req := request{ctx: ctx, requestID: string(requestID), threshold: threshold, reply: out}
		select {
		case <-ctx.Done():
		case reqSignc <- req:
		}
	}()
	return out
}

func recoverSign(ctx context.Context, signc chan *vss.Signature, suite suites.Suite, pubPoly *share.PubPoly, nbThreshold int, nbParticipants int, logger log.Logger) (chan *vss.Signature, chan error) {
	out := make(chan *vss.Signature)
	errc := make(chan error)
	go func() {
		var signShares [][]byte
		defer close(out)
		defer close(errc)

		for {
			select {
			case sign, ok := <-signc:
				if !ok {
					return
				}
				logger.Event("recoverSign", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})

				logger.Info(fmt.Sprintf("recoverSign %d", len(signShares)))
				if len(signShares) == 0 {
					defer logger.TimeTrack(time.Now(), "RecoverSign", map[string]interface{}{"GroupID": ctx.Value(ctxKey("GroupID")), "RequestID": ctx.Value(ctxKey("RequestID"))})
				}

				if sign == nil || sign.Signature == nil || sign.Content == nil {
					err := errors.New("Detected nil pointer and skipped")
					logger.Error(err)
					errc <- err
					continue
				}

				signShares = append(signShares, sign.Signature)
				if len(signShares) >= nbThreshold {
					sig, err := tbls.Recover(suite, pubPoly, sign.Content, signShares, nbThreshold, nbParticipants)
					if err != nil {
						logger.Error(err)
						errc <- err
						continue
					}

					if err = bls.Verify(suite, pubPoly.Commit(), sign.Content, sig); err != nil {
						logger.Error(err)
						errc <- err
						continue
					}
					x, y := sign.ToBigInt()
					logger.Info(fmt.Sprintf("Verify success signature %s %s", x.String(), y.String()))

					//Contract will append sender address to content to verify if it is a right submitter
					t := len(sign.Content) - addrLen
					if t < 0 {
						errc <- errors.New("length of content less than 0")
					}

					queryResult := make([]byte, t)
					copy(queryResult, sign.Content)
					select {
					case out <- &vss.Signature{Index: sign.Index, RequestId: sign.RequestId, Content: queryResult, Signature: sig}:
					case <-ctx.Done():
					}
					return
				}

			case <-ctx.Done():
				return
			}
		}
	}()
	return out, errc
}

func registerGroup(ctx context.Context, chain onchain.ProxyAdapter, IdWithPubKeys chan [5]*big.Int) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		var err error
		select {
		case idPubkey, ok := <-IdWithPubKeys:
			if ok {
				err = chain.RegisterGroupPubKey(idPubkey)
			} else {
				err = errors.New("no publickey")
			}
		case <-ctx.Done():
			err = ctx.Err()
			return
		}
		if err != nil {
			select {
			case errc <- err:
			case <-ctx.Done():
			}
		}
	}()
	return
}
func reportQueryResult(ctx context.Context, chain onchain.ProxyAdapter, queryType uint32, signC chan *vss.Signature) (errc chan error) {
	errc = make(chan error)
	go func() {
		defer close(errc)
		var err error
		select {
		case signature, ok := <-signC:
			if ok {
				if queryType == onchain.TrafficSystemRandom {
					err = chain.UpdateRandomness(signature)
				} else {
					err = chain.DataReturn(signature)
				}
			} else {
				err = errors.New("no signature")
			}
		case <-ctx.Done():
			return
		}
		if err != nil {
			select {
			case errc <- err:
			case <-ctx.Done():
			}
		}
	}()
	return
}
func padOrTrim(bb []byte, size int) []byte {
	l := len(bb)
	if l == size {
		return bb
	}
	if l > size {
		return bb[l-size:]
	}
	tmp := make([]byte, size)
	copy(tmp[size-l:], bb)
	return tmp
}
