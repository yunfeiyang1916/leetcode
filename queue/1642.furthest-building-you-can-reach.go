package main

import (
	"container/heap"
	"fmt"
)

// 可以到达的最远建筑测试
// 输入：heights =[1,5,1,2,3,4,10000], bricks = 4, ladders = 1
// 输出：5
func FurthestBuildingTest() {
	heights := []int{1, 5, 1, 2, 3, 4, 10000}
	bricks := 4
	ladders := 1
	fmt.Println(furthestBuilding(heights, bricks, ladders))
}

// 可以到达的最远建筑,从低处往高处跳，借助砖块与梯子
// @param	heights		高度
// @param	bricks		砖块数量
// @param	ladders		梯子数量
// @remark 最差的情况下，我们需要把所有的高度差记录下来。在这种情况下，每个高度差都需要执行 push 操作。那么时间复杂度为 O(NlgN)，空间复杂度为O(N)。
func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	var (
		// 使用大堆序
		p = NewPriorityQueue(true)
		// 最远位置
		lastPos   = 0
		preHeight = heights[0]
		// 跳跃的高度差总和
		sum = 0
	)
	for i := 1; i < n; i++ {
		curHeight := heights[i]
		// 如果是从高处往低处跳，直接跳即可
		if preHeight >= curHeight {
			lastPos = i
		} else {
			// 从低处往高处跳，记录差值
			diff := curHeight - preHeight
			// 将高度差压入队列
			heap.Push(&p, Item{Priority: diff})
			sum += diff
			// 判断高度差是否大于砖块总和并且还有梯子
			for sum > bricks && ladders > 0 {
				// 使用梯子
				// 弹出队列最大高度差
				sum -= heap.Pop(&p).(Item).Priority
				// 梯子-1
				ladders--
			}
			// 经过一番处理后，还能不能跳到i这个位置
			if sum <= bricks {
				lastPos = i
			} else {
				// 结束
				break
			}
		}
		preHeight = curHeight
	}
	return lastPos
}
