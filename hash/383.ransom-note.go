package main

// 赎金信，求字符A是否在字符串B中存在
func canConstruct(ransomNote string, magazine string) bool {
	// 因为题目假设都为小写字母，则可以用26个字母数组表示
	// 因为字符a-z是顺序存储的，所以v-'a'肯定是落在0-25之间,数组的值为该字符出现的次数
	arr := make([]int, 26)
	for _, v := range magazine {
		arr[v-'a']++
	}
	for _, v := range ransomNote {
		arr[v-'a']--
		if arr[v-'a'] < 0 {
			return false
		}
	}
	return true
}
