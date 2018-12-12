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

func dosNodeDir() string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "DOSNetwork", "core", "testing", "dosUser")
}

func isTestFail(test string) (error, bool) {
	filename := dosNodeDir() + "/" + test + ".log"
	f, err := os.Open(filename)
	if err != nil {
		return err, true
	}
	content, _ := ioutil.ReadAll(f)
	contentstr := string(content)
	return nil, strings.Index(contentstr, ErrorCode) != -1
}


func TestDosNodeStart(test *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel() // The cancel should be deferred so resources are cleaned up
	cmdOutput := &bytes.Buffer{}
	cmd := exec.CommandContext(ctx, "docker-compose", "-f", "dos-docker-compose.yml", "up", "--scale", "dosnode=3")
	cmd.Stdout = cmdOutput
	cmd.Stderr = os.Stderr
	cmd.Dir = dosNodeDir()

	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
		return
	}

	if err != nil {
		os.Stderr.WriteString(err.Error())
		test.Fail()
	}
	f, err := os.Create(cmd.Dir + "/DosNodeStartMessage.log")
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

	err, isfail := isTestFail("DosNodeStartMessage")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isfail {
		test.Fail()
	}
}


