package main

//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
//在杨辉三角中，每个数是它左上方和上方的数的和。
//
//示例:
//
//输入: 5
//输出:
//[
//[1],
//[1,1],
//[1,2,1],
//[1,3,3,1],
//[1,4,6,4,1]
//]
func generate(numRows int) [][]int {
	out := make([][]int, numRows)
	// 生成三角边
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		// 将对角线和左侧置1
		tmp[0],tmp[i]=1,1
		out[i]=tmp


	}
	// 计算杨辉三角内的其他数
	for i := 2; i < numRows; i++ {
		// i为行数
		for j := 1; j < i; j++ {
			out[i][j] = out[i-1][j] + out[i-1][j-1]
		}
	}
	return out
}
