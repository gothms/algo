package main

import (
	"fmt"
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
func constructMaximumBinaryTree(nums []int) *TreeNode {

}

// leetcode submit region end(Prohibit modification and deletion)

func constructMaximumBinaryTree_(nums []int) *TreeNode {
	// 单调栈
	//st := []*TreeNode{{Val: math.MaxInt32}}
	//for _, v := range nums {
	//	for v > st[len(st)-1].Val {
	//		st = st[:len(st)-1]
	//	}
	//	cur := &TreeNode{Val: v}
	//	cur.Left, st[len(st)-1].Right = st[len(st)-1].Right, cur
	//	st = append(st, cur)
	//}
	//return st[0].Right

	// dfs

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
	fmt.Println(left)
	fmt.Println(right)
	fmt.Println(stack)
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
