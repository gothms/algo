package main

import (
	"algo/tree"
	"fmt"
)

func main() {
	// SegmentTree
	array := []int{1, 2, 3, 4, 5}
	array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	n := len(array)
	st := make([]int, n<<2)
	tree.Build(0, n-1, array, 0, st)
	// [15 6 9 3 3 4 5 1 2 0 0 0 0 0 0 0 0 0 0 0]
	// [66 21 45 6 15 24 21 3 3 9 6 15 9 10 11 1 2 0 0 4 5 0 0 7 8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//fmt.Println(st)
	//sum := tree.STRangeSum(1, 3, 0, n-1, 0, st) // 33
	//fmt.Println(sum)
	//tree.STUpdate(3, 1, 0, n-1, 0, st)
	//fmt.Println(st)

	// SegmentTree & Lazy
	lazy := make([]int, n<<2)
	tree.STUpdateLazy(1, 7, 1, 0, n-1, 0, st, lazy)
	fmt.Println(st)
	tree.STUpdateLazy(3, 8, 2, 0, n-1, 0, st, lazy)
	fmt.Println(st)
	sumLazy := tree.STRangeSumLazy(2, 8, 0, n-1, 0, st, lazy) // 42+6+12
	fmt.Println(sumLazy)

	// Morris 遍历测试
	root := &tree.TreeNode{5, &tree.TreeNode{3, &tree.TreeNode{1, nil, &tree.TreeNode{2, nil, nil}}, &tree.TreeNode{4, nil, nil}},
		&tree.TreeNode{7, &tree.TreeNode{6, nil, nil}, &tree.TreeNode{9, &tree.TreeNode{8, nil, nil}, nil}}}
	traversal := tree.InorderTraversal(root)
	fmt.Println(traversal)

	// 检查二叉树结构是否遭到破坏
	inorder, stack := make([]int, 0), make([]*tree.TreeNode, 0)
	curr := root
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
