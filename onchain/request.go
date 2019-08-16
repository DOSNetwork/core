package onchain

import (
	"context"

	"github.com/DOSNetwork/core/onchain/commitreveal"
	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/core/types"
)

type request struct {
	ctx    context.Context
	idx    int
	proxy  *dosproxy.DosproxySession
	cr     *commitreveal.CommitrevealSession
	f      setFunc
	params []interface{}
	reply  chan *response
}

type response struct {
	idx int
	tx  *types.Transaction
	err error
}
