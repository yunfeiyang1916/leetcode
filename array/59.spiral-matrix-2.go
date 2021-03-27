package main

import "fmt"

// 生成螺旋矩阵2
func GenerateMatrixTest() {
	fmt.Println(generateMatrix(3))
}

// 生成螺旋矩阵 II
func generateMatrix(n int) [][]int {
	// 定义左右下上边界
	var (
		left   = 0
		right  = n - 1
		top    = 0
		bottom = n - 1
		result = make([][]int, 0, n)
		// 初始值
		start = 1
		end   = n * n
	)
	// 初始化二维数组
	for i := 0; i < n; i++ {
		arr := make([]int, n)
		result = append(result, arr)
	}
	for start <= end {
		// 先构造从左到右
		for i := left; i <= right; i++ {
			result[top][i] = start
			start++
		}
		top++
		// 构造从上到下
		for i := top; i <= bottom; i++ {
			result[i][right] = start
			start++
		}
		right--
		// 构造从右到左
		for i := right; i >= left; i-- {
			result[bottom][i] = start
			start++
		}
		bottom--
		// 构造从下到上
		for i := bottom; i >= top; i-- {
			result[i][left] = start
			start++
		}
		left++
	}
	return result
}
