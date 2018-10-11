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
        uint query_id;
        G2Point handled_group;
        // User contract issued the query.
        address callback_addr;
    }

    uint nextGroupID = 0;
    uint[] nodeId;
    mapping(uint => uint[]) groupMapping;
    // calling query_id => PendingQuery metadata
    mapping(uint => PendingQuery) pending_queries;
    // Note: Update to groupPubKeys and groups must be made together and atomic.
    G2Point[] groupPubKeys;
    // group_identifier => is_existed
    mapping(bytes32 => bool) groups;
    // Note: Update to randomness metadata must be made atomic.
    // last block number within contains the last updated randomness.
    uint public last_updated_blk = block.number;
    uint public last_randomness = uint(keccak256(abi.encodePacked(blockhash(last_updated_blk), blockhash(last_updated_blk), blockhash(last_updated_blk))));
    G2Point last_handled_group;

    // Log struct is an experimental feature, use with care.
    // event LogUrl(uint queryId, string url, uint timeout, G2Point dispatched_group);
    event LogUrl(uint queryId, string url, uint timeout, uint[4] dispatched_group);
    event LogNonSupportedType(string query_type);
    event LogNonContractCall(address from);
    event LogCallbackTriggeredFor(address callback_addr, bytes result);
    event LogQueryFromNonExistentUC();
    event LogUpdateRandom(uint last_randomness, uint last_blknum, uint[4] dispatched_group);
    event LogInvalidSignature();
    event LogInsufficientGroupNumber();
    event LogGrouping(uint GroupId, uint[] NodeId);

    function getCodeSize(address addr) internal constant returns (uint size) {
        assembly {
            size := extcodesize(addr)
        }
    }

    function strEqual(string a, string b) internal pure returns (bool) {
        bytes memory a_bytes = bytes(a);
        bytes memory b_bytes = bytes(b);
        if (a_bytes.length != b_bytes.length) {
            return false;
        }
        for(uint i = 0; i < a_bytes.length; i++) {
            if (a_bytes[i] != b_bytes[i]) {
                return false;
            }
        }
        return true;
    }

    // @return query id.
    // TODO: restrict query from subscribed/paid calling contract address.
    function query(address from, uint blk_num, uint timeout, string query_type, string query_path)
    external
    returns (uint)
    {
        if (getCodeSize(from) > 0) {
            // Only supporting api/url for alpha release.
            if (strEqual(query_type, 'API')) {
                uint query_id = uint(keccak256(abi.encodePacked(
                        from, blk_num, timeout, query_type, query_path)));
                uint idx = last_randomness % groupPubKeys.length;
                pending_queries[query_id] = PendingQuery(query_id, groupPubKeys[idx], from);
                emit LogUrl(query_id, query_path, timeout, getGroupPubKey(idx));
                return query_id;
            } else {
                emit LogNonSupportedType(query_type);
                return 0x0;
            }
        } else {
            // Skip if @from is not contract address.
            emit LogNonContractCall(from);
            return 0x0;
        }
    }

    function triggerCallback(uint query_id, bytes result, uint[2] sig) external {
        // TODO
        // 1. Check msg.sender from registered and staked node operator. (post alpha)
        // 2. Check msg.sender belongs to pending_queries[query_id].handled_group (alpha)
        // 3. Check whether group signature is valid or not (alpha)
        // Only 3) is implemented below, 1 & 2 & 3 can be implemented in modifier
        // and reused in updateRandomness().
        G1Point memory signature = G1Point(sig[0], sig[1]);
        // TODO: change to sha3(result) after off-chain clients signs on sha3(result)
        if (!verifyGroupSignature(result, signature,
            pending_queries[query_id].handled_group)) {
            emit LogInvalidSignature();
            return;
        }
        address uc_addr = pending_queries[query_id].callback_addr;
        if (uc_addr == 0x0) {
            emit LogQueryFromNonExistentUC();
            return;
        }
        emit LogCallbackTriggeredFor(uc_addr, result);

        UserContractInterface(uc_addr).__callback__(query_id, result);
        delete pending_queries[query_id];
    }

    function updateRandomness(uint[2] sig) external {
        // TODO
        // 1. Check msg.sender from registered and staked node operator. (post alpha)
        // 2. Check msg.sender belongs to last_handled_group (alpha)
        // 3. Check whether group signature is valid or not (alpha)
        // Only 3) is implemented below, 1 & 2 & 3 can be implemented in modifier
        // and reused in triggerCallback().

        // TODO: The message off-chain clients signed: (last_blockhash || last_randomness)
        G1Point memory signature = G1Point(sig[0], sig[1]);
        if (!verifyGroupSignature(toBytes(last_randomness), signature, last_handled_group)) {
            emit LogInvalidSignature();
            return;
        }
        // Update new randomness = sha3(group signature)
        last_randomness = uint(keccak256(abi.encodePacked(signature.x, signature.y, blockhash(last_updated_blk))));
        last_updated_blk = block.number;
        uint idx = last_randomness % groupPubKeys.length;
        last_handled_group = groupPubKeys[idx];
        // Signal off-chain clients
        emit LogUpdateRandom(last_randomness, last_updated_blk, getGroupPubKey(idx));
    }

    function toBytes(uint num) internal returns (bytes numBytes) {
        numBytes = new bytes(32);
        assembly { mstore(add(numBytes, 32), num) }
    }

    // TODO: change to bytes32 in future.
    function verifyGroupSignature(bytes message, G1Point signature, G2Point grpPubKey)
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
        bytes32 group_id = keccak256(abi.encodePacked(x1, x2, y1, y2));
        require(!groups[group_id], "group has already registered");

        groupPubKeys.push(G2Point([x1, x2], [y1, y2]));
        groups[group_id] = true;
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
        uint groupId = nextGroupID++;
        groupMapping[groupId] = toBeGrouped;
        emit LogGrouping(groupId, toBeGrouped);
    }

    function resetContract() {
        nextGroupID = 0;
        nodeId.length = 0;
    }
}