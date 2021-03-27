package main

import "fmt"

// 爬楼梯
func ClimbStairs() {
	// fmt.Println(climbStairs(3))
	fmt.Println(climbStairs2(3))
}

// 爬楼梯
// @remark 时间复杂度o(n),空间复杂度o(n)
func climbStairs(n int) int {
	// 1和2是初始值
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	// 设置初始值
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 爬楼梯优化版
// @remark 时间复杂度o(n),空间复杂度o(1)
func climbStairs2(n int) int {
	// 1和2是初始值
	if n <= 2 {
		return n
	}
	// 借用三个局部变量表示dp[1] db[2] dp[i]
	dp1, dp2, r := 1, 2, 0
	for i := 3; i <= n; i++ {
		// dp1表示dp[i-2] dp2表示dp[i-1] r表示dp[i]
		r = dp1 + dp2
		// 保存下一轮dp[i-2] dp[i-1],循环利用
		dp1 = dp2
		dp2 = r
	}
	return r
}
