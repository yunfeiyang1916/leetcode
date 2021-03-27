package main

import (
	"fmt"
)

// 对称二叉树测试
func IsSymmetricTest() {
	queue := []*TreeNode{&TreeNode{Val: 1}, nil, nil}
	fmt.Println(len(queue))
}

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

// 比较左右子树是否对称
func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	// 有一个为nil则不对称
	if left == nil || right == nil {
		return false
	}
	// 再排除值
	if left.Val != right.Val {
		return false
	}
	// 走到这说明：左右结点都不为空，且数值相同的情况
	// 此时才做递归，做下一层的判断
	outside := compare(left.Left, right.Right) // 左子树左结点与右子树右结点比较，也就是最外侧结点比较
	inside := compare(left.Right, right.Left)  // 左子树右结点与右子树左结点比较，也就是最内侧结点比较
	return outside && inside
}

// 对称二叉树,使用队列和层序遍历
func isSymmetricWithQueue(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	// nil也能入队成功
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		// 都是成对入队，不会出现数组越界
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		// 只要有一个为空，则不对称
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		// 最外侧结点都不为空，则入队
		queue = append(queue, left.Left, right.Right)
		// 内侧结点都不为空，则入队
		queue = append(queue, left.Right, right.Left)
	}
	return true
}

// 100.给你两棵二叉树的根结点 p 和 q ，编写一个函数来检验这两棵树是否相同。
// @remark 时间复杂度o(min(m,n))，空间复杂度o(min(m,n))
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// 只要有一个为nil则不对称
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	// 两个树的左子树是否相等
	left := isSameTree(p.Left, q.Left)
	// 两个树的右子树是否相等
	right := isSameTree(p.Right, q.Right)
	return left && right
}
