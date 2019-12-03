package main

//给定一个二进制数组， 计算其中最大连续1的个数。
//
//示例 1:
//
//输入: [1,1,0,1,1,1]
//输出: 3
//解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
//注意：
//
//输入的数组只包含 0 和1。
//输入数组的长度是正整数，且不超过 10,000。

//func findMaxConsecutiveOnes(nums []int) int {
//	// 思路：指针找0，快指针j指向当前0的位置，慢指针i指向前一个0的位置，当前连续1的长度为j-i-1
//	numLen := len(nums)
//	jLimit := numLen - 1
//	// 出于效率优化，maxLen中存j-i的结果，最后退出是统一做 -1 结算
//	var maxLen int
//
//LOOP:
//	for i, j := -1, 0; j < numLen; j++ {
//		for nums[j] != 0 {
//			if j < jLimit {
//				j++
//			} else {
//				// 快指针已经到结尾,且最后元素为1
//				lastCheck := j - i + 1
//				if lastCheck > maxLen {
//					maxLen = lastCheck
//				}
//				break LOOP
//			}
//		}
//		check := j - i
//		if check > maxLen {
//			maxLen = check
//		}
//		// 慢指针更替
//		i = j
//	}
//	return maxLen - 1
//}

func findMaxConsecutiveOnes(nums []int) int {
	// 思路：j指针找1，每找到一个1，自加。当j碰到非1时候，结算i与maxLen来决定是否更新最长长度。
	numLen := len(nums)
	jLimit := numLen - 1
	var maxLen int

	for i, j := 0, 0; j < numLen; j++ {
		for nums[j] == 1 {
			i++
			if j < jLimit {
				j++
			} else {
				// j走到最后一个元素，直接结算
				break
			}
		}
		if i > maxLen {
			maxLen = i
		}
		// 当前长度清零
		i = 0
	}
	return maxLen
}

// 慢死的方法
//func findMaxConsecutiveOnes(nums []int) int {
//	// 思路：j指针找1，每找到一个1，自加。当j碰到非1时候，结算i与maxLen来决定是否更新最长长度。
//	numLen := len(nums)
//
//	var maxLen, curLen int
//	for i := 0; i < numLen; i++ {
//		if nums[i] == 1 {
//			curLen++
//		} else {
//			curLen = 0
//		}
//		if curLen > maxLen {
//			maxLen = curLen
//		}
//
//	}
//	return maxLen
//}
