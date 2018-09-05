package bn256

import (
	"testing"

	"github.com/dedis/kyber"
	"github.com/dedis/kyber/group/mod"
	"github.com/dedis/kyber/util/random"
	"github.com/ethereum/go-ethereum/crypto/bn256"
	"github.com/stretchr/testify/require"
)

func TestScalarMarshal(t *testing.T) {
	suite := NewSuite()
	a := suite.G1().Scalar().Pick(random.New())
	b := suite.G1().Scalar()
	am, err := a.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	if err := b.UnmarshalBinary(am); err != nil {
		t.Fatal(err)
	}
	if !a.Equal(b) {
		t.Fatal("bn256: scalars not equal")
	}
}

func TestScalarOps(t *testing.T) {
	suite := NewSuite()
	a := suite.Scalar().Pick(random.New())
	b := suite.Scalar().Pick(random.New())
	c := suite.G1().Scalar().Pick(random.New())
	d := suite.G1().Scalar()
	e := suite.G1().Scalar()
	// check that (a+b)-c == (a-c)+b
	d.Add(a, b)
	d.Sub(d, c)
	e.Sub(a, c)
	e.Add(e, b)
	require.True(t, d.Equal(e))
	// check that (a*b)*c^-1 == (a*c^-1)*b
	d.One()
	e.One()
	d.Mul(a, b)
	d.Div(d, c)
	e.Div(a, c)
	e.Mul(e, b)
	require.True(t, d.Equal(e))
	// check that (a*b*c)^-1*(a*b*c) == 1
	d.One()
	e.One()
	d.Mul(a, b)
	d.Mul(d, c)
	d.Inv(d)
	e.Mul(a, b)
	e.Mul(e, c)
	e.Mul(e, d)
	require.True(t, e.Equal(suite.G1().Scalar().One()))
}

func TestG1(t *testing.T) {
	suite := NewSuite()
	k := suite.G1().Scalar().Pick(random.New())
	pa := suite.G1().Point().Mul(k, nil)
	ma, err := pa.MarshalBinary()
	require.Nil(t, err)

	pb := new(bn256.G1).ScalarBaseMult(&k.(*mod.Int).V)
	mb := pb.Marshal()

	require.Equal(t, ma, mb)
}

func TestG1Marshal(t *testing.T) {
	suite := NewSuite()
	k := suite.G1().Scalar().Pick(random.New())
	pa := suite.G1().Point().Mul(k, nil)
	ma, err := pa.MarshalBinary()
	require.Nil(t, err)

	pb := suite.G1().Point()
	err = pb.UnmarshalBinary(ma)
	require.Nil(t, err)

	mb, err := pb.MarshalBinary()
	require.Nil(t, err)

	require.Equal(t, ma, mb)
}

func TestG1Ops(t *testing.T) {
	suite := NewSuite()
	a := suite.G1().Point().Pick(random.New())
	b := suite.G1().Point().Pick(random.New())
	c := a.Clone()
	a.Neg(a)
	a.Neg(a)
	if !a.Equal(c) {
		t.Fatal("bn256.G1: neg failed")
	}
	a.Add(a, b)
	a.Sub(a, b)
	if !a.Equal(c) {
		t.Fatal("bn256.G1: add sub failed")
	}
	a.Add(a, suite.G1().Point().Null())
	if !a.Equal(c) {
		t.Fatal("bn256.G1: add with neutral element failed")
	}
}

func TestG2(t *testing.T) {
	suite := NewSuite()
	k := suite.G2().Scalar().Pick(random.New())
	pa := suite.G2().Point().Mul(k, nil)
	ma, err := pa.MarshalBinary()
	require.Nil(t, err)
	pb := new(bn256.G2).ScalarBaseMult(&k.(*mod.Int).V)
	mb := pb.Marshal()
	mb = append([]byte{0x01}, mb...)
	require.Equal(t, ma, mb)
}

func TestG2Marshal(t *testing.T) {
	suite := NewSuite()
	k := suite.G2().Scalar().Pick(random.New())
	pa := suite.G2().Point().Mul(k, nil)
	ma, err := pa.MarshalBinary()
	require.Nil(t, err)
	pb := suite.G2().Point()
	err = pb.UnmarshalBinary(ma)
	require.Nil(t, err)
	mb, err := pb.MarshalBinary()
	require.Nil(t, err)
	require.Equal(t, ma, mb)
}

func TestG2Ops(t *testing.T) {
	suite := NewSuite()
	a := suite.Point().Pick(random.New())
	b := suite.Point().Pick(random.New())
	c := a.Clone()
	a.Neg(a)
	a.Neg(a)
	if !a.Equal(c) {
		t.Fatal("bn256.G2: neg failed")
	}
	a.Add(a, b)
	a.Sub(a, b)
	if !a.Equal(c) {
		t.Fatal("bn256.G2: add sub failed")
	}
	a.Add(a, suite.G2().Point().Null())
	if !a.Equal(c) {
		t.Fatal("bn256.G2: add with neutral element failed")
	}
}

func TestPairingCheck(t *testing.T) {
	suite := NewSuite()
	//Get G2 public key
	x := suite.G2().Scalar().Pick(random.New())
	X1 := suite.G2().Point().Mul(x, nil)
	m1, err := X1.MarshalBinary()
	require.Nil(t, err)
	//Get G2 base point
	base1 := suite.G2().Point().Base()
	mbase1, err := base1.MarshalBinary()
	require.Nil(t, err)
	//Get G1 H(m)
	h := suite.G1().Scalar().Pick(random.New())
	hm := suite.G1().Point().Mul(h, nil)
	mh1, err := hm.MarshalBinary()
	require.Nil(t, err)
	//Get G1 signature
	sig := suite.G1().Point().Mul(x, hm)
	sig.Neg(sig)
	msig1, err := sig.MarshalBinary()
	require.Nil(t, err)
	if !suite.PairingCheck([]kyber.Point{sig, hm}, []kyber.Point{base1, X1}) {
		t.Fatal("bn256.GT: PairingCheck failed")
	}

	//Get G1 H(m) on ethereum
	hm2 := new(bn256.G1).ScalarBaseMult(&h.(*mod.Int).V)
	mh2 := hm2.Marshal()
	require.Equal(t, mh1, mh2)

	//Get G1 signature on ethereum
	sig2 := new(bn256.G1).ScalarMult(hm2, &x.(*mod.Int).V)
	sig2.Neg(sig2)
	msig2 := sig2.Marshal()
	require.Equal(t, msig1, msig2)

	//Get G2 public key on ethereum
	X2 := new(bn256.G2).ScalarBaseMult(&x.(*mod.Int).V)
	m2 := X2.Marshal()
	m2 = append([]byte{0x01}, m2...)
	require.Equal(t, m1, m2)

	//Get G2 base point on ethereum
	x.SetInt64(1)
	base2 := new(bn256.G2).ScalarBaseMult(&x.(*mod.Int).V)
	mbase2 := base2.Marshal()
	mbase2 = append([]byte{0x01}, mbase2...)
	require.Equal(t, mbase1, mbase2)
	if !bn256.PairingCheck([]*bn256.G1{sig2, hm2}, []*bn256.G2{base2, X2}) {
		t.Fatal("bn256.GT: PairingCheck failed")
	}
}

func TestGTMarshal(t *testing.T) {
	suite := NewSuite()
	k := suite.GT().Scalar().Pick(random.New())
	pa := suite.GT().Point().Mul(k, nil)
	ma, err := pa.MarshalBinary()
	require.Nil(t, err)
	pb := suite.GT().Point()
	err = pb.UnmarshalBinary(ma)
	require.Nil(t, err)
	mb, err := pb.MarshalBinary()
	require.Nil(t, err)
	require.Equal(t, ma, mb)
}

func TestGTOps(t *testing.T) {
	suite := NewSuite()
	a := suite.GT().Point().Pick(random.New())
	b := suite.GT().Point().Pick(random.New())
	c := a.Clone()
	a.Neg(a)
	a.Neg(a)
	if !a.Equal(c) {
		t.Fatal("bn256.GT: neg failed")
	}
	a.Add(a, b)
	a.Sub(a, b)
	if !a.Equal(c) {
		t.Fatal("bn256.GT: add sub failed")
	}
	a.Add(a, suite.GT().Point().Null())
	if !a.Equal(c) {
		t.Fatal("bn256.GT: add with neutral element failed")
	}
}

func TestBilinearity(t *testing.T) {
	suite := NewSuite()
	a := suite.G1().Scalar().Pick(random.New())
	pa := suite.G1().Point().Mul(a, nil)
	b := suite.G2().Scalar().Pick(random.New())
	pb := suite.G2().Point().Mul(b, nil)
	pc := suite.Pair(pa, pb)
	pd := suite.Pair(suite.G1().Point().Base(), suite.G2().Point().Base())
	pd = suite.GT().Point().Mul(a, pd)
	pd = suite.GT().Point().Mul(b, pd)
	require.Equal(t, pc, pd)
}

func TestTripartiteDiffieHellman(t *testing.T) {
	suite := NewSuite()
	a := suite.G1().Scalar().Pick(random.New())
	b := suite.G1().Scalar().Pick(random.New())
	c := suite.G1().Scalar().Pick(random.New())
	pa, pb, pc := suite.G1().Point().Mul(a, nil), suite.G1().Point().Mul(b, nil), suite.G1().Point().Mul(c, nil)
	qa, qb, qc := suite.G2().Point().Mul(a, nil), suite.G2().Point().Mul(b, nil), suite.G2().Point().Mul(c, nil)
	k1 := suite.Pair(pb, qc)
	k1 = suite.GT().Point().Mul(a, k1)
	k2 := suite.Pair(pc, qa)
	k2 = suite.GT().Point().Mul(b, k2)
	k3 := suite.Pair(pa, qb)
	k3 = suite.GT().Point().Mul(c, k3)
	require.Equal(t, k1, k2)
	require.Equal(t, k2, k3)
}

func TestCombined(t *testing.T) {
	// Making sure we can do some basic arithmetic with the suites without having
	// to extract the suite using .G1(), .G2(), .GT()
	basicPointTest(t, NewSuiteG1())
	basicPointTest(t, NewSuiteG2())
	basicPointTest(t, NewSuiteGT())
}

func basicPointTest(t *testing.T, s *Suite) {
	a := s.Scalar().Pick(random.New())
	pa := s.Point().Mul(a, nil)

	b := s.Scalar().Add(a, s.Scalar().One())
	pb1 := s.Point().Mul(b, nil)
	pb2 := s.Point().Add(pa, s.Point().Base())
	require.True(t, pb1.Equal(pb2))
}
