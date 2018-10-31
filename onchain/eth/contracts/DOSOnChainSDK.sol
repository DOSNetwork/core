pragma solidity ^0.4.24;

import "./lib/utils.sol";

interface DOSProxyInterface {
    function query(address, uint, string, string) external returns (uint);
    function requestRandom(address, uint8, uint) external returns (uint);
}

interface DOSAddressBridgeInterface {
    function getProxyAddress() external view returns (address);
}

contract DOSOnChainSDK {
    using utils for *;

    DOSProxyInterface dosProxy;
    DOSAddressBridgeInterface dosAddrBridge =
        DOSAddressBridgeInterface(0xe987926A226932DFB1f71FA316461db272E05317);

    modifier resolveAddress {
        dosProxy = DOSProxyInterface(dosAddrBridge.getProxyAddress());
        _;
    }

    function fromDOSProxyContract() internal view returns (address) {
        return dosAddrBridge.getProxyAddress();
    }

    // @dev: Call this function to get a unique queryId to differentiate
    //       parallel requests. A return value of 0x0 stands for error and a
    //       related event would be emitted.
    // @timeout: Estimated timeout in seconds specified by caller; e.g. 15.
    //           Response is not guaranteed if processing time exceeds this.
    // @dataSource: Data source destination specified by caller.
    //              E.g.: 'https://api.coinbase.com/v2/prices/ETH-USD/spot'
    // @selector: A selector expression provided by caller to filter out
    //            specific data fields out of the raw response. The response
    //            data format (json, xml/html, and more) is identified from
    //            the selector expression.
    //            E.g. Use "$.data.amount" to extract "194.22" out.
    //             {
    //                  "data":{
    //                          "base":"ETH",
    //                          "currency":"USD",
    //                          "amount":"194.22"
    //                  }
    //             }
    //            Check below documentation for details.
    //            (https://dosnetwork.github.io/docs/#/contents/blockchains/ethereum?id=selector).
    function DOSQuery(uint timeout, string dataSource, string selector)
        resolveAddress
        internal
        returns (uint)
    {
        return dosProxy.query(this, timeout, dataSource, selector);
    }

    // @dev: Must override __callback__ to process a corresponding response. A
    //       user-defined event could be added to notify the Dapp frontend that
    //       the response is ready.
    // @queryId: A unique queryId returned by DOSQuery() for callers to
    //           differentiate parallel responses.
    // @result: Response for the specified queryId.
    function __callback__(uint queryId, bytes result) external {
        // To be overridden in the caller contract.
    }

    // @dev: Call this function to request either a fast but insecure random
    //       number or a safe and secure random number delivered back
    //       asynchronously through the __callback__ function.
    //       Depending on the mode, the return value would be a random number
    //       (for fast mode) or a requestId (for safe mode).
    // @mode: 1: safe mode - The asynchronous but safe way to generate a new
    //                       secure random number by a group of off-chain
    //                       clients using VRF and Threshold Signature. There
    //                       would be a fee to run in safe mode.
    //        0: fast mode - Return a random number in one invocation directly.
    //                       The returned random is the sha3 hash of latest
    //                       generated random number by DOS Network combining
    //                       with the optional seed.
    //                       Thus the result should NOT be considered safe and
    //                       is for testing purpose only. It's free of charge.
    // @seed: Optional random seed provided by caller.
    function DOSRandom(uint8 mode, uint seed)
        resolveAddress
        internal
        returns (uint)
    {
        return dosProxy.requestRandom(this, mode, seed);
    }

    // @dev: Must override __callback__ to process a corresponding random
    //       number. A user-defined event could be added to notify the Dapp
    //       frontend that a new secure random number is generated.
    // @requestId: A unique requestId returned by DOSRandom() for requester to
    //             differentiate random numbers generated concurrently.
    // @generatedRandom: Generated secure random number for the specific
    //                   requestId.
    function __callback__(uint requestId, uint generatedRandom) external {
        // To be overridden in the caller contract.
    }
}
