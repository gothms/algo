package tree

// lc-144
func preorderTraversal(root *TreeNode) []int {
	// Morris：未破坏原二叉树结构
	preorder, curr := make([]int, 0), root
	var pre *TreeNode
	for curr != nil {
		if curr.Left == nil {
			preorder = append(preorder, curr.Val)
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}
			if pre.Right == nil {
				preorder = append(preorder, curr.Val)
				pre.Right, curr = curr, curr.Left
			} else {
				pre.Right, curr = nil, curr.Right
			}
		}
	}
	return preorder

	// stack
	//if root == nil {
	//	return nil
	//}
	//preorder, stack := make([]int, 0), []*TreeNode{root}
	//var curr *TreeNode
	//for l := len(stack) - 1; l >= 0; l = len(stack) - 1 {
	//	curr, stack = stack[l], stack[:l]
	//	preorder = append(preorder, curr.Val)
	//	if curr.Right != nil {
	//		stack = append(stack, curr.Right)
	//	}
	//	if curr.Left != nil {
	//		stack = append(stack, curr.Left)
	//	}
	//}
	//return preorder

	// dfs
	//preorder := make([]int, 0)
	//dfsPreorder(root, &preorder)
	//return preorder
}
func dfsPreorder(root *TreeNode, preorder *[]int) {
	if root == nil {
		return
	}
	*preorder = append(*preorder, root.Val)
	dfsPreorder(root.Left, preorder)
	dfsPreorder(root.Right, preorder)
}
