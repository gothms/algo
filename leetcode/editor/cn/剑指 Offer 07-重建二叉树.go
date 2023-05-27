//输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
//
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//
//
// 示例 1:
//
//
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
//Input: preorder = [-1], inorder = [-1]
//Output: [-1]
//
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 5000
//
//
//
// 注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-
//preorder-and-inorder-traversal/
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 1071 👎 0

package main

import "fmt"

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	//preorder = []int{4, 1, 2, 3}
	//inorder = []int{1, 2, 3, 4}
	tree := buildTree(preorder, inorder)
	fmt.Println(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	// 迭代：stack

	// dfs
	n := len(preorder) - 1
	var dfs func(int, int, int) *TreeNode
	dfs = func(pi, ii, ij int) *TreeNode {
		switch {
		case ii == ij:
			return &TreeNode{inorder[ii], nil, nil}
		case ii > ij:
			return nil
		}
		i := ii
		for i <= ij && inorder[i] != preorder[pi] {
			i++
		}
		// preorder 的索引移动的距离，等于 inorder 的索引 i 移动的距离
		return &TreeNode{inorder[i], dfs(pi+1, ii, i-1), dfs(pi+i-ii+1, i+1, ij)}
	}
	return dfs(0, 0, n)
}

//leetcode submit region end(Prohibit modification and deletion)
