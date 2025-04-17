package main

/*
107. 二叉树的层序遍历 II
中等
给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
和102解法一样，只是最后结果反转一下
*/

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
