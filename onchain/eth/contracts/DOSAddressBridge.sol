pragma solidity ^0.4.24;

import "./Ownable.sol";

contract DOSAddressBridge is Ownable {
  // Deployed DOSProxy contract address.
  address private _proxyAddress;
  // Deployed DOSPayment contract address.
  address private _paymentAddress;
  // Deployed DOSRegistry contract address.
  address private _registryAddress;

  event ProxyAddressUpdated(address previousProxy, address newProxy);
  event PaymentAddressUpdated(address previousPayment, address newPayment);
  event RegistryAddressUpdated(address previousRegistry, address newRegistry);

  function getProxyAddress() external view returns (address) {
    return _proxyAddress;
  }

  function setProxyAddress(address newAddr) public onlyOwner {
    emit ProxyAddressUpdated(_proxyAddress, newAddr);
    _proxyAddress = newAddr;
  }

  function getPaymentAddress() external view returns (address) {
    return _paymentAddress;
  }

  function setPaymentAddress(address newAddr) public onlyOwner {
    emit PaymentAddressUpdated(_paymentAddress, newAddr);
    _paymentAddress = newAddr;
  }

  function getRegistryAddress() external view returns (address) {
    return _registryAddress;
  }

  function setRegistryAddress(address newAddr) public onlyOwner {
    emit RegistryAddressUpdated(_registryAddress, newAddr);
    _registryAddress = newAddr;
  }
}
