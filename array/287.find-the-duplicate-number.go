package main

// 寻找重复数
// 时间复杂度O(N),空间复杂度O(N)
func findduplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 借用map
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = struct{}{}
	}
	return 0
}

// 寻找重复数,因为整数的范围是1-n,而长度为n+1,则多了一个重复的元素，可以使用二分查找
// 时间复杂度O(NlogN),空间复杂度O(N)
func findduplicate2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var (
		left  = 0
		right = n - 1
	)
	// 使用左闭右闭区间[left,right]
	for left < right {
		mid := left + (right-left)/2
		// 小于等于mid的数量
		count := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				count++
			}
		}
		// 说明重复元素落在了左区间
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 寻找重复数，可以将重复元素之间的内容构成一个环，使用快慢指针查找
func findduplicate3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		fast = 2
		slow = 1
	)
	// 快指针一次走两步，慢指针一次走一步，当二者第一次相遇

}
