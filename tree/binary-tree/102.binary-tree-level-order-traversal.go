package main

import "fmt"

// 二叉树的层序遍历
func LevelOrderTest() {
	var (
		preorder = []int{3, 9, 20, 15, 7}
		inorder  = []int{9, 3, 15, 20, 7}
	)
	tree := buildTreeFromPre(preorder, inorder)
	fmt.Println(levelOrder(tree))
}

// 二叉树的层序遍历，就是广度优先遍历
// @remark 每个结点进队出队各一次，所以时间复杂度为o(n),队列中的原始个数不超过n，所以空间复杂度为o(n)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 队列总是存储一层的结点
	queue := make([]*TreeNode, 0)
	result := make([][]int, 0)
	// 根结点入队
	queue = append(queue, root)
	for len(queue) > 0 {
		// 一层一层遍历
		list := make([]int, 0)
		n := len(queue)
		for i := 0; i < n; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			// 左子树入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// 右子树入队
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)
	}
	return result
}

// 给定一个二叉树，返回其结点值自底向上的层次遍历
func levelOrderBottom(root *TreeNode) [][]int {
	// 先从上到下层序访问，再反转即可
	result := levelOrder(root)
	reverse2(result)
	return result
}

// 翻转二维数组
func reverse2(nums [][]int) {
	n := len(nums)
	if n == 0 {
		return
	}
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
