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
  event CallbackReady(uint queryId, string result);
  event QuerySent(bool succ, uint queryId);

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
        emit QuerySent(true, id);
    } else {
        revert("Invalid query id.");
        emit QuerySent(false, id);
    }
  }

  // User-defined callback function to take and process response.
  function __callback__(uint queryId, bytes result) external {
    require(msg.sender == fromDOSProxyContract(), "Unauthenticated response from non-DOS.");
    require(_valid[queryId], "Response with invalid query id!");

    emit CallbackReady(queryId, string(result));
//    response = string(result);
    delete _valid[queryId];

    if (repeatedCall) {
        AMA(lastQueriedUrl);
    }
  }
}
