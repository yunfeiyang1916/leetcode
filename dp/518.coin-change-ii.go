package main

/*
518. 零钱兑换 II
中等
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。

题目数据保证结果符合 32 位带符号整数。
*/
func change(amount int, coins []int) int {
	// dp[i]表示凑出金额i的硬币组合数
	dp := make([]int, amount+1)
	// 初始化，0大小的硬币组合数为1，就是不装任何东西
	dp[0] = 1
	for _, v := range coins {
		for i := v; i <= amount; i++ {
			dp[i] += dp[i-v]
		}
	}
	return dp[amount]
}
