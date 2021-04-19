package main

import "fmt"

// 找出数组中右边比我小的元素测试
func FindRightSmallTest() {
	nums := []int{1, 2, 4, 9, 4, 0, 5}
	// 输出[5,5,5,4,5,-1,-1]
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
		return nil
	}
	// 单调递增栈,栈中存的是下标
	stack := make([]int, 0)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		cur := nums[i]
		// 需要保持递增栈的单调性，遇到小的元素需要弹出栈顶元素，此时表明栈顶元素遇到了右侧第一个比我小的元素了
		for len(stack) > 0 && nums[stack[len(stack)-1]] > cur {
			res[stack[len(stack)-1]] = i
			// 出栈
			stack = stack[:len(stack)-1]
		}
		// 当前元素索引入栈
		stack = append(stack, i)
	}
	// 剩下未出栈的元素未找到右边第一个比我小的元素，置为-1
	for len(stack) > 0 {
		res[stack[len(stack)-1]] = -1
		//	出栈
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
	// 使用单调递减栈，存储下标
	stack := make([]int, 0)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		cur := nums[i]
		// 栈顶元素小于当前元素，说明栈顶元素遇到了第一个比我大的元素了，栈顶出栈
		for len(stack) > 0 && nums[stack[len(stack)-1]] < cur {
			res[stack[len(stack)-1]] = i
			// 出栈
			stack = stack[:len(stack)-1]
		}
		// 当前元素下标入栈
		stack = append(stack, i)
	}
	// 剩余元素未找到，赋值-1
	for len(stack) > 0 {
		res[stack[len(stack)-1]] = -1
		// 出栈
		stack = stack[:len(stack)-1]
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
