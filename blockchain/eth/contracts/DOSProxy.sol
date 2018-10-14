pragma solidity ^0.4.24;
// Not enabled for production yet.
//pragma experimental ABIEncoderV2;

contract DOSProxy {
    using BN256 for *;

    struct PendingQuery {
        uint queryId;
        BN256.G2Point handledGroup;
        // User contract issued the query.
        address callbackAddr;
    }

    uint groupSize;
    uint[] nodeId;
    // calling queryId => PendingQuery metadata
    mapping(uint => PendingQuery) pendingQueries;
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

    event LogUrl(
        uint queryId,
        string url,
        uint timeout,
        uint randomness,
    // Log G2Point struct directly is an experimental feature, use with care.
        uint[4] dispatchedGroup
    );
    event LogNonSupportedType(string queryType);
    event LogNonContractCall(address from);
    event LogCallbackTriggeredFor(address callbackAddr);
    event LogQueryFromNonExistentUC();
    event LogUpdateRandom(
        uint lastRandomness,
        uint lastUpdatedBlock,
        uint[4] dispatchedGroup
    );
    // 0: query; 1: random
    event LogValidationResult(
        uint8 trafficType,
        uint trafficId,
        bytes data,
        uint[2] signature,
        uint[4] pubKey,
        bool pass
    );
    event LogInsufficientGroupNumber();
    event LogGrouping(uint[] NodeId);
    event LogPublicKeyAccepted(uint x1, uint x2, uint y1, uint y2);

    function getCodeSize(address addr) internal constant returns (uint size) {
        assembly {
            size := extcodesize(addr)
        }
    }

    function strEqual(string a, string b) internal pure returns (bool) {
        bytes memory aBytes = bytes(a);
        bytes memory bBytes = bytes(b);
        if (aBytes.length != bBytes.length) {
            return false;
        }
        for(uint i = 0; i < aBytes.length; i++) {
            if (aBytes[i] != bBytes[i]) {
                return false;
            }
        }
        return true;
    }

    // @return query id.
    // TODO: restrict query from subscribed/paid calling contract address.
    function query(
        address from,
        uint blkNum,
        uint timeout,
        string queryType,
        string queryPath
    )
    external
    returns (uint)
    {
        if (getCodeSize(from) > 0) {
            // Only supporting api/url for alpha release.
            if (strEqual(queryType, 'API')) {
                uint queryId = uint(keccak256(abi.encodePacked(
                        from, blkNum, timeout, queryType, queryPath)));
                uint idx = lastRandomness % groupPubKeys.length;
                pendingQueries[queryId] =
                PendingQuery(queryId, groupPubKeys[idx], from);
                emit LogUrl(
                    queryId,
                    queryPath,
                    timeout,
                    lastRandomness,
                    getGroupPubKey(idx)
                );
                return queryId;
            } else {
                emit LogNonSupportedType(queryType);
                return 0x0;
            }
        } else {
            // Skip if @from is not contract address.
            emit LogNonContractCall(from);
            return 0x0;
        }
    }

    // Random submitter validation + group signature verification.
    function validateAndVeriry(
        uint8 trafficType,
        uint trafficId,
        bytes data,
        BN256.G1Point signature,
        BN256.G2Point grpPubKey
    )
    internal
    returns (bool)
    {
        // Verification
        BN256.G1Point[] memory p1 = new BN256.G1Point[](2);
        BN256.G2Point[] memory p2 = new BN256.G2Point[](2);
        // The signature has already been applied neg() function offchainly to
        // fit requirement of pairingCheck function
        p1[0] = signature;
        p1[1] = BN256.hashToG1(data);
        p2[0] = BN256.P2();
        p2[1] = grpPubKey;
        if (!BN256.pairingCheck(p1, p2)) {
            emit LogValidationResult(
                trafficType,
                trafficId,
                data,
                [signature.x, signature.y],
                [grpPubKey.x[0], grpPubKey.x[1], grpPubKey.y[0], grpPubKey.y[1]],
                false
            );
            return false;
        } else {
            emit LogValidationResult(
                trafficType,
                trafficId,
                data,
                [signature.x, signature.y],
                [grpPubKey.x[0], grpPubKey.x[1], grpPubKey.y[0], grpPubKey.y[1]],
                true
            );
            return true;
        }
    }

    function triggerCallback(uint queryId, bytes result, uint[2] sig) external {
        address ucAddr = pendingQueries[queryId].callbackAddr;
        if (ucAddr == 0x0) {
            emit LogQueryFromNonExistentUC();
            return;
        }

        if (!validateAndVeriry(
            0,
            queryId,
            result,
            BN256.G1Point(sig[0], sig[1]),
            pendingQueries[queryId].handledGroup))
        {
            return;
        }
        emit LogCallbackTriggeredFor(ucAddr);

        UserContractInterface(ucAddr).__callback__(queryId, result);
        delete pendingQueries[queryId];
    }

    function toBytes(bytes32[2] data) internal pure returns (bytes) {
        bytes memory result = new bytes(32 * 2);
        assembly {
            mstore(add(result, 0x20), mload(data))
            mstore(add(result, 0x40), mload(add(data, 0x20)))
        }
        return result;
    }

    function updateRandomness(uint[2] sig) external {
        if (!validateAndVeriry(
            1,
            lastRandomness,
        // (lastBlockhash || lastRandomness)
            toBytes([blockhash(lastUpdatedBlock), bytes32(lastRandomness)]),
            BN256.G1Point(sig[0], sig[1]),
            lastHandledGroup))
        {
            return;
        }
        // Update new randomness = sha3(group signature)
        lastRandomness = uint(keccak256(abi.encodePacked(sig[0], sig[1])));
        lastUpdatedBlock = block.number;
        uint idx = lastRandomness % groupPubKeys.length;
        lastHandledGroup = groupPubKeys[idx];
        // Signal off-chain clients
        emit LogUpdateRandom(lastRandomness, lastUpdatedBlock, getGroupPubKey(idx));
    }

    // TODO(jonny): fix all function ACLs below.
    // For test. To trigger first random number after first grouping has done
    function fireRandom() {
        lastRandomness = uint(keccak256(abi.encode(blockhash(block.number - 1))));
        lastUpdatedBlock = block.number;
        uint idx = lastRandomness % groupPubKeys.length;
        lastHandledGroup = groupPubKeys[idx];
        // Signal off-chain clients
        emit LogUpdateRandom(lastRandomness, lastUpdatedBlock, getGroupPubKey(idx));
    }

    function setPublicKey(uint x1, uint x2, uint y1, uint y2) {
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

    function uploadNodeId(uint id) {
        nodeId.push(id);
    }

    function grouping(uint size) {
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

    function resetContract() {
        nodeId.length = 0;
        groupPubKeys.length = 0;
    }
}

library BN256 {
    struct G1Point {
        uint x;
        uint y;
    }

    struct G2Point {
        uint[2] x;
        uint[2] y;
    }

    function P1() internal pure returns (G1Point) {
        return G1Point(1, 2);
    }

    function P2() internal pure returns (G2Point) {
        return G2Point(
            [11559732032986387107991004021392285783925812861821192530917403151452391805634,
            10857046999023057135944570762232829481370756359578518086990519993285655852781],

            [4082367875863433681332203403145435568316851327593401208105741076214120093531,
            8495653923123431417604973247489272438418190587263600148770280649306958101930]
        );
    }

    function pointAdd(G1Point p1, G1Point p2) internal returns (G1Point r) {
        uint[4] memory input;
        input[0] = p1.x;
        input[1] = p1.y;
        input[2] = p2.x;
        input[3] = p2.y;
        assembly {
            if iszero(call(sub(gas, 2000), 0x6, 0, input, 0x80, r, 0x40)) {
                revert(0, 0)
            }
        }
    }

    function scalarMul(G1Point p, uint s) internal returns (G1Point r) {
        uint[3] memory input;
        input[0] = p.x;
        input[1] = p.y;
        input[2] = s;
        assembly {
            if iszero(call(sub(gas, 2000), 0x7, 0, input, 0x60, r, 0x40)) {
                revert(0, 0)
            }
        }
    }

    function hashToG1(bytes data) internal returns (G1Point) {
        uint256 h = uint256(keccak256(data));
        return scalarMul(P1(), h);
    }

    // @return the result of computing the pairing check
    // check passes if e(p1[0], p2[0]) *  .... * e(p1[n], p2[n]) == 1
    function pairingCheck(G1Point[] p1, G2Point[] p2) internal returns (bool) {
        require(p1.length == p2.length);
        uint elements = p1.length;
        uint inputSize = elements * 6;
        uint[] memory input = new uint[](inputSize);

        for (uint i = 0; i < elements; i++)
        {
            input[i * 6 + 0] = p1[i].x;
            input[i * 6 + 1] = p1[i].y;
            input[i * 6 + 2] = p2[i].x[0];
            input[i * 6 + 3] = p2[i].x[1];
            input[i * 6 + 4] = p2[i].y[0];
            input[i * 6 + 5] = p2[i].y[1];
        }

        uint[1] memory out;
        bool success;
        assembly {
            success := call(
            sub(gas, 2000),
            0x8,
            0,
            add(input, 0x20),
            mul(inputSize, 0x20),
            out, 0x20
            )
        }
        return success && (out[0] != 0);
    }
}

interface UserContractInterface {
    function __callback__(uint, bytes) external;
}