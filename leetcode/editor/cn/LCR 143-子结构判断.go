package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	// lc
	return B != nil && dfs143(A, B)

	// 个人
	//if B == nil {
	//	return false
	//}
	//var check func(*TreeNode, *TreeNode) bool
	//check = func(a, b *TreeNode) bool {
	//	if b == nil {
	//		return true
	//	}
	//	return a != nil && a.Val == b.Val && check(a.Left, b.Left) && check(a.Right, b.Right)
	//}
	//var dfs func(root *TreeNode) bool
	//dfs = func(root *TreeNode) bool {
	//	if root == nil {
	//		return false
	//	}
	//	return root.Val == B.Val && check(root.Left, B.Left) && check(root.Right, B.Right) ||
	//		dfs(root.Left) ||
	//		dfs(root.Right)
	//}
	//return dfs(A)
}
func dfs143(A, B *TreeNode) bool {
	return A != nil &&
		(check(A, B) || dfs143(A.Left, B) || dfs143(A.Right, B))
}

func check(a, b *TreeNode) bool {
	return b == nil ||
		a != nil && a.Val == b.Val && check(a.Left, b.Left) && check(a.Right, b.Right)
}

//leetcode submit region end(Prohibit modification and deletion)
