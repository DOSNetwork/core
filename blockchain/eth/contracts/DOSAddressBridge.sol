pragma solidity ^0.4.14;

import "Ownable.sol";

contract DOSAddressBridge is Ownable {
    // Deployed DOSProxy contract address.
    address public proxy_address;

    function getProxyAddress() returns (address) {
        return proxy_address;
    }

    function setProxyAddress(address new_addr) onlyOwner {
        proxy_address = new_addr;
    }
}    
