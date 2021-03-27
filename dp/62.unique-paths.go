package main

import "fmt"

// 最多路径数,到达m✖️n的位置共有多少种路径
func UniquePaths() {
	fmt.Println(uniquePaths2(4, 4))
}

// 最多路径数,到达m✖️n的位置共有多少种路径
// @remark 时间复杂度o(m*n),空间复杂度o(m*n)
func uniquePaths(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	// 当机器人从左上角走到(i,j) 这个位置时，一共有 dp[i][j] 种路径
	dp := make([][]int, m)
	// 设置初始值
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			for j := 0; j < n; j++ {
				dp[i][j] = 1
			}
			continue
		}
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 最多路径数优化版,到达m✖️n的位置共有多少种路径
// @remark 时间复杂度o(m*n),空间复杂度o(n)
func uniquePaths2(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	// 当机器人从左上角走到(i,j) 这个位置时，一共有 dp[i][j] 种路径
	// dp用于填充每一行的数据，上一行的值被使用过后在也用不到了，所以可以被覆盖，转移方程就是dp[j]=dp[j-1]+d[j]
	dp := make([]int, n)
	// 初始值，第一行的初始值均为1
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	// 公式dp[j]=dp[j-1]+dp[j]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}
