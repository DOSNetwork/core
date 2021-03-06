# <img align="center" width=40 src="media/logo-white.jpg"> DOS Client and Core Libraries
[![Go Report Card](https://goreportcard.com/badge/github.com/DOSNetwork/core)](https://goreportcard.com/report/github.com/DOSNetwork/core)
[![Maintainability](https://api.codeclimate.com/v1/badges/a2eb5767f8984835fb3b/maintainability)](https://codeclimate.com/github/DOSNetwork/core/maintainability)
[![GoDoc](https://godoc.org/github.com/DOSNetwork/core?status.svg)](https://godoc.org/github.com/DOSNetwork/core)


## Prerequisites:
##### Cloud Server / VPS Recommendations
- [AWS Lightsail](https://aws.amazon.com/lightsail/pricing) - Linux virtual server (1 cpu, 1GB memory, 40GB ssd, 2TB bandwidth)
- [Google Cloud Platform](https://cloud.google.com) - Compute Engine General purpose (N2)
- [Vultr](https://www.vultr.com/products/cloud-compute/) - Cloud Compute (1 cpu, 1GB memory, 25GB ssd, 1TB bandwidth)
- [Digital Ocean](https://www.digitalocean.com/products/droplets/) - Droplet (1 cpu, 25GB ssd, 1TB bandwidth)
- [Linode](https://www.linode.com/products/shared/) - Shared virtual Instances (1 cpu, 1GB memory, 25GB ssd, 1TB bandwidth)
- [Others]

##### Verified and recommended installation environment
- Ubuntu 18.04 x64 LTS or higher 
- An IPv4 address
  - Run `$ dig +short myip.opendns.com @resolver1.opendns.com`
  - Or get it from cloud server providers. Most vps / cloud server 
- With below ports open:
  - **udp** port `7946`
  - **tcp** port `7946`,`9501`
- It's recommended to generate ssh login key pairs and setup public key authentication instead of using password login for server security and funds safety:
  - Learn [how to](https://www.digitalocean.com/community/tutorials/how-to-set-up-ssh-keys-on-ubuntu-1604) setup SSH public key authentication on Ubuntu 18.04 and disable password logins.


##### Bonding enough DOS tokens to run a node
- Acquire at least 800,000 [DOS tokens](https://etherscan.io/address/0x70861e862e1ac0c96f853c8231826e469ead37b1), (and optional - acquire several [DropBurn token](https://etherscan.io/address/0x68423B3B0769c739D1fe4C398C3d91F0d646424f) to reduce (up to 30%) the amount of DOS tokens needed to start a node, distribution plan be out later).
- Replace your node ip address in [config.json](https://github.com/DOSNetwork/core/blob/master/config.json#L3)
- Node runners currently earn three types of incomes: (Self-bonded) staking rewards, (other delegators') staking reward shares, oracle request processing fees.


##### Register and setup Infura api key
- Register and get [Infura api key](https://ethereumico.io/knowledge-base/infura-api-key-guide/)
- Replace your infura api key in [config.json](https://github.com/DOSNetwork/core/blob/master/config.json#L11)



## Run with docker image
- https://medium.com/dos-network/instructions-of-launching-a-node-in-dos-network-932e73a91a75


## Building binary from source
- [Install](https://golang.org/doc/install) **go1.11 or higher** and setup golang workingspace like below:
    ```sh
    $ sudo apt-get install golang 
    $ sudo apt-get install build-essential
    ```
    
- Open `~/.bashrc` and set `$GOPATH` and `$PATH` environmental variables:
    ```sh
    $ vim ~/.bashrc
      export GOPATH=$HOME/go
      export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
    $ source ~/.bashrc
    ```

- Download source code:
    ```sh
    $ mkdir -p $GOPATH/src/github.com/DOSNetwork
    $ cd $GOPATH/src/github.com/DOSNetwork && git clone https://github.com/DOSNetwork/core.git
    $ cd core
    $ git checkout v1.0.4-m
    ```

- Build:
  - `$ GO111MODULE=on go mod download` - to download building dependencies
  - Dev dependencies: solc-v0.5.17, abigen-v1.9.18-stable
  - `$ make` - to build release version client

- Run:
    ```sh
    $ ./dos.sh start
    ```

- Dev tips:
  - `$ go fmt ./...` to reformat go source code.
  - `$ golint` to fix style mistakes conflicting with [effective go](https://golang.org/doc/effective_go.html). ([golint](https://github.com/golang/lint) tool for vim users.)
  - `$ make devClient` to build develoment version client.
  - `$ make updateSubmodule` to fetch latest system contracts from [repo](https://github.com/DOSNetwork/eth-contracts), instead of making contract modifications locally.
  - `$ make gen` to generate binding files for system contracts.
  - `$ make clean` to remove built binaries or unnecessary generated files.
  - Use [solc-select](https://github.com/crytic/solc-select) tool to select / pin to specific solidity compiler version.
