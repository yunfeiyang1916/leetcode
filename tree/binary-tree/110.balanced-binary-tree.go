package main

// 是否平衡二叉树，一个二叉树每个结点 的左右两个子树的高度差的绝对值不超过 1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getDepth(root) != -1
}

// 返回以该结点为根结点的二叉树的高度，如果不是二叉平衡树了则返回-1,
// @remark 使用自底向上(后序)的方式递归,时间复杂度o(n),空间复杂度o(n)
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 左子树的高度
	leftDepth := getDepth(root.Left)
	// 说明左子树已经不是平衡二叉树了
	if leftDepth == -1 {
		return -1
	}
	// 右子树高度
	rightDepth := getDepth(root.Right)
	// 说明右子树已经不是平衡树了
	if rightDepth == -1 {
		return -1
	}
	// 左右子树的高度差的绝对值超过1了，说明不是平衡二叉树
	if abs(leftDepth-rightDepth) > 1 {
		return -1
	}
	return 1 + max(leftDepth, rightDepth)
}

// 求绝对值
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
