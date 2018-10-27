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

    // Developers call this function directly, a unique queryId will be
    // returned to callers for them to identify parallel requests.
    // @timeout: Estimated timeout in seconds specified by caller; e.g. 15.
    // Response is not guaranteed if processing time exceeds this value.
    // @queryType: Type of request specified by caller. E.g. 'API'.
    // @queryString: Data source destination specified by caller.
    // E.g.: 'https://api.coinbase.com/v2/prices/ETH-USD/spot'
    // A return value of 0x0 stands for error and a related event would
    // be emitted.
    // TODO: Working on response parser.
    function DOSQuery(uint timeout, string queryType, string queryString)
        resolveAddress
        internal
        returns (uint)
    {
        return dosProxy.query(this, timeout, queryType, queryString);
    }


    // Developers need to override __callback__ to process a corresponding
    // response. A user-defined event could be added to notify the Dapp
    // frontend that the response is ready.
    // @queryId: The unique queryId returned by DOSQuery() for callers to
    // identify parallel responses.
    // @result: Response for the specified queryId.
    function __callback__(uint queryId, bytes result) external {
        // To be overriden in the caller contract.
    }
}
