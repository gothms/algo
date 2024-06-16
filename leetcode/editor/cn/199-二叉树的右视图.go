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
func rightSideView(root *TreeNode) []int {
	// dfs
	ans := make([]int, 0)
	var dfs func(int, *TreeNode)
	dfs = func(d int, cur *TreeNode) {
		if cur == nil {
			return
		}
		if d == len(ans) {
			ans = append(ans, cur.Val)
		}
		d++
		dfs(d, cur.Right)
		dfs(d, cur.Left)
	}
	dfs(0, root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
