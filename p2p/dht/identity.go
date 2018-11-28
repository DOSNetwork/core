package dht

import (
	"bytes"
	"encoding/hex"
	"math/bits"

	"golang.org/x/crypto/blake2b"

	"github.com/DOSNetwork/core/p2p/internal"
)

// internal.ID is an identity of nodes, using its public key hash and network address.
//type internal.ID internal.internal.ID

// Createinternal.ID is a factory function creating internal.ID.
func CreateID(address string, publicKey []byte) internal.ID {
	id := blake2b.Sum256(publicKey)
	return internal.ID{Address: address, PublicKey: publicKey, Id: id[:]}
}

// Equals determines if two peer internal.IDs are equal to each other based on the contents of their public keys.
func Equals(a, b internal.ID) bool {
	return bytes.Equal(a.Id, b.Id)
}

// Less determines if this peer internal.ID's public key is less than other internal.ID's public key.
func Less(a, b internal.ID) bool {
	return bytes.Compare(a.Id, b.Id) == -1
}

// PublicKeyHex generates a hex-encoded string of public key hash of this given peer internal.ID.
func PublicKeyHex(a internal.ID) string {
	return hex.EncodeToString(a.PublicKey)
}

// Xor performs XOR (^) over another peer internal.ID's public key.
func Xor(a, b internal.ID) internal.ID {
	result := make([]byte, len(a.PublicKey))

	for i := 0; i < len(a.PublicKey) && i < len(b.PublicKey); i++ {
		result[i] = a.PublicKey[i] ^ b.PublicKey[i]
	}
	return internal.ID{Address: a.Address, PublicKey: result}
}

// Xorinternal.ID performs XOR (^) over another peer internal.ID's public key hash.
func XorID(a, b internal.ID) internal.ID {
	result := make([]byte, len(a.Id))

	for i := 0; i < len(a.Id) && i < len(b.Id); i++ {
		result[i] = a.Id[i] ^ b.Id[i]
	}
	return internal.ID{Address: a.Address, Id: result}
}

// PrefixLen returns the number of prefixed zeros in a peer internal.ID.
func PrefixLen(a internal.ID) int {
	for i, b := range a.Id {
		if b != 0 {
			return i*8 + bits.LeadingZeros8(uint8(b))
		}
	}
	return len(a.Id)*8 - 1
}
