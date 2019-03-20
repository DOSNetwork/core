package p2p

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"time"

	"github.com/dedis/kyber"

	"github.com/golang/protobuf/proto"

	"reflect"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p/internal"
	"github.com/DOSNetwork/core/suites"

	"github.com/hashicorp/serf/serf"
)

const DefaultModelName = "Undefined"

type P2P struct {
	identity internal.ID
	//Map of ID (string) <-> *p2p.PeerConn
	replyPeers   *sync.Map
	requestPeers *sync.Map
	messages     chan P2PMessage
	suite        suites.Suite
	port         string
	secKey       kyber.Scalar
	pubKey       kyber.Point
	//	routingTable *dht.RoutingTable
	logger     log.Logger
	msgChanMap sync.Map

	cluster *serf.Serf
}

func SetupCluster(advertiseAddr string, id []byte) (*serf.Serf, error) {
	conf := serf.DefaultConfig()
	conf.Init()
	conf.LogOutput = ioutil.Discard
	conf.MemberlistConfig.AdvertiseAddr = advertiseAddr
	conf.NodeName = string(id)
	cluster, err := serf.Create(conf)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (n *P2P) SetID(id []byte) {
	n.identity.Id = id
}

func (n *P2P) GetID() []byte {
	return n.identity.Id
}

func (n *P2P) GetIP() string {
	return n.identity.Address
}

func (n *P2P) Listen() (err error) {
	var ip string
	var listener net.Listener
	var pubKeyBytes []byte

	//FOR DOCKER AWS TESTING
	//response, err := http.Get("http://ipconfig.me")
	//if err != nil {
	//	n.logger.Error(err)
	//	return
	//}

	//ipBytes, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	n.logger.Error(err)
	//	return
	//}
	//ip = string(ipBytes)
	//fmt.Println(ip)
	//////////////////////////////

	if ip, err = GetLocalIP(); err != nil {
		n.logger.Error(err)
		return
	}

	p := fmt.Sprintf(":%s", n.port)
	if listener, err = net.Listen("tcp", p); err != nil {
		n.logger.Error(err)
		return
	}

	//NAT discover

	//isPrivateIp, err := nat.IsPrivateIp()
	//if err != nil {
	//	return err
	//}
	//
	//if isPrivateIp {
	//	externalPort := nat.RandomPort()
	//	nat, err := nat.SetMapping("tcp", externalPort, listener.Addr().(*net.TCPAddr).Port, "DosNode")
	//	if err != nil {
	//		return err
	//	}
	//
	//	externalIp, err := nat.GetExternalAddress()
	//	if err != nil {
	//		return err
	//	}
	//
	//	n.port = externalPort
	//	ip = externalIp.String() + ":" + strconv.Itoa(n.port)
	//} else {
	//	n.port = listener.Addr().(*net.TCPAddr).Port
	//	ip = ip + ":" + strconv.Itoa(n.port)
	//}

	//n.port = listener.Addr().(*net.TCPAddr).Port
	n.identity.Address = ip

	n.secKey, n.pubKey = genPair()
	if pubKeyBytes, err = n.pubKey.MarshalBinary(); err != nil {
		n.logger.Error(err)
		return err
	}

	n.identity.PublicKey = pubKeyBytes

	//n.routingTable = dht.CreateRoutingTable(n.identity)
	go n.messageHanding()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				n.logger.Error(err)
				continue
			}

			go func() {
				peer, err := NewPeerConn(n, &conn, n.messages, true)
				if err != nil {
					peer.EndWithoutDelete()
					n.logger.Error(err)
					return
				}
				fmt.Println(n.identity.Id, " accept from ", log.ByteTohex(peer.identity.Id))
				//n.replyPeers.LoadOrStore(string(peer.identity.Id), peer)

			}()
		}
	}()

	return nil
}

func (n *P2P) Join(bootstrapIp string) (err error) {
	_, err = n.cluster.Join([]string{bootstrapIp}, true)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (n *P2P) Leave() {
	/*
		n.peers.Range(func(key, value interface{}) bool {
			peer := value.(*PeerConn)
			peer.End()
			return true
		})
	*/
	return
}
func (n *P2P) Members() int {
	return len(n.cluster.Members())
}
func (n *P2P) findPeer(id []byte, role int) (peer *PeerConn, found bool) {
	var value interface{}
	var err error
	//	var err error

	defer n.logger.TimeTrack(time.Now(), "FindPeer", nil)
	if role == 1 {

		// Find Peer from existing peerConn
		for retry := 0; retry < 10; retry++ {
			if value, found = n.replyPeers.Load(string(id)); found {
				peer = value.(*PeerConn)
				f := map[string]interface{}{
					"fromID": log.ByteTohex(id),
					"Time":   time.Now()}
				n.logger.Event("ReplyFoundFromPeers", f)
				return peer, true
			}
			time.Sleep(1 * time.Second)
		}
		f := map[string]interface{}{
			"fromID": log.ByteTohex(id),
			"Time":   time.Now()}
		n.logger.Event("ReplyNotFoundPeer", f)
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!Not found peer form replyPeers")
		return nil, false
	} else {
		// Find Peer from existing peerConn
		if value, found = n.requestPeers.Load(string(id)); found {
			peer = value.(*PeerConn)
			f := map[string]interface{}{
				"fromID": log.ByteTohex(id),
				"Time":   time.Now()}
			n.logger.Event("SendFoundFromPeers", f)
			return
		}
		// Find Peer from routing map
		members := n.cluster.Members()
		var addr net.IP
		for i := 0; i < len(members); i++ {
			if members[i].Name == string(id) {
				addr = members[i].Addr
				peer, err = n.connectTo(addr.String())
				if err != nil {
					fmt.Println("!!!!!!!!!!!!!!!!!!Can't connectTo ", err)
					f := map[string]interface{}{
						"fromID": log.ByteTohex(id),
						"Time":   time.Now()}
					n.logger.Event("NotAbleToConnect", f)
					return nil, false
				}
				f := map[string]interface{}{
					"fromID": log.ByteTohex(id),
					"Time":   time.Now()}
				n.logger.Event("SendFoundFromSWIM", f)
				found = true
				return
			}
		}
	}
	return nil, false
}
func (n *P2P) SendMessage(id []byte, m proto.Message) (err error) {
	var peer *PeerConn
	var found bool
	defer n.logger.TimeTrack(time.Now(), "Request", nil)

	if peer, found = n.findPeer(id, 0); found {
		request := new(Request)
		request.SetMessage(m)
		request.SetTimeout(180 * time.Second)

		if _, err = peer.Request(request); err != nil {
			fmt.Println("Request err ", err)
			n.logger.Error(err)
			return
		}
	}

	return
}
func (n *P2P) Request(id []byte, m proto.Message) (msg proto.Message, err error) {
	var peer *PeerConn
	var found bool
	//defer n.logger.Metrics(time.Since(startTime).Seconds() * 1000)
	///defer n.logger.TimeTrack(time.Now(), "Request", nil)
	if peer, found = n.findPeer(id, 0); found {
		request := new(Request)
		request.SetMessage(m)
		request.SetTimeout(180 * time.Second)

		if msg, err = peer.Request(request); err != nil {
			fmt.Println("Request ", err)
			n.logger.Error(err)
		}

		return
	} else {
		return nil, errors.New("p2pRequest: peer not found")
	}
}

func (n *P2P) Reply(id []byte, nonce uint64, response proto.Message) (err error) {
	fmt.Println("Reply to ", id)
	if peer, found := n.findPeer(id, 1); found {
		err = peer.Reply(nonce, response)
		return
	} else {
		return errors.New("p2pReply: peer not found")
	}
}

/*
This is a block call
*/
func (n *P2P) ConnectTo(addr string) (id []byte, err error) {
	var peer *PeerConn

	if peer, err = n.connectTo(addr); err != nil {
		n.logger.Error(err)
		return
	}

	id = peer.identity.Id
	return
}

func (n *P2P) connectTo(addr string) (peer *PeerConn, err error) {
	addr = addr + ":" + n.port
	var conn net.Conn
	fmt.Println(n.identity.Address, " connectTo ", addr)
	for retry := 0; retry < 5; retry++ {
		if conn, err = net.Dial("tcp", addr); err != nil {
			n.logger.Error(err)
			time.Sleep(1 * time.Second)
			continue
		}

		if peer, err = NewPeerConn(n, &conn, n.messages, false); err != nil {

			//peer.EndWithoutDelete()
			n.logger.Error(err)
			continue
		}
		if err = peer.SayHi(); err != nil {

			//peer.EndWithoutDelete()
			n.logger.Error(err)
			continue
		}

		n.requestPeers.LoadOrStore(string(peer.identity.Id), peer)
		fmt.Println("connectTo retry ", retry)
		return
	}

	err = errors.New("Retry connection over the limit")
	n.logger.Error(err)
	return
}

func (n *P2P) messageHanding() {
	for message := range n.messages {
		messagetype := reflect.TypeOf(message.Msg.Message).String()
		if len(messagetype) > 0 && messagetype[0] == '*' {
			messagetype = messagetype[1:]
		}
		if ch, ok := n.msgChanMap.Load(messagetype); ok {
			go func() {
				select {
				case ch.(chan P2PMessage) <- message:
				case <-time.After(5 * time.Second):
					fmt.Println("messageHanding timeout !!!!!!!!!!!!!!!")
				}
			}()
		}

	}
}

func (n *P2P) SubscribeEvent(chanBuffer int, messages ...interface{}) (outch chan P2PMessage, err error) {
	if chanBuffer > 0 {
		outch = make(chan P2PMessage, chanBuffer)
	} else {
		outch = make(chan P2PMessage)
	}

	errstr := ""
	for _, m := range messages {
		_, l := n.msgChanMap.LoadOrStore(reflect.TypeOf(m).String(), outch)
		if l {
			if errstr != "" {
				errstr = errstr + ", " + reflect.TypeOf(m).String()
			} else {
				errstr = reflect.TypeOf(m).String()
			}
		}
	}
	if errstr != "" {
		err = errors.New("The messages:[" + errstr + "]has been subscribed")
		n.logger.Error(err)
	}
	return
}

func (n *P2P) UnSubscribeEvent(messages ...interface{}) {
	for _, m := range messages {
		ch, ok := n.msgChanMap.Load(reflect.TypeOf(m).String())
		if ok {
			n.msgChanMap.Delete(reflect.TypeOf(m).String())
			find := false
			n.msgChanMap.Range(func(key, value interface{}) bool {
				if value.(chan P2PMessage) == ch.(chan P2PMessage) {
					find = true
					return false
				}
				return true
			})
			if !find {
				close(ch.(chan P2PMessage))
			}
		}
	}
	return
}

func (n *P2P) CloseMessagesChannel() {
	close(n.messages)
}
