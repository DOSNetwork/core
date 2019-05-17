# <img align="left" width=30 src="media/logo-white.jpg"> DOS Client and Core Libraries


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
- Dev:
  - `$ go fmt .` for indentation and basic coding styles.
  - `$ make clean` to remove binaries or unnecessary generated files.
  - `$ make build && cd testing && make deploy && make buildDockers && cd ../ && docker-compose up --scale dosnode=3` to do local tests.



## Running a Beta DOS Client on a VPS (Ubuntu 16.04 LTS):
### Requirements
1) A Ubuntu 16.04 LTS VPS
- A public IP
- Open 7946,8545,8546 and 9501 ports
- A ssh private key file for the VPS ( [How to set up ssh for Amazon Lightsail...](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/lightsail-how-to-set-up-ssh))
2) An Ethereum wallet with enough ether and DOS token.
- An Ethereum keystore file ( [Create a keystore by geth ](https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts) or [Export a keystore through MyEtherwWllet Extension](https://bitcointalk.org/index.php?topic=3014688.0)])
- [Get Rinkeby test Ether from faucet](https://faucet.rinkeby.io/)
- DOS Token 

### 1) Prepare the environment
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
### 2) Install and run the client (Docker or Binary)
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

### 2) Build the client locally and run the client on a VPS
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
