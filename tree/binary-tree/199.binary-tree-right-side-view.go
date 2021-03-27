package main

import "fmt"

// 二叉树的右视图，给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的结点值。
func RightSideViewTest() {
	var (
		preorder = []int{3, 9, 20, 15, 7}
		inorder  = []int{9, 3, 15, 20, 7}
	)
	tree := buildTreeFromPre(preorder, inorder)
	fmt.Println(rightSideView(tree))
}

// 二叉树的右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	// 根结点入队
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			// 说明该结点是该层最右侧结点
			if i == n-1 {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}
