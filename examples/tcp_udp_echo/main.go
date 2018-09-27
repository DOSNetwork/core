package main

import (
	"os"
)

func main() {

	done:= make(chan bool)

	if os.Args[1] == "tcp" {
		if os.Args[2] == "server" {
			TcpServer()
		} else if os.Args[2] == "client" {
			TcpClient()
		}
	} else if os.Args[1] == "udp" {
		if os.Args[2] == "server" {
			UdpServer()
		} else if os.Args[2] == "client" {
			UdpClient()
		}
	}

	<- done
}