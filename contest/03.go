package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 9
	nums = []int{4, 1, 3, 2, 1, 5}
	target = 7
	//nums = []int{3, 88, 93, 38, 97, 87, 5, 75, 56, 67, 25, 3, 4, 45, 8, 27, 60, 94, 70, 41, 2, 91, 89, 83, 39, 64, 83, 83, 29, 90, 27, 25, 65, 89, 100, 76, 7, 20, 94, 68, 73, 100, 25, 77, 72, 93, 85, 59, 60, 100, 96, 20, 50, 45, 10, 61, 49, 45, 40, 18, 48, 36, 15, 9, 82, 78, 28, 86, 100, 61, 11, 3, 80, 83, 35, 59, 39, 62, 8, 10, 53, 91, 19, 11, 73, 18, 18, 70, 62, 1, 57, 10, 2, 22, 29, 83, 93, 55, 27, 52}
	//target = 66 // 13
	//nums = []int{100}
	//target = 100
	subsequence := lengthOfLongestSubsequence(nums, target)
	fmt.Println(subsequence)
}
func lengthOfLongestSubsequence(nums []int, target int) int {
	sort.Ints(nums)
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})
	memo := make([]int, target+1)
	for _, v := range nums[:idx] {
		for i := target - v; i >= 0; i-- {
			s := i + v
			if memo[i] != 0 {
				memo[s] = maxVal(memo[i]+1, memo[s])
			}
		}
		memo[v] = maxVal(memo[v], 1)
	}
	if memo[target] > 0 {
		return memo[target]
	}
	return -1

	// dfs
	//sort.Ints(nums)
	//idx := sort.Search(len(nums), func(i int) bool {
	//	return nums[i] > target
	//})
	//memo := map[int]int{0: 0}
	//for _, v := range nums[:idx] {
	//	temp := make(map[int]int)
	//	for k, c := range memo {
	//		if s := k + v; s <= target {
	//			temp[s] = maxVal(memo[s], c+1)
	//		}
	//	}
	//	for k, c := range temp {
	//		memo[k] = c
	//	}
	//}
	//fmt.Println(memo)
	//if c, ok := memo[target]; ok {
	//	return c
	//}
	//return -1
}
func maxVal(a, b int) int {
	if b > a {
		return b
	}
	return a
}
