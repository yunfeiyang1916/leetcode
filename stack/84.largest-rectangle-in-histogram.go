package main

import "fmt"

// 求柱状图中最大的矩形
func LargestRectangleAreaTest() {
	// 输入[2,1,5,6,2,3]，输出10
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
	fmt.Println(largestRectangleArea1(heights))
	fmt.Println(largestRectangleArea2(heights))
}

// 柱状图中最大的矩形,暴力解法
// @remark 时间复杂度O(N²)，空间复杂度O(1)
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	res := 0
	for i := 0; i < n; i++ {
		cur := heights[i]
		// 向左找到最后一个大于等于当前元素的下标
		left := i - 1
		for left >= 0 && heights[left] >= cur {
			left--
		}
		// 向右找到最后一个大于等于当前元素的下标
		right := i + 1
		for right < n && heights[right] >= cur {
			right++
		}
		width := right - left - 1
		res = max(res, cur*width)
	}
	return res
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

// 柱状图中最大的矩形优化版1,先找出数组中左边第一个比我小的元素索引下标，再找出数组中右边第一个比我小的元素索引下标
// @remark 时间复杂度O(N),空间复杂度O(N)
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	res := 0
	// 先找出数组中左边第一个比我小的元素索引下标
	leftSmall := findLeftSmall(heights)
	// 再找出数组中右边第一个比我小的元素索引下标
	rightSmall := findRightSmall(heights)
	for i := 0; i < n; i++ {
		// 左边比我小的位置
		leftPos := leftSmall[i]
		// 右边比我小的位置，如果右边没有比我小的位置，说明该元素可以扩散到最右边
		rightPos := rightSmall[i]
		if rightPos == -1 {
			rightPos = n
		}
		// 现在我们确定区间(leftPos, rightPos)
		// 注意两边都是开区间。在这个区间里面，所有的数肯定都是 >= heights[i]的。
		// 那么底部的宽度就是
		width := rightPos - leftPos - 1
		if width == 0 {
			width = 1
		}
		res = max(res, heights[i]*width)
	}
	return res
}

// 柱状图中最大的矩形优化版2,只用一遍循环
// @remark 时间复杂度O(N),空间复杂度O(N)
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	res := 0
	// 单调递增栈
	stack := make([]int, 0)
	// 注意，这里我们取到了i == n
	// 按理说，不应该取到i == n的。但是这时候，主要是为了处理这种数组
	// A = [1, 2, 3]
	// 没有任何元素会出栈。
	// 那么最后我们用一个0元素，把所有的元素都削出栈。
	// 这样代码就可以统一处理掉。
	for i := 0; i <= n; i++ {
		cur := -1
		if i < n {
			cur = heights[i]
		}
		// 栈不为空并且栈顶元素大于当前值，说明栈顶元素的下一个更小值就是当前值，栈顶元素出栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] > cur {
			// 因为是递增栈，则说明栈中剩余元素的值都比栈顶小
			height := heights[stack[len(stack)-1]]
			// 出栈
			stack = stack[:len(stack)-1]
			// 找到右侧第一个小于栈顶元素的值的索引了
			rightPos := i
			// 左侧第一个小于栈顶的元素就是栈顶的上一个元素
			leftPos := -1
			if len(stack) > 0 {
				leftPos = stack[len(stack)-1]
			}
			width := rightPos - leftPos - 1
			res = max(res, height*width)
		}
		// 当前元素下标入栈
		stack = append(stack, i)
	}
	return res
}
