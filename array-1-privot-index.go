package main

//我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
//
//如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
//
//示例 1:
//
//输入:
//nums = [1, 7, 3, 6, 5, 6]
//输出: 3
//解释:
//索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
//同时, 3 也是第一个符合要求的中心索引。
//示例 2:
//
//输入:
//nums = [1, 2, 3]
//输出: -1
//解释:
//数组中不存在满足此条件的中心索引。

//示例 3:
//输入:
//nums = [-1,-1,-1,0,1,1]
//输出: 0
//说明:
//nums 的长度范围为 [0, 10000]。
//任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。
func pivotIndex(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	if numsLen == 1 {
		return 0
	}

	var totalSum int
	for i := 0; i < numsLen; i++ {
		totalSum += nums[i]
	}

	// 边界初始值
	var leftSum int
	rightSum := totalSum - nums[0]

	for i := 0; i < numsLen-1; i++ {
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
		rightSum -= nums[i+1]
	}
	// 判断中间索引为numsLen-1的特殊情况
	if leftSum == 0 {
		return numsLen - 1
	}
	return -1

}
