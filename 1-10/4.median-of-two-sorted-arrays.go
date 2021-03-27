package main

import "fmt"

// 寻找两个正序数组的中位数
func FindMedianSortedArrays() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2, 5}
	fmt.Println(findMedianSortedArrays1(nums1, nums2))
}

// 寻找两个正序数组的中位数,暴力循环，合并两个数组
// @remark 时间复杂度： O(m+n) 空间复杂度：开辟了新的数组，所以空间复杂度是 O(m+n)
func findMedianSortedArrays1(nums1, nums2 []int) float64 {
	var (
		count1 = len(nums1)
		count2 = len(nums2)
		count  = count1 + count2
		// 商
		q   = count / 2
		arr = make([]int, 0, count)
		i   = 0
		j   = 0
		r   float64
	)
	if count == 0 {
		return 0
	}
	for {
		if i == count1 {
			if j < count2 {
				arr = append(arr, nums2[j:count2]...)
			}
			break
		}
		if j == count2 {
			if i < count1 {
				arr = append(arr, nums1[i:count1]...)
			}
			break
		}
		if nums1[i] <= nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else {
			arr = append(arr, nums2[j])
			j++
		}
	}
	fmt.Println(arr)
	// 说明总长度是奇数
	if count%2 == 1 {
		r = float64(arr[q])
	} else {
		r = float64(arr[q-1]+arr[q]) / 2
	}

	return r
}
