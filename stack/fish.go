package main

import "fmt"

// 大鱼吃小鱼测试
func FishTest() {
	size := []int{4, 3, 2, 1, 5}
	dir := []int{1, 1, 1, 1, 0}
	fmt.Println(fish(size, dir))
}

// 大鱼吃小鱼，拉勾算法题。返回剩余鱼的数量
// @param	size	鱼的大小
// @param	dir		鱼游动的方向，0表示向左游，1表示向右游
// @remark	时间复杂度O(N),空间复杂度O(N)
func fish(size, dir []int) int {
	n := len(size)
	// 小于等于1直接返回
	if n <= 1 {
		return n
	}
	var (
		// 0表示向左游，1表示向右游
		left  = 0
		right = 1
		// 栈的内容为size的下标
		stack = make([]int, 0, n)
	)
	for i := 0; i < n; i++ {
		curSize := size[i]
		curDir := dir[i]
		// 当前鱼是否被吃掉
		isEat := false
		// 两条鱼必须相对而游才会相遇，也就是第一条需要向右游，第二条需要向左游才行
		// 如果栈不为空并且栈顶元素向右游，当前鱼向左游，则大鱼需要吃掉小鱼
		for len(stack) > 0 && dir[stack[len(stack)-1]] == right && curDir == left {
			// 哪个鱼大吃掉哪个
			topSize := size[stack[len(stack)-1]]
			if topSize > curSize {
				// 吃掉当前鱼，不需要出栈了
				isEat = true
				break
			} else {
				// 吃掉栈顶的鱼，出栈,继续循环
				stack = stack[:len(stack)-1]
			}
		}
		// 当前鱼如果未被吃掉，则入栈
		if !isEat {
			stack = append(stack, i)
		}
	}
	return len(stack)
}
