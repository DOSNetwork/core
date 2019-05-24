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

	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
	"github.com/antchfx/xmlquery"
	"github.com/oliveagle/jsonpath"
)

const (
	RANDOMNUMBERSIZE = 32
	ADDRESSLENGTH    = 20
)

func mergeErrors(ctx context.Context, cs ...<-chan error) <-chan error {
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

func fanIn(ctx context.Context, channels ...<-chan *vss.Signature) <-chan *vss.Signature {
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
func choseSubmitter(ctx context.Context, p p2p.P2PInterface, lastSysRand *big.Int, ids [][]byte, outCount int) ([]chan []byte, <-chan error) {
	errc := make(chan error)
	var outs []chan []byte
	for i := 0; i < outCount; i++ {
		outs = append(outs, make(chan []byte, 1))
	}

	go func() {
		defer close(errc)
		start := time.Now()
		lastRand := int(lastSysRand.Uint64())
		if lastRand < 0 {
			lastRand = 0 - lastRand
		}

		submitter := -1
		//Check to see if submitter is reachable
		for i := 0; i < len(ids); i++ {
			idx := (lastRand + i) % len(ids)
			if !bytes.Equal(p.GetID(), ids[idx]) {
				if _, err := p.ConnectTo("", ids[i]); err != nil {
					continue
				}
			}
			submitter = idx
			break
		}

		if submitter == -1 {
			select {
			case errc <- errors.New("No reachable submitter"):
			case <-ctx.Done():
			}
		} else {
			for _, out := range outs {
				select {
				case out <- ids[submitter]:
				case <-ctx.Done():
				}
			}
		}
		for _, out := range outs {
			close(out)
		}
		logger.TimeTrack(start, "ChoseSubmitter", map[string]interface{}{"submitter": submitter, "GroupID": ctx.Value("GroupID"), "RequestID": ctx.Value("RequestID")})
		return
	}()
	return outs, errc
}

func requestSign(
	ctx context.Context,
	submitterc <-chan []byte,
	contentc <-chan []byte,
	p p2p.P2PInterface,
	nodeId []byte,
	requestId []byte,
	trafficType uint32,
	id []byte,
	nonce []byte) (<-chan *vss.Signature, <-chan error) {
	out := make(chan *vss.Signature)
	errc := make(chan error)
	go func() {
		defer close(out)
		defer close(errc)

		select {
		case submitter, ok := <-submitterc:
			if !ok {
				return
			}
			if r := bytes.Compare(nodeId, submitter); r != 0 {
				return
			}
			defer logger.TimeTrack(time.Now(), "RequestSign", map[string]interface{}{"GroupID": ctx.Value("GroupID"), "RequestID": ctx.Value("RequestID")})
			fmt.Println("requestSign nonce ", nonce)
			sign := &vss.Signature{
				Index:     trafficType,
				RequestId: requestId,
				Nonce:     nonce,
			}

			retryCount := 0
			for retryCount < 30 {
				if msg, err := p.Request(id, sign); err == nil {
					switch content := msg.Msg.Message.(type) {
					case *vss.Signature:
						sign.Content = content.Content
						sign.Signature = content.Signature
						select {
						case out <- sign:
						case <-ctx.Done():
						}
						return
					default:
					}
				}
				retryCount++
			}
			err := errors.New("Retry limit exceeded")
			logger.Error(err)
			select {
			case errc <- err:
			case <-ctx.Done():
			}
			return
		case <-ctx.Done():
			return
		}
	}()
	return out, errc
}

func genSign(
	ctx context.Context,
	contentc <-chan []byte,
	cSignToPeer chan *vss.Signature,
	sec *share.PriShare,
	suite suites.Suite,
	nodeID []byte,
	groupID string,
	requestId []byte,
	index uint32,
	nonce []byte) (<-chan *vss.Signature, <-chan error) {
	out := make(chan *vss.Signature)
	errc := make(chan error)
	go func() {
		var submitter []byte
		var content []byte
		defer close(out)
		defer close(errc)

		select {
		case value, ok := <-contentc:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "GenSign", map[string]interface{}{"GroupID": ctx.Value("GroupID"), "RequestID": ctx.Value("RequestID")})

			content = value
			submitter = content[len(content)-20:]

			sig, err := tbls.Sign(suite, sec, content)
			if err != nil {
				logger.Error(err)
				select {
				case errc <- err:
				case <-ctx.Done():
					return
				}
			}

			sign := &vss.Signature{
				Index:     index,
				RequestId: requestId,
				Nonce:     nonce,
				Content:   content,
				Signature: sig,
			}
			select {
			case cSignToPeer <- sign:
			case <-ctx.Done():
			}
			if r := bytes.Compare(nodeID, submitter); r == 0 {
				select {
				case out <- sign:
				case <-ctx.Done():
				}
			}
			return

		case <-ctx.Done():
			return
		}
	}()
	return out, errc
}

func genUserRandom(
	ctx context.Context,
	submitterc <-chan []byte,
	requestId []byte,
	lastSysRand []byte,
	userSeed []byte,
) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)

		select {
		case submitter, ok := <-submitterc:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "GenUserRandom", map[string]interface{}{"GroupID": ctx.Value("GroupID"), "RequestID": ctx.Value("RequestID")})

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

func genSysRandom(
	ctx context.Context,
	submitterc <-chan []byte,
	lastSysRand []byte,
) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)

		select {
		case submitter, ok := <-submitterc:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "GenSysRandom", map[string]interface{}{"GroupID": ctx.Value("GroupID"), "RequestID": ctx.Value("RequestID")})

			// signed message: concat(lastSystemRandom, submitter address)
			paddedLastSysRand := padOrTrim(lastSysRand, RANDOMNUMBERSIZE)
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
		var rawMsgJson, msgJson interface{}
		if err = json.Unmarshal(rawMsg, &rawMsgJson); err != nil {
			return
		}

		if msgJson, err = jsonpath.JsonPathLookup(rawMsgJson, pathStr); err != nil {
			return
		}

		msg, err = json.Marshal(msgJson)
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

func genQueryResult(ctx context.Context, submitterc chan []byte, url string, pathStr string) (<-chan []byte, chan error) {
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
		logger.TimeTrack(startTime, "TFetch", map[string]interface{}{"GroupID": ctx.Value("GroupID"), "RequestID": ctx.Value("RequestID")})
		select {
		case submitter, ok := <-submitterc:
			if !ok {
				return
			}
			defer logger.TimeTrack(time.Now(), "GenQueryResult", map[string]interface{}{"RequestID": ctx.Value("RequestID")})

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

func recoverSign(ctx context.Context, signc <-chan *vss.Signature, suite suites.Suite, pubPoly *share.PubPoly, nbThreshold int, nbParticipants int) (<-chan *vss.Signature, <-chan error) {
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
				if len(signShares) == 0 {
					defer logger.TimeTrack(time.Now(), "RecoverSign", map[string]interface{}{"GroupID": ctx.Value("GroupID"), "RequestID": ctx.Value("RequestID")})
				}

				signShares = append(signShares, sign.Signature)

				if len(signShares) >= nbThreshold {
					sig, err := tbls.Recover(
						suite,
						pubPoly,
						sign.Content,
						signShares,
						nbThreshold,
						nbParticipants)
					if err != nil {
						logger.Error(err)
						errc <- err
						continue
					}

					if err = bls.Verify(
						suite,
						pubPoly.Commit(),
						sign.Content,
						sig); err != nil {
						logger.Error(err)
						errc <- err
						continue
					}
					x, y := onchain.DecodeSig(sig)
					fmt.Println("Verify success signature ", x.String(), y.String())

					//Contract will append sender address to content to verify if it is a right submitter
					t := len(sign.Content) - ADDRESSLENGTH
					if t < 0 {
						errc <- errors.New("length of content less than 0")
					}

					queryResult := make([]byte, t)
					copy(queryResult, sign.Content)
					select {
					case out <- &vss.Signature{
						Index:     sign.Index,
						RequestId: sign.RequestId,
						Content:   queryResult,
						Signature: sig,
					}:
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
