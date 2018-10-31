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

	"github.com/oliveagle/jsonpath"

	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
)

type DosNodeInterface interface {
	PipeQueries(queries ...<-chan interface{}) <-chan Report
	PipeSignAndBroadcast(reports <-chan Report) <-chan Report
	PipeRecoverAndVerify(chan vss.Signature, <-chan Report) (<-chan Report, <-chan Report)
	PipeSendToOnchain(<-chan Report)
	PipeCleanFinishMap(chan interface{}, <-chan Report) <-chan interface{}
	PipeGrouping(<-chan interface{})
	SetParticipants(int)
}

// TODO: Move constants to some unified places.
const WATCHDOGTIMEOUT = 20
const VALIDATIONTIMEOUT = 20
const RANDOMNUMBERSIZE = 32
const NETMSGTIMEOUT = 20

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

func CreateDosNode(suite suites.Suite, nbParticipants int, p p2p.P2PInterface, chainConn onchain.ChainInterface, p2pDkg dkg.P2PDkgInterface) (DosNodeInterface, error) {
	d := &DosNode{
		suite:          suite,
		quit:           make(chan struct{}),
		nbParticipants: nbParticipants,
		nbThreshold:    nbParticipants/2 + 1,
		network:        &p,
		chainConn:      chainConn,
		p2pDkg:         p2pDkg,
	}
	return d, nil
}

func (d *DosNode) SetParticipants(nbParticipants int) {
	d.nbParticipants = nbParticipants
	d.nbThreshold = nbParticipants/2 + 1
}

func dataFetch(url string) (body []byte, err error) {
	sTime := time.Now()
	client := &http.Client{Timeout: 60 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	r.Body.Close()

	fmt.Println("fetched data: ", string(body))
	fmt.Println("dataFetch End:	took", time.Now().Sub(sTime))

	return

}

func dataParse(rawMsg []byte, pathStr string) (msg []byte, err error) {
	if pathStr == "" {
		msg = rawMsg
	} else if strings.HasPrefix(pathStr, "$") {
		var rawMsgJson, msgJson interface{}
		err = json.Unmarshal(rawMsg, &rawMsgJson)
		if err != nil {
			return
		}

		msgJson, err = jsonpath.JsonPathLookup(rawMsgJson, pathStr)
		if err != nil {
			return
		}

		switch content := msgJson.(type) {
		case interface{}:
			msg = dataExtract(content)
		}
	} else if strings.HasPrefix(pathStr, "/") {
		var rawMsgXml *xmlquery.Node
		rawMsgXml, err = xmlquery.Parse(bytes.NewReader(rawMsg))
		if err != nil {
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

func dataExtract(rawData interface{}) (dataBytes []byte) {
	switch content := rawData.(type) {
	case float64:
		dataString := strconv.FormatFloat(content, 'f', -1, 64)
		dataBytes = []byte(dataString)
	case string:
		dataBytes = []byte(content)
	case []interface{}:
		for _, rawData := range content {
			switch innerContent := rawData.(type) {
			case interface{}:
				dataBytes = append(dataBytes, dataExtract(innerContent)...)
				dataBytes = append(dataBytes, "\n"...)
			}
		}
	}
	return
}

func (d *DosNode) isMember(dispatchedGroup [4]*big.Int) bool {
	if d.p2pDkg.IsCetified() {
		temp := append(dispatchedGroup[2].Bytes(), dispatchedGroup[3].Bytes()...)
		temp = append(dispatchedGroup[1].Bytes(), temp...)
		temp = append(dispatchedGroup[0].Bytes(), temp...)
		temp = append([]byte{0x01}, temp...)

		groupPub, err := d.p2pDkg.GetGroupPublicPoly().Commit().MarshalBinary()
		if err != nil {
			fmt.Println(err)
			return false
		}
		r := bytes.Compare(groupPub, temp)
		if r == 0 {
			fmt.Println("isMember TRUE")
			return true
		}
	}
	fmt.Println("isMember false, isCertified:", d.p2pDkg.IsCetified())
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
						if string(id) == string((*d.network).GetId().Id) {
							isMember = true
						}
						fmt.Println("DOSProxyLogGrouping member i= ", i, " id ", id, " ", isMember)
						groupIds = append(groupIds, id)
					}
					nbParticipants := len(groupIds)
					d.SetParticipants(nbParticipants)
					d.p2pDkg.SetNbParticipants(nbParticipants)
					if isMember {
						d.groupIds = groupIds
						d.p2pDkg.SetGroupMembers(groupIds)
						d.p2pDkg.RunDKG()
					}
				default:
					fmt.Println("DOSProxyLogUpdateRandom type mismatch")
				}
			case <-d.quit:
				break
			}
		}
	}()
	return
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
						queryId := content.QueryId
						fmt.Println("PipeCheckURL!!", queryId)
						url := content.DataSource
						submitter := d.choseSubmitter(content.Randomness)
						rawMsg, err := dataFetch(url)
						if err != nil {
							fmt.Println(err)
						}
						msgReturn, err := dataParse(rawMsg, content.Selector)
						// signed message = concat(msgReturn, submitter address)
						msgReturn = append(msgReturn, submitter...)
						//combine result with submitter
						sign := &vss.Signature{
							Index:   uint32(onchain.TrafficUserQuery),
							QueryId: queryId.String(),
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
						requestId := content.RequestId
						fmt.Println("PipeUserRandom!!", requestId)
						submitter := d.choseSubmitter(content.LastSystemRandomness)
						// signed message: concat(requestId, lastSystemRandom, userSeed, submitter address)
						msgReturn := append(content.RequestId.Bytes(), content.LastSystemRandomness.Bytes()...)
						msgReturn = append(msgReturn, content.UserSeed.Bytes()...)
						msgReturn = append(msgReturn, submitter...)
						//combine result with submitter
						sign := &vss.Signature{
							Index:   uint32(onchain.TrafficUserRandom),
							QueryId: requestId.String(),
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
						fmt.Println("1 ) GenerateRandomNumber preRandom ", preRandom)

						submitter := d.choseSubmitter(preRandom)
						paddedRandomBytes := padOrTrim(preRandom.Bytes(), RANDOMNUMBERSIZE)
						randomNum := append(paddedRandomBytes, submitter...)
						fmt.Println("GenerateRandomNumber content = ", randomNum)

						sign := &vss.Signature{
							Index:   uint32(onchain.TrafficSystemRandom),
							QueryId: preRandom.String(),
							Content: randomNum,
						}
						report := &Report{}
						report.selfSign = *sign
						report.submitter = submitter
						out <- *report
					}
				default:
					fmt.Println("query type mismatch", msg)
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

func (d *DosNode) PipeSignAndBroadcast(reports <-chan Report) <-chan Report {
	out := make(chan Report)

	go func() {
		for {
			select {
			case report := <-reports:
				fmt.Println("2) PipeSignAndBroadcast")
				sign := report.selfSign
				sig, err := tbls.Sign(d.suite, d.p2pDkg.GetShareSecuirty(), sign.Content)
				if err != nil {
					fmt.Println(err)
					continue
				}
				sign.Signature = sig
				report.signShares = append(report.signShares, sig)
				out <- report
				for _, member := range d.groupIds {
					if string(member) != string((*d.network).GetId().Id) {
						//Todo:Need to check to see if it is thread safe
						(*d.network).SendMessageById(member, &sign)
						fmt.Println("send to ", member)
					}
				}

			case <-d.quit:
				close(out)
				break
			}
		}
	}()

	return out
}

//receiveSignature is thread safe.
func (d *DosNode) PipeRecoverAndVerify(cSignatureFromPeer chan vss.Signature, fromEth <-chan Report) (<-chan Report, <-chan Report) {
	outForSubmit := make(chan Report, 100)
	outForValidate := make(chan Report, 100)
	reportMap := map[string]Report{}
	signatureAwait := map[string]time.Time{}

	go func() {
		for {
			select {
			case report := <-fromEth:
				fmt.Println("3) PipeRecoverAndVerify fromEth")
				reportMap[report.selfSign.QueryId] = report
			case sign := <-cSignatureFromPeer:
				var sigShares [][]byte
				signIdentityBytes := append([]byte(sign.QueryId), sign.Signature...)
				signIdentity := new(big.Int).SetBytes(signIdentityBytes).String()
				if report, ok := reportMap[sign.QueryId]; ok {
					delete(signatureAwait, signIdentity)
					//TODO:Should check if content is always the same
					r := bytes.Compare(report.selfSign.Content, sign.Content)
					if r != 0 {
						fmt.Println("report query id ", report.selfSign.QueryId)
						fmt.Println("report Content ", string(report.selfSign.Content))
						fmt.Println("sign query id ", sign.QueryId)
						fmt.Println("sign Content ", string(sign.Content))
					}

					report.signShares = append(report.signShares, sign.Signature)

					sigShares = report.signShares
					if report.signGroup == nil && len(sigShares) >= d.nbThreshold {
						fmt.Println("4) PipeRecoverAndVerify recover")
						sig, err := tbls.Recover(d.suite, d.p2pDkg.GetGroupPublicPoly(), sign.Content, sigShares, d.nbThreshold, d.nbParticipants)
						if err != nil {
							fmt.Println(err)
							fmt.Println("recover failed!!!!!!!!!!!!!!!!! report Content ", string(report.selfSign.Content))
							fmt.Println("recover failed!!!!!!!!!!!!!!!!! len(sigShares) = ", len(sigShares), " Content ", string(sign.Content))
							fmt.Println("recover failed!!!!!!!!!!!!!!!!!")
							continue
						}
						err = bls.Verify(d.suite, d.p2pDkg.GetGroupPublicPoly().Commit(), sign.Content, sig)
						if err != nil {
							fmt.Println(err)
							fmt.Println("Verify failed!!!!!!!!!!!!!!!!!")
							continue
						}
						x, y := onchain.DecodeSig(sig)
						fmt.Println("5) Verify success signature ", x.String(), y.String())
						report.signGroup = sig
						report.timeStamp = time.Now()
						r := bytes.Compare(d.chainConn.GetId(), report.submitter)
						if r == 0 {
							fmt.Println("I'm submitter !!!!!!!!!!!!!!!!!!!", sign.QueryId)
							fmt.Println("I'm submitter !!!!!!!!!!!!!!!!!!!", d.chainConn.GetId())
							fmt.Println("I'm submitter !!!!!!!!!!!!!!!!!!!", report.submitter)

							report.selfSign.Content = sign.Content
							outForSubmit <- report
							fmt.Println("submit to outForSubmit", sign.QueryId)

						}
						outForValidate <- report
						fmt.Println("submit to outForValidate", sign.QueryId)
						delete(reportMap, sign.QueryId)
					} else {
						reportMap[sign.QueryId] = report
					}
				} else {
					//Other nodes has received blockchain event and send signature to other nodes
					//Node put back these signatures until it received blockchain event.
					//TODO:It should use a timeStamp that can be used to see if it is a stale signature
					if preTime, reenter := signatureAwait[signIdentity]; reenter {
						if time.Since(preTime).Minutes() >= NETMSGTIMEOUT {
							fmt.Println("remove stale signature")
							delete(signatureAwait, signIdentity)
							continue
						}
					} else {
						signatureAwait[signIdentity] = time.Now()
					}

					go func() {
						cSignatureFromPeer <- sign
					}()
				}
			case <-d.quit:
				close(outForSubmit)
				close(outForValidate)
				break
			}
		}
	}()
	return outForSubmit, outForValidate
}

func (d *DosNode) PipeSendToOnchain(chReport <-chan Report) {
	go func() {
		for {
			select {
			case report := <-chReport:
				switch report.selfSign.Index {
				case onchain.TrafficSystemRandom:
					err := d.chainConn.SetRandomNum(report.signGroup)
					if err != nil {
						fmt.Println("SetRandomNum err ", err)
					} else {
						fmt.Println("randomNumber Set for ", report.selfSign.QueryId)
					}
				default:
					fmt.Println("PipeSendToOnchain URL/usrRandom result = ", report.selfSign.Content)
					qID := big.NewInt(0)
					qID.SetString(report.selfSign.QueryId, 10)

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
					err := d.chainConn.DataReturn(qID, uint8(report.selfSign.Index), queryResult, report.signGroup)
					if err != nil {
						fmt.Println("DataReturn err ", err)
					}
					fmt.Println("urlCallback Set for ", report.selfSign.QueryId)
				}
			case <-d.quit:
				break
			}
		}
	}()
	return
}

func (d *DosNode) PipeCleanFinishMap(chValidation chan interface{}, forValidate <-chan Report) <-chan interface{} {
	ticker := time.NewTicker(2 * time.Minute)
	validateMap := map[string]Report{}
	chValidationAwait := map[interface{}]time.Time{}
	lastValidatedRandom := time.Now()
	queries := make(chan interface{})
	go func() {
		for {
			select {
			case report := <-forValidate:
				validateMap[report.selfSign.QueryId] = report
			case msg := <-chValidation:
				switch content := msg.(type) {
				case *onchain.DOSProxyLogValidationResult:
					if report, match := validateMap[content.TrafficId.String()]; match {
						fmt.Println("event DOSProxyLogValidationResult========================")
						fmt.Println("DOSProxyLogValidationResult pass ", content.Pass)
						fmt.Println("DOSProxyLogValidationResult TrafficType ", content.TrafficType)
						fmt.Println("DOSProxyLogValidationResult TrafficId ", content.TrafficId)
						fmt.Println("DOSProxyLogValidationResult Signature ", content.Signature)
						fmt.Println("DOSProxyLogValidationResult GroupKey ", content.PubKey)
						fmt.Println("DOSProxyLogValidationResult Message ", content.Message)
						x, y, z, t, _ := onchain.DecodePubKey(d.p2pDkg.GetGroupPublicPoly().Commit())
						fmt.Println("GroupKey ", x.String(), y.String(), z.String(), t.String())
						delete(chValidationAwait, msg)
						if !content.Pass {
							//switch uint32(content.TrafficType) {
							//case ForRandomNumber:
							fmt.Println("Invalide Signature.........")
							_ = report
							//default:
							//	fmt.Println("===========================================")
							//	fmt.Println("Retrigger query......... time-out", report.backupTask.Timeout)
							//	tmp := big.NewInt(0)
							//	if report.backupTask.Timeout.Cmp(tmp) == 0 {
							//		fmt.Println("give up this query.........")
							//	}
							//	tmp = big.NewInt(1)
							//	report.backupTask.Randomness = report.backupTask.Randomness.Sub(report.backupTask.Randomness, tmp)
							//	report.backupTask.Timeout = report.backupTask.Timeout.Sub(report.backupTask.Timeout, tmp)
							//	chUrl <- &blockchain.DOSProxyLogUrl{
							//		QueryId:         report.backupTask.QueryId,
							//		Url:             report.backupTask.Url,
							//		Timeout:         report.backupTask.Timeout,
							//		Randomness:      report.backupTask.Randomness,
							//		DispatchedGroup: report.backupTask.DispatchedGroup,
							//	}
							//}
						} else {
							fmt.Println("Validated ", content.TrafficId.String())
							if content.TrafficType == onchain.TrafficSystemRandom {
								lastValidatedRandom = time.Now()
							}
						}
						delete(validateMap, content.TrafficId.String())
					} else {
						if preTime, reenter := chValidationAwait[msg]; reenter {
							if time.Since(preTime).Minutes() >= NETMSGTIMEOUT {
								fmt.Println("remove stale chValidation")
								delete(chValidationAwait, msg)
								continue
							}
						} else {
							chValidationAwait[msg] = time.Now()
						}
						go func() {
							chValidation <- msg
						}()
					}
				default:
					fmt.Println("DOSProxyLogValidationResult type mismatch")
				}
			case <-ticker.C:
				if d.p2pDkg.IsCetified() {
					dur := time.Since(lastValidatedRandom)
					if dur.Minutes() >= WATCHDOGTIMEOUT {
						fmt.Println("WatchDog Random timeout.........")
						d.chainConn.RandomNumberTimeOut()
					}
				}
				for queryId, report := range validateMap {
					fmt.Println("Time to check and clean .........")
					dur := time.Since(report.timeStamp)
					//If a submitter is not alive then 3 minutes timeout will be trigged
					if dur.Minutes() >= VALIDATIONTIMEOUT {
						//switch report.selfSign.Index {
						//case ForRandomNumber:
						fmt.Println("Validation timeout.........")
						//default:
						//	fmt.Println("Retrigger query.........")
						//	tmp := big.NewInt(0)
						//	if report.backupTask.Timeout.Cmp(tmp) == 0 {
						//		fmt.Println("give up this query.........")
						//	}
						//	tmp = big.NewInt(1)
						//	report.backupTask.Randomness = report.backupTask.Randomness.Sub(report.backupTask.Randomness, tmp)
						//	report.backupTask.Timeout = report.backupTask.Timeout.Sub(report.backupTask.Timeout, tmp)
						//	chUrl <- &blockchain.DOSProxyLogUrl{
						//		QueryId:         report.backupTask.QueryId,
						//		Url:             report.backupTask.Url,
						//		Timeout:         report.backupTask.Timeout,
						//		Randomness:      report.backupTask.Randomness,
						//		DispatchedGroup: report.backupTask.DispatchedGroup,
						//	}
						//}
						delete(validateMap, queryId)
					}
				}
			case <-d.quit:
				close(queries)
				ticker.Stop()
				return
			}
		}
	}()
	return queries
}
