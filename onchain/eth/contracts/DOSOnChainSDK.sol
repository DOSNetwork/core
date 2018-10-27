pragma solidity ^0.4.24;

interface DOSProxyInterface {
    function query(address, uint, string, string) external returns (uint);
    function requestRandom(uint8, uint) external returns (uint);
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

    // Developers call this function to get a unique queryId to differenciate
    // parallel requests. A return value of 0x0 stands for error and a related
    // event would be emitted.
    // @timeout: Estimated timeout in seconds specified by caller; e.g. 15.
    //           Response is not guaranteed if processing time exceeds this value.
    // @queryType: Type of request specified by caller. E.g. 'API'.
    // @queryString: Data source destination specified by caller.
    //               E.g.: 'https://api.coinbase.com/v2/prices/ETH-USD/spot'
    // TODO: Working on response parser.
    function DOSQuery(uint timeout, string queryType, string queryString)
        resolveAddress
        internal
        returns (uint)
    {
        return dosProxy.query(this, timeout, queryType, queryString);
    }

    // Developers call this function to request either: a fast but unsecure
    // random number or: a safe and secure random number delivered back
    // asynchronously through the __callback__ function.
    // Depending on the mode, the return value would be a random number
    // (for fast mode) or a requestId (for safe mode).
    // @mode: 0: fast mode - Return a random number in one invocation directly.
    //                       The returned random is the sha3 hash of latest
    //                       generated random number by DOS Network combining
    //                       with the optional seed.
    //                       Thus the result should NOT be considered safe and is
    //                       for testing purpose only. It's free of charge.
    //        1: safe mode - The asynchronous but safe way to generate a new
    //                       secure random number by a group of off-chain clients
    //                       using VRF and Threshold Signature. There would be a
    //                       fee to run in safe mode.
    // @seed: Optional random seed provided by caller.
    function DOSRandom(uint8 mode, uint seed)
        resolveAddress
        internal
        returns (uint)
    {
        return dosProxy.requestRandom(mode, seed);
    }

    // Developers need to override __callback__ to process a corresponding
    // response. A user-defined event could be added to notify the Dapp
    // frontend that the response is ready.
    // @requestId: The unique requestId returned by DOSQuery() or DOSRandom()
    //             for callers to differenciate parallel responses.
    // @result: Response for the specified requestId.
    function __callback__(uint requestId, bytes result) external {
        // To be overriden in the caller contract.
    }
}
