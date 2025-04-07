package main

import "math"

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
func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(root *TreeNode, l, r int) bool {
		if root == nil {
			return true
		}
		v := root.Val
		if v <= l || v >= r {
			return false
		}
		return dfs(root.Left, l, v) && dfs(root.Right, v, r)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

//leetcode submit region end(Prohibit modification and deletion)
