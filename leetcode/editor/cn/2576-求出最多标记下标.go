package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 5, 2, 4}
	nums = []int{6, 8, 7}
	nums = []int{4, 8, 7, 16, 2, 3}
	nums = []int{42, 83, 48, 10, 24, 55, 9, 100, 10, 17, 17, 99, 51, 32, 16, 98, 99, 31, 28, 68, 71, 14, 64, 29, 15, 40}

	indices := maxNumOfMarkedIndices(nums)
	fmt.Println(indices)
}

/*
反证法：
	1
		如果 2 * nums ≤ nums[j]，则称 nums[i] 和 nums[j] 匹配
		如果可以匹配 kkk 对，那么也可以匹配小于 kkk 对，去掉一些数对即可做到
		如果无法匹配 kkk 对，那么也无法匹配大于 kkk 对（反证法）
		因此答案有单调性，可以二分
	2
		检测能不能匹配 kkk 对
		要让哪些数匹配呢？
		结论：从小到大排序后，如果存在 kkk 对匹配，那么一定可以让最小的 kkk 个数和最大的 kkk 个数匹配

*/
// leetcode submit region begin(Prohibit modification and deletion)
func maxNumOfMarkedIndices(nums []int) int {
	// 双指针
	sort.Ints(nums)
	j, n := 0, len(nums)
	for i := (n + 1) >> 1; i < n; i++ {
		if nums[j]<<1 <= nums[i] {
			j++
		}
	}
	return j << 1

	// lc
	//sort.Ints(nums)
	//n := len(nums)
	//return 2 * sort.Search(n/2, func(k int) bool {
	//	k++
	//	for i, x := range nums[:k] {
	//		if x*2 > nums[n-k+i] {
	//			return true
	//		}
	//	}
	//	return false
	//})
}

//leetcode submit region end(Prohibit modification and deletion)
