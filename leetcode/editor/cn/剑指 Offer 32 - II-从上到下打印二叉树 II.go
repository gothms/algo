//从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
//
//
//
// 例如: 给定二叉树: [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层次遍历结果：
//
// [
//  [3],
//  [9,20],
//  [15,7]
//]
//
//
//
//
// 提示：
//
//
// 节点总数 <= 1000
//
//
// 注意：本题与主站 102 题相同：https://leetcode-cn.com/problems/binary-tree-level-order-
//traversal/
//
// Related Topics 树 广度优先搜索 二叉树 👍 289 👎 0

package main

func main() {

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
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(r *TreeNode, l int) {
		if r == nil {
			return
		}
		if len(ans) == l {
			ans = append(ans, []int{r.Val})
		} else {
			ans[l] = append(ans[l], r.Val)
		}
		dfs(r.Left, l+1)
		dfs(r.Right, l+1)
	}
	dfs(root, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
