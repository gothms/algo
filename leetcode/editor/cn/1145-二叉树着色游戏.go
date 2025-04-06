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
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	l, r := 0, 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lc := dfs(root.Left)
		rc := dfs(root.Right)
		if root.Val == x {
			l, r = lc, rc
		}
		return lc + rc + 1
	}
	dfs(root)
	return max(l, r, n-l-r-1)<<1 > n

	// 个人
	//l, r := 0, 0
	//var cnt func(*TreeNode) int
	//cnt = func(root *TreeNode) int {
	//	if root == nil {
	//		return 0
	//	}
	//	return cnt(root.Left) + cnt(root.Right) + 1
	//}
	//var dfs func(*TreeNode)
	//dfs = func(root *TreeNode) {
	//	if root == nil {
	//		return
	//	}
	//	if root.Val == x {
	//		l, r = cnt(root.Left), cnt(root.Right)
	//		return
	//	}
	//	dfs(root.Left)
	//	dfs(root.Right)
	//}
	//dfs(root)
	//y := max(l, r, n-l-r-1)
	//return y<<1 > n
}

//leetcode submit region end(Prohibit modification and deletion)
