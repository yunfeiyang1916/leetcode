package main

import "fmt"

// 二叉树的层平均值
func AverageOfLevelsTest() {
	var (
		preorder = []int{3, 9, 20, 15, 7}
		inorder  = []int{9, 3, 15, 20, 7}
	)
	tree := buildTreeFromPre(preorder, inorder)
	fmt.Println(averageOfLevels(tree))
}

// 二叉树的层平均值，给定一个非空二叉树, 返回一个由每层结点平均值组成的数组。
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	result := make([]float64, 0)
	queue := make([]*TreeNode, 0)
	// 根结点入队
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		// 先求每层的和
		for i := 0; i < n; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 计算每层平均值
		result = append(result, float64(sum)/float64(n))
	}
	return result
}
