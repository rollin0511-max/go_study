// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// 引入openzeppelin 权限控制
import "@openzeppelin/contracts/access/Ownable.sol";

// 讨饭合约
contract BeggingContract is Ownable {
    // 记录每个捐赠者的捐赠金额
    mapping(address => uint256) public donations;

    // 前 3 名捐赠者地址（额外挑战：排行榜）
    address[3] public topDonors;

    // 时间限制（额外挑战）
    uint256 public startTime;
    uint256 public endTime;

    // 事件：记录每次捐赠（额外挑战）
    event Donation(address indexed donor, uint256 amount);

    constructor() Ownable(msg.sender) {
        // 构造函数：设置合约所有者、时间窗口（1 周）
        startTime = block.timestamp;
        endTime = startTime + 1 weeks; // 1 周后结束捐赠
    }

    // 捐赠函数：允许用户发送 ETH，并记录金额（包含时间限制和排行榜更新）
    function donate() public payable {
        // 捐赠金额必须大于0
        require(msg.value > 0, "Donation must be greater than 0");
        // 当前时间不在捐赠时间范围内
        require(block.timestamp >= startTime && block.timestamp <= endTime, "Donation window is closed");

        // 更新捐赠记录
        donations[msg.sender] += msg.value;

        // 更新排行榜（额外挑战：简单逻辑，检查是否进入前 3）
        updateTopDonors(msg.sender, donations[msg.sender]);

        emit Donation(msg.sender, msg.value);
    }

    // 提取函数：仅所有者可调用，提取所有资金
    function withdraw() public onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");
        payable(owner()).transfer(balance);
    }

    // 查询函数：查看指定地址的捐赠金额
    function getDonation(address donor) public view returns (uint256) {
        return donations[donor];
    }

    // 获取前 3 名捐赠者（额外挑战：排行榜）
    function getTopDonors() public view returns (address[3] memory) {
        return topDonors;
    }

    // 内部函数：更新前 3 名（额外挑战）
    function updateTopDonors(address donor, uint256 amount) internal {
        // 初始化 topDonors（如果为空）
        bool hasDonor = false;
        for (uint i = 0; i < 3; i++) {
            if (topDonors[i] == donor) {
                hasDonor = true;
                break;
            }
        }

        if (hasDonor) return; // 已在前 3，不变

        // 找到最低捐赠者（简化：假设我们只比较新捐赠者与当前 top 3 的捐赠金额）
        uint256 minAmount = type(uint256).max;
        uint minIndex = 3;
        for (uint i = 0; i < 3; i++) {
            if (topDonors[i] == address(0)) {
                // 空位，直接插入
                topDonors[i] = donor;
                return;
            }
            uint256 currentAmount = donations[topDonors[i]];
            if (currentAmount < minAmount) {
                minAmount = currentAmount;
                minIndex = i;
            }
        }

        // 如果新捐赠 > 最低，替换
        if (amount > minAmount && minIndex < 3) {
            topDonors[minIndex] = donor;
        }
    }
}