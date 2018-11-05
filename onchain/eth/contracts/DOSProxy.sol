pragma solidity ^0.4.24;
// Not enabled for production yet.
//pragma experimental ABIEncoderV2;

import "./lib/bn256.sol";

interface UserContractInterface {
    // Query callback.
    function __callback__(uint, bytes) external;
    // Random number callback.
    function __callback__(uint, uint) external;
}

contract DOSProxy {
    using BN256 for *;

    struct PendingRequest {
        uint requestId;
        BN256.G2Point handledGroup;
        // User contract issued the query.
        address callbackAddr;
    }

    uint requestIdSeed;
    uint groupSize;
    uint[] nodeId;
    // calling requestId => PendingQuery metadata
    mapping(uint => PendingRequest) PendingRequests;
    // Note: Make atomic changes to group metadata below.
    BN256.G2Point[] groupPubKeys;
    // groupIdentifier => isExisted
    mapping(bytes32 => bool) groups;
    //publicKey => publicKey appearance
    mapping(bytes32 => uint) pubKeyCounter;
    // Note: Make atomic changes to randomness metadata below.
    uint public lastUpdatedBlock;
    uint public lastRandomness;
    BN256.G2Point lastHandledGroup;
    uint8 constant TrafficSystemRandom = 0;
    uint8 constant TrafficUserRandom = 1;
    uint8 constant TrafficUserQuery = 2;

    event LogUrl(
        uint queryId,
        uint timeout,
        string dataSource,
        string selector,
        uint randomness,
        // Log G2Point struct directly is an experimental feature, use with care.
        uint[4] dispatchedGroup
    );
    event LogRequestUserRandom(
        uint requestId,
        uint lastSystemRandomness,
        uint userSeed,
        uint[4] dispatchedGroup
    );
    event LogNonSupportedType(string invalidSelector);
    event LogNonContractCall(address from);
    event LogCallbackTriggeredFor(address callbackAddr);
    event LogRequestFromNonExistentUC();
    event LogUpdateRandom(uint lastRandomness, uint[4] dispatchedGroup);
    event LogValidationResult(
        uint8 trafficType,
        uint trafficId,
        bytes message,
        uint[2] signature,
        uint[4] pubKey,
        bool pass
    );
    event LogInsufficientGroupNumber();
    event LogGrouping(uint[] NodeId);
    event LogPublicKeyAccepted(uint x1, uint x2, uint y1, uint y2);

    // whitelist state variables used only for alpha release.
    // Index starting from 1.
    address[22] whitelists;
    // whitelisted address => index in whitelists.
    mapping(address => uint) isWhitelisted;
    bool public whitelistInitialized = false;
    event WhitelistAddressTransferred(address previous, address curr);

    modifier onlyWhitelisted {
        uint idx = isWhitelisted[msg.sender];
        require(idx != 0 && whitelists[idx] == msg.sender, "Not whitelisted!");
        _;
    }

    function initWhitelist(address[21] addresses) public {
        require(!whitelistInitialized, "Whitelist already initialized!");

        for (uint idx = 0; idx < 21; idx++) {
            whitelists[idx+1] = addresses[idx];
            isWhitelisted[addresses[idx]] = idx+1;
        }
        whitelistInitialized = true;
    }

    function getWhitelistAddress(uint idx) public view returns (address) {
        require(idx > 0 && idx <= 21, "Index out of range");
        return whitelists[idx];
    }

    function transferWhitelistAddress(address newWhitelistedAddr)
        public
        onlyWhitelisted
    {
        require(newWhitelistedAddr != 0x0 && newWhitelistedAddr != msg.sender);

        emit WhitelistAddressTransferred(msg.sender, newWhitelistedAddr);
        whitelists[isWhitelisted[msg.sender]] = newWhitelistedAddr;
    }

    function getCodeSize(address addr) internal constant returns (uint size) {
        assembly {
            size := extcodesize(addr)
        }
    }

    // Returns query id.
    // TODO: restrict query from subscribed/paid calling contracts.
    function query(
        address from,
        uint timeout,
        string dataSource,
        string selector
    )
        external
        returns (uint)
    {
        if (getCodeSize(from) > 0) {
            bytes memory bs = bytes(selector);
            // '': Return whole raw response;
            // Starts with '$': response format is parsed as json.
            // Starts with '/': response format is parsed as xml/html.
            if (bs.length == 0 || bs[0] == '$' || bs[0] == '/') {
                uint queryId = uint(keccak256(abi.encodePacked(
                    ++requestIdSeed, from, timeout, dataSource, selector)));
                uint idx = lastRandomness % groupPubKeys.length;
                PendingRequests[queryId] =
                    PendingRequest(queryId, groupPubKeys[idx], from);
                emit LogUrl(
                    queryId,
                    timeout,
                    dataSource,
                    selector,
                    lastRandomness,
                    getGroupPubKey(idx)
                );
                return queryId;
            } else {
                emit LogNonSupportedType(selector);
                return 0x0;
            }
        } else {
            // Skip if @from is not contract address.
            emit LogNonContractCall(from);
            return 0x0;
        }
    }

    // Request a new user-level random number.
    function requestRandom(address from, uint8 mode, uint userSeed)
        external
        returns (uint)
    {
        // fast mode
        if (mode == 0) {
            return uint(keccak256(abi.encodePacked(
                ++requestIdSeed,lastRandomness, userSeed)));
        } else if (mode == 1) {
            // safe mode
            // TODO: restrict request from paid calling contract address.
            uint requestId = uint(keccak256(abi.encodePacked(
                ++requestIdSeed, from, userSeed)));
            uint idx = lastRandomness % groupPubKeys.length;
            PendingRequests[requestId] =
                PendingRequest(requestId, groupPubKeys[idx], from);
            // sign(requestId ||lastSystemRandomness || userSeed) with
            // selected group
            emit LogRequestUserRandom(
                requestId,
                lastRandomness,
                userSeed,
                getGroupPubKey(idx)
            );
            return requestId;
        } else {
            revert("Non-supported random request");
        }
    }

    // Random submitter validation + group signature verification.
    function validateAndVerify(
        uint8 trafficType,
        uint trafficId,
        bytes data,
        BN256.G1Point signature,
        BN256.G2Point grpPubKey
    )
        internal
        onlyWhitelisted
        returns (bool)
    {
        // Validation
        // TODO
        // 1. Check msg.sender from registered and staked node operator.
        // 2. Check msg.sender is a member in Group(grpPubKey).
        // Clients actually signs (data || addr(selected_submitter)).
        // TODO: Sync and change to sign ( sha256(data) || address )
        bytes memory message = abi.encodePacked(data, msg.sender);

        // Verification
        BN256.G1Point[] memory p1 = new BN256.G1Point[](2);
        BN256.G2Point[] memory p2 = new BN256.G2Point[](2);
        // The signature has already been applied neg() function offchainly to
        // fit requirement of pairingCheck function
        p1[0] = signature;
        p1[1] = BN256.hashToG1(message);
        p2[0] = BN256.P2();
        p2[1] = grpPubKey;
        bool passVerify = BN256.pairingCheck(p1, p2);
        emit LogValidationResult(
            trafficType,
            trafficId,
            message,
            [signature.x, signature.y],
            [grpPubKey.x[0], grpPubKey.x[1], grpPubKey.y[0], grpPubKey.y[1]],
            passVerify
        );
        return passVerify;
    }

    function triggerCallback(
        uint requestId,
        uint8 trafficType,
        bytes result,
        uint[2] sig
    )
        external
    {
        if (!validateAndVerify(
                trafficType,
                requestId,
                result,
                BN256.G1Point(sig[0], sig[1]),
                PendingRequests[requestId].handledGroup))
        {
            return;
        }

        address ucAddr = PendingRequests[requestId].callbackAddr;
        if (ucAddr == 0x0) {
            emit LogRequestFromNonExistentUC();
            return;
        }

        emit LogCallbackTriggeredFor(ucAddr);
        delete PendingRequests[requestId];
        if (trafficType == TrafficUserQuery) {
            UserContractInterface(ucAddr).__callback__(requestId, result);
        } else if (trafficType == TrafficUserRandom) {
            // Safe random number is the collectively signed threshold signature
            // of the message (requestId || lastRandomness || userSeed ||
            // selected sender in group).
            UserContractInterface(ucAddr).__callback__(
                requestId, uint(keccak256(abi.encodePacked(sig[0], sig[1]))));
        } else {
            revert("Unsupported traffic type");
        }
    }

    function toBytes(uint x) internal pure returns (bytes b) {
        b = new bytes(32);
        assembly { mstore(add(b, 32), x) }
    }

    // System-level secure distributed random number generator.
    function updateRandomness(uint[2] sig) external {
        if (!validateAndVerify(
                TrafficSystemRandom,
                lastRandomness,
                toBytes(lastRandomness),
                BN256.G1Point(sig[0], sig[1]),
                lastHandledGroup))
        {
            return;
        }
        // Update new randomness = sha3(collectively signed group signature)
        lastRandomness = uint(keccak256(abi.encodePacked(sig[0], sig[1])));
        lastUpdatedBlock = block.number - 1;
        uint idx = lastRandomness % groupPubKeys.length;
        lastHandledGroup = groupPubKeys[idx];
        // Signal selected off-chain clients to collectively generate a new
        // system level random number for next round.
        emit LogUpdateRandom(lastRandomness, getGroupPubKey(idx));
    }

    // For alpha. To trigger first random number after grouping has done
    // or timeout.
    function fireRandom() public onlyWhitelisted {
        lastRandomness = uint(keccak256(abi.encode(blockhash(block.number - 1))));
        lastUpdatedBlock = block.number - 1;
        uint idx = lastRandomness % groupPubKeys.length;
        lastHandledGroup = groupPubKeys[idx];
        // Signal off-chain clients
        emit LogUpdateRandom(lastRandomness, getGroupPubKey(idx));
    }

    function handleTimeout() public onlyWhitelisted {
        uint currentBlockNumber = block.number - 1;
        if (currentBlockNumber - lastUpdatedBlock > 5) {
            fireRandom();
        }
    }

    function setPublicKey(uint x1, uint x2, uint y1, uint y2)
        public
        onlyWhitelisted
    {
        bytes32 groupId = keccak256(abi.encodePacked(x1, x2, y1, y2));
        require(!groups[groupId], "group has already registered");

        pubKeyCounter[groupId] = pubKeyCounter[groupId] + 1;
        if (pubKeyCounter[groupId] > groupSize / 2) {
            groupPubKeys.push(BN256.G2Point([x1, x2], [y1, y2]));
            groups[groupId] = true;
            delete(pubKeyCounter[groupId]);
            emit LogPublicKeyAccepted(x1, x2, y1, y2);
        }
    }

    function getGroupPubKey(uint idx) public constant returns (uint[4]) {
        require(idx < groupPubKeys.length, "group index out of range");

        return [
            groupPubKeys[idx].x[0], groupPubKeys[idx].x[1],
            groupPubKeys[idx].y[0], groupPubKeys[idx].y[1]
        ];
    }

    function uploadNodeId(uint id) public onlyWhitelisted {
        nodeId.push(id);
    }

    function grouping(uint size) public onlyWhitelisted {
        groupSize = size;
        uint[] memory toBeGrouped = new uint[](size);
        if (nodeId.length < size) {
            emit LogInsufficientGroupNumber();
            return;
        }
        for (uint i = 0; i < size; i++) {
            toBeGrouped[i] = nodeId[nodeId.length - 1];
            nodeId.length--;
        }
        emit LogGrouping(toBeGrouped);
    }

    function resetContract() public onlyWhitelisted {
        nodeId.length = 0;
        groupPubKeys.length = 0;
    }
}
