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
func findTilt(root *TreeNode) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 0
		}
		l, r := dfs(cur.Left), dfs(cur.Right)
		ans += abs(l - r)
		return l + r + cur.Val
	}
	dfs(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
