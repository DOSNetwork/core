// Package bls implements the Boneh-Lynn-Shacham (BLS) signature scheme which
// was introduced in the paper "Short Signatures from the Weil Pairing". BLS
// requires pairing-based cryptography.
package bls

import (
	"crypto/cipher"
	"errors"

	"github.com/DOSNetwork/core/suites"
	"github.com/dedis/kyber"
	"golang.org/x/crypto/sha3"
)

// NewKeyPair creates a new BLS signing key pair. The private key x is a scalar
// and the public key X is a point on curve G2.
func NewKeyPair(suite suites.Suite, random cipher.Stream) (kyber.Scalar, kyber.Point) {
	x := suite.G2().Scalar().Pick(random)
	X := suite.G2().Point().Mul(x, nil)
	return x, X
}

// Sign creates a BLS signature S = x * H(m) on a message m using the private
// key x. The signature S is a point on curve G1.
func Sign(suite suites.Suite, x kyber.Scalar, msg []byte) ([]byte, error) {
	HM := hashToPoint(suite, msg)
	xHM := HM.Mul(x, HM)
	s, err := xHM.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Verify checks the given BLS signature S on the message m using the public
// key X by verifying that the equality e(H(m), X) == e(H(m), x*B2) ==
// e(x*H(m), B2) == e(S, B2) holds where e is the pairing operation and B2 is
// the base point from curve G2.
func Verify(suite suites.Suite, X kyber.Point, msg, sig []byte) error {
	HM := hashToPoint(suite, msg)
	s := suite.G1().Point()
	if err := s.UnmarshalBinary(sig); err != nil {
		return err
	}
	s.Neg(s)
	if !suite.PairingCheck([]kyber.Point{s, HM}, []kyber.Point{suite.G2().Point().Base(), X}) {
		return errors.New("bls: invalid signature")
	}
	return nil
}

// hashToPoint hashes a message to a point on curve G1. XXX: This should be replaced
// eventually by a proper hash-to-point mapping like Elligator.
func hashToPoint(suite suites.Suite, msg []byte) kyber.Point {
	hash := sha3.NewLegacyKeccak256()
	var buf []byte
	hash.Write(msg)
	buf = hash.Sum(buf)
	x := suite.G1().Scalar().SetBytes(buf)
	point := suite.G1().Point().Mul(x, nil)
	return point
}
