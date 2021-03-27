package main

// 给定一个前序遍历和中序遍历构建二叉树
func BuildTreeFromPre() {
	var (
		preorder = []int{3, 9, 20, 15, 7}
		inorder  = []int{9, 3, 15, 20, 7}
	)
	tree := buildTreeFromPre(preorder, inorder)
	preorderTraversalPrint(tree)
}

// 给定一个前序遍历和中序遍历构建二叉树
// @param 	前序序列 [根结点,[左子树的前序遍历结果],[右子树的前序遍历结果]]
// @param 	中序序列 [[左子树的中序遍历结果],根结点,[右子树的中序遍历结果]]
// @remark	时间复杂度o(n),空间复杂度o(n)
func buildTreeFromPre(preorder, inorder []int) *TreeNode {
	preLen := len(preorder)
	inLen := len(inorder)
	if preLen != inLen {
		return nil
	}
	// 在中序序列里，数值与下标的映射
	inMap := make(map[int]int)
	for i := 0; i < inLen; i++ {
		inMap[inorder[i]] = i
	}
	return buildTreeFromPreCore(preorder, 0, preLen-1, inMap, 0, inLen-1)
}

// 给定一个前序遍历和中序遍历构建二叉树
// @param	前序序列		[根结点,[左子树的前序遍历结果],[右子树的前序遍历结果]]
// @param	preStart	前序序列子区间起始下标
// @param	preEnd		前序序列子区间结束下标
// @param	inMap		在中序序列里，数值与下标的映射
// @param	inStart		中序序列子区间起始下标
// @param	inEnd		中序序列子区间结束下标
func buildTreeFromPreCore(preorder []int, preStart, preEnd int, inMap map[int]int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	// 根结点
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	// 根结点在中序序列中的下标
	index := inMap[rootVal]
	// 左子树的长度
	leftLen := index - inStart
	// 因为在前序序列中与中序序列中的左子树的长度相等，则左子树右边界为preLeft+leftLen
	root.Left = buildTreeFromPreCore(preorder, preStart+1, preStart+leftLen, inMap, inStart, index-1)
	root.Right = buildTreeFromPreCore(preorder, preStart+leftLen+1, preEnd, inMap, index+1, inEnd)
	return root
}
