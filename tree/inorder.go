package tree

/*
lc-94
lc-99
*/

// InorderTraversal 二叉树中序遍历
func InorderTraversal(root *TreeNode) []int {
	// Morris：未破坏原二叉树结构
	inorder, curr := make([]int, 0), root
	var pre *TreeNode
	for curr != nil {
		if curr.Left == nil {
			inorder = append(inorder, curr.Val)
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right, curr = curr, curr.Left // 破坏结构
			} else {
				inorder = append(inorder, curr.Val)
				pre.Right, curr = nil, curr.Right // 恢复结构
			}
		}
	}
	return inorder

	// Morris：破坏了原二叉树结构（多平行单链）
	//inorder, curr := make([]int, 0), root
	//var pre *TreeNode
	//for curr != nil {
	//	if curr.Left == nil {
	//		inorder = append(inorder, curr.Val)
	//		curr = curr.Right
	//	} else {
	//		pre = curr.Left
	//		for pre.Right != nil {
	//			pre = pre.Right
	//		}
	//		pre.Right, curr, curr.Left = curr, curr.Left, nil	// 破坏结构
	//	}
	//}
	//return inorder

	// stack
	//inorder, stack := make([]int, 0), make([]*TreeNode, 0)
	//curr := root
	//for curr != nil || len(stack) > 0 {
	//	if curr != nil {
	//		for curr.Left != nil {
	//			curr, stack = curr.Left, append(stack, curr)
	//		}
	//	} else {
	//		curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
	//	}
	//	inorder, curr = append(inorder, curr.Val), curr.Right
	//}
	//return inorder

	// dfs
	//inorder := make([]int, 0)
	//dfsInorder(root, &inorder)
	//return inorder
}
func dfsInorder(root *TreeNode, inorder *[]int) {
	if root == nil {
		return
	}
	dfsInorder(root.Left, inorder)
	*inorder = append(*inorder, root.Val)
	dfsInorder(root.Right, inorder)
}
func InorderStackTest(root *TreeNode) []int {
	inorder, stack := make([]int, 0), make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stack) > 0 {
		if curr != nil {
			for curr.Left != nil {
				curr, stack = curr.Left, append(stack, curr)
			}
		} else {
			curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
		}
		inorder, curr = append(inorder, curr.Val), curr.Right
	}
	return inorder
}
