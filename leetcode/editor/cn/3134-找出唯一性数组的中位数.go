package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	//nums = []int{3, 4, 3, 4, 5}
	//nums = []int{4, 3, 5, 4}
	array := medianOfUniquenessArray(nums)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func medianOfUniquenessArray(nums []int) int {
	//二分 + 滑动窗体
	// https://leetcode.cn/problems/find-the-median-of-the-uniqueness-array/solutions/2759114/er-fen-da-an-hua-dong-chuang-kou-pythonj-ykg9/
	n := len(nums)
	k := (n*(n+1)>>1 + 1) >> 1 // 第 k 大
	return sort.Search(n, func(i int) bool {
		cnt, l := 0, 0
		memo := make(map[int]int)
		for r, v := range nums {
			memo[v]++
			for ; len(memo) > i; l++ {
				if memo[nums[l]] == 1 {
					delete(memo, nums[l])
				} else {
					memo[nums[l]]--
				}
			}
			cnt += r - l + 1
			if cnt >= k {
				return true
			}
		}
		return false
	})
}

//leetcode submit region end(Prohibit modification and deletion)

//func medianOfUniquenessArray(nums []int) int {
//	// 二分 + 滑动窗体
//	// https://leetcode.cn/problems/find-the-median-of-the-uniqueness-array/solutions/2759114/er-fen-da-an-hua-dong-chuang-kou-pythonj-ykg9/
//	n := len(nums)
//	k := (n*(n+1)>>1 + 1) >> 1 // 第 k 大
//	return sort.Search(n, func(i int) bool {
//		cnt, l := 0, 0
//		memo := make(map[int]int)
//		for r, v := range nums {
//			memo[v]++
//			for ; len(memo) > i; l++ {
//				if memo[nums[l]] == 1 {
//					delete(memo, nums[l])
//				} else {
//					memo[nums[l]]--
//				}
//			}
//			cnt += r - l + 1
//			if cnt >= k {
//				return true
//			}
//		}
//		return false
//	})
//}
