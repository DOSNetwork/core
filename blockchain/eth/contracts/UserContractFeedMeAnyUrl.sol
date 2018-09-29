pragma solidity ^0.4.14;

import "github.com/OpenZeppelin/zeppelin-solidity/contracts/ownership/Ownable.sol";
import "./DOSOnChainSDK.sol";

// A user contract asks anything from off-chain world through a url.
contract UserContractFeedMeAnyUrl is Ownable, DOSOnChainSDK {
    bytes public api_result;
    uint public timestamp;
    bool public repeated_call = false;
    string public last_queried_url;
    event EventCallbackReady(bytes result, uint256 time);

    function setQueryMode(bool new_mode) onlyOwner {
        repeated_call = new_mode;
    }

    function checkAPI(string url) public {
        // With 15 seconds timeout
        last_queried_url = url;
        queryDOS(15, "API", url);
    }

    function __callback__(bytes result) {
        require(msg.sender == fromDOSProxyContract());

        api_result = result;
        timestamp = now;

        EventCallbackReady(result, now);

        if (repeated_call) {
            checkAPI(last_queried_url);
        }
    }
}