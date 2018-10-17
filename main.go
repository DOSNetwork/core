package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"

	"github.com/DOSNetwork/core/blockchain"
	"github.com/DOSNetwork/core/blockchain/eth"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/p2p/dht"
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/share/vss/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
)

var mux sync.Mutex

const (
	ForRandomNumber = uint32(1)
	ForCheckURL     = uint32(0)
)

type dosNode struct {
	suite          suites.Suite
	quit           chan struct{}
	nbParticipants int
	nbThreshold    int
	groupPubPoly   share.PubPoly
	shareSec       share.PriShare
	chainConn      blockchain.ChainInterface
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
	backupTask eth.DOSProxyLogUrl
}

func (d *dosNode) dataFetch(url string) ([]byte, error) {
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

func (d *dosNode) isMember(dispatchedGroup [4]*big.Int) bool {
	if d.p2pDkg.IsCetified() {
		temp := append(dispatchedGroup[2].Bytes(), dispatchedGroup[3].Bytes()...)
		temp = append(dispatchedGroup[1].Bytes(), temp...)
		temp = append(dispatchedGroup[0].Bytes(), temp...)
		temp = append([]byte{0x01}, temp...)
		//fmt.Println("isMember from eth : ", temp)
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

func (d *dosNode) choseSubmitter(r *big.Int) (id []byte) {
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

func (d *dosNode) PipeCheckURL(chLogUrl <-chan interface{}) <-chan Report {
	out := make(chan Report)
	go func() {
		for {
			select {
			//event from blockChain
			case msg := <-chLogUrl:
				switch content := msg.(type) {
				case *eth.DOSProxyLogUrl:
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
						msg = append(msg[:], new([12]byte)[:]...)
						msg = append(msg[:], submitter...)
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
					fmt.Println("type mismatch")
				}
			case <-d.quit:
				close(out)
				break
			}
		}
	}()
	return out
}

// TODO: error handling and logging
// Note: Signed messages keep synced with on-chain contracts
func (d *dosNode) PipeGenerateRandomNumber(chRandom <-chan interface{}) <-chan Report {
	out := make(chan Report)
	go func() {
		for {
			select {
			//event from blockChain
			case msg := <-chRandom:
				switch content := msg.(type) {
				case *eth.DOSProxyLogUpdateRandom:

					if d.isMember(content.DispatchedGroup) {
						preRandom := content.LastRandomness
						preBlock := content.LastUpdatedBlock
						fmt.Printf("0 ) DOSProxyLogUpdateRandom #%v \n", preBlock)
						fmt.Printf("1 ) GenerateRandomNumber #%v \n", preRandom)
						//To avoid duplicate query
						//_, ok := (*d.reportMap).Load(preRandom.String())
						//if !ok {
						submitter := d.choseSubmitter(preRandom)
						hash, err := d.chainConn.GetBlockHashByNumber(preBlock)
						if err != nil {
							fmt.Printf("GetBlockHashByNumber #%v fails\n", preBlock)
							return
						}
						// message = concat(lastUpdatedBlockhash, lastRandomness, submitter)
						msg := append(hash[:], preRandom.Bytes()...)
						msg = append(msg[:], new([12]byte)[:]...)
						msg = append(msg[:], submitter...)
						fmt.Println("GenerateRandomNumber content = ", msg)

						sign := &vss.Signature{
							Index:   uint32(ForRandomNumber),
							QueryId: preRandom.String(),
							Content: msg,
						}
						report := &Report{}
						report.selfSign = *sign
						report.submitter = submitter
						out <- *report
						//}
					}
				default:
					fmt.Println("type mismatch")
				}
			case <-d.quit:
				close(out)
				break
			}
		}
	}()
	return out
}

func (d *dosNode) PipeSignAndBroadcast(reports ...<-chan Report) <-chan Report {
	merged := make(chan Report)
	var wg sync.WaitGroup
	wg.Add(len(reports))
	out := make(chan Report)
	for _, c := range reports {
		go func(c <-chan Report) {
			for v := range c {
				merged <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		for {
			select {
			case report := <-merged:
				fmt.Println("2) PipeSignAndBroadcast")
				sign := report.selfSign
				sig, _ := tbls.Sign(d.suite, d.p2pDkg.GetShareSecuirty(), sign.Content)
				sign.Signature = sig
				report.signShares = append(report.signShares, sig)
				out <- report
				for i, member := range d.groupIds {
					if string(member) != string((*d.network).GetId().Id) {
						//Todo:Need to check to see if it is thread safe
						fmt.Println(i, " : send to ", member)
						(*d.network).SendMessageById(member, &sign)
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
func (d *dosNode) PipeRecoverAndVerify(cSignatureFromPeer chan vss.Signature, fromEth <-chan Report) (<-chan Report, <-chan Report) {
	outForSubmit := make(chan Report)
	outForValidate := make(chan Report)
	ticker := time.NewTicker(10 * time.Minute)
	reportMap := new(sync.Map)
	go func() {
		for {
			select {
			case report := <-fromEth:
				fmt.Println("3) PipeRecoverAndVerify fromEth")
				reportMap.Store(report.selfSign.QueryId, report)
			case sign := <-cSignatureFromPeer:
				var sigShares [][]byte
				result, ok := reportMap.Load(sign.QueryId)
				if ok {
					fmt.Println("4) PipeRecoverAndVerify start")
					report, ok := result.(Report)
					if !ok {
						fmt.Println("cast report failed ", sign.QueryId)
						continue
					}
					//TODO:Should check if content is always the same
					report.signShares = append(report.signShares, sign.Signature)
					reportMap.Store(sign.QueryId, report)
					sigShares = report.signShares
					fmt.Println("receiveSignature id ", sign.QueryId, " len ", len(sigShares), " nbParticipants ", d.nbParticipants)
					if report.signGroup == nil && len(sigShares) >= d.nbThreshold {
						fmt.Println("5) PipeRecoverAndVerify recover")
						sig, err := tbls.Recover(d.suite, d.p2pDkg.GetGroupPublicPoly(), sign.Content, sigShares, d.nbThreshold, d.nbParticipants)
						if err != nil {
							continue
						}
						err = bls.Verify(d.suite, d.p2pDkg.GetGroupPublicPoly().Commit(), sign.Content, sig)
						if err != nil {
							continue
						}
						fmt.Println("6) Verify success")
						report.signGroup = sig
						report.timeStamp = time.Now()
						if string(d.chainConn.GetId()) == string(report.submitter) {
							outForSubmit <- report
						}
						outForValidate <- report
						reportMap.Store(sign.QueryId, report)
					}
				} else {
					//Other nodes has received eth event and send signature to other nodes
					//Node put back these signatures until it received eth event.
					//TODO:It should use a timeStamp that can be used to see if it is a stale signature
					cSignatureFromPeer <- sign
				}
			case <-ticker.C:
				reportMap.Range(func(key, value interface{}) bool {
					fmt.Println("Tiem to clean .........")
					report := value.(Report)
					dur := time.Since(report.timeStamp)
					if dur.Minutes() >= 10 {
						reportMap.Delete(key)
						fmt.Println("clean report key ", key, "time dur ", dur)
					}
					return true
				})
			case <-d.quit:
				close(outForSubmit)
				close(outForValidate)
				ticker.Stop()
				break
			}
		}
	}()
	return outForSubmit, outForValidate
}

func (d *dosNode) PipeSendToOnchain(chReport <-chan Report) {
	go func() {
		for {
			select {
			case report := <-chReport:
				switch report.selfSign.Index {
				case ForRandomNumber:
					fmt.Println("PipeSendToOnchain = ", report.signGroup)
					//Test Case 1
					os.Exit(0)
					//Test Case 2
					//report.signGroup[10] ^= 0x01
					err := d.chainConn.SetRandomNum(report.signGroup)
					if err != nil {
						fmt.Println("SetRandomNum err ", err)
					}
				default:
					fmt.Println("PipeSendToOnchain checkURL result = ", string(report.selfSign.Content))
					fmt.Println("PipeSendToOnchain checkURL result = ", report.signGroup)
					qID := big.NewInt(0)
					qID.SetString(report.selfSign.QueryId, 10)
					err := d.chainConn.DataReturn(qID, report.selfSign.Content, report.signGroup)
					if err != nil {
						fmt.Println("DataReturn err ", err)
					}
				}
			case <-d.quit:
				break
			}
		}
	}()
	return
}

func (d *dosNode) PipeCleanFinishMap(chUrl chan interface{}, chValidation <-chan interface{}, forValidate <-chan Report) {
	ticker := time.NewTicker(2 * time.Minute)
	validateMap := new(sync.Map)
	lastValidatedRandom := time.Now()
	go func() {
		for {
			select {
			case report := <-forValidate:
				fmt.Println("PipeCleanFinishMap ", report.timeStamp)
				validateMap.Store(report.selfSign.QueryId, report)
			case msg := <-chValidation:
				switch content := msg.(type) {
				case *eth.DOSProxyLogValidationResult:
					if !content.Pass {
						switch uint32(content.TrafficType) {
						case ForRandomNumber:
							fmt.Println("Invalide Signature.........")
						default:
							fmt.Println("Retrigger query.........")
							result, ok := validateMap.Load(content.TrafficId.String())
							if ok {
								report, _ := result.(Report)
								tmp := big.NewInt(1)
								report.backupTask.Randomness = report.backupTask.Randomness.Sub(report.backupTask.Randomness, tmp)
								report.backupTask.Timeout = report.backupTask.Timeout.Sub(report.backupTask.Timeout, tmp)
								chUrl <- report.backupTask
							}
						}
					} else {
						_, ok := validateMap.Load(content.TrafficId.String())
						if ok {
							fmt.Println("Validated ", content.TrafficId.String())
							validateMap.Delete(content.TrafficId.String())
							lastValidatedRandom = time.Now()
						}
					}
				default:
					fmt.Println("type mismatch")
				}
			case <-ticker.C:
				if d.p2pDkg.IsCetified() {
					dur := time.Since(lastValidatedRandom)
					if dur.Minutes() >= 5 {
						fmt.Println("WatchDog Random tiemout.........")
						d.chainConn.RandomNumberTimeOut()
					}
				}
				validateMap.Range(func(key, value interface{}) bool {
					fmt.Println("Tiem to clean .........")
					report := value.(Report)
					dur := time.Since(report.timeStamp)
					if dur.Minutes() >= 1 {
						switch report.selfSign.Index {
						case ForRandomNumber:
							fmt.Println("Random tiemout.........")
						default:
							fmt.Println("Retrigger query.........")
							result, ok := validateMap.Load(report.selfSign.QueryId)
							if ok {
								report, _ := result.(Report)
								tmp := big.NewInt(0)
								if report.backupTask.Timeout.Cmp(tmp) == 0 {
									validateMap.Delete(key)
									fmt.Println("give up this query.........")
								}
								tmp = big.NewInt(1)
								report.backupTask.Randomness = report.backupTask.Randomness.Sub(report.backupTask.Randomness, tmp)
								report.backupTask.Timeout = report.backupTask.Timeout.Sub(report.backupTask.Timeout, tmp)
								chUrl <- report.backupTask
							}
						}
					}
					if dur.Minutes() >= 10 {
						//Failed too long ,just delete
						validateMap.Delete(key)
					}
					return true
				})
			case <-d.quit:
				ticker.Stop()
				return
			}
		}
	}()
	return
}

// main
func main() {
	roleFlag := flag.String("role", "", "BootstrapNode or not")
	nbParticipantsFlag := flag.Int("nbVerifiers", 3, "Number of Participants")
	portFlag := flag.Int("port", 0, "port number")
	bootstrapIpFlag := flag.String("bootstrapIp", "67.207.98.117:42745", "bootstrapIp")

	flag.Parse()
	role := *roleFlag
	nbParticipants := *nbParticipantsFlag
	port := *portFlag
	bootstrapIp := *bootstrapIpFlag
	//1)Connect to Eth and Set node ID
	chainConn, err := blockchain.AdaptTo(blockchain.ETH, true, eth.Rinkeby)
	if err != nil {
		log.Fatal(err)
	}

	err = chainConn.UploadID()
	if err != nil {
		log.Fatal(err)
	}

	//1)Build a p2p network
	peerEvent := make(chan p2p.P2PMessage, 100)
	defer close(peerEvent)
	p, _ := p2p.CreateP2PNetwork(peerEvent, port)
	p.SetId(chainConn.GetId())
	p.Listen()

	//3)Dial to peers to build peerClient
	if role == "" {
		fmt.Println(bootstrapIp)
		p.CreatePeer(bootstrapIp, nil)
		results := p.FindNode(p.GetId(), dht.BucketSize, 20)
		for _, result := range results {
			p.GetRoutingTable().Update(result)
			fmt.Println(p.GetId().Address, "Update peer: ", result.Address)
		}
	}

	suite := suites.MustFind("bn256")
	peerEventForDKG := make(chan p2p.P2PMessage, 100)
	defer close(peerEventForDKG)
	p2pDkg, _ := dkg.CreateP2PDkg(p, suite, peerEventForDKG, nbParticipants)
	go p2pDkg.EventLoop()
	dkgEvent := make(chan string, 100)
	p2pDkg.SubscribeEvent(dkgEvent)
	defer close(dkgEvent)

	chGroup := make(chan interface{}, 100)
	defer close(chGroup)
	chUrl := make(chan interface{}, 100)
	defer close(chUrl)
	chRandom := make(chan interface{}, 100)
	defer close(chRandom)
	cSignatureFromPeer := make(chan vss.Signature, 100)
	defer close(cSignatureFromPeer)
	chValidation := make(chan interface{}, 100)
	defer close(chValidation)
	chainConn.SubscribeEvent(chUrl, eth.SubscribeDOSProxyLogUrl)
	err = chainConn.SubscribeEvent(chGroup, eth.SubscribeDOSProxyLogGrouping)
	chainConn.SubscribeEvent(chRandom, eth.SubscribeDOSProxyLogUpdateRandom)
	chainConn.SubscribeEvent(chValidation, eth.SubscribeDOSProxyLogValidationResult)
	toOnChainQueue := make(chan string, 100)
	defer close(toOnChainQueue)

	d := &dosNode{
		suite:          suite,
		quit:           make(chan struct{}),
		nbParticipants: nbParticipants,
		nbThreshold:    nbParticipants/2 + 1,
		network:        &p,
		chainConn:      chainConn,
		p2pDkg:         p2pDkg,
	}
	out1 := d.PipeCheckURL(chUrl)
	out2 := d.PipeGenerateRandomNumber(chRandom)
	out3 := d.PipeSignAndBroadcast(out1, out2)
	outForSubmit, outForValidate := d.PipeRecoverAndVerify(cSignatureFromPeer, out3)
	d.PipeSendToOnchain(outForSubmit)
	d.PipeCleanFinishMap(chUrl, chValidation, outForValidate)
	_ = outForValidate
	for {
		select {
		//event from peer
		case msg := <-peerEvent:
			switch content := msg.Msg.Message.(type) {
			case *vss.PublicKey:
				peerEventForDKG <- msg
			case *dkg.Deal:
				peerEventForDKG <- msg
			case *dkg.Response:
				peerEventForDKG <- msg
			case *vss.Signature:
				cSignatureFromPeer <- *content
				fmt.Println("signature form peer")
			default:
				fmt.Println("unknown", content)
			}
		case msg := <-dkgEvent:
			if msg == "cetified" {
				gId := new(big.Int)
				gId.SetBytes(d.p2pDkg.GetGroupId())
				d.chainConn.UploadPubKey(d.p2pDkg.GetGroupPublicPoly().Commit())
			}
		case msg := <-chGroup:
			fmt.Println("chGroup")
			switch content := msg.(type) {
			case *eth.DOSProxyLogGrouping:
				isMember := false
				groupIds := [][]byte{}
				for i, node := range content.NodeId {
					id := node.Bytes()
					if string(id) == string(p.GetId().Id) {
						isMember = true
					}
					fmt.Println("DOSProxyLogGrouping member i= ", i, " id ", id, " ", isMember)
					groupIds = append(groupIds, id)
				}
				d.nbParticipants = len(groupIds)
				d.nbThreshold = d.nbParticipants/2 + 1
				d.p2pDkg.SetNbParticipants(d.nbParticipants)
				if isMember {
					d.groupIds = groupIds
					d.p2pDkg.SetGroupMembers(groupIds)
					d.p2pDkg.RunDKG()
				}
			default:
				fmt.Println("type mismatch")
			}
		}
	}
}
