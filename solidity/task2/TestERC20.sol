// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 参考 openzeppelin的IERC20
//import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

// IERC20 接口定义（参考 OpenZeppelin）
interface IERC20 {
    // 代币的总量
    function totalSupply() external view returns (uint256);
    // 查询指定账户的代币余额
    function balanceOf(address account) external view returns (uint256);
    // 调用者将代币转账给接收者
    function transfer(address recipient, uint256 amount) external returns (bool);
    // 查询某个账户授权给另一账户的代币数量
    function allowance(address owner, address spender) external view returns (uint256);
    // 调用者授权 spender 消费者转移指定数量的代币
    function approve(address spender, uint256 amount) external returns (bool);
    // spender 消费者 从 sender 发送者账户转移代币
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    // 定义 Transfer 事件，记录代币转账操作
    event Transfer(address indexed from, address indexed to, uint256 value);
    // 定义 Approval 事件，记录授权操作。
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

// 实现测试合约
contract TestERC20 is IERC20 {
    // 定义私有 mapping 存储每个账户的代币余额
    mapping(address => uint256) private _balances;
    // 定义私有嵌套 mapping，存储每个账户对其他账户的授权金额
    mapping(address => mapping(address => uint256)) private _allowances;
    // 存储代币总供给量
    uint256 private _totalSupply;
    // 存储合约所有者地址，用于权限控制（如 mint）
    address private _owner;

    // 存储代币名称（例如 "MyToken"）
    string public name;
    // 存储代币符号（例如 "MTK"）
    string public symbol;

    // 构造函数：设置代币名称、符号、初始供给，并指定所有者
    constructor(string memory tokenName, string memory tokenSymbol, uint256 initialSupply) {
        name = tokenName;
        symbol = tokenSymbol;
        _owner = msg.sender;
        // 调用内部 _mint 函数，为部署者增发初始供给的代币。
        _mint(msg.sender, initialSupply);
    }

    // 修饰符：仅允许所有者调用
    modifier onlyOwner() {
        require(msg.sender == _owner, "Caller is not the owner");
        _;
    }

    // totalSupply：返回总供给
    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }

    // balanceOf：查询账户余额
    function balanceOf(address account) external view override returns (uint256) {
        return _balances[account];
    }

    // transfer：转账
    function transfer(address recipient, uint256 amount) external override returns (bool) {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    // allowance：查询授权金额
    function allowance(address owner, address spender) external view override returns (uint256) {
        return _allowances[owner][spender];
    }

    // approve：授权
    function approve(address spender, uint256 amount) external override returns (bool) {
        _approve(msg.sender, spender, amount);
        return true;
    }

    // transferFrom：代扣转账
    function transferFrom(address sender, address recipient, uint256 amount) external override returns (bool) {
        _transfer(sender, recipient, amount);
        _approve(sender, msg.sender, _allowances[sender][msg.sender] - amount);
        return true;
    }

    // mint：增发代币（仅所有者）
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    // 内部函数：转账逻辑
    function _transfer(address sender, address recipient, uint256 amount) internal {
        require(sender != address(0), "Transfer from the zero address");
        require(recipient != address(0), "Transfer to the zero address");
        require(_balances[sender] >= amount, "Transfer amount exceeds balance");

        _balances[sender] -= amount;
        _balances[recipient] += amount;
        emit Transfer(sender, recipient, amount);
    }

    // 内部函数：授权逻辑
    function _approve(address owner, address spender, uint256 amount) internal {
        require(owner != address(0), "Approve from the zero address");
        require(spender != address(0), "Approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    // 内部函数：增发逻辑
    function _mint(address account, uint256 amount) internal {
        require(account != address(0), "Mint to the zero address");

        _totalSupply += amount;
        _balances[account] += amount;
        emit Transfer(address(0), account, amount);
    }
}