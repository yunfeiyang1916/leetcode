package main

import (
	"fmt"
	"math"
)

// 最小栈测试
func MinStackTest() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}

// 最小栈,借助辅助栈存储最小值
type MinStack struct {
	// 栈
	stack []int
	// 栈顶为最小值的栈
	min []int
}

// 构造
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		// 最小栈里压入一个最大值
		min: []int{math.MaxInt32},
	}
}

// 压栈,如果当前值大于栈，则继续压入栈顶元素，否则压入当前值
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	min := this.GetMin()
	// 当前值小于栈顶元素，将当前值压入栈
	if val < min {
		this.min = append(this.min, val)
	} else {
		// 将最小值压入栈，
		this.min = append(this.min, min)
	}
}

// 出栈
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	// 辅助栈也出栈
	this.min = this.min[:len(this.min)-1]
}

// 返回栈顶元素
func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

// 获取最小值,辅助栈的栈顶就是最小值
func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
