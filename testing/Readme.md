1)Check if dosenv is existing.
  Only rm dosenv when project has new dependency
$docker image ls | grep dosenv
$docker image rm dosenv

2)Get code from github

DOSNetwork$
DOSNetwork$git clone https://github.com/DOSNetwork/core.git
DOSNetwork$git checkout grouping_xxx
DOSNetwork$cd core

3)Set up the test accounts
  Add passPhrase and move testAccounts under core

DOSNetwork/core$vi credential/passPhrase
DOSNetwork/core$mv ../testAccounts .

4)Build client node
  Use maek build to update dependency ,generate DOSAddressBridge.go and DOSProxy.go and build  client node
DOSNetwork/core$make build

5)Build docker testing images
  Set ENVIRONMENT variable CHAINNODE to indicate which eth full node we want to access
  fish shell use set and export to set ENVIRONMENT variable
#For fish shell
DOSNetwork/core$cd testing
DOSNetwork/core/testing$set CHAINNODE rinkebyInfura
DOSNetwork/core/testing$export CHAINNODE

#For bash shell
DOSNetwork/core/testing$export CHAINNODE="rinkebyInfura"

#Make deploy build go files from .sol ,deploy it to test net and update onChain.json
DOSNetwork/core/testing$make deploy
#Make dock-all build all testing docker images
DOSNetwork/core/testing$make dock-all

6)Run unit test
DOSNetwork/core/$go test -v ./...

7)Run all unit test and integration test 
DOSNetwork/core$alias gtest="go test \$(go list ./... | grep -v /vendor/) -tags=integration"
DOSNetwork/core$alias gtest -v
