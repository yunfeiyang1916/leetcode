package main

// 给定一个后序遍历和中序遍历构建二叉树
func BuildTreeFromPost() {
	var (
		postorder = []int{9, 15, 7, 20, 3}
		inorder   = []int{9, 3, 15, 20, 7}
	)
	tree := buildTreeFromPost(inorder, postorder)
	preorderTraversalPrint(tree)
}

// 给定一个后序遍历和中序遍历构建二叉树
func buildTreeFromPost(inorder, postorder []int) *TreeNode {
	inLen := len(inorder)
	postLen := len(postorder)
	if inLen != postLen {
		return nil
	}
	// 中序序列值与下标的映射
	inMap := make(map[int]int)
	for i := 0; i < inLen; i++ {
		inMap[inorder[i]] = i
	}
	return buildTreeFromPostCore(postorder, 0, postLen-1, inMap, 0, inLen-1)
}

// 给定一个后序遍历和中序遍历构建二叉树核心算法
// @param	postorder	后序遍历序列
// @param	postStart	后序序列起始下标
// @param	postEnd		后序序列结束下标
// @param	inMap		中序序列值与下标的映射
// @param	inStart		中序序列起始下标
// @param	inEnd		中序序列结束下标
func buildTreeFromPostCore(postorder []int, postStart, postEnd int, inMap map[int]int, inStart, inEnd int) *TreeNode {
	if postStart > postEnd || inStart > inEnd {
		return nil
	}
	rootVal := postorder[postEnd]
	root := &TreeNode{Val: rootVal}
	index := inMap[rootVal]
	// 左子树在中序序列中的长度
	leftLen := index - inStart
	root.Left = buildTreeFromPostCore(postorder, postStart, postStart+leftLen-1, inMap, inStart, index-1)
	root.Right = buildTreeFromPostCore(postorder, postStart+leftLen, postEnd-1, inMap, index+1, inEnd)
	return root
}
