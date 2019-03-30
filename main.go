package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/DOSNetwork/core/dosnode"

	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

var PIDFile = "./dosclient.pid"

func savePID(pid int) {

	file, err := os.Create(PIDFile)
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

func runDos(credentialPath, passphrase string) {

	// Make arrangement to remove PID file upon receiving the SIGTERM from kill command
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		signalType := <-ch
		signal.Stop(ch)

		fmt.Println("Received signal type : ", signalType)

		// remove PID file
		os.Remove(PIDFile)

		os.Exit(0)

	}()
	mux := http.NewServeMux()
	dosclient, err := dosnode.NewDosNode(credentialPath, passphrase)
	if err != nil {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			content := "Error:" + err.Error() + "\n"
			w.Write([]byte(content))
		})
	} else {
		mux.HandleFunc("/", dosclient.Status)
		mux.HandleFunc("/balance", dosclient.Balance)
		mux.HandleFunc("/groups", dosclient.Groups)
	}
	http.ListenAndServe("localhost:8080", mux)
}

func makeRequest(f string, args []byte) ([]byte, error) {

	tServer := "http://localhost:8080/" + f

	req, err := http.NewRequest("POST", tServer, bytes.NewBuffer(args))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	return r, err
}

// main
func main() {
	if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "run" {
		path := ""
		pass := ""
		if len(os.Args) > 3 {
			path = strings.ToLower(os.Args[2])
			pass = strings.ToLower(os.Args[3])
		}
		fmt.Println(path, pass)
		runDos(path, pass)
	}

	app := cli.NewApp()
	app.Name = "client"
	app.Usage = "the dos-client command line interface"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start a dos-client daemon",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "credential.path,c",
					Usage:  "credential.path",
					EnvVar: "CREDENTIALPATH",
				},
				cli.StringFlag{
					Name:   "credential.passphrase,p",
					Usage:  "credential.passPhrase",
					EnvVar: "PASSPHRASE",
				},
			},
			Action: func(c *cli.Context) error {
				password := c.String("credential.passphrase")
				fmt.Println(password)
				for password == "" {
					fmt.Print("Enter Password: ")
					bytePassword, _ := terminal.ReadPassword(0)
					password = strings.TrimSpace(string(bytePassword))
				}

				// check if daemon already running.
				if _, err := os.Stat(PIDFile); err == nil {
					fmt.Println("Already running or /tmp/dosclient.pid file exist.")
					os.Exit(1)
				}
				cmd := exec.Command(os.Args[0], "run", "", password)
				cmd.Start()
				//runDos(c.String("credential.path"), password)
				savePID(cmd.Process.Pid)
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "Stop a daemon",
			Action: func(c *cli.Context) error {
				if _, err := os.Stat(PIDFile); err == nil {
					data, err := ioutil.ReadFile(PIDFile)
					if err != nil {
						fmt.Println("Not running")
						return nil
					}
					ProcessID, err := strconv.Atoi(string(data))

					if err != nil {
						fmt.Println("Unable to read and parse process id found in ", PIDFile)
						return nil
					}

					process, err := os.FindProcess(ProcessID)

					if err != nil {
						fmt.Printf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
						return nil
					}
					// remove PID file
					os.Remove(PIDFile)

					fmt.Printf("Killing process ID [%v] now.\n", ProcessID)
					// kill process and exit immediately
					err = process.Kill()

					if err != nil {
						fmt.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
						return nil
					} else {
						fmt.Printf("Killed process ID [%v]\n", ProcessID)
						return nil
					}

				} else {

					fmt.Println("Not running.")
					return nil
				}
				return nil
			},
		},
		{
			Name:  "show",
			Usage: "Show dos-client status",
			Subcommands: []cli.Command{
				{
					Name:  "status",
					Usage: "show status",
					Action: func(c *cli.Context) error {
						r, _ := makeRequest("/", []byte{})
						fmt.Println(string(r))
						return nil
					},
				},
				{
					Name:  "balance",
					Usage: "show balance",
					Action: func(c *cli.Context) error {
						r, _ := makeRequest("balance", []byte{})
						fmt.Println("show balance: ", string(r))
						return nil
					},
				},
				{
					Name:  "groups",
					Usage: "show group number",
					Action: func(c *cli.Context) error {
						r, _ := makeRequest("groups", []byte{})
						fmt.Println("show group number: ", string(r))
						return nil
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
