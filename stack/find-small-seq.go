package main

import "fmt"

// 查找字典序最小的 k 个数的子序列
func FindSmallSeqTest() {
	// 输入：[9, 2, 4, 5, 1, 2, 3, 0], k = 3，输出：[1,2,3]
	nums := []int{9, 2, 4, 5, 1, 2, 3, 0}
	fmt.Println(findSmallSeq(nums, 3))
}

// 查找字典序最小的 k 个数的子序列
func findSmallSeq(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || n < k {
		return nil
	}
	// 单调栈，递增栈，存储值
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		x := nums[i]
		// 剩余多少元素
		left := n - i
		// 如果栈不为空，并且栈的长度+剩余元素大于k，并且栈顶元素小于当前元素，则需要出栈
		for len(stack) > 0 && (len(stack)+left) > k && stack[len(stack)-1] > x {
			// 消除栈顶元素
			stack = stack[:len(stack)-1]
		}
		// 保证每个元素都能入栈
		stack = append(stack, x)
	}
	// 如果数组是一个递增数组，则单调栈的长度会超过k的长度，需要截取栈尾元素
	if len(stack) > k {
		stack = stack[:k]
	}
	return stack
}
