package main

import "math"

/*
76. 最小覆盖子串
困难
提示
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

滑动窗口解法
*/

// 76. 最小覆盖子串
// @remark 时间复杂度 O(n)
// @remark 空间复杂度 O(n)
func minWindow(s string, t string) string {
	var (
		// 记录所需的字符出现次数
		need = make(map[byte]int)
		// 记录 window 中的字符出现次数
		window = make(map[byte]int)
		// 表示窗口中满足 need 条件的字符个数
		valid int
		// 左右指针
		left, right   int
		start, length = 0, math.MaxInt32
	)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	for right < len(s) {
		// c是将移入窗口的字符
		c := s[right]
		if _, ok := need[c]; ok {
			window[c]++
			// 窗口内的字符个数与需要的字符个数相等，则满足条件
			if window[c] == need[c] {
				valid++
			}
		}
		right++
		// 判断是否需要收缩窗口
		for valid == len(need) {
			// 更新最新小子覆盖串的起始位置和长度
			if right-left < length {
				start = left
				length = right - left
			}

			// 将要移出窗口的字符
			d := s[left]
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
			// 缩小窗口
			left++
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
