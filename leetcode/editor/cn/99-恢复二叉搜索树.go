//给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,3,null,null,2]
//输出：[3,1,null,null,2]
//解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
//
//
// 示例 2：
//
//
//输入：root = [3,1,4,null,null,2]
//输出：[2,1,4,null,null,3]
//解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
//
//
//
// 提示：
//
//
// 树上节点的数目在范围 [2, 1000] 内
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
//
// 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 895 👎 0

package main

func main() {
	root := &TreeNode{5, &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{8, nil, nil}},
		&TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, &TreeNode{4, nil, nil}, nil}}}
	recoverTree(root)
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
func recoverTree(root *TreeNode) {
}

//leetcode submit region end(Prohibit modification and deletion)

//func recoverTree(root *TreeNode) {
//	// Morris 遍历
//	//var (
//	//	preTN = &TreeNode{Val: math.MinInt32} // 中序遍历的前一个节点
//	//	x     *TreeNode                       // 先遍历到的，要交换的节点
//	//	y     *TreeNode                       // 后遍历到的，要交换的节点
//	//	pre   *TreeNode
//	//)
//	//check := func(curr *TreeNode) { // 检查是否符合二叉搜索树规则
//	//	if preTN.Val > curr.Val {
//	//		y = curr
//	//		if x == nil {
//	//			x = preTN
//	//		}
//	//	}
//	//	preTN = curr
//	//}
//	//// Morris 遍历：未破坏结构
//	//for curr := root; curr != nil; {
//	//	if curr.Left == nil {
//	//		check(curr)
//	//		curr = curr.Right
//	//	} else {
//	//		pre = curr.Left
//	//		for pre.Right != nil && pre.Right != curr {
//	//			pre = pre.Right
//	//		}
//	//		if pre.Right == nil {
//	//			pre.Right, curr = curr, curr.Left // 破坏结构
//	//		} else {
//	//			check(curr)
//	//			pre.Right, curr = nil, curr.Right // 恢复结构
//	//		}
//	//	}
//	//}
//	//x.Val, y.Val = y.Val, x.Val
//
//	// dfs：中序遍历中，记录逆序数对
//	var x, y *TreeNode                  // 记录需要交换的节点
//	last := &TreeNode{Val: math.MinInt} // 中序遍历的前一个节点
//	var dfs func(*TreeNode)
//	dfs = func(cur *TreeNode) {
//		if cur == nil {
//			return
//		}
//		dfs(cur.Left)
//		if cur.Val < last.Val {
//			if x == nil { // 第一次，记录两个节点。last 为确定的第一个目标
//				x, y = last, cur
//			} else { // 第二次，记录第二个结点
//				y = cur
//				return
//			}
//		}
//		last = cur
//		dfs(cur.Right)
//	}
//	dfs(root)
//	x.Val, y.Val = y.Val, x.Val
//}

//leetcode submit region end(Prohibit modification and deletion)
