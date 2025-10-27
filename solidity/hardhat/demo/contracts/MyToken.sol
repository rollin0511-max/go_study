// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// 自定义代币合约
contract MyToken is ERC20 {
    // 构造函数
    constructor(uint256 initialSupply) ERC20("MyToken", "MTK") {
        // 铸造初始供应量给合约部署者
        _mint(msg.sender, initialSupply);
    }
}
