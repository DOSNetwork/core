package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/dosnode"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

// Caching running node's process id.
const pidFile string = "./vault/dosclient.pid"
const logFile string = "./vault/doslog.txt"
const dosPath string = "./vault"

func savePID(pid int) {

	file, err := os.Create(pidFile)
	if err != nil {
		fmt.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))
	if err != nil {
		fmt.Printf("Unable to create pid file : %v\n", err)
		os.Exit(1)
	}

	file.Sync() // flush to disk

}

func runDos() {
	defer os.Remove(pidFile)
	// Check for the directory's existence and create it if it doesn't exist
	if _, err := os.Stat(dosPath); os.IsNotExist(err) {
		os.Mkdir(dosPath, os.ModeDir)
	}

	// check if there is an account
	n := onchain.NumOfAccounts(dosPath)
	if n == 0 {
		fmt.Println("Please run 'client wallet create' to create a new wallet for the node")
		return
	}
	// check if password is in env variable
	password := os.Getenv("PASSPHRASE")
	if len(password) == 0 {
		fmt.Println("Please run 'client start' to start a client daemon and provide a password")
		return
	}
	key, err := onchain.ReadEthKey(dosPath, password)
	if err != nil {
		fmt.Println("ReadEthKey Error : ", err)
		return
	}
	log.Init(key.Address.Bytes()[:])

	//Read Configuration
	config := configuration.Config{}
	if err := config.LoadConfig(); err != nil {
		fmt.Println("LoadConfig Error : ", err)
		return
	}
	//Check if config has all information
	if len(config.ChainNodePool) == 0 {
		fmt.Println("Please provides geth ws address in config.json")
		return
	}
	hasWsAddr := false
	for _, gethAddr := range config.ChainNodePool {
		if strings.Contains(gethAddr, "ws://") {
			hasWsAddr = true
			break
		}

	}
	if !hasWsAddr {
		fmt.Println("Please provides at least one geth ws address in config.json")
		return
	}

	// Make arrangement to remove PID file upon receiving the SIGTERM from kill command
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		//defer profile.Start().Stop()
		defer os.Exit(0)
		signalType := <-ch
		signal.Stop(ch)
		os.Remove(pidFile)
		fmt.Println("Received signal type : ", signalType)
	}()

	fErr, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("OpenLogFile err ", err)
		return
	}
	syscall.Dup2(int(fErr.Fd()), 1) /* -- stdout */
	syscall.Dup2(int(fErr.Fd()), 2) /* -- stderr */

	dosclient, err := dosnode.NewDosNode(key, config)
	if err != nil {
		fmt.Println("NewDosNode err", err)
		return
	}
	dosclient.Start()
	dosclient.Listen()
}

func makeRequest(f string) ([]byte, error) {

	tServer := "http://localhost:8080/" + f

	req, err := http.NewRequest("GET", tServer, nil)
	if err != nil {
		fmt.Println("makeRequest err ", err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("makeRequest err ", err)
		return nil, err
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("makeRequest err ", err)
		return nil, err
	}
	return r, err
}

func getkey() (key *keystore.Key, err error) {
	//Check if there is a keystore
	password := ""
	key, err = onchain.ReadEthKey(dosPath, password)
	if err != nil && err.Error() == "No Account" {
		return
	}
	//Get password from env variable
	password = os.Getenv("PASSPHRASE")
	if len(password) == 0 {
		//Get password from terminal
		password = getPassword("Enter password :")
	}
	key, err = onchain.ReadEthKey(dosPath, password)

	return
}
func getPassword(s string) (p string) {
	for p == "" {
		fmt.Print(s)
		bytePassword, _ := terminal.ReadPassword(0)
		p = strings.TrimSpace(string(bytePassword))
	}
	fmt.Println("")
	return
}
func actionStart(c *cli.Context) error {
	// check if daemon already running.
	if _, err := os.Stat(pidFile); err == nil {
		fmt.Println("Already running or ${PWD}/dosclient.pid file exist.")
		os.Exit(1)
	}
	// check if there is a account
	n := onchain.NumOfAccounts(dosPath)
	if n == 0 {
		fmt.Println("Please run 'client wallet create' to create a new wallet for the node")
		os.Exit(1)
	}

	// check if password is in env variable
	password := os.Getenv("PASSPHRASE")
	if len(password) == 0 {
		os.Setenv("PASSPHRASE", getPassword("Enter Password: "))
	}

	cmd := exec.Command(os.Args[0], "run")
	cmd.Stdout = os.Stdout
	cmd.Start()
	savePID(cmd.Process.Pid)
	return nil
}

func actionStop(c *cli.Context) error {
	_, err := os.Stat(pidFile)
	if err == nil {
		data, err := ioutil.ReadFile(pidFile)
		if err != nil {
			fmt.Println("Not running")
			return err
		}

		ProcessID, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Println("Unable to read and parse process id found in ", pidFile)
			return err
		}

		process, err := os.FindProcess(ProcessID)
		if err != nil {
			fmt.Printf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
			return err
		}

		// remove PID file
		os.Remove(pidFile)
		fmt.Printf("Killing process ID [%v] now.\n", ProcessID)

		// kill process and exit immediately
		err = process.Kill()
		if err != nil {
			fmt.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
			return err
		}
		fmt.Printf("Killed process ID [%v]\n", ProcessID)
		return nil
	}
	fmt.Println("Not running.")
	return err
}

func actionShowStatus(c *cli.Context) error {
	r, err := makeRequest("/")
	if err == nil {
		fmt.Println(string(r))
		return nil
	}
	return err
}

func actionGroupFormation(c *cli.Context) error {
	r, err := makeRequest("/signalGroupFormation")
	if err == nil {
		fmt.Println("trigger guardian functions : \n", string(r))
		return nil
	}
	return err
}

func actionGroupDissolve(c *cli.Context) error {
	r, err := makeRequest("/signalGroupDissolve")
	if err == nil {
		fmt.Println("trigger guardian functions : \n", string(r))
		return nil
	}
	return err
}

func actionBootstrap(c *cli.Context) error {
	c.String("cid")
	r, err := makeRequest("/signalBootstrap?cid=" + c.String("cid"))
	if err == nil {
		fmt.Println("trigger guardian functions : \n", string(r))
		return nil
	}
	return err
}

func actionRnadom(c *cli.Context) error {
	r, err := makeRequest("/signalRandom")
	if err == nil {
		fmt.Println("trigger guardian functions : \n", string(r))
		return nil
	}
	return err
}

func actionCreateWallet(c *cli.Context) error {

	first := getPassword("Generating node wallet...\nEnter passphrase (empty is not allowed): ")
	second := getPassword("Confirm passphrase again: ")
	if first != second {
		fmt.Println("Unmatched Password")
		return errors.New("Unmatched Password")
	}
	err := onchain.GenEthkey(dosPath, first)
	if err != nil {
		fmt.Println("GenEthkey error : ", err)
	} else {
		key, err := onchain.ReadEthKey(dosPath, first)
		if err != nil {
			fmt.Println("Error :", err)
			return err
		}
		fmt.Println("wallet keystore file has been saved under", dosPath)
		fmt.Println("Your node wallet address is:", fmt.Sprintf("0x%x", key.Address))
	}
	return nil
}

func actionWalletInfo(c *cli.Context) error {
	return nil
}

func actionWalletBalance(c *cli.Context) error {
	r, err := makeRequest("/balance")
	if err == nil {
		fmt.Println("show balance: ", string(r))
		return nil
	}
	return err
}

// main
func main() {
	if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "run" {
		runDos()
		return
	}

	app := cli.NewApp()
	app.Name = "client"
	app.Version = "beta"
	app.Usage = "CLI for dos client"

	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "Start a dos client daemon",
			Action: actionStart,
		},
		{
			Name:   "stop",
			Usage:  "Stop a dos client daemon",
			Action: actionStop,
		},
		{
			Name:   "status",
			Usage:  "show dos client status",
			Action: actionShowStatus,
		},
		{
			Name:  "guardian",
			Usage: "Guardian functions",
			Subcommands: []cli.Command{
				{
					Name:   "groupFormation",
					Usage:  "signal proxy to generate a new group",
					Action: actionGroupFormation,
				},
				{
					Name:   "groupDissolve",
					Usage:  "signal proxy to dissolve expired groups",
					Action: actionGroupDissolve,
				},
				{
					Name:   "bootStrap",
					Usage:  "signal proxy to bootstrape",
					Action: actionBootstrap,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "cid, bootstrap campaign id"},
					},
				},
				{
					Name:   "random",
					Usage:  "signal proxy to generate a random number",
					Action: actionRnadom,
				},
			},
		},
		{
			Name:  "wallet",
			Usage: "Manage Node Wallet",
			Subcommands: []cli.Command{
				{
					Name:   "create",
					Usage:  "show wallet status",
					Action: actionCreateWallet,
				},
				{
					Name:   "balance",
					Usage:  "show wallet balance",
					Action: actionWalletBalance,
				},
				{
					Name:   "info",
					Usage:  "show wallet info",
					Action: actionWalletInfo,
				},
			},
		},
	}

	app.Run(os.Args)
}
