package main

import (
	"fmt"
	"sort"
)

// 136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
	// 定义map 记录每个元素出现的次数 map[key]value
	// key 元素 value 出现次数
	countMap := make(map[int]int)
	for index, num := range nums {
		fmt.Printf("index：%v，num：%v \n", index, num)
		countMap[num] = countMap[num] + 1
	}
	// 打印map
	fmt.Println("countMap:", countMap)

	// 遍历map
	for num, count := range countMap {
		// 如果出现次数为1 则返回
		if count == 1 {
			return num
		}
	}

	// 默认返回0
	return 0
}

// 9. 回文数：给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func isPalindrome(x int) bool {
	// 负数不是回文数
	if x < 0 {
		return false
	}
	// 个位数必然为回文数
	if x > 0 && x < 10 {
		return true
	}
	// 定义原始变量
	original := x
	// 定义翻转后的变量
	reversed := 0

	for x > 0 {
		// 取出个位数
		lastNum := x % 10
		fmt.Println("本次lastNum：", lastNum)
		// 翻转
		reversed = reversed*10 + lastNum
		fmt.Println("当前翻转后数值reversed：", reversed)
		// 去掉个位数
		x = x / 10
		fmt.Println("当前反转后原数值x：", x)
	}
	fmt.Println("original value:", original)
	fmt.Println("reversed value:", reversed)
	// 比较原始数和翻转后的数是否相等
	if original == reversed {
		return true
	}
	return false

}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValidStr(s string) bool {
	// 使用切片模拟栈
	var stack []rune

	// 遍历字符串的每个字符
	for _, char := range s {
		// 如果是左括号，压入栈
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			// 如果是右括号，检查栈是否为空
			if len(stack) == 0 {
				return false
			}
			// 取出栈顶括号
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 检查是否匹配
			if (char == ')' && top != '(') ||
				(char == '}' && top != '{') ||
				(char == ']' && top != '[') {
				return false
			}
		}
	}

	// 检查栈是否为空
	return len(stack) == 0

}

// 编写一个函数来查找字符串数组中的最长公共前缀。 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	// 如果数组长度为0 则返回空字符串
	if len(strs) == 0 {
		return ""
	}
	// 如果数组长度为1 则返回该字符串
	if len(strs) == 1 {
		return strs[0]
	}

	// 循环遍历 数组中的每个字符串 再把字符串拆分为字符数组进行嵌套比对
	// 定义公共前缀
	// 以第一个字符串为初始前缀
	commonPrefix := strs[0]
	// 遍历其他字符串，逐步缩短前缀
	for i := 1; i < len(strs); i++ {
		// 如果当前字符串为空或前缀为空，直接返回 ""
		if strs[i] == "" || commonPrefix == "" {
			return ""
		}
		// 如果公共前缀的字符长度比当前字符串的字符长度长 或者 截取当前字符串从0下标到公共前缀字符串的长度的字符串不等于公共前缀字符串
		for len(commonPrefix) > len(strs[i]) || strs[i][:len(commonPrefix)] != commonPrefix {
			// 缩短公共前缀 然后重新比较 如果公共前缀最终为空直接返回空
			commonPrefix = commonPrefix[:len(commonPrefix)-1]
			if commonPrefix == "" {
				return ""
			}
		}
	}
	// 打印
	fmt.Println("commonPrefix:", commonPrefix)
	return commonPrefix

}

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(nums []int) []int {
	// 从数组的最后一个元素开始遍历
	for i := len(nums) - 1; i >= 0; i-- {
		// 如果当前元素小于9，直接加1并返回
		if nums[i] < 9 {
			nums[i]++
			return nums
		}
		// 如果当前元素为9，设为0并继续遍历前一个元素
		nums[i] = 0
	}
	// 如果所有元素都是9，需要在数组最前面插入1
	nums = append([]int{1}, nums...)
	return nums
}

// 26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
// 可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
func removeDuplicates(nums []int) int {
	// 如果是空数组 或者 是只有一个元素 则直接返回 长度
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	// 定义一个 i 用于记录不重复元素的位置
	i := 0
	// 遍历数组
	for j := 1; j < len(nums); j++ {
		fmt.Printf("nums[j]元素：%v，nums[i]：%v \n", nums[j], nums[i])
		if nums[j] != nums[i] {
			// i 下标自增
			i++
			// 不相等时 赋值给数组 i+1下标 表示没有重复 插入过去
			nums[i] = nums[j]
		} else {
			// 不满足则为重复元素
			fmt.Printf("重复元素：%v，对应原数组中第：%v 个元素 \n", nums[j], j)
		}
		fmt.Println()
		fmt.Println()
	}
	// 最后返回剔除重复后的 数组下标+1 为数组长度
	return i + 1

}

// 56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中
func mergeArea(intervals [][]int) [][]int {
	// 如果空数组 直接返回
	if len(intervals) == 0 {
		return [][]int{}
	}
	// 按起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 定义一个切片来存储合并后的区间
	result := [][]int{intervals[0]}

	// 遍历区间 从第二个元素开始 因为result中已经放置了第一个元素
	for i := 1; i < len(intervals); i++ {
		// 当前区间
		currStart, currEnd := intervals[i][0], intervals[i][1]
		// result 中最后一个区间 的 end值
		lastEnd := result[len(result)-1][1]

		// 如果当前区间与最后一个区间重叠，合并 【当前区间的开始值小于等于result中最后一个区间的结束值】
		if currStart <= lastEnd {
			// 合并区间 【赋值最后一个区间的结束值为 当前区间的结束值 与 result中最后一个区间的结束值 中的较大值】
			result[len(result)-1][1] = max(lastEnd, currEnd)
		} else {
			// 无重叠，追加当前区间到结果区间数组中
			result = append(result, []int{currStart, currEnd})
		}
	}

	return result
}

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func matchNumber(nums []int, target int) []int {
	// 定义map 存储数组元素 和 下标
	numMap := make(map[int]int)
	// 遍历传入数值数组
	for index, num := range nums {
		// 计算所需补数
		complement := target - num
		// 判断是否存在于numMap中 存在则认为找到了 返回对应的数组下标数组
		if j, ok := numMap[complement]; ok {
			return []int{j, index}
		}
		// 不存在则把数据的下标和值 存放到numMap中
		numMap[num] = index
	}
	// 未找到返回空数组
	return []int{}
}

func main() {
	// 只出现一次的数字
	//fmt.Println("[]int{1, 2, 2, 4, 5, 4, 5, 5}数组中只出现一次的数字是：", singleNumber([]int{1, 2, 2, 4, 5, 4, 5, 5}))

	// 验证回文数
	//fmt.Println("-1 是回文数吗？", isPalindrome(-1))
	//fmt.Println("1 是回文数吗？", isPalindrome(1))
	//fmt.Println("121 是回文数吗？", isPalindrome(121))
	//fmt.Println("123 是回文数吗？", isPalindrome(123))

	// 验证括号是否匹配
	//fmt.Println("()", isValidStr("()"))
	//fmt.Println("()[]{}", isValidStr("()[]{}"))
	//fmt.Println("(]", isValidStr("(]"))
	//fmt.Println("([)]", isValidStr("([)]"))
	//fmt.Println("{[]}", isValidStr("{[]}"))

	// 获取字符串数组的公共前缀
	//fmt.Println("strs = [\"flower\",\"flow\",\"flight\"] 的最长公共前缀为：", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println("strs = [\"flower\",\"flow\",\"flight\"] 的最长公共前缀为：", longestCommonPrefix([]string{"dog", "racecar", "car"}))

	// 加一
	//fmt.Println("[]int{1, 2, 3} 加一后的结果为：", plusOne([]int{1, 2, 3}))
	//fmt.Println("[]int{9, 9, 9} 加一后的结果为：", plusOne([]int{9, 9, 9}))

	// 移除数组中的重复元素 后返回数组长度
	// fmt.Println("nums = [0,0,1,1,1,2,2,3,3,4] 移除重复元素后为：", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	// 合并区间
	//fmt.Println("intervals = [[1,3],[2,6],[8,10],[15,18]] 合并重叠区间后：", mergeArea([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))

	// 寻找传入数组与目标值 重复的两个值的前两个下标
	//fmt.Println("nums = [3，2，5] ,target = 8，最终找到的target对应的两个下标分别为：", matchNumber([]int{3, 2, 5}, 8))
}
