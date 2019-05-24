package dosnode

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p"
)

//func choseSubmitter(ctx context.Context, p p2p.P2PInterface, lastSysRand *big.Int, ids [][]byte, outCount int) ([]chan []byte, <-chan error) {

func TestChoseSubmitter(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ids := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	lastSysRand := big.NewInt(2)
	outCount := 1
	nodeId := []byte{'a', 'b', 'c'}
	os.Setenv("PUBLICIP", "0.0.0.0")
	log.Init(nodeId)
	logger := log.New("module", "dostage")
	p, err := p2p.CreateP2PNetwork(nodeId, "9905", p2p.SWIM)
	if err != nil {
		t.Errorf("Error %s", err)
	}
	p.Listen()

	submitterc, errc := choseSubmitter(ctx, p, lastSysRand, ids, outCount, logger)
	select {
	case s := <-submitterc[0]:
		fmt.Println("Submitter ", string(s))
		if string(s) != string(nodeId) {
			t.Errorf("TestChoseSubmitter ,Expected %s Actual %s", string(nodeId), string(s))
		}
	case err := <-errc:
		t.Errorf("Error %s", err)
	case <-ctx.Done():
		t.Errorf("Error %s", ctx.Err())
	}
}
