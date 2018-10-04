## <img align="left" width=30 src="media/logo-white.jpg"> Core libraries


### Components and file structure (TODO)
- Shamir's secret sharing 
- Feldman/Pedersen's DKG approach
- Paring Library and BLS Signature
- DHT & Gossip implementation
- P2P NAT Support
- On-chain Verification Contracts
- .



### DEV setup and workflow:
- Install Go [recommended version 1.10 and above](https://blog.golang.org/go1.10) and [setup golang workingspace](https://golang.org/doc/install), specifically environment variables like GOPATH, GOROOT, et al.
- [How to go-get from a private repo](https://blog.wilianto.com/go-get-from-private-repository.html)
- Install [dep](https://golang.github.io/dep/docs/daily-dep.html#key-takeaways) to manage package dependencies and versions.
  - [Visualize package dependencies](https://golang.github.io/dep/docs/daily-dep.html#visualizing-dependencies)
- Use $ go fmt .; or plugin before commit.
- Download: $ go get -d github.com/DOSNetwork/core/...
- Build: $ make
- Install: $ make install; (makefile todo) 
- ... Hack ...




### Deploy 
- Dockerize ... (TODO)

