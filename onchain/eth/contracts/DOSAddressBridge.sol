pragma solidity ^0.4.24;

import "./Ownable.sol";

contract DOSAddressBridge is Ownable {
  // Deployed DOSProxy contract address.
  address private _proxy_address;
  // Deployed DOSPayment contract address.
  address private _payment_address;
  // Deployed DOSRegistry contract address.
  address private _registry_address;

  event ProxyAddressUpdated(address previousProxy, address newProxy);
  event PaymentAddressUpdated(address previousPayment, address newPayment);
  event RegistryAddressUpdated(address previousRegistry, address newRegistry);

  function getProxyAddress() external view returns (address) {
    return _proxy_address;
  }

  function setProxyAddress(address newAddr) public onlyOwner {
    emit ProxyAddressUpdated(_proxy_address, newAddr);
    _proxy_address = newAddr;
  }

  function getPaymentAddress() external view returns (address) {
    return _payment_address;
  }

  function setPaymentAddress(address newAddr) public onlyOwner {
    emit PaymentAddressUpdated(_payment_address, newAddr);
    _payment_address = newAddr;
  }

  function getRegistryAddress() external view returns (address) {
    return _registry_address;
  }

  function setRegistryAddress(address newAddr) public onlyOwner {
    emit RegistryAddressUpdated(_registry_address, newAddr);
    _registry_address = newAddr;
  }
}
