// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract MergeSortedArrays {
    // 合并两个有序数组为一个有序数组
    function mergeSortedArrays(uint256[] memory arr1, uint256[] memory arr2) public pure returns (uint256[] memory) {
        // 检查输入数组是否有效
        require(arr1.length > 0 || arr2.length > 0, unicode"至少一个输入数组不能为空");

        // 创建结果数组，大小为两个输入数组长度之和
        uint256[] memory result = new uint256[](arr1.length + arr2.length);
        uint256 i = 0; // arr1 的索引
        uint256 j = 0; // arr2 的索引
        uint256 k = 0; // result 的索引

        // 比较并合并，直到一个数组耗尽
        while (i < arr1.length && j < arr2.length) {
            if (arr1[i] <= arr2[j]) {
                result[k] = arr1[i];
                i++;
            } else {
                result[k] = arr2[j];
                j++;
            }
            k++;
        }

        // 复制 arr1 剩余元素（如果有）
        while (i < arr1.length) {
            result[k] = arr1[i];
            i++;
            k++;
        }

        // 复制 arr2 剩余元素（如果有）
        while (j < arr2.length) {
            result[k] = arr2[j];
            j++;
            k++;
        }

        return result;
    }
}