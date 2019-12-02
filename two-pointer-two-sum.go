package main

//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//
//函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//
//说明:
//
//返回的下标值（index1 和 index2）不是从零开始的。
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
//示例:
//
//输入: numbers = [2, 7, 11, 15], target = 9
//输出: [1,2]
//解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

func twoSum1(numbers []int, target int) []int {
	for cursorLeft, cursorRight := 0, len(numbers)-1; cursorLeft < cursorRight; cursorLeft++ {
		// 本轮所需要找的值
		numNeed := target - numbers[cursorLeft]
		for {
			if numbers[cursorRight] > numNeed {
				// 因为有序，所以该值一定在右指针的当下或左边
				cursorRight--
			} else {
				// 当右指针停止递减的时候判断是否满足返回要求
				// 需要注意：如果不满足要求，左指针+1。由于是升序数列，那么下一个numNeed一定在当下的右指针左侧或当下
				if numbers[cursorRight] == numNeed {
					return []int{cursorLeft + 1, cursorRight + 1}
				}
				break
			}

		}
	}
	// 当左指针和右指针相遇时，说明该数组无满足条件的元素。
	// 这样遍历一遍，就可以找到解

	return nil
}
