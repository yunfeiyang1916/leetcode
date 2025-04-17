package main

/*
567. 字符串的排列
中等
提示
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的 排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

滑动窗口解法
*/

func checkInclusion(s1 string, s2 string) bool {
	var (
		// 记录所需的字符出现次数
		need = make(map[byte]int)
		// 记录 window 中的字符出现次数
		window = make(map[byte]int)
		// 表示窗口中满足 need 条件的字符个数
		valid int
		// 左右指针
		left, right int
	)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	for right < len(s2) {
		c := s2[right]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		right++
		// 缩小窗口
		if right-left >= len(s1) {
			// 匹配了直接返回
			if valid == len(need) {
				return true
			}
			d := s2[left]
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
			left++
		}
	}
	return false
}
