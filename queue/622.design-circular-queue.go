package main

// 设计循环队列
type MyCircularQueue struct {
	// 队首
	front int
	// 队尾
	rear int
	// 长度，表示已经使用的元素个数
	len int
	// 容量
	cap int
	// 数据，一个切片
	data []int
}

// 构造器，设置队列长度为 k
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		cap:  k,
		data: make([]int, k),
	}
}

// 入队
func (this *MyCircularQueue) EnQueue(value int) bool {
	// 队列是否满了
	if this.len == this.cap {
		return false
	}
	this.data[this.rear] = value
	// 队尾加1，注意需要取模
	this.rear = (this.rear + 1) % this.cap
	this.len++
	return true
}

// 从队首删除一个元素
func (this *MyCircularQueue) DeQueue() bool {
	// 空队列，不需要出队
	if this.len == 0 {
		return false
	}
	// 第一个元素出队,队首元素不需要置为0
	// 队首向后移动
	this.front = (this.front + 1) % this.cap
	this.len--
	return true
}

// 从队首获取元素，不删除队首元素。如果队列为空，返回 -1
func (this *MyCircularQueue) Front() int {
	if this.len == 0 {
		return -1
	}
	// 取出队首元素
	return this.data[this.front]
}

// 获取队尾元素。如果队列为空，返回 -1
func (this *MyCircularQueue) Rear() int {
	if this.len == 0 {
		return -1
	}
	// 注意：这里不能使用rear-1,因为rear=0时，就会出现-1
	rear := (this.rear - 1 + this.cap) % this.cap
	return this.data[rear]
}

// 队列是否为空
func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

// 队列是否满了
func (this *MyCircularQueue) IsFull() bool {
	return this.len == this.cap
}
