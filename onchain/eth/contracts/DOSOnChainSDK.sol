pragma solidity ^0.4.24;

interface DOSProxyInterface {
  function query(address, uint, string, string) external returns (uint);
}

interface DOSAddressBridgeInterface {
  function getProxyAddress() external view returns (address);
}

contract DOSOnChainSDK {
  DOSProxyInterface dosProxy;
  DOSAddressBridgeInterface dosAddrBridge =
    DOSAddressBridgeInterface(0xf3c4695fe4DE7BCC79F3d4A9e7A1cC7d9ED7Dd98);

  modifier resolveAddress {
      dosProxy = DOSProxyInterface(dosAddrBridge.getProxyAddress());
      _;
  }

  function fromDOSProxyContract() internal view returns (address) {
      return dosAddrBridge.getProxyAddress();
  }

  // TODO: Working on response parser.
  // @return a unique query_id for parallel requests. A zeroed (0x0) query_id
  // represents error.
  function DOSQuery(uint timeout, string queryType, string queryString)
    resolveAddress
    internal
    returns (uint)
    {
      return dosProxy.query(this, timeout, queryType, queryString);
    }
}
