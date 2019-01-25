## <img align="left" width=30 src="media/logo-white.jpg"> P2P Client of DOS Network Core Protocols
[![Go Report Card](https://goreportcard.com/badge/github.com/DOSNetwork/core)](https://goreportcard.com/report/github.com/DOSNetwork/core)
[![godoc](https://godoc.org/github.com/DOSNetwork/core?status.svg)](https://godoc.org/github.com/DOSNetwork/core)


### Download and build from source:
- Install Go [recommended version 1.10 and above](https://blog.golang.org/go1.10) and [setup golang workingspace](https://golang.org/doc/install), specifically environment variables like GOPATH.
- Use [dep](https://golang.github.io/dep/docs/daily-dep.html#key-takeaways) to manage package dependencies and generate [package dependency graphs](//golang.github.io/dep/docs/daily-dep.html#visualizing-dependencies).
- Download: `$ go get -d github.com/DOSNetwork/core/...`
- Build: `$ make build`
- Local Test workflow: `$ make build && cd testing && make deploy && make buildDockers && cd ../ && docker-compose up --scale dosnode=3`    


### Development & contribution notice:	
- `$ go fmt .` or use IDE plugins format source code before commit.
- `$ make clean` to remove unnecessary changes before commit.
- If it's necessary to update on-chain smart contracts, update `github.com/DOSNetwork/eth-contracts` instead of modifying locally cloned submodules, `$ git submodule update --remote --merge` to checkout latest changes.
- Golang bindings for the on-chain contract are generated using tool `abigen`.
- For common errors like "cannot find package "xxx" in any of: ...": Try: `$ unset GOROOT` and/or `$ dep ensure -update`.
- Create a new branch to upload pull requests for contributions.


### Core functionalities and components:
- Shamir secret sharing 
- Pedersen DKG algorithm
- Paring library and threshold bls signature
- Layer-2 P2P network based on Kademlia DHT with NAT support
- On-chain verification contracts and Ethereum chain adaptor
- Jsonpath and xmlpath parser
- ...
