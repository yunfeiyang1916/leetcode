package main

/*
704. 二分查找
简单
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
*/

// 二分查找，时间复杂度 O(logN)，空间复杂度 O(1)
// 使用左闭右闭的区间，当 left > right 时，区间为空，返回 -1
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// 等价于 (left+right)/2，防止溢出
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
