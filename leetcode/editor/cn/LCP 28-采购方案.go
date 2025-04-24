package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func purchasePlans(nums []int, target int) int {
	// 双指针
	sort.Ints(nums)
	ans, j := 0, len(nums)-1
	for i, v := range nums[:sort.SearchInts(nums, target+1)] {
		for j >= 0 && v+nums[j] > target {
			j--
		}
		ans += min(i, j+1)
	}
	return ans % 1_000_000_007
}

//leetcode submit region end(Prohibit modification and deletion)
