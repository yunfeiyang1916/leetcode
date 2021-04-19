package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 数组中最小的k个数测试，使用大堆
// 输入：[9, 3, 2, 8, 1, 7, 5, 6, 4], k = 4
// 输出：[1,2,3,4]
func GetLeastNumbersTest() {
	arr := []int{9, 3, 2, 8, 1, 7, 5, 6, 4}
	k := 4
	fmt.Println(getLeastNumbers(arr, k))
	fmt.Println(getLeastNumbers2(arr, k))
	fmt.Println(getLeastNumbersBySort(arr, k))
}

// 数组中最小的k个数，使用大堆
// 因为控制入堆的元素的数量为k，push的时间复杂度为O(logK),所以总的时间复杂度为O(NlogK)
func getLeastNumbers(arr []int, k int) []int {
	n := len(arr)
	if n == 0 || k <= 0 || n < k {
		return nil
	}
	// 构造堆结构,容量为k+1
	h := NewHeap(k + 1)
	for i := 0; i < n; i++ {
		h.Push(arr[i])
		if h.Size() > k {
			// 最大元素需要弹出
			h.Pop()
		}
	}
	// 去除最后一个元素
	return h.data[:k]
}

// 数组中最小的k个数，使用系统中的堆
// 因为控制入堆的元素的数量为k，push的时间复杂度为O(logK),所以总的时间复杂度为O(NlogK)
func getLeastNumbers2(arr []int, k int) []int {
	n := len(arr)
	if n == 0 || k <= 0 || n < k {
		return nil
	}
	// 构造堆结构,容量为k+1
	h := make(IntHeap, 0)
	heap.Init(&h)
	for i := 0; i < n; i++ {
		heap.Push(&h, arr[i])
		if h.Len() > k {
			// 最大元素需要弹出
			heap.Pop(&h)
		}
	}
	// 去除最后一个元素
	return h
}

// 数组中最小的k个数，使用快排
// 时间复杂度O(NlogN)
func getLeastNumbersBySort(arr []int, k int) []int {
	n := len(arr)
	if n == 0 || k <= 0 || n < k {
		return nil
	}
	sort.Ints(arr)
	return arr[:k]
}
