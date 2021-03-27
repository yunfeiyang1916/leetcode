package main

import "fmt"

// 找出数组中右边比我小的元素测试
func FindRightSmallTest() {
	nums := []int{1, 2, 4, 9, 4, 0, 5}
	fmt.Println(findRightSmall(nums))
	nums = []int{5, 6}
	fmt.Println(findRightLarge(nums))
	// nums = []int{6, 5}
	fmt.Println(findLeftLarge(nums))
	fmt.Println(findLeftSmall(nums))
}

// 找出数组中右边第一个比我小的元素,使用单调栈
// @remark 时间复杂度O(N),空间复杂度O(N)
func findRightSmall(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	res := make([]int, n)
	// 单调递增栈，栈中存储的是下标
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		x := nums[i]
		// 每个元素都向左遍历栈中的元素完成消除动作
		for len(stack) > 0 && nums[stack[len(stack)-1]] > x {
			// 被消除的值的下标就是右边第一个比我小的元素
			res[stack[len(stack)-1]] = i
			// 消除的时候，值更大的需要出栈
			stack = stack[:len(stack)-1]
		}
		// 剩下的入栈
		stack = append(stack, i)
	}
	// 栈中剩下的元素，由于没有人能消除他们，因此，只能将结果设置为-1。
	for len(stack) > 0 {
		res[stack[len(stack)-1]] = -1
		// 出栈
		stack = stack[:len(stack)-1]
	}
	return res
}

// 找出数组中右边第一个比我大的元素,使用单调栈
// @remark 时间复杂度O(N),空间复杂度O(N)
func findRightLarge(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	res := make([]int, n)
	// 单调栈，递减栈，栈中存储的是下标
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		x := nums[i]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < x {
			// 栈顶元素小于当前元素，则需要消除栈顶元素
			// 同时当前值是栈顶元素的右边第一个大值（因为栈顶元素被消除了）
			res[stack[len(stack)-1]] = i
			// 出栈
			stack = stack[:len(stack)-1]
		}
		// 入栈下标
		stack = append(stack, i)
	}
	// 栈中剩下的元素，由于没有人能消除它们，因此，只能将结果设置为-1
	if len(stack) > 0 {
		for i := 0; i < len(stack); i++ {
			res[stack[i]] = -1
		}
	}
	return res
}

// 找出数组中左边第一个比我大的元素,使用单调栈
// @remark 时间复杂度O(N),空间复杂度O(N)
func findLeftLarge(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	res := make([]int, n)
	// 单调栈，递减栈,栈中存的是下标
	stack := make([]int, 0)
	// 因为要找左边比我大的元素，所以需要倒序入栈
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < x {
			res[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		// 剩余元素入栈
		stack = append(stack, i)
	}
	if len(stack) > 0 {
		for _, v := range stack {
			res[v] = -1
		}
	}
	return res
}

// 找出数组中左边第一个比我小的元素,使用单调栈
// @remark 时间复杂度O(N),空间复杂度O(N)
func findLeftSmall(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	res := make([]int, n)
	// 单调栈，递增栈,栈中存的是下标
	stack := make([]int, 0)
	// 因为要找左边比我小的元素，所以需要倒序入栈
	for i := n - 1; i >= 0; i-- {
		x := nums[i]
		for len(stack) > 0 && nums[stack[len(stack)-1]] > x {
			res[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		// 剩余元素入栈
		stack = append(stack, i)
	}
	if len(stack) > 0 {
		for _, v := range stack {
			res[v] = -1
		}
	}
	return res
}
