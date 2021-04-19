package main

import "container/heap"

// 优先级队列元素
type Item struct {
	// 值，如果是纯整型，这里可以使用整型
	Value interface{}
	// 优先级
	Priority int
	// 在堆中的索引
	Index int
}

// 优先级队列
type PriorityQueue struct {
	// 数据
	Data []Item
	// 是否是大堆
	IsLarge bool
}

// 构造优先级队列
// @param	isLarge		是否是大堆
func NewPriorityQueue(isLarge bool) PriorityQueue {
	return PriorityQueue{
		IsLarge: isLarge,
	}
}

// 实现sort
// Len is the number of elements in the collection.
func (p PriorityQueue) Len() int {
	n := len(p.Data)
	return n
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p PriorityQueue) Less(i, j int) bool {
	if p.IsLarge {
		return p.Data[i].Priority > p.Data[j].Priority
	}
	return p.Data[i].Priority < p.Data[j].Priority
}

// Swap swaps the elements with indexes i and j.
func (p PriorityQueue) Swap(i, j int) {
	p.Data[i], p.Data[j] = p.Data[j], p.Data[i]
}

// 实现堆接口
func (p *PriorityQueue) Push(x interface{}) {
	n := p.Len()
	item := x.(Item)
	item.Index = n
	p.Data = append(p.Data, item)
}

// 弹出
func (p *PriorityQueue) Pop() interface{} {
	n := p.Len()
	item := p.Data[n-1]
	item.Index = -1
	p.Data = p.Data[:n-1]
	return item
}

// 更新队列元素中的值与优先级
func (p *PriorityQueue) Update(item *Item, value int, priority int) {
	item.Priority = priority
	item.Value = value
	heap.Fix(p, item.Index)
}
