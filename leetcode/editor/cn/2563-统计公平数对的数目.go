package main

import (
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countFairPairs(nums []int, lower int, upper int) int64 {
	// 离散化树状数组
	// nums[j] 取值区间：[lower-nums[j], upper-nums[j]]

	// 三指针
	sort.Ints(nums)
	ans, l, r := 0, len(nums)-1, len(nums)
	for j, v := range nums { // 枚举 j
		for r > 0 && upper-v < nums[r-1] {
			r--
		}
		for l > 0 && lower-v <= nums[l-1] {
			l--
		}
		ans += min(r, j) - min(l, j) // 核心算法
	}
	return int64(ans)

	// 二分查找
	//sort.Ints(nums)
	//ans := 0
	//for j, v := range nums {
	//	ans += sort.SearchInts(nums[:j], upper-v+1) - sort.SearchInts(nums[:j], lower-v)
	//}
	//return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
