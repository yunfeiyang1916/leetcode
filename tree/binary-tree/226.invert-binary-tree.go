package main

// 翻转二叉树
func InvertTreeTest() {

}

// 使用层序遍历的方式翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// 使用先序遍历的方式翻转二叉树
func invertTreeByPre(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 根结点左右结点先交换顺序
	root.Left, root.Right = root.Right, root.Left
	// 左子树交换顺序
	invertTreeByPre(root.Left)
	// 右子树交换顺序
	invertTreeByPre(root.Right)
	return root
}

// 借用栈使用先序遍历的方式翻转二叉树
func invertTreeByPreWithStack(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	// 根结点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		// 出栈
		n := len(stack)
		node := stack[n-1]
		stack = stack[:n-1]
		node.Left, node.Right = node.Right, node.Left
		// 右结点入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		// 左结点入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return root
}
