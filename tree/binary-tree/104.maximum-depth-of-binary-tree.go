package main

// 二叉树最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 左子树深度
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// 使用层序遍历获取二叉树最大深度
func maxDepthWithLevelOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}

// N叉树最大深度
func maxDepthN(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 0
	if root.Children != nil {
		for _, v := range root.Children {
			depth = max(depth, maxDepthN(v))
		}
	}
	return 1 + depth
}

// 使用层序遍历获取N叉树最大深度
func maxDepthNWithLevelOrder(root *Node) int {
	if root == nil {
		return 0
	}
	queue := []*Node{root}
	depth := 0
	for len(queue) > 0 {
		depth++
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if len(node.Children) > 0 {
				for _, v := range node.Children {
					queue = append(queue, v)
				}
			}
		}
	}
	return depth
}
