package dosnode

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
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
)

type DosNodeInterface interface {
	PipeCheckURL(<-chan interface{}) <-chan Report
	PipeGenerateRandomNumber(<-chan interface{}) <-chan Report
	PipeSignAndBroadcast(reports ...<-chan Report) <-chan Report
	PipeRecoverAndVerify(chan vss.Signature, <-chan Report) (<-chan Report, <-chan Report)
	PipeSendToOnchain(<-chan Report)
	PipeCleanFinishMap(chan interface{}, <-chan Report) <-chan interface{}
	PipeGrouping(<-chan interface{})
	SetParticipants(int)
}

const WATCHDOGTIMEOUT = 20
const VALIDATIONTIMEOUT = 20
const RANDOMNUMBERSIZE = 32
const NETMSGTIMEOUT = 20

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

const (
	ForRandomNumber = uint32(1)
	ForCheckURL     = uint32(0)
)

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
	backupTask onchain.DOSProxyLogUrl
}

func (d *DosNode) SetParticipants(nbParticipants int) {
	d.nbParticipants = nbParticipants
	d.nbThreshold = nbParticipants/2 + 1
}
func (d *DosNode) dataFetch(url string) ([]byte, error) {
	sTime := time.Now()
	client := &http.Client{Timeout: 60 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body.Close()

	fmt.Println("fetched data: ", string(body))
	fmt.Println("dataFetch End:	took", time.Now().Sub(sTime))

	return body, nil

}

func (d *DosNode) isMember(dispatchedGroup [4]*big.Int) bool {
	if d.p2pDkg.IsCetified() {
		temp := append(dispatchedGroup[2].Bytes(), dispatchedGroup[3].Bytes()...)
		temp = append(dispatchedGroup[1].Bytes(), temp...)
		temp = append(dispatchedGroup[0].Bytes(), temp...)
		temp = append([]byte{0x01}, temp...)

		//fmt.Println("isMember from onchain : ", temp)
		groupPub, err := d.p2pDkg.GetGroupPublicPoly().Commit().MarshalBinary()
		//fmt.Println("isMember : ", groupPub)
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
	fmt.Println("isMember false")
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
	//id = append(new([12]byte)[:], d.groupIds[submitter]...)
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
					groupIds := [][]byte{}
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
func (d *DosNode) PipeCheckURL(chLogUrl <-chan interface{}) <-chan Report {
	out := make(chan Report)
	go func() {
		for {
			select {
			//event from blockChain
			case msg := <-chLogUrl:
				switch content := msg.(type) {
				case *onchain.DOSProxyLogUrl:
					//Check to see if this request is for node's group
					if d.isMember(content.DispatchedGroup) {
						queryId := content.QueryId
						fmt.Println("PipeCheckURL!!", queryId)
						url := content.Url
						submitter := d.choseSubmitter(content.Randomness)
						msg, err := d.dataFetch(url)
						if err != nil {
							fmt.Println(err)
						}
						msg = append(msg, make([]byte, 12)...)
						msg = append(msg, submitter...)
						//combine result with submitter
						sign := &vss.Signature{
							Index:   uint32(ForCheckURL),
							QueryId: queryId.String(),
							Content: msg,
						}
						report := &Report{}
						report.selfSign = *sign
						report.submitter = submitter
						report.backupTask = *content
						out <- *report
					}
				default:
					fmt.Println("DOSProxyLogUrl type mismatch", msg)
				}
			case <-d.quit:
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

// TODO: error handling and logging
// Note: Signed messages keep synced with on-chain contracts
func (d *DosNode) PipeGenerateRandomNumber(chRandom <-chan interface{}) <-chan Report {
	out := make(chan Report)
	go func() {
		for {
			select {
			//event from blockChain
			case msg := <-chRandom:
				switch content := msg.(type) {
				case *onchain.DOSProxyLogUpdateRandom:
					if d.isMember(content.DispatchedGroup) {
						preRandom := content.LastRandomness
						fmt.Printf("1 ) GenerateRandomNumber preRandom #%v \n", preRandom)
						fmt.Println("1 ) GenerateRandomNumber preRandom ", preRandom)
						//To avoid duplicate query
						//_, ok := (*d.reportMap).Load(preRandom.String())
						//if !ok {
						submitter := d.choseSubmitter(preRandom)
						//hash, err := d.chainConn.GetBlockHashByNumber(preBlock)
						//if err != nil {
						//	fmt.Printf("GetBlockHashByNumber #%v fails\n", preBlock)
						//	return
						//}
						// message = concat(lastUpdatedBlockhash, lastRandomness, submitter)
						preRandomBytes := padOrTrim(preRandom.Bytes(), RANDOMNUMBERSIZE)
						randomNum := append(preRandomBytes, make([]byte, 12)...)
						randomNum = append(randomNum, submitter...)
						fmt.Println("GenerateRandomNumber content = ", randomNum)

						sign := &vss.Signature{
							Index:   uint32(ForRandomNumber),
							QueryId: preRandom.String(),
							Content: randomNum,
						}
						report := &Report{}
						report.selfSign = *sign
						report.submitter = submitter
						out <- *report
						//}
					}
				default:
					fmt.Println("DOSProxyLogUpdateRandom type mismatch")
				}
			case <-d.quit:
				close(out)
				break
			}
		}
	}()
	return out
}

func (d *DosNode) PipeSignAndBroadcast(reports ...<-chan Report) <-chan Report {
	merged := make(chan Report)
	var wg sync.WaitGroup
	wg.Add(len(reports))
	for _, c := range reports {
		go func(c <-chan Report) {
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
			case report := <-merged:
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
						//fmt.Println(i, " : send to ", member)
						(*d.network).SendMessageById(member, &sign)
						fmt.Println("send to ", member)
					}
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
				//fmt.Println("PipeRecoverAndVerify fromPeer")
				var sigShares [][]byte
				signIdentityBytes := append([]byte(sign.QueryId), sign.Signature...)
				signIdentity := new(big.Int).SetBytes(signIdentityBytes).String()
				if report, ok := reportMap[sign.QueryId]; ok {
					//fmt.Println("4) PipeRecoverAndVerify start")
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
					//fmt.Println("receiveSignature id ", sign.QueryId, " len ", len(sigShares), " nbParticipants ", d.nbParticipants)
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
						//if string(d.chainConn.GetId()) == string(report.submitter) {
						//	outForSubmit <- report
						//}
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
						//fmt.Println("4) PipeRecoverAndVerify not enough ->")
						cSignatureFromPeer <- sign
						//fmt.Println("4) PipeRecoverAndVerify not enough <-")
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
				case ForRandomNumber:
					//fmt.Println("PipeSendToOnchain = ", report.signGroup)
					//Test Case 1
					//os.Exit(0)
					//Test Case 2
					//report.signGroup[10] ^= 0x01
					err := d.chainConn.SetRandomNum(report.signGroup)
					if err != nil {
						fmt.Println("SetRandomNum err ", err)
					} else {
						fmt.Println("randomNumber Set for ", report.selfSign.QueryId)
					}
				default:
					fmt.Println("PipeSendToOnchain checkURL result = ", report.selfSign.Content)
					//fmt.Println("PipeSendToOnchain checkURL result = ", report.signGroup)
					qID := big.NewInt(0)
					qID.SetString(report.selfSign.QueryId, 10)

					fmt.Println("content", report.selfSign.Content)
					t := len(report.selfSign.Content) - 32
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
					err := d.chainConn.DataReturn(qID, queryResult, report.signGroup)
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
	chUrl := make(chan interface{})
	go func() {
		for {
			select {
			case report := <-forValidate:
				//fmt.Println("PipeCleanFinishMap ", report.timeStamp)
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
							if uint32(content.TrafficType) == ForRandomNumber {
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
							//fmt.Println("4) PipeRecoverAndVerify not enough ->")
							chValidation <- msg
							//fmt.Println("4) PipeRecoverAndVerify not enough <-")
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
				close(chUrl)
				ticker.Stop()
				return
			}
		}
	}()
	return chUrl
}
