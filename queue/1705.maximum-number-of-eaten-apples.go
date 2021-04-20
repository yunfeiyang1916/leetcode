package main

import (
	"container/heap"
	"fmt"
)

// 吃苹果的最大数目测试
// 输入：apples = [2,1,10], days = [2,10,1]
// 输出：4
func EatenApplesTest() {
	apples := []int{2, 1, 10}
	days := []int{2, 10, 1}
	fmt.Println(eatenApples(apples, days))
}

// 吃苹果的最大数目
// @param	apples		每天长出的苹果数量
// @param	days		苹果坏掉的时间
func eatenApples(apples []int, days []int) int {
	n := len(apples)
	if n == 0 {
		return 0
	}
	var (
		// 优先级队列，使用小堆序，将快过期的数据放到根
		p = NewPriorityQueue(false)
		// 天数
		i   = 1
		res = 0
	)
	for i <= n || p.Len() > 0 {
		// 第i天得到 num 个苹果
		// 会在 bad 那天坏掉
		if i <= n {
			num := apples[i-1]
			bad := days[i-1] + i
			if num > 0 {
				heap.Push(&p, Item{Value: num, Priority: bad})
			}
		}
		// 弹出
		for p.Len() > 0 {
			cur := heap.Pop(&p).(Item)
			// 过期的直接扔掉
			if cur.Priority <= i {
				continue
			}
			// 吃掉一个
			num := cur.Value.(int)
			res++
			num--
			if num > 0 {
				cur.Value = num
				// 在加入队列
				heap.Push(&p, cur)
			}
			break
		}
		i++
	}
	return res
}
