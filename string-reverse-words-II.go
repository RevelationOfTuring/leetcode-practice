package main

import "bytes"

//反转字符串中的单词 III
//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//示例 1:
//
//输入: "Let's take LeetCode contest"
//输出: "s'teL ekat edoCteeL tsetnoc"
//注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

// 效率略低
func reverseWordsII(s string) string {
	// string换成[]byte 会大大提高效率
	srcBytes := []byte(s)
	// 分组
	bytesSlice := bytes.Split(srcBytes, []byte{' '})
	var dstBytes []byte
	sliceLen := len(bytesSlice)
	for i := 0; i < sliceLen; i++ {
		bytesLen := len(bytesSlice[i])
		// 剔除[]byte{''}的元素
		if bytesLen == 0 {
			continue
		}
		// 倒序填充
		for j := bytesLen - 1; j >= 0; j-- {
			dstBytes = append(dstBytes, bytesSlice[i][j])
		}

		if i != sliceLen-1 {
			dstBytes = append(dstBytes, ' ')
		}
	}

	// 判断dstBytes==[]byte{''}的情况
	dstLen := len(dstBytes)
	if dstLen != 0 {
		if dstBytes[dstLen-1] == ' ' {
			dstBytes = dstBytes[:dstLen-1]
		}
	}
	return string(dstBytes)

}
