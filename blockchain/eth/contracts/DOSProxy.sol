pragma solidity ^0.4.24;
// Not enabled for production yet.
//pragma experimental ABIEncoderV2;

interface UserContractInterface {
    function __callback__(uint, bytes) external;
}

contract DOSProxy {
    // TODO(jonny): Convert to library.
    struct G1Point {
        uint x;
        uint y;
    }

    struct G2Point {
        uint[2] x;
        uint[2] y;
    }

    struct PendingQuery {
        uint queryId;
        G2Point handledGroup;
        // User contract issued the query.
        address callbackAddr;
    }

    uint[] nodeId;
    // calling queryId => PendingQuery metadata
    mapping(uint => PendingQuery) pendingQueries;
    // Note: Update to groupPubKeys and groups must be made together and atomic.
    G2Point[] groupPubKeys;
    // groupIdentifier => isExisted
    mapping(bytes32 => bool) groups;
    // Note: Update to randomness metadata must be made atomic.
    // last block number within contains the last updated randomness.
    uint public lastUpdatedBlock;
    uint public lastRandomness;
    G2Point lastHandledGroup;

    /* Log struct is an experimental feature, use with care.
    event LogUrl(
        uint queryId,
        string url,
        uint timeout,
        G2Point dispatchedGroup
    );
    */
    event LogUrl(
        uint queryId,
        string url,
        uint timeout,
        uint[4] dispatchedGroup
    );
    event LogNonSupportedType(string queryType);
    event LogNonContractCall(address from);
    event LogCallbackTriggeredFor(address callbackAddr, bytes result);
    event LogQueryFromNonExistentUC();
    event LogUpdateRandom(
        uint lastRandomness,
        uint lastUpdatedBlock,
        uint[4] dispatchedGroup
    );
    event LogInvalidSignature(string err);
    event LogInsufficientGroupNumber();
    event LogGrouping(uint[] NodeId);

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
                emit LogUrl(queryId, queryPath, timeout, getGroupPubKey(idx));
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

    function triggerCallback(uint queryId, bytes result, uint[2] sig) external {
        // TODO
        // 1. Check msg.sender from registered and staked node operator. (post alpha)
        // 2. Check msg.sender belongs to pendingQueries[queryId].handledGroup (alpha)
        // 3. Check whether group signature is valid or not (alpha)
        // Only 3) is implemented below, 1 & 2 & 3 can be implemented in modifier
        // and reused in updateRandomness().
        G1Point memory signature = G1Point(sig[0], sig[1]);
        // TODO: change to sha3(result) after off-chain clients signs on sha3(result)
        if (!verifyGroupSignature(
                result, signature, pendingQueries[queryId].handledGroup)) {
            emit LogInvalidSignature("Callback");
            return;
        }
        address ucAddr = pendingQueries[queryId].callbackAddr;
        if (ucAddr == 0x0) {
            emit LogQueryFromNonExistentUC();
            return;
        }
        emit LogCallbackTriggeredFor(ucAddr, result);

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
        // TODO
        // 1. Check msg.sender from registered and staked node operator. (post alpha)
        // 2. Check msg.sender belongs to lastHandledGroup (alpha)
        // 3. Check whether group signature is valid or not (alpha)
        // Only 3) is implemented below, 1 & 2 & 3 can be implemented in modifier
        // and reused in triggerCallback().

        // TODO: The message off-chain clients signed: concat(lastBlockhash, lastRandomness)
        bytes memory message =
            toBytes([blockhash(lastUpdatedBlock), bytes32(lastRandomness)]);
        G1Point memory signature = G1Point(sig[0], sig[1]);
        if (!verifyGroupSignature(message, signature, lastHandledGroup)) {
            emit LogInvalidSignature("Randomness");
            return;
        }
        // Update new randomness = sha3(group signature)
        lastRandomness =
            uint(keccak256(abi.encodePacked(signature.x, signature.y)));
        lastUpdatedBlock = block.number;
        uint idx = lastRandomness % groupPubKeys.length;
        lastHandledGroup = groupPubKeys[idx];
        // Signal off-chain clients
        emit LogUpdateRandom(lastRandomness, lastUpdatedBlock, getGroupPubKey(idx));
    }

    // For test. To trigger first random number after first grouping has done
    function fireRandom() {
        lastRandomness =
            uint(keccak256(abi.encodePacked(blockhash(block.number), blockhash(block.number), blockhash(block.number))));
        lastUpdatedBlock = block.number;
        uint idx = lastRandomness % groupPubKeys.length;
        lastHandledGroup = groupPubKeys[idx];
        // Signal off-chain clients
        emit LogUpdateRandom(lastRandomness, lastUpdatedBlock, getGroupPubKey(idx));
    }

    // TODO: change to bytes32 in future.
    function verifyGroupSignature(
        bytes message,
        G1Point signature,
        G2Point grpPubKey
    )
        internal
        returns (bool)
    {
        G1Point[] memory p1 = new G1Point[](2);
        G2Point[] memory p2 = new G2Point[](2);
        // The signature has already been applied neg() function offchainly to
        // fit requirement of pairingCheck function
        p1[0] = signature;
        p1[1] = hashToG1(message);
        p2[0] = P2();
        p2[1] = grpPubKey;

        return pairingCheck(p1, p2);
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
            success := call(sub(gas, 2000), 8, 0, add(input, 0x20), mul(inputSize, 0x20), out, 0x20)
            // Use "invalid" to make gas estimation work
            switch success case 0 {invalid}
        }
        require(success);
        return out[0] != 0;
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

    function hashToG1(bytes message) internal returns (G1Point) {
        uint256 h = uint256(keccak256(message));
        return scalarMul(P1(), h);
    }

    function pointAdd(G1Point p1, G1Point p2) internal returns (G1Point r) {
        uint[4] memory input;
        input[0] = p1.x;
        input[1] = p1.y;
        input[2] = p2.x;
        input[3] = p2.y;
        bool success;
        assembly {
            success := call(sub(gas, 2000), 6, 0, input, 0x80, r, 0x40)
            // Use "invalid" to make gas estimation work
            switch success case 0 {invalid}
        }
        require(success);
    }

    function scalarMul(G1Point p, uint s) internal returns (G1Point r) {
        uint[3] memory input;
        input[0] = p.x;
        input[1] = p.y;
        input[2] = s;
        bool success;
        assembly {
            success := call(sub(gas, 2000), 7, 0, input, 0x60, r, 0x40)
            // Use "invalid" to make gas estimation work
            switch success case 0 {invalid}
        }
        require(success);
    }

    // TODO(jonny): fix all function ACLs below.
    function setPublicKey(uint x1, uint x2, uint y1, uint y2) {
        bytes32 groupId = keccak256(abi.encodePacked(x1, x2, y1, y2));
        require(!groups[groupId], "group has already registered");

        groupPubKeys.push(G2Point([x1, x2], [y1, y2]));
        groups[groupId] = true;
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
