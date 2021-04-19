package main

// 环形链表2,给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null
func detectCycle(head *ListNode) *ListNode {
	// 定义快慢指针，快指针一次走两步，慢指针一次走一步，等快指针追上慢指针说明是环形
	var (
		fast = head
		slow = head
	)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 相遇了说明是环形链表
		if fast == slow {
			// 继续从头遍历，当index与slow相遇时，则相遇处即为环形入口
			index := head
			for index != slow {
				index = index.Next
				slow = slow.Next
			}
			return index
		}
	}
	return nil
}
