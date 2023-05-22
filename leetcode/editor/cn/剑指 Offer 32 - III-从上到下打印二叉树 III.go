//请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
//  [20,9],
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
// Related Topics 树 广度优先搜索 二叉树 👍 289 👎 0

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
func levelOrder(root *TreeNode) [][]int {
	// bfs
	if root == nil {
		return nil
	}
	ans, memo, level := make([][]int, 0), []*TreeNode{root}, 0
	addTN := func(node *TreeNode) {
		if node != nil {
			memo = append(memo, node)
		}
	}
	for l := len(memo); l > 0; l = len(memo) {
		lVal := make([]int, l)
		for i := 0; i < l; i++ {
			addTN(memo[i].Left)
			addTN(memo[i].Right)
			if level&1 > 0 {
				lVal[l-i-1] = memo[i].Val // 从后往前
			} else {
				lVal[i] = memo[i].Val
			}
		}
		level++
		memo = memo[l:]
		ans = append(ans, lVal)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
