# Usage:
# $ make == $ make deploy

.PHONY: buildDockers deploy

.DEFAULT_GOAL := deploy
ETH_CONTRACTS := ../onchain/eth/contracts
REMOTE_USER_CONTRACT := ../onchain/eth/contracts/examples
USER_CONTRACTS := dosUser/contract
BOOT_CREDENTIAL := ../testAccounts/bootCredential/fundKey
AMA_CONFIGPATH := dosUser/ama.json
DOSPROXY_GOPATH := onchain/dosproxy
DOSPAYMENT_GOPATH := onchain/dospayment
DOSBRIDGE_GOPATH := onchain/dosbridge
COMMITREVEAL_GOPATH := onchain/commitreveal
UNAME_S := $(shell uname)

clean:
	cd .. && make clean

deploy-bridge:
	rm ../$(DOSBRIDGE_GOPATH)/*
	solc --optimize --abi --bin $(ETH_CONTRACTS)/DOSAddressBridge.sol -o ../$(DOSBRIDGE_GOPATH)
	abigen --abi="../$(DOSBRIDGE_GOPATH)/DOSAddressBridge.abi" --bin="../$(DOSBRIDGE_GOPATH)/DOSAddressBridge.bin" --pkg dosbridge --out ../$(DOSBRIDGE_GOPATH)/DOSAddressBridge.go
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployBridge
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target updateBridgeAddrToSDK

deploy-payment:
	rm ../$(DOSPAYMENT_GOPATH)/*
	solc --optimize --abi --bin $(ETH_CONTRACTS)/DOSPayment.sol -o ../$(DOSPAYMENT_GOPATH)
	abigen --abi="../$(DOSPAYMENT_GOPATH)/DOSPayment.abi" --bin="../$(DOSPAYMENT_GOPATH)/DOSPayment.bin" --pkg dospayment --out ../$(DOSPAYMENT_GOPATH)/DOSPayment.go
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployPayment

deploy-proxy:
	rm ../$(DOSPROXY_GOPATH)/*
	solc --optimize --abi --bin $(ETH_CONTRACTS)/DOSProxy.sol -o ../$(DOSPROXY_GOPATH)
	abigen --abi="../$(DOSPROXY_GOPATH)/DOSProxy.abi" --bin="../$(DOSPROXY_GOPATH)/DOSProxy.bin" --pkg dosproxy --out ../$(DOSPROXY_GOPATH)/DOSProxy.go 
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployProxy

deploy-cr:
	rm ../$(COMMITREVEAL_GOPATH)/*
	solc --optimize --abi --bin $(ETH_CONTRACTS)/CommitReveal.sol -o ../$(COMMITREVEAL_GOPATH)
	abigen --abi="../$(COMMITREVEAL_GOPATH)/CommitReveal.abi" --bin="../$(COMMITREVEAL_GOPATH)/CommitReveal.bin" --pkg commitreveal --out ../$(COMMITREVEAL_GOPATH)/CommitReveal.go
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployCommitReveal

setup-dos:
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target setUpDOSNetwork

deploy-dos:
	$(MAKE) deploy-bridge
	$(MAKE) deploy-payment
	$(MAKE) deploy-proxy
	$(MAKE) deploy-cr
	$(MAKE) setup-dos

build-client:
ifeq ($(UNAME_S),Linux)
	go build -o client ../
	cp ../config.json .
endif
ifeq ($(UNAME_S),Darwin)
	xgo --targets=linux/amd64 -out client ../
	mv client-linux-amd64 client
	cp ../config.json .
endif

build-client-docker:
	cd .. && make client-docker
	docker tag dosnode dosnetwork/dosnode:latest
	docker push dosnetwork/dosnode:latest

dos-all:
	$(MAKE) deploy-dos
	$(MAKE) build-client
	$(MAKE) build-client-docker

gen-dosuser:
	mkdir -p $(USER_CONTRACTS)
	cp $(ETH_CONTRACTS)/DOSOnChainSDK.sol $(USER_CONTRACTS)/
	cp $(ETH_CONTRACTS)/Ownable.sol $(USER_CONTRACTS)/
	mkdir -p $(USER_CONTRACTS)/lib/
	cp $(ETH_CONTRACTS)/lib/utils.sol $(USER_CONTRACTS)/lib/
	cp $(REMOTE_USER_CONTRACT)/AskMeAnything.sol $(USER_CONTRACTS)
	abigen -sol $(USER_CONTRACTS)/AskMeAnything.sol --pkg dosUser --out $(USER_CONTRACTS)/AskMeAnything.go
	rm $(USER_CONTRACTS)/*.sol
	rm -r $(USER_CONTRACTS)/lib

deploy-dosuser:
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployAMA -configPath $(AMA_CONFIGPATH)

build-dosuser:
	cp ../config.json dosUser/
	go build -o dosUser/dos_user dosUser/dos_user.go

dosuser-all:
	$(MAKE) gen-dosuser
	$(MAKE) deploy-dosuser
	$(MAKE) build-dosuser

update-dosenv:
	-docker rmi dosenv
	-docker rmi dosnetwork/dosenv
	cd .. && docker build -t dosenv -f Dockerfile.env .
	docker tag dosenv dosnetwork/dosenv:latest
	docker push dosnetwork/dosenv:latest

docker-bootnode:
	cp bootStrapNode/Dockerfile  ../Dockerfile.bootstrap	
	cd .. && docker build -t bootstrap -f Dockerfile.bootstrap .
	rm ../Dockerfile.bootstrap

docker-usernode:
	cp dosUser/Dockerfile ../Dockerfile.usernode
	cd .. && docker build -t usernode -f Dockerfile.usernode .
	rm ../Dockerfile.usernode

docker-peernode:
	cp peerNode/Dockerfile ../Dockerfile.peernode
	cd .. && docker build -t peernode -f Dockerfile.peernode .
	rm ../Dockerfile.peernode

docker-banker:
	cp banker/Dockerfile ../Dockerfile.banker
	cd .. && docker build -t banker -f Dockerfile.banker .
	rm ../Dockerfile.banker

docker-all: docker-bootnode docker-dosnode docker-usernode docker-banker

docker-aws: docker-bootnode docker-dosnode docker-usernode docker-banker
	docker image prune -f
	docker tag usernode dosnetwork/usernode:latest
	docker tag banker dosnetwork/banker:latest
	docker tag bootstrap dosnetwork/bootstrap:latest
	docker tag dosnode dosnetwork/dosnode:latest
	docker push dosnetwork/usernode:latest
	docker push dosnetwork/banker:latest
	docker push dosnetwork/bootstrap:latest
	docker push dosnetwork/dosnode:latest
