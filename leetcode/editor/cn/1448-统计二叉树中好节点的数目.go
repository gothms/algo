//给你一棵根为 root 的二叉树，请你返回二叉树中好节点的数目。
//
// 「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。
//
//
//
// 示例 1：
//
//
//
// 输入：root = [3,1,4,3,null,1,5]
//输出：4
//解释：图中蓝色节点为好节点。
//根节点 (3) 永远是个好节点。
//节点 4 -> (3,4) 是路径中的最大值。
//节点 5 -> (3,4,5) 是路径中的最大值。
//节点 3 -> (3,1,3) 是路径中的最大值。
//
// 示例 2：
//
//
//
// 输入：root = [3,3,null,4,2]
//输出：3
//解释：节点 2 -> (3, 3, 2) 不是好节点，因为 "3" 比它大。
//
// 示例 3：
//
// 输入：root = [1]
//输出：1
//解释：根节点是好节点。
//
//
//
// 提示：
//
//
// 二叉树中节点数目范围是 [1, 10^5] 。
// 每个节点权值的范围是 [-10^4, 10^4] 。
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 90 👎 0

package main

import "math"

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
/*
思路：dfs
	1.翻译下好节点：当某个节点，大于等于路径上所有节点的值时，这个节点就是好节点
	2.dfs
		2.1.终止条件
			当前节点为空，返回0
		2.2.当前节点 root.Val >= 路径节点最大值 v
			v = root.Val，好节点数 +1
		2.3.dfs 下探
			dfs(r.Left, v) + dfs(r.Right, v)
*/
func goodNodes(root *TreeNode) int {
	// dfs
	var dfs func(*TreeNode, int) int
	dfs = func(r *TreeNode, v int) int {
		if r == nil {
			return 0
		}
		if r.Val < v {
			return dfs(r.Left, v) + dfs(r.Right, v)
		}
		return 1 + dfs(r.Left, r.Val) + dfs(r.Right, r.Val)
	}
	return dfs(root, math.MinInt32)
}

//leetcode submit region end(Prohibit modification and deletion)
