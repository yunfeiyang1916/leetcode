package main

import "fmt"

// 下一个更大的元素
func NextGreaterElementTest() {
	num1 := []int{4, 1, 2}
	num2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(num1, num2))
}

// 下一个更大的元素
func nextGreaterElement(nums1, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 == 0 || n2 == 0 {
		return nil
	}
	m := make(map[int]int)
	// 单调栈，使用递增栈，栈存储的是值
	stack := make([]int, 0)
	for i := 0; i < n2; i++ {
		x := nums2[i]
		// 栈不为空并且栈顶元素小于当前值，则需要消减
		for len(stack) > 0 && stack[len(stack)-1] < x {
			m[stack[len(stack)-1]] = x
			// 出栈
			stack = stack[:len(stack)-1]
		}
		// 入栈
		stack = append(stack, x)
	}
	// 组装结果
	res := make([]int, n1)
	for i := 0; i < n1; i++ {
		res[i] = m[nums1[i]]
		// 如果map不存在，则说明该元素右侧没有最大元素了，置为-1
		if res[i] == 0 {
			res[i] = -1
		}
	}
	return res
}
