package main

// 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 借用虚拟结点
	dummyHead := &ListNode{Val: -1}
	dummyHead.Next = head
	// 使用快慢指针，找到倒数第N个结点
	var (
		fast = head
		// 慢指针指向虚拟头结点，这样等快指针遍历结束，则慢指针指向倒数n个结点的前一个节点
		slow = dummyHead
	)
	// fast移动n次
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// fast移动结束，则slow正好移动到倒数第n个结点的前一个节点
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	return head
}
