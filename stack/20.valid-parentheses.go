package main

import "fmt"

// 有效的括号测试
func IsValidTest() {
	s := "()(())()"
	fmt.Println(isValidOne(s))
	fmt.Println(isValidOne2(s))
}

// 有效的括号,就是左右括号要配对,只有单个括号
// @remark 时间复杂度O(N),空间复杂度O(N)
func isValidOne(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	// 如果长度是奇数，说明肯定不匹配
	if n%2 != 0 {
		return false
	}
	// 借用栈
	stack := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		c := s[i]
		// 如果是左括号，则入栈
		if c == '(' {
			stack = append(stack, c)
		} else if c == ')' {
			m := len(stack)
			// 栈如果为空，则无法消除右括号，返回false
			if m == 0 {
				return false
			}
			// 因为栈里只有左括号，直接出栈即可消除
			stack = stack[:m-1]
		} else {
			return false
		}
	}
	// 栈为空则表明左右括号全消除完
	return len(stack) == 0
}

// 单括号优化版本
// @remark 时间复杂度O(N),空间复杂度O(1)
func isValidOne2(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	// 如果长度是奇数，说明不匹配
	if n%2 != 0 {
		return false
	}
	// 因为要入栈的内容一样，所以可以只记录入栈左括号的数量
	leftCount := 0
	for i := 0; i < n; i++ {
		c := s[i]
		// 如果是左括号，数量+1
		if c == '(' {
			leftCount++
		} else if c == ')' {
			// 栈如果为空，则无法消除右括号，返回false
			if leftCount <= 0 {
				return false
			}
			leftCount--
		} else {
			return false
		}
	}
	return leftCount == 0
}

// 多括号版本，比如(),{},[]
// @remark 时间复杂度O(N),空间复杂度O(N)
func isValid(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	// 如果长度为奇数，肯定不匹配
	if n%2 != 0 {
		return false
	}
	// 借用栈
	stack := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		c := s[i]
		// 栈长度
		m := len(stack)
		if c == '(' || c == '{' || c == '[' {
			// 入栈
			stack = append(stack, c)
		} else if c == ')' {
			// 栈为空，则无法消除，直接返回
			if m == 0 {
				return false
			}
			// 判栈顶元素是不是左括号
			if stack[m-1] != '(' {
				return false
			}
			// 出栈
			stack = stack[:m-1]
		} else if c == '}' {
			// 栈为空，则无法消除，直接返回
			if m == 0 {
				return false
			}
			// 判栈顶元素是不是左括号
			if stack[m-1] != '{' {
				return false
			}
			// 出栈
			stack = stack[:m-1]
		} else if c == ']' {
			// 栈为空，则无法消除，直接返回
			if m == 0 {
				return false
			}
			// 判栈顶元素是不是左括号
			if stack[m-1] != '[' {
				return false
			}
			// 出栈
			stack = stack[:m-1]
		} else {
			return false
		}
	}
	// 栈为空，则说明左右括号完全消除
	return len(stack) == 0
}

// 多括号优化版本,这里假设输入字符全为(),{},[]
// @remark 时间复杂度O(N),空间复杂度O(N+6)
func isValid2(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	// 长度为奇数，不满足
	if n%2 != 0 {
		return false
	}
	// 借助map,以右括号为键，左括号为值
	m := map[byte]byte{')': '(', '}': '{', ']': '['}
	// 栈
	stack := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		c := s[i]
		// 比较ascii值，如果大于0说明命中右括号，需要入栈
		if m[c] > 0 {
			stackLen := len(stack)
			// 栈为空或者栈顶元素不等于配对的左括号，不匹配
			if stackLen == 0 || stack[stackLen-1] != m[c] {
				return false
			}
			// 出栈
			stack = stack[:stackLen-1]
		} else {
			// 入栈
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
