package main

import (
	"container/heap"
	"fmt"
)

// 查找和最小的K对数字
// 输入：[1,1,2],[1,2,3],k=10
// 输出：[[1,1],[1,1],[2,1],[1,2],[1,2],[2,2],[1,3],[1,3],[2,3]]
func KSmallestPairsTest() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{1, 2, 3}
	k := 10
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

// 查找和最小的K对数字,使用小堆序优先级队列
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 == 0 || n2 == 0 || k <= 0 {
		return nil
	}
	// 小堆序优先级队列
	p := NewPriorityQueue(false)
	heap.Init(&p)
	// 先将num1组装出最小的数组,因为都是有序数组，所以(num2[0],num1[0..n1-1])一定是最小值
	// 优先级队列中的元素值要存储下标
	for i := 0; i < n1; i++ {
		cur := []int{i, 0}
		heap.Push(&p, Item{Value: cur, Priority: nums1[cur[0]] + nums2[cur[1]]})
	}

	var res = make([][]int, 0)
	for i := 0; i < k && p.Len() > 0; i++ {
		// 最小值弹出
		min := heap.Pop(&p).(Item).Value.([]int)
		res = append(res, []int{nums1[min[0]], nums2[min[1]]})
		if min[1]+1 < n2 {
			// 继续入队
			cur := []int{min[0], min[1] + 1}
			heap.Push(&p, Item{Value: cur, Priority: nums1[cur[0]] + nums2[cur[1]]})
		}
	}
	return res
}
