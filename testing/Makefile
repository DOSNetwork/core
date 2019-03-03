# Usage:
# $ make == $ make deploy

.PHONY: buildDockers deploy

.DEFAULT_GOAL := deploy
ETH_CONTRACTS := ../onchain/eth/contracts
REMOTE_USER_CONTRACT := ../onchain/eth/contracts/examples
USER_CONTRACTS := dosUser/contract
BOOT_CREDENTIAL := ../testAccounts/bootCredential
AMA_CONFIGPATH := dosUser/ama.json

gen:
	cd .. && make gen
	mkdir -p $(USER_CONTRACTS)
	cp $(ETH_CONTRACTS)/DOSOnChainSDK.sol $(USER_CONTRACTS)/
	cp $(ETH_CONTRACTS)/Ownable.sol $(USER_CONTRACTS)/
	mkdir -p $(USER_CONTRACTS)/lib/
	cp $(ETH_CONTRACTS)/lib/utils.sol $(USER_CONTRACTS)/lib/
	cp $(REMOTE_USER_CONTRACT)/AskMeAnything.sol $(USER_CONTRACTS)/
	abigen -sol $(USER_CONTRACTS)/AskMeAnything.sol --pkg dosUser --out $(USER_CONTRACTS)/AskMeAnything.go
	rm $(USER_CONTRACTS)/*.sol
	rm -r $(USER_CONTRACTS)/lib

clean:
	cd .. && make clean

deploy: gen
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -contract DOSProxy
	cp $(ETH_CONTRACTS)/DOSOnChainSDK.sol $(USER_CONTRACTS)/
	cp $(ETH_CONTRACTS)/Ownable.sol $(USER_CONTRACTS)/
	mkdir -p $(USER_CONTRACTS)/lib/
	cp $(ETH_CONTRACTS)/lib/utils.sol $(USER_CONTRACTS)/lib/
	cp $(REMOTE_USER_CONTRACT)/AskMeAnything.sol $(USER_CONTRACTS)/
	abigen -sol $(USER_CONTRACTS)/AskMeAnything.sol --pkg dosUser --out $(USER_CONTRACTS)/AskMeAnything.go
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -contract AMA -configPath $(AMA_CONFIGPATH)
	rm $(USER_CONTRACTS)/*.sol
	rm -r $(USER_CONTRACTS)/lib

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

docker-peerTalk:
	cp peerTalk/Dockerfile ../Dockerfile.peertalk
	cd .. && docker build -t peertalk -f Dockerfile.peertalk .
	rm ../Dockerfile.peertalk

docker-all: docker-bootnode docker-dosnode docker-usernode docker-peernode docker-peerTalk
	docker image prune -f

update-dosenv:
	-docker rmi dosenv
	-docker rmi dosnetwork/dosenv
	cd .. && docker build -t dosenv -f Dockerfile.env .
	docker tag dosenv dosnetwork/dosenv:latest
	docker push dosnetwork/dosenv:latest
