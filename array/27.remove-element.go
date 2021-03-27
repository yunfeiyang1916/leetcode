package main

// 移除元素，使用双指针法
// @remark 时间复杂度o(n),空间复杂度o(1)
func removeElement(nums []int, val int) int {
	// 慢指针
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}
