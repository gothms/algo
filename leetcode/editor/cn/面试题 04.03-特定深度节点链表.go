//给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
//
//
//
// 示例：
//
// 输入：[1,2,3,4,5,null,7,8]
//
//        1
//       /  \
//      2    3
//     / \    \
//    4   5    7
//   /
//  8
//
//输出：[[1],[2,3],[4,5,7],[8]]
//
//
// Related Topics 树 广度优先搜索 链表 二叉树 👍 100 👎 0

package main

func main() {

}

/*
思路：dfs
	1.本质是树的层序遍历，可以bfs，也dfs
		将每层的节点按链表串起来
	2.dfs
		头节点得是每层的最左边节点，所以按照 根-右-左 遍历节点
		并且每次将每层的头节点替换为当前节点
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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
	ret := make([]*ListNode, 0)
	var dfs func(*TreeNode, int)
	dfs = func(r *TreeNode, l int) {
		if r == nil {
			return
		}
		if len(ret) == l {
			ret = append(ret, &ListNode{Val: r.Val}) // 新层
		} else {
			ret[l] = &ListNode{r.Val, ret[l]} // 头节点为当前节点
		}
		dfs(r.Right, l+1) // 根-右-左
		dfs(r.Left, l+1)
	}
	dfs(tree, 0)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
