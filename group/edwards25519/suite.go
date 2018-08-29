package edwards25519

import (
	"crypto/cipher"
	"crypto/sha256"
	"hash"
	"io"
	"reflect"

	"github.com/DOSNetwork/core/group/internal/marshalling"
	"github.com/dedis/fixbuf"
	"github.com/dedis/kyber"
	"github.com/dedis/kyber/util/random"
	"github.com/dedis/kyber/xof/blake2xb"
)

// SuiteEd25519 implements some basic functionalities such as Group, HashFactory,
// and XOFFactory.
type SuiteEd25519 struct {
	Curve
	r cipher.Stream
}

// G1 returns the group G1 of the BN256 pairing.
func (s *SuiteEd25519) G1() kyber.Group {
	return s
}

// G2 returns the group G2 of the BN256 pairing.
func (s *SuiteEd25519) G2() kyber.Group {
	return s
}

// GT returns the group GT of the BN256 pairing.
func (s *SuiteEd25519) GT() kyber.Group {
	return s
}

// Pair takes the points p1 and p2 in groups G1 and G2, respectively, as input
// and computes their pairing in GT.
func (s *SuiteEd25519) Pair(p1 kyber.Point, p2 kyber.Point) kyber.Point {
	return s.Point()
}

// Pair takes the points p1 and p2 in groups G1 and G2, respectively, as input
// and computes their pairing in GT.
func (s *SuiteEd25519) PairingCheck(a []kyber.Point, b []kyber.Point) bool {
	return false
}

// Hash returns a newly instanciated sha256 hash function.
func (s *SuiteEd25519) Hash() hash.Hash {
	return sha256.New()
}

// XOF returns an XOF which is implemented via the Blake2b hash.
func (s *SuiteEd25519) XOF(key []byte) kyber.XOF {
	return blake2xb.New(key)
}

func (s *SuiteEd25519) Read(r io.Reader, objs ...interface{}) error {
	return fixbuf.Read(r, s, objs...)
}

func (s *SuiteEd25519) Write(w io.Writer, objs ...interface{}) error {
	return fixbuf.Write(w, objs)
}

// New implements the kyber.Encoding interface
func (s *SuiteEd25519) New(t reflect.Type) interface{} {
	return marshalling.GroupNew(s, t)
}

// RandomStream returns a cipher.Stream that returns a key stream
// from crypto/rand.
func (s *SuiteEd25519) RandomStream() cipher.Stream {
	if s.r != nil {
		return s.r
	}
	return random.New()
}

// NewBlakeSHA256Ed25519 returns a cipher suite based on package
// github.com/dedis/kyber/xof/blake2xb, SHA-256, and the Ed25519 curve.
// It produces cryptographically random numbers via package crypto/rand.
func NewBlakeSHA256Ed25519() *SuiteEd25519 {
	suite := new(SuiteEd25519)
	return suite
}

// NewBlakeSHA256Ed25519WithRand returns a cipher suite based on package
// github.com/dedis/kyber/xof/blake2xb, SHA-256, and the Ed25519 curve.
// It produces cryptographically random numbers via the provided stream r.
func NewBlakeSHA256Ed25519WithRand(r cipher.Stream) *SuiteEd25519 {
	suite := new(SuiteEd25519)
	suite.r = r
	return suite
}
