package onchain

import (
	"context"

	//	"github.com/DOSNetwork/core/onchain/commitreveal"
	//	"github.com/DOSNetwork/core/onchain/dosproxy"

	"github.com/ethereum/go-ethereum/core/types"
)

type request struct {
	opCtx context.Context
	f     func(ctx context.Context) (tx *types.Transaction, err error)
	reply chan *response
}

type response struct {
	idx int
	tx  *types.Transaction
	err error
}
