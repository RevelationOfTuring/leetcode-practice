package main

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 思路：滑动窗口
func lengthOfLongestSubstring(s string) int {
	bytes := []byte(s)
	book := make([]bool, 128)
	var curLen, maxLen int
	// left 滑动窗口左边
	// right 滑动窗口邮编
	bytesLen:=len(bytes)
	for left, right := 0, 0; right < bytesLen; right++ {

		if !book[bytes[right]] {
			book[bytes[right]] = true
			curLen++
		} else {
			repetition := bytes[right]
			if curLen > maxLen {
				maxLen = curLen
			}

			for {
				if bytes[left] != repetition {
					book[bytes[left]] = false
					curLen--
					left++
				} else {
					left++
					break
				}
			}
		}
	}

	if curLen>maxLen{
		maxLen=curLen
	}

	return maxLen

}
