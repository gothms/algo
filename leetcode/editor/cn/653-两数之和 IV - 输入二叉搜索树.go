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
func findTarget(root *TreeNode, k int) bool {
	memo := make(map[int]struct{})
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if _, ok := memo[root.Val]; ok {
			return true
		}
		memo[k-root.Val] = struct{}{}
		return dfs(root.Left) || dfs(root.Right)
	}
	return dfs(root)
}

//leetcode submit region end(Prohibit modification and deletion)
