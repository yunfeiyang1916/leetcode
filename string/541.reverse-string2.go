package main

import "fmt"

// 反转字符串2测试
func ReverseStrTest() {
	// 输出bacdfeg
	fmt.Println(reverseStr("abcdefg", 2))
}

// 反转字符串2，每隔 2k 个字符的前 k 个字符进行反转
func reverseStr(s string, k int) string {
	n := len(s)
	if n == 0 {
		return s
	}
	buf := []byte(s)
	for i := 0; i < n; i += 2 * k {
		// 每隔2k个字符的前k个字符反转或者剩余字符数小于2k但大于等于k个也反转前k个字符
		if i+k <= n {
			// 反转字符串
			reverse(buf, i, i+k)
			continue
		}
		// 走到这说明剩余字符数量小于k个，则剩余字符都需要反转
		reverse(buf, i, n)
	}
	return string(buf)
}

// 反转，左闭右开区间
func reverse(s []byte, start, end int) {
	var (
		i   = start
		j   = end - 1
		mid = (end-start)/2 + start // 防止溢出
	)
	for i < mid {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
