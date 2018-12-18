package p2p

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func NewPeerSendMessageTest(test *testing.T) {
	cmdOutput := &bytes.Buffer{}
	cmd := exec.Command("docker-compose", "-f", "peer-docker-compose.yml", "up", "--scale", "peernode=5")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "DOSNetwork", "core", "testing", "peerNode")

	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	f, err := os.Create(cmd.Dir + "/log")
	defer f.Close()
	w := bufio.NewWriter(f)
	n4, err := w.WriteString(string(cmdOutput.Bytes()))
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
