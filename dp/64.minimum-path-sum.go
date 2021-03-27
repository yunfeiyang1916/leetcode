package main

import (
	"fmt"
)

// 最小路径和
func MinPathSum() {
	arr := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	arr = [][]int{{1, 2, 5}, {3, 2, 1}}
	fmt.Println(minPathSum2(arr))
}

// 最小路径和
// @remark 时间复杂度o(m*n),空间复杂度o(m*n)
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	// 当机器人从左上角走到(i,j) 这个位置时，最小的路径和是 dp[i][j]
	dp := make([][]int, m)
	// 设置初始值，第一行及第一列的初始值均为路径对应的数字
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			dp[i][0] = grid[i][0]
			for j := 1; j < n; j++ {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	// 计算最小路径和
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// grid[i][j]表示网格中的值
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// 最小路径和优化版
// @remark 时间复杂度o(m*n),空间复杂度o(n)
func minPathSum2(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	// dp[i]表示走到该位置时的最小路径和
	// dp用于填充每一行的数据，上一行的值被使用过后在也用不到了，所以可以被覆盖，转移方程就是dp[j]=dp[j-1]+d[j]
	dp := make([]int, n)
	dp[0] = grid[0][0]
	// 初始化值，使用第一行的值填充
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		// 设置每行的第一列的值
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[n-1]
}
