package main

// 环形链表，给定一个链表，判断链表中是否有环
// @remark 时间复杂度o(n),空间复杂度o(1)
func hasCycle(head *ListNode) bool {
	// 使用快慢指针
	var (
		fast = head
		slow = head
	)
	// 每次遍历时，fast移动2步，slow移动1步，如果是环形链表，则fast还会追上slow
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
