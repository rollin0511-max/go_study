// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Reverse {
    // 反转一个字符串。输入 "abcde"，输出 "edcba"
    function reverse(string memory input) public pure returns (string memory) {
        // 将字符串转换为 bytes 数组
        bytes memory bytesInput = bytes(input);
        // 获取输入字符串的长度
        uint length = bytesInput.length;
        // 创建一个新的 bytes 数组来存储反转后的字符串
        bytes memory reversedBytes = new bytes(length);
        // 遍历输入字符串，将每个字符从后向前复制到新数组中
        for (uint i = 0; i < length; i++) {
            reversedBytes[i] = bytesInput[length - 1 - i];
        }
        // 将 bytes 数组转换为字符串并返回
        return string(reversedBytes);
    }

}