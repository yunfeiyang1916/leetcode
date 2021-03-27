package main

// 删除排序链表中的重复元素
// @remakr 时间复杂度o(n)，空间复杂度o(1)
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	// 不会删除头结点的
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
