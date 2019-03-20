package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/DOSNetwork/core/testing/peerNode/node"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

// main
func main() {
	var ip string
	var err error
	var peerSize int
	var numMessages int
	debug.FreeOSMemory()
	//need to connect to bootstrap node to get credential
	noderole := os.Getenv("NODEROLE")
	bootStrapIP := os.Getenv("BOOTSTRAPIP")
	port := os.Getenv("NODEPORT")
	peerSize, err = strconv.Atoi(os.Getenv("PEERSIZE"))
	if err != nil {
		fmt.Println("PEERSIZE ", err)
	}
	numMessages, err = strconv.Atoi(os.Getenv("NUMOFMESSAGS"))
	if err != nil {
		fmt.Println("NUMOFMESSAGS ", err)
	}
	tStrategy := os.Getenv("TESTSTRATEGY")

	//0)initial log module
	//	log = logrus.New()
	//boot node
	if noderole == "boot" {
		b := new(node.BootNode)
		b.Init(port, peerSize)
		b.EventLoop()
	} else {
		s := strings.Split(bootStrapIP, ":")
		ip, _ = s[0], s[1]
		d := new(node.PeerNode)
		d.Init(ip, port, peerSize, numMessages, tStrategy)
		d.EventLoop()
	}
}
