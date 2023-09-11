//给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。注意，路径不一定非得从二叉树的
//根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。
//
// 示例: 给定如下二叉树，以及目标和 sum = 22，
//
//               5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \    / \
//        7    2  5   1
//
//
// 返回:
//
// 3
//解释：和为 22 的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]
//
// 提示：
//
//
// 节点总数 <= 10000
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 137 👎 0

package main

func main() {

}

/*
思路：dfs
	1.
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	// dfs
	var dfs func(*TreeNode, map[int]int, int) int
	dfs = func(r *TreeNode, memo map[int]int, s int) int {
		if r == nil {
			return 0
		}
		s += r.Val
		cnt := memo[s-sum] // 查询当前节点和的路径数
		memo[s]++          // 更新节点和
		cnt += dfs(r.Left, memo, s) + dfs(r.Right, memo, s)
		memo[s]-- // 回溯
		return cnt
	}
	return dfs(root, map[int]int{0: 1}, 0) // 初始化 [0:1]
}

//leetcode submit region end(Prohibit modification and deletion)
