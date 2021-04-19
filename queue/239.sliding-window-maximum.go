package main

import "fmt"

// 滑动窗口最大值测试
// 输入：nums = [1,3,-1,-3,5,3], k = 3
// 输出：[3,3,5,5]
func MaxSlidingWindowTest() {
	nums := []int{1, 3, -1, -3, 5, 3}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}

// 滑动窗口最大值,使用单调递减队列
// @remark 时间复杂度O(N)，空间复杂度O(K),因为最多只有k个元素入队、出队
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return nil
	}
	// 单调队列
	queue := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		// 入队，需要保持队列为单调递减的,重复元素需要入队
		for len(queue) > 0 && queue[len(queue)-1] < cur {
			// 队尾出队
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, cur)
		// 凑满滑动窗口了
		if (i + 1) >= k {
			// 此时队首即为该区间的最大的元素
			max := queue[0]
			res = append(res, max)
			// 当队首元素的值等于滑动窗口起始值时需要出队
			if max == nums[i+1-k] {
				queue = queue[1:]
			}
		}
	}
	return res
}
