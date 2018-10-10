package example

import (
	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/sign/tbls"
	"github.com/DOSNetwork/core/suites"
	"math/big"

	"crypto/sha256"
	"errors"
	"github.com/dedis/kyber"
)

type RandomNumberGenerator struct {
	suite suites.Suite

	nbParticipants int

	partPubs []kyber.Point
	partSec  []kyber.Scalar

	dkgs []*dkg.DistKeyGenerator
	dkss []*dkg.DistKeyShare
}

// InitialRandomNumberGenerator sets up one RandomNumberGenerator for a DKG group according to chosen suite
// and participant count, It also runs DKG to generate group public key and private key shares for each group member
func InitialRandomNumberGenerator(suite suites.Suite, nbParticipants int) (*RandomNumberGenerator, error) {
	partPubs := make([]kyber.Point, nbParticipants)
	partSec := make([]kyber.Scalar, nbParticipants)
	for i := 0; i < nbParticipants; i++ {
		sec, pub := genPair(suite)
		partPubs[i] = pub
		partSec[i] = sec
	}

	dkgs, err := dkgGen(suite, nbParticipants, partPubs, partSec)
	if err != nil {
		return nil, err
	}

	err = fullExchange(dkgs, nbParticipants)
	if err != nil {
		return nil, err
	}

	dkss := make([]*dkg.DistKeyShare, nbParticipants)
	for i, dkg := range dkgs {
		dks, err := dkg.DistKeyShare()
		if err != nil {
			return nil, err
		}
		dkss[i] = dks
	}

	return &RandomNumberGenerator{
		suite:          suite,
		nbParticipants: nbParticipants,
		partPubs:       partPubs,
		partSec:        partSec,
		dkgs:           dkgs,
		dkss:           dkss,
	}, nil
}

func (r *RandomNumberGenerator) GetPubKey() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	pubKey := share.NewPubPoly(r.suite, r.suite.Point().Base(), r.dkss[0].Commitments()).Commit()
	pubKeyMar, err := pubKey.MarshalBinary()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	x0 := big.NewInt(0)
	x1 := big.NewInt(0)
	y0 := big.NewInt(0)
	y1 := big.NewInt(0)
	x0.SetBytes(pubKeyMar[1:33])
	x1.SetBytes(pubKeyMar[33:65])
	y0.SetBytes(pubKeyMar[65:97])
	y1.SetBytes(pubKeyMar[97:])
	return x0, x1, y0, y1, nil
}

// generate takes a seed to generate random number by SHA256 using the tBLS signature to the seed
func (r *RandomNumberGenerator) generate(seed []byte) ([]byte, error) {
	sig, err := r.TBlsSign(seed)
	if err == nil {
		hashSig := sha256.Sum256(sig)
		converted := hashSig[:]
		return converted, nil
	} else {
		return nil, err
	}
}

func (r *RandomNumberGenerator) TBlsSign(msg []byte) ([]byte, error) {
	shares := make([]*share.PriShare, r.nbParticipants)
	sigShares := make([][]byte, 0)
	for i, dks := range r.dkss {
		shares[i] = dks.Share
		if i < (r.nbParticipants/2 + 1) {
			sig, _ := tbls.Sign(r.suite, shares[i], msg)
			sigShares = append(sigShares, sig)
		}
	}
	dks := r.dkss[0]
	pubPoly := share.NewPubPoly(r.suite, r.suite.Point().Base(), dks.Commitments())
	sig, _ := tbls.Recover(r.suite, pubPoly, msg, sigShares, r.nbParticipants/2+1, r.nbParticipants)
	err := bls.Verify(r.suite, pubPoly.Commit(), msg, sig)
	if err != nil {
		return nil, err
	} else {
		return sig, nil
	}
}

func fullExchange(dkgs []*dkg.DistKeyGenerator, nbParticipants int) error {

	// full secret sharing exchange

	// 1. broadcast deals
	resps := make([]*dkg.Response, 0, nbParticipants*nbParticipants)
	for _, dkg := range dkgs {
		deals, err := dkg.Deals()
		if err != nil {
			return err
		}
		for i, d := range deals {
			resp, err := dkgs[i].ProcessDeal(d)
			if err != nil {
				return err
			}
			resps = append(resps, resp)
		}
	}

	// 2. Broadcast responses
	for _, resp := range resps {
		for index, dkg := range dkgs {
			// Ignore messages about ourselves
			if resp.Response.Index == uint32(index) {
				continue
			}
			j, err := dkg.ProcessResponse(resp)
			if err != nil {
				return err
			}
			if j != nil {
				return errors.New("dkg: Complaint received, stop at secret sharing exchanging stage")
			}
		}
	}

	// 3. make sure everyone has the same QUAL set
	//for _, dkg := range dkgs {
	//	for _, dkg2 := range dkgs {
	//		require.True(t, dkg.isInQUAL(dkg2.index))
	//	}
	//}

	return nil
}

func genPair(suite suites.Suite) (kyber.Scalar, kyber.Point) {
	sc := suite.Scalar().Pick(suite.RandomStream())
	return sc, suite.Point().Mul(sc, nil)
}

func dkgGen(suite suites.Suite, nbParticipants int, partPubs []kyber.Point, partSec []kyber.Scalar) ([]*dkg.DistKeyGenerator, error) {
	dkgs := make([]*dkg.DistKeyGenerator, nbParticipants)
	for i := 0; i < nbParticipants; i++ {
		dkg, err := dkg.NewDistKeyGenerator(suite, partSec[i], partPubs, nbParticipants/2+1)
		if err != nil {
			return nil, err
		}
		dkgs[i] = dkg
	}
	return dkgs, nil
}
