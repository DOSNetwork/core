package endpoint

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/dedis/kyber"
	"log"
	"math/bits"
)

// ID is an identity of nodes, using its public key hash and network address.
type endpoint struct {
	address		string
	pubKey      kyber.Point
	id          []byte
}

// CreateID is a factory function creating ID.
func CreateEndpoint(address string, publicKey kyber.Point) *endpoint {
	publicKeyBytes, err := publicKey.MarshalBinary()
	if err != nil {
		log.Fatal("Unable to marshal public key")
	}

	return &endpoint{address: address, pubKey: publicKey, id:sha1.Sum(publicKeyBytes)[:]}
}

// String returns the identity address and public key.
func (endPoint *endpoint) String() string {
	return fmt.Sprintf("ID{Address: %v, Id: %v}", endPoint.address, endPoint.id)
}

// Equals determines if two peer IDs are equal to each other based on the contents of their public keys.
func (endpoint *endpoint) Equals(other *endpoint) bool {
	return bytes.Equal(endpoint.id, other.id)
}

// Less determines if this peer ID's public key is less than other ID's public key.
func (endpoint *endpoint) Less(other interface{}) bool {
	if other, is := other.(node); is {
		return bytes.Compare(id.Id, other.Id) == -1
	}
	return false
}

// PublicKeyHex generates a hex-encoded string of public key hash of this given peer ID.
func (id ID) PublicKeyHex() string {
	return hex.EncodeToString(id.PublicKey)
}

// Xor performs XOR (^) over another peer ID's public key.
func (id ID) Xor(other ID) ID {
	result := make([]byte, len(id.PublicKey))

	for i := 0; i < len(id.PublicKey) && i < len(other.PublicKey); i++ {
		result[i] = id.PublicKey[i] ^ other.PublicKey[i]
	}
	return ID{Address: id.Address, PublicKey: result}
}

// XorID performs XOR (^) over another peer ID's public key hash.
func (id ID) XorID(other ID) ID {
	result := make([]byte, len(id.Id))

	for i := 0; i < len(id.Id) && i < len(other.Id); i++ {
		result[i] = id.Id[i] ^ other.Id[i]
	}
	return ID{Address: id.Address, Id: result}
}

// PrefixLen returns the number of prefixed zeros in a peer ID.
func (id ID) PrefixLen() int {
	for i, b := range id.Id {
		if b != 0 {
			return i*8 + bits.LeadingZeros8(uint8(b))
		}
	}
	return len(id.Id)*8 - 1
}
