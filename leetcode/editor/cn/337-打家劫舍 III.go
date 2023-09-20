//小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为
// root 。
//
// 除了
// root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的
//房子在同一天晚上被打劫 ，房屋将自动报警。
//
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
//
//
//
// 示例 1:
//
//
//
//
//输入: root = [3,2,3,null,3,null,1]
//输出: 7
//解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
//
// 示例 2:
//
//
//
//
//输入: root = [3,4,5,1,3,null,1]
//输出: 9
//解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
//
//
//
//
// 提示：
//
//
//
//
//
// 树的节点数在 [1, 10⁴] 范围内
// 0 <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 动态规划 二叉树 👍 1787 👎 0

package main

func main() {

}

/*
思路：dfs
	同 lc-198：打家劫舍 I，lc-213：打家劫舍 II
	不同之处是，一个遍历切片，一个遍历二叉树
	注意点：将需要的数据，通过递归的“归”来往上传递
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
func rob(root *TreeNode) int {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var dfs func(*TreeNode) (int, int)
	dfs = func(r *TreeNode) (int, int) {
		if r == nil {
			return 0, 0
		}
		lv, lcv := dfs(r.Left)
		rv, rcv := dfs(r.Right)
		return maxVal(lcv+rcv+r.Val, lv+rv), lv + rv
	}
	_, v := dfs(root)
	return v

	// leetcode-337
	//maxVal := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//var dfs func(*TreeNode) (int, int)
	//dfs = func(r *TreeNode) (int, int) {
	//	if r == nil {
	//		return 0, 0
	//	}
	//	ccL, cL := dfs(r.Left)
	//	ccR, cR := dfs(r.Right)
	//	// cL + cR：记录不打劫 r 的值，用于计算 r 的父节点
	//	// maxVal(cL+cR, ccL+ccR+r.Val)：不打劫 r，或打劫 r
	//	return cL + cR, maxVal(cL+cR, ccL+ccR+r.Val)
	//}
	//_, c := dfs(root)
	//return c
}

//leetcode submit region end(Prohibit modification and deletion)
