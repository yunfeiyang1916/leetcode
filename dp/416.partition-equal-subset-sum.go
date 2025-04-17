package main

/*
416. 分割等和子集
中等
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

类似01背包问题，但是要求每个子集的元素和相等，所以需要使用动态规划解决。

时间复杂度：O(n*target)，空间复杂度：O(target)
*/
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// 如果数组元素和为奇数，则无法平分成两个子集
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	// dp[i]表示能否凑到目标值为i的子集
	dp := make([]int, target+1)
	for _, v := range nums {
		// 使用倒序遍历，保证v只被添加一次
		for j := target; j >= v; j-- {
			if dp[j] < dp[j-v]+v {
				dp[j] = max(dp[j], dp[j-v]+v)
			}
		}
	}
	return dp[target] == target
}
