pragma solidity ^0.4.24;

import "./Ownable.sol";
import "./DOSOnChainSDK.sol";

// A user contract asks anything from off-chain world through a url.
contract AskMeAnything is Ownable, DOSOnChainSDK {
    string public response;
    uint public random;
    // query_id -> valid_status
    mapping(uint => bool) private _valid;
    bool public repeatedCall = false;
    // Default timeout in seconds: Two blocks.
    uint public timeout = 14 * 2;
    string public lastQueriedUrl;
    uint public lastRequestedRandom;

    event SetTimeout(uint previousTimeout, uint newTimeout);
    event QueryResponseReady(uint queryId, string result);
    event RequestSent(bool succ, uint requestId);
    event RandomReady(uint generatedRandom);

    modifier auth(uint id) {
        require(msg.sender == fromDOSProxyContract(),
                "Unauthenticated response from non-DOS.");
        require(_valid[id], "Response with invalid request id!");
        _;
    }

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
            emit RequestSent(true, id);
        } else {
            revert("Invalid query id.");
            emit RequestSent(false, id);
        }
    }

    // User-defined callback function handling query response.
    function __callback__(uint queryId, bytes result) external auth(queryId) {
        response = string(result);
        emit QueryResponseReady(queryId, response);
        delete _valid[queryId];

        if (repeatedCall) {
            AMA(lastQueriedUrl);
        }
    }

    // Request a fast but insecure random number to use directly.
    function requestFastRandom() public {
        lastRequestedRandom = random;
        random = DOSRandom(0, now);
    }

    function requestSafeRandom() public {
        lastRequestedRandom = random;
        uint requestId = DOSRandom(1, now);
        _valid[requestId] = true;
        emit RequestSent(true, requestId);
    }

    // User-defined callback function handling newly generated secure
    // random number.
    function __callback__(uint requestId, uint generatedRandom)
        external
        auth(requestId)
    {
        random = generatedRandom;
        emit RandomReady(generatedRandom);
        delete _valid[requestId];
    }
}
