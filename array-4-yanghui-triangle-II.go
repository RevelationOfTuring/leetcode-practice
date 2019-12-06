package main

//给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
//在杨辉三角中，每个数是它左上方和右上方的数的和。
//
//示例:
//
//输入: 3
//输出: [1,3,3,1]

func getRow(rowIndex int) []int {
	// 第0行
	nums := []int{1}
	for i := 1; i <= rowIndex; i++ {
		// 尾部追加1
		nums = append(nums, 1)
		// 倒序计算杨辉三角当前行
		for j:=i-1;j>0;j--{
			nums[j]+=nums[j-1]
		}
	}
	return nums
}


