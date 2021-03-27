package main

// 反转字符串
func reverseString(s []byte) {
	var (
		n = len(s)
		i = 0
		j = n - 1
	)
	for i < n/2 {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
