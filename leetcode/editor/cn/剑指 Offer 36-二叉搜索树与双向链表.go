//输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
//
//
//
// 为了让您更好地理解问题，以下面的二叉搜索树为例：
//
//
//
//
//
//
//
// 我们希望将这个二叉搜索树转化为双向循环链表。链表中的每个节点都有一个前驱和后继指针。对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是
//第一个节点。
//
// 下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。
//
//
//
//
//
//
//
// 特别地，我们希望可以就地完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中的第一个节点的指针。
//
//
//
// 注意：本题与主站 426 题相同：https://leetcode-cn.com/problems/convert-binary-search-tree-
//to-sorted-doubly-linked-list/
//
// 注意：此题对比原题有改动。
//
// Related Topics 栈 树 深度优先搜索 二叉搜索树 链表 二叉树 双向链表 👍 693 👎 0

package main

import "fmt"

func main() {
	l := &TreeNode{1, nil, nil}
	root := &TreeNode{4, &TreeNode{2, l, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}}
	convert := Convert(root)
	fmt.Println(convert.Val)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Right
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：dfs
	1.二叉搜索树的中序遍历，即为升序排序
	2.dfs 中序遍历时，先访问左节点，再处理当前值，最后访问右节点
		所以，把二叉搜索树变成双向链表的关键就在于，处理当前值的逻辑
	3.用一个变量 p 记录当前最大的节点，结合上面分析，有：
		当前遍历的节点 curr 即为在排序中，仅大于 p 的节点
	4.指针变更
		a)p.Right 为下一个更大的节点，为当前遍历的节点 curr。curr 的左节点为 p
		b)p 指向下一个更大的节点 curr
		c)利用 p 节点初始值为 nil，寻找到双向链表的 头节点 minTreeNode
			即最小的节点，为二叉搜索树最“左”的节点
*/

// There is no code of Go type for this problem
func Convert(root *TreeNode) *TreeNode {
	// dfs：二叉查找树的顺序
	var p, minTN *TreeNode
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		if p != nil { // p：二叉查找树的 当前最大数值
			p.Right = r // Right 指向下一个 更大数值
		} else {
			minTN = r
		}
		r.Left, p = p, r // Left 指向上一个 最大数值，并保留 当前最大数值
		dfs(r.Right)
	}
	dfs(root)
	return minTN
}
