package main

import (
	"container/heap"
	"fmt"
)

// 前 K 个高频元素测试
func TopKFrequentTest() {
	// 输入: nums = [1,1,1,2,2,3], k = 2
	// 输出: [1,2]
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

// 前 K 个高频元素,使用优先级队列，使用小堆
// 因为优先级队列的push和pop操作时间复杂度都是O(logK),所以总的时间复杂度为O(NlogK),空间复杂度为O(N)
func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k <= 0 || n < k {
		return nil
	}
	// 使用map存储数量
	var m = make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	// 使用优先级队列,使用小堆
	// 优先级为某个元素出现的次数
	p := NewPriorityQueue(false)
	for key, v := range m {
		heap.Push(&p, Item{Value: key, Priority: v})
		if p.Len() > k {
			// 弹出
			heap.Pop(&p)
		}
	}
	var result = make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(&p).(Item).Value.(int)
	}

	return result
}
