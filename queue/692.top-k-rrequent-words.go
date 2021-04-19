package main

import (
	"container/heap"
	"fmt"
)

// 前K个高频单词测试
func TopKFrequentWordsTest() {
	// 输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
	// 输出: ["i", "love"]
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequentWords(words, k))
}

// 前K个高频单词,使用优先级队列（借助小堆）
func topKFrequentWords(words []string, k int) []string {
	n := len(words)
	if n == 0 || k <= 0 || n < k {
		return nil
	}
	// 以字符串为键，出现次数为值的map
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}
	// 使用优先级队列（借助小堆）
	p := NewPriorityQueue(false)
	for key, v := range m {
		heap.Push(&p, Item{Value: key, Priority: v})
		if p.Len() > k {
			// 弹出栈顶最小值
			heap.Pop(&p)
		}
	}
	result := make([]string, k)
	for i := 0; i < k; i++ {
		// 每次都是弹出最小值,所以可以倒序赋值
		result[k-i-1] = heap.Pop(&p).(Item).Value.(string)
	}
	return result
}
