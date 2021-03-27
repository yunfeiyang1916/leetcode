package main

import "fmt"

func main() {
	// 有效的括号
	// IsValidTest()
	// 大鱼吃小鱼
	// FishTest()
	// 找出数组中右边比我小的元素
	// FindRightSmallTest()
	// 下一个更大元素
	// NextGreaterElementTest()
	// 在循环数组中查找下一个更大的元素
	// NextGreaterElementsTest()
	// 查找字典序最小的 k 个数的子序列
	// FindSmallSeqTest()
	// 求柱状图中最大的矩形测试
	// LargestRectangleAreaTest()
	// 内存逃逸测试
	escapeTest()
}

// 分析内存逃逸
type A struct {
	s string
	n int
	i interface{}
}

func foo(s *string, n *int) *A {
	a := &A{}
	a.s = *s
	a.n = *n
	return a
}

func foo2(s ...interface{}) A {
	a := new(A)
	if len(s) == 0 {
		return *a
	}

	switch k := s[0].(type) {
	case string:
		a.s = k
	case *int:
		a.n = *k
	default:
		a.i = k
	}
	return *a
}

// 内存逃逸测试
func escapeTest() {
	s := "hello"
	n := 23
	a := foo(&s, &n)
	*a = foo2(&n)
	//b := a.s + " world"
	//c := b + "!"
	fmt.Println(a)

}
