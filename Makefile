# Usage:
# $ make == $ make build
# $ make install

.PHONY: dep build client install clean gen

.DEFAULT_GOAL := build
DOSPROXY_PATH := onchain/eth/contracts/DOSProxy.sol
DOSPROXY_GOPATH := onchain/dosproxy
DOSBRIDGE_GOPATH := onchain/dosbridge
COMMITREVEAL_GOPATH := onchain/commitreveal
TEST_CONTRACTS_GOPATH := testing/dosUser/contract
GENERATED_FILES := $(shell find $(DOSPROXY_GOPATH) $(DOSBRIDGE_GOPATH) $(TEST_CONTRACTS_GOPATH) -name '*.go')
ETH_CONTRACTS := onchain/eth/contracts
BOOT_CREDENTIAL := testAccounts/bootCredential

build: dep client

dep:
	dep ensure -vendor-only

client: gen
ifeq ($(UNAME_S),Linux)
        go build -o client
endif
ifeq ($(UNAME_S),Darwin)
        xgo --targets=linux/amd64 -out client .
endif

client-docker:
	docker build -t dosnode -f Dockerfile .

install: dep client
	go install

updateSubmodule:
	test -f $(DOSPROXY_PATH) || git submodule update --init --recursive
	git submodule update --remote

gen: updateSubmodule
	abigen -sol $(ETH_CONTRACTS)/DOSAddressBridge.sol --pkg dosbridge --out $(DOSBRIDGE_GOPATH)/DOSAddressBridge.go
	abigen -sol $(ETH_CONTRACTS)/DOSProxy.sol --pkg dosproxy --out $(DOSPROXY_GOPATH)/DOSProxy.go
	abigen -sol $(ETH_CONTRACTS)/CommitReveal.sol --pkg commitreveal --out $(COMMITREVEAL_GOPATH)/CommitReveal.go

clean:
	rm -f client
	rm -f $(GENERATED_FILES)
