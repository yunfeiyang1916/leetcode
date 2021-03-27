package main

import "fmt"

// 零钱兑换，给定不同面额的硬币 coins 和一个总金额 amount。计算可以凑成总金额所需的最少的硬币个数
func CoinChange() {
	coins := []int{2, 3, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

// 零钱兑换，给定不同面额的硬币 coins 和一个总金额 amount。计算可以凑成总金额所需的最少的硬币个数
// @remark	时间复杂度o(amount*n),空间复杂度o(amount)
func coinChange(coins []int, amount int) int {
	n := len(coins)
	if n == 0 || amount < 0 {
		return -1
	}
	if amount == 0 {
		return amount
	}
	// dp[i]表示构造金额i所需要的最少硬币个数
	dp := make([]int, amount+1)
	// 判断金额凑不出的小技巧：先初始化dp各个元素为amount + 1（代表不可能存在的情况），在遍历时如果金额凑不出则不更新，若最后结果仍然是amount + 1，则表示金额凑不出
	maxAmount := amount + 1
	for i := 1; i <= amount; i++ {
		dp[i] = maxAmount
	}
	// dp[0]置为0
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	// 表示未凑出
	if dp[amount] == maxAmount {
		return -1
	}
	return dp[amount]
}
