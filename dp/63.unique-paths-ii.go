package main

/*
63. 不同路径 II
中等
提示
给定一个 m x n 的整数数组 grid。一个机器人初始位于 左上角（即 grid[0][0]）。机器人尝试移动到 右下角（即 grid[m - 1][n - 1]）。机器人每次只能向下或者向右移动一步。

网格中的障碍物和空位置分别用 1 和 0 来表示。机器人的移动路径中不能包含 任何 有障碍物的方格。

返回机器人能够到达右下角的不同路径数量。

测试用例保证答案小于等于 2 * 109。
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// 如果在起点或终点有障碍物，则直接返回0
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	dp := make([][]int, m)
	// 初始化
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if i == 0 {
			// 有障碍物，当前行都不用计算了，退出循环了
			for j := 0; j < n && obstacleGrid[i][j] == 0; j++ {
				dp[i][j] = 1
			}
			continue
		}
	}
	// 初始化第一列，需要单独处理
	for i := 0; i < m; i++ {
		// 有障碍物，当前列都不用计算了，退出循环了
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func UniquePathsWithObstaclesTest() {
	obstacleGrid := [][]int{
		{0, 0},
		{1, 1},
		{0, 0},
	}
	uniquePathsWithObstacles(obstacleGrid)
}
