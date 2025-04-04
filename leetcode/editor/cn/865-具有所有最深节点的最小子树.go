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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var (
		p  *TreeNode
		mx int
	)
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, level int) int {
		if root == nil {
			return level
		}
		level++
		l, r := dfs(root.Left, level), dfs(root.Right, level)
		if l == r && l >= mx {
			p, mx = root, l
		}
		return max(l, r)
	}
	dfs(root, 0)
	return p
}

//leetcode submit region end(Prohibit modification and deletion)
