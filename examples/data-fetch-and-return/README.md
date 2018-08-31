# Demo

#### Deployed on-chain contracts in rinkeby testnet:
* UserContractFeedMeAnyUrl.sol (Updated):
https://rinkeby.etherscan.io/address/0xFfb6A23dE33eB7efE126A844882F40411a02A21b
* DOSOnChainSDK.sol:
https://rinkeby.etherscan.io/address/0x813b5999c6a80019cd2c0c76d2bd27c6c0fe1a93
* DOSAddressBridge.sol:
https://rinkeby.etherscan.io/address/0x593bce0faf2d3d0863324fffb1a1c988cd22d5e5
* DOSProxy.sol (Updated):  
https://rinkeby.etherscan.io/address/0xbD5784b224D40213df1F9eeb572961E2a859Cb80


#### Steps:
1. Open `fetching-and-return.go`, copy content of keyStore json file to `key`, and corresponding passphrase to `passphrase`.
2. `$ go run fetching-and return.go`. Please retry if stuck at "Establishing listen channel to group public key...".
3. Please wait until "Group set-up finished, start listening to query..." is printed.
4. Go to [myetherwallet.com](https://www.myetherwallet.com/#contracts)
5. Switch to `rinkeby net`
6. Contract Address: `0xFfb6A23dE33eB7efE126A844882F40411a02A21b`
7. ABI / JSON Interface:`[{"constant":false,"inputs":[{"name":"url","type":"string"}],"name":"checkAPI","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"last_queried_url","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"result","type":"bytes"}],"name":"__callback__","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"new_mode","type":"bool"}],"name":"setQueryMode","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"api_result","outputs":[{"name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"timestamp","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"repeated_call","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"result","type":"bytes"},{"indexed":false,"name":"time","type":"uint256"}],"name":"EventCallbackReady","type":"event"},{"anonymous":false,"inputs":[],"name":"LogQueriedDOS","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"previousOwner","type":"address"},{"indexed":true,"name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"}]`
8. Access
9. Choose checkAPI function. Give any api/url, e.g. https://api.coinbase.com/v2/prices/ETH-USD/spot, or https://api.coinmarketcap.com/v1/global/, etc.
10. Check https://rinkeby.etherscan.io/address/0xFfb6A23dE33eB7efE126A844882F40411a02A21b and wait for api_result to be updated. Developer/Application would be notitified with EventCallbackReady() when api_result is ready.