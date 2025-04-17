package main

/*
438. 找到字符串中所有字母异位词
中等
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

与【567. 字符串的排列】解法一致
滑动窗口解法
*/

func findAnagrams(s string, p string) []int {
	var (
		// 记录所需的字符出现次数
		need = make(map[byte]int)
		// 记录 window 中的字符出现次数
		window = make(map[byte]int)
		// 表示窗口中满足 need 条件的字符个数
		valid int
		// 左右指针
		left, right int
		// 记录结果
		res = make([]int, 0)
	)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	for right < len(s) {
		c := s[right]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		right++
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
			left++
		}
	}
	return res
}
