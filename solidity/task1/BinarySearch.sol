// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BinarySearch {
    // 在有序数组中使用二分查找目标值，返回索引（未找到返回 -1）
    function binarySearch(uint256[] memory arr, uint256 target) public pure returns (int256) {
        // 检查输入数组是否为空
        require(arr.length > 0, unicode"输入数组不能为空");

        // 二分查找
        uint256 left = 0;
        uint256 right = arr.length - 1;

        // 二分查找循环 当 left 小于等于 right 时继续查找
        while (left <= right) {
            uint256 mid = left + (right - left) / 2; // 防止溢出

            if (arr[mid] == target) {
                return int256(mid); // 找到目标，返回索引
            } else if (arr[mid] < target) {
                left = mid + 1; // 目标在右半部分
            } else {
                right = mid - 1; // 目标在左半部分
            }
        }
        return -1; // 未找到目标
    }

}