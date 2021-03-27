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
// 		   空间复杂度：O(∣Σ∣)
func lengthOfLongestSubstring(s string) int {
	var (
		// 记录每个字符是否出现过
		m     = make(map[byte]int)
		count = len(s)
		// 右指针，初始值为-1，相当于我们在字符串的左边界的左侧，还没有开始移动
		rk = -1
		// 结果
		ans = 0
	)
	for i := 0; i < count; i++ {
		if i != 0 {
			// 向右移动1格
			delete(m, s[i-1])
		}
		for rk+1 < count && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符就是一个无重复字符子串
		ans = Max(ans, rk+1-i)
	}
	return ans
}

// 返回最大值
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
