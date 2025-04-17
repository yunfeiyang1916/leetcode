package main

import "fmt"

// 无重复字符的最长子串,指的是子串中没有重复的字符串
func LengthOfLongestSubstring() {
	str := "abcabcbb"
	str = "abbccabbc"
	fmt.Println(lengthOfLongestSubstring(str))
}

// 无重复字符的最长子串算法
// @remark 时间复杂度：O(N)，其中 N 是字符串的长度。左指针和右指针分别会遍历整个字符串一次。
//
//	空间复杂度：O(N)
func lengthOfLongestSubstring(s string) int {
	var (
		window      = make(map[byte]int)
		left, right int
		res         int
	)
	for right < len(s) {
		c := s[right]
		window[c]++
		right++
		// 判断左侧窗口是否需要收缩
		for window[c] > 1 {
			d := s[left]
			window[d]--
			left++
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}

// 返回最大值
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
