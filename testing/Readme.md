### Build and run test:
- Only remove dosenv when the project has new dependency

  `$docker image ls | grep dosenv`
  
  `$docker image rm dosenv`

- Set up the test accounts.

  `DOSNetwork/core$vi credential/passPhrase`
  
  `DOSNetwork/core$mv Downloads/testAccounts .`

- Build client node

  `DOSNetwork/core$make build`

- Build all docker testing images

  - Set ENVIRONMENT variable CHAINNODE to indicate which eth full node we want to access
 
    `DOSNetwork/core/testing$export CHAINNODE="rinkebyInfura"`

  - Build go files from .sol ,deploy contracts to test net and update onChain.json
  
    `DOSNetwork/core/testing$make deploy`
 
  - Build all testing docker images
  
    `DOSNetwork/core/testing$make dock-all`
 
- Run all test

  - Run all unit test
  
    `DOSNetwork/core/$go test -v ./...`

  - Run all unit test and integration test 
  
    `DOSNetwork/core$alias gtest="go test \$(go list ./... | grep -v /vendor/) -tags=integration"`
 
    `DOSNetwork/core$gtest -v`
