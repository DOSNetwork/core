pragma solidity ^0.4.24;

import "./Ownable.sol";
import "./DOSOnChainSDK.sol";

// A user contract asks anything from off-chain world through a url.
contract AskMeAnything is Ownable, DOSOnChainSDK {
    string public response;
    // query_id -> valid_status
    mapping(uint => bool) private _valid;
    bool public repeatedCall = false;
    // Default timeout in seconds: Two blocks.
    uint public timeout = 14 * 2;
    string public lastQueriedUrl;

    event SetTimeout(uint previousTimeout, uint newTimeout);
    event CallbackReady(uint queryId, string result, uint randomNumber);
    // 0: Random-fast; 1: Random-safe; 2: AMA
    event QuerySent(uint8 trafficType, bool succ, uint queryId);

    function setQueryMode(bool newMode) public onlyOwner {
        repeatedCall = newMode;
    }

    function setTimeout(uint newTimeout) public onlyOwner {
        emit SetTimeout(timeout, newTimeout);
        timeout = newTimeout;
    }

    // Ask me anything (AMA) off-chain through an api/url.
    function AMA(string url) public {
        lastQueriedUrl = url;
        uint id = DOSQuery(timeout, "API", url);
        if (id != 0x0) {
            _valid[id] = true;
            emit QuerySent(2, true, id);
        } else {
            revert("Invalid query id.");
            emit QuerySent(2, false, id);
        }
    }

    function requestRandom(uint8 mode, uint seed) {
        uint receipt = DOSRandom(mode, seed, timeout);
        if (mode == 1) {
            _valid[receipt] = true;
        }
        emit QuerySent(mode, true, receipt);
    }

    // User-defined callback function to take and process response.
    function __callback__(uint queryId, uint8 trafficType, bytes result) external {
        require(msg.sender == fromDOSProxyContract(),
                "Unauthenticated response from non-DOS.");
        require(_valid[queryId], "Response with invalid query id!");

        if (trafficType == 1) {
            emit CallbackReady(queryId, "", bytesToUint(result));
        } else {
            response = string(result);
            emit CallbackReady(queryId, response, 0);
        }
        delete _valid[queryId];

        if (repeatedCall) {
            AMA(lastQueriedUrl);
        }
    }

    function bytesToUint(bytes result) internal pure returns (uint) {
        uint randomNumber;
        for (uint i = 0; i < result.length; i++){
            randomNumber = randomNumber + uint(result[i]) * (2 ** (8 * (result.length - (i + 1))));
        }
        return randomNumber;
    }
}
