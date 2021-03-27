package main

// N叉树
type Node struct {
	Val      int
	Children []*Node
}

// N叉树层序遍历
func levelOrderN(root *Node) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*Node, 0)
	// 根结点入队
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		list := make([]int, 0, n)
		for i := 0; i < n; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			// 所有子结点入队
			if len(node.Children) > 0 {
				for _, v := range node.Children {
					queue = append(queue, v)
				}
			}
		}
		result = append(result, list)
	}
	return result
}
