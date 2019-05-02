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

### Running a DOS Client on a VPS (Ubuntu 16.04 LTS):
#### Requirements
- A Ubuntu 16.04 LTS VPS with a public IP
- Opening 7946,8545,8546 and 9501 port in the VPS
- A ssh private key file for the VPS
- A ETH keystore

#### 1) Using a docker
- Download vps_docker.sh and setup the following setting in vps_docker.sh
	- USER : VPS user name
	- VPSIP : VPS public IP
	- VPSKEY : ssh private key for VPS  ( [Ex:Amazon Lightsail...](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/lightsail-how-to-set-up-ssh))
	- ETHKEY : Ethereum keystore file ( [Create a geth keystore by geth ](https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts))
- Download config.json and setup the following setting in config.sh
	- RemoteNodeAddressPool : geth RPC address
		
			"RemoteNodeAddressPool": [
			    "https://rinkeby.infura.io/projectid",
			    "ws://xxx.xxx.xxx.xxx:8546",
			    "ws://xxx.xxx.xxx.xxx:8546"
			]
			
- Download geth.service
- Put vps_docker.sh,config.json and geth.service under the same directory
- Install Geth Light Node : `$ bash vps_docker.sh install_lightnode`
- Install docker : `$ bash vps_docker.sh install_docker`
- Install client : `$ bash vps_docker.sh install_client`
- Run the client : `$ bash vps_docker.sh run_client`
- Show the client status : `$ bash vps_docker.sh show_status`
- Stop the client : `$ bash vps_docker.sh run_client`

#### 2) Build a binary 


### Trouble shooting and Deploy
- Run `$ dep ensure -update` when it complains about missing dependencies/packages, and commits updated Gopkg.lock file.
- Dockerize ... (TODO)
