pragma solidity ^0.4.24;

import "./Ownable.sol";

contract DOSAddressBridge is Ownable {
  // Deployed DOSProxy contract address.
  address private _proxy_address;

  event ProxyAddressUpdated(address previous_proxy, address new_proxy);

  function getProxyAddress() external view returns (address) {
    return _proxy_address;
  }

  function setProxyAddress(address new_addr) public onlyOwner {
    emit ProxyAddressUpdated(_proxy_address, new_addr);
    _proxy_address = new_addr;
  }
}
