package main

import (
	"fmt"
	"sort"
)

func main() {
	n1 := []int{0, 0, 0, 0, 1}
	n2 := []int{1, 0, 0, 0, 0}
	n1 = []int{0, 0, 0, 0, 0}
	n2 = []int{0, 0, 0, 0, 0}
	n1 = []int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}
	n2 = []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0} // 9
	n1 = []int{1, 2, 3, 2, 1}
	n2 = []int{3, 2, 1, 4, 7} // 3
	length := findLength(n1, n2)
	fmt.Println(length)

	//const prime = 16777619
	//var v, pow uint32 = 1, 1055306571
	//v = v*prime + 2
	//v = v*prime + 3
	//fmt.Println(v)
	//v = v*prime + 2 - 1*pow
	//fmt.Println(v)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findLength(nums1 []int, nums2 []int) int {
	// 滚动 hash
	const prime = 16777619
	check := func(i int) bool {
		memo := make(map[uint32]struct{}, len(nums1)-i+1)
		var hash, pow uint32 = 0, 1
		for _, v := range nums1[:i] {
			hash = hash*prime + uint32(v)
			pow = pow * prime
		}
		memo[hash] = struct{}{}
		for j, v := range nums1[i:] {
			hash = hash*prime + uint32(v) - uint32(nums1[j])*pow
			memo[hash] = struct{}{}
		}
		hash = 0
		for _, v := range nums2[:i] {
			hash = hash*prime + uint32(v)
		}
		if _, ok := memo[hash]; ok {
			return true
		}
		for j, v := range nums2[i:] {
			hash = hash*prime + uint32(v) - uint32(nums2[j])*pow
			if _, ok := memo[hash]; ok {
				return true
			}
		}
		return false
	}
	ans := 0
	sort.Search(min(len(nums1), len(nums2))+1, func(i int) bool {
		if i <= ans {
			return false
		}
		if check(i) {
			ans = i
			return false
		}
		return true
	})
	return ans

	// dp
	//ans, m, n := 0, len(nums1), len(nums2)
	//dp := make([][]int, m+1)
	//for i := range dp {
	//	dp[i] = make([]int, n+1)
	//}
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		if nums1[i] == nums2[j] {
	//			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j]+1)
	//			ans = max(ans, dp[i+1][j+1])
	//		}
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
