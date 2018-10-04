package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

var port = ":1200"

func TcpServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on ", listener.Addr(), " through TCP")

	for true {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go tcpMsgHandler(conn)
	}
}

func UdpServer() {
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on ", conn.LocalAddr(), " through UDP")
	for true {
		udpMsgHandler(conn)
	}
}

func udpMsgHandler(conn *net.UDPConn) {
	receivedRaw := make([]byte, 1024)
	_, remoteAddr, _ := conn.ReadFromUDP(receivedRaw)
	fmt.Println("msg from: ", remoteAddr)

	msg := packMsg(unpackMsg(receivedRaw), true)
	conn.WriteToUDP(msg, remoteAddr)
}

func tcpMsgHandler(conn *net.TCPConn) {
	fmt.Println("tcp connection established with ", conn.RemoteAddr())

	msgSizePrefix := make([]byte, 4)

	for true {
		_, err := conn.Read(msgSizePrefix)
		if err == io.EOF {
			break
		}
		fmt.Println("msg from ", conn.RemoteAddr())
		msgSize := binary.LittleEndian.Uint32(msgSizePrefix)
		fmt.Println("msgSize: ", msgSize)
		msg := make([]byte, msgSize)
		conn.Read(msg)
		fmt.Print("msg: ", string(msg))
		conn.Write(packMsg(msg, true))
	}
}
