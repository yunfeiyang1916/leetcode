package main

/*
5. 最长回文子串
中等
给你一个字符串 s，找到 s 中最长的 回文 子串。
*/

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		// 先以i为中心扩展查找最长回文串
		r1 := palindrome(s, i, i)
		// 在以i,i+1为中心扩展查找最长回文串
		r2 := palindrome(s, i, i+1)
		if len(r1) < len(r2) {
			r1 = r2
		}
		if len(res) < len(r1) {
			res = r1
		}
	}
	return res
}

// 以i,j为中心点扩散，查找回文串
func palindrome(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}
