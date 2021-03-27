package main

import "fmt"

// 在循环数组中查找下一个更大的元素测试
func NextGreaterElementsTest() {
	nums := []int{1, 2, 3, 4, 5}
	// 预期结果[2,3,4,5,-1]
	fmt.Println(nextGreaterElements(nums))
}

// 在循环数组中查找下一个更大的元素
// @remark 时间复杂度O(N)，空间复杂度O(N)
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	res := make([]int, n)
	// 先将结果全置为-1
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	// 单调栈，递增栈，栈存储的是下标
	stack := make([]int, 0)
	// 因为是循环数组，所以需要多循环一次
	for i := 0; i < 2*n-1; i++ {
		// 真实的下标需要取模
		index := i % n
		x := nums[index]
		// 栈不为空并且栈顶元素小于当前元素，则需要消除栈顶元素
		for len(stack) > 0 && nums[stack[len(stack)-1]] < x {
			res[stack[len(stack)-1]] = x
			// 出栈
			stack = stack[:len(stack)-1]
		}
		// 当前值入栈
		stack = append(stack, index)
	}
	return res
}
