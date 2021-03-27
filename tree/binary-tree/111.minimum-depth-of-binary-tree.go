package main

// 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		depth = 0
		queue = []*TreeNode{root}
	)
	for len(queue) > 0 {
		depth++
		n := len(queue)
		for i := 0; i < n; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			// 如果是叶子结点则说明找到最小深度了
			if node.Left == nil && node.Right == nil {
				return depth
			}
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
