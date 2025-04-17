package main

/*
746. 使用最小花费爬楼梯
简单
提示
给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费
*/
func minCostClimbingStairs(cost []int) int {
	var a, b = 0, 0
	for i := 2; i <= len(cost); i++ {
		b, a = min(b+cost[i-1], a+cost[i-2]), b
	}
	return b
}
