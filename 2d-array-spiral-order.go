package main

//给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
//
//示例 1:
//
//输入:
//[
//[ 1, 2, 3 ],
//[ 4, 5, 6 ],
//[ 7, 8, 9 ]
//]
//输出: [1,2,3,6,9,8,7,4,5]
//示例 2:
//
//输入:
//[
//[1, 2, 3, 4],
//[5, 6, 7, 8],
//[9,10,11,12]
//]
//输出: [1,2,3,4,8,12,11,10,9,5,6,7]
func spiralOrder(matrix [][]int) []int {
	// 思路：
	// round为轮次(第一轮为0)
	// 1.每一轮顺时针遍历的起点都是(round,round)
	// 2.每一轮矩形的右下角点为(n-1-round,m-1-round)
	// 3.一共能进行多少轮:(m+1)/2和(n+1)/2的最小值

	m := len(matrix)
	// 边界检查
	if m==0{
		return nil
	}
	n := len(matrix[0])
	out := make([]int, m*n)
	var outIndex int
	var totalRound int
	if (m+1)/2 < (n+1)/2 {
		totalRound = (m + 1) / 2
	} else {
		totalRound = (n + 1) / 2
	}

	for round := 0; round < totalRound; round++ {
		minX, minY, maxX, maxY := round, round, n-1-round, m-1-round
		//	向右
		for i := minX; i <= maxX; i++ {
			out[outIndex] = matrix[minY][i]
			outIndex++
		}
		//	向下
		for i := minY + 1; i <= maxY; i++ {
			out[outIndex] = matrix[i][maxX]
			outIndex++
		}

		// 还有横向回程的条件
		if maxY-minY >= 1 {
			//	向左
			for i := maxX - 1; i >= minX; i-- {
				out[outIndex] = matrix[maxY][i]
				outIndex++
			}
		}
		// 还有纵向回程的条件
		if maxX-minX >= 1 {
			//	向上
			for i := maxY - 1; i > minY; i-- {
				out[outIndex] = matrix[i][minX]
				outIndex++
			}
		}
	}
	return out
}
