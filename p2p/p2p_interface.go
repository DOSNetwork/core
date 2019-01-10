package p2p

import (
	"encoding/hex"
	"errors"
	"net"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/crypto/sha3"
)

var (
	suite = suites.MustFind("bn256")
)

func genPair() (kyber.Scalar, kyber.Point) {
	secret := suite.Scalar().Pick(suite.RandomStream())
	public := suite.Point().Mul(secret, nil)
	return secret, public
}

func Hex(a []byte) string {
	unchecksummed := hex.EncodeToString(a[:])
	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(unchecksummed))
	hash := sha.Sum(nil)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return "0x" + string(result)
}

func CreateP2PNetwork(id []byte, port int) (P2PInterface, chan P2PMessage, error) {
	p := &P2P{
		peers:    CreatePeerConnManager(),
		suite:    suite,
		messages: make(chan P2PMessage, 100),
		port:     port,
		logger:   log.New("module", "p2p"),
	}
	p.identity.Id = id
	return p, p.messages, nil
}

func GetLocalIP() (ip string, err error) {
	var addrs []net.Addr

	if addrs, err = net.InterfaceAddrs(); err != nil {
		return "", err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("IP not found")
}

type P2PMessage struct {
	Msg    ptypes.DynamicAny
	Sender []byte
}

type P2PInterface interface {
	GetIP() string
	GetID() []byte
	Listen() error
	Join(bootstrapIp string) error
	ConnectTo(IpAddr string) (id []byte, err error)
	Leave()
	SendMessage(id []byte, msg proto.Message) error
}
