//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//
// Related Topics 栈 树 深度优先搜索 二叉树 👍 2013 👎 0

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
func inorderTraversal(root *TreeNode) []int {
	ret, stack := make([]int, 0), make([]*TreeNode, 0)
	for cur := root; cur != nil || len(stack) > 0; cur = cur.Right {
		if cur == nil {
			cur, stack = stack[len(stack)-1], stack[:len(stack)-1]
		} else {
			for ; cur.Left != nil; cur = cur.Left {
				stack = append(stack, cur)
			}
		}
		ret = append(ret, cur.Val)
	}
	return ret

	//inorder, curr := make([]int, 0), root
	//var pre *TreeNode
	//for curr != nil {
	//	if curr.Left == nil {
	//		inorder = append(inorder, curr.Val)
	//		curr = curr.Right
	//	} else {
	//		pre = curr.Left
	//		for pre.Right != nil && pre.Right != curr {
	//			pre = pre.Right
	//		}
	//		if pre.Right == nil {
	//			pre.Right, curr = curr, curr.Left // 破坏结构
	//		} else {
	//			inorder = append(inorder, curr.Val)
	//			pre.Right, curr = nil, curr.Right // 恢复结构
	//		}
	//	}
	//}
	//return inorder
}

//leetcode submit region end(Prohibit modification and deletion)
