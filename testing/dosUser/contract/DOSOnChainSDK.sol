pragma solidity ^0.5.0;

// Comment out utils library if you don't need it to save gas. (L4 and L17)
import "./lib/utils.sol";
import "./Ownable.sol";

contract DOSProxyInterface {
    function query(address, uint, string memory, string memory) public returns (uint);
    function requestRandom(address, uint) public returns (uint);
}

contract DOSPaymentInterface {
    function isSupportedToken(address) public view returns(bool);
    function setPaymentMethod(address consumer,address tokenAddr) public;
}

contract DOSAddressBridgeInterface {
    function getProxyAddress() public view returns (address);
    function getPaymentAddress() public view returns (address);
}

contract ERC20I {
    function balanceOf(address who) public view returns (uint);
    function transfer(address to, uint value) public returns (bool);
    function approve(address spender, uint value) public returns (bool);
}

contract DOSOnChainSDK is Ownable{
    // Comment out utils library if you don't need it to save gas. (L4 and L17)
    using utils for *;

    DOSProxyInterface dosProxy;
    DOSAddressBridgeInterface dosAddrBridge =
        DOSAddressBridgeInterface(0x848C0Bb953755293230b705464654F647967639A);
    address _tokenAddr = 0x214e79c85744CD2eBBc64dDc0047131496871bEe;

    modifier resolveAddress {
        dosProxy = DOSProxyInterface(dosAddrBridge.getProxyAddress());
        _;
    }

    modifier onlySupportedToken(address tokenAddr) {
        DOSPaymentInterface payment = DOSPaymentInterface(dosAddrBridge.getPaymentAddress());
        require(payment.isSupportedToken(tokenAddr), "Not supported token address!");
        _;
    }

    constructor() public {
        address paymentAddr = dosAddrBridge.getPaymentAddress();
        ERC20I(_tokenAddr).approve(paymentAddr, uint(-1));
        DOSPaymentInterface payment = DOSPaymentInterface(dosAddrBridge.getPaymentAddress());
        payment.setPaymentMethod(address(this),_tokenAddr);
    }

    function fromDOSProxyContract() internal view returns (address) {
        return dosAddrBridge.getProxyAddress();
    }

    function DOSWithdraw() public onlyOwner{
        uint amount = ERC20I(_tokenAddr).balanceOf(address(this));
        ERC20I(_tokenAddr).transfer(msg.sender, amount);
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
    function DOSQuery(uint timeout, string memory dataSource, string memory selector)
        internal
        resolveAddress
        returns (uint)
    {
        return dosProxy.query(address(this), timeout, dataSource, selector);
    }

    // @dev: Must override __callback__ to process a corresponding response. A
    //       user-defined event could be added to notify the Dapp frontend that
    //       the response is ready.
    // @queryId: A unique queryId returned by DOSQuery() for callers to
    //           differentiate parallel responses.
    // @result: Response for the specified queryId.
    function __callback__(uint queryId, bytes calldata result) external {
        // To be overridden in the caller contract.
    }

    // @dev: Call this function to request either a fast but insecure random
    //       number or a safe and secure random number delivered back
    //       asynchronously through the __callback__ function.
    //       Depending on the mode, the return value would be a random number
    //       (for fast mode) or a requestId (for safe mode).
    // @seed: Optional random seed provided by caller.
    function DOSRandom(uint seed)
        internal
        resolveAddress
        returns (uint)
    {
        return dosProxy.requestRandom(address(this), seed);
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
