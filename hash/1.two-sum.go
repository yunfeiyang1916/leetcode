package main

import "fmt"

func TwoSumTest() {
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum2(arr, target))
}

// 计算两数之和，两次暴力循环
// @remark	时间复杂度o(n²)，空间复杂度o(1)
func twoSum(arr []int, target int) []int {
	count := len(arr)
	// 数组长度小于1，直接返回
	if count <= 1 {
		return nil
	}
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 计算两数之和，借用map
// @remark	时间复杂度o(n),空间复杂度o(n)
func twoSum2(nums []int, target int) []int {
	if len(nums) <= 1 {
		return nil
	}
	// 以数组值为键，下标为值
	m := make(map[int]int)
	for k, v := range nums {
		if k2, ok := m[target-v]; ok {
			return []int{k2, k}
		}
		m[v] = k
	}
	return nil
}
