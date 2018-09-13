package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	nat2 "github.com/DOSNetwork/core/p2p/nat"
)

var externalAddress net.IP
var internalAddress net.IP
var externalPort = 3323
var internalPort = 9854
var err error

func server() {
	private, err := nat2.IsPrivateIp()
	if err != nil {
		log.Fatal(err)
	}

	if private {
		fmt.Println("private ip detected, searching for nat device...")
	} else {
		fmt.Println("public ip, no nat needed")
		os.Exit(0)
	}

	//use setMapping to search for nat device and build port mapping
	nat, err := nat2.SetMapping("tcp", externalPort, internalPort, "DOSNode")
	if err != nil {
		log.Fatal(err)
	}

	externalAddress, err = nat.GetExternalAddress()
	if err != nil {
		log.Fatal(err)
	}

	internalAddress, err = nat.GetInternalAddress()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("external address: ", externalAddress)
	fmt.Println("external port: ", externalPort)
	fmt.Println("internal address: ", internalAddress)
	fmt.Println("internal port: ", internalPort)
	fmt.Println("nat type: ", nat.GetType())

	address := internalAddress.String() + ":" + strconv.Itoa(internalPort)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server listening on ", address)

	go client()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("server accepts")

	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("server receive msg: ", msg)

	conn.Write([]byte("echo " + msg))

	conn.Close()
	listener.Close()
	nat.CloseMapping()

}

func client()  {
	time.Sleep(2 * time.Second)
	address := externalAddress.String() + ":" + strconv.Itoa(externalPort)
	fmt.Println("client dialing ", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	conn.Write([]byte("hello\n"))

	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("client receive msg: ", msg)
	conn.Close()
}

func main() {
	server()
}
