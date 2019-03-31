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
DOSBRIDGE_GOPATH := onchain/dosbridge
DOSCOMMITREVEAL_GOPATH := onchain/doscommitreveal
clean:
	cd .. && make clean
deployAll:
	$(MAKE) deployBridge
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployProxy
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployCommitReveal
	$(MAKE) deployAMA

deployBridge:
	abigen -sol $(ETH_CONTRACTS)/DOSAddressBridge.sol --pkg dosbridge --out ../$(DOSBRIDGE_GOPATH)/DOSAddressBridge.go
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployBridge
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target updateBridgeAddrToOtherContract
	abigen -sol $(ETH_CONTRACTS)/DOSAddressBridge.sol --pkg dosbridge --out ../$(DOSBRIDGE_GOPATH)/DOSAddressBridge.go
	abigen -sol $(ETH_CONTRACTS)/DOSProxy.sol --pkg dosproxy --out ../$(DOSPROXY_GOPATH)/DOSProxy.go
	abigen -sol $(ETH_CONTRACTS)/DOSCommitReveal.sol --pkg doscommitreveal --out ../$(DOSCOMMITREVEAL_GOPATH)/DOSCommitReveal.go
genProxy:
	abigen -sol $(ETH_CONTRACTS)/DOSProxy.sol --pkg dosproxy --out ../$(DOSPROXY_GOPATH)/DOSProxy.go
deployPoxy:genProxy
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployProxy

genCommitReveal:
	abigen -sol $(ETH_CONTRACTS)/DOSCommitReveal.sol --pkg doscommitreveal --out ../$(DOSCOMMITREVEAL_GOPATH)/DOSCommitReveal.go
deployCommitReveal:genCommitReveal
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployCommitReveal

genAMA:
	mkdir -p $(USER_CONTRACTS)
	cp $(ETH_CONTRACTS)/DOSOnChainSDK.sol $(USER_CONTRACTS)/
	cp $(ETH_CONTRACTS)/Ownable.sol $(USER_CONTRACTS)/
	mkdir -p $(USER_CONTRACTS)/lib/
	cp $(ETH_CONTRACTS)/lib/utils.sol $(USER_CONTRACTS)/lib/
	cp $(REMOTE_USER_CONTRACT)/AskMeAnything.sol $(USER_CONTRACTS)/
	abigen -sol $(USER_CONTRACTS)/AskMeAnything.sol --pkg dosUser --out $(USER_CONTRACTS)/AskMeAnything.go
	rm $(USER_CONTRACTS)/*.sol
	rm -r $(USER_CONTRACTS)/lib
deployAMA:genAMA
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -target deployAMA -configPath $(AMA_CONFIGPATH)

buildAll:deployAll
	xgo --targets=linux/amd64 -out .testEnv/dosclient ../
	cp ../config.json testEnv/
	go build -o testEnv/dos_user dosUser/dos_user.go
	cp dosUser/ama.json testEnv/

docker-bootnode:
	cp bootStrapNode/Dockerfile  ../Dockerfile.bootstrap	
	cd .. && docker build -t bootstrap -f Dockerfile.bootstrap .
	rm ../Dockerfile.bootstrap

docker-dosnode:
	cd .. && make client-docker

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

update-dosenv:
	-docker rmi dosenv
	-docker rmi dosnetwork/dosenv
	cd .. && docker build -t dosenv -f Dockerfile.env .
	docker tag dosenv dosnetwork/dosenv:latest
	docker push dosnetwork/dosenv:latest
