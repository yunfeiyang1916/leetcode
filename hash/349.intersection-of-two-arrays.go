package main

// 两个数组求交集
// @remark 时间复杂度o(m+n),空间复杂度o(m+n)
func intersection(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	var (
		result = make([]int, 0)
		// 用于判断是否有交集
		m1 = make(map[int]struct{})
		// 用于去重
		m2 = make(map[int]struct{})
	)
	for _, v := range nums1 {
		m1[v] = struct{}{}
	}
	for _, v := range nums2 {
		// 说明有交集
		if _, ok := m1[v]; ok {
			m2[v] = struct{}{}
		}
	}
	for k, _ := range m2 {
		result = append(result, k)
	}
	return result
}
