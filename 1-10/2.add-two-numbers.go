package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两数相加
func AddTwoNumbers() {
	l1 := initListNode(2, 4, 3)
	l2 := initListNode(5, 6, 4)
	PrintListNode(addTwoNumbers(l1, l2))
}

// 两数相加
// @remark 时间复杂度：O(max(m,n))O(max(m,n))，其中 m,nm,n 为两个链表的长度。我们要遍历两个链表的全部位置，而处理每个位置只需要 O(1)O(1) 的时间
//		   空间复杂度：O(max(m,n))O(max(m,n))。结果链表的长度最多为较长链表的长度 +1+1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		// 头指针
		head *ListNode
		tail *ListNode
		// 进位
		carry = 0
	)
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		// 进位需要除以10
		carry = sum / 10
		// 和只能用一位表示，需要模10
		sum = sum % 10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	// 如果进位数有值，则说明还需要进位
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

// 初始化链表
func initListNode(arr ...int) *ListNode {
	var (
		head *ListNode
		next *ListNode
	)
	for _, v := range arr {
		if head == nil {
			head = &ListNode{Val: v}
			next = head
		} else {
			next.Next = &ListNode{Val: v}
			next = next.Next
		}
	}
	return head
}

// 打印链表
func PrintListNode(head *ListNode) {
	node := head
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}
