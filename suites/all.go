package suites

import (
	"github.com/DOSNetwork/core-lib/group/edwards25519"
	"github.com/DOSNetwork/core-lib/group/bn256"
)

func init() {
	register(edwards25519.NewBlakeSHA256Ed25519())
	register(bn256.NewSuite())
}
