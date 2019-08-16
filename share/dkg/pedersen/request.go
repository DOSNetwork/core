package dkg

import (
	"context"
)

type request struct {
	ctx        context.Context
	reqType    int
	sessionID  string
	numOfResps int
	reply      chan []interface{}
}
