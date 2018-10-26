pragma solidity ^0.4.24;

interface DOSProxyInterface {
  function query(address, uint, uint, string, string) external returns (bool);
}

interface DOSAddressBridgeInterface {
  function getProxyAddress() external view returns (address);
}

contract DOSOnChainSDK {
  DOSProxyInterface dosProxy;
  DOSAddressBridgeInterface dosAddrBridge =
    DOSAddressBridgeInterface(0x87095a8115b8385E6A4852640eC9852cD9b6ad9E);
  uint queryIdSeed;

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
      uint curQueryIdSeed = ++queryIdSeed;
      uint curQueryId = uint(keccak256(abi.encodePacked(curQueryIdSeed, this)));
      if (!dosProxy.query(this, curQueryId, timeout, queryType, queryString)) {
        curQueryId = 0x0;
      }
      return curQueryId;
    }
}
