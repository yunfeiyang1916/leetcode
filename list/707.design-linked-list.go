package main

import "fmt"

// 设计链表测试
func MyLinkedListTest() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)
	fmt.Println(linkedList.Get(1))
	linkedList.DeleteAtIndex(1)
	fmt.Println(linkedList.Get(1))
}

// 设计链表
type MyLinkedList struct {
	size int
	// 虚拟头结点
	head *ListNode
}

// 初始化
func Constructor() MyLinkedList {
	linkedList := MyLinkedList{}
	linkedList.size = 0
	linkedList.head = &ListNode{}
	return linkedList
}

// 获取链表中第 index 个结点的值。如果索引无效，则返回-1
func (this *MyLinkedList) Get(index int) int {
	if index >= this.size || index < 0 {
		return -1
	}
	current := this.head
	for i := 0; i <= index; i++ {
		current = current.Next
	}
	return current.Val
}

// addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的结点。插入后，新结点将成为链表的第一个结点。
func (this *MyLinkedList) AddAtHead(val int) {
	head := &ListNode{Val: val}
	head.Next = this.head.Next
	this.head.Next = head
	this.size++
}

// 将值为 val 的结点追加到链表的最后一个元素
func (this *MyLinkedList) AddAtTail(val int) {
	current := this.head
	for current.Next != nil {
		current = current.Next
	}
	tail := &ListNode{Val: val}
	current.Next = tail
	this.size++
}

// 在链表中的第 index 个结点之前添加值为 val  的结点。
// 如果 index 等于链表的长度，则该结点将附加到链表的末尾。
// 如果 index 大于链表长度，则不会插入结点。
// 如果index小于0，则在头部插入结点
func (this *MyLinkedList) AddAtIndex(index, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	current := this.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	node := &ListNode{Val: val}
	node.Next = current.Next
	current.Next = node
	this.size++
}

// 如果索引 index 有效，则删除链表中的第 index 个结点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size || index < 0 {
		return
	}
	current := this.head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	this.size--
}
