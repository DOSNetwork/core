.DEFAULT_GOAL := build

DOSPROXY_PATH := onchain/eth/contracts/DOSProxy.sol
DOSPROXY_GOPATH := onchain/dosproxy
DOSPAYMENT_GOPATH := onchain/dospayment
DOSBRIDGE_GOPATH := onchain/dosbridge
COMMITREVEAL_GOPATH := onchain/commitreveal
DOSSTAKING_GOPATH := onchain/dosstaking
GENERATED_FILES := $(shell find $(DOSPROXY_GOPATH) $(DOSBRIDGE_GOPATH) $(DOSPAYMENT_GOPATH) $(DOSSTAKING_GOPATH) $(COMMITREVEAL_GOPATH) -name '*.go')
ETH_CONTRACTS := onchain/eth/contracts
BOOT_CREDENTIAL := testAccounts/bootCredential
UNAME_S := $(shell uname -s)

VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS_STATIC=-ldflags "-linkmode external -extldflags -static -w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

.PHONY: build
build: vendor client


.PHONY: vendor
vendor:
	@ dep ensure -vendor-only


.PHONY: devClient
# Build a development version client node
devClient: gen
	go build -a -o client.dev


.PHONY: client
# Build a prod/release version client node
client:
	go build ${LDFLAGS} -o dosclient

client-static:
	go build ${LDFLAGS_STATIC} -o dosclient

.PHONY: client-docker
client-docker:
	docker build -t dosnode -f Dockerfile .


.PHONY: install
install: vendor client
	go install ${LDFLAGS}


.PHONY: updateSubmodule
updateSubmodule:
	test -f $(DOSPROXY_PATH) || git submodule update --init --recursive
	git submodule update --remote


.PHONY: gen
gen: updateSubmodule
	solc --optimize --overwrite --abi  --bin $(ETH_CONTRACTS)/DOSAddressBridge.sol -o $(DOSBRIDGE_GOPATH)
	abigen --abi="$(DOSBRIDGE_GOPATH)/DOSAddressBridge.abi" --bin="$(DOSBRIDGE_GOPATH)/DOSAddressBridge.bin" --pkg dosbridge --out $(DOSBRIDGE_GOPATH)/DOSAddressBridge.go
	solc --optimize --overwrite --abi  --bin $(ETH_CONTRACTS)/DOSProxy.sol -o $(DOSPROXY_GOPATH)
	abigen --abi="$(DOSPROXY_GOPATH)/DOSProxy.abi" --bin="$(DOSPROXY_GOPATH)/DOSProxy.bin" --pkg dosproxy --out $(DOSPROXY_GOPATH)/DOSProxy.go
	solc --optimize --overwrite --abi  --bin $(ETH_CONTRACTS)/DOSPayment.sol -o $(DOSPAYMENT_GOPATH)
	abigen --abi="$(DOSPAYMENT_GOPATH)/DOSPayment.abi" --bin="$(DOSPAYMENT_GOPATH)/DOSPayment.bin" --pkg dospayment --out $(DOSPAYMENT_GOPATH)/DOSPayment.go
	solc --optimize --overwrite --abi  --bin $(ETH_CONTRACTS)/CommitReveal.sol -o $(COMMITREVEAL_GOPATH)
	abigen --abi="$(COMMITREVEAL_GOPATH)/CommitReveal.abi" --bin="$(COMMITREVEAL_GOPATH)/CommitReveal.bin" --pkg commitreveal --out $(COMMITREVEAL_GOPATH)/CommitReveal.go
	solc --optimize --overwrite --abi  --bin $(ETH_CONTRACTS)/Staking.sol -o $(DOSSTAKING_GOPATH)
	abigen --abi="$(DOSSTAKING_GOPATH)/Staking.abi" --bin="$(DOSSTAKING_GOPATH)/Staking.bin" --pkg staking --out $(DOSSTAKING_GOPATH)/Staking.go


.PHONY: clean
clean:
	@ rm -f dosclient*
	@ # rm -f $(GENERATED_FILES)
