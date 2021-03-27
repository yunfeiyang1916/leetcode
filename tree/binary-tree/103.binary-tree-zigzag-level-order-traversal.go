package main

// 二叉树的锯齿形层序遍历,就是奇偶层交错输出
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	// 根结点入队
	queue = append(queue, root)
	// 是否切换
	toggle := false
	for len(queue) > 0 {
		n := len(queue)
		list := make([]int, 0, n)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if toggle {
			// 反转
			for i := 0; i < n/2; i++ {
				list[i], list[n-1-i] = list[n-1-i], list[i]
			}
		}
		result = append(result, list)
		toggle = !toggle
	}
	return result
}
