# <img align="left" width=30 src="media/logo-white.jpg"> DOS Client and Core Libraries
 [![Go Report Card](https://goreportcard.com/badge/github.com/DOSNetwork/core)](https://goreportcard.com/report/github.com/DOSNetwork/core)
 [![Maintainability](https://api.codeclimate.com/v1/badges/a2eb5767f8984835fb3b/maintainability)](https://codeclimate.com/github/DOSNetwork/core/maintainability)
 [![GoDoc](https://godoc.org/github.com/DOSNetwork/core?status.svg)](https://godoc.org/github.com/DOSNetwork/core)

## Development Setup:
- [Install](https://golang.org/doc/install) Go (recommended version 1.10+) and setup golang workingspace, specifically by adding environment variable [GOPATH](https://golang.org/doc/code.html#GOPATH) into PATH.
- Install [dep](https://golang.github.io/dep/docs/daily-dep.html#key-takeaways) to manage package dependencies and versions.
  - Run `$ dep ensure` to update missing dependencies/packages.
  - [Visualize package dependencies](https://golang.github.io/dep/docs/daily-dep.html#visualizing-dependencies)
- Download:
  - `$ go get -d github.com/DOSNetwork/core/...` or
  - `$ git clone git@github.com:DOSNetwork/core.git`
- Build:
  - `$ make` or `$ make client` to build release version client.
  - `$ make devClient` to build develoment version client.
  - `$ make updateSubmodule` to fetch latest system contracts from [repo](https://github.com/DOSNetwork/eth-contracts), instead of making contract modifications locally.
  - `$ make gen` to generate binding files for system contracts.
- Dev tips:
  - `$ go fmt ./...` to reformat go source code.
  - `$ golint` to fix style mistakes conflicting with [effective go](https://golang.org/doc/effective_go.html). ([golint](https://github.com/golang/lint) tool for vim users.)
  - `$ make clean` to remove built binaries or unnecessary generated files.



## Running a Beta DOS node on a cloud server or VPS:
### Requirements
- **Cloud Server / VPS Recommendations**
  - [Vultr](https://www.vultr.com/?ref=7806004-4F) - Cloud Compute $5 monthly plan (1CPU, 1GB Memory, 25GB SSD, 1TB Bandwidth)
  - [AWS Lightsail](https://aws.amazon.com/lightsail/pricing/?opdp1=pricing) - $5 monthly plan (1CPU, 1GB Memory, 40GB SSD, 2TB Bandwidth)
  - [DigitalOcean](https://m.do.co/c/a912bdc08b78) - Droplet $5 monthly plan (1CPU, 25GB SSD, 1TB Bandwidth)
  - [Linode](https://www.linode.com/?r=35c0c22d412b3fc8bd98b4c7c6f5ac42ae3bc2e2) - $5 monthly plan (1CPU, 1GB Memory, 25GB SSD, 1TB Bandwidth)
  - .

- **Verified and recommended installation environment**
  - Ubuntu 16.04 x64 LTS or higher 
  - A static IPv4 address
  - Open port `7946, 8545, 8546 and 9501`
  - A ssh private key file for the VPS ( [How to set up ssh for Amazon Lightsail...](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/lightsail-how-to-set-up-ssh))

- **An Ethereum wallet with enough ether and DOS token**
  - An Ethereum keystore file ( [Create a keystore by geth ](https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts) or [Export a keystore through MyEtherwWllet Extension](https://bitcointalk.org/index.php?topic=3014688.0)])
  - Acquire testnet ether from rinkeby [faucet](https://faucet.rinkeby.io/)
  - Acquire 50,000 [testnet DOS token](https://rinkeby.etherscan.io/address/0x214e79c85744cd2ebbc64ddc0047131496871bee)
  - (Optional - acquire several [testnet DropBurn token](https://rinkeby.etherscan.io/address/0x9bfe8f5749d90eb4049ad94cc4de9b6c4c31f822))



### Prepare the environment
- Download [vps_docker.sh](https://raw.githubusercontent.com/DOSNetwork/core/Beta/vps_docker.sh) or [vps.sh](https://raw.githubusercontent.com/DOSNetwork/core/Beta/vps.sh),[dos.setting](https://raw.githubusercontent.com/DOSNetwork/core/Beta/dos.setting),[static-nodes.json](https://raw.githubusercontent.com/DOSNetwork/core/Beta/static-nodes.json),[config.json](https://raw.githubusercontent.com/DOSNetwork/core/Beta/config.json),[rinkeby.json](https://www.rinkeby.io/rinkeby.json) and [geth.service](https://raw.githubusercontent.com/DOSNetwork/core/Beta/geth.service)
- setup the following setting in the [dos.setting](https://raw.githubusercontent.com/DOSNetwork/core/Beta/dos.setting)
	- USER : VPS user name
	- VPSIP : VPS public IP
	- VPSKEY : VPS ssh private key location
	- KEYSTORE : Ethereum keystore file location
	- GETHPOOL : User can add his own geth full node here.More geth node could improve performance and stability of DOS.
	             Please note that ws is only for eveny subsciption.DOS need at least one valid geth http url.
		
			DOSIMAGE=dosnetwork/dosnode:beta
			USER=tester
			VPSIP=xxx.xxx.xxx.xxx
			KEYSTORE=xxx
			GETHPOOL="https://rinkeby.infura.io/projectid;ws://xxx.xxx.xxx.xxx:8546"

- change the following setting in the [geth.service](https://raw.githubusercontent.com/DOSNetwork/core/Beta/geth.service)
	- WorkingDirectory : 
	- User : VPS user name
	- datadir :
		
			[Service]
			WorkingDirectory=/home/tester
			User=tester
			ExecStart=/usr/bin/geth --datadir /home/tester/.rinkeby ...
### Install and run the client (Docker or Binary)
- Install and setup directorys for client
```sh
$ bash vps_docker.sh install
or
$ bash vps.sh install
```
- Run the client
```sh
$ bash vps_docker.sh run
or
$ bash vps.sh run
```
- Stop the client
```sh
$ bash vps_docker.sh stop
or
$ bash vps.sh stop
```
- Show the client status
```sh
$ bash vps_docker.sh clientInfo
or
$ bash vps.sh clientInfo
```

### Build the client locally and run the client on a VPS
- Follow the section [DEV setup and workflow] to build client
- 



## Status
- ☑️ Secret Sharing
- ☑️ Distributed Key Generation (Pedersen's DKG approach)
- ☑️ Paring Library and BLS Signature
- ☑️ Distributed Randomness Engine with VRF
- ☑️ Gossip & DHT Implementation
- ☑️ P2P NAT Support
- ☑️ Json / Xml / Html Parser
- ☑️ Dockerize and Deployment Script
- ☑️ Integration with Ethereum On-chain [System Contracts](https://github.com/DOSNetwork/eth-contracts)
- :white_large_square: P2P Network Performance Tuning
- :white_large_square: Network Status Scanner/Explorer
- :white_large_square: Staking & Delegation Contracts with a User-friendly Frontend
