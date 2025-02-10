package main

import (
	"slices"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	n := len(nums1)
	arr := slices.Clone(nums1)
	sort.Ints(arr)
	s, maxV := 0, 0
	for i, v := range nums2 {
		d := abs(nums1[i] - v)
		s += d
		j := sort.SearchInts(arr, v)
		//if j == n {
		//	maxV = max(maxV, d-v+arr[n-1])
		//} else if arr[j] == v {
		//	maxV = max(maxV, d)
		//} else if j == 0 {
		//	maxV = max(maxV, d-arr[0]+v)
		//} else {
		//	maxV = max(maxV, d-min(arr[j]-v, v-arr[j-1]))
		//}
		// 简化
		if j < n {
			maxV = max(maxV, d-arr[j]+v)
		}
		if j > 0 {
			maxV = max(maxV, d-v+arr[j-1])
		}
	}
	return (s - maxV) % 1_000_000_007
}

//leetcode submit region end(Prohibit modification and deletion)
