# Usage:
# $ make == $ make build
# $ make install

.PHONY: dep build client install clean gen

.DEFAULT_GOAL := build
ETH_CONTRACTS := onchain/eth/contracts
GENERATED_FILES := $(filter-out $(shell find $(ETH_CONTRACTS) -name '*_test.go'), $(shell find $(ETH_CONTRACTS) -name '*.go'))


build: dep client


dep:
	dep ensure -vendor-only


client: gen
	go build -o client


gen:
	abigen -sol $(ETH_CONTRACTS)/DOSAddressBridge.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSAddressBridge.go
	abigen -sol $(ETH_CONTRACTS)/DOSOnChainSDK.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSOnChainSDK.go
	abigen -sol $(ETH_CONTRACTS)/DOSProxy.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSProxy.go


install: dep client
	go install


clean:
	rm -f client
	rm -f $(GENERATED_FILES)
