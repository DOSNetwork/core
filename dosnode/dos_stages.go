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
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
	"github.com/antchfx/xmlquery"
	"github.com/ethereum/go-ethereum/common"
	"github.com/oliveagle/jsonpath"
)

const (
	RANDOMNUMBERSIZE = 32
	ADDRESSLENGTH    = 20
)

func mergeErrors(cs ...<-chan error) <-chan error {
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
			out <- n
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
		defer logger.Event("EndFanIn", nil)

		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func choseSubmitter(ctx context.Context, chain onchain.ChainInterface, lastSysRand *big.Int, ids [][]byte, outCount int) ([]chan []byte, <-chan error) {
	x := int(lastSysRand.Uint64())
	if x < 0 {
		x = 0 - x
	}
	y := len(ids)
	submitter := x % y
	reChose := 0
	checkBalance := make(chan int, 1)
	checkAlive := make(chan int, 1)
	errc := make(chan error, 1)
	var outs []chan []byte
	for i := 0; i < outCount; i++ {
		outs = append(outs, make(chan []byte, 1))
	}
	go func() {
		defer close(checkBalance)
		defer close(checkAlive)
		defer close(errc)

		for {
			select {
			case idx := <-checkBalance:
				fmt.Println("checkBalance ")
				if reChose == y {
					errc <- errors.New("No Suitable Submitter")
					return
				}
				if !chain.EnoughBalance(common.BytesToAddress(ids[idx])) {
					reChose++
					idx = (idx + 1) % y
					checkBalance <- idx
				} else {
					checkAlive <- idx
				}
			case idx := <-checkAlive:

				if reChose == y {
					errc <- errors.New("No Suitable Submitter")
					return
				}
				//TODO: Use ping/pong to check if submitter is alive
				//out <- ids[idx]
				for _, out := range outs {
					out <- ids[idx]
					close(out)
				}
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	checkBalance <- submitter
	return outs, errc
}
func requestSign(
	ctx context.Context,
	submitterc <-chan []byte,
	p p2p.P2PInterface,
	nodeId []byte,
	requestId string,
	trafficType uint32,
	id []byte) (<-chan *vss.Signature, <-chan error) {
	out := make(chan *vss.Signature)
	errc := make(chan error)
	retry := make(chan bool, 1)
	go func() {
		defer close(retry)
		defer close(out)
		defer close(errc)

		retryCount := 0
		submitter := <-submitterc
		if r := bytes.Compare(nodeId, submitter); r != 0 {
			return
		}

		sign := &vss.Signature{
			Index:   trafficType,
			QueryId: requestId,
		}
		retry <- true
		for {
			select {
			case <-retry:
				retryCount++
				msg, err := p.Request(id, sign)
				if err != nil {
					if retryCount < 5 {
						retry <- true
					} else {
						return
					}
				} else {
					switch content := msg.(type) {
					case *vss.Signature:
						sign.Content = content.Content
						sign.Signature = content.Signature
						out <- sign
						return
					default:
					}
				}
			case <-ctx.Done():
				return
			}
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
		case submitter := <-submitterc:
			// signed message: concat(requestId, lastSystemRandom, userSeed, submitter address)
			random := append(requestId, lastSysRand...)
			random = append(random, userSeed...)
			random = append(random, submitter...)
			out <- random
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
		case submitter := <-submitterc:
			// signed message: concat(lastSystemRandom, submitter address)
			paddedLastSysRand := padOrTrim(lastSysRand, RANDOMNUMBERSIZE)
			random := append(paddedLastSysRand, submitter...)
			out <- random
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
		defer close(out)
		defer close(errc)

		rawMsg, err := dataFetch(url)
		if err != nil {
			errc <- err
			return
		}
		msgReturn, err := dataParse(rawMsg, pathStr)
		if err != nil {
			errc <- err
			return
		}
		select {
		case submitter := <-submitterc:
			// signed message = concat(msgReturn, submitter address)
			msgReturn = append(msgReturn, submitter...)
			out <- msgReturn
			return
		case <-ctx.Done():
			return
		}
	}()
	return out, errc
}
func genSign(
	ctx context.Context,
	submitterc <-chan []byte,
	contentc <-chan []byte,
	signc chan *vss.Signature,
	dkg dkg.P2PDkgInterface,
	suite suites.Suite,
	nodeID []byte,
	pubkey [4]*big.Int,
	requestId string,
	index uint32) (<-chan *vss.Signature, <-chan error) {
	out := make(chan *vss.Signature)
	errc := make(chan error)
	wait := make(chan bool, 1)
	go func() {
		var submitter []byte
		var content []byte
		defer close(out)
		defer close(errc)
		defer close(wait)

		for {
			select {
			case value, ok := <-submitterc:
				if ok {
					submitter = value
					if len(submitter) != 0 && len(content) != 0 {
						wait <- true
					}
				}
			case value, ok := <-contentc:
				if ok {
					content = value

					if len(submitter) != 0 && len(content) != 0 {
						wait <- true
					}
				}
			case <-wait:

				sig, err := tbls.Sign(suite, dkg.GetShareSecurity(pubkey), content)
				if err != nil {
					errc <- err
				}

				sign := &vss.Signature{
					Index:     index,
					QueryId:   requestId,
					Content:   content,
					Signature: sig,
				}
				signc <- sign
				if r := bytes.Compare(nodeID, submitter); r == 0 {
					out <- sign
				}
				return
			case <-ctx.Done():
				return
			}
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
				if ok {
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
							continue
						}

						if err = bls.Verify(
							suite,
							pubPoly.Commit(),
							sign.Content,
							sig); err != nil {
							continue
						}
						x, y := onchain.DecodeSig(sig)
						fmt.Println("Verify success signature ", x.String(), y.String())

						//Contract will append sender address to content to verify if it is a right submitter
						t := len(sign.Content) - ADDRESSLENGTH
						if t < 0 {
							fmt.Println("Error : length of content less than 0", t)
						}

						queryResult := make([]byte, t)
						copy(queryResult, sign.Content)

						out <- &vss.Signature{
							Index:     sign.Index,
							QueryId:   sign.QueryId,
							Content:   queryResult,
							Signature: sig,
						}
					}
					if len(signShares) == nbParticipants {
						return
					}
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
