package main

import "fmt"

// 实现各种排序算法

func main() {
	var nums = []int{5, 2, 3, 1, 4}
	//bubbleSort(nums)
	//selectionSort(nums)
	nums = quickSort(nums)
	fmt.Println(nums)
}

// 1.冒泡排序，每次比较两个相邻的元素，如果前一个比后一个大，则交换位置
// 和选择排序的区别是，冒泡排序每次循环都把最大的元素放在最后面，选择排序每次循环都把最小的元素放在最前面
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		// 每次循环，都把最大的元素放在最后面
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 2.选择排序，将每次最小的元素放在最前面
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// 3.插入排序，从第一个元素开始，视为已排序部分,逐个将未排序元素插入到已排序部分的正确位置
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

// 4.快速排序，时间复杂度O(nlogn)，空间复杂度O(logn)
// 解题思路（分治法）
// 1.选择择基准值：从数组中选择一个元素作为基准（通常选第一个或中间元素）。
// 2.分区操作：将数组分为两部分，左边小于基准，右边大于基准。
// 3.递归排序：对左右子数组重复上述步骤，直到子数组长度为1。
func quickSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	// 选择基准值
	pivot := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < n; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	// 对左右子数组递归排序并合并
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}
