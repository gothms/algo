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
}
