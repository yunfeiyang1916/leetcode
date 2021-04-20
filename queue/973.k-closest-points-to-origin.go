package main

import "container/heap"

// 最接近原点的k个点，使用大堆弹出最大值
// 时间复杂度O(NlogK),空间复杂度O(K)
func kClosest(points [][]int, k int) [][]int {
	n := len(points)
	if n == 0 || k <= 0 || n < k {
		return nil
	}
	// 大堆序优先级队列
	p := NewPriorityQueue(true)
	heap.Init(&p)
	for i := 0; i < n; i++ {
		// 计算距离的平方，a²+b²=c²
		cur := points[i]
		dist := cur[0]*cur[0] + cur[1]*cur[1]
		heap.Push(&p, Item{Priority: dist, Value: cur})
		if p.Len() > k {
			// 将最大值弹出
			heap.Pop(&p)
		}
	}
	res := make([][]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(&p).(Item).Value.([]int)
	}
	return res
}
