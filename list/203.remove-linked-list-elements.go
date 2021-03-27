package main

// 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
	// 因为可能会删除头结点，所以需要虚拟一个头结点
	dummyHead := &ListNode{Next: head}
	// 当前结点
	current := dummyHead
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return dummyHead.Next
}
