package main

import "fmt"

// 最大子序和
func MaxSubArray() {
	//arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	arr := []int{-2, -1}
	fmt.Println(maxSubArray2(arr))
}

// 最大子序和
// @remark	时间复杂度o(n),空间复杂度o(n)
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	// dp[i]为以nums[i]结尾的最大子序和（注意并不一定是整个序列的最大子序和）
	dp := make([]int, n)
	// 最大子序和
	temp := 0
	// 初始值
	dp[0], temp = nums[0], nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		// 最大子序和
		temp = max(temp, dp[i])
	}
	return temp
}

// 最大子序和优化版
// @remark	时间复杂度o(n),空间复杂度o(1)
func maxSubArray2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	// 表示以nums[i]结尾的最大子序和（注意并不一定是整个序列的最大子序和）
	dp := 0
	// 最大子序和
	temp := 0
	// 初始值
	dp, temp = nums[0], nums[0]
	for i := 1; i < n; i++ {
		dp = max(dp+nums[i], nums[i])
		// 最大子序和
		temp = max(temp, dp)
	}
	return temp
}
