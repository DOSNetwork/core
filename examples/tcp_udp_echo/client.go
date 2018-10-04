package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
)

var serverPort = ":1200"

func TcpClient() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverPort)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to ", conn.RemoteAddr(), " through TCP.")

	go tcpMsgReader(conn)

	go tcpMsgSender(conn)

}

func UdpClient() {
	udpAddr, err := net.ResolveUDPAddr("udp", serverPort)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to ", conn.RemoteAddr(), " through UDP.")

	go udpMsgReader(conn)

	go udpMsgSender(conn)
}

func udpMsgReader(conn *net.UDPConn) {
	receivedRaw := make([]byte, 1024)

	for true {
		conn.ReadFromUDP(receivedRaw)
		unpackMsg(receivedRaw)
	}
}

func udpMsgSender(conn *net.UDPConn) {
	fmt.Print("Enter text: ")

	for true {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		msgContent := []byte(text)
		msg := packMsg(msgContent, false)
		conn.Write(msg)
	}
}

func tcpMsgReader(conn *net.TCPConn) {
	msgSizePrefix := make([]byte, 4)
	for true {
		conn.Read(msgSizePrefix)
		msgSize := binary.LittleEndian.Uint32(msgSizePrefix)
		fmt.Println("msgSize: ", msgSize)
		msg := make([]byte, msgSize)
		conn.Read(msg)
		fmt.Print("msg: ", string(msg))
	}
}

func tcpMsgSender(conn *net.TCPConn) {
	fmt.Print("Enter text: ")

	for true {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		msgContent := []byte(text)
		msg := packMsg(msgContent, false)
		conn.Write(msg)
	}
}
