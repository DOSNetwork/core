## <img align="left" width=30 src="media/logo-white.jpg"> Core libraries


### Components and file structure (TODO)
- Shamir's secret sharing 
- Pedersen's DKG approach
- Paring Library and BLS Signature
- DHT & Gossip implementation
- P2P NAT Support
- On-chain Verification Contracts
- Json / xml / html parser
- .


### DEV setup and workflow:
- Install Go [recommended version 1.10 and above](https://blog.golang.org/go1.10) and [setup golang workingspace](https://golang.org/doc/install), specifically environment variables like GOPATH, GOROOT, et al.
- [How to go-get from a private repo](https://blog.wilianto.com/go-get-from-private-repository.html)
- Install [dep](https://golang.github.io/dep/docs/daily-dep.html#key-takeaways) to manage package dependencies and versions.
  - [Visualize package dependencies](https://golang.github.io/dep/docs/daily-dep.html#visualizing-dependencies)
- Download: `$ go get -d github.com/DOSNetwork/core/...` or `git clone git@github.com:DOSNetwork/core.git`
- Build: `$ make build`
- Local Test workflow: `$ make build && cd testing && make deploy && make buildDockers && cd ../ && docker-compose up --scale dosnode=3`    
- before commit,you should do:
	- vi .git/config 
	- add 	
		
			[alias]
			ignore = update-index --assume-unchanged
			unignore = update-index --no-assume-unchanged
			ignored = !git ls-files -v | grep "^[[:lower:]]"
			
	- `git ignore onChain.json offChain.json testing/dosUser/ama.json`. If you want to change these file. exec `git unignore *.json`
	- Use `$ go fmt .`; or plugin before commit.
	- `$ make clean` before commit
- Note that Changing `github.com/DOSNetwork/eth-contracts` instead of modifying locally cloned submodules, and using `$ git submodule update --remote --merge` to checkout latest changes.

### Running a Beta DOS Client on a VPS (Ubuntu 16.04 LTS):
#### Requirements
1) A Ubuntu 16.04 LTS VPS
- A public IP
- Open 7946,8545,8546 and 9501 ports
- A ssh private key file for the VPS ( [How to set up ssh for Amazon Lightsail...](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/lightsail-how-to-set-up-ssh))
2) An Ethereum wallet with enough ether and DOS token.
- An Ethereum keystore file ( [Create a keystore by geth ](https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts) or [Export a keystore through MyEtherwWllet Extension](https://bitcointalk.org/index.php?topic=3014688.0)])
- [Get Rinkeby test Ether from faucet](https://faucet.rinkeby.io/)
- DOS Token 

#### 1) Prepare the environment
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
#### 2) Install and run the client (Docker or Binary)
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

#### 2) Build the client locally and run the client on a VPS
- Follow the section [DEV setup and workflow] to build client
- 

### Trouble shooting and Deploy
- Run `$ dep ensure -update` when it complains about missing dependencies/packages, and commits updated Gopkg.lock file.
- Dockerize ... (TODO)
