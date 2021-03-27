package main

// 获取链表的中间结点
// 输入1->2>3,输出2->3
// 输入1->2->3->4,输出3->4
// @remark 时间复杂度o(n),空间复杂度o(1)
func middleNode(head *ListNode) *ListNode {
	// 使用快慢指针
	var (
		fast = head
		slow = head
	)
	// 每次遍历时fast向后走2次，slow走1次，直到fast无法向后走两次
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 如果是奇数正好返回的是中间结点，如果是偶数，返回的是中间结点靠右结点
	return slow
}
