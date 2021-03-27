package main

import "fmt"

// 快乐数测试
// @remark 时间内复杂度O(logn),空间复杂度O(logn)
func IsHappyTest() {
	fmt.Println(isHappy(19))
}

// 快乐数
func isHappy(n int) bool {
	m := make(map[int]bool)
	for {
		if n == 1 {
			return true
		}
		// 有重复的了，说明不是快乐数
		if m[n] {
			return false
		}
		m[n] = true
		n = getSum(n)
	}
}

// 取数值各个位上的单数之和
func getSum(n int) int {
	sum := 0
	for n > 0 {
		m := n % 10
		sum += m * m
		n = n / 10
	}
	return sum
}
