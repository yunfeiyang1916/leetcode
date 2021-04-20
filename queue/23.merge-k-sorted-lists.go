package main

import (
	"container/heap"
	"fmt"
)

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并K个升序链表测试
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
func MergeKListsTest() {
	lists := []*ListNode{
		buildList([]int{1, 4, 5}),
		buildList([]int{1, 3, 4}),
		buildList([]int{2, 6}),
	}
	res := mergeKLists(lists)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
	fmt.Println()
}

// 组装链表
func buildList(arr []int) *ListNode {
	var (
		dummy = &ListNode{}
		tail  = dummy
	)
	for _, v := range arr {
		node := &ListNode{
			Val: v,
		}
		tail.Next = node
		tail = tail.Next
	}
	return dummy.Next
}

// 合并K个升序链表,使用优先级队列（小堆序）
// 最多有n个结点，所以时间复杂度为O(NlogK),空间复杂度O(K)
func mergeKLists(lists []*ListNode) *ListNode {
	// 优先级队列（小堆序）
	p := NewPriorityQueue(false)
	heap.Init(&p)
	// 先将链表头入堆(也就是最小值入堆)
	for _, v := range lists {
		if v == nil {
			continue
		}
		heap.Push(&p, Item{Value: v, Priority: v.Val})
	}
	var (
		// 假头
		dummy = &ListNode{}
		// 尾指针
		tail = dummy
	)
	for p.Len() > 0 {
		// 弹出最小值
		min := heap.Pop(&p).(Item).Value.(*ListNode)
		tail.Next = &ListNode{Val: min.Val}
		tail = tail.Next
		if min.Next != nil {
			// 将下一个节点压入优先级队列
			heap.Push(&p, Item{Value: min.Next, Priority: min.Next.Val})
		}
	}
	return dummy.Next
}
