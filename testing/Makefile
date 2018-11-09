# Usage:
# $ make == $ make deploy

.PHONY: buildDockers deploy

.DEFAULT_GOAL := deploy
ETH_CONTRACTS := ../onchain/eth/contracts
USER_CONTRACTS := dosUser/contract
BOOT_CREDENTIAL := ../testAccounts/bootCredential
AMA_CONFIGPATH := dosUser/ama.json


deploy:
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -step ProxyAndBridge
	go run contracts_deploy/deploy.go -credentialPath $(BOOT_CREDENTIAL) -contractPath $(ETH_CONTRACTS) -step SetProxyAddress
	cp $(ETH_CONTRACTS)/DOSOnChainSDK.sol $(USER_CONTRACTS)/
	cp $(ETH_CONTRACTS)/Ownable.sol $(USER_CONTRACTS)/
	mkdir -p $(USER_CONTRACTS)/lib/
	cp $(ETH_CONTRACTS)/lib/utils.sol $(USER_CONTRACTS)/lib/
	abigen -sol $(USER_CONTRACTS)/AskMeAnything.sol --pkg dosUser --out $(USER_CONTRACTS)/AskMeAnything.go
	rm $(USER_CONTRACTS)/DOSOnChainSDK.sol
	rm $(USER_CONTRACTS)/Ownable.sol
	rm -r $(USER_CONTRACTS)/lib/
	go run contracts_deploy/depoy_ama.go -AMAPath $(AMA_CONFIGPATH) -credentialPath $(BOOT_CREDENTIAL)


buildDockers:
	cp Dockerfile ../Dockerfile.base
	cp ../Dockerfile ../Dockerfile.dosnode
	cp bootStrapNode/Dockerfile  ../Dockerfile.bootstrap
	cp dosUser/Dockerfile ../Dockerfile.usernode
	cd ../;docker build -t dosbase -f Dockerfile.base .
	cd ../;docker build -t dosnode -f Dockerfile.dosnode .
	cd ../;docker build -t bootstrap -f Dockerfile.bootstrap  .
	cd ../;docker build -t usernode -f Dockerfile.usernode  .
	docker image prune -f
	docker rmi dosbase
	cd ../;rm -f Dockerfile.bootstrap Dockerfile.dosnode Dockerfile.usernode Dockerfile.base