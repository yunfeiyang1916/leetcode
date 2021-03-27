package main

// 有效的字母异位词,其实就是判断两个字符串的所有相同字符数量是否相等
func isAnagram(s, t string) bool {
	// 长度不相等则表明不是异位词
	if len(s) != len(t) {
		return false
	}
	// 总共有26个字母，则设置数组长度为26
	arr := make([]int, 26)
	// 因为字符a-z是顺序存储的，所以v-'a'肯定是落在0-25之间
	// 遍历s，将相同字符数量递增
	for _, v := range s {
		arr[v-'a']++
	}
	// 遍历t，将相同字符数量递减
	for _, v := range t {
		arr[v-'a']--
	}
	// 遍历字符数组，如果每个字符数量全为0表明是异位词
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

// 进阶版，如果s和t不是纯字母的，包含中文等字符，借用map处理
func isAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
