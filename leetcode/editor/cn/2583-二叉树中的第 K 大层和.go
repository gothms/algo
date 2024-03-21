//给你一棵二叉树的根节点 root 和一个正整数 k 。
//
// 树中的 层和 是指 同一层 上节点值的总和。
//
// 返回树中第 k 大的层和（不一定不同）。如果树少于 k 层，则返回 -1 。
//
// 注意，如果两个节点与根节点的距离相同，则认为它们在同一层。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [5,8,9,2,1,3,7,4,6], k = 2
//输出：13
//解释：树中每一层的层和分别是：
//- Level 1: 5
//- Level 2: 8 + 9 = 17
//- Level 3: 2 + 1 + 3 + 7 = 13
//- Level 4: 4 + 6 = 10
//第 2 大的层和等于 13 。
//
//
// 示例 2：
//
//
//
//
//输入：root = [1,2,null,3], k = 1
//输出：3
//解释：最大的层和是 3 。
//
//
//
//
// 提示：
//
//
// 树中的节点数为 n
// 2 <= n <= 10⁵
// 1 <= Node.val <= 10⁶
// 1 <= k <= n
//
//
// Related Topics 树 广度优先搜索 二叉树 排序 👍 15 👎 0

package main

import (
	"container/heap"
	"sort"
)

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
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q, h := []*TreeNode{root}, &hp2583{}
	add := func(node *TreeNode) {
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	for l := len(q); l > 0; l = len(q) {
		var v int
		for i := 0; i < l; i++ {
			v += q[i].Val
			add(q[i])
		}
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
		q = q[l:]
	}
	if h.Len() < k {
		return -1
	}
	return int64(h.IntSlice[0])

	//hp, q := &kllsHp{}, []*TreeNode{root}
	//for l := len(q); l > 0; l = len(q) {
	//	var s int
	//	for i := 0; i < l; i++ {
	//		tn := q[i]
	//		s += tn.Val
	//		if tn.Left != nil {
	//			q = append(q, tn.Left)
	//		}
	//		if tn.Right != nil {
	//			q = append(q, tn.Right)
	//		}
	//	}
	//	q = q[l:]
	//	heap.Push(hp, s)
	//	if hp.Len() > k {
	//		heap.Pop(hp)
	//	}
	//}
	//if hp.Len() < k {
	//	return -1
	//}
	//return int64(hp.IntSlice[0])
}

type hp2583 struct {
	sort.IntSlice
}

func (h *hp2583) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp2583) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

//type kllsHp struct {
//	sort.IntSlice
//}
//
//func (k *kllsHp) Push(x any) {
//	k.IntSlice = append(k.IntSlice, x.(int))
//}
//
//func (k *kllsHp) Pop() any {
//	v := k.IntSlice[len(k.IntSlice)-1]
//	k.IntSlice = k.IntSlice[:len(k.IntSlice)-1]
//	return v
//}

//leetcode submit region end(Prohibit modification and deletion)
