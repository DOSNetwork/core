pragma solidity ^0.4.14;

interface DOSProxyInterface {
    function query(address, uint, uint, string, string) ;
}

interface DOSAddressBridgeInterface {
    function getProxyAddress() returns (address);
}

contract DOSOnChainSDK {
    DOSProxyInterface dos_proxy;
    DOSAddressBridgeInterface dos_addr_bridge = DOSAddressBridgeInterface(0x593bce0faf2d3d0863324fffb1a1c988cd22d5e5);
    event LogQueriedDOS();

    modifier resolveAddress {
        dos_proxy = DOSProxyInterface(dos_addr_bridge.getProxyAddress());
        _;
    }

    function fromDOSProxyContract() internal returns (address) {
        return dos_addr_bridge.getProxyAddress();
    }

    function queryDOS(uint timeout, string query_type, string query_string) resolveAddress internal {
        LogQueriedDOS();
        dos_proxy.query(this, block.number, timeout, query_type, query_string);
    }
}
