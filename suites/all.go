package suites

import (
	"github.com/DOSNetwork/core/group/edwards25519"
	"github.com/DOSNetwork/core/group/bn256"
)

func init() {
	register(edwards25519.NewBlakeSHA256Ed25519())
	register(bn256.NewSuite())
}
