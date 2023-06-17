//给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
//
// 二叉搜索树的定义如下：
//
//
// 任意节点的左子树中的键值都 小于 此节点的键值。
// 任意节点的右子树中的键值都 大于 此节点的键值。
// 任意节点的左子树和右子树都是二叉搜索树。
//
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
//输出：20
//解释：键值为 3 的子树是和最大的二叉搜索树。
//
//
// 示例 2：
//
//
//
//
//输入：root = [4,3,null,1,2]
//输出：2
//解释：键值为 2 的单节点子树是和最大的二叉搜索树。
//
//
// 示例 3：
//
//
//输入：root = [-4,-2,-5]
//输出：0
//解释：所有节点键值都为负数，和最大的二叉搜索树为空。
//
//
// 示例 4：
//
//
//输入：root = [2,1,3]
//输出：6
//
//
// 示例 5：
//
//
//输入：root = [5,4,8,3,null,6,3]
//输出：7
//
//
//
//
// 提示：
//
//
// 每棵树有 1 到 40000 个节点。
// 每个节点的键值在 [-4 * 10^4 , 4 * 10^4] 之间。
//
//
// Related Topics 树 深度优先搜索 二叉搜索树 动态规划 二叉树 👍 120 👎 0

package main

import (
	"math"
)

func main() {

}

/*
思路：采用 dfs 的要点
	1.任一节点 curr 对应的二叉树是二叉搜索树的条件是：左、右子树都是二叉搜索树
		当左、右子树的节点的 Val 范围分别是 [minl,maxl] 和 [minr,maxr] 时，若 curr 为 二叉搜索树，有：
			curr.Val > maxl：大于左子树节点 Val 的最大值
			curr.Val < minr：小于右子树节点 Val 的最小值
		而要满足上述条件，只有在分别遍历了 curr 的左右子树，是否是二叉搜索树后，才可以进行上述两对值的比较
			所以当采用 dfs 来解决此题时，需要将子树的 Val 取值区间范围返回
			也就是 [minl,maxl] 和 [minr,maxr]
	2.当一个节点是叶子节点 leaf 时，怎么处理返回值问题？
		很显然，当一个 leaf 节点返回时，它的 Val 取值区间是 [leaf.Val,leaf.Val]
		但是在 dfs 中，怎么处理 leaf 的两个 nil 子节点？可预见的方式有两种
		2.1.当一个节点的子节点是 nil 时，不往下遍历 nil 节点，而是提前判断是左右子节点否为 nil
			a)当 curr.left == nil && curr.right == nil，curr 为 leaf 节点，返回 [curr.Val,curr.Val]
			b)当 curr.left == nil && curr.right != nil，则返回的 minl = curr.Val
				curr.right == nil && curr.left != nil，同理
			c)左右子节点均不为 nil，则从左右返回的 [minl,maxl] 和 [minr,maxr]，取 [minl,maxr] 返回
			这3种情况，都是 curr 对应的二叉树是二叉搜索树
		2.2.不判断当前 curr 节点的左右节点是否为 nil
			这种处理方式的好处是，不会带来子节点为 nil 的各种情况判断
			坏处是：（详见代码）
				当 curr == nil 时，返回 [math.MaxInt32, math.MinInt32]
				则处理 2.1.c 的情况时，需要小技巧来返回 [minl,maxr]
	3.假设所求值 二叉搜索树最大和 为 maxSum
		由于节点值可能是负数，所以可能存在当二叉搜索树 curr 的最大和为 sum 时
		curr 的某个 二叉搜索子树 的 最大和 > sum
		所以需要实时比较每个二叉搜索树的和，以求出最大值
		而不是返回到根节点再取值
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 *
 */
func maxSumBST(root *TreeNode) int {
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	maxSum := 0
	var dfs func(*TreeNode) (int, int, bool, int)
	dfs = func(r *TreeNode) (int, int, bool, int) {
		if r == nil {
			return math.MaxInt32, math.MinInt32, true, 0
		}
		minl, maxl, bstl, vl := dfs(r.Left)
		minr, maxr, bstr, vr := dfs(r.Right)
		if bstl && bstr && r.Val > maxl && r.Val < minr {
			sum := vl + vr + r.Val
			maxSum = maxVal(maxSum, sum)
			return minVal(r.Val, minl), maxVal(r.Val, maxr), true, sum
		}
		return 0, 0, false, 0
	}
	dfs(root)
	return maxSum
}

//leetcode submit region end(Prohibit modification and deletion)
