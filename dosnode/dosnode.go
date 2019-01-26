package dosnode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/antchfx/xmlquery"

	"github.com/ethereum/go-ethereum/common"

	"github.com/oliveagle/jsonpath"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
)

// TODO: Move constants to some unified places.
const (
	WATCHDOGTIMEOUT   = 20
	VALIDATIONTIMEOUT = 20
	RANDOMNUMBERSIZE  = 32
	UINT256SIZE       = 32
	NETMSGTIMEOUT     = 20
)

type DosNodeInterface interface {
	PipeGrouping(<-chan interface{})
	PipeQueries(queries ...<-chan interface{}) <-chan Report
	PipeSignAndBroadcast(reports <-chan Report) (<-chan Report, <-chan Report)
	PipeRecoverAndVerify(chan p2p.P2PMessage, <-chan Report) <-chan Report
	PipeSendToOnchain(chReport <-chan Report) <-chan Report
	PipeCleanFinishMap(chValidation chan interface{}, request ...<-chan Report) <-chan interface{}
	SetParticipants(groupIds [][]byte)
}

type DosNode struct {
	suite          suites.Suite
	quit           chan struct{}
	nbParticipants int
	nbThreshold    int
	groupPubPoly   share.PubPoly
	shareSec       share.PriShare
	chainConn      onchain.ChainInterface
	p2pDkg         dkg.P2PDkgInterface
	network        *p2p.P2PInterface
	groupIds       [][]byte
	timeCostMap    *sync.Map
	requestTracker *sync.Map
	logger         log.Logger
}

type Report struct {
	submitter  []byte
	signShares [][]byte
	signGroup  []byte
	selfSign   vss.Signature
	timeStamp  time.Time
	//For retrigger
	backupTask interface{}
}

type timeRecord struct {
	lastUpdateTime    time.Time
	sendToChannelTime time.Time
	dataProcessCost   float64
	dataSignCost      float64
	dataChannelCost   float64
	dataUploadCost    float64
	dataConfirmCost   float64
	trafficType       uint8
	pass              bool
	requestTx         string
	requestBlkNb      uint64
	replyTx           string
	replyBlkNb        uint64
}

type request struct {
	lastUpdateTime time.Time
	version        uint8
}

func CreateDosNode(suite suites.Suite, p p2p.P2PInterface, chainConn onchain.ChainInterface, p2pDkg dkg.P2PDkgInterface) (dosNode DosNodeInterface) {
	return &DosNode{
		suite:          suite,
		quit:           make(chan struct{}),
		network:        &p,
		chainConn:      chainConn,
		p2pDkg:         p2pDkg,
		timeCostMap:    new(sync.Map),
		requestTracker: new(sync.Map),
		logger:         log.New("module", "dosclient"),
	}
}

func (d *DosNode) SetParticipants(groupIds [][]byte) {
	d.groupIds = groupIds
	d.nbParticipants = len(d.groupIds)
	d.nbThreshold = d.nbParticipants/2 + 1
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

func (d *DosNode) isMember(dispatchedGroup [4]*big.Int) bool {
	if d.p2pDkg.IsCertified() {
		temp := append(padOrTrim(dispatchedGroup[2].Bytes(), UINT256SIZE), padOrTrim(dispatchedGroup[3].Bytes(), UINT256SIZE)...)
		temp = append(padOrTrim(dispatchedGroup[1].Bytes(), UINT256SIZE), temp...)
		temp = append(padOrTrim(dispatchedGroup[0].Bytes(), UINT256SIZE), temp...)
		temp = append([]byte{0x01}, temp...)

		groupPub, err := d.p2pDkg.GetGroupPublicPoly().Commit().MarshalBinary()
		if err != nil {
			return false
		}

		fmt.Println(groupPub)
		fmt.Println(temp)

		if r := bytes.Compare(groupPub, temp); r == 0 {
			fmt.Println("isMember TRUE")
			return true
		}
	}
	fmt.Println("isMember false, isCertified:", d.p2pDkg.IsCertified())
	return false
}

func (d *DosNode) choseSubmitter(r *big.Int) (id []byte) {
	x := int(r.Uint64())
	if x < 0 {
		x = 0 - x
	}
	y := len(d.groupIds)
	submitter := x % y
	//TODO:Check to see if submitter is alive
	for !d.chainConn.EnoughBalance(common.BytesToAddress(d.groupIds[submitter])) {
		fmt.Println("choose next submitter due to balance low")
		submitter = (submitter + 1) % y
	}
	id = d.groupIds[submitter]
	fmt.Println("choseSubmitter ", id)
	return
}

// Note: Signed messages keep synced with on-chain contracts
func (d *DosNode) PipeGrouping(chGroup <-chan interface{}) {
	go func() {
		for {
			select {
			//event from blockChain
			case msg := <-chGroup:
				switch content := msg.(type) {
				case *onchain.DOSProxyLogGrouping:
					isMember := false
					var groupIds [][]byte
					for i, node := range content.NodeId {
						id := node.Bytes()
						if string(id) == string((*d.network).GetID()) {
							isMember = true
						}
						fmt.Println("DOSProxyLogGrouping member i= ", i, " id ", id, " ", isMember)
						groupIds = append(groupIds, id)
					}
					d.SetParticipants(groupIds)
					if isMember {
						dkgRes, errc := d.p2pDkg.Start(&dkg.DkgSession{SessionId: "0", GroupIds: groupIds})
						go func() {
							for err := range errc {
								fmt.Println(err)
							}
						}()
						select {
						case certified := <-dkgRes:
							if certified == true {
								if err := d.chainConn.UploadPubKey(d.p2pDkg.GetGroupPublicPoly().Commit()); err != nil {
								}
							}
						}
					}
				default:
					fmt.Println("DOSProxyLogGrouping type mismatch")
				}
			case <-d.quit:
				break
			}
		}
	}()
}

// TODO: error handling and logging
// Note: Signed messages keep synced with on-chain contracts
func (d *DosNode) PipeQueries(queries ...<-chan interface{}) <-chan Report {
	merged := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(queries))
	for _, c := range queries {
		go func(c <-chan interface{}) {
			for v := range c {
				merged <- v
			}
			wg.Done()
		}(c)
	}
	out := make(chan Report)

	go func() {
		for {
			select {
			//event from blockChain
			case msg := <-merged:
				switch content := msg.(type) {
				case *onchain.DOSProxyLogUrl:
					//Check to see if this request is for node's group
					if d.isMember(content.DispatchedGroup) {
						preVersion, loaded := d.requestTracker.LoadOrStore(content.QueryId.String(), &request{lastUpdateTime: time.Now(), version: 0})
						if loaded {
							preVersion.(*request).version++
							preVersion.(*request).lastUpdateTime = time.Now()
						}
						queryId := content.QueryId.String() + ":" + strconv.Itoa(int(preVersion.(*request).version))
						fmt.Println("PipeCheckURL!!", queryId)
						submitter := d.choseSubmitter(content.Randomness)
						newTimeRecord := &timeRecord{
							lastUpdateTime: time.Now(),
							requestTx:      content.Tx,
							requestBlkNb:   content.BlockN,
						}

						url := content.DataSource
						rawMsg, err := dataFetch(url)
						if err != nil {
						}
						fmt.Println("fetched data: ", rawMsg, string(rawMsg))

						msgReturn, err := dataParse(rawMsg, content.Selector)
						if err != nil {
						}
						fmt.Println("Data to return:", msgReturn, string(msgReturn))

						newTimeRecord.dataProcessCost = time.Since(newTimeRecord.lastUpdateTime).Seconds()
						newTimeRecord.lastUpdateTime = time.Now()
						d.timeCostMap.Store(queryId, newTimeRecord)

						// signed message = concat(msgReturn, submitter address)
						msgReturn = append(msgReturn, submitter...)
						//combine result with submitter
						sign := &vss.Signature{
							Index:   uint32(onchain.TrafficUserQuery),
							QueryId: queryId,
							Content: msgReturn,
						}
						report := &Report{}
						report.selfSign = *sign
						report.submitter = submitter
						report.backupTask = *content
						out <- *report
					}
				case *onchain.DOSProxyLogRequestUserRandom:
					//Check to see if this request is for node's group
					if d.isMember(content.DispatchedGroup) {
						preVersion, loaded := d.requestTracker.LoadOrStore(content.RequestId.String(), &request{lastUpdateTime: time.Now(), version: 0})
						if loaded {
							preVersion.(*request).version++
							preVersion.(*request).lastUpdateTime = time.Now()
						}
						requestId := content.RequestId.String() + ":" + strconv.Itoa(int(preVersion.(*request).version))
						fmt.Println("PipeUserRandom!!", requestId)
						submitter := d.choseSubmitter(content.LastSystemRandomness)
						d.timeCostMap.Store(requestId, &timeRecord{
							lastUpdateTime: time.Now(),
							requestTx:      content.Tx,
							requestBlkNb:   content.BlockN,
						})

						// signed message: concat(requestId, lastSystemRandom, userSeed, submitter address)
						msgReturn := append(content.RequestId.Bytes(), content.LastSystemRandomness.Bytes()...)
						msgReturn = append(msgReturn, content.UserSeed.Bytes()...)
						msgReturn = append(msgReturn, submitter...)
						//combine result with submitter
						sign := &vss.Signature{
							Index:   uint32(onchain.TrafficUserRandom),
							QueryId: requestId,
							Content: msgReturn,
						}
						report := &Report{}
						report.selfSign = *sign
						report.submitter = submitter
						report.backupTask = *content
						out <- *report
					}
				case *onchain.DOSProxyLogUpdateRandom:
					if d.isMember(content.DispatchedGroup) {
						preRandom := content.LastRandomness
						fmt.Printf("1 ) GenerateRandomNumber preRandom #%v \n", preRandom)

						preVersion, loaded := d.requestTracker.LoadOrStore(preRandom.String(), &request{lastUpdateTime: time.Now(), version: 0})
						if loaded {
							preVersion.(*request).version++
							preVersion.(*request).lastUpdateTime = time.Now()
						}
						requestId := preRandom.String() + ":" + strconv.Itoa(int(preVersion.(*request).version))
						d.timeCostMap.Store(requestId, &timeRecord{
							lastUpdateTime: time.Now(),
							requestTx:      content.Tx,
							requestBlkNb:   content.BlockN,
						})

						submitter := d.choseSubmitter(preRandom)
						paddedRandomBytes := padOrTrim(preRandom.Bytes(), RANDOMNUMBERSIZE)
						randomNum := append(paddedRandomBytes, submitter...)
						fmt.Println("GenerateRandomNumber content = ", randomNum)

						sign := &vss.Signature{
							Index:   uint32(onchain.TrafficSystemRandom),
							QueryId: requestId,
							Content: randomNum,
						}
						report := &Report{}
						report.selfSign = *sign
						report.submitter = submitter
						out <- *report
					}
				default:
				}
			case <-d.quit:
				wg.Wait()
				close(merged)
				close(out)
				break
			}
		}
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

func (d *DosNode) PipeSignAndBroadcast(reports <-chan Report) (<-chan Report, <-chan Report) {
	outForRecover := make(chan Report, 100)
	outForValidate := make(chan Report, 100)

	go func() {
		for {
			select {
			case report := <-reports:
				fmt.Println("2) PipeSignAndBroadcast")
				sign := report.selfSign
				sig, err := tbls.Sign(d.suite, d.p2pDkg.GetShareSecurity(), sign.Content)
				if err != nil {
					continue
				}

				sign.Signature = sig
				report.signShares = append(report.signShares, sig)
				if r := bytes.Compare(d.chainConn.GetId(), report.submitter); r == 0 {
					outForRecover <- report
				} else {
					for _, member := range d.groupIds {
						if r := bytes.Compare(member, report.submitter); r == 0 {
							//Todo:Need to check to see if it is thread safe
							memberAddress := common.BytesToAddress(member).Hex()
							if err = (*d.network).SendMessage(member, &sign); err != nil {

							} else {
								fmt.Println("send to ", memberAddress)
							}
							break
						}
					}
					outForValidate <- report
				}
			case <-d.quit:
				close(outForValidate)
				close(outForRecover)
				break
			}
		}
	}()

	return outForRecover, outForValidate
}

//receiveSignature is thread safe.
func (d *DosNode) PipeRecoverAndVerify(cSignatureFromPeer chan p2p.P2PMessage, fromEth <-chan Report) <-chan Report {
	outForSubmit := make(chan Report, 100)
	reportMap := map[string]Report{}
	signatureAwait := map[string]time.Time{}

	go func() {
		for {
			select {
			case report := <-fromEth:
				fmt.Println("3) PipeRecoverAndVerify fromEth")
				reportMap[report.selfSign.QueryId] = report
			case msg := <-cSignatureFromPeer:
				if err := (*d.network).Reply(msg.Sender, msg.RequestNonce, &vss.Signature{}); err != nil {
				}
				sign, ok := msg.Msg.Message.(*vss.Signature)
				if !ok {
					continue
				}
				var sigShares [][]byte
				if report, ok := reportMap[sign.QueryId]; ok {
					delete(signatureAwait, sign.QueryId)

					//TODO:Should check if content is always the same
					if r := bytes.Compare(report.selfSign.Content, sign.Content); r != 0 {
						fmt.Println("report query id ", report.selfSign.QueryId)
						fmt.Println("report Content ", report.selfSign.Content, string(report.selfSign.Content))
						fmt.Println("sign query id ", sign.QueryId)
						fmt.Println("sign Content ", sign.Content, string(sign.Content))
					}

					report.signShares = append(report.signShares, sign.Signature)
					sigShares = report.signShares
					if report.signGroup == nil && len(sigShares) >= d.nbThreshold {
						fmt.Println("4) PipeRecoverAndVerify recover")
						sig, err := tbls.Recover(d.suite, d.p2pDkg.GetGroupPublicPoly(), sign.Content, sigShares, d.nbThreshold, d.nbParticipants)
						if err != nil {
							fmt.Println("recover failed!!!!!!!!!!!!!!!!! report Content ", report.selfSign.Content, string(report.selfSign.Content))
							fmt.Println("recover failed!!!!!!!!!!!!!!!!! len(sigShares) = ", len(sigShares), " Content ", sign.Content, string(sign.Content))
							fmt.Println("recover failed!!!!!!!!!!!!!!!!!")
							continue
						}

						if err = bls.Verify(d.suite, d.p2pDkg.GetGroupPublicPoly().Commit(), sign.Content, sig); err != nil {
							fmt.Println("Verify failed!!!!!!!!!!!!!!!!!")
							continue
						}

						x, y := onchain.DecodeSig(sig)
						fmt.Println("5) Verify success signature ", x.String(), y.String())
						report.signGroup = sig
						record, _ := d.timeCostMap.Load(report.selfSign.QueryId)
						record.(*timeRecord).dataSignCost = time.Since(record.(*timeRecord).lastUpdateTime).Seconds()
						record.(*timeRecord).lastUpdateTime = time.Now()
						record.(*timeRecord).sendToChannelTime = time.Now()

						fmt.Println("I'm submitter !!!!!!!!!!!!!!!!!!!", sign.QueryId)
						fmt.Println("I'm submitter !!!!!!!!!!!!!!!!!!!", d.chainConn.GetId())
						fmt.Println("I'm submitter !!!!!!!!!!!!!!!!!!!", report.submitter)

						report.selfSign.Content = sign.Content
						outForSubmit <- report
						fmt.Println("submit to outForSubmit", sign.QueryId)
						delete(reportMap, sign.QueryId)
					} else {
						reportMap[sign.QueryId] = report
					}
				} else {
					//Other nodes has received blockchain event and send signature to other nodes
					//Node put back these signatures until it received blockchain event.
					//TODO:It should use a timeStamp that can be used to see if it is a stale signature
					if preTime, reenter := signatureAwait[sign.QueryId]; reenter {
						if time.Since(preTime).Minutes() >= NETMSGTIMEOUT {
							fmt.Println("remove stale signature")
							delete(signatureAwait, sign.QueryId)
							continue
						}
					} else {
						signatureAwait[sign.QueryId] = time.Now()
					}
					go func() {
						cSignatureFromPeer <- msg
					}()
				}
			case <-d.quit:
				close(outForSubmit)
				break
			}
		}
	}()
	return outForSubmit
}

func (d *DosNode) PipeSendToOnchain(chReport <-chan Report) <-chan Report {
	outForValidate := make(chan Report, 100)

	go func() {
		for {
			select {
			case report := <-chReport:
				record, _ := d.timeCostMap.Load(report.selfSign.QueryId)
				record.(*timeRecord).dataChannelCost = time.Since(record.(*timeRecord).sendToChannelTime).Seconds()

				qIDArray := strings.Split(report.selfSign.QueryId, ":")
				qID, succ := new(big.Int).SetString(qIDArray[0], 10)
				if !succ {
				}
				qVersion, err := strconv.Atoi(qIDArray[1])
				if err != nil {
				}

				switch report.selfSign.Index {
				case onchain.TrafficSystemRandom:
					fmt.Println("PipeSendToOnchain sysRandom QueryId = ", report.selfSign.QueryId)
					if err := d.chainConn.SetRandomNum(report.signGroup, uint8(qVersion)); err != nil {
					} else {
						fmt.Println("randomNumber Set for ", report.selfSign.QueryId)
						record.(*timeRecord).dataUploadCost = time.Since(record.(*timeRecord).sendToChannelTime).Seconds() - record.(*timeRecord).dataChannelCost
						outForValidate <- report
						fmt.Println("submit to outForValidate", report.selfSign.QueryId)
					}
				default:
					fmt.Println("PipeSendToOnchain URL/usrRandom QueryId = ", report.selfSign.QueryId)

					t := len(report.selfSign.Content) - 20
					if t < 0 {
						fmt.Println("Error : length of content less than 0", t)
					}

					queryResult := make([]byte, t)
					copy(queryResult, report.selfSign.Content)

					/*//Test Case 3
					tmp := big.NewInt(20)
					if report.backupTask.Timeout.Cmp(tmp) > 0 {
					   //x := []byte{184, 194}
					   queryResult[0] ^= 0x01
					}
					*/
					//TODO:chainCoo should use a sendToOnChain(protobuf message) instead of DataReturn with mutex
					//sendToOnChain receive a message from channel then call the corresponding function
					if err = d.chainConn.DataReturn(qID, uint8(report.selfSign.Index), queryResult, report.signGroup, uint8(qVersion)); err != nil {
					} else {
						fmt.Println("urlCallback Set for ", report.selfSign.QueryId)
						record.(*timeRecord).dataUploadCost = time.Since(record.(*timeRecord).sendToChannelTime).Seconds() - record.(*timeRecord).dataChannelCost
						outForValidate <- report
						fmt.Println("submit to outForValidate", report.selfSign.QueryId)
					}
				}
			case <-d.quit:
				break
			}
		}
	}()

	return outForValidate
}

func (d *DosNode) PipeCleanFinishMap(chValidation chan interface{}, request ...<-chan Report) <-chan interface{} {
	forValidate := make(chan Report)
	var wg sync.WaitGroup
	wg.Add(len(request))
	for _, c := range request {
		go func(c <-chan Report) {
			for v := range c {
				forValidate <- v
			}
			wg.Done()
		}(c)
	}

	ticker := time.NewTicker(2 * time.Minute)
	validateMap := map[string]Report{}
	chValidationAwait := map[string]time.Time{}
	lastValidatedRandom := time.Now()
	queries := make(chan interface{})
	go func() {
		for {
			select {
			case report := <-forValidate:
				report.timeStamp = time.Now()
				validateMap[report.selfSign.QueryId] = report
			case msg := <-chValidation:
				switch content := msg.(type) {
				case *onchain.DOSProxyLogValidationResult:
					trafficId := content.TrafficId.String() + ":" + strconv.Itoa(int(content.Version))
					if report, match := validateMap[trafficId]; match {
						record, _ := d.timeCostMap.Load(trafficId)
						record.(*timeRecord).dataConfirmCost = time.Since(record.(*timeRecord).lastUpdateTime).Seconds()
						record.(*timeRecord).trafficType = content.TrafficType
						record.(*timeRecord).pass = content.Pass
						d.timeCostMap.Delete(trafficId)
						delete(chValidationAwait, trafficId)

						if !content.Pass {
							fmt.Println("event DOSProxyLogValidationResult========================")
							fmt.Println("DOSProxyLogValidationResult pass ", content.Pass)
							fmt.Println("DOSProxyLogValidationResult TrafficType ", content.TrafficType)
							fmt.Println("DOSProxyLogValidationResult TrafficId ", content.TrafficId)
							fmt.Println("DOSProxyLogValidationResult Signature ", content.Signature)
							fmt.Println("DOSProxyLogValidationResult GroupKey ", content.PubKey)
							fmt.Println("DOSProxyLogValidationResult Message ", content.Message)
							fmt.Println("DOSProxyLogValidationResult Version ", content.Version)
							x, y, z, t, _ := onchain.DecodePubKey(d.p2pDkg.GetGroupPublicPoly().Commit())
							fmt.Println("GroupKey ", x.String(), y.String(), z.String(), t.String())
							fmt.Println("Signature", report.selfSign.Signature)
							fmt.Println("Content", report.selfSign.Content)

							//switch uint32(content.TrafficType) {
							//case ForRandomNumber:
							fmt.Println("Invalide Signature.........")
							_ = report
							//default:
							// fmt.Println("===========================================")
							// fmt.Println("Retrigger query......... time-out", report.backupTask.Timeout)
							// tmp := big.NewInt(0)
							// if report.backupTask.Timeout.Cmp(tmp) == 0 {
							//    fmt.Println("give up this query.........")
							// }
							// tmp = big.NewInt(1)
							// report.backupTask.Randomness = report.backupTask.Randomness.Sub(report.backupTask.Randomness, tmp)
							// report.backupTask.Timeout = report.backupTask.Timeout.Sub(report.backupTask.Timeout, tmp)
							// chUrl <- &blockchain.DOSProxyLogUrl{
							//    QueryId:         report.backupTask.QueryId,
							//    Url:             report.backupTask.Url,
							//    Timeout:         report.backupTask.Timeout,
							//    Randomness:      report.backupTask.Randomness,
							//    DispatchedGroup: report.backupTask.DispatchedGroup,
							// }
							//}
						} else {
							fmt.Println("Validated ", trafficId)
							if content.TrafficType == onchain.TrafficSystemRandom {
								lastValidatedRandom = time.Now()
							}
						}
						delete(validateMap, trafficId)
					} else {
						if preTime, reenter := chValidationAwait[trafficId]; reenter {
							if time.Since(preTime).Minutes() >= NETMSGTIMEOUT {
								fmt.Println("remove stale chValidation", trafficId)
								delete(chValidationAwait, trafficId)
								continue
							}
						} else {
							chValidationAwait[trafficId] = time.Now()
						}
						go func() {
							chValidation <- msg
						}()
					}
				default:
					fmt.Println("DOSProxyLogValidationResult type mismatch")
				}
			case <-ticker.C:
				if d.p2pDkg.IsCertified() {
					dur := time.Since(lastValidatedRandom)
					if dur.Minutes() >= WATCHDOGTIMEOUT {
						fmt.Println("WatchDog Random timeout.........")
						if err := d.chainConn.RandomNumberTimeOut(); err != nil {

						}
					}
				}
				for queryId, report := range validateMap {
					fmt.Println("Time to check and clean .........")
					dur := time.Since(report.timeStamp)
					//If a submitter is not alive then 3 minutes timeout will be trigged
					if dur.Minutes() >= VALIDATIONTIMEOUT {

						fmt.Println("Validation timeout.........")

						delete(validateMap, queryId)
					}
				}
				d.timeCostMap.Range(d.checkLog)
				d.requestTracker.Range(d.checkLog)
			case <-d.quit:
				close(queries)
				ticker.Stop()
				return
			}
		}
	}()
	return queries
}

func (d *DosNode) checkLog(key interface{}, record interface{}) (deleted bool) {
	switch record.(type) {
	case *timeRecord:
		if time.Since(record.(*timeRecord).lastUpdateTime).Minutes() > VALIDATIONTIMEOUT {

			d.timeCostMap.Delete(key)
		}
	case *request:
		if time.Since(record.(*request).lastUpdateTime).Minutes() > VALIDATIONTIMEOUT {

			d.requestTracker.Delete(key)
		}
	}
	return true
}
