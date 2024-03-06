//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 1885 👎 0

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
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	memo := []*TreeNode{root}
	for i, n := 0, len(memo); n > 0; i, n = i+1, len(memo) {
		ret = append(ret, make([]int, 0))
		for j := 0; j < n; j++ {
			tn := memo[j]
			ret[i] = append(ret[i], tn.Val)
			if tn.Left != nil {
				memo = append(memo, tn.Left)
			}
			if tn.Right != nil {
				memo = append(memo, tn.Right)
			}
		}
		memo = memo[n:]
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
