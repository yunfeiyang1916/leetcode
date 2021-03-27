package main

import "fmt"

// 搜索插入位置
func SearchInsertTest() {
	fmt.Println(searchInsert([]int{1, 2, 4, 5, 6}, 0))
}

// 搜索插入位置，给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// @remark 时间复杂度o(logn),空间复杂度o(1)
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1
	// 假设target在[left,right]左闭右闭区间处理，而不是[left,right)左闭右开
	for left <= right {
		// 中间结点
		middle := (right-left)/2 + left // 防止溢出，等同于(left+right)/2
		val := nums[middle]
		if target == val {
			return middle
		} else if target < val {
			right = middle - 1
		} else if target > val {
			left = middle + 1
		}
	}
	// 如果target落在所有元素之前，则区间为[0,-1],落的位置为right+1
	// 如果target等于数组中的某个值，则return middle
	// 如果target插入数组中的某个位置[left,right]，则结果会是left>right,落的位置则为right+1
	// 如果target落在所有元素之后，落的位置也是right+1
	return right + 1
}
