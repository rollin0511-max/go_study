// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {
    // 一个mapping来存储候选人的得票数
    mapping (string name => uint256 count) public votes;
    // 存储所有候选人名称的数组
    string[] public candidateList;
    // 一个vote函数，允许用户投票给某个候选人
    function vote( string memory name ) public returns (uint256 currentCount){
        // 如果候选人不在列表中，则添加到列表中
        if (votes[name] == 0 && !isCandidate(name)) {
            candidateList.push(name);
        }
        // 投票数+1
        votes[name] += 1;
        // 返回候选人当前投票数
        return votes[name];
    }
    // 一个getVotes函数，返回某个候选人的得票数
    function getVotes(string memory name) public view returns (uint256) {
        return votes[name];
    }
    // 一个resetVotes函数，重置所有候选人的得票数
    function resetVotes() public {
        for (uint256 i = 0; i < candidateList.length; i++) {
            votes[candidateList[i]] = 0;
        }
    }

    // 辅助函数：检查候选人是否已存在
    function isCandidate(string memory name) private view returns (bool) {
        for (uint i = 0; i < candidateList.length; i++) {
            if (keccak256(bytes(candidateList[i])) == keccak256(bytes(name))) {
                return true;
            }
        }
        return false;
    }

}
