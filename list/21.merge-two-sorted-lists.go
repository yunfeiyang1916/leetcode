package main

import (
	"fmt"
)

// 合并两个有序链表测试
func MergeTwoListsTest() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 3}
	l1.Next.Next = &ListNode{Val: 5}
	l2 := &ListNode{Val: 2}
	l2.Next = &ListNode{Val: 4}
	l2.Next.Next = &ListNode{Val: 6}

	r := mergeTwoLists(l2, l1)
	for r != nil {
		fmt.Printf("%d ", r.Val)
		r = r.Next
	}
}

// 合并两个有序链表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	// 借助虚拟头结点
	dummyHead := &ListNode{
		Val: 0,
	}
	tmp := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}
	return dummyHead.Next
}
