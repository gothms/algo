//给你一棵二叉树的根 root ，请你将每个节点的值替换成该节点的所有 堂兄弟节点值的和 。
//
// 如果两个节点在树中有相同的深度且它们的父节点不同，那么它们互为 堂兄弟 。
//
// 请你返回修改值之后，树的根 root 。
//
// 注意，一个节点的深度指的是从树根节点到这个节点经过的边数。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [5,4,9,1,10,null,7]
//输出：[0,0,0,7,7,null,11]
//解释：上图展示了初始的二叉树和修改每个节点的值之后的二叉树。
//- 值为 5 的节点没有堂兄弟，所以值修改为 0 。
//- 值为 4 的节点没有堂兄弟，所以值修改为 0 。
//- 值为 9 的节点没有堂兄弟，所以值修改为 0 。
//- 值为 1 的节点有一个堂兄弟，值为 7 ，所以值修改为 7 。
//- 值为 10 的节点有一个堂兄弟，值为 7 ，所以值修改为 7 。
//- 值为 7 的节点有两个堂兄弟，值分别为 1 和 10 ，所以值修改为 11 。
//
//
// 示例 2：
//
//
//
//
//输入：root = [3,1,2]
//输出：[0,0,0]
//解释：上图展示了初始的二叉树和修改每个节点的值之后的二叉树。
//- 值为 3 的节点没有堂兄弟，所以值修改为 0 。
//- 值为 1 的节点没有堂兄弟，所以值修改为 0 。
//- 值为 2 的节点没有堂兄弟，所以值修改为 0 。
//
//
//
//
// 提示：
//
//
// 树中节点数目的范围是 [1, 10⁵] 。
// 1 <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 哈希表 二叉树 👍 20 👎 0

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
func replaceValueInTree(root *TreeNode) *TreeNode {
	deep := make([]int, 0)
	var dfs func(r *TreeNode, d int) int // 每层的和
	dfs = func(r *TreeNode, d int) int {
		if r == nil {
			return 0
		}
		v := r.Val
		if d == len(deep) {
			deep = append(deep, 0)
		}
		deep[d] += v
		r.Val = dfs(r.Left, d+1) + dfs(r.Right, d+1)
		return v
	}
	dfs(root, 0)
	var Dfs func(r *TreeNode, d, v int) // 修正
	Dfs = func(r *TreeNode, d, v int) {
		if r == nil {
			return
		}
		val := r.Val
		Dfs(r.Left, d+1, val)
		Dfs(r.Right, d+1, val)
		r.Val = deep[d] - v
	}
	Dfs(root, 0, deep[0])
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
