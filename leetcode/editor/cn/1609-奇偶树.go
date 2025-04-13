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
func isEvenOddTree(root *TreeNode) bool {
	// lc：bfs

	// 个人
	h := make([]int, 0)
	var dfs func(*TreeNode, int) bool
	dfs = func(root *TreeNode, i int) bool {
		if root == nil {
			return true
		}
		if (root.Val^i)&1 == 0 {
			return false
		}
		if len(h) == i {
			h = append(h, root.Val)
		} else if i&1 == 0 && root.Val <= h[i] ||
			i&1 == 1 && root.Val >= h[i] {
			return false
		} else {
			h[i] = root.Val
		}
		return dfs(root.Left, i+1) && dfs(root.Right, i+1)
	}
	return dfs(root, 0)
}

//leetcode submit region end(Prohibit modification and deletion)
