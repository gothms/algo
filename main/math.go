package main

import (
	"algo/math"
	"fmt"
	"math/rand"
	"sort"
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
	permute := math.Permute(nums)
	//fmt.Println(permute)
	//permute = math.PermuteDFS(nums)
	//fmt.Println(permute)
	permute = math.PermuteUnique(nums)
	fmt.Println(permute, len(permute))

	// 子集
	//nums = []int{1, 2, 3, 4}
	subsets := math.Subsets(nums)
	fmt.Println(subsets, len(subsets))
	// 子集 & 组合
	dfs := math.SubsetsWithDupDFS(nums)
	sort.Slice(dfs, func(i, j int) bool {
		return len(dfs[i]) < len(dfs[j]) || len(dfs[i]) == len(dfs[j]) && func(i, j int) bool {
			for idx := range dfs[i] {
				if dfs[i][idx] < dfs[j][idx] {
					return true
				}
			}
			return false
		}(i, j)
	})
	fmt.Println(dfs)
	counter := math.SubsetsAndCombineCounter(nums)
	fmt.Println(counter)

	// 组合：重复元素
	nums = []int{1, 1, 1, 2, 3, 4, 5, 6, 6}
	nums = []int{1, 1, 1}       // [[1] [1 1] [1 1 1] [1 1 1 2] [1 1 2] [1 2] [2]] 7
	nums = []int{1, 1, 1, 2}    // [[1] [1 1] [1 1 1] [1 1 1 2] [1 1 2] [1 2] [2]] 7
	nums = []int{1, 1, 1, 2, 2} // [[1] [1 1] [1 1 1] [1 1 1 2] [1 1 1 2 2] [1 1 2] [1 1 2 2] [1 2] [1 2 2] [2] [2 2]] 11
	//nums = []int{1, 1, 1, 2, 2, 2} // [[1] [1 1] [1 1 1] [1 1 1 2] [1 1 1 2 2] [1 1 2] [1 1 2 2] [1 2] [1 2 2] [2] [2 2]] 15
	//nums = []int{1, 1, 1, 2, 3}    // [[1] [1 1] [1 1 1] [1 1 1 2] [1 1 1 2 3] [1 1 1 3] [1 1 2] [1 1 2 3] [1 1 3] [1 2] [1 2 3] [1 3] [2] [2 3] [3]] 15
	//nums = []int{1, 1, 1, 2, 2, 3} // [[1] [1 1] [1 1 1] [1 1 1 2] [1 1 1 2 2] [1 1 1 2 2 3] [1 1 1 2 3] [1 1 1 3] [1 1 2] [1 1 2 2] [1 1 2 2 3] [1 1 2 3] [1 1 3] [1 2] [1 2 2] [1 2 2 3] [1 2 3] [1 3] [2] [2 2] [2 2 3] [2 3] [3]] 23
	//nums = []int{1, 2, 2, 3, 3, 3}
	nums = make([]int, 0, 10)
	for i := 0; i < 4; i++ {
		nums = append(nums, rand.Intn(10))
	}
	k := 3
	k = 2
	dup := math.CombineWithDup(nums, k)
	//fmt.Println(dup, len(dup)) // 2：17，3：31，4：41，5：41，6：31，7：17
	dup = math.CombineWithDupDFS(nums, k)
	fmt.Println(dup, len(dup)) // 2：17，3：31，4：41，5：41，6：31，7：17
	counterK := math.SubsetsAndCombineCounterK(nums, k)
	fmt.Println(counterK)

	// 子集 & 排列
	//andPermute := math.SubsetsAndPermute(nums)
	////fmt.Println(andPermute, len(andPermute))
	//for _, ap := range andPermute {
	//	fmt.Println(ap, len(ap))
	//}
	permuteCounter := math.SubsetsAndPermuteCounter(nums)
	fmt.Println(permuteCounter)

	// 排列：k/任意长度
	nums = []int{1, 2, 3, 4, 5}
	nums = []int{1, 2, 3}
	k = 2
	permuteK := math.PermuteK(nums, k)
	fmt.Println(permuteK, len(permuteK))
	permuteK = math.PermuteAll(nums)
	fmt.Println(permuteK, len(permuteK))
	nums = []int{1, 2, 3, 3, 1}
	permuteK = math.PermuteUniqueK(nums, k)
	fmt.Println(permuteK, len(permuteK))
}
