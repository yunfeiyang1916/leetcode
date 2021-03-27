package main

import (
	"fmt"
	"math"
)

// 求长度最小的子数组
func MinSubArrayLenTest() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 7, 8, 4, 3}))
}

// 求长度最小的子数组，使用滑动窗口的方式
// 时间复杂度o(n),空间复杂度o(1)
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var (
		result = math.MaxInt32
		// 起始指针
		start int
		// 结束指针
		end int
		// 和
		sum int
	)
	for end < n {
		sum += nums[end]
		for sum >= target {
			result = min(result, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if result == math.MaxInt32 {
		return 0
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
