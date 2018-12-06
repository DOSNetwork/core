// +build integration

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"context"
	"time"
)

const ErrorCode = "[ERROR]"

func peerNodeDir() string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "DOSNetwork", "core", "testing", "peerNode")
}

func isTestFail(test string) (error, bool) {
	filename := peerNodeDir() + "/" + test + ".log"
	f, err := os.Open(filename)
	if err != nil {
		return err, true
	}
	content, _ := ioutil.ReadAll(f)
	contentstr := string(content)
	return nil, strings.Index(contentstr, ErrorCode) != -1
}

func TestNewPeerAndSendMessage(test *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel() // The cancel should be deferred so resources are cleaned up
	cmdOutput := &bytes.Buffer{}
	cmd := exec.CommandContext(ctx, "docker-compose", "-f", "peer-docker-compose.yml", "up", "--scale", "peernode=5")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = peerNodeDir()

	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		return
	}

	if err != nil {
		os.Stderr.WriteString(err.Error())
		test.Fail()
	}
	f, err := os.Create(cmd.Dir + "/NewPeerAndSendMessage.log")
	if err != nil {
		test.Fail()
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(cmdOutput.Bytes()))
	if err != nil {
		fmt.Println(err.Error())
		test.Fail()
	}
	w.Flush()

	output, err := exec.Command("docker-compose", "-f", "peer-docker-compose.yml", "down").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))

	err, isfail := isTestFail("NewPeerAndSendMessage")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isfail {
		test.Fail()
	}
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
		test.Fail()
	}
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		return
	}
	f, err := os.Create(cmd.Dir + "/FindNode.log")
	if err != nil {
		test.Fail()
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(cmdOutput.Bytes()))
	if err != nil {
		test.Fail()
	}
	w.Flush()
	fmt.Println("finishin test")
	output, err := exec.Command("docker-compose", "-f", "findnode-docker-compose.yml", "down").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))

	err, isfail := isTestFail("FindNode")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isfail {
		test.Fail()
	}
}
