package main

import (
	"fmt"
	"sort"
)

// 二叉树
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 构建二叉树测试
func Test() {
	var (
		preorder = []int{3, 9, 20, 15, 7}
		inorder  = []int{9, 3, 15, 20, 7}
	)
	tree := buildTreeFromPre(preorder, inorder)
	// 前序
	fmt.Println(preorderTraversalWithStack(tree))
	fmt.Println(preorderTraversalWithStack2(tree))
	// 中序
	//fmt.Println(inorderTraversal(tree))
	// 中序，借用栈
	fmt.Println(inorderTraversalWithStack(tree))
	fmt.Println(inorderTraversalWithStack2(tree))
	// 后序遍历，递归
	fmt.Println(postorderTraversal(tree))
	// 后序，借用栈
	fmt.Println(postorderTraversalWithStack(tree))
}

// 给定一串数字生成二叉树
func buildTree(arr []int) *TreeNode {
	arrLen := len(arr)
	if arrLen == 0 {
		return nil
	}
	// 先排序
	sort.Ints(arr)
	return buildTreeCore(arr, 0, arrLen)
}

// 给定一串数字生成平衡二叉树
func buildTreeCore(arr []int, start, end int) *TreeNode {
	if end < start {
		return nil
	}
	mid := (start + end) / 2
	root := &TreeNode{Val: arr[mid]}
	root.Left = buildTreeCore(arr, start, mid-1)
	root.Right = buildTreeCore(arr, mid+1, end)
	return root
}

// 前序遍历打印
func preorderTraversalPrint(root *TreeNode) {
	fmt.Println(preorderTraversal(root))
}

// 前序遍历,递归方式
// @remark 时间复杂度o(n),空间复杂度o(n)
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	// 先打印根结点
	// fmt.Printf("%d ", root.Val)
	result = append(result, root.Val)
	// 再遍历左子树
	result = append(result, preorderTraversal(root.Left)...)
	// 再遍历右子树
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

// 前序遍历，借用栈
func preorderTraversalWithStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 栈
	stack := make([]*TreeNode, 0)
	// 结果
	result := make([]int, 0)
	// 根结点入栈
	stack = append(stack, root)
	for {
		sLen := len(stack)
		// 栈为空，结束循环
		if sLen == 0 {
			break
		}
		// 出栈
		node := stack[sLen-1]
		stack = stack[:sLen-1]
		result = append(result, node.Val)
		// 先右子树入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		// 再左子树入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

// 前序遍历，嵌套循环
func preorderTraversalWithStack2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			// 先保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈
		sLen := len(stack)
		root = stack[sLen-1]
		stack = stack[:sLen-1]
		root = root.Right
	}
	return result
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)
	return result
}

// 中序遍历，借用栈
func inorderTraversalWithStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	node := root
	for {
		if node != nil {
			// 访问到最底层
			stack = append(stack, node)
			node = node.Left
		} else {
			sLen := len(stack)
			if sLen == 0 {
				break
			}
			// 开始出栈，出栈的结点就是需要打印的结点
			node = stack[sLen-1]
			stack = stack[:sLen-1]
			result = append(result, node.Val)
			node = node.Right
		}
	}
	return result
}

// 中序遍历，嵌套循环
func inorderTraversalWithStack2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		sLen := len(stack)
		// 出栈
		root = stack[sLen-1]
		stack = stack[:sLen-1]
		result = append(result, root.Val)
		root = root.Right
	}
	return result
}

// 后序遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

// 后序遍历，借用栈
func postorderTraversalWithStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		stack  = make([]*TreeNode, 0)
		result = make([]int, 0)
		// 上一个结点
		prev *TreeNode
		// 当前结点
		current = root
	)
	for current != nil || len(stack) > 0 {
		// 左子树入栈
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		// 出栈
		sLen := len(stack)
		current = stack[sLen-1]
		stack = stack[:sLen-1]
		// 右结点为空或者右结点被访问过
		if current.Right == nil || current.Right == prev {
			result = append(result, current.Val)
			prev = current
			current = nil
		} else {
			// 右结点尚未被访问，则根结点需要继续入栈
			stack = append(stack, current)
			current = current.Right
		}
	}
	return result
}

// 后序遍历，使用翻转方式实现，
// @remark 先序遍历是中左右，后续遍历是左右中，那么我们只需要调整一下先序遍历的代码顺序，就变成中右左的遍历顺序，然后在反转result数组
func postorderTraversalWithReverse(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		// 出栈
		n := len(stack)
		node := stack[n-1]
		stack = stack[:n-1]
		result = append(result, node.Val)
		// 先入栈左结点
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		// 再入栈右结点
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	// 翻转
	reverse(result)
	return result
}

// 翻转一维数组
func reverse(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
