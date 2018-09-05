package bn256

// For details of the algorithms used, see "Multiplication and Squaring on
// Pairing-Friendly Fields, Devegili et al.
// http://eprint.iacr.org/2006/471.pdf.

import (
	"math/big"
)

// gfP12 implements the field of size p¹² as a quadratic extension of gfP6
// where ω²=τ.
type gfP12 struct {
	x, y gfP6 // value is xω + y
}

var gfP12Gen = &gfP12{
	x: gfP6{
		x: gfP2{
			x: gfP{0xD34BAB373157AA84, 0x3511ED44FD0D8598, 0x67E42A0BC2CED972, 0x2B8F1D5DFD20C55B},
			y: gfP{0x988AE2485B36CF53, 0x5091CC0581334E54, 0xDA7903229312CA0F, 0x2A2341538EAEE95C},
		},
		y: gfP2{
			x: gfP{0x7002907C28EBFE11, 0x7B0591D3D080DA67, 0xDE7E5AA2181F138E, 0x210E437DFC43D951},
			y: gfP{0x8975B68A2BAB1F9C, 0x2FDD826B796E0F35, 0x6A90A35FA03DFAA5, 0x1FFEF4581607FC37},
		},
		z: gfP2{
			x: gfP{0x9458ABCB56D24998, 0xB17540BD2A9E5ADB, 0x9A9983C82E401A9F, 0x1614817A84C16291},
			y: gfP{0x1BB0CE0DEF1B82A1, 0x4C4C9FE1CADEFA95, 0x746D9990CB12B27E, 0x13495C08E5D415C5},
		},
	},
	y: gfP6{
		x: gfP2{
			x: gfP{0xF16C96D081754CDB, 0xCE0394312BCEEB55, 0x644E4DCF1F01FF0A, 0xCBEA85EE0B236CC},
			y: gfP{0x5CF9CC917DA86724, 0xC799DC487A0B2753, 0xDF2027BF1DE17A7, 0x197CDA6CC3E20636},
		},
		y: gfP2{
			x: gfP{0x172D1F257A4D598E, 0xDDF5BC7B7FFB5AC0, 0xAE0B22C0BBB0F602, 0x1B158F3C2FAE9B18},
			y: gfP{0x2306E4312363B991, 0x465F6072D4023BF4, 0xA2FF062A4A77E736, 0x76EA6F18435864A},
		},
		z: gfP2{
			x: gfP{0x2E02A64ACBD60549, 0xD618018EA58E4ADD, 0x14D585F1A45BA647, 0x1832226987C434FC},
			y: gfP{0xC556F62B2A98671D, 0x23A59AC167BCF363, 0x5EF208445F5F6F37, 0x12ADF27CCB29382A},
		},
	},
}

var gfP12Inf = &gfP12{
	x: gfP6{
		x: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
		y: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
		z: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
	},
	y: gfP6{
		x: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
		y: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		},
		z: gfP2{
			x: gfP{0x0000000000000000, 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
			y: gfP{0xD35D438DC58F0D9D, 0xA78EB28F5C70B3D, 0x666EA36F7879462C, 0xE0A77C19A07DF2F},
		},
	},
}

func (e *gfP12) String() string {
	return "(" + e.x.String() + "," + e.y.String() + ")"
}

func (e *gfP12) Set(a *gfP12) *gfP12 {
	e.x.Set(&a.x)
	e.y.Set(&a.y)
	return e
}

func (e *gfP12) SetZero() *gfP12 {
	e.x.SetZero()
	e.y.SetZero()
	return e
}

func (e *gfP12) SetOne() *gfP12 {
	e.x.SetZero()
	e.y.SetOne()
	return e
}

func (e *gfP12) IsZero() bool {
	return e.x.IsZero() && e.y.IsZero()
}

func (e *gfP12) IsOne() bool {
	return e.x.IsZero() && e.y.IsOne()
}

func (e *gfP12) Conjugate(a *gfP12) *gfP12 {
	e.x.Neg(&a.x)
	e.y.Set(&a.y)
	return e
}

func (e *gfP12) Neg(a *gfP12) *gfP12 {
	e.x.Neg(&a.x)
	e.y.Neg(&a.y)
	return e
}

// Frobenius computes (xω+y)^p = x^p ω·ξ^((p-1)/6) + y^p
func (e *gfP12) Frobenius(a *gfP12) *gfP12 {
	e.x.Frobenius(&a.x)
	e.y.Frobenius(&a.y)
	e.x.MulScalar(&e.x, xiToPMinus1Over6)
	return e
}

// FrobeniusP2 computes (xω+y)^p² = x^p² ω·ξ^((p²-1)/6) + y^p²
func (e *gfP12) FrobeniusP2(a *gfP12) *gfP12 {
	e.x.FrobeniusP2(&a.x)
	e.x.MulGFP(&e.x, xiToPSquaredMinus1Over6)
	e.y.FrobeniusP2(&a.y)
	return e
}

func (e *gfP12) FrobeniusP4(a *gfP12) *gfP12 {
	e.x.FrobeniusP4(&a.x)
	e.x.MulGFP(&e.x, xiToPSquaredMinus1Over3)
	e.y.FrobeniusP4(&a.y)
	return e
}

func (e *gfP12) Add(a, b *gfP12) *gfP12 {
	e.x.Add(&a.x, &b.x)
	e.y.Add(&a.y, &b.y)
	return e
}

func (e *gfP12) Sub(a, b *gfP12) *gfP12 {
	e.x.Sub(&a.x, &b.x)
	e.y.Sub(&a.y, &b.y)
	return e
}

func (e *gfP12) Mul(a, b *gfP12) *gfP12 {
	tx := (&gfP6{}).Mul(&a.x, &b.y)
	t := (&gfP6{}).Mul(&b.x, &a.y)
	tx.Add(tx, t)

	ty := (&gfP6{}).Mul(&a.y, &b.y)
	t.Mul(&a.x, &b.x).MulTau(t)

	e.x.Set(tx)
	e.y.Add(ty, t)
	return e
}

func (e *gfP12) MulScalar(a *gfP12, b *gfP6) *gfP12 {
	e.x.Mul(&e.x, b)
	e.y.Mul(&e.y, b)
	return e
}

func (e *gfP12) Exp(a *gfP12, power *big.Int) *gfP12 {
	sum := (&gfP12{}).SetOne()
	t := &gfP12{}

	for i := power.BitLen() - 1; i >= 0; i-- {
		t.Square(sum)
		if power.Bit(i) != 0 {
			sum.Mul(t, a)
		} else {
			sum.Set(t)
		}
	}

	e.Set(sum)
	return e
}

func (e *gfP12) Square(a *gfP12) *gfP12 {
	// Complex squaring algorithm
	v0 := (&gfP6{}).Mul(&a.x, &a.y)

	t := (&gfP6{}).MulTau(&a.x)
	t.Add(&a.y, t)
	ty := (&gfP6{}).Add(&a.x, &a.y)
	ty.Mul(ty, t).Sub(ty, v0)
	t.MulTau(v0)
	ty.Sub(ty, t)

	e.x.Add(v0, v0)
	e.y.Set(ty)
	return e
}

func (e *gfP12) Invert(a *gfP12) *gfP12 {
	// See "Implementing cryptographic pairings", M. Scott, section 3.2.
	// ftp://136.206.11.249/pub/crypto/pairings.pdf
	t1, t2 := &gfP6{}, &gfP6{}

	t1.Square(&a.x)
	t2.Square(&a.y)
	t1.MulTau(t1).Sub(t2, t1)
	t2.Invert(t1)

	e.x.Neg(&a.x)
	e.y.Set(&a.y)
	e.MulScalar(e, t2)
	return e
}
