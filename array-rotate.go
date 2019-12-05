package main

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//示例 2:
//
//输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释:
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]
//说明:
//
//尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
//要求使用空间复杂度为 O(1) 的 原地 算法。

//遍历k次，效率低
//func rotate(nums []int, k int) {
//	numLen := len(nums)
//	for i := 0; i < k; i++ {
//		// 单次向右旋转
//		// 临时存放元素的变量(初始化为最右侧元素)
//		var tmp = nums[numLen-1]
//		// 因为是右旋转，从后开始遍历
//		for j := numLen - 1; j > 0; j-- {
//			nums[j] = nums[j-1]
//		}
//		nums[0] = tmp
//	}
//
//}

// 高效！
// 思路：当我们旋转数组 k 次， k%n 个尾部元素会被移动到头部，剩下的元素会被向后移动。
//
//在这个方法中，我们首先将所有元素反转。然后反转前 k 个元素，再反转后面 n-k 个元素，就能得到想要的结果。

func rotate(nums []int, k int) {
	numMoveFront:=k%len(nums)
	// 整体反转数组
	reverseNums(nums)
	// 反转前numMoveFront个元素，即将需要移到前面的元素移动完毕
	reverseNums(nums[:numMoveFront])
	// 反转后面的元素，即将不需要移到开头的元素全部后移
	reverseNums(nums[numMoveFront:])
}

// 反转数组
func reverseNums(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
