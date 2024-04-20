package main

import (
	"algo/tree"
	"fmt"
	"math/bits"
)

func main() {
	// SegmentTree
	//array := []int{1, 2, 3, 4, 5}
	//array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	//array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	// 多次测试：对线段树的长度精确赋值
	//const N = 10000
	//array := make([]int, 0)
	//for i := 1; i <= N; i++ {
	//	array = append(array, i)
	//	n := len(array)
	//	k := bits.Len(uint(n - 1))
	//	stLen := 1 << (k + 1)
	//	st := make([]int, stLen)
	//	//tree.Build(0, n-1, array, 0, st)
	//	tree.Build(0, n-1, array, 1, st) // 从 1 开始
	//	//fmt.Println("线段树：", st, len(st))
	//
	//	left := 1<<(k-bits.Len(uint(n-stLen>>2))+1) - 2
	//	j := len(st) - 1
	//	for st[j] == 0 {
	//		j--
	//	}
	//	if left != len(st)-j-1 {
	//		//fmt.Println(bits.Len(uint(len(st) >> 2)))
	//		fmt.Println("error:", n, k, left, len(st)-j-1)
	//	}
	//}

	// 单次测试
	const N = 15
	array := make([]int, N)
	for i := 0; i < N; i++ {
		array[i] = i + 1
	}
	n := len(array)
	stLen := 1 << (bits.Len(uint(n-1)) + 1)
	if n > 1 {
		k := stLen >> 2
		fmt.Println(bits.Len(uint(n - stLen>>2)))
		stLen -= 1<<(bits.Len(uint(k))-bits.Len(uint(n-k))+1) - 2
	}
	fmt.Println(stLen)
	/*
		st n  d
		32 9  7 1 7 14
		32 10 6 2 3 6
		32 11 5 3 3 6
		32 12 4 4 1 2
		32 13 3 5 1 2
		32 14 2 6 1 2
		32 15 1 7 1 2
		32 16 0 8 0 0

		(32>>1-n)
	*/
	//st := make([]int, n<<2)
	st := make([]int, stLen)
	//tree.Build(0, n-1, array, 0, st)
	tree.Build(0, n-1, array, 1, st) // 从 1 开始
	fmt.Println("线段树：", st, len(st))

	// [15 6 9 3 3 4 5 1 2 0 0 0 0 0 0 0 0 0 0 0]
	// [66 21 45 6 15 24 21 3 3 9 6 15 9 10 11 1 2 0 0 4 5 0 0 7 8 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//fmt.Println(st)
	//sum := tree.STRangeSum(1, 3, 0, n-1, 0, st) // 33
	//fmt.Println(sum)
	//tree.STUpdate(3, 1, 0, n-1, 0, st)
	//fmt.Println(st)

	// SegmentTree & Lazy
	//lazy := make([]int, n<<2)
	//tree.STUpdateLazy(1, 7, 1, 0, n-1, 1, st, lazy)
	//fmt.Println(st)
	//tree.STUpdateLazy(3, 8, 2, 0, n-1, 1, st, lazy)
	//fmt.Println(st)
	//sumLazy := tree.STRangeSumLazy(2, 8, 0, n-1, 1, st, lazy) // 42+6+12
	//fmt.Println(sumLazy)

	// Morris 遍历测试
	//root := &tree.TreeNode{5, &tree.TreeNode{3, &tree.TreeNode{1, nil, &tree.TreeNode{2, nil, nil}}, &tree.TreeNode{4, nil, nil}},
	//	&tree.TreeNode{7, &tree.TreeNode{6, nil, nil}, &tree.TreeNode{9, &tree.TreeNode{8, nil, nil}, nil}}}
	//traversal := tree.InorderTraversal(root)
	//fmt.Println(traversal)
	//
	//// 检查二叉树结构是否遭到破坏
	//inorder, stack := make([]int, 0), make([]*tree.TreeNode, 0)
	//curr := root
	//for curr != nil || len(stack) > 0 {
	//	if curr != nil {
	//		for curr.Left != nil {
	//			curr, stack = curr.Left, append(stack, curr)
	//		}
	//	} else {
	//		curr, stack = stack[len(stack)-1], stack[:len(stack)-1]
	//	}
	//	inorder, curr = append(inorder, curr.Val), curr.Right
	//}
	//fmt.Println(inorder)
}
