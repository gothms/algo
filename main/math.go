package main

import (
	"algo/math"
	"fmt"
)

func main() {
	// 组合
	//n, k := 4, 2
	//combine := math.Combine(n, k)
	//fmt.Println(combine)
	//combine = math.CombineBit(n, k)
	//fmt.Println(combine)
	//combine = math.CombineDFS(n, k)
	//fmt.Println(combine)
	//combine = math.CombineDFSChoose(n, k)
	//fmt.Println(combine)

	// 排列
	nums := []int{1, 2, 3, 4}
	//permute := math.Permute(nums)
	//fmt.Println(permute)
	//permute = math.PermuteDFS(nums)
	//fmt.Println(permute)

	// 子集
	nums = []int{1, 2, 3, 4}
	subsets := math.Subsets(nums)
	fmt.Println(subsets)
}
