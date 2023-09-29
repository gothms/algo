//给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
//输出：15
//
//
// 示例 2：
//
//
//输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
//输出：19
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 10⁴] 之间。
// 1 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 167 👎 0

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
func deepestLeavesSum(root *TreeNode) int {
	v, d := 0, 0
	var dfs func(*TreeNode, int)
	dfs = func(r *TreeNode, i int) {
		if r == nil {
			return
		}
		switch {
		case i == d:
			v = v + r.Val
		case i > d:
			v, d = r.Val, i
		}
		dfs(r.Left, i+1)
		dfs(r.Right, i+1)
	}
	dfs(root, 0)
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
