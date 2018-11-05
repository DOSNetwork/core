# Usage:
# $ make == $ make build
# $ make install

.PHONY: dep build client install clean gen deploy

.DEFAULT_GOAL := build
ETH_CONTRACTS := onchain/eth/contracts
GENERATED_FILES := $(filter-out $(shell find $(ETH_CONTRACTS) -name '*_test.go'), $(shell find $(ETH_CONTRACTS) -name '*.go'))
#For testing
DOCKER_IMAGES := dockerImages
TEST_ACCOUNTS := testAccounts

build: dep client


dep:
	dep ensure -vendor-only


client:
	go build -o client

install: dep client
	go install

genDockers:
	mkdir -p $(DOCKER_IMAGES)
	cp -r $(TEST_ACCOUNTS) $(DOCKER_IMAGES)/
	cp -r credential $(DOCKER_IMAGES)/
	cp config.json $(DOCKER_IMAGES)/
	go build -ldflags "-linkmode external -extldflags -static" -a -o clientNode main.go
	mv clientNode $(DOCKER_IMAGES)/
	cd testing/bootStrapNode/;go build -ldflags "-linkmode external -extldflags -static" -a -o bootstrapNode boot_strap_node.go
	mv testing/bootStrapNode/bootstrapNode $(DOCKER_IMAGES)/
	cd testing/dosUser/;go build -ldflags "-linkmode external -extldflags -static" -a -o userNode dos_user.go
	mv testing/dosUser/userNode $(DOCKER_IMAGES)/
	cp Dockerfile $(DOCKER_IMAGES)/Dockerfile.dosnode
	cp testing/bootStrapNode/Dockerfile $(DOCKER_IMAGES)/Dockerfile.bootstrap
	cp testing/dosUser/Dockerfile $(DOCKER_IMAGES)/Dockerfile.usernode

buildDockers:genDockers
	cd $(DOCKER_IMAGES);docker build -t bootstrap -f Dockerfile.bootstrap  .
	cd $(DOCKER_IMAGES);docker build -t dosnode -f Dockerfile.dosnode .
	cd $(DOCKER_IMAGES);docker build -t usernode -f Dockerfile.usernode  .

#Only used for deploy a new contracts for testing
deploy:
	abigen -sol $(ETH_CONTRACTS)/DOSProxy.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSProxy.go
	abigen -sol $(ETH_CONTRACTS)/DOSAddressBridge.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSAddressBridge.go
	abigen -sol $(ETH_CONTRACTS)/DOSOnChainSDK.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSOnChainSDK.go
	cp $(ETH_CONTRACTS)/DOSOnChainSDK.sol $(ETH_CONTRACTS)/userContract/
	cp $(ETH_CONTRACTS)/Ownable.sol $(ETH_CONTRACTS)/userContract/
	mkdir $(ETH_CONTRACTS)/userContract/lib/
	cp $(ETH_CONTRACTS)/lib/utils.sol $(ETH_CONTRACTS)/userContract/lib/
	abigen -sol $(ETH_CONTRACTS)/userContract/AskMeAnything.sol --pkg userContract --out $(ETH_CONTRACTS)/userContract/AskMeAnything.go
	rm $(ETH_CONTRACTS)/userContract/DOSOnChainSDK.sol
	rm $(ETH_CONTRACTS)/userContract/Ownable.sol
	go run onchain/eth/deploy.go -contractPath $(ETH_CONTRACTS) -step ProxyAndBridge
	abigen -sol $(ETH_CONTRACTS)/DOSOnChainSDK.sol --pkg dosproxy --out $(ETH_CONTRACTS)/DOSOnChainSDK.go
	cp $(ETH_CONTRACTS)/DOSOnChainSDK.sol $(ETH_CONTRACTS)/userContract/
	cp $(ETH_CONTRACTS)/Ownable.sol $(ETH_CONTRACTS)/userContract/
	abigen -sol $(ETH_CONTRACTS)/userContract/AskMeAnything.sol --pkg userContract --out $(ETH_CONTRACTS)/userContract/AskMeAnything.go
	rm $(ETH_CONTRACTS)/userContract/DOSOnChainSDK.sol
	rm $(ETH_CONTRACTS)/userContract/Ownable.sol
	rm -r $(ETH_CONTRACTS)/userContract/lib/
	go run onchain/eth/deploy.go -contractPath $(ETH_CONTRACTS) -step SDKAndAMA
	go run onchain/eth/deploy.go -contractPath $(ETH_CONTRACTS) -step SetProxyAddress


clean:
	rm -f client
	rm -f $(GENERATED_FILES)
