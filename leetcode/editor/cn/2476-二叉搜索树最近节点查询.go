//给你一个 二叉搜索树 的根节点 root ，和一个由正整数组成、长度为 n 的数组 queries 。
//
// 请你找出一个长度为 n 的 二维 答案数组 answer ，其中 answer[i] = [mini, maxi] ：
//
//
// mini 是树中小于等于 queries[i] 的 最大值 。如果不存在这样的值，则使用 -1 代替。
// maxi 是树中大于等于 queries[i] 的 最小值 。如果不存在这样的值，则使用 -1 代替。
//
//
// 返回数组 answer 。
//
//
//
// 示例 1 ：
//
//
//
//
//输入：root = [6,2,13,1,4,9,15,null,null,null,null,null,null,14], queries = [2,5,1
//6]
//输出：[[2,2],[4,6],[15,-1]]
//解释：按下面的描述找出并返回查询的答案：
//- 树中小于等于 2 的最大值是 2 ，且大于等于 2 的最小值也是 2 。所以第一个查询的答案是 [2,2] 。
//- 树中小于等于 5 的最大值是 4 ，且大于等于 5 的最小值是 6 。所以第二个查询的答案是 [4,6] 。
//- 树中小于等于 16 的最大值是 15 ，且大于等于 16 的最小值不存在。所以第三个查询的答案是 [15,-1] 。
//
//
// 示例 2 ：
//
//
//
//
//输入：root = [4,null,9], queries = [3]
//输出：[[-1,4]]
//解释：树中不存在小于等于 3 的最大值，且大于等于 3 的最小值是 4 。所以查询的答案是 [-1,4] 。
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [2, 10⁵] 内
// 1 <= Node.val <= 10⁶
// n == queries.length
// 1 <= n <= 10⁵
// 1 <= queries[i] <= 10⁶
//
//
// Related Topics 树 深度优先搜索 二叉搜索树 数组 二分查找 二叉树 👍 32 👎 0

package main

import "slices"

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
func closestNodes(root *TreeNode, queries []int) [][]int {
	arr := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		arr = append(arr, r.Val)
		dfs(r.Right)
	}
	dfs(root)
	ret := make([][]int, len(queries))
	n := len(arr)
	for i, q := range queries {
		l, r := -1, -1
		j, ok := slices.BinarySearch(arr, q)
		if j < n {
			r = arr[j]
		}
		if !ok {
			j--
		}
		if j >= 0 {
			l = arr[j]
		}
		ret[i] = []int{l, r}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
