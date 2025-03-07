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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 迭代：lc

	// 递归
	var dfs func(int, int, int) *TreeNode
	dfs = func(i, l, r int) *TreeNode {
		if l > r {
			return nil
		}
		j := l
		for j <= r && inorder[j] != preorder[i] {
			j++
		}
		return &TreeNode{preorder[i], dfs(i+1, l, j-1), dfs(i+j-l+1, j+1, r)}
	}
	return dfs(0, 0, len(inorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)
