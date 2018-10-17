pragma solidity ^0.4.24;

import "github.com/OpenZeppelin/zeppelin-solidity/contracts/ownership/Ownable.sol";
import "../DOSOnChainSDK.sol";

// A user contract asks anything from off-chain world through a url.
contract AskMeAnything is Ownable, DOSOnChainSDK {
  string public response;
  // query_id -> valid_status
  mapping(uint => bool) private _valid;
  bool public repeated_call = false;
  // Default timeout in seconds: Two blocks.
  uint public timeout = 14 * 2;
  string public last_queried_url;

  event SetTimeout(uint previous_timeout, uint new_timeout);
  event CallbackReady(uint query_id);

  function setQueryMode(bool new_mode) public onlyOwner {
    repeated_call = new_mode;
  }

  function setTimeout(uint new_timeout) public onlyOwner {
    emit SetTimeout(timeout, new_timeout);
    timeout = new_timeout;
  }

  // Ask me anything (AMA) off-chain through an api/url.
  function AMA(string url) public {
    last_queried_url = url;
    uint id = DOSQuery(timeout, "API", url);
    if (id != 0x0) {
        _valid[id] = true;
    } else {
        revert("Invalid query id.");
    }
  }

  // User-defined callback function to take and process response.
  function __callback__(uint query_id, bytes result) external {
    require(msg.sender == fromDOSProxyContract(), "Unauthenticated response from non-DOS.");
    require(_valid[query_id], "Response with invalid query id!");

    emit CallbackReady(query_id);
    response = string(result);
    delete _valid[query_id];

    if (repeated_call) {
        AMA(last_queried_url);
    }
  }
}
