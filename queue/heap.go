package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

// 大堆，大堆的根是最大值，小堆的根是最小值。
// 堆是一个完全二叉树的结构
// i结点的父结点下标(i-1)/2
// i结点的左子结点2i+1
// i结点的右子结点2i+2
type Heap struct {
	// 数据
	data []int
	// 已使用的长度
	len int
}

func NewHeap(n int) *Heap {
	return &Heap{
		data: make([]int, n),
		len:  0,
	}
}

// 压入
func (h *Heap) Push(value int) {
	// 先把元素追加到数组尾巴上，然后在上浮
	h.data[h.len] = value
	h.swim(h.len)
	h.len++
}

// 弹出堆顶元素
func (h *Heap) Pop() int {
	if h.len == 0 {
		return 0
	}
	res := h.data[0]
	h.len--
	// 将len-1放到堆顶，然后执行下沉操作
	h.data[0] = h.data[h.len]
	h.data[h.len] = 0
	h.sink(0)
	return res
}

// 索引i的值上浮
// 时间复杂度为树的高度O(logN)
func (h *Heap) swim(i int) {
	var (
		t = h.data[i]
		// 父结点下标
		parentIndex = 0
	)
	// 是否还存在父结点
	for i > 0 {
		// 父结点下标
		parentIndex = (i - 1) / 2
		// 如果父结点大于等于t，说明已经满足大堆了，结束循环
		if h.data[parentIndex] >= t {
			break
		}
		// 父结点值比t值小，则向下移动父结点的值
		h.data[i] = h.data[parentIndex]
		i = parentIndex
	}
	// 如果i的值上浮了，则需要重新赋值
	h.data[i] = t
}

// 下沉
func (h *Heap) sink(i int) {
	var (
		t = h.data[i]
	)
	for {
		// 找到i结点的左子结点
		j := 2*i + 1
		if j >= h.len {
			break
		}
		// j<h.len-1 判断是否有右子结点，如果有，并且右子结点更大，则j指向右子结点
		if j < h.len-1 && h.data[j] < h.data[j+1] {
			j++
		}
		// 如果子结点比t大，则t还需要下沉
		if h.data[j] > t {
			h.data[i] = h.data[j]
			i = j
		} else {
			// 找到t的位置了，此时t是大于所有的子结点的
			break
		}
	}
	// 将t放在找到的位置那里
	h.data[i] = t
}

// 堆大小
func (h *Heap) Size() int {
	return h.len
}

func (h *Heap) String() string {
	sb := strings.Builder{}
	for i := 0; i < h.len; i++ {
		sb.WriteString(strconv.Itoa(h.data[i]) + " ")
	}
	return sb.String()
}

// 堆测试
func HeapTest() {
	h := NewHeap(20)
	h.Push(9)
	h.Push(5)
	h.Push(6)
	h.Push(3)
	h.Push(4)
	h.Push(1)
	fmt.Println(h)
	// 系统堆
	h2 := IntHeap{3, 4, 1, 9, 5, 6}
	heap.Init(&h2)
	fmt.Println(h2)
	fmt.Println("出栈一次")
	h.Pop()
	heap.Pop(&h2)
	fmt.Println(h)
	fmt.Println(h2)
}

// 使用系统堆
type IntHeap []int

// 实现sort
func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 实现堆接口
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// 要弹出栈顶
func (h *IntHeap) Pop() interface{} {
	res := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return res
}
