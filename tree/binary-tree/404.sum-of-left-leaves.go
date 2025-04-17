package main

/*
404. 左叶子之和
简单
给定二叉树的根节点 root ，返回所有左叶子之和。
*/
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 叶子节点也可以提前返回
	if root.Left == nil && root.Right == nil {
		return 0
	}
	leftValue := 0
	// 说明该节点是左叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val
	}
	return leftValue + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
