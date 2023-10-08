//给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：
//
//
// 该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）
//
//
// 如果不存在祖父节点值为偶数的节点，那么返回 0 。
//
//
//
// 示例：
//
//
//
// 输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
//输出：18
//解释：图中红色节点的祖父节点的值为偶数，蓝色节点为这些红色节点的祖父节点。
//
//
//
//
// 提示：
//
//
// 树中节点的数目在 1 到 10^4 之间。
// 每个节点的值在 1 到 100 之间。
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 91 👎 0

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
func sumEvenGrandparent(root *TreeNode) int {
	var dfs func(*TreeNode, int, int) int
	dfs = func(r *TreeNode, pp, p int) int {
		if r == nil {
			return 0
		}
		var sum int
		if pp == 0 || p == 0 {
			sum = r.Val
		}
		pp = p - 1        // 祖父
		if r.Val&1 == 0 { // 父节点
			p = 1
		} else {
			p = -1
		}
		return sum + dfs(r.Left, pp, p) + dfs(r.Right, pp, p)
	}
	return dfs(root, -1, -1)
}

//leetcode submit region end(Prohibit modification and deletion)
