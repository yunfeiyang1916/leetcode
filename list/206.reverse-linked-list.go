package main

// 反转链表
// 输入1->2->3->4>-5->NULL
// 输出5->4->3->2->1->NULL
func reverseList(head *ListNode) *ListNode {
	var (
		// 当前结点
		cur = head
		// 上一个结点
		pre *ListNode
		// 下一个结点
		next *ListNode
	)
	for cur != nil {
		next = cur.Next
		// 下一个结点指向上一个结点
		cur.Next = pre
		pre = cur
		// 继续循环
		cur = next
	}
	return pre
}
