package main

// 删除排序链表中的重复元素
// @remakr 时间复杂度o(n)，空间复杂度o(1)
func deleteDuplicates(head *ListNode) *ListNode {
	// 使用双指针
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 断开链表
	slow.Next = nil
	return head
}
