package main

import (
	"fmt"

	"github.com/DOSNetwork/core/examples/protobuf/messages"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

//protoc -I=. -I=$GOPATH/src/github.com/golang --go_out=. message.proto
func main() {

	//raw, _ := types.MarshalAny(&messages.Company{Id: 2})
	raw, _ := ptypes.MarshalAny(&messages.Person{Name: "Eric"})
	msg := &messages.Message{
		Raw: raw,
	}
	bytes, _ := proto.Marshal(msg)
	//----------------------------------------------------------------------
	recemsg := new(messages.Message)
	_ = proto.Unmarshal(bytes, recemsg)

	var ptr ptypes.DynamicAny
	_ = ptypes.UnmarshalAny(recemsg.Raw, &ptr)
	switch msgRaw := ptr.Message.(type) {
	case *messages.Person:
		fmt.Println("msgRaw messages.Person", messages.Person(*msgRaw).Name)
	case *messages.Company:
		fmt.Println("msgRaw messages.Company", messages.Company(*msgRaw).Id)
	default:
		fmt.Println("msgRaw default")
	}
}
