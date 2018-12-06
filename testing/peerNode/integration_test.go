// +build integration

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"context"
	"time"
)

func TestNewPeerAndSendMessage(test *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel() // The cancel should be deferred so resources are cleaned up
	cmdOutput := &bytes.Buffer{}
	cmd := exec.CommandContext(ctx, "docker-compose", "-f", "peer-docker-compose.yml", "up", "--scale", "peernode=5")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "DOSNetwork", "core", "testing", "peerNode")

	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		return
	}

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	f, err := os.Create(cmd.Dir + "/NewPeerAndSendMessage.log")
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(cmdOutput.Bytes()))
	w.Flush()
	fmt.Println("finishin test")

	output, err := exec.Command("docker-compose", "-f", "peer-docker-compose.yml", "down").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))
}

func TestFindNode(test *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel() // The cancel should be deferred so resources are cleaned up
	cmdOutput := &bytes.Buffer{}
	cmd := exec.CommandContext("docker-compose", "-f", "findnode-docker-compose.yml", "up", "--scale", "peernode=5")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "DOSNetwork", "core", "testing", "peerNode")

	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		return
	}
	f, err := os.Create(cmd.Dir + "/FindNode.log")
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(cmdOutput.Bytes()))
	w.Flush()
	fmt.Println("finishin test")
	output, err := exec.Command("docker-compose", "-f", "findnode-docker-compose.yml", "down").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))
}
