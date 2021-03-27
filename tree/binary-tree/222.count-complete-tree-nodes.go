package main

// 完全二叉树的结点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		result = 0
		queue  = []*TreeNode{root}
	)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			result++
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
