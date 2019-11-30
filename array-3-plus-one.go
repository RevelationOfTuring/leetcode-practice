package main

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//示例 1:
//
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//示例 2:
//
//输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。

func plusOne(digits []int) []int {
	numsLen := len(digits)
	// 准备一个长度多1的数组（防止最高位进位）
	out := make([]int, numsLen+1)
	var bitSum int
	var carryBit = 1

	for i := numsLen - 1; i >= 0; i-- {
		bitSum = digits[i] + carryBit
		carryBit = bitSum / 10
		out[i+1] = bitSum % 10
	}
	// 处理最高位
	out[0] = carryBit
	// 取巧
	return out[1-carryBit:]
}
