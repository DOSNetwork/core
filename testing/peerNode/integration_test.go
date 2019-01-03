// +build integration

package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
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
	cmd := exec.CommandContext(ctx, "docker-compose", "-f", "peer-docker-compose.yml", "up", "--scale", "peernode=10")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = peerNodeDir()

	if err := cmd.Run(); err != nil {
		os.Stderr.WriteString(err.Error())
		test.FailNow()
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		test.FailNow()
	}

	f, err := os.Create(cmd.Dir + "/NewPeerAndSendMessage.log")
	if err != nil {
		fmt.Println("unable to create log file", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(cmdOutput.Bytes()))
	if err != nil {
		fmt.Println("unable to write to log file", err)
	}
	w.Flush()

	fmt.Println("finishing NewPeerAndSendMessage")
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
	cmd := exec.CommandContext(ctx, "docker-compose", "-f", "findnode-docker-compose.yml", "up", "--scale", "peernode=10")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = peerNodeDir()

	if err := cmd.Run(); err != nil {
		os.Stderr.WriteString(err.Error())
		test.FailNow()
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		test.FailNow()
	}

	f, err := os.Create(cmd.Dir + "/FindNode.log")
	if err != nil {
		fmt.Println("unable to create log file", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(cmdOutput.Bytes()))
	if err != nil {
		fmt.Println("unable to write to log file", err)
	}
	w.Flush()

	fmt.Println("finishing FindNode")
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

func TestDKG(test *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel() // The cancel should be deferred so resources are cleaned up
	cmdOutput := &bytes.Buffer{}
	cmd := exec.CommandContext(ctx, "docker-compose", "-f", "dkg-docker-compose.yml", "up", "--scale", "peernode=10")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = peerNodeDir()

	if err := cmd.Run(); err != nil {
		os.Stderr.WriteString(err.Error())
		test.FailNow()
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		test.FailNow()
	}

	f, err := os.Create(cmd.Dir + "/dkg.log")
	if err != nil {
		fmt.Println("unable to create log file", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(cmdOutput.Bytes()))
	if err != nil {
		fmt.Println("unable to write to log file", err)
	}
	w.Flush()

	fmt.Println("finishing dkg")
	output, err := exec.Command("docker-compose", "-f", "dkg-docker-compose.yml", "down").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))

	err, isfail := isTestFail("dkg")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isfail {
		test.Fail()
	}
}

func TestTBLS(test *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel() // The cancel should be deferred so resources are cleaned up
	cmdOutput := &bytes.Buffer{}
	cmd := exec.CommandContext(ctx, "docker-compose", "-f", "tbls-docker-compose.yml", "up", "--scale", "peernode=10")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = peerNodeDir()

	if err := cmd.Run(); err != nil {
		os.Stderr.WriteString(err.Error())
		test.FailNow()
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		test.FailNow()
	}

	f, err := os.Create(cmd.Dir + "/tbls.log")
	if err != nil {
		fmt.Println("unable to create log file", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(cmdOutput.Bytes()))
	if err != nil {
		fmt.Println("unable to write to log file", err)
	}
	w.Flush()

	fmt.Println("finishing tbls")
	output, err := exec.Command("docker-compose", "-f", "tbls-docker-compose.yml", "down").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))

	err, isfail := isTestFail("tbls")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isfail {
		test.Fail()
	}
}

func TestDKGMultiGrouping(test *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel() // The cancel should be deferred so resources are cleaned up
	cmdOutput := &bytes.Buffer{}
	cmd := exec.CommandContext(ctx, "docker-compose", "-f", "dkgMultiGrouping-docker-compose.yml", "up", "--scale", "peernode=10")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = peerNodeDir()

	if err := cmd.Run(); err != nil {
		os.Stderr.WriteString(err.Error())
		test.FailNow()
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		test.FailNow()
	}

	f, err := os.Create(cmd.Dir + "/dkgMultiGrouping.log")
	if err != nil {
		fmt.Println("unable to create log file", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(cmdOutput.Bytes()))
	if err != nil {
		fmt.Println("unable to write to log file", err)
	}
	w.Flush()

	fmt.Println("finishing dkgMultiGrouping")
	output, err := exec.Command("docker-compose", "-f", "dkgMultiGrouping-docker-compose.yml", "down").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))

	err, isfail := isTestFail("dkgMultiGrouping")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isfail {
		test.Fail()
	}
}
