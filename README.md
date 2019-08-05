# <img align="center" width=40 src="media/logo-white.jpg"> DOS Client and Core Libraries
[![Go Report Card](https://goreportcard.com/badge/github.com/DOSNetwork/core)](https://goreportcard.com/report/github.com/DOSNetwork/core)
[![Maintainability](https://api.codeclimate.com/v1/badges/a2eb5767f8984835fb3b/maintainability)](https://codeclimate.com/github/DOSNetwork/core/maintainability)
[![GoDoc](https://godoc.org/github.com/DOSNetwork/core?status.svg)](https://godoc.org/github.com/DOSNetwork/core)
[![](https://img.shields.io/static/v1.svg?label=chat&message=Telegram&color=brightgreen)](https://t.me/joinchat/KhcP5BQXgWLyojui9BCGfQ)


## Prerequisites:
##### Cloud Server / VPS Recommendations
- [AWS Lightsail](https://aws.amazon.com/lightsail/pricing/?opdp1=pricing) - $5 monthly plan (1CPU, 1GB Memory, 40GB SSD, 2TB Bandwidth)
- [Vultr](https://www.vultr.com/?ref=7806004-4F) - Cloud Compute $5 monthly plan (1CPU, 1GB Memory, 25GB SSD, 1TB Bandwidth)
- [Digital Ocean](https://m.do.co/c/a912bdc08b78) - Droplet $5 monthly plan (1CPU, 25GB SSD, 1TB Bandwidth)
- [Linode](https://www.linode.com/?r=35c0c22d412b3fc8bd98b4c7c6f5ac42ae3bc2e2) - $5 monthly plan (1CPU, 1GB Memory, 25GB SSD, 1TB Bandwidth)
- .

##### Verified and recommended installation environment
- Ubuntu 16.04 x64 LTS or higher 
- WebSocket of Geth endpoints (at least one)
  - sign up with Infura
    - Visit the following URL: https://infura.io/register
    - Submit registration form by filling in the relevant details
    - Confirm the email receipt
  - creating a new project
    - Create your project by providing the project name 
    - Go to VIEW PROJECT
    - Chose ENDPOINT as RINKEBY
    - Copy PROJECT ID to config.json
- An IPv4 address
  - Run `$ dig +short myip.opendns.com @resolver1.opendns.com`
  - Or get it from cloud server providers. Most vps / cloud server 
- With below ports open:
  - **udp** port `7946`
  - **tcp** port `7946`,`9501`
- It's recommended to generate ssh login key pairs and setup public key authentication instead of using password login for server security and funds safety:
  - Learn [how to](https://www.digitalocean.com/community/tutorials/how-to-set-up-ssh-keys-on-ubuntu-1604) setup SSH public key authentication on Ubuntu 16.04 and disable password logins.


##### Acquire testnet ether and testnet tokens
- Acquire testnet ether from rinkeby [faucet](https://faucet.rinkeby.io/).
- Acquire 50,000 [testnet DOS token](https://rinkeby.etherscan.io/address/0x214e79c85744cd2ebbc64ddc0047131496871bee), (and optional - acquire several [testnet DropBurn token](https://rinkeby.etherscan.io/address/0x9bfe8f5749d90eb4049ad94cc4de9b6c4c31f822)).
- Please fill in [this](https://docs.google.com/forms/d/e/1FAIpQLSdiWuVdyxpVozEC0uWZIj9HCBX9COBYFj8Dxp2C2qX4Qv5U9g/viewform) form to request testnet tokens.

## Building from source
- [Install](https://golang.org/doc/install) Go and setup golang workingspace like below:
    ```sh
    $ cd /usr/local
    $ wget https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz
    $ tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz
    ```
    
- Open `~/.bashrc` and set `$GOPATH` and `$PATH` environmental variables:
    ```sh
    $ vim ~/.bashrc
      export GOPATH=$HOME/go
      export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
    $ source ~/.bashrc
    ```
- Install [dep](https://golang.github.io/dep/docs/installation.html#binary-installation) to manage package dependencies.

- Download source code:
    ```sh
    $ mkdir -p $GOPATH/src/github.com/DOSNetwork
    $ cd $GOPATH/src/github.com/DOSNetwork && git clone https://github.com/DOSNetwork/core.git
    $ cd core
    ```

- Install `build-essential` if not already done
    ```sh
    $ sudo apt-get install build-essential
    ```
- Install `go-dep` if not already done
    ```sh
    $ cd $GOPATH
    $ mkdir -p bin
    $ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
    ```    
- Build:
  - `$ make vendor` - to prepare dependencies for building 
  - `$ make` - to build release version client

- Dev tips:
  - `$ go fmt ./...` to reformat go source code.
  - `$ golint` to fix style mistakes conflicting with [effective go](https://golang.org/doc/effective_go.html). ([golint](https://github.com/golang/lint) tool for vim users.)
  - `$ make devClient` to build develoment version client.
  - `$ make updateSubmodule` to fetch latest system contracts from [repo](https://github.com/DOSNetwork/eth-contracts), instead of making contract modifications locally.
  - `$ make gen` to generate binding files for system contracts.
  - `$ make clean` to remove built binaries or unnecessary generated files.

## Building using Docker
- Install `docker` if not already done
    ```sh
    $ sudo apt-get purge docker lxc-docker docker-engine docker.io
    $ sudo apt-get install  curl  apt-transport-https ca-certificates software-properties-common
    $ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add 
    $ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
    $ sudo apt-get update
    $ sudo apt-get install docker-ce
    ```
- Build:
  - `$ make client-docker`
  
## Configure the client
- Config following fields in [`config.json`](https://github.com/DOSNetwork/core/blob/master/config.json) file:
  - `NodeIP`: Public ip address of the server/vps. 
  - `ChainNodePool`: User can add more ws (web socket) of infura endpoints and more ws (web socket) of geth full nodes here.
  - Example:
    Assuming one has an infura account that INFURA-PROJECT-ID is "abcdefghijk"
	```
	"NodeIP":"111.111.111.111",
	"ChainNodePool": [
                "wss://rinkeby.infura.io/ws/v3/abcdefghijk"
        ]
	```
    Assuming one has a geth full node that ip is 22.22.22.22
	```
	"NodeIP":"111.111.111.111",
	"ChainNodePool": [
                "ws://22.22.22.22:8546"
        ]
	```
    Assuming one has a geth full node and an infura account
	```
	"NodeIP":"111.111.111.111",
	"ChainNodePool": [
                "ws://22.22.22.22:8546",
		"wss://rinkeby.infura.io/ws/v3/abcdefghijk"
        ]
	```
## Run standalone binary
- Create an node wallet (Only for the first time):
    ```sh
    $ ./client wallet create
    ```	
- Run client :
    ```sh
    $ ./client start
    ```
- Run client on the background (Recommended):
    ```sh
    $ echo -n Password:;read -s password ;export PASSPHRASE=$password ;nohup ./client start &
    ```
- Check client status :
	./client status
    ```sh
    $ ./client status
    ```
- Stop client :
    ```sh
    $ ./client stop
    ```
- Debuging an issue :
    ```sh
    $ cat vault/doslog.txt
    ```

### Run client node using Docker
- Run client on the background (Recommended):
    ```sh
    $ ./vps_docker.sh start
    ```
- Check client status :
    ```sh
    $ ./vps_docker.sh status
    ```
- Stop client :
    ```sh
    $ ./vps_docker.sh stop
    ```
- Debuging an issue :
    ```sh
    $ cat vault/doslog.txt
    ```


## Status
- [x] Verifiable Secret Sharing
- [x] Distributed Key Generation (Pedersen's DKG approach)
- [x] Paring Library and Threshold BLS Signature
- [x] Distributed Randomness Generation
- [x] Gossip & DHT Implementation
- [x] P2P NAT Support
- [x] Json / Xml / Html Request Parser
- [x] Dockerization and Client Deployment Script
- [x] Integration with Ethereum On-chain [System Contracts](https://github.com/DOSNetwork/eth-contracts)
- [x] P2P Network Performance Tuning
- [ ] Test with geth lightnode mode and experiment with parity clients
- [ ] Staking & Delegation Contracts with a User-friendly Dashboard
- [ ] Network Status Scanner/Explorer
