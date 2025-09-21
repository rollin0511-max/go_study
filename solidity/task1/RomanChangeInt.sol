// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract RomanChangeInt {
    // 实现整数转罗马数字
    function intToRoman(uint256 num) public pure returns (string memory) {
        // 使用 require 替换 if 检查，提供中文错误提示
        require(num >= 1 && num <= 3999, unicode"输入的数字不在有效范围内");
        // 定义罗马数字符号和对应的值（从大到小）
        uint256[] memory values = new uint256[](13);
        string[] memory symbols = new string[](13);
        values[0] = 1000; symbols[0] = "M";
        values[1] = 900;  symbols[1] = "CM";
        values[2] = 500;  symbols[2] = "D";
        values[3] = 400;  symbols[3] = "CD";
        values[4] = 100;  symbols[4] = "C";
        values[5] = 90;   symbols[5] = "XC";
        values[6] = 50;   symbols[6] = "L";
        values[7] = 40;   symbols[7] = "XL";
        values[8] = 10;   symbols[8] = "X";
        values[9] = 9;    symbols[9] = "IX";
        values[10] = 5;   symbols[10] = "V";
        values[11] = 4;   symbols[11] = "IV";
        values[12] = 1;   symbols[12] = "I";

        // 动态构建结果字符串
        bytes memory result = "";
        for (uint i = 0; i < values.length; i++) {
            while (num >= values[i]) {
                result = abi.encodePacked(result, symbols[i]);
                num -= values[i];
            }
        }

        string memory roman = string(result);
        return roman;
    }

    // 罗马数字转整数
    function romanToInt(string memory s) public pure returns (uint256) {
        require(bytes(s).length > 0, unicode"输入字符串不能为空");

        bytes memory input = bytes(s);
        uint256 result = 0;

        for (uint i = 0; i < input.length; i++) {
            uint256 current = charToValue(input[i]);
            require(current > 0, unicode"无效的罗马数字字符");
            if (i < input.length - 1 && current < charToValue(input[i + 1])) {
                result -= current;
            } else {
                result += current;
            }
        }

        require(result >= 1 && result <= 3999, unicode"罗马数字值不在有效范围内");
        return result;
    }

    // 辅助函数：将罗马字符转换为对应的值
    function charToValue(bytes1 c) private pure returns (uint256) {
        if (c == 'I') return 1;
        if (c == 'V') return 5;
        if (c == 'X') return 10;
        if (c == 'L') return 50;
        if (c == 'C') return 100;
        if (c == 'D') return 500;
        if (c == 'M') return 1000;
        return 0; // 无效字符
    }
}