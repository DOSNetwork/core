package example

import (
	"testing"
	"fmt"
	"time"

	"github.com/DOSNetwork/core/suites"
	"github.com/stretchr/testify/require"
)

func TestRandomNumberGenerator(t *testing.T) {
	suite := suites.MustFind("bn256")
	nbParticipants := 15
	randomNumberGenerator, err := InitialRandomNumberGenerator(suite, nbParticipants)
	require.Nil(t, err)
	require.NotNil(t, randomNumberGenerator)

	cur := []byte("Hello threshold Boneh-Lynn-Shacham")

	for i := 0; i < 10; i++ {
		start := time.Now()
		cur,err = randomNumberGenerator.generate(cur)
		require.Nil(t, err)
		fmt.Println(cur)
		fmt.Println(time.Since(start))
	}
}