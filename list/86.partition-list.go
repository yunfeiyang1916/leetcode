package main

/*
86. 分隔链表
难度：中等

给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。
*/

func partition(head *ListNode, x int) *ListNode {
	// 定义两个链表，分别存储小于x的节点和大于等于x的节点
	var leftDummy, rightDummy = &ListNode{}, &ListNode{}
	var left, right = leftDummy, rightDummy
	for head != nil {
		if head.Val < x {
			leftDummy.Next = head
			leftDummy = leftDummy.Next
		} else {
			rightDummy.Next = head
			rightDummy = rightDummy.Next
		}
		head = head.Next
	}
	// 我们将 large 的 next 指针置空，这是因为当前节点复用的是原链表的节点，而其 next 指针可能指向一个小于 x 的节点，我们需要切断这个引用。
	rightDummy.Next = nil
	leftDummy.Next = right.Next
	return left.Next
}

func partitionTest() {
	nums := []int{1, 4, 3, 2, 5, 2}
	head := &ListNode{}
	for i := len(nums) - 1; i >= 0; i-- {
		head = &ListNode{Val: nums[i], Next: head}
	}
	partition(head, 3)
}
