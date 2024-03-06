//给定两个整数数组，preorder 和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵
//树的后序遍历，重构并返回二叉树。
//
// 如果存在多个答案，您可以返回其中 任何 一个。
//
//
//
// 示例 1：
//
//
//
//
//输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
//输出：[1,2,3,4,5,6,7]
//
//
// 示例 2:
//
//
//输入: preorder = [1], postorder = [1]
//输出: [1]
//
//
//
//
// 提示：
//
//
// 1 <= preorder.length <= 30
// 1 <= preorder[i] <= preorder.length
// preorder 中所有值都 不同
// postorder.length == preorder.length
// 1 <= postorder[i] <= postorder.length
// postorder 中所有值都 不同
// 保证 preorder 和 postorder 是同一棵二叉树的前序遍历和后序遍历
//
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 341 👎 0

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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	// dfs
	n := len(preorder)
	idx := make([]int, n+1)
	for i, v := range postorder {
		idx[v] = i
	}
	var dfs func(preL, preR, postL, postR int) *TreeNode
	dfs = func(preL, preR, postL, postR int) *TreeNode {
		if preL == preR {
			return nil
		} else if preL+1 == preR {
			return &TreeNode{Val: preorder[preL]}
		} else {
			leftLen := idx[preorder[preL+1]] - postL + 1
			return &TreeNode{preorder[preL],
				dfs(preL+1, preL+leftLen+1, postL, postL+leftLen),
				dfs(preL+leftLen+1, preR, postL+leftLen, postR-1)}
		}
	}
	return dfs(0, n, 0, n)
}

//leetcode submit region end(Prohibit modification and deletion)
