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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// lc：迭代

	// 个人
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val < low {
			return dfs(root.Right)
		} else if root.Val > high {
			return dfs(root.Left)
		}
		root.Left, root.Right = dfs(root.Left), dfs(root.Right)
		return root
	}
	return dfs(root)
}

//leetcode submit region end(Prohibit modification and deletion)
