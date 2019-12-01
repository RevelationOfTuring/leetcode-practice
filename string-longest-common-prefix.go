package main

//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//示例 1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"

//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//说明:
//
//所有输入只包含小写字母 a-z 。

func longestCommonPrefix(strs []string) string {
	lenStrs := len(strs)
	// 边界判断
	if lenStrs == 0 {
		return ""
	}
	// 找到最短元素长度和索引
	var minLen, minStrIndex = len(strs[0]), 0
	for i := 1; i < lenStrs; i++ {
		curLen := len(strs[i])
		if curLen < minLen {
			minLen = curLen
			minStrIndex = i
		}
	}
	// 最短元素
	minLenStr := strs[minStrIndex]
	var cursor int
	Flag:
	for ; cursor < minLen; cursor++ {
		for j := 0; j < lenStrs; j++ {
			if strs[j][cursor] != minLenStr[cursor] {
				break Flag
			}
		}
	}
	return minLenStr[:cursor]

}
