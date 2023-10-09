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
	nums := []int{1, 2, 1, 4}
	permute := math.Permute(nums)
	//fmt.Println(permute)
	//permute = math.PermuteDFS(nums)
	//fmt.Println(permute)
	permute = math.PermuteUnique(nums)
	fmt.Println(permute)

	// 子集
	//nums = []int{1, 2, 3, 4}
	//subsets := math.Subsets(nums)
	//fmt.Println(subsets)

	// 组合：重复元素
	nums = []int{1, 1, 1, 2, 3, 4, 5, 6, 6}
	//nums = []int{1, 1, 1, 1, 1, 2, 3, 4, 5, 6, 6, 6, 6}
	k := 3
	dup := math.CombineWithDup(nums, k)
	fmt.Println(dup, len(dup)) // 2：17，3：31，4：41，5：41，6：31，7：17
	dup = math.CombineWithDupDFS(nums, k)
	fmt.Println(dup, len(dup)) // 2：17，3：31，4：41，5：41，6：31，7：17
}
