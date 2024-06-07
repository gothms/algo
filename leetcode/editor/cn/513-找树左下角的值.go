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
func findBottomLeftValue(root *TreeNode) int {
	ans, last := root.Val, 0
	var dfs func(int, *TreeNode)
	dfs = func(d int, r *TreeNode) {
		if r == nil {
			return
		}
		if d > last {
			ans, last = r.Val, d
		}
		dfs(d+1, r.Left)
		dfs(d+1, r.Right)
	}
	dfs(0, root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
