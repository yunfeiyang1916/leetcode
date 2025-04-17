package main

/*
647. 回文子串
中等
提示
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

使用双指针，从中心开始向两边扩展，判断是否是回文串。
*/

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		// 以i为中心
		res += extend(s, i, i)
		// 以i和i+1为中心
		res += extend(s, i, i+1)
	}
	return res
}

// 以i,j为中心向两边扩展
func extend(s string, i, j int) int {
	res := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		res++
		i--
		j++
	}
	return res
}
