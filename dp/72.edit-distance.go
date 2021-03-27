package main

import "fmt"

// 编辑距离，将一个单词转成另一个单词的最小操作数
func MinDistance() {
	//word1 := "horse"
	//word2 := "ros"
	word1 := "sea"
	word2 := "eat"
	fmt.Println(minDistance2(word1, word2))
}

// 编辑距离，将一个单词(word1)转成另一个单词(word2)的最小操作数
// @remark 时间复杂度o(m*n),空间复杂度o(m*n)
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	// 当字符串 word1 的长度为 i，字符串 word2 的长度为 j 时，将 word1 转化为 word2 所使用的最少操作次数为 dp[i] [j]
	// 注意下标的含义表示长度，0表示空字符，所以总长度需要+1
	dp := make([][]int, m+1)
	// 设置初始值
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		if i == 0 {
			dp[0][0] = 0
			for j := 1; j <= n; j++ {
				dp[0][j] = dp[0][j-1] + 1
			}
		} else {
			dp[i][0] = dp[i-1][0] + 1
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 第i个字符对应的下标为i-1
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

// 最小编辑距离优化版
// @remark 时间复杂度o(m*n),空间复杂度o(n)
func minDistance2(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	dp := make([]int, n+1)
	// 临时变量存储上一个元素，表示dp[i-1][j-1]
	pre := 0
	// 初始化
	for j := 0; j <= n; j++ {
		dp[j] = j
	}
	for i := 1; i <= m; i++ {
		dp[0], pre = i, dp[0]
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[j], pre = pre, dp[j]
			} else {
				dp[j], pre = min(min(dp[j], dp[j-1]), pre)+1, dp[j]
			}
		}
	}
	return dp[n]
}
