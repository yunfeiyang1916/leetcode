package main

// 反转链表2,反转部分链表
func ReverseBetweenTest() {
	linked := Constructor()
	linked.AddAtTail(1)
	linked.AddAtTail(2)
	linked.AddAtTail(3)
	linked.AddAtTail(4)
	linked.AddAtTail(5)
	reverseBetween(linked.head.Next, 2, 4)
}

// 反转链表2,反转部分链表，从left到right
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 借用虚拟头结点,此时链表变成0->1->2->3->4->5->NULL
	dummyHead := &ListNode{Val: 0}
	dummyHead.Next = head
	var (
		// 上一个结点
		pre *ListNode
		// 当前结点
		cur *ListNode
		// 下一个结点
		next *ListNode
		// 最左连接点
		leftLink *ListNode
		i        int
	)
	cur = dummyHead
	for cur != nil && i < left {
		leftLink = cur
		cur = cur.Next
		i++
	}
	// 此时i=left,leftLink=1,leftLink.next=2,cur=2
	for cur != nil && i <= right {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		i++
	}
	// 此时，pre=4,cur=5,leftLink=1,leftLink.next=2,注意leftLink.next并没有断开，仍然指向2
	if leftLink != nil && leftLink.Next != nil {
		if leftLink.Next != nil {
			// 此时cur为最右连接点
			leftLink.Next.Next = cur
		}
		leftLink.Next = pre
	}
	return dummyHead.Next
}
