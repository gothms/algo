//给你两棵二叉树： root1 和 root2 。
//
// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠
//，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。
//
// 返回合并后的二叉树。
//
// 注意: 合并过程必须从两个树的根节点开始。
//
//
//
// 示例 1：
//
//
//输入：root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
//输出：[3,4,5,5,4,null,7]
//
//
// 示例 2：
//
//
//输入：root1 = [1], root2 = [1,2]
//输出：[2,2]
//
//
//
//
// 提示：
//
//
// 两棵树中的节点数目在范围 [0, 2000] 内
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1260 👎 0

package main

func main() {

}

/*
思路：dfs
	将两棵树合并到 root1，考虑两种情况：
	r1 和 r2 对应位置都有节点，那么：r1.Val + r2.Val
	r1 和 r2 对应位置只有一个节点（不论是 r1 还是 r2 为 nil），那么：
		假设 r1 为nil，那么把 r2 拼接到 r1 为nil的位置即可
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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// dfs
	if root1 == nil {
		return root2 // r1为nil，把r2拼接过去
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left, root1.Right =
		mergeTrees(root1.Left, root2.Left), mergeTrees(root1.Right, root2.Right)
	return root1
}

//leetcode submit region end(Prohibit modification and deletion)
