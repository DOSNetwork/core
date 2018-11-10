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


### Trouble shooting and Deploy
- Run `$ dep ensure -update` when it complains about missing dependencies/packages, and commits updated Gopkg.lock file.
- Dockerize ... (TODO)
