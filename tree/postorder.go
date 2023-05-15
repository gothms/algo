package tree

// lc-145
func PostorderTraversal(root *TreeNode) []int {
	// Morris
	postorder, curr := make([]int, 0), root
	var pre *TreeNode
	for curr != nil {
		if curr.Right == nil {
			postorder = append(postorder, curr.Val)
			curr = curr.Left
		} else {
			pre = curr.Right
			for pre.Left != nil && pre.Left != curr {
				pre = pre.Left
			}
			if pre.Left == nil {
				postorder = append(postorder, curr.Val)
				pre.Left, curr = curr, curr.Right
			} else {
				pre.Left, curr = nil, curr.Left
			}
		}
	}
	reverse(postorder)
	return postorder

	// stack
	//if root == nil {
	//	return nil
	//}
	//postorder, stack := make([]int, 0), []*TreeNode{root}
	//var curr *TreeNode
	//for l := len(stack) - 1; l >= 0; l = len(stack) - 1 {
	//	curr, stack = stack[l], stack[:l]
	//	postorder = append(postorder, curr.Val)
	//	if curr.Left != nil {
	//		stack = append(stack, curr.Left)
	//	}
	//	if curr.Right != nil {
	//		stack = append(stack, curr.Right)
	//	}
	//}
	//reverse(postorder)
	//return postorder

	// dfs
	//postorder := make([]int, 0)
	//dfsPostorder(root, &postorder)
	//return postorder
}
func dfsPostorder(root *TreeNode, post *[]int) {
	if root == nil {
		return
	}
	dfsPostorder(root.Left, post)
	dfsPostorder(root.Right, post)
	*post = append(*post, root.Val)
}
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
