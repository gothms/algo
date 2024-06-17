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
func buildTree(inorder []int, postorder []int) *TreeNode {

}

//leetcode submit region end(Prohibit modification and deletion)

func buildTree_(inorder []int, postorder []int) *TreeNode {
	// 迭代

	// dfs
	var dfs func(int, int, int) *TreeNode
	dfs = func(i, l, r int) *TreeNode {
		if l > r {
			return nil
		}
		j := l
		for j <= r && inorder[j] != postorder[i] {
			j++
		}
		return &TreeNode{postorder[i], dfs(i-r+j-1, l, j-1), dfs(i-1, j+1, r)}
	}
	return dfs(len(postorder)-1, 0, len(postorder)-1)
}
