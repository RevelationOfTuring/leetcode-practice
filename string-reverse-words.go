package main

import "bytes"

//翻转字符串里的单词
//给定一个字符串，逐个翻转字符串中的每个单词。
//
//
//
//示例 1：
//
//输入: "the sky is blue"
//输出: "blue is sky the"
//示例 2：
//
//输入: "  hello world!  "
//输出: "world! hello"
//解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//示例 3：
//
//输入: "a good   example"
//输出: "example good a"
//解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
//
//
//说明：
//
//无空格字符构成一个单词。
//输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

// 效率最低
//func reverseWords(s string) string {
//	strs := strings.Split(s, " ")
//
//	var ret string
//	for i := len(strs) - 1; i >= 0; i-- {
//		if strs[i] == "" {
//			continue
//		}
//		ret += strs[i]
//		if i > 0 {
//			ret += " "
//
//		}
//	}
//
//	strLen:=len(ret)
//	if strLen!=0{
//		if ret[strLen-1]==' '{
//			return ret[:strLen-1]
//		}
//	}
//
//	return ret
//}

func reverseWords(s string) string {
	// string换成[]byte 会大大提高效率
	srcBytes := []byte(s)
	// 分组
	bytesSlice := bytes.Split(srcBytes, []byte{' '})
	var dstBytes []byte
	for i := len(bytesSlice) - 1; i >= 0; i-- {
		// 剔除[]byte{''}的元素
		if len(bytesSlice[i]) == 0 {
			continue
		}
		dstBytes = append(dstBytes, bytesSlice[i]...)
		if i > 0 {
			dstBytes = append(dstBytes, ' ')
		}
	}

	l := len(dstBytes)
	// 判断dstBytes==[]byte{''}的情况
	if l != 0 {
		if dstBytes[l-1] == ' ' {
			dstBytes = dstBytes[:l-1]
		}
	}
	return string(dstBytes)
}
