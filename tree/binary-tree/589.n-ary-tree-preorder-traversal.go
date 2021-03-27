package main

// N叉树的前序遍历
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	// 先访问根结点
	result = append(result, root.Val)
	if len(root.Children) > 0 {
		for _, v := range root.Children {
			result = append(result, preorder(v)...)
		}
	}
	return result
}

// N叉树的前序遍历,使用栈
func preorderWithStack(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*Node, 0)
	// 根结点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		n := len(stack)
		// 出栈
		node := stack[n-1]
		stack = stack[:n-1]
		result = append(result, node.Val)
		// 倒序入栈
		n = len(node.Children)
		if n > 0 {
			for i := n - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
		}
	}
	return result
}

// N叉树的后序遍历
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	if len(root.Children) > 0 {
		for _, v := range root.Children {
			result = append(result, postorder(v)...)
		}
	}
	result = append(result, root.Val)
	return result
}

// N叉树后序遍历，使用栈
// @remark 先序遍历是中左右，后续遍历是左右中，那么我们只需要调整一下先序遍历的代码顺序，就变成中右左的遍历顺序，然后在反转result数组
func postorderWithStack(root *Node) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		n := len(stack)
		// 出栈
		node := stack[n-1]
		stack = stack[:n-1]
		result = append(result, node.Val)
		for _, v := range node.Children {
			stack = append(stack, v)
		}
	}
	// 翻转
	reverse(result)
	return result
}
