package main

import "container/heap"

// 最低加油次数
// @param		target		目标距离
// @param		startFuel	起始邮箱
// @param		stations	加油站
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	var (
		n = len(stations)
		i = 0
		// 当前汽车的位置
		curPos = 0
		// 汽车中还有多少汽油
		fuelLeft = startFuel
		// 优先级队列，大堆序
		p = NewPriorityQueue(true)
		// 加油的次数
		res = 0
	)
	// 当前位置+剩余汽油数小于目的地距离了，需要加油了
	for curPos+fuelLeft < target {
		var (
			// 默认期望的下一站，初始设置为目标地址（将目标地址当成一个0升汽油的加油站）
			pos  = target
			fuel = 0
		)
		// 如果有位于target之前的站点，那么更新加油站的位置以及能加到副油箱的油量
		if i < n && stations[i][0] <= target {
			pos = stations[i][0]
			fuel = stations[i][1]
		}
		// 如果当前汽车的状态不能到达期望的下一站
		for curPos+fuelLeft < pos {
			// 副油箱为空，说明走不动下一站了，结束
			if p.Len() == 0 {
				return -1
			}
			// 从优先级队列里面取出最大的油桶加油
			maxFuel := heap.Pop(&p).(Item).Priority
			fuelLeft += maxFuel
			// 加油次数++
			res++
		}

		// 现在可以到达期望的下一站了，把消耗的汽油扣掉
		fuelLeft -= (pos - curPos)
		curPos = pos
		// 将当前站点汽油入队
		if fuel > 0 {
			heap.Push(&p, Item{Priority: fuel})
		}
		// 走过当前站点
		i++
	}
	// 剩余汽油是否能到达终点
	if curPos+fuelLeft >= target {
		return res
	}
	return -1
}
