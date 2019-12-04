package main

//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
//
//示例:
//
//输入: s = 7, nums = [2,3,1,2,4,3]
//输出: 2
//解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
//进阶:
//
//如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

// 效率有点慢
//func minSubArrayLen(s int, nums []int) int {
//	numLen := len(nums)
//	// +1是为了帮助判断整个数组的所有元素加和也达不到s的情况出现
//	minLen := numLen + 1
//	for i := 0; i < numLen; i++ {
//		tmpSum, j := 0, i
//		for j < numLen {
//			tmpSum += nums[j]
//			if tmpSum >= s {
//				curLen := j - i + 1
//				if curLen < minLen {
//					minLen = curLen
//				}
//				break
//			}
//			j++
//		}
//
//	}
//
//	if minLen == numLen+1 {
//		// 表明在遍历了整个数组之后，也没有达到条件：tmpSum >= s，即数组里所有元素加起来也没有到达s
//		minLen = 0
//	}
//	return minLen
//}

// 双指针 + 滑动窗口
func minSubArrayLen(s int, nums []int) int {
	numLen := len(nums)
	// +1是为了帮助判断整个数组的所有元素加和也达不到s的情况出现
	minLen := numLen + 1
	// 记录窗口内连续和的变量
	var tmpSum int
	// 快指针
	var j int
	for i := 0; i < numLen; i, j = i+1, j+1 {
		for j < numLen {
			// 窗口右扩
			if tmpSum >= s {
				curLen := j - i + 1
				if curLen < minLen {
					minLen = curLen
				}

				break
			}
			j++
		}

		for {
			// 窗口左缩
			tmpSum -= nums[i]
			// 对窗口缩小一位后总和仍大于s的追加检查
			// e.g:
			// .... 1,1,100,....        s=50
			// 第一次进入窗口是 1,1,100
			// 窗口左移：... 1,100,... 和仍大于s
			if tmpSum >= s {
				i++
			} else {
				break
			}
		}
		curLen := j - i + 1
		if curLen < minLen {
			minLen = curLen
		}
	}

	if minLen == numLen+1 {
		// 表明在遍历了整个数组之后，也没有达到条件：tmpSum >= s，即数组里所有元素加起来也没有到达s
		minLen = 0
	}
	return minLen
}

//func minSubArrayLen(s int, nums []int) int {
//	numLen := len(nums)
//	// 慢指针和窗口中的连续和
//	var left, tmpSum int
//	var minLen = numLen + 1
//	for right := 0; right < numLen; right++ {
//		tmpSum += nums[right]
//		for tmpSum >= s {
//			curLen := right - left + 1
//			if curLen < minLen {
//				minLen = curLen
//			}
//			tmpSum -= nums[left]
//			left++
//		}
//	}
//
//	if minLen == numLen+1 {
//		minLen = 0
//	}
//
//	return minLen
//
//}
