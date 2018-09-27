package main

import (
	"encoding/binary"
	"fmt"
)

func packMsg(msg []byte, server bool) []byte {
	if server {
		msg = append([]byte("Echo: "), msg...)
	}
	msgSizePrefix := make([]byte, 4)
	binary.LittleEndian.PutUint32(msgSizePrefix, uint32(len(msg)))
	return append(msgSizePrefix, msg...)
}

func unpackMsg(msg []byte) []byte {
	msgSize := binary.LittleEndian.Uint32(msg[:4])
	fmt.Println("msgSize: ", msgSize)
	receivedMsg := msg[4:msgSize + 4]
	fmt.Print("Msg: ", string(receivedMsg))
	return receivedMsg
}