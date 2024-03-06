//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
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
// -100 <= Node.val <= 100
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 862 👎 0

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret, stack := make([][]int, 0), []*TreeNode{root}
	for l, idx := len(stack), 0; l > 0; l, idx = len(stack), idx+1 {
		ret = append(ret, make([]int, 0))
		for i := l - 1; i >= 0; i-- {
			ret[idx] = append(ret[idx], stack[i].Val)
			left, right := stack[i].Left, stack[i].Right
			if idx&1 == 1 {
				left, right = right, left
			}
			if left != nil {
				stack = append(stack, left)
			}
			if right != nil {
				stack = append(stack, right)
			}
		}
		stack = stack[l:]
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
