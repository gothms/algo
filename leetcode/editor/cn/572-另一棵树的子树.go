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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(r, s *TreeNode) bool {
		return r == nil && s == nil ||
			r != nil && s != nil && r.Val == s.Val && check(r.Left, s.Left) && check(r.Right, s.Right)
	}
	var dfs func(*TreeNode) bool
	dfs = func(cur *TreeNode) bool {
		if cur == nil {
			return false
		}
		if check(cur, subRoot) {
			return true
		}
		return dfs(cur.Left) || dfs(cur.Right)
	}
	return dfs(root)
}

//leetcode submit region end(Prohibit modification and deletion)
