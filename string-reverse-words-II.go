package main

//反转字符串中的单词 III
//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//
//示例 1:
//
//输入: "Let's take LeetCode contest"
//输出: "s'teL ekat edoCteeL tsetnoc"
//注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

// 效率略低
//func reverseWordsII(s string) string {
//	// string换成[]byte 会大大提高效率
//	srcBytes := []byte(s)
//	// 分组
//	bytesSlice := bytes.Split(srcBytes, []byte{' '})
//	var dstBytes []byte
//	sliceLen := len(bytesSlice)
//	for i := 0; i < sliceLen; i++ {
//		bytesLen := len(bytesSlice[i])
//		// 剔除[]byte{''}的元素
//		if bytesLen == 0 {
//			continue
//		}
//		// 倒序填充
//		for j := bytesLen - 1; j >= 0; j-- {
//			dstBytes = append(dstBytes, bytesSlice[i][j])
//		}
//
//		if i != sliceLen-1 {
//			dstBytes = append(dstBytes, ' ')
//		}
//	}
//
//	// 判断dstBytes==[]byte{''}的情况
//	dstLen := len(dstBytes)
//	if dstLen != 0 {
//		if dstBytes[dstLen-1] == ' ' {
//			dstBytes = dstBytes[:dstLen-1]
//		}
//	}
//	return string(dstBytes)
//
//}

// 效率可以
func reverseWordsII(s string) string {
	sLen := len(s)
	srcBytes := []byte(s)

	var dstBytes []byte
	for i := 0; i < sLen; {
		// 寻找每一个单词的开头索引
		for srcBytes[i] == ' ' {
			i++
		}

		// 寻找每一个单词的结束索引
		indexEnd := i
		for indexEnd < sLen && srcBytes[indexEnd] != ' ' {
			indexEnd++
		}

		// 单词倒序填充
		for curIndex := indexEnd - 1; curIndex >= i; curIndex-- {
			dstBytes = append(dstBytes, srcBytes[curIndex])
		}
		dstBytes = append(dstBytes, ' ')

		// 开始下一轮搜索
		i = indexEnd
	}

	dstLen:=len(dstBytes)
	if dstLen!=0{
		// dstBytes最后一定是' ',剔除返回string
		dstBytes=dstBytes[:dstLen-1]
	}
	return string(dstBytes)
}

// 不用append,考虑到切片自动扩容的效率浪费
//func reverseWordsII(s string) string {
//	sLen := len(s)
//	srcBytes := []byte(s)
//	// 为什么sLen+1:
//	// 防止输入非常规整，没有多余空格，这样输出在最后一个单词填充完后添加' '会指针越界
//	dstBytes := make([]byte, sLen+1)
//	// 计数器
//	var dstCount int
//	for i := 0; i < sLen; {
//		// 寻找每一个单词的开头索引
//		for srcBytes[i] == ' ' {
//			i++
//		}
//
//		// 寻找每一个单词的结束索引
//		indexEnd := i
//		for indexEnd < sLen && srcBytes[indexEnd] != ' ' {
//			indexEnd++
//		}
//
//		// 单词倒序填充
//		for curIndex := indexEnd - 1; curIndex >= i; curIndex-- {
//			dstBytes[dstCount] = srcBytes[curIndex]
//			dstCount++
//		}
//		dstBytes[dstCount] = ' '
//		dstCount++
//
//		// 开始下一轮搜索
//		i = indexEnd
//	}
//
//	// 排除输入中没有单词的情况
//	if dstCount == 0 {
//		dstCount = 1
//	}
//	// dstBytes最后一定是' ',剔除返回string
//	return string(dstBytes[:dstCount-1])
//}
