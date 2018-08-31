pragma solidity ^0.4.14;

/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
    address public owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);


    /**
     * @dev The Ownable constructor sets the original `owner` of the contract to the sender
     * account.
     */
    function Ownable() public {
        owner = msg.sender;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    /**
     * @dev Allows the current owner to transfer control of the contract to a newOwner.
     * @param newOwner The address to transfer ownership to.
     */
    function transferOwnership(address newOwner) public onlyOwner {
        require(newOwner != address(0));
        OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }
}

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

// A user contract asks anything from off-chain world through a url.
contract UserContractFeedMeAnyUrl is Ownable, DOSOnChainSDK {
    bytes public api_result;
    uint public timestamp;
    bool public repeated_call = false;
    string public last_queried_url;
    event EventCallbackReady(bytes result, uint256 time);

    function setQueryMode(bool new_mode) onlyOwner {
        repeated_call = new_mode;
    }

    function checkAPI(string url) public {
        // With 15 seconds timeout
        last_queried_url = url;
        queryDOS(15, "API", url);
    }

    function __callback__(bytes result) {
        require(msg.sender == fromDOSProxyContract());

        api_result = result;
        timestamp = now;

        EventCallbackReady(result, now);

        if (repeated_call) {
            checkAPI(last_queried_url);
        }
    }
}