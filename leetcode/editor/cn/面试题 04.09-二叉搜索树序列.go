//从左向右遍历一个数组，通过不断将其中的元素插入树中可以逐步地生成一棵二叉搜索树。
//
// 给定一个由不同节点组成的二叉搜索树 root，输出所有可能生成此树的数组。
//
//
//
// 示例 1:
//
//
//输入: root = [2,1,3]
//输出: [[2,1,3],[2,3,1]]
//解释: 数组 [2,1,3]、[2,3,1] 均可以通过从左向右遍历元素插入树中形成以下二叉搜索树
//       2
//      / \
//     1   3
//
//
//
//
//
// 示例 2:
//
//
//输入: root = [4,1,null,null,3,2]
//输出: [[4,1,3,2]]
//
//
//
//
// 提示：
//
//
// 二叉搜索树中的节点数在
// [0, 1000] 的范围内
// 1 <= 节点值 <= 10^6
// 用例保证符合要求的数组数量不超过 5000
//
//
// Related Topics 树 二叉搜索树 回溯 二叉树 👍 116 👎 0

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

func BSTSequences(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{nil} // [[]]
	}
	ret := make([][]int, 0)
	path := []int{root.Val}
	var dfs func(*TreeNode, []*TreeNode)
	dfs = func(r *TreeNode, bfs []*TreeNode) {
		if r.Left != nil { // bfs
			bfs = append(bfs, r.Left)
		}
		if r.Right != nil {
			bfs = append(bfs, r.Right)
		}
		l := len(bfs)
		if l == 0 { // 已遍历完
			ret = append(ret, append([]int(nil), path...))
			return
		}
		for i := 0; i < l; i++ {
			path = append(path, bfs[i].Val)
			bfs[0], bfs[i] = bfs[i], bfs[0] // 并没有打乱顺序，[0,l) 的节点是同一层级
			dfs(bfs[0], bfs[1:])            // 下一个节点现在是 0
			bfs[0], bfs[i] = bfs[i], bfs[0]
			path = path[:len(path)-1] // 回溯
		}
	}
	dfs(root, make([]*TreeNode, 0))
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
