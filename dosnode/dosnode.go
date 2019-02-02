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

	"github.com/oliveagle/jsonpath"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"

	"github.com/ethereum/go-ethereum/common"

	"github.com/DOSNetwork/core/suites"
	"github.com/antchfx/xmlquery"
)

type DosNode struct {
	suite  suites.Suite
	chain  onchain.ChainInterface
	dkg    dkg.P2PDkgInterface
	p      p2p.P2PInterface
	logger log.Logger
	done   chan interface{}
	signc  chan *vss.Signature
}

func NewDosNode(suite suites.Suite, p p2p.P2PInterface, chain onchain.ChainInterface, dkg dkg.P2PDkgInterface) (dosNode *DosNode) {
	return &DosNode{
		suite:  suite,
		p:      p,
		chain:  chain,
		dkg:    dkg,
		logger: log.New("module", "dosclient"),
		done:   make(chan interface{}),
		signc:  make(chan *vss.Signature, 50),
	}
}

func (d *DosNode) Start() (err error) {
	if err = d.listen(); err != nil {
		d.logger.Error(err)
		return
	}
	//Register nodeID to onchain
	if err = d.chain.UploadID(); err != nil {
		d.logger.Error(err)
	}
	return
}

func (d *DosNode) End() {
	close(d.done)
}

func MergeErrors(cs ...<-chan error) <-chan error {
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
		fmt.Println("End fanin")

		close(multiplexedStream)
	}()

	return multiplexedStream
}

func (d *DosNode) waitForGrouping(errs ...<-chan error) {
	errc := MergeErrors(errs...)
	for err := range errc {
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (d *DosNode) choseSubmitter(ctx context.Context, lastSysRand *big.Int, ids [][]byte, outCount int) ([]chan []byte, <-chan error) {
	x := int(lastSysRand.Uint64())
	if x < 0 {
		x = 0 - x
	}
	y := len(ids)
	submitter := x % y
	fmt.Println("checkBalance submitter ", submitter)
	reChose := 0
	out := make(chan []byte)
	checkBalance := make(chan int, 1)
	checkAlive := make(chan int, 1)
	errc := make(chan error, 1)
	var outs []chan []byte
	for i := 0; i < outCount; i++ {
		outs = append(outs, make(chan []byte, 1))
	}
	go func() {
		defer close(out)
		defer close(checkBalance)
		defer close(checkAlive)
		defer close(errc)
		defer fmt.Println("End choseSubmitter")

		for {
			select {
			case idx := <-checkBalance:
				fmt.Println("checkBalance ")
				if reChose == y {
					errc <- errors.New("No Suitable Submitter")
					return
				}
				if !d.chain.EnoughBalance(common.BytesToAddress(ids[idx])) {
					reChose++
					idx = (idx + 1) % y
					checkBalance <- idx
				} else {
					checkAlive <- idx
				}
			case idx := <-checkAlive:
				fmt.Println("checkAlive ")

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
func (d *DosNode) requestSign(
	ctx context.Context,
	submitterc <-chan []byte,
	requestId string,
	trafficType uint32,
	id []byte) (<-chan *vss.Signature, <-chan error) {
	out := make(chan *vss.Signature, 1)
	errc := make(chan error, 1)
	//ticker := time.NewTicker(3 * time.Second)
	retry := make(chan bool, 1)
	go func() {
		defer close(retry)
		defer close(out)
		defer close(errc)
		defer fmt.Println("End requestSign")
		retryCount := 0
		submitter := <-submitterc
		if r := bytes.Compare(d.chain.GetId(), submitter); r != 0 {
			fmt.Println("requestSign : Not a Submitter ")
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
				fmt.Println("requestSign from ", id)
				retryCount++
				msg, err := d.p.Request(id, sign)
				if err != nil {
					fmt.Println("requestSign err ", err)
					d.logger.Error(err)
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
						fmt.Println("requestSign : Got a signature from ", id)
						out <- sign
						return
					default:
						fmt.Println("requestSign : not *vss.Signature:")
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, errc
}

func (d *DosNode) generateRandom(
	ctx context.Context,
	submitterc <-chan []byte,
	requestId []byte,
	lastSysRand []byte,
	userSeed []byte,
) <-chan []byte {
	out := make(chan []byte, 1)
	go func() {
		defer close(out)
		select {
		case submitter := <-submitterc:
			// signed message: concat(requestId, lastSystemRandom, userSeed, submitter address)
			random := append(requestId, lastSysRand...)
			random = append(random, userSeed...)
			random = append(random, submitter...)
			out <- random
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

func (d *DosNode) genQueryResult(ctx context.Context, url string, pathStr string) (<-chan []byte, chan error) {
	out := make(chan []byte, 1)
	errc := make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errc)
		rawMsg, err := dataFetch(url)
		if err != nil {
			fmt.Println("genQueryResult err", err)

			errc <- err
			return
		}
		msgReturn, err := dataParse(rawMsg, pathStr)
		if err != nil {
			fmt.Println("genQueryResult err", err)

			errc <- err
			return
		}
		fmt.Println("genQueryResult ", msgReturn)
		out <- msgReturn
	}()
	return out, errc
}
func (d *DosNode) generateSign(
	ctx context.Context,
	submitterc <-chan []byte,
	contentc <-chan []byte,
	signc chan *vss.Signature,
	pubkey [4]*big.Int,
	requestId *big.Int,
	index uint32) (<-chan *vss.Signature, <-chan error) {
	out := make(chan *vss.Signature, 1)
	errc := make(chan error, 1)
	wait := make(chan bool, 1)
	go func() {
		var submitter []byte
		var content []byte
		defer close(out)
		defer close(errc)
		defer close(wait)
		defer fmt.Println("End generateSign")
		for {
			select {
			case value, ok := <-submitterc:
				if ok {
					submitter = value
					fmt.Println("generateSign from submitterc", len(submitter), len(content))
					if len(submitter) != 0 && len(content) != 0 {
						wait <- true
					}
				}
			case value, ok := <-contentc:
				if ok {
					content = value
					fmt.Println("generateSign from contentc ", len(submitter), len(content))

					if len(submitter) != 0 && len(content) != 0 {
						wait <- true
					}
				}
			case <-wait:

				sig, err := tbls.Sign(d.suite, d.dkg.GetShareSecurity(pubkey), content)
				if err != nil {
					d.logger.Error(err)
					errc <- err
				}

				sign := &vss.Signature{
					Index:     index,
					QueryId:   requestId.String(),
					Content:   content,
					Signature: sig,
				}
				signc <- sign
				if r := bytes.Compare(d.chain.GetId(), submitter); r == 0 {
					fmt.Println("generateSign : I'm a submitter ")
					out <- sign
				} else {
					fmt.Println("generateSign I'm  not a submitter ")

				}
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, errc
}

func (d *DosNode) recoverSignature(ctx context.Context, signc <-chan *vss.Signature, pubPoly *share.PubPoly, nbThreshold int, nbParticipants int) (<-chan *vss.Signature, <-chan error) {
	out := make(chan *vss.Signature, 1)
	errc := make(chan error, 1)
	go func() {
		var signShares [][]byte
		defer close(out)
		defer close(errc)
		defer fmt.Println("End recoverSignature!!!!!!!")

		for {
			select {
			case sign, ok := <-signc:
				if ok {
					signShares = append(signShares, sign.Signature)
					fmt.Println("recoverSignature receive a sign.Index =", sign.Index)

					if len(signShares) >= nbThreshold {
						sig, err := tbls.Recover(
							d.suite,
							pubPoly,
							sign.Content,
							signShares,
							nbThreshold,
							nbParticipants)
						if err != nil {
							fmt.Println("recoverSignatur Recover err", err)

							continue
						}

						if err = bls.Verify(
							d.suite,
							pubPoly.Commit(),
							sign.Content,
							sig); err != nil {
							fmt.Println("recoverSignatur Verifyerr", err)
							continue
						}
						x, y := onchain.DecodeSig(sig)
						fmt.Println("Verify success signature ", x.String(), y.String())

						out <- &vss.Signature{
							Index:     sign.Index,
							QueryId:   sign.QueryId,
							Content:   sign.Content,
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

func (d *DosNode) listen() (err error) {
	eventGrouping := make(chan interface{})
	chUrl := make(chan interface{})
	chRandom := make(chan interface{})
	chUsrRandom := make(chan interface{})
	eventValidation := make(chan interface{})
	if err = d.chain.SubscribeEvent(eventGrouping, onchain.SubscribeDOSProxyLogGrouping); err != nil {
		return err
	}

	if err = d.chain.SubscribeEvent(chUrl, onchain.SubscribeDOSProxyLogUrl); err != nil {
		return err
	}

	if err = d.chain.SubscribeEvent(chRandom, onchain.SubscribeDOSProxyLogUpdateRandom); err != nil {
		return err
	}

	if err = d.chain.SubscribeEvent(chUsrRandom, onchain.SubscribeDOSProxyLogRequestUserRandom); err != nil {
		return err
	}

	if err = d.chain.SubscribeEvent(eventValidation, onchain.SubscribeDOSProxyLogValidationResult); err != nil {
		return err
	}
	peerEvent, err := d.p.SubscribeEvent(100, vss.Signature{})

	go func() {
		groupSession := make(map[string]context.CancelFunc)
		userRandRequests := make(map[string]context.CancelFunc)
		peerSignMap := make(map[string]*vss.Signature)
		defer close(eventGrouping)
		defer close(chUrl)
		defer close(chRandom)
		defer close(chUsrRandom)
		defer close(eventValidation)
		defer d.p.UnSubscribeEvent(vss.Signature{})
		for {
			select {
			case msg := <-d.signc:
				fmt.Println("Pipeline Output : msg ", msg)
				peerSignMap[msg.QueryId] = msg
			case msg := <-peerEvent:
				//fmt.Println("peerEvent : msg ")
				switch content := msg.Msg.Message.(type) {
				case *vss.Signature:
					d.p.Reply(msg.Sender, msg.RequestNonce, peerSignMap[content.QueryId])
				}
			case msg := <-eventGrouping:
				fmt.Println("EthLog : eventGrouping")

				content, ok := msg.(*onchain.DOSProxyLogGrouping)
				if !ok {
					log.Error(err)
				}
				isMember := false
				var groupIds [][]byte
				for _, node := range content.NodeId {
					id := node.Bytes()
					if r := bytes.Compare(d.chain.GetId(), id); r == 0 {
						isMember = true
					}
					fmt.Println("EthLog : eventGrouping compare ", bytes.Compare(d.chain.GetId(), id))

					groupIds = append(groupIds, id)
				}
				fmt.Println("EthLog : eventGrouping ", isMember, " ", content.Removed)

				if isMember {
					sessionId := dkg.GIdsToSessionID(groupIds)
					if !content.Removed {
						ctx, cancelFunc := context.WithCancel(context.Background())
						groupSession[sessionId] = cancelFunc
						fmt.Println("EthLog : groupIds", groupIds, " isMember ", isMember)
						var errcList []<-chan error
						outFromDkg, errc := d.dkg.Start(ctx, groupIds, sessionId)
						errcList = append(errcList, errc)
						errc = d.chain.UploadPubKey(ctx, outFromDkg)
						errcList = append(errcList, errc)
						go d.waitForGrouping(errcList...)
					} else { //if chain reorg then call
						if groupSession[sessionId] != nil {
							groupSession[sessionId]()
							groupSession[sessionId] = nil
						}
					}
				}
			case msg := <-chUsrRandom:
				content, ok := msg.(*onchain.DOSProxyLogRequestUserRandom)
				if !ok {
					log.Error(err)
				}
				if d.isMember(content.DispatchedGroup) {
					fmt.Println("EthLog : userRandom")
					ctx, cancelFunc := context.WithCancel(context.Background())
					userRandRequests[content.RequestId.String()] = cancelFunc

					var errcList []<-chan error
					ids := d.dkg.GetGroupIDs(content.DispatchedGroup)
					lastRand := content.LastSystemRandomness
					requestID := content.RequestId
					pubkey := content.DispatchedGroup
					useSeed := content.UserSeed
					var signShares []<-chan *vss.Signature
					pubPoly := d.dkg.GetGroupPublicPoly(pubkey)

					//Build a pipeline
					submitterc, errc := d.choseSubmitter(ctx, lastRand, ids, len(ids)+1)
					errcList = append(errcList, errc)
					for i, id := range ids {
						var signc <-chan *vss.Signature
						var errc <-chan error
						if r := bytes.Compare(d.chain.GetId(), id); r != 0 {
							signc, errc = d.requestSign(ctx, submitterc[i], requestID.String(), uint32(onchain.TrafficUserRandom), id)
						} else {
							contentc := d.generateRandom(ctx, submitterc[len(ids)], requestID.Bytes(), lastRand.Bytes(), useSeed.Bytes())
							signc, errc = d.generateSign(ctx, submitterc[i], contentc, d.signc, pubkey, requestID, uint32(onchain.TrafficUserRandom))
						}
						signShares = append(signShares, signc)
						errcList = append(errcList, errc)
					}
					signc, errc := d.recoverSignature(ctx, fanIn(ctx, signShares...), pubPoly, (len(ids)/2 + 1), len(ids))
					errcList = append(errcList, errc)
					errc = d.chain.DataReturn(ctx, signc)
					errcList = append(errcList, errc)
				}
			case msg := <-chUrl:
				content, ok := msg.(*onchain.DOSProxyLogUrl)
				if !ok {
					log.Error(err)
				}
				if d.isMember(content.DispatchedGroup) {
					fmt.Println("EthLog : queryLog")
					ctx, cancelFunc := context.WithCancel(context.Background())
					userRandRequests[content.QueryId.String()] = cancelFunc
					var errcList []<-chan error

					lastRand := content.Randomness
					url := content.DataSource
					selector := content.Selector
					requestID := content.QueryId
					ids := d.dkg.GetGroupIDs(content.DispatchedGroup)
					pubkey := content.DispatchedGroup
					var signShares []<-chan *vss.Signature
					pubPoly := d.dkg.GetGroupPublicPoly(pubkey)

					//Build a pipeline
					submitterc, errc := d.choseSubmitter(ctx, lastRand, ids, len(ids))
					errcList = append(errcList, errc)
					for i, id := range ids {
						var signc <-chan *vss.Signature
						var contentc <-chan []byte
						var errc <-chan error
						if r := bytes.Compare(d.chain.GetId(), id); r != 0 {
							signc, errc = d.requestSign(ctx, submitterc[i], requestID.String(), uint32(onchain.TrafficUserQuery), id)
						} else {
							contentc, errc = d.genQueryResult(ctx, url, selector)
							errcList = append(errcList, errc)
							signc, errc = d.generateSign(ctx, submitterc[i], contentc, d.signc, pubkey, requestID, uint32(onchain.TrafficUserQuery))
						}
						signShares = append(signShares, signc)
						errcList = append(errcList, errc)
					}
					signc, errc := d.recoverSignature(ctx, fanIn(ctx, signShares...), pubPoly, (len(ids)/2 + 1), len(ids))
					errcList = append(errcList, errc)
					errc = d.chain.DataReturn(ctx, signc)
					errcList = append(errcList, errc)
				}

			case msg := <-chRandom:
				content, ok := msg.(*onchain.DOSProxyLogUpdateRandom)
				if !ok {
					log.Error(err)
				}
				if d.isMember(content.DispatchedGroup) {
					fmt.Println("EthLog : systemRandom")

				}

			case msg := <-eventValidation:
				content, ok := msg.(*onchain.DOSProxyLogValidationResult)
				if !ok {
					log.Error(err)
				}
				_ = content
				if d.isMember(content.PubKey) {
					fmt.Println("EthLog : validationResult", content.Pass)
				}
			case <-d.done:
				return
			default:
			}
		}
	}()
	return
}

func (d *DosNode) isMember(pubkey [4]*big.Int) bool {
	return d.dkg.GetShareSecurity(pubkey) != nil
}
