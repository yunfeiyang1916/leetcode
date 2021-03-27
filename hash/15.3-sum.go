package main

import "sort"

// 三数之和
// @remark 时间复杂度O(N²),空间复杂度O(logN)
func threeSum(nums []int) [][]int {
	var (
		n      = len(nums)
		result = make([][]int, 0)
	)
	// 数量小于3，直接返回
	if n < 3 {
		return nil
	}
	// 排序
	sort.Ints(nums)
	// 找出a+b+c=0,使用双指针+排序
	// a=nums[i],b=nums[left],c=nums[right]
	for i := 0; i < n-2; i++ {
		// 排序后如果第一个元素大于0，说明后面的已经不满足条件了，直接返回结果
		if nums[i] > 0 {
			return result
		}
		// 去重,需要和上一次枚举的数不相同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 左指针
		left := i + 1
		// 右指针
		right := n - 1
		// 双指针相遇扔没找到，则说明没有与i坐标组成三数之和的值
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			// 三数之和大于0，则右指针需要向左移动
			if sum > 0 {
				right--
			} else if sum < 0 {
				// 三数之和小于0，则左指针需要向右移动
				left++
			} else {
				// 说明三数之和等于0
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 去重右指针
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// 去重左指针
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 找到结果了，则需要同时收缩左右指针
				left++
				right--
			}
		}
	}
	return result
}
