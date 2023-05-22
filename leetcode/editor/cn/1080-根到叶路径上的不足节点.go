//给你二叉树的根节点 root 和一个整数 limit ，请你同时删除树中所有 不足节点 ，并返回最终二叉树的根节点。
//
// 假如通过节点 node 的每种可能的 “根-叶” 路径上值的总和全都小于给定的 limit，则该节点被称之为 不足节点 ，需要被删除。
//
// 叶子节点，就是没有子节点的节点。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
//输出：[1,2,3,4,null,null,7,8,9,null,14]
//
//
// 示例 2：
//
//
//输入：root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
//输出：[5,4,8,11,null,17,4,7,null,null,null,5]
//
//
// 示例 3：
//
//
//输入：root = [1,2,-3,-5,null,4,null], limit = -1
//输出：[1,null,-3,4]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 5000] 内
// -10⁵ <= Node.val <= 10⁵
// -10⁹ <= limit <= 10⁹
//
//
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 82 👎 0

package main

func main() {

}

/*
思路：典型的 dfs
	1.如何判断和处理 根-叶 的路径和 >= limit
		1.1.只有在遍历某条路径的 叶子 节点求出 sum，才能知道
			sum>=limit：保留该 叶-根
			sum<limit：废弃该 叶-根
		1.2.任意节点 curr 的 left路径 和 right路径:
			a)均不满足所求，则 curr 节点所在子树 pass，并向上返回 false
			b)left 满足，right 不满足，pass right，left路径所有节点 全保留，并向上返回 true
				left 不满足，right 满足，亦然
			c)left、right 均满足，均保留，并向上返回 true
		1.3.结合上面，可知：
			1.1.自顶向下遍历节点，求sum
			1.2.自底向上返回该路径是否满足，从而判断是否保留，并返回
	2.注意点在于各种 nil 判断
		本代码采用较简洁的判断方式，但不一定最高效
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
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(r *TreeNode, limit int, sum int) bool {
		sum += r.Val
		if r.Left == nil && r.Right == nil {
			return sum >= limit
		}
		sL := r.Left != nil && dfs(r.Left, limit, sum)
		sR := r.Right != nil && dfs(r.Right, limit, sum)
		if !sL {
			r.Left = nil
		}
		if !sR {
			r.Right = nil
		}
		return sL || sR
	}
	if !dfs(root, limit, 0) {
		return nil
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
