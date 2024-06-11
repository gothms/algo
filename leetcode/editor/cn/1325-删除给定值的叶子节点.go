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
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var dfs func(*TreeNode) int
	dfs = func(cur *TreeNode) int {
		if cur == nil {
			return 0
		}
		l, r := dfs(cur.Left), dfs(cur.Right)
		if l == 1 {
			cur.Left = nil
		}
		if r == 1 {
			cur.Right = nil
		}
		if cur.Left == nil && cur.Right == nil && cur.Val == target {
			return 1
		}
		return 0
	}
	p := &TreeNode{Left: root}
	dfs(p)
	return p.Left
}

//leetcode submit region end(Prohibit modification and deletion)
