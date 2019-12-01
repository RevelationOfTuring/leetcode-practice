package main

//数组拆分 I
//给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。
//
//示例 1:
//
//输入: [1,4,3,2]
//
//输出: 4
//解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4).
//提示:
//
//n 是正整数,范围在 [1, 10000].
//数组中的元素范围在 [-10000, 10000].

// bubble sort
//func arrayPairSum(nums []int) (sum int) {
//	// 先升序排列
//	numsLen := len(nums)
//	for i := 0; i < numsLen-1; i++ {
//		for j := 0; j < numsLen-i-1; j++ {
//			if nums[j+1] < nums[j] {
//				nums[j+1], nums[j] = nums[j], nums[j+1]
//			}
//		}
//	}
//
//	var maxSum int
//	// 从头开始，紧挨着的两个为一对，求每对的最小值之和
//	for i := 0; i < numsLen; i = i + 2 {
//		maxSum += nums[i]
//	}
//
//	return maxSum
//
//}

// quick sort
func arrayPairSum(nums []int) (sum int) {
	// 先升序排列
	numsLen := len(nums)
	quickSort(0, numsLen-1, nums)

	var maxSum int
	// 从头开始，紧挨着的两个为一对，求每对的最小值之和
	for i := 0; i < numsLen; i = i + 2 {
		maxSum += nums[i]
	}

	return maxSum

}

func quickSort(left, right int, nums []int) {
	if left >= right {
		return
	}

	i, j, base := left, right, nums[left]
	for i < j {
		for ; base <= nums[j] && i != j; {
			j--
		}
		for ; base >= nums[i] && i != j; {
			i++
		}
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			nums[left], nums[i] = nums[i], nums[left]
			break
		}
	}
	quickSort(left, i-1, nums)
	quickSort(i+1, right, nums)
}
