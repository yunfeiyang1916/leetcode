package main

import "fmt"

// 打家劫舍，在不触动警报装置的情况下，一夜之内能够偷窃到的最高金额
func Rob() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(rob2(arr))
}

// 打家劫舍，在不触动警报装置的情况下，一夜之内能够偷窃到的最高金额
// @remark	时间复杂度o(n),空间复杂度o(n)
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	// dp[i]表示前i个房间偷到的最大钱数
	dp := make([]int, n)
	// 初始化
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

// 打家劫舍优化版，在不触动警报装置的情况下，一夜之内能够偷窃到的最高金额
// @remark	时间复杂度o(n),空间复杂度o(1)
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	} else if n == 2 {
		return max(nums[0], nums[1])
	}
	// 使用三个局部变量表示dp[i-2] dp[i-1] dp[i]
	var dp2, dp1, dp int
	dp1 = max(nums[0], nums[1])
	dp2 = nums[0]
	for i := 2; i < n; i++ {
		dp = max(dp2+nums[i], dp1)
		// 交换值
		dp2, dp1 = dp1, dp
	}
	return dp
}
