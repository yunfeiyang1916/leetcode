package main

// 删除排序链表中的重复元素2
func deleteDuplicates2(head *ListNode) *ListNode {
	// 因为可能要删除头结点，所以需要虚拟一个头结点
	dummyHead := &ListNode{Next: head}
	// 当前结点
	current := dummyHead
	// 记录已经删除的值
	var val int
	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			val = current.Next.Val
			current.Next = current.Next.Next.Next
			for current.Next != nil && current.Next.Val == val {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}
	}
	return dummyHead.Next
}
