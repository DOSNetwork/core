pragma solidity ^0.4.14;

contract UserContractInterface {
    function __callback__(bytes) external;
}

contract DOSProxy {

    struct G1Point {
        uint x;
        uint y;
    }

    struct G2Point {
        uint[2] x;
        uint[2] y;
    }

    uint randomNum = uint(keccak256(block.number, block.number, block.number));

    uint nextGroupID = 0;

    uint currentGroup;

    string bootstrapIp;

    uint[] nodeId;

    mapping(uint => uint[]) groupMapping;

    // calling query_id => user contract
    mapping(uint => address) pending_queries;

    // group_id (random number generator) => public key
    mapping(uint => G2Point) groupKeyMapping;

    // query_id => public key
    mapping(uint => G2Point) queryKeyMapping;

    // randomNumber_id => randomNumber
    mapping(uint => uint) randomNumberMapping;

    event LogUrl(uint groupId, uint queryId, string url, uint timeout);
    event LogNonSupportedType(string query_type);
    event LogNonContractCall(address from);
    event LogCallbackTriggeredFor(address user_contract_addr, bytes result);
    event LogQueryFromNonExistentUC();
    event LogInvalidSignature();
    event LogInsufficientGroupNumber();
    event LogGrouping(uint GroupId, uint[] NodeId);
    event LogUpdateRandom(uint randomId, uint groupId, uint preRandomNumber);

    function () public payable {}

    function getCodeSize(address addr) constant internal returns (uint size) {
        assembly {
            size := extcodesize(addr)
        }
    }

    function strEqual(string a, string b) constant internal returns (bool) {
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

    function query(address from, uint block_number, uint timeout, string query_type, string query_path) {
        if (getCodeSize(from) > 0) {
            uint query_id = uint(keccak256(from, block_number, timeout, query_type, query_path));
            pending_queries[query_id] = from;
            // Only supporting api/url for demos.
            if (strEqual(query_type, 'API')) {
                uint assignGroup = currentGroup;
                queryKeyMapping[query_id] = groupKeyMapping[assignGroup];
                LogUrl(assignGroup, query_id, query_path, timeout);
            } else {
                LogNonSupportedType(query_type);
            }
        } else {
            // Skip if @from is not contract address.
            LogNonContractCall(from);
        }
    }

    function triggerCallback(uint query_id, bytes result, uint x, uint y) {
        G1Point memory signature = G1Point(x, y);
        address uc_addr = pending_queries[query_id];
        if (uc_addr == 0x0) {
            LogQueryFromNonExistentUC();
            return;
        }
        if (!verifyBLSTest(query_id, result, signature)) {
            LogInvalidSignature();
            return;
        }
        LogCallbackTriggeredFor(uc_addr, result);

        UserContractInterface(uc_addr).__callback__(result);
        // delete pending_queries[query_id];
    }


    function verifyBLSTest(uint query_id, bytes result, G1Point signature) internal returns (bool) {
        G1Point[] memory p1 = new G1Point[](2);
        G2Point[] memory p2 = new G2Point[](2);
        //The signature has already been applied neg() function offchainly to fit requirement of pairingCheck function
        p1[0] = signature;
        p1[1] = hashToG1(result);
        p2[0] = P2();
        p2[1] = queryKeyMapping[query_id];

        return pairingCheck(p1, p2);

    }

    /// @return the result of computing the pairing check
    /// check passes if e(p1[0], p2[0]) *  .... * e(p1[n], p2[n]) == 1
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


    function P1() internal returns (G1Point) {
        return G1Point(1, 2);
    }


    function P2() internal returns (G2Point) {
        return G2Point(
            [11559732032986387107991004021392285783925812861821192530917403151452391805634,
            10857046999023057135944570762232829481370756359578518086990519993285655852781],

            [4082367875863433681332203403145435568316851327593401208105741076214120093531,
            8495653923123431417604973247489272438418190587263600148770280649306958101930]
        );
    }

    function hashToG1(bytes message) internal returns (G1Point) {
        uint256 h = uint256(keccak256(message));
        return mul(P1(), h);
    }

    function add(G1Point p1, G1Point p2) internal returns (G1Point r) {
        uint[4] memory input;
        input[0] = p1.x;
        input[1] = p1.y;
        input[2] = p2.x;
        input[3] = p2.y;
        bool success;
        assembly {
            success := call(sub(gas, 2000), 6, 0, input, 0xc0, r, 0x60)
        // Use "invalid" to make gas estimation work
            switch success case 0 {invalid}
        }
        require(success);
    }

    function mul(G1Point p, uint s) internal returns (G1Point r) {
        uint[3] memory input;
        input[0] = p.x;
        input[1] = p.y;
        input[2] = s;
        bool success;
        assembly {
            success := call(sub(gas, 2000), 7, 0, input, 0x80, r, 0x60)
        // Use "invalid" to make gas estimation work
            switch success case 0 {invalid}
        }
        require(success);
    }

    function setPublicKey(uint group_id, uint x1, uint x2, uint y1, uint y2) {
        uint[] memory x = new uint[](2);
        uint[] memory y = new uint[](2);
        x[0] = x1;
        x[1] = x2;
        y[0] = y1;
        y[1] = y2;
        groupKeyMapping[group_id] = G2Point([x1,x2], [y1,y2]);
        currentGroup = group_id;
    }

    function getPublicKey(uint group_id) public constant returns(uint, uint, uint, uint) {
        return (groupKeyMapping[group_id].x[0], groupKeyMapping[group_id].x[1], groupKeyMapping[group_id].y[0], groupKeyMapping[group_id].y[1]);
    }

    function setBootstrapIp(string ip) {
        bootstrapIp = ip;
    }

    function getBootstrapIp() public constant returns(string){
        return bootstrapIp;
    }

    function getRandomNum() public constant returns(uint) {
        return randomNum;
    }

    function setRandomNum(uint randomId, uint group_id, uint x, uint y) {
        G1Point memory signature = G1Point(x, y);
        bytes memory randomBytes = toBytes(randomNumberMapping[randomId]);
        if (!verifyRandomNum(group_id, randomBytes, signature)) {
            LogInvalidSignature();
            return;
        }
        randomNum = uint(keccak256(x, y, block.number));
        genRandomNum();
    }

    function genRandomNum() {
        uint randomId = block.number;
        randomNumberMapping[randomId] = randomNum;
        LogUpdateRandom(randomId, currentGroup, randomNum);
    }

    function toBytes(uint num) returns (bytes numBytes) {
        numBytes = new bytes(32);
        assembly { mstore(add(numBytes, 32), num) }
    }

    function verifyRandomNum(uint group_id, bytes randomBytes, G1Point signature) internal returns (bool) {
        G1Point[] memory p1 = new G1Point[](2);
        G2Point[] memory p2 = new G2Point[](2);
        //The signature has already been applied neg() function offchainly to fit requirement of pairingCheck function
        p1[0] = signature;
        p1[1] = hashToG1(randomBytes);
        p2[0] = P2();
        p2[1] = groupKeyMapping[group_id];

        return pairingCheck(p1, p2);

    }

    function uploadNodeId(uint id) {
        nodeId.push(id);
    }

    function grouping(uint size) {
        uint[] memory toBeGrouped = new uint[](size);
        if (nodeId.length < size) {
            LogInsufficientGroupNumber();
            return;
        }
        for (uint i = 0; i < size; i++) {
            toBeGrouped[i] = nodeId[nodeId.length - 1];
            nodeId.length--;
        }
        uint groupId = nextGroupID++;
        groupMapping[groupId] = toBeGrouped;
        LogGrouping(groupId, toBeGrouped);
    }

    function resetContract() {
        nextGroupID = 0;
        nodeId.length = 0;
        bootstrapIp = "";
    }
}