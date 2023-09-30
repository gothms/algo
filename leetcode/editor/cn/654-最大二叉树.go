//给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:
//
//
// 创建一个根节点，其值为 nums 中的最大值。
// 递归地在最大值 左边 的 子数组前缀上 构建左子树。
// 递归地在最大值 右边 的 子数组后缀上 构建右子树。
//
//
// 返回 nums 构建的 最大二叉树 。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,1,6,0,5]
//输出：[6,3,5,null,2,0,null,null,1]
//解释：递归调用如下所示：
//- [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
//    - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
//        - 空数组，无子节点。
//        - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
//            - 空数组，无子节点。
//            - 只有一个元素，所以子节点是一个值为 1 的节点。
//    - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
//        - 只有一个元素，所以子节点是一个值为 0 的节点。
//        - 空数组，无子节点。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,1]
//输出：[3,null,2,null,1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 1000
// nums 中的所有整数 互不相同
//
//
// Related Topics 栈 树 数组 分治 二叉树 单调栈 👍 732 👎 0

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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// 参考：lc-503
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = -1 // 否则 right 没法取值 -1
	}
	stack, tns := make([]int, 1, n+1), make([]*TreeNode, n)
	stack[0] = -1 // 哨兵 & 为 left 提供 -1
	for i, v := range nums {
		tns[i] = &TreeNode{Val: v}
		last := len(stack) - 1
		for last > 0 && v > nums[stack[last]] { // 右侧比 nums[stack[last]] 大的节点为 nums[i]
			right[stack[last]] = i
			last--
		}
		left[i] = stack[last] // 左侧比 nums[i] 大的节点
		stack = stack[:last+1]
		stack = append(stack, i)
	}
	var root *TreeNode
	for i, t := range tns {
		l, r := left[i], right[i]
		if l == -1 && r == -1 { // 根节点
			root = t
		} else if l == -1 || r != -1 && nums[l] > nums[r] {
			tns[r].Left = t // 左右比 t 大的节点，较小的为其父节点
		} else {
			tns[l].Right = t // 左小右大，则 t 是左侧节点的右节点
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
