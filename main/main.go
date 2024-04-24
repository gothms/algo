package main

import (
	"algo/binary"
	"algo/bit"
	"algo/tree"
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	// tree
	zero := &tree.TreeNode{0, nil, nil} // 加上测试，但不满足平衡二叉树
	one := &tree.TreeNode{1, nil, nil}
	six := &tree.TreeNode{6, nil, nil}
	root := &tree.TreeNode{5,
		&tree.TreeNode{3, &tree.TreeNode{2, one, nil}, &tree.TreeNode{4, zero, nil}},
		&tree.TreeNode{9,
			&tree.TreeNode{7, six, &tree.TreeNode{8, nil, nil}},
			&tree.TreeNode{11, &tree.TreeNode{10, nil, nil}, nil}}}
	//traversal := tree.PostorderTraversal(root)
	//fmt.Println(traversal)

	inorderTraversal := tree.InorderTraversal(root)
	fmt.Println(inorderTraversal)

	test1 := tree.InorderStackTest(one)
	test2 := tree.InorderStackTest(six)
	test3 := tree.InorderStackTest(zero)
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)

	c := 25
	b := bit.GetHighestOneBit(c)
	fmt.Println(b)

	b = bit.GetLeadingZeros(c)
	fmt.Println(b)
	b = bits.LeadingZeros(uint(c))
	fmt.Println(b)

	//bits.Reverse()

	add := bit.Add(9, 5)
	fmt.Println(add)
	add = bit.Substract(9, 5)
	fmt.Println(add)
	add = bit.Multiply(9, 5)
	fmt.Println(add)
	q, r := bit.Divide(64, 1)
	fmt.Println(q, r)

	//var x string = nil	// 下面两种可以
	//var x interface{} = nil
	//var x error = nil

	// 二分
	arr := []int{1, 2, 3, 3, 3, 3, 4, 4, 5, 7, 7}
	ret := binary.BinarySearchFirstEqual(arr, 3)
	fmt.Println("BinarySearchFirstEqual:", ret)
	ret = binary.BinarySearchLastEqual(arr, 1)
	fmt.Println(ret)
	for i := 0; i < 8; i++ {
		ret = binary.BinarySearchLastSmaller(arr, i)
		idx := sort.Search(len(arr), func(x int) bool { // 通过 API 验证 BinarySearchLastSmaller
			return arr[x] > i
		}) - 1
		fmt.Println(i, ret, ret == idx)
	}

	//fmt.Println("mod", 5%-3) // 2
	//fmt.Println("mod", -7%4) // -3
}
