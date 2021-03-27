package main

import (
	"strconv"
)

// 给定一个二叉树，返回所有从根结点到叶子结点的路径。
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	return traversal(root, "")
}

// 先序回溯访问
func traversal(root *TreeNode, path string) []string {
	result := make([]string, 0)
	path += strconv.Itoa(root.Val)
	// 遍历到叶子结点，则结束
	if root.Left == nil && root.Right == nil {
		result = append(result, path)
		return result
	}
	if root.Left != nil {
		result = append(result, traversal(root.Left, path+"->")...)
	}
	if root.Right != nil {
		result = append(result, traversal(root.Right, path+"->")...)
	}
	return result
}
