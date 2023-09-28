//二叉树数据结构TreeNode可用来表示单向链表（其中left置空，right为下一个链表节点）。实现一个方法，把二叉搜索树转换为单向链表，要求依然符合二叉
//搜索树的性质，转换操作应是原址的，也就是在原始的二叉搜索树上直接修改。
//
// 返回转换后的单向链表的头节点。
//
// 注意：本题相对原题稍作改动
//
//
//
// 示例：
//
// 输入： [4,2,5,1,3,null,6,0]
//输出： [0,null,1,null,2,null,3,null,4,null,5,null,6]
//
//
// 提示：
//
//
// 节点数量不会超过 100000。
//
//
// Related Topics 栈 树 深度优先搜索 二叉搜索树 链表 二叉树 👍 139 👎 0

package main

import (
	"fmt"
)

func main() {
	r := &TreeNode{0, nil, nil}
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, r, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{5, nil, &TreeNode{6, nil, nil}}}
	convertBiNode(root)

	inorder, stack := make([]int, 0), make([]*TreeNode, 0)
	curr := r
	for curr != nil || len(stack) > 0 {
		if curr != nil {
			for curr.Left != nil {
				curr, stack = curr.Left, append(stack, curr)
			}
		} else {
			curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
		}
		inorder, curr = append(inorder, curr.Val), curr.Right
	}
	fmt.Println(inorder)
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
func convertBiNode(root *TreeNode) *TreeNode {
	var tn TreeNode
	pre, cur := &tn, &tn // 一个用于哨兵，一个用于遍历
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		cur.Right, cur, r.Left = r, r, nil // 中序遍历
		dfs(r.Right)
	}
	dfs(root)
	return pre.Right
}

//leetcode submit region end(Prohibit modification and deletion)
