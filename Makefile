.DEFAULT_GOAL := build

DOSPROXY_PATH := onchain/eth/contracts/DOSProxy.sol
DOSPROXY_GOPATH := onchain/dosproxy
DOSPAYMENT_GOPATH := onchain/dospayment
DOSBRIDGE_GOPATH := onchain/dosbridge
COMMITREVEAL_GOPATH := onchain/commitreveal
TEST_CONTRACTS_GOPATH := testing/dosUser/contract
GENERATED_FILES := $(shell find $(DOSPROXY_GOPATH) $(DOSBRIDGE_GOPATH) $(DOSPAYMENT_GOPATH) $(COMMITREVEAL_GOPATH) $(TEST_CONTRACTS_GOPATH) -name '*.go')
ETH_CONTRACTS := onchain/eth/contracts
BOOT_CREDENTIAL := testAccounts/bootCredential
UNAME_S := $(shell uname -s)


.PHONY: build
build: vendor client


.PHONY: vendor
vendor:
	@ dep ensure -vendor-only


.PHONY: devClient
# Build a development version client node
devClient: gen
	go build -o client.dev


.PHONY: client
# Build a prod/release version client node
client:
	go build -o dosclient

.PHONY: client-docker
client-docker:
	docker build -t dosnode -f Dockerfile .


.PHONY: install
install: vendor client
	go install


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


.PHONY: clean
clean:
	@ rm -f dosclient*
	@ # rm -f $(GENERATED_FILES)
