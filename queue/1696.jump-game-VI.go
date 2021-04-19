package main

import (
	"fmt"
	"math"
)

// 跳跃游戏或捡金币游戏测试，使用动态规划+滑动窗口+单调队列
// 输入nums = [1,-1,-2,4,-7,3], k = 2，输出7
func MaxResultTest() {
	nums := []int{1, -1, -2, 4, -7, 3}
	nums = []int{100, -1, -100, -1, 100}
	// 输出0
	nums = []int{100, -100, -300, -300, -300, -100, 100}
	k := 2
	k = 4
	//
	nums = []int{40, 30, -100, -100, -10, -7, -3, -3}
	k = 2
	fmt.Println(maxResultByDP(nums, k))
	fmt.Println(maxResult(nums, k))
	fmt.Println(maxResult2(nums, k))
}

// 跳跃游戏，只使用动态规划
// 时间复杂度O(kn)，容易超时
func maxResultByDP(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	// 初始状态
	dp[0] = nums[0]
	// 公式 dp[i]=max(dp[i-k],..dp[i-2],dp[i-1])+nums[i]
	for i := 1; i < n; i++ {
		var (
			sub     = i - k
			maxTemp = math.MinInt32
		)
		// 差值小于0，则设置为0
		if sub < 0 {
			sub = 0
		}
		for ; sub < i; sub++ {
			maxTemp = max(maxTemp, dp[sub])
		}
		dp[i] = nums[i] + maxTemp
	}
	return dp[n-1]
}

// 跳跃游戏或捡金币游戏，使用动态规划+滑动窗口+单调队列
func maxResult(nums []int, k int) int {
	n := len(nums)
	// 处理边界
	if n == 0 || k <= 0 {
		return 0
	}
	// 动态规划，公式：dp[i]=max(dp[i-k],..dp[i-2],dp[i-1])+nums[i]
	dp := make([]int, n)
	// 单调递减队列，注意并不是严格递减
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		// 在取最大值之前，需要保证单调队列中都是有效值
		// 也就是都在区间里面的值
		// 对于i而言,[i-k,i-1]都可以调到i
		// 最远的距离为i-(i-k)=k,因此超出这个范围的需要出队
		if i-k > 0 {
			// 如果队首元素等于边界元素，则出队队首
			if len(queue) > 0 && queue[0] == dp[i-k-1] {
				queue = queue[1:]
			}
		}
		// 从单调队列取得最大值
		maxTemp := 0
		// 队列为空，说明是第一次循环
		if len(queue) > 0 {
			maxTemp = queue[0]
		}
		dp[i] = maxTemp + nums[i]
		// 入队的时候，保持单调队列有序
		for len(queue) > 0 && queue[len(queue)-1] < dp[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, dp[i])
	}
	return dp[n-1]
}

// 跳跃游戏或捡金币游戏，使用动态规划+滑动窗口+单调队列
func maxResult2(nums []int, k int) int {
	n := len(nums)
	// 处理边界
	if n == 0 || k <= 0 {
		return 0
	}
	dp := make([]int, n)
	// 借用递减单调队列
	queue := make([]int, 0)
	// 初始状态
	dp[0] = nums[0]
	queue = append(queue, dp[0])
	for i := 1; i < n; i++ {
		// 凑满滑动窗口了，则需要出队首元素
		// 因为循环是从1开始的
		if i-k > 0 {
			if dp[i-k-1] == queue[0] {
				queue = queue[1:]
			}
		}
		// 队首为[i-1,i-k]之间的最大元素
		dp[i] = nums[i] + queue[0]
		// 保证单调队列有序
		for len(queue) > 0 && queue[len(queue)-1] < dp[i] {
			// 出队
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, dp[i])
	}
	return dp[n-1]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
