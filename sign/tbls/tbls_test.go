package tbls

import (
	"fmt"
	"testing"

	"github.com/DOSNetwork/core/share"
	"github.com/DOSNetwork/core/sign/bls"
	"github.com/DOSNetwork/core/suites"

	"github.com/stretchr/testify/require"
)

var suite = suites.MustFind("bn256")

func TestTBLS(test *testing.T) {
	var err error
	msg := []byte("Hello threshold Boneh-Lynn-Shacham")
	n := 10
	t := n/2 + 1
	secret := suite.Scalar().Pick(suite.RandomStream())
	priPoly := share.NewPriPoly(suite, t, secret, suite.RandomStream())
	pubPoly := priPoly.Commit(suite.Point().Base())
	sigShares := make([][]byte, 0)
	for i, x := range priPoly.Shares(n) {
		sig, err := Sign(suite, x, msg)
		require.Nil(test, err)
		sigShares = append(sigShares, sig)
		//To simulate wrong signatures case
		if i <= 3 {
			sig[0] ^= 0x01
		}
		//To simulate duplicate signatures case
		sigShares = append(sigShares, sig)
		sigShares = append(sigShares, sig)
	}
	fmt.Println(len(sigShares))
	sig, err := Recover(suite, pubPoly, msg, sigShares, t, n)
	require.Nil(test, err)
	err = bls.Verify(suite, pubPoly.Commit(), msg, sig)
	require.Nil(test, err)
}
